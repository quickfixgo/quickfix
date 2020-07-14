package quickfix

import (
	"bytes"
	"testing"
	"time"

	"github.com/quickfixgo/quickfix/internal"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func newFIXString(val string) *FIXString {
	s := FIXString(val)
	return &s
}

type SessionSuite struct {
	SessionSuiteRig
}

func TestSessionSuite(t *testing.T) {
	suite.Run(t, new(SessionSuite))
}

func (s *SessionSuite) SetupTest() {
	s.Init()
	s.Require().Nil(s.session.store.Reset())
	s.session.State = latentState{}
}

func (s *SessionSuite) TestFillDefaultHeader() {
	s.session.sessionID.BeginString = "FIX.4.2"
	s.session.sessionID.TargetCompID = "TAR"
	s.session.sessionID.SenderCompID = "SND"

	msg := NewMessage()
	s.session.fillDefaultHeader(msg, nil)
	s.FieldEquals(tagBeginString, "FIX.4.2", msg.Header)
	s.FieldEquals(tagTargetCompID, "TAR", msg.Header)
	s.FieldEquals(tagSenderCompID, "SND", msg.Header)
	s.False(msg.Header.Has(tagSenderSubID))
	s.False(msg.Header.Has(tagSenderLocationID))
	s.False(msg.Header.Has(tagTargetSubID))
	s.False(msg.Header.Has(tagTargetLocationID))

	s.session.sessionID.BeginString = "FIX.4.3"
	s.session.sessionID.TargetCompID = "TAR"
	s.session.sessionID.TargetSubID = "TARS"
	s.session.sessionID.TargetLocationID = "TARL"
	s.session.sessionID.SenderCompID = "SND"
	s.session.sessionID.SenderSubID = "SNDS"
	s.session.sessionID.SenderLocationID = "SNDL"

	msg = NewMessage()
	s.session.fillDefaultHeader(msg, nil)
	s.FieldEquals(tagBeginString, "FIX.4.3", msg.Header)
	s.FieldEquals(tagTargetCompID, "TAR", msg.Header)
	s.FieldEquals(tagTargetSubID, "TARS", msg.Header)
	s.FieldEquals(tagTargetLocationID, "TARL", msg.Header)
	s.FieldEquals(tagSenderCompID, "SND", msg.Header)
	s.FieldEquals(tagSenderSubID, "SNDS", msg.Header)
	s.FieldEquals(tagSenderLocationID, "SNDL", msg.Header)
}

func (s *SessionSuite) TestInsertSendingTime() {
	var tests = []struct {
		BeginString       string
		Precision         TimestampPrecision
		ExpectedPrecision TimestampPrecision
	}{
		{BeginStringFIX40, Millis, Seconds}, //config is ignored for fix < 4.2
		{BeginStringFIX41, Millis, Seconds},

		{BeginStringFIX42, Millis, Millis},
		{BeginStringFIX42, Micros, Micros},
		{BeginStringFIX42, Nanos, Nanos},

		{BeginStringFIX43, Nanos, Nanos},
		{BeginStringFIX44, Nanos, Nanos},
		{BeginStringFIXT11, Nanos, Nanos},
	}

	for _, test := range tests {
		s.session.sessionID.BeginString = test.BeginString
		s.timestampPrecision = test.Precision

		msg := NewMessage()
		s.session.insertSendingTime(msg)

		var f FIXUTCTimestamp
		s.Nil(msg.Header.GetField(tagSendingTime, &f))
		s.Equal(f.Precision, test.ExpectedPrecision)
	}
}

func (s *SessionSuite) TestCheckCorrectCompID() {
	s.session.sessionID.TargetCompID = "TAR"
	s.session.sessionID.SenderCompID = "SND"

	var testCases = []struct {
		senderCompID *FIXString
		targetCompID *FIXString
		returnsError bool
		rejectReason int
	}{
		{returnsError: true, rejectReason: rejectReasonRequiredTagMissing},
		{senderCompID: newFIXString("TAR"),
			returnsError: true,
			rejectReason: rejectReasonRequiredTagMissing},
		{senderCompID: newFIXString("TAR"),
			targetCompID: newFIXString("JCD"),
			returnsError: true,
			rejectReason: rejectReasonCompIDProblem},
		{senderCompID: newFIXString("JCD"),
			targetCompID: newFIXString("SND"),
			returnsError: true,
			rejectReason: rejectReasonCompIDProblem},
		{senderCompID: newFIXString("TAR"),
			targetCompID: newFIXString("SND"),
			returnsError: false},
	}

	for _, tc := range testCases {
		msg := NewMessage()

		if tc.senderCompID != nil {
			msg.Header.SetField(tagSenderCompID, tc.senderCompID)
		}

		if tc.targetCompID != nil {
			msg.Header.SetField(tagTargetCompID, tc.targetCompID)
		}

		rej := s.session.checkCompID(msg)

		if !tc.returnsError {
			s.Require().Nil(rej)
			continue
		}

		s.NotNil(rej)
		s.Equal(tc.rejectReason, rej.RejectReason())
	}
}

func (s *SessionSuite) TestCheckBeginString() {
	msg := NewMessage()

	msg.Header.SetField(tagBeginString, FIXString("FIX.4.4"))
	err := s.session.checkBeginString(msg)
	s.Require().NotNil(err, "wrong begin string should return error")
	s.IsType(incorrectBeginString{}, err)

	msg.Header.SetField(tagBeginString, FIXString(s.session.sessionID.BeginString))
	s.Nil(s.session.checkBeginString(msg))
}

func (s *SessionSuite) TestCheckTargetTooHigh() {
	msg := NewMessage()
	s.Require().Nil(s.session.store.SetNextTargetMsgSeqNum(45))

	err := s.session.checkTargetTooHigh(msg)
	s.Require().NotNil(err, "missing sequence number should return error")
	s.Equal(rejectReasonRequiredTagMissing, err.RejectReason())

	msg.Header.SetField(tagMsgSeqNum, FIXInt(47))
	err = s.session.checkTargetTooHigh(msg)
	s.Require().NotNil(err, "sequence number too high should return an error")
	s.IsType(targetTooHigh{}, err)

	//spot on
	msg.Header.SetField(tagMsgSeqNum, FIXInt(45))
	s.Nil(s.session.checkTargetTooHigh(msg))
}

func (s *SessionSuite) TestCheckSendingTime() {
	s.session.MaxLatency = time.Duration(120) * time.Second
	msg := NewMessage()

	err := s.session.checkSendingTime(msg)
	s.Require().NotNil(err, "sending time is a required field")
	s.Equal(rejectReasonRequiredTagMissing, err.RejectReason())

	sendingTime := time.Now().Add(time.Duration(-200) * time.Second)
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime})

	err = s.session.checkSendingTime(msg)
	s.Require().NotNil(err, "sending time too late should give error")
	s.Equal(rejectReasonSendingTimeAccuracyProblem, err.RejectReason())

	sendingTime = time.Now().Add(time.Duration(200) * time.Second)
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime})

	err = s.session.checkSendingTime(msg)
	s.Require().NotNil(err, "future sending time should give error")
	s.Equal(rejectReasonSendingTimeAccuracyProblem, err.RejectReason())

	sendingTime = time.Now()
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime})

	s.Nil(s.session.checkSendingTime(msg), "sending time should be ok")

	s.session.SkipCheckLatency = true
	sendingTime = time.Now().Add(time.Duration(-200) * time.Second)
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime})
	err = s.session.checkSendingTime(msg)
	s.Require().Nil(err, "should skip latency check")
}

