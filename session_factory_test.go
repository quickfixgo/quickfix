package quickfix

import (
	"testing"
	"time"

	"github.com/quickfixgo/quickfix/config"
	"github.com/quickfixgo/quickfix/internal"
	"github.com/stretchr/testify/suite"
)

type SessionFactorySuite struct {
	suite.Suite

	sessionFactory
	SessionID
	MessageStoreFactory
	*SessionSettings
	LogFactory
	App *MockApp
}

func TestSessionFactorySuite(t *testing.T) {
	suite.Run(t, new(SessionFactorySuite))
}

func (s *SessionFactorySuite) SetupTest() {
	s.sessionFactory = sessionFactory{}
	s.SessionID = SessionID{BeginString: "FIX.4.2", TargetCompID: "TW", SenderCompID: "ISLD"}
	s.MessageStoreFactory = NewMemoryStoreFactory()
	s.SessionSettings = NewSessionSettings()
	s.LogFactory = nullLogFactory{}
	s.App = new(MockApp)
}

func (s *SessionFactorySuite) TestDefaults() {
	session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.Nil(err)
	s.NotNil(session)

	s.False(session.ResetOnLogon)
	s.False(session.RefreshOnLogon)
	s.False(session.ResetOnLogout)
	s.False(session.ResetOnDisconnect)
	s.Nil(session.SessionTime, "By default, start and end time unset")
	s.Equal("", session.DefaultApplVerID)
	s.False(session.InitiateLogon)
	s.Equal(0, session.ResendRequestChunkSize)
	s.False(session.EnableLastMsgSeqNumProcessed)
	s.False(session.SkipCheckLatency)
	s.Equal(Millis, session.timestampPrecision)
	s.Equal(120*time.Second, session.MaxLatency)
	s.False(session.DisableMessagePersist)
}

func (s *SessionFactorySuite) TestResetOnLogon() {
	var tests = []struct {
		setting  string
		expected bool
	}{{"Y", true}, {"N", false}}

	for _, test := range tests {
		s.SetupTest()
		s.SessionSettings.Set(config.ResetOnLogon, test.setting)
		session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
		s.Nil(err)
		s.NotNil(session)

		s.Equal(test.expected, session.ResetOnLogon)
	}
}

func (s *SessionFactorySuite) TestRefreshOnLogon() {
	var tests = []struct {
		setting  string
		expected bool
	}{{"Y", true}, {"N", false}}

	for _, test := range tests {
		s.SetupTest()
		s.SessionSettings.Set(config.RefreshOnLogon, test.setting)
		session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
		s.Nil(err)
		s.NotNil(session)

		s.Equal(test.expected, session.RefreshOnLogon)
	}
}

func (s *SessionFactorySuite) TestResetOnLogout() {
	var tests = []struct {
		setting  string
		expected bool
	}{{"Y", true}, {"N", false}}

	for _, test := range tests {
		s.SetupTest()
		s.SessionSettings.Set(config.ResetOnLogout, test.setting)
		session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
		s.Nil(err)
		s.NotNil(session)

		s.Equal(test.expected, session.ResetOnLogout)
	}
}

func (s *SessionFactorySuite) TestResetOnDisconnect() {
	var tests = []struct {
		setting  string
		expected bool
	}{{"Y", true}, {"N", false}}

	for _, test := range tests {
		s.SetupTest()
		s.SessionSettings.Set(config.ResetOnDisconnect, test.setting)
		session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
		s.Nil(err)
		s.NotNil(session)

		s.Equal(test.expected, session.ResetOnDisconnect)
	}
}

func (s *SessionFactorySuite) TestResendRequestChunkSize() {
	s.SessionSettings.Set(config.ResendRequestChunkSize, "2500")
	session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.Nil(err)
	s.NotNil(session)
	s.Equal(2500, session.ResendRequestChunkSize)

	s.SessionSettings.Set(config.ResendRequestChunkSize, "notanint")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)
}

