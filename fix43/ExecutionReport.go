package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ExecutionReport msg type = 8.
type ExecutionReport struct {
	message.Message
}

//ExecutionReportBuilder builds ExecutionReport messages.
type ExecutionReportBuilder struct {
	message.MessageBuilder
}

//CreateExecutionReportBuilder returns an initialized ExecutionReportBuilder with specified required fields.
func CreateExecutionReportBuilder(
	orderid *field.OrderIDField,
	execid *field.ExecIDField,
	exectype *field.ExecTypeField,
	ordstatus *field.OrdStatusField,
	side *field.SideField,
	leavesqty *field.LeavesQtyField,
	cumqty *field.CumQtyField,
	avgpx *field.AvgPxField) ExecutionReportBuilder {
	var builder ExecutionReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header.Set(field.NewMsgType("8"))
	builder.Body.Set(orderid)
	builder.Body.Set(execid)
	builder.Body.Set(exectype)
	builder.Body.Set(ordstatus)
	builder.Body.Set(side)
	builder.Body.Set(leavesqty)
	builder.Body.Set(cumqty)
	builder.Body.Set(avgpx)
	return builder
}

//OrderID is a required field for ExecutionReport.
func (m ExecutionReport) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from ExecutionReport.
func (m ExecutionReport) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryOrderID() (*field.SecondaryOrderIDField, errors.MessageRejectError) {
	f := &field.SecondaryOrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from ExecutionReport.
func (m ExecutionReport) GetSecondaryOrderID(f *field.SecondaryOrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryClOrdID() (*field.SecondaryClOrdIDField, errors.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from ExecutionReport.
func (m ExecutionReport) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryExecID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryExecID() (*field.SecondaryExecIDField, errors.MessageRejectError) {
	f := &field.SecondaryExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryExecID reads a SecondaryExecID from ExecutionReport.
func (m ExecutionReport) GetSecondaryExecID(f *field.SecondaryExecIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from ExecutionReport.
func (m ExecutionReport) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) OrigClOrdID() (*field.OrigClOrdIDField, errors.MessageRejectError) {
	f := &field.OrigClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from ExecutionReport.
func (m ExecutionReport) GetOrigClOrdID(f *field.OrigClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for ExecutionReport.
func (m ExecutionReport) ClOrdLinkID() (*field.ClOrdLinkIDField, errors.MessageRejectError) {
	f := &field.ClOrdLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from ExecutionReport.
func (m ExecutionReport) GetClOrdLinkID(f *field.ClOrdLinkIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for ExecutionReport.
func (m ExecutionReport) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from ExecutionReport.
func (m ExecutionReport) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for ExecutionReport.
func (m ExecutionReport) TradeOriginationDate() (*field.TradeOriginationDateField, errors.MessageRejectError) {
	f := &field.TradeOriginationDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from ExecutionReport.
func (m ExecutionReport) GetTradeOriginationDate(f *field.TradeOriginationDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoContraBrokers is a non-required field for ExecutionReport.
func (m ExecutionReport) NoContraBrokers() (*field.NoContraBrokersField, errors.MessageRejectError) {
	f := &field.NoContraBrokersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoContraBrokers reads a NoContraBrokers from ExecutionReport.
func (m ExecutionReport) GetNoContraBrokers(f *field.NoContraBrokersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for ExecutionReport.
func (m ExecutionReport) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ExecutionReport.
func (m ExecutionReport) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossID is a non-required field for ExecutionReport.
func (m ExecutionReport) CrossID() (*field.CrossIDField, errors.MessageRejectError) {
	f := &field.CrossIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossID reads a CrossID from ExecutionReport.
func (m ExecutionReport) GetCrossID(f *field.CrossIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigCrossID is a non-required field for ExecutionReport.
func (m ExecutionReport) OrigCrossID() (*field.OrigCrossIDField, errors.MessageRejectError) {
	f := &field.OrigCrossIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigCrossID reads a OrigCrossID from ExecutionReport.
func (m ExecutionReport) GetOrigCrossID(f *field.OrigCrossIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossType is a non-required field for ExecutionReport.
func (m ExecutionReport) CrossType() (*field.CrossTypeField, errors.MessageRejectError) {
	f := &field.CrossTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossType reads a CrossType from ExecutionReport.
func (m ExecutionReport) GetCrossType(f *field.CrossTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecID is a required field for ExecutionReport.
func (m ExecutionReport) ExecID() (*field.ExecIDField, errors.MessageRejectError) {
	f := &field.ExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecID reads a ExecID from ExecutionReport.
func (m ExecutionReport) GetExecID(f *field.ExecIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecRefID is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecRefID() (*field.ExecRefIDField, errors.MessageRejectError) {
	f := &field.ExecRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecRefID reads a ExecRefID from ExecutionReport.
func (m ExecutionReport) GetExecRefID(f *field.ExecRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecType is a required field for ExecutionReport.
func (m ExecutionReport) ExecType() (*field.ExecTypeField, errors.MessageRejectError) {
	f := &field.ExecTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecType reads a ExecType from ExecutionReport.
func (m ExecutionReport) GetExecType(f *field.ExecTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdStatus is a required field for ExecutionReport.
func (m ExecutionReport) OrdStatus() (*field.OrdStatusField, errors.MessageRejectError) {
	f := &field.OrdStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdStatus reads a OrdStatus from ExecutionReport.
func (m ExecutionReport) GetOrdStatus(f *field.OrdStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//WorkingIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) WorkingIndicator() (*field.WorkingIndicatorField, errors.MessageRejectError) {
	f := &field.WorkingIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetWorkingIndicator reads a WorkingIndicator from ExecutionReport.
func (m ExecutionReport) GetWorkingIndicator(f *field.WorkingIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdRejReason is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdRejReason() (*field.OrdRejReasonField, errors.MessageRejectError) {
	f := &field.OrdRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdRejReason reads a OrdRejReason from ExecutionReport.
func (m ExecutionReport) GetOrdRejReason(f *field.OrdRejReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecRestatementReason is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecRestatementReason() (*field.ExecRestatementReasonField, errors.MessageRejectError) {
	f := &field.ExecRestatementReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecRestatementReason reads a ExecRestatementReason from ExecutionReport.
func (m ExecutionReport) GetExecRestatementReason(f *field.ExecRestatementReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for ExecutionReport.
func (m ExecutionReport) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from ExecutionReport.
func (m ExecutionReport) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for ExecutionReport.
func (m ExecutionReport) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from ExecutionReport.
func (m ExecutionReport) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayBookingInst is a non-required field for ExecutionReport.
func (m ExecutionReport) DayBookingInst() (*field.DayBookingInstField, errors.MessageRejectError) {
	f := &field.DayBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayBookingInst reads a DayBookingInst from ExecutionReport.
func (m ExecutionReport) GetDayBookingInst(f *field.DayBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingUnit is a non-required field for ExecutionReport.
func (m ExecutionReport) BookingUnit() (*field.BookingUnitField, errors.MessageRejectError) {
	f := &field.BookingUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingUnit reads a BookingUnit from ExecutionReport.
func (m ExecutionReport) GetBookingUnit(f *field.BookingUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreallocMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) PreallocMethod() (*field.PreallocMethodField, errors.MessageRejectError) {
	f := &field.PreallocMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreallocMethod reads a PreallocMethod from ExecutionReport.
func (m ExecutionReport) GetPreallocMethod(f *field.PreallocMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlmntTyp() (*field.SettlmntTypField, errors.MessageRejectError) {
	f := &field.SettlmntTypField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from ExecutionReport.
func (m ExecutionReport) GetSettlmntTyp(f *field.SettlmntTypField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for ExecutionReport.
func (m ExecutionReport) FutSettDate() (*field.FutSettDateField, errors.MessageRejectError) {
	f := &field.FutSettDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from ExecutionReport.
func (m ExecutionReport) GetFutSettDate(f *field.FutSettDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashMargin is a non-required field for ExecutionReport.
func (m ExecutionReport) CashMargin() (*field.CashMarginField, errors.MessageRejectError) {
	f := &field.CashMarginField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashMargin reads a CashMargin from ExecutionReport.
func (m ExecutionReport) GetCashMargin(f *field.CashMarginField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) ClearingFeeIndicator() (*field.ClearingFeeIndicatorField, errors.MessageRejectError) {
	f := &field.ClearingFeeIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from ExecutionReport.
func (m ExecutionReport) GetClearingFeeIndicator(f *field.ClearingFeeIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for ExecutionReport.
func (m ExecutionReport) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from ExecutionReport.
func (m ExecutionReport) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for ExecutionReport.
func (m ExecutionReport) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from ExecutionReport.
func (m ExecutionReport) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from ExecutionReport.
func (m ExecutionReport) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from ExecutionReport.
func (m ExecutionReport) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for ExecutionReport.
func (m ExecutionReport) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from ExecutionReport.
func (m ExecutionReport) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for ExecutionReport.
func (m ExecutionReport) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from ExecutionReport.
func (m ExecutionReport) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for ExecutionReport.
func (m ExecutionReport) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from ExecutionReport.
func (m ExecutionReport) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from ExecutionReport.
func (m ExecutionReport) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from ExecutionReport.
func (m ExecutionReport) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from ExecutionReport.
func (m ExecutionReport) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for ExecutionReport.
func (m ExecutionReport) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from ExecutionReport.
func (m ExecutionReport) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for ExecutionReport.
func (m ExecutionReport) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from ExecutionReport.
func (m ExecutionReport) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for ExecutionReport.
func (m ExecutionReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from ExecutionReport.
func (m ExecutionReport) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for ExecutionReport.
func (m ExecutionReport) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from ExecutionReport.
func (m ExecutionReport) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for ExecutionReport.
func (m ExecutionReport) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from ExecutionReport.
func (m ExecutionReport) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for ExecutionReport.
func (m ExecutionReport) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from ExecutionReport.
func (m ExecutionReport) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for ExecutionReport.
func (m ExecutionReport) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from ExecutionReport.
func (m ExecutionReport) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for ExecutionReport.
func (m ExecutionReport) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from ExecutionReport.
func (m ExecutionReport) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for ExecutionReport.
func (m ExecutionReport) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from ExecutionReport.
func (m ExecutionReport) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for ExecutionReport.
func (m ExecutionReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from ExecutionReport.
func (m ExecutionReport) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for ExecutionReport.
func (m ExecutionReport) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from ExecutionReport.
func (m ExecutionReport) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for ExecutionReport.
func (m ExecutionReport) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from ExecutionReport.
func (m ExecutionReport) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from ExecutionReport.
func (m ExecutionReport) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for ExecutionReport.
func (m ExecutionReport) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from ExecutionReport.
func (m ExecutionReport) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for ExecutionReport.
func (m ExecutionReport) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from ExecutionReport.
func (m ExecutionReport) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for ExecutionReport.
func (m ExecutionReport) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from ExecutionReport.
func (m ExecutionReport) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from ExecutionReport.
func (m ExecutionReport) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for ExecutionReport.
func (m ExecutionReport) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from ExecutionReport.
func (m ExecutionReport) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from ExecutionReport.
func (m ExecutionReport) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from ExecutionReport.
func (m ExecutionReport) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from ExecutionReport.
func (m ExecutionReport) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from ExecutionReport.
func (m ExecutionReport) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from ExecutionReport.
func (m ExecutionReport) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for ExecutionReport.
func (m ExecutionReport) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from ExecutionReport.
func (m ExecutionReport) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for ExecutionReport.
func (m ExecutionReport) NoStipulations() (*field.NoStipulationsField, errors.MessageRejectError) {
	f := &field.NoStipulationsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from ExecutionReport.
func (m ExecutionReport) GetNoStipulations(f *field.NoStipulationsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuantityType is a non-required field for ExecutionReport.
func (m ExecutionReport) QuantityType() (*field.QuantityTypeField, errors.MessageRejectError) {
	f := &field.QuantityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuantityType reads a QuantityType from ExecutionReport.
func (m ExecutionReport) GetQuantityType(f *field.QuantityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderQty() (*field.OrderQtyField, errors.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from ExecutionReport.
func (m ExecutionReport) GetOrderQty(f *field.OrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) CashOrderQty() (*field.CashOrderQtyField, errors.MessageRejectError) {
	f := &field.CashOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from ExecutionReport.
func (m ExecutionReport) GetCashOrderQty(f *field.CashOrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderPercent() (*field.OrderPercentField, errors.MessageRejectError) {
	f := &field.OrderPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from ExecutionReport.
func (m ExecutionReport) GetOrderPercent(f *field.OrderPercentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for ExecutionReport.
func (m ExecutionReport) RoundingDirection() (*field.RoundingDirectionField, errors.MessageRejectError) {
	f := &field.RoundingDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from ExecutionReport.
func (m ExecutionReport) GetRoundingDirection(f *field.RoundingDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for ExecutionReport.
func (m ExecutionReport) RoundingModulus() (*field.RoundingModulusField, errors.MessageRejectError) {
	f := &field.RoundingModulusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from ExecutionReport.
func (m ExecutionReport) GetRoundingModulus(f *field.RoundingModulusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdType() (*field.OrdTypeField, errors.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from ExecutionReport.
func (m ExecutionReport) GetOrdType(f *field.OrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from ExecutionReport.
func (m ExecutionReport) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for ExecutionReport.
func (m ExecutionReport) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from ExecutionReport.
func (m ExecutionReport) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for ExecutionReport.
func (m ExecutionReport) StopPx() (*field.StopPxField, errors.MessageRejectError) {
	f := &field.StopPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from ExecutionReport.
func (m ExecutionReport) GetStopPx(f *field.StopPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegDifference is a non-required field for ExecutionReport.
func (m ExecutionReport) PegDifference() (*field.PegDifferenceField, errors.MessageRejectError) {
	f := &field.PegDifferenceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegDifference reads a PegDifference from ExecutionReport.
func (m ExecutionReport) GetPegDifference(f *field.PegDifferenceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionInst() (*field.DiscretionInstField, errors.MessageRejectError) {
	f := &field.DiscretionInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from ExecutionReport.
func (m ExecutionReport) GetDiscretionInst(f *field.DiscretionInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffset is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionOffset() (*field.DiscretionOffsetField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffset reads a DiscretionOffset from ExecutionReport.
func (m ExecutionReport) GetDiscretionOffset(f *field.DiscretionOffsetField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for ExecutionReport.
func (m ExecutionReport) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from ExecutionReport.
func (m ExecutionReport) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for ExecutionReport.
func (m ExecutionReport) ComplianceID() (*field.ComplianceIDField, errors.MessageRejectError) {
	f := &field.ComplianceIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from ExecutionReport.
func (m ExecutionReport) GetComplianceID(f *field.ComplianceIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for ExecutionReport.
func (m ExecutionReport) SolicitedFlag() (*field.SolicitedFlagField, errors.MessageRejectError) {
	f := &field.SolicitedFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from ExecutionReport.
func (m ExecutionReport) GetSolicitedFlag(f *field.SolicitedFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeInForce() (*field.TimeInForceField, errors.MessageRejectError) {
	f := &field.TimeInForceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from ExecutionReport.
func (m ExecutionReport) GetTimeInForce(f *field.TimeInForceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for ExecutionReport.
func (m ExecutionReport) EffectiveTime() (*field.EffectiveTimeField, errors.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from ExecutionReport.
func (m ExecutionReport) GetEffectiveTime(f *field.EffectiveTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for ExecutionReport.
func (m ExecutionReport) ExpireDate() (*field.ExpireDateField, errors.MessageRejectError) {
	f := &field.ExpireDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from ExecutionReport.
func (m ExecutionReport) GetExpireDate(f *field.ExpireDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for ExecutionReport.
func (m ExecutionReport) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from ExecutionReport.
func (m ExecutionReport) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecInst() (*field.ExecInstField, errors.MessageRejectError) {
	f := &field.ExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from ExecutionReport.
func (m ExecutionReport) GetExecInst(f *field.ExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderCapacity() (*field.OrderCapacityField, errors.MessageRejectError) {
	f := &field.OrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from ExecutionReport.
func (m ExecutionReport) GetOrderCapacity(f *field.OrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderRestrictions() (*field.OrderRestrictionsField, errors.MessageRejectError) {
	f := &field.OrderRestrictionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from ExecutionReport.
func (m ExecutionReport) GetOrderRestrictions(f *field.OrderRestrictionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for ExecutionReport.
func (m ExecutionReport) CustOrderCapacity() (*field.CustOrderCapacityField, errors.MessageRejectError) {
	f := &field.CustOrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from ExecutionReport.
func (m ExecutionReport) GetCustOrderCapacity(f *field.CustOrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Rule80A is a non-required field for ExecutionReport.
func (m ExecutionReport) Rule80A() (*field.Rule80AField, errors.MessageRejectError) {
	f := &field.Rule80AField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRule80A reads a Rule80A from ExecutionReport.
func (m ExecutionReport) GetRule80A(f *field.Rule80AField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastQty is a non-required field for ExecutionReport.
func (m ExecutionReport) LastQty() (*field.LastQtyField, errors.MessageRejectError) {
	f := &field.LastQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastQty reads a LastQty from ExecutionReport.
func (m ExecutionReport) GetLastQty(f *field.LastQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLastQty is a non-required field for ExecutionReport.
func (m ExecutionReport) UnderlyingLastQty() (*field.UnderlyingLastQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingLastQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLastQty reads a UnderlyingLastQty from ExecutionReport.
func (m ExecutionReport) GetUnderlyingLastQty(f *field.UnderlyingLastQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a non-required field for ExecutionReport.
func (m ExecutionReport) LastPx() (*field.LastPxField, errors.MessageRejectError) {
	f := &field.LastPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from ExecutionReport.
func (m ExecutionReport) GetLastPx(f *field.LastPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLastPx is a non-required field for ExecutionReport.
func (m ExecutionReport) UnderlyingLastPx() (*field.UnderlyingLastPxField, errors.MessageRejectError) {
	f := &field.UnderlyingLastPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLastPx reads a UnderlyingLastPx from ExecutionReport.
func (m ExecutionReport) GetUnderlyingLastPx(f *field.UnderlyingLastPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastSpotRate is a non-required field for ExecutionReport.
func (m ExecutionReport) LastSpotRate() (*field.LastSpotRateField, errors.MessageRejectError) {
	f := &field.LastSpotRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastSpotRate reads a LastSpotRate from ExecutionReport.
func (m ExecutionReport) GetLastSpotRate(f *field.LastSpotRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastForwardPoints is a non-required field for ExecutionReport.
func (m ExecutionReport) LastForwardPoints() (*field.LastForwardPointsField, errors.MessageRejectError) {
	f := &field.LastForwardPointsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastForwardPoints reads a LastForwardPoints from ExecutionReport.
func (m ExecutionReport) GetLastForwardPoints(f *field.LastForwardPointsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for ExecutionReport.
func (m ExecutionReport) LastMkt() (*field.LastMktField, errors.MessageRejectError) {
	f := &field.LastMktField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from ExecutionReport.
func (m ExecutionReport) GetLastMkt(f *field.LastMktField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for ExecutionReport.
func (m ExecutionReport) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from ExecutionReport.
func (m ExecutionReport) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for ExecutionReport.
func (m ExecutionReport) TradingSessionSubID() (*field.TradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from ExecutionReport.
func (m ExecutionReport) GetTradingSessionSubID(f *field.TradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastCapacity is a non-required field for ExecutionReport.
func (m ExecutionReport) LastCapacity() (*field.LastCapacityField, errors.MessageRejectError) {
	f := &field.LastCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastCapacity reads a LastCapacity from ExecutionReport.
func (m ExecutionReport) GetLastCapacity(f *field.LastCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LeavesQty is a required field for ExecutionReport.
func (m ExecutionReport) LeavesQty() (*field.LeavesQtyField, errors.MessageRejectError) {
	f := &field.LeavesQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLeavesQty reads a LeavesQty from ExecutionReport.
func (m ExecutionReport) GetLeavesQty(f *field.LeavesQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CumQty is a required field for ExecutionReport.
func (m ExecutionReport) CumQty() (*field.CumQtyField, errors.MessageRejectError) {
	f := &field.CumQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCumQty reads a CumQty from ExecutionReport.
func (m ExecutionReport) GetCumQty(f *field.CumQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a required field for ExecutionReport.
func (m ExecutionReport) AvgPx() (*field.AvgPxField, errors.MessageRejectError) {
	f := &field.AvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from ExecutionReport.
func (m ExecutionReport) GetAvgPx(f *field.AvgPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayOrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DayOrderQty() (*field.DayOrderQtyField, errors.MessageRejectError) {
	f := &field.DayOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayOrderQty reads a DayOrderQty from ExecutionReport.
func (m ExecutionReport) GetDayOrderQty(f *field.DayOrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayCumQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DayCumQty() (*field.DayCumQtyField, errors.MessageRejectError) {
	f := &field.DayCumQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayCumQty reads a DayCumQty from ExecutionReport.
func (m ExecutionReport) GetDayCumQty(f *field.DayCumQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayAvgPx is a non-required field for ExecutionReport.
func (m ExecutionReport) DayAvgPx() (*field.DayAvgPxField, errors.MessageRejectError) {
	f := &field.DayAvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayAvgPx reads a DayAvgPx from ExecutionReport.
func (m ExecutionReport) GetDayAvgPx(f *field.DayAvgPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for ExecutionReport.
func (m ExecutionReport) GTBookingInst() (*field.GTBookingInstField, errors.MessageRejectError) {
	f := &field.GTBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from ExecutionReport.
func (m ExecutionReport) GetGTBookingInst(f *field.GTBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for ExecutionReport.
func (m ExecutionReport) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from ExecutionReport.
func (m ExecutionReport) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for ExecutionReport.
func (m ExecutionReport) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ExecutionReport.
func (m ExecutionReport) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ReportToExch is a non-required field for ExecutionReport.
func (m ExecutionReport) ReportToExch() (*field.ReportToExchField, errors.MessageRejectError) {
	f := &field.ReportToExchField{}
	err := m.Body.Get(f)
	return f, err
}

//GetReportToExch reads a ReportToExch from ExecutionReport.
func (m ExecutionReport) GetReportToExch(f *field.ReportToExchField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for ExecutionReport.
func (m ExecutionReport) Commission() (*field.CommissionField, errors.MessageRejectError) {
	f := &field.CommissionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from ExecutionReport.
func (m ExecutionReport) GetCommission(f *field.CommissionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for ExecutionReport.
func (m ExecutionReport) CommType() (*field.CommTypeField, errors.MessageRejectError) {
	f := &field.CommTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from ExecutionReport.
func (m ExecutionReport) GetCommType(f *field.CommTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) CommCurrency() (*field.CommCurrencyField, errors.MessageRejectError) {
	f := &field.CommCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommCurrency reads a CommCurrency from ExecutionReport.
func (m ExecutionReport) GetCommCurrency(f *field.CommCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FundRenewWaiv is a non-required field for ExecutionReport.
func (m ExecutionReport) FundRenewWaiv() (*field.FundRenewWaivField, errors.MessageRejectError) {
	f := &field.FundRenewWaivField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFundRenewWaiv reads a FundRenewWaiv from ExecutionReport.
func (m ExecutionReport) GetFundRenewWaiv(f *field.FundRenewWaivField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for ExecutionReport.
func (m ExecutionReport) Spread() (*field.SpreadField, errors.MessageRejectError) {
	f := &field.SpreadField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from ExecutionReport.
func (m ExecutionReport) GetSpread(f *field.SpreadField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrencyField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from ExecutionReport.
func (m ExecutionReport) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkCurveName() (*field.BenchmarkCurveNameField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from ExecutionReport.
func (m ExecutionReport) GetBenchmarkCurveName(f *field.BenchmarkCurveNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkCurvePoint() (*field.BenchmarkCurvePointField, errors.MessageRejectError) {
	f := &field.BenchmarkCurvePointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from ExecutionReport.
func (m ExecutionReport) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldType() (*field.YieldTypeField, errors.MessageRejectError) {
	f := &field.YieldTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from ExecutionReport.
func (m ExecutionReport) GetYieldType(f *field.YieldTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for ExecutionReport.
func (m ExecutionReport) Yield() (*field.YieldField, errors.MessageRejectError) {
	f := &field.YieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from ExecutionReport.
func (m ExecutionReport) GetYield(f *field.YieldField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) GrossTradeAmt() (*field.GrossTradeAmtField, errors.MessageRejectError) {
	f := &field.GrossTradeAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from ExecutionReport.
func (m ExecutionReport) GetGrossTradeAmt(f *field.GrossTradeAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NumDaysInterest is a non-required field for ExecutionReport.
func (m ExecutionReport) NumDaysInterest() (*field.NumDaysInterestField, errors.MessageRejectError) {
	f := &field.NumDaysInterestField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNumDaysInterest reads a NumDaysInterest from ExecutionReport.
func (m ExecutionReport) GetNumDaysInterest(f *field.NumDaysInterestField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDate is a non-required field for ExecutionReport.
func (m ExecutionReport) ExDate() (*field.ExDateField, errors.MessageRejectError) {
	f := &field.ExDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDate reads a ExDate from ExecutionReport.
func (m ExecutionReport) GetExDate(f *field.ExDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestRate is a non-required field for ExecutionReport.
func (m ExecutionReport) AccruedInterestRate() (*field.AccruedInterestRateField, errors.MessageRejectError) {
	f := &field.AccruedInterestRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestRate reads a AccruedInterestRate from ExecutionReport.
func (m ExecutionReport) GetAccruedInterestRate(f *field.AccruedInterestRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) AccruedInterestAmt() (*field.AccruedInterestAmtField, errors.MessageRejectError) {
	f := &field.AccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from ExecutionReport.
func (m ExecutionReport) GetAccruedInterestAmt(f *field.AccruedInterestAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradedFlatSwitch is a non-required field for ExecutionReport.
func (m ExecutionReport) TradedFlatSwitch() (*field.TradedFlatSwitchField, errors.MessageRejectError) {
	f := &field.TradedFlatSwitchField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradedFlatSwitch reads a TradedFlatSwitch from ExecutionReport.
func (m ExecutionReport) GetTradedFlatSwitch(f *field.TradedFlatSwitchField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BasisFeatureDate is a non-required field for ExecutionReport.
func (m ExecutionReport) BasisFeatureDate() (*field.BasisFeatureDateField, errors.MessageRejectError) {
	f := &field.BasisFeatureDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBasisFeatureDate reads a BasisFeatureDate from ExecutionReport.
func (m ExecutionReport) GetBasisFeatureDate(f *field.BasisFeatureDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BasisFeaturePrice is a non-required field for ExecutionReport.
func (m ExecutionReport) BasisFeaturePrice() (*field.BasisFeaturePriceField, errors.MessageRejectError) {
	f := &field.BasisFeaturePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBasisFeaturePrice reads a BasisFeaturePrice from ExecutionReport.
func (m ExecutionReport) GetBasisFeaturePrice(f *field.BasisFeaturePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Concession is a non-required field for ExecutionReport.
func (m ExecutionReport) Concession() (*field.ConcessionField, errors.MessageRejectError) {
	f := &field.ConcessionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConcession reads a Concession from ExecutionReport.
func (m ExecutionReport) GetConcession(f *field.ConcessionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalTakedown is a non-required field for ExecutionReport.
func (m ExecutionReport) TotalTakedown() (*field.TotalTakedownField, errors.MessageRejectError) {
	f := &field.TotalTakedownField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalTakedown reads a TotalTakedown from ExecutionReport.
func (m ExecutionReport) GetTotalTakedown(f *field.TotalTakedownField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for ExecutionReport.
func (m ExecutionReport) NetMoney() (*field.NetMoneyField, errors.MessageRejectError) {
	f := &field.NetMoneyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from ExecutionReport.
func (m ExecutionReport) GetNetMoney(f *field.NetMoneyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrAmt() (*field.SettlCurrAmtField, errors.MessageRejectError) {
	f := &field.SettlCurrAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrAmt reads a SettlCurrAmt from ExecutionReport.
func (m ExecutionReport) GetSettlCurrAmt(f *field.SettlCurrAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrency() (*field.SettlCurrencyField, errors.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from ExecutionReport.
func (m ExecutionReport) GetSettlCurrency(f *field.SettlCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRate is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrFxRate() (*field.SettlCurrFxRateField, errors.MessageRejectError) {
	f := &field.SettlCurrFxRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRate reads a SettlCurrFxRate from ExecutionReport.
func (m ExecutionReport) GetSettlCurrFxRate(f *field.SettlCurrFxRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRateCalc is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalcField, errors.MessageRejectError) {
	f := &field.SettlCurrFxRateCalcField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRateCalc reads a SettlCurrFxRateCalc from ExecutionReport.
func (m ExecutionReport) GetSettlCurrFxRateCalc(f *field.SettlCurrFxRateCalcField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a non-required field for ExecutionReport.
func (m ExecutionReport) HandlInst() (*field.HandlInstField, errors.MessageRejectError) {
	f := &field.HandlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from ExecutionReport.
func (m ExecutionReport) GetHandlInst(f *field.HandlInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for ExecutionReport.
func (m ExecutionReport) MinQty() (*field.MinQtyField, errors.MessageRejectError) {
	f := &field.MinQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from ExecutionReport.
func (m ExecutionReport) GetMinQty(f *field.MinQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxFloor() (*field.MaxFloorField, errors.MessageRejectError) {
	f := &field.MaxFloorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from ExecutionReport.
func (m ExecutionReport) GetMaxFloor(f *field.MaxFloorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for ExecutionReport.
func (m ExecutionReport) PositionEffect() (*field.PositionEffectField, errors.MessageRejectError) {
	f := &field.PositionEffectField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from ExecutionReport.
func (m ExecutionReport) GetPositionEffect(f *field.PositionEffectField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxShow() (*field.MaxShowField, errors.MessageRejectError) {
	f := &field.MaxShowField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from ExecutionReport.
func (m ExecutionReport) GetMaxShow(f *field.MaxShowField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ExecutionReport.
func (m ExecutionReport) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ExecutionReport.
func (m ExecutionReport) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ExecutionReport.
func (m ExecutionReport) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ExecutionReport.
func (m ExecutionReport) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate2 is a non-required field for ExecutionReport.
func (m ExecutionReport) FutSettDate2() (*field.FutSettDate2Field, errors.MessageRejectError) {
	f := &field.FutSettDate2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate2 reads a FutSettDate2 from ExecutionReport.
func (m ExecutionReport) GetFutSettDate2(f *field.FutSettDate2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty2 is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderQty2() (*field.OrderQty2Field, errors.MessageRejectError) {
	f := &field.OrderQty2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty2 reads a OrderQty2 from ExecutionReport.
func (m ExecutionReport) GetOrderQty2(f *field.OrderQty2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastForwardPoints2 is a non-required field for ExecutionReport.
func (m ExecutionReport) LastForwardPoints2() (*field.LastForwardPoints2Field, errors.MessageRejectError) {
	f := &field.LastForwardPoints2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastForwardPoints2 reads a LastForwardPoints2 from ExecutionReport.
func (m ExecutionReport) GetLastForwardPoints2(f *field.LastForwardPoints2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegReportingType is a non-required field for ExecutionReport.
func (m ExecutionReport) MultiLegReportingType() (*field.MultiLegReportingTypeField, errors.MessageRejectError) {
	f := &field.MultiLegReportingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegReportingType reads a MultiLegReportingType from ExecutionReport.
func (m ExecutionReport) GetMultiLegReportingType(f *field.MultiLegReportingTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for ExecutionReport.
func (m ExecutionReport) CancellationRights() (*field.CancellationRightsField, errors.MessageRejectError) {
	f := &field.CancellationRightsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from ExecutionReport.
func (m ExecutionReport) GetCancellationRights(f *field.CancellationRightsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for ExecutionReport.
func (m ExecutionReport) MoneyLaunderingStatus() (*field.MoneyLaunderingStatusField, errors.MessageRejectError) {
	f := &field.MoneyLaunderingStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from ExecutionReport.
func (m ExecutionReport) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for ExecutionReport.
func (m ExecutionReport) RegistID() (*field.RegistIDField, errors.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from ExecutionReport.
func (m ExecutionReport) GetRegistID(f *field.RegistIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for ExecutionReport.
func (m ExecutionReport) Designation() (*field.DesignationField, errors.MessageRejectError) {
	f := &field.DesignationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from ExecutionReport.
func (m ExecutionReport) GetDesignation(f *field.DesignationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransBkdTime is a non-required field for ExecutionReport.
func (m ExecutionReport) TransBkdTime() (*field.TransBkdTimeField, errors.MessageRejectError) {
	f := &field.TransBkdTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransBkdTime reads a TransBkdTime from ExecutionReport.
func (m ExecutionReport) GetTransBkdTime(f *field.TransBkdTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecValuationPoint is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecValuationPoint() (*field.ExecValuationPointField, errors.MessageRejectError) {
	f := &field.ExecValuationPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecValuationPoint reads a ExecValuationPoint from ExecutionReport.
func (m ExecutionReport) GetExecValuationPoint(f *field.ExecValuationPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecPriceType() (*field.ExecPriceTypeField, errors.MessageRejectError) {
	f := &field.ExecPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecPriceType reads a ExecPriceType from ExecutionReport.
func (m ExecutionReport) GetExecPriceType(f *field.ExecPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecPriceAdjustment is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecPriceAdjustment() (*field.ExecPriceAdjustmentField, errors.MessageRejectError) {
	f := &field.ExecPriceAdjustmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecPriceAdjustment reads a ExecPriceAdjustment from ExecutionReport.
func (m ExecutionReport) GetExecPriceAdjustment(f *field.ExecPriceAdjustmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriorityIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) PriorityIndicator() (*field.PriorityIndicatorField, errors.MessageRejectError) {
	f := &field.PriorityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriorityIndicator reads a PriorityIndicator from ExecutionReport.
func (m ExecutionReport) GetPriorityIndicator(f *field.PriorityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceImprovement is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceImprovement() (*field.PriceImprovementField, errors.MessageRejectError) {
	f := &field.PriceImprovementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceImprovement reads a PriceImprovement from ExecutionReport.
func (m ExecutionReport) GetPriceImprovement(f *field.PriceImprovementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoContAmts is a non-required field for ExecutionReport.
func (m ExecutionReport) NoContAmts() (*field.NoContAmtsField, errors.MessageRejectError) {
	f := &field.NoContAmtsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoContAmts reads a NoContAmts from ExecutionReport.
func (m ExecutionReport) GetNoContAmts(f *field.NoContAmtsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for ExecutionReport.
func (m ExecutionReport) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from ExecutionReport.
func (m ExecutionReport) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
