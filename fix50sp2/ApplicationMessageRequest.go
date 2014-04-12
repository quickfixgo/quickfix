package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type ApplicationMessageRequest struct {
	message.Message
}

func (m *ApplicationMessageRequest) ApplReqID() (*field.ApplReqID, error) {
	f := new(field.ApplReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequest) ApplReqType() (*field.ApplReqType, error) {
	f := new(field.ApplReqType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequest) NoApplIDs() (*field.NoApplIDs, error) {
	f := new(field.NoApplIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequest) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
