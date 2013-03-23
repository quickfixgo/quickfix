package session

import(
  "time"
  . "launchpad.net/gocheck"
  "quickfixgo/fix"
  "quickfixgo/reject"
  "quickfixgo/message/basic"
    )

var _ = Suite(&SessionTests{})
type SessionTests struct {
  session
}

func (s *SessionTests) TestCheckTargetTooHigh(c *C) {
  msg:=basic.NewMessage()
  s.session.expectedSeqNum=45

  //missing seq number
  err:=s.session.checkTargetTooHigh(msg)
  c.Check(err, NotNil)
  c.Check(err.RejectReason(), Equals, reject.RequiredTagMissing)

  //too low
  msg.MsgHeader.Set(basic.NewIntField(fix.MsgSeqNum, 47))
  err=s.session.checkTargetTooHigh(msg)
  c.Check(err, NotNil)
  c.Check(err, FitsTypeOf, reject.TargetTooHigh{})

  //spot on
  msg.MsgHeader.Set(basic.NewIntField(fix.MsgSeqNum, 45))
  err=s.session.checkTargetTooHigh(msg)
  c.Check(err, IsNil)
}

func (s *SessionTests) TestCheckSendingTime(c *C) {
  msg:=basic.NewMessage()

  //missing sending time
  err:=s.session.checkSendingTime(msg)
  c.Check(err, NotNil)
  c.Check(err.RejectReason(), Equals, reject.RequiredTagMissing)

  //sending time too late
  sendingTime:=time.Now().Add(time.Duration(-200)*time.Second)
  msg.MsgHeader.Set(basic.NewUTCTimestampField(fix.SendingTime, sendingTime))
  err=s.session.checkSendingTime(msg)
  c.Check(err, NotNil)
  c.Check(err.RejectReason(), Equals, reject.SendingTimeAccuracyProblem)

  //future sending time
  sendingTime=time.Now().Add(time.Duration(200)*time.Second)
  msg.MsgHeader.Set(basic.NewUTCTimestampField(fix.SendingTime, sendingTime))
  err=s.session.checkSendingTime(msg)
  c.Check(err, NotNil)
  c.Check(err.RejectReason(), Equals, reject.SendingTimeAccuracyProblem)

  //sending time ok
  sendingTime=time.Now()
  msg.MsgHeader.Set(basic.NewUTCTimestampField(fix.SendingTime, sendingTime))
  err=s.session.checkSendingTime(msg)
  c.Check(err, IsNil)
}

func (s *SessionTests) TestCheckTargetTooLow(c *C) {
  msg:=basic.NewMessage()
  s.session.expectedSeqNum=45

  //missing seq number
  err:=s.session.checkTargetTooLow(msg)
  c.Check(err, NotNil)
  c.Check(err.RejectReason(), Equals, reject.RequiredTagMissing)

  //too low
  msg.MsgHeader.Set(basic.NewIntField(fix.MsgSeqNum, 43))
  err=s.session.checkTargetTooLow(msg)
  c.Check(err, NotNil)
  c.Check(err, FitsTypeOf, reject.TargetTooLow{})

  //spot on
  msg.MsgHeader.Set(basic.NewIntField(fix.MsgSeqNum, 45))
  err=s.session.checkTargetTooLow(msg)
  c.Check(err, IsNil)
}




