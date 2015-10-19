//Package orderstatusrequest msg type = H.
package orderstatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a OrderStatusRequest wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//OrderID is a non-required field for OrderStatusRequest.
func (m Message) OrderID() (*field.OrderIDField, quickfix.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderStatusRequest.
func (m Message) GetOrderID(f *field.OrderIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for OrderStatusRequest.
func (m Message) ClOrdID() (*field.ClOrdIDField, quickfix.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderStatusRequest.
func (m Message) GetClOrdID(f *field.ClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for OrderStatusRequest.
func (m Message) SecondaryClOrdID() (*field.SecondaryClOrdIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from OrderStatusRequest.
func (m Message) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for OrderStatusRequest.
func (m Message) ClOrdLinkID() (*field.ClOrdLinkIDField, quickfix.MessageRejectError) {
	f := &field.ClOrdLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from OrderStatusRequest.
func (m Message) GetClOrdLinkID(f *field.ClOrdLinkIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for OrderStatusRequest.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from OrderStatusRequest.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrdStatusReqID is a non-required field for OrderStatusRequest.
func (m Message) OrdStatusReqID() (*field.OrdStatusReqIDField, quickfix.MessageRejectError) {
	f := &field.OrdStatusReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdStatusReqID reads a OrdStatusReqID from OrderStatusRequest.
func (m Message) GetOrdStatusReqID(f *field.OrdStatusReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for OrderStatusRequest.
func (m Message) Account() (*field.AccountField, quickfix.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from OrderStatusRequest.
func (m Message) GetAccount(f *field.AccountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for OrderStatusRequest.
func (m Message) AcctIDSource() (*field.AcctIDSourceField, quickfix.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from OrderStatusRequest.
func (m Message) GetAcctIDSource(f *field.AcctIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for OrderStatusRequest.
func (m Message) Symbol() (*field.SymbolField, quickfix.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderStatusRequest.
func (m Message) GetSymbol(f *field.SymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderStatusRequest.
func (m Message) SymbolSfx() (*field.SymbolSfxField, quickfix.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderStatusRequest.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderStatusRequest.
func (m Message) SecurityID() (*field.SecurityIDField, quickfix.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderStatusRequest.
func (m Message) GetSecurityID(f *field.SecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for OrderStatusRequest.
func (m Message) SecurityIDSource() (*field.SecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from OrderStatusRequest.
func (m Message) GetSecurityIDSource(f *field.SecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for OrderStatusRequest.
func (m Message) NoSecurityAltID() (*field.NoSecurityAltIDField, quickfix.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from OrderStatusRequest.
func (m Message) GetNoSecurityAltID(f *field.NoSecurityAltIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for OrderStatusRequest.
func (m Message) Product() (*field.ProductField, quickfix.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from OrderStatusRequest.
func (m Message) GetProduct(f *field.ProductField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for OrderStatusRequest.
func (m Message) CFICode() (*field.CFICodeField, quickfix.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from OrderStatusRequest.
func (m Message) GetCFICode(f *field.CFICodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for OrderStatusRequest.
func (m Message) SecurityType() (*field.SecurityTypeField, quickfix.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from OrderStatusRequest.
func (m Message) GetSecurityType(f *field.SecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for OrderStatusRequest.
func (m Message) SecuritySubType() (*field.SecuritySubTypeField, quickfix.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from OrderStatusRequest.
func (m Message) GetSecuritySubType(f *field.SecuritySubTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for OrderStatusRequest.
func (m Message) MaturityMonthYear() (*field.MaturityMonthYearField, quickfix.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from OrderStatusRequest.
func (m Message) GetMaturityMonthYear(f *field.MaturityMonthYearField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for OrderStatusRequest.
func (m Message) MaturityDate() (*field.MaturityDateField, quickfix.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from OrderStatusRequest.
func (m Message) GetMaturityDate(f *field.MaturityDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for OrderStatusRequest.
func (m Message) CouponPaymentDate() (*field.CouponPaymentDateField, quickfix.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from OrderStatusRequest.
func (m Message) GetCouponPaymentDate(f *field.CouponPaymentDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for OrderStatusRequest.
func (m Message) IssueDate() (*field.IssueDateField, quickfix.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from OrderStatusRequest.
func (m Message) GetIssueDate(f *field.IssueDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for OrderStatusRequest.
func (m Message) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, quickfix.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from OrderStatusRequest.
func (m Message) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for OrderStatusRequest.
func (m Message) RepurchaseTerm() (*field.RepurchaseTermField, quickfix.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from OrderStatusRequest.
func (m Message) GetRepurchaseTerm(f *field.RepurchaseTermField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for OrderStatusRequest.
func (m Message) RepurchaseRate() (*field.RepurchaseRateField, quickfix.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from OrderStatusRequest.
func (m Message) GetRepurchaseRate(f *field.RepurchaseRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for OrderStatusRequest.
func (m Message) Factor() (*field.FactorField, quickfix.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from OrderStatusRequest.
func (m Message) GetFactor(f *field.FactorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for OrderStatusRequest.
func (m Message) CreditRating() (*field.CreditRatingField, quickfix.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from OrderStatusRequest.
func (m Message) GetCreditRating(f *field.CreditRatingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for OrderStatusRequest.
func (m Message) InstrRegistry() (*field.InstrRegistryField, quickfix.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from OrderStatusRequest.
func (m Message) GetInstrRegistry(f *field.InstrRegistryField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for OrderStatusRequest.
func (m Message) CountryOfIssue() (*field.CountryOfIssueField, quickfix.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from OrderStatusRequest.
func (m Message) GetCountryOfIssue(f *field.CountryOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for OrderStatusRequest.
func (m Message) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, quickfix.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from OrderStatusRequest.
func (m Message) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for OrderStatusRequest.
func (m Message) LocaleOfIssue() (*field.LocaleOfIssueField, quickfix.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from OrderStatusRequest.
func (m Message) GetLocaleOfIssue(f *field.LocaleOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for OrderStatusRequest.
func (m Message) RedemptionDate() (*field.RedemptionDateField, quickfix.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from OrderStatusRequest.
func (m Message) GetRedemptionDate(f *field.RedemptionDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for OrderStatusRequest.
func (m Message) StrikePrice() (*field.StrikePriceField, quickfix.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from OrderStatusRequest.
func (m Message) GetStrikePrice(f *field.StrikePriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for OrderStatusRequest.
func (m Message) StrikeCurrency() (*field.StrikeCurrencyField, quickfix.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from OrderStatusRequest.
func (m Message) GetStrikeCurrency(f *field.StrikeCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for OrderStatusRequest.
func (m Message) OptAttribute() (*field.OptAttributeField, quickfix.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from OrderStatusRequest.
func (m Message) GetOptAttribute(f *field.OptAttributeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for OrderStatusRequest.
func (m Message) ContractMultiplier() (*field.ContractMultiplierField, quickfix.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from OrderStatusRequest.
func (m Message) GetContractMultiplier(f *field.ContractMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for OrderStatusRequest.
func (m Message) CouponRate() (*field.CouponRateField, quickfix.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from OrderStatusRequest.
func (m Message) GetCouponRate(f *field.CouponRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for OrderStatusRequest.
func (m Message) SecurityExchange() (*field.SecurityExchangeField, quickfix.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from OrderStatusRequest.
func (m Message) GetSecurityExchange(f *field.SecurityExchangeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderStatusRequest.
func (m Message) Issuer() (*field.IssuerField, quickfix.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderStatusRequest.
func (m Message) GetIssuer(f *field.IssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for OrderStatusRequest.
func (m Message) EncodedIssuerLen() (*field.EncodedIssuerLenField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from OrderStatusRequest.
func (m Message) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for OrderStatusRequest.
func (m Message) EncodedIssuer() (*field.EncodedIssuerField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from OrderStatusRequest.
func (m Message) GetEncodedIssuer(f *field.EncodedIssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderStatusRequest.
func (m Message) SecurityDesc() (*field.SecurityDescField, quickfix.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderStatusRequest.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for OrderStatusRequest.
func (m Message) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from OrderStatusRequest.
func (m Message) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for OrderStatusRequest.
func (m Message) EncodedSecurityDesc() (*field.EncodedSecurityDescField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from OrderStatusRequest.
func (m Message) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for OrderStatusRequest.
func (m Message) Pool() (*field.PoolField, quickfix.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from OrderStatusRequest.
func (m Message) GetPool(f *field.PoolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for OrderStatusRequest.
func (m Message) ContractSettlMonth() (*field.ContractSettlMonthField, quickfix.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from OrderStatusRequest.
func (m Message) GetContractSettlMonth(f *field.ContractSettlMonthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for OrderStatusRequest.
func (m Message) CPProgram() (*field.CPProgramField, quickfix.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from OrderStatusRequest.
func (m Message) GetCPProgram(f *field.CPProgramField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for OrderStatusRequest.
func (m Message) CPRegType() (*field.CPRegTypeField, quickfix.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from OrderStatusRequest.
func (m Message) GetCPRegType(f *field.CPRegTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for OrderStatusRequest.
func (m Message) NoEvents() (*field.NoEventsField, quickfix.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from OrderStatusRequest.
func (m Message) GetNoEvents(f *field.NoEventsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for OrderStatusRequest.
func (m Message) DatedDate() (*field.DatedDateField, quickfix.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from OrderStatusRequest.
func (m Message) GetDatedDate(f *field.DatedDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for OrderStatusRequest.
func (m Message) InterestAccrualDate() (*field.InterestAccrualDateField, quickfix.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from OrderStatusRequest.
func (m Message) GetInterestAccrualDate(f *field.InterestAccrualDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for OrderStatusRequest.
func (m Message) SecurityStatus() (*field.SecurityStatusField, quickfix.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from OrderStatusRequest.
func (m Message) GetSecurityStatus(f *field.SecurityStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for OrderStatusRequest.
func (m Message) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, quickfix.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from OrderStatusRequest.
func (m Message) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for OrderStatusRequest.
func (m Message) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, quickfix.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from OrderStatusRequest.
func (m Message) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for OrderStatusRequest.
func (m Message) StrikeMultiplier() (*field.StrikeMultiplierField, quickfix.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from OrderStatusRequest.
func (m Message) GetStrikeMultiplier(f *field.StrikeMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for OrderStatusRequest.
func (m Message) StrikeValue() (*field.StrikeValueField, quickfix.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from OrderStatusRequest.
func (m Message) GetStrikeValue(f *field.StrikeValueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for OrderStatusRequest.
func (m Message) MinPriceIncrement() (*field.MinPriceIncrementField, quickfix.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from OrderStatusRequest.
func (m Message) GetMinPriceIncrement(f *field.MinPriceIncrementField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for OrderStatusRequest.
func (m Message) PositionLimit() (*field.PositionLimitField, quickfix.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from OrderStatusRequest.
func (m Message) GetPositionLimit(f *field.PositionLimitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for OrderStatusRequest.
func (m Message) NTPositionLimit() (*field.NTPositionLimitField, quickfix.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from OrderStatusRequest.
func (m Message) GetNTPositionLimit(f *field.NTPositionLimitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for OrderStatusRequest.
func (m Message) NoInstrumentParties() (*field.NoInstrumentPartiesField, quickfix.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from OrderStatusRequest.
func (m Message) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for OrderStatusRequest.
func (m Message) UnitOfMeasure() (*field.UnitOfMeasureField, quickfix.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from OrderStatusRequest.
func (m Message) GetUnitOfMeasure(f *field.UnitOfMeasureField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for OrderStatusRequest.
func (m Message) TimeUnit() (*field.TimeUnitField, quickfix.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from OrderStatusRequest.
func (m Message) GetTimeUnit(f *field.TimeUnitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for OrderStatusRequest.
func (m Message) MaturityTime() (*field.MaturityTimeField, quickfix.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from OrderStatusRequest.
func (m Message) GetMaturityTime(f *field.MaturityTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for OrderStatusRequest.
func (m Message) SecurityGroup() (*field.SecurityGroupField, quickfix.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from OrderStatusRequest.
func (m Message) GetSecurityGroup(f *field.SecurityGroupField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for OrderStatusRequest.
func (m Message) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, quickfix.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from OrderStatusRequest.
func (m Message) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for OrderStatusRequest.
func (m Message) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, quickfix.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from OrderStatusRequest.
func (m Message) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for OrderStatusRequest.
func (m Message) SecurityXMLLen() (*field.SecurityXMLLenField, quickfix.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from OrderStatusRequest.
func (m Message) GetSecurityXMLLen(f *field.SecurityXMLLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for OrderStatusRequest.
func (m Message) SecurityXML() (*field.SecurityXMLField, quickfix.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from OrderStatusRequest.
func (m Message) GetSecurityXML(f *field.SecurityXMLField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for OrderStatusRequest.
func (m Message) SecurityXMLSchema() (*field.SecurityXMLSchemaField, quickfix.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from OrderStatusRequest.
func (m Message) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for OrderStatusRequest.
func (m Message) ProductComplex() (*field.ProductComplexField, quickfix.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from OrderStatusRequest.
func (m Message) GetProductComplex(f *field.ProductComplexField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for OrderStatusRequest.
func (m Message) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, quickfix.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from OrderStatusRequest.
func (m Message) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for OrderStatusRequest.
func (m Message) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, quickfix.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from OrderStatusRequest.
func (m Message) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for OrderStatusRequest.
func (m Message) SettlMethod() (*field.SettlMethodField, quickfix.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from OrderStatusRequest.
func (m Message) GetSettlMethod(f *field.SettlMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for OrderStatusRequest.
func (m Message) ExerciseStyle() (*field.ExerciseStyleField, quickfix.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from OrderStatusRequest.
func (m Message) GetExerciseStyle(f *field.ExerciseStyleField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for OrderStatusRequest.
func (m Message) OptPayoutAmount() (*field.OptPayoutAmountField, quickfix.MessageRejectError) {
	f := &field.OptPayoutAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from OrderStatusRequest.
func (m Message) GetOptPayoutAmount(f *field.OptPayoutAmountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for OrderStatusRequest.
func (m Message) PriceQuoteMethod() (*field.PriceQuoteMethodField, quickfix.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from OrderStatusRequest.
func (m Message) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for OrderStatusRequest.
func (m Message) ListMethod() (*field.ListMethodField, quickfix.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from OrderStatusRequest.
func (m Message) GetListMethod(f *field.ListMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for OrderStatusRequest.
func (m Message) CapPrice() (*field.CapPriceField, quickfix.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from OrderStatusRequest.
func (m Message) GetCapPrice(f *field.CapPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for OrderStatusRequest.
func (m Message) FloorPrice() (*field.FloorPriceField, quickfix.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from OrderStatusRequest.
func (m Message) GetFloorPrice(f *field.FloorPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for OrderStatusRequest.
func (m Message) PutOrCall() (*field.PutOrCallField, quickfix.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from OrderStatusRequest.
func (m Message) GetPutOrCall(f *field.PutOrCallField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for OrderStatusRequest.
func (m Message) FlexibleIndicator() (*field.FlexibleIndicatorField, quickfix.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from OrderStatusRequest.
func (m Message) GetFlexibleIndicator(f *field.FlexibleIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for OrderStatusRequest.
func (m Message) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, quickfix.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from OrderStatusRequest.
func (m Message) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for OrderStatusRequest.
func (m Message) ValuationMethod() (*field.ValuationMethodField, quickfix.MessageRejectError) {
	f := &field.ValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from OrderStatusRequest.
func (m Message) GetValuationMethod(f *field.ValuationMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for OrderStatusRequest.
func (m Message) ContractMultiplierUnit() (*field.ContractMultiplierUnitField, quickfix.MessageRejectError) {
	f := &field.ContractMultiplierUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from OrderStatusRequest.
func (m Message) GetContractMultiplierUnit(f *field.ContractMultiplierUnitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for OrderStatusRequest.
func (m Message) FlowScheduleType() (*field.FlowScheduleTypeField, quickfix.MessageRejectError) {
	f := &field.FlowScheduleTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from OrderStatusRequest.
func (m Message) GetFlowScheduleType(f *field.FlowScheduleTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for OrderStatusRequest.
func (m Message) RestructuringType() (*field.RestructuringTypeField, quickfix.MessageRejectError) {
	f := &field.RestructuringTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from OrderStatusRequest.
func (m Message) GetRestructuringType(f *field.RestructuringTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for OrderStatusRequest.
func (m Message) Seniority() (*field.SeniorityField, quickfix.MessageRejectError) {
	f := &field.SeniorityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from OrderStatusRequest.
func (m Message) GetSeniority(f *field.SeniorityField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for OrderStatusRequest.
func (m Message) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstandingField, quickfix.MessageRejectError) {
	f := &field.NotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from OrderStatusRequest.
func (m Message) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstandingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for OrderStatusRequest.
func (m Message) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstandingField, quickfix.MessageRejectError) {
	f := &field.OriginalNotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from OrderStatusRequest.
func (m Message) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstandingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for OrderStatusRequest.
func (m Message) AttachmentPoint() (*field.AttachmentPointField, quickfix.MessageRejectError) {
	f := &field.AttachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from OrderStatusRequest.
func (m Message) GetAttachmentPoint(f *field.AttachmentPointField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for OrderStatusRequest.
func (m Message) DetachmentPoint() (*field.DetachmentPointField, quickfix.MessageRejectError) {
	f := &field.DetachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from OrderStatusRequest.
func (m Message) GetDetachmentPoint(f *field.DetachmentPointField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for OrderStatusRequest.
func (m Message) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethodField, quickfix.MessageRejectError) {
	f := &field.StrikePriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from OrderStatusRequest.
func (m Message) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for OrderStatusRequest.
func (m Message) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethodField, quickfix.MessageRejectError) {
	f := &field.StrikePriceBoundaryMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from OrderStatusRequest.
func (m Message) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for OrderStatusRequest.
func (m Message) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecisionField, quickfix.MessageRejectError) {
	f := &field.StrikePriceBoundaryPrecisionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from OrderStatusRequest.
func (m Message) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecisionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for OrderStatusRequest.
func (m Message) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethodField, quickfix.MessageRejectError) {
	f := &field.UnderlyingPriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from OrderStatusRequest.
func (m Message) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for OrderStatusRequest.
func (m Message) OptPayoutType() (*field.OptPayoutTypeField, quickfix.MessageRejectError) {
	f := &field.OptPayoutTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from OrderStatusRequest.
func (m Message) GetOptPayoutType(f *field.OptPayoutTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for OrderStatusRequest.
func (m Message) NoComplexEvents() (*field.NoComplexEventsField, quickfix.MessageRejectError) {
	f := &field.NoComplexEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from OrderStatusRequest.
func (m Message) GetNoComplexEvents(f *field.NoComplexEventsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for OrderStatusRequest.
func (m Message) AgreementDesc() (*field.AgreementDescField, quickfix.MessageRejectError) {
	f := &field.AgreementDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from OrderStatusRequest.
func (m Message) GetAgreementDesc(f *field.AgreementDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for OrderStatusRequest.
func (m Message) AgreementID() (*field.AgreementIDField, quickfix.MessageRejectError) {
	f := &field.AgreementIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from OrderStatusRequest.
func (m Message) GetAgreementID(f *field.AgreementIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for OrderStatusRequest.
func (m Message) AgreementDate() (*field.AgreementDateField, quickfix.MessageRejectError) {
	f := &field.AgreementDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from OrderStatusRequest.
func (m Message) GetAgreementDate(f *field.AgreementDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for OrderStatusRequest.
func (m Message) AgreementCurrency() (*field.AgreementCurrencyField, quickfix.MessageRejectError) {
	f := &field.AgreementCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from OrderStatusRequest.
func (m Message) GetAgreementCurrency(f *field.AgreementCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for OrderStatusRequest.
func (m Message) TerminationType() (*field.TerminationTypeField, quickfix.MessageRejectError) {
	f := &field.TerminationTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from OrderStatusRequest.
func (m Message) GetTerminationType(f *field.TerminationTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for OrderStatusRequest.
func (m Message) StartDate() (*field.StartDateField, quickfix.MessageRejectError) {
	f := &field.StartDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from OrderStatusRequest.
func (m Message) GetStartDate(f *field.StartDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for OrderStatusRequest.
func (m Message) EndDate() (*field.EndDateField, quickfix.MessageRejectError) {
	f := &field.EndDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from OrderStatusRequest.
func (m Message) GetEndDate(f *field.EndDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for OrderStatusRequest.
func (m Message) DeliveryType() (*field.DeliveryTypeField, quickfix.MessageRejectError) {
	f := &field.DeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from OrderStatusRequest.
func (m Message) GetDeliveryType(f *field.DeliveryTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for OrderStatusRequest.
func (m Message) MarginRatio() (*field.MarginRatioField, quickfix.MessageRejectError) {
	f := &field.MarginRatioField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from OrderStatusRequest.
func (m Message) GetMarginRatio(f *field.MarginRatioField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for OrderStatusRequest.
func (m Message) NoUnderlyings() (*field.NoUnderlyingsField, quickfix.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from OrderStatusRequest.
func (m Message) GetNoUnderlyings(f *field.NoUnderlyingsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for OrderStatusRequest.
func (m Message) Side() (*field.SideField, quickfix.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderStatusRequest.
func (m Message) GetSide(f *field.SideField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds OrderStatusRequest messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for OrderStatusRequest.
func Builder(
	side *field.SideField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("H"))
	builder.Body.Set(side)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "H", r
}