func (s *SessionSuite) TestCheckTargetTooLow() {
	msg := NewMessage()
	s.Require().Nil(s.session.store.SetNextTargetMsgSeqNum(45))

	err := s.session.checkTargetTooLow(msg)
	s.Require().NotNil(err, "sequence number is required")
	s.Equal(rejectReasonRequiredTagMissing, err.RejectReason())

	//too low
	msg.Header.SetField(tagMsgSeqNum, FIXInt(43))
	err = s.session.checkTargetTooLow(msg)
	s.NotNil(err, "sequence number too low should return error")
	s.IsType(targetTooLow{}, err)

	//spot on
	msg.Header.SetField(tagMsgSeqNum, FIXInt(45))
	s.Nil(s.session.checkTargetTooLow(msg))
}

func (s *SessionSuite) TestShouldSendReset() {
	var tests = []struct {
		BeginString         string
		ResetOnLogon        bool
		ResetOnDisconnect   bool
		ResetOnLogout       bool
		NextSenderMsgSeqNum int
		NextTargetMsgSeqNum int
		Expected            bool
	}{
		{BeginStringFIX40, true, false, false, 1, 1, false}, //ResetSeqNumFlag not available < fix41

		{BeginStringFIX41, true, false, false, 1, 1, true}, //session must be configured to reset on logon
		{BeginStringFIX42, true, false, false, 1, 1, true},
		{BeginStringFIX43, true, false, false, 1, 1, true},
		{BeginStringFIX44, true, false, false, 1, 1, true},
		{BeginStringFIXT11, true, false, false, 1, 1, true},

		{BeginStringFIX41, false, true, false, 1, 1, true}, //or disconnect
		{BeginStringFIX42, false, true, false, 1, 1, true},
		{BeginStringFIX43, false, true, false, 1, 1, true},
		{BeginStringFIX44, false, true, false, 1, 1, true},
		{BeginStringFIXT11, false, true, false, 1, 1, true},

		{BeginStringFIX41, false, false, true, 1, 1, true}, //or logout
		{BeginStringFIX42, false, false, true, 1, 1, true},
		{BeginStringFIX43, false, false, true, 1, 1, true},
		{BeginStringFIX44, false, false, true, 1, 1, true},
		{BeginStringFIXT11, false, false, true, 1, 1, true},

		{BeginStringFIX41, true, true, false, 1, 1, true}, //or combo
		{BeginStringFIX42, false, true, true, 1, 1, true},
		{BeginStringFIX43, true, false, true, 1, 1, true},
		{BeginStringFIX44, true, true, true, 1, 1, true},

		{BeginStringFIX41, false, false, false, 1, 1, false}, //or will not be set

		{BeginStringFIX41, true, false, false, 1, 10, false}, //session seq numbers should be reset at the time of check
		{BeginStringFIX42, true, false, false, 2, 1, false},
		{BeginStringFIX43, true, false, false, 14, 100, false},
	}

	for _, test := range tests {
		s.session.sessionID.BeginString = test.BeginString
		s.session.ResetOnLogon = test.ResetOnLogon
		s.session.ResetOnDisconnect = test.ResetOnDisconnect
		s.session.ResetOnLogout = test.ResetOnLogout

		s.Require().Nil(s.MockStore.SetNextSenderMsgSeqNum(test.NextSenderMsgSeqNum))
		s.Require().Nil(s.MockStore.SetNextTargetMsgSeqNum(test.NextTargetMsgSeqNum))

		s.Equal(s.shouldSendReset(), test.Expected)
	}
}

