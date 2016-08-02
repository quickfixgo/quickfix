package quickfix

import (
	"time"

	"github.com/quickfixgo/quickfix/enum"
	"github.com/stretchr/testify/mock"
)

type messageFactory struct {
	seqNum int
}

func (m *messageFactory) buildMessage(msgType string) Message {
	m.seqNum++
	msg := NewMessage()
	msg.Header.
		SetField(tagBeginString, FIXString(enum.BeginStringFIX42)).
		SetField(tagSenderCompID, FIXString("TW")).
		SetField(tagTargetCompID, FIXString("ISLD")).
		SetField(tagSendingTime, FIXUTCTimestamp{Time: time.Now()}).
		SetField(tagMsgSeqNum, FIXInt(m.seqNum)).
		SetField(tagMsgType, FIXString(msgType))
	return msg
}

func (m *messageFactory) Logout() Message {
	return m.buildMessage("5")
}

func (m *messageFactory) NewOrderSingle() Message {
	return m.buildMessage("D")
}

func (m *messageFactory) Heartbeat() Message {
	return m.buildMessage("0")
}

func (m *messageFactory) Logon() Message {
	return m.buildMessage("A")
}

type mockApp struct {
	mock.Mock

	lastToAdmin Message
}

func (e *mockApp) OnCreate(sessionID SessionID) {
}

func (e *mockApp) OnLogon(sessionID SessionID) {
}

func (e *mockApp) OnLogout(sessionID SessionID) {
}

func (e *mockApp) FromAdmin(msg Message, sessionID SessionID) (reject MessageRejectError) {
	if err, ok := e.Called().Get(0).(MessageRejectError); ok {
		return err
	}

	return nil
}

func (e *mockApp) ToAdmin(msg Message, sessionID SessionID) {
	e.Called()
	e.lastToAdmin = msg
}

func (e *mockApp) ToApp(msg Message, sessionID SessionID) (err error) {
	return e.Called().Error(0)
}

func (e *mockApp) FromApp(msg Message, sessionID SessionID) (reject MessageRejectError) {
	if err, ok := e.Called().Get(0).(MessageRejectError); ok {
		return err
	}

	return nil
}

type mockSessionReceiver struct {
	sendChannel chan []byte
}

func newMockSessionReceiver() *mockSessionReceiver {
	return &mockSessionReceiver{
		sendChannel: make(chan []byte, 10),
	}
}

func (p *mockSessionReceiver) LastMessage() (msg []byte) {
	select {
	case msg = <-p.sendChannel:
	default:
	}

	return
}
