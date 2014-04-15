package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TradingSessionStatus msg type = h.
type TradingSessionStatus struct {
	message.Message
}

//TradingSessionStatusBuilder builds TradingSessionStatus messages.
type TradingSessionStatusBuilder struct {
	message.MessageBuilder
}

//CreateTradingSessionStatusBuilder returns an initialized TradingSessionStatusBuilder with specified required fields.
func CreateTradingSessionStatusBuilder(
	tradingsessionid field.TradingSessionID,
	tradsesstatus field.TradSesStatus) TradingSessionStatusBuilder {
	var builder TradingSessionStatusBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(tradingsessionid)
	builder.Body.Set(tradsesstatus)
	return builder
}

//TradSesReqID is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesReqID() (field.TradSesReqID, error) {
	var f field.TradSesReqID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a required field for TradingSessionStatus.
func (m TradingSessionStatus) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//TradSesMethod is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesMethod() (field.TradSesMethod, error) {
	var f field.TradSesMethod
	err := m.Body.Get(&f)
	return f, err
}

//TradSesMode is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesMode() (field.TradSesMode, error) {
	var f field.TradSesMode
	err := m.Body.Get(&f)
	return f, err
}

//UnsolicitedIndicator is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) UnsolicitedIndicator() (field.UnsolicitedIndicator, error) {
	var f field.UnsolicitedIndicator
	err := m.Body.Get(&f)
	return f, err
}

//TradSesStatus is a required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesStatus() (field.TradSesStatus, error) {
	var f field.TradSesStatus
	err := m.Body.Get(&f)
	return f, err
}

//TradSesStatusRejReason is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesStatusRejReason() (field.TradSesStatusRejReason, error) {
	var f field.TradSesStatusRejReason
	err := m.Body.Get(&f)
	return f, err
}

//TradSesStartTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesStartTime() (field.TradSesStartTime, error) {
	var f field.TradSesStartTime
	err := m.Body.Get(&f)
	return f, err
}

//TradSesOpenTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesOpenTime() (field.TradSesOpenTime, error) {
	var f field.TradSesOpenTime
	err := m.Body.Get(&f)
	return f, err
}

//TradSesPreCloseTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesPreCloseTime() (field.TradSesPreCloseTime, error) {
	var f field.TradSesPreCloseTime
	err := m.Body.Get(&f)
	return f, err
}

//TradSesCloseTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesCloseTime() (field.TradSesCloseTime, error) {
	var f field.TradSesCloseTime
	err := m.Body.Get(&f)
	return f, err
}

//TradSesEndTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesEndTime() (field.TradSesEndTime, error) {
	var f field.TradSesEndTime
	err := m.Body.Get(&f)
	return f, err
}

//TotalVolumeTraded is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TotalVolumeTraded() (field.TotalVolumeTraded, error) {
	var f field.TotalVolumeTraded
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}
