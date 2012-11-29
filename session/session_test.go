package session

import(
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



