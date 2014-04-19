package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//RegistrationInstructionsResponse msg type = p.
type RegistrationInstructionsResponse struct {
	message.Message
}

//RegistrationInstructionsResponseBuilder builds RegistrationInstructionsResponse messages.
type RegistrationInstructionsResponseBuilder struct {
	message.MessageBuilder
}

//CreateRegistrationInstructionsResponseBuilder returns an initialized RegistrationInstructionsResponseBuilder with specified required fields.
func CreateRegistrationInstructionsResponseBuilder(
	registid field.RegistID,
	registtranstype field.RegistTransType,
	registrefid field.RegistRefID,
	registstatus field.RegistStatus) RegistrationInstructionsResponseBuilder {
	var builder RegistrationInstructionsResponseBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(registid)
	builder.Body.Set(registtranstype)
	builder.Body.Set(registrefid)
	builder.Body.Set(registstatus)
	return builder
}

//RegistID is a required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistID() (field.RegistID, errors.MessageRejectError) {
	var f field.RegistID
	err := m.Body.Get(&f)
	return f, err
}

//RegistTransType is a required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistTransType() (field.RegistTransType, errors.MessageRejectError) {
	var f field.RegistTransType
	err := m.Body.Get(&f)
	return f, err
}

//RegistRefID is a required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistRefID() (field.RegistRefID, errors.MessageRejectError) {
	var f field.RegistRefID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) ClOrdID() (field.ClOrdID, errors.MessageRejectError) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) NoPartyIDs() (field.NoPartyIDs, errors.MessageRejectError) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) Account() (field.Account, errors.MessageRejectError) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) AcctIDSource() (field.AcctIDSource, errors.MessageRejectError) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//RegistStatus is a required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistStatus() (field.RegistStatus, errors.MessageRejectError) {
	var f field.RegistStatus
	err := m.Body.Get(&f)
	return f, err
}

//RegistRejReasonCode is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistRejReasonCode() (field.RegistRejReasonCode, errors.MessageRejectError) {
	var f field.RegistRejReasonCode
	err := m.Body.Get(&f)
	return f, err
}

//RegistRejReasonText is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistRejReasonText() (field.RegistRejReasonText, errors.MessageRejectError) {
	var f field.RegistRejReasonText
	err := m.Body.Get(&f)
	return f, err
}
