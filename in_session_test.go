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
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/quickfixgo/quickfix/internal"
)

type InSessionTestSuite struct {
	SessionSuiteRig
}

func TestInSessionTestSuite(t *testing.T) {
	suite.Run(t, new(InSessionTestSuite))
}

func (s *InSessionTestSuite) SetupTest() {
	s.Init()
	s.session.State = inSession{}
}

func (s *InSessionTestSuite) TestPreliminary() {
	s.True(s.session.IsLoggedOn())
	s.True(s.session.IsConnected())
	s.True(s.session.IsSessionTime())
}

func (s *InSessionTestSuite) TestLogout() {
	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("ToAdmin")
	s.MockApp.On("OnLogout")
	s.session.fixMsgIn(s.session, s.Logout())

	s.MockApp.AssertExpectations(s.T())
	s.State(latentState{})

	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeLogout), s.MockApp.lastToAdmin)
	s.NextTargetMsgSeqNum(2)
	s.NextSenderMsgSeqNum(2)
}

func (s *InSessionTestSuite) TestLogoutEnableLastMsgSeqNumProcessed() {
	s.session.EnableLastMsgSeqNumProcessed = true

	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("ToAdmin")
	s.MockApp.On("OnLogout")
	s.session.fixMsgIn(s.session, s.Logout())

	s.MockApp.AssertExpectations(s.T())
	s.LastToAdminMessageSent()

	s.MessageType(string(msgTypeLogout), s.MockApp.lastToAdmin)
	s.FieldEquals(tagLastMsgSeqNumProcessed, 1, s.MockApp.lastToAdmin.Header)
}

func (s *InSessionTestSuite) TestLogoutResetOnLogout() {
	s.session.ResetOnLogout = true

	s.MockApp.On("ToApp").Return(nil)
	s.Nil(s.queueForSend(s.NewOrderSingle()))
	s.MockApp.AssertExpectations(s.T())

	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("ToAdmin")
	s.MockApp.On("OnLogout")
	s.session.fixMsgIn(s.session, s.Logout())

	s.MockApp.AssertExpectations(s.T())
	s.State(latentState{})
	s.LastToAppMessageSent()
	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeLogout), s.MockApp.lastToAdmin)

	s.NextTargetMsgSeqNum(1)
	s.NextSenderMsgSeqNum(1)
	s.NoMessageQueued()
}

func (s *InSessionTestSuite) TestLogoutTargetTooHigh() {
	s.MessageFactory.seqNum = 5

	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("ToAdmin")
	s.MockApp.On("OnLogout")
	s.session.fixMsgIn(s.session, s.Logout())

	s.MockApp.AssertExpectations(s.T())
	s.State(latentState{})

	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeLogout), s.MockApp.lastToAdmin)
	s.NextTargetMsgSeqNum(1)
	s.NextSenderMsgSeqNum(2)
}

func (s *InSessionTestSuite) TestTimeoutNeedHeartbeat() {
	s.MockApp.On("ToAdmin").Return(nil)
	s.session.Timeout(s.session, internal.NeedHeartbeat)

	s.MockApp.AssertExpectations(s.T())
	s.State(inSession{})
	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeHeartbeat), s.MockApp.lastToAdmin)
	s.NextSenderMsgSeqNum(2)
}

func (s *InSessionTestSuite) TestTimeoutPeerTimeout() {
	s.MockApp.On("ToAdmin").Return(nil)
	s.session.Timeout(s.session, internal.PeerTimeout)

	s.MockApp.AssertExpectations(s.T())
	s.State(pendingTimeout{inSession{}})
	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeTestRequest), s.MockApp.lastToAdmin)
	s.NextSenderMsgSeqNum(2)
}

func (s *InSessionTestSuite) TestDisconnected() {
	s.MockApp.On("OnLogout").Return(nil)
	s.session.Disconnected(s.session)
	s.MockApp.AssertExpectations(s.T())
	s.State(latentState{})
}

func (s *InSessionTestSuite) TestStop() {
	s.MockApp.On("ToAdmin")
	s.session.Stop(s.session)

	s.MockApp.AssertExpectations(s.T())
	s.State(logoutState{})
	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeLogout), s.MockApp.lastToAdmin)

	s.MockApp.On("OnLogout")
	s.session.Timeout(s.session, <-s.sessionEvent)
	s.MockApp.AssertExpectations(s.T())
	s.Stopped()
	s.Disconnected()
}

