package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	nosides field.NoSides) TradeCaptureReportAckBuilder {
	var builder TradeCaptureReportAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.BuildMsgType("AR"))
	builder.Body.Set(nosides)
	return builder
}

//TradeReportID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportID() (*field.TradeReportID, errors.MessageRejectError) {
	f := new(field.TradeReportID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportID reads a TradeReportID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeReportID(f *field.TradeReportID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportTransType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportTransType() (*field.TradeReportTransType, errors.MessageRejectError) {
	f := new(field.TradeReportTransType)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportTransType reads a TradeReportTransType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeReportTransType(f *field.TradeReportTransType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportType() (*field.TradeReportType, errors.MessageRejectError) {
	f := new(field.TradeReportType)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportType reads a TradeReportType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeReportType(f *field.TradeReportType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdType() (*field.TrdType, errors.MessageRejectError) {
	f := new(field.TrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetTrdType reads a TrdType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTrdType(f *field.TrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdSubType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdSubType() (*field.TrdSubType, errors.MessageRejectError) {
	f := new(field.TrdSubType)
	err := m.Body.Get(f)
	return f, err
}

//GetTrdSubType reads a TrdSubType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTrdSubType(f *field.TrdSubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTrdType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryTrdType() (*field.SecondaryTrdType, errors.MessageRejectError) {
	f := new(field.SecondaryTrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTrdType reads a SecondaryTrdType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecondaryTrdType(f *field.SecondaryTrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransferReason is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TransferReason() (*field.TransferReason, errors.MessageRejectError) {
	f := new(field.TransferReason)
	err := m.Body.Get(f)
	return f, err
}

//GetTransferReason reads a TransferReason from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTransferReason(f *field.TransferReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ExecType() (*field.ExecType, errors.MessageRejectError) {
	f := new(field.ExecType)
	err := m.Body.Get(f)
	return f, err
}

//GetExecType reads a ExecType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetExecType(f *field.ExecType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportRefID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportRefID() (*field.TradeReportRefID, errors.MessageRejectError) {
	f := new(field.TradeReportRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportRefID reads a TradeReportRefID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeReportRefID(f *field.TradeReportRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTradeReportRefID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryTradeReportRefID() (*field.SecondaryTradeReportRefID, errors.MessageRejectError) {
	f := new(field.SecondaryTradeReportRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTradeReportRefID reads a SecondaryTradeReportRefID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecondaryTradeReportRefID(f *field.SecondaryTradeReportRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdRptStatus is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdRptStatus() (*field.TrdRptStatus, errors.MessageRejectError) {
	f := new(field.TrdRptStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetTrdRptStatus reads a TrdRptStatus from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTrdRptStatus(f *field.TrdRptStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportRejectReason is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportRejectReason() (*field.TradeReportRejectReason, errors.MessageRejectError) {
	f := new(field.TradeReportRejectReason)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportRejectReason reads a TradeReportRejectReason from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeReportRejectReason(f *field.TradeReportRejectReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTradeReportID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryTradeReportID() (*field.SecondaryTradeReportID, errors.MessageRejectError) {
	f := new(field.SecondaryTradeReportID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTradeReportID reads a SecondaryTradeReportID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecondaryTradeReportID(f *field.SecondaryTradeReportID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SubscriptionRequestType() (*field.SubscriptionRequestType, errors.MessageRejectError) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSubscriptionRequestType(f *field.SubscriptionRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeLinkID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeLinkID() (*field.TradeLinkID, errors.MessageRejectError) {
	f := new(field.TradeLinkID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeLinkID reads a TradeLinkID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeLinkID(f *field.TradeLinkID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdMatchID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdMatchID() (*field.TrdMatchID, errors.MessageRejectError) {
	f := new(field.TrdMatchID)
	err := m.Body.Get(f)
	return f, err
}

//GetTrdMatchID reads a TrdMatchID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTrdMatchID(f *field.TrdMatchID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ExecID() (*field.ExecID, errors.MessageRejectError) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}

//GetExecID reads a ExecID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetExecID(f *field.ExecID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryExecID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryExecID() (*field.SecondaryExecID, errors.MessageRejectError) {
	f := new(field.SecondaryExecID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryExecID reads a SecondaryExecID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecondaryExecID(f *field.SecondaryExecID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRegTimestamps is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoTrdRegTimestamps() (*field.NoTrdRegTimestamps, errors.MessageRejectError) {
	f := new(field.NoTrdRegTimestamps)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRegTimestamps reads a NoTrdRegTimestamps from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetNoTrdRegTimestamps(f *field.NoTrdRegTimestamps) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseTransportType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ResponseTransportType() (*field.ResponseTransportType, errors.MessageRejectError) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}

//GetResponseTransportType reads a ResponseTransportType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetResponseTransportType(f *field.ResponseTransportType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseDestination is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ResponseDestination() (*field.ResponseDestination, errors.MessageRejectError) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetResponseDestination reads a ResponseDestination from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetResponseDestination(f *field.ResponseDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ClearingFeeIndicator() (*field.ClearingFeeIndicator, errors.MessageRejectError) {
	f := new(field.ClearingFeeIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetClearingFeeIndicator(f *field.ClearingFeeIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdStatus is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OrdStatus() (*field.OrdStatus, errors.MessageRejectError) {
	f := new(field.OrdStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdStatus reads a OrdStatus from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetOrdStatus(f *field.OrdStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecRestatementReason is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ExecRestatementReason() (*field.ExecRestatementReason, errors.MessageRejectError) {
	f := new(field.ExecRestatementReason)
	err := m.Body.Get(f)
	return f, err
}

//GetExecRestatementReason reads a ExecRestatementReason from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetExecRestatementReason(f *field.ExecRestatementReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreviouslyReported is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PreviouslyReported() (*field.PreviouslyReported, errors.MessageRejectError) {
	f := new(field.PreviouslyReported)
	err := m.Body.Get(f)
	return f, err
}

//GetPreviouslyReported reads a PreviouslyReported from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetPreviouslyReported(f *field.PreviouslyReported) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTradingSessionID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) UnderlyingTradingSessionID() (*field.UnderlyingTradingSessionID, errors.MessageRejectError) {
	f := new(field.UnderlyingTradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTradingSessionID reads a UnderlyingTradingSessionID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetUnderlyingTradingSessionID(f *field.UnderlyingTradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) QtyType() (*field.QtyType, errors.MessageRejectError) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetQtyType(f *field.QtyType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTradingSessionSubID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) UnderlyingTradingSessionSubID() (*field.UnderlyingTradingSessionSubID, errors.MessageRejectError) {
	f := new(field.UnderlyingTradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTradingSessionSubID reads a UnderlyingTradingSessionSubID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetUnderlyingTradingSessionSubID(f *field.UnderlyingTradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastQty is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastQty() (*field.LastQty, errors.MessageRejectError) {
	f := new(field.LastQty)
	err := m.Body.Get(f)
	return f, err
}

//GetLastQty reads a LastQty from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetLastQty(f *field.LastQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastPx() (*field.LastPx, errors.MessageRejectError) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetLastPx(f *field.LastPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastParPx is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastParPx() (*field.LastParPx, errors.MessageRejectError) {
	f := new(field.LastParPx)
	err := m.Body.Get(f)
	return f, err
}

//GetLastParPx reads a LastParPx from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetLastParPx(f *field.LastParPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastSpotRate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastSpotRate() (*field.LastSpotRate, errors.MessageRejectError) {
	f := new(field.LastSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//GetLastSpotRate reads a LastSpotRate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetLastSpotRate(f *field.LastSpotRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastForwardPoints is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastForwardPoints() (*field.LastForwardPoints, errors.MessageRejectError) {
	f := new(field.LastForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//GetLastForwardPoints reads a LastForwardPoints from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetLastForwardPoints(f *field.LastForwardPoints) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastMkt() (*field.LastMkt, errors.MessageRejectError) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetLastMkt(f *field.LastMkt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) AvgPx() (*field.AvgPx, errors.MessageRejectError) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetAvgPx(f *field.AvgPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxIndicator is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) AvgPxIndicator() (*field.AvgPxIndicator, errors.MessageRejectError) {
	f := new(field.AvgPxIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxIndicator reads a AvgPxIndicator from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetAvgPxIndicator(f *field.AvgPxIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegReportingType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MultiLegReportingType() (*field.MultiLegReportingType, errors.MessageRejectError) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegReportingType reads a MultiLegReportingType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetMultiLegReportingType(f *field.MultiLegReportingType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeLegRefID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeLegRefID() (*field.TradeLegRefID, errors.MessageRejectError) {
	f := new(field.TradeLegRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeLegRefID reads a TradeLegRefID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeLegRefID(f *field.TradeLegRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SettlType() (*field.SettlType, errors.MessageRejectError) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSettlType(f *field.SettlType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MatchStatus() (*field.MatchStatus, errors.MessageRejectError) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetMatchStatus(f *field.MatchStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MatchType() (*field.MatchType, errors.MessageRejectError) {
	f := new(field.MatchType)
	err := m.Body.Get(f)
	return f, err
}

//GetMatchType reads a MatchType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetMatchType(f *field.MatchType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CopyMsgIndicator is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CopyMsgIndicator() (*field.CopyMsgIndicator, errors.MessageRejectError) {
	f := new(field.CopyMsgIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetCopyMsgIndicator reads a CopyMsgIndicator from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCopyMsgIndicator(f *field.CopyMsgIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PublishTrdIndicator is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PublishTrdIndicator() (*field.PublishTrdIndicator, errors.MessageRejectError) {
	f := new(field.PublishTrdIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetPublishTrdIndicator reads a PublishTrdIndicator from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetPublishTrdIndicator(f *field.PublishTrdIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ShortSaleReason is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ShortSaleReason() (*field.ShortSaleReason, errors.MessageRejectError) {
	f := new(field.ShortSaleReason)
	err := m.Body.Get(f)
	return f, err
}

//GetShortSaleReason reads a ShortSaleReason from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetShortSaleReason(f *field.ShortSaleReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SettlDate() (*field.SettlDate, errors.MessageRejectError) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSettlDate(f *field.SettlDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SettlSessID() (*field.SettlSessID, errors.MessageRejectError) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSettlSessID(f *field.SettlSessID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SettlSessSubID() (*field.SettlSessSubID, errors.MessageRejectError) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSettlSessSubID(f *field.SettlSessSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPosAmt is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoPosAmt() (*field.NoPosAmt, errors.MessageRejectError) {
	f := new(field.NoPosAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPosAmt reads a NoPosAmt from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetNoPosAmt(f *field.NoPosAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TierCode is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TierCode() (*field.TierCode, errors.MessageRejectError) {
	f := new(field.TierCode)
	err := m.Body.Get(f)
	return f, err
}

//GetTierCode reads a TierCode from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTierCode(f *field.TierCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageEventSource is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MessageEventSource() (*field.MessageEventSource, errors.MessageRejectError) {
	f := new(field.MessageEventSource)
	err := m.Body.Get(f)
	return f, err
}

//GetMessageEventSource reads a MessageEventSource from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetMessageEventSource(f *field.MessageEventSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastUpdateTime is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastUpdateTime() (*field.LastUpdateTime, errors.MessageRejectError) {
	f := new(field.LastUpdateTime)
	err := m.Body.Get(f)
	return f, err
}

//GetLastUpdateTime reads a LastUpdateTime from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetLastUpdateTime(f *field.LastUpdateTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RndPx is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RndPx() (*field.RndPx, errors.MessageRejectError) {
	f := new(field.RndPx)
	err := m.Body.Get(f)
	return f, err
}

//GetRndPx reads a RndPx from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetRndPx(f *field.RndPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSides is a required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoSides() (*field.NoSides, errors.MessageRejectError) {
	f := new(field.NoSides)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSides reads a NoSides from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetNoSides(f *field.NoSides) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AsOfIndicator is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) AsOfIndicator() (*field.AsOfIndicator, errors.MessageRejectError) {
	f := new(field.AsOfIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetAsOfIndicator reads a AsOfIndicator from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetAsOfIndicator(f *field.AsOfIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeID() (*field.TradeID, errors.MessageRejectError) {
	f := new(field.TradeID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeID reads a TradeID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeID(f *field.TradeID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTradeID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryTradeID() (*field.SecondaryTradeID, errors.MessageRejectError) {
	f := new(field.SecondaryTradeID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTradeID reads a SecondaryTradeID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecondaryTradeID(f *field.SecondaryTradeID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FirmTradeID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) FirmTradeID() (*field.FirmTradeID, errors.MessageRejectError) {
	f := new(field.FirmTradeID)
	err := m.Body.Get(f)
	return f, err
}

//GetFirmTradeID reads a FirmTradeID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetFirmTradeID(f *field.FirmTradeID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryFirmTradeID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryFirmTradeID() (*field.SecondaryFirmTradeID, errors.MessageRejectError) {
	f := new(field.SecondaryFirmTradeID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryFirmTradeID reads a SecondaryFirmTradeID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecondaryFirmTradeID(f *field.SecondaryFirmTradeID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CalculatedCcyLastQty is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CalculatedCcyLastQty() (*field.CalculatedCcyLastQty, errors.MessageRejectError) {
	f := new(field.CalculatedCcyLastQty)
	err := m.Body.Get(f)
	return f, err
}

//GetCalculatedCcyLastQty reads a CalculatedCcyLastQty from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCalculatedCcyLastQty(f *field.CalculatedCcyLastQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastSwapPoints is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastSwapPoints() (*field.LastSwapPoints, errors.MessageRejectError) {
	f := new(field.LastSwapPoints)
	err := m.Body.Get(f)
	return f, err
}

//GetLastSwapPoints reads a LastSwapPoints from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetLastSwapPoints(f *field.LastSwapPoints) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) GrossTradeAmt() (*field.GrossTradeAmt, errors.MessageRejectError) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetGrossTradeAmt(f *field.GrossTradeAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRootPartyIDs is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoRootPartyIDs() (*field.NoRootPartyIDs, errors.MessageRejectError) {
	f := new(field.NoRootPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRootPartyIDs reads a NoRootPartyIDs from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetNoRootPartyIDs(f *field.NoRootPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeHandlingInstr is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeHandlingInstr() (*field.TradeHandlingInstr, errors.MessageRejectError) {
	f := new(field.TradeHandlingInstr)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeHandlingInstr reads a TradeHandlingInstr from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeHandlingInstr(f *field.TradeHandlingInstr) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTradeHandlingInstr is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OrigTradeHandlingInstr() (*field.OrigTradeHandlingInstr, errors.MessageRejectError) {
	f := new(field.OrigTradeHandlingInstr)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTradeHandlingInstr reads a OrigTradeHandlingInstr from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetOrigTradeHandlingInstr(f *field.OrigTradeHandlingInstr) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTradeDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OrigTradeDate() (*field.OrigTradeDate, errors.MessageRejectError) {
	f := new(field.OrigTradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTradeDate reads a OrigTradeDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetOrigTradeDate(f *field.OrigTradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTradeID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OrigTradeID() (*field.OrigTradeID, errors.MessageRejectError) {
	f := new(field.OrigTradeID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTradeID reads a OrigTradeID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetOrigTradeID(f *field.OrigTradeID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigSecondaryTradeID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OrigSecondaryTradeID() (*field.OrigSecondaryTradeID, errors.MessageRejectError) {
	f := new(field.OrigSecondaryTradeID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigSecondaryTradeID reads a OrigSecondaryTradeID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetOrigSecondaryTradeID(f *field.OrigSecondaryTradeID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RptSys is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RptSys() (*field.RptSys, errors.MessageRejectError) {
	f := new(field.RptSys)
	err := m.Body.Get(f)
	return f, err
}

//GetRptSys reads a RptSys from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetRptSys(f *field.RptSys) errors.MessageRejectError {
	return m.Body.Get(f)
}
