package quickfix

import (
	"strings"
	"testing"

	"github.com/quickfixgo/quickfix/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSettings_New(t *testing.T) {
	s := NewSettings()
	assert.NotNil(t, s)

	globalSettings := s.GlobalSettings()
	assert.NotNil(t, globalSettings)

	sessionSettings := s.SessionSettings()
	assert.NotNil(t, sessionSettings)
	assert.Empty(t, sessionSettings)
}

func TestSettings_AddSession(t *testing.T) {
	s := NewSettings()
	globalSettings := s.GlobalSettings()

	globalSettings.Set(config.BeginString, "FIX.4.0")
	globalSettings.Set(config.SocketAcceptPort, "1000")

	s1 := NewSessionSettings()
	s1.Set(config.BeginString, "FIX.4.1")
	s1.Set(config.SenderCompID, "CB")
	s1.Set(config.TargetCompID, "SS")

	sessionID1 := SessionID{BeginString: "FIX.4.1", SenderCompID: "CB", TargetCompID: "SS"}

	s2 := NewSessionSettings()
	s2.Set(config.ResetOnLogon, "Y")
	s2.Set(config.SenderCompID, "CB")
	s2.Set(config.TargetCompID, "SS")

	sessionID2 := SessionID{BeginString: "FIX.4.0", SenderCompID: "CB", TargetCompID: "SS"}

	var addCases = []struct {
		settings          *SessionSettings
		expectedSessionID SessionID
	}{
		{s1, sessionID1},
		{s2, sessionID2},
	}

	for _, tc := range addCases {
		sid, err := s.AddSession(tc.settings)
		if err != nil {
			t.Fatal("Got Error adding session", err)
		}

		if sid != tc.expectedSessionID {
			t.Fatalf("expected %v got %v", tc.expectedSessionID, sid)
		}
	}

	s3 := NewSessionSettings()
	s3.Set(config.ResetOnLogon, "Y")
	s3.Set(config.SenderCompID, "CB")
	s3.Set(config.TargetCompID, "SS")
	if _, err := s.AddSession(s3); err == nil {
		t.Error("Expected error for adding duplicate session")
	}

	sessionSettings := s.SessionSettings()

	if len(sessionSettings) != 2 {
		t.Errorf("Expected %v settings, got %v", 2, len(sessionSettings))
	}

	var cases = []struct {
		sessionID SessionID
		input     string
		expected  string
	}{
		{sessionID1, config.BeginString, "FIX.4.1"},
		{sessionID1, config.SocketAcceptPort, "1000"},
		{sessionID2, config.BeginString, "FIX.4.0"},
		{sessionID2, config.SocketAcceptPort, "1000"},
		{sessionID2, config.ResetOnLogon, "Y"},
	}

	for _, tc := range cases {
		settings := sessionSettings[tc.sessionID]

		actual, err := settings.Setting(tc.input)
		if err != nil {
			t.Error("Unexpected Error", err)
		}

		if actual != tc.expected {
			t.Errorf("Got %v, expected %v", actual, tc.expected)
		}
	}
}

