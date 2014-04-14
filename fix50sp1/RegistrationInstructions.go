package fix50sp1

import (
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

//NewRegistrationInstructionsBuilder returns an initialized RegistrationInstructionsBuilder with specified required fields.
func NewRegistrationInstructionsBuilder(
	registid field.RegistID,
	registtranstype field.RegistTransType,
	registrefid field.RegistRefID) *RegistrationInstructionsBuilder {
	builder := new(RegistrationInstructionsBuilder)
	builder.Body.Set(registid)
	builder.Body.Set(registtranstype)
	builder.Body.Set(registrefid)
	return builder
}

//RegistID is a required field for RegistrationInstructions.
func (m *RegistrationInstructions) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//RegistTransType is a required field for RegistrationInstructions.
func (m *RegistrationInstructions) RegistTransType() (*field.RegistTransType, error) {
	f := new(field.RegistTransType)
	err := m.Body.Get(f)
	return f, err
}

//RegistRefID is a required field for RegistrationInstructions.
func (m *RegistrationInstructions) RegistRefID() (*field.RegistRefID, error) {
	f := new(field.RegistRefID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a non-required field for RegistrationInstructions.
func (m *RegistrationInstructions) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for RegistrationInstructions.
func (m *RegistrationInstructions) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for RegistrationInstructions.
func (m *RegistrationInstructions) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//AcctIDSource is a non-required field for RegistrationInstructions.
func (m *RegistrationInstructions) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//RegistAcctType is a non-required field for RegistrationInstructions.
func (m *RegistrationInstructions) RegistAcctType() (*field.RegistAcctType, error) {
	f := new(field.RegistAcctType)
	err := m.Body.Get(f)
	return f, err
}

//TaxAdvantageType is a non-required field for RegistrationInstructions.
func (m *RegistrationInstructions) TaxAdvantageType() (*field.TaxAdvantageType, error) {
	f := new(field.TaxAdvantageType)
	err := m.Body.Get(f)
	return f, err
}

//OwnershipType is a non-required field for RegistrationInstructions.
func (m *RegistrationInstructions) OwnershipType() (*field.OwnershipType, error) {
	f := new(field.OwnershipType)
	err := m.Body.Get(f)
	return f, err
}

//NoRegistDtls is a non-required field for RegistrationInstructions.
func (m *RegistrationInstructions) NoRegistDtls() (*field.NoRegistDtls, error) {
	f := new(field.NoRegistDtls)
	err := m.Body.Get(f)
	return f, err
}

//NoDistribInsts is a non-required field for RegistrationInstructions.
func (m *RegistrationInstructions) NoDistribInsts() (*field.NoDistribInsts, error) {
	f := new(field.NoDistribInsts)
	err := m.Body.Get(f)
	return f, err
}
