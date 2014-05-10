package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//RegistrationInstructions msg type = o.
type RegistrationInstructions struct {
	message.Message
}

//RegistrationInstructionsBuilder builds RegistrationInstructions messages.
type RegistrationInstructionsBuilder struct {
	message.MessageBuilder
}

//CreateRegistrationInstructionsBuilder returns an initialized RegistrationInstructionsBuilder with specified required fields.
func CreateRegistrationInstructionsBuilder(
	registid *field.RegistIDField,
	registtranstype *field.RegistTransTypeField,
	registrefid *field.RegistRefIDField) RegistrationInstructionsBuilder {
	var builder RegistrationInstructionsBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header.Set(field.NewMsgType("o"))
	builder.Body.Set(registid)
	builder.Body.Set(registtranstype)
	builder.Body.Set(registrefid)
	return builder
}

//RegistID is a required field for RegistrationInstructions.
func (m RegistrationInstructions) RegistID() (*field.RegistIDField, errors.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from RegistrationInstructions.
func (m RegistrationInstructions) GetRegistID(f *field.RegistIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistTransType is a required field for RegistrationInstructions.
func (m RegistrationInstructions) RegistTransType() (*field.RegistTransTypeField, errors.MessageRejectError) {
	f := &field.RegistTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistTransType reads a RegistTransType from RegistrationInstructions.
func (m RegistrationInstructions) GetRegistTransType(f *field.RegistTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistRefID is a required field for RegistrationInstructions.
func (m RegistrationInstructions) RegistRefID() (*field.RegistRefIDField, errors.MessageRejectError) {
	f := &field.RegistRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistRefID reads a RegistRefID from RegistrationInstructions.
func (m RegistrationInstructions) GetRegistRefID(f *field.RegistRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from RegistrationInstructions.
func (m RegistrationInstructions) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from RegistrationInstructions.
func (m RegistrationInstructions) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from RegistrationInstructions.
func (m RegistrationInstructions) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) AcctIDSource() (*field.AcctIDSourceField, errors.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from RegistrationInstructions.
func (m RegistrationInstructions) GetAcctIDSource(f *field.AcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistAcctType is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) RegistAcctType() (*field.RegistAcctTypeField, errors.MessageRejectError) {
	f := &field.RegistAcctTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistAcctType reads a RegistAcctType from RegistrationInstructions.
func (m RegistrationInstructions) GetRegistAcctType(f *field.RegistAcctTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TaxAdvantageType is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) TaxAdvantageType() (*field.TaxAdvantageTypeField, errors.MessageRejectError) {
	f := &field.TaxAdvantageTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTaxAdvantageType reads a TaxAdvantageType from RegistrationInstructions.
func (m RegistrationInstructions) GetTaxAdvantageType(f *field.TaxAdvantageTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OwnershipType is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) OwnershipType() (*field.OwnershipTypeField, errors.MessageRejectError) {
	f := &field.OwnershipTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOwnershipType reads a OwnershipType from RegistrationInstructions.
func (m RegistrationInstructions) GetOwnershipType(f *field.OwnershipTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRegistDtls is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) NoRegistDtls() (*field.NoRegistDtlsField, errors.MessageRejectError) {
	f := &field.NoRegistDtlsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRegistDtls reads a NoRegistDtls from RegistrationInstructions.
func (m RegistrationInstructions) GetNoRegistDtls(f *field.NoRegistDtlsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDistribInsts is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) NoDistribInsts() (*field.NoDistribInstsField, errors.MessageRejectError) {
	f := &field.NoDistribInstsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoDistribInsts reads a NoDistribInsts from RegistrationInstructions.
func (m RegistrationInstructions) GetNoDistribInsts(f *field.NoDistribInstsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
