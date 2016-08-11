package quickfix

import (
	"testing"
	"time"

	"github.com/quickfixgo/quickfix/config"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/internal"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func newFIXString(val string) *FIXString {
	s := FIXString(val)
	return &s
}

func TestSession_CheckCorrectCompID(t *testing.T) {
	session := session{}
	session.sessionID.TargetCompID = "TAR"
	session.sessionID.SenderCompID = "SND"

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

		rej := session.checkCompID(msg)

		if !tc.returnsError {
			require.Nil(t, rej)
			continue
		}

		require.NotNil(t, rej)
		assert.Equal(t, tc.rejectReason, rej.RejectReason())
	}
}

func TestSession_CheckBeginString(t *testing.T) {
	session := session{
		sessionID: SessionID{BeginString: "FIX.4.2"},
	}

	msg := NewMessage()

	//wrong value
	msg.Header.SetField(tagBeginString, FIXString("FIX.4.4"))
	err := session.checkBeginString(msg)
	require.NotNil(t, err, "wrong begin string should return error")
	assert.IsType(t, incorrectBeginString{}, err)

	msg.Header.SetField(tagBeginString, FIXString(session.sessionID.BeginString))
	err = session.checkBeginString(msg)
	assert.Nil(t, err)
}

func TestSession_CheckTargetTooHigh(t *testing.T) {
	store := new(memoryStore)
	session := session{store: store}

	msg := NewMessage()
	require.Nil(t, store.SetNextTargetMsgSeqNum(45))

	err := session.checkTargetTooHigh(msg)
	require.NotNil(t, err, "missing sequence number should return error")
	assert.Equal(t, rejectReasonRequiredTagMissing, err.RejectReason())

	msg.Header.SetField(tagMsgSeqNum, FIXInt(47))
	err = session.checkTargetTooHigh(msg)
	require.NotNil(t, err, "sequence number too high should return an error")
	assert.IsType(t, targetTooHigh{}, err)

	//spot on
	msg.Header.SetField(tagMsgSeqNum, FIXInt(45))
	err = session.checkTargetTooHigh(msg)
	assert.Nil(t, err)
}

func TestSession_CheckSendingTime(t *testing.T) {
	session := session{}
	msg := NewMessage()

	err := session.checkSendingTime(msg)
	require.NotNil(t, err, "sending time is a required field")
	assert.Equal(t, rejectReasonRequiredTagMissing, err.RejectReason())

	sendingTime := time.Now().Add(time.Duration(-200) * time.Second)
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime})

	err = session.checkSendingTime(msg)
	require.NotNil(t, err, "sending time too late should give error")
	assert.Equal(t, rejectReasonSendingTimeAccuracyProblem, err.RejectReason())

	sendingTime = time.Now().Add(time.Duration(200) * time.Second)
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime})

	err = session.checkSendingTime(msg)
	require.NotNil(t, err, "future sending time should give error")
	assert.Equal(t, rejectReasonSendingTimeAccuracyProblem, err.RejectReason())

	sendingTime = time.Now()
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime})

	err = session.checkSendingTime(msg)
	assert.Nil(t, err, "sending time should be ok")
}

func TestSession_CheckTargetTooLow(t *testing.T) {
	store, _ := NewMemoryStoreFactory().Create(SessionID{})
	session := session{store: store}

	msg := NewMessage()
	require.Nil(t, store.SetNextTargetMsgSeqNum(45))

	err := session.checkTargetTooLow(msg)
	require.NotNil(t, err, "sequence number is required")
	assert.Equal(t, rejectReasonRequiredTagMissing, err.RejectReason())

	//too low
	msg.Header.SetField(tagMsgSeqNum, FIXInt(43))
	err = session.checkTargetTooLow(msg)
	require.NotNil(t, err, "sequence number too low should return error")
	assert.IsType(t, targetTooLow{}, err)

	//spot on
	msg.Header.SetField(tagMsgSeqNum, FIXInt(45))
	err = session.checkTargetTooLow(msg)
	assert.Nil(t, err)
}

