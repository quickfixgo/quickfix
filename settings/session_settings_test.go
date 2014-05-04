package settings

import (
	"testing"
)

func TestSessionSettings_StringSettings(t *testing.T) {
	s := NewSessionSettings()
	if s == nil {
		t.Error("NewSessionSettings returned nil")
	}

	s.Set(BeginString, "foo")
	s.Set(BeginString, "blah")
	s.Set(SenderCompID, "bar")

	if ok := s.HasSetting("DoesNotExist"); ok {
		t.Error("HasSetting returned true for setting that doesn't exist")
	}

	if ok := s.HasSetting(BeginString); !ok {
		t.Error("HasSetting returned false for setting that does exist")
	}

	if val, err := s.Setting(BeginString); err != nil {
		t.Error("Got error requesing setting", err)
	} else {
		if val != "blah" {
			t.Errorf("Expected %v got %v", "blah", val)
		}
	}
}

func TestSessionSettings_IntSettings(t *testing.T) {
	s := NewSessionSettings()
	if _, err := s.IntSetting(SocketAcceptPort); err == nil {
		t.Error("Expected error for unknown setting")
	}

	s.Set(SocketAcceptPort, "notanint")
	if _, err := s.IntSetting(SocketAcceptPort); err == nil {
		t.Error("Expected error for unparsable value")
	}

	s.Set(SocketAcceptPort, "1005")
	val, err := s.IntSetting(SocketAcceptPort)
	if err != nil {
		t.Error("Unexpected err", err)
	}

	if val != 1005 {
		t.Errorf("Expected %v, got %v", 1005, val)
	}
}

func TestSessionSettings_BoolSettins(t *testing.T) {
	s := NewSessionSettings()
	if _, err := s.BoolSetting(ResetOnLogon); err == nil {
		t.Error("Expected error for unknown setting")
	}

	s.Set(ResetOnLogon, "notabool")
	if _, err := s.BoolSetting(ResetOnLogon); err == nil {
		t.Error("Expected error for unparsable value")
	}

	var boolTests = []struct {
		input    string
		expected bool
	}{
		{"Y", true},
		{"y", true},
		{"N", false},
		{"n", false},
	}

	for _, bt := range boolTests {
		s.Set(ResetOnLogon, bt.input)

		actual, err := s.BoolSetting(ResetOnLogon)
		if err != nil {
			t.Error("Unexpected err", err)
		}

		if actual != bt.expected {
			t.Errorf("Expected %v, got %v", bt.expected, actual)
		}
	}
}

func TestSessionSettings_Clone(t *testing.T) {
	s := NewSessionSettings()

	var cloneTests = []struct {
		input    string
		expected string
	}{
		{SocketAcceptPort, "101"},
		{BeginString, "foo"},
		{ResetOnLogon, "N"},
	}

	for _, ct := range cloneTests {
		s.Set(ct.input, ct.expected)
	}

	cloned := s.clone()
	if cloned == nil {
		t.Error("clone returned nil")
	}

	for _, ct := range cloneTests {
		actual, err := cloned.Setting(ct.input)

		if err != nil {
			t.Error("Unexpected Error", err)
		}

		if ct.expected != actual {
			t.Errorf("Expected %v got %v", ct.expected, actual)
		}
	}
}

func TestSessionSettings_Overlay(t *testing.T) {
	s := NewSessionSettings()
	overlay := NewSessionSettings()

	s.Set(SocketAcceptPort, "101")
	s.Set(BeginString, "foo")

	overlay.Set(SocketAcceptPort, "102")
	overlay.Set(SenderCompID, "blah")

	var overlayTests = []struct {
		input    string
		expected string
	}{
		{SocketAcceptPort, "102"},
		{BeginString, "foo"},
		{SenderCompID, "blah"},
	}

	s.overlay(overlay)

	for _, ot := range overlayTests {
		actual, err := s.Setting(ot.input)
		if err != nil {
			t.Error("Unexpected error", err)
		}

		if actual != ot.expected {
			t.Errorf("expected %v got %v", ot.expected, actual)
		}
	}
}