func (s *SessionSuite) TestCheckSessionTimeNoStartTimeEndTime() {
	var tests = []struct {
		before, after sessionState
	}{
		{before: latentState{}},
		{before: logonState{}},
		{before: logoutState{}},
		{before: inSession{}},
		{before: resendState{}},
		{before: pendingTimeout{resendState{}}},
		{before: pendingTimeout{inSession{}}},
		{before: notSessionTime{}, after: latentState{}},
	}

	for _, test := range tests {
		s.SetupTest()
		s.session.SessionTime = nil
		s.session.State = test.before

		s.session.CheckSessionTime(s.session, time.Now())
		if test.after != nil {
			s.State(test.after)
		} else {
			s.State(test.before)
		}
	}
}

func (s *SessionSuite) TestCheckSessionTimeInRange() {
	var tests = []struct {
		before, after sessionState
		expectReset   bool
	}{
		{before: latentState{}},
		{before: logonState{}},
		{before: logoutState{}},
		{before: inSession{}},
		{before: resendState{}},
		{before: pendingTimeout{resendState{}}},
		{before: pendingTimeout{inSession{}}},
		{before: notSessionTime{}, after: latentState{}, expectReset: true},
	}

	for _, test := range tests {
		s.SetupTest()
		s.session.State = test.before

		now := time.Now().UTC()
		store := new(memoryStore)
		if test.before.IsSessionTime() {
			s.Require().Nil(store.Reset())
		} else {
			store.creationTime = now.Add(time.Duration(-1) * time.Minute)
		}
		s.session.store = store
		s.IncrNextSenderMsgSeqNum()
		s.IncrNextTargetMsgSeqNum()

		s.session.SessionTime = internal.NewUTCTimeRange(
			internal.NewTimeOfDay(now.Clock()),
			internal.NewTimeOfDay(now.Add(time.Hour).Clock()),
		)

		s.session.CheckSessionTime(s.session, now)
		if test.after != nil {
			s.State(test.after)
		} else {
			s.State(test.before)
		}

		if test.expectReset {
			s.ExpectStoreReset()
		} else {
			s.NextSenderMsgSeqNum(2)
			s.NextSenderMsgSeqNum(2)
		}
	}
}

