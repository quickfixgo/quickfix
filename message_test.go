package quickfix

import (
	"bytes"
	"github.com/quickfixgo/quickfix/enum"
	"testing"
)

var msgResult Message

func BenchmarkParseMessage(b *testing.B) {
	rawMsg := []byte("8=FIX.4.29=10435=D34=249=TW52=20140515-19:49:56.65956=ISLD11=10021=140=154=155=TSLA60=00010101-00:00:00.00010=039")

	var msg Message
	for i := 0; i < b.N; i++ {
		msg, _ = ParseMessage(rawMsg)
	}

	msgResult = msg
}

func TestMessage_ParseMessage(t *testing.T) {
	rawMsg := []byte("8=FIX.4.29=10435=D34=249=TW52=20140515-19:49:56.65956=ISLD11=10021=140=154=155=TSLA60=00010101-00:00:00.00010=039")

	msg, err := ParseMessage(rawMsg)

	if err != nil {
		t.Error("Unexpected error, ", err)
	}

	if !bytes.Equal(rawMsg, msg.rawMessage) {
		t.Error("Expected msg bytes to equal raw bytes")
	}

	expectedBodyBytes := []byte("11=10021=140=154=155=TSLA60=00010101-00:00:00.000")

	if !bytes.Equal(msg.bodyBytes, expectedBodyBytes) {
		t.Error("Incorrect body bytes, got ", string(msg.bodyBytes))
	}

	expectedLenFields := 14
	if len(msg.fields) != expectedLenFields {
		t.Errorf("Expected %v fields, got %v", expectedLenFields, len(msg.fields))
	}

	msgType, err := msg.Header.GetMsgType()
	if err != nil {
		t.Error("Unexpected error, ", err)
	}

	if msgType != enum.MsgType_ORDER_SINGLE {
		t.Errorf("Expected msgType MsgType_ORDER_SINGLE, got %#v", msgType)
	}

	if !msg.Header.HasMsgTypeOf(enum.MsgType_ORDER_SINGLE) {
		t.Errorf("Expected Header.HasMsgTypeOf(MsgType_ORDER_SINGLE) is true")
	}

	if msg.Header.HasMsgTypeOf(enum.MsgType_LOGON) {
		t.Errorf("Expected Header.HasMsgTypeOf(MsgType_LOGON) is false")
	}
}

func TestMessage_parseOutOfOrder(t *testing.T) {
	//allow fields out of order, save for validation
	rawMsg := []byte("8=FIX.4.09=8135=D11=id21=338=10040=154=155=MSFT34=249=TW52=20140521-22:07:0956=ISLD10=250")
	_, err := ParseMessage(rawMsg)

	if err != nil {
		t.Error("Should not have gotten error, got ", err)
	}
}

func TestMessage_Build(t *testing.T) {
	builder := NewMessage()

	builder.Header.SetField(tagBeginString, FIXString(enum.BeginStringFIX44))
	builder.Header.SetField(tagMsgType, FIXString("A"))
	builder.Header.SetField(tagSendingTime, FIXString("20140615-19:49:56"))

	builder.Body.SetField(Tag(553), FIXString("my_user"))
	builder.Body.SetField(Tag(554), FIXString("secret"))

	expectedBytes := []byte("8=FIX.4.49=4935=A52=20140615-19:49:56553=my_user554=secret10=072")
	result, err := builder.Build()
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if !bytes.Equal(expectedBytes, result) {
		t.Error("Unexpected bytes, got ", string(result))
	}
}

func TestMessage_ReBuild(t *testing.T) {
	rawMsg := []byte("8=FIX.4.29=10435=D34=249=TW52=20140515-19:49:56.65956=ISLD11=10021=140=154=155=TSLA60=00010101-00:00:00.00010=039")

	msg, _ := ParseMessage(rawMsg)

	msg.Header.SetField(tagOrigSendingTime, FIXString("20140515-19:49:56.659"))
	msg.Header.SetField(tagSendingTime, FIXString("20140615-19:49:56"))

	msg.Build()

	expectedBytes := []byte("8=FIX.4.29=12635=D34=249=TW52=20140615-19:49:5656=ISLD122=20140515-19:49:56.65911=10021=140=154=155=TSLA60=00010101-00:00:00.00010=128")

	if !bytes.Equal(expectedBytes, msg.rawMessage) {
		t.Errorf("Unexpected bytes,\n +%s\n-%s", msg.rawMessage, expectedBytes)
	}

	expectedBodyBytes := []byte("11=10021=140=154=155=TSLA60=00010101-00:00:00.000")

	if !bytes.Equal(msg.bodyBytes, expectedBodyBytes) {
		t.Error("Incorrect body bytes, got ", string(msg.bodyBytes))
	}
}

func TestMessage_reverseRoute(t *testing.T) {
	msg, _ := ParseMessage([]byte("8=FIX.4.29=17135=D34=249=TW50=KK52=20060102-15:04:0556=ISLD57=AP144=BB115=JCD116=CS128=MG129=CB142=JV143=RY145=BH11=ID21=338=10040=w54=155=INTC60=20060102-15:04:0510=123"))

	builder := msg.reverseRoute()

	var testCases = []struct {
		tag           Tag
		expectedValue string
	}{
		{tagTargetCompID, "TW"},
		{tagTargetSubID, "KK"},
		{tagTargetLocationID, "JV"},
		{tagSenderCompID, "ISLD"},
		{tagSenderSubID, "AP"},
		{tagSenderLocationID, "RY"},
		{tagDeliverToCompID, "JCD"},
		{tagDeliverToSubID, "CS"},
		{tagDeliverToLocationID, "BB"},
		{tagOnBehalfOfCompID, "MG"},
		{tagOnBehalfOfSubID, "CB"},
		{tagOnBehalfOfLocationID, "BH"},
	}

	for _, tc := range testCases {
		var field FIXString
		err := builder.Header.GetField(tc.tag, &field)
		if err != nil {
			t.Error("Unexpected error, ", err)
		}

		if string(field) != tc.expectedValue {
			t.Errorf("Expected %v got %v", tc.expectedValue, field)
		}
	}
}

func TestMessage_reverseRouteIgnoreEmpty(t *testing.T) {
	msg, _ := ParseMessage([]byte("8=FIX.4.09=12835=D34=249=TW52=20060102-15:04:0556=ISLD115=116=CS128=MG129=CB11=ID21=338=10040=w54=155=INTC60=20060102-15:04:0510=123"))
	builder := msg.reverseRoute()

	if builder.Header.Has(tagDeliverToCompID) {
		t.Error("Should not reverse if empty")
	}
}

func TestMessage_reverseRouteFIX40(t *testing.T) {
	//onbehalfof/deliverto location id not supported in fix 4.0

	msg, _ := ParseMessage([]byte("8=FIX.4.09=17135=D34=249=TW50=KK52=20060102-15:04:0556=ISLD57=AP144=BB115=JCD116=CS128=MG129=CB142=JV143=RY145=BH11=ID21=338=10040=w54=155=INTC60=20060102-15:04:0510=123"))

	builder := msg.reverseRoute()

	if builder.Header.Has(tagDeliverToLocationID) {
		t.Error("delivertolocation id not supported in fix40")
	}

	if builder.Header.Has(tagOnBehalfOfLocationID) {
		t.Error("onbehalfof location id not supported in fix40")
	}
}
