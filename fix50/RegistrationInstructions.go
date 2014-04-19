package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
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
	registid field.RegistID,
	registtranstype field.RegistTransType,
	registrefid field.RegistRefID) RegistrationInstructionsBuilder {
	var builder RegistrationInstructionsBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(registid)
	builder.Body.Set(registtranstype)
	builder.Body.Set(registrefid)
	return builder
}

//RegistID is a required field for RegistrationInstructions.
func (m RegistrationInstructions) RegistID() (field.RegistID, errors.MessageRejectError) {
	var f field.RegistID
	err := m.Body.Get(&f)
	return f, err
}

//RegistTransType is a required field for RegistrationInstructions.
func (m RegistrationInstructions) RegistTransType() (field.RegistTransType, errors.MessageRejectError) {
	var f field.RegistTransType
	err := m.Body.Get(&f)
	return f, err
}

//RegistRefID is a required field for RegistrationInstructions.
func (m RegistrationInstructions) RegistRefID() (field.RegistRefID, errors.MessageRejectError) {
	var f field.RegistRefID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) ClOrdID() (field.ClOrdID, errors.MessageRejectError) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) NoPartyIDs() (field.NoPartyIDs, errors.MessageRejectError) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) Account() (field.Account, errors.MessageRejectError) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) AcctIDSource() (field.AcctIDSource, errors.MessageRejectError) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//RegistAcctType is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) RegistAcctType() (field.RegistAcctType, errors.MessageRejectError) {
	var f field.RegistAcctType
	err := m.Body.Get(&f)
	return f, err
}

//TaxAdvantageType is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) TaxAdvantageType() (field.TaxAdvantageType, errors.MessageRejectError) {
	var f field.TaxAdvantageType
	err := m.Body.Get(&f)
	return f, err
}

//OwnershipType is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) OwnershipType() (field.OwnershipType, errors.MessageRejectError) {
	var f field.OwnershipType
	err := m.Body.Get(&f)
	return f, err
}

//NoRegistDtls is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) NoRegistDtls() (field.NoRegistDtls, errors.MessageRejectError) {
	var f field.NoRegistDtls
	err := m.Body.Get(&f)
	return f, err
}

//NoDistribInsts is a non-required field for RegistrationInstructions.
func (m RegistrationInstructions) NoDistribInsts() (field.NoDistribInsts, errors.MessageRejectError) {
	var f field.NoDistribInsts
	err := m.Body.Get(&f)
	return f, err
}
