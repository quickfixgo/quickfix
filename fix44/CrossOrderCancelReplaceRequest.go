package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CrossOrderCancelReplaceRequest msg type = t.
type CrossOrderCancelReplaceRequest struct {
	message.Message
}

//CrossOrderCancelReplaceRequestBuilder builds CrossOrderCancelReplaceRequest messages.
type CrossOrderCancelReplaceRequestBuilder struct {
	message.MessageBuilder
}

//CreateCrossOrderCancelReplaceRequestBuilder returns an initialized CrossOrderCancelReplaceRequestBuilder with specified required fields.
func CreateCrossOrderCancelReplaceRequestBuilder(
	crossid *field.CrossIDField,
	origcrossid *field.OrigCrossIDField,
	crosstype *field.CrossTypeField,
	crossprioritization *field.CrossPrioritizationField,
	nosides *field.NoSidesField,
	transacttime *field.TransactTimeField,
	ordtype *field.OrdTypeField) CrossOrderCancelReplaceRequestBuilder {
	var builder CrossOrderCancelReplaceRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header.Set(field.NewMsgType("t"))
	builder.Body.Set(crossid)
	builder.Body.Set(origcrossid)
	builder.Body.Set(crosstype)
	builder.Body.Set(crossprioritization)
	builder.Body.Set(nosides)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//OrderID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossID is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CrossID() (*field.CrossIDField, errors.MessageRejectError) {
	f := &field.CrossIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossID reads a CrossID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCrossID(f *field.CrossIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigCrossID is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OrigCrossID() (*field.OrigCrossIDField, errors.MessageRejectError) {
	f := &field.OrigCrossIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigCrossID reads a OrigCrossID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetOrigCrossID(f *field.OrigCrossIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossType is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CrossType() (*field.CrossTypeField, errors.MessageRejectError) {
	f := &field.CrossTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossType reads a CrossType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCrossType(f *field.CrossTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossPrioritization is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CrossPrioritization() (*field.CrossPrioritizationField, errors.MessageRejectError) {
	f := &field.CrossPrioritizationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossPrioritization reads a CrossPrioritization from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCrossPrioritization(f *field.CrossPrioritizationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSides is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoSides() (*field.NoSidesField, errors.MessageRejectError) {
	f := &field.NoSidesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSides reads a NoSides from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetNoSides(f *field.NoSidesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SettlType() (*field.SettlTypeField, errors.MessageRejectError) {
	f := &field.SettlTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSettlType(f *field.SettlTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SettlDate() (*field.SettlDateField, errors.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSettlDate(f *field.SettlDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) HandlInst() (*field.HandlInstField, errors.MessageRejectError) {
	f := &field.HandlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetHandlInst(f *field.HandlInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExecInst() (*field.ExecInstField, errors.MessageRejectError) {
	f := &field.ExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetExecInst(f *field.ExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MinQty() (*field.MinQtyField, errors.MessageRejectError) {
	f := &field.MinQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetMinQty(f *field.MinQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaxFloor() (*field.MaxFloorField, errors.MessageRejectError) {
	f := &field.MaxFloorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetMaxFloor(f *field.MaxFloorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExDestination() (*field.ExDestinationField, errors.MessageRejectError) {
	f := &field.ExDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetExDestination(f *field.ExDestinationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoTradingSessions() (*field.NoTradingSessionsField, errors.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetNoTradingSessions(f *field.NoTradingSessionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ProcessCode() (*field.ProcessCodeField, errors.MessageRejectError) {
	f := &field.ProcessCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetProcessCode(f *field.ProcessCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PrevClosePx() (*field.PrevClosePxField, errors.MessageRejectError) {
	f := &field.PrevClosePxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetPrevClosePx(f *field.PrevClosePxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) LocateReqd() (*field.LocateReqdField, errors.MessageRejectError) {
	f := &field.LocateReqdField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetLocateReqd(f *field.LocateReqdField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoStipulations() (*field.NoStipulationsField, errors.MessageRejectError) {
	f := &field.NoStipulationsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetNoStipulations(f *field.NoStipulationsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OrdType() (*field.OrdTypeField, errors.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetOrdType(f *field.OrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StopPx() (*field.StopPxField, errors.MessageRejectError) {
	f := &field.StopPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetStopPx(f *field.StopPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Spread() (*field.SpreadField, errors.MessageRejectError) {
	f := &field.SpreadField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSpread(f *field.SpreadField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrencyField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkCurveName() (*field.BenchmarkCurveNameField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetBenchmarkCurveName(f *field.BenchmarkCurveNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkCurvePoint() (*field.BenchmarkCurvePointField, errors.MessageRejectError) {
	f := &field.BenchmarkCurvePointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkPrice() (*field.BenchmarkPriceField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetBenchmarkPrice(f *field.BenchmarkPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkPriceType() (*field.BenchmarkPriceTypeField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetBenchmarkPriceType(f *field.BenchmarkPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkSecurityID() (*field.BenchmarkSecurityIDField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetBenchmarkSecurityID(f *field.BenchmarkSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldType() (*field.YieldTypeField, errors.MessageRejectError) {
	f := &field.YieldTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetYieldType(f *field.YieldTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Yield() (*field.YieldField, errors.MessageRejectError) {
	f := &field.YieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetYield(f *field.YieldField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldCalcDate() (*field.YieldCalcDateField, errors.MessageRejectError) {
	f := &field.YieldCalcDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetYieldCalcDate(f *field.YieldCalcDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldRedemptionDate() (*field.YieldRedemptionDateField, errors.MessageRejectError) {
	f := &field.YieldRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetYieldRedemptionDate(f *field.YieldRedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldRedemptionPrice() (*field.YieldRedemptionPriceField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetYieldRedemptionPrice(f *field.YieldRedemptionPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldRedemptionPriceType() (*field.YieldRedemptionPriceTypeField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ComplianceID() (*field.ComplianceIDField, errors.MessageRejectError) {
	f := &field.ComplianceIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetComplianceID(f *field.ComplianceIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) IOIID() (*field.IOIIDField, errors.MessageRejectError) {
	f := &field.IOIIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIID reads a IOIID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetIOIID(f *field.IOIIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) QuoteID() (*field.QuoteIDField, errors.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetQuoteID(f *field.QuoteIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TimeInForce() (*field.TimeInForceField, errors.MessageRejectError) {
	f := &field.TimeInForceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetTimeInForce(f *field.TimeInForceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EffectiveTime() (*field.EffectiveTimeField, errors.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetEffectiveTime(f *field.EffectiveTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExpireDate() (*field.ExpireDateField, errors.MessageRejectError) {
	f := &field.ExpireDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetExpireDate(f *field.ExpireDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GTBookingInst() (*field.GTBookingInstField, errors.MessageRejectError) {
	f := &field.GTBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetGTBookingInst(f *field.GTBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaxShow() (*field.MaxShowField, errors.MessageRejectError) {
	f := &field.MaxShowField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetMaxShow(f *field.MaxShowField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetValue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegOffsetValue() (*field.PegOffsetValueField, errors.MessageRejectError) {
	f := &field.PegOffsetValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetValue reads a PegOffsetValue from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetPegOffsetValue(f *field.PegOffsetValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegMoveType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegMoveType() (*field.PegMoveTypeField, errors.MessageRejectError) {
	f := &field.PegMoveTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegMoveType reads a PegMoveType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetPegMoveType(f *field.PegMoveTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegOffsetType() (*field.PegOffsetTypeField, errors.MessageRejectError) {
	f := &field.PegOffsetTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetType reads a PegOffsetType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetPegOffsetType(f *field.PegOffsetTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegLimitType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegLimitType() (*field.PegLimitTypeField, errors.MessageRejectError) {
	f := &field.PegLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegLimitType reads a PegLimitType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetPegLimitType(f *field.PegLimitTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegRoundDirection is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegRoundDirection() (*field.PegRoundDirectionField, errors.MessageRejectError) {
	f := &field.PegRoundDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegRoundDirection reads a PegRoundDirection from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetPegRoundDirection(f *field.PegRoundDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegScope is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegScope() (*field.PegScopeField, errors.MessageRejectError) {
	f := &field.PegScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegScope reads a PegScope from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetPegScope(f *field.PegScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionInst() (*field.DiscretionInstField, errors.MessageRejectError) {
	f := &field.DiscretionInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetDiscretionInst(f *field.DiscretionInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetValue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionOffsetValue() (*field.DiscretionOffsetValueField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetValue reads a DiscretionOffsetValue from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetDiscretionOffsetValue(f *field.DiscretionOffsetValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionMoveType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionMoveType() (*field.DiscretionMoveTypeField, errors.MessageRejectError) {
	f := &field.DiscretionMoveTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionMoveType reads a DiscretionMoveType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetDiscretionMoveType(f *field.DiscretionMoveTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionOffsetType() (*field.DiscretionOffsetTypeField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetType reads a DiscretionOffsetType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetDiscretionOffsetType(f *field.DiscretionOffsetTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionLimitType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionLimitType() (*field.DiscretionLimitTypeField, errors.MessageRejectError) {
	f := &field.DiscretionLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionLimitType reads a DiscretionLimitType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetDiscretionLimitType(f *field.DiscretionLimitTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionRoundDirection is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionRoundDirection() (*field.DiscretionRoundDirectionField, errors.MessageRejectError) {
	f := &field.DiscretionRoundDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionRoundDirection reads a DiscretionRoundDirection from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetDiscretionRoundDirection(f *field.DiscretionRoundDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionScope is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionScope() (*field.DiscretionScopeField, errors.MessageRejectError) {
	f := &field.DiscretionScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionScope reads a DiscretionScope from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetDiscretionScope(f *field.DiscretionScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategy is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TargetStrategy() (*field.TargetStrategyField, errors.MessageRejectError) {
	f := &field.TargetStrategyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategy reads a TargetStrategy from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetTargetStrategy(f *field.TargetStrategyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategyParameters is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TargetStrategyParameters() (*field.TargetStrategyParametersField, errors.MessageRejectError) {
	f := &field.TargetStrategyParametersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategyParameters reads a TargetStrategyParameters from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetTargetStrategyParameters(f *field.TargetStrategyParametersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParticipationRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ParticipationRate() (*field.ParticipationRateField, errors.MessageRejectError) {
	f := &field.ParticipationRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetParticipationRate reads a ParticipationRate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetParticipationRate(f *field.ParticipationRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CancellationRights() (*field.CancellationRightsField, errors.MessageRejectError) {
	f := &field.CancellationRightsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCancellationRights(f *field.CancellationRightsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MoneyLaunderingStatus() (*field.MoneyLaunderingStatusField, errors.MessageRejectError) {
	f := &field.MoneyLaunderingStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RegistID() (*field.RegistIDField, errors.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetRegistID(f *field.RegistIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Designation() (*field.DesignationField, errors.MessageRejectError) {
	f := &field.DesignationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetDesignation(f *field.DesignationField) errors.MessageRejectError {
	return m.Body.Get(f)
}
