package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//IOI msg type = 6.
type IOI struct {
	message.Message
}

//IOIBuilder builds IOI messages.
type IOIBuilder struct {
	message.MessageBuilder
}

//CreateIOIBuilder returns an initialized IOIBuilder with specified required fields.
func CreateIOIBuilder(
	ioiid *field.IOIidField,
	ioitranstype *field.IOITransTypeField,
	side *field.SideField,
	ioiqty *field.IOIQtyField) IOIBuilder {
	var builder IOIBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header.Set(field.NewMsgType("6"))
	builder.Body.Set(ioiid)
	builder.Body.Set(ioitranstype)
	builder.Body.Set(side)
	builder.Body.Set(ioiqty)
	return builder
}

//IOIid is a required field for IOI.
func (m IOI) IOIid() (*field.IOIidField, errors.MessageRejectError) {
	f := &field.IOIidField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIid reads a IOIid from IOI.
func (m IOI) GetIOIid(f *field.IOIidField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOITransType is a required field for IOI.
func (m IOI) IOITransType() (*field.IOITransTypeField, errors.MessageRejectError) {
	f := &field.IOITransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOITransType reads a IOITransType from IOI.
func (m IOI) GetIOITransType(f *field.IOITransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIRefID is a non-required field for IOI.
func (m IOI) IOIRefID() (*field.IOIRefIDField, errors.MessageRejectError) {
	f := &field.IOIRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIRefID reads a IOIRefID from IOI.
func (m IOI) GetIOIRefID(f *field.IOIRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for IOI.
func (m IOI) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from IOI.
func (m IOI) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for IOI.
func (m IOI) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from IOI.
func (m IOI) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for IOI.
func (m IOI) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from IOI.
func (m IOI) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for IOI.
func (m IOI) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from IOI.
func (m IOI) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for IOI.
func (m IOI) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from IOI.
func (m IOI) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for IOI.
func (m IOI) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from IOI.
func (m IOI) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for IOI.
func (m IOI) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from IOI.
func (m IOI) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for IOI.
func (m IOI) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from IOI.
func (m IOI) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for IOI.
func (m IOI) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from IOI.
func (m IOI) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for IOI.
func (m IOI) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from IOI.
func (m IOI) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for IOI.
func (m IOI) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from IOI.
func (m IOI) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for IOI.
func (m IOI) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from IOI.
func (m IOI) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for IOI.
func (m IOI) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from IOI.
func (m IOI) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for IOI.
func (m IOI) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from IOI.
func (m IOI) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for IOI.
func (m IOI) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from IOI.
func (m IOI) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for IOI.
func (m IOI) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from IOI.
func (m IOI) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for IOI.
func (m IOI) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from IOI.
func (m IOI) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for IOI.
func (m IOI) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from IOI.
func (m IOI) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for IOI.
func (m IOI) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from IOI.
func (m IOI) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for IOI.
func (m IOI) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from IOI.
func (m IOI) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for IOI.
func (m IOI) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from IOI.
func (m IOI) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for IOI.
func (m IOI) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from IOI.
func (m IOI) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for IOI.
func (m IOI) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from IOI.
func (m IOI) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for IOI.
func (m IOI) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from IOI.
func (m IOI) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for IOI.
func (m IOI) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from IOI.
func (m IOI) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for IOI.
func (m IOI) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from IOI.
func (m IOI) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for IOI.
func (m IOI) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from IOI.
func (m IOI) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for IOI.
func (m IOI) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from IOI.
func (m IOI) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for IOI.
func (m IOI) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from IOI.
func (m IOI) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for IOI.
func (m IOI) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from IOI.
func (m IOI) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for IOI.
func (m IOI) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from IOI.
func (m IOI) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for IOI.
func (m IOI) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from IOI.
func (m IOI) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for IOI.
func (m IOI) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from IOI.
func (m IOI) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for IOI.
func (m IOI) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from IOI.
func (m IOI) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuantityType is a non-required field for IOI.
func (m IOI) QuantityType() (*field.QuantityTypeField, errors.MessageRejectError) {
	f := &field.QuantityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuantityType reads a QuantityType from IOI.
func (m IOI) GetQuantityType(f *field.QuantityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIQty is a required field for IOI.
func (m IOI) IOIQty() (*field.IOIQtyField, errors.MessageRejectError) {
	f := &field.IOIQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIQty reads a IOIQty from IOI.
func (m IOI) GetIOIQty(f *field.IOIQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for IOI.
func (m IOI) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from IOI.
func (m IOI) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for IOI.
func (m IOI) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from IOI.
func (m IOI) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for IOI.
func (m IOI) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from IOI.
func (m IOI) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValidUntilTime is a non-required field for IOI.
func (m IOI) ValidUntilTime() (*field.ValidUntilTimeField, errors.MessageRejectError) {
	f := &field.ValidUntilTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetValidUntilTime reads a ValidUntilTime from IOI.
func (m IOI) GetValidUntilTime(f *field.ValidUntilTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIQltyInd is a non-required field for IOI.
func (m IOI) IOIQltyInd() (*field.IOIQltyIndField, errors.MessageRejectError) {
	f := &field.IOIQltyIndField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIQltyInd reads a IOIQltyInd from IOI.
func (m IOI) GetIOIQltyInd(f *field.IOIQltyIndField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOINaturalFlag is a non-required field for IOI.
func (m IOI) IOINaturalFlag() (*field.IOINaturalFlagField, errors.MessageRejectError) {
	f := &field.IOINaturalFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOINaturalFlag reads a IOINaturalFlag from IOI.
func (m IOI) GetIOINaturalFlag(f *field.IOINaturalFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoIOIQualifiers is a non-required field for IOI.
func (m IOI) NoIOIQualifiers() (*field.NoIOIQualifiersField, errors.MessageRejectError) {
	f := &field.NoIOIQualifiersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoIOIQualifiers reads a NoIOIQualifiers from IOI.
func (m IOI) GetNoIOIQualifiers(f *field.NoIOIQualifiersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for IOI.
func (m IOI) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from IOI.
func (m IOI) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for IOI.
func (m IOI) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from IOI.
func (m IOI) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for IOI.
func (m IOI) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from IOI.
func (m IOI) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for IOI.
func (m IOI) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from IOI.
func (m IOI) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//URLLink is a non-required field for IOI.
func (m IOI) URLLink() (*field.URLLinkField, errors.MessageRejectError) {
	f := &field.URLLinkField{}
	err := m.Body.Get(f)
	return f, err
}

//GetURLLink reads a URLLink from IOI.
func (m IOI) GetURLLink(f *field.URLLinkField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRoutingIDs is a non-required field for IOI.
func (m IOI) NoRoutingIDs() (*field.NoRoutingIDsField, errors.MessageRejectError) {
	f := &field.NoRoutingIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRoutingIDs reads a NoRoutingIDs from IOI.
func (m IOI) GetNoRoutingIDs(f *field.NoRoutingIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for IOI.
func (m IOI) Spread() (*field.SpreadField, errors.MessageRejectError) {
	f := &field.SpreadField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from IOI.
func (m IOI) GetSpread(f *field.SpreadField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for IOI.
func (m IOI) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrencyField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from IOI.
func (m IOI) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for IOI.
func (m IOI) BenchmarkCurveName() (*field.BenchmarkCurveNameField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from IOI.
func (m IOI) GetBenchmarkCurveName(f *field.BenchmarkCurveNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for IOI.
func (m IOI) BenchmarkCurvePoint() (*field.BenchmarkCurvePointField, errors.MessageRejectError) {
	f := &field.BenchmarkCurvePointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from IOI.
func (m IOI) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Benchmark is a non-required field for IOI.
func (m IOI) Benchmark() (*field.BenchmarkField, errors.MessageRejectError) {
	f := &field.BenchmarkField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmark reads a Benchmark from IOI.
func (m IOI) GetBenchmark(f *field.BenchmarkField) errors.MessageRejectError {
	return m.Body.Get(f)
}
