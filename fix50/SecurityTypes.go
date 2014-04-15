package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityTypes msg type = w.
type SecurityTypes struct {
	message.Message
}

//SecurityTypesBuilder builds SecurityTypes messages.
type SecurityTypesBuilder struct {
	message.MessageBuilder
}

//CreateSecurityTypesBuilder returns an initialized SecurityTypesBuilder with specified required fields.
func CreateSecurityTypesBuilder(
	securityreqid field.SecurityReqID,
	securityresponseid field.SecurityResponseID,
	securityresponsetype field.SecurityResponseType) SecurityTypesBuilder {
	var builder SecurityTypesBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(securityreqid)
	builder.Body.Set(securityresponseid)
	builder.Body.Set(securityresponsetype)
	return builder
}

//SecurityReqID is a required field for SecurityTypes.
func (m SecurityTypes) SecurityReqID() (field.SecurityReqID, error) {
	var f field.SecurityReqID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityResponseID is a required field for SecurityTypes.
func (m SecurityTypes) SecurityResponseID() (field.SecurityResponseID, error) {
	var f field.SecurityResponseID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityResponseType is a required field for SecurityTypes.
func (m SecurityTypes) SecurityResponseType() (field.SecurityResponseType, error) {
	var f field.SecurityResponseType
	err := m.Body.Get(&f)
	return f, err
}

//TotNoSecurityTypes is a non-required field for SecurityTypes.
func (m SecurityTypes) TotNoSecurityTypes() (field.TotNoSecurityTypes, error) {
	var f field.TotNoSecurityTypes
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for SecurityTypes.
func (m SecurityTypes) LastFragment() (field.LastFragment, error) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityTypes is a non-required field for SecurityTypes.
func (m SecurityTypes) NoSecurityTypes() (field.NoSecurityTypes, error) {
	var f field.NoSecurityTypes
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for SecurityTypes.
func (m SecurityTypes) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for SecurityTypes.
func (m SecurityTypes) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for SecurityTypes.
func (m SecurityTypes) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for SecurityTypes.
func (m SecurityTypes) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for SecurityTypes.
func (m SecurityTypes) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for SecurityTypes.
func (m SecurityTypes) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}
