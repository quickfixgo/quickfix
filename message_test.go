package quickfix

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/stretchr/testify/suite"
)

func BenchmarkParseMessage(b *testing.B) {
	rawMsg := bytes.NewBufferString("8=FIX.4.29=10435=D34=249=TW52=20140515-19:49:56.65956=ISLD11=10021=140=154=155=TSLA60=00010101-00:00:00.00010=039")

	var msg Message
	for i := 0; i < b.N; i++ {
		_ = ParseMessage(&msg, rawMsg)
	}
}

type MessageSuite struct {
	QuickFIXSuite
	msg *Message
}

func TestMessageSuite(t *testing.T) {
	suite.Run(t, new(MessageSuite))
}

func (s *MessageSuite) SetupTest() {
	s.msg = NewMessage()
}

func (s *MessageSuite) TestParseMessageEmpty() {
	rawMsg := bytes.NewBufferString("")

	err := ParseMessage(s.msg, rawMsg)
	s.NotNil(err)
}

func (s *MessageSuite) TestParseMessage() {
	rawMsg := bytes.NewBufferString("8=FIX.4.29=10435=D34=249=TW52=20140515-19:49:56.65956=ISLD11=10021=140=154=155=TSLA60=00010101-00:00:00.00010=039")

	err := ParseMessage(s.msg, rawMsg)
	s.Nil(err)

	s.True(bytes.Equal(rawMsg.Bytes(), s.msg.rawMessage.Bytes()), "Expected msg bytes to equal raw bytes")

	expectedBodyBytes := []byte("11=10021=140=154=155=TSLA60=00010101-00:00:00.000")

	s.True(bytes.Equal(s.msg.bodyBytes, expectedBodyBytes), "Incorrect body bytes, got %s", string(s.msg.bodyBytes))

	s.Equal(14, len(s.msg.fields))

	msgType, err := s.msg.MsgType()
	s.Nil(err)

	s.Equal("D", msgType)
	s.True(s.msg.IsMsgTypeOf("D"))

	s.False(s.msg.IsMsgTypeOf("A"))
}

func (s *MessageSuite) TestParseMessageWithDataDictionary() {
	dict := new(datadictionary.DataDictionary)
	dict.Header = &datadictionary.MessageDef{
		Fields: map[int]*datadictionary.FieldDef{
			10030: nil,
		},
	}
	dict.Trailer = &datadictionary.MessageDef{
		Fields: map[int]*datadictionary.FieldDef{
			5050: nil,
		},
	}
	rawMsg := bytes.NewBufferString("8=FIX.4.29=12635=D34=249=TW52=20140515-19:49:56.65956=ISLD10030=CUST11=10021=140=154=155=TSLA60=00010101-00:00:00.0005050=HELLO10=039")

	err := ParseMessageWithDataDictionary(s.msg, rawMsg, dict, dict)
	s.Nil(err)
	s.FieldEquals(Tag(10030), "CUST", s.msg.Header)
	s.FieldEquals(Tag(5050), "HELLO", s.msg.Trailer)
}

func (s *MessageSuite) TestParseOutOfOrder() {
	//allow fields out of order, save for validation
	rawMsg := bytes.NewBufferString("8=FIX.4.09=8135=D11=id21=338=10040=154=155=MSFT34=249=TW52=20140521-22:07:0956=ISLD10=250")
	s.Nil(ParseMessage(s.msg, rawMsg))
}

func (s *MessageSuite) TestBuild() {
	s.msg.Header.SetField(tagBeginString, FIXString(BeginStringFIX44))
	s.msg.Header.SetField(tagMsgType, FIXString("A"))
	s.msg.Header.SetField(tagSendingTime, FIXString("20140615-19:49:56"))

	s.msg.Body.SetField(Tag(553), FIXString("my_user"))
	s.msg.Body.SetField(Tag(554), FIXString("secret"))

	expectedBytes := []byte("8=FIX.4.49=4935=A52=20140615-19:49:56553=my_user554=secret10=072")
	result := s.msg.build()
	s.True(bytes.Equal(expectedBytes, result), "Unexpected bytes, got %s", string(result))
}

