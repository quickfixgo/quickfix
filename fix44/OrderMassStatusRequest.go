package fix44

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type OrderMassStatusRequest struct {
	quickfixgo.Message
}

func (m *OrderMassStatusRequest) MassStatusReqID() (*field.MassStatusReqID, error) {
	f := new(field.MassStatusReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassStatusRequest) MassStatusReqType() (*field.MassStatusReqType, error) {
	f := new(field.MassStatusReqType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassStatusRequest) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassStatusRequest) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassStatusRequest) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassStatusRequest) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassStatusRequest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
