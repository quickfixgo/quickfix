package fix44

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type BidRequest struct {
	quickfix.Message
}

func (m *BidRequest) ListName() (*field.ListName, error) {
	f := new(field.ListName)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) NoBidDescriptors() (*field.NoBidDescriptors, error) {
	f := new(field.NoBidDescriptors)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) ExchangeForPhysical() (*field.ExchangeForPhysical, error) {
	f := new(field.ExchangeForPhysical)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) ProgPeriodInterval() (*field.ProgPeriodInterval, error) {
	f := new(field.ProgPeriodInterval)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) IncTaxInd() (*field.IncTaxInd, error) {
	f := new(field.IncTaxInd)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) BidID() (*field.BidID, error) {
	f := new(field.BidID)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) NoBidComponents() (*field.NoBidComponents, error) {
	f := new(field.NoBidComponents)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) SideValue1() (*field.SideValue1, error) {
	f := new(field.SideValue1)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) NumBidders() (*field.NumBidders, error) {
	f := new(field.NumBidders)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) StrikeTime() (*field.StrikeTime, error) {
	f := new(field.StrikeTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) ClientBidID() (*field.ClientBidID, error) {
	f := new(field.ClientBidID)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) BidRequestTransType() (*field.BidRequestTransType, error) {
	f := new(field.BidRequestTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) OutMainCntryUIndex() (*field.OutMainCntryUIndex, error) {
	f := new(field.OutMainCntryUIndex)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) BasisPxType() (*field.BasisPxType, error) {
	f := new(field.BasisPxType)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) WtAverageLiquidity() (*field.WtAverageLiquidity, error) {
	f := new(field.WtAverageLiquidity)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) ForexReq() (*field.ForexReq, error) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) BidTradeType() (*field.BidTradeType, error) {
	f := new(field.BidTradeType)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) BidType() (*field.BidType, error) {
	f := new(field.BidType)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) SideValue2() (*field.SideValue2, error) {
	f := new(field.SideValue2)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) LiquidityIndType() (*field.LiquidityIndType, error) {
	f := new(field.LiquidityIndType)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) CrossPercent() (*field.CrossPercent, error) {
	f := new(field.CrossPercent)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) TotNoRelatedSym() (*field.TotNoRelatedSym, error) {
	f := new(field.TotNoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) NumTickets() (*field.NumTickets, error) {
	f := new(field.NumTickets)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidRequest) ProgRptReqs() (*field.ProgRptReqs, error) {
	f := new(field.ProgRptReqs)
	err := m.Body.Get(f)
	return f, err
}
