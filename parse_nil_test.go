package quickfix

import (
	"bytes"
	"testing"
)

// TestParseMessageZeroValue ensures ParseMessage works on a zero-value Message
// without panicking. Prior to the fix, passing &msg (zero-value) caused a nil
// pointer dereference on rwLock because FieldMap.rwLock is a *sync.RWMutex that
// is only set by Init().
func TestParseMessageZeroValue(t *testing.T) {
	rawMsg := bytes.NewBufferString("8=FIX.4.2\x019=104\x0135=D\x0134=2\x0149=TW\x0152=20140515-19:49:56.659\x0156=ISLD\x0111=100\x0121=1\x0140=1\x0154=1\x0155=TSLA\x0160=00010101-00:00:00.000\x0110=039\x01")

	var msg Message
	err := ParseMessage(&msg, rawMsg)
	// A parse error is acceptable; a panic is not.
	if err != nil {
		t.Logf("ParseMessage returned error (acceptable): %v", err)
	}
}
