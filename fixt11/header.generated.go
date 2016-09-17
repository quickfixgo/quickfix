package fixt11

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/tag"
)

//Header is the fixt11 Header type
type Header struct {
	quickfix.Header
}

//NewHeader returns a new, initialized Header instance
func NewHeader() (h Header) {
	h.Init()
	h.SetBeginString("FIXT.1.1")
	return
}

//SetBeginString sets BeginString, Tag 8
func (h Header) SetBeginString(v string) {
	h.Set(field.NewBeginString(v))
}

//SetBodyLength sets BodyLength, Tag 9
func (h Header) SetBodyLength(v int) {
	h.Set(field.NewBodyLength(v))
}

//SetMsgSeqNum sets MsgSeqNum, Tag 34
func (h Header) SetMsgSeqNum(v int) {
	h.Set(field.NewMsgSeqNum(v))
}

//SetMsgType sets MsgType, Tag 35
func (h Header) SetMsgType(v enum.MsgType) {
	h.Set(field.NewMsgType(v))
}

//SetPossDupFlag sets PossDupFlag, Tag 43
func (h Header) SetPossDupFlag(v bool) {
	h.Set(field.NewPossDupFlag(v))
}

//SetSenderCompID sets SenderCompID, Tag 49
func (h Header) SetSenderCompID(v string) {
	h.Set(field.NewSenderCompID(v))
}

//SetSenderSubID sets SenderSubID, Tag 50
func (h Header) SetSenderSubID(v string) {
	h.Set(field.NewSenderSubID(v))
}

//SetSendingTime sets SendingTime, Tag 52
func (h Header) SetSendingTime(v time.Time) {
	h.Set(field.NewSendingTime(v))
}

//SetTargetCompID sets TargetCompID, Tag 56
func (h Header) SetTargetCompID(v string) {
	h.Set(field.NewTargetCompID(v))
}

//SetTargetSubID sets TargetSubID, Tag 57
func (h Header) SetTargetSubID(v string) {
	h.Set(field.NewTargetSubID(v))
}

//SetSecureDataLen sets SecureDataLen, Tag 90
func (h Header) SetSecureDataLen(v int) {
	h.Set(field.NewSecureDataLen(v))
}

//SetSecureData sets SecureData, Tag 91
func (h Header) SetSecureData(v string) {
	h.Set(field.NewSecureData(v))
}

//SetPossResend sets PossResend, Tag 97
func (h Header) SetPossResend(v bool) {
	h.Set(field.NewPossResend(v))
}

//SetOnBehalfOfCompID sets OnBehalfOfCompID, Tag 115
func (h Header) SetOnBehalfOfCompID(v string) {
	h.Set(field.NewOnBehalfOfCompID(v))
}

//SetOnBehalfOfSubID sets OnBehalfOfSubID, Tag 116
func (h Header) SetOnBehalfOfSubID(v string) {
	h.Set(field.NewOnBehalfOfSubID(v))
}

//SetOrigSendingTime sets OrigSendingTime, Tag 122
func (h Header) SetOrigSendingTime(v time.Time) {
	h.Set(field.NewOrigSendingTime(v))
}

//SetDeliverToCompID sets DeliverToCompID, Tag 128
func (h Header) SetDeliverToCompID(v string) {
	h.Set(field.NewDeliverToCompID(v))
}

//SetDeliverToSubID sets DeliverToSubID, Tag 129
func (h Header) SetDeliverToSubID(v string) {
	h.Set(field.NewDeliverToSubID(v))
}

//SetSenderLocationID sets SenderLocationID, Tag 142
func (h Header) SetSenderLocationID(v string) {
	h.Set(field.NewSenderLocationID(v))
}

//SetTargetLocationID sets TargetLocationID, Tag 143
func (h Header) SetTargetLocationID(v string) {
	h.Set(field.NewTargetLocationID(v))
}

//SetOnBehalfOfLocationID sets OnBehalfOfLocationID, Tag 144
func (h Header) SetOnBehalfOfLocationID(v string) {
	h.Set(field.NewOnBehalfOfLocationID(v))
}

//SetDeliverToLocationID sets DeliverToLocationID, Tag 145
func (h Header) SetDeliverToLocationID(v string) {
	h.Set(field.NewDeliverToLocationID(v))
}

//SetXmlDataLen sets XmlDataLen, Tag 212
func (h Header) SetXmlDataLen(v int) {
	h.Set(field.NewXmlDataLen(v))
}

//SetXmlData sets XmlData, Tag 213
func (h Header) SetXmlData(v string) {
	h.Set(field.NewXmlData(v))
}

