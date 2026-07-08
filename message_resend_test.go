package quickfix

// Reproduction + regression test for the resend BodyLength bug.
//
// Drop this into the root of the clear-street/quickfix fork (package quickfix)
// so it can reach the unexported buildWithBodyBytes / bodyBytes.
//
//   - Against the ORIGINAL buildWithBodyBytes (which calls m.cook()): FAILS,
//     reporting declared BodyLength 9=<collapsed> < actual body bytes — the
//     exact "Framing error: Checksum not found" NSCC saw.
//   - Against the PATCHED buildWithBodyBytes: PASSES.

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const soh = "\x01"

// assembleStored builds a valid stored FIX message (correct tag 9 and tag 10)
// from a begin string, the header fields after BodyLength, and the body fields.
func assembleStored(begin string, headerRest, body []string) string {
	inner := ""
	for _, f := range headerRest {
		inner += f + soh
	}
	for _, f := range body {
		inner += f + soh
	}
	pre := begin + soh + "9=" + strconv.Itoa(len(inner)) + soh + inner
	ck := 0
	for i := 0; i < len(pre); i++ {
		ck += int(pre[i])
	}
	ck %= 256
	return pre + "10=" + fmt.Sprintf("%03d", ck) + soh
}

// frameCheck independently re-derives BodyLength and CheckSum from raw bytes,
// the same way a receiving FIX engine frames a message.
func frameCheck(t *testing.T, raw []byte) (declaredLen, actualLen, declaredCk, computedCk int) {
	t.Helper()
	s := string(raw)

	// BodyLength: the "9=<n>" field is second.
	i := strings.Index(s, soh+"9=")
	if i < 0 {
		t.Fatalf("no BodyLength field in %q", s)
	}
	j := strings.Index(s[i+1:], soh)
	if j < 0 {
		t.Fatalf("unterminated BodyLength field")
	}
	declaredLen, _ = strconv.Atoi(s[i+len(soh)+len("9=") : i+1+j])
	startOfBody := i + 1 + j + 1 // first byte after "9=<n>\x01"

	// CheckSum: last field "10=<ccc>". Body runs up to (incl) the SOH before it.
	c := strings.LastIndex(s, soh+"10=")
	if c < 0 {
		t.Fatalf("no CheckSum field in %q", s)
	}
	pos10 := c + 1 // index of the '1' in "10="
	actualLen = pos10 - startOfBody

	end := strings.Index(s[pos10:], soh)
	declaredCk, _ = strconv.Atoi(s[pos10+len("10=") : pos10+end])
	for k := 0; k < pos10; k++ {
		computedCk += int(raw[k])
	}
	computedCk %= 256
	return
}

func TestBuildWithBodyBytes_RepeatingGroupFraming(t *testing.T) {
	// A TradeCaptureReport-shaped message with repeating groups: repeated
	// tags (54, 453, 448, 452) are what FieldMap.add collapses on parse when
	// no app DataDictionary is configured.
	begin := "8=FIX.4.4"
	headerRest := []string{
		"35=AE", "34=2", "49=CLSTFIX1", "52=20260706-08:00:00.104", "56=NSCC", "115=CLST",
	}
	body := []string{
		"22=1", "30=100", "31=639.18", "32=2", "552=2",
		"54=1", "453=2", "448=695", "452=4", "448=ALG1", "452=1",
		"54=2", "453=1", "448=9132", "452=18",
		"15=USD", "381=1278.36", "118=1278.36", "570=N", "571=1",
	}

	stored := assembleStored(begin, headerRest, body)

	// Sanity: the stored message itself must frame cleanly.
	if dl, al, dc, cc := frameCheck(t, []byte(stored)); dl != al || dc != cc {
		t.Fatalf("stored message is not self-consistent: 9=%d actual=%d 10=%03d computed=%03d", dl, al, dc, cc)
	}

	msg := NewMessage()
	if err := ParseMessage(msg, bytes.NewBufferString(stored)); err != nil {
		t.Fatalf("ParseMessage: %v", err)
	}

	// Mimic session.resend(): stamp PossDupFlag + OrigSendingTime on the header.
	msg.Header.SetField(tagPossDupFlag, FIXBoolean(true))
	var origSendingTime FIXString
	if err := msg.Header.GetField(tagSendingTime, &origSendingTime); err == nil {
		msg.Header.SetField(tagOrigSendingTime, origSendingTime)
	}

	out := msg.buildWithBodyBytes(msg.bodyBytes)

	declaredLen, actualLen, declaredCk, computedCk := frameCheck(t, out)
	if declaredLen != actualLen {
		t.Fatalf("BodyLength wrong: header says 9=%d but body is %d bytes -> peer reports 'Checksum not found'.\nout=%q",
			declaredLen, actualLen, strings.ReplaceAll(string(out), soh, "|"))
	}
	if declaredCk != computedCk {
		t.Fatalf("CheckSum wrong: header says 10=%03d but computed %03d", declaredCk, computedCk)
	}

	// The resent body must still contain every repeating-group entry.
	os := strings.ReplaceAll(string(out), soh, "|")
	for _, want := range []string{"|448=695|", "|448=ALG1|", "|448=9132|", "|54=1|", "|54=2|"} {
		if !strings.Contains(os, want) {
			t.Fatalf("resent message dropped repeating-group field %q: %s", want, os)
		}
	}
}
