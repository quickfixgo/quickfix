package message

import (
	"github.com/quickfixgo/quickfix/fix"
)

type RejectReason int

type MessageReject interface {
	RejectedMessage() Message
	RejectReason() RejectReason
	IsBusinessReject() bool
	RefTagID() *fix.Tag
	error
}

type messageRejectBase struct {
	rejectedMessage Message
	rejectReason    RejectReason
	text            string
	businessReject  bool
	refTagID        *fix.Tag
}

func (reject messageRejectBase) RejectedMessage() Message   { return reject.rejectedMessage }
func (reject messageRejectBase) RejectReason() RejectReason { return reject.rejectReason }
func (reject messageRejectBase) Error() string              { return reject.text }
func (reject messageRejectBase) IsBusinessReject() bool     { return reject.businessReject }
func (reject messageRejectBase) RefTagID() *fix.Tag         { return reject.refTagID }
