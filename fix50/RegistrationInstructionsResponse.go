package fix50

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
	builder.Header.Set(field.BuildMsgType("p"))
	builder.Body.Set(registid)
	builder.Body.Set(registtranstype)
	builder.Body.Set(registrefid)
	builder.Body.Set(registstatus)
	return builder
}

//RegistID is a required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistID() (*field.RegistID, errors.MessageRejectError) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetRegistID(f *field.RegistID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistTransType is a required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistTransType() (*field.RegistTransType, errors.MessageRejectError) {
	f := new(field.RegistTransType)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistTransType reads a RegistTransType from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetRegistTransType(f *field.RegistTransType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistRefID is a required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistRefID() (*field.RegistRefID, errors.MessageRejectError) {
	f := new(field.RegistRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistRefID reads a RegistRefID from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetRegistRefID(f *field.RegistRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistStatus is a required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistStatus() (*field.RegistStatus, errors.MessageRejectError) {
	f := new(field.RegistStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistStatus reads a RegistStatus from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetRegistStatus(f *field.RegistStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistRejReasonCode is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistRejReasonCode() (*field.RegistRejReasonCode, errors.MessageRejectError) {
	f := new(field.RegistRejReasonCode)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistRejReasonCode reads a RegistRejReasonCode from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetRegistRejReasonCode(f *field.RegistRejReasonCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistRejReasonText is a non-required field for RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) RegistRejReasonText() (*field.RegistRejReasonText, errors.MessageRejectError) {
	f := new(field.RegistRejReasonText)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistRejReasonText reads a RegistRejReasonText from RegistrationInstructionsResponse.
func (m RegistrationInstructionsResponse) GetRegistRejReasonText(f *field.RegistRejReasonText) errors.MessageRejectError {
	return m.Body.Get(f)
}
