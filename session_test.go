package quickfix

import (
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/fix/tag"
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

var _ = Suite(&SessionTests{})

type SessionTests struct {
	session Session
}

func getBuilder() MessageBuilder {
	builder := NewMessageBuilder()
	builder.Header().Set(fix.NewStringField(tag.BeginString, fix.BeginString_FIX40))
	builder.Header().Set(fix.NewStringField(tag.MsgType, "D"))
	return builder
}

func TestSession_CheckCorrectCompID(t *testing.T) {
	session := Session{}
	session.sessionID.TargetCompID = "TAR"
	session.sessionID.SenderCompID = "SND"

	var testCases = []struct {
		senderCompID *field.SenderCompIDField
		targetCompID *field.TargetCompIDField
		returnsError bool
		rejectReason RejectReason
	}{
		{returnsError: true, rejectReason: RejectReasonRequiredTagMissing},
		{senderCompID: field.NewSenderCompID("TAR"),
			returnsError: true,
			rejectReason: RejectReasonRequiredTagMissing},
		{senderCompID: field.NewSenderCompID("TAR"),
			targetCompID: field.NewTargetCompID("JCD"),
			returnsError: true,
			rejectReason: RejectReasonCompIDProblem},
		{senderCompID: field.NewSenderCompID("JCD"),
			targetCompID: field.NewTargetCompID("SND"),
			returnsError: true,
			rejectReason: RejectReasonCompIDProblem},
		{senderCompID: field.NewSenderCompID("TAR"),
			targetCompID: field.NewTargetCompID("SND"),
			returnsError: false},
	}

	for _, tc := range testCases {
		builder := getBuilder()

		if tc.senderCompID != nil {
			builder.Header().Set(tc.senderCompID)
		}

		if tc.targetCompID != nil {
			builder.Header().Set(tc.targetCompID)
		}

		msg, _ := builder.Build()
		err := session.checkCompID(*msg)

		if err == nil {
			if tc.returnsError {
				t.Error("expected error")
			}

			return
		}

		if !tc.returnsError {
			t.Fatal("unexpected error", err)
		}

		if err.RejectReason() != tc.rejectReason {
			t.Errorf("expected %v got %v", tc.rejectReason, err.RejectReason())
		}
	}
}

func (s *SessionTests) TestCheckBeginString(c *C) {
	s.session.sessionID.BeginString = "FIX.4.2"
	builder := getBuilder()

	//wrong value
	builder.Header().Set(fix.NewStringField(tag.BeginString, "FIX.4.4"))
	msg, _ := builder.Build()
	err := s.session.checkBeginString(*msg)
	c.Check(err, NotNil)
	c.Check(err, FitsTypeOf, incorrectBeginString{})

	builder.Header().Set(fix.NewStringField(tag.BeginString, s.session.sessionID.BeginString))
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
	c.Check(err.RejectReason(), Equals, RejectReasonRequiredTagMissing)

	//too low
	builder.Header().Set(fix.NewIntField(tag.MsgSeqNum, 47))
	msg, _ = builder.Build()
	err = s.session.checkTargetTooHigh(*msg)
	c.Check(err, NotNil)
	c.Check(err, FitsTypeOf, targetTooHigh{})

	//spot on
	builder.Header().Set(fix.NewIntField(tag.MsgSeqNum, 45))
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
	c.Check(err.RejectReason(), Equals, RejectReasonRequiredTagMissing)

	//sending time too late
	sendingTime := time.Now().Add(time.Duration(-200) * time.Second)
	builder.Header().Set(fix.NewUTCTimestampField(tag.SendingTime, sendingTime))
	msg, _ = builder.Build()
	err = s.session.checkSendingTime(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, RejectReasonSendingTimeAccuracyProblem)

	//future sending time
	sendingTime = time.Now().Add(time.Duration(200) * time.Second)
	builder.Header().Set(fix.NewUTCTimestampField(tag.SendingTime, sendingTime))
	msg, _ = builder.Build()
	err = s.session.checkSendingTime(*msg)
	c.Check(err, NotNil)
	c.Check(err.RejectReason(), Equals, RejectReasonSendingTimeAccuracyProblem)

	//sending time ok
	sendingTime = time.Now()
	builder.Header().Set(fix.NewUTCTimestampField(tag.SendingTime, sendingTime))
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
	c.Check(err.RejectReason(), Equals, RejectReasonRequiredTagMissing)

	//too low
	builder.Header().Set(fix.NewIntField(tag.MsgSeqNum, 43))
	msg, _ = builder.Build()
	err = s.session.checkTargetTooLow(*msg)
	c.Check(err, NotNil)
	c.Check(err, FitsTypeOf, targetTooLow{})

	//spot on
	builder.Header().Set(fix.NewIntField(tag.MsgSeqNum, 45))
	msg, _ = builder.Build()
	err = s.session.checkTargetTooLow(*msg)
	c.Check(err, IsNil)
}