type NewSessionTestSuite struct {
	suite.Suite

	SessionID
	MessageStoreFactory
	*SessionSettings
	LogFactory
	App *mockApp
}

func TestNewSessionTestSuite(t *testing.T) {
	suite.Run(t, new(NewSessionTestSuite))
}

func (s *NewSessionTestSuite) SetupTest() {
	s.SessionID = SessionID{BeginString: "FIX.4.2", TargetCompID: "TW", SenderCompID: "ISLD"}
	s.MessageStoreFactory = NewMemoryStoreFactory()
	s.SessionSettings = NewSessionSettings()
	s.LogFactory = nullLogFactory{}
	s.App = new(mockApp)
}

func (s *NewSessionTestSuite) TestDefaults() {
	session, err := newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.Nil(err)
	s.NotNil(session)

	s.False(session.resetOnLogon)
	s.False(session.resetOnLogout)
	s.Nil(session.sessionTime, "By default, start and end time unset")
}

func (s *NewSessionTestSuite) TestResetOnLogon() {
	var tests = []struct {
		setting  string
		expected bool
	}{{"Y", true}, {"N", false}}

	for _, test := range tests {
		s.SetupTest()
		s.SessionSettings.Set(config.ResetOnLogon, test.setting)
		session, err := newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
		s.Nil(err)
		s.NotNil(session)

		s.Equal(test.expected, session.resetOnLogon)
	}
}

func (s *NewSessionTestSuite) TestResetOnLogout() {
	var tests = []struct {
		setting  string
		expected bool
	}{{"Y", true}, {"N", false}}

	for _, test := range tests {
		s.SetupTest()
		s.SessionSettings.Set(config.ResetOnLogout, test.setting)
		session, err := newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
		s.Nil(err)
		s.NotNil(session)

		s.Equal(test.expected, session.resetOnLogout)
	}
}

func (s *NewSessionTestSuite) TestStartAndEndTime() {
	s.SessionSettings.Set(config.StartTime, "12:00:00")
	s.SessionSettings.Set(config.EndTime, "14:00:00")
	session, err := newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.Nil(err)
	s.NotNil(session.sessionTime)

	s.Equal(
		*internal.NewUTCTimeRange(internal.NewTimeOfDay(12, 0, 0), internal.NewTimeOfDay(14, 0, 0)),
		*session.sessionTime,
	)
}

func (s *NewSessionTestSuite) TestStartAndEndTimeAndTimeZone() {
	s.SessionSettings.Set(config.StartTime, "12:00:00")
	s.SessionSettings.Set(config.EndTime, "14:00:00")
	s.SessionSettings.Set(config.TimeZone, "Local")

	session, err := newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.Nil(err)
	s.NotNil(session.sessionTime)

	s.Equal(
		*internal.NewTimeRangeInLocation(internal.NewTimeOfDay(12, 0, 0), internal.NewTimeOfDay(14, 0, 0), time.Local),
		*session.sessionTime,
	)
}

func (s *NewSessionTestSuite) TestMissingStartOrEndTime() {
	s.SessionSettings.Set(config.StartTime, "12:00:00")
	_, err := newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)

	s.SetupTest()
	s.SessionSettings.Set(config.EndTime, "14:00:00")
	_, err = newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)
}

func (s *NewSessionTestSuite) TestStartOrEndTimeParseError() {
	s.SessionSettings.Set(config.StartTime, "1200:00")
	s.SessionSettings.Set(config.EndTime, "14:00:00")

	_, err := newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)

	s.SessionSettings.Set(config.StartTime, "12:00:00")
	s.SessionSettings.Set(config.EndTime, "")

	_, err = newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)
}

func (s *NewSessionTestSuite) TestInvalidTimeZone() {
	s.SessionSettings.Set(config.StartTime, "12:00:00")
	s.SessionSettings.Set(config.EndTime, "14:00:00")
	s.SessionSettings.Set(config.TimeZone, "not valid")

	_, err := newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)
}

type CheckSessionTimeTestSuite struct {
	SessionSuite
}

