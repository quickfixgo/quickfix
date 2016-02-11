package fix43

import (
	"time"
)

//NoHops is a repeating group in Header
type NoHops struct {
	//HopCompID is a non-required field for NoHops.
	HopCompID *string `fix:"628"`
	//HopSendingTime is a non-required field for NoHops.
	HopSendingTime *time.Time `fix:"629"`
	//HopRefID is a non-required field for NoHops.
	HopRefID *int `fix:"630"`
}

//Header is the fix43 Header type
type Header struct {
	//BeginString is a required field for Header.
	BeginString string `fix:"8,default=FIX.4.3"`
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
	//OnBehalfOfSendingTime is a non-required field for Header.
	OnBehalfOfSendingTime *time.Time `fix:"370"`
	//NoHops is a non-required field for Header.
	NoHops []NoHops `fix:"627,omitempty"`
}
