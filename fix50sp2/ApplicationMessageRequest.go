package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ApplicationMessageRequest msg type = BW.
type ApplicationMessageRequest struct {
	message.Message
}

//ApplicationMessageRequestBuilder builds ApplicationMessageRequest messages.
type ApplicationMessageRequestBuilder struct {
	message.MessageBuilder
}

//NewApplicationMessageRequestBuilder returns an initialized ApplicationMessageRequestBuilder with specified required fields.
func NewApplicationMessageRequestBuilder(
	applreqid field.ApplReqID,
	applreqtype field.ApplReqType) *ApplicationMessageRequestBuilder {
	builder := new(ApplicationMessageRequestBuilder)
	builder.Body.Set(applreqid)
	builder.Body.Set(applreqtype)
	return builder
}

//ApplReqID is a required field for ApplicationMessageRequest.
func (m *ApplicationMessageRequest) ApplReqID() (*field.ApplReqID, error) {
	f := new(field.ApplReqID)
	err := m.Body.Get(f)
	return f, err
}

//ApplReqType is a required field for ApplicationMessageRequest.
func (m *ApplicationMessageRequest) ApplReqType() (*field.ApplReqType, error) {
	f := new(field.ApplReqType)
	err := m.Body.Get(f)
	return f, err
}

//NoApplIDs is a non-required field for ApplicationMessageRequest.
func (m *ApplicationMessageRequest) NoApplIDs() (*field.NoApplIDs, error) {
	f := new(field.NoApplIDs)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for ApplicationMessageRequest.
func (m *ApplicationMessageRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for ApplicationMessageRequest.
func (m *ApplicationMessageRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for ApplicationMessageRequest.
func (m *ApplicationMessageRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for ApplicationMessageRequest.
func (m *ApplicationMessageRequest) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
