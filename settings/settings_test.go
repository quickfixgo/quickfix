package settings

import (
	"strings"
	"testing"
)

func TestSettings_New(t *testing.T) {
	s := New()

	if s == nil {
		t.Error("New returned nil")
	}

	globalSettings := s.GlobalSettings()
	if globalSettings == nil {
		t.Error("global settings is nil")
	}

	sessionSettings := s.SessionSettings()
	if sessionSettings == nil {
		t.Error("session settings is nil")
	}

	if len(sessionSettings) != 0 {
		t.Errorf("Expected %v settings, got %v", 0, len(sessionSettings))
	}
}

func TestSettings_AddSession(t *testing.T) {
	s := New()
	globalSettings := s.GlobalSettings()

	globalSettings.Set(BeginString, "blah")
	globalSettings.Set(SocketAcceptPort, "1000")

	s1 := NewSessionSettings()
	s1.Set(BeginString, "foo")

	s2 := NewSessionSettings()
	s2.Set(ResetOnLogon, "Y")

	s.AddSession(s1)
	s.AddSession(s2)

	sessionSettings := s.SessionSettings()

	if len(sessionSettings) != 2 {
		t.Errorf("Expected %v settings, got %v", 2, len(sessionSettings))
	}

	var cases = []struct {
		index    int
		input    string
		expected string
	}{
		{0, BeginString, "foo"},
		{0, SocketAcceptPort, "1000"},
		{1, BeginString, "blah"},
		{1, SocketAcceptPort, "1000"},
		{1, ResetOnLogon, "Y"},
	}

	for _, tc := range cases {
		settings := sessionSettings[tc.index]

		actual, err := settings.Setting(tc.input)
		if err != nil {
			t.Error("Unexpected Error", err)
		}

		if actual != tc.expected {
			t.Errorf("Got %v, expected %v", actual, tc.expected)
		}
	}
}

func TestSettings_Read(t *testing.T) {
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
TargetCompID=INCA
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
	s, err := Read(stringReader)

	if err != nil {
		t.Error("Error in Read: ", err)
	}

	if s == nil {
		t.Error("settings is nil")
	}

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

		if err != nil {
			t.Error("Got Error checking global", err)
		}

		if actual != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, actual)
		}
	}

	sessionSettings := s.SessionSettings()

	if len(sessionSettings) != 3 {
		t.Errorf("Expected %v sessions, got %v", 3, len(sessionSettings))
	}

	var sessionTCs = []struct {
		index    int
		setting  string
		expected string
	}{
		{0, "ConnectionType", "initiator"},
		{0, "ReconnectInterval", "60"},
		{0, "SenderCompID", "TW"},
		{0, "BeginString", "FIX.4.1"},
		{0, "TargetCompID", "ARCA"},
		{0, "StartTime", "12:30:00"},
		{0, "EndTime", "23:30:00"},
		{0, "HeartBtInt", "20"},
		{0, "SocketConnectPort", "9823"},
		{0, "SocketConnectHost", "123.123.123.123"},
		{0, "DataDictionary", "somewhere/FIX41.xml"},

		{1, "ConnectionType", "initiator"},
		{1, "ReconnectInterval", "60"},
		{1, "SenderCompID", "TW"},
		{1, "BeginString", "FIX.4.0"},
		{1, "TargetCompID", "ISLD"},
		{1, "StartTime", "12:00:00"},
		{1, "EndTime", "23:00:00"},
		{1, "HeartBtInt", "30"},
		{1, "SocketConnectPort", "8323"},
		{1, "SocketConnectHost", "23.23.23.23"},
		{1, "DataDictionary", "somewhere/FIX40.xml"},

		{2, "ConnectionType", "initiator"},
		{2, "BeginString", "FIX.4.2"},
		{2, "SenderCompID", "TW"},
		{2, "TargetCompID", "INCA"},
		{2, "StartTime", "12:30:00"},
		{2, "EndTime", "21:30:00"},
		{2, "ReconnectInterval", "30"},
		{2, "HeartBtInt", "30"},
		{2, "SocketConnectPort", "6523"},
		{2, "SocketConnectHost", "3.3.3.3"},
		{2, "SocketConnectPort1", "8392"},
		{2, "SocketConnectHost1", "8.8.8.8"},
		{2, "SocketConnectPort2", "2932"},
		{2, "SocketConnectHost2", "12.12.12.12"},
		{2, "DataDictionary", "somewhere/FIX42.xml"},
	}

	for _, tc := range sessionTCs {
		settings := sessionSettings[tc.index]
		actual, err := settings.Setting(tc.setting)

		if err != nil {
			t.Error("Got Error", err)
		}

		if tc.expected != actual {
			t.Errorf("Expected %v, got %v", tc.expected, actual)
		}
	}

}