func (s *MessageSuite) TestReBuild() {
	rawMsg := bytes.NewBufferString("8=FIX.4.29=10435=D34=249=TW52=20140515-19:49:56.65956=ISLD11=10021=140=154=155=TSLA60=00010101-00:00:00.00010=039")

	s.Nil(ParseMessage(s.msg, rawMsg))

	s.msg.Header.SetField(tagOrigSendingTime, FIXString("20140515-19:49:56.659"))
	s.msg.Header.SetField(tagSendingTime, FIXString("20140615-19:49:56"))
	s.msg.Header.SetField(tagPossDupFlag, FIXBoolean(true))

	rebuildBytes := s.msg.build()

	expectedBytes := []byte("8=FIX.4.29=13135=D34=243=Y49=TW52=20140615-19:49:5656=ISLD122=20140515-19:49:56.65911=10021=140=154=155=TSLA60=00010101-00:00:00.00010=122")

	s.True(bytes.Equal(expectedBytes, rebuildBytes), "Unexpected bytes,\n +%s\n -%s", rebuildBytes, expectedBytes)

	expectedBodyBytes := []byte("11=10021=140=154=155=TSLA60=00010101-00:00:00.000")

	s.True(bytes.Equal(s.msg.bodyBytes, expectedBodyBytes), "Incorrect body bytes, got %s", string(s.msg.bodyBytes))
}

func (s *MessageSuite) TestReBuildWithRepeatingGroupForResend() {
	// Given the following message with a repeating group
	origHeader := "8=FIXT.1.19=16135=834=349=ISLD52=20240415-03:43:17.92356=TW"
	origBody := "6=1.0011=114=1.0017=131=1.0032=1.0037=138=1.0039=254=155=1150=2151=0.00453=1448=xyzzy447=D452=1"
	origTrailer := "10=014"
	rawMsg := bytes.NewBufferString(origHeader + origBody + origTrailer)

	// When I reparse the message from the store during a resend request
	s.Nil(ParseMessage(s.msg, rawMsg))

	// And I update the headers for resend
	s.msg.Header.SetField(tagOrigSendingTime, FIXString("20240415-03:43:17.923"))
	s.msg.Header.SetField(tagSendingTime, FIXString("20240415-14:41:23.456"))
	s.msg.Header.SetField(tagPossDupFlag, FIXBoolean(true))

	// When I rebuild the message
	rebuildBytes := s.msg.build()

	// Then the repeating groups will not be in the correct order in the rebuilt message (note tags 447, 448, 452, 453)
	expectedBytes := []byte("8=FIXT.1.19=19235=834=343=Y49=ISLD52=20240415-14:41:23.45656=TW122=20240415-03:43:17.9236=1.0011=114=1.0017=131=1.0032=1.0037=138=1.0039=254=155=1150=2151=0.00453=1448=xyzzy447=D452=110=018")
	s.False(bytes.Equal(expectedBytes, rebuildBytes), "Unexpected bytes,\n expected: %s\n  but was: %s", expectedBytes, rebuildBytes)
	expectedOutOfOrderBytes := []byte("8=FIXT.1.19=19235=834=343=Y49=ISLD52=20240415-14:41:23.45656=TW122=20240415-03:43:17.9236=1.0011=114=1.0017=131=1.0032=1.0037=138=1.0039=254=155=1150=2151=0.00447=D448=xyzzy452=1453=110=018")
	s.True(bytes.Equal(expectedOutOfOrderBytes, rebuildBytes), "Unexpected bytes,\n expected: %s\n  but was: %s", expectedOutOfOrderBytes, rebuildBytes)

	// But the bodyBytes will still be correct
	origBodyBytes := []byte(origBody)
	s.True(bytes.Equal(origBodyBytes, s.msg.bodyBytes), "Incorrect body bytes, \n expected: %s\n  but was: %s", origBodyBytes, s.msg.bodyBytes)

	// So when I combine the updated header + the original bodyBytes + the as-is trailer
	resendBytes := s.msg.buildWithBodyBytes(s.msg.bodyBytes)

	// Then the reparsed, rebuilt message will retain the correct ordering of repeating group tags during resend
	s.True(bytes.Equal(expectedBytes, resendBytes), "Unexpected bytes,\n expected: %s\n  but was: %s", expectedBytes, resendBytes)
}

