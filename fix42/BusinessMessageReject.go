package fix42

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//BusinessMessageReject msg type = j.
type BusinessMessageReject struct {
	message.Message
}

//BusinessMessageRejectBuilder builds BusinessMessageReject messages.
type BusinessMessageRejectBuilder struct {
	message.MessageBuilder
}

//NewBusinessMessageRejectBuilder returns an initialized BusinessMessageRejectBuilder with specified required fields.
func NewBusinessMessageRejectBuilder(
	refmsgtype field.RefMsgType,
	businessrejectreason field.BusinessRejectReason) *BusinessMessageRejectBuilder {
	builder := new(BusinessMessageRejectBuilder)
	builder.Body.Set(refmsgtype)
	builder.Body.Set(businessrejectreason)
	return builder
}

//RefSeqNum is a non-required field for BusinessMessageReject.
func (m *BusinessMessageReject) RefSeqNum() (*field.RefSeqNum, error) {
	f := new(field.RefSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//RefMsgType is a required field for BusinessMessageReject.
func (m *BusinessMessageReject) RefMsgType() (*field.RefMsgType, error) {
	f := new(field.RefMsgType)
	err := m.Body.Get(f)
	return f, err
}

//BusinessRejectRefID is a non-required field for BusinessMessageReject.
func (m *BusinessMessageReject) BusinessRejectRefID() (*field.BusinessRejectRefID, error) {
	f := new(field.BusinessRejectRefID)
	err := m.Body.Get(f)
	return f, err
}

//BusinessRejectReason is a required field for BusinessMessageReject.
func (m *BusinessMessageReject) BusinessRejectReason() (*field.BusinessRejectReason, error) {
	f := new(field.BusinessRejectReason)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for BusinessMessageReject.
func (m *BusinessMessageReject) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for BusinessMessageReject.
func (m *BusinessMessageReject) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for BusinessMessageReject.
func (m *BusinessMessageReject) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
