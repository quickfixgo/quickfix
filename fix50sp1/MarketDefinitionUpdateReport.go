package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
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
	marketreportid field.MarketReportID,
	marketid field.MarketID) MarketDefinitionUpdateReportBuilder {
	var builder MarketDefinitionUpdateReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildMsgType("BV"))
	builder.Body.Set(marketreportid)
	builder.Body.Set(marketid)
	return builder
}

//MarketReportID is a required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketReportID() (*field.MarketReportID, errors.MessageRejectError) {
	f := new(field.MarketReportID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketReportID reads a MarketReportID from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMarketReportID(f *field.MarketReportID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketReqID is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketReqID() (*field.MarketReqID, errors.MessageRejectError) {
	f := new(field.MarketReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketReqID reads a MarketReqID from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMarketReqID(f *field.MarketReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketUpdateAction is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketUpdateAction() (*field.MarketUpdateAction, errors.MessageRejectError) {
	f := new(field.MarketUpdateAction)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketUpdateAction reads a MarketUpdateAction from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMarketUpdateAction(f *field.MarketUpdateAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketID() (*field.MarketID, errors.MessageRejectError) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMarketID(f *field.MarketID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketSegmentID() (*field.MarketSegmentID, errors.MessageRejectError) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMarketSegmentID(f *field.MarketSegmentID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentDesc is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketSegmentDesc() (*field.MarketSegmentDesc, errors.MessageRejectError) {
	f := new(field.MarketSegmentDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentDesc reads a MarketSegmentDesc from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMarketSegmentDesc(f *field.MarketSegmentDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedMktSegmDescLen is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) EncodedMktSegmDescLen() (*field.EncodedMktSegmDescLen, errors.MessageRejectError) {
	f := new(field.EncodedMktSegmDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedMktSegmDescLen reads a EncodedMktSegmDescLen from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetEncodedMktSegmDescLen(f *field.EncodedMktSegmDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedMktSegmDesc is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) EncodedMktSegmDesc() (*field.EncodedMktSegmDesc, errors.MessageRejectError) {
	f := new(field.EncodedMktSegmDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedMktSegmDesc reads a EncodedMktSegmDesc from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetEncodedMktSegmDesc(f *field.EncodedMktSegmDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParentMktSegmID is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ParentMktSegmID() (*field.ParentMktSegmID, errors.MessageRejectError) {
	f := new(field.ParentMktSegmID)
	err := m.Body.Get(f)
	return f, err
}

//GetParentMktSegmID reads a ParentMktSegmID from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetParentMktSegmID(f *field.ParentMktSegmID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTickRules is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) NoTickRules() (*field.NoTickRules, errors.MessageRejectError) {
	f := new(field.NoTickRules)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTickRules reads a NoTickRules from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetNoTickRules(f *field.NoTickRules) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLotTypeRules is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) NoLotTypeRules() (*field.NoLotTypeRules, errors.MessageRejectError) {
	f := new(field.NoLotTypeRules)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLotTypeRules reads a NoLotTypeRules from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetNoLotTypeRules(f *field.NoLotTypeRules) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceLimitType is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) PriceLimitType() (*field.PriceLimitType, errors.MessageRejectError) {
	f := new(field.PriceLimitType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceLimitType reads a PriceLimitType from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetPriceLimitType(f *field.PriceLimitType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LowLimitPrice is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) LowLimitPrice() (*field.LowLimitPrice, errors.MessageRejectError) {
	f := new(field.LowLimitPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetLowLimitPrice reads a LowLimitPrice from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetLowLimitPrice(f *field.LowLimitPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HighLimitPrice is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) HighLimitPrice() (*field.HighLimitPrice, errors.MessageRejectError) {
	f := new(field.HighLimitPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetHighLimitPrice reads a HighLimitPrice from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetHighLimitPrice(f *field.HighLimitPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingReferencePrice is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) TradingReferencePrice() (*field.TradingReferencePrice, errors.MessageRejectError) {
	f := new(field.TradingReferencePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingReferencePrice reads a TradingReferencePrice from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetTradingReferencePrice(f *field.TradingReferencePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpirationCycle is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ExpirationCycle() (*field.ExpirationCycle, errors.MessageRejectError) {
	f := new(field.ExpirationCycle)
	err := m.Body.Get(f)
	return f, err
}

//GetExpirationCycle reads a ExpirationCycle from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetExpirationCycle(f *field.ExpirationCycle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinTradeVol is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MinTradeVol() (*field.MinTradeVol, errors.MessageRejectError) {
	f := new(field.MinTradeVol)
	err := m.Body.Get(f)
	return f, err
}

//GetMinTradeVol reads a MinTradeVol from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMinTradeVol(f *field.MinTradeVol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxTradeVol is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MaxTradeVol() (*field.MaxTradeVol, errors.MessageRejectError) {
	f := new(field.MaxTradeVol)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxTradeVol reads a MaxTradeVol from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMaxTradeVol(f *field.MaxTradeVol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceVariation is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MaxPriceVariation() (*field.MaxPriceVariation, errors.MessageRejectError) {
	f := new(field.MaxPriceVariation)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceVariation reads a MaxPriceVariation from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMaxPriceVariation(f *field.MaxPriceVariation) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ImpliedMarketIndicator is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ImpliedMarketIndicator() (*field.ImpliedMarketIndicator, errors.MessageRejectError) {
	f := new(field.ImpliedMarketIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetImpliedMarketIndicator reads a ImpliedMarketIndicator from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetImpliedMarketIndicator(f *field.ImpliedMarketIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingCurrency is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) TradingCurrency() (*field.TradingCurrency, errors.MessageRejectError) {
	f := new(field.TradingCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingCurrency reads a TradingCurrency from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetTradingCurrency(f *field.TradingCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundLot is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) RoundLot() (*field.RoundLot, errors.MessageRejectError) {
	f := new(field.RoundLot)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundLot reads a RoundLot from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetRoundLot(f *field.RoundLot) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegModel is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MultilegModel() (*field.MultilegModel, errors.MessageRejectError) {
	f := new(field.MultilegModel)
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegModel reads a MultilegModel from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMultilegModel(f *field.MultilegModel) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegPriceMethod is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MultilegPriceMethod() (*field.MultilegPriceMethod, errors.MessageRejectError) {
	f := new(field.MultilegPriceMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegPriceMethod reads a MultilegPriceMethod from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetMultilegPriceMethod(f *field.MultilegPriceMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrdTypeRules is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) NoOrdTypeRules() (*field.NoOrdTypeRules, errors.MessageRejectError) {
	f := new(field.NoOrdTypeRules)
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrdTypeRules reads a NoOrdTypeRules from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetNoOrdTypeRules(f *field.NoOrdTypeRules) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTimeInForceRules is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) NoTimeInForceRules() (*field.NoTimeInForceRules, errors.MessageRejectError) {
	f := new(field.NoTimeInForceRules)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTimeInForceRules reads a NoTimeInForceRules from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetNoTimeInForceRules(f *field.NoTimeInForceRules) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecInstRules is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) NoExecInstRules() (*field.NoExecInstRules, errors.MessageRejectError) {
	f := new(field.NoExecInstRules)
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecInstRules reads a NoExecInstRules from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetNoExecInstRules(f *field.NoExecInstRules) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ApplID() (*field.ApplID, errors.MessageRejectError) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetApplID(f *field.ApplID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ApplSeqNum() (*field.ApplSeqNum, errors.MessageRejectError) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetApplSeqNum(f *field.ApplSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ApplLastSeqNum() (*field.ApplLastSeqNum, errors.MessageRejectError) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetApplLastSeqNum(f *field.ApplLastSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ApplResendFlag() (*field.ApplResendFlag, errors.MessageRejectError) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) GetApplResendFlag(f *field.ApplResendFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}
