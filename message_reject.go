package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
)

type RejectReason int

type MessageReject interface {
	RejectedMessage() Message
	RejectReason() RejectReason
	IsBusinessReject() bool
	RefTagID() tag.Tag
	error
}

type messageRejectBase struct {
	rejectedMessage Message
	rejectReason    RejectReason
	text            string
	businessReject  bool
	refTagID        tag.Tag
}

func (reject messageRejectBase) RejectedMessage() Message { return reject.rejectedMessage }
func (reject messageRejectBase) RejectReason() RejectReason       { return reject.rejectReason }
func (reject messageRejectBase) Error() string                    { return reject.text }
func (reject messageRejectBase) IsBusinessReject() bool           { return reject.businessReject }
func (reject messageRejectBase) RefTagID() tag.Tag                { return reject.refTagID }
