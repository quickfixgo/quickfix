package quickfix

import "testing"

func TestSessionID_String(t *testing.T) {
	s := SessionID{BeginString: "FIX.4.2", SenderCompID: "SND", TargetCompID: "TAR"}

	expect := "FIX.4.2:SND->TAR"
	actual := s.String()

	if expect != actual {
		t.Errorf("Expected %v got %v ", expect, actual)
	}
}
