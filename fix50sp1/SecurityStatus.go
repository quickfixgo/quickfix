package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityStatus msg type = f.
type SecurityStatus struct {
	message.Message
}

//SecurityStatusBuilder builds SecurityStatus messages.
type SecurityStatusBuilder struct {
	message.MessageBuilder
}

//NewSecurityStatusBuilder returns an initialized SecurityStatusBuilder with specified required fields.
func NewSecurityStatusBuilder() *SecurityStatusBuilder {
	builder := new(SecurityStatusBuilder)
	return builder
}

//SecurityStatusReqID is a non-required field for SecurityStatus.
func (m *SecurityStatus) SecurityStatusReqID() (*field.SecurityStatusReqID, error) {
	f := new(field.SecurityStatusReqID)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for SecurityStatus.
func (m *SecurityStatus) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for SecurityStatus.
func (m *SecurityStatus) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for SecurityStatus.
func (m *SecurityStatus) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for SecurityStatus.
func (m *SecurityStatus) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for SecurityStatus.
func (m *SecurityStatus) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for SecurityStatus.
func (m *SecurityStatus) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for SecurityStatus.
func (m *SecurityStatus) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for SecurityStatus.
func (m *SecurityStatus) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for SecurityStatus.
func (m *SecurityStatus) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for SecurityStatus.
func (m *SecurityStatus) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for SecurityStatus.
func (m *SecurityStatus) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for SecurityStatus.
func (m *SecurityStatus) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for SecurityStatus.
func (m *SecurityStatus) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for SecurityStatus.
func (m *SecurityStatus) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for SecurityStatus.
func (m *SecurityStatus) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for SecurityStatus.
func (m *SecurityStatus) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for SecurityStatus.
func (m *SecurityStatus) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for SecurityStatus.
func (m *SecurityStatus) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for SecurityStatus.
func (m *SecurityStatus) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for SecurityStatus.
func (m *SecurityStatus) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for SecurityStatus.
func (m *SecurityStatus) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for SecurityStatus.
func (m *SecurityStatus) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for SecurityStatus.
func (m *SecurityStatus) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for SecurityStatus.
func (m *SecurityStatus) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for SecurityStatus.
func (m *SecurityStatus) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for SecurityStatus.
func (m *SecurityStatus) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for SecurityStatus.
func (m *SecurityStatus) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for SecurityStatus.
func (m *SecurityStatus) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for SecurityStatus.
func (m *SecurityStatus) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for SecurityStatus.
func (m *SecurityStatus) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for SecurityStatus.
func (m *SecurityStatus) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for SecurityStatus.
func (m *SecurityStatus) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for SecurityStatus.
func (m *SecurityStatus) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for SecurityStatus.
func (m *SecurityStatus) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for SecurityStatus.
func (m *SecurityStatus) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for SecurityStatus.
func (m *SecurityStatus) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for SecurityStatus.
func (m *SecurityStatus) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for SecurityStatus.
func (m *SecurityStatus) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for SecurityStatus.
func (m *SecurityStatus) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for SecurityStatus.
func (m *SecurityStatus) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for SecurityStatus.
func (m *SecurityStatus) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for SecurityStatus.
func (m *SecurityStatus) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for SecurityStatus.
func (m *SecurityStatus) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for SecurityStatus.
func (m *SecurityStatus) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for SecurityStatus.
func (m *SecurityStatus) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for SecurityStatus.
func (m *SecurityStatus) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for SecurityStatus.
func (m *SecurityStatus) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for SecurityStatus.
func (m *SecurityStatus) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for SecurityStatus.
func (m *SecurityStatus) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for SecurityStatus.
func (m *SecurityStatus) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for SecurityStatus.
func (m *SecurityStatus) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for SecurityStatus.
func (m *SecurityStatus) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for SecurityStatus.
func (m *SecurityStatus) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for SecurityStatus.
func (m *SecurityStatus) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//SecurityGroup is a non-required field for SecurityStatus.
func (m *SecurityStatus) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for SecurityStatus.
func (m *SecurityStatus) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for SecurityStatus.
func (m *SecurityStatus) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLLen is a non-required field for SecurityStatus.
func (m *SecurityStatus) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXML is a non-required field for SecurityStatus.
func (m *SecurityStatus) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLSchema is a non-required field for SecurityStatus.
func (m *SecurityStatus) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//ProductComplex is a non-required field for SecurityStatus.
func (m *SecurityStatus) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for SecurityStatus.
func (m *SecurityStatus) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for SecurityStatus.
func (m *SecurityStatus) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SettlMethod is a non-required field for SecurityStatus.
func (m *SecurityStatus) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//ExerciseStyle is a non-required field for SecurityStatus.
func (m *SecurityStatus) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//OptPayAmount is a non-required field for SecurityStatus.
func (m *SecurityStatus) OptPayAmount() (*field.OptPayAmount, error) {
	f := new(field.OptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//PriceQuoteMethod is a non-required field for SecurityStatus.
func (m *SecurityStatus) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//ListMethod is a non-required field for SecurityStatus.
func (m *SecurityStatus) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//CapPrice is a non-required field for SecurityStatus.
func (m *SecurityStatus) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//FloorPrice is a non-required field for SecurityStatus.
func (m *SecurityStatus) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for SecurityStatus.
func (m *SecurityStatus) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//FlexibleIndicator is a non-required field for SecurityStatus.
func (m *SecurityStatus) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for SecurityStatus.
func (m *SecurityStatus) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FuturesValuationMethod is a non-required field for SecurityStatus.
func (m *SecurityStatus) FuturesValuationMethod() (*field.FuturesValuationMethod, error) {
	f := new(field.FuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryForm is a non-required field for SecurityStatus.
func (m *SecurityStatus) DeliveryForm() (*field.DeliveryForm, error) {
	f := new(field.DeliveryForm)
	err := m.Body.Get(f)
	return f, err
}

//PctAtRisk is a non-required field for SecurityStatus.
func (m *SecurityStatus) PctAtRisk() (*field.PctAtRisk, error) {
	f := new(field.PctAtRisk)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrAttrib is a non-required field for SecurityStatus.
func (m *SecurityStatus) NoInstrAttrib() (*field.NoInstrAttrib, error) {
	f := new(field.NoInstrAttrib)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for SecurityStatus.
func (m *SecurityStatus) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for SecurityStatus.
func (m *SecurityStatus) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for SecurityStatus.
func (m *SecurityStatus) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for SecurityStatus.
func (m *SecurityStatus) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for SecurityStatus.
func (m *SecurityStatus) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//UnsolicitedIndicator is a non-required field for SecurityStatus.
func (m *SecurityStatus) UnsolicitedIndicator() (*field.UnsolicitedIndicator, error) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}

//SecurityTradingStatus is a non-required field for SecurityStatus.
func (m *SecurityStatus) SecurityTradingStatus() (*field.SecurityTradingStatus, error) {
	f := new(field.SecurityTradingStatus)
	err := m.Body.Get(f)
	return f, err
}

//FinancialStatus is a non-required field for SecurityStatus.
func (m *SecurityStatus) FinancialStatus() (*field.FinancialStatus, error) {
	f := new(field.FinancialStatus)
	err := m.Body.Get(f)
	return f, err
}

//CorporateAction is a non-required field for SecurityStatus.
func (m *SecurityStatus) CorporateAction() (*field.CorporateAction, error) {
	f := new(field.CorporateAction)
	err := m.Body.Get(f)
	return f, err
}

//HaltReasonChar is a non-required field for SecurityStatus.
func (m *SecurityStatus) HaltReasonChar() (*field.HaltReasonChar, error) {
	f := new(field.HaltReasonChar)
	err := m.Body.Get(f)
	return f, err
}

//InViewOfCommon is a non-required field for SecurityStatus.
func (m *SecurityStatus) InViewOfCommon() (*field.InViewOfCommon, error) {
	f := new(field.InViewOfCommon)
	err := m.Body.Get(f)
	return f, err
}

//DueToRelated is a non-required field for SecurityStatus.
func (m *SecurityStatus) DueToRelated() (*field.DueToRelated, error) {
	f := new(field.DueToRelated)
	err := m.Body.Get(f)
	return f, err
}

//BuyVolume is a non-required field for SecurityStatus.
func (m *SecurityStatus) BuyVolume() (*field.BuyVolume, error) {
	f := new(field.BuyVolume)
	err := m.Body.Get(f)
	return f, err
}

//SellVolume is a non-required field for SecurityStatus.
func (m *SecurityStatus) SellVolume() (*field.SellVolume, error) {
	f := new(field.SellVolume)
	err := m.Body.Get(f)
	return f, err
}

//HighPx is a non-required field for SecurityStatus.
func (m *SecurityStatus) HighPx() (*field.HighPx, error) {
	f := new(field.HighPx)
	err := m.Body.Get(f)
	return f, err
}

//LowPx is a non-required field for SecurityStatus.
func (m *SecurityStatus) LowPx() (*field.LowPx, error) {
	f := new(field.LowPx)
	err := m.Body.Get(f)
	return f, err
}

//LastPx is a non-required field for SecurityStatus.
func (m *SecurityStatus) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for SecurityStatus.
func (m *SecurityStatus) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//Adjustment is a non-required field for SecurityStatus.
func (m *SecurityStatus) Adjustment() (*field.Adjustment, error) {
	f := new(field.Adjustment)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for SecurityStatus.
func (m *SecurityStatus) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for SecurityStatus.
func (m *SecurityStatus) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for SecurityStatus.
func (m *SecurityStatus) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//FirstPx is a non-required field for SecurityStatus.
func (m *SecurityStatus) FirstPx() (*field.FirstPx, error) {
	f := new(field.FirstPx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityTradingEvent is a non-required field for SecurityStatus.
func (m *SecurityStatus) SecurityTradingEvent() (*field.SecurityTradingEvent, error) {
	f := new(field.SecurityTradingEvent)
	err := m.Body.Get(f)
	return f, err
}

//MDBookType is a non-required field for SecurityStatus.
func (m *SecurityStatus) MDBookType() (*field.MDBookType, error) {
	f := new(field.MDBookType)
	err := m.Body.Get(f)
	return f, err
}

//MarketDepth is a non-required field for SecurityStatus.
func (m *SecurityStatus) MarketDepth() (*field.MarketDepth, error) {
	f := new(field.MarketDepth)
	err := m.Body.Get(f)
	return f, err
}

//MarketID is a non-required field for SecurityStatus.
func (m *SecurityStatus) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//MarketSegmentID is a non-required field for SecurityStatus.
func (m *SecurityStatus) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//ApplID is a non-required field for SecurityStatus.
func (m *SecurityStatus) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//ApplSeqNum is a non-required field for SecurityStatus.
func (m *SecurityStatus) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplLastSeqNum is a non-required field for SecurityStatus.
func (m *SecurityStatus) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplResendFlag is a non-required field for SecurityStatus.
func (m *SecurityStatus) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
