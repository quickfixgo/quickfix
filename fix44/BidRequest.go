package fix44

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
	clientbidid *field.ClientBidIDField,
	bidrequesttranstype *field.BidRequestTransTypeField,
	totnorelatedsym *field.TotNoRelatedSymField,
	bidtype *field.BidTypeField,
	bidtradetype *field.BidTradeTypeField,
	basispxtype *field.BasisPxTypeField) BidRequestBuilder {
	var builder BidRequestBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(field.NewMsgType("k"))
	builder.Body().Set(clientbidid)
	builder.Body().Set(bidrequesttranstype)
	builder.Body().Set(totnorelatedsym)
	builder.Body().Set(bidtype)
	builder.Body().Set(bidtradetype)
	builder.Body().Set(basispxtype)
	return builder
}

//BidID is a non-required field for BidRequest.
func (m BidRequest) BidID() (*field.BidIDField, errors.MessageRejectError) {
	f := &field.BidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidID reads a BidID from BidRequest.
func (m BidRequest) GetBidID(f *field.BidIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientBidID is a required field for BidRequest.
func (m BidRequest) ClientBidID() (*field.ClientBidIDField, errors.MessageRejectError) {
	f := &field.ClientBidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientBidID reads a ClientBidID from BidRequest.
func (m BidRequest) GetClientBidID(f *field.ClientBidIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidRequestTransType is a required field for BidRequest.
func (m BidRequest) BidRequestTransType() (*field.BidRequestTransTypeField, errors.MessageRejectError) {
	f := &field.BidRequestTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidRequestTransType reads a BidRequestTransType from BidRequest.
func (m BidRequest) GetBidRequestTransType(f *field.BidRequestTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListName is a non-required field for BidRequest.
func (m BidRequest) ListName() (*field.ListNameField, errors.MessageRejectError) {
	f := &field.ListNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListName reads a ListName from BidRequest.
func (m BidRequest) GetListName(f *field.ListNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoRelatedSym is a required field for BidRequest.
func (m BidRequest) TotNoRelatedSym() (*field.TotNoRelatedSymField, errors.MessageRejectError) {
	f := &field.TotNoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoRelatedSym reads a TotNoRelatedSym from BidRequest.
func (m BidRequest) GetTotNoRelatedSym(f *field.TotNoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidType is a required field for BidRequest.
func (m BidRequest) BidType() (*field.BidTypeField, errors.MessageRejectError) {
	f := &field.BidTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidType reads a BidType from BidRequest.
func (m BidRequest) GetBidType(f *field.BidTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NumTickets is a non-required field for BidRequest.
func (m BidRequest) NumTickets() (*field.NumTicketsField, errors.MessageRejectError) {
	f := &field.NumTicketsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNumTickets reads a NumTickets from BidRequest.
func (m BidRequest) GetNumTickets(f *field.NumTicketsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for BidRequest.
func (m BidRequest) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from BidRequest.
func (m BidRequest) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SideValue1 is a non-required field for BidRequest.
func (m BidRequest) SideValue1() (*field.SideValue1Field, errors.MessageRejectError) {
	f := &field.SideValue1Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetSideValue1 reads a SideValue1 from BidRequest.
func (m BidRequest) GetSideValue1(f *field.SideValue1Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SideValue2 is a non-required field for BidRequest.
func (m BidRequest) SideValue2() (*field.SideValue2Field, errors.MessageRejectError) {
	f := &field.SideValue2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetSideValue2 reads a SideValue2 from BidRequest.
func (m BidRequest) GetSideValue2(f *field.SideValue2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoBidDescriptors is a non-required field for BidRequest.
func (m BidRequest) NoBidDescriptors() (*field.NoBidDescriptorsField, errors.MessageRejectError) {
	f := &field.NoBidDescriptorsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoBidDescriptors reads a NoBidDescriptors from BidRequest.
func (m BidRequest) GetNoBidDescriptors(f *field.NoBidDescriptorsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoBidComponents is a non-required field for BidRequest.
func (m BidRequest) NoBidComponents() (*field.NoBidComponentsField, errors.MessageRejectError) {
	f := &field.NoBidComponentsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoBidComponents reads a NoBidComponents from BidRequest.
func (m BidRequest) GetNoBidComponents(f *field.NoBidComponentsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LiquidityIndType is a non-required field for BidRequest.
func (m BidRequest) LiquidityIndType() (*field.LiquidityIndTypeField, errors.MessageRejectError) {
	f := &field.LiquidityIndTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLiquidityIndType reads a LiquidityIndType from BidRequest.
func (m BidRequest) GetLiquidityIndType(f *field.LiquidityIndTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//WtAverageLiquidity is a non-required field for BidRequest.
func (m BidRequest) WtAverageLiquidity() (*field.WtAverageLiquidityField, errors.MessageRejectError) {
	f := &field.WtAverageLiquidityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetWtAverageLiquidity reads a WtAverageLiquidity from BidRequest.
func (m BidRequest) GetWtAverageLiquidity(f *field.WtAverageLiquidityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExchangeForPhysical is a non-required field for BidRequest.
func (m BidRequest) ExchangeForPhysical() (*field.ExchangeForPhysicalField, errors.MessageRejectError) {
	f := &field.ExchangeForPhysicalField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExchangeForPhysical reads a ExchangeForPhysical from BidRequest.
func (m BidRequest) GetExchangeForPhysical(f *field.ExchangeForPhysicalField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OutMainCntryUIndex is a non-required field for BidRequest.
func (m BidRequest) OutMainCntryUIndex() (*field.OutMainCntryUIndexField, errors.MessageRejectError) {
	f := &field.OutMainCntryUIndexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOutMainCntryUIndex reads a OutMainCntryUIndex from BidRequest.
func (m BidRequest) GetOutMainCntryUIndex(f *field.OutMainCntryUIndexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossPercent is a non-required field for BidRequest.
func (m BidRequest) CrossPercent() (*field.CrossPercentField, errors.MessageRejectError) {
	f := &field.CrossPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossPercent reads a CrossPercent from BidRequest.
func (m BidRequest) GetCrossPercent(f *field.CrossPercentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProgRptReqs is a non-required field for BidRequest.
func (m BidRequest) ProgRptReqs() (*field.ProgRptReqsField, errors.MessageRejectError) {
	f := &field.ProgRptReqsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProgRptReqs reads a ProgRptReqs from BidRequest.
func (m BidRequest) GetProgRptReqs(f *field.ProgRptReqsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProgPeriodInterval is a non-required field for BidRequest.
func (m BidRequest) ProgPeriodInterval() (*field.ProgPeriodIntervalField, errors.MessageRejectError) {
	f := &field.ProgPeriodIntervalField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProgPeriodInterval reads a ProgPeriodInterval from BidRequest.
func (m BidRequest) GetProgPeriodInterval(f *field.ProgPeriodIntervalField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IncTaxInd is a non-required field for BidRequest.
func (m BidRequest) IncTaxInd() (*field.IncTaxIndField, errors.MessageRejectError) {
	f := &field.IncTaxIndField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIncTaxInd reads a IncTaxInd from BidRequest.
func (m BidRequest) GetIncTaxInd(f *field.IncTaxIndField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for BidRequest.
func (m BidRequest) ForexReq() (*field.ForexReqField, errors.MessageRejectError) {
	f := &field.ForexReqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from BidRequest.
func (m BidRequest) GetForexReq(f *field.ForexReqField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NumBidders is a non-required field for BidRequest.
func (m BidRequest) NumBidders() (*field.NumBiddersField, errors.MessageRejectError) {
	f := &field.NumBiddersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNumBidders reads a NumBidders from BidRequest.
func (m BidRequest) GetNumBidders(f *field.NumBiddersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for BidRequest.
func (m BidRequest) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from BidRequest.
func (m BidRequest) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidTradeType is a required field for BidRequest.
func (m BidRequest) BidTradeType() (*field.BidTradeTypeField, errors.MessageRejectError) {
	f := &field.BidTradeTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidTradeType reads a BidTradeType from BidRequest.
func (m BidRequest) GetBidTradeType(f *field.BidTradeTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BasisPxType is a required field for BidRequest.
func (m BidRequest) BasisPxType() (*field.BasisPxTypeField, errors.MessageRejectError) {
	f := &field.BasisPxTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBasisPxType reads a BasisPxType from BidRequest.
func (m BidRequest) GetBasisPxType(f *field.BasisPxTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeTime is a non-required field for BidRequest.
func (m BidRequest) StrikeTime() (*field.StrikeTimeField, errors.MessageRejectError) {
	f := &field.StrikeTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeTime reads a StrikeTime from BidRequest.
func (m BidRequest) GetStrikeTime(f *field.StrikeTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for BidRequest.
func (m BidRequest) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from BidRequest.
func (m BidRequest) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for BidRequest.
func (m BidRequest) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from BidRequest.
func (m BidRequest) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for BidRequest.
func (m BidRequest) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from BidRequest.
func (m BidRequest) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
