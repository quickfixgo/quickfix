package fix50sp2

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

//NewMarketDefinitionBuilder returns an initialized MarketDefinitionBuilder with specified required fields.
func NewMarketDefinitionBuilder(
	marketreportid field.MarketReportID,
	marketid field.MarketID) *MarketDefinitionBuilder {
	builder := new(MarketDefinitionBuilder)
	builder.Body.Set(marketreportid)
	builder.Body.Set(marketid)
	return builder
}

//MarketReportID is a required field for MarketDefinition.
func (m *MarketDefinition) MarketReportID() (*field.MarketReportID, error) {
	f := new(field.MarketReportID)
	err := m.Body.Get(f)
	return f, err
}

//MarketReqID is a non-required field for MarketDefinition.
func (m *MarketDefinition) MarketReqID() (*field.MarketReqID, error) {
	f := new(field.MarketReqID)
	err := m.Body.Get(f)
	return f, err
}

//MarketID is a required field for MarketDefinition.
func (m *MarketDefinition) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//MarketSegmentID is a non-required field for MarketDefinition.
func (m *MarketDefinition) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//MarketSegmentDesc is a non-required field for MarketDefinition.
func (m *MarketDefinition) MarketSegmentDesc() (*field.MarketSegmentDesc, error) {
	f := new(field.MarketSegmentDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedMktSegmDescLen is a non-required field for MarketDefinition.
func (m *MarketDefinition) EncodedMktSegmDescLen() (*field.EncodedMktSegmDescLen, error) {
	f := new(field.EncodedMktSegmDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedMktSegmDesc is a non-required field for MarketDefinition.
func (m *MarketDefinition) EncodedMktSegmDesc() (*field.EncodedMktSegmDesc, error) {
	f := new(field.EncodedMktSegmDesc)
	err := m.Body.Get(f)
	return f, err
}

//ParentMktSegmID is a non-required field for MarketDefinition.
func (m *MarketDefinition) ParentMktSegmID() (*field.ParentMktSegmID, error) {
	f := new(field.ParentMktSegmID)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for MarketDefinition.
func (m *MarketDefinition) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//NoTickRules is a non-required field for MarketDefinition.
func (m *MarketDefinition) NoTickRules() (*field.NoTickRules, error) {
	f := new(field.NoTickRules)
	err := m.Body.Get(f)
	return f, err
}

//NoLotTypeRules is a non-required field for MarketDefinition.
func (m *MarketDefinition) NoLotTypeRules() (*field.NoLotTypeRules, error) {
	f := new(field.NoLotTypeRules)
	err := m.Body.Get(f)
	return f, err
}

//PriceLimitType is a non-required field for MarketDefinition.
func (m *MarketDefinition) PriceLimitType() (*field.PriceLimitType, error) {
	f := new(field.PriceLimitType)
	err := m.Body.Get(f)
	return f, err
}

//LowLimitPrice is a non-required field for MarketDefinition.
func (m *MarketDefinition) LowLimitPrice() (*field.LowLimitPrice, error) {
	f := new(field.LowLimitPrice)
	err := m.Body.Get(f)
	return f, err
}

//HighLimitPrice is a non-required field for MarketDefinition.
func (m *MarketDefinition) HighLimitPrice() (*field.HighLimitPrice, error) {
	f := new(field.HighLimitPrice)
	err := m.Body.Get(f)
	return f, err
}

//TradingReferencePrice is a non-required field for MarketDefinition.
func (m *MarketDefinition) TradingReferencePrice() (*field.TradingReferencePrice, error) {
	f := new(field.TradingReferencePrice)
	err := m.Body.Get(f)
	return f, err
}

//ExpirationCycle is a non-required field for MarketDefinition.
func (m *MarketDefinition) ExpirationCycle() (*field.ExpirationCycle, error) {
	f := new(field.ExpirationCycle)
	err := m.Body.Get(f)
	return f, err
}

//MinTradeVol is a non-required field for MarketDefinition.
func (m *MarketDefinition) MinTradeVol() (*field.MinTradeVol, error) {
	f := new(field.MinTradeVol)
	err := m.Body.Get(f)
	return f, err
}

//MaxTradeVol is a non-required field for MarketDefinition.
func (m *MarketDefinition) MaxTradeVol() (*field.MaxTradeVol, error) {
	f := new(field.MaxTradeVol)
	err := m.Body.Get(f)
	return f, err
}

//MaxPriceVariation is a non-required field for MarketDefinition.
func (m *MarketDefinition) MaxPriceVariation() (*field.MaxPriceVariation, error) {
	f := new(field.MaxPriceVariation)
	err := m.Body.Get(f)
	return f, err
}

//ImpliedMarketIndicator is a non-required field for MarketDefinition.
func (m *MarketDefinition) ImpliedMarketIndicator() (*field.ImpliedMarketIndicator, error) {
	f := new(field.ImpliedMarketIndicator)
	err := m.Body.Get(f)
	return f, err
}

//TradingCurrency is a non-required field for MarketDefinition.
func (m *MarketDefinition) TradingCurrency() (*field.TradingCurrency, error) {
	f := new(field.TradingCurrency)
	err := m.Body.Get(f)
	return f, err
}

//RoundLot is a non-required field for MarketDefinition.
func (m *MarketDefinition) RoundLot() (*field.RoundLot, error) {
	f := new(field.RoundLot)
	err := m.Body.Get(f)
	return f, err
}

//MultilegModel is a non-required field for MarketDefinition.
func (m *MarketDefinition) MultilegModel() (*field.MultilegModel, error) {
	f := new(field.MultilegModel)
	err := m.Body.Get(f)
	return f, err
}

//MultilegPriceMethod is a non-required field for MarketDefinition.
func (m *MarketDefinition) MultilegPriceMethod() (*field.MultilegPriceMethod, error) {
	f := new(field.MultilegPriceMethod)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for MarketDefinition.
func (m *MarketDefinition) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//NoOrdTypeRules is a non-required field for MarketDefinition.
func (m *MarketDefinition) NoOrdTypeRules() (*field.NoOrdTypeRules, error) {
	f := new(field.NoOrdTypeRules)
	err := m.Body.Get(f)
	return f, err
}

//NoTimeInForceRules is a non-required field for MarketDefinition.
func (m *MarketDefinition) NoTimeInForceRules() (*field.NoTimeInForceRules, error) {
	f := new(field.NoTimeInForceRules)
	err := m.Body.Get(f)
	return f, err
}

//NoExecInstRules is a non-required field for MarketDefinition.
func (m *MarketDefinition) NoExecInstRules() (*field.NoExecInstRules, error) {
	f := new(field.NoExecInstRules)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for MarketDefinition.
func (m *MarketDefinition) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for MarketDefinition.
func (m *MarketDefinition) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for MarketDefinition.
func (m *MarketDefinition) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for MarketDefinition.
func (m *MarketDefinition) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//ApplID is a non-required field for MarketDefinition.
func (m *MarketDefinition) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//ApplSeqNum is a non-required field for MarketDefinition.
func (m *MarketDefinition) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplLastSeqNum is a non-required field for MarketDefinition.
func (m *MarketDefinition) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplResendFlag is a non-required field for MarketDefinition.
func (m *MarketDefinition) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
