package fixt11

import (
	"github.com/quickfixgo/quickfix/fixt11/hopgrp"
	"time"
)

//Header is the fixt11 Header type
type Header struct {
	//BeginString is a required field for Header.
	BeginString string `fix:"8,default=FIXT.1.1"`
	//BodyLength is a required field for Header.
	BodyLength int `fix:"9"`
	//MsgType is a required field for Header.
	MsgType string `fix:"35"`
	//SenderCompID is a required field for Header.
	SenderCompID string `fix:"49"`
	//TargetCompID is a required field for Header.
	TargetCompID string `fix:"56"`
	//OnBehalfOfCompID is a non-required field for Header.
	OnBehalfOfCompID *string `fix:"115"`
	//DeliverToCompID is a non-required field for Header.
	DeliverToCompID *string `fix:"128"`
	//SecureDataLen is a non-required field for Header.
	SecureDataLen *int `fix:"90"`
	//SecureData is a non-required field for Header.
	SecureData *string `fix:"91"`
	//MsgSeqNum is a required field for Header.
	MsgSeqNum int `fix:"34"`
	//SenderSubID is a non-required field for Header.
	SenderSubID *string `fix:"50"`
	//SenderLocationID is a non-required field for Header.
	SenderLocationID *string `fix:"142"`
	//TargetSubID is a non-required field for Header.
	TargetSubID *string `fix:"57"`
	//TargetLocationID is a non-required field for Header.
	TargetLocationID *string `fix:"143"`
	//OnBehalfOfSubID is a non-required field for Header.
	OnBehalfOfSubID *string `fix:"116"`
	//OnBehalfOfLocationID is a non-required field for Header.
	OnBehalfOfLocationID *string `fix:"144"`
	//DeliverToSubID is a non-required field for Header.
	DeliverToSubID *string `fix:"129"`
	//DeliverToLocationID is a non-required field for Header.
	DeliverToLocationID *string `fix:"145"`
	//PossDupFlag is a non-required field for Header.
	PossDupFlag *bool `fix:"43"`
	//PossResend is a non-required field for Header.
	PossResend *bool `fix:"97"`
	//SendingTime is a required field for Header.
	SendingTime time.Time `fix:"52"`
	//OrigSendingTime is a non-required field for Header.
	OrigSendingTime *time.Time `fix:"122"`
	//XmlDataLen is a non-required field for Header.
	XmlDataLen *int `fix:"212"`
	//XmlData is a non-required field for Header.
	XmlData *string `fix:"213"`
	//MessageEncoding is a non-required field for Header.
	MessageEncoding *string `fix:"347"`
	//LastMsgSeqNumProcessed is a non-required field for Header.
	LastMsgSeqNumProcessed *int `fix:"369"`
	//HopGrp Component
	HopGrp hopgrp.Component
	//ApplVerID is a non-required field for Header.
	ApplVerID *string `fix:"1128"`
	//CstmApplVerID is a non-required field for Header.
	CstmApplVerID *string `fix:"1129"`
}

func (m *Header) SetBeginString(v string)          { m.BeginString = v }
func (m *Header) SetBodyLength(v int)              { m.BodyLength = v }
func (m *Header) SetMsgType(v string)              { m.MsgType = v }
func (m *Header) SetSenderCompID(v string)         { m.SenderCompID = v }
func (m *Header) SetTargetCompID(v string)         { m.TargetCompID = v }
func (m *Header) SetOnBehalfOfCompID(v string)     { m.OnBehalfOfCompID = &v }
func (m *Header) SetDeliverToCompID(v string)      { m.DeliverToCompID = &v }
func (m *Header) SetSecureDataLen(v int)           { m.SecureDataLen = &v }
func (m *Header) SetSecureData(v string)           { m.SecureData = &v }
func (m *Header) SetMsgSeqNum(v int)               { m.MsgSeqNum = v }
func (m *Header) SetSenderSubID(v string)          { m.SenderSubID = &v }
func (m *Header) SetSenderLocationID(v string)     { m.SenderLocationID = &v }
func (m *Header) SetTargetSubID(v string)          { m.TargetSubID = &v }
func (m *Header) SetTargetLocationID(v string)     { m.TargetLocationID = &v }
func (m *Header) SetOnBehalfOfSubID(v string)      { m.OnBehalfOfSubID = &v }
func (m *Header) SetOnBehalfOfLocationID(v string) { m.OnBehalfOfLocationID = &v }
func (m *Header) SetDeliverToSubID(v string)       { m.DeliverToSubID = &v }
func (m *Header) SetDeliverToLocationID(v string)  { m.DeliverToLocationID = &v }
func (m *Header) SetPossDupFlag(v bool)            { m.PossDupFlag = &v }
func (m *Header) SetPossResend(v bool)             { m.PossResend = &v }
func (m *Header) SetSendingTime(v time.Time)       { m.SendingTime = v }
func (m *Header) SetOrigSendingTime(v time.Time)   { m.OrigSendingTime = &v }
func (m *Header) SetXmlDataLen(v int)              { m.XmlDataLen = &v }
func (m *Header) SetXmlData(v string)              { m.XmlData = &v }
func (m *Header) SetMessageEncoding(v string)      { m.MessageEncoding = &v }
func (m *Header) SetLastMsgSeqNumProcessed(v int)  { m.LastMsgSeqNumProcessed = &v }
func (m *Header) SetApplVerID(v string)            { m.ApplVerID = &v }
func (m *Header) SetCstmApplVerID(v string)        { m.CstmApplVerID = &v }
