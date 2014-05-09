package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Quote msg type = S.
type Quote struct {
	message.Message
}

//QuoteBuilder builds Quote messages.
type QuoteBuilder struct {
	message.MessageBuilder
}

//CreateQuoteBuilder returns an initialized QuoteBuilder with specified required fields.
func CreateQuoteBuilder(
	quoteid *field.QuoteIDField) QuoteBuilder {
	var builder QuoteBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header.Set(field.NewMsgType("S"))
	builder.Body.Set(quoteid)
	return builder
}

//QuoteReqID is a non-required field for Quote.
func (m Quote) QuoteReqID() (*field.QuoteReqIDField, errors.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from Quote.
func (m Quote) GetQuoteReqID(f *field.QuoteReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a required field for Quote.
func (m Quote) QuoteID() (*field.QuoteIDField, errors.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from Quote.
func (m Quote) GetQuoteID(f *field.QuoteIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteType is a non-required field for Quote.
func (m Quote) QuoteType() (*field.QuoteTypeField, errors.MessageRejectError) {
	f := &field.QuoteTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteType reads a QuoteType from Quote.
func (m Quote) GetQuoteType(f *field.QuoteTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for Quote.
func (m Quote) QuoteResponseLevel() (*field.QuoteResponseLevelField, errors.MessageRejectError) {
	f := &field.QuoteResponseLevelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from Quote.
func (m Quote) GetQuoteResponseLevel(f *field.QuoteResponseLevelField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for Quote.
func (m Quote) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from Quote.
func (m Quote) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for Quote.
func (m Quote) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from Quote.
func (m Quote) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for Quote.
func (m Quote) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from Quote.
func (m Quote) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for Quote.
func (m Quote) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from Quote.
func (m Quote) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for Quote.
func (m Quote) TradingSessionSubID() (*field.TradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from Quote.
func (m Quote) GetTradingSessionSubID(f *field.TradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for Quote.
func (m Quote) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from Quote.
func (m Quote) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for Quote.
func (m Quote) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from Quote.
func (m Quote) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for Quote.
func (m Quote) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from Quote.
func (m Quote) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for Quote.
func (m Quote) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from Quote.
func (m Quote) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for Quote.
func (m Quote) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from Quote.
func (m Quote) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for Quote.
func (m Quote) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from Quote.
func (m Quote) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for Quote.
func (m Quote) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from Quote.
func (m Quote) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for Quote.
func (m Quote) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from Quote.
func (m Quote) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for Quote.
func (m Quote) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from Quote.
func (m Quote) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for Quote.
func (m Quote) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from Quote.
func (m Quote) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for Quote.
func (m Quote) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from Quote.
func (m Quote) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for Quote.
func (m Quote) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from Quote.
func (m Quote) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for Quote.
func (m Quote) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from Quote.
func (m Quote) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for Quote.
func (m Quote) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from Quote.
func (m Quote) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for Quote.
func (m Quote) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from Quote.
func (m Quote) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for Quote.
func (m Quote) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from Quote.
func (m Quote) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for Quote.
func (m Quote) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from Quote.
func (m Quote) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for Quote.
func (m Quote) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from Quote.
func (m Quote) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for Quote.
func (m Quote) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from Quote.
func (m Quote) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for Quote.
func (m Quote) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from Quote.
func (m Quote) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for Quote.
func (m Quote) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from Quote.
func (m Quote) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for Quote.
func (m Quote) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from Quote.
func (m Quote) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for Quote.
func (m Quote) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from Quote.
func (m Quote) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for Quote.
func (m Quote) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from Quote.
func (m Quote) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for Quote.
func (m Quote) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from Quote.
func (m Quote) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for Quote.
func (m Quote) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from Quote.
func (m Quote) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for Quote.
func (m Quote) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from Quote.
func (m Quote) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for Quote.
func (m Quote) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from Quote.
func (m Quote) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for Quote.
func (m Quote) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from Quote.
func (m Quote) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for Quote.
func (m Quote) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from Quote.
func (m Quote) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for Quote.
func (m Quote) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from Quote.
func (m Quote) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for Quote.
func (m Quote) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from Quote.
func (m Quote) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for Quote.
func (m Quote) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from Quote.
func (m Quote) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidPx is a non-required field for Quote.
func (m Quote) BidPx() (*field.BidPxField, errors.MessageRejectError) {
	f := &field.BidPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidPx reads a BidPx from Quote.
func (m Quote) GetBidPx(f *field.BidPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferPx is a non-required field for Quote.
func (m Quote) OfferPx() (*field.OfferPxField, errors.MessageRejectError) {
	f := &field.OfferPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOfferPx reads a OfferPx from Quote.
func (m Quote) GetOfferPx(f *field.OfferPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MktBidPx is a non-required field for Quote.
func (m Quote) MktBidPx() (*field.MktBidPxField, errors.MessageRejectError) {
	f := &field.MktBidPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMktBidPx reads a MktBidPx from Quote.
func (m Quote) GetMktBidPx(f *field.MktBidPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MktOfferPx is a non-required field for Quote.
func (m Quote) MktOfferPx() (*field.MktOfferPxField, errors.MessageRejectError) {
	f := &field.MktOfferPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMktOfferPx reads a MktOfferPx from Quote.
func (m Quote) GetMktOfferPx(f *field.MktOfferPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinBidSize is a non-required field for Quote.
func (m Quote) MinBidSize() (*field.MinBidSizeField, errors.MessageRejectError) {
	f := &field.MinBidSizeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinBidSize reads a MinBidSize from Quote.
func (m Quote) GetMinBidSize(f *field.MinBidSizeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidSize is a non-required field for Quote.
func (m Quote) BidSize() (*field.BidSizeField, errors.MessageRejectError) {
	f := &field.BidSizeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidSize reads a BidSize from Quote.
func (m Quote) GetBidSize(f *field.BidSizeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinOfferSize is a non-required field for Quote.
func (m Quote) MinOfferSize() (*field.MinOfferSizeField, errors.MessageRejectError) {
	f := &field.MinOfferSizeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinOfferSize reads a MinOfferSize from Quote.
func (m Quote) GetMinOfferSize(f *field.MinOfferSizeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferSize is a non-required field for Quote.
func (m Quote) OfferSize() (*field.OfferSizeField, errors.MessageRejectError) {
	f := &field.OfferSizeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOfferSize reads a OfferSize from Quote.
func (m Quote) GetOfferSize(f *field.OfferSizeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValidUntilTime is a non-required field for Quote.
func (m Quote) ValidUntilTime() (*field.ValidUntilTimeField, errors.MessageRejectError) {
	f := &field.ValidUntilTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetValidUntilTime reads a ValidUntilTime from Quote.
func (m Quote) GetValidUntilTime(f *field.ValidUntilTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidSpotRate is a non-required field for Quote.
func (m Quote) BidSpotRate() (*field.BidSpotRateField, errors.MessageRejectError) {
	f := &field.BidSpotRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidSpotRate reads a BidSpotRate from Quote.
func (m Quote) GetBidSpotRate(f *field.BidSpotRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferSpotRate is a non-required field for Quote.
func (m Quote) OfferSpotRate() (*field.OfferSpotRateField, errors.MessageRejectError) {
	f := &field.OfferSpotRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOfferSpotRate reads a OfferSpotRate from Quote.
func (m Quote) GetOfferSpotRate(f *field.OfferSpotRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidForwardPoints is a non-required field for Quote.
func (m Quote) BidForwardPoints() (*field.BidForwardPointsField, errors.MessageRejectError) {
	f := &field.BidForwardPointsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidForwardPoints reads a BidForwardPoints from Quote.
func (m Quote) GetBidForwardPoints(f *field.BidForwardPointsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferForwardPoints is a non-required field for Quote.
func (m Quote) OfferForwardPoints() (*field.OfferForwardPointsField, errors.MessageRejectError) {
	f := &field.OfferForwardPointsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOfferForwardPoints reads a OfferForwardPoints from Quote.
func (m Quote) GetOfferForwardPoints(f *field.OfferForwardPointsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MidPx is a non-required field for Quote.
func (m Quote) MidPx() (*field.MidPxField, errors.MessageRejectError) {
	f := &field.MidPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMidPx reads a MidPx from Quote.
func (m Quote) GetMidPx(f *field.MidPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidYield is a non-required field for Quote.
func (m Quote) BidYield() (*field.BidYieldField, errors.MessageRejectError) {
	f := &field.BidYieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidYield reads a BidYield from Quote.
func (m Quote) GetBidYield(f *field.BidYieldField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MidYield is a non-required field for Quote.
func (m Quote) MidYield() (*field.MidYieldField, errors.MessageRejectError) {
	f := &field.MidYieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMidYield reads a MidYield from Quote.
func (m Quote) GetMidYield(f *field.MidYieldField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferYield is a non-required field for Quote.
func (m Quote) OfferYield() (*field.OfferYieldField, errors.MessageRejectError) {
	f := &field.OfferYieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOfferYield reads a OfferYield from Quote.
func (m Quote) GetOfferYield(f *field.OfferYieldField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for Quote.
func (m Quote) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from Quote.
func (m Quote) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for Quote.
func (m Quote) SettlmntTyp() (*field.SettlmntTypField, errors.MessageRejectError) {
	f := &field.SettlmntTypField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from Quote.
func (m Quote) GetSettlmntTyp(f *field.SettlmntTypField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for Quote.
func (m Quote) FutSettDate() (*field.FutSettDateField, errors.MessageRejectError) {
	f := &field.FutSettDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from Quote.
func (m Quote) GetFutSettDate(f *field.FutSettDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a non-required field for Quote.
func (m Quote) OrdType() (*field.OrdTypeField, errors.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from Quote.
func (m Quote) GetOrdType(f *field.OrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate2 is a non-required field for Quote.
func (m Quote) FutSettDate2() (*field.FutSettDate2Field, errors.MessageRejectError) {
	f := &field.FutSettDate2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate2 reads a FutSettDate2 from Quote.
func (m Quote) GetFutSettDate2(f *field.FutSettDate2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty2 is a non-required field for Quote.
func (m Quote) OrderQty2() (*field.OrderQty2Field, errors.MessageRejectError) {
	f := &field.OrderQty2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty2 reads a OrderQty2 from Quote.
func (m Quote) GetOrderQty2(f *field.OrderQty2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidForwardPoints2 is a non-required field for Quote.
func (m Quote) BidForwardPoints2() (*field.BidForwardPoints2Field, errors.MessageRejectError) {
	f := &field.BidForwardPoints2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidForwardPoints2 reads a BidForwardPoints2 from Quote.
func (m Quote) GetBidForwardPoints2(f *field.BidForwardPoints2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferForwardPoints2 is a non-required field for Quote.
func (m Quote) OfferForwardPoints2() (*field.OfferForwardPoints2Field, errors.MessageRejectError) {
	f := &field.OfferForwardPoints2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetOfferForwardPoints2 reads a OfferForwardPoints2 from Quote.
func (m Quote) GetOfferForwardPoints2(f *field.OfferForwardPoints2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for Quote.
func (m Quote) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from Quote.
func (m Quote) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrBidFxRate is a non-required field for Quote.
func (m Quote) SettlCurrBidFxRate() (*field.SettlCurrBidFxRateField, errors.MessageRejectError) {
	f := &field.SettlCurrBidFxRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrBidFxRate reads a SettlCurrBidFxRate from Quote.
func (m Quote) GetSettlCurrBidFxRate(f *field.SettlCurrBidFxRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrOfferFxRate is a non-required field for Quote.
func (m Quote) SettlCurrOfferFxRate() (*field.SettlCurrOfferFxRateField, errors.MessageRejectError) {
	f := &field.SettlCurrOfferFxRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrOfferFxRate reads a SettlCurrOfferFxRate from Quote.
func (m Quote) GetSettlCurrOfferFxRate(f *field.SettlCurrOfferFxRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRateCalc is a non-required field for Quote.
func (m Quote) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalcField, errors.MessageRejectError) {
	f := &field.SettlCurrFxRateCalcField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRateCalc reads a SettlCurrFxRateCalc from Quote.
func (m Quote) GetSettlCurrFxRateCalc(f *field.SettlCurrFxRateCalcField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for Quote.
func (m Quote) Commission() (*field.CommissionField, errors.MessageRejectError) {
	f := &field.CommissionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from Quote.
func (m Quote) GetCommission(f *field.CommissionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for Quote.
func (m Quote) CommType() (*field.CommTypeField, errors.MessageRejectError) {
	f := &field.CommTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from Quote.
func (m Quote) GetCommType(f *field.CommTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for Quote.
func (m Quote) CustOrderCapacity() (*field.CustOrderCapacityField, errors.MessageRejectError) {
	f := &field.CustOrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from Quote.
func (m Quote) GetCustOrderCapacity(f *field.CustOrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for Quote.
func (m Quote) ExDestination() (*field.ExDestinationField, errors.MessageRejectError) {
	f := &field.ExDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from Quote.
func (m Quote) GetExDestination(f *field.ExDestinationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for Quote.
func (m Quote) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from Quote.
func (m Quote) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for Quote.
func (m Quote) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from Quote.
func (m Quote) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for Quote.
func (m Quote) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from Quote.
func (m Quote) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
