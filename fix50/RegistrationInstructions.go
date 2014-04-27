package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	registid field.RegistID,
	registtranstype field.RegistTransType,
	registrefid field.RegistRefID) RegistrationInstructionsBuilder {
	var builder RegistrationInstructionsBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.BuildMsgType("o"))
	builder.Body.Set(registid)
	builder.Body.Set(registtranstype)
	builder.Body.Set(registrefid)
	return builder
}

//RegistID is a required field for RegistrationInstructions.
func (m RegistrationInstructions) RegistID() (*field.RegistID, errors.MessageRejectError) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from RegistrationInstructions.
func (m RegistrationInstructions) GetRegistID(f *field.RegistID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistTransType is a required field for RegistrationInstructions.
func (m RegistrationInstructions) RegistTransType() (*field.RegistTransType, errors.MessageRejectError) {
	f := new(field.RegistTransType)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistTransType reads a RegistTransType from RegistrationInstructions.
func (m RegistrationInstructions) GetRegistTransType(f *field.RegistTransType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistRefID is a required field for RegistrationInstructions.
func (m RegistrationInstructions) RegistRefID() (*field.RegistRefID, errors.MessageRejectError) {
	f := new(field.RegistRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistRefID reads a RegistRefID from RegistrationInstructions.
func (m RegistrationInstructions) GetRegistRefID(f *field.RegistRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from RegistrationInstructions.
func (m RegistrationInstructions) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from RegistrationInstructions.
func (m RegistrationInstructions) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from RegistrationInstructions.
func (m RegistrationInstructions) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from RegistrationInstructions.
func (m RegistrationInstructions) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistAcctType is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) RegistAcctType() (*field.RegistAcctType, errors.MessageRejectError) {
	f := new(field.RegistAcctType)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistAcctType reads a RegistAcctType from RegistrationInstructions.
func (m RegistrationInstructions) GetRegistAcctType(f *field.RegistAcctType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TaxAdvantageType is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) TaxAdvantageType() (*field.TaxAdvantageType, errors.MessageRejectError) {
	f := new(field.TaxAdvantageType)
	err := m.Body.Get(f)
	return f, err
}

//GetTaxAdvantageType reads a TaxAdvantageType from RegistrationInstructions.
func (m RegistrationInstructions) GetTaxAdvantageType(f *field.TaxAdvantageType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OwnershipType is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) OwnershipType() (*field.OwnershipType, errors.MessageRejectError) {
	f := new(field.OwnershipType)
	err := m.Body.Get(f)
	return f, err
}

//GetOwnershipType reads a OwnershipType from RegistrationInstructions.
func (m RegistrationInstructions) GetOwnershipType(f *field.OwnershipType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRegistDtls is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) NoRegistDtls() (*field.NoRegistDtls, errors.MessageRejectError) {
	f := new(field.NoRegistDtls)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRegistDtls reads a NoRegistDtls from RegistrationInstructions.
func (m RegistrationInstructions) GetNoRegistDtls(f *field.NoRegistDtls) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDistribInsts is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) NoDistribInsts() (*field.NoDistribInsts, errors.MessageRejectError) {
	f := new(field.NoDistribInsts)
	err := m.Body.Get(f)
	return f, err
}

//GetNoDistribInsts reads a NoDistribInsts from RegistrationInstructions.
func (m RegistrationInstructions) GetNoDistribInsts(f *field.NoDistribInsts) errors.MessageRejectError {
	return m.Body.Get(f)
}
