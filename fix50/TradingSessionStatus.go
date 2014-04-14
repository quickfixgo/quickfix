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

//NewTradingSessionStatusBuilder returns an initialized TradingSessionStatusBuilder with specified required fields.
func NewTradingSessionStatusBuilder(
	tradingsessionid field.TradingSessionID,
	tradsesstatus field.TradSesStatus) *TradingSessionStatusBuilder {
	builder := new(TradingSessionStatusBuilder)
	builder.Body.Set(tradingsessionid)
	builder.Body.Set(tradsesstatus)
	return builder
}

//TradSesReqID is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) TradSesReqID() (*field.TradSesReqID, error) {
	f := new(field.TradSesReqID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a required field for TradingSessionStatus.
func (m *TradingSessionStatus) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//TradSesMethod is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) TradSesMethod() (*field.TradSesMethod, error) {
	f := new(field.TradSesMethod)
	err := m.Body.Get(f)
	return f, err
}

//TradSesMode is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) TradSesMode() (*field.TradSesMode, error) {
	f := new(field.TradSesMode)
	err := m.Body.Get(f)
	return f, err
}

//UnsolicitedIndicator is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) UnsolicitedIndicator() (*field.UnsolicitedIndicator, error) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}

//TradSesStatus is a required field for TradingSessionStatus.
func (m *TradingSessionStatus) TradSesStatus() (*field.TradSesStatus, error) {
	f := new(field.TradSesStatus)
	err := m.Body.Get(f)
	return f, err
}

//TradSesStatusRejReason is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) TradSesStatusRejReason() (*field.TradSesStatusRejReason, error) {
	f := new(field.TradSesStatusRejReason)
	err := m.Body.Get(f)
	return f, err
}

//TradSesStartTime is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) TradSesStartTime() (*field.TradSesStartTime, error) {
	f := new(field.TradSesStartTime)
	err := m.Body.Get(f)
	return f, err
}

//TradSesOpenTime is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) TradSesOpenTime() (*field.TradSesOpenTime, error) {
	f := new(field.TradSesOpenTime)
	err := m.Body.Get(f)
	return f, err
}

//TradSesPreCloseTime is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) TradSesPreCloseTime() (*field.TradSesPreCloseTime, error) {
	f := new(field.TradSesPreCloseTime)
	err := m.Body.Get(f)
	return f, err
}

//TradSesCloseTime is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) TradSesCloseTime() (*field.TradSesCloseTime, error) {
	f := new(field.TradSesCloseTime)
	err := m.Body.Get(f)
	return f, err
}

//TradSesEndTime is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) TradSesEndTime() (*field.TradSesEndTime, error) {
	f := new(field.TradSesEndTime)
	err := m.Body.Get(f)
	return f, err
}

//TotalVolumeTraded is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) TotalVolumeTraded() (*field.TotalVolumeTraded, error) {
	f := new(field.TotalVolumeTraded)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for TradingSessionStatus.
func (m *TradingSessionStatus) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}
