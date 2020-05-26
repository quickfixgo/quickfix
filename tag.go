package quickfix

//Tag is a typed int representing a FIX tag
type Tag int

const (
	tagBeginString            Tag = 8
	tagBodyLength             Tag = 9
	tagMsgType                Tag = 35
	tagSenderCompID           Tag = 49
	tagTargetCompID           Tag = 56
	tagOnBehalfOfCompID       Tag = 115
	tagDeliverToCompID        Tag = 128
	tagSecureDataLen          Tag = 90
	tagMsgSeqNum              Tag = 34
	tagSenderSubID            Tag = 50
	tagSenderLocationID       Tag = 142
	tagTargetSubID            Tag = 57
	tagTargetLocationID       Tag = 143
	tagOnBehalfOfSubID        Tag = 116
	tagOnBehalfOfLocationID   Tag = 144
	tagDeliverToSubID         Tag = 129
	tagDeliverToLocationID    Tag = 145
	tagPossDupFlag            Tag = 43
	tagPossResend             Tag = 97
	tagSendingTime            Tag = 52
	tagOrigSendingTime        Tag = 122
	tagXMLDataLen             Tag = 212
	tagXMLData                Tag = 213
	tagMessageEncoding        Tag = 347
	tagLastMsgSeqNumProcessed Tag = 369
	tagOnBehalfOfSendingTime  Tag = 370
	tagApplVerID              Tag = 1128
	tagCstmApplVerID          Tag = 1129
	tagNoHops                 Tag = 627
	tagApplExtID              Tag = 1156
	tagSecureData             Tag = 91
	tagHopCompID              Tag = 628
	tagHopSendingTime         Tag = 629
	tagHopRefID               Tag = 630

	tagHeartBtInt           Tag = 108
	tagBusinessRejectReason Tag = 380
	tagSessionRejectReason  Tag = 373
	tagRefMsgType           Tag = 372
	tagBusinessRejectRefID  Tag = 379
	tagRefTagID             Tag = 371
	tagRefSeqNum            Tag = 45
	tagEncryptMethod        Tag = 98
	tagResetSeqNumFlag      Tag = 141
	tagDefaultApplVerID     Tag = 1137
	tagText                 Tag = 58
	tagTestReqID            Tag = 112
	tagGapFillFlag          Tag = 123
	tagNewSeqNo             Tag = 36
	tagBeginSeqNo           Tag = 7
	tagEndSeqNo             Tag = 16

	tagSignatureLength Tag = 93
	tagSignature       Tag = 89
	tagCheckSum        Tag = 10
)

//IsTrailer returns true if tag belongs in the message trailer
func (t Tag) IsTrailer() bool {
	switch t {
	case tagSignatureLength, tagSignature, tagCheckSum:
		return true
	}
	return false
}

//IsHeader returns true if tag belongs in the message header
func (t Tag) IsHeader() bool {
	switch t {
	case tagBeginString,
		tagBodyLength,
		tagMsgType,
		tagSenderCompID,
		tagTargetCompID,
		tagOnBehalfOfCompID,
		tagDeliverToCompID,
		tagSecureDataLen,
		tagMsgSeqNum,
		tagSenderSubID,
		tagSenderLocationID,
		tagTargetSubID,
		tagTargetLocationID,
		tagOnBehalfOfSubID,
		tagOnBehalfOfLocationID,
		tagDeliverToSubID,
		tagDeliverToLocationID,
		tagPossDupFlag,
		tagPossResend,
		tagSendingTime,
		tagOrigSendingTime,
		tagXMLDataLen,
		tagXMLData,
		tagMessageEncoding,
		tagLastMsgSeqNumProcessed,
		tagOnBehalfOfSendingTime,
		tagApplVerID,
		tagCstmApplVerID,
		tagNoHops,
		tagApplExtID,
		tagSecureData,
		tagHopCompID,
		tagHopSendingTime,
		tagHopRefID:

		return true
	}

	return false
}
