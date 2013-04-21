package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
	. "launchpad.net/gocheck"
	"time"
)

var _ = Suite(&SessionTests{})

type SessionTests struct {
	session
}

func (s *SessionTests) TestCheckCorrectCompID(c *C) {
	s.session.SessionID.TargetCompID = "TAR"
	s.session.SessionID.SenderCompID = "SND"

	msg := NewMessage()
	//missing target comp id or sender comp id
	err := s.session.checkCompID(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, RequiredTagMissing)

	msg.Header.SetField(NewStringField(tag.SenderCompID, "TAR"))
	err = s.session.checkCompID(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, RequiredTagMissing)

	//comp wrong
	msg.Header.SetField(NewStringField(tag.TargetCompID, "JCD"))
	err = s.session.checkCompID(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, CompIDProblem)

	msg.Header.SetField(NewStringField(tag.TargetCompID, "SND"))
	msg.Header.SetField(NewStringField(tag.SenderCompID, "JCD"))
	err = s.session.checkCompID(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, CompIDProblem)

	msg.Header.SetField(NewStringField(tag.TargetCompID, "SND"))
	msg.Header.SetField(NewStringField(tag.SenderCompID, "TAR"))
	err = s.session.checkCompID(*msg)
	c.Check(err, IsNil)
}

func (s *SessionTests) TestCheckBeginString(c *C) {
	s.session.SessionID.BeginString = "FIX.4.2"
	msg := NewMessage()

	//missing begin string
	err := s.session.checkBeginString(*msg)
	c.Check(err.RejectReason(), Equals, RequiredTagMissing)

	//wrong value
	msg.Header.SetField(NewStringField(tag.BeginString, "FIX.4.4"))
	err = s.session.checkBeginString(*msg)
	c.Check(err, NotNil)
	c.Check(err, FitsTypeOf, IncorrectBeginString{})

	msg.Header.SetField(NewStringField(tag.BeginString, s.session.SessionID.BeginString))
	err = s.session.checkBeginString(*msg)
	c.Check(err, IsNil)

}

func (s *SessionTests) TestCheckTargetTooHigh(c *C) {
	msg := NewMessage()
	s.session.expectedSeqNum = 45

	//missing seq number
	err := s.session.checkTargetTooHigh(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, RequiredTagMissing)

	//too low
	msg.Header.SetField(NewIntField(tag.MsgSeqNum, 47))
	err = s.session.checkTargetTooHigh(*msg)
	c.Check(err, NotNil)
	c.Check(err, FitsTypeOf, TargetTooHigh{})

	//spot on
	msg.Header.SetField(NewIntField(tag.MsgSeqNum, 45))
	err = s.session.checkTargetTooHigh(*msg)
	c.Check(err, IsNil)
}

func (s *SessionTests) TestCheckSendingTime(c *C) {
	msg := NewMessage()

	//missing sending time
	err := s.session.checkSendingTime(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, RequiredTagMissing)

	//sending time too late
	sendingTime := time.Now().Add(time.Duration(-200) * time.Second)
	msg.Header.SetField(NewUTCTimestampField(tag.SendingTime, sendingTime))
	err = s.session.checkSendingTime(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, SendingTimeAccuracyProblem)

	//future sending time
	sendingTime = time.Now().Add(time.Duration(200) * time.Second)
	msg.Header.SetField(NewUTCTimestampField(tag.SendingTime, sendingTime))
	err = s.session.checkSendingTime(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, SendingTimeAccuracyProblem)

	//sending time ok
	sendingTime = time.Now()
	msg.Header.SetField(NewUTCTimestampField(tag.SendingTime, sendingTime))
	err = s.session.checkSendingTime(*msg)
	c.Check(err, IsNil)
}

func (s *SessionTests) TestCheckTargetTooLow(c *C) {
	msg := NewMessage()
	s.session.expectedSeqNum = 45

	//missing seq number
	err := s.session.checkTargetTooLow(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, RequiredTagMissing)

	//too low
	msg.Header.SetField(NewIntField(tag.MsgSeqNum, 43))
	err = s.session.checkTargetTooLow(*msg)
	c.Check(err, NotNil)
	c.Check(err, FitsTypeOf, TargetTooLow{})

	//spot on
	msg.Header.SetField(NewIntField(tag.MsgSeqNum, 45))
	err = s.session.checkTargetTooLow(*msg)
	c.Check(err, IsNil)
}
