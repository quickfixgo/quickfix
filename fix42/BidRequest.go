package fix42

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//BidRequest msg type = k.
type BidRequest struct {
	message.Message
}

//BidRequestBuilder builds BidRequest messages.
type BidRequestBuilder struct {
	message.MessageBuilder
}

//NewBidRequestBuilder returns an initialized BidRequestBuilder with specified required fields.
func NewBidRequestBuilder(
	clientbidid field.ClientBidID,
	bidrequesttranstype field.BidRequestTransType,
	totalnumsecurities field.TotalNumSecurities,
	bidtype field.BidType,
	tradetype field.TradeType,
	basispxtype field.BasisPxType) *BidRequestBuilder {
	builder := new(BidRequestBuilder)
	builder.Body.Set(clientbidid)
	builder.Body.Set(bidrequesttranstype)
	builder.Body.Set(totalnumsecurities)
	builder.Body.Set(bidtype)
	builder.Body.Set(tradetype)
	builder.Body.Set(basispxtype)
	return builder
}

//BidID is a non-required field for BidRequest.
func (m *BidRequest) BidID() (*field.BidID, error) {
	f := new(field.BidID)
	err := m.Body.Get(f)
	return f, err
}

//ClientBidID is a required field for BidRequest.
func (m *BidRequest) ClientBidID() (*field.ClientBidID, error) {
	f := new(field.ClientBidID)
	err := m.Body.Get(f)
	return f, err
}

//BidRequestTransType is a required field for BidRequest.
func (m *BidRequest) BidRequestTransType() (*field.BidRequestTransType, error) {
	f := new(field.BidRequestTransType)
	err := m.Body.Get(f)
	return f, err
}

//ListName is a non-required field for BidRequest.
func (m *BidRequest) ListName() (*field.ListName, error) {
	f := new(field.ListName)
	err := m.Body.Get(f)
	return f, err
}

//TotalNumSecurities is a required field for BidRequest.
func (m *BidRequest) TotalNumSecurities() (*field.TotalNumSecurities, error) {
	f := new(field.TotalNumSecurities)
	err := m.Body.Get(f)
	return f, err
}

//BidType is a required field for BidRequest.
func (m *BidRequest) BidType() (*field.BidType, error) {
	f := new(field.BidType)
	err := m.Body.Get(f)
	return f, err
}

//NumTickets is a non-required field for BidRequest.
func (m *BidRequest) NumTickets() (*field.NumTickets, error) {
	f := new(field.NumTickets)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for BidRequest.
func (m *BidRequest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//SideValue1 is a non-required field for BidRequest.
func (m *BidRequest) SideValue1() (*field.SideValue1, error) {
	f := new(field.SideValue1)
	err := m.Body.Get(f)
	return f, err
}

//SideValue2 is a non-required field for BidRequest.
func (m *BidRequest) SideValue2() (*field.SideValue2, error) {
	f := new(field.SideValue2)
	err := m.Body.Get(f)
	return f, err
}

//NoBidDescriptors is a non-required field for BidRequest.
func (m *BidRequest) NoBidDescriptors() (*field.NoBidDescriptors, error) {
	f := new(field.NoBidDescriptors)
	err := m.Body.Get(f)
	return f, err
}

//NoBidComponents is a non-required field for BidRequest.
func (m *BidRequest) NoBidComponents() (*field.NoBidComponents, error) {
	f := new(field.NoBidComponents)
	err := m.Body.Get(f)
	return f, err
}

//LiquidityIndType is a non-required field for BidRequest.
func (m *BidRequest) LiquidityIndType() (*field.LiquidityIndType, error) {
	f := new(field.LiquidityIndType)
	err := m.Body.Get(f)
	return f, err
}

//WtAverageLiquidity is a non-required field for BidRequest.
func (m *BidRequest) WtAverageLiquidity() (*field.WtAverageLiquidity, error) {
	f := new(field.WtAverageLiquidity)
	err := m.Body.Get(f)
	return f, err
}

//ExchangeForPhysical is a non-required field for BidRequest.
func (m *BidRequest) ExchangeForPhysical() (*field.ExchangeForPhysical, error) {
	f := new(field.ExchangeForPhysical)
	err := m.Body.Get(f)
	return f, err
}

//OutMainCntryUIndex is a non-required field for BidRequest.
func (m *BidRequest) OutMainCntryUIndex() (*field.OutMainCntryUIndex, error) {
	f := new(field.OutMainCntryUIndex)
	err := m.Body.Get(f)
	return f, err
}

//CrossPercent is a non-required field for BidRequest.
func (m *BidRequest) CrossPercent() (*field.CrossPercent, error) {
	f := new(field.CrossPercent)
	err := m.Body.Get(f)
	return f, err
}

//ProgRptReqs is a non-required field for BidRequest.
func (m *BidRequest) ProgRptReqs() (*field.ProgRptReqs, error) {
	f := new(field.ProgRptReqs)
	err := m.Body.Get(f)
	return f, err
}

//ProgPeriodInterval is a non-required field for BidRequest.
func (m *BidRequest) ProgPeriodInterval() (*field.ProgPeriodInterval, error) {
	f := new(field.ProgPeriodInterval)
	err := m.Body.Get(f)
	return f, err
}

//IncTaxInd is a non-required field for BidRequest.
func (m *BidRequest) IncTaxInd() (*field.IncTaxInd, error) {
	f := new(field.IncTaxInd)
	err := m.Body.Get(f)
	return f, err
}

//ForexReq is a non-required field for BidRequest.
func (m *BidRequest) ForexReq() (*field.ForexReq, error) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}

//NumBidders is a non-required field for BidRequest.
func (m *BidRequest) NumBidders() (*field.NumBidders, error) {
	f := new(field.NumBidders)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a non-required field for BidRequest.
func (m *BidRequest) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//TradeType is a required field for BidRequest.
func (m *BidRequest) TradeType() (*field.TradeType, error) {
	f := new(field.TradeType)
	err := m.Body.Get(f)
	return f, err
}

//BasisPxType is a required field for BidRequest.
func (m *BidRequest) BasisPxType() (*field.BasisPxType, error) {
	f := new(field.BasisPxType)
	err := m.Body.Get(f)
	return f, err
}

//StrikeTime is a non-required field for BidRequest.
func (m *BidRequest) StrikeTime() (*field.StrikeTime, error) {
	f := new(field.StrikeTime)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for BidRequest.
func (m *BidRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for BidRequest.
func (m *BidRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for BidRequest.
func (m *BidRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
