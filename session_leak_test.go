package quickfix

import (
	"bytes"
	"runtime/pprof"
	"strings"
	"testing"
	"time"

	"github.com/quickfixgo/quickfix/internal"
)

type testLog struct{}

func (testLog) OnIncoming([]byte)               {}
func (testLog) OnOutgoing([]byte)               {}
func (testLog) OnEvent(string)                  {}
func (testLog) OnEventf(string, ...interface{}) {}

// testApp is a no-op Application for tests.
type testApp struct{}

func (testApp) OnCreate(SessionID)                               {}
func (testApp) OnLogon(SessionID)                                {}
func (testApp) OnLogout(SessionID)                               {}
func (testApp) ToAdmin(*Message, SessionID)                      {}
func (testApp) FromAdmin(*Message, SessionID) MessageRejectError { return nil }
func (testApp) ToApp(*Message, SessionID) error                  { return nil }
func (testApp) FromApp(*Message, SessionID) MessageRejectError   { return nil }

func newTimerOnlySession() *session {
	tr, _ := internal.NewUTCTimeRange(internal.NewTimeOfDay(0, 0, 0), internal.NewTimeOfDay(23, 59, 59), nil)
	s := &session{
		store:        &memoryStore{},
		log:          testLog{},
		sessionID:    SessionID{BeginString: BeginStringFIX44, SenderCompID: "S", TargetCompID: "T"},
		messageOut:   make(chan []byte, 2),
		messageIn:    make(chan fixIn),
		sessionEvent: make(chan internal.Event),
		messageEvent: make(chan bool),
		application:  testApp{},
		stopCh:       make(chan struct{}),
		SessionSettings: internal.SessionSettings{
			SessionTime: tr,
		},
	}
	return s
}

func countGoroutinesContaining(substr string) int {
	var buf bytes.Buffer
	_ = pprof.Lookup("goroutine").WriteTo(&buf, 2)
	return strings.Count(buf.String(), substr)
}

func TestLogonTimeoutDoesNotLeakGoroutine(t *testing.T) {
	s := newTimerOnlySession()
	s.InitiateLogon = true
	s.LogonTimeout = 10 * time.Millisecond

	baseline := countGoroutinesContaining("stateMachine).Connect.func1")

	s.stateMachine.Connect(s)
	close(s.stopCh)
	time.Sleep(4 * s.LogonTimeout)

	if got := countGoroutinesContaining("stateMachine).Connect.func1"); got > baseline {
		t.Fatalf("logon timeout goroutine leaked: baseline=%d current=%d", baseline, got)
	}
}

func TestLogoutTimeoutDoesNotLeakGoroutine(t *testing.T) {
	s := newTimerOnlySession()
	s.LogoutTimeout = 10 * time.Millisecond
	baseline := countGoroutinesContaining("initiateLogoutInReplyTo")
	s.stateMachine.Connect(s)

	if err := s.initiateLogout("bye"); err != nil {
		t.Fatalf("initiateLogout returned error: %v", err)
	}
	close(s.stopCh)
	time.Sleep(4 * s.LogoutTimeout)

	if got := countGoroutinesContaining("initiateLogoutInReplyTo"); got > baseline {
		t.Fatalf("logout timeout goroutine leaked: baseline=%d current=%d", baseline, got)
	}
}
