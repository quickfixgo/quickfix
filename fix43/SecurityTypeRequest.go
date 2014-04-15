package fix43

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

//CreateSecurityTypeRequestBuilder returns an initialized SecurityTypeRequestBuilder with specified required fields.
func CreateSecurityTypeRequestBuilder(
	securityreqid field.SecurityReqID) SecurityTypeRequestBuilder {
	var builder SecurityTypeRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(securityreqid)
	return builder
}

//SecurityReqID is a required field for SecurityTypeRequest.
func (m SecurityTypeRequest) SecurityReqID() (field.SecurityReqID, error) {
	var f field.SecurityReqID
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for SecurityTypeRequest.
func (m SecurityTypeRequest) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for SecurityTypeRequest.
func (m SecurityTypeRequest) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for SecurityTypeRequest.
func (m SecurityTypeRequest) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for SecurityTypeRequest.
func (m SecurityTypeRequest) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for SecurityTypeRequest.
func (m SecurityTypeRequest) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}
