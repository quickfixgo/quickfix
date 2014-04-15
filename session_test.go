package quickfix

import (
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/tag"
	"github.com/quickfixgo/quickfix/message"
	. "launchpad.net/gocheck"
	"time"
)

var _ = Suite(&SessionTests{})

type SessionTests struct {
	session
}

func getBuilder() message.MessageBuilder {
	builder := message.CreateMessageBuilder()
	builder.Header.Set(message.NewStringField(tag.BeginString, fix.BeginString_FIX40))
	builder.Header.Set(message.NewStringField(tag.MsgType, "D"))
	return builder
}

func (s *SessionTests) TestCheckCorrectCompID(c *C) {
	s.session.SessionID.TargetCompID = "TAR"
	s.session.SessionID.SenderCompID = "SND"

	builder := getBuilder()
	msg, _ := builder.Build()
	//missing target comp id or sender comp id
	err := s.session.checkCompID(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, message.RequiredTagMissing)

	builder.Header.Set(message.NewStringField(tag.SenderCompID, "TAR"))
	msg, _ = builder.Build()
	err = s.session.checkCompID(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, message.RequiredTagMissing)

	//comp wrong
	builder.Header.Set(message.NewStringField(tag.TargetCompID, "JCD"))
	msg, _ = builder.Build()
	err = s.session.checkCompID(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, message.CompIDProblem)

	builder.Header.Set(message.NewStringField(tag.TargetCompID, "SND"))
	builder.Header.Set(message.NewStringField(tag.SenderCompID, "JCD"))
	msg, _ = builder.Build()
	err = s.session.checkCompID(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, message.CompIDProblem)

	builder.Header.Set(message.NewStringField(tag.TargetCompID, "SND"))
	builder.Header.Set(message.NewStringField(tag.SenderCompID, "TAR"))
	msg, _ = builder.Build()
	err = s.session.checkCompID(*msg)
	c.Check(err, IsNil)
}

func (s *SessionTests) TestCheckBeginString(c *C) {
	s.session.SessionID.BeginString = "FIX.4.2"
	builder := getBuilder()

	//wrong value
	builder.Header.Set(message.NewStringField(tag.BeginString, "FIX.4.4"))
	msg, _ := builder.Build()
	err := s.session.checkBeginString(*msg)
	c.Check(err, NotNil)
	c.Check(err, FitsTypeOf, message.IncorrectBeginString{})

	builder.Header.Set(message.NewStringField(tag.BeginString, s.session.SessionID.BeginString))
	msg, _ = builder.Build()
	err = s.session.checkBeginString(*msg)
	c.Check(err, IsNil)

}

func (s *SessionTests) TestCheckTargetTooHigh(c *C) {
	builder := getBuilder()
	msg, _ := builder.Build()
	s.session.expectedSeqNum = 45

	//missing seq number
	err := s.session.checkTargetTooHigh(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, message.RequiredTagMissing)

	//too low
	builder.Header.Set(message.NewIntField(tag.MsgSeqNum, 47))
	msg, _ = builder.Build()
	err = s.session.checkTargetTooHigh(*msg)
	c.Check(err, NotNil)
	c.Check(err, FitsTypeOf, message.TargetTooHigh{})

	//spot on
	builder.Header.Set(message.NewIntField(tag.MsgSeqNum, 45))
	msg, _ = builder.Build()
	err = s.session.checkTargetTooHigh(*msg)
	c.Check(err, IsNil)
}

func (s *SessionTests) TestCheckSendingTime(c *C) {
	builder := getBuilder()
	msg, _ := builder.Build()

	//missing sending time
	err := s.session.checkSendingTime(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, message.RequiredTagMissing)

	//sending time too late
	sendingTime := time.Now().Add(time.Duration(-200) * time.Second)
	builder.Header.Set(message.NewUTCTimestampField(tag.SendingTime, sendingTime))
	msg, _ = builder.Build()
	err = s.session.checkSendingTime(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, message.SendingTimeAccuracyProblem)

	//future sending time
	sendingTime = time.Now().Add(time.Duration(200) * time.Second)
	builder.Header.Set(message.NewUTCTimestampField(tag.SendingTime, sendingTime))
	msg, _ = builder.Build()
	err = s.session.checkSendingTime(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, message.SendingTimeAccuracyProblem)

	//sending time ok
	sendingTime = time.Now()
	builder.Header.Set(message.NewUTCTimestampField(tag.SendingTime, sendingTime))
	msg, _ = builder.Build()
	err = s.session.checkSendingTime(*msg)
	c.Check(err, IsNil)
}

func (s *SessionTests) TestCheckTargetTooLow(c *C) {
	builder := getBuilder()
	msg, _ := builder.Build()
	s.session.expectedSeqNum = 45

	//missing seq number
	err := s.session.checkTargetTooLow(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, message.RequiredTagMissing)

	//too low
	builder.Header.Set(message.NewIntField(tag.MsgSeqNum, 43))
	msg, _ = builder.Build()
	err = s.session.checkTargetTooLow(*msg)
	c.Check(err, NotNil)
	c.Check(err, FitsTypeOf, message.TargetTooLow{})

	//spot on
	builder.Header.Set(message.NewIntField(tag.MsgSeqNum, 45))
	msg, _ = builder.Build()
	err = s.session.checkTargetTooLow(*msg)
	c.Check(err, IsNil)
}
