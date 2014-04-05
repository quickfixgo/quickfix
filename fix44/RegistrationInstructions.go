package fix44

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type RegistrationInstructions struct {
	quickfixgo.Message
}

func (m *RegistrationInstructions) RegistTransType() (*field.RegistTransType, error) {
	f := new(field.RegistTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructions) RegistRefID() (*field.RegistRefID, error) {
	f := new(field.RegistRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructions) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructions) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructions) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructions) TaxAdvantageType() (*field.TaxAdvantageType, error) {
	f := new(field.TaxAdvantageType)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructions) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructions) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructions) RegistAcctType() (*field.RegistAcctType, error) {
	f := new(field.RegistAcctType)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructions) OwnershipType() (*field.OwnershipType, error) {
	f := new(field.OwnershipType)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructions) NoRegistDtls() (*field.NoRegistDtls, error) {
	f := new(field.NoRegistDtls)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructions) NoDistribInsts() (*field.NoDistribInsts, error) {
	f := new(field.NoDistribInsts)
	err := m.Body.Get(f)
	return f, err
}
