//Package marketdefinition msg type = BU.
package marketdefinition

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a MarketDefinition wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//MarketReportID is a required field for MarketDefinition.
func (m Message) MarketReportID() (*field.MarketReportIDField, quickfix.MessageRejectError) {
	f := &field.MarketReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketReportID reads a MarketReportID from MarketDefinition.
func (m Message) GetMarketReportID(f *field.MarketReportIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarketReqID is a non-required field for MarketDefinition.
func (m Message) MarketReqID() (*field.MarketReqIDField, quickfix.MessageRejectError) {
	f := &field.MarketReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketReqID reads a MarketReqID from MarketDefinition.
func (m Message) GetMarketReqID(f *field.MarketReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a required field for MarketDefinition.
func (m Message) MarketID() (*field.MarketIDField, quickfix.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from MarketDefinition.
func (m Message) GetMarketID(f *field.MarketIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for MarketDefinition.
func (m Message) MarketSegmentID() (*field.MarketSegmentIDField, quickfix.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from MarketDefinition.
func (m Message) GetMarketSegmentID(f *field.MarketSegmentIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentDesc is a non-required field for MarketDefinition.
func (m Message) MarketSegmentDesc() (*field.MarketSegmentDescField, quickfix.MessageRejectError) {
	f := &field.MarketSegmentDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentDesc reads a MarketSegmentDesc from MarketDefinition.
func (m Message) GetMarketSegmentDesc(f *field.MarketSegmentDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedMktSegmDescLen is a non-required field for MarketDefinition.
func (m Message) EncodedMktSegmDescLen() (*field.EncodedMktSegmDescLenField, quickfix.MessageRejectError) {
	f := &field.EncodedMktSegmDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedMktSegmDescLen reads a EncodedMktSegmDescLen from MarketDefinition.
func (m Message) GetEncodedMktSegmDescLen(f *field.EncodedMktSegmDescLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedMktSegmDesc is a non-required field for MarketDefinition.
func (m Message) EncodedMktSegmDesc() (*field.EncodedMktSegmDescField, quickfix.MessageRejectError) {
	f := &field.EncodedMktSegmDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedMktSegmDesc reads a EncodedMktSegmDesc from MarketDefinition.
func (m Message) GetEncodedMktSegmDesc(f *field.EncodedMktSegmDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ParentMktSegmID is a non-required field for MarketDefinition.
func (m Message) ParentMktSegmID() (*field.ParentMktSegmIDField, quickfix.MessageRejectError) {
	f := &field.ParentMktSegmIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetParentMktSegmID reads a ParentMktSegmID from MarketDefinition.
func (m Message) GetParentMktSegmID(f *field.ParentMktSegmIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for MarketDefinition.
func (m Message) Currency() (*field.CurrencyField, quickfix.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from MarketDefinition.
func (m Message) GetCurrency(f *field.CurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoTickRules is a non-required field for MarketDefinition.
func (m Message) NoTickRules() (*field.NoTickRulesField, quickfix.MessageRejectError) {
	f := &field.NoTickRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTickRules reads a NoTickRules from MarketDefinition.
func (m Message) GetNoTickRules(f *field.NoTickRulesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoLotTypeRules is a non-required field for MarketDefinition.
func (m Message) NoLotTypeRules() (*field.NoLotTypeRulesField, quickfix.MessageRejectError) {
	f := &field.NoLotTypeRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLotTypeRules reads a NoLotTypeRules from MarketDefinition.
func (m Message) GetNoLotTypeRules(f *field.NoLotTypeRulesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceLimitType is a non-required field for MarketDefinition.
func (m Message) PriceLimitType() (*field.PriceLimitTypeField, quickfix.MessageRejectError) {
	f := &field.PriceLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceLimitType reads a PriceLimitType from MarketDefinition.
func (m Message) GetPriceLimitType(f *field.PriceLimitTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LowLimitPrice is a non-required field for MarketDefinition.
func (m Message) LowLimitPrice() (*field.LowLimitPriceField, quickfix.MessageRejectError) {
	f := &field.LowLimitPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLowLimitPrice reads a LowLimitPrice from MarketDefinition.
func (m Message) GetLowLimitPrice(f *field.LowLimitPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//HighLimitPrice is a non-required field for MarketDefinition.
func (m Message) HighLimitPrice() (*field.HighLimitPriceField, quickfix.MessageRejectError) {
	f := &field.HighLimitPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHighLimitPrice reads a HighLimitPrice from MarketDefinition.
func (m Message) GetHighLimitPrice(f *field.HighLimitPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingReferencePrice is a non-required field for MarketDefinition.
func (m Message) TradingReferencePrice() (*field.TradingReferencePriceField, quickfix.MessageRejectError) {
	f := &field.TradingReferencePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingReferencePrice reads a TradingReferencePrice from MarketDefinition.
func (m Message) GetTradingReferencePrice(f *field.TradingReferencePriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExpirationCycle is a non-required field for MarketDefinition.
func (m Message) ExpirationCycle() (*field.ExpirationCycleField, quickfix.MessageRejectError) {
	f := &field.ExpirationCycleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpirationCycle reads a ExpirationCycle from MarketDefinition.
func (m Message) GetExpirationCycle(f *field.ExpirationCycleField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MinTradeVol is a non-required field for MarketDefinition.
func (m Message) MinTradeVol() (*field.MinTradeVolField, quickfix.MessageRejectError) {
	f := &field.MinTradeVolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinTradeVol reads a MinTradeVol from MarketDefinition.
func (m Message) GetMinTradeVol(f *field.MinTradeVolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaxTradeVol is a non-required field for MarketDefinition.
func (m Message) MaxTradeVol() (*field.MaxTradeVolField, quickfix.MessageRejectError) {
	f := &field.MaxTradeVolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxTradeVol reads a MaxTradeVol from MarketDefinition.
func (m Message) GetMaxTradeVol(f *field.MaxTradeVolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceVariation is a non-required field for MarketDefinition.
func (m Message) MaxPriceVariation() (*field.MaxPriceVariationField, quickfix.MessageRejectError) {
	f := &field.MaxPriceVariationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceVariation reads a MaxPriceVariation from MarketDefinition.
func (m Message) GetMaxPriceVariation(f *field.MaxPriceVariationField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ImpliedMarketIndicator is a non-required field for MarketDefinition.
func (m Message) ImpliedMarketIndicator() (*field.ImpliedMarketIndicatorField, quickfix.MessageRejectError) {
	f := &field.ImpliedMarketIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetImpliedMarketIndicator reads a ImpliedMarketIndicator from MarketDefinition.
func (m Message) GetImpliedMarketIndicator(f *field.ImpliedMarketIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingCurrency is a non-required field for MarketDefinition.
func (m Message) TradingCurrency() (*field.TradingCurrencyField, quickfix.MessageRejectError) {
	f := &field.TradingCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingCurrency reads a TradingCurrency from MarketDefinition.
func (m Message) GetTradingCurrency(f *field.TradingCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RoundLot is a non-required field for MarketDefinition.
func (m Message) RoundLot() (*field.RoundLotField, quickfix.MessageRejectError) {
	f := &field.RoundLotField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundLot reads a RoundLot from MarketDefinition.
func (m Message) GetRoundLot(f *field.RoundLotField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegModel is a non-required field for MarketDefinition.
func (m Message) MultilegModel() (*field.MultilegModelField, quickfix.MessageRejectError) {
	f := &field.MultilegModelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegModel reads a MultilegModel from MarketDefinition.
func (m Message) GetMultilegModel(f *field.MultilegModelField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegPriceMethod is a non-required field for MarketDefinition.
func (m Message) MultilegPriceMethod() (*field.MultilegPriceMethodField, quickfix.MessageRejectError) {
	f := &field.MultilegPriceMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegPriceMethod reads a MultilegPriceMethod from MarketDefinition.
func (m Message) GetMultilegPriceMethod(f *field.MultilegPriceMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for MarketDefinition.
func (m Message) PriceType() (*field.PriceTypeField, quickfix.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from MarketDefinition.
func (m Message) GetPriceType(f *field.PriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrdTypeRules is a non-required field for MarketDefinition.
func (m Message) NoOrdTypeRules() (*field.NoOrdTypeRulesField, quickfix.MessageRejectError) {
	f := &field.NoOrdTypeRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrdTypeRules reads a NoOrdTypeRules from MarketDefinition.
func (m Message) GetNoOrdTypeRules(f *field.NoOrdTypeRulesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoTimeInForceRules is a non-required field for MarketDefinition.
func (m Message) NoTimeInForceRules() (*field.NoTimeInForceRulesField, quickfix.MessageRejectError) {
	f := &field.NoTimeInForceRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTimeInForceRules reads a NoTimeInForceRules from MarketDefinition.
func (m Message) GetNoTimeInForceRules(f *field.NoTimeInForceRulesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecInstRules is a non-required field for MarketDefinition.
func (m Message) NoExecInstRules() (*field.NoExecInstRulesField, quickfix.MessageRejectError) {
	f := &field.NoExecInstRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecInstRules reads a NoExecInstRules from MarketDefinition.
func (m Message) GetNoExecInstRules(f *field.NoExecInstRulesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for MarketDefinition.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from MarketDefinition.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MarketDefinition.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MarketDefinition.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MarketDefinition.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MarketDefinition.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MarketDefinition.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MarketDefinition.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for MarketDefinition.
func (m Message) ApplID() (*field.ApplIDField, quickfix.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from MarketDefinition.
func (m Message) GetApplID(f *field.ApplIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for MarketDefinition.
func (m Message) ApplSeqNum() (*field.ApplSeqNumField, quickfix.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from MarketDefinition.
func (m Message) GetApplSeqNum(f *field.ApplSeqNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for MarketDefinition.
func (m Message) ApplLastSeqNum() (*field.ApplLastSeqNumField, quickfix.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from MarketDefinition.
func (m Message) GetApplLastSeqNum(f *field.ApplLastSeqNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for MarketDefinition.
func (m Message) ApplResendFlag() (*field.ApplResendFlagField, quickfix.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from MarketDefinition.
func (m Message) GetApplResendFlag(f *field.ApplResendFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for MarketDefinition.
func New(
	marketreportid *field.MarketReportIDField,
	marketid *field.MarketIDField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("BU"))
	builder.Body.Set(marketreportid)
	builder.Body.Set(marketid)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "BU", r
}
