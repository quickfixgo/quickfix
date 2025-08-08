// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package quickfix

import (
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/quickfixgo/quickfix/internal"
)

type QuickFIXSuite struct {
	suite.Suite
}

type KnowsFieldMap interface {
	Has(Tag) bool
	GetString(Tag) (string, MessageRejectError)
	GetInt(Tag) (int, MessageRejectError)
	GetUint64(Tag) (uint64, MessageRejectError)
	GetField(Tag, FieldValueReader) MessageRejectError
}

func (s *QuickFIXSuite) MessageType(msgType string, msg *Message) {
	s.FieldEquals(tagMsgType, msgType, msg.Header)
}

func (s *QuickFIXSuite) FieldEquals(tag Tag, expectedValue interface{}, fieldMap KnowsFieldMap) {
	s.Require().True(fieldMap.Has(tag), "Tag %v not set", tag)

	switch expected := expectedValue.(type) {
	case string:
		val, err := fieldMap.GetString(tag)
		s.Nil(err)
		s.Equal(expected, val)
	case int:
		val, err := fieldMap.GetInt(tag)
		s.Nil(err)
		s.Equal(expected, val)
	case uint64:
		val, err := fieldMap.GetUint64(tag)
		s.Nil(err)
		s.Equal(expected, val)
	case bool:
		var val FIXBoolean
		err := fieldMap.GetField(tag, &val)
		s.Nil(err)
		s.Equal(expected, val.Bool())
	default:
		s.FailNow("Field type not handled")
	}
}

func (s *QuickFIXSuite) MessageEqualsBytes(expectedBytes []byte, msg *Message) {
	actualBytes := msg.build()
	s.Equal(string(actualBytes), string(expectedBytes))
}

// MockStore wraps a memory store and mocks Refresh for convenience.
type MockStore struct {
	mock.Mock
	memoryStore
}

func (s *MockStore) Refresh() error {
	return s.Called().Error(0)
}

type MockApp struct {
	mock.Mock

	decorateToAdmin func(*Message)
	lastToAdmin     *Message
	lastToApp       *Message
}

func (e *MockApp) OnCreate(_ SessionID) {
}

func (e *MockApp) OnLogon(_ SessionID) {
	e.Called()
}

func (e *MockApp) OnLogout(_ SessionID) {
	e.Called()
}

func (e *MockApp) FromAdmin(_ *Message, _ SessionID) (reject MessageRejectError) {
	if err, ok := e.Called().Get(0).(MessageRejectError); ok {
		return err
	}

	return nil
}

func (e *MockApp) ToAdmin(msg *Message, _ SessionID) {
	e.Called()

	if e.decorateToAdmin != nil {
		e.decorateToAdmin(msg)
	}

	e.lastToAdmin = msg
}

func (e *MockApp) ToApp(msg *Message, _ SessionID) (err error) {
	e.lastToApp = msg
	return e.Called().Error(0)
}

func (e *MockApp) FromApp(_ *Message, _ SessionID) (reject MessageRejectError) {
	if err, ok := e.Called().Get(0).(MessageRejectError); ok {
		return err
	}

	return nil
}

type MessageFactory struct {
	seqNum uint64
}

func (m *MessageFactory) SetNextSeqNum(next uint64) {
	m.seqNum = next - 1
}

func (m *MessageFactory) buildMessage(msgType string) *Message {
	m.seqNum++
	msg := NewMessage()
	msg.Header.
		SetField(tagBeginString, FIXString(string(BeginStringFIX42))).
		SetField(tagSenderCompID, FIXString("TW")).
		SetField(tagTargetCompID, FIXString("ISLD")).
		SetField(tagSendingTime, FIXUTCTimestamp{Time: time.Now()}).
		SetField(tagMsgSeqNum, FIXUint64(m.seqNum)).
		SetField(tagMsgType, FIXString(msgType))
	return msg
}

func (m *MessageFactory) Logout() *Message {
	return m.buildMessage(string(msgTypeLogout))
}

func (m *MessageFactory) NewOrderSingle() *Message {
	return m.buildMessage("D")
}

func (m *MessageFactory) Heartbeat() *Message {
	return m.buildMessage(string(msgTypeHeartbeat))
}

func (m *MessageFactory) Logon() *Message {
	return m.buildMessage(string(msgTypeLogon))
}

func (m *MessageFactory) ResendRequest(beginSeqNo uint64) *Message {
	msg := m.buildMessage(string(msgTypeResendRequest))
	msg.Body.SetField(tagBeginSeqNo, FIXUint64(beginSeqNo))
	msg.Body.SetField(tagEndSeqNo, FIXUint64(0))

	return msg
}