func (s *SessionFactorySuite) TestEnableLastMsgSeqNumProcessed() {
	var tests = []struct {
		setting  string
		expected bool
	}{{"Y", true}, {"N", false}}

	for _, test := range tests {
		s.SetupTest()
		s.SessionSettings.Set(config.EnableLastMsgSeqNumProcessed, test.setting)
		session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
		s.Nil(err)
		s.NotNil(session)

		s.Equal(test.expected, session.EnableLastMsgSeqNumProcessed)
	}
}

func (s *SessionFactorySuite) TestCheckLatency() {
	var tests = []struct {
		setting  string
		expected bool
	}{{"Y", false}, {"N", true}}

	for _, test := range tests {
		s.SetupTest()
		s.SessionSettings.Set(config.CheckLatency, test.setting)
		session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
		s.Nil(err)
		s.NotNil(session)

		s.Equal(test.expected, session.SkipCheckLatency)
	}
}

func (s *SessionFactorySuite) TestStartAndEndTime() {
	s.SessionSettings.Set(config.StartTime, "12:00:00")
	s.SessionSettings.Set(config.EndTime, "14:00:00")
	session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.Nil(err)
	s.NotNil(session.SessionTime)

	s.Equal(
		*internal.NewUTCTimeRange(internal.NewTimeOfDay(12, 0, 0), internal.NewTimeOfDay(14, 0, 0)),
		*session.SessionTime,
	)
}

func (s *SessionFactorySuite) TestStartAndEndTimeAndTimeZone() {
	s.SessionSettings.Set(config.StartTime, "12:00:00")
	s.SessionSettings.Set(config.EndTime, "14:00:00")
	s.SessionSettings.Set(config.TimeZone, "Local")

	session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.Nil(err)
	s.NotNil(session.SessionTime)

	s.Equal(
		*internal.NewTimeRangeInLocation(internal.NewTimeOfDay(12, 0, 0), internal.NewTimeOfDay(14, 0, 0), time.Local),
		*session.SessionTime,
	)
}

func (s *SessionFactorySuite) TestStartAndEndTimeAndStartAndEndDay() {
	var tests = []struct {
		startDay, endDay string
	}{
		{"Sunday", "Thursday"},
		{"Sun", "Thu"},
	}

	for _, test := range tests {
		s.SetupTest()

		s.SessionSettings.Set(config.StartTime, "12:00:00")
		s.SessionSettings.Set(config.EndTime, "14:00:00")
		s.SessionSettings.Set(config.StartDay, test.startDay)
		s.SessionSettings.Set(config.EndDay, test.endDay)

		session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
		s.Nil(err)
		s.NotNil(session.SessionTime)

		s.Equal(
			*internal.NewUTCWeekRange(
				internal.NewTimeOfDay(12, 0, 0), internal.NewTimeOfDay(14, 0, 0),
				time.Sunday, time.Thursday,
			),
			*session.SessionTime,
		)
	}
}

func (s *SessionFactorySuite) TestStartAndEndTimeAndStartAndEndDayAndTimeZone() {
	s.SessionSettings.Set(config.StartTime, "12:00:00")
	s.SessionSettings.Set(config.EndTime, "14:00:00")
	s.SessionSettings.Set(config.StartDay, "Sunday")
	s.SessionSettings.Set(config.EndDay, "Thursday")
	s.SessionSettings.Set(config.TimeZone, "Local")

	session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.Nil(err)
	s.NotNil(session.SessionTime)

	s.Equal(
		*internal.NewWeekRangeInLocation(
			internal.NewTimeOfDay(12, 0, 0), internal.NewTimeOfDay(14, 0, 0),
			time.Sunday, time.Thursday, time.Local,
		),
		*session.SessionTime,
	)
}

func (s *SessionFactorySuite) TestMissingStartOrEndTime() {
	s.SessionSettings.Set(config.StartTime, "12:00:00")
	_, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)

	s.SetupTest()
	s.SessionSettings.Set(config.EndTime, "14:00:00")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)
}

