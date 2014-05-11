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

//MarketDefinitionUpdateReport msg type = BV.
type MarketDefinitionUpdateReport struct {
	message.Message
}

//MarketDefinitionUpdateReportBuilder builds MarketDefinitionUpdateReport messages.
type MarketDefinitionUpdateReportBuilder struct {
	message.MessageBuilder
}

//CreateMarketDefinitionUpdateReportBuilder returns an initialized MarketDefinitionUpdateReportBuilder with specified required fields.
func CreateMarketDefinitionUpdateReportBuilder(
	marketreportid *field.MarketReportIDField,
	marketid *field.MarketIDField) MarketDefinitionUpdateReportBuilder {
	var builder MarketDefinitionUpdateReportBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("BV"))
	builder.Body().Set(marketreportid)
	builder.Body().Set(marketid)
	return builder
}

//MarketReportID is a required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketReportID() (*field.MarketReportIDField, errors.MessageRejectError) {
	f := &field.MarketReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketReportID reads a MarketReportID from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMarketReportID(f *field.MarketReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketReqID is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketReqID() (*field.MarketReqIDField, errors.MessageRejectError) {
	f := &field.MarketReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketReqID reads a MarketReqID from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMarketReqID(f *field.MarketReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketUpdateAction is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketUpdateAction() (*field.MarketUpdateActionField, errors.MessageRejectError) {
	f := &field.MarketUpdateActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketUpdateAction reads a MarketUpdateAction from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMarketUpdateAction(f *field.MarketUpdateActionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketID() (*field.MarketIDField, errors.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMarketID(f *field.MarketIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketSegmentID() (*field.MarketSegmentIDField, errors.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMarketSegmentID(f *field.MarketSegmentIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentDesc is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketSegmentDesc() (*field.MarketSegmentDescField, errors.MessageRejectError) {
	f := &field.MarketSegmentDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentDesc reads a MarketSegmentDesc from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMarketSegmentDesc(f *field.MarketSegmentDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedMktSegmDescLen is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) EncodedMktSegmDescLen() (*field.EncodedMktSegmDescLenField, errors.MessageRejectError) {
	f := &field.EncodedMktSegmDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedMktSegmDescLen reads a EncodedMktSegmDescLen from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetEncodedMktSegmDescLen(f *field.EncodedMktSegmDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedMktSegmDesc is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) EncodedMktSegmDesc() (*field.EncodedMktSegmDescField, errors.MessageRejectError) {
	f := &field.EncodedMktSegmDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedMktSegmDesc reads a EncodedMktSegmDesc from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetEncodedMktSegmDesc(f *field.EncodedMktSegmDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParentMktSegmID is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ParentMktSegmID() (*field.ParentMktSegmIDField, errors.MessageRejectError) {
	f := &field.ParentMktSegmIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetParentMktSegmID reads a ParentMktSegmID from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetParentMktSegmID(f *field.ParentMktSegmIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTickRules is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) NoTickRules() (*field.NoTickRulesField, errors.MessageRejectError) {
	f := &field.NoTickRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTickRules reads a NoTickRules from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetNoTickRules(f *field.NoTickRulesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLotTypeRules is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) NoLotTypeRules() (*field.NoLotTypeRulesField, errors.MessageRejectError) {
	f := &field.NoLotTypeRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLotTypeRules reads a NoLotTypeRules from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetNoLotTypeRules(f *field.NoLotTypeRulesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceLimitType is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) PriceLimitType() (*field.PriceLimitTypeField, errors.MessageRejectError) {
	f := &field.PriceLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceLimitType reads a PriceLimitType from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetPriceLimitType(f *field.PriceLimitTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LowLimitPrice is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) LowLimitPrice() (*field.LowLimitPriceField, errors.MessageRejectError) {
	f := &field.LowLimitPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLowLimitPrice reads a LowLimitPrice from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetLowLimitPrice(f *field.LowLimitPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HighLimitPrice is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) HighLimitPrice() (*field.HighLimitPriceField, errors.MessageRejectError) {
	f := &field.HighLimitPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHighLimitPrice reads a HighLimitPrice from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetHighLimitPrice(f *field.HighLimitPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingReferencePrice is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) TradingReferencePrice() (*field.TradingReferencePriceField, errors.MessageRejectError) {
	f := &field.TradingReferencePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingReferencePrice reads a TradingReferencePrice from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetTradingReferencePrice(f *field.TradingReferencePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpirationCycle is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ExpirationCycle() (*field.ExpirationCycleField, errors.MessageRejectError) {
	f := &field.ExpirationCycleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpirationCycle reads a ExpirationCycle from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetExpirationCycle(f *field.ExpirationCycleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinTradeVol is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MinTradeVol() (*field.MinTradeVolField, errors.MessageRejectError) {
	f := &field.MinTradeVolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinTradeVol reads a MinTradeVol from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMinTradeVol(f *field.MinTradeVolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxTradeVol is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MaxTradeVol() (*field.MaxTradeVolField, errors.MessageRejectError) {
	f := &field.MaxTradeVolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxTradeVol reads a MaxTradeVol from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMaxTradeVol(f *field.MaxTradeVolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceVariation is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MaxPriceVariation() (*field.MaxPriceVariationField, errors.MessageRejectError) {
	f := &field.MaxPriceVariationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceVariation reads a MaxPriceVariation from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMaxPriceVariation(f *field.MaxPriceVariationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ImpliedMarketIndicator is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ImpliedMarketIndicator() (*field.ImpliedMarketIndicatorField, errors.MessageRejectError) {
	f := &field.ImpliedMarketIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetImpliedMarketIndicator reads a ImpliedMarketIndicator from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetImpliedMarketIndicator(f *field.ImpliedMarketIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingCurrency is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) TradingCurrency() (*field.TradingCurrencyField, errors.MessageRejectError) {
	f := &field.TradingCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingCurrency reads a TradingCurrency from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetTradingCurrency(f *field.TradingCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundLot is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) RoundLot() (*field.RoundLotField, errors.MessageRejectError) {
	f := &field.RoundLotField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundLot reads a RoundLot from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetRoundLot(f *field.RoundLotField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegModel is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MultilegModel() (*field.MultilegModelField, errors.MessageRejectError) {
	f := &field.MultilegModelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegModel reads a MultilegModel from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMultilegModel(f *field.MultilegModelField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegPriceMethod is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MultilegPriceMethod() (*field.MultilegPriceMethodField, errors.MessageRejectError) {
	f := &field.MultilegPriceMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegPriceMethod reads a MultilegPriceMethod from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMultilegPriceMethod(f *field.MultilegPriceMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrdTypeRules is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) NoOrdTypeRules() (*field.NoOrdTypeRulesField, errors.MessageRejectError) {
	f := &field.NoOrdTypeRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrdTypeRules reads a NoOrdTypeRules from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetNoOrdTypeRules(f *field.NoOrdTypeRulesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTimeInForceRules is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) NoTimeInForceRules() (*field.NoTimeInForceRulesField, errors.MessageRejectError) {
	f := &field.NoTimeInForceRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTimeInForceRules reads a NoTimeInForceRules from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetNoTimeInForceRules(f *field.NoTimeInForceRulesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecInstRules is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) NoExecInstRules() (*field.NoExecInstRulesField, errors.MessageRejectError) {
	f := &field.NoExecInstRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecInstRules reads a NoExecInstRules from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetNoExecInstRules(f *field.NoExecInstRulesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}