func (m *MessageFactory) SequenceReset(seqNo uint64) *Message {
	msg := m.buildMessage(string(msgTypeSequenceReset))
	msg.Body.SetField(tagNewSeqNo, FIXUint64(seqNo))

	return msg
}

type MockSessionReceiver struct {
	sendChannel chan []byte
}

func newMockSessionReceiver() MockSessionReceiver {
	return MockSessionReceiver{
		sendChannel: make(chan []byte, 10),
	}
}

func (p *MockSessionReceiver) LastMessage() (msg []byte, ok bool) {
	select {
	case msg, ok = <-p.sendChannel:
	default:
		ok = true
	}

	return
}

type SessionSuiteRig struct {
	QuickFIXSuite
	MessageFactory
	MockApp   MockApp
	MockStore MockStore
	*session
	Receiver MockSessionReceiver
}

func (s *SessionSuiteRig) Init() {
	s.MockApp = MockApp{}
	s.MockStore = MockStore{}
	s.MessageFactory = MessageFactory{}
	s.Receiver = newMockSessionReceiver()
	s.session = &session{
		sessionID:    SessionID{BeginString: "FIX.4.2", TargetCompID: "TW", SenderCompID: "ISLD"},
		store:        &s.MockStore,
		application:  &s.MockApp,
		log:          nullLog{},
		messageOut:   s.Receiver.sendChannel,
		sessionEvent: make(chan internal.Event),
	}
	s.MaxLatency = 120 * time.Second
}

func (s *SessionSuiteRig) State(state sessionState) {
	s.IsType(state, s.session.State, "session state should be %v", state)
}

func (s *SessionSuiteRig) MessageSentEquals(msg *Message) {
	msgBytes, ok := s.Receiver.LastMessage()
	s.True(ok, "Should be connected")
	s.NotNil(msgBytes, "Message should have been sent")
	s.MessageEqualsBytes(msgBytes, msg)
}

func (s *SessionSuiteRig) LastToAppMessageSent() {
	s.MessageSentEquals(s.MockApp.lastToApp)
}

func (s *SessionSuiteRig) LastToAdminMessageSent() {
	require.NotNil(s.T(), s.MockApp.lastToAdmin, "No ToAdmin received")
	s.MessageSentEquals(s.MockApp.lastToAdmin)
}

func (s *SessionSuiteRig) NotStopped() {
	s.False(s.session.Stopped(), "session should not be stopped")
}

func (s *SessionSuiteRig) Stopped() {
	s.True(s.session.Stopped(), "session should be stopped")
}

func (s *SessionSuiteRig) Disconnected() {
	msg, ok := s.Receiver.LastMessage()
	s.Nil(msg, "Expect disconnect, not message")
	s.False(ok, "Expect disconnect")
}

func (s *SessionSuiteRig) NoMessageSent() {
	msg, _ := s.Receiver.LastMessage()
	s.Nil(msg, "no message should be sent but got %s", msg)
}

func (s *SessionSuiteRig) NoMessageQueued() {
	s.Empty(s.session.toSend, "no messages should be queueud")
}

func (s *SessionSuiteRig) ExpectStoreReset() {
	s.NextSenderMsgSeqNum(1)
	s.NextTargetMsgSeqNum(1)
}

func (s *SessionSuiteRig) NextTargetMsgSeqNum(expected uint64) {
	s.Equal(expected, s.session.store.NextTargetMsgSeqNum(), "NextTargetMsgSeqNum should be %v ", expected)
}

func (s *SessionSuiteRig) NextSenderMsgSeqNum(expected uint64) {
	s.Equal(expected, s.session.store.NextSenderMsgSeqNum(), "NextSenderMsgSeqNum should be %v", expected)
}

func (s *SessionSuiteRig) IncrNextSenderMsgSeqNum() {
	s.Require().Nil(s.session.store.IncrNextSenderMsgSeqNum())
}

func (s *SessionSuiteRig) IncrNextTargetMsgSeqNum() {
	s.Require().Nil(s.session.store.IncrNextTargetMsgSeqNum())
}

func (s *SessionSuiteRig) NoMessagePersisted(seqNum uint64) {
	persistedMessages, err := s.session.store.GetMessages(seqNum, seqNum)
	s.Nil(err)
	s.Empty(persistedMessages, "The message should not be persisted")
}

func (s *SessionSuiteRig) MessagePersisted(msg *Message) {
	var err error
	seqNum, err := msg.Header.GetUint64(tagMsgSeqNum)
	s.Nil(err, "message should have seq num")

	persistedMessages, err := s.session.store.GetMessages(seqNum, seqNum)
	s.Nil(err)
	s.Len(persistedMessages, 1, "a message should be stored at %v", seqNum)
	s.MessageEqualsBytes(persistedMessages[0], msg)
}