func (s *SessionFactorySuite) TestStartOrEndTimeParseError() {
	s.SessionSettings.Set(config.StartTime, "1200:00")
	s.SessionSettings.Set(config.EndTime, "14:00:00")

	_, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)

	s.SessionSettings.Set(config.StartTime, "12:00:00")
	s.SessionSettings.Set(config.EndTime, "")

	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)
}

func (s *SessionFactorySuite) TestInvalidTimeZone() {
	s.SessionSettings.Set(config.StartTime, "12:00:00")
	s.SessionSettings.Set(config.EndTime, "14:00:00")
	s.SessionSettings.Set(config.TimeZone, "not valid")

	_, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)
}

func (s *SessionFactorySuite) TestMissingStartOrEndDay() {
	s.SessionSettings.Set(config.StartTime, "12:00:00")
	s.SessionSettings.Set(config.EndTime, "14:00:00")
	s.SessionSettings.Set(config.StartDay, "Thursday")
	_, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)

	s.SetupTest()
	s.SessionSettings.Set(config.StartTime, "12:00:00")
	s.SessionSettings.Set(config.EndTime, "14:00:00")
	s.SessionSettings.Set(config.EndDay, "Sunday")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)
}

func (s *SessionFactorySuite) TestStartOrEndDayParseError() {
	s.SessionSettings.Set(config.StartTime, "12:00:00")
	s.SessionSettings.Set(config.EndTime, "14:00:00")
	s.SessionSettings.Set(config.StartDay, "notvalid")
	s.SessionSettings.Set(config.EndDay, "Sunday")
	_, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)

	s.SetupTest()
	s.SessionSettings.Set(config.StartTime, "12:00:00")
	s.SessionSettings.Set(config.EndTime, "14:00:00")
	s.SessionSettings.Set(config.StartDay, "Sunday")
	s.SessionSettings.Set(config.EndDay, "blah")

	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)
}

func (s *SessionFactorySuite) TestDefaultApplVerID() {
	s.SessionID = SessionID{BeginString: BeginStringFIXT11, TargetCompID: "TW", SenderCompID: "ISLD"}

	var tests = []struct{ expected, config string }{
		{"2", "2"},
		{"2", "FIX.4.0"},
		{"3", "3"},
		{"3", "FIX.4.1"},
		{"4", "4"},
		{"4", "FIX.4.2"},
		{"5", "5"},
		{"5", "FIX.4.3"},
		{"6", "6"},
		{"6", "FIX.4.4"},
		{"7", "7"},
		{"7", "FIX.5.0"},
		{"8", "8"},
		{"8", "FIX.5.0SP1"},
		{"9", "9"},
		{"9", "FIX.5.0SP2"},
	}

	for _, test := range tests {
		s.SessionSettings.Set(config.DefaultApplVerID, test.config)
		session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
		s.Nil(err)
		s.Equal(test.expected, session.DefaultApplVerID)
	}
}

func (s *SessionFactorySuite) TestNewSessionBuildInitiators() {
	s.sessionFactory.BuildInitiators = true
	s.SessionSettings.Set(config.HeartBtInt, "34")
	s.SessionSettings.Set(config.SocketConnectHost, "127.0.0.1")
	s.SessionSettings.Set(config.SocketConnectPort, "5000")

	session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.Nil(err)
	s.True(session.InitiateLogon)
	s.Equal(34*time.Second, session.HeartBtInt)
	s.Equal(30*time.Second, session.ReconnectInterval)
	s.Equal(10*time.Second, session.LogonTimeout)
	s.Equal(2*time.Second, session.LogoutTimeout)
	s.Equal("127.0.0.1:5000", session.SocketConnectAddress[0])
}

