package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIX43))
	builder.Header.Set(field.BuildMsgType("k"))
	builder.Body.Set(clientbidid)
	builder.Body.Set(bidrequesttranstype)
	builder.Body.Set(totalnumsecurities)
	builder.Body.Set(bidtype)
	builder.Body.Set(tradetype)
	builder.Body.Set(basispxtype)
	return builder
}

//BidID is a non-required field for BidRequest.
func (m BidRequest) BidID() (*field.BidID, errors.MessageRejectError) {
	f := new(field.BidID)
	err := m.Body.Get(f)
	return f, err
}

//GetBidID reads a BidID from BidRequest.
func (m BidRequest) GetBidID(f *field.BidID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientBidID is a required field for BidRequest.
func (m BidRequest) ClientBidID() (*field.ClientBidID, errors.MessageRejectError) {
	f := new(field.ClientBidID)
	err := m.Body.Get(f)
	return f, err
}

//GetClientBidID reads a ClientBidID from BidRequest.
func (m BidRequest) GetClientBidID(f *field.ClientBidID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidRequestTransType is a required field for BidRequest.
func (m BidRequest) BidRequestTransType() (*field.BidRequestTransType, errors.MessageRejectError) {
	f := new(field.BidRequestTransType)
	err := m.Body.Get(f)
	return f, err
}

//GetBidRequestTransType reads a BidRequestTransType from BidRequest.
func (m BidRequest) GetBidRequestTransType(f *field.BidRequestTransType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListName is a non-required field for BidRequest.
func (m BidRequest) ListName() (*field.ListName, errors.MessageRejectError) {
	f := new(field.ListName)
	err := m.Body.Get(f)
	return f, err
}

//GetListName reads a ListName from BidRequest.
func (m BidRequest) GetListName(f *field.ListName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalNumSecurities is a required field for BidRequest.
func (m BidRequest) TotalNumSecurities() (*field.TotalNumSecurities, errors.MessageRejectError) {
	f := new(field.TotalNumSecurities)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalNumSecurities reads a TotalNumSecurities from BidRequest.
func (m BidRequest) GetTotalNumSecurities(f *field.TotalNumSecurities) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidType is a required field for BidRequest.
func (m BidRequest) BidType() (*field.BidType, errors.MessageRejectError) {
	f := new(field.BidType)
	err := m.Body.Get(f)
	return f, err
}

//GetBidType reads a BidType from BidRequest.
func (m BidRequest) GetBidType(f *field.BidType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NumTickets is a non-required field for BidRequest.
func (m BidRequest) NumTickets() (*field.NumTickets, errors.MessageRejectError) {
	f := new(field.NumTickets)
	err := m.Body.Get(f)
	return f, err
}

//GetNumTickets reads a NumTickets from BidRequest.
func (m BidRequest) GetNumTickets(f *field.NumTickets) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for BidRequest.
func (m BidRequest) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from BidRequest.
func (m BidRequest) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SideValue1 is a non-required field for BidRequest.
func (m BidRequest) SideValue1() (*field.SideValue1, errors.MessageRejectError) {
	f := new(field.SideValue1)
	err := m.Body.Get(f)
	return f, err
}

//GetSideValue1 reads a SideValue1 from BidRequest.
func (m BidRequest) GetSideValue1(f *field.SideValue1) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SideValue2 is a non-required field for BidRequest.
func (m BidRequest) SideValue2() (*field.SideValue2, errors.MessageRejectError) {
	f := new(field.SideValue2)
	err := m.Body.Get(f)
	return f, err
}

//GetSideValue2 reads a SideValue2 from BidRequest.
func (m BidRequest) GetSideValue2(f *field.SideValue2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoBidDescriptors is a non-required field for BidRequest.
func (m BidRequest) NoBidDescriptors() (*field.NoBidDescriptors, errors.MessageRejectError) {
	f := new(field.NoBidDescriptors)
	err := m.Body.Get(f)
	return f, err
}

//GetNoBidDescriptors reads a NoBidDescriptors from BidRequest.
func (m BidRequest) GetNoBidDescriptors(f *field.NoBidDescriptors) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoBidComponents is a non-required field for BidRequest.
func (m BidRequest) NoBidComponents() (*field.NoBidComponents, errors.MessageRejectError) {
	f := new(field.NoBidComponents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoBidComponents reads a NoBidComponents from BidRequest.
func (m BidRequest) GetNoBidComponents(f *field.NoBidComponents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LiquidityIndType is a non-required field for BidRequest.
func (m BidRequest) LiquidityIndType() (*field.LiquidityIndType, errors.MessageRejectError) {
	f := new(field.LiquidityIndType)
	err := m.Body.Get(f)
	return f, err
}

//GetLiquidityIndType reads a LiquidityIndType from BidRequest.
func (m BidRequest) GetLiquidityIndType(f *field.LiquidityIndType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//WtAverageLiquidity is a non-required field for BidRequest.
func (m BidRequest) WtAverageLiquidity() (*field.WtAverageLiquidity, errors.MessageRejectError) {
	f := new(field.WtAverageLiquidity)
	err := m.Body.Get(f)
	return f, err
}

//GetWtAverageLiquidity reads a WtAverageLiquidity from BidRequest.
func (m BidRequest) GetWtAverageLiquidity(f *field.WtAverageLiquidity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExchangeForPhysical is a non-required field for BidRequest.
func (m BidRequest) ExchangeForPhysical() (*field.ExchangeForPhysical, errors.MessageRejectError) {
	f := new(field.ExchangeForPhysical)
	err := m.Body.Get(f)
	return f, err
}

//GetExchangeForPhysical reads a ExchangeForPhysical from BidRequest.
func (m BidRequest) GetExchangeForPhysical(f *field.ExchangeForPhysical) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OutMainCntryUIndex is a non-required field for BidRequest.
func (m BidRequest) OutMainCntryUIndex() (*field.OutMainCntryUIndex, errors.MessageRejectError) {
	f := new(field.OutMainCntryUIndex)
	err := m.Body.Get(f)
	return f, err
}

//GetOutMainCntryUIndex reads a OutMainCntryUIndex from BidRequest.
func (m BidRequest) GetOutMainCntryUIndex(f *field.OutMainCntryUIndex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossPercent is a non-required field for BidRequest.
func (m BidRequest) CrossPercent() (*field.CrossPercent, errors.MessageRejectError) {
	f := new(field.CrossPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetCrossPercent reads a CrossPercent from BidRequest.
func (m BidRequest) GetCrossPercent(f *field.CrossPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProgRptReqs is a non-required field for BidRequest.
func (m BidRequest) ProgRptReqs() (*field.ProgRptReqs, errors.MessageRejectError) {
	f := new(field.ProgRptReqs)
	err := m.Body.Get(f)
	return f, err
}

//GetProgRptReqs reads a ProgRptReqs from BidRequest.
func (m BidRequest) GetProgRptReqs(f *field.ProgRptReqs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProgPeriodInterval is a non-required field for BidRequest.
func (m BidRequest) ProgPeriodInterval() (*field.ProgPeriodInterval, errors.MessageRejectError) {
	f := new(field.ProgPeriodInterval)
	err := m.Body.Get(f)
	return f, err
}

//GetProgPeriodInterval reads a ProgPeriodInterval from BidRequest.
func (m BidRequest) GetProgPeriodInterval(f *field.ProgPeriodInterval) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IncTaxInd is a non-required field for BidRequest.
func (m BidRequest) IncTaxInd() (*field.IncTaxInd, errors.MessageRejectError) {
	f := new(field.IncTaxInd)
	err := m.Body.Get(f)
	return f, err
}

//GetIncTaxInd reads a IncTaxInd from BidRequest.
func (m BidRequest) GetIncTaxInd(f *field.IncTaxInd) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for BidRequest.
func (m BidRequest) ForexReq() (*field.ForexReq, errors.MessageRejectError) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from BidRequest.
func (m BidRequest) GetForexReq(f *field.ForexReq) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NumBidders is a non-required field for BidRequest.
func (m BidRequest) NumBidders() (*field.NumBidders, errors.MessageRejectError) {
	f := new(field.NumBidders)
	err := m.Body.Get(f)
	return f, err
}

//GetNumBidders reads a NumBidders from BidRequest.
func (m BidRequest) GetNumBidders(f *field.NumBidders) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for BidRequest.
func (m BidRequest) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from BidRequest.
func (m BidRequest) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeType is a required field for BidRequest.
func (m BidRequest) TradeType() (*field.TradeType, errors.MessageRejectError) {
	f := new(field.TradeType)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeType reads a TradeType from BidRequest.
func (m BidRequest) GetTradeType(f *field.TradeType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BasisPxType is a required field for BidRequest.
func (m BidRequest) BasisPxType() (*field.BasisPxType, errors.MessageRejectError) {
	f := new(field.BasisPxType)
	err := m.Body.Get(f)
	return f, err
}

//GetBasisPxType reads a BasisPxType from BidRequest.
func (m BidRequest) GetBasisPxType(f *field.BasisPxType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeTime is a non-required field for BidRequest.
func (m BidRequest) StrikeTime() (*field.StrikeTime, errors.MessageRejectError) {
	f := new(field.StrikeTime)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeTime reads a StrikeTime from BidRequest.
func (m BidRequest) GetStrikeTime(f *field.StrikeTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for BidRequest.
func (m BidRequest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from BidRequest.
func (m BidRequest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for BidRequest.
func (m BidRequest) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from BidRequest.
func (m BidRequest) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for BidRequest.
func (m BidRequest) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from BidRequest.
func (m BidRequest) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