func (s *SessionSuite) TestCheckSessionTimeNotInRange() {
	var tests = []struct {
		before           sessionState
		initiateLogon    bool
		expectOnLogout   bool
		expectSendLogout bool
	}{
		{before: latentState{}},
		{before: logonState{}},
		{before: logonState{}, initiateLogon: true, expectOnLogout: true},
		{before: logoutState{}, expectOnLogout: true},
		{before: inSession{}, expectOnLogout: true, expectSendLogout: true},
		{before: resendState{}, expectOnLogout: true, expectSendLogout: true},
		{before: pendingTimeout{resendState{}}, expectOnLogout: true, expectSendLogout: true},
		{before: pendingTimeout{inSession{}}, expectOnLogout: true, expectSendLogout: true},
		{before: notSessionTime{}},
	}

	for _, test := range tests {
		s.SetupTest()
		s.session.State = test.before
		s.session.InitiateLogon = test.initiateLogon
		s.IncrNextSenderMsgSeqNum()
		s.IncrNextTargetMsgSeqNum()

		now := time.Now().UTC()
		s.session.SessionTime = internal.NewUTCTimeRange(
			internal.NewTimeOfDay(now.Add(time.Hour).Clock()),
			internal.NewTimeOfDay(now.Add(time.Duration(2)*time.Hour).Clock()),
		)

		if test.expectOnLogout {
			s.MockApp.On("OnLogout")
		}
		if test.expectSendLogout {
			s.MockApp.On("ToAdmin")
		}
		s.session.CheckSessionTime(s.session, now)

		s.MockApp.AssertExpectations(s.T())
		s.State(notSessionTime{})

		s.NextTargetMsgSeqNum(2)
		if test.expectSendLogout {
			s.LastToAdminMessageSent()
			s.MessageType(string(msgTypeLogout), s.MockApp.lastToAdmin)
			s.NextSenderMsgSeqNum(3)
		} else {
			s.NextSenderMsgSeqNum(2)
		}
	}
}

func (s *SessionSuite) TestCheckSessionTimeInRangeButNotSameRangeAsStore() {
	var tests = []struct {
		before           sessionState
		initiateLogon    bool
		expectOnLogout   bool
		expectSendLogout bool
	}{
		{before: latentState{}},
		{before: logonState{}},
		{before: logonState{}, initiateLogon: true, expectOnLogout: true},
		{before: logoutState{}, expectOnLogout: true},
		{before: inSession{}, expectOnLogout: true, expectSendLogout: true},
		{before: resendState{}, expectOnLogout: true, expectSendLogout: true},
		{before: pendingTimeout{resendState{}}, expectOnLogout: true, expectSendLogout: true},
		{before: pendingTimeout{inSession{}}, expectOnLogout: true, expectSendLogout: true},
		{before: notSessionTime{}},
	}

	for _, test := range tests {
		s.SetupTest()
		s.session.State = test.before
		s.session.InitiateLogon = test.initiateLogon
		s.Require().Nil(s.store.Reset())
		s.IncrNextSenderMsgSeqNum()
		s.IncrNextTargetMsgSeqNum()

		now := time.Now().UTC()
		s.session.SessionTime = internal.NewUTCTimeRange(
			internal.NewTimeOfDay(now.Add(time.Duration(-1)*time.Hour).Clock()),
			internal.NewTimeOfDay(now.Add(time.Hour).Clock()),
		)

		if test.expectOnLogout {
			s.MockApp.On("OnLogout")
		}
		if test.expectSendLogout {
			s.MockApp.On("ToAdmin")
		}
		s.session.CheckSessionTime(s.session, now.AddDate(0, 0, 1))

		s.MockApp.AssertExpectations(s.T())
		s.State(latentState{})
		if test.expectSendLogout {
			s.LastToAdminMessageSent()
			s.MessageType(string(msgTypeLogout), s.MockApp.lastToAdmin)
			s.FieldEquals(tagMsgSeqNum, 2, s.MockApp.lastToAdmin.Header)
		}
		s.ExpectStoreReset()
	}
}

