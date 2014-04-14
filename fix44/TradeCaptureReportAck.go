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

//NewTradeCaptureReportAckBuilder returns an initialized TradeCaptureReportAckBuilder with specified required fields.
func NewTradeCaptureReportAckBuilder(
	tradereportid field.TradeReportID,
	exectype field.ExecType) *TradeCaptureReportAckBuilder {
	builder := new(TradeCaptureReportAckBuilder)
	builder.Body.Set(tradereportid)
	builder.Body.Set(exectype)
	return builder
}

//TradeReportID is a required field for TradeCaptureReportAck.
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

//ExecType is a required field for TradeCaptureReportAck.
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

//OrderCapacity is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) OrderCapacity() (*field.OrderCapacity, error) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//OrderRestrictions is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) OrderRestrictions() (*field.OrderRestrictions, error) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}

//CustOrderCapacity is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) CustOrderCapacity() (*field.CustOrderCapacity, error) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//AcctIDSource is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//AccountType is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//PositionEffect is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) PositionEffect() (*field.PositionEffect, error) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}

//PreallocMethod is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) PreallocMethod() (*field.PreallocMethod, error) {
	f := new(field.PreallocMethod)
	err := m.Body.Get(f)
	return f, err
}

//NoAllocs is a non-required field for TradeCaptureReportAck.
func (m *TradeCaptureReportAck) NoAllocs() (*field.NoAllocs, error) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}
