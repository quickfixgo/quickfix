//Package tradecapturereport msg type = AE.
package tradecapturereport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a TradeCaptureReport wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//TradeReportID is a non-required field for TradeCaptureReport.
func (m Message) TradeReportID() (*field.TradeReportIDField, quickfix.MessageRejectError) {
	f := &field.TradeReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportID reads a TradeReportID from TradeCaptureReport.
func (m Message) GetTradeReportID(f *field.TradeReportIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportTransType is a non-required field for TradeCaptureReport.
func (m Message) TradeReportTransType() (*field.TradeReportTransTypeField, quickfix.MessageRejectError) {
	f := &field.TradeReportTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportTransType reads a TradeReportTransType from TradeCaptureReport.
func (m Message) GetTradeReportTransType(f *field.TradeReportTransTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportType is a non-required field for TradeCaptureReport.
func (m Message) TradeReportType() (*field.TradeReportTypeField, quickfix.MessageRejectError) {
	f := &field.TradeReportTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportType reads a TradeReportType from TradeCaptureReport.
func (m Message) GetTradeReportType(f *field.TradeReportTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeRequestID is a non-required field for TradeCaptureReport.
func (m Message) TradeRequestID() (*field.TradeRequestIDField, quickfix.MessageRejectError) {
	f := &field.TradeRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeRequestID reads a TradeRequestID from TradeCaptureReport.
func (m Message) GetTradeRequestID(f *field.TradeRequestIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TrdType is a non-required field for TradeCaptureReport.
func (m Message) TrdType() (*field.TrdTypeField, quickfix.MessageRejectError) {
	f := &field.TrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdType reads a TrdType from TradeCaptureReport.
func (m Message) GetTrdType(f *field.TrdTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TrdSubType is a non-required field for TradeCaptureReport.
func (m Message) TrdSubType() (*field.TrdSubTypeField, quickfix.MessageRejectError) {
	f := &field.TrdSubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdSubType reads a TrdSubType from TradeCaptureReport.
func (m Message) GetTrdSubType(f *field.TrdSubTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTrdType is a non-required field for TradeCaptureReport.
func (m Message) SecondaryTrdType() (*field.SecondaryTrdTypeField, quickfix.MessageRejectError) {
	f := &field.SecondaryTrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTrdType reads a SecondaryTrdType from TradeCaptureReport.
func (m Message) GetSecondaryTrdType(f *field.SecondaryTrdTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransferReason is a non-required field for TradeCaptureReport.
func (m Message) TransferReason() (*field.TransferReasonField, quickfix.MessageRejectError) {
	f := &field.TransferReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransferReason reads a TransferReason from TradeCaptureReport.
func (m Message) GetTransferReason(f *field.TransferReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecType is a non-required field for TradeCaptureReport.
func (m Message) ExecType() (*field.ExecTypeField, quickfix.MessageRejectError) {
	f := &field.ExecTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecType reads a ExecType from TradeCaptureReport.
func (m Message) GetExecType(f *field.ExecTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TotNumTradeReports is a non-required field for TradeCaptureReport.
func (m Message) TotNumTradeReports() (*field.TotNumTradeReportsField, quickfix.MessageRejectError) {
	f := &field.TotNumTradeReportsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNumTradeReports reads a TotNumTradeReports from TradeCaptureReport.
func (m Message) GetTotNumTradeReports(f *field.TotNumTradeReportsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastRptRequested is a non-required field for TradeCaptureReport.
func (m Message) LastRptRequested() (*field.LastRptRequestedField, quickfix.MessageRejectError) {
	f := &field.LastRptRequestedField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastRptRequested reads a LastRptRequested from TradeCaptureReport.
func (m Message) GetLastRptRequested(f *field.LastRptRequestedField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnsolicitedIndicator is a non-required field for TradeCaptureReport.
func (m Message) UnsolicitedIndicator() (*field.UnsolicitedIndicatorField, quickfix.MessageRejectError) {
	f := &field.UnsolicitedIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnsolicitedIndicator reads a UnsolicitedIndicator from TradeCaptureReport.
func (m Message) GetUnsolicitedIndicator(f *field.UnsolicitedIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for TradeCaptureReport.
func (m Message) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, quickfix.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from TradeCaptureReport.
func (m Message) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportRefID is a non-required field for TradeCaptureReport.
func (m Message) TradeReportRefID() (*field.TradeReportRefIDField, quickfix.MessageRejectError) {
	f := &field.TradeReportRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportRefID reads a TradeReportRefID from TradeCaptureReport.
func (m Message) GetTradeReportRefID(f *field.TradeReportRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTradeReportRefID is a non-required field for TradeCaptureReport.
func (m Message) SecondaryTradeReportRefID() (*field.SecondaryTradeReportRefIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryTradeReportRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTradeReportRefID reads a SecondaryTradeReportRefID from TradeCaptureReport.
func (m Message) GetSecondaryTradeReportRefID(f *field.SecondaryTradeReportRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTradeReportID is a non-required field for TradeCaptureReport.
func (m Message) SecondaryTradeReportID() (*field.SecondaryTradeReportIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryTradeReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTradeReportID reads a SecondaryTradeReportID from TradeCaptureReport.
func (m Message) GetSecondaryTradeReportID(f *field.SecondaryTradeReportIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeLinkID is a non-required field for TradeCaptureReport.
func (m Message) TradeLinkID() (*field.TradeLinkIDField, quickfix.MessageRejectError) {
	f := &field.TradeLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeLinkID reads a TradeLinkID from TradeCaptureReport.
func (m Message) GetTradeLinkID(f *field.TradeLinkIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TrdMatchID is a non-required field for TradeCaptureReport.
func (m Message) TrdMatchID() (*field.TrdMatchIDField, quickfix.MessageRejectError) {
	f := &field.TrdMatchIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdMatchID reads a TrdMatchID from TradeCaptureReport.
func (m Message) GetTrdMatchID(f *field.TrdMatchIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecID is a non-required field for TradeCaptureReport.
func (m Message) ExecID() (*field.ExecIDField, quickfix.MessageRejectError) {
	f := &field.ExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecID reads a ExecID from TradeCaptureReport.
func (m Message) GetExecID(f *field.ExecIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrdStatus is a non-required field for TradeCaptureReport.
func (m Message) OrdStatus() (*field.OrdStatusField, quickfix.MessageRejectError) {
	f := &field.OrdStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdStatus reads a OrdStatus from TradeCaptureReport.
func (m Message) GetOrdStatus(f *field.OrdStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryExecID is a non-required field for TradeCaptureReport.
func (m Message) SecondaryExecID() (*field.SecondaryExecIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryExecID reads a SecondaryExecID from TradeCaptureReport.
func (m Message) GetSecondaryExecID(f *field.SecondaryExecIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecRestatementReason is a non-required field for TradeCaptureReport.
func (m Message) ExecRestatementReason() (*field.ExecRestatementReasonField, quickfix.MessageRejectError) {
	f := &field.ExecRestatementReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecRestatementReason reads a ExecRestatementReason from TradeCaptureReport.
func (m Message) GetExecRestatementReason(f *field.ExecRestatementReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PreviouslyReported is a non-required field for TradeCaptureReport.
func (m Message) PreviouslyReported() (*field.PreviouslyReportedField, quickfix.MessageRejectError) {
	f := &field.PreviouslyReportedField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreviouslyReported reads a PreviouslyReported from TradeCaptureReport.
func (m Message) GetPreviouslyReported(f *field.PreviouslyReportedField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for TradeCaptureReport.
func (m Message) PriceType() (*field.PriceTypeField, quickfix.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from TradeCaptureReport.
func (m Message) GetPriceType(f *field.PriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for TradeCaptureReport.
func (m Message) Symbol() (*field.SymbolField, quickfix.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from TradeCaptureReport.
func (m Message) GetSymbol(f *field.SymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for TradeCaptureReport.
func (m Message) SymbolSfx() (*field.SymbolSfxField, quickfix.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from TradeCaptureReport.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for TradeCaptureReport.
func (m Message) SecurityID() (*field.SecurityIDField, quickfix.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from TradeCaptureReport.
func (m Message) GetSecurityID(f *field.SecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for TradeCaptureReport.
func (m Message) SecurityIDSource() (*field.SecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from TradeCaptureReport.
func (m Message) GetSecurityIDSource(f *field.SecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for TradeCaptureReport.
func (m Message) NoSecurityAltID() (*field.NoSecurityAltIDField, quickfix.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from TradeCaptureReport.
func (m Message) GetNoSecurityAltID(f *field.NoSecurityAltIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for TradeCaptureReport.
func (m Message) Product() (*field.ProductField, quickfix.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from TradeCaptureReport.
func (m Message) GetProduct(f *field.ProductField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for TradeCaptureReport.
func (m Message) CFICode() (*field.CFICodeField, quickfix.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from TradeCaptureReport.
func (m Message) GetCFICode(f *field.CFICodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for TradeCaptureReport.
func (m Message) SecurityType() (*field.SecurityTypeField, quickfix.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from TradeCaptureReport.
func (m Message) GetSecurityType(f *field.SecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for TradeCaptureReport.
func (m Message) SecuritySubType() (*field.SecuritySubTypeField, quickfix.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from TradeCaptureReport.
func (m Message) GetSecuritySubType(f *field.SecuritySubTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for TradeCaptureReport.
func (m Message) MaturityMonthYear() (*field.MaturityMonthYearField, quickfix.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from TradeCaptureReport.
func (m Message) GetMaturityMonthYear(f *field.MaturityMonthYearField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for TradeCaptureReport.
func (m Message) MaturityDate() (*field.MaturityDateField, quickfix.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from TradeCaptureReport.
func (m Message) GetMaturityDate(f *field.MaturityDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for TradeCaptureReport.
func (m Message) CouponPaymentDate() (*field.CouponPaymentDateField, quickfix.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from TradeCaptureReport.
func (m Message) GetCouponPaymentDate(f *field.CouponPaymentDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for TradeCaptureReport.
func (m Message) IssueDate() (*field.IssueDateField, quickfix.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from TradeCaptureReport.
func (m Message) GetIssueDate(f *field.IssueDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReport.
func (m Message) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, quickfix.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from TradeCaptureReport.
func (m Message) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for TradeCaptureReport.
func (m Message) RepurchaseTerm() (*field.RepurchaseTermField, quickfix.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from TradeCaptureReport.
func (m Message) GetRepurchaseTerm(f *field.RepurchaseTermField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for TradeCaptureReport.
func (m Message) RepurchaseRate() (*field.RepurchaseRateField, quickfix.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from TradeCaptureReport.
func (m Message) GetRepurchaseRate(f *field.RepurchaseRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for TradeCaptureReport.
func (m Message) Factor() (*field.FactorField, quickfix.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from TradeCaptureReport.
func (m Message) GetFactor(f *field.FactorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for TradeCaptureReport.
func (m Message) CreditRating() (*field.CreditRatingField, quickfix.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from TradeCaptureReport.
func (m Message) GetCreditRating(f *field.CreditRatingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for TradeCaptureReport.
func (m Message) InstrRegistry() (*field.InstrRegistryField, quickfix.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from TradeCaptureReport.
func (m Message) GetInstrRegistry(f *field.InstrRegistryField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for TradeCaptureReport.
func (m Message) CountryOfIssue() (*field.CountryOfIssueField, quickfix.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from TradeCaptureReport.
func (m Message) GetCountryOfIssue(f *field.CountryOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReport.
func (m Message) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, quickfix.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from TradeCaptureReport.
func (m Message) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for TradeCaptureReport.
func (m Message) LocaleOfIssue() (*field.LocaleOfIssueField, quickfix.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from TradeCaptureReport.
func (m Message) GetLocaleOfIssue(f *field.LocaleOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for TradeCaptureReport.
func (m Message) RedemptionDate() (*field.RedemptionDateField, quickfix.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from TradeCaptureReport.
func (m Message) GetRedemptionDate(f *field.RedemptionDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for TradeCaptureReport.
func (m Message) StrikePrice() (*field.StrikePriceField, quickfix.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from TradeCaptureReport.
func (m Message) GetStrikePrice(f *field.StrikePriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for TradeCaptureReport.
func (m Message) StrikeCurrency() (*field.StrikeCurrencyField, quickfix.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from TradeCaptureReport.
func (m Message) GetStrikeCurrency(f *field.StrikeCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for TradeCaptureReport.
func (m Message) OptAttribute() (*field.OptAttributeField, quickfix.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from TradeCaptureReport.
func (m Message) GetOptAttribute(f *field.OptAttributeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for TradeCaptureReport.
func (m Message) ContractMultiplier() (*field.ContractMultiplierField, quickfix.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from TradeCaptureReport.
func (m Message) GetContractMultiplier(f *field.ContractMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for TradeCaptureReport.
func (m Message) CouponRate() (*field.CouponRateField, quickfix.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from TradeCaptureReport.
func (m Message) GetCouponRate(f *field.CouponRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for TradeCaptureReport.
func (m Message) SecurityExchange() (*field.SecurityExchangeField, quickfix.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from TradeCaptureReport.
func (m Message) GetSecurityExchange(f *field.SecurityExchangeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for TradeCaptureReport.
func (m Message) Issuer() (*field.IssuerField, quickfix.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from TradeCaptureReport.
func (m Message) GetIssuer(f *field.IssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for TradeCaptureReport.
func (m Message) EncodedIssuerLen() (*field.EncodedIssuerLenField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from TradeCaptureReport.
func (m Message) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for TradeCaptureReport.
func (m Message) EncodedIssuer() (*field.EncodedIssuerField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from TradeCaptureReport.
func (m Message) GetEncodedIssuer(f *field.EncodedIssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for TradeCaptureReport.
func (m Message) SecurityDesc() (*field.SecurityDescField, quickfix.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from TradeCaptureReport.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReport.
func (m Message) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from TradeCaptureReport.
func (m Message) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReport.
func (m Message) EncodedSecurityDesc() (*field.EncodedSecurityDescField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from TradeCaptureReport.
func (m Message) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for TradeCaptureReport.
func (m Message) Pool() (*field.PoolField, quickfix.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from TradeCaptureReport.
func (m Message) GetPool(f *field.PoolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for TradeCaptureReport.
func (m Message) ContractSettlMonth() (*field.ContractSettlMonthField, quickfix.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from TradeCaptureReport.
func (m Message) GetContractSettlMonth(f *field.ContractSettlMonthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for TradeCaptureReport.
func (m Message) CPProgram() (*field.CPProgramField, quickfix.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from TradeCaptureReport.
func (m Message) GetCPProgram(f *field.CPProgramField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for TradeCaptureReport.
func (m Message) CPRegType() (*field.CPRegTypeField, quickfix.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from TradeCaptureReport.
func (m Message) GetCPRegType(f *field.CPRegTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for TradeCaptureReport.
func (m Message) NoEvents() (*field.NoEventsField, quickfix.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from TradeCaptureReport.
func (m Message) GetNoEvents(f *field.NoEventsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for TradeCaptureReport.
func (m Message) DatedDate() (*field.DatedDateField, quickfix.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from TradeCaptureReport.
func (m Message) GetDatedDate(f *field.DatedDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for TradeCaptureReport.
func (m Message) InterestAccrualDate() (*field.InterestAccrualDateField, quickfix.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from TradeCaptureReport.
func (m Message) GetInterestAccrualDate(f *field.InterestAccrualDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for TradeCaptureReport.
func (m Message) SecurityStatus() (*field.SecurityStatusField, quickfix.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from TradeCaptureReport.
func (m Message) GetSecurityStatus(f *field.SecurityStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for TradeCaptureReport.
func (m Message) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, quickfix.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from TradeCaptureReport.
func (m Message) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for TradeCaptureReport.
func (m Message) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, quickfix.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from TradeCaptureReport.
func (m Message) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for TradeCaptureReport.
func (m Message) StrikeMultiplier() (*field.StrikeMultiplierField, quickfix.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from TradeCaptureReport.
func (m Message) GetStrikeMultiplier(f *field.StrikeMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for TradeCaptureReport.
func (m Message) StrikeValue() (*field.StrikeValueField, quickfix.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from TradeCaptureReport.
func (m Message) GetStrikeValue(f *field.StrikeValueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for TradeCaptureReport.
func (m Message) MinPriceIncrement() (*field.MinPriceIncrementField, quickfix.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from TradeCaptureReport.
func (m Message) GetMinPriceIncrement(f *field.MinPriceIncrementField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for TradeCaptureReport.
func (m Message) PositionLimit() (*field.PositionLimitField, quickfix.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from TradeCaptureReport.
func (m Message) GetPositionLimit(f *field.PositionLimitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for TradeCaptureReport.
func (m Message) NTPositionLimit() (*field.NTPositionLimitField, quickfix.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from TradeCaptureReport.
func (m Message) GetNTPositionLimit(f *field.NTPositionLimitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for TradeCaptureReport.
func (m Message) NoInstrumentParties() (*field.NoInstrumentPartiesField, quickfix.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from TradeCaptureReport.
func (m Message) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for TradeCaptureReport.
func (m Message) UnitOfMeasure() (*field.UnitOfMeasureField, quickfix.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from TradeCaptureReport.
func (m Message) GetUnitOfMeasure(f *field.UnitOfMeasureField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for TradeCaptureReport.
func (m Message) TimeUnit() (*field.TimeUnitField, quickfix.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from TradeCaptureReport.
func (m Message) GetTimeUnit(f *field.TimeUnitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for TradeCaptureReport.
func (m Message) MaturityTime() (*field.MaturityTimeField, quickfix.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from TradeCaptureReport.
func (m Message) GetMaturityTime(f *field.MaturityTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for TradeCaptureReport.
func (m Message) AgreementDesc() (*field.AgreementDescField, quickfix.MessageRejectError) {
	f := &field.AgreementDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from TradeCaptureReport.
func (m Message) GetAgreementDesc(f *field.AgreementDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for TradeCaptureReport.
func (m Message) AgreementID() (*field.AgreementIDField, quickfix.MessageRejectError) {
	f := &field.AgreementIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from TradeCaptureReport.
func (m Message) GetAgreementID(f *field.AgreementIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for TradeCaptureReport.
func (m Message) AgreementDate() (*field.AgreementDateField, quickfix.MessageRejectError) {
	f := &field.AgreementDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from TradeCaptureReport.
func (m Message) GetAgreementDate(f *field.AgreementDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for TradeCaptureReport.
func (m Message) AgreementCurrency() (*field.AgreementCurrencyField, quickfix.MessageRejectError) {
	f := &field.AgreementCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from TradeCaptureReport.
func (m Message) GetAgreementCurrency(f *field.AgreementCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for TradeCaptureReport.
func (m Message) TerminationType() (*field.TerminationTypeField, quickfix.MessageRejectError) {
	f := &field.TerminationTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from TradeCaptureReport.
func (m Message) GetTerminationType(f *field.TerminationTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for TradeCaptureReport.
func (m Message) StartDate() (*field.StartDateField, quickfix.MessageRejectError) {
	f := &field.StartDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from TradeCaptureReport.
func (m Message) GetStartDate(f *field.StartDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for TradeCaptureReport.
func (m Message) EndDate() (*field.EndDateField, quickfix.MessageRejectError) {
	f := &field.EndDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from TradeCaptureReport.
func (m Message) GetEndDate(f *field.EndDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for TradeCaptureReport.
func (m Message) DeliveryType() (*field.DeliveryTypeField, quickfix.MessageRejectError) {
	f := &field.DeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from TradeCaptureReport.
func (m Message) GetDeliveryType(f *field.DeliveryTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for TradeCaptureReport.
func (m Message) MarginRatio() (*field.MarginRatioField, quickfix.MessageRejectError) {
	f := &field.MarginRatioField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from TradeCaptureReport.
func (m Message) GetMarginRatio(f *field.MarginRatioField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for TradeCaptureReport.
func (m Message) OrderQty() (*field.OrderQtyField, quickfix.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from TradeCaptureReport.
func (m Message) GetOrderQty(f *field.OrderQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for TradeCaptureReport.
func (m Message) CashOrderQty() (*field.CashOrderQtyField, quickfix.MessageRejectError) {
	f := &field.CashOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from TradeCaptureReport.
func (m Message) GetCashOrderQty(f *field.CashOrderQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for TradeCaptureReport.
func (m Message) OrderPercent() (*field.OrderPercentField, quickfix.MessageRejectError) {
	f := &field.OrderPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from TradeCaptureReport.
func (m Message) GetOrderPercent(f *field.OrderPercentField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for TradeCaptureReport.
func (m Message) RoundingDirection() (*field.RoundingDirectionField, quickfix.MessageRejectError) {
	f := &field.RoundingDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from TradeCaptureReport.
func (m Message) GetRoundingDirection(f *field.RoundingDirectionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for TradeCaptureReport.
func (m Message) RoundingModulus() (*field.RoundingModulusField, quickfix.MessageRejectError) {
	f := &field.RoundingModulusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from TradeCaptureReport.
func (m Message) GetRoundingModulus(f *field.RoundingModulusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for TradeCaptureReport.
func (m Message) QtyType() (*field.QtyTypeField, quickfix.MessageRejectError) {
	f := &field.QtyTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from TradeCaptureReport.
func (m Message) GetQtyType(f *field.QtyTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for TradeCaptureReport.
func (m Message) YieldType() (*field.YieldTypeField, quickfix.MessageRejectError) {
	f := &field.YieldTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from TradeCaptureReport.
func (m Message) GetYieldType(f *field.YieldTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for TradeCaptureReport.
func (m Message) Yield() (*field.YieldField, quickfix.MessageRejectError) {
	f := &field.YieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from TradeCaptureReport.
func (m Message) GetYield(f *field.YieldField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for TradeCaptureReport.
func (m Message) YieldCalcDate() (*field.YieldCalcDateField, quickfix.MessageRejectError) {
	f := &field.YieldCalcDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from TradeCaptureReport.
func (m Message) GetYieldCalcDate(f *field.YieldCalcDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for TradeCaptureReport.
func (m Message) YieldRedemptionDate() (*field.YieldRedemptionDateField, quickfix.MessageRejectError) {
	f := &field.YieldRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from TradeCaptureReport.
func (m Message) GetYieldRedemptionDate(f *field.YieldRedemptionDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for TradeCaptureReport.
func (m Message) YieldRedemptionPrice() (*field.YieldRedemptionPriceField, quickfix.MessageRejectError) {
	f := &field.YieldRedemptionPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from TradeCaptureReport.
func (m Message) GetYieldRedemptionPrice(f *field.YieldRedemptionPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for TradeCaptureReport.
func (m Message) YieldRedemptionPriceType() (*field.YieldRedemptionPriceTypeField, quickfix.MessageRejectError) {
	f := &field.YieldRedemptionPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from TradeCaptureReport.
func (m Message) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for TradeCaptureReport.
func (m Message) NoUnderlyings() (*field.NoUnderlyingsField, quickfix.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from TradeCaptureReport.
func (m Message) GetNoUnderlyings(f *field.NoUnderlyingsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTradingSessionID is a non-required field for TradeCaptureReport.
func (m Message) UnderlyingTradingSessionID() (*field.UnderlyingTradingSessionIDField, quickfix.MessageRejectError) {
	f := &field.UnderlyingTradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTradingSessionID reads a UnderlyingTradingSessionID from TradeCaptureReport.
func (m Message) GetUnderlyingTradingSessionID(f *field.UnderlyingTradingSessionIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTradingSessionSubID is a non-required field for TradeCaptureReport.
func (m Message) UnderlyingTradingSessionSubID() (*field.UnderlyingTradingSessionSubIDField, quickfix.MessageRejectError) {
	f := &field.UnderlyingTradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTradingSessionSubID reads a UnderlyingTradingSessionSubID from TradeCaptureReport.
func (m Message) GetUnderlyingTradingSessionSubID(f *field.UnderlyingTradingSessionSubIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastQty is a required field for TradeCaptureReport.
func (m Message) LastQty() (*field.LastQtyField, quickfix.MessageRejectError) {
	f := &field.LastQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastQty reads a LastQty from TradeCaptureReport.
func (m Message) GetLastQty(f *field.LastQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a required field for TradeCaptureReport.
func (m Message) LastPx() (*field.LastPxField, quickfix.MessageRejectError) {
	f := &field.LastPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from TradeCaptureReport.
func (m Message) GetLastPx(f *field.LastPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastParPx is a non-required field for TradeCaptureReport.
func (m Message) LastParPx() (*field.LastParPxField, quickfix.MessageRejectError) {
	f := &field.LastParPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastParPx reads a LastParPx from TradeCaptureReport.
func (m Message) GetLastParPx(f *field.LastParPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastSpotRate is a non-required field for TradeCaptureReport.
func (m Message) LastSpotRate() (*field.LastSpotRateField, quickfix.MessageRejectError) {
	f := &field.LastSpotRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastSpotRate reads a LastSpotRate from TradeCaptureReport.
func (m Message) GetLastSpotRate(f *field.LastSpotRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastForwardPoints is a non-required field for TradeCaptureReport.
func (m Message) LastForwardPoints() (*field.LastForwardPointsField, quickfix.MessageRejectError) {
	f := &field.LastForwardPointsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastForwardPoints reads a LastForwardPoints from TradeCaptureReport.
func (m Message) GetLastForwardPoints(f *field.LastForwardPointsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for TradeCaptureReport.
func (m Message) LastMkt() (*field.LastMktField, quickfix.MessageRejectError) {
	f := &field.LastMktField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from TradeCaptureReport.
func (m Message) GetLastMkt(f *field.LastMktField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for TradeCaptureReport.
func (m Message) TradeDate() (*field.TradeDateField, quickfix.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from TradeCaptureReport.
func (m Message) GetTradeDate(f *field.TradeDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for TradeCaptureReport.
func (m Message) ClearingBusinessDate() (*field.ClearingBusinessDateField, quickfix.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from TradeCaptureReport.
func (m Message) GetClearingBusinessDate(f *field.ClearingBusinessDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a non-required field for TradeCaptureReport.
func (m Message) AvgPx() (*field.AvgPxField, quickfix.MessageRejectError) {
	f := &field.AvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from TradeCaptureReport.
func (m Message) GetAvgPx(f *field.AvgPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for TradeCaptureReport.
func (m Message) Spread() (*field.SpreadField, quickfix.MessageRejectError) {
	f := &field.SpreadField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from TradeCaptureReport.
func (m Message) GetSpread(f *field.SpreadField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for TradeCaptureReport.
func (m Message) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrencyField, quickfix.MessageRejectError) {
	f := &field.BenchmarkCurveCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from TradeCaptureReport.
func (m Message) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for TradeCaptureReport.
func (m Message) BenchmarkCurveName() (*field.BenchmarkCurveNameField, quickfix.MessageRejectError) {
	f := &field.BenchmarkCurveNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from TradeCaptureReport.
func (m Message) GetBenchmarkCurveName(f *field.BenchmarkCurveNameField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for TradeCaptureReport.
func (m Message) BenchmarkCurvePoint() (*field.BenchmarkCurvePointField, quickfix.MessageRejectError) {
	f := &field.BenchmarkCurvePointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from TradeCaptureReport.
func (m Message) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePointField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for TradeCaptureReport.
func (m Message) BenchmarkPrice() (*field.BenchmarkPriceField, quickfix.MessageRejectError) {
	f := &field.BenchmarkPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from TradeCaptureReport.
func (m Message) GetBenchmarkPrice(f *field.BenchmarkPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for TradeCaptureReport.
func (m Message) BenchmarkPriceType() (*field.BenchmarkPriceTypeField, quickfix.MessageRejectError) {
	f := &field.BenchmarkPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from TradeCaptureReport.
func (m Message) GetBenchmarkPriceType(f *field.BenchmarkPriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for TradeCaptureReport.
func (m Message) BenchmarkSecurityID() (*field.BenchmarkSecurityIDField, quickfix.MessageRejectError) {
	f := &field.BenchmarkSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from TradeCaptureReport.
func (m Message) GetBenchmarkSecurityID(f *field.BenchmarkSecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for TradeCaptureReport.
func (m Message) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.BenchmarkSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from TradeCaptureReport.
func (m Message) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxIndicator is a non-required field for TradeCaptureReport.
func (m Message) AvgPxIndicator() (*field.AvgPxIndicatorField, quickfix.MessageRejectError) {
	f := &field.AvgPxIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxIndicator reads a AvgPxIndicator from TradeCaptureReport.
func (m Message) GetAvgPxIndicator(f *field.AvgPxIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPosAmt is a non-required field for TradeCaptureReport.
func (m Message) NoPosAmt() (*field.NoPosAmtField, quickfix.MessageRejectError) {
	f := &field.NoPosAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPosAmt reads a NoPosAmt from TradeCaptureReport.
func (m Message) GetNoPosAmt(f *field.NoPosAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegReportingType is a non-required field for TradeCaptureReport.
func (m Message) MultiLegReportingType() (*field.MultiLegReportingTypeField, quickfix.MessageRejectError) {
	f := &field.MultiLegReportingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegReportingType reads a MultiLegReportingType from TradeCaptureReport.
func (m Message) GetMultiLegReportingType(f *field.MultiLegReportingTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeLegRefID is a non-required field for TradeCaptureReport.
func (m Message) TradeLegRefID() (*field.TradeLegRefIDField, quickfix.MessageRejectError) {
	f := &field.TradeLegRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeLegRefID reads a TradeLegRefID from TradeCaptureReport.
func (m Message) GetTradeLegRefID(f *field.TradeLegRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for TradeCaptureReport.
func (m Message) NoLegs() (*field.NoLegsField, quickfix.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from TradeCaptureReport.
func (m Message) GetNoLegs(f *field.NoLegsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for TradeCaptureReport.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from TradeCaptureReport.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRegTimestamps is a non-required field for TradeCaptureReport.
func (m Message) NoTrdRegTimestamps() (*field.NoTrdRegTimestampsField, quickfix.MessageRejectError) {
	f := &field.NoTrdRegTimestampsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRegTimestamps reads a NoTrdRegTimestamps from TradeCaptureReport.
func (m Message) GetNoTrdRegTimestamps(f *field.NoTrdRegTimestampsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for TradeCaptureReport.
func (m Message) SettlType() (*field.SettlTypeField, quickfix.MessageRejectError) {
	f := &field.SettlTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from TradeCaptureReport.
func (m Message) GetSettlType(f *field.SettlTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for TradeCaptureReport.
func (m Message) SettlDate() (*field.SettlDateField, quickfix.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from TradeCaptureReport.
func (m Message) GetSettlDate(f *field.SettlDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for TradeCaptureReport.
func (m Message) MatchStatus() (*field.MatchStatusField, quickfix.MessageRejectError) {
	f := &field.MatchStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from TradeCaptureReport.
func (m Message) GetMatchStatus(f *field.MatchStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MatchType is a non-required field for TradeCaptureReport.
func (m Message) MatchType() (*field.MatchTypeField, quickfix.MessageRejectError) {
	f := &field.MatchTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchType reads a MatchType from TradeCaptureReport.
func (m Message) GetMatchType(f *field.MatchTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoSides is a required field for TradeCaptureReport.
func (m Message) NoSides() (*field.NoSidesField, quickfix.MessageRejectError) {
	f := &field.NoSidesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSides reads a NoSides from TradeCaptureReport.
func (m Message) GetNoSides(f *field.NoSidesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CopyMsgIndicator is a non-required field for TradeCaptureReport.
func (m Message) CopyMsgIndicator() (*field.CopyMsgIndicatorField, quickfix.MessageRejectError) {
	f := &field.CopyMsgIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCopyMsgIndicator reads a CopyMsgIndicator from TradeCaptureReport.
func (m Message) GetCopyMsgIndicator(f *field.CopyMsgIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PublishTrdIndicator is a non-required field for TradeCaptureReport.
func (m Message) PublishTrdIndicator() (*field.PublishTrdIndicatorField, quickfix.MessageRejectError) {
	f := &field.PublishTrdIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPublishTrdIndicator reads a PublishTrdIndicator from TradeCaptureReport.
func (m Message) GetPublishTrdIndicator(f *field.PublishTrdIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ShortSaleReason is a non-required field for TradeCaptureReport.
func (m Message) ShortSaleReason() (*field.ShortSaleReasonField, quickfix.MessageRejectError) {
	f := &field.ShortSaleReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetShortSaleReason reads a ShortSaleReason from TradeCaptureReport.
func (m Message) GetShortSaleReason(f *field.ShortSaleReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TrdRptStatus is a non-required field for TradeCaptureReport.
func (m Message) TrdRptStatus() (*field.TrdRptStatusField, quickfix.MessageRejectError) {
	f := &field.TrdRptStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdRptStatus reads a TrdRptStatus from TradeCaptureReport.
func (m Message) GetTrdRptStatus(f *field.TrdRptStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AsOfIndicator is a non-required field for TradeCaptureReport.
func (m Message) AsOfIndicator() (*field.AsOfIndicatorField, quickfix.MessageRejectError) {
	f := &field.AsOfIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAsOfIndicator reads a AsOfIndicator from TradeCaptureReport.
func (m Message) GetAsOfIndicator(f *field.AsOfIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for TradeCaptureReport.
func (m Message) SettlSessID() (*field.SettlSessIDField, quickfix.MessageRejectError) {
	f := &field.SettlSessIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from TradeCaptureReport.
func (m Message) GetSettlSessID(f *field.SettlSessIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for TradeCaptureReport.
func (m Message) SettlSessSubID() (*field.SettlSessSubIDField, quickfix.MessageRejectError) {
	f := &field.SettlSessSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from TradeCaptureReport.
func (m Message) GetSettlSessSubID(f *field.SettlSessSubIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TierCode is a non-required field for TradeCaptureReport.
func (m Message) TierCode() (*field.TierCodeField, quickfix.MessageRejectError) {
	f := &field.TierCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTierCode reads a TierCode from TradeCaptureReport.
func (m Message) GetTierCode(f *field.TierCodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageEventSource is a non-required field for TradeCaptureReport.
func (m Message) MessageEventSource() (*field.MessageEventSourceField, quickfix.MessageRejectError) {
	f := &field.MessageEventSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMessageEventSource reads a MessageEventSource from TradeCaptureReport.
func (m Message) GetMessageEventSource(f *field.MessageEventSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastUpdateTime is a non-required field for TradeCaptureReport.
func (m Message) LastUpdateTime() (*field.LastUpdateTimeField, quickfix.MessageRejectError) {
	f := &field.LastUpdateTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastUpdateTime reads a LastUpdateTime from TradeCaptureReport.
func (m Message) GetLastUpdateTime(f *field.LastUpdateTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RndPx is a non-required field for TradeCaptureReport.
func (m Message) RndPx() (*field.RndPxField, quickfix.MessageRejectError) {
	f := &field.RndPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRndPx reads a RndPx from TradeCaptureReport.
func (m Message) GetRndPx(f *field.RndPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeID is a non-required field for TradeCaptureReport.
func (m Message) TradeID() (*field.TradeIDField, quickfix.MessageRejectError) {
	f := &field.TradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeID reads a TradeID from TradeCaptureReport.
func (m Message) GetTradeID(f *field.TradeIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTradeID is a non-required field for TradeCaptureReport.
func (m Message) SecondaryTradeID() (*field.SecondaryTradeIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryTradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTradeID reads a SecondaryTradeID from TradeCaptureReport.
func (m Message) GetSecondaryTradeID(f *field.SecondaryTradeIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FirmTradeID is a non-required field for TradeCaptureReport.
func (m Message) FirmTradeID() (*field.FirmTradeIDField, quickfix.MessageRejectError) {
	f := &field.FirmTradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFirmTradeID reads a FirmTradeID from TradeCaptureReport.
func (m Message) GetFirmTradeID(f *field.FirmTradeIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryFirmTradeID is a non-required field for TradeCaptureReport.
func (m Message) SecondaryFirmTradeID() (*field.SecondaryFirmTradeIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryFirmTradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryFirmTradeID reads a SecondaryFirmTradeID from TradeCaptureReport.
func (m Message) GetSecondaryFirmTradeID(f *field.SecondaryFirmTradeIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CalculatedCcyLastQty is a non-required field for TradeCaptureReport.
func (m Message) CalculatedCcyLastQty() (*field.CalculatedCcyLastQtyField, quickfix.MessageRejectError) {
	f := &field.CalculatedCcyLastQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCalculatedCcyLastQty reads a CalculatedCcyLastQty from TradeCaptureReport.
func (m Message) GetCalculatedCcyLastQty(f *field.CalculatedCcyLastQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastSwapPoints is a non-required field for TradeCaptureReport.
func (m Message) LastSwapPoints() (*field.LastSwapPointsField, quickfix.MessageRejectError) {
	f := &field.LastSwapPointsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastSwapPoints reads a LastSwapPoints from TradeCaptureReport.
func (m Message) GetLastSwapPoints(f *field.LastSwapPointsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlementDate is a non-required field for TradeCaptureReport.
func (m Message) UnderlyingSettlementDate() (*field.UnderlyingSettlementDateField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSettlementDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlementDate reads a UnderlyingSettlementDate from TradeCaptureReport.
func (m Message) GetUnderlyingSettlementDate(f *field.UnderlyingSettlementDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a non-required field for TradeCaptureReport.
func (m Message) GrossTradeAmt() (*field.GrossTradeAmtField, quickfix.MessageRejectError) {
	f := &field.GrossTradeAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from TradeCaptureReport.
func (m Message) GetGrossTradeAmt(f *field.GrossTradeAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRootPartyIDs is a non-required field for TradeCaptureReport.
func (m Message) NoRootPartyIDs() (*field.NoRootPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoRootPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRootPartyIDs reads a NoRootPartyIDs from TradeCaptureReport.
func (m Message) GetNoRootPartyIDs(f *field.NoRootPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCategory is a non-required field for TradeCaptureReport.
func (m Message) OrderCategory() (*field.OrderCategoryField, quickfix.MessageRejectError) {
	f := &field.OrderCategoryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCategory reads a OrderCategory from TradeCaptureReport.
func (m Message) GetOrderCategory(f *field.OrderCategoryField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeHandlingInstr is a non-required field for TradeCaptureReport.
func (m Message) TradeHandlingInstr() (*field.TradeHandlingInstrField, quickfix.MessageRejectError) {
	f := &field.TradeHandlingInstrField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeHandlingInstr reads a TradeHandlingInstr from TradeCaptureReport.
func (m Message) GetTradeHandlingInstr(f *field.TradeHandlingInstrField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTradeHandlingInstr is a non-required field for TradeCaptureReport.
func (m Message) OrigTradeHandlingInstr() (*field.OrigTradeHandlingInstrField, quickfix.MessageRejectError) {
	f := &field.OrigTradeHandlingInstrField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTradeHandlingInstr reads a OrigTradeHandlingInstr from TradeCaptureReport.
func (m Message) GetOrigTradeHandlingInstr(f *field.OrigTradeHandlingInstrField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTradeDate is a non-required field for TradeCaptureReport.
func (m Message) OrigTradeDate() (*field.OrigTradeDateField, quickfix.MessageRejectError) {
	f := &field.OrigTradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTradeDate reads a OrigTradeDate from TradeCaptureReport.
func (m Message) GetOrigTradeDate(f *field.OrigTradeDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTradeID is a non-required field for TradeCaptureReport.
func (m Message) OrigTradeID() (*field.OrigTradeIDField, quickfix.MessageRejectError) {
	f := &field.OrigTradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTradeID reads a OrigTradeID from TradeCaptureReport.
func (m Message) GetOrigTradeID(f *field.OrigTradeIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrigSecondaryTradeID is a non-required field for TradeCaptureReport.
func (m Message) OrigSecondaryTradeID() (*field.OrigSecondaryTradeIDField, quickfix.MessageRejectError) {
	f := &field.OrigSecondaryTradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigSecondaryTradeID reads a OrigSecondaryTradeID from TradeCaptureReport.
func (m Message) GetOrigSecondaryTradeID(f *field.OrigSecondaryTradeIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TZTransactTime is a non-required field for TradeCaptureReport.
func (m Message) TZTransactTime() (*field.TZTransactTimeField, quickfix.MessageRejectError) {
	f := &field.TZTransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTZTransactTime reads a TZTransactTime from TradeCaptureReport.
func (m Message) GetTZTransactTime(f *field.TZTransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ReportedPxDiff is a non-required field for TradeCaptureReport.
func (m Message) ReportedPxDiff() (*field.ReportedPxDiffField, quickfix.MessageRejectError) {
	f := &field.ReportedPxDiffField{}
	err := m.Body.Get(f)
	return f, err
}

//GetReportedPxDiff reads a ReportedPxDiff from TradeCaptureReport.
func (m Message) GetReportedPxDiff(f *field.ReportedPxDiffField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds TradeCaptureReport messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for TradeCaptureReport.
func Builder(
	lastqty *field.LastQtyField,
	lastpx *field.LastPxField,
	tradedate *field.TradeDateField,
	nosides *field.NoSidesField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.NewMsgType("AE"))
	builder.Body.Set(lastqty)
	builder.Body.Set(lastpx)
	builder.Body.Set(tradedate)
	builder.Body.Set(nosides)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX50, "AE", r
}