func (s *InSessionTestSuite) TestFIXMsgInTargetTooHighEnableLastMsgSeqNumProcessed() {
	s.session.EnableLastMsgSeqNumProcessed = true
	s.MessageFactory.seqNum = 5

	s.MockApp.On("ToAdmin")
	msgSeqNumTooHigh := s.NewOrderSingle()
	s.fixMsgIn(s.session, msgSeqNumTooHigh)

	s.MockApp.AssertExpectations(s.T())
	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeResendRequest), s.MockApp.lastToAdmin)
	s.FieldEquals(tagLastMsgSeqNumProcessed, 0, s.MockApp.lastToAdmin.Header)
}

func (s *InSessionTestSuite) TestFIXMsgInTargetTooHigh() {
	s.MessageFactory.seqNum = 5

	s.MockApp.On("ToAdmin")
	msgSeqNumTooHigh := s.NewOrderSingle()
	s.fixMsgIn(s.session, msgSeqNumTooHigh)

	s.MockApp.AssertExpectations(s.T())
	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeResendRequest), s.MockApp.lastToAdmin)
	s.FieldEquals(tagBeginSeqNo, 1, s.MockApp.lastToAdmin.Body)
	s.FieldEquals(tagEndSeqNo, 0, s.MockApp.lastToAdmin.Body)

	resendState, ok := s.session.State.(resendState)
	s.True(ok)
	s.NextTargetMsgSeqNum(1)

	stashedMsg, ok := resendState.messageStash[6]
	s.True(ok)

	rawMsg := msgSeqNumTooHigh.build()
	stashedRawMsg := stashedMsg.build()
	s.Equal(string(rawMsg), string(stashedRawMsg))
}
func (s *InSessionTestSuite) TestFIXMsgInTargetTooHighResendRequestChunkSize() {
	var tests = []struct {
		chunkSize        int
		expectedEndSeqNo int
	}{
		{0, 0},
		{10, 0},
		{5, 0},
		{2, 2},
		{3, 3},
	}

	for _, test := range tests {
		s.SetupTest()
		s.MessageFactory.seqNum = 5
		s.session.ResendRequestChunkSize = test.chunkSize

		s.MockApp.On("ToAdmin")
		msgSeqNumTooHigh := s.NewOrderSingle()
		s.fixMsgIn(s.session, msgSeqNumTooHigh)

		s.MockApp.AssertExpectations(s.T())
		s.LastToAdminMessageSent()
		s.MessageType(string(msgTypeResendRequest), s.MockApp.lastToAdmin)
		s.FieldEquals(tagBeginSeqNo, 1, s.MockApp.lastToAdmin.Body)
		s.FieldEquals(tagEndSeqNo, test.expectedEndSeqNo, s.MockApp.lastToAdmin.Body)

		resendState, ok := s.session.State.(resendState)
		s.True(ok)
		s.NextTargetMsgSeqNum(1)

		stashedMsg, ok := resendState.messageStash[6]
		s.True(ok)

		rawMsg := msgSeqNumTooHigh.build()
		stashedRawMsg := stashedMsg.build()
		s.Equal(string(rawMsg), string(stashedRawMsg))
	}
}

func (s *InSessionTestSuite) TestFIXMsgInResendRequestAllAdminExpectGapFill() {
	s.MockApp.On("ToAdmin")
	s.session.Timeout(s.session, internal.NeedHeartbeat)
	s.LastToAdminMessageSent()

	s.session.Timeout(s.session, internal.NeedHeartbeat)
	s.LastToAdminMessageSent()

	s.session.Timeout(s.session, internal.NeedHeartbeat)
	s.LastToAdminMessageSent()

	s.MockApp.AssertNumberOfCalls(s.T(), "ToAdmin", 3)
	s.NextSenderMsgSeqNum(4)

	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("ToAdmin")
	s.fixMsgIn(s.session, s.ResendRequest(1))

	s.MockApp.AssertExpectations(s.T())
	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeSequenceReset), s.MockApp.lastToAdmin)
	s.FieldEquals(tagMsgSeqNum, 1, s.MockApp.lastToAdmin.Header)
	s.FieldEquals(tagPossDupFlag, true, s.MockApp.lastToAdmin.Header)
	s.FieldEquals(tagNewSeqNo, 4, s.MockApp.lastToAdmin.Body)
	s.FieldEquals(tagGapFillFlag, true, s.MockApp.lastToAdmin.Body)

	s.NextSenderMsgSeqNum(4)
	s.State(inSession{})
}

