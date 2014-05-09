package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	registid *field.RegistIDField,
	registtranstype *field.RegistTransTypeField,
	registrefid *field.RegistRefIDField,
	registstatus *field.RegistStatusField) RegistrationInstructionsResponseBuilder {
	var builder RegistrationInstructionsResponseBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("p"))
	builder.Body.Set(registid)
	builder.Body.Set(registtranstype)
	builder.Body.Set(registrefid)
	builder.Body.Set(registstatus)
	return builder
}

//RegistID is a required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistID() (*field.RegistIDField, errors.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetRegistID(f *field.RegistIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistTransType is a required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistTransType() (*field.RegistTransTypeField, errors.MessageRejectError) {
	f := &field.RegistTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistTransType reads a RegistTransType from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetRegistTransType(f *field.RegistTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistRefID is a required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistRefID() (*field.RegistRefIDField, errors.MessageRejectError) {
	f := &field.RegistRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistRefID reads a RegistRefID from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetRegistRefID(f *field.RegistRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) AcctIDSource() (*field.AcctIDSourceField, errors.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetAcctIDSource(f *field.AcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistStatus is a required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistStatus() (*field.RegistStatusField, errors.MessageRejectError) {
	f := &field.RegistStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistStatus reads a RegistStatus from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetRegistStatus(f *field.RegistStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistRejReasonCode is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistRejReasonCode() (*field.RegistRejReasonCodeField, errors.MessageRejectError) {
	f := &field.RegistRejReasonCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistRejReasonCode reads a RegistRejReasonCode from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetRegistRejReasonCode(f *field.RegistRejReasonCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistRejReasonText is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistRejReasonText() (*field.RegistRejReasonTextField, errors.MessageRejectError) {
	f := &field.RegistRejReasonTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistRejReasonText reads a RegistRejReasonText from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetRegistRejReasonText(f *field.RegistRejReasonTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
