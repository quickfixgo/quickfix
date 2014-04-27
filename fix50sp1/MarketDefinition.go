package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
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
	marketreportid field.MarketReportID,
	marketid field.MarketID) MarketDefinitionBuilder {
	var builder MarketDefinitionBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildMsgType("BU"))
	builder.Body.Set(marketreportid)
	builder.Body.Set(marketid)
	return builder
}

//MarketReportID is a required field for MarketDefinition.
func (m MarketDefinition) MarketReportID() (*field.MarketReportID, errors.MessageRejectError) {
	f := new(field.MarketReportID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketReportID reads a MarketReportID from MarketDefinition.
func (m MarketDefinition) GetMarketReportID(f *field.MarketReportID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketReqID is a non-required field for MarketDefinition.
func (m MarketDefinition) MarketReqID() (*field.MarketReqID, errors.MessageRejectError) {
	f := new(field.MarketReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketReqID reads a MarketReqID from MarketDefinition.
func (m MarketDefinition) GetMarketReqID(f *field.MarketReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a required field for MarketDefinition.
func (m MarketDefinition) MarketID() (*field.MarketID, errors.MessageRejectError) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from MarketDefinition.
func (m MarketDefinition) GetMarketID(f *field.MarketID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for MarketDefinition.
func (m MarketDefinition) MarketSegmentID() (*field.MarketSegmentID, errors.MessageRejectError) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from MarketDefinition.
func (m MarketDefinition) GetMarketSegmentID(f *field.MarketSegmentID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentDesc is a non-required field for MarketDefinition.
func (m MarketDefinition) MarketSegmentDesc() (*field.MarketSegmentDesc, errors.MessageRejectError) {
	f := new(field.MarketSegmentDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentDesc reads a MarketSegmentDesc from MarketDefinition.
func (m MarketDefinition) GetMarketSegmentDesc(f *field.MarketSegmentDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedMktSegmDescLen is a non-required field for MarketDefinition.
func (m MarketDefinition) EncodedMktSegmDescLen() (*field.EncodedMktSegmDescLen, errors.MessageRejectError) {
	f := new(field.EncodedMktSegmDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedMktSegmDescLen reads a EncodedMktSegmDescLen from MarketDefinition.
func (m MarketDefinition) GetEncodedMktSegmDescLen(f *field.EncodedMktSegmDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedMktSegmDesc is a non-required field for MarketDefinition.
func (m MarketDefinition) EncodedMktSegmDesc() (*field.EncodedMktSegmDesc, errors.MessageRejectError) {
	f := new(field.EncodedMktSegmDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedMktSegmDesc reads a EncodedMktSegmDesc from MarketDefinition.
func (m MarketDefinition) GetEncodedMktSegmDesc(f *field.EncodedMktSegmDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParentMktSegmID is a non-required field for MarketDefinition.
func (m MarketDefinition) ParentMktSegmID() (*field.ParentMktSegmID, errors.MessageRejectError) {
	f := new(field.ParentMktSegmID)
	err := m.Body.Get(f)
	return f, err
}

//GetParentMktSegmID reads a ParentMktSegmID from MarketDefinition.
func (m MarketDefinition) GetParentMktSegmID(f *field.ParentMktSegmID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for MarketDefinition.
func (m MarketDefinition) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from MarketDefinition.
func (m MarketDefinition) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTickRules is a non-required field for MarketDefinition.
func (m MarketDefinition) NoTickRules() (*field.NoTickRules, errors.MessageRejectError) {
	f := new(field.NoTickRules)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTickRules reads a NoTickRules from MarketDefinition.
func (m MarketDefinition) GetNoTickRules(f *field.NoTickRules) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLotTypeRules is a non-required field for MarketDefinition.
func (m MarketDefinition) NoLotTypeRules() (*field.NoLotTypeRules, errors.MessageRejectError) {
	f := new(field.NoLotTypeRules)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLotTypeRules reads a NoLotTypeRules from MarketDefinition.
func (m MarketDefinition) GetNoLotTypeRules(f *field.NoLotTypeRules) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceLimitType is a non-required field for MarketDefinition.
func (m MarketDefinition) PriceLimitType() (*field.PriceLimitType, errors.MessageRejectError) {
	f := new(field.PriceLimitType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceLimitType reads a PriceLimitType from MarketDefinition.
func (m MarketDefinition) GetPriceLimitType(f *field.PriceLimitType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LowLimitPrice is a non-required field for MarketDefinition.
func (m MarketDefinition) LowLimitPrice() (*field.LowLimitPrice, errors.MessageRejectError) {
	f := new(field.LowLimitPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetLowLimitPrice reads a LowLimitPrice from MarketDefinition.
func (m MarketDefinition) GetLowLimitPrice(f *field.LowLimitPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HighLimitPrice is a non-required field for MarketDefinition.
func (m MarketDefinition) HighLimitPrice() (*field.HighLimitPrice, errors.MessageRejectError) {
	f := new(field.HighLimitPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetHighLimitPrice reads a HighLimitPrice from MarketDefinition.
func (m MarketDefinition) GetHighLimitPrice(f *field.HighLimitPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingReferencePrice is a non-required field for MarketDefinition.
func (m MarketDefinition) TradingReferencePrice() (*field.TradingReferencePrice, errors.MessageRejectError) {
	f := new(field.TradingReferencePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingReferencePrice reads a TradingReferencePrice from MarketDefinition.
func (m MarketDefinition) GetTradingReferencePrice(f *field.TradingReferencePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpirationCycle is a non-required field for MarketDefinition.
func (m MarketDefinition) ExpirationCycle() (*field.ExpirationCycle, errors.MessageRejectError) {
	f := new(field.ExpirationCycle)
	err := m.Body.Get(f)
	return f, err
}

//GetExpirationCycle reads a ExpirationCycle from MarketDefinition.
func (m MarketDefinition) GetExpirationCycle(f *field.ExpirationCycle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinTradeVol is a non-required field for MarketDefinition.
func (m MarketDefinition) MinTradeVol() (*field.MinTradeVol, errors.MessageRejectError) {
	f := new(field.MinTradeVol)
	err := m.Body.Get(f)
	return f, err
}

//GetMinTradeVol reads a MinTradeVol from MarketDefinition.
func (m MarketDefinition) GetMinTradeVol(f *field.MinTradeVol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxTradeVol is a non-required field for MarketDefinition.
func (m MarketDefinition) MaxTradeVol() (*field.MaxTradeVol, errors.MessageRejectError) {
	f := new(field.MaxTradeVol)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxTradeVol reads a MaxTradeVol from MarketDefinition.
func (m MarketDefinition) GetMaxTradeVol(f *field.MaxTradeVol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceVariation is a non-required field for MarketDefinition.
func (m MarketDefinition) MaxPriceVariation() (*field.MaxPriceVariation, errors.MessageRejectError) {
	f := new(field.MaxPriceVariation)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceVariation reads a MaxPriceVariation from MarketDefinition.
func (m MarketDefinition) GetMaxPriceVariation(f *field.MaxPriceVariation) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ImpliedMarketIndicator is a non-required field for MarketDefinition.
func (m MarketDefinition) ImpliedMarketIndicator() (*field.ImpliedMarketIndicator, errors.MessageRejectError) {
	f := new(field.ImpliedMarketIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetImpliedMarketIndicator reads a ImpliedMarketIndicator from MarketDefinition.
func (m MarketDefinition) GetImpliedMarketIndicator(f *field.ImpliedMarketIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingCurrency is a non-required field for MarketDefinition.
func (m MarketDefinition) TradingCurrency() (*field.TradingCurrency, errors.MessageRejectError) {
	f := new(field.TradingCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingCurrency reads a TradingCurrency from MarketDefinition.
func (m MarketDefinition) GetTradingCurrency(f *field.TradingCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundLot is a non-required field for MarketDefinition.
func (m MarketDefinition) RoundLot() (*field.RoundLot, errors.MessageRejectError) {
	f := new(field.RoundLot)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundLot reads a RoundLot from MarketDefinition.
func (m MarketDefinition) GetRoundLot(f *field.RoundLot) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegModel is a non-required field for MarketDefinition.
func (m MarketDefinition) MultilegModel() (*field.MultilegModel, errors.MessageRejectError) {
	f := new(field.MultilegModel)
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegModel reads a MultilegModel from MarketDefinition.
func (m MarketDefinition) GetMultilegModel(f *field.MultilegModel) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegPriceMethod is a non-required field for MarketDefinition.
func (m MarketDefinition) MultilegPriceMethod() (*field.MultilegPriceMethod, errors.MessageRejectError) {
	f := new(field.MultilegPriceMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegPriceMethod reads a MultilegPriceMethod from MarketDefinition.
func (m MarketDefinition) GetMultilegPriceMethod(f *field.MultilegPriceMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for MarketDefinition.
func (m MarketDefinition) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from MarketDefinition.
func (m MarketDefinition) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrdTypeRules is a non-required field for MarketDefinition.
func (m MarketDefinition) NoOrdTypeRules() (*field.NoOrdTypeRules, errors.MessageRejectError) {
	f := new(field.NoOrdTypeRules)
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrdTypeRules reads a NoOrdTypeRules from MarketDefinition.
func (m MarketDefinition) GetNoOrdTypeRules(f *field.NoOrdTypeRules) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTimeInForceRules is a non-required field for MarketDefinition.
func (m MarketDefinition) NoTimeInForceRules() (*field.NoTimeInForceRules, errors.MessageRejectError) {
	f := new(field.NoTimeInForceRules)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTimeInForceRules reads a NoTimeInForceRules from MarketDefinition.
func (m MarketDefinition) GetNoTimeInForceRules(f *field.NoTimeInForceRules) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecInstRules is a non-required field for MarketDefinition.
func (m MarketDefinition) NoExecInstRules() (*field.NoExecInstRules, errors.MessageRejectError) {
	f := new(field.NoExecInstRules)
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecInstRules reads a NoExecInstRules from MarketDefinition.
func (m MarketDefinition) GetNoExecInstRules(f *field.NoExecInstRules) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for MarketDefinition.
func (m MarketDefinition) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from MarketDefinition.
func (m MarketDefinition) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MarketDefinition.
func (m MarketDefinition) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MarketDefinition.
func (m MarketDefinition) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MarketDefinition.
func (m MarketDefinition) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MarketDefinition.
func (m MarketDefinition) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MarketDefinition.
func (m MarketDefinition) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MarketDefinition.
func (m MarketDefinition) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for MarketDefinition.
func (m MarketDefinition) ApplID() (*field.ApplID, errors.MessageRejectError) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from MarketDefinition.
func (m MarketDefinition) GetApplID(f *field.ApplID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for MarketDefinition.
func (m MarketDefinition) ApplSeqNum() (*field.ApplSeqNum, errors.MessageRejectError) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from MarketDefinition.
func (m MarketDefinition) GetApplSeqNum(f *field.ApplSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for MarketDefinition.
func (m MarketDefinition) ApplLastSeqNum() (*field.ApplLastSeqNum, errors.MessageRejectError) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from MarketDefinition.
func (m MarketDefinition) GetApplLastSeqNum(f *field.ApplLastSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for MarketDefinition.
func (m MarketDefinition) ApplResendFlag() (*field.ApplResendFlag, errors.MessageRejectError) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from MarketDefinition.
func (m MarketDefinition) GetApplResendFlag(f *field.ApplResendFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}
