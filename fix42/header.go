package fix42

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/tag"
)

//Header is the fix42 Header type
type Header struct {
	quickfix.Header
}

//NewHeader returns a new, initialized Header instance
func NewHeader() (h Header) {
	h.Init()
	h.SetBeginString("FIX.4.2")
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
func (h Header) SetMsgType(v string) {
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
func (h Header) SetMessageEncoding(v string) {
	h.Set(field.NewMessageEncoding(v))
}

//SetLastMsgSeqNumProcessed sets LastMsgSeqNumProcessed, Tag 369
func (h Header) SetLastMsgSeqNumProcessed(v int) {
	h.Set(field.NewLastMsgSeqNumProcessed(v))
}

//SetOnBehalfOfSendingTime sets OnBehalfOfSendingTime, Tag 370
func (h Header) SetOnBehalfOfSendingTime(v time.Time) {
	h.Set(field.NewOnBehalfOfSendingTime(v))
}

//GetBeginString gets BeginString, Tag 8
func (h Header) GetBeginString() (f field.BeginStringField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetBodyLength gets BodyLength, Tag 9
func (h Header) GetBodyLength() (f field.BodyLengthField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetMsgSeqNum gets MsgSeqNum, Tag 34
func (h Header) GetMsgSeqNum() (f field.MsgSeqNumField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetMsgType gets MsgType, Tag 35
func (h Header) GetMsgType() (f field.MsgTypeField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetPossDupFlag gets PossDupFlag, Tag 43
func (h Header) GetPossDupFlag() (f field.PossDupFlagField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetSenderCompID gets SenderCompID, Tag 49
func (h Header) GetSenderCompID() (f field.SenderCompIDField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetSenderSubID gets SenderSubID, Tag 50
func (h Header) GetSenderSubID() (f field.SenderSubIDField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetSendingTime gets SendingTime, Tag 52
func (h Header) GetSendingTime() (f field.SendingTimeField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetTargetCompID gets TargetCompID, Tag 56
func (h Header) GetTargetCompID() (f field.TargetCompIDField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetTargetSubID gets TargetSubID, Tag 57
func (h Header) GetTargetSubID() (f field.TargetSubIDField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetSecureDataLen gets SecureDataLen, Tag 90
func (h Header) GetSecureDataLen() (f field.SecureDataLenField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetSecureData gets SecureData, Tag 91
func (h Header) GetSecureData() (f field.SecureDataField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetPossResend gets PossResend, Tag 97
func (h Header) GetPossResend() (f field.PossResendField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetOnBehalfOfCompID gets OnBehalfOfCompID, Tag 115
func (h Header) GetOnBehalfOfCompID() (f field.OnBehalfOfCompIDField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetOnBehalfOfSubID gets OnBehalfOfSubID, Tag 116
func (h Header) GetOnBehalfOfSubID() (f field.OnBehalfOfSubIDField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetOrigSendingTime gets OrigSendingTime, Tag 122
func (h Header) GetOrigSendingTime() (f field.OrigSendingTimeField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetDeliverToCompID gets DeliverToCompID, Tag 128
func (h Header) GetDeliverToCompID() (f field.DeliverToCompIDField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetDeliverToSubID gets DeliverToSubID, Tag 129
func (h Header) GetDeliverToSubID() (f field.DeliverToSubIDField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetSenderLocationID gets SenderLocationID, Tag 142
func (h Header) GetSenderLocationID() (f field.SenderLocationIDField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetTargetLocationID gets TargetLocationID, Tag 143
func (h Header) GetTargetLocationID() (f field.TargetLocationIDField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetOnBehalfOfLocationID gets OnBehalfOfLocationID, Tag 144
func (h Header) GetOnBehalfOfLocationID() (f field.OnBehalfOfLocationIDField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetDeliverToLocationID gets DeliverToLocationID, Tag 145
func (h Header) GetDeliverToLocationID() (f field.DeliverToLocationIDField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetXmlDataLen gets XmlDataLen, Tag 212
func (h Header) GetXmlDataLen() (f field.XmlDataLenField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetXmlData gets XmlData, Tag 213
func (h Header) GetXmlData() (f field.XmlDataField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetMessageEncoding gets MessageEncoding, Tag 347
func (h Header) GetMessageEncoding() (f field.MessageEncodingField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetLastMsgSeqNumProcessed gets LastMsgSeqNumProcessed, Tag 369
func (h Header) GetLastMsgSeqNumProcessed() (f field.LastMsgSeqNumProcessedField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
	return
}

//GetOnBehalfOfSendingTime gets OnBehalfOfSendingTime, Tag 370
func (h Header) GetOnBehalfOfSendingTime() (f field.OnBehalfOfSendingTimeField, err quickfix.MessageRejectError) {
	err = h.Get(&f)
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

//HasOnBehalfOfSendingTime returns true if OnBehalfOfSendingTime is present, Tag 370
func (h Header) HasOnBehalfOfSendingTime() bool {
	return h.Has(tag.OnBehalfOfSendingTime)
}