func TestCheckSessionTimeTestSuite(t *testing.T) {
	suite.Run(t, new(CheckSessionTimeTestSuite))
}

func (s *CheckSessionTimeTestSuite) SetupTest() {
	s.Init()
	s.session.store.Reset()
	s.session.State = latentState{}
}

func (s *CheckSessionTimeTestSuite) TestNoStartTimeEndTime() {
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
		s.session.sessionTime = nil
		s.session.State = test.before

		s.session.CheckSessionTime(s.session, time.Now())
		if test.after != nil {
			s.State(test.after)
		} else {
			s.State(test.before)
		}
	}
}

func (s *CheckSessionTimeTestSuite) TestInRange() {
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
			store.Reset()
		} else {
			store.creationTime = now.Add(time.Duration(-1) * time.Minute)
		}
		s.session.store = store
		s.Nil(s.session.store.IncrNextSenderMsgSeqNum())
		s.Nil(s.session.store.IncrNextTargetMsgSeqNum())

		s.session.sessionTime = internal.NewUTCTimeRange(
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

func (s *CheckSessionTimeTestSuite) TestNotInRange() {
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
		s.session.initiateLogon = test.initiateLogon
		s.Nil(s.session.store.IncrNextSenderMsgSeqNum())
		s.Nil(s.session.store.IncrNextTargetMsgSeqNum())

		now := time.Now().UTC()
		s.session.sessionTime = internal.NewUTCTimeRange(
			internal.NewTimeOfDay(now.Add(time.Hour).Clock()),
			internal.NewTimeOfDay(now.Add(time.Duration(2)*time.Hour).Clock()),
		)

		if test.expectOnLogout {
			s.mockApp.On("OnLogout")
		}
		if test.expectSendLogout {
			s.mockApp.On("ToAdmin")
		}
		s.session.CheckSessionTime(s.session, now)

		s.mockApp.AssertExpectations(s.T())
		s.State(notSessionTime{})

		s.NextTargetMsgSeqNum(2)
		if test.expectSendLogout {
			s.LastToAdminMessageSent()
			s.MessageType(enum.MsgType_LOGOUT, s.mockApp.lastToAdmin)
			s.NextSenderMsgSeqNum(3)
		} else {
			s.NextSenderMsgSeqNum(2)
		}
	}
}

func (s *CheckSessionTimeTestSuite) TestInRangeButNotSameRangeAsStore() {
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
		s.session.initiateLogon = test.initiateLogon
		s.store.Reset()
		s.Nil(s.session.store.IncrNextSenderMsgSeqNum())
		s.Nil(s.session.store.IncrNextTargetMsgSeqNum())

		now := time.Now().UTC()
		s.session.sessionTime = internal.NewUTCTimeRange(
			internal.NewTimeOfDay(now.Add(time.Duration(-1)*time.Hour).Clock()),
			internal.NewTimeOfDay(now.Add(time.Hour).Clock()),
		)

		if test.expectOnLogout {
			s.mockApp.On("OnLogout")
		}
		if test.expectSendLogout {
			s.mockApp.On("ToAdmin")
		}
		s.session.CheckSessionTime(s.session, now.AddDate(0, 0, 1))

		s.mockApp.AssertExpectations(s.T())
		s.State(latentState{})
		if test.expectSendLogout {
			s.LastToAdminMessageSent()
			s.MessageType(enum.MsgType_LOGOUT, s.mockApp.lastToAdmin)
			s.FieldEquals(tagMsgSeqNum, 2, s.mockApp.lastToAdmin.Header)
		}
		s.ExpectStoreReset()
	}
}

type SessionAdminTestSuite struct {
	SessionSuite
}

func TestSessionAdminTestSuite(t *testing.T) {
	suite.Run(t, new(SessionAdminTestSuite))
}

func (s *SessionAdminTestSuite) SetupTest() {
	s.Init()
}