func (s *SessionSuite) TestIncomingNotInSessionTime() {
	var tests = []struct {
		before           sessionState
		initiateLogon    bool
		expectOnLogout   bool
		expectSendLogout bool
	}{
		{before: logonState{}},
		{before: logonState{}, initiateLogon: true, expectOnLogout: true},
		{before: logoutState{}, expectOnLogout: true},
		{before: inSession{}, expectOnLogout: true, expectSendLogout: true},
		{before: resendState{}, expectOnLogout: true, expectSendLogout: true},
		{before: pendingTimeout{resendState{}}, expectOnLogout: true, expectSendLogout: true},
		{before: pendingTimeout{inSession{}}, expectOnLogout: true, expectSendLogout: true},
	}

	for _, test := range tests {
		s.SetupTest()

		s.session.State = test.before
		s.session.InitiateLogon = test.initiateLogon
		s.IncrNextSenderMsgSeqNum()
		s.IncrNextTargetMsgSeqNum()

		now := time.Now().UTC()
		s.session.SessionTime = internal.NewUTCTimeRange(
			internal.NewTimeOfDay(now.Add(time.Hour).Clock()),
			internal.NewTimeOfDay(now.Add(time.Duration(2)*time.Hour).Clock()),
		)
		if test.expectOnLogout {
			s.MockApp.On("OnLogout")
		}
		if test.expectSendLogout {
			s.MockApp.On("ToAdmin")
		}

		msg := s.NewOrderSingle()
		msgBytes := msg.build()

		s.session.Incoming(s.session, fixIn{bytes: bytes.NewBuffer(msgBytes)})
		s.MockApp.AssertExpectations(s.T())
		s.State(notSessionTime{})
	}
}

func (s *SessionSuite) TestSendAppMessagesNotInSessionTime() {
	var tests = []struct {
		before           sessionState
		initiateLogon    bool
		expectOnLogout   bool
		expectSendLogout bool
	}{
		{before: logonState{}},
		{before: logonState{}, initiateLogon: true, expectOnLogout: true},
		{before: logoutState{}, expectOnLogout: true},
		{before: inSession{}, expectOnLogout: true, expectSendLogout: true},
		{before: resendState{}, expectOnLogout: true, expectSendLogout: true},
		{before: pendingTimeout{resendState{}}, expectOnLogout: true, expectSendLogout: true},
		{before: pendingTimeout{inSession{}}, expectOnLogout: true, expectSendLogout: true},
	}

	for _, test := range tests {
		s.SetupTest()

		s.session.State = test.before
		s.session.InitiateLogon = test.initiateLogon
		s.IncrNextSenderMsgSeqNum()
		s.IncrNextTargetMsgSeqNum()

		s.MockApp.On("ToApp").Return(nil)
		s.Require().Nil(s.queueForSend(s.NewOrderSingle()))
		s.MockApp.AssertExpectations(s.T())

		now := time.Now().UTC()
		s.session.SessionTime = internal.NewUTCTimeRange(
			internal.NewTimeOfDay(now.Add(time.Hour).Clock()),
			internal.NewTimeOfDay(now.Add(time.Duration(2)*time.Hour).Clock()),
		)
		if test.expectOnLogout {
			s.MockApp.On("OnLogout")
		}
		if test.expectSendLogout {
			s.MockApp.On("ToAdmin")
		}

		s.session.SendAppMessages(s.session)
		s.MockApp.AssertExpectations(s.T())
		s.State(notSessionTime{})
	}
}

func (s *SessionSuite) TestTimeoutNotInSessionTime() {
	var tests = []struct {
		before           sessionState
		initiateLogon    bool
		expectOnLogout   bool
		expectSendLogout bool
	}{
		{before: logonState{}},
		{before: logonState{}, initiateLogon: true, expectOnLogout: true},
		{before: logoutState{}, expectOnLogout: true},
		{before: inSession{}, expectOnLogout: true, expectSendLogout: true},
		{before: resendState{}, expectOnLogout: true, expectSendLogout: true},
		{before: pendingTimeout{resendState{}}, expectOnLogout: true, expectSendLogout: true},
		{before: pendingTimeout{inSession{}}, expectOnLogout: true, expectSendLogout: true},
	}

	var events = []internal.Event{internal.PeerTimeout, internal.NeedHeartbeat, internal.LogonTimeout, internal.LogoutTimeout}

	for _, test := range tests {
		for _, event := range events {
			s.SetupTest()

			s.session.State = test.before
			s.session.InitiateLogon = test.initiateLogon
			s.IncrNextSenderMsgSeqNum()
			s.IncrNextTargetMsgSeqNum()

			now := time.Now().UTC()
			s.session.SessionTime = internal.NewUTCTimeRange(
				internal.NewTimeOfDay(now.Add(time.Hour).Clock()),
				internal.NewTimeOfDay(now.Add(time.Duration(2)*time.Hour).Clock()),
			)
			if test.expectOnLogout {
				s.MockApp.On("OnLogout")
			}
			if test.expectSendLogout {
				s.MockApp.On("ToAdmin")
			}

			s.session.Timeout(s.session, event)
			s.MockApp.AssertExpectations(s.T())
			s.State(notSessionTime{})
		}
	}
}

