package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	tradereportid *field.TradeReportIDField,
	exectype *field.ExecTypeField) TradeCaptureReportAckBuilder {
	var builder TradeCaptureReportAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header.Set(field.NewMsgType("AR"))
	builder.Body.Set(tradereportid)
	builder.Body.Set(exectype)
	return builder
}

//TradeReportID is a required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportID() (*field.TradeReportIDField, errors.MessageRejectError) {
	f := &field.TradeReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportID reads a TradeReportID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeReportID(f *field.TradeReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportTransType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportTransType() (*field.TradeReportTransTypeField, errors.MessageRejectError) {
	f := &field.TradeReportTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportTransType reads a TradeReportTransType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeReportTransType(f *field.TradeReportTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportType() (*field.TradeReportTypeField, errors.MessageRejectError) {
	f := &field.TradeReportTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportType reads a TradeReportType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeReportType(f *field.TradeReportTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdType() (*field.TrdTypeField, errors.MessageRejectError) {
	f := &field.TrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdType reads a TrdType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTrdType(f *field.TrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdSubType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdSubType() (*field.TrdSubTypeField, errors.MessageRejectError) {
	f := &field.TrdSubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdSubType reads a TrdSubType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTrdSubType(f *field.TrdSubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTrdType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryTrdType() (*field.SecondaryTrdTypeField, errors.MessageRejectError) {
	f := &field.SecondaryTrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTrdType reads a SecondaryTrdType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecondaryTrdType(f *field.SecondaryTrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransferReason is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TransferReason() (*field.TransferReasonField, errors.MessageRejectError) {
	f := &field.TransferReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransferReason reads a TransferReason from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTransferReason(f *field.TransferReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecType is a required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ExecType() (*field.ExecTypeField, errors.MessageRejectError) {
	f := &field.ExecTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecType reads a ExecType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetExecType(f *field.ExecTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportRefID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportRefID() (*field.TradeReportRefIDField, errors.MessageRejectError) {
	f := &field.TradeReportRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportRefID reads a TradeReportRefID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeReportRefID(f *field.TradeReportRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTradeReportRefID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryTradeReportRefID() (*field.SecondaryTradeReportRefIDField, errors.MessageRejectError) {
	f := &field.SecondaryTradeReportRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTradeReportRefID reads a SecondaryTradeReportRefID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecondaryTradeReportRefID(f *field.SecondaryTradeReportRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdRptStatus is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdRptStatus() (*field.TrdRptStatusField, errors.MessageRejectError) {
	f := &field.TrdRptStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdRptStatus reads a TrdRptStatus from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTrdRptStatus(f *field.TrdRptStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportRejectReason is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportRejectReason() (*field.TradeReportRejectReasonField, errors.MessageRejectError) {
	f := &field.TradeReportRejectReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportRejectReason reads a TradeReportRejectReason from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeReportRejectReason(f *field.TradeReportRejectReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTradeReportID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryTradeReportID() (*field.SecondaryTradeReportIDField, errors.MessageRejectError) {
	f := &field.SecondaryTradeReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTradeReportID reads a SecondaryTradeReportID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecondaryTradeReportID(f *field.SecondaryTradeReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeLinkID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeLinkID() (*field.TradeLinkIDField, errors.MessageRejectError) {
	f := &field.TradeLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeLinkID reads a TradeLinkID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTradeLinkID(f *field.TradeLinkIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdMatchID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdMatchID() (*field.TrdMatchIDField, errors.MessageRejectError) {
	f := &field.TrdMatchIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdMatchID reads a TrdMatchID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTrdMatchID(f *field.TrdMatchIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ExecID() (*field.ExecIDField, errors.MessageRejectError) {
	f := &field.ExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecID reads a ExecID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetExecID(f *field.ExecIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryExecID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryExecID() (*field.SecondaryExecIDField, errors.MessageRejectError) {
	f := &field.SecondaryExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryExecID reads a SecondaryExecID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecondaryExecID(f *field.SecondaryExecIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRegTimestamps is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoTrdRegTimestamps() (*field.NoTrdRegTimestampsField, errors.MessageRejectError) {
	f := &field.NoTrdRegTimestampsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRegTimestamps reads a NoTrdRegTimestamps from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetNoTrdRegTimestamps(f *field.NoTrdRegTimestampsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseTransportType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ResponseTransportType() (*field.ResponseTransportTypeField, errors.MessageRejectError) {
	f := &field.ResponseTransportTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetResponseTransportType reads a ResponseTransportType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetResponseTransportType(f *field.ResponseTransportTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseDestination is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ResponseDestination() (*field.ResponseDestinationField, errors.MessageRejectError) {
	f := &field.ResponseDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetResponseDestination reads a ResponseDestination from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetResponseDestination(f *field.ResponseDestinationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ClearingFeeIndicator() (*field.ClearingFeeIndicatorField, errors.MessageRejectError) {
	f := &field.ClearingFeeIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetClearingFeeIndicator(f *field.ClearingFeeIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OrderCapacity() (*field.OrderCapacityField, errors.MessageRejectError) {
	f := &field.OrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetOrderCapacity(f *field.OrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OrderRestrictions() (*field.OrderRestrictionsField, errors.MessageRejectError) {
	f := &field.OrderRestrictionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetOrderRestrictions(f *field.OrderRestrictionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CustOrderCapacity() (*field.CustOrderCapacityField, errors.MessageRejectError) {
	f := &field.CustOrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetCustOrderCapacity(f *field.CustOrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) AcctIDSource() (*field.AcctIDSourceField, errors.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetAcctIDSource(f *field.AcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PositionEffect() (*field.PositionEffectField, errors.MessageRejectError) {
	f := &field.PositionEffectField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetPositionEffect(f *field.PositionEffectField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreallocMethod is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PreallocMethod() (*field.PreallocMethodField, errors.MessageRejectError) {
	f := &field.PreallocMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreallocMethod reads a PreallocMethod from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetPreallocMethod(f *field.PreallocMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoAllocs() (*field.NoAllocsField, errors.MessageRejectError) {
	f := &field.NoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from TradeCaptureReportAck.
func (m TradeCaptureReportAck) GetNoAllocs(f *field.NoAllocsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
