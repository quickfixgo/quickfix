package quickfix

import (
	"bytes"
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

	rebuildBytes := s.msg.build()

	expectedBytes := []byte("8=FIX.4.29=12635=D34=249=TW52=20140615-19:49:5656=ISLD122=20140515-19:49:56.65911=10021=140=154=155=TSLA60=00010101-00:00:00.00010=128")

	s.True(bytes.Equal(expectedBytes, rebuildBytes), "Unexpected bytes,\n +%s\n-%s", rebuildBytes, expectedBytes)

	expectedBodyBytes := []byte("11=10021=140=154=155=TSLA60=00010101-00:00:00.000")

	s.True(bytes.Equal(s.msg.bodyBytes, expectedBodyBytes), "Incorrect body bytes, got %s", string(s.msg.bodyBytes))
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
