package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityTypeRequest msg type = v.
type SecurityTypeRequest struct {
	message.Message
}

//SecurityTypeRequestBuilder builds SecurityTypeRequest messages.
type SecurityTypeRequestBuilder struct {
	message.MessageBuilder
}

//NewSecurityTypeRequestBuilder returns an initialized SecurityTypeRequestBuilder with specified required fields.
func NewSecurityTypeRequestBuilder(
	securityreqid field.SecurityReqID) *SecurityTypeRequestBuilder {
	builder := new(SecurityTypeRequestBuilder)
	builder.Body.Set(securityreqid)
	return builder
}

//SecurityReqID is a required field for SecurityTypeRequest.
func (m *SecurityTypeRequest) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for SecurityTypeRequest.
func (m *SecurityTypeRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for SecurityTypeRequest.
func (m *SecurityTypeRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for SecurityTypeRequest.
func (m *SecurityTypeRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for SecurityTypeRequest.
func (m *SecurityTypeRequest) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for SecurityTypeRequest.
func (m *SecurityTypeRequest) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for SecurityTypeRequest.
func (m *SecurityTypeRequest) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for SecurityTypeRequest.
func (m *SecurityTypeRequest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for SecurityTypeRequest.
func (m *SecurityTypeRequest) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}
