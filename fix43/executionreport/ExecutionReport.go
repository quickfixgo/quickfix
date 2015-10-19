//Package executionreport msg type = 8.
package executionreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a ExecutionReport wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//OrderID is a required field for ExecutionReport.
func (m Message) OrderID() (*field.OrderIDField, quickfix.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from ExecutionReport.
func (m Message) GetOrderID(f *field.OrderIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for ExecutionReport.
func (m Message) SecondaryOrderID() (*field.SecondaryOrderIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryOrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from ExecutionReport.
func (m Message) GetSecondaryOrderID(f *field.SecondaryOrderIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for ExecutionReport.
func (m Message) SecondaryClOrdID() (*field.SecondaryClOrdIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from ExecutionReport.
func (m Message) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryExecID is a non-required field for ExecutionReport.
func (m Message) SecondaryExecID() (*field.SecondaryExecIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryExecID reads a SecondaryExecID from ExecutionReport.
func (m Message) GetSecondaryExecID(f *field.SecondaryExecIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for ExecutionReport.
func (m Message) ClOrdID() (*field.ClOrdIDField, quickfix.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from ExecutionReport.
func (m Message) GetClOrdID(f *field.ClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a non-required field for ExecutionReport.
func (m Message) OrigClOrdID() (*field.OrigClOrdIDField, quickfix.MessageRejectError) {
	f := &field.OrigClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from ExecutionReport.
func (m Message) GetOrigClOrdID(f *field.OrigClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for ExecutionReport.
func (m Message) ClOrdLinkID() (*field.ClOrdLinkIDField, quickfix.MessageRejectError) {
	f := &field.ClOrdLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from ExecutionReport.
func (m Message) GetClOrdLinkID(f *field.ClOrdLinkIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for ExecutionReport.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from ExecutionReport.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for ExecutionReport.
func (m Message) TradeOriginationDate() (*field.TradeOriginationDateField, quickfix.MessageRejectError) {
	f := &field.TradeOriginationDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from ExecutionReport.
func (m Message) GetTradeOriginationDate(f *field.TradeOriginationDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoContraBrokers is a non-required field for ExecutionReport.
func (m Message) NoContraBrokers() (*field.NoContraBrokersField, quickfix.MessageRejectError) {
	f := &field.NoContraBrokersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoContraBrokers reads a NoContraBrokers from ExecutionReport.
func (m Message) GetNoContraBrokers(f *field.NoContraBrokersField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for ExecutionReport.
func (m Message) ListID() (*field.ListIDField, quickfix.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ExecutionReport.
func (m Message) GetListID(f *field.ListIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CrossID is a non-required field for ExecutionReport.
func (m Message) CrossID() (*field.CrossIDField, quickfix.MessageRejectError) {
	f := &field.CrossIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossID reads a CrossID from ExecutionReport.
func (m Message) GetCrossID(f *field.CrossIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrigCrossID is a non-required field for ExecutionReport.
func (m Message) OrigCrossID() (*field.OrigCrossIDField, quickfix.MessageRejectError) {
	f := &field.OrigCrossIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigCrossID reads a OrigCrossID from ExecutionReport.
func (m Message) GetOrigCrossID(f *field.OrigCrossIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CrossType is a non-required field for ExecutionReport.
func (m Message) CrossType() (*field.CrossTypeField, quickfix.MessageRejectError) {
	f := &field.CrossTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossType reads a CrossType from ExecutionReport.
func (m Message) GetCrossType(f *field.CrossTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecID is a required field for ExecutionReport.
func (m Message) ExecID() (*field.ExecIDField, quickfix.MessageRejectError) {
	f := &field.ExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecID reads a ExecID from ExecutionReport.
func (m Message) GetExecID(f *field.ExecIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecRefID is a non-required field for ExecutionReport.
func (m Message) ExecRefID() (*field.ExecRefIDField, quickfix.MessageRejectError) {
	f := &field.ExecRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecRefID reads a ExecRefID from ExecutionReport.
func (m Message) GetExecRefID(f *field.ExecRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecType is a required field for ExecutionReport.
func (m Message) ExecType() (*field.ExecTypeField, quickfix.MessageRejectError) {
	f := &field.ExecTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecType reads a ExecType from ExecutionReport.
func (m Message) GetExecType(f *field.ExecTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrdStatus is a required field for ExecutionReport.
func (m Message) OrdStatus() (*field.OrdStatusField, quickfix.MessageRejectError) {
	f := &field.OrdStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdStatus reads a OrdStatus from ExecutionReport.
func (m Message) GetOrdStatus(f *field.OrdStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//WorkingIndicator is a non-required field for ExecutionReport.
func (m Message) WorkingIndicator() (*field.WorkingIndicatorField, quickfix.MessageRejectError) {
	f := &field.WorkingIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetWorkingIndicator reads a WorkingIndicator from ExecutionReport.
func (m Message) GetWorkingIndicator(f *field.WorkingIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrdRejReason is a non-required field for ExecutionReport.
func (m Message) OrdRejReason() (*field.OrdRejReasonField, quickfix.MessageRejectError) {
	f := &field.OrdRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdRejReason reads a OrdRejReason from ExecutionReport.
func (m Message) GetOrdRejReason(f *field.OrdRejReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecRestatementReason is a non-required field for ExecutionReport.
func (m Message) ExecRestatementReason() (*field.ExecRestatementReasonField, quickfix.MessageRejectError) {
	f := &field.ExecRestatementReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecRestatementReason reads a ExecRestatementReason from ExecutionReport.
func (m Message) GetExecRestatementReason(f *field.ExecRestatementReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for ExecutionReport.
func (m Message) Account() (*field.AccountField, quickfix.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from ExecutionReport.
func (m Message) GetAccount(f *field.AccountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for ExecutionReport.
func (m Message) AccountType() (*field.AccountTypeField, quickfix.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from ExecutionReport.
func (m Message) GetAccountType(f *field.AccountTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DayBookingInst is a non-required field for ExecutionReport.
func (m Message) DayBookingInst() (*field.DayBookingInstField, quickfix.MessageRejectError) {
	f := &field.DayBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayBookingInst reads a DayBookingInst from ExecutionReport.
func (m Message) GetDayBookingInst(f *field.DayBookingInstField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BookingUnit is a non-required field for ExecutionReport.
func (m Message) BookingUnit() (*field.BookingUnitField, quickfix.MessageRejectError) {
	f := &field.BookingUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingUnit reads a BookingUnit from ExecutionReport.
func (m Message) GetBookingUnit(f *field.BookingUnitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PreallocMethod is a non-required field for ExecutionReport.
func (m Message) PreallocMethod() (*field.PreallocMethodField, quickfix.MessageRejectError) {
	f := &field.PreallocMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreallocMethod reads a PreallocMethod from ExecutionReport.
func (m Message) GetPreallocMethod(f *field.PreallocMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for ExecutionReport.
func (m Message) SettlmntTyp() (*field.SettlmntTypField, quickfix.MessageRejectError) {
	f := &field.SettlmntTypField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from ExecutionReport.
func (m Message) GetSettlmntTyp(f *field.SettlmntTypField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for ExecutionReport.
func (m Message) FutSettDate() (*field.FutSettDateField, quickfix.MessageRejectError) {
	f := &field.FutSettDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from ExecutionReport.
func (m Message) GetFutSettDate(f *field.FutSettDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CashMargin is a non-required field for ExecutionReport.
func (m Message) CashMargin() (*field.CashMarginField, quickfix.MessageRejectError) {
	f := &field.CashMarginField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashMargin reads a CashMargin from ExecutionReport.
func (m Message) GetCashMargin(f *field.CashMarginField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for ExecutionReport.
func (m Message) ClearingFeeIndicator() (*field.ClearingFeeIndicatorField, quickfix.MessageRejectError) {
	f := &field.ClearingFeeIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from ExecutionReport.
func (m Message) GetClearingFeeIndicator(f *field.ClearingFeeIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for ExecutionReport.
func (m Message) Symbol() (*field.SymbolField, quickfix.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from ExecutionReport.
func (m Message) GetSymbol(f *field.SymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for ExecutionReport.
func (m Message) SymbolSfx() (*field.SymbolSfxField, quickfix.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from ExecutionReport.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for ExecutionReport.
func (m Message) SecurityID() (*field.SecurityIDField, quickfix.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from ExecutionReport.
func (m Message) GetSecurityID(f *field.SecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for ExecutionReport.
func (m Message) SecurityIDSource() (*field.SecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from ExecutionReport.
func (m Message) GetSecurityIDSource(f *field.SecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for ExecutionReport.
func (m Message) NoSecurityAltID() (*field.NoSecurityAltIDField, quickfix.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from ExecutionReport.
func (m Message) GetNoSecurityAltID(f *field.NoSecurityAltIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for ExecutionReport.
func (m Message) Product() (*field.ProductField, quickfix.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from ExecutionReport.
func (m Message) GetProduct(f *field.ProductField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for ExecutionReport.
func (m Message) CFICode() (*field.CFICodeField, quickfix.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from ExecutionReport.
func (m Message) GetCFICode(f *field.CFICodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for ExecutionReport.
func (m Message) SecurityType() (*field.SecurityTypeField, quickfix.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from ExecutionReport.
func (m Message) GetSecurityType(f *field.SecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for ExecutionReport.
func (m Message) MaturityMonthYear() (*field.MaturityMonthYearField, quickfix.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from ExecutionReport.
func (m Message) GetMaturityMonthYear(f *field.MaturityMonthYearField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for ExecutionReport.
func (m Message) MaturityDate() (*field.MaturityDateField, quickfix.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from ExecutionReport.
func (m Message) GetMaturityDate(f *field.MaturityDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for ExecutionReport.
func (m Message) CouponPaymentDate() (*field.CouponPaymentDateField, quickfix.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from ExecutionReport.
func (m Message) GetCouponPaymentDate(f *field.CouponPaymentDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for ExecutionReport.
func (m Message) IssueDate() (*field.IssueDateField, quickfix.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from ExecutionReport.
func (m Message) GetIssueDate(f *field.IssueDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for ExecutionReport.
func (m Message) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, quickfix.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from ExecutionReport.
func (m Message) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for ExecutionReport.
func (m Message) RepurchaseTerm() (*field.RepurchaseTermField, quickfix.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from ExecutionReport.
func (m Message) GetRepurchaseTerm(f *field.RepurchaseTermField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for ExecutionReport.
func (m Message) RepurchaseRate() (*field.RepurchaseRateField, quickfix.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from ExecutionReport.
func (m Message) GetRepurchaseRate(f *field.RepurchaseRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for ExecutionReport.
func (m Message) Factor() (*field.FactorField, quickfix.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from ExecutionReport.
func (m Message) GetFactor(f *field.FactorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for ExecutionReport.
func (m Message) CreditRating() (*field.CreditRatingField, quickfix.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from ExecutionReport.
func (m Message) GetCreditRating(f *field.CreditRatingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for ExecutionReport.
func (m Message) InstrRegistry() (*field.InstrRegistryField, quickfix.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from ExecutionReport.
func (m Message) GetInstrRegistry(f *field.InstrRegistryField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for ExecutionReport.
func (m Message) CountryOfIssue() (*field.CountryOfIssueField, quickfix.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from ExecutionReport.
func (m Message) GetCountryOfIssue(f *field.CountryOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for ExecutionReport.
func (m Message) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, quickfix.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from ExecutionReport.
func (m Message) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for ExecutionReport.
func (m Message) LocaleOfIssue() (*field.LocaleOfIssueField, quickfix.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from ExecutionReport.
func (m Message) GetLocaleOfIssue(f *field.LocaleOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for ExecutionReport.
func (m Message) RedemptionDate() (*field.RedemptionDateField, quickfix.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from ExecutionReport.
func (m Message) GetRedemptionDate(f *field.RedemptionDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for ExecutionReport.
func (m Message) StrikePrice() (*field.StrikePriceField, quickfix.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from ExecutionReport.
func (m Message) GetStrikePrice(f *field.StrikePriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for ExecutionReport.
func (m Message) OptAttribute() (*field.OptAttributeField, quickfix.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from ExecutionReport.
func (m Message) GetOptAttribute(f *field.OptAttributeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for ExecutionReport.
func (m Message) ContractMultiplier() (*field.ContractMultiplierField, quickfix.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from ExecutionReport.
func (m Message) GetContractMultiplier(f *field.ContractMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for ExecutionReport.
func (m Message) CouponRate() (*field.CouponRateField, quickfix.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from ExecutionReport.
func (m Message) GetCouponRate(f *field.CouponRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for ExecutionReport.
func (m Message) SecurityExchange() (*field.SecurityExchangeField, quickfix.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from ExecutionReport.
func (m Message) GetSecurityExchange(f *field.SecurityExchangeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for ExecutionReport.
func (m Message) Issuer() (*field.IssuerField, quickfix.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from ExecutionReport.
func (m Message) GetIssuer(f *field.IssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for ExecutionReport.
func (m Message) EncodedIssuerLen() (*field.EncodedIssuerLenField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from ExecutionReport.
func (m Message) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for ExecutionReport.
func (m Message) EncodedIssuer() (*field.EncodedIssuerField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from ExecutionReport.
func (m Message) GetEncodedIssuer(f *field.EncodedIssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for ExecutionReport.
func (m Message) SecurityDesc() (*field.SecurityDescField, quickfix.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from ExecutionReport.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for ExecutionReport.
func (m Message) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from ExecutionReport.
func (m Message) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for ExecutionReport.
func (m Message) EncodedSecurityDesc() (*field.EncodedSecurityDescField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from ExecutionReport.
func (m Message) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for ExecutionReport.
func (m Message) Side() (*field.SideField, quickfix.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from ExecutionReport.
func (m Message) GetSide(f *field.SideField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for ExecutionReport.
func (m Message) NoStipulations() (*field.NoStipulationsField, quickfix.MessageRejectError) {
	f := &field.NoStipulationsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from ExecutionReport.
func (m Message) GetNoStipulations(f *field.NoStipulationsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QuantityType is a non-required field for ExecutionReport.
func (m Message) QuantityType() (*field.QuantityTypeField, quickfix.MessageRejectError) {
	f := &field.QuantityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuantityType reads a QuantityType from ExecutionReport.
func (m Message) GetQuantityType(f *field.QuantityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for ExecutionReport.
func (m Message) OrderQty() (*field.OrderQtyField, quickfix.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from ExecutionReport.
func (m Message) GetOrderQty(f *field.OrderQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for ExecutionReport.
func (m Message) CashOrderQty() (*field.CashOrderQtyField, quickfix.MessageRejectError) {
	f := &field.CashOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from ExecutionReport.
func (m Message) GetCashOrderQty(f *field.CashOrderQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for ExecutionReport.
func (m Message) OrderPercent() (*field.OrderPercentField, quickfix.MessageRejectError) {
	f := &field.OrderPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from ExecutionReport.
func (m Message) GetOrderPercent(f *field.OrderPercentField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for ExecutionReport.
func (m Message) RoundingDirection() (*field.RoundingDirectionField, quickfix.MessageRejectError) {
	f := &field.RoundingDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from ExecutionReport.
func (m Message) GetRoundingDirection(f *field.RoundingDirectionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for ExecutionReport.
func (m Message) RoundingModulus() (*field.RoundingModulusField, quickfix.MessageRejectError) {
	f := &field.RoundingModulusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from ExecutionReport.
func (m Message) GetRoundingModulus(f *field.RoundingModulusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a non-required field for ExecutionReport.
func (m Message) OrdType() (*field.OrdTypeField, quickfix.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from ExecutionReport.
func (m Message) GetOrdType(f *field.OrdTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for ExecutionReport.
func (m Message) PriceType() (*field.PriceTypeField, quickfix.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from ExecutionReport.
func (m Message) GetPriceType(f *field.PriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for ExecutionReport.
func (m Message) Price() (*field.PriceField, quickfix.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from ExecutionReport.
func (m Message) GetPrice(f *field.PriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for ExecutionReport.
func (m Message) StopPx() (*field.StopPxField, quickfix.MessageRejectError) {
	f := &field.StopPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from ExecutionReport.
func (m Message) GetStopPx(f *field.StopPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PegDifference is a non-required field for ExecutionReport.
func (m Message) PegDifference() (*field.PegDifferenceField, quickfix.MessageRejectError) {
	f := &field.PegDifferenceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegDifference reads a PegDifference from ExecutionReport.
func (m Message) GetPegDifference(f *field.PegDifferenceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for ExecutionReport.
func (m Message) DiscretionInst() (*field.DiscretionInstField, quickfix.MessageRejectError) {
	f := &field.DiscretionInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from ExecutionReport.
func (m Message) GetDiscretionInst(f *field.DiscretionInstField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffset is a non-required field for ExecutionReport.
func (m Message) DiscretionOffset() (*field.DiscretionOffsetField, quickfix.MessageRejectError) {
	f := &field.DiscretionOffsetField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffset reads a DiscretionOffset from ExecutionReport.
func (m Message) GetDiscretionOffset(f *field.DiscretionOffsetField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for ExecutionReport.
func (m Message) Currency() (*field.CurrencyField, quickfix.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from ExecutionReport.
func (m Message) GetCurrency(f *field.CurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for ExecutionReport.
func (m Message) ComplianceID() (*field.ComplianceIDField, quickfix.MessageRejectError) {
	f := &field.ComplianceIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from ExecutionReport.
func (m Message) GetComplianceID(f *field.ComplianceIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for ExecutionReport.
func (m Message) SolicitedFlag() (*field.SolicitedFlagField, quickfix.MessageRejectError) {
	f := &field.SolicitedFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from ExecutionReport.
func (m Message) GetSolicitedFlag(f *field.SolicitedFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for ExecutionReport.
func (m Message) TimeInForce() (*field.TimeInForceField, quickfix.MessageRejectError) {
	f := &field.TimeInForceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from ExecutionReport.
func (m Message) GetTimeInForce(f *field.TimeInForceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for ExecutionReport.
func (m Message) EffectiveTime() (*field.EffectiveTimeField, quickfix.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from ExecutionReport.
func (m Message) GetEffectiveTime(f *field.EffectiveTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for ExecutionReport.
func (m Message) ExpireDate() (*field.ExpireDateField, quickfix.MessageRejectError) {
	f := &field.ExpireDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from ExecutionReport.
func (m Message) GetExpireDate(f *field.ExpireDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for ExecutionReport.
func (m Message) ExpireTime() (*field.ExpireTimeField, quickfix.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from ExecutionReport.
func (m Message) GetExpireTime(f *field.ExpireTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for ExecutionReport.
func (m Message) ExecInst() (*field.ExecInstField, quickfix.MessageRejectError) {
	f := &field.ExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from ExecutionReport.
func (m Message) GetExecInst(f *field.ExecInstField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for ExecutionReport.
func (m Message) OrderCapacity() (*field.OrderCapacityField, quickfix.MessageRejectError) {
	f := &field.OrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from ExecutionReport.
func (m Message) GetOrderCapacity(f *field.OrderCapacityField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for ExecutionReport.
func (m Message) OrderRestrictions() (*field.OrderRestrictionsField, quickfix.MessageRejectError) {
	f := &field.OrderRestrictionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from ExecutionReport.
func (m Message) GetOrderRestrictions(f *field.OrderRestrictionsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for ExecutionReport.
func (m Message) CustOrderCapacity() (*field.CustOrderCapacityField, quickfix.MessageRejectError) {
	f := &field.CustOrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from ExecutionReport.
func (m Message) GetCustOrderCapacity(f *field.CustOrderCapacityField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Rule80A is a non-required field for ExecutionReport.
func (m Message) Rule80A() (*field.Rule80AField, quickfix.MessageRejectError) {
	f := &field.Rule80AField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRule80A reads a Rule80A from ExecutionReport.
func (m Message) GetRule80A(f *field.Rule80AField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastQty is a non-required field for ExecutionReport.
func (m Message) LastQty() (*field.LastQtyField, quickfix.MessageRejectError) {
	f := &field.LastQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastQty reads a LastQty from ExecutionReport.
func (m Message) GetLastQty(f *field.LastQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLastQty is a non-required field for ExecutionReport.
func (m Message) UnderlyingLastQty() (*field.UnderlyingLastQtyField, quickfix.MessageRejectError) {
	f := &field.UnderlyingLastQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLastQty reads a UnderlyingLastQty from ExecutionReport.
func (m Message) GetUnderlyingLastQty(f *field.UnderlyingLastQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a non-required field for ExecutionReport.
func (m Message) LastPx() (*field.LastPxField, quickfix.MessageRejectError) {
	f := &field.LastPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from ExecutionReport.
func (m Message) GetLastPx(f *field.LastPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLastPx is a non-required field for ExecutionReport.
func (m Message) UnderlyingLastPx() (*field.UnderlyingLastPxField, quickfix.MessageRejectError) {
	f := &field.UnderlyingLastPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLastPx reads a UnderlyingLastPx from ExecutionReport.
func (m Message) GetUnderlyingLastPx(f *field.UnderlyingLastPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastSpotRate is a non-required field for ExecutionReport.
func (m Message) LastSpotRate() (*field.LastSpotRateField, quickfix.MessageRejectError) {
	f := &field.LastSpotRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastSpotRate reads a LastSpotRate from ExecutionReport.
func (m Message) GetLastSpotRate(f *field.LastSpotRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastForwardPoints is a non-required field for ExecutionReport.
func (m Message) LastForwardPoints() (*field.LastForwardPointsField, quickfix.MessageRejectError) {
	f := &field.LastForwardPointsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastForwardPoints reads a LastForwardPoints from ExecutionReport.
func (m Message) GetLastForwardPoints(f *field.LastForwardPointsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for ExecutionReport.
func (m Message) LastMkt() (*field.LastMktField, quickfix.MessageRejectError) {
	f := &field.LastMktField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from ExecutionReport.
func (m Message) GetLastMkt(f *field.LastMktField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for ExecutionReport.
func (m Message) TradingSessionID() (*field.TradingSessionIDField, quickfix.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from ExecutionReport.
func (m Message) GetTradingSessionID(f *field.TradingSessionIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for ExecutionReport.
func (m Message) TradingSessionSubID() (*field.TradingSessionSubIDField, quickfix.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from ExecutionReport.
func (m Message) GetTradingSessionSubID(f *field.TradingSessionSubIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastCapacity is a non-required field for ExecutionReport.
func (m Message) LastCapacity() (*field.LastCapacityField, quickfix.MessageRejectError) {
	f := &field.LastCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastCapacity reads a LastCapacity from ExecutionReport.
func (m Message) GetLastCapacity(f *field.LastCapacityField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LeavesQty is a required field for ExecutionReport.
func (m Message) LeavesQty() (*field.LeavesQtyField, quickfix.MessageRejectError) {
	f := &field.LeavesQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLeavesQty reads a LeavesQty from ExecutionReport.
func (m Message) GetLeavesQty(f *field.LeavesQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CumQty is a required field for ExecutionReport.
func (m Message) CumQty() (*field.CumQtyField, quickfix.MessageRejectError) {
	f := &field.CumQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCumQty reads a CumQty from ExecutionReport.
func (m Message) GetCumQty(f *field.CumQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a required field for ExecutionReport.
func (m Message) AvgPx() (*field.AvgPxField, quickfix.MessageRejectError) {
	f := &field.AvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from ExecutionReport.
func (m Message) GetAvgPx(f *field.AvgPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DayOrderQty is a non-required field for ExecutionReport.
func (m Message) DayOrderQty() (*field.DayOrderQtyField, quickfix.MessageRejectError) {
	f := &field.DayOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayOrderQty reads a DayOrderQty from ExecutionReport.
func (m Message) GetDayOrderQty(f *field.DayOrderQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DayCumQty is a non-required field for ExecutionReport.
func (m Message) DayCumQty() (*field.DayCumQtyField, quickfix.MessageRejectError) {
	f := &field.DayCumQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayCumQty reads a DayCumQty from ExecutionReport.
func (m Message) GetDayCumQty(f *field.DayCumQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DayAvgPx is a non-required field for ExecutionReport.
func (m Message) DayAvgPx() (*field.DayAvgPxField, quickfix.MessageRejectError) {
	f := &field.DayAvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayAvgPx reads a DayAvgPx from ExecutionReport.
func (m Message) GetDayAvgPx(f *field.DayAvgPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for ExecutionReport.
func (m Message) GTBookingInst() (*field.GTBookingInstField, quickfix.MessageRejectError) {
	f := &field.GTBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from ExecutionReport.
func (m Message) GetGTBookingInst(f *field.GTBookingInstField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for ExecutionReport.
func (m Message) TradeDate() (*field.TradeDateField, quickfix.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from ExecutionReport.
func (m Message) GetTradeDate(f *field.TradeDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for ExecutionReport.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ExecutionReport.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ReportToExch is a non-required field for ExecutionReport.
func (m Message) ReportToExch() (*field.ReportToExchField, quickfix.MessageRejectError) {
	f := &field.ReportToExchField{}
	err := m.Body.Get(f)
	return f, err
}

//GetReportToExch reads a ReportToExch from ExecutionReport.
func (m Message) GetReportToExch(f *field.ReportToExchField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for ExecutionReport.
func (m Message) Commission() (*field.CommissionField, quickfix.MessageRejectError) {
	f := &field.CommissionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from ExecutionReport.
func (m Message) GetCommission(f *field.CommissionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for ExecutionReport.
func (m Message) CommType() (*field.CommTypeField, quickfix.MessageRejectError) {
	f := &field.CommTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from ExecutionReport.
func (m Message) GetCommType(f *field.CommTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CommCurrency is a non-required field for ExecutionReport.
func (m Message) CommCurrency() (*field.CommCurrencyField, quickfix.MessageRejectError) {
	f := &field.CommCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommCurrency reads a CommCurrency from ExecutionReport.
func (m Message) GetCommCurrency(f *field.CommCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FundRenewWaiv is a non-required field for ExecutionReport.
func (m Message) FundRenewWaiv() (*field.FundRenewWaivField, quickfix.MessageRejectError) {
	f := &field.FundRenewWaivField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFundRenewWaiv reads a FundRenewWaiv from ExecutionReport.
func (m Message) GetFundRenewWaiv(f *field.FundRenewWaivField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for ExecutionReport.
func (m Message) Spread() (*field.SpreadField, quickfix.MessageRejectError) {
	f := &field.SpreadField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from ExecutionReport.
func (m Message) GetSpread(f *field.SpreadField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for ExecutionReport.
func (m Message) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrencyField, quickfix.MessageRejectError) {
	f := &field.BenchmarkCurveCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from ExecutionReport.
func (m Message) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for ExecutionReport.
func (m Message) BenchmarkCurveName() (*field.BenchmarkCurveNameField, quickfix.MessageRejectError) {
	f := &field.BenchmarkCurveNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from ExecutionReport.
func (m Message) GetBenchmarkCurveName(f *field.BenchmarkCurveNameField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for ExecutionReport.
func (m Message) BenchmarkCurvePoint() (*field.BenchmarkCurvePointField, quickfix.MessageRejectError) {
	f := &field.BenchmarkCurvePointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from ExecutionReport.
func (m Message) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePointField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for ExecutionReport.
func (m Message) YieldType() (*field.YieldTypeField, quickfix.MessageRejectError) {
	f := &field.YieldTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from ExecutionReport.
func (m Message) GetYieldType(f *field.YieldTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for ExecutionReport.
func (m Message) Yield() (*field.YieldField, quickfix.MessageRejectError) {
	f := &field.YieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from ExecutionReport.
func (m Message) GetYield(f *field.YieldField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a non-required field for ExecutionReport.
func (m Message) GrossTradeAmt() (*field.GrossTradeAmtField, quickfix.MessageRejectError) {
	f := &field.GrossTradeAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from ExecutionReport.
func (m Message) GetGrossTradeAmt(f *field.GrossTradeAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NumDaysInterest is a non-required field for ExecutionReport.
func (m Message) NumDaysInterest() (*field.NumDaysInterestField, quickfix.MessageRejectError) {
	f := &field.NumDaysInterestField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNumDaysInterest reads a NumDaysInterest from ExecutionReport.
func (m Message) GetNumDaysInterest(f *field.NumDaysInterestField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExDate is a non-required field for ExecutionReport.
func (m Message) ExDate() (*field.ExDateField, quickfix.MessageRejectError) {
	f := &field.ExDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDate reads a ExDate from ExecutionReport.
func (m Message) GetExDate(f *field.ExDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestRate is a non-required field for ExecutionReport.
func (m Message) AccruedInterestRate() (*field.AccruedInterestRateField, quickfix.MessageRejectError) {
	f := &field.AccruedInterestRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestRate reads a AccruedInterestRate from ExecutionReport.
func (m Message) GetAccruedInterestRate(f *field.AccruedInterestRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for ExecutionReport.
func (m Message) AccruedInterestAmt() (*field.AccruedInterestAmtField, quickfix.MessageRejectError) {
	f := &field.AccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from ExecutionReport.
func (m Message) GetAccruedInterestAmt(f *field.AccruedInterestAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradedFlatSwitch is a non-required field for ExecutionReport.
func (m Message) TradedFlatSwitch() (*field.TradedFlatSwitchField, quickfix.MessageRejectError) {
	f := &field.TradedFlatSwitchField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradedFlatSwitch reads a TradedFlatSwitch from ExecutionReport.
func (m Message) GetTradedFlatSwitch(f *field.TradedFlatSwitchField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BasisFeatureDate is a non-required field for ExecutionReport.
func (m Message) BasisFeatureDate() (*field.BasisFeatureDateField, quickfix.MessageRejectError) {
	f := &field.BasisFeatureDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBasisFeatureDate reads a BasisFeatureDate from ExecutionReport.
func (m Message) GetBasisFeatureDate(f *field.BasisFeatureDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BasisFeaturePrice is a non-required field for ExecutionReport.
func (m Message) BasisFeaturePrice() (*field.BasisFeaturePriceField, quickfix.MessageRejectError) {
	f := &field.BasisFeaturePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBasisFeaturePrice reads a BasisFeaturePrice from ExecutionReport.
func (m Message) GetBasisFeaturePrice(f *field.BasisFeaturePriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Concession is a non-required field for ExecutionReport.
func (m Message) Concession() (*field.ConcessionField, quickfix.MessageRejectError) {
	f := &field.ConcessionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConcession reads a Concession from ExecutionReport.
func (m Message) GetConcession(f *field.ConcessionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TotalTakedown is a non-required field for ExecutionReport.
func (m Message) TotalTakedown() (*field.TotalTakedownField, quickfix.MessageRejectError) {
	f := &field.TotalTakedownField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalTakedown reads a TotalTakedown from ExecutionReport.
func (m Message) GetTotalTakedown(f *field.TotalTakedownField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for ExecutionReport.
func (m Message) NetMoney() (*field.NetMoneyField, quickfix.MessageRejectError) {
	f := &field.NetMoneyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from ExecutionReport.
func (m Message) GetNetMoney(f *field.NetMoneyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrAmt is a non-required field for ExecutionReport.
func (m Message) SettlCurrAmt() (*field.SettlCurrAmtField, quickfix.MessageRejectError) {
	f := &field.SettlCurrAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrAmt reads a SettlCurrAmt from ExecutionReport.
func (m Message) GetSettlCurrAmt(f *field.SettlCurrAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for ExecutionReport.
func (m Message) SettlCurrency() (*field.SettlCurrencyField, quickfix.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from ExecutionReport.
func (m Message) GetSettlCurrency(f *field.SettlCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRate is a non-required field for ExecutionReport.
func (m Message) SettlCurrFxRate() (*field.SettlCurrFxRateField, quickfix.MessageRejectError) {
	f := &field.SettlCurrFxRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRate reads a SettlCurrFxRate from ExecutionReport.
func (m Message) GetSettlCurrFxRate(f *field.SettlCurrFxRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRateCalc is a non-required field for ExecutionReport.
func (m Message) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalcField, quickfix.MessageRejectError) {
	f := &field.SettlCurrFxRateCalcField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRateCalc reads a SettlCurrFxRateCalc from ExecutionReport.
func (m Message) GetSettlCurrFxRateCalc(f *field.SettlCurrFxRateCalcField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a non-required field for ExecutionReport.
func (m Message) HandlInst() (*field.HandlInstField, quickfix.MessageRejectError) {
	f := &field.HandlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from ExecutionReport.
func (m Message) GetHandlInst(f *field.HandlInstField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for ExecutionReport.
func (m Message) MinQty() (*field.MinQtyField, quickfix.MessageRejectError) {
	f := &field.MinQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from ExecutionReport.
func (m Message) GetMinQty(f *field.MinQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for ExecutionReport.
func (m Message) MaxFloor() (*field.MaxFloorField, quickfix.MessageRejectError) {
	f := &field.MaxFloorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from ExecutionReport.
func (m Message) GetMaxFloor(f *field.MaxFloorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for ExecutionReport.
func (m Message) PositionEffect() (*field.PositionEffectField, quickfix.MessageRejectError) {
	f := &field.PositionEffectField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from ExecutionReport.
func (m Message) GetPositionEffect(f *field.PositionEffectField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for ExecutionReport.
func (m Message) MaxShow() (*field.MaxShowField, quickfix.MessageRejectError) {
	f := &field.MaxShowField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from ExecutionReport.
func (m Message) GetMaxShow(f *field.MaxShowField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ExecutionReport.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ExecutionReport.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ExecutionReport.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ExecutionReport.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ExecutionReport.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ExecutionReport.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate2 is a non-required field for ExecutionReport.
func (m Message) FutSettDate2() (*field.FutSettDate2Field, quickfix.MessageRejectError) {
	f := &field.FutSettDate2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate2 reads a FutSettDate2 from ExecutionReport.
func (m Message) GetFutSettDate2(f *field.FutSettDate2Field) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty2 is a non-required field for ExecutionReport.
func (m Message) OrderQty2() (*field.OrderQty2Field, quickfix.MessageRejectError) {
	f := &field.OrderQty2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty2 reads a OrderQty2 from ExecutionReport.
func (m Message) GetOrderQty2(f *field.OrderQty2Field) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastForwardPoints2 is a non-required field for ExecutionReport.
func (m Message) LastForwardPoints2() (*field.LastForwardPoints2Field, quickfix.MessageRejectError) {
	f := &field.LastForwardPoints2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastForwardPoints2 reads a LastForwardPoints2 from ExecutionReport.
func (m Message) GetLastForwardPoints2(f *field.LastForwardPoints2Field) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegReportingType is a non-required field for ExecutionReport.
func (m Message) MultiLegReportingType() (*field.MultiLegReportingTypeField, quickfix.MessageRejectError) {
	f := &field.MultiLegReportingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegReportingType reads a MultiLegReportingType from ExecutionReport.
func (m Message) GetMultiLegReportingType(f *field.MultiLegReportingTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for ExecutionReport.
func (m Message) CancellationRights() (*field.CancellationRightsField, quickfix.MessageRejectError) {
	f := &field.CancellationRightsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from ExecutionReport.
func (m Message) GetCancellationRights(f *field.CancellationRightsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for ExecutionReport.
func (m Message) MoneyLaunderingStatus() (*field.MoneyLaunderingStatusField, quickfix.MessageRejectError) {
	f := &field.MoneyLaunderingStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from ExecutionReport.
func (m Message) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for ExecutionReport.
func (m Message) RegistID() (*field.RegistIDField, quickfix.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from ExecutionReport.
func (m Message) GetRegistID(f *field.RegistIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for ExecutionReport.
func (m Message) Designation() (*field.DesignationField, quickfix.MessageRejectError) {
	f := &field.DesignationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from ExecutionReport.
func (m Message) GetDesignation(f *field.DesignationField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransBkdTime is a non-required field for ExecutionReport.
func (m Message) TransBkdTime() (*field.TransBkdTimeField, quickfix.MessageRejectError) {
	f := &field.TransBkdTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransBkdTime reads a TransBkdTime from ExecutionReport.
func (m Message) GetTransBkdTime(f *field.TransBkdTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecValuationPoint is a non-required field for ExecutionReport.
func (m Message) ExecValuationPoint() (*field.ExecValuationPointField, quickfix.MessageRejectError) {
	f := &field.ExecValuationPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecValuationPoint reads a ExecValuationPoint from ExecutionReport.
func (m Message) GetExecValuationPoint(f *field.ExecValuationPointField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecPriceType is a non-required field for ExecutionReport.
func (m Message) ExecPriceType() (*field.ExecPriceTypeField, quickfix.MessageRejectError) {
	f := &field.ExecPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecPriceType reads a ExecPriceType from ExecutionReport.
func (m Message) GetExecPriceType(f *field.ExecPriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecPriceAdjustment is a non-required field for ExecutionReport.
func (m Message) ExecPriceAdjustment() (*field.ExecPriceAdjustmentField, quickfix.MessageRejectError) {
	f := &field.ExecPriceAdjustmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecPriceAdjustment reads a ExecPriceAdjustment from ExecutionReport.
func (m Message) GetExecPriceAdjustment(f *field.ExecPriceAdjustmentField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriorityIndicator is a non-required field for ExecutionReport.
func (m Message) PriorityIndicator() (*field.PriorityIndicatorField, quickfix.MessageRejectError) {
	f := &field.PriorityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriorityIndicator reads a PriorityIndicator from ExecutionReport.
func (m Message) GetPriorityIndicator(f *field.PriorityIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceImprovement is a non-required field for ExecutionReport.
func (m Message) PriceImprovement() (*field.PriceImprovementField, quickfix.MessageRejectError) {
	f := &field.PriceImprovementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceImprovement reads a PriceImprovement from ExecutionReport.
func (m Message) GetPriceImprovement(f *field.PriceImprovementField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoContAmts is a non-required field for ExecutionReport.
func (m Message) NoContAmts() (*field.NoContAmtsField, quickfix.MessageRejectError) {
	f := &field.NoContAmtsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoContAmts reads a NoContAmts from ExecutionReport.
func (m Message) GetNoContAmts(f *field.NoContAmtsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for ExecutionReport.
func (m Message) NoLegs() (*field.NoLegsField, quickfix.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from ExecutionReport.
func (m Message) GetNoLegs(f *field.NoLegsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds ExecutionReport messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for ExecutionReport.
func Builder(
	orderid *field.OrderIDField,
	execid *field.ExecIDField,
	exectype *field.ExecTypeField,
	ordstatus *field.OrdStatusField,
	side *field.SideField,
	leavesqty *field.LeavesQtyField,
	cumqty *field.CumQtyField,
	avgpx *field.AvgPxField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
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

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX43, "8", r
}