func (s *SessionSuite) TestOnAdminConnectInitiateLogon() {
	adminMsg := connect{
		messageOut: s.Receiver.sendChannel,
	}
	s.session.State = latentState{}
	s.session.HeartBtInt = time.Duration(45) * time.Second
	s.IncrNextSenderMsgSeqNum()
	s.session.InitiateLogon = true

	s.MockApp.On("ToAdmin")
	s.session.onAdmin(adminMsg)

	s.MockApp.AssertExpectations(s.T())
	s.True(s.session.InitiateLogon)
	s.False(s.sentReset)
	s.State(logonState{})
	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeLogon), s.MockApp.lastToAdmin)
	s.FieldEquals(tagHeartBtInt, 45, s.MockApp.lastToAdmin.Body)
	s.FieldEquals(tagMsgSeqNum, 2, s.MockApp.lastToAdmin.Header)
	s.NextSenderMsgSeqNum(3)
}

func (s *SessionSuite) TestInitiateLogonResetSeqNumFlag() {
	adminMsg := connect{
		messageOut: s.Receiver.sendChannel,
	}
	s.session.State = latentState{}
	s.session.HeartBtInt = time.Duration(45) * time.Second
	s.IncrNextTargetMsgSeqNum()
	s.IncrNextSenderMsgSeqNum()
	s.session.InitiateLogon = true

	s.MockApp.On("ToAdmin")
	s.MockApp.decorateToAdmin = func(msg *Message) {
		msg.Body.SetField(tagResetSeqNumFlag, FIXBoolean(true))
	}
	s.session.onAdmin(adminMsg)

	s.MockApp.AssertExpectations(s.T())
	s.True(s.session.InitiateLogon)
	s.True(s.sentReset)
	s.State(logonState{})
	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeLogon), s.MockApp.lastToAdmin)
	s.FieldEquals(tagMsgSeqNum, 1, s.MockApp.lastToAdmin.Header)
	s.FieldEquals(tagResetSeqNumFlag, true, s.MockApp.lastToAdmin.Body)
	s.NextSenderMsgSeqNum(2)
	s.NextTargetMsgSeqNum(1)
}

func (s *SessionSuite) TestOnAdminConnectInitiateLogonFIXT11() {
	s.session.sessionID.BeginString = string(BeginStringFIXT11)
	s.session.DefaultApplVerID = "8"
	s.session.InitiateLogon = true

	adminMsg := connect{
		messageOut: s.Receiver.sendChannel,
	}
	s.session.State = latentState{}

	s.MockApp.On("ToAdmin")
	s.session.onAdmin(adminMsg)

	s.MockApp.AssertExpectations(s.T())
	s.True(s.session.InitiateLogon)
	s.State(logonState{})
	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeLogon), s.MockApp.lastToAdmin)
	s.FieldEquals(tagDefaultApplVerID, "8", s.MockApp.lastToAdmin.Body)
}

func (s *SessionSuite) TestOnAdminConnectRefreshOnLogon() {
	var tests = []bool{true, false}

	for _, doRefresh := range tests {
		s.SetupTest()
		s.session.RefreshOnLogon = doRefresh

		adminMsg := connect{
			messageOut: s.Receiver.sendChannel,
		}
		s.session.State = latentState{}
		s.session.InitiateLogon = true

		if doRefresh {
			s.MockStore.On("Refresh").Return(nil)
		}
		s.MockApp.On("ToAdmin")
		s.session.onAdmin(adminMsg)

		s.MockStore.AssertExpectations(s.T())
	}
}

func (s *SessionSuite) TestOnAdminConnectAccept() {
	adminMsg := connect{
		messageOut: s.Receiver.sendChannel,
	}
	s.session.State = latentState{}
	s.IncrNextSenderMsgSeqNum()

	s.session.onAdmin(adminMsg)
	s.False(s.session.InitiateLogon)
	s.State(logonState{})
	s.NoMessageSent()
	s.NextSenderMsgSeqNum(2)
}

