package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ApplicationMessageRequestAck msg type = BX.
type ApplicationMessageRequestAck struct {
	message.Message
}

//ApplicationMessageRequestAckBuilder builds ApplicationMessageRequestAck messages.
type ApplicationMessageRequestAckBuilder struct {
	message.MessageBuilder
}

//NewApplicationMessageRequestAckBuilder returns an initialized ApplicationMessageRequestAckBuilder with specified required fields.
func NewApplicationMessageRequestAckBuilder(
	applresponseid field.ApplResponseID) *ApplicationMessageRequestAckBuilder {
	builder := new(ApplicationMessageRequestAckBuilder)
	builder.Body.Set(applresponseid)
	return builder
}

//ApplResponseID is a required field for ApplicationMessageRequestAck.
func (m *ApplicationMessageRequestAck) ApplResponseID() (*field.ApplResponseID, error) {
	f := new(field.ApplResponseID)
	err := m.Body.Get(f)
	return f, err
}

//ApplReqID is a non-required field for ApplicationMessageRequestAck.
func (m *ApplicationMessageRequestAck) ApplReqID() (*field.ApplReqID, error) {
	f := new(field.ApplReqID)
	err := m.Body.Get(f)
	return f, err
}

//ApplReqType is a non-required field for ApplicationMessageRequestAck.
func (m *ApplicationMessageRequestAck) ApplReqType() (*field.ApplReqType, error) {
	f := new(field.ApplReqType)
	err := m.Body.Get(f)
	return f, err
}

//ApplResponseType is a non-required field for ApplicationMessageRequestAck.
func (m *ApplicationMessageRequestAck) ApplResponseType() (*field.ApplResponseType, error) {
	f := new(field.ApplResponseType)
	err := m.Body.Get(f)
	return f, err
}

//ApplTotalMessageCount is a non-required field for ApplicationMessageRequestAck.
func (m *ApplicationMessageRequestAck) ApplTotalMessageCount() (*field.ApplTotalMessageCount, error) {
	f := new(field.ApplTotalMessageCount)
	err := m.Body.Get(f)
	return f, err
}

//NoApplIDs is a non-required field for ApplicationMessageRequestAck.
func (m *ApplicationMessageRequestAck) NoApplIDs() (*field.NoApplIDs, error) {
	f := new(field.NoApplIDs)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for ApplicationMessageRequestAck.
func (m *ApplicationMessageRequestAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for ApplicationMessageRequestAck.
func (m *ApplicationMessageRequestAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for ApplicationMessageRequestAck.
func (m *ApplicationMessageRequestAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for ApplicationMessageRequestAck.
func (m *ApplicationMessageRequestAck) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
