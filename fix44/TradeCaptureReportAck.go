package fix44

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

//CreateTradeCaptureReportAckBuilder returns an initialized TradeCaptureReportAckBuilder with specified required fields.
func CreateTradeCaptureReportAckBuilder(
	tradereportid field.TradeReportID,
	exectype field.ExecType) TradeCaptureReportAckBuilder {
	var builder TradeCaptureReportAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(tradereportid)
	builder.Body.Set(exectype)
	return builder
}

//TradeReportID is a required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportID() (field.TradeReportID, error) {
	var f field.TradeReportID
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportTransType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportTransType() (field.TradeReportTransType, error) {
	var f field.TradeReportTransType
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportType() (field.TradeReportType, error) {
	var f field.TradeReportType
	err := m.Body.Get(&f)
	return f, err
}

//TrdType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdType() (field.TrdType, error) {
	var f field.TrdType
	err := m.Body.Get(&f)
	return f, err
}

//TrdSubType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdSubType() (field.TrdSubType, error) {
	var f field.TrdSubType
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTrdType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryTrdType() (field.SecondaryTrdType, error) {
	var f field.SecondaryTrdType
	err := m.Body.Get(&f)
	return f, err
}

//TransferReason is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TransferReason() (field.TransferReason, error) {
	var f field.TransferReason
	err := m.Body.Get(&f)
	return f, err
}

//ExecType is a required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ExecType() (field.ExecType, error) {
	var f field.ExecType
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportRefID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportRefID() (field.TradeReportRefID, error) {
	var f field.TradeReportRefID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeReportRefID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryTradeReportRefID() (field.SecondaryTradeReportRefID, error) {
	var f field.SecondaryTradeReportRefID
	err := m.Body.Get(&f)
	return f, err
}

//TrdRptStatus is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdRptStatus() (field.TrdRptStatus, error) {
	var f field.TrdRptStatus
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportRejectReason is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportRejectReason() (field.TradeReportRejectReason, error) {
	var f field.TradeReportRejectReason
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeReportID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryTradeReportID() (field.SecondaryTradeReportID, error) {
	var f field.SecondaryTradeReportID
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//TradeLinkID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeLinkID() (field.TradeLinkID, error) {
	var f field.TradeLinkID
	err := m.Body.Get(&f)
	return f, err
}

//TrdMatchID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdMatchID() (field.TrdMatchID, error) {
	var f field.TrdMatchID
	err := m.Body.Get(&f)
	return f, err
}

//ExecID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ExecID() (field.ExecID, error) {
	var f field.ExecID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryExecID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryExecID() (field.SecondaryExecID, error) {
	var f field.SecondaryExecID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoTrdRegTimestamps() (field.NoTrdRegTimestamps, error) {
	var f field.NoTrdRegTimestamps
	err := m.Body.Get(&f)
	return f, err
}

//ResponseTransportType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ResponseTransportType() (field.ResponseTransportType, error) {
	var f field.ResponseTransportType
	err := m.Body.Get(&f)
	return f, err
}

//ResponseDestination is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ResponseDestination() (field.ResponseDestination, error) {
	var f field.ResponseDestination
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//ClearingFeeIndicator is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ClearingFeeIndicator() (field.ClearingFeeIndicator, error) {
	var f field.ClearingFeeIndicator
	err := m.Body.Get(&f)
	return f, err
}

//OrderCapacity is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OrderCapacity() (field.OrderCapacity, error) {
	var f field.OrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//OrderRestrictions is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OrderRestrictions() (field.OrderRestrictions, error) {
	var f field.OrderRestrictions
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderCapacity is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CustOrderCapacity() (field.CustOrderCapacity, error) {
	var f field.CustOrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) AcctIDSource() (field.AcctIDSource, error) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//PositionEffect is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PositionEffect() (field.PositionEffect, error) {
	var f field.PositionEffect
	err := m.Body.Get(&f)
	return f, err
}

//PreallocMethod is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PreallocMethod() (field.PreallocMethod, error) {
	var f field.PreallocMethod
	err := m.Body.Get(&f)
	return f, err
}

//NoAllocs is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoAllocs() (field.NoAllocs, error) {
	var f field.NoAllocs
	err := m.Body.Get(&f)
	return f, err
}
