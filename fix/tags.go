package fix

import (
	"github.com/cbusbey/quickfixgo/message"
)

const (
	ClOrdID message.Tag = 11
)

const (
	ApplVerID_FIX50SP1 = "8"
	ApplVerID_FIX27    = "0"
	ApplVerID_FIX50SP2 = "9"
	ApplVerID_FIX50    = "7"
	ApplVerID_FIX40    = "2"
	ApplVerID_FIX41    = "3"
	ApplVerID_FIX30    = "1"
	ApplVerID_FIX42    = "4"
	ApplVerID_FIX43    = "5"
	ApplVerID_FIX44    = "6"
)

const (
	BeginString_FIX40  = "FIX.4.0"
	BeginString_FIX41  = "FIX.4.1"
	BeginString_FIX42  = "FIX.4.2"
	BeginString_FIX43  = "FIX.4.3"
	BeginString_FIX44  = "FIX.4.4"
	BeginString_FIX50  = "FIX.5.0"
	BeginString_FIXT11 = "FIXT.1.1"
)

const (
	BeginString            message.Tag = 8
	BodyLength             message.Tag = 9
	MsgType                message.Tag = 35
	SenderCompID           message.Tag = 49
	TargetCompID           message.Tag = 56
	OnBehalfOfCompID       message.Tag = 115
	DeliverToCompID        message.Tag = 128
	SecureDataLen          message.Tag = 90
	MsgSeqNum              message.Tag = 34
	SenderSubID            message.Tag = 50
	SenderLocationID       message.Tag = 142
	TargetSubID            message.Tag = 57
	TargetLocationID       message.Tag = 143
	OnBehalfOfSubID        message.Tag = 116
	OnBehalfOfLocationID   message.Tag = 144
	DeliverToSubID         message.Tag = 129
	DeliverToLocationID    message.Tag = 145
	PossDupFlag            message.Tag = 43
	PossResend             message.Tag = 97
	SendingTime            message.Tag = 52
	OrigSendingTime        message.Tag = 122
	XmlDataLen             message.Tag = 212
	XmlData                message.Tag = 213
	MessageEncoding        message.Tag = 347
	LastMsgSeqNumProcessed message.Tag = 369
	OnBehalfOfSendingTime  message.Tag = 370
	ApplVerID              message.Tag = 1128
	CstmApplVerID          message.Tag = 1129
	NoHops                 message.Tag = 627

	BeginSeqNo           message.Tag = 7
	EndSeqNo             message.Tag = 16
	NewSeqNo             message.Tag = 36
	RefSeqNum            message.Tag = 45
	Text                 message.Tag = 58
	EncryptMethod        message.Tag = 98
	HeartBtInt           message.Tag = 108
	TestReqID            message.Tag = 112
	GapFillFlag          message.Tag = 123
	RefTagID             message.Tag = 371
	RefMsgType           message.Tag = 372
	SessionRejectReason  message.Tag = 373
	BusinessRejectReason message.Tag = 380
	DefaultApplVerID     message.Tag = 1137

	SignatureLength message.Tag = 93
	Signature       message.Tag = 89
	CheckSum        message.Tag = 10
)

//Returns true if tag belongs in the message trailer
func IsTrailer(t message.Tag) bool {
	switch t {
	case SignatureLength, Signature, CheckSum:
		return true
	}
	return false
}

//Returns true if tag belongs in the message header
func IsHeader(t message.Tag) bool {
	switch t {
	case BeginString,
		BodyLength,
		MsgType,
		SenderCompID,
		TargetCompID,
		OnBehalfOfCompID,
		DeliverToCompID,
		SecureDataLen,
		MsgSeqNum,
		SenderSubID,
		SenderLocationID,
		TargetSubID,
		TargetLocationID,
		OnBehalfOfSubID,
		OnBehalfOfLocationID,
		DeliverToSubID,
		DeliverToLocationID,
		PossDupFlag,
		PossResend,
		SendingTime,
		OrigSendingTime,
		XmlDataLen,
		XmlData,
		MessageEncoding,
		LastMsgSeqNumProcessed,
		OnBehalfOfSendingTime,
		ApplVerID,
		CstmApplVerID,
		NoHops:
		return true
	}

	return false
}
