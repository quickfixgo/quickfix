package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TradeCaptureReportAck msg type = AR.
type TradeCaptureReportAck struct {
	message.Message
}

//TradeCaptureReportAckBuilder builds TradeCaptureReportAck messages.
type TradeCaptureReportAckBuilder struct {
	message.MessageBuilder
}

//NewTradeCaptureReportAckBuilder returns an initialized TradeCaptureReportAckBuilder with specified required fields.
func NewTradeCaptureReportAckBuilder(
	nosides field.NoSides) *TradeCaptureReportAckBuilder {
	builder := new(TradeCaptureReportAckBuilder)
	builder.Body.Set(nosides)
	return builder
}

//TradeReportID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TradeReportID() (*field.TradeReportID, error) {
	f := new(field.TradeReportID)
	err := m.Body.Get(f)
	return f, err
}

//TradeReportTransType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TradeReportTransType() (*field.TradeReportTransType, error) {
	f := new(field.TradeReportTransType)
	err := m.Body.Get(f)
	return f, err
}

//TradeReportType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TradeReportType() (*field.TradeReportType, error) {
	f := new(field.TradeReportType)
	err := m.Body.Get(f)
	return f, err
}

//TrdType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TrdType() (*field.TrdType, error) {
	f := new(field.TrdType)
	err := m.Body.Get(f)
	return f, err
}

//TrdSubType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TrdSubType() (*field.TrdSubType, error) {
	f := new(field.TrdSubType)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryTrdType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecondaryTrdType() (*field.SecondaryTrdType, error) {
	f := new(field.SecondaryTrdType)
	err := m.Body.Get(f)
	return f, err
}

//TransferReason is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TransferReason() (*field.TransferReason, error) {
	f := new(field.TransferReason)
	err := m.Body.Get(f)
	return f, err
}

//ExecType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) ExecType() (*field.ExecType, error) {
	f := new(field.ExecType)
	err := m.Body.Get(f)
	return f, err
}

//TradeReportRefID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TradeReportRefID() (*field.TradeReportRefID, error) {
	f := new(field.TradeReportRefID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryTradeReportRefID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecondaryTradeReportRefID() (*field.SecondaryTradeReportRefID, error) {
	f := new(field.SecondaryTradeReportRefID)
	err := m.Body.Get(f)
	return f, err
}

//TrdRptStatus is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TrdRptStatus() (*field.TrdRptStatus, error) {
	f := new(field.TrdRptStatus)
	err := m.Body.Get(f)
	return f, err
}

//TradeReportRejectReason is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TradeReportRejectReason() (*field.TradeReportRejectReason, error) {
	f := new(field.TradeReportRejectReason)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryTradeReportID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecondaryTradeReportID() (*field.SecondaryTradeReportID, error) {
	f := new(field.SecondaryTradeReportID)
	err := m.Body.Get(f)
	return f, err
}

//SubscriptionRequestType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//TradeLinkID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TradeLinkID() (*field.TradeLinkID, error) {
	f := new(field.TradeLinkID)
	err := m.Body.Get(f)
	return f, err
}

//TrdMatchID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TrdMatchID() (*field.TrdMatchID, error) {
	f := new(field.TrdMatchID)
	err := m.Body.Get(f)
	return f, err
}

