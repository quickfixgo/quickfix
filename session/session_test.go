package session

import (
	. "launchpad.net/gocheck"
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/message/basic"
	"github.com/cbusbey/quickfixgo/reject"
	"time"
)

var _ = Suite(&SessionTests{})

type SessionTests struct {
	session
}

func (s *SessionTests) TestCheckCorrectCompID(c *C) {
	s.session.ID.TargetCompID = "TAR"
	s.session.ID.SenderCompID = "SND"

	msg := basic.NewMessage()
	//missing target comp id or sender comp id
	err := s.session.checkCompID(msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, reject.RequiredTagMissing)

	msg.MsgHeader.Set(basic.NewStringField(fix.SenderCompID, "TAR"))
	err = s.session.checkCompID(msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, reject.RequiredTagMissing)

	//comp wrong
	msg.MsgHeader.Set(basic.NewStringField(fix.TargetCompID, "JCD"))
	err = s.session.checkCompID(msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, reject.CompIDProblem)

	msg.MsgHeader.Set(basic.NewStringField(fix.TargetCompID, "SND"))
	msg.MsgHeader.Set(basic.NewStringField(fix.SenderCompID, "JCD"))
	err = s.session.checkCompID(msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, reject.CompIDProblem)

	msg.MsgHeader.Set(basic.NewStringField(fix.TargetCompID, "SND"))
	msg.MsgHeader.Set(basic.NewStringField(fix.SenderCompID, "TAR"))
	err = s.session.checkCompID(msg)
	c.Check(err, IsNil)
}

func (s *SessionTests) TestCheckBeginString(c *C) {
	s.session.ID.BeginString = "FIX.4.2"
	msg := basic.NewMessage()

	//missing begin string
	err := s.session.checkBeginString(msg)
	c.Check(err.RejectReason(), Equals, reject.RequiredTagMissing)

	//wrong value
	msg.MsgHeader.Set(basic.NewStringField(fix.BeginString, "FIX.4.4"))
	err = s.session.checkBeginString(msg)
	c.Check(err, NotNil)
	c.Check(err, FitsTypeOf, reject.IncorrectBeginString{})

	msg.MsgHeader.Set(basic.NewStringField(fix.BeginString, s.session.ID.BeginString))
	err = s.session.checkBeginString(msg)
	c.Check(err, IsNil)

}

func (s *SessionTests) TestCheckTargetTooHigh(c *C) {
	msg := basic.NewMessage()
	s.session.expectedSeqNum = 45

	//missing seq number
	err := s.session.checkTargetTooHigh(msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, reject.RequiredTagMissing)

	//too low
	msg.MsgHeader.Set(basic.NewIntField(fix.MsgSeqNum, 47))
	err = s.session.checkTargetTooHigh(msg)
	c.Check(err, NotNil)
	c.Check(err, FitsTypeOf, reject.TargetTooHigh{})

	//spot on
	msg.MsgHeader.Set(basic.NewIntField(fix.MsgSeqNum, 45))
	err = s.session.checkTargetTooHigh(msg)
	c.Check(err, IsNil)
}

func (s *SessionTests) TestCheckSendingTime(c *C) {
	msg := basic.NewMessage()

	//missing sending time
	err := s.session.checkSendingTime(msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, reject.RequiredTagMissing)

	//sending time too late
	sendingTime := time.Now().Add(time.Duration(-200) * time.Second)
	msg.MsgHeader.Set(basic.NewUTCTimestampField(fix.SendingTime, sendingTime))
	err = s.session.checkSendingTime(msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, reject.SendingTimeAccuracyProblem)

	//future sending time
	sendingTime = time.Now().Add(time.Duration(200) * time.Second)
	msg.MsgHeader.Set(basic.NewUTCTimestampField(fix.SendingTime, sendingTime))
	err = s.session.checkSendingTime(msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, reject.SendingTimeAccuracyProblem)

	//sending time ok
	sendingTime = time.Now()
	msg.MsgHeader.Set(basic.NewUTCTimestampField(fix.SendingTime, sendingTime))
	err = s.session.checkSendingTime(msg)
	c.Check(err, IsNil)
}

func (s *SessionTests) TestCheckTargetTooLow(c *C) {
	msg := basic.NewMessage()
	s.session.expectedSeqNum = 45

	//missing seq number
	err := s.session.checkTargetTooLow(msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, reject.RequiredTagMissing)

	//too low
	msg.MsgHeader.Set(basic.NewIntField(fix.MsgSeqNum, 43))
	err = s.session.checkTargetTooLow(msg)
	c.Check(err, NotNil)
	c.Check(err, FitsTypeOf, reject.TargetTooLow{})

	//spot on
	msg.MsgHeader.Set(basic.NewIntField(fix.MsgSeqNum, 45))
	err = s.session.checkTargetTooLow(msg)
	c.Check(err, IsNil)
}
