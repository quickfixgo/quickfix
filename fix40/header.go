package fix40

import (
	"time"
)

//Header is the fix40 Header type
type Header struct {
	//BeginString is a required field for Header.
	BeginString string `fix:"8,default=FIX.4.0"`
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
	//TargetSubID is a non-required field for Header.
	TargetSubID *string `fix:"57"`
	//OnBehalfOfSubID is a non-required field for Header.
	OnBehalfOfSubID *string `fix:"116"`
	//DeliverToSubID is a non-required field for Header.
	DeliverToSubID *string `fix:"129"`
	//PossDupFlag is a non-required field for Header.
	PossDupFlag *string `fix:"43"`
	//PossResend is a non-required field for Header.
	PossResend *string `fix:"97"`
	//SendingTime is a required field for Header.
	SendingTime time.Time `fix:"52"`
	//OrigSendingTime is a non-required field for Header.
	OrigSendingTime *time.Time `fix:"122"`
}
