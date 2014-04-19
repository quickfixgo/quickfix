package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
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
	totalnumsecurities field.TotalNumSecurities,
	bidtype field.BidType,
	tradetype field.TradeType,
	basispxtype field.BasisPxType) BidRequestBuilder {
	var builder BidRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(clientbidid)
	builder.Body.Set(bidrequesttranstype)
	builder.Body.Set(totalnumsecurities)
	builder.Body.Set(bidtype)
	builder.Body.Set(tradetype)
	builder.Body.Set(basispxtype)
	return builder
}

//BidID is a non-required field for BidRequest.
func (m BidRequest) BidID() (field.BidID, errors.MessageRejectError) {
	var f field.BidID
	err := m.Body.Get(&f)
	return f, err
}

//ClientBidID is a required field for BidRequest.
func (m BidRequest) ClientBidID() (field.ClientBidID, errors.MessageRejectError) {
	var f field.ClientBidID
	err := m.Body.Get(&f)
	return f, err
}

//BidRequestTransType is a required field for BidRequest.
func (m BidRequest) BidRequestTransType() (field.BidRequestTransType, errors.MessageRejectError) {
	var f field.BidRequestTransType
	err := m.Body.Get(&f)
	return f, err
}

//ListName is a non-required field for BidRequest.
func (m BidRequest) ListName() (field.ListName, errors.MessageRejectError) {
	var f field.ListName
	err := m.Body.Get(&f)
	return f, err
}

//TotalNumSecurities is a required field for BidRequest.
func (m BidRequest) TotalNumSecurities() (field.TotalNumSecurities, errors.MessageRejectError) {
	var f field.TotalNumSecurities
	err := m.Body.Get(&f)
	return f, err
}

//BidType is a required field for BidRequest.
func (m BidRequest) BidType() (field.BidType, errors.MessageRejectError) {
	var f field.BidType
	err := m.Body.Get(&f)
	return f, err
}

//NumTickets is a non-required field for BidRequest.
func (m BidRequest) NumTickets() (field.NumTickets, errors.MessageRejectError) {
	var f field.NumTickets
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for BidRequest.
func (m BidRequest) Currency() (field.Currency, errors.MessageRejectError) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//SideValue1 is a non-required field for BidRequest.
func (m BidRequest) SideValue1() (field.SideValue1, errors.MessageRejectError) {
	var f field.SideValue1
	err := m.Body.Get(&f)
	return f, err
}

//SideValue2 is a non-required field for BidRequest.
func (m BidRequest) SideValue2() (field.SideValue2, errors.MessageRejectError) {
	var f field.SideValue2
	err := m.Body.Get(&f)
	return f, err
}

//NoBidDescriptors is a non-required field for BidRequest.
func (m BidRequest) NoBidDescriptors() (field.NoBidDescriptors, errors.MessageRejectError) {
	var f field.NoBidDescriptors
	err := m.Body.Get(&f)
	return f, err
}

//NoBidComponents is a non-required field for BidRequest.
func (m BidRequest) NoBidComponents() (field.NoBidComponents, errors.MessageRejectError) {
	var f field.NoBidComponents
	err := m.Body.Get(&f)
	return f, err
}

//LiquidityIndType is a non-required field for BidRequest.
func (m BidRequest) LiquidityIndType() (field.LiquidityIndType, errors.MessageRejectError) {
	var f field.LiquidityIndType
	err := m.Body.Get(&f)
	return f, err
}

//WtAverageLiquidity is a non-required field for BidRequest.
func (m BidRequest) WtAverageLiquidity() (field.WtAverageLiquidity, errors.MessageRejectError) {
	var f field.WtAverageLiquidity
	err := m.Body.Get(&f)
	return f, err
}

//ExchangeForPhysical is a non-required field for BidRequest.
func (m BidRequest) ExchangeForPhysical() (field.ExchangeForPhysical, errors.MessageRejectError) {
	var f field.ExchangeForPhysical
	err := m.Body.Get(&f)
	return f, err
}

//OutMainCntryUIndex is a non-required field for BidRequest.
func (m BidRequest) OutMainCntryUIndex() (field.OutMainCntryUIndex, errors.MessageRejectError) {
	var f field.OutMainCntryUIndex
	err := m.Body.Get(&f)
	return f, err
}

//CrossPercent is a non-required field for BidRequest.
func (m BidRequest) CrossPercent() (field.CrossPercent, errors.MessageRejectError) {
	var f field.CrossPercent
	err := m.Body.Get(&f)
	return f, err
}

//ProgRptReqs is a non-required field for BidRequest.
func (m BidRequest) ProgRptReqs() (field.ProgRptReqs, errors.MessageRejectError) {
	var f field.ProgRptReqs
	err := m.Body.Get(&f)
	return f, err
}

//ProgPeriodInterval is a non-required field for BidRequest.
func (m BidRequest) ProgPeriodInterval() (field.ProgPeriodInterval, errors.MessageRejectError) {
	var f field.ProgPeriodInterval
	err := m.Body.Get(&f)
	return f, err
}

//IncTaxInd is a non-required field for BidRequest.
func (m BidRequest) IncTaxInd() (field.IncTaxInd, errors.MessageRejectError) {
	var f field.IncTaxInd
	err := m.Body.Get(&f)
	return f, err
}

//ForexReq is a non-required field for BidRequest.
func (m BidRequest) ForexReq() (field.ForexReq, errors.MessageRejectError) {
	var f field.ForexReq
	err := m.Body.Get(&f)
	return f, err
}

//NumBidders is a non-required field for BidRequest.
func (m BidRequest) NumBidders() (field.NumBidders, errors.MessageRejectError) {
	var f field.NumBidders
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for BidRequest.
func (m BidRequest) TradeDate() (field.TradeDate, errors.MessageRejectError) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//TradeType is a required field for BidRequest.
func (m BidRequest) TradeType() (field.TradeType, errors.MessageRejectError) {
	var f field.TradeType
	err := m.Body.Get(&f)
	return f, err
}

//BasisPxType is a required field for BidRequest.
func (m BidRequest) BasisPxType() (field.BasisPxType, errors.MessageRejectError) {
	var f field.BasisPxType
	err := m.Body.Get(&f)
	return f, err
}

//StrikeTime is a non-required field for BidRequest.
func (m BidRequest) StrikeTime() (field.StrikeTime, errors.MessageRejectError) {
	var f field.StrikeTime
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for BidRequest.
func (m BidRequest) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for BidRequest.
func (m BidRequest) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for BidRequest.
func (m BidRequest) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