func (s *SessionAdminTestSuite) TestConnectInitiateLogon() {
	adminMsg := connect{
		messageOut:    s.Receiver.sendChannel,
		initiateLogon: true,
	}
	s.session.State = latentState{}
	s.session.heartBtInt = time.Duration(45) * time.Second
	s.session.store.IncrNextSenderMsgSeqNum()

	s.mockApp.On("ToAdmin")
	s.session.onAdmin(adminMsg)

	s.mockApp.AssertExpectations(s.T())
	s.True(s.session.initiateLogon)
	s.State(logonState{})
	s.LastToAdminMessageSent()
	s.MessageType(enum.MsgType_LOGON, s.mockApp.lastToAdmin)
	s.FieldEquals(tagHeartBtInt, 45, s.mockApp.lastToAdmin.Body)
	s.FieldEquals(tagMsgSeqNum, 2, s.mockApp.lastToAdmin.Header)
	s.NextSenderMsgSeqNum(3)
}

func (s *SessionAdminTestSuite) TestConnectAccept() {
	adminMsg := connect{
		messageOut: s.Receiver.sendChannel,
	}
	s.session.State = latentState{}
	s.session.store.IncrNextSenderMsgSeqNum()

	s.session.onAdmin(adminMsg)
	s.False(s.session.initiateLogon)
	s.State(logonState{})
	s.NoMessageSent()
	s.NextSenderMsgSeqNum(2)
}

func (s *SessionAdminTestSuite) TestConnectNotInSession() {
	var tests = []bool{true, false}

	for _, doInitiateLogon := range tests {
		s.SetupTest()
		s.session.State = notSessionTime{}
		s.session.store.IncrNextSenderMsgSeqNum()

		adminMsg := connect{
			messageOut:    s.Receiver.sendChannel,
			initiateLogon: doInitiateLogon,
		}

		s.session.onAdmin(adminMsg)

		s.State(notSessionTime{})
		s.NoMessageSent()
		s.Disconnected()
		s.NextSenderMsgSeqNum(2)
	}
}

func (s *SessionAdminTestSuite) TestDisconnect() {
	s.session.State = logonState{}

	req := make(disconnectReq)
	go s.session.onAdmin(req)
	_, ok := <-req
	s.False(ok)
}

type SessionSendTestSuite struct {
	SessionSuite
}

func TestSessionSendTestSuite(t *testing.T) {
	suite.Run(t, new(SessionSendTestSuite))
}

func (suite *SessionSendTestSuite) SetupTest() {
	suite.Init()
	suite.session.State = inSession{}
}

func (suite *SessionSendTestSuite) TestQueueForSendAppMessage() {
	suite.mockApp.On("ToApp").Return(nil)
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))

	suite.mockApp.AssertExpectations(suite.T())
	suite.NoMessageSent()
	suite.MessagePersisted(suite.mockApp.lastToApp)
	suite.FieldEquals(tagMsgSeqNum, 1, suite.mockApp.lastToApp.Header)
	suite.NextSenderMsgSeqNum(2)
}

func (suite *SessionSendTestSuite) TestQueueForSendDoNotSendAppMessage() {
	suite.mockApp.On("ToApp").Return(DoNotSend)
	suite.Equal(DoNotSend, suite.queueForSend(suite.NewOrderSingle()))

	suite.AssertExpectations(suite.T())
	suite.NoMessagePersisted(1)
	suite.NoMessageSent()
	suite.NextSenderMsgSeqNum(1)

	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.send(suite.Heartbeat()))

	suite.mockApp.AssertExpectations(suite.T())
	suite.LastToAdminMessageSent()
	suite.MessagePersisted(suite.mockApp.lastToAdmin)
	suite.NextSenderMsgSeqNum(2)
}

func (suite *SessionSendTestSuite) TestQueueForSendAdminMessage() {
	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))

	suite.mockApp.AssertExpectations(suite.T())
	suite.MessagePersisted(suite.mockApp.lastToAdmin)
	suite.NoMessageSent()
	suite.NextSenderMsgSeqNum(2)
}