func (s *SessionFactorySuite) TestNewSessionBuildInitiatorsValidHeartBtInt() {
	s.sessionFactory.BuildInitiators = true

	_, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "HeartBtInt should be required for initiators")

	s.SessionSettings.Set(config.HeartBtInt, "not a number")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "HeartBtInt must be a number")

	s.SessionSettings.Set(config.HeartBtInt, "0")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "HeartBtInt must be greater than zero")

	s.SessionSettings.Set(config.HeartBtInt, "-20")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "HeartBtInt must be greater than zero")
}

func (s *SessionFactorySuite) TestNewSessionBuildInitiatorsValidReconnectInterval() {
	s.sessionFactory.BuildInitiators = true
	s.SessionSettings.Set(config.HeartBtInt, "34")
	s.SessionSettings.Set(config.SocketConnectHost, "127.0.0.1")
	s.SessionSettings.Set(config.SocketConnectPort, "3000")

	s.SessionSettings.Set(config.ReconnectInterval, "45")
	session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.Nil(err)
	s.Equal(45*time.Second, session.ReconnectInterval)

	s.SessionSettings.Set(config.ReconnectInterval, "not a number")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "ReconnectInterval must be a number")

	s.SessionSettings.Set(config.ReconnectInterval, "0")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "ReconnectInterval must be greater than zero")

	s.SessionSettings.Set(config.ReconnectInterval, "-20")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "ReconnectInterval must be greater than zero")
}

func (s *SessionFactorySuite) TestNewSessionBuildInitiatorsValidLogoutTimeout() {
	s.sessionFactory.BuildInitiators = true
	s.SessionSettings.Set(config.HeartBtInt, "34")
	s.SessionSettings.Set(config.SocketConnectHost, "127.0.0.1")
	s.SessionSettings.Set(config.SocketConnectPort, "3000")

	s.SessionSettings.Set(config.LogoutTimeout, "45")
	session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.Nil(err)
	s.Equal(45*time.Second, session.LogoutTimeout)

	s.SessionSettings.Set(config.LogoutTimeout, "not a number")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "LogoutTimeout must be a number")

	s.SessionSettings.Set(config.LogoutTimeout, "0")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "LogoutTimeout must be greater than zero")

	s.SessionSettings.Set(config.LogoutTimeout, "-20")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "LogoutTimeout must be greater than zero")
}

func (s *SessionFactorySuite) TestNewSessionBuildInitiatorsValidLogonTimeout() {
	s.sessionFactory.BuildInitiators = true
	s.SessionSettings.Set(config.HeartBtInt, "34")
	s.SessionSettings.Set(config.SocketConnectHost, "127.0.0.1")
	s.SessionSettings.Set(config.SocketConnectPort, "3000")

	s.SessionSettings.Set(config.LogonTimeout, "45")
	session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.Nil(err)
	s.Equal(45*time.Second, session.LogonTimeout)

	s.SessionSettings.Set(config.LogonTimeout, "not a number")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "LogonTimeout must be a number")

	s.SessionSettings.Set(config.LogonTimeout, "0")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "LogonTimeout must be greater than zero")

	s.SessionSettings.Set(config.LogonTimeout, "-20")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "LogonTimeout must be greater than zero")
}

func (s *SessionFactorySuite) TestConfigureSocketConnectAddress() {
	sess := new(session)
	err := s.configureSocketConnectAddress(sess, s.SessionSettings)
	s.NotNil(err, "SocketConnectHost and SocketConnectPort should be required")

	s.SessionSettings.Set(config.SocketConnectHost, "127.0.0.1")
	err = s.configureSocketConnectAddress(sess, s.SessionSettings)
	s.NotNil(err, "SocketConnectHost and SocketConnectPort should be required")

	s.SessionSettings = NewSessionSettings()
	s.SessionSettings.Set(config.SocketConnectPort, "5000")
	err = s.configureSocketConnectAddress(sess, s.SessionSettings)
	s.NotNil(err, "SocketConnectHost and SocketConnectPort should be required")

	var tests = []struct{ host, port, expected string }{
		{"127.0.0.1", "3000", "127.0.0.1:3000"},
		{"example.com", "5000", "example.com:5000"},
		{"2001:db8:a0b:12f0::1", "3001", "[2001:db8:a0b:12f0::1]:3001"},
	}

	for _, test := range tests {
		sess = new(session)
		s.SessionSettings.Set(config.SocketConnectHost, test.host)
		s.SessionSettings.Set(config.SocketConnectPort, test.port)
		err = s.configureSocketConnectAddress(sess, s.SessionSettings)
		s.Nil(err)
		s.Len(sess.SocketConnectAddress, 1)
		s.Equal(test.expected, sess.SocketConnectAddress[0])
	}
}

