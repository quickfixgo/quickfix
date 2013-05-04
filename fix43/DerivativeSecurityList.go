package fix43

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type DerivativeSecurityList struct {
	quickfixgo.Message
}

func (m *DerivativeSecurityList) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *DerivativeSecurityList) SecurityResponseID() (*field.SecurityResponseID, error) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}
func (m *DerivativeSecurityList) SecurityRequestResult() (*field.SecurityRequestResult, error) {
	f := new(field.SecurityRequestResult)
	err := m.Body.Get(f)
	return f, err
}
func (m *DerivativeSecurityList) TotalNumSecurities() (*field.TotalNumSecurities, error) {
	f := new(field.TotalNumSecurities)
	err := m.Body.Get(f)
	return f, err
}