func (s *InSessionTestSuite) TestFIXMsgInResendRequestAllAdminThenApp() {
	s.MockApp.On("ToAdmin")
	s.session.Timeout(s.session, internal.NeedHeartbeat)
	s.LastToAdminMessageSent()

	s.session.Timeout(s.session, internal.NeedHeartbeat)
	s.LastToAdminMessageSent()

	s.MockApp.On("ToApp").Return(nil)
	s.Require().Nil(s.session.send(s.NewOrderSingle()))
	s.LastToAppMessageSent()

	s.MockApp.AssertNumberOfCalls(s.T(), "ToAdmin", 2)
	s.MockApp.AssertNumberOfCalls(s.T(), "ToApp", 1)
	s.NextSenderMsgSeqNum(4)

	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("ToAdmin")
	s.MockApp.On("ToApp").Return(nil)
	s.fixMsgIn(s.session, s.ResendRequest(1))

	s.MockApp.AssertNumberOfCalls(s.T(), "ToAdmin", 3)
	s.MockApp.AssertNumberOfCalls(s.T(), "ToApp", 2)

	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeSequenceReset), s.MockApp.lastToAdmin)
	s.FieldEquals(tagMsgSeqNum, 1, s.MockApp.lastToAdmin.Header)
	s.FieldEquals(tagPossDupFlag, true, s.MockApp.lastToAdmin.Header)
	s.FieldEquals(tagNewSeqNo, 3, s.MockApp.lastToAdmin.Body)
	s.FieldEquals(tagGapFillFlag, true, s.MockApp.lastToAdmin.Body)

	s.LastToAppMessageSent()
	s.MessageType("D", s.MockApp.lastToApp)
	s.FieldEquals(tagMsgSeqNum, 3, s.MockApp.lastToApp.Header)
	s.FieldEquals(tagPossDupFlag, true, s.MockApp.lastToApp.Header)

	s.NextSenderMsgSeqNum(4)
	s.State(inSession{})
}

func (s *InSessionTestSuite) TestFIXMsgInResendRequestNoMessagePersist() {
	s.session.DisableMessagePersist = true

	s.MockApp.On("ToApp").Return(nil)
	s.Require().Nil(s.session.send(s.NewOrderSingle()))
	s.LastToAppMessageSent()

	s.MockApp.AssertNumberOfCalls(s.T(), "ToApp", 1)
	s.NextSenderMsgSeqNum(2)

	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("ToAdmin")
	s.fixMsgIn(s.session, s.ResendRequest(1))

	s.MockApp.AssertNumberOfCalls(s.T(), "ToAdmin", 1)
	s.MockApp.AssertNumberOfCalls(s.T(), "ToApp", 1)

	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeSequenceReset), s.MockApp.lastToAdmin)
	s.FieldEquals(tagMsgSeqNum, 1, s.MockApp.lastToAdmin.Header)
	s.FieldEquals(tagPossDupFlag, true, s.MockApp.lastToAdmin.Header)
	s.FieldEquals(tagNewSeqNo, 2, s.MockApp.lastToAdmin.Body)
	s.FieldEquals(tagGapFillFlag, true, s.MockApp.lastToAdmin.Body)

	s.NextSenderMsgSeqNum(2)
	s.State(inSession{})
}

func (s *InSessionTestSuite) TestFIXMsgInResendRequestDoNotSendApp() {
	s.MockApp.On("ToAdmin")
	s.session.Timeout(s.session, internal.NeedHeartbeat)
	s.LastToAdminMessageSent()

	s.MockApp.On("ToApp").Return(nil)
	s.Require().Nil(s.session.send(s.NewOrderSingle()))
	s.LastToAppMessageSent()

	s.session.Timeout(s.session, internal.NeedHeartbeat)
	s.LastToAdminMessageSent()

	s.MockApp.AssertNumberOfCalls(s.T(), "ToAdmin", 2)
	s.MockApp.AssertNumberOfCalls(s.T(), "ToApp", 1)
	s.NextSenderMsgSeqNum(4)

	// NOTE: a cheat here, need to reset mock.
	s.MockApp = MockApp{}
	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("ToApp").Return(ErrDoNotSend)
	s.MockApp.On("ToAdmin")
	s.fixMsgIn(s.session, s.ResendRequest(1))

	s.MockApp.AssertNumberOfCalls(s.T(), "ToAdmin", 1)
	s.MockApp.AssertNumberOfCalls(s.T(), "ToApp", 1)

	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeSequenceReset), s.MockApp.lastToAdmin)
	s.FieldEquals(tagMsgSeqNum, 1, s.MockApp.lastToAdmin.Header)
	s.FieldEquals(tagPossDupFlag, true, s.MockApp.lastToAdmin.Header)
	s.FieldEquals(tagNewSeqNo, 4, s.MockApp.lastToAdmin.Body)
	s.FieldEquals(tagGapFillFlag, true, s.MockApp.lastToAdmin.Body)

	s.NoMessageSent()

	s.NextSenderMsgSeqNum(4)
	s.State(inSession{})
}

// resendInterceptStore delegates to a real store but hooks IterateMessages so a
// test can simulate an application goroutine sending live traffic mid-replay.
type resendInterceptStore struct {
	MessageStore
	afterFirstMsg func()
}

