//Package tag defines the tag type and known tag values
package tag

//Abstract type representing a fix tag
type Tag int

//Returns true if tag belongs in the message trailer
func (t Tag) IsTrailer() bool {
	switch t {
	case SignatureLength, Signature, CheckSum:
		return true
	}
	return false
}

//Returns true if tag belongs in the message header
func (t Tag) IsHeader() bool {
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
