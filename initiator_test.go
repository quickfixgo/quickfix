package quickfix

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	"github.com/quickfixgo/quickfix/config"
)

func TestNewInitiatorKeepReconnectingAfterLogonError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logonCount := 0
	app := &mockApplication{}
	storeFactory := &mockMessageStoreFactory{saveMessageAndIncrError: errDBError}
	logFactory := &mockLogFactory{
		onEvent: func(s string) {
			if s == "Sending logon request" {
				logonCount++
				if logonCount >= 5 {
					cancel()
				}
			}
		},
	}

	settings := NewSettings()
	sessionSettings := newSession()
	sessionID, err := settings.AddSession(sessionSettings)
	if err != nil {
		t.Fatalf("Expected no error adding session, got %v", err)
	}

	initiator, err := NewInitiator(app, storeFactory, settings, logFactory)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	s, ok := initiator.sessions[sessionID]
	if !ok {
		t.Fatal("Expected session to be created")
	}

	initiator.stopChan = make(chan interface{})
	go initiator.handleConnection(s, nil, &mockDialer{})

	select {
	case <-ctx.Done():
		initiator.Stop()
		return
	case <-time.After(10 * time.Second):
		t.Error("retry stopped after logon error")
		return
	}
}

func newSession() *SessionSettings {
	sessionSettings := NewSessionSettings()
	sessionSettings.Set(config.BeginString, "FIX.4.4")
	sessionSettings.Set(config.SenderCompID, "X")
	sessionSettings.Set(config.TargetCompID, "X")
	sessionSettings.Set(config.HeartBtInt, "30")
	sessionSettings.Set(config.SocketConnectHost, "localhost")
	sessionSettings.Set(config.SocketConnectPort, "9878")
	sessionSettings.Set(config.ReconnectInterval, "1")
	return sessionSettings
}

type mockApplication struct{}

func (m *mockApplication) OnCreate(_ SessionID)                {}
func (m *mockApplication) OnLogon(_ SessionID)                 {}
func (m *mockApplication) OnLogout(_ SessionID)                {}
func (m *mockApplication) ToAdmin(_ *Message, _ SessionID)     {}
func (m *mockApplication) ToApp(_ *Message, _ SessionID) error { return nil }
func (m *mockApplication) FromAdmin(_ *Message, _ SessionID) MessageRejectError {
	return nil
}
func (m *mockApplication) FromApp(_ *Message, _ SessionID) MessageRejectError {
	return nil
}

type mockMessageStoreFactory struct {
	saveMessageAndIncrError error
}

func (m *mockMessageStoreFactory) Create(_ SessionID) (MessageStore, error) {
	return &mockMessageStore{saveMessageAndIncrError: m.saveMessageAndIncrError}, nil
}

var errDBError = errors.New("db error")

type mockMessageStore struct {
	saveMessageAndIncrError error
}

func (m *mockMessageStore) NextSenderMsgSeqNum() int                           { return 1 }
func (m *mockMessageStore) NextTargetMsgSeqNum() int                           { return 1 }
func (m *mockMessageStore) IncrSenderMsgSeqNum() error                         { return nil }
func (m *mockMessageStore) IncrTargetMsgSeqNum() error                         { return nil }
func (m *mockMessageStore) SetNextSenderMsgSeqNum(_ int) error                 { return nil }
func (m *mockMessageStore) SetNextTargetMsgSeqNum(_ int) error                 { return nil }
func (m *mockMessageStore) CreationTime() time.Time                            { return time.Now() }
func (m *mockMessageStore) SaveMessage(_ int, _ []byte) error                  { return nil }
func (m *mockMessageStore) GetMessages(_, _ int) ([][]byte, error)             { return nil, nil }
func (m *mockMessageStore) Refresh() error                                     { return nil }
func (m *mockMessageStore) Reset() error                                       { return nil }
func (m *mockMessageStore) Close() error                                       { return nil }
func (m *mockMessageStore) IncrNextSenderMsgSeqNum() error                     { return nil }
func (m *mockMessageStore) IncrNextTargetMsgSeqNum() error                     { return nil }
func (m *mockMessageStore) IterateMessages(int, int, func([]byte) error) error { return nil }
func (m *mockMessageStore) SaveMessageAndIncrNextSenderMsgSeqNum(_ int, _ []byte) error {
	return m.saveMessageAndIncrError
}
func (m *mockMessageStore) SetCreationTime(time.Time) {}

type mockLogFactory struct {
	shouldFail bool
	onEvent    func(string)
}

func (m *mockLogFactory) Create() (Log, error) {
	if m.shouldFail {
		return nil, errors.New("log factory error")
	}
	return &mockLog{
		onEvent: m.onEvent,
	}, nil
}

func (m *mockLogFactory) CreateSessionLog(_ SessionID) (Log, error) {
	return &mockLog{
		onEvent: m.onEvent,
	}, nil
}

type mockDialer struct{}

type mockAddr struct {
	network string
	address string
}

func (m *mockAddr) Network() string { return m.network }
func (m *mockAddr) String() string  { return m.address }

type mockConn struct{}

func (m *mockConn) Read(_ []byte) (n int, err error)   { return 0, nil }
func (m *mockConn) Write(_ []byte) (n int, err error)  { return 0, nil }
func (m *mockConn) Close() error                       { return nil }
func (m *mockConn) LocalAddr() net.Addr                { return &mockAddr{network: "tcp", address: "127.0.0.1:8080"} }
func (m *mockConn) RemoteAddr() net.Addr               { return &mockAddr{network: "tcp", address: "127.0.0.1:9090"} }
func (m *mockConn) SetDeadline(_ time.Time) error      { return nil }
func (m *mockConn) SetReadDeadline(_ time.Time) error  { return nil }
func (m *mockConn) SetWriteDeadline(_ time.Time) error { return nil }

func (m *mockDialer) DialContext(_ context.Context, _, _ string) (net.Conn, error) {
	return &mockConn{}, nil
}

type mockLog struct {
	onEvent func(string)
}

func (m *mockLog) OnIncoming(_ []byte) {}
func (m *mockLog) OnOutgoing(_ []byte) {}
func (m *mockLog) OnEvent(s string) {
	if m.onEvent != nil {
		m.onEvent(s)
	}
}
func (m *mockLog) OnEventf(_ string, _ ...interface{}) {}
