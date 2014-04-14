package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TradeCaptureReportRequestAck msg type = AQ.
type TradeCaptureReportRequestAck struct {
	message.Message
}

//TradeCaptureReportRequestAckBuilder builds TradeCaptureReportRequestAck messages.
type TradeCaptureReportRequestAckBuilder struct {
	message.MessageBuilder
}

//NewTradeCaptureReportRequestAckBuilder returns an initialized TradeCaptureReportRequestAckBuilder with specified required fields.
func NewTradeCaptureReportRequestAckBuilder(
	traderequestid field.TradeRequestID,
	traderequesttype field.TradeRequestType,
	traderequestresult field.TradeRequestResult,
	traderequeststatus field.TradeRequestStatus) *TradeCaptureReportRequestAckBuilder {
	builder := new(TradeCaptureReportRequestAckBuilder)
	builder.Body.Set(traderequestid)
	builder.Body.Set(traderequesttype)
	builder.Body.Set(traderequestresult)
	builder.Body.Set(traderequeststatus)
	return builder
}

//TradeRequestID is a required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) TradeRequestID() (*field.TradeRequestID, error) {
	f := new(field.TradeRequestID)
	err := m.Body.Get(f)
	return f, err
}

//TradeRequestType is a required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) TradeRequestType() (*field.TradeRequestType, error) {
	f := new(field.TradeRequestType)
	err := m.Body.Get(f)
	return f, err
}

//SubscriptionRequestType is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//TotNumTradeReports is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) TotNumTradeReports() (*field.TotNumTradeReports, error) {
	f := new(field.TotNumTradeReports)
	err := m.Body.Get(f)
	return f, err
}

//TradeRequestResult is a required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) TradeRequestResult() (*field.TradeRequestResult, error) {
	f := new(field.TradeRequestResult)
	err := m.Body.Get(f)
	return f, err
}

//TradeRequestStatus is a required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) TradeRequestStatus() (*field.TradeRequestStatus, error) {
	f := new(field.TradeRequestStatus)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//SecurityGroup is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLLen is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXML is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLSchema is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//ProductComplex is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SettlMethod is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//ExerciseStyle is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//OptPayoutAmount is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) OptPayoutAmount() (*field.OptPayoutAmount, error) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}

//PriceQuoteMethod is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//ListMethod is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//CapPrice is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//FloorPrice is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//FlexibleIndicator is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//ValuationMethod is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) ValuationMethod() (*field.ValuationMethod, error) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) ContractMultiplierUnit() (*field.ContractMultiplierUnit, error) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//FlowScheduleType is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) FlowScheduleType() (*field.FlowScheduleType, error) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//RestructuringType is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) RestructuringType() (*field.RestructuringType, error) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}

//Seniority is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) Seniority() (*field.Seniority, error) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, error) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, error) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//AttachmentPoint is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) AttachmentPoint() (*field.AttachmentPoint, error) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//DetachmentPoint is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) DetachmentPoint() (*field.DetachmentPoint, error) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, error) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, error) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, error) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, error) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//OptPayoutType is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) OptPayoutType() (*field.OptPayoutType, error) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}

//NoComplexEvents is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) NoComplexEvents() (*field.NoComplexEvents, error) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//MultiLegReportingType is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) MultiLegReportingType() (*field.MultiLegReportingType, error) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}

//ResponseTransportType is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) ResponseTransportType() (*field.ResponseTransportType, error) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}

//ResponseDestination is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) ResponseDestination() (*field.ResponseDestination, error) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//MessageEventSource is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) MessageEventSource() (*field.MessageEventSource, error) {
	f := new(field.MessageEventSource)
	err := m.Body.Get(f)
	return f, err
}

//TradeID is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) TradeID() (*field.TradeID, error) {
	f := new(field.TradeID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryTradeID is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SecondaryTradeID() (*field.SecondaryTradeID, error) {
	f := new(field.SecondaryTradeID)
	err := m.Body.Get(f)
	return f, err
}

//FirmTradeID is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) FirmTradeID() (*field.FirmTradeID, error) {
	f := new(field.FirmTradeID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryFirmTradeID is a non-required field for TradeCaptureReportRequestAck.
func (m *TradeCaptureReportRequestAck) SecondaryFirmTradeID() (*field.SecondaryFirmTradeID, error) {
	f := new(field.SecondaryFirmTradeID)
	err := m.Body.Get(f)
	return f, err
}