//ExecID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryExecID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecondaryExecID() (*field.SecondaryExecID, error) {
	f := new(field.SecondaryExecID)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//SecurityGroup is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLLen is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXML is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLSchema is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//ProductComplex is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SettlMethod is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//ExerciseStyle is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//OptPayAmount is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) OptPayAmount() (*field.OptPayAmount, error) {
	f := new(field.OptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//PriceQuoteMethod is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//ListMethod is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//CapPrice is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//FloorPrice is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//FlexibleIndicator is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FuturesValuationMethod is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) FuturesValuationMethod() (*field.FuturesValuationMethod, error) {
	f := new(field.FuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) NoTrdRegTimestamps() (*field.NoTrdRegTimestamps, error) {
	f := new(field.NoTrdRegTimestamps)
	err := m.Body.Get(f)
	return f, err
}

//ResponseTransportType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) ResponseTransportType() (*field.ResponseTransportType, error) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}

//ResponseDestination is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) ResponseDestination() (*field.ResponseDestination, error) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//ClearingFeeIndicator is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) ClearingFeeIndicator() (*field.ClearingFeeIndicator, error) {
	f := new(field.ClearingFeeIndicator)
	err := m.Body.Get(f)
	return f, err
}

//OrdStatus is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) OrdStatus() (*field.OrdStatus, error) {
	f := new(field.OrdStatus)
	err := m.Body.Get(f)
	return f, err
}

//ExecRestatementReason is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) ExecRestatementReason() (*field.ExecRestatementReason, error) {
	f := new(field.ExecRestatementReason)
	err := m.Body.Get(f)
	return f, err
}

//PreviouslyReported is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) PreviouslyReported() (*field.PreviouslyReported, error) {
	f := new(field.PreviouslyReported)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingTradingSessionID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) UnderlyingTradingSessionID() (*field.UnderlyingTradingSessionID, error) {
	f := new(field.UnderlyingTradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//QtyType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingTradingSessionSubID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) UnderlyingTradingSessionSubID() (*field.UnderlyingTradingSessionSubID, error) {
	f := new(field.UnderlyingTradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//LastQty is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) LastQty() (*field.LastQty, error) {
	f := new(field.LastQty)
	err := m.Body.Get(f)
	return f, err
}

//LastPx is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}

//LastParPx is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) LastParPx() (*field.LastParPx, error) {
	f := new(field.LastParPx)
	err := m.Body.Get(f)
	return f, err
}

//LastSpotRate is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) LastSpotRate() (*field.LastSpotRate, error) {
	f := new(field.LastSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//LastForwardPoints is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) LastForwardPoints() (*field.LastForwardPoints, error) {
	f := new(field.LastForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//LastMkt is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//ClearingBusinessDate is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//AvgPx is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//AvgPxIndicator is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) AvgPxIndicator() (*field.AvgPxIndicator, error) {
	f := new(field.AvgPxIndicator)
	err := m.Body.Get(f)
	return f, err
}

//MultiLegReportingType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) MultiLegReportingType() (*field.MultiLegReportingType, error) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}

//TradeLegRefID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TradeLegRefID() (*field.TradeLegRefID, error) {
	f := new(field.TradeLegRefID)
	err := m.Body.Get(f)
	return f, err
}

//SettlType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//MatchStatus is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) MatchStatus() (*field.MatchStatus, error) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}

//MatchType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) MatchType() (*field.MatchType, error) {
	f := new(field.MatchType)
	err := m.Body.Get(f)
	return f, err
}

//CopyMsgIndicator is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) CopyMsgIndicator() (*field.CopyMsgIndicator, error) {
	f := new(field.CopyMsgIndicator)
	err := m.Body.Get(f)
	return f, err
}

//PublishTrdIndicator is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) PublishTrdIndicator() (*field.PublishTrdIndicator, error) {
	f := new(field.PublishTrdIndicator)
	err := m.Body.Get(f)
	return f, err
}

//ShortSaleReason is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) ShortSaleReason() (*field.ShortSaleReason, error) {
	f := new(field.ShortSaleReason)
	err := m.Body.Get(f)
	return f, err
}

//SettlDate is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//SettlSessID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//SettlSessSubID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//NoPosAmt is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) NoPosAmt() (*field.NoPosAmt, error) {
	f := new(field.NoPosAmt)
	err := m.Body.Get(f)
	return f, err
}

//TierCode is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TierCode() (*field.TierCode, error) {
	f := new(field.TierCode)
	err := m.Body.Get(f)
	return f, err
}

