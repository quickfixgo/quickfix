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

//OrderStatusRequest msg type = H.
type OrderStatusRequest struct {
	message.Message
}

//OrderStatusRequestBuilder builds OrderStatusRequest messages.
type OrderStatusRequestBuilder struct {
	message.MessageBuilder
}

//CreateOrderStatusRequestBuilder returns an initialized OrderStatusRequestBuilder with specified required fields.
func CreateOrderStatusRequestBuilder(
	side *field.SideField) OrderStatusRequestBuilder {
	var builder OrderStatusRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("H"))
	builder.Body.Set(side)
	return builder
}

//OrderID is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderStatusRequest.
func (m OrderStatusRequest) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderStatusRequest.
func (m OrderStatusRequest) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecondaryClOrdID() (*field.SecondaryClOrdIDField, errors.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from OrderStatusRequest.
func (m OrderStatusRequest) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) ClOrdLinkID() (*field.ClOrdLinkIDField, errors.MessageRejectError) {
	f := &field.ClOrdLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from OrderStatusRequest.
func (m OrderStatusRequest) GetClOrdLinkID(f *field.ClOrdLinkIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from OrderStatusRequest.
func (m OrderStatusRequest) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdStatusReqID is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) OrdStatusReqID() (*field.OrdStatusReqIDField, errors.MessageRejectError) {
	f := &field.OrdStatusReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdStatusReqID reads a OrdStatusReqID from OrderStatusRequest.
func (m OrderStatusRequest) GetOrdStatusReqID(f *field.OrdStatusReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from OrderStatusRequest.
func (m OrderStatusRequest) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) AcctIDSource() (*field.AcctIDSourceField, errors.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from OrderStatusRequest.
func (m OrderStatusRequest) GetAcctIDSource(f *field.AcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderStatusRequest.
func (m OrderStatusRequest) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderStatusRequest.
func (m OrderStatusRequest) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderStatusRequest.
func (m OrderStatusRequest) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from OrderStatusRequest.
func (m OrderStatusRequest) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from OrderStatusRequest.
func (m OrderStatusRequest) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from OrderStatusRequest.
func (m OrderStatusRequest) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from OrderStatusRequest.
func (m OrderStatusRequest) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from OrderStatusRequest.
func (m OrderStatusRequest) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from OrderStatusRequest.
func (m OrderStatusRequest) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from OrderStatusRequest.
func (m OrderStatusRequest) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from OrderStatusRequest.
func (m OrderStatusRequest) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from OrderStatusRequest.
func (m OrderStatusRequest) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from OrderStatusRequest.
func (m OrderStatusRequest) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from OrderStatusRequest.
func (m OrderStatusRequest) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from OrderStatusRequest.
func (m OrderStatusRequest) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from OrderStatusRequest.
func (m OrderStatusRequest) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from OrderStatusRequest.
func (m OrderStatusRequest) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from OrderStatusRequest.
func (m OrderStatusRequest) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from OrderStatusRequest.
func (m OrderStatusRequest) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from OrderStatusRequest.
func (m OrderStatusRequest) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from OrderStatusRequest.
func (m OrderStatusRequest) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from OrderStatusRequest.
func (m OrderStatusRequest) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from OrderStatusRequest.
func (m OrderStatusRequest) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from OrderStatusRequest.
func (m OrderStatusRequest) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from OrderStatusRequest.
func (m OrderStatusRequest) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from OrderStatusRequest.
func (m OrderStatusRequest) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from OrderStatusRequest.
func (m OrderStatusRequest) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from OrderStatusRequest.
func (m OrderStatusRequest) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from OrderStatusRequest.
func (m OrderStatusRequest) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderStatusRequest.
func (m OrderStatusRequest) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from OrderStatusRequest.
func (m OrderStatusRequest) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from OrderStatusRequest.
func (m OrderStatusRequest) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderStatusRequest.
func (m OrderStatusRequest) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from OrderStatusRequest.
func (m OrderStatusRequest) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from OrderStatusRequest.
func (m OrderStatusRequest) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from OrderStatusRequest.
func (m OrderStatusRequest) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from OrderStatusRequest.
func (m OrderStatusRequest) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from OrderStatusRequest.
func (m OrderStatusRequest) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from OrderStatusRequest.
func (m OrderStatusRequest) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from OrderStatusRequest.
func (m OrderStatusRequest) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from OrderStatusRequest.
func (m OrderStatusRequest) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from OrderStatusRequest.
func (m OrderStatusRequest) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityStatus() (*field.SecurityStatusField, errors.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from OrderStatusRequest.
func (m OrderStatusRequest) GetSecurityStatus(f *field.SecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from OrderStatusRequest.
func (m OrderStatusRequest) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from OrderStatusRequest.
func (m OrderStatusRequest) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) StrikeMultiplier() (*field.StrikeMultiplierField, errors.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from OrderStatusRequest.
func (m OrderStatusRequest) GetStrikeMultiplier(f *field.StrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) StrikeValue() (*field.StrikeValueField, errors.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from OrderStatusRequest.
func (m OrderStatusRequest) GetStrikeValue(f *field.StrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) MinPriceIncrement() (*field.MinPriceIncrementField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from OrderStatusRequest.
func (m OrderStatusRequest) GetMinPriceIncrement(f *field.MinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) PositionLimit() (*field.PositionLimitField, errors.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from OrderStatusRequest.
func (m OrderStatusRequest) GetPositionLimit(f *field.PositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) NTPositionLimit() (*field.NTPositionLimitField, errors.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from OrderStatusRequest.
func (m OrderStatusRequest) GetNTPositionLimit(f *field.NTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) NoInstrumentParties() (*field.NoInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from OrderStatusRequest.
func (m OrderStatusRequest) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) UnitOfMeasure() (*field.UnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from OrderStatusRequest.
func (m OrderStatusRequest) GetUnitOfMeasure(f *field.UnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) TimeUnit() (*field.TimeUnitField, errors.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from OrderStatusRequest.
func (m OrderStatusRequest) GetTimeUnit(f *field.TimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) MaturityTime() (*field.MaturityTimeField, errors.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from OrderStatusRequest.
func (m OrderStatusRequest) GetMaturityTime(f *field.MaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityGroup() (*field.SecurityGroupField, errors.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from OrderStatusRequest.
func (m OrderStatusRequest) GetSecurityGroup(f *field.SecurityGroupField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from OrderStatusRequest.
func (m OrderStatusRequest) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from OrderStatusRequest.
func (m OrderStatusRequest) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityXMLLen() (*field.SecurityXMLLenField, errors.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from OrderStatusRequest.
func (m OrderStatusRequest) GetSecurityXMLLen(f *field.SecurityXMLLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityXML() (*field.SecurityXMLField, errors.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from OrderStatusRequest.
func (m OrderStatusRequest) GetSecurityXML(f *field.SecurityXMLField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityXMLSchema() (*field.SecurityXMLSchemaField, errors.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from OrderStatusRequest.
func (m OrderStatusRequest) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) ProductComplex() (*field.ProductComplexField, errors.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from OrderStatusRequest.
func (m OrderStatusRequest) GetProductComplex(f *field.ProductComplexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from OrderStatusRequest.
func (m OrderStatusRequest) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from OrderStatusRequest.
func (m OrderStatusRequest) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SettlMethod() (*field.SettlMethodField, errors.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from OrderStatusRequest.
func (m OrderStatusRequest) GetSettlMethod(f *field.SettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) ExerciseStyle() (*field.ExerciseStyleField, errors.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from OrderStatusRequest.
func (m OrderStatusRequest) GetExerciseStyle(f *field.ExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) OptPayoutAmount() (*field.OptPayoutAmountField, errors.MessageRejectError) {
	f := &field.OptPayoutAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from OrderStatusRequest.
func (m OrderStatusRequest) GetOptPayoutAmount(f *field.OptPayoutAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) PriceQuoteMethod() (*field.PriceQuoteMethodField, errors.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from OrderStatusRequest.
func (m OrderStatusRequest) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) ListMethod() (*field.ListMethodField, errors.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from OrderStatusRequest.
func (m OrderStatusRequest) GetListMethod(f *field.ListMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) CapPrice() (*field.CapPriceField, errors.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from OrderStatusRequest.
func (m OrderStatusRequest) GetCapPrice(f *field.CapPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) FloorPrice() (*field.FloorPriceField, errors.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from OrderStatusRequest.
func (m OrderStatusRequest) GetFloorPrice(f *field.FloorPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from OrderStatusRequest.
func (m OrderStatusRequest) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) FlexibleIndicator() (*field.FlexibleIndicatorField, errors.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from OrderStatusRequest.
func (m OrderStatusRequest) GetFlexibleIndicator(f *field.FlexibleIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, errors.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from OrderStatusRequest.
func (m OrderStatusRequest) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) ValuationMethod() (*field.ValuationMethodField, errors.MessageRejectError) {
	f := &field.ValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from OrderStatusRequest.
func (m OrderStatusRequest) GetValuationMethod(f *field.ValuationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) ContractMultiplierUnit() (*field.ContractMultiplierUnitField, errors.MessageRejectError) {
	f := &field.ContractMultiplierUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from OrderStatusRequest.
func (m OrderStatusRequest) GetContractMultiplierUnit(f *field.ContractMultiplierUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) FlowScheduleType() (*field.FlowScheduleTypeField, errors.MessageRejectError) {
	f := &field.FlowScheduleTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from OrderStatusRequest.
func (m OrderStatusRequest) GetFlowScheduleType(f *field.FlowScheduleTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) RestructuringType() (*field.RestructuringTypeField, errors.MessageRejectError) {
	f := &field.RestructuringTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from OrderStatusRequest.
func (m OrderStatusRequest) GetRestructuringType(f *field.RestructuringTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) Seniority() (*field.SeniorityField, errors.MessageRejectError) {
	f := &field.SeniorityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from OrderStatusRequest.
func (m OrderStatusRequest) GetSeniority(f *field.SeniorityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstandingField, errors.MessageRejectError) {
	f := &field.NotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from OrderStatusRequest.
func (m OrderStatusRequest) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstandingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstandingField, errors.MessageRejectError) {
	f := &field.OriginalNotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from OrderStatusRequest.
func (m OrderStatusRequest) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstandingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) AttachmentPoint() (*field.AttachmentPointField, errors.MessageRejectError) {
	f := &field.AttachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from OrderStatusRequest.
func (m OrderStatusRequest) GetAttachmentPoint(f *field.AttachmentPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) DetachmentPoint() (*field.DetachmentPointField, errors.MessageRejectError) {
	f := &field.DetachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from OrderStatusRequest.
func (m OrderStatusRequest) GetDetachmentPoint(f *field.DetachmentPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethodField, errors.MessageRejectError) {
	f := &field.StrikePriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from OrderStatusRequest.
func (m OrderStatusRequest) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethodField, errors.MessageRejectError) {
	f := &field.StrikePriceBoundaryMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from OrderStatusRequest.
func (m OrderStatusRequest) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecisionField, errors.MessageRejectError) {
	f := &field.StrikePriceBoundaryPrecisionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from OrderStatusRequest.
func (m OrderStatusRequest) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecisionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethodField, errors.MessageRejectError) {
	f := &field.UnderlyingPriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from OrderStatusRequest.
func (m OrderStatusRequest) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) OptPayoutType() (*field.OptPayoutTypeField, errors.MessageRejectError) {
	f := &field.OptPayoutTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from OrderStatusRequest.
func (m OrderStatusRequest) GetOptPayoutType(f *field.OptPayoutTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) NoComplexEvents() (*field.NoComplexEventsField, errors.MessageRejectError) {
	f := &field.NoComplexEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from OrderStatusRequest.
func (m OrderStatusRequest) GetNoComplexEvents(f *field.NoComplexEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) AgreementDesc() (*field.AgreementDescField, errors.MessageRejectError) {
	f := &field.AgreementDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from OrderStatusRequest.
func (m OrderStatusRequest) GetAgreementDesc(f *field.AgreementDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) AgreementID() (*field.AgreementIDField, errors.MessageRejectError) {
	f := &field.AgreementIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from OrderStatusRequest.
func (m OrderStatusRequest) GetAgreementID(f *field.AgreementIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) AgreementDate() (*field.AgreementDateField, errors.MessageRejectError) {
	f := &field.AgreementDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from OrderStatusRequest.
func (m OrderStatusRequest) GetAgreementDate(f *field.AgreementDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) AgreementCurrency() (*field.AgreementCurrencyField, errors.MessageRejectError) {
	f := &field.AgreementCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from OrderStatusRequest.
func (m OrderStatusRequest) GetAgreementCurrency(f *field.AgreementCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) TerminationType() (*field.TerminationTypeField, errors.MessageRejectError) {
	f := &field.TerminationTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from OrderStatusRequest.
func (m OrderStatusRequest) GetTerminationType(f *field.TerminationTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) StartDate() (*field.StartDateField, errors.MessageRejectError) {
	f := &field.StartDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from OrderStatusRequest.
func (m OrderStatusRequest) GetStartDate(f *field.StartDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) EndDate() (*field.EndDateField, errors.MessageRejectError) {
	f := &field.EndDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from OrderStatusRequest.
func (m OrderStatusRequest) GetEndDate(f *field.EndDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) DeliveryType() (*field.DeliveryTypeField, errors.MessageRejectError) {
	f := &field.DeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from OrderStatusRequest.
func (m OrderStatusRequest) GetDeliveryType(f *field.DeliveryTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) MarginRatio() (*field.MarginRatioField, errors.MessageRejectError) {
	f := &field.MarginRatioField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from OrderStatusRequest.
func (m OrderStatusRequest) GetMarginRatio(f *field.MarginRatioField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from OrderStatusRequest.
func (m OrderStatusRequest) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for OrderStatusRequest.
func (m OrderStatusRequest) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderStatusRequest.
func (m OrderStatusRequest) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}