//SetMessageEncoding sets MessageEncoding, Tag 347
func (h Header) SetMessageEncoding(v enum.MessageEncoding) {
	h.Set(field.NewMessageEncoding(v))
}

//SetLastMsgSeqNumProcessed sets LastMsgSeqNumProcessed, Tag 369
func (h Header) SetLastMsgSeqNumProcessed(v int) {
	h.Set(field.NewLastMsgSeqNumProcessed(v))
}

//SetNoHops sets NoHops, Tag 627
func (h Header) SetNoHops(f NoHopsRepeatingGroup) {
	h.SetGroup(f)
}

//SetApplVerID sets ApplVerID, Tag 1128
func (h Header) SetApplVerID(v enum.ApplVerID) {
	h.Set(field.NewApplVerID(v))
}

//SetCstmApplVerID sets CstmApplVerID, Tag 1129
func (h Header) SetCstmApplVerID(v string) {
	h.Set(field.NewCstmApplVerID(v))
}

//GetBeginString gets BeginString, Tag 8
func (h Header) GetBeginString() (v string, err quickfix.MessageRejectError) {
	var f field.BeginStringField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBodyLength gets BodyLength, Tag 9
func (h Header) GetBodyLength() (v int, err quickfix.MessageRejectError) {
	var f field.BodyLengthField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMsgSeqNum gets MsgSeqNum, Tag 34
func (h Header) GetMsgSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.MsgSeqNumField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMsgType gets MsgType, Tag 35
func (h Header) GetMsgType() (v enum.MsgType, err quickfix.MessageRejectError) {
	var f field.MsgTypeField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPossDupFlag gets PossDupFlag, Tag 43
func (h Header) GetPossDupFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.PossDupFlagField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSenderCompID gets SenderCompID, Tag 49
func (h Header) GetSenderCompID() (v string, err quickfix.MessageRejectError) {
	var f field.SenderCompIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSenderSubID gets SenderSubID, Tag 50
func (h Header) GetSenderSubID() (v string, err quickfix.MessageRejectError) {
	var f field.SenderSubIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSendingTime gets SendingTime, Tag 52
func (h Header) GetSendingTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.SendingTimeField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTargetCompID gets TargetCompID, Tag 56
func (h Header) GetTargetCompID() (v string, err quickfix.MessageRejectError) {
	var f field.TargetCompIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTargetSubID gets TargetSubID, Tag 57
func (h Header) GetTargetSubID() (v string, err quickfix.MessageRejectError) {
	var f field.TargetSubIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecureDataLen gets SecureDataLen, Tag 90
func (h Header) GetSecureDataLen() (v int, err quickfix.MessageRejectError) {
	var f field.SecureDataLenField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecureData gets SecureData, Tag 91
func (h Header) GetSecureData() (v string, err quickfix.MessageRejectError) {
	var f field.SecureDataField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPossResend gets PossResend, Tag 97
func (h Header) GetPossResend() (v bool, err quickfix.MessageRejectError) {
	var f field.PossResendField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOnBehalfOfCompID gets OnBehalfOfCompID, Tag 115
func (h Header) GetOnBehalfOfCompID() (v string, err quickfix.MessageRejectError) {
	var f field.OnBehalfOfCompIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOnBehalfOfSubID gets OnBehalfOfSubID, Tag 116
func (h Header) GetOnBehalfOfSubID() (v string, err quickfix.MessageRejectError) {
	var f field.OnBehalfOfSubIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrigSendingTime gets OrigSendingTime, Tag 122
func (h Header) GetOrigSendingTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.OrigSendingTimeField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDeliverToCompID gets DeliverToCompID, Tag 128
func (h Header) GetDeliverToCompID() (v string, err quickfix.MessageRejectError) {
	var f field.DeliverToCompIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDeliverToSubID gets DeliverToSubID, Tag 129
func (h Header) GetDeliverToSubID() (v string, err quickfix.MessageRejectError) {
	var f field.DeliverToSubIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSenderLocationID gets SenderLocationID, Tag 142
func (h Header) GetSenderLocationID() (v string, err quickfix.MessageRejectError) {
	var f field.SenderLocationIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTargetLocationID gets TargetLocationID, Tag 143
func (h Header) GetTargetLocationID() (v string, err quickfix.MessageRejectError) {
	var f field.TargetLocationIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOnBehalfOfLocationID gets OnBehalfOfLocationID, Tag 144
func (h Header) GetOnBehalfOfLocationID() (v string, err quickfix.MessageRejectError) {
	var f field.OnBehalfOfLocationIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDeliverToLocationID gets DeliverToLocationID, Tag 145
func (h Header) GetDeliverToLocationID() (v string, err quickfix.MessageRejectError) {
	var f field.DeliverToLocationIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetXmlDataLen gets XmlDataLen, Tag 212
func (h Header) GetXmlDataLen() (v int, err quickfix.MessageRejectError) {
	var f field.XmlDataLenField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetXmlData gets XmlData, Tag 213
func (h Header) GetXmlData() (v string, err quickfix.MessageRejectError) {
	var f field.XmlDataField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMessageEncoding gets MessageEncoding, Tag 347
func (h Header) GetMessageEncoding() (v enum.MessageEncoding, err quickfix.MessageRejectError) {
	var f field.MessageEncodingField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastMsgSeqNumProcessed gets LastMsgSeqNumProcessed, Tag 369
func (h Header) GetLastMsgSeqNumProcessed() (v int, err quickfix.MessageRejectError) {
	var f field.LastMsgSeqNumProcessedField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoHops gets NoHops, Tag 627
func (h Header) GetNoHops() (NoHopsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoHopsRepeatingGroup()
	err := h.GetGroup(f)
	return f, err
}

//GetApplVerID gets ApplVerID, Tag 1128
func (h Header) GetApplVerID() (v enum.ApplVerID, err quickfix.MessageRejectError) {
	var f field.ApplVerIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCstmApplVerID gets CstmApplVerID, Tag 1129
func (h Header) GetCstmApplVerID() (v string, err quickfix.MessageRejectError) {
	var f field.CstmApplVerIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasBeginString returns true if BeginString is present, Tag 8
func (h Header) HasBeginString() bool {
	return h.Has(tag.BeginString)
}

//HasBodyLength returns true if BodyLength is present, Tag 9
func (h Header) HasBodyLength() bool {
	return h.Has(tag.BodyLength)
}

//HasMsgSeqNum returns true if MsgSeqNum is present, Tag 34
func (h Header) HasMsgSeqNum() bool {
	return h.Has(tag.MsgSeqNum)
}

//HasMsgType returns true if MsgType is present, Tag 35
func (h Header) HasMsgType() bool {
	return h.Has(tag.MsgType)
}

//HasPossDupFlag returns true if PossDupFlag is present, Tag 43
func (h Header) HasPossDupFlag() bool {
	return h.Has(tag.PossDupFlag)
}

//HasSenderCompID returns true if SenderCompID is present, Tag 49
func (h Header) HasSenderCompID() bool {
	return h.Has(tag.SenderCompID)
}

//HasSenderSubID returns true if SenderSubID is present, Tag 50
func (h Header) HasSenderSubID() bool {
	return h.Has(tag.SenderSubID)
}

//HasSendingTime returns true if SendingTime is present, Tag 52
func (h Header) HasSendingTime() bool {
	return h.Has(tag.SendingTime)
}

//HasTargetCompID returns true if TargetCompID is present, Tag 56
func (h Header) HasTargetCompID() bool {
	return h.Has(tag.TargetCompID)
}

//HasTargetSubID returns true if TargetSubID is present, Tag 57
func (h Header) HasTargetSubID() bool {
	return h.Has(tag.TargetSubID)
}

//HasSecureDataLen returns true if SecureDataLen is present, Tag 90
func (h Header) HasSecureDataLen() bool {
	return h.Has(tag.SecureDataLen)
}

//HasSecureData returns true if SecureData is present, Tag 91
func (h Header) HasSecureData() bool {
	return h.Has(tag.SecureData)
}

//HasPossResend returns true if PossResend is present, Tag 97
func (h Header) HasPossResend() bool {
	return h.Has(tag.PossResend)
}

//HasOnBehalfOfCompID returns true if OnBehalfOfCompID is present, Tag 115
func (h Header) HasOnBehalfOfCompID() bool {
	return h.Has(tag.OnBehalfOfCompID)
}

//HasOnBehalfOfSubID returns true if OnBehalfOfSubID is present, Tag 116
func (h Header) HasOnBehalfOfSubID() bool {
	return h.Has(tag.OnBehalfOfSubID)
}

//HasOrigSendingTime returns true if OrigSendingTime is present, Tag 122
func (h Header) HasOrigSendingTime() bool {
	return h.Has(tag.OrigSendingTime)
}

//HasDeliverToCompID returns true if DeliverToCompID is present, Tag 128
func (h Header) HasDeliverToCompID() bool {
	return h.Has(tag.DeliverToCompID)
}

//HasDeliverToSubID returns true if DeliverToSubID is present, Tag 129
func (h Header) HasDeliverToSubID() bool {
	return h.Has(tag.DeliverToSubID)
}

//HasSenderLocationID returns true if SenderLocationID is present, Tag 142
func (h Header) HasSenderLocationID() bool {
	return h.Has(tag.SenderLocationID)
}

//HasTargetLocationID returns true if TargetLocationID is present, Tag 143
func (h Header) HasTargetLocationID() bool {
	return h.Has(tag.TargetLocationID)
}

//HasOnBehalfOfLocationID returns true if OnBehalfOfLocationID is present, Tag 144
func (h Header) HasOnBehalfOfLocationID() bool {
	return h.Has(tag.OnBehalfOfLocationID)
}

//HasDeliverToLocationID returns true if DeliverToLocationID is present, Tag 145
func (h Header) HasDeliverToLocationID() bool {
	return h.Has(tag.DeliverToLocationID)
}

//HasXmlDataLen returns true if XmlDataLen is present, Tag 212
func (h Header) HasXmlDataLen() bool {
	return h.Has(tag.XmlDataLen)
}

//HasXmlData returns true if XmlData is present, Tag 213
func (h Header) HasXmlData() bool {
	return h.Has(tag.XmlData)
}

//HasMessageEncoding returns true if MessageEncoding is present, Tag 347
func (h Header) HasMessageEncoding() bool {
	return h.Has(tag.MessageEncoding)
}

//HasLastMsgSeqNumProcessed returns true if LastMsgSeqNumProcessed is present, Tag 369
func (h Header) HasLastMsgSeqNumProcessed() bool {
	return h.Has(tag.LastMsgSeqNumProcessed)
}

//HasNoHops returns true if NoHops is present, Tag 627
func (h Header) HasNoHops() bool {
	return h.Has(tag.NoHops)
}

//HasApplVerID returns true if ApplVerID is present, Tag 1128
func (h Header) HasApplVerID() bool {
	return h.Has(tag.ApplVerID)
}

//HasCstmApplVerID returns true if CstmApplVerID is present, Tag 1129
func (h Header) HasCstmApplVerID() bool {
	return h.Has(tag.CstmApplVerID)
}

//NoHops is a repeating group element, Tag 627
type NoHops struct {
	quickfix.Group
}

//SetHopCompID sets HopCompID, Tag 628
func (h NoHops) SetHopCompID(v string) {
	h.Set(field.NewHopCompID(v))
}

//SetHopSendingTime sets HopSendingTime, Tag 629
func (h NoHops) SetHopSendingTime(v time.Time) {
	h.Set(field.NewHopSendingTime(v))
}

//SetHopRefID sets HopRefID, Tag 630
func (h NoHops) SetHopRefID(v int) {
	h.Set(field.NewHopRefID(v))
}

//GetHopCompID gets HopCompID, Tag 628
func (h NoHops) GetHopCompID() (v string, err quickfix.MessageRejectError) {
	var f field.HopCompIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetHopSendingTime gets HopSendingTime, Tag 629
func (h NoHops) GetHopSendingTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.HopSendingTimeField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetHopRefID gets HopRefID, Tag 630
func (h NoHops) GetHopRefID() (v int, err quickfix.MessageRejectError) {
	var f field.HopRefIDField
	if err = h.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasHopCompID returns true if HopCompID is present, Tag 628
func (h NoHops) HasHopCompID() bool {
	return h.Has(tag.HopCompID)
}

//HasHopSendingTime returns true if HopSendingTime is present, Tag 629
func (h NoHops) HasHopSendingTime() bool {
	return h.Has(tag.HopSendingTime)
}

//HasHopRefID returns true if HopRefID is present, Tag 630
func (h NoHops) HasHopRefID() bool {
	return h.Has(tag.HopRefID)
}

//NoHopsRepeatingGroup is a repeating group, Tag 627
type NoHopsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoHopsRepeatingGroup returns an initialized, NoHopsRepeatingGroup
func NewNoHopsRepeatingGroup() NoHopsRepeatingGroup {
	return NoHopsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoHops,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.HopCompID), quickfix.GroupElement(tag.HopSendingTime), quickfix.GroupElement(tag.HopRefID)})}
}

//Add create and append a new NoHops to this group
func (h NoHopsRepeatingGroup) Add() NoHops {
	g := h.RepeatingGroup.Add()
	return NoHops{g}
}

//Get returns the ith NoHops in the NoHopsRepeatinGroup
func (h NoHopsRepeatingGroup) Get(i int) NoHops {
	return NoHops{h.RepeatingGroup.Get(i)}
}
