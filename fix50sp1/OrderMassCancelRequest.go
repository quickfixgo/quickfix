package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//OrderMassCancelRequest msg type = q.
type OrderMassCancelRequest struct {
	message.Message
}

//OrderMassCancelRequestBuilder builds OrderMassCancelRequest messages.
type OrderMassCancelRequestBuilder struct {
	message.MessageBuilder
}

//CreateOrderMassCancelRequestBuilder returns an initialized OrderMassCancelRequestBuilder with specified required fields.
func CreateOrderMassCancelRequestBuilder(
	clordid *field.ClOrdIDField,
	masscancelrequesttype *field.MassCancelRequestTypeField,
	transacttime *field.TransactTimeField) OrderMassCancelRequestBuilder {
	var builder OrderMassCancelRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("q"))
	builder.Body.Set(clordid)
	builder.Body.Set(masscancelrequesttype)
	builder.Body.Set(transacttime)
	return builder
}

//ClOrdID is a required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecondaryClOrdID() (*field.SecondaryClOrdIDField, errors.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MassCancelRequestType is a required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) MassCancelRequestType() (*field.MassCancelRequestTypeField, errors.MessageRejectError) {
	f := &field.MassCancelRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMassCancelRequestType reads a MassCancelRequestType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetMassCancelRequestType(f *field.MassCancelRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) TradingSessionSubID() (*field.TradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetTradingSessionSubID(f *field.TradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityStatus() (*field.SecurityStatusField, errors.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityStatus(f *field.SecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) StrikeMultiplier() (*field.StrikeMultiplierField, errors.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetStrikeMultiplier(f *field.StrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) StrikeValue() (*field.StrikeValueField, errors.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetStrikeValue(f *field.StrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) MinPriceIncrement() (*field.MinPriceIncrementField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetMinPriceIncrement(f *field.MinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) PositionLimit() (*field.PositionLimitField, errors.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetPositionLimit(f *field.PositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NTPositionLimit() (*field.NTPositionLimitField, errors.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNTPositionLimit(f *field.NTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NoInstrumentParties() (*field.NoInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnitOfMeasure() (*field.UnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnitOfMeasure(f *field.UnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) TimeUnit() (*field.TimeUnitField, errors.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetTimeUnit(f *field.TimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) MaturityTime() (*field.MaturityTimeField, errors.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetMaturityTime(f *field.MaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityGroup() (*field.SecurityGroupField, errors.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityGroup(f *field.SecurityGroupField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityXMLLen() (*field.SecurityXMLLenField, errors.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityXMLLen(f *field.SecurityXMLLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityXML() (*field.SecurityXMLField, errors.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityXML(f *field.SecurityXMLField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityXMLSchema() (*field.SecurityXMLSchemaField, errors.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) ProductComplex() (*field.ProductComplexField, errors.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetProductComplex(f *field.ProductComplexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SettlMethod() (*field.SettlMethodField, errors.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSettlMethod(f *field.SettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) ExerciseStyle() (*field.ExerciseStyleField, errors.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetExerciseStyle(f *field.ExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayAmount is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) OptPayAmount() (*field.OptPayAmountField, errors.MessageRejectError) {
	f := &field.OptPayAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayAmount reads a OptPayAmount from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetOptPayAmount(f *field.OptPayAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) PriceQuoteMethod() (*field.PriceQuoteMethodField, errors.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) ListMethod() (*field.ListMethodField, errors.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetListMethod(f *field.ListMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) CapPrice() (*field.CapPriceField, errors.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetCapPrice(f *field.CapPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) FloorPrice() (*field.FloorPriceField, errors.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetFloorPrice(f *field.FloorPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) FlexibleIndicator() (*field.FlexibleIndicatorField, errors.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetFlexibleIndicator(f *field.FlexibleIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, errors.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FuturesValuationMethod is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) FuturesValuationMethod() (*field.FuturesValuationMethodField, errors.MessageRejectError) {
	f := &field.FuturesValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFuturesValuationMethod reads a FuturesValuationMethod from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetFuturesValuationMethod(f *field.FuturesValuationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbol is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSymbol() (*field.UnderlyingSymbolField, errors.MessageRejectError) {
	f := &field.UnderlyingSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbol reads a UnderlyingSymbol from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSymbol(f *field.UnderlyingSymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbolSfx is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfxField, errors.MessageRejectError) {
	f := &field.UnderlyingSymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbolSfx reads a UnderlyingSymbolSfx from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSymbolSfx(f *field.UnderlyingSymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSecurityID() (*field.UnderlyingSecurityIDField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityID reads a UnderlyingSecurityID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSecurityID(f *field.UnderlyingSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityIDSource is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityIDSource reads a UnderlyingSecurityIDSource from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSecurityIDSource(f *field.UnderlyingSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingSecurityAltID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoUnderlyingSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingSecurityAltID reads a NoUnderlyingSecurityAltID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNoUnderlyingSecurityAltID(f *field.NoUnderlyingSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingProduct is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingProduct() (*field.UnderlyingProductField, errors.MessageRejectError) {
	f := &field.UnderlyingProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingProduct reads a UnderlyingProduct from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingProduct(f *field.UnderlyingProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCFICode is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCFICode() (*field.UnderlyingCFICodeField, errors.MessageRejectError) {
	f := &field.UnderlyingCFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCFICode reads a UnderlyingCFICode from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCFICode(f *field.UnderlyingCFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSecurityType() (*field.UnderlyingSecurityTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityType reads a UnderlyingSecurityType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSecurityType(f *field.UnderlyingSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecuritySubType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingSecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecuritySubType reads a UnderlyingSecuritySubType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSecuritySubType(f *field.UnderlyingSecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityMonthYear is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYearField, errors.MessageRejectError) {
	f := &field.UnderlyingMaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityMonthYear reads a UnderlyingMaturityMonthYear from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingMaturityMonthYear(f *field.UnderlyingMaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingMaturityDate() (*field.UnderlyingMaturityDateField, errors.MessageRejectError) {
	f := &field.UnderlyingMaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityDate reads a UnderlyingMaturityDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingMaturityDate(f *field.UnderlyingMaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponPaymentDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDateField, errors.MessageRejectError) {
	f := &field.UnderlyingCouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponPaymentDate reads a UnderlyingCouponPaymentDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCouponPaymentDate(f *field.UnderlyingCouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssueDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingIssueDate() (*field.UnderlyingIssueDateField, errors.MessageRejectError) {
	f := &field.UnderlyingIssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssueDate reads a UnderlyingIssueDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingIssueDate(f *field.UnderlyingIssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepoCollateralSecurityType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingRepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepoCollateralSecurityType reads a UnderlyingRepoCollateralSecurityType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingRepoCollateralSecurityType(f *field.UnderlyingRepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseTerm is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTermField, errors.MessageRejectError) {
	f := &field.UnderlyingRepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseTerm reads a UnderlyingRepurchaseTerm from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingRepurchaseTerm(f *field.UnderlyingRepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseRate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRateField, errors.MessageRejectError) {
	f := &field.UnderlyingRepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseRate reads a UnderlyingRepurchaseRate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingRepurchaseRate(f *field.UnderlyingRepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFactor is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingFactor() (*field.UnderlyingFactorField, errors.MessageRejectError) {
	f := &field.UnderlyingFactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFactor reads a UnderlyingFactor from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingFactor(f *field.UnderlyingFactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCreditRating is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCreditRating() (*field.UnderlyingCreditRatingField, errors.MessageRejectError) {
	f := &field.UnderlyingCreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCreditRating reads a UnderlyingCreditRating from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCreditRating(f *field.UnderlyingCreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingInstrRegistry is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistryField, errors.MessageRejectError) {
	f := &field.UnderlyingInstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingInstrRegistry reads a UnderlyingInstrRegistry from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingInstrRegistry(f *field.UnderlyingInstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCountryOfIssue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssueField, errors.MessageRejectError) {
	f := &field.UnderlyingCountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCountryOfIssue reads a UnderlyingCountryOfIssue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCountryOfIssue(f *field.UnderlyingCountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.UnderlyingStateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStateOrProvinceOfIssue reads a UnderlyingStateOrProvinceOfIssue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingStateOrProvinceOfIssue(f *field.UnderlyingStateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLocaleOfIssue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssueField, errors.MessageRejectError) {
	f := &field.UnderlyingLocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLocaleOfIssue reads a UnderlyingLocaleOfIssue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingLocaleOfIssue(f *field.UnderlyingLocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRedemptionDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDateField, errors.MessageRejectError) {
	f := &field.UnderlyingRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRedemptionDate reads a UnderlyingRedemptionDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingRedemptionDate(f *field.UnderlyingRedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikePrice is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingStrikePrice() (*field.UnderlyingStrikePriceField, errors.MessageRejectError) {
	f := &field.UnderlyingStrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikePrice reads a UnderlyingStrikePrice from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingStrikePrice(f *field.UnderlyingStrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikeCurrency is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrencyField, errors.MessageRejectError) {
	f := &field.UnderlyingStrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikeCurrency reads a UnderlyingStrikeCurrency from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingStrikeCurrency(f *field.UnderlyingStrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOptAttribute is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingOptAttribute() (*field.UnderlyingOptAttributeField, errors.MessageRejectError) {
	f := &field.UnderlyingOptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOptAttribute reads a UnderlyingOptAttribute from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingOptAttribute(f *field.UnderlyingOptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplier is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplierField, errors.MessageRejectError) {
	f := &field.UnderlyingContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplier reads a UnderlyingContractMultiplier from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingContractMultiplier(f *field.UnderlyingContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponRate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCouponRate() (*field.UnderlyingCouponRateField, errors.MessageRejectError) {
	f := &field.UnderlyingCouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponRate reads a UnderlyingCouponRate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCouponRate(f *field.UnderlyingCouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityExchange is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchangeField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityExchange reads a UnderlyingSecurityExchange from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSecurityExchange(f *field.UnderlyingSecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssuer is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingIssuer() (*field.UnderlyingIssuerField, errors.MessageRejectError) {
	f := &field.UnderlyingIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssuer reads a UnderlyingIssuer from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingIssuer(f *field.UnderlyingIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuerLen is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuerLen reads a EncodedUnderlyingIssuerLen from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedUnderlyingIssuerLen(f *field.EncodedUnderlyingIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuer is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuerField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuer reads a EncodedUnderlyingIssuer from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedUnderlyingIssuer(f *field.EncodedUnderlyingIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityDesc is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDescField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityDesc reads a UnderlyingSecurityDesc from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSecurityDesc(f *field.UnderlyingSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDescLen is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDescLen reads a EncodedUnderlyingSecurityDescLen from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedUnderlyingSecurityDescLen(f *field.EncodedUnderlyingSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDesc is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDesc reads a EncodedUnderlyingSecurityDesc from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedUnderlyingSecurityDesc(f *field.EncodedUnderlyingSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPProgram is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCPProgram() (*field.UnderlyingCPProgramField, errors.MessageRejectError) {
	f := &field.UnderlyingCPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPProgram reads a UnderlyingCPProgram from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCPProgram(f *field.UnderlyingCPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPRegType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCPRegType() (*field.UnderlyingCPRegTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingCPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPRegType reads a UnderlyingCPRegType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCPRegType(f *field.UnderlyingCPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrency is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCurrency() (*field.UnderlyingCurrencyField, errors.MessageRejectError) {
	f := &field.UnderlyingCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrency reads a UnderlyingCurrency from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCurrency(f *field.UnderlyingCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingQty is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingQty() (*field.UnderlyingQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingQty reads a UnderlyingQty from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingQty(f *field.UnderlyingQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPx is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingPx() (*field.UnderlyingPxField, errors.MessageRejectError) {
	f := &field.UnderlyingPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPx reads a UnderlyingPx from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingPx(f *field.UnderlyingPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDirtyPrice is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPriceField, errors.MessageRejectError) {
	f := &field.UnderlyingDirtyPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDirtyPrice reads a UnderlyingDirtyPrice from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingDirtyPrice(f *field.UnderlyingDirtyPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndPrice is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingEndPrice() (*field.UnderlyingEndPriceField, errors.MessageRejectError) {
	f := &field.UnderlyingEndPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndPrice reads a UnderlyingEndPrice from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingEndPrice(f *field.UnderlyingEndPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStartValue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingStartValue() (*field.UnderlyingStartValueField, errors.MessageRejectError) {
	f := &field.UnderlyingStartValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStartValue reads a UnderlyingStartValue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingStartValue(f *field.UnderlyingStartValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrentValue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCurrentValue() (*field.UnderlyingCurrentValueField, errors.MessageRejectError) {
	f := &field.UnderlyingCurrentValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrentValue reads a UnderlyingCurrentValue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCurrentValue(f *field.UnderlyingCurrentValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndValue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingEndValue() (*field.UnderlyingEndValueField, errors.MessageRejectError) {
	f := &field.UnderlyingEndValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndValue reads a UnderlyingEndValue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingEndValue(f *field.UnderlyingEndValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingStips is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NoUnderlyingStips() (*field.NoUnderlyingStipsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingStipsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingStips reads a NoUnderlyingStips from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNoUnderlyingStips(f *field.NoUnderlyingStipsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAllocationPercent is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercentField, errors.MessageRejectError) {
	f := &field.UnderlyingAllocationPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAllocationPercent reads a UnderlyingAllocationPercent from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingAllocationPercent(f *field.UnderlyingAllocationPercentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlementType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSettlementType() (*field.UnderlyingSettlementTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingSettlementTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlementType reads a UnderlyingSettlementType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSettlementType(f *field.UnderlyingSettlementTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashAmount is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCashAmount() (*field.UnderlyingCashAmountField, errors.MessageRejectError) {
	f := &field.UnderlyingCashAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashAmount reads a UnderlyingCashAmount from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCashAmount(f *field.UnderlyingCashAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCashType() (*field.UnderlyingCashTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingCashTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashType reads a UnderlyingCashType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCashType(f *field.UnderlyingCashTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasure is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnderlyingUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasure reads a UnderlyingUnitOfMeasure from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingUnitOfMeasure(f *field.UnderlyingUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTimeUnit is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingTimeUnit() (*field.UnderlyingTimeUnitField, errors.MessageRejectError) {
	f := &field.UnderlyingTimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTimeUnit reads a UnderlyingTimeUnit from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingTimeUnit(f *field.UnderlyingTimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCapValue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCapValue() (*field.UnderlyingCapValueField, errors.MessageRejectError) {
	f := &field.UnderlyingCapValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCapValue reads a UnderlyingCapValue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCapValue(f *field.UnderlyingCapValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUndlyInstrumentParties is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoUndlyInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUndlyInstrumentParties reads a NoUndlyInstrumentParties from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNoUndlyInstrumentParties(f *field.NoUndlyInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlMethod is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSettlMethod() (*field.UnderlyingSettlMethodField, errors.MessageRejectError) {
	f := &field.UnderlyingSettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlMethod reads a UnderlyingSettlMethod from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSettlMethod(f *field.UnderlyingSettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAdjustedQuantity is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantityField, errors.MessageRejectError) {
	f := &field.UnderlyingAdjustedQuantityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAdjustedQuantity reads a UnderlyingAdjustedQuantity from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingAdjustedQuantity(f *field.UnderlyingAdjustedQuantityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingFXRate() (*field.UnderlyingFXRateField, errors.MessageRejectError) {
	f := &field.UnderlyingFXRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRate reads a UnderlyingFXRate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingFXRate(f *field.UnderlyingFXRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRateCalc is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalcField, errors.MessageRejectError) {
	f := &field.UnderlyingFXRateCalcField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRateCalc reads a UnderlyingFXRateCalc from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingFXRateCalc(f *field.UnderlyingFXRateCalcField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityTime is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingMaturityTime() (*field.UnderlyingMaturityTimeField, errors.MessageRejectError) {
	f := &field.UnderlyingMaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityTime reads a UnderlyingMaturityTime from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingMaturityTime(f *field.UnderlyingMaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPutOrCall is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingPutOrCall() (*field.UnderlyingPutOrCallField, errors.MessageRejectError) {
	f := &field.UnderlyingPutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPutOrCall reads a UnderlyingPutOrCall from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingPutOrCall(f *field.UnderlyingPutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingExerciseStyle is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingExerciseStyle() (*field.UnderlyingExerciseStyleField, errors.MessageRejectError) {
	f := &field.UnderlyingExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingExerciseStyle reads a UnderlyingExerciseStyle from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingExerciseStyle(f *field.UnderlyingExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasureQty is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingUnitOfMeasureQty() (*field.UnderlyingUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasureQty reads a UnderlyingUnitOfMeasureQty from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingUnitOfMeasureQty(f *field.UnderlyingUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasure is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingPriceUnitOfMeasure() (*field.UnderlyingPriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnderlyingPriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasure reads a UnderlyingPriceUnitOfMeasure from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingPriceUnitOfMeasure(f *field.UnderlyingPriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingPriceUnitOfMeasureQty() (*field.UnderlyingPriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingPriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasureQty reads a UnderlyingPriceUnitOfMeasureQty from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingPriceUnitOfMeasureQty(f *field.UnderlyingPriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) MarketID() (*field.MarketIDField, errors.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetMarketID(f *field.MarketIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) MarketSegmentID() (*field.MarketSegmentIDField, errors.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetMarketSegmentID(f *field.MarketSegmentIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}
