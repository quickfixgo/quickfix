package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
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

//Returns true if tag belongs in the message trailer
func IsTrailer(t tag.Tag) bool {
	switch t {
	case tag.SignatureLength, tag.Signature, tag.CheckSum:
		return true
	}
	return false
}

//Returns true if tag belongs in the message header
func IsHeader(t tag.Tag) bool {
	switch t {
	case tag.BeginString,
		tag.BodyLength,
		tag.MsgType,
		tag.SenderCompID,
		tag.TargetCompID,
		tag.OnBehalfOfCompID,
		tag.DeliverToCompID,
		tag.SecureDataLen,
		tag.MsgSeqNum,
		tag.SenderSubID,
		tag.SenderLocationID,
		tag.TargetSubID,
		tag.TargetLocationID,
		tag.OnBehalfOfSubID,
		tag.OnBehalfOfLocationID,
		tag.DeliverToSubID,
		tag.DeliverToLocationID,
		tag.PossDupFlag,
		tag.PossResend,
		tag.SendingTime,
		tag.OrigSendingTime,
		tag.XmlDataLen,
		tag.XmlData,
		tag.MessageEncoding,
		tag.LastMsgSeqNumProcessed,
		tag.OnBehalfOfSendingTime,
		tag.ApplVerID,
		tag.CstmApplVerID,
		tag.NoHops:
		return true
	}

	return false
}
