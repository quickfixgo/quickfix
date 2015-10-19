//Package assignmentreport msg type = AW.
package assignmentreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a AssignmentReport wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//AsgnRptID is a required field for AssignmentReport.
func (m Message) AsgnRptID() (*field.AsgnRptIDField, quickfix.MessageRejectError) {
	f := &field.AsgnRptIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAsgnRptID reads a AsgnRptID from AssignmentReport.
func (m Message) GetAsgnRptID(f *field.AsgnRptIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TotNumAssignmentReports is a non-required field for AssignmentReport.
func (m Message) TotNumAssignmentReports() (*field.TotNumAssignmentReportsField, quickfix.MessageRejectError) {
	f := &field.TotNumAssignmentReportsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNumAssignmentReports reads a TotNumAssignmentReports from AssignmentReport.
func (m Message) GetTotNumAssignmentReports(f *field.TotNumAssignmentReportsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastRptRequested is a non-required field for AssignmentReport.
func (m Message) LastRptRequested() (*field.LastRptRequestedField, quickfix.MessageRejectError) {
	f := &field.LastRptRequestedField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastRptRequested reads a LastRptRequested from AssignmentReport.
func (m Message) GetLastRptRequested(f *field.LastRptRequestedField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for AssignmentReport.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AssignmentReport.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for AssignmentReport.
func (m Message) Account() (*field.AccountField, quickfix.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from AssignmentReport.
func (m Message) GetAccount(f *field.AccountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a required field for AssignmentReport.
func (m Message) AccountType() (*field.AccountTypeField, quickfix.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from AssignmentReport.
func (m Message) GetAccountType(f *field.AccountTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for AssignmentReport.
func (m Message) Symbol() (*field.SymbolField, quickfix.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from AssignmentReport.
func (m Message) GetSymbol(f *field.SymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for AssignmentReport.
func (m Message) SymbolSfx() (*field.SymbolSfxField, quickfix.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from AssignmentReport.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for AssignmentReport.
func (m Message) SecurityID() (*field.SecurityIDField, quickfix.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from AssignmentReport.
func (m Message) GetSecurityID(f *field.SecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for AssignmentReport.
func (m Message) SecurityIDSource() (*field.SecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from AssignmentReport.
func (m Message) GetSecurityIDSource(f *field.SecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for AssignmentReport.
func (m Message) NoSecurityAltID() (*field.NoSecurityAltIDField, quickfix.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from AssignmentReport.
func (m Message) GetNoSecurityAltID(f *field.NoSecurityAltIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for AssignmentReport.
func (m Message) Product() (*field.ProductField, quickfix.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from AssignmentReport.
func (m Message) GetProduct(f *field.ProductField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for AssignmentReport.
func (m Message) CFICode() (*field.CFICodeField, quickfix.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from AssignmentReport.
func (m Message) GetCFICode(f *field.CFICodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for AssignmentReport.
func (m Message) SecurityType() (*field.SecurityTypeField, quickfix.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from AssignmentReport.
func (m Message) GetSecurityType(f *field.SecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for AssignmentReport.
func (m Message) SecuritySubType() (*field.SecuritySubTypeField, quickfix.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from AssignmentReport.
func (m Message) GetSecuritySubType(f *field.SecuritySubTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for AssignmentReport.
func (m Message) MaturityMonthYear() (*field.MaturityMonthYearField, quickfix.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from AssignmentReport.
func (m Message) GetMaturityMonthYear(f *field.MaturityMonthYearField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for AssignmentReport.
func (m Message) MaturityDate() (*field.MaturityDateField, quickfix.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from AssignmentReport.
func (m Message) GetMaturityDate(f *field.MaturityDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for AssignmentReport.
func (m Message) CouponPaymentDate() (*field.CouponPaymentDateField, quickfix.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from AssignmentReport.
func (m Message) GetCouponPaymentDate(f *field.CouponPaymentDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for AssignmentReport.
func (m Message) IssueDate() (*field.IssueDateField, quickfix.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from AssignmentReport.
func (m Message) GetIssueDate(f *field.IssueDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for AssignmentReport.
func (m Message) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, quickfix.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from AssignmentReport.
func (m Message) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for AssignmentReport.
func (m Message) RepurchaseTerm() (*field.RepurchaseTermField, quickfix.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from AssignmentReport.
func (m Message) GetRepurchaseTerm(f *field.RepurchaseTermField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for AssignmentReport.
func (m Message) RepurchaseRate() (*field.RepurchaseRateField, quickfix.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from AssignmentReport.
func (m Message) GetRepurchaseRate(f *field.RepurchaseRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for AssignmentReport.
func (m Message) Factor() (*field.FactorField, quickfix.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from AssignmentReport.
func (m Message) GetFactor(f *field.FactorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for AssignmentReport.
func (m Message) CreditRating() (*field.CreditRatingField, quickfix.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from AssignmentReport.
func (m Message) GetCreditRating(f *field.CreditRatingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for AssignmentReport.
func (m Message) InstrRegistry() (*field.InstrRegistryField, quickfix.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from AssignmentReport.
func (m Message) GetInstrRegistry(f *field.InstrRegistryField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for AssignmentReport.
func (m Message) CountryOfIssue() (*field.CountryOfIssueField, quickfix.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from AssignmentReport.
func (m Message) GetCountryOfIssue(f *field.CountryOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for AssignmentReport.
func (m Message) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, quickfix.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from AssignmentReport.
func (m Message) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for AssignmentReport.
func (m Message) LocaleOfIssue() (*field.LocaleOfIssueField, quickfix.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from AssignmentReport.
func (m Message) GetLocaleOfIssue(f *field.LocaleOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for AssignmentReport.
func (m Message) RedemptionDate() (*field.RedemptionDateField, quickfix.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from AssignmentReport.
func (m Message) GetRedemptionDate(f *field.RedemptionDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for AssignmentReport.
func (m Message) StrikePrice() (*field.StrikePriceField, quickfix.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from AssignmentReport.
func (m Message) GetStrikePrice(f *field.StrikePriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for AssignmentReport.
func (m Message) StrikeCurrency() (*field.StrikeCurrencyField, quickfix.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from AssignmentReport.
func (m Message) GetStrikeCurrency(f *field.StrikeCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for AssignmentReport.
func (m Message) OptAttribute() (*field.OptAttributeField, quickfix.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from AssignmentReport.
func (m Message) GetOptAttribute(f *field.OptAttributeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for AssignmentReport.
func (m Message) ContractMultiplier() (*field.ContractMultiplierField, quickfix.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from AssignmentReport.
func (m Message) GetContractMultiplier(f *field.ContractMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for AssignmentReport.
func (m Message) CouponRate() (*field.CouponRateField, quickfix.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from AssignmentReport.
func (m Message) GetCouponRate(f *field.CouponRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for AssignmentReport.
func (m Message) SecurityExchange() (*field.SecurityExchangeField, quickfix.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from AssignmentReport.
func (m Message) GetSecurityExchange(f *field.SecurityExchangeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for AssignmentReport.
func (m Message) Issuer() (*field.IssuerField, quickfix.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from AssignmentReport.
func (m Message) GetIssuer(f *field.IssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for AssignmentReport.
func (m Message) EncodedIssuerLen() (*field.EncodedIssuerLenField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from AssignmentReport.
func (m Message) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for AssignmentReport.
func (m Message) EncodedIssuer() (*field.EncodedIssuerField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from AssignmentReport.
func (m Message) GetEncodedIssuer(f *field.EncodedIssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for AssignmentReport.
func (m Message) SecurityDesc() (*field.SecurityDescField, quickfix.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from AssignmentReport.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for AssignmentReport.
func (m Message) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from AssignmentReport.
func (m Message) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for AssignmentReport.
func (m Message) EncodedSecurityDesc() (*field.EncodedSecurityDescField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from AssignmentReport.
func (m Message) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for AssignmentReport.
func (m Message) Pool() (*field.PoolField, quickfix.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from AssignmentReport.
func (m Message) GetPool(f *field.PoolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for AssignmentReport.
func (m Message) ContractSettlMonth() (*field.ContractSettlMonthField, quickfix.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from AssignmentReport.
func (m Message) GetContractSettlMonth(f *field.ContractSettlMonthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for AssignmentReport.
func (m Message) CPProgram() (*field.CPProgramField, quickfix.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from AssignmentReport.
func (m Message) GetCPProgram(f *field.CPProgramField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for AssignmentReport.
func (m Message) CPRegType() (*field.CPRegTypeField, quickfix.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from AssignmentReport.
func (m Message) GetCPRegType(f *field.CPRegTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for AssignmentReport.
func (m Message) NoEvents() (*field.NoEventsField, quickfix.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from AssignmentReport.
func (m Message) GetNoEvents(f *field.NoEventsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for AssignmentReport.
func (m Message) DatedDate() (*field.DatedDateField, quickfix.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from AssignmentReport.
func (m Message) GetDatedDate(f *field.DatedDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for AssignmentReport.
func (m Message) InterestAccrualDate() (*field.InterestAccrualDateField, quickfix.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from AssignmentReport.
func (m Message) GetInterestAccrualDate(f *field.InterestAccrualDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for AssignmentReport.
func (m Message) Currency() (*field.CurrencyField, quickfix.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from AssignmentReport.
func (m Message) GetCurrency(f *field.CurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for AssignmentReport.
func (m Message) NoLegs() (*field.NoLegsField, quickfix.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from AssignmentReport.
func (m Message) GetNoLegs(f *field.NoLegsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for AssignmentReport.
func (m Message) NoUnderlyings() (*field.NoUnderlyingsField, quickfix.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from AssignmentReport.
func (m Message) GetNoUnderlyings(f *field.NoUnderlyingsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPositions is a non-required field for AssignmentReport.
func (m Message) NoPositions() (*field.NoPositionsField, quickfix.MessageRejectError) {
	f := &field.NoPositionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPositions reads a NoPositions from AssignmentReport.
func (m Message) GetNoPositions(f *field.NoPositionsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPosAmt is a non-required field for AssignmentReport.
func (m Message) NoPosAmt() (*field.NoPosAmtField, quickfix.MessageRejectError) {
	f := &field.NoPosAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPosAmt reads a NoPosAmt from AssignmentReport.
func (m Message) GetNoPosAmt(f *field.NoPosAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ThresholdAmount is a non-required field for AssignmentReport.
func (m Message) ThresholdAmount() (*field.ThresholdAmountField, quickfix.MessageRejectError) {
	f := &field.ThresholdAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetThresholdAmount reads a ThresholdAmount from AssignmentReport.
func (m Message) GetThresholdAmount(f *field.ThresholdAmountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlPrice is a required field for AssignmentReport.
func (m Message) SettlPrice() (*field.SettlPriceField, quickfix.MessageRejectError) {
	f := &field.SettlPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlPrice reads a SettlPrice from AssignmentReport.
func (m Message) GetSettlPrice(f *field.SettlPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlPriceType is a required field for AssignmentReport.
func (m Message) SettlPriceType() (*field.SettlPriceTypeField, quickfix.MessageRejectError) {
	f := &field.SettlPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlPriceType reads a SettlPriceType from AssignmentReport.
func (m Message) GetSettlPriceType(f *field.SettlPriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlPrice is a required field for AssignmentReport.
func (m Message) UnderlyingSettlPrice() (*field.UnderlyingSettlPriceField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSettlPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlPrice reads a UnderlyingSettlPrice from AssignmentReport.
func (m Message) GetUnderlyingSettlPrice(f *field.UnderlyingSettlPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for AssignmentReport.
func (m Message) ExpireDate() (*field.ExpireDateField, quickfix.MessageRejectError) {
	f := &field.ExpireDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from AssignmentReport.
func (m Message) GetExpireDate(f *field.ExpireDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AssignmentMethod is a required field for AssignmentReport.
func (m Message) AssignmentMethod() (*field.AssignmentMethodField, quickfix.MessageRejectError) {
	f := &field.AssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAssignmentMethod reads a AssignmentMethod from AssignmentReport.
func (m Message) GetAssignmentMethod(f *field.AssignmentMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AssignmentUnit is a non-required field for AssignmentReport.
func (m Message) AssignmentUnit() (*field.AssignmentUnitField, quickfix.MessageRejectError) {
	f := &field.AssignmentUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAssignmentUnit reads a AssignmentUnit from AssignmentReport.
func (m Message) GetAssignmentUnit(f *field.AssignmentUnitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OpenInterest is a required field for AssignmentReport.
func (m Message) OpenInterest() (*field.OpenInterestField, quickfix.MessageRejectError) {
	f := &field.OpenInterestField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOpenInterest reads a OpenInterest from AssignmentReport.
func (m Message) GetOpenInterest(f *field.OpenInterestField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseMethod is a required field for AssignmentReport.
func (m Message) ExerciseMethod() (*field.ExerciseMethodField, quickfix.MessageRejectError) {
	f := &field.ExerciseMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseMethod reads a ExerciseMethod from AssignmentReport.
func (m Message) GetExerciseMethod(f *field.ExerciseMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a required field for AssignmentReport.
func (m Message) SettlSessID() (*field.SettlSessIDField, quickfix.MessageRejectError) {
	f := &field.SettlSessIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from AssignmentReport.
func (m Message) GetSettlSessID(f *field.SettlSessIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a required field for AssignmentReport.
func (m Message) SettlSessSubID() (*field.SettlSessSubIDField, quickfix.MessageRejectError) {
	f := &field.SettlSessSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from AssignmentReport.
func (m Message) GetSettlSessSubID(f *field.SettlSessSubIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a required field for AssignmentReport.
func (m Message) ClearingBusinessDate() (*field.ClearingBusinessDateField, quickfix.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from AssignmentReport.
func (m Message) GetClearingBusinessDate(f *field.ClearingBusinessDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for AssignmentReport.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from AssignmentReport.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for AssignmentReport.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from AssignmentReport.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for AssignmentReport.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from AssignmentReport.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds AssignmentReport messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for AssignmentReport.
func Builder(
	asgnrptid *field.AsgnRptIDField,
	accounttype *field.AccountTypeField,
	settlprice *field.SettlPriceField,
	settlpricetype *field.SettlPriceTypeField,
	underlyingsettlprice *field.UnderlyingSettlPriceField,
	assignmentmethod *field.AssignmentMethodField,
	openinterest *field.OpenInterestField,
	exercisemethod *field.ExerciseMethodField,
	settlsessid *field.SettlSessIDField,
	settlsesssubid *field.SettlSessSubIDField,
	clearingbusinessdate *field.ClearingBusinessDateField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header.Set(field.NewMsgType("AW"))
	builder.Body.Set(asgnrptid)
	builder.Body.Set(accounttype)
	builder.Body.Set(settlprice)
	builder.Body.Set(settlpricetype)
	builder.Body.Set(underlyingsettlprice)
	builder.Body.Set(assignmentmethod)
	builder.Body.Set(openinterest)
	builder.Body.Set(exercisemethod)
	builder.Body.Set(settlsessid)
	builder.Body.Set(settlsesssubid)
	builder.Body.Set(clearingbusinessdate)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX44, "AW", r
}
