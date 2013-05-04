package fix50sp2

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type OrderMassActionReport struct {
	quickfixgo.Message
}

func (m *OrderMassActionReport) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) MassActionReportID() (*field.MassActionReportID, error) {
	f := new(field.MassActionReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) MassActionType() (*field.MassActionType, error) {
	f := new(field.MassActionType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) MassActionScope() (*field.MassActionScope, error) {
	f := new(field.MassActionScope)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) MassActionResponse() (*field.MassActionResponse, error) {
	f := new(field.MassActionResponse)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) MassActionRejectReason() (*field.MassActionRejectReason, error) {
	f := new(field.MassActionRejectReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) TotalAffectedOrders() (*field.TotalAffectedOrders, error) {
	f := new(field.TotalAffectedOrders)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
