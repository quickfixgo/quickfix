package fix50sp2

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type MarketDefinition struct {
	quickfix.Message
}

func (m *MarketDefinition) MarketReqID() (*field.MarketReqID, error) {
	f := new(field.MarketReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) EncodedMktSegmDesc() (*field.EncodedMktSegmDesc, error) {
	f := new(field.EncodedMktSegmDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) ImpliedMarketIndicator() (*field.ImpliedMarketIndicator, error) {
	f := new(field.ImpliedMarketIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) NoLotTypeRules() (*field.NoLotTypeRules, error) {
	f := new(field.NoLotTypeRules)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) NoExecInstRules() (*field.NoExecInstRules, error) {
	f := new(field.NoExecInstRules)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) RoundLot() (*field.RoundLot, error) {
	f := new(field.RoundLot)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) NoOrdTypeRules() (*field.NoOrdTypeRules, error) {
	f := new(field.NoOrdTypeRules)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) MarketReportID() (*field.MarketReportID, error) {
	f := new(field.MarketReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) MarketSegmentDesc() (*field.MarketSegmentDesc, error) {
	f := new(field.MarketSegmentDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) EncodedMktSegmDescLen() (*field.EncodedMktSegmDescLen, error) {
	f := new(field.EncodedMktSegmDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) ExpirationCycle() (*field.ExpirationCycle, error) {
	f := new(field.ExpirationCycle)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) TradingCurrency() (*field.TradingCurrency, error) {
	f := new(field.TradingCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) MultilegPriceMethod() (*field.MultilegPriceMethod, error) {
	f := new(field.MultilegPriceMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) NoTickRules() (*field.NoTickRules, error) {
	f := new(field.NoTickRules)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) PriceLimitType() (*field.PriceLimitType, error) {
	f := new(field.PriceLimitType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) MinTradeVol() (*field.MinTradeVol, error) {
	f := new(field.MinTradeVol)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) MultilegModel() (*field.MultilegModel, error) {
	f := new(field.MultilegModel)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) NoTimeInForceRules() (*field.NoTimeInForceRules, error) {
	f := new(field.NoTimeInForceRules)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) MaxTradeVol() (*field.MaxTradeVol, error) {
	f := new(field.MaxTradeVol)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) ParentMktSegmID() (*field.ParentMktSegmID, error) {
	f := new(field.ParentMktSegmID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) LowLimitPrice() (*field.LowLimitPrice, error) {
	f := new(field.LowLimitPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) HighLimitPrice() (*field.HighLimitPrice, error) {
	f := new(field.HighLimitPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) TradingReferencePrice() (*field.TradingReferencePrice, error) {
	f := new(field.TradingReferencePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) MaxPriceVariation() (*field.MaxPriceVariation, error) {
	f := new(field.MaxPriceVariation)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
