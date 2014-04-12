//Package tag declares known tag values
package tag

import (
	"github.com/quickfixgo/quickfix/fix"
)

//IsTrailer returns true if tag belongs in the message trailer
func IsTrailer(t fix.Tag) bool {
	switch t {
	case SignatureLength, Signature, CheckSum:
		return true
	}
	return false
}

//IsHeader returns true if tag belongs in the message header
func IsHeader(t fix.Tag) bool {
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