//MessageEventSource is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) MessageEventSource() (*field.MessageEventSource, error) {
	f := new(field.MessageEventSource)
	err := m.Body.Get(f)
	return f, err
}

//LastUpdateTime is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) LastUpdateTime() (*field.LastUpdateTime, error) {
	f := new(field.LastUpdateTime)
	err := m.Body.Get(f)
	return f, err
}

//RndPx is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) RndPx() (*field.RndPx, error) {
	f := new(field.RndPx)
	err := m.Body.Get(f)
	return f, err
}

//NoSides is a required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) NoSides() (*field.NoSides, error) {
	f := new(field.NoSides)
	err := m.Body.Get(f)
	return f, err
}

//AsOfIndicator is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) AsOfIndicator() (*field.AsOfIndicator, error) {
	f := new(field.AsOfIndicator)
	err := m.Body.Get(f)
	return f, err
}

//TradeID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TradeID() (*field.TradeID, error) {
	f := new(field.TradeID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryTradeID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecondaryTradeID() (*field.SecondaryTradeID, error) {
	f := new(field.SecondaryTradeID)
	err := m.Body.Get(f)
	return f, err
}

//FirmTradeID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) FirmTradeID() (*field.FirmTradeID, error) {
	f := new(field.FirmTradeID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryFirmTradeID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SecondaryFirmTradeID() (*field.SecondaryFirmTradeID, error) {
	f := new(field.SecondaryFirmTradeID)
	err := m.Body.Get(f)
	return f, err
}

//CalculatedCcyLastQty is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) CalculatedCcyLastQty() (*field.CalculatedCcyLastQty, error) {
	f := new(field.CalculatedCcyLastQty)
	err := m.Body.Get(f)
	return f, err
}

//LastSwapPoints is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) LastSwapPoints() (*field.LastSwapPoints, error) {
	f := new(field.LastSwapPoints)
	err := m.Body.Get(f)
	return f, err
}

//GrossTradeAmt is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) GrossTradeAmt() (*field.GrossTradeAmt, error) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}

//NoRootPartyIDs is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) NoRootPartyIDs() (*field.NoRootPartyIDs, error) {
	f := new(field.NoRootPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//TradeHandlingInstr is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TradeHandlingInstr() (*field.TradeHandlingInstr, error) {
	f := new(field.TradeHandlingInstr)
	err := m.Body.Get(f)
	return f, err
}

//OrigTradeHandlingInstr is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) OrigTradeHandlingInstr() (*field.OrigTradeHandlingInstr, error) {
	f := new(field.OrigTradeHandlingInstr)
	err := m.Body.Get(f)
	return f, err
}

//OrigTradeDate is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) OrigTradeDate() (*field.OrigTradeDate, error) {
	f := new(field.OrigTradeDate)
	err := m.Body.Get(f)
	return f, err
}

//OrigTradeID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) OrigTradeID() (*field.OrigTradeID, error) {
	f := new(field.OrigTradeID)
	err := m.Body.Get(f)
	return f, err
}

//OrigSecondaryTradeID is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) OrigSecondaryTradeID() (*field.OrigSecondaryTradeID, error) {
	f := new(field.OrigSecondaryTradeID)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//RptSys is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) RptSys() (*field.RptSys, error) {
	f := new(field.RptSys)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrency is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//FeeMultiplier is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) FeeMultiplier() (*field.FeeMultiplier, error) {
	f := new(field.FeeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//NoTrdRepIndicators is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) NoTrdRepIndicators() (*field.NoTrdRepIndicators, error) {
	f := new(field.NoTrdRepIndicators)
	err := m.Body.Get(f)
	return f, err
}

//TradePublishIndicator is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) TradePublishIndicator() (*field.TradePublishIndicator, error) {
	f := new(field.TradePublishIndicator)
	err := m.Body.Get(f)
	return f, err
}