func (s *SessionFactorySuite) TestConfigureSocketConnectAddressMulti() {
	session := new(session)
	s.SessionSettings.Set(config.SocketConnectHost, "127.0.0.1")
	s.SessionSettings.Set(config.SocketConnectPort, "3000")

	s.SessionSettings.Set(config.SocketConnectHost+"1", "127.0.0.2")
	s.SessionSettings.Set(config.SocketConnectPort+"1", "4000")

	s.SessionSettings.Set(config.SocketConnectHost+"2", "127.0.0.3")
	s.SessionSettings.Set(config.SocketConnectPort+"2", "5000")

	err := s.configureSocketConnectAddress(session, s.SessionSettings)
	s.Require().Nil(err)
	s.Require().Len(session.SocketConnectAddress, 3)
	var tests = []string{
		"127.0.0.1:3000",
		"127.0.0.2:4000",
		"127.0.0.3:5000",
	}

	for i, expected := range tests {
		s.Equal(expected, session.SocketConnectAddress[i])
	}

	s.SessionSettings.Set(config.SocketConnectHost+"3", "127.0.0.4")
	err = s.configureSocketConnectAddress(session, s.SessionSettings)
	s.NotNil(err, "must have both host and port to be valid")

	s.SessionSettings.Set(config.SocketConnectPort+"3", "5000")
	s.SessionSettings.Set(config.SocketConnectPort+"4", "5000")

	err = s.configureSocketConnectAddress(session, s.SessionSettings)
	s.NotNil(err, "must have both host and port to be valid")
}

func (s *SessionFactorySuite) TestNewSessionTimestampPrecision() {
	s.SessionSettings.Set(config.TimeStampPrecision, "blah")

	_, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err)

	var tests = []struct {
		config    string
		precision TimestampPrecision
	}{
		{"SECONDS", Seconds},
		{"MILLIS", Millis},
		{"MICROS", Micros},
		{"NANOS", Nanos},
	}

	for _, test := range tests {
		s.SessionSettings.Set(config.TimeStampPrecision, test.config)
		session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
		s.Nil(err)

		s.Equal(session.timestampPrecision, test.precision)
	}
}

func (s *SessionFactorySuite) TestNewSessionMaxLatency() {
	s.SessionSettings.Set(config.MaxLatency, "not a number")
	_, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "MaxLatency must be a number")

	s.SessionSettings.Set(config.MaxLatency, "-20")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "MaxLatency must be positive")

	s.SessionSettings.Set(config.MaxLatency, "0")
	_, err = s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.NotNil(err, "MaxLatency must be positive")

	s.SessionSettings.Set(config.MaxLatency, "20")
	session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
	s.Nil(err)
	s.Equal(session.MaxLatency, 20*time.Second)
}

func (s *SessionFactorySuite) TestPersistMessages() {
	var tests = []struct {
		setting  string
		expected bool
	}{{"Y", false}, {"N", true}}

	for _, test := range tests {
		s.SetupTest()
		s.SessionSettings.Set(config.PersistMessages, test.setting)
		session, err := s.newSession(s.SessionID, s.MessageStoreFactory, s.SessionSettings, s.LogFactory, s.App)
		s.Nil(err)
		s.NotNil(session)

		s.Equal(test.expected, session.DisableMessagePersist)
	}
}
