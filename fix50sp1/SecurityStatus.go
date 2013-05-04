package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type SecurityStatus struct {
	quickfixgo.Message
}

func (m *SecurityStatus) SecurityStatusReqID() (*field.SecurityStatusReqID, error) {
	f := new(field.SecurityStatusReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) UnsolicitedIndicator() (*field.UnsolicitedIndicator, error) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) SecurityTradingStatus() (*field.SecurityTradingStatus, error) {
	f := new(field.SecurityTradingStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) FinancialStatus() (*field.FinancialStatus, error) {
	f := new(field.FinancialStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) CorporateAction() (*field.CorporateAction, error) {
	f := new(field.CorporateAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) HaltReasonChar() (*field.HaltReasonChar, error) {
	f := new(field.HaltReasonChar)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) InViewOfCommon() (*field.InViewOfCommon, error) {
	f := new(field.InViewOfCommon)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) DueToRelated() (*field.DueToRelated, error) {
	f := new(field.DueToRelated)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) BuyVolume() (*field.BuyVolume, error) {
	f := new(field.BuyVolume)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) SellVolume() (*field.SellVolume, error) {
	f := new(field.SellVolume)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) HighPx() (*field.HighPx, error) {
	f := new(field.HighPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) LowPx() (*field.LowPx, error) {
	f := new(field.LowPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) Adjustment() (*field.Adjustment, error) {
	f := new(field.Adjustment)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) FirstPx() (*field.FirstPx, error) {
	f := new(field.FirstPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) SecurityTradingEvent() (*field.SecurityTradingEvent, error) {
	f := new(field.SecurityTradingEvent)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) MDBookType() (*field.MDBookType, error) {
	f := new(field.MDBookType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) MarketDepth() (*field.MarketDepth, error) {
	f := new(field.MarketDepth)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
