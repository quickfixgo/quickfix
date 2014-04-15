package fix50sp1

import (
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
	builder.Body.Set(marketreportid)
	builder.Body.Set(marketid)
	return builder
}

//MarketReportID is a required field for MarketDefinition.
func (m MarketDefinition) MarketReportID() (field.MarketReportID, error) {
	var f field.MarketReportID
	err := m.Body.Get(&f)
	return f, err
}

//MarketReqID is a non-required field for MarketDefinition.
func (m MarketDefinition) MarketReqID() (field.MarketReqID, error) {
	var f field.MarketReqID
	err := m.Body.Get(&f)
	return f, err
}

//MarketID is a required field for MarketDefinition.
func (m MarketDefinition) MarketID() (field.MarketID, error) {
	var f field.MarketID
	err := m.Body.Get(&f)
	return f, err
}

//MarketSegmentID is a non-required field for MarketDefinition.
func (m MarketDefinition) MarketSegmentID() (field.MarketSegmentID, error) {
	var f field.MarketSegmentID
	err := m.Body.Get(&f)
	return f, err
}

//MarketSegmentDesc is a non-required field for MarketDefinition.
func (m MarketDefinition) MarketSegmentDesc() (field.MarketSegmentDesc, error) {
	var f field.MarketSegmentDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedMktSegmDescLen is a non-required field for MarketDefinition.
func (m MarketDefinition) EncodedMktSegmDescLen() (field.EncodedMktSegmDescLen, error) {
	var f field.EncodedMktSegmDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedMktSegmDesc is a non-required field for MarketDefinition.
func (m MarketDefinition) EncodedMktSegmDesc() (field.EncodedMktSegmDesc, error) {
	var f field.EncodedMktSegmDesc
	err := m.Body.Get(&f)
	return f, err
}

//ParentMktSegmID is a non-required field for MarketDefinition.
func (m MarketDefinition) ParentMktSegmID() (field.ParentMktSegmID, error) {
	var f field.ParentMktSegmID
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for MarketDefinition.
func (m MarketDefinition) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//NoTickRules is a non-required field for MarketDefinition.
func (m MarketDefinition) NoTickRules() (field.NoTickRules, error) {
	var f field.NoTickRules
	err := m.Body.Get(&f)
	return f, err
}

//NoLotTypeRules is a non-required field for MarketDefinition.
func (m MarketDefinition) NoLotTypeRules() (field.NoLotTypeRules, error) {
	var f field.NoLotTypeRules
	err := m.Body.Get(&f)
	return f, err
}

//PriceLimitType is a non-required field for MarketDefinition.
func (m MarketDefinition) PriceLimitType() (field.PriceLimitType, error) {
	var f field.PriceLimitType
	err := m.Body.Get(&f)
	return f, err
}

//LowLimitPrice is a non-required field for MarketDefinition.
func (m MarketDefinition) LowLimitPrice() (field.LowLimitPrice, error) {
	var f field.LowLimitPrice
	err := m.Body.Get(&f)
	return f, err
}

//HighLimitPrice is a non-required field for MarketDefinition.
func (m MarketDefinition) HighLimitPrice() (field.HighLimitPrice, error) {
	var f field.HighLimitPrice
	err := m.Body.Get(&f)
	return f, err
}

//TradingReferencePrice is a non-required field for MarketDefinition.
func (m MarketDefinition) TradingReferencePrice() (field.TradingReferencePrice, error) {
	var f field.TradingReferencePrice
	err := m.Body.Get(&f)
	return f, err
}

//ExpirationCycle is a non-required field for MarketDefinition.
func (m MarketDefinition) ExpirationCycle() (field.ExpirationCycle, error) {
	var f field.ExpirationCycle
	err := m.Body.Get(&f)
	return f, err
}

//MinTradeVol is a non-required field for MarketDefinition.
func (m MarketDefinition) MinTradeVol() (field.MinTradeVol, error) {
	var f field.MinTradeVol
	err := m.Body.Get(&f)
	return f, err
}

//MaxTradeVol is a non-required field for MarketDefinition.
func (m MarketDefinition) MaxTradeVol() (field.MaxTradeVol, error) {
	var f field.MaxTradeVol
	err := m.Body.Get(&f)
	return f, err
}

//MaxPriceVariation is a non-required field for MarketDefinition.
func (m MarketDefinition) MaxPriceVariation() (field.MaxPriceVariation, error) {
	var f field.MaxPriceVariation
	err := m.Body.Get(&f)
	return f, err
}

//ImpliedMarketIndicator is a non-required field for MarketDefinition.
func (m MarketDefinition) ImpliedMarketIndicator() (field.ImpliedMarketIndicator, error) {
	var f field.ImpliedMarketIndicator
	err := m.Body.Get(&f)
	return f, err
}

//TradingCurrency is a non-required field for MarketDefinition.
func (m MarketDefinition) TradingCurrency() (field.TradingCurrency, error) {
	var f field.TradingCurrency
	err := m.Body.Get(&f)
	return f, err
}

//RoundLot is a non-required field for MarketDefinition.
func (m MarketDefinition) RoundLot() (field.RoundLot, error) {
	var f field.RoundLot
	err := m.Body.Get(&f)
	return f, err
}

//MultilegModel is a non-required field for MarketDefinition.
func (m MarketDefinition) MultilegModel() (field.MultilegModel, error) {
	var f field.MultilegModel
	err := m.Body.Get(&f)
	return f, err
}

//MultilegPriceMethod is a non-required field for MarketDefinition.
func (m MarketDefinition) MultilegPriceMethod() (field.MultilegPriceMethod, error) {
	var f field.MultilegPriceMethod
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for MarketDefinition.
func (m MarketDefinition) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//NoOrdTypeRules is a non-required field for MarketDefinition.
func (m MarketDefinition) NoOrdTypeRules() (field.NoOrdTypeRules, error) {
	var f field.NoOrdTypeRules
	err := m.Body.Get(&f)
	return f, err
}

//NoTimeInForceRules is a non-required field for MarketDefinition.
func (m MarketDefinition) NoTimeInForceRules() (field.NoTimeInForceRules, error) {
	var f field.NoTimeInForceRules
	err := m.Body.Get(&f)
	return f, err
}

//NoExecInstRules is a non-required field for MarketDefinition.
func (m MarketDefinition) NoExecInstRules() (field.NoExecInstRules, error) {
	var f field.NoExecInstRules
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for MarketDefinition.
func (m MarketDefinition) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for MarketDefinition.
func (m MarketDefinition) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for MarketDefinition.
func (m MarketDefinition) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for MarketDefinition.
func (m MarketDefinition) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//ApplID is a non-required field for MarketDefinition.
func (m MarketDefinition) ApplID() (field.ApplID, error) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for MarketDefinition.
func (m MarketDefinition) ApplSeqNum() (field.ApplSeqNum, error) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for MarketDefinition.
func (m MarketDefinition) ApplLastSeqNum() (field.ApplLastSeqNum, error) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for MarketDefinition.
func (m MarketDefinition) ApplResendFlag() (field.ApplResendFlag, error) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}
