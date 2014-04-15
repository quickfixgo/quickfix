package fix50sp2

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

//CreateBidRequestBuilder returns an initialized BidRequestBuilder with specified required fields.
func CreateBidRequestBuilder(
	clientbidid field.ClientBidID,
	bidrequesttranstype field.BidRequestTransType,
	totnorelatedsym field.TotNoRelatedSym,
	bidtype field.BidType,
	bidtradetype field.BidTradeType,
	basispxtype field.BasisPxType) BidRequestBuilder {
	var builder BidRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(clientbidid)
	builder.Body.Set(bidrequesttranstype)
	builder.Body.Set(totnorelatedsym)
	builder.Body.Set(bidtype)
	builder.Body.Set(bidtradetype)
	builder.Body.Set(basispxtype)
	return builder
}

//BidID is a non-required field for BidRequest.
func (m BidRequest) BidID() (field.BidID, error) {
	var f field.BidID
	err := m.Body.Get(&f)
	return f, err
}

//ClientBidID is a required field for BidRequest.
func (m BidRequest) ClientBidID() (field.ClientBidID, error) {
	var f field.ClientBidID
	err := m.Body.Get(&f)
	return f, err
}

//BidRequestTransType is a required field for BidRequest.
func (m BidRequest) BidRequestTransType() (field.BidRequestTransType, error) {
	var f field.BidRequestTransType
	err := m.Body.Get(&f)
	return f, err
}

//ListName is a non-required field for BidRequest.
func (m BidRequest) ListName() (field.ListName, error) {
	var f field.ListName
	err := m.Body.Get(&f)
	return f, err
}

//TotNoRelatedSym is a required field for BidRequest.
func (m BidRequest) TotNoRelatedSym() (field.TotNoRelatedSym, error) {
	var f field.TotNoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//BidType is a required field for BidRequest.
func (m BidRequest) BidType() (field.BidType, error) {
	var f field.BidType
	err := m.Body.Get(&f)
	return f, err
}

//NumTickets is a non-required field for BidRequest.
func (m BidRequest) NumTickets() (field.NumTickets, error) {
	var f field.NumTickets
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for BidRequest.
func (m BidRequest) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//SideValue1 is a non-required field for BidRequest.
func (m BidRequest) SideValue1() (field.SideValue1, error) {
	var f field.SideValue1
	err := m.Body.Get(&f)
	return f, err
}

//SideValue2 is a non-required field for BidRequest.
func (m BidRequest) SideValue2() (field.SideValue2, error) {
	var f field.SideValue2
	err := m.Body.Get(&f)
	return f, err
}

//NoBidDescriptors is a non-required field for BidRequest.
func (m BidRequest) NoBidDescriptors() (field.NoBidDescriptors, error) {
	var f field.NoBidDescriptors
	err := m.Body.Get(&f)
	return f, err
}

//NoBidComponents is a non-required field for BidRequest.
func (m BidRequest) NoBidComponents() (field.NoBidComponents, error) {
	var f field.NoBidComponents
	err := m.Body.Get(&f)
	return f, err
}

//LiquidityIndType is a non-required field for BidRequest.
func (m BidRequest) LiquidityIndType() (field.LiquidityIndType, error) {
	var f field.LiquidityIndType
	err := m.Body.Get(&f)
	return f, err
}

//WtAverageLiquidity is a non-required field for BidRequest.
func (m BidRequest) WtAverageLiquidity() (field.WtAverageLiquidity, error) {
	var f field.WtAverageLiquidity
	err := m.Body.Get(&f)
	return f, err
}

//ExchangeForPhysical is a non-required field for BidRequest.
func (m BidRequest) ExchangeForPhysical() (field.ExchangeForPhysical, error) {
	var f field.ExchangeForPhysical
	err := m.Body.Get(&f)
	return f, err
}

//OutMainCntryUIndex is a non-required field for BidRequest.
func (m BidRequest) OutMainCntryUIndex() (field.OutMainCntryUIndex, error) {
	var f field.OutMainCntryUIndex
	err := m.Body.Get(&f)
	return f, err
}

//CrossPercent is a non-required field for BidRequest.
func (m BidRequest) CrossPercent() (field.CrossPercent, error) {
	var f field.CrossPercent
	err := m.Body.Get(&f)
	return f, err
}

//ProgRptReqs is a non-required field for BidRequest.
func (m BidRequest) ProgRptReqs() (field.ProgRptReqs, error) {
	var f field.ProgRptReqs
	err := m.Body.Get(&f)
	return f, err
}

//ProgPeriodInterval is a non-required field for BidRequest.
func (m BidRequest) ProgPeriodInterval() (field.ProgPeriodInterval, error) {
	var f field.ProgPeriodInterval
	err := m.Body.Get(&f)
	return f, err
}

//IncTaxInd is a non-required field for BidRequest.
func (m BidRequest) IncTaxInd() (field.IncTaxInd, error) {
	var f field.IncTaxInd
	err := m.Body.Get(&f)
	return f, err
}

//ForexReq is a non-required field for BidRequest.
func (m BidRequest) ForexReq() (field.ForexReq, error) {
	var f field.ForexReq
	err := m.Body.Get(&f)
	return f, err
}

//NumBidders is a non-required field for BidRequest.
func (m BidRequest) NumBidders() (field.NumBidders, error) {
	var f field.NumBidders
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for BidRequest.
func (m BidRequest) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//BidTradeType is a required field for BidRequest.
func (m BidRequest) BidTradeType() (field.BidTradeType, error) {
	var f field.BidTradeType
	err := m.Body.Get(&f)
	return f, err
}

//BasisPxType is a required field for BidRequest.
func (m BidRequest) BasisPxType() (field.BasisPxType, error) {
	var f field.BasisPxType
	err := m.Body.Get(&f)
	return f, err
}

//StrikeTime is a non-required field for BidRequest.
func (m BidRequest) StrikeTime() (field.StrikeTime, error) {
	var f field.StrikeTime
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for BidRequest.
func (m BidRequest) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for BidRequest.
func (m BidRequest) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for BidRequest.
func (m BidRequest) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