func (s *resendInterceptStore) IterateMessages(beginSeqNum, endSeqNum int, cb func([]byte) error) error {
	n := 0
	return s.MessageStore.IterateMessages(beginSeqNum, endSeqNum, func(msgBytes []byte) error {
		err := cb(msgBytes)
		if n++; n == 1 && s.afterFirstMsg != nil {
			s.afterFirstMsg()
		}
		return err
	})
}

func (s *InSessionTestSuite) TestFIXMsgInResendRequestLiveMessagesDoNotInterleave() {
	s.MockApp.On("ToApp").Return(nil)
	for i := 0; i < 3; i++ {
		s.Require().Nil(s.session.send(s.NewOrderSingle()))
	}
	s.NextSenderMsgSeqNum(4)
	for {
		if msgBytes, _ := s.Receiver.LastMessage(); msgBytes == nil {
			break
		}
	}

	// While the replay is in progress, an application goroutine sends a live
	// message the same way SendToTarget does. The send must block until the
	// replay completes so the message cannot reach the wire mid-replay.
	liveSent := make(chan error, 1)
	s.session.store = &resendInterceptStore{
		MessageStore: &s.MockStore,
		afterFirstMsg: func() {
			go func() {
				liveSent <- s.session.queueForSend(s.NewOrderSingle())
			}()
			select {
			case err := <-liveSent:
				s.Failf("", "live send completed during replay (err=%v)", err)
			case <-time.After(20 * time.Millisecond):
			}
		},
	}

	s.MockApp.On("FromAdmin").Return(nil)
	s.fixMsgIn(s.session, s.ResendRequest(1))

	select {
	case err := <-liveSent:
		s.Require().Nil(err)
	case <-time.After(5 * time.Second):
		s.FailNow("live send never completed after the replay finished")
	}

	for expectedSeqNum := 1; expectedSeqNum <= 3; expectedSeqNum++ {
		msgBytes, _ := s.Receiver.LastMessage()
		s.Require().NotNil(msgBytes)
		msg := NewMessage()
		s.Require().Nil(ParseMessage(msg, bytes.NewBuffer(msgBytes)))
		s.FieldEquals(tagMsgSeqNum, expectedSeqNum, msg.Header)
		s.FieldEquals(tagPossDupFlag, true, msg.Header)
	}
	msgBytes, _ := s.Receiver.LastMessage()
	s.Nil(msgBytes, "live message must not be sent during the replay")

	// The live message stays queued and goes out once the event loop resumes.
	s.session.SendAppMessages(s.session)
	msgBytes, _ = s.Receiver.LastMessage()
	s.Require().NotNil(msgBytes)
	msg := NewMessage()
	s.Require().Nil(ParseMessage(msg, bytes.NewBuffer(msgBytes)))
	s.FieldEquals(tagMsgSeqNum, 4, msg.Header)
	s.False(msg.Header.Has(tagPossDupFlag))
	s.NextSenderMsgSeqNum(5)
}

func (s *InSessionTestSuite) TestFIXMsgInTargetTooLow() {
	s.IncrNextTargetMsgSeqNum()

	s.MockApp.On("ToAdmin")
	s.fixMsgIn(s.session, s.NewOrderSingle())
	s.MockApp.AssertExpectations(s.T())
	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeLogout), s.MockApp.lastToAdmin)
	s.FieldEquals(tagText, "MsgSeqNum too low, expecting 2 but received 1", s.MockApp.lastToAdmin.Body)
	s.State(logoutState{})
}

func (s *InSessionTestSuite) TestFIXMsgInTargetTooLowPossDup() {
	s.IncrNextTargetMsgSeqNum()

	s.MockApp.On("ToAdmin")
	nos := s.NewOrderSingle()
	nos.Header.SetField(tagPossDupFlag, FIXBoolean(true))

	s.fixMsgIn(s.session, nos)
	s.MockApp.AssertExpectations(s.T())
	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeReject), s.MockApp.lastToAdmin)
	s.FieldEquals(tagText, "Required tag missing", s.MockApp.lastToAdmin.Body)
	s.FieldEquals(tagRefTagID, int(tagOrigSendingTime), s.MockApp.lastToAdmin.Body)
	s.State(inSession{})

	nos.Header.SetField(tagOrigSendingTime, FIXUTCTimestamp{Time: time.Now().Add(time.Duration(-1) * time.Minute)})
	nos.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: time.Now()})
	s.fixMsgIn(s.session, nos)
	s.MockApp.AssertExpectations(s.T())
	s.NoMessageSent()
	s.State(inSession{})
	s.NextTargetMsgSeqNum(2)
}
