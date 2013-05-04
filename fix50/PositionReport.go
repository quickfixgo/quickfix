package fix50

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type PositionReport struct {
	quickfixgo.Message
}

func (m *PositionReport) PosMaintRptID() (*field.PosMaintRptID, error) {
	f := new(field.PosMaintRptID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) PosReqID() (*field.PosReqID, error) {
	f := new(field.PosReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) PosReqType() (*field.PosReqType, error) {
	f := new(field.PosReqType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) TotalNumPosReports() (*field.TotalNumPosReports, error) {
	f := new(field.TotalNumPosReports)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) UnsolicitedIndicator() (*field.UnsolicitedIndicator, error) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) PosReqResult() (*field.PosReqResult, error) {
	f := new(field.PosReqResult)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SettlPrice() (*field.SettlPrice, error) {
	f := new(field.SettlPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SettlPriceType() (*field.SettlPriceType, error) {
	f := new(field.SettlPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) PriorSettlPrice() (*field.PriorSettlPrice, error) {
	f := new(field.PriorSettlPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) RegistStatus() (*field.RegistStatus, error) {
	f := new(field.RegistStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) DeliveryDate() (*field.DeliveryDate, error) {
	f := new(field.DeliveryDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) MatchStatus() (*field.MatchStatus, error) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) MessageEventSource() (*field.MessageEventSource, error) {
	f := new(field.MessageEventSource)
	err := m.Body.Get(f)
	return f, err
}
