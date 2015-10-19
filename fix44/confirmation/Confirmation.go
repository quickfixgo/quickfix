//Package confirmation msg type = AK.
package confirmation

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a Confirmation wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//ConfirmID is a required field for Confirmation.
func (m Message) ConfirmID() (*field.ConfirmIDField, quickfix.MessageRejectError) {
	f := &field.ConfirmIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmID reads a ConfirmID from Confirmation.
func (m Message) GetConfirmID(f *field.ConfirmIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ConfirmRefID is a non-required field for Confirmation.
func (m Message) ConfirmRefID() (*field.ConfirmRefIDField, quickfix.MessageRejectError) {
	f := &field.ConfirmRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmRefID reads a ConfirmRefID from Confirmation.
func (m Message) GetConfirmRefID(f *field.ConfirmRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ConfirmReqID is a non-required field for Confirmation.
func (m Message) ConfirmReqID() (*field.ConfirmReqIDField, quickfix.MessageRejectError) {
	f := &field.ConfirmReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmReqID reads a ConfirmReqID from Confirmation.
func (m Message) GetConfirmReqID(f *field.ConfirmReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ConfirmTransType is a required field for Confirmation.
func (m Message) ConfirmTransType() (*field.ConfirmTransTypeField, quickfix.MessageRejectError) {
	f := &field.ConfirmTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmTransType reads a ConfirmTransType from Confirmation.
func (m Message) GetConfirmTransType(f *field.ConfirmTransTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ConfirmType is a required field for Confirmation.
func (m Message) ConfirmType() (*field.ConfirmTypeField, quickfix.MessageRejectError) {
	f := &field.ConfirmTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmType reads a ConfirmType from Confirmation.
func (m Message) GetConfirmType(f *field.ConfirmTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CopyMsgIndicator is a non-required field for Confirmation.
func (m Message) CopyMsgIndicator() (*field.CopyMsgIndicatorField, quickfix.MessageRejectError) {
	f := &field.CopyMsgIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCopyMsgIndicator reads a CopyMsgIndicator from Confirmation.
func (m Message) GetCopyMsgIndicator(f *field.CopyMsgIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LegalConfirm is a non-required field for Confirmation.
func (m Message) LegalConfirm() (*field.LegalConfirmField, quickfix.MessageRejectError) {
	f := &field.LegalConfirmField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLegalConfirm reads a LegalConfirm from Confirmation.
func (m Message) GetLegalConfirm(f *field.LegalConfirmField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ConfirmStatus is a required field for Confirmation.
func (m Message) ConfirmStatus() (*field.ConfirmStatusField, quickfix.MessageRejectError) {
	f := &field.ConfirmStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmStatus reads a ConfirmStatus from Confirmation.
func (m Message) GetConfirmStatus(f *field.ConfirmStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for Confirmation.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from Confirmation.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a non-required field for Confirmation.
func (m Message) NoOrders() (*field.NoOrdersField, quickfix.MessageRejectError) {
	f := &field.NoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from Confirmation.
func (m Message) GetNoOrders(f *field.NoOrdersField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for Confirmation.
func (m Message) AllocID() (*field.AllocIDField, quickfix.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from Confirmation.
func (m Message) GetAllocID(f *field.AllocIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryAllocID is a non-required field for Confirmation.
func (m Message) SecondaryAllocID() (*field.SecondaryAllocIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryAllocID reads a SecondaryAllocID from Confirmation.
func (m Message) GetSecondaryAllocID(f *field.SecondaryAllocIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IndividualAllocID is a non-required field for Confirmation.
func (m Message) IndividualAllocID() (*field.IndividualAllocIDField, quickfix.MessageRejectError) {
	f := &field.IndividualAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIndividualAllocID reads a IndividualAllocID from Confirmation.
func (m Message) GetIndividualAllocID(f *field.IndividualAllocIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for Confirmation.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from Confirmation.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for Confirmation.
func (m Message) TradeDate() (*field.TradeDateField, quickfix.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from Confirmation.
func (m Message) GetTradeDate(f *field.TradeDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRegTimestamps is a non-required field for Confirmation.
func (m Message) NoTrdRegTimestamps() (*field.NoTrdRegTimestampsField, quickfix.MessageRejectError) {
	f := &field.NoTrdRegTimestampsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRegTimestamps reads a NoTrdRegTimestamps from Confirmation.
func (m Message) GetNoTrdRegTimestamps(f *field.NoTrdRegTimestampsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for Confirmation.
func (m Message) Symbol() (*field.SymbolField, quickfix.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from Confirmation.
func (m Message) GetSymbol(f *field.SymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for Confirmation.
func (m Message) SymbolSfx() (*field.SymbolSfxField, quickfix.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from Confirmation.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for Confirmation.
func (m Message) SecurityID() (*field.SecurityIDField, quickfix.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from Confirmation.
func (m Message) GetSecurityID(f *field.SecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for Confirmation.
func (m Message) SecurityIDSource() (*field.SecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from Confirmation.
func (m Message) GetSecurityIDSource(f *field.SecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for Confirmation.
func (m Message) NoSecurityAltID() (*field.NoSecurityAltIDField, quickfix.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from Confirmation.
func (m Message) GetNoSecurityAltID(f *field.NoSecurityAltIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for Confirmation.
func (m Message) Product() (*field.ProductField, quickfix.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from Confirmation.
func (m Message) GetProduct(f *field.ProductField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for Confirmation.
func (m Message) CFICode() (*field.CFICodeField, quickfix.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from Confirmation.
func (m Message) GetCFICode(f *field.CFICodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for Confirmation.
func (m Message) SecurityType() (*field.SecurityTypeField, quickfix.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from Confirmation.
func (m Message) GetSecurityType(f *field.SecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for Confirmation.
func (m Message) SecuritySubType() (*field.SecuritySubTypeField, quickfix.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from Confirmation.
func (m Message) GetSecuritySubType(f *field.SecuritySubTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for Confirmation.
func (m Message) MaturityMonthYear() (*field.MaturityMonthYearField, quickfix.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from Confirmation.
func (m Message) GetMaturityMonthYear(f *field.MaturityMonthYearField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for Confirmation.
func (m Message) MaturityDate() (*field.MaturityDateField, quickfix.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from Confirmation.
func (m Message) GetMaturityDate(f *field.MaturityDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for Confirmation.
func (m Message) CouponPaymentDate() (*field.CouponPaymentDateField, quickfix.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from Confirmation.
func (m Message) GetCouponPaymentDate(f *field.CouponPaymentDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for Confirmation.
func (m Message) IssueDate() (*field.IssueDateField, quickfix.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from Confirmation.
func (m Message) GetIssueDate(f *field.IssueDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for Confirmation.
func (m Message) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, quickfix.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from Confirmation.
func (m Message) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for Confirmation.
func (m Message) RepurchaseTerm() (*field.RepurchaseTermField, quickfix.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from Confirmation.
func (m Message) GetRepurchaseTerm(f *field.RepurchaseTermField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for Confirmation.
func (m Message) RepurchaseRate() (*field.RepurchaseRateField, quickfix.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from Confirmation.
func (m Message) GetRepurchaseRate(f *field.RepurchaseRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for Confirmation.
func (m Message) Factor() (*field.FactorField, quickfix.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from Confirmation.
func (m Message) GetFactor(f *field.FactorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for Confirmation.
func (m Message) CreditRating() (*field.CreditRatingField, quickfix.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from Confirmation.
func (m Message) GetCreditRating(f *field.CreditRatingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for Confirmation.
func (m Message) InstrRegistry() (*field.InstrRegistryField, quickfix.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from Confirmation.
func (m Message) GetInstrRegistry(f *field.InstrRegistryField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for Confirmation.
func (m Message) CountryOfIssue() (*field.CountryOfIssueField, quickfix.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from Confirmation.
func (m Message) GetCountryOfIssue(f *field.CountryOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for Confirmation.
func (m Message) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, quickfix.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from Confirmation.
func (m Message) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for Confirmation.
func (m Message) LocaleOfIssue() (*field.LocaleOfIssueField, quickfix.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from Confirmation.
func (m Message) GetLocaleOfIssue(f *field.LocaleOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for Confirmation.
func (m Message) RedemptionDate() (*field.RedemptionDateField, quickfix.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from Confirmation.
func (m Message) GetRedemptionDate(f *field.RedemptionDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for Confirmation.
func (m Message) StrikePrice() (*field.StrikePriceField, quickfix.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from Confirmation.
func (m Message) GetStrikePrice(f *field.StrikePriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for Confirmation.
func (m Message) StrikeCurrency() (*field.StrikeCurrencyField, quickfix.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from Confirmation.
func (m Message) GetStrikeCurrency(f *field.StrikeCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for Confirmation.
func (m Message) OptAttribute() (*field.OptAttributeField, quickfix.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from Confirmation.
func (m Message) GetOptAttribute(f *field.OptAttributeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for Confirmation.
func (m Message) ContractMultiplier() (*field.ContractMultiplierField, quickfix.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from Confirmation.
func (m Message) GetContractMultiplier(f *field.ContractMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for Confirmation.
func (m Message) CouponRate() (*field.CouponRateField, quickfix.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from Confirmation.
func (m Message) GetCouponRate(f *field.CouponRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for Confirmation.
func (m Message) SecurityExchange() (*field.SecurityExchangeField, quickfix.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from Confirmation.
func (m Message) GetSecurityExchange(f *field.SecurityExchangeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for Confirmation.
func (m Message) Issuer() (*field.IssuerField, quickfix.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from Confirmation.
func (m Message) GetIssuer(f *field.IssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for Confirmation.
func (m Message) EncodedIssuerLen() (*field.EncodedIssuerLenField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from Confirmation.
func (m Message) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for Confirmation.
func (m Message) EncodedIssuer() (*field.EncodedIssuerField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from Confirmation.
func (m Message) GetEncodedIssuer(f *field.EncodedIssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for Confirmation.
func (m Message) SecurityDesc() (*field.SecurityDescField, quickfix.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from Confirmation.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for Confirmation.
func (m Message) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from Confirmation.
func (m Message) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for Confirmation.
func (m Message) EncodedSecurityDesc() (*field.EncodedSecurityDescField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from Confirmation.
func (m Message) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for Confirmation.
func (m Message) Pool() (*field.PoolField, quickfix.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from Confirmation.
func (m Message) GetPool(f *field.PoolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for Confirmation.
func (m Message) ContractSettlMonth() (*field.ContractSettlMonthField, quickfix.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from Confirmation.
func (m Message) GetContractSettlMonth(f *field.ContractSettlMonthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for Confirmation.
func (m Message) CPProgram() (*field.CPProgramField, quickfix.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from Confirmation.
func (m Message) GetCPProgram(f *field.CPProgramField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for Confirmation.
func (m Message) CPRegType() (*field.CPRegTypeField, quickfix.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from Confirmation.
func (m Message) GetCPRegType(f *field.CPRegTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for Confirmation.
func (m Message) NoEvents() (*field.NoEventsField, quickfix.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from Confirmation.
func (m Message) GetNoEvents(f *field.NoEventsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for Confirmation.
func (m Message) DatedDate() (*field.DatedDateField, quickfix.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from Confirmation.
func (m Message) GetDatedDate(f *field.DatedDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for Confirmation.
func (m Message) InterestAccrualDate() (*field.InterestAccrualDateField, quickfix.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from Confirmation.
func (m Message) GetInterestAccrualDate(f *field.InterestAccrualDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryForm is a non-required field for Confirmation.
func (m Message) DeliveryForm() (*field.DeliveryFormField, quickfix.MessageRejectError) {
	f := &field.DeliveryFormField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryForm reads a DeliveryForm from Confirmation.
func (m Message) GetDeliveryForm(f *field.DeliveryFormField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PctAtRisk is a non-required field for Confirmation.
func (m Message) PctAtRisk() (*field.PctAtRiskField, quickfix.MessageRejectError) {
	f := &field.PctAtRiskField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPctAtRisk reads a PctAtRisk from Confirmation.
func (m Message) GetPctAtRisk(f *field.PctAtRiskField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrAttrib is a non-required field for Confirmation.
func (m Message) NoInstrAttrib() (*field.NoInstrAttribField, quickfix.MessageRejectError) {
	f := &field.NoInstrAttribField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrAttrib reads a NoInstrAttrib from Confirmation.
func (m Message) GetNoInstrAttrib(f *field.NoInstrAttribField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for Confirmation.
func (m Message) AgreementDesc() (*field.AgreementDescField, quickfix.MessageRejectError) {
	f := &field.AgreementDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from Confirmation.
func (m Message) GetAgreementDesc(f *field.AgreementDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for Confirmation.
func (m Message) AgreementID() (*field.AgreementIDField, quickfix.MessageRejectError) {
	f := &field.AgreementIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from Confirmation.
func (m Message) GetAgreementID(f *field.AgreementIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for Confirmation.
func (m Message) AgreementDate() (*field.AgreementDateField, quickfix.MessageRejectError) {
	f := &field.AgreementDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from Confirmation.
func (m Message) GetAgreementDate(f *field.AgreementDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for Confirmation.
func (m Message) AgreementCurrency() (*field.AgreementCurrencyField, quickfix.MessageRejectError) {
	f := &field.AgreementCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from Confirmation.
func (m Message) GetAgreementCurrency(f *field.AgreementCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for Confirmation.
func (m Message) TerminationType() (*field.TerminationTypeField, quickfix.MessageRejectError) {
	f := &field.TerminationTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from Confirmation.
func (m Message) GetTerminationType(f *field.TerminationTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for Confirmation.
func (m Message) StartDate() (*field.StartDateField, quickfix.MessageRejectError) {
	f := &field.StartDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from Confirmation.
func (m Message) GetStartDate(f *field.StartDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for Confirmation.
func (m Message) EndDate() (*field.EndDateField, quickfix.MessageRejectError) {
	f := &field.EndDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from Confirmation.
func (m Message) GetEndDate(f *field.EndDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for Confirmation.
func (m Message) DeliveryType() (*field.DeliveryTypeField, quickfix.MessageRejectError) {
	f := &field.DeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from Confirmation.
func (m Message) GetDeliveryType(f *field.DeliveryTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for Confirmation.
func (m Message) MarginRatio() (*field.MarginRatioField, quickfix.MessageRejectError) {
	f := &field.MarginRatioField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from Confirmation.
func (m Message) GetMarginRatio(f *field.MarginRatioField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a required field for Confirmation.
func (m Message) NoUnderlyings() (*field.NoUnderlyingsField, quickfix.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from Confirmation.
func (m Message) GetNoUnderlyings(f *field.NoUnderlyingsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a required field for Confirmation.
func (m Message) NoLegs() (*field.NoLegsField, quickfix.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from Confirmation.
func (m Message) GetNoLegs(f *field.NoLegsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for Confirmation.
func (m Message) YieldType() (*field.YieldTypeField, quickfix.MessageRejectError) {
	f := &field.YieldTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from Confirmation.
func (m Message) GetYieldType(f *field.YieldTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for Confirmation.
func (m Message) Yield() (*field.YieldField, quickfix.MessageRejectError) {
	f := &field.YieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from Confirmation.
func (m Message) GetYield(f *field.YieldField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for Confirmation.
func (m Message) YieldCalcDate() (*field.YieldCalcDateField, quickfix.MessageRejectError) {
	f := &field.YieldCalcDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from Confirmation.
func (m Message) GetYieldCalcDate(f *field.YieldCalcDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for Confirmation.
func (m Message) YieldRedemptionDate() (*field.YieldRedemptionDateField, quickfix.MessageRejectError) {
	f := &field.YieldRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from Confirmation.
func (m Message) GetYieldRedemptionDate(f *field.YieldRedemptionDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for Confirmation.
func (m Message) YieldRedemptionPrice() (*field.YieldRedemptionPriceField, quickfix.MessageRejectError) {
	f := &field.YieldRedemptionPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from Confirmation.
func (m Message) GetYieldRedemptionPrice(f *field.YieldRedemptionPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for Confirmation.
func (m Message) YieldRedemptionPriceType() (*field.YieldRedemptionPriceTypeField, quickfix.MessageRejectError) {
	f := &field.YieldRedemptionPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from Confirmation.
func (m Message) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocQty is a required field for Confirmation.
func (m Message) AllocQty() (*field.AllocQtyField, quickfix.MessageRejectError) {
	f := &field.AllocQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocQty reads a AllocQty from Confirmation.
func (m Message) GetAllocQty(f *field.AllocQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for Confirmation.
func (m Message) QtyType() (*field.QtyTypeField, quickfix.MessageRejectError) {
	f := &field.QtyTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from Confirmation.
func (m Message) GetQtyType(f *field.QtyTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for Confirmation.
func (m Message) Side() (*field.SideField, quickfix.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from Confirmation.
func (m Message) GetSide(f *field.SideField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for Confirmation.
func (m Message) Currency() (*field.CurrencyField, quickfix.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from Confirmation.
func (m Message) GetCurrency(f *field.CurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for Confirmation.
func (m Message) LastMkt() (*field.LastMktField, quickfix.MessageRejectError) {
	f := &field.LastMktField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from Confirmation.
func (m Message) GetLastMkt(f *field.LastMktField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoCapacities is a required field for Confirmation.
func (m Message) NoCapacities() (*field.NoCapacitiesField, quickfix.MessageRejectError) {
	f := &field.NoCapacitiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoCapacities reads a NoCapacities from Confirmation.
func (m Message) GetNoCapacities(f *field.NoCapacitiesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAccount is a required field for Confirmation.
func (m Message) AllocAccount() (*field.AllocAccountField, quickfix.MessageRejectError) {
	f := &field.AllocAccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAccount reads a AllocAccount from Confirmation.
func (m Message) GetAllocAccount(f *field.AllocAccountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAcctIDSource is a non-required field for Confirmation.
func (m Message) AllocAcctIDSource() (*field.AllocAcctIDSourceField, quickfix.MessageRejectError) {
	f := &field.AllocAcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAcctIDSource reads a AllocAcctIDSource from Confirmation.
func (m Message) GetAllocAcctIDSource(f *field.AllocAcctIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAccountType is a non-required field for Confirmation.
func (m Message) AllocAccountType() (*field.AllocAccountTypeField, quickfix.MessageRejectError) {
	f := &field.AllocAccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAccountType reads a AllocAccountType from Confirmation.
func (m Message) GetAllocAccountType(f *field.AllocAccountTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a required field for Confirmation.
func (m Message) AvgPx() (*field.AvgPxField, quickfix.MessageRejectError) {
	f := &field.AvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from Confirmation.
func (m Message) GetAvgPx(f *field.AvgPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxPrecision is a non-required field for Confirmation.
func (m Message) AvgPxPrecision() (*field.AvgPxPrecisionField, quickfix.MessageRejectError) {
	f := &field.AvgPxPrecisionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxPrecision reads a AvgPxPrecision from Confirmation.
func (m Message) GetAvgPxPrecision(f *field.AvgPxPrecisionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for Confirmation.
func (m Message) PriceType() (*field.PriceTypeField, quickfix.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from Confirmation.
func (m Message) GetPriceType(f *field.PriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AvgParPx is a non-required field for Confirmation.
func (m Message) AvgParPx() (*field.AvgParPxField, quickfix.MessageRejectError) {
	f := &field.AvgParPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgParPx reads a AvgParPx from Confirmation.
func (m Message) GetAvgParPx(f *field.AvgParPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for Confirmation.
func (m Message) Spread() (*field.SpreadField, quickfix.MessageRejectError) {
	f := &field.SpreadField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from Confirmation.
func (m Message) GetSpread(f *field.SpreadField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for Confirmation.
func (m Message) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrencyField, quickfix.MessageRejectError) {
	f := &field.BenchmarkCurveCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from Confirmation.
func (m Message) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for Confirmation.
func (m Message) BenchmarkCurveName() (*field.BenchmarkCurveNameField, quickfix.MessageRejectError) {
	f := &field.BenchmarkCurveNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from Confirmation.
func (m Message) GetBenchmarkCurveName(f *field.BenchmarkCurveNameField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for Confirmation.
func (m Message) BenchmarkCurvePoint() (*field.BenchmarkCurvePointField, quickfix.MessageRejectError) {
	f := &field.BenchmarkCurvePointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from Confirmation.
func (m Message) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePointField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for Confirmation.
func (m Message) BenchmarkPrice() (*field.BenchmarkPriceField, quickfix.MessageRejectError) {
	f := &field.BenchmarkPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from Confirmation.
func (m Message) GetBenchmarkPrice(f *field.BenchmarkPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for Confirmation.
func (m Message) BenchmarkPriceType() (*field.BenchmarkPriceTypeField, quickfix.MessageRejectError) {
	f := &field.BenchmarkPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from Confirmation.
func (m Message) GetBenchmarkPriceType(f *field.BenchmarkPriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for Confirmation.
func (m Message) BenchmarkSecurityID() (*field.BenchmarkSecurityIDField, quickfix.MessageRejectError) {
	f := &field.BenchmarkSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from Confirmation.
func (m Message) GetBenchmarkSecurityID(f *field.BenchmarkSecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for Confirmation.
func (m Message) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.BenchmarkSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from Confirmation.
func (m Message) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ReportedPx is a non-required field for Confirmation.
func (m Message) ReportedPx() (*field.ReportedPxField, quickfix.MessageRejectError) {
	f := &field.ReportedPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetReportedPx reads a ReportedPx from Confirmation.
func (m Message) GetReportedPx(f *field.ReportedPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for Confirmation.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from Confirmation.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for Confirmation.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from Confirmation.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for Confirmation.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from Confirmation.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for Confirmation.
func (m Message) ProcessCode() (*field.ProcessCodeField, quickfix.MessageRejectError) {
	f := &field.ProcessCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from Confirmation.
func (m Message) GetProcessCode(f *field.ProcessCodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a required field for Confirmation.
func (m Message) GrossTradeAmt() (*field.GrossTradeAmtField, quickfix.MessageRejectError) {
	f := &field.GrossTradeAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from Confirmation.
func (m Message) GetGrossTradeAmt(f *field.GrossTradeAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NumDaysInterest is a non-required field for Confirmation.
func (m Message) NumDaysInterest() (*field.NumDaysInterestField, quickfix.MessageRejectError) {
	f := &field.NumDaysInterestField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNumDaysInterest reads a NumDaysInterest from Confirmation.
func (m Message) GetNumDaysInterest(f *field.NumDaysInterestField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExDate is a non-required field for Confirmation.
func (m Message) ExDate() (*field.ExDateField, quickfix.MessageRejectError) {
	f := &field.ExDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDate reads a ExDate from Confirmation.
func (m Message) GetExDate(f *field.ExDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestRate is a non-required field for Confirmation.
func (m Message) AccruedInterestRate() (*field.AccruedInterestRateField, quickfix.MessageRejectError) {
	f := &field.AccruedInterestRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestRate reads a AccruedInterestRate from Confirmation.
func (m Message) GetAccruedInterestRate(f *field.AccruedInterestRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for Confirmation.
func (m Message) AccruedInterestAmt() (*field.AccruedInterestAmtField, quickfix.MessageRejectError) {
	f := &field.AccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from Confirmation.
func (m Message) GetAccruedInterestAmt(f *field.AccruedInterestAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAtMaturity is a non-required field for Confirmation.
func (m Message) InterestAtMaturity() (*field.InterestAtMaturityField, quickfix.MessageRejectError) {
	f := &field.InterestAtMaturityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAtMaturity reads a InterestAtMaturity from Confirmation.
func (m Message) GetInterestAtMaturity(f *field.InterestAtMaturityField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EndAccruedInterestAmt is a non-required field for Confirmation.
func (m Message) EndAccruedInterestAmt() (*field.EndAccruedInterestAmtField, quickfix.MessageRejectError) {
	f := &field.EndAccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndAccruedInterestAmt reads a EndAccruedInterestAmt from Confirmation.
func (m Message) GetEndAccruedInterestAmt(f *field.EndAccruedInterestAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StartCash is a non-required field for Confirmation.
func (m Message) StartCash() (*field.StartCashField, quickfix.MessageRejectError) {
	f := &field.StartCashField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartCash reads a StartCash from Confirmation.
func (m Message) GetStartCash(f *field.StartCashField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EndCash is a non-required field for Confirmation.
func (m Message) EndCash() (*field.EndCashField, quickfix.MessageRejectError) {
	f := &field.EndCashField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndCash reads a EndCash from Confirmation.
func (m Message) GetEndCash(f *field.EndCashField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Concession is a non-required field for Confirmation.
func (m Message) Concession() (*field.ConcessionField, quickfix.MessageRejectError) {
	f := &field.ConcessionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConcession reads a Concession from Confirmation.
func (m Message) GetConcession(f *field.ConcessionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TotalTakedown is a non-required field for Confirmation.
func (m Message) TotalTakedown() (*field.TotalTakedownField, quickfix.MessageRejectError) {
	f := &field.TotalTakedownField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalTakedown reads a TotalTakedown from Confirmation.
func (m Message) GetTotalTakedown(f *field.TotalTakedownField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a required field for Confirmation.
func (m Message) NetMoney() (*field.NetMoneyField, quickfix.MessageRejectError) {
	f := &field.NetMoneyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from Confirmation.
func (m Message) GetNetMoney(f *field.NetMoneyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityNetMoney is a non-required field for Confirmation.
func (m Message) MaturityNetMoney() (*field.MaturityNetMoneyField, quickfix.MessageRejectError) {
	f := &field.MaturityNetMoneyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityNetMoney reads a MaturityNetMoney from Confirmation.
func (m Message) GetMaturityNetMoney(f *field.MaturityNetMoneyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrAmt is a non-required field for Confirmation.
func (m Message) SettlCurrAmt() (*field.SettlCurrAmtField, quickfix.MessageRejectError) {
	f := &field.SettlCurrAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrAmt reads a SettlCurrAmt from Confirmation.
func (m Message) GetSettlCurrAmt(f *field.SettlCurrAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for Confirmation.
func (m Message) SettlCurrency() (*field.SettlCurrencyField, quickfix.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from Confirmation.
func (m Message) GetSettlCurrency(f *field.SettlCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRate is a non-required field for Confirmation.
func (m Message) SettlCurrFxRate() (*field.SettlCurrFxRateField, quickfix.MessageRejectError) {
	f := &field.SettlCurrFxRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRate reads a SettlCurrFxRate from Confirmation.
func (m Message) GetSettlCurrFxRate(f *field.SettlCurrFxRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRateCalc is a non-required field for Confirmation.
func (m Message) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalcField, quickfix.MessageRejectError) {
	f := &field.SettlCurrFxRateCalcField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRateCalc reads a SettlCurrFxRateCalc from Confirmation.
func (m Message) GetSettlCurrFxRateCalc(f *field.SettlCurrFxRateCalcField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for Confirmation.
func (m Message) SettlType() (*field.SettlTypeField, quickfix.MessageRejectError) {
	f := &field.SettlTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from Confirmation.
func (m Message) GetSettlType(f *field.SettlTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for Confirmation.
func (m Message) SettlDate() (*field.SettlDateField, quickfix.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from Confirmation.
func (m Message) GetSettlDate(f *field.SettlDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDeliveryType is a non-required field for Confirmation.
func (m Message) SettlDeliveryType() (*field.SettlDeliveryTypeField, quickfix.MessageRejectError) {
	f := &field.SettlDeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDeliveryType reads a SettlDeliveryType from Confirmation.
func (m Message) GetSettlDeliveryType(f *field.SettlDeliveryTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbType is a non-required field for Confirmation.
func (m Message) StandInstDbType() (*field.StandInstDbTypeField, quickfix.MessageRejectError) {
	f := &field.StandInstDbTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbType reads a StandInstDbType from Confirmation.
func (m Message) GetStandInstDbType(f *field.StandInstDbTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbName is a non-required field for Confirmation.
func (m Message) StandInstDbName() (*field.StandInstDbNameField, quickfix.MessageRejectError) {
	f := &field.StandInstDbNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbName reads a StandInstDbName from Confirmation.
func (m Message) GetStandInstDbName(f *field.StandInstDbNameField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbID is a non-required field for Confirmation.
func (m Message) StandInstDbID() (*field.StandInstDbIDField, quickfix.MessageRejectError) {
	f := &field.StandInstDbIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbID reads a StandInstDbID from Confirmation.
func (m Message) GetStandInstDbID(f *field.StandInstDbIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoDlvyInst is a non-required field for Confirmation.
func (m Message) NoDlvyInst() (*field.NoDlvyInstField, quickfix.MessageRejectError) {
	f := &field.NoDlvyInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoDlvyInst reads a NoDlvyInst from Confirmation.
func (m Message) GetNoDlvyInst(f *field.NoDlvyInstField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for Confirmation.
func (m Message) Commission() (*field.CommissionField, quickfix.MessageRejectError) {
	f := &field.CommissionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from Confirmation.
func (m Message) GetCommission(f *field.CommissionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for Confirmation.
func (m Message) CommType() (*field.CommTypeField, quickfix.MessageRejectError) {
	f := &field.CommTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from Confirmation.
func (m Message) GetCommType(f *field.CommTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CommCurrency is a non-required field for Confirmation.
func (m Message) CommCurrency() (*field.CommCurrencyField, quickfix.MessageRejectError) {
	f := &field.CommCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommCurrency reads a CommCurrency from Confirmation.
func (m Message) GetCommCurrency(f *field.CommCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FundRenewWaiv is a non-required field for Confirmation.
func (m Message) FundRenewWaiv() (*field.FundRenewWaivField, quickfix.MessageRejectError) {
	f := &field.FundRenewWaivField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFundRenewWaiv reads a FundRenewWaiv from Confirmation.
func (m Message) GetFundRenewWaiv(f *field.FundRenewWaivField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SharedCommission is a non-required field for Confirmation.
func (m Message) SharedCommission() (*field.SharedCommissionField, quickfix.MessageRejectError) {
	f := &field.SharedCommissionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSharedCommission reads a SharedCommission from Confirmation.
func (m Message) GetSharedCommission(f *field.SharedCommissionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for Confirmation.
func (m Message) NoStipulations() (*field.NoStipulationsField, quickfix.MessageRejectError) {
	f := &field.NoStipulationsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from Confirmation.
func (m Message) GetNoStipulations(f *field.NoStipulationsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoMiscFees is a non-required field for Confirmation.
func (m Message) NoMiscFees() (*field.NoMiscFeesField, quickfix.MessageRejectError) {
	f := &field.NoMiscFeesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMiscFees reads a NoMiscFees from Confirmation.
func (m Message) GetNoMiscFees(f *field.NoMiscFeesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds Confirmation messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for Confirmation.
func Builder(
	confirmid *field.ConfirmIDField,
	confirmtranstype *field.ConfirmTransTypeField,
	confirmtype *field.ConfirmTypeField,
	confirmstatus *field.ConfirmStatusField,
	transacttime *field.TransactTimeField,
	tradedate *field.TradeDateField,
	nounderlyings *field.NoUnderlyingsField,
	nolegs *field.NoLegsField,
	allocqty *field.AllocQtyField,
	side *field.SideField,
	nocapacities *field.NoCapacitiesField,
	allocaccount *field.AllocAccountField,
	avgpx *field.AvgPxField,
	grosstradeamt *field.GrossTradeAmtField,
	netmoney *field.NetMoneyField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header.Set(field.NewMsgType("AK"))
	builder.Body.Set(confirmid)
	builder.Body.Set(confirmtranstype)
	builder.Body.Set(confirmtype)
	builder.Body.Set(confirmstatus)
	builder.Body.Set(transacttime)
	builder.Body.Set(tradedate)
	builder.Body.Set(nounderlyings)
	builder.Body.Set(nolegs)
	builder.Body.Set(allocqty)
	builder.Body.Set(side)
	builder.Body.Set(nocapacities)
	builder.Body.Set(allocaccount)
	builder.Body.Set(avgpx)
	builder.Body.Set(grosstradeamt)
	builder.Body.Set(netmoney)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX44, "AK", r
}
