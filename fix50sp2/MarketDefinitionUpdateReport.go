package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
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
	builder.Body.Set(marketreportid)
	builder.Body.Set(marketid)
	return builder
}

//MarketReportID is a required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketReportID() (field.MarketReportID, errors.MessageRejectError) {
	var f field.MarketReportID
	err := m.Body.Get(&f)
	return f, err
}

//MarketReqID is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketReqID() (field.MarketReqID, errors.MessageRejectError) {
	var f field.MarketReqID
	err := m.Body.Get(&f)
	return f, err
}

//MarketUpdateAction is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketUpdateAction() (field.MarketUpdateAction, errors.MessageRejectError) {
	var f field.MarketUpdateAction
	err := m.Body.Get(&f)
	return f, err
}

//MarketID is a required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketID() (field.MarketID, errors.MessageRejectError) {
	var f field.MarketID
	err := m.Body.Get(&f)
	return f, err
}

//MarketSegmentID is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketSegmentID() (field.MarketSegmentID, errors.MessageRejectError) {
	var f field.MarketSegmentID
	err := m.Body.Get(&f)
	return f, err
}

//MarketSegmentDesc is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MarketSegmentDesc() (field.MarketSegmentDesc, errors.MessageRejectError) {
	var f field.MarketSegmentDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedMktSegmDescLen is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) EncodedMktSegmDescLen() (field.EncodedMktSegmDescLen, errors.MessageRejectError) {
	var f field.EncodedMktSegmDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedMktSegmDesc is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) EncodedMktSegmDesc() (field.EncodedMktSegmDesc, errors.MessageRejectError) {
	var f field.EncodedMktSegmDesc
	err := m.Body.Get(&f)
	return f, err
}

//ParentMktSegmID is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ParentMktSegmID() (field.ParentMktSegmID, errors.MessageRejectError) {
	var f field.ParentMktSegmID
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) Currency() (field.Currency, errors.MessageRejectError) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//NoTickRules is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) NoTickRules() (field.NoTickRules, errors.MessageRejectError) {
	var f field.NoTickRules
	err := m.Body.Get(&f)
	return f, err
}

//NoLotTypeRules is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) NoLotTypeRules() (field.NoLotTypeRules, errors.MessageRejectError) {
	var f field.NoLotTypeRules
	err := m.Body.Get(&f)
	return f, err
}

//PriceLimitType is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) PriceLimitType() (field.PriceLimitType, errors.MessageRejectError) {
	var f field.PriceLimitType
	err := m.Body.Get(&f)
	return f, err
}

//LowLimitPrice is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) LowLimitPrice() (field.LowLimitPrice, errors.MessageRejectError) {
	var f field.LowLimitPrice
	err := m.Body.Get(&f)
	return f, err
}

//HighLimitPrice is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) HighLimitPrice() (field.HighLimitPrice, errors.MessageRejectError) {
	var f field.HighLimitPrice
	err := m.Body.Get(&f)
	return f, err
}

//TradingReferencePrice is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) TradingReferencePrice() (field.TradingReferencePrice, errors.MessageRejectError) {
	var f field.TradingReferencePrice
	err := m.Body.Get(&f)
	return f, err
}

//ExpirationCycle is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ExpirationCycle() (field.ExpirationCycle, errors.MessageRejectError) {
	var f field.ExpirationCycle
	err := m.Body.Get(&f)
	return f, err
}

//MinTradeVol is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MinTradeVol() (field.MinTradeVol, errors.MessageRejectError) {
	var f field.MinTradeVol
	err := m.Body.Get(&f)
	return f, err
}

//MaxTradeVol is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MaxTradeVol() (field.MaxTradeVol, errors.MessageRejectError) {
	var f field.MaxTradeVol
	err := m.Body.Get(&f)
	return f, err
}

//MaxPriceVariation is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MaxPriceVariation() (field.MaxPriceVariation, errors.MessageRejectError) {
	var f field.MaxPriceVariation
	err := m.Body.Get(&f)
	return f, err
}

//ImpliedMarketIndicator is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ImpliedMarketIndicator() (field.ImpliedMarketIndicator, errors.MessageRejectError) {
	var f field.ImpliedMarketIndicator
	err := m.Body.Get(&f)
	return f, err
}

//TradingCurrency is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) TradingCurrency() (field.TradingCurrency, errors.MessageRejectError) {
	var f field.TradingCurrency
	err := m.Body.Get(&f)
	return f, err
}

//RoundLot is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) RoundLot() (field.RoundLot, errors.MessageRejectError) {
	var f field.RoundLot
	err := m.Body.Get(&f)
	return f, err
}

//MultilegModel is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MultilegModel() (field.MultilegModel, errors.MessageRejectError) {
	var f field.MultilegModel
	err := m.Body.Get(&f)
	return f, err
}

//MultilegPriceMethod is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) MultilegPriceMethod() (field.MultilegPriceMethod, errors.MessageRejectError) {
	var f field.MultilegPriceMethod
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) PriceType() (field.PriceType, errors.MessageRejectError) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//NoOrdTypeRules is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) NoOrdTypeRules() (field.NoOrdTypeRules, errors.MessageRejectError) {
	var f field.NoOrdTypeRules
	err := m.Body.Get(&f)
	return f, err
}

//NoTimeInForceRules is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) NoTimeInForceRules() (field.NoTimeInForceRules, errors.MessageRejectError) {
	var f field.NoTimeInForceRules
	err := m.Body.Get(&f)
	return f, err
}

//NoExecInstRules is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) NoExecInstRules() (field.NoExecInstRules, errors.MessageRejectError) {
	var f field.NoExecInstRules
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//ApplID is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ApplID() (field.ApplID, errors.MessageRejectError) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ApplSeqNum() (field.ApplSeqNum, errors.MessageRejectError) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ApplLastSeqNum() (field.ApplLastSeqNum, errors.MessageRejectError) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for MarketDefinitionUpdateReport.
func (m MarketDefinitionUpdateReport) ApplResendFlag() (field.ApplResendFlag, errors.MessageRejectError) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}
