package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//TradeCaptureReport msg type = AE.
type TradeCaptureReport struct {
	message.Message
}

//TradeCaptureReportBuilder builds TradeCaptureReport messages.
type TradeCaptureReportBuilder struct {
	message.MessageBuilder
}

//CreateTradeCaptureReportBuilder returns an initialized TradeCaptureReportBuilder with specified required fields.
func CreateTradeCaptureReportBuilder(
	lastqty *field.LastQtyField,
	lastpx *field.LastPxField,
	nosides *field.NoSidesField) TradeCaptureReportBuilder {
	var builder TradeCaptureReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("AE"))
	builder.Body.Set(lastqty)
	builder.Body.Set(lastpx)
	builder.Body.Set(nosides)
	return builder
}

//TradeReportID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportID() (*field.TradeReportIDField, errors.MessageRejectError) {
	f := &field.TradeReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportID reads a TradeReportID from TradeCaptureReport.
func (m TradeCaptureReport) GetTradeReportID(f *field.TradeReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportTransType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportTransType() (*field.TradeReportTransTypeField, errors.MessageRejectError) {
	f := &field.TradeReportTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportTransType reads a TradeReportTransType from TradeCaptureReport.
func (m TradeCaptureReport) GetTradeReportTransType(f *field.TradeReportTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportType() (*field.TradeReportTypeField, errors.MessageRejectError) {
	f := &field.TradeReportTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportType reads a TradeReportType from TradeCaptureReport.
func (m TradeCaptureReport) GetTradeReportType(f *field.TradeReportTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeRequestID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeRequestID() (*field.TradeRequestIDField, errors.MessageRejectError) {
	f := &field.TradeRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeRequestID reads a TradeRequestID from TradeCaptureReport.
func (m TradeCaptureReport) GetTradeRequestID(f *field.TradeRequestIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TrdType() (*field.TrdTypeField, errors.MessageRejectError) {
	f := &field.TrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdType reads a TrdType from TradeCaptureReport.
func (m TradeCaptureReport) GetTrdType(f *field.TrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdSubType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TrdSubType() (*field.TrdSubTypeField, errors.MessageRejectError) {
	f := &field.TrdSubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdSubType reads a TrdSubType from TradeCaptureReport.
func (m TradeCaptureReport) GetTrdSubType(f *field.TrdSubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTrdType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryTrdType() (*field.SecondaryTrdTypeField, errors.MessageRejectError) {
	f := &field.SecondaryTrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTrdType reads a SecondaryTrdType from TradeCaptureReport.
func (m TradeCaptureReport) GetSecondaryTrdType(f *field.SecondaryTrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransferReason is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TransferReason() (*field.TransferReasonField, errors.MessageRejectError) {
	f := &field.TransferReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransferReason reads a TransferReason from TradeCaptureReport.
func (m TradeCaptureReport) GetTransferReason(f *field.TransferReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ExecType() (*field.ExecTypeField, errors.MessageRejectError) {
	f := &field.ExecTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecType reads a ExecType from TradeCaptureReport.
func (m TradeCaptureReport) GetExecType(f *field.ExecTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNumTradeReports is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TotNumTradeReports() (*field.TotNumTradeReportsField, errors.MessageRejectError) {
	f := &field.TotNumTradeReportsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNumTradeReports reads a TotNumTradeReports from TradeCaptureReport.
func (m TradeCaptureReport) GetTotNumTradeReports(f *field.TotNumTradeReportsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastRptRequested is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastRptRequested() (*field.LastRptRequestedField, errors.MessageRejectError) {
	f := &field.LastRptRequestedField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastRptRequested reads a LastRptRequested from TradeCaptureReport.
func (m TradeCaptureReport) GetLastRptRequested(f *field.LastRptRequestedField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnsolicitedIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnsolicitedIndicator() (*field.UnsolicitedIndicatorField, errors.MessageRejectError) {
	f := &field.UnsolicitedIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnsolicitedIndicator reads a UnsolicitedIndicator from TradeCaptureReport.
func (m TradeCaptureReport) GetUnsolicitedIndicator(f *field.UnsolicitedIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from TradeCaptureReport.
func (m TradeCaptureReport) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportRefID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportRefID() (*field.TradeReportRefIDField, errors.MessageRejectError) {
	f := &field.TradeReportRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportRefID reads a TradeReportRefID from TradeCaptureReport.
func (m TradeCaptureReport) GetTradeReportRefID(f *field.TradeReportRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTradeReportRefID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryTradeReportRefID() (*field.SecondaryTradeReportRefIDField, errors.MessageRejectError) {
	f := &field.SecondaryTradeReportRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTradeReportRefID reads a SecondaryTradeReportRefID from TradeCaptureReport.
func (m TradeCaptureReport) GetSecondaryTradeReportRefID(f *field.SecondaryTradeReportRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTradeReportID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryTradeReportID() (*field.SecondaryTradeReportIDField, errors.MessageRejectError) {
	f := &field.SecondaryTradeReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTradeReportID reads a SecondaryTradeReportID from TradeCaptureReport.
func (m TradeCaptureReport) GetSecondaryTradeReportID(f *field.SecondaryTradeReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeLinkID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeLinkID() (*field.TradeLinkIDField, errors.MessageRejectError) {
	f := &field.TradeLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeLinkID reads a TradeLinkID from TradeCaptureReport.
func (m TradeCaptureReport) GetTradeLinkID(f *field.TradeLinkIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdMatchID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TrdMatchID() (*field.TrdMatchIDField, errors.MessageRejectError) {
	f := &field.TrdMatchIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdMatchID reads a TrdMatchID from TradeCaptureReport.
func (m TradeCaptureReport) GetTrdMatchID(f *field.TrdMatchIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ExecID() (*field.ExecIDField, errors.MessageRejectError) {
	f := &field.ExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecID reads a ExecID from TradeCaptureReport.
func (m TradeCaptureReport) GetExecID(f *field.ExecIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryExecID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryExecID() (*field.SecondaryExecIDField, errors.MessageRejectError) {
	f := &field.SecondaryExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryExecID reads a SecondaryExecID from TradeCaptureReport.
func (m TradeCaptureReport) GetSecondaryExecID(f *field.SecondaryExecIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecRestatementReason is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ExecRestatementReason() (*field.ExecRestatementReasonField, errors.MessageRejectError) {
	f := &field.ExecRestatementReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecRestatementReason reads a ExecRestatementReason from TradeCaptureReport.
func (m TradeCaptureReport) GetExecRestatementReason(f *field.ExecRestatementReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreviouslyReported is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PreviouslyReported() (*field.PreviouslyReportedField, errors.MessageRejectError) {
	f := &field.PreviouslyReportedField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreviouslyReported reads a PreviouslyReported from TradeCaptureReport.
func (m TradeCaptureReport) GetPreviouslyReported(f *field.PreviouslyReportedField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from TradeCaptureReport.
func (m TradeCaptureReport) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from TradeCaptureReport.
func (m TradeCaptureReport) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from TradeCaptureReport.
func (m TradeCaptureReport) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from TradeCaptureReport.
func (m TradeCaptureReport) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from TradeCaptureReport.
func (m TradeCaptureReport) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from TradeCaptureReport.
func (m TradeCaptureReport) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from TradeCaptureReport.
func (m TradeCaptureReport) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from TradeCaptureReport.
func (m TradeCaptureReport) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from TradeCaptureReport.
func (m TradeCaptureReport) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from TradeCaptureReport.
func (m TradeCaptureReport) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from TradeCaptureReport.
func (m TradeCaptureReport) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from TradeCaptureReport.
func (m TradeCaptureReport) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from TradeCaptureReport.
func (m TradeCaptureReport) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from TradeCaptureReport.
func (m TradeCaptureReport) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from TradeCaptureReport.
func (m TradeCaptureReport) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from TradeCaptureReport.
func (m TradeCaptureReport) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from TradeCaptureReport.
func (m TradeCaptureReport) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from TradeCaptureReport.
func (m TradeCaptureReport) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from TradeCaptureReport.
func (m TradeCaptureReport) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from TradeCaptureReport.
func (m TradeCaptureReport) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from TradeCaptureReport.
func (m TradeCaptureReport) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from TradeCaptureReport.
func (m TradeCaptureReport) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from TradeCaptureReport.
func (m TradeCaptureReport) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from TradeCaptureReport.
func (m TradeCaptureReport) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from TradeCaptureReport.
func (m TradeCaptureReport) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from TradeCaptureReport.
func (m TradeCaptureReport) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from TradeCaptureReport.
func (m TradeCaptureReport) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from TradeCaptureReport.
func (m TradeCaptureReport) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from TradeCaptureReport.
func (m TradeCaptureReport) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from TradeCaptureReport.
func (m TradeCaptureReport) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from TradeCaptureReport.
func (m TradeCaptureReport) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from TradeCaptureReport.
func (m TradeCaptureReport) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from TradeCaptureReport.
func (m TradeCaptureReport) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from TradeCaptureReport.
func (m TradeCaptureReport) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from TradeCaptureReport.
func (m TradeCaptureReport) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from TradeCaptureReport.
func (m TradeCaptureReport) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from TradeCaptureReport.
func (m TradeCaptureReport) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from TradeCaptureReport.
func (m TradeCaptureReport) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from TradeCaptureReport.
func (m TradeCaptureReport) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from TradeCaptureReport.
func (m TradeCaptureReport) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from TradeCaptureReport.
func (m TradeCaptureReport) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from TradeCaptureReport.
func (m TradeCaptureReport) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from TradeCaptureReport.
func (m TradeCaptureReport) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityStatus() (*field.SecurityStatusField, errors.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from TradeCaptureReport.
func (m TradeCaptureReport) GetSecurityStatus(f *field.SecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from TradeCaptureReport.
func (m TradeCaptureReport) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from TradeCaptureReport.
func (m TradeCaptureReport) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikeMultiplier() (*field.StrikeMultiplierField, errors.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from TradeCaptureReport.
func (m TradeCaptureReport) GetStrikeMultiplier(f *field.StrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikeValue() (*field.StrikeValueField, errors.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from TradeCaptureReport.
func (m TradeCaptureReport) GetStrikeValue(f *field.StrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MinPriceIncrement() (*field.MinPriceIncrementField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from TradeCaptureReport.
func (m TradeCaptureReport) GetMinPriceIncrement(f *field.MinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PositionLimit() (*field.PositionLimitField, errors.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from TradeCaptureReport.
func (m TradeCaptureReport) GetPositionLimit(f *field.PositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NTPositionLimit() (*field.NTPositionLimitField, errors.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from TradeCaptureReport.
func (m TradeCaptureReport) GetNTPositionLimit(f *field.NTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoInstrumentParties() (*field.NoInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from TradeCaptureReport.
func (m TradeCaptureReport) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnitOfMeasure() (*field.UnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from TradeCaptureReport.
func (m TradeCaptureReport) GetUnitOfMeasure(f *field.UnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TimeUnit() (*field.TimeUnitField, errors.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from TradeCaptureReport.
func (m TradeCaptureReport) GetTimeUnit(f *field.TimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MaturityTime() (*field.MaturityTimeField, errors.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from TradeCaptureReport.
func (m TradeCaptureReport) GetMaturityTime(f *field.MaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityGroup() (*field.SecurityGroupField, errors.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from TradeCaptureReport.
func (m TradeCaptureReport) GetSecurityGroup(f *field.SecurityGroupField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from TradeCaptureReport.
func (m TradeCaptureReport) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from TradeCaptureReport.
func (m TradeCaptureReport) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityXMLLen() (*field.SecurityXMLLenField, errors.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from TradeCaptureReport.
func (m TradeCaptureReport) GetSecurityXMLLen(f *field.SecurityXMLLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityXML() (*field.SecurityXMLField, errors.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from TradeCaptureReport.
func (m TradeCaptureReport) GetSecurityXML(f *field.SecurityXMLField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityXMLSchema() (*field.SecurityXMLSchemaField, errors.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from TradeCaptureReport.
func (m TradeCaptureReport) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ProductComplex() (*field.ProductComplexField, errors.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from TradeCaptureReport.
func (m TradeCaptureReport) GetProductComplex(f *field.ProductComplexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from TradeCaptureReport.
func (m TradeCaptureReport) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from TradeCaptureReport.
func (m TradeCaptureReport) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlMethod() (*field.SettlMethodField, errors.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from TradeCaptureReport.
func (m TradeCaptureReport) GetSettlMethod(f *field.SettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ExerciseStyle() (*field.ExerciseStyleField, errors.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from TradeCaptureReport.
func (m TradeCaptureReport) GetExerciseStyle(f *field.ExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OptPayoutAmount() (*field.OptPayoutAmountField, errors.MessageRejectError) {
	f := &field.OptPayoutAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from TradeCaptureReport.
func (m TradeCaptureReport) GetOptPayoutAmount(f *field.OptPayoutAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PriceQuoteMethod() (*field.PriceQuoteMethodField, errors.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from TradeCaptureReport.
func (m TradeCaptureReport) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ListMethod() (*field.ListMethodField, errors.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from TradeCaptureReport.
func (m TradeCaptureReport) GetListMethod(f *field.ListMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CapPrice() (*field.CapPriceField, errors.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from TradeCaptureReport.
func (m TradeCaptureReport) GetCapPrice(f *field.CapPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) FloorPrice() (*field.FloorPriceField, errors.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from TradeCaptureReport.
func (m TradeCaptureReport) GetFloorPrice(f *field.FloorPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from TradeCaptureReport.
func (m TradeCaptureReport) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) FlexibleIndicator() (*field.FlexibleIndicatorField, errors.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from TradeCaptureReport.
func (m TradeCaptureReport) GetFlexibleIndicator(f *field.FlexibleIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, errors.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from TradeCaptureReport.
func (m TradeCaptureReport) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ValuationMethod() (*field.ValuationMethodField, errors.MessageRejectError) {
	f := &field.ValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from TradeCaptureReport.
func (m TradeCaptureReport) GetValuationMethod(f *field.ValuationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ContractMultiplierUnit() (*field.ContractMultiplierUnitField, errors.MessageRejectError) {
	f := &field.ContractMultiplierUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from TradeCaptureReport.
func (m TradeCaptureReport) GetContractMultiplierUnit(f *field.ContractMultiplierUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) FlowScheduleType() (*field.FlowScheduleTypeField, errors.MessageRejectError) {
	f := &field.FlowScheduleTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from TradeCaptureReport.
func (m TradeCaptureReport) GetFlowScheduleType(f *field.FlowScheduleTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RestructuringType() (*field.RestructuringTypeField, errors.MessageRejectError) {
	f := &field.RestructuringTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from TradeCaptureReport.
func (m TradeCaptureReport) GetRestructuringType(f *field.RestructuringTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Seniority() (*field.SeniorityField, errors.MessageRejectError) {
	f := &field.SeniorityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from TradeCaptureReport.
func (m TradeCaptureReport) GetSeniority(f *field.SeniorityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstandingField, errors.MessageRejectError) {
	f := &field.NotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from TradeCaptureReport.
func (m TradeCaptureReport) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstandingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstandingField, errors.MessageRejectError) {
	f := &field.OriginalNotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from TradeCaptureReport.
func (m TradeCaptureReport) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstandingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AttachmentPoint() (*field.AttachmentPointField, errors.MessageRejectError) {
	f := &field.AttachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from TradeCaptureReport.
func (m TradeCaptureReport) GetAttachmentPoint(f *field.AttachmentPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) DetachmentPoint() (*field.DetachmentPointField, errors.MessageRejectError) {
	f := &field.DetachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from TradeCaptureReport.
func (m TradeCaptureReport) GetDetachmentPoint(f *field.DetachmentPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethodField, errors.MessageRejectError) {
	f := &field.StrikePriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from TradeCaptureReport.
func (m TradeCaptureReport) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethodField, errors.MessageRejectError) {
	f := &field.StrikePriceBoundaryMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from TradeCaptureReport.
func (m TradeCaptureReport) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecisionField, errors.MessageRejectError) {
	f := &field.StrikePriceBoundaryPrecisionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from TradeCaptureReport.
func (m TradeCaptureReport) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecisionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethodField, errors.MessageRejectError) {
	f := &field.UnderlyingPriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from TradeCaptureReport.
func (m TradeCaptureReport) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OptPayoutType() (*field.OptPayoutTypeField, errors.MessageRejectError) {
	f := &field.OptPayoutTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from TradeCaptureReport.
func (m TradeCaptureReport) GetOptPayoutType(f *field.OptPayoutTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoComplexEvents() (*field.NoComplexEventsField, errors.MessageRejectError) {
	f := &field.NoComplexEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from TradeCaptureReport.
func (m TradeCaptureReport) GetNoComplexEvents(f *field.NoComplexEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementDesc() (*field.AgreementDescField, errors.MessageRejectError) {
	f := &field.AgreementDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from TradeCaptureReport.
func (m TradeCaptureReport) GetAgreementDesc(f *field.AgreementDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementID() (*field.AgreementIDField, errors.MessageRejectError) {
	f := &field.AgreementIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from TradeCaptureReport.
func (m TradeCaptureReport) GetAgreementID(f *field.AgreementIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementDate() (*field.AgreementDateField, errors.MessageRejectError) {
	f := &field.AgreementDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from TradeCaptureReport.
func (m TradeCaptureReport) GetAgreementDate(f *field.AgreementDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementCurrency() (*field.AgreementCurrencyField, errors.MessageRejectError) {
	f := &field.AgreementCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from TradeCaptureReport.
func (m TradeCaptureReport) GetAgreementCurrency(f *field.AgreementCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TerminationType() (*field.TerminationTypeField, errors.MessageRejectError) {
	f := &field.TerminationTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from TradeCaptureReport.
func (m TradeCaptureReport) GetTerminationType(f *field.TerminationTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StartDate() (*field.StartDateField, errors.MessageRejectError) {
	f := &field.StartDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from TradeCaptureReport.
func (m TradeCaptureReport) GetStartDate(f *field.StartDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EndDate() (*field.EndDateField, errors.MessageRejectError) {
	f := &field.EndDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from TradeCaptureReport.
func (m TradeCaptureReport) GetEndDate(f *field.EndDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) DeliveryType() (*field.DeliveryTypeField, errors.MessageRejectError) {
	f := &field.DeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from TradeCaptureReport.
func (m TradeCaptureReport) GetDeliveryType(f *field.DeliveryTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MarginRatio() (*field.MarginRatioField, errors.MessageRejectError) {
	f := &field.MarginRatioField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from TradeCaptureReport.
func (m TradeCaptureReport) GetMarginRatio(f *field.MarginRatioField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) QtyType() (*field.QtyTypeField, errors.MessageRejectError) {
	f := &field.QtyTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from TradeCaptureReport.
func (m TradeCaptureReport) GetQtyType(f *field.QtyTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldType() (*field.YieldTypeField, errors.MessageRejectError) {
	f := &field.YieldTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from TradeCaptureReport.
func (m TradeCaptureReport) GetYieldType(f *field.YieldTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Yield() (*field.YieldField, errors.MessageRejectError) {
	f := &field.YieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from TradeCaptureReport.
func (m TradeCaptureReport) GetYield(f *field.YieldField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldCalcDate() (*field.YieldCalcDateField, errors.MessageRejectError) {
	f := &field.YieldCalcDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from TradeCaptureReport.
func (m TradeCaptureReport) GetYieldCalcDate(f *field.YieldCalcDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldRedemptionDate() (*field.YieldRedemptionDateField, errors.MessageRejectError) {
	f := &field.YieldRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from TradeCaptureReport.
func (m TradeCaptureReport) GetYieldRedemptionDate(f *field.YieldRedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldRedemptionPrice() (*field.YieldRedemptionPriceField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from TradeCaptureReport.
func (m TradeCaptureReport) GetYieldRedemptionPrice(f *field.YieldRedemptionPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldRedemptionPriceType() (*field.YieldRedemptionPriceTypeField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from TradeCaptureReport.
func (m TradeCaptureReport) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from TradeCaptureReport.
func (m TradeCaptureReport) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTradingSessionID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnderlyingTradingSessionID() (*field.UnderlyingTradingSessionIDField, errors.MessageRejectError) {
	f := &field.UnderlyingTradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTradingSessionID reads a UnderlyingTradingSessionID from TradeCaptureReport.
func (m TradeCaptureReport) GetUnderlyingTradingSessionID(f *field.UnderlyingTradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTradingSessionSubID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnderlyingTradingSessionSubID() (*field.UnderlyingTradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.UnderlyingTradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTradingSessionSubID reads a UnderlyingTradingSessionSubID from TradeCaptureReport.
func (m TradeCaptureReport) GetUnderlyingTradingSessionSubID(f *field.UnderlyingTradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastQty is a required field for TradeCaptureReport.
func (m TradeCaptureReport) LastQty() (*field.LastQtyField, errors.MessageRejectError) {
	f := &field.LastQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastQty reads a LastQty from TradeCaptureReport.
func (m TradeCaptureReport) GetLastQty(f *field.LastQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a required field for TradeCaptureReport.
func (m TradeCaptureReport) LastPx() (*field.LastPxField, errors.MessageRejectError) {
	f := &field.LastPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from TradeCaptureReport.
func (m TradeCaptureReport) GetLastPx(f *field.LastPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastParPx is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastParPx() (*field.LastParPxField, errors.MessageRejectError) {
	f := &field.LastParPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastParPx reads a LastParPx from TradeCaptureReport.
func (m TradeCaptureReport) GetLastParPx(f *field.LastParPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastSpotRate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastSpotRate() (*field.LastSpotRateField, errors.MessageRejectError) {
	f := &field.LastSpotRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastSpotRate reads a LastSpotRate from TradeCaptureReport.
func (m TradeCaptureReport) GetLastSpotRate(f *field.LastSpotRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastForwardPoints is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastForwardPoints() (*field.LastForwardPointsField, errors.MessageRejectError) {
	f := &field.LastForwardPointsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastForwardPoints reads a LastForwardPoints from TradeCaptureReport.
func (m TradeCaptureReport) GetLastForwardPoints(f *field.LastForwardPointsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastMkt() (*field.LastMktField, errors.MessageRejectError) {
	f := &field.LastMktField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from TradeCaptureReport.
func (m TradeCaptureReport) GetLastMkt(f *field.LastMktField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from TradeCaptureReport.
func (m TradeCaptureReport) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from TradeCaptureReport.
func (m TradeCaptureReport) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AvgPx() (*field.AvgPxField, errors.MessageRejectError) {
	f := &field.AvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from TradeCaptureReport.
func (m TradeCaptureReport) GetAvgPx(f *field.AvgPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Spread() (*field.SpreadField, errors.MessageRejectError) {
	f := &field.SpreadField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from TradeCaptureReport.
func (m TradeCaptureReport) GetSpread(f *field.SpreadField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrencyField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from TradeCaptureReport.
func (m TradeCaptureReport) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkCurveName() (*field.BenchmarkCurveNameField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from TradeCaptureReport.
func (m TradeCaptureReport) GetBenchmarkCurveName(f *field.BenchmarkCurveNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkCurvePoint() (*field.BenchmarkCurvePointField, errors.MessageRejectError) {
	f := &field.BenchmarkCurvePointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from TradeCaptureReport.
func (m TradeCaptureReport) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkPrice() (*field.BenchmarkPriceField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from TradeCaptureReport.
func (m TradeCaptureReport) GetBenchmarkPrice(f *field.BenchmarkPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkPriceType() (*field.BenchmarkPriceTypeField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from TradeCaptureReport.
func (m TradeCaptureReport) GetBenchmarkPriceType(f *field.BenchmarkPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkSecurityID() (*field.BenchmarkSecurityIDField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from TradeCaptureReport.
func (m TradeCaptureReport) GetBenchmarkSecurityID(f *field.BenchmarkSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from TradeCaptureReport.
func (m TradeCaptureReport) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AvgPxIndicator() (*field.AvgPxIndicatorField, errors.MessageRejectError) {
	f := &field.AvgPxIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxIndicator reads a AvgPxIndicator from TradeCaptureReport.
func (m TradeCaptureReport) GetAvgPxIndicator(f *field.AvgPxIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPosAmt is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoPosAmt() (*field.NoPosAmtField, errors.MessageRejectError) {
	f := &field.NoPosAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPosAmt reads a NoPosAmt from TradeCaptureReport.
func (m TradeCaptureReport) GetNoPosAmt(f *field.NoPosAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegReportingType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MultiLegReportingType() (*field.MultiLegReportingTypeField, errors.MessageRejectError) {
	f := &field.MultiLegReportingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegReportingType reads a MultiLegReportingType from TradeCaptureReport.
func (m TradeCaptureReport) GetMultiLegReportingType(f *field.MultiLegReportingTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeLegRefID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeLegRefID() (*field.TradeLegRefIDField, errors.MessageRejectError) {
	f := &field.TradeLegRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeLegRefID reads a TradeLegRefID from TradeCaptureReport.
func (m TradeCaptureReport) GetTradeLegRefID(f *field.TradeLegRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from TradeCaptureReport.
func (m TradeCaptureReport) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from TradeCaptureReport.
func (m TradeCaptureReport) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRegTimestamps is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoTrdRegTimestamps() (*field.NoTrdRegTimestampsField, errors.MessageRejectError) {
	f := &field.NoTrdRegTimestampsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRegTimestamps reads a NoTrdRegTimestamps from TradeCaptureReport.
func (m TradeCaptureReport) GetNoTrdRegTimestamps(f *field.NoTrdRegTimestampsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlType() (*field.SettlTypeField, errors.MessageRejectError) {
	f := &field.SettlTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from TradeCaptureReport.
func (m TradeCaptureReport) GetSettlType(f *field.SettlTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlDate() (*field.SettlDateField, errors.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from TradeCaptureReport.
func (m TradeCaptureReport) GetSettlDate(f *field.SettlDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MatchStatus() (*field.MatchStatusField, errors.MessageRejectError) {
	f := &field.MatchStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from TradeCaptureReport.
func (m TradeCaptureReport) GetMatchStatus(f *field.MatchStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MatchType() (*field.MatchTypeField, errors.MessageRejectError) {
	f := &field.MatchTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchType reads a MatchType from TradeCaptureReport.
func (m TradeCaptureReport) GetMatchType(f *field.MatchTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSides is a required field for TradeCaptureReport.
func (m TradeCaptureReport) NoSides() (*field.NoSidesField, errors.MessageRejectError) {
	f := &field.NoSidesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSides reads a NoSides from TradeCaptureReport.
func (m TradeCaptureReport) GetNoSides(f *field.NoSidesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CopyMsgIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CopyMsgIndicator() (*field.CopyMsgIndicatorField, errors.MessageRejectError) {
	f := &field.CopyMsgIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCopyMsgIndicator reads a CopyMsgIndicator from TradeCaptureReport.
func (m TradeCaptureReport) GetCopyMsgIndicator(f *field.CopyMsgIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PublishTrdIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PublishTrdIndicator() (*field.PublishTrdIndicatorField, errors.MessageRejectError) {
	f := &field.PublishTrdIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPublishTrdIndicator reads a PublishTrdIndicator from TradeCaptureReport.
func (m TradeCaptureReport) GetPublishTrdIndicator(f *field.PublishTrdIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ShortSaleReason is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ShortSaleReason() (*field.ShortSaleReasonField, errors.MessageRejectError) {
	f := &field.ShortSaleReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetShortSaleReason reads a ShortSaleReason from TradeCaptureReport.
func (m TradeCaptureReport) GetShortSaleReason(f *field.ShortSaleReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdRptStatus is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TrdRptStatus() (*field.TrdRptStatusField, errors.MessageRejectError) {
	f := &field.TrdRptStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdRptStatus reads a TrdRptStatus from TradeCaptureReport.
func (m TradeCaptureReport) GetTrdRptStatus(f *field.TrdRptStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AsOfIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AsOfIndicator() (*field.AsOfIndicatorField, errors.MessageRejectError) {
	f := &field.AsOfIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAsOfIndicator reads a AsOfIndicator from TradeCaptureReport.
func (m TradeCaptureReport) GetAsOfIndicator(f *field.AsOfIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlSessID() (*field.SettlSessIDField, errors.MessageRejectError) {
	f := &field.SettlSessIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from TradeCaptureReport.
func (m TradeCaptureReport) GetSettlSessID(f *field.SettlSessIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlSessSubID() (*field.SettlSessSubIDField, errors.MessageRejectError) {
	f := &field.SettlSessSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from TradeCaptureReport.
func (m TradeCaptureReport) GetSettlSessSubID(f *field.SettlSessSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TierCode is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TierCode() (*field.TierCodeField, errors.MessageRejectError) {
	f := &field.TierCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTierCode reads a TierCode from TradeCaptureReport.
func (m TradeCaptureReport) GetTierCode(f *field.TierCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageEventSource is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MessageEventSource() (*field.MessageEventSourceField, errors.MessageRejectError) {
	f := &field.MessageEventSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMessageEventSource reads a MessageEventSource from TradeCaptureReport.
func (m TradeCaptureReport) GetMessageEventSource(f *field.MessageEventSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastUpdateTime is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastUpdateTime() (*field.LastUpdateTimeField, errors.MessageRejectError) {
	f := &field.LastUpdateTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastUpdateTime reads a LastUpdateTime from TradeCaptureReport.
func (m TradeCaptureReport) GetLastUpdateTime(f *field.LastUpdateTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RndPx is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RndPx() (*field.RndPxField, errors.MessageRejectError) {
	f := &field.RndPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRndPx reads a RndPx from TradeCaptureReport.
func (m TradeCaptureReport) GetRndPx(f *field.RndPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeID() (*field.TradeIDField, errors.MessageRejectError) {
	f := &field.TradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeID reads a TradeID from TradeCaptureReport.
func (m TradeCaptureReport) GetTradeID(f *field.TradeIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryTradeID() (*field.SecondaryTradeIDField, errors.MessageRejectError) {
	f := &field.SecondaryTradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTradeID reads a SecondaryTradeID from TradeCaptureReport.
func (m TradeCaptureReport) GetSecondaryTradeID(f *field.SecondaryTradeIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FirmTradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) FirmTradeID() (*field.FirmTradeIDField, errors.MessageRejectError) {
	f := &field.FirmTradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFirmTradeID reads a FirmTradeID from TradeCaptureReport.
func (m TradeCaptureReport) GetFirmTradeID(f *field.FirmTradeIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryFirmTradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryFirmTradeID() (*field.SecondaryFirmTradeIDField, errors.MessageRejectError) {
	f := &field.SecondaryFirmTradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryFirmTradeID reads a SecondaryFirmTradeID from TradeCaptureReport.
func (m TradeCaptureReport) GetSecondaryFirmTradeID(f *field.SecondaryFirmTradeIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CalculatedCcyLastQty is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CalculatedCcyLastQty() (*field.CalculatedCcyLastQtyField, errors.MessageRejectError) {
	f := &field.CalculatedCcyLastQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCalculatedCcyLastQty reads a CalculatedCcyLastQty from TradeCaptureReport.
func (m TradeCaptureReport) GetCalculatedCcyLastQty(f *field.CalculatedCcyLastQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastSwapPoints is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastSwapPoints() (*field.LastSwapPointsField, errors.MessageRejectError) {
	f := &field.LastSwapPointsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastSwapPoints reads a LastSwapPoints from TradeCaptureReport.
func (m TradeCaptureReport) GetLastSwapPoints(f *field.LastSwapPointsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlementDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnderlyingSettlementDate() (*field.UnderlyingSettlementDateField, errors.MessageRejectError) {
	f := &field.UnderlyingSettlementDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlementDate reads a UnderlyingSettlementDate from TradeCaptureReport.
func (m TradeCaptureReport) GetUnderlyingSettlementDate(f *field.UnderlyingSettlementDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) GrossTradeAmt() (*field.GrossTradeAmtField, errors.MessageRejectError) {
	f := &field.GrossTradeAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from TradeCaptureReport.
func (m TradeCaptureReport) GetGrossTradeAmt(f *field.GrossTradeAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRootPartyIDs is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoRootPartyIDs() (*field.NoRootPartyIDsField, errors.MessageRejectError) {
	f := &field.NoRootPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRootPartyIDs reads a NoRootPartyIDs from TradeCaptureReport.
func (m TradeCaptureReport) GetNoRootPartyIDs(f *field.NoRootPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeHandlingInstr is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeHandlingInstr() (*field.TradeHandlingInstrField, errors.MessageRejectError) {
	f := &field.TradeHandlingInstrField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeHandlingInstr reads a TradeHandlingInstr from TradeCaptureReport.
func (m TradeCaptureReport) GetTradeHandlingInstr(f *field.TradeHandlingInstrField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTradeHandlingInstr is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrigTradeHandlingInstr() (*field.OrigTradeHandlingInstrField, errors.MessageRejectError) {
	f := &field.OrigTradeHandlingInstrField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTradeHandlingInstr reads a OrigTradeHandlingInstr from TradeCaptureReport.
func (m TradeCaptureReport) GetOrigTradeHandlingInstr(f *field.OrigTradeHandlingInstrField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTradeDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrigTradeDate() (*field.OrigTradeDateField, errors.MessageRejectError) {
	f := &field.OrigTradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTradeDate reads a OrigTradeDate from TradeCaptureReport.
func (m TradeCaptureReport) GetOrigTradeDate(f *field.OrigTradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrigTradeID() (*field.OrigTradeIDField, errors.MessageRejectError) {
	f := &field.OrigTradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTradeID reads a OrigTradeID from TradeCaptureReport.
func (m TradeCaptureReport) GetOrigTradeID(f *field.OrigTradeIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigSecondaryTradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrigSecondaryTradeID() (*field.OrigSecondaryTradeIDField, errors.MessageRejectError) {
	f := &field.OrigSecondaryTradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigSecondaryTradeID reads a OrigSecondaryTradeID from TradeCaptureReport.
func (m TradeCaptureReport) GetOrigSecondaryTradeID(f *field.OrigSecondaryTradeIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TZTransactTime is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TZTransactTime() (*field.TZTransactTimeField, errors.MessageRejectError) {
	f := &field.TZTransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTZTransactTime reads a TZTransactTime from TradeCaptureReport.
func (m TradeCaptureReport) GetTZTransactTime(f *field.TZTransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ReportedPxDiff is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ReportedPxDiff() (*field.ReportedPxDiffField, errors.MessageRejectError) {
	f := &field.ReportedPxDiffField{}
	err := m.Body.Get(f)
	return f, err
}

//GetReportedPxDiff reads a ReportedPxDiff from TradeCaptureReport.
func (m TradeCaptureReport) GetReportedPxDiff(f *field.ReportedPxDiffField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from TradeCaptureReport.
func (m TradeCaptureReport) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlCurrency() (*field.SettlCurrencyField, errors.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from TradeCaptureReport.
func (m TradeCaptureReport) GetSettlCurrency(f *field.SettlCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RejectText is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RejectText() (*field.RejectTextField, errors.MessageRejectError) {
	f := &field.RejectTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRejectText reads a RejectText from TradeCaptureReport.
func (m TradeCaptureReport) GetRejectText(f *field.RejectTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FeeMultiplier is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) FeeMultiplier() (*field.FeeMultiplierField, errors.MessageRejectError) {
	f := &field.FeeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFeeMultiplier reads a FeeMultiplier from TradeCaptureReport.
func (m TradeCaptureReport) GetFeeMultiplier(f *field.FeeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Volatility is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Volatility() (*field.VolatilityField, errors.MessageRejectError) {
	f := &field.VolatilityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetVolatility reads a Volatility from TradeCaptureReport.
func (m TradeCaptureReport) GetVolatility(f *field.VolatilityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DividendYield is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) DividendYield() (*field.DividendYieldField, errors.MessageRejectError) {
	f := &field.DividendYieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDividendYield reads a DividendYield from TradeCaptureReport.
func (m TradeCaptureReport) GetDividendYield(f *field.DividendYieldField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RiskFreeRate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RiskFreeRate() (*field.RiskFreeRateField, errors.MessageRejectError) {
	f := &field.RiskFreeRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRiskFreeRate reads a RiskFreeRate from TradeCaptureReport.
func (m TradeCaptureReport) GetRiskFreeRate(f *field.RiskFreeRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CurrencyRatio is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CurrencyRatio() (*field.CurrencyRatioField, errors.MessageRejectError) {
	f := &field.CurrencyRatioField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrencyRatio reads a CurrencyRatio from TradeCaptureReport.
func (m TradeCaptureReport) GetCurrencyRatio(f *field.CurrencyRatioField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRepIndicators is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoTrdRepIndicators() (*field.NoTrdRepIndicatorsField, errors.MessageRejectError) {
	f := &field.NoTrdRepIndicatorsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRepIndicators reads a NoTrdRepIndicators from TradeCaptureReport.
func (m TradeCaptureReport) GetNoTrdRepIndicators(f *field.NoTrdRepIndicatorsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradePublishIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradePublishIndicator() (*field.TradePublishIndicatorField, errors.MessageRejectError) {
	f := &field.TradePublishIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradePublishIndicator reads a TradePublishIndicator from TradeCaptureReport.
func (m TradeCaptureReport) GetTradePublishIndicator(f *field.TradePublishIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from TradeCaptureReport.
func (m TradeCaptureReport) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from TradeCaptureReport.
func (m TradeCaptureReport) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from TradeCaptureReport.
func (m TradeCaptureReport) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from TradeCaptureReport.
func (m TradeCaptureReport) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//VenueType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) VenueType() (*field.VenueTypeField, errors.MessageRejectError) {
	f := &field.VenueTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetVenueType reads a VenueType from TradeCaptureReport.
func (m TradeCaptureReport) GetVenueType(f *field.VenueTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MarketSegmentID() (*field.MarketSegmentIDField, errors.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from TradeCaptureReport.
func (m TradeCaptureReport) GetMarketSegmentID(f *field.MarketSegmentIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MarketID() (*field.MarketIDField, errors.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from TradeCaptureReport.
func (m TradeCaptureReport) GetMarketID(f *field.MarketIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}