func (suite *SessionSendTestSuite) TestSendAppMessage() {
	suite.mockApp.On("ToApp").Return(nil)
	require.Nil(suite.T(), suite.send(suite.NewOrderSingle()))

	suite.mockApp.AssertExpectations(suite.T())
	suite.MessagePersisted(suite.mockApp.lastToApp)
	suite.LastToAppMessageSent()
	suite.NextSenderMsgSeqNum(2)
}

func (suite *SessionSendTestSuite) TestSendAppDoNotSendMessage() {
	suite.mockApp.On("ToApp").Return(DoNotSend)
	suite.Equal(DoNotSend, suite.send(suite.NewOrderSingle()))

	suite.mockApp.AssertExpectations(suite.T())
	suite.NextSenderMsgSeqNum(1)
	suite.NoMessageSent()
}

func (suite *SessionSendTestSuite) TestSendAdminMessage() {
	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.send(suite.Heartbeat()))
	suite.mockApp.AssertExpectations(suite.T())

	suite.LastToAdminMessageSent()
	suite.MessagePersisted(suite.mockApp.lastToAdmin)
}

func (suite *SessionSendTestSuite) TestSendFlushesQueue() {
	suite.mockApp.On("ToApp").Return(nil)
	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))

	order1 := suite.mockApp.lastToApp
	heartbeat := suite.mockApp.lastToAdmin

	suite.mockApp.AssertExpectations(suite.T())
	suite.NoMessageSent()

	suite.mockApp.On("ToApp").Return(nil)
	require.Nil(suite.T(), suite.send(suite.NewOrderSingle()))
	suite.mockApp.AssertExpectations(suite.T())
	order2 := suite.mockApp.lastToApp
	suite.MessageSentEquals(order1)
	suite.MessageSentEquals(heartbeat)
	suite.MessageSentEquals(order2)
	suite.NoMessageSent()
}

func (suite *SessionSendTestSuite) TestSendNotLoggedOn() {
	suite.mockApp.On("ToApp").Return(nil)
	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))

	suite.mockApp.AssertExpectations(suite.T())
	suite.NoMessageSent()

	var tests = []sessionState{logoutState{}, latentState{}, logonState{}}

	for _, test := range tests {
		suite.mockApp.On("ToApp").Return(nil)
		suite.session.State = test
		require.Nil(suite.T(), suite.send(suite.NewOrderSingle()))
		suite.mockApp.AssertExpectations(suite.T())
		suite.NoMessageSent()
	}
}

func (suite *SessionSendTestSuite) TestDropAndSendAdminMessage() {
	suite.mockApp.On("ToAdmin")
	suite.Require().Nil(suite.dropAndSend(suite.Heartbeat(), false))
	suite.mockApp.AssertExpectations(suite.T())

	suite.MessagePersisted(suite.mockApp.lastToAdmin)
	suite.LastToAdminMessageSent()
}

func (suite *SessionSendTestSuite) TestDropAndSendDropsQueue() {
	suite.mockApp.On("ToApp").Return(nil)
	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))
	suite.mockApp.AssertExpectations(suite.T())

	suite.NoMessageSent()

	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.dropAndSend(suite.Logon(), false))
	suite.mockApp.AssertExpectations(suite.T())

	msg := suite.mockApp.lastToAdmin
	suite.MessageType(enum.MsgType_LOGON, msg)
	suite.FieldEquals(tagMsgSeqNum, 3, msg.Header)

	//only one message sent
	suite.LastToAdminMessageSent()
	suite.NoMessageSent()
}

func (suite *SessionSendTestSuite) TestDropAndSendDropsQueueWithReset() {
	suite.mockApp.On("ToApp").Return(nil)
	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))
	suite.mockApp.AssertExpectations(suite.T())
	suite.NoMessageSent()

	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.dropAndSend(suite.Logon(), true))
	suite.mockApp.AssertExpectations(suite.T())
	msg := suite.mockApp.lastToAdmin

	suite.MessageType(enum.MsgType_LOGON, msg)
	suite.FieldEquals(tagMsgSeqNum, 1, msg.Header)

	//only one message sent
	suite.LastToAdminMessageSent()
	suite.NoMessageSent()
}