func (s *SessionSuite) TestOnAdminConnectNotInSession() {
	var tests = []bool{true, false}

	for _, doInitiateLogon := range tests {
		s.SetupTest()
		s.session.State = notSessionTime{}
		s.IncrNextSenderMsgSeqNum()
		s.session.InitiateLogon = doInitiateLogon

		adminMsg := connect{
			messageOut: s.Receiver.sendChannel,
		}

		s.session.onAdmin(adminMsg)

		s.State(notSessionTime{})
		s.NoMessageSent()
		s.Disconnected()
		s.NextSenderMsgSeqNum(2)
	}
}

func (s *SessionSuite) TestOnAdminStop() {
	s.session.State = logonState{}

	s.session.onAdmin(stopReq{})
	s.Disconnected()
	s.Stopped()
}

func (s *SessionSuite) TestResetOnDisconnect() {
	s.IncrNextSenderMsgSeqNum()
	s.IncrNextTargetMsgSeqNum()

	s.session.ResetOnDisconnect = false
	s.session.onDisconnect()
	s.NextSenderMsgSeqNum(2)
	s.NextTargetMsgSeqNum(2)

	s.session.ResetOnDisconnect = true
	s.session.onDisconnect()
	s.ExpectStoreReset()
}

type SessionSendTestSuite struct {
	SessionSuiteRig
}

func TestSessionSendTestSuite(t *testing.T) {
	suite.Run(t, new(SessionSendTestSuite))
}

func (suite *SessionSendTestSuite) SetupTest() {
	suite.Init()
	suite.session.State = inSession{}
}

func (suite *SessionSendTestSuite) TestQueueForSendAppMessage() {
	suite.MockApp.On("ToApp").Return(nil)
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))

	suite.MockApp.AssertExpectations(suite.T())
	suite.NoMessageSent()
	suite.MessagePersisted(suite.MockApp.lastToApp)
	suite.FieldEquals(tagMsgSeqNum, 1, suite.MockApp.lastToApp.Header)
	suite.NextSenderMsgSeqNum(2)
}

func (suite *SessionSendTestSuite) TestQueueForSendDoNotSendAppMessage() {
	suite.MockApp.On("ToApp").Return(ErrDoNotSend)
	suite.Equal(ErrDoNotSend, suite.queueForSend(suite.NewOrderSingle()))

	suite.MockApp.AssertExpectations(suite.T())
	suite.NoMessagePersisted(1)
	suite.NoMessageSent()
	suite.NextSenderMsgSeqNum(1)

	suite.MockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.send(suite.Heartbeat()))

	suite.MockApp.AssertExpectations(suite.T())
	suite.LastToAdminMessageSent()
	suite.MessagePersisted(suite.MockApp.lastToAdmin)
	suite.NextSenderMsgSeqNum(2)
}

func (suite *SessionSendTestSuite) TestQueueForSendAdminMessage() {
	suite.MockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))

	suite.MockApp.AssertExpectations(suite.T())
	suite.MessagePersisted(suite.MockApp.lastToAdmin)
	suite.NoMessageSent()
	suite.NextSenderMsgSeqNum(2)
}

func (suite *SessionSendTestSuite) TestSendAppMessage() {
	suite.MockApp.On("ToApp").Return(nil)
	require.Nil(suite.T(), suite.send(suite.NewOrderSingle()))

	suite.MockApp.AssertExpectations(suite.T())
	suite.MessagePersisted(suite.MockApp.lastToApp)
	suite.LastToAppMessageSent()
	suite.NextSenderMsgSeqNum(2)
}

func (suite *SessionSendTestSuite) TestSendAppDoNotSendMessage() {
	suite.MockApp.On("ToApp").Return(ErrDoNotSend)
	suite.Equal(ErrDoNotSend, suite.send(suite.NewOrderSingle()))

	suite.MockApp.AssertExpectations(suite.T())
	suite.NextSenderMsgSeqNum(1)
	suite.NoMessageSent()
}

func (suite *SessionSendTestSuite) TestSendAdminMessage() {
	suite.MockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.send(suite.Heartbeat()))
	suite.MockApp.AssertExpectations(suite.T())

	suite.LastToAdminMessageSent()
	suite.MessagePersisted(suite.MockApp.lastToAdmin)
}

