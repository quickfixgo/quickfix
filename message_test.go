// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package quickfix

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/quickfixgo/quickfix/datadictionary"
)

func BenchmarkParseMessage(b *testing.B) {
	rawMsg := bytes.NewBufferString("8=FIX.4.29=10435=D34=249=TW52=20140515-19:49:56.65956=ISLD11=10021=140=154=155=TSLA60=00010101-00:00:00.00010=039")

	msg := NewMessage()
	for i := 0; i < b.N; i++ {
		_ = ParseMessage(msg, rawMsg)
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

func TestXMLNonFIX(t *testing.T) {
	rawMsg := bytes.NewBufferString("8=FIX.4.29=37235=n34=25512369=148152=20200522-07:05:33.75649=CME50=G56=OAEAAAN57=TRADE_CAPTURE143=US,IL212=261213=<RTRF>8=FIX.4.29=22535=BZ34=6549369=651852=20200522-07:05:33.74649=CME50=G56=9Q5000N57=DUMMY143=US,IL11=ACP159013113373460=20200522-07:05:33.734533=0893=Y1028=Y1300=991369=99612:325081373=31374=91375=15979=159013113373461769710=167</RTRF>10=245\"")
	msg := NewMessage()
	_ = ParseMessage(msg, rawMsg)

	if !msg.Header.Has(tagXMLData) {
		t.Error("Expected xmldata tag")
	}
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
	// Allow fields out of order, save for validation.
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
	// The onbehalfof/deliverto location id not supported in fix 4.0.
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
	s.Equal(string(s.msg.Bytes()), newMsgString)

	// clear the source buffer also
	msgBuf.Reset()

	s.True(dest.IsMsgTypeOf("D"))
	s.Equal(dest.String(), renderedString)
	s.Equal(string(dest.Bytes()), renderedString)
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
