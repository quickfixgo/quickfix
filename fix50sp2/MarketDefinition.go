package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//MarketDefinition msg type = BU.
type MarketDefinition struct {
	message.Message
}

//MarketDefinitionBuilder builds MarketDefinition messages.
type MarketDefinitionBuilder struct {
	message.MessageBuilder
}

//CreateMarketDefinitionBuilder returns an initialized MarketDefinitionBuilder with specified required fields.
func CreateMarketDefinitionBuilder(
	marketreportid *field.MarketReportIDField,
	marketid *field.MarketIDField) MarketDefinitionBuilder {
	var builder MarketDefinitionBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("BU"))
	builder.Body().Set(marketreportid)
	builder.Body().Set(marketid)
	return builder
}

//MarketReportID is a required field for MarketDefinition.
func (m MarketDefinition) MarketReportID() (*field.MarketReportIDField, errors.MessageRejectError) {
	f := &field.MarketReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketReportID reads a MarketReportID from MarketDefinition.
func (m MarketDefinition) GetMarketReportID(f *field.MarketReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketReqID is a non-required field for MarketDefinition.
func (m MarketDefinition) MarketReqID() (*field.MarketReqIDField, errors.MessageRejectError) {
	f := &field.MarketReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketReqID reads a MarketReqID from MarketDefinition.
func (m MarketDefinition) GetMarketReqID(f *field.MarketReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a required field for MarketDefinition.
func (m MarketDefinition) MarketID() (*field.MarketIDField, errors.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from MarketDefinition.
func (m MarketDefinition) GetMarketID(f *field.MarketIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for MarketDefinition.
func (m MarketDefinition) MarketSegmentID() (*field.MarketSegmentIDField, errors.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from MarketDefinition.
func (m MarketDefinition) GetMarketSegmentID(f *field.MarketSegmentIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentDesc is a non-required field for MarketDefinition.
func (m MarketDefinition) MarketSegmentDesc() (*field.MarketSegmentDescField, errors.MessageRejectError) {
	f := &field.MarketSegmentDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentDesc reads a MarketSegmentDesc from MarketDefinition.
func (m MarketDefinition) GetMarketSegmentDesc(f *field.MarketSegmentDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedMktSegmDescLen is a non-required field for MarketDefinition.
func (m MarketDefinition) EncodedMktSegmDescLen() (*field.EncodedMktSegmDescLenField, errors.MessageRejectError) {
	f := &field.EncodedMktSegmDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedMktSegmDescLen reads a EncodedMktSegmDescLen from MarketDefinition.
func (m MarketDefinition) GetEncodedMktSegmDescLen(f *field.EncodedMktSegmDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedMktSegmDesc is a non-required field for MarketDefinition.
func (m MarketDefinition) EncodedMktSegmDesc() (*field.EncodedMktSegmDescField, errors.MessageRejectError) {
	f := &field.EncodedMktSegmDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedMktSegmDesc reads a EncodedMktSegmDesc from MarketDefinition.
func (m MarketDefinition) GetEncodedMktSegmDesc(f *field.EncodedMktSegmDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParentMktSegmID is a non-required field for MarketDefinition.
func (m MarketDefinition) ParentMktSegmID() (*field.ParentMktSegmIDField, errors.MessageRejectError) {
	f := &field.ParentMktSegmIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetParentMktSegmID reads a ParentMktSegmID from MarketDefinition.
func (m MarketDefinition) GetParentMktSegmID(f *field.ParentMktSegmIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for MarketDefinition.
func (m MarketDefinition) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from MarketDefinition.
func (m MarketDefinition) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTickRules is a non-required field for MarketDefinition.
func (m MarketDefinition) NoTickRules() (*field.NoTickRulesField, errors.MessageRejectError) {
	f := &field.NoTickRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTickRules reads a NoTickRules from MarketDefinition.
func (m MarketDefinition) GetNoTickRules(f *field.NoTickRulesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLotTypeRules is a non-required field for MarketDefinition.
func (m MarketDefinition) NoLotTypeRules() (*field.NoLotTypeRulesField, errors.MessageRejectError) {
	f := &field.NoLotTypeRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLotTypeRules reads a NoLotTypeRules from MarketDefinition.
func (m MarketDefinition) GetNoLotTypeRules(f *field.NoLotTypeRulesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceLimitType is a non-required field for MarketDefinition.
func (m MarketDefinition) PriceLimitType() (*field.PriceLimitTypeField, errors.MessageRejectError) {
	f := &field.PriceLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceLimitType reads a PriceLimitType from MarketDefinition.
func (m MarketDefinition) GetPriceLimitType(f *field.PriceLimitTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LowLimitPrice is a non-required field for MarketDefinition.
func (m MarketDefinition) LowLimitPrice() (*field.LowLimitPriceField, errors.MessageRejectError) {
	f := &field.LowLimitPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLowLimitPrice reads a LowLimitPrice from MarketDefinition.
func (m MarketDefinition) GetLowLimitPrice(f *field.LowLimitPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HighLimitPrice is a non-required field for MarketDefinition.
func (m MarketDefinition) HighLimitPrice() (*field.HighLimitPriceField, errors.MessageRejectError) {
	f := &field.HighLimitPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHighLimitPrice reads a HighLimitPrice from MarketDefinition.
func (m MarketDefinition) GetHighLimitPrice(f *field.HighLimitPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingReferencePrice is a non-required field for MarketDefinition.
func (m MarketDefinition) TradingReferencePrice() (*field.TradingReferencePriceField, errors.MessageRejectError) {
	f := &field.TradingReferencePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingReferencePrice reads a TradingReferencePrice from MarketDefinition.
func (m MarketDefinition) GetTradingReferencePrice(f *field.TradingReferencePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpirationCycle is a non-required field for MarketDefinition.
func (m MarketDefinition) ExpirationCycle() (*field.ExpirationCycleField, errors.MessageRejectError) {
	f := &field.ExpirationCycleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpirationCycle reads a ExpirationCycle from MarketDefinition.
func (m MarketDefinition) GetExpirationCycle(f *field.ExpirationCycleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinTradeVol is a non-required field for MarketDefinition.
func (m MarketDefinition) MinTradeVol() (*field.MinTradeVolField, errors.MessageRejectError) {
	f := &field.MinTradeVolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinTradeVol reads a MinTradeVol from MarketDefinition.
func (m MarketDefinition) GetMinTradeVol(f *field.MinTradeVolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxTradeVol is a non-required field for MarketDefinition.
func (m MarketDefinition) MaxTradeVol() (*field.MaxTradeVolField, errors.MessageRejectError) {
	f := &field.MaxTradeVolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxTradeVol reads a MaxTradeVol from MarketDefinition.
func (m MarketDefinition) GetMaxTradeVol(f *field.MaxTradeVolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceVariation is a non-required field for MarketDefinition.
func (m MarketDefinition) MaxPriceVariation() (*field.MaxPriceVariationField, errors.MessageRejectError) {
	f := &field.MaxPriceVariationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceVariation reads a MaxPriceVariation from MarketDefinition.
func (m MarketDefinition) GetMaxPriceVariation(f *field.MaxPriceVariationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ImpliedMarketIndicator is a non-required field for MarketDefinition.
func (m MarketDefinition) ImpliedMarketIndicator() (*field.ImpliedMarketIndicatorField, errors.MessageRejectError) {
	f := &field.ImpliedMarketIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetImpliedMarketIndicator reads a ImpliedMarketIndicator from MarketDefinition.
func (m MarketDefinition) GetImpliedMarketIndicator(f *field.ImpliedMarketIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingCurrency is a non-required field for MarketDefinition.
func (m MarketDefinition) TradingCurrency() (*field.TradingCurrencyField, errors.MessageRejectError) {
	f := &field.TradingCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingCurrency reads a TradingCurrency from MarketDefinition.
func (m MarketDefinition) GetTradingCurrency(f *field.TradingCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundLot is a non-required field for MarketDefinition.
func (m MarketDefinition) RoundLot() (*field.RoundLotField, errors.MessageRejectError) {
	f := &field.RoundLotField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundLot reads a RoundLot from MarketDefinition.
func (m MarketDefinition) GetRoundLot(f *field.RoundLotField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegModel is a non-required field for MarketDefinition.
func (m MarketDefinition) MultilegModel() (*field.MultilegModelField, errors.MessageRejectError) {
	f := &field.MultilegModelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegModel reads a MultilegModel from MarketDefinition.
func (m MarketDefinition) GetMultilegModel(f *field.MultilegModelField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegPriceMethod is a non-required field for MarketDefinition.
func (m MarketDefinition) MultilegPriceMethod() (*field.MultilegPriceMethodField, errors.MessageRejectError) {
	f := &field.MultilegPriceMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegPriceMethod reads a MultilegPriceMethod from MarketDefinition.
func (m MarketDefinition) GetMultilegPriceMethod(f *field.MultilegPriceMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for MarketDefinition.
func (m MarketDefinition) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from MarketDefinition.
func (m MarketDefinition) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrdTypeRules is a non-required field for MarketDefinition.
func (m MarketDefinition) NoOrdTypeRules() (*field.NoOrdTypeRulesField, errors.MessageRejectError) {
	f := &field.NoOrdTypeRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrdTypeRules reads a NoOrdTypeRules from MarketDefinition.
func (m MarketDefinition) GetNoOrdTypeRules(f *field.NoOrdTypeRulesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTimeInForceRules is a non-required field for MarketDefinition.
func (m MarketDefinition) NoTimeInForceRules() (*field.NoTimeInForceRulesField, errors.MessageRejectError) {
	f := &field.NoTimeInForceRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTimeInForceRules reads a NoTimeInForceRules from MarketDefinition.
func (m MarketDefinition) GetNoTimeInForceRules(f *field.NoTimeInForceRulesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecInstRules is a non-required field for MarketDefinition.
func (m MarketDefinition) NoExecInstRules() (*field.NoExecInstRulesField, errors.MessageRejectError) {
	f := &field.NoExecInstRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecInstRules reads a NoExecInstRules from MarketDefinition.
func (m MarketDefinition) GetNoExecInstRules(f *field.NoExecInstRulesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for MarketDefinition.
func (m MarketDefinition) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from MarketDefinition.
func (m MarketDefinition) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MarketDefinition.
func (m MarketDefinition) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MarketDefinition.
func (m MarketDefinition) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MarketDefinition.
func (m MarketDefinition) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MarketDefinition.
func (m MarketDefinition) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MarketDefinition.
func (m MarketDefinition) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MarketDefinition.
func (m MarketDefinition) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for MarketDefinition.
func (m MarketDefinition) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from MarketDefinition.
func (m MarketDefinition) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for MarketDefinition.
func (m MarketDefinition) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from MarketDefinition.
func (m MarketDefinition) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for MarketDefinition.
func (m MarketDefinition) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from MarketDefinition.
func (m MarketDefinition) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for MarketDefinition.
func (m MarketDefinition) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from MarketDefinition.
func (m MarketDefinition) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}
