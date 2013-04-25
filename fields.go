package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
)

//TODO: Generate me

type BeginString struct{ StringValue }

func (f BeginString) Tag() tag.Tag { return tag.BeginString }

type SenderCompID struct{ StringValue }

func (f SenderCompID) Tag() tag.Tag { return tag.SenderCompID }

type TargetCompID struct{ StringValue }

func (f TargetCompID) Tag() tag.Tag { return tag.TargetCompID }

type DefaultApplVerID struct{ StringValue }

func (f DefaultApplVerID) Tag() tag.Tag { return tag.DefaultApplVerID }

type ApplVerID struct{ StringValue }

func (f ApplVerID) Tag() tag.Tag { return tag.ApplVerID }

type MsgType struct{ StringValue }

func (f MsgType) Tag() tag.Tag { return tag.MsgType }

//Returns true if the message type is a session level type
func (m MsgType) IsAdminMessageType() bool {
	switch m.Value {
	case "0", "A", "1", "2", "3", "4", "5":
		return true
	}

	return false
}

type TestReqID struct{ StringValue }

func (f TestReqID) Tag() tag.Tag { return tag.TestReqID }

type GapFillFlag struct{ BooleanValue }

func (f GapFillFlag) Tag() tag.Tag { return tag.GapFillFlag }

type PossDupFlag struct{ BooleanValue }

func (f PossDupFlag) Tag() tag.Tag { return tag.PossDupFlag }

type PossResend struct{ BooleanValue }

func (f PossResend) Tag() tag.Tag { return tag.PossResend }

type BodyLength struct{ IntValue }

func (f BodyLength) Tag() tag.Tag { return tag.BodyLength }

type NewSeqNo struct{ IntValue }

func (f NewSeqNo) Tag() tag.Tag { return tag.NewSeqNo }

type HeartBtInt struct{ IntValue }

func (f HeartBtInt) Tag() tag.Tag { return tag.HeartBtInt }

type ClOrdID struct{ StringValue }

func (f ClOrdID) Tag() tag.Tag { return tag.ClOrdID }

type MsgSeqNum IntField

func (f *MsgSeqNum) Tag() tag.Tag { return tag.MsgSeqNum }

type BeginSeqNo IntField

func (f *BeginSeqNo) Tag() tag.Tag { return tag.BeginSeqNo }

type EndSeqNo IntField

func (f *EndSeqNo) Tag() tag.Tag { return tag.EndSeqNo }

type SendingTime struct{ UTCTimestampValue }

func (f SendingTime) Tag() tag.Tag { return tag.SendingTime }

type OrigSendingTime struct{ UTCTimestampValue }

func (f OrigSendingTime) Tag() tag.Tag { return tag.OrigSendingTime }
