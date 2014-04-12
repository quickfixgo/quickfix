package fix42

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type BusinessMessageReject struct {
	message.Message
}

func (m *BusinessMessageReject) RefSeqNum() (*field.RefSeqNum, error) {
	f := new(field.RefSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *BusinessMessageReject) RefMsgType() (*field.RefMsgType, error) {
	f := new(field.RefMsgType)
	err := m.Body.Get(f)
	return f, err
}
func (m *BusinessMessageReject) BusinessRejectRefID() (*field.BusinessRejectRefID, error) {
	f := new(field.BusinessRejectRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *BusinessMessageReject) BusinessRejectReason() (*field.BusinessRejectReason, error) {
	f := new(field.BusinessRejectReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *BusinessMessageReject) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *BusinessMessageReject) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *BusinessMessageReject) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