func TestSettings_ParseSettings(t *testing.T) {
	cfg := `
# default settings for sessions
[DEFAULT]
ConnectionType=initiator
ReconnectInterval=60
SenderCompID=TW


# session definition
[SESSION]
# inherit ConnectionType, ReconnectInterval and SenderCompID from default

BeginString=FIX.4.1
TargetCompID=ARCA
StartTime=12:30:00
EndTime=23:30:00
HeartBtInt=20
SocketConnectPort=9823
SocketConnectHost=123.123.123.123
DataDictionary=somewhere/FIX41.xml


[SESSION]
BeginString=FIX.4.0
TargetCompID=ISLD
StartTime=12:00:00
EndTime=23:00:00
HeartBtInt=30
SocketConnectPort=8323
SocketConnectHost=23.23.23.23
DataDictionary=somewhere/FIX40.xml


[SESSION]
BeginString=FIX.4.2
SenderSubID=TWSub
SenderLocationID=TWLoc
TargetCompID=INCA
TargetSubID=INCASub
TargetLocationID=INCALoc
StartTime=12:30:00
EndTime=21:30:00
# overide default setting for RecconnectInterval
ReconnectInterval=30
HeartBtInt=30
SocketConnectPort=6523
SocketConnectHost=3.3.3.3
# (optional) alternate connection ports and hosts to cycle through on failover
SocketConnectPort1=8392
SocketConnectHost1=8.8.8.8
SocketConnectPort2=2932
SocketConnectHost2=12.12.12.12
DataDictionary=somewhere/FIX42.xml
`

	stringReader := strings.NewReader(cfg)
	s, err := ParseSettings(stringReader)
	assert.Nil(t, err)
	assert.NotNil(t, s)

	var globalTCs = []struct {
		setting  string
		expected string
	}{
		{"ConnectionType", "initiator"},
		{"ReconnectInterval", "60"},
		{"SenderCompID", "TW"},
	}

	globalSettings := s.GlobalSettings()
	for _, tc := range globalTCs {
		actual, err := globalSettings.Setting(tc.setting)
		assert.Nil(t, err)

		assert.Equal(t, tc.expected, actual)
	}

	sessionSettings := s.SessionSettings()
	assert.Len(t, sessionSettings, 3)

	sessionID1 := SessionID{BeginString: "FIX.4.1", SenderCompID: "TW", TargetCompID: "ARCA"}
	sessionID2 := SessionID{BeginString: "FIX.4.0", SenderCompID: "TW", TargetCompID: "ISLD"}
	sessionID3 := SessionID{
		BeginString:  "FIX.4.2",
		SenderCompID: "TW", SenderSubID: "TWSub", SenderLocationID: "TWLoc",
		TargetCompID: "INCA", TargetSubID: "INCASub", TargetLocationID: "INCALoc"}

	var sessionTCs = []struct {
		sessionID SessionID
		setting   string
		expected  string
	}{
		{sessionID1, "ConnectionType", "initiator"},
		{sessionID1, "ReconnectInterval", "60"},
		{sessionID1, "SenderCompID", "TW"},
		{sessionID1, "BeginString", "FIX.4.1"},
		{sessionID1, "TargetCompID", "ARCA"},
		{sessionID1, "StartTime", "12:30:00"},
		{sessionID1, "EndTime", "23:30:00"},
		{sessionID1, "HeartBtInt", "20"},
		{sessionID1, "SocketConnectPort", "9823"},
		{sessionID1, "SocketConnectHost", "123.123.123.123"},
		{sessionID1, "DataDictionary", "somewhere/FIX41.xml"},

		{sessionID2, "ConnectionType", "initiator"},
		{sessionID2, "ReconnectInterval", "60"},
		{sessionID2, "SenderCompID", "TW"},
		{sessionID2, "BeginString", "FIX.4.0"},
		{sessionID2, "TargetCompID", "ISLD"},
		{sessionID2, "StartTime", "12:00:00"},
		{sessionID2, "EndTime", "23:00:00"},
		{sessionID2, "HeartBtInt", "30"},
		{sessionID2, "SocketConnectPort", "8323"},
		{sessionID2, "SocketConnectHost", "23.23.23.23"},
		{sessionID2, "DataDictionary", "somewhere/FIX40.xml"},

		{sessionID3, "ConnectionType", "initiator"},
		{sessionID3, "BeginString", "FIX.4.2"},
		{sessionID3, "SenderCompID", "TW"},
		{sessionID3, "TargetCompID", "INCA"},
		{sessionID3, "StartTime", "12:30:00"},
		{sessionID3, "EndTime", "21:30:00"},
		{sessionID3, "ReconnectInterval", "30"},
		{sessionID3, "HeartBtInt", "30"},
		{sessionID3, "SocketConnectPort", "6523"},
		{sessionID3, "SocketConnectHost", "3.3.3.3"},
		{sessionID3, "SocketConnectPort1", "8392"},
		{sessionID3, "SocketConnectHost1", "8.8.8.8"},
		{sessionID3, "SocketConnectPort2", "2932"},
		{sessionID3, "SocketConnectHost2", "12.12.12.12"},
		{sessionID3, "DataDictionary", "somewhere/FIX42.xml"},
	}

	for _, tc := range sessionTCs {
		settings, ok := sessionSettings[tc.sessionID]
		require.True(t, ok, "No Session recalled for %v", tc.sessionID)
		actual, err := settings.Setting(tc.setting)

		assert.Nil(t, err)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestSettings_ParseSettings_WithEqualsSignInValue(t *testing.T) {
	s, err := ParseSettings(strings.NewReader(`
[DEFAULT]
ConnectionType=initiator
SQLDriver=mysql
SQLDataSourceName=root:root@/quickfix?parseTime=true&loc=UTC

[SESSION]
BeginString=FIX.4.2
SenderCompID=SENDER
TargetCompID=TARGET`))

	require.Nil(t, err)
	require.NotNil(t, s)

	sessionSettings := s.SessionSettings()[SessionID{BeginString: "FIX.4.2", SenderCompID: "SENDER", TargetCompID: "TARGET"}]
	val, err := sessionSettings.Setting("SQLDataSourceName")
	assert.Nil(t, err)
	assert.Equal(t, `root:root@/quickfix?parseTime=true&loc=UTC`, val)
}

func TestSettings_SessionIDFromSessionSettings(t *testing.T) {
	var testCases = []struct {
		globalBeginString   string
		globalTargetCompID  string
		globalSenderCompID  string
		sessionBeginString  string
		sessionTargetCompID string
		sessionSenderCompID string
		sessionQualifier    string
		expectedSessionID   SessionID
	}{
		{globalBeginString: "FIX.4.0", globalTargetCompID: "CB", globalSenderCompID: "SS",
			expectedSessionID: SessionID{BeginString: "FIX.4.0", TargetCompID: "CB", SenderCompID: "SS"}},

		{sessionBeginString: "FIX.4.1", sessionTargetCompID: "GE", sessionSenderCompID: "LE",
			expectedSessionID: SessionID{BeginString: "FIX.4.1", TargetCompID: "GE", SenderCompID: "LE"}},

		{globalBeginString: "FIX.4.2", globalTargetCompID: "CB", sessionTargetCompID: "GE", sessionSenderCompID: "LE", sessionQualifier: "J",
			expectedSessionID: SessionID{BeginString: "FIX.4.2", TargetCompID: "GE", SenderCompID: "LE", Qualifier: "J"}},
	}

	for _, tc := range testCases {
		globalSettings := NewSessionSettings()
		sessionSettings := NewSessionSettings()

		if tc.globalBeginString != "" {
			globalSettings.Set(config.BeginString, tc.globalBeginString)
		}

		if tc.sessionBeginString != "" {
			sessionSettings.Set(config.BeginString, tc.sessionBeginString)
		}

		if tc.globalTargetCompID != "" {
			globalSettings.Set(config.TargetCompID, tc.globalTargetCompID)
		}

		if tc.sessionTargetCompID != "" {
			sessionSettings.Set(config.TargetCompID, tc.sessionTargetCompID)
		}

		if tc.globalSenderCompID != "" {
			globalSettings.Set(config.SenderCompID, tc.globalSenderCompID)
		}

		if tc.sessionSenderCompID != "" {
			sessionSettings.Set(config.SenderCompID, tc.sessionSenderCompID)
		}

		if len(tc.sessionQualifier) > 0 {
			sessionSettings.Set(config.SessionQualifier, tc.sessionQualifier)
		}

		actualSessionID := sessionIDFromSessionSettings(globalSettings, sessionSettings)

		if tc.expectedSessionID != actualSessionID {
			t.Errorf("Expected %v, got %v", tc.expectedSessionID, actualSessionID)
		}
	}
}
