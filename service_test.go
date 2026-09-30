package quickfix_test

import (
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/phayes/freeport"
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	report "github.com/quickfixgo/fix44/tradecapturereport"
	ack "github.com/quickfixgo/fix44/tradecapturereportack"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

type Client struct {
	*quickfix.Initiator
	quickfix.EmptyApplication
	isLoggedIn atomic.Int32
	t          *testing.T
	*quickfix.MessageRouter
	reports   atomic.Int32
	sessionID quickfix.SessionID
}

func (c *Client) OnLogon(_ quickfix.SessionID) {
	c.isLoggedIn.Add(1)
}
func (c *Client) OnLogoff(_ quickfix.SessionID) {
	c.isLoggedIn.Add(-1)
}

func (c *Client) FromApp(msg *quickfix.Message, id quickfix.SessionID) quickfix.MessageRejectError {
	err := c.Route(msg, id)
	require.NoError(c.t, err)
	return nil
}

var _ quickfix.Application = &Client{}

type Server struct {
	*quickfix.Acceptor
	quickfix.EmptyApplication
	t          *testing.T
	isLoggedIn atomic.Int32
	*quickfix.MessageRouter
	acks      atomic.Int32
	sessionID quickfix.SessionID
}

func (s *Server) OnLogon(_ quickfix.SessionID) {
	s.isLoggedIn.Add(1)
}
func (s *Server) OnLogoff(_ quickfix.SessionID) {
	s.isLoggedIn.Add(-1)
}
func (s *Server) FromApp(msg *quickfix.Message, id quickfix.SessionID) quickfix.MessageRejectError {
	err := s.Route(msg, id)
	require.NoError(s.t, err)
	return nil
}

var _ quickfix.Application = &Server{}

func TestBackToBackTest(t *testing.T) {
	port, err := freeport.GetFreePort()
	require.NoError(t, err)
	//This is to catch any default registry that might be used by getting a segv
	quickfix.SetDefaultRegistry(nil)
	s := NewServer(t, port).Start()
	c := NewClient(t, port).Start()
	t.Cleanup(c.Stop)
	t.Cleanup(s.Stop)

	require.Eventually(t, func() bool { return c.isLoggedIn.Load() > 0 }, 10*time.Second, 100*time.Millisecond)

	m := report.New(
		field.NewTradeReportID(strconv.Itoa(1)),
		field.NewPreviouslyReported(false),
		field.NewLastQty(decimal.New(1, 1), 2),
		field.NewLastPx(decimal.New(1, 1), 2),
		field.NewTradeDate(time.Now().Format("20060102")),
		field.NewTransactTime(time.Now()),
	)
	m.SetString(tag.Account, "account")
	m.SetExecID("execid")
	m.SetSymbol("symbol")
	m.SetSendingTime(time.Now())
	m.SetOrderQty(decimal.New(1, 1), 2)

	err = s.SendToTarget(m, s.sessionID)
	require.NoError(t, err)

	require.Eventually(t, func() bool { return s.acks.Load() > 0 }, 10*time.Second, 100*time.Millisecond)
	require.Equal(t, int32(1), s.acks.Load())
	require.Equal(t, int32(1), c.isLoggedIn.Load())
	require.Equal(t, int32(1), s.isLoggedIn.Load())
	require.Equal(t, int32(1), c.reports.Load())

}

func NewServer(t *testing.T, port int) *Server {
	var err error
	s := &Server{t: t}
	settings := quickfix.NewSettings()
	settings.GlobalSettings().Set("HeartBtInt", "1")
	session := quickfix.NewSessionSettings()
	session.Set("BeginString", "FIX.4.4")
	session.Set("SocketAcceptPort", strconv.Itoa(port))
	session.Set("TargetCompID", "sender")
	session.Set("SenderCompID", "target")
	s.sessionID, err = settings.AddSession(session)
	require.NoError(t, err)
	logger := quickfix.NewSlogLogger()
	logger.Name = "server"
	registry := quickfix.NewRegistry()

	s.Acceptor, err = quickfix.NewAcceptor(s, quickfix.NewMemoryStoreFactory(), settings, logger, quickfix.WithAcceptorRegistry(registry))
	s.MessageRouter = quickfix.NewMessageRouter(quickfix.WithMessageRouterRegistry(registry))
	require.NoError(t, err)
	s.AddRoute(ack.Route(s.OnAck))
	return s
}
func (s *Server) Start() *Server {
	s.t.Logf("Starting server")
	require.NoError(s.t, s.Acceptor.Start())
	return s
}

func (s *Server) OnAck(_ ack.TradeCaptureReportAck, _ quickfix.SessionID) quickfix.MessageRejectError {
	s.acks.Add(1)
	return nil
}

func NewClient(t *testing.T, port int) *Client {
	var err error
	c := &Client{t: t}
	settings := quickfix.NewSettings()
	session := quickfix.NewSessionSettings()
	session.Set("BeginString", "FIX.4.4")
	session.Set("SocketConnectPort", strconv.Itoa(port))
	session.Set("SocketConnectHost", "localhost")
	session.Set("TargetCompID", "target")
	session.Set("SenderCompID", "sender")
	session.Set("HeartBtInt", "1")
	c.sessionID, err = settings.AddSession(session)
	require.NoError(t, err)

	logger := quickfix.NewSlogLogger()
	logger.Name = "client"
	registry := quickfix.NewRegistry()
	c.Initiator, err = quickfix.NewInitiator(c, quickfix.NewMemoryStoreFactory(), settings, logger, quickfix.WithInitiatorRegistry(registry))
	require.NoError(t, err)
	c.MessageRouter = quickfix.NewMessageRouter(quickfix.WithMessageRouterRegistry(registry))
	c.AddRoute(report.Route(c.OnReport))

	return c
}
func (c *Client) Start() *Client {
	c.t.Logf("Starting client")
	require.NoError(c.t, c.Initiator.Start())
	return c
}

func (c *Client) OnReport(_ report.TradeCaptureReport, id quickfix.SessionID) quickfix.MessageRejectError {
	err := c.SendToTarget(ack.New(field.NewTradeReportID("1"), field.NewExecType(enum.ExecType_NEW)), id)
	require.NoError(c.t, err)
	c.reports.Add(1)
	return nil
}
