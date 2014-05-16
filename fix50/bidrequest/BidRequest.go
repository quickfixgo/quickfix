//Package bidrequest msg type = k.
package bidrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a BidRequest wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//BidID is a non-required field for BidRequest.
func (m Message) BidID() (*field.BidIDField, errors.MessageRejectError) {
	f := &field.BidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidID reads a BidID from BidRequest.
func (m Message) GetBidID(f *field.BidIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientBidID is a required field for BidRequest.
func (m Message) ClientBidID() (*field.ClientBidIDField, errors.MessageRejectError) {
	f := &field.ClientBidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientBidID reads a ClientBidID from BidRequest.
func (m Message) GetClientBidID(f *field.ClientBidIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidRequestTransType is a required field for BidRequest.
func (m Message) BidRequestTransType() (*field.BidRequestTransTypeField, errors.MessageRejectError) {
	f := &field.BidRequestTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidRequestTransType reads a BidRequestTransType from BidRequest.
func (m Message) GetBidRequestTransType(f *field.BidRequestTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListName is a non-required field for BidRequest.
func (m Message) ListName() (*field.ListNameField, errors.MessageRejectError) {
	f := &field.ListNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListName reads a ListName from BidRequest.
func (m Message) GetListName(f *field.ListNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoRelatedSym is a required field for BidRequest.
func (m Message) TotNoRelatedSym() (*field.TotNoRelatedSymField, errors.MessageRejectError) {
	f := &field.TotNoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoRelatedSym reads a TotNoRelatedSym from BidRequest.
func (m Message) GetTotNoRelatedSym(f *field.TotNoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidType is a required field for BidRequest.
func (m Message) BidType() (*field.BidTypeField, errors.MessageRejectError) {
	f := &field.BidTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidType reads a BidType from BidRequest.
func (m Message) GetBidType(f *field.BidTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NumTickets is a non-required field for BidRequest.
func (m Message) NumTickets() (*field.NumTicketsField, errors.MessageRejectError) {
	f := &field.NumTicketsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNumTickets reads a NumTickets from BidRequest.
func (m Message) GetNumTickets(f *field.NumTicketsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for BidRequest.
func (m Message) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from BidRequest.
func (m Message) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SideValue1 is a non-required field for BidRequest.
func (m Message) SideValue1() (*field.SideValue1Field, errors.MessageRejectError) {
	f := &field.SideValue1Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetSideValue1 reads a SideValue1 from BidRequest.
func (m Message) GetSideValue1(f *field.SideValue1Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SideValue2 is a non-required field for BidRequest.
func (m Message) SideValue2() (*field.SideValue2Field, errors.MessageRejectError) {
	f := &field.SideValue2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetSideValue2 reads a SideValue2 from BidRequest.
func (m Message) GetSideValue2(f *field.SideValue2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoBidDescriptors is a non-required field for BidRequest.
func (m Message) NoBidDescriptors() (*field.NoBidDescriptorsField, errors.MessageRejectError) {
	f := &field.NoBidDescriptorsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoBidDescriptors reads a NoBidDescriptors from BidRequest.
func (m Message) GetNoBidDescriptors(f *field.NoBidDescriptorsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoBidComponents is a non-required field for BidRequest.
func (m Message) NoBidComponents() (*field.NoBidComponentsField, errors.MessageRejectError) {
	f := &field.NoBidComponentsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoBidComponents reads a NoBidComponents from BidRequest.
func (m Message) GetNoBidComponents(f *field.NoBidComponentsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LiquidityIndType is a non-required field for BidRequest.
func (m Message) LiquidityIndType() (*field.LiquidityIndTypeField, errors.MessageRejectError) {
	f := &field.LiquidityIndTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLiquidityIndType reads a LiquidityIndType from BidRequest.
func (m Message) GetLiquidityIndType(f *field.LiquidityIndTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//WtAverageLiquidity is a non-required field for BidRequest.
func (m Message) WtAverageLiquidity() (*field.WtAverageLiquidityField, errors.MessageRejectError) {
	f := &field.WtAverageLiquidityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetWtAverageLiquidity reads a WtAverageLiquidity from BidRequest.
func (m Message) GetWtAverageLiquidity(f *field.WtAverageLiquidityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExchangeForPhysical is a non-required field for BidRequest.
func (m Message) ExchangeForPhysical() (*field.ExchangeForPhysicalField, errors.MessageRejectError) {
	f := &field.ExchangeForPhysicalField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExchangeForPhysical reads a ExchangeForPhysical from BidRequest.
func (m Message) GetExchangeForPhysical(f *field.ExchangeForPhysicalField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OutMainCntryUIndex is a non-required field for BidRequest.
func (m Message) OutMainCntryUIndex() (*field.OutMainCntryUIndexField, errors.MessageRejectError) {
	f := &field.OutMainCntryUIndexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOutMainCntryUIndex reads a OutMainCntryUIndex from BidRequest.
func (m Message) GetOutMainCntryUIndex(f *field.OutMainCntryUIndexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossPercent is a non-required field for BidRequest.
func (m Message) CrossPercent() (*field.CrossPercentField, errors.MessageRejectError) {
	f := &field.CrossPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossPercent reads a CrossPercent from BidRequest.
func (m Message) GetCrossPercent(f *field.CrossPercentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProgRptReqs is a non-required field for BidRequest.
func (m Message) ProgRptReqs() (*field.ProgRptReqsField, errors.MessageRejectError) {
	f := &field.ProgRptReqsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProgRptReqs reads a ProgRptReqs from BidRequest.
func (m Message) GetProgRptReqs(f *field.ProgRptReqsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProgPeriodInterval is a non-required field for BidRequest.
func (m Message) ProgPeriodInterval() (*field.ProgPeriodIntervalField, errors.MessageRejectError) {
	f := &field.ProgPeriodIntervalField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProgPeriodInterval reads a ProgPeriodInterval from BidRequest.
func (m Message) GetProgPeriodInterval(f *field.ProgPeriodIntervalField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IncTaxInd is a non-required field for BidRequest.
func (m Message) IncTaxInd() (*field.IncTaxIndField, errors.MessageRejectError) {
	f := &field.IncTaxIndField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIncTaxInd reads a IncTaxInd from BidRequest.
func (m Message) GetIncTaxInd(f *field.IncTaxIndField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for BidRequest.
func (m Message) ForexReq() (*field.ForexReqField, errors.MessageRejectError) {
	f := &field.ForexReqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from BidRequest.
func (m Message) GetForexReq(f *field.ForexReqField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NumBidders is a non-required field for BidRequest.
func (m Message) NumBidders() (*field.NumBiddersField, errors.MessageRejectError) {
	f := &field.NumBiddersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNumBidders reads a NumBidders from BidRequest.
func (m Message) GetNumBidders(f *field.NumBiddersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for BidRequest.
func (m Message) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from BidRequest.
func (m Message) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidTradeType is a required field for BidRequest.
func (m Message) BidTradeType() (*field.BidTradeTypeField, errors.MessageRejectError) {
	f := &field.BidTradeTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidTradeType reads a BidTradeType from BidRequest.
func (m Message) GetBidTradeType(f *field.BidTradeTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BasisPxType is a required field for BidRequest.
func (m Message) BasisPxType() (*field.BasisPxTypeField, errors.MessageRejectError) {
	f := &field.BasisPxTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBasisPxType reads a BasisPxType from BidRequest.
func (m Message) GetBasisPxType(f *field.BasisPxTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeTime is a non-required field for BidRequest.
func (m Message) StrikeTime() (*field.StrikeTimeField, errors.MessageRejectError) {
	f := &field.StrikeTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeTime reads a StrikeTime from BidRequest.
func (m Message) GetStrikeTime(f *field.StrikeTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for BidRequest.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from BidRequest.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for BidRequest.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from BidRequest.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for BidRequest.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from BidRequest.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds BidRequest messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for BidRequest.
func Builder(
	clientbidid *field.ClientBidIDField,
	bidrequesttranstype *field.BidRequestTransTypeField,
	totnorelatedsym *field.TotNoRelatedSymField,
	bidtype *field.BidTypeField,
	bidtradetype *field.BidTradeTypeField,
	basispxtype *field.BasisPxTypeField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header().Set(field.NewMsgType("k"))
	builder.Body().Set(clientbidid)
	builder.Body().Set(bidrequesttranstype)
	builder.Body().Set(totnorelatedsym)
	builder.Body().Set(bidtype)
	builder.Body().Set(bidtradetype)
	builder.Body().Set(basispxtype)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX50, "k", r
}