func (suite *SessionSendTestSuite) TestSendFlushesQueue() {
	suite.MockApp.On("ToApp").Return(nil)
	suite.MockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))

	order1 := suite.MockApp.lastToApp
	heartbeat := suite.MockApp.lastToAdmin

	suite.MockApp.AssertExpectations(suite.T())
	suite.NoMessageSent()

	suite.MockApp.On("ToApp").Return(nil)
	require.Nil(suite.T(), suite.send(suite.NewOrderSingle()))
	suite.MockApp.AssertExpectations(suite.T())
	order2 := suite.MockApp.lastToApp
	suite.MessageSentEquals(order1)
	suite.MessageSentEquals(heartbeat)
	suite.MessageSentEquals(order2)
	suite.NoMessageSent()
}

func (suite *SessionSendTestSuite) TestSendNotLoggedOn() {
	suite.MockApp.On("ToApp").Return(nil)
	suite.MockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))

	suite.MockApp.AssertExpectations(suite.T())
	suite.NoMessageSent()

	var tests = []sessionState{logoutState{}, latentState{}, logonState{}}

	for _, test := range tests {
		suite.MockApp.On("ToApp").Return(nil)
		suite.session.State = test
		require.Nil(suite.T(), suite.send(suite.NewOrderSingle()))
		suite.MockApp.AssertExpectations(suite.T())
		suite.NoMessageSent()
	}
}

func (suite *SessionSendTestSuite) TestSendEnableLastMsgSeqNumProcessed() {
	suite.session.State = inSession{}
	suite.session.EnableLastMsgSeqNumProcessed = true

	suite.Require().Nil(suite.session.store.SetNextTargetMsgSeqNum(45))

	suite.MockApp.On("ToApp").Return(nil)
	require.Nil(suite.T(), suite.send(suite.NewOrderSingle()))
	suite.MockApp.AssertExpectations(suite.T())
	suite.LastToAppMessageSent()

	suite.FieldEquals(tagLastMsgSeqNumProcessed, 44, suite.MockApp.lastToApp.Header)
}

func (suite *SessionSendTestSuite) TestSendDisableMessagePersist() {
	suite.session.State = inSession{}
	suite.session.DisableMessagePersist = true

	suite.MockApp.On("ToApp").Return(nil)
	require.Nil(suite.T(), suite.send(suite.NewOrderSingle()))
	suite.MockApp.AssertExpectations(suite.T())
	suite.LastToAppMessageSent()
	suite.NoMessagePersisted(1)
	suite.NextSenderMsgSeqNum(2)
}

func (suite *SessionSendTestSuite) TestDropAndSendAdminMessage() {
	suite.MockApp.On("ToAdmin")
	suite.Require().Nil(suite.dropAndSend(suite.Heartbeat()))
	suite.MockApp.AssertExpectations(suite.T())

	suite.MessagePersisted(suite.MockApp.lastToAdmin)
	suite.LastToAdminMessageSent()
}

func (suite *SessionSendTestSuite) TestDropAndSendDropsQueue() {
	suite.MockApp.On("ToApp").Return(nil)
	suite.MockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))
	suite.MockApp.AssertExpectations(suite.T())

	suite.NoMessageSent()

	suite.MockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.dropAndSend(suite.Logon()))
	suite.MockApp.AssertExpectations(suite.T())

	msg := suite.MockApp.lastToAdmin
	suite.MessageType(string(msgTypeLogon), msg)
	suite.FieldEquals(tagMsgSeqNum, 3, msg.Header)

	//only one message sent
	suite.LastToAdminMessageSent()
	suite.NoMessageSent()
}

func (suite *SessionSendTestSuite) TestDropAndSendDropsQueueWithReset() {
	suite.MockApp.On("ToApp").Return(nil)
	suite.MockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))
	suite.MockApp.AssertExpectations(suite.T())
	suite.NoMessageSent()

	suite.MockApp.On("ToAdmin")
	suite.Require().Nil(suite.MockStore.Reset())
	require.Nil(suite.T(), suite.dropAndSend(suite.Logon()))
	suite.MockApp.AssertExpectations(suite.T())
	msg := suite.MockApp.lastToAdmin

	suite.MessageType(string(msgTypeLogon), msg)
	suite.FieldEquals(tagMsgSeqNum, 1, msg.Header)

	//only one message sent
	suite.LastToAdminMessageSent()
	suite.NoMessageSent()
}
