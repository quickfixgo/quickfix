package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Reject msg type = 3.
type Reject struct {
	message.Message
}

//RejectBuilder builds Reject messages.
type RejectBuilder struct {
	message.MessageBuilder
}

//NewRejectBuilder returns an initialized RejectBuilder with specified required fields.
func NewRejectBuilder(
	refseqnum field.RefSeqNum) *RejectBuilder {
	builder := new(RejectBuilder)
	builder.Body.Set(refseqnum)
	return builder
}

//RefSeqNum is a required field for Reject.
func (m *Reject) RefSeqNum() (*field.RefSeqNum, error) {
	f := new(field.RefSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//RefTagID is a non-required field for Reject.
func (m *Reject) RefTagID() (*field.RefTagID, error) {
	f := new(field.RefTagID)
	err := m.Body.Get(f)
	return f, err
}

//RefMsgType is a non-required field for Reject.
func (m *Reject) RefMsgType() (*field.RefMsgType, error) {
	f := new(field.RefMsgType)
	err := m.Body.Get(f)
	return f, err
}

//SessionRejectReason is a non-required field for Reject.
func (m *Reject) SessionRejectReason() (*field.SessionRejectReason, error) {
	f := new(field.SessionRejectReason)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for Reject.
func (m *Reject) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for Reject.
func (m *Reject) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for Reject.
func (m *Reject) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