func (s *MessageSuite) TestReverseRoute() {
	s.Nil(ParseMessage(s.msg, bytes.NewBufferString("8=FIX.4.29=17135=D34=249=TW50=KK52=20060102-15:04:0556=ISLD57=AP144=BB115=JCD116=CS128=MG129=CB142=JV143=RY145=BH11=ID21=338=10040=w54=155=INTC60=20060102-15:04:0510=123")))

	builder := s.msg.reverseRoute()

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
		s.Nil(builder.Header.GetField(tc.tag, &field))

		s.Equal(tc.expectedValue, string(field))
	}
}

func (s *MessageSuite) TestReverseRouteIgnoreEmpty() {
	s.Nil(ParseMessage(s.msg, bytes.NewBufferString("8=FIX.4.09=12835=D34=249=TW52=20060102-15:04:0556=ISLD115=116=CS128=MG129=CB11=ID21=338=10040=w54=155=INTC60=20060102-15:04:0510=123")))
	builder := s.msg.reverseRoute()

	s.False(builder.Header.Has(tagDeliverToCompID), "Should not reverse if empty")
}

func (s *MessageSuite) TestReverseRouteFIX40() {
	//onbehalfof/deliverto location id not supported in fix 4.0
	s.Nil(ParseMessage(s.msg, bytes.NewBufferString("8=FIX.4.09=17135=D34=249=TW50=KK52=20060102-15:04:0556=ISLD57=AP144=BB115=JCD116=CS128=MG129=CB142=JV143=RY145=BH11=ID21=338=10040=w54=155=INTC60=20060102-15:04:0510=123")))

	builder := s.msg.reverseRoute()

	s.False(builder.Header.Has(tagDeliverToLocationID), "delivertolocation id not supported in fix40")

	s.False(builder.Header.Has(tagOnBehalfOfLocationID), "onbehalfof location id not supported in fix40")
}

func (s *MessageSuite) TestCopyIntoMessage() {
	msgString := "8=FIX.4.29=17135=D34=249=TW50=KK52=20060102-15:04:0556=ISLD57=AP144=BB115=JCD116=CS128=MG129=CB142=JV143=RY145=BH11=ID21=338=10040=w54=155=INTC60=20060102-15:04:0510=123"
	msgBuf := bytes.NewBufferString(msgString)
	s.Nil(ParseMessage(s.msg, msgBuf))

	dest := NewMessage()
	s.msg.CopyInto(dest)

	checkFieldInt(s, dest.Header.FieldMap, int(tagMsgSeqNum), 2)
	checkFieldInt(s, dest.Body.FieldMap, 21, 3)
	checkFieldString(s, dest.Body.FieldMap, 11, "ID")
	s.Equal(len(dest.bodyBytes), len(s.msg.bodyBytes))

	// copying decouples the message from its input buffer, so the raw message will be re-rendered
	renderedString := "8=FIX.4.29=17135=D34=249=TW50=KK52=20060102-15:04:0556=ISLD57=AP115=JCD116=CS128=MG129=CB142=JV143=RY144=BB145=BH11=ID21=338=10040=w54=155=INTC60=20060102-15:04:0510=033"
	s.Equal(dest.String(), renderedString)

	s.True(reflect.DeepEqual(s.msg.bodyBytes, dest.bodyBytes))
	s.True(s.msg.IsMsgTypeOf("D"))
	s.Equal(s.msg.ReceiveTime, dest.ReceiveTime)

	s.True(reflect.DeepEqual(s.msg.fields, dest.fields))

	// update the source message to validate the copy is truly deep
	newMsgString := "8=FIX.4.49=4935=A52=20140615-19:49:56553=my_user554=secret10=072"
	s.Nil(ParseMessage(s.msg, bytes.NewBufferString(newMsgString)))
	s.True(s.msg.IsMsgTypeOf("A"))
	s.Equal(s.msg.String(), newMsgString)

	// clear the source buffer also
	msgBuf.Reset()

	s.True(dest.IsMsgTypeOf("D"))
	s.Equal(dest.String(), renderedString)
}

func checkFieldInt(s *MessageSuite, fields FieldMap, tag, expected int) {
	toCheck, _ := fields.GetInt(Tag(tag))
	s.Equal(expected, toCheck)
}

func checkFieldString(s *MessageSuite, fields FieldMap, tag int, expected string) {
	toCheck, err := fields.GetString(Tag(tag))
	s.NoError(err)
	s.Equal(expected, toCheck)
}
