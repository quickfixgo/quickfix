package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type DerivativeSecurityListUpdateReport struct {
	quickfixgo.Message
}

func (m *DerivativeSecurityListUpdateReport) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *DerivativeSecurityListUpdateReport) SecurityResponseID() (*field.SecurityResponseID, error) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}
func (m *DerivativeSecurityListUpdateReport) SecurityRequestResult() (*field.SecurityRequestResult, error) {
	f := new(field.SecurityRequestResult)
	err := m.Body.Get(f)
	return f, err
}
func (m *DerivativeSecurityListUpdateReport) SecurityUpdateAction() (*field.SecurityUpdateAction, error) {
	f := new(field.SecurityUpdateAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *DerivativeSecurityListUpdateReport) TotNoRelatedSym() (*field.TotNoRelatedSym, error) {
	f := new(field.TotNoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}
func (m *DerivativeSecurityListUpdateReport) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}
