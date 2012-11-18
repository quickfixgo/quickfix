package quickfixgo

import(
    "quickfixgo/message"
    )

type MessageReject interface {
  RejectedMessage() message.Message
  RejectReason() int
  IsBusinessReject() bool
  RefTagID() message.Tag
  error
}

/*type messageRejectBase struct {
  rejectedMessage message.Message
  rejectReason int
  text string
  businessReject bool
  refTagID tag.Tag
}

func (reject messageRejectBase) RejectedMessage() message.Message {return reject.rejectedMessage}
func (reject messageRejectBase) RejectReason() int {return reject.rejectReason}
func (reject messageRejectBase) Error() string {return reject.text}
func (reject messageRejectBase) IsBusinessReject() bool {return reject.businessReject}
func (reject messageRejectBase) RefTagID() tag.Tag {return reject.refTagID}
*/
