package reject

import (
	"github.com/cbusbey/quickfixgo/message"
)

type RejectReason int

type MessageReject interface {
	RejectedMessage() message.Message
	RejectReason() RejectReason
	IsBusinessReject() bool
	RefTagID() message.Tag
	error
}

type messageRejectBase struct {
	rejectedMessage message.Message
	rejectReason    RejectReason
	text            string
	businessReject  bool
	refTagID        message.Tag
}

func (reject messageRejectBase) RejectedMessage() message.Message { return reject.rejectedMessage }
func (reject messageRejectBase) RejectReason() RejectReason       { return reject.rejectReason }
func (reject messageRejectBase) Error() string                    { return reject.text }
func (reject messageRejectBase) IsBusinessReject() bool           { return reject.businessReject }
func (reject messageRejectBase) RefTagID() message.Tag            { return reject.refTagID }
