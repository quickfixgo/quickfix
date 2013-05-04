package fix44

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type CollateralInquiryAck struct {
	quickfixgo.Message
}

func (m *CollateralInquiryAck) CollInquiryID() (*field.CollInquiryID, error) {
	f := new(field.CollInquiryID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) CollInquiryStatus() (*field.CollInquiryStatus, error) {
	f := new(field.CollInquiryStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) CollInquiryResult() (*field.CollInquiryResult, error) {
	f := new(field.CollInquiryResult)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) TotNumReports() (*field.TotNumReports, error) {
	f := new(field.TotNumReports)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) Quantity() (*field.Quantity, error) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) ResponseTransportType() (*field.ResponseTransportType, error) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) ResponseDestination() (*field.ResponseDestination, error) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralInquiryAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
