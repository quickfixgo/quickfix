package fix44

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type CollateralInquiry struct {
	quickfixgo.Message
}

func (m *CollateralInquiry) CollInquiryID() (*field.CollInquiryID, error) {
	f := new(field.CollInquiryID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) ResponseTransportType() (*field.ResponseTransportType, error) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) ResponseDestination() (*field.ResponseDestination, error) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) Quantity() (*field.Quantity, error) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) MarginExcess() (*field.MarginExcess, error) {
	f := new(field.MarginExcess)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) TotalNetValue() (*field.TotalNetValue, error) {
	f := new(field.TotalNetValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) CashOutstanding() (*field.CashOutstanding, error) {
	f := new(field.CashOutstanding)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) AccruedInterestAmt() (*field.AccruedInterestAmt, error) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) EndAccruedInterestAmt() (*field.EndAccruedInterestAmt, error) {
	f := new(field.EndAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) StartCash() (*field.StartCash, error) {
	f := new(field.StartCash)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) EndCash() (*field.EndCash, error) {
	f := new(field.EndCash)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiry) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
