package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//QuoteStatusReport msg type = AI.
type QuoteStatusReport struct {
	message.Message
}

//QuoteStatusReportBuilder builds QuoteStatusReport messages.
type QuoteStatusReportBuilder struct {
	message.MessageBuilder
}

//CreateQuoteStatusReportBuilder returns an initialized QuoteStatusReportBuilder with specified required fields.
func CreateQuoteStatusReportBuilder(
	quoteid field.QuoteID) QuoteStatusReportBuilder {
	var builder QuoteStatusReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("AI"))
	builder.Body.Set(quoteid)
	return builder
}

//QuoteStatusReqID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteStatusReqID() (*field.QuoteStatusReqID, errors.MessageRejectError) {
	f := new(field.QuoteStatusReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteStatusReqID reads a QuoteStatusReqID from QuoteStatusReport.
func (m QuoteStatusReport) GetQuoteStatusReqID(f *field.QuoteStatusReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteReqID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteReqID() (*field.QuoteReqID, errors.MessageRejectError) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteStatusReport.
func (m QuoteStatusReport) GetQuoteReqID(f *field.QuoteReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteID() (*field.QuoteID, errors.MessageRejectError) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from QuoteStatusReport.
func (m QuoteStatusReport) GetQuoteID(f *field.QuoteID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteType() (*field.QuoteType, errors.MessageRejectError) {
	f := new(field.QuoteType)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteType reads a QuoteType from QuoteStatusReport.
func (m QuoteStatusReport) GetQuoteType(f *field.QuoteType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from QuoteStatusReport.
func (m QuoteStatusReport) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from QuoteStatusReport.
func (m QuoteStatusReport) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from QuoteStatusReport.
func (m QuoteStatusReport) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from QuoteStatusReport.
func (m QuoteStatusReport) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from QuoteStatusReport.
func (m QuoteStatusReport) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from QuoteStatusReport.
func (m QuoteStatusReport) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from QuoteStatusReport.
func (m QuoteStatusReport) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from QuoteStatusReport.
func (m QuoteStatusReport) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from QuoteStatusReport.
func (m QuoteStatusReport) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from QuoteStatusReport.
func (m QuoteStatusReport) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from QuoteStatusReport.
func (m QuoteStatusReport) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from QuoteStatusReport.
func (m QuoteStatusReport) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from QuoteStatusReport.
func (m QuoteStatusReport) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from QuoteStatusReport.
func (m QuoteStatusReport) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from QuoteStatusReport.
func (m QuoteStatusReport) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from QuoteStatusReport.
func (m QuoteStatusReport) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from QuoteStatusReport.
func (m QuoteStatusReport) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from QuoteStatusReport.
func (m QuoteStatusReport) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from QuoteStatusReport.
func (m QuoteStatusReport) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from QuoteStatusReport.
func (m QuoteStatusReport) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from QuoteStatusReport.
func (m QuoteStatusReport) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from QuoteStatusReport.
func (m QuoteStatusReport) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from QuoteStatusReport.
func (m QuoteStatusReport) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from QuoteStatusReport.
func (m QuoteStatusReport) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from QuoteStatusReport.
func (m QuoteStatusReport) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from QuoteStatusReport.
func (m QuoteStatusReport) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from QuoteStatusReport.
func (m QuoteStatusReport) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from QuoteStatusReport.
func (m QuoteStatusReport) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from QuoteStatusReport.
func (m QuoteStatusReport) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from QuoteStatusReport.
func (m QuoteStatusReport) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from QuoteStatusReport.
func (m QuoteStatusReport) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from QuoteStatusReport.
func (m QuoteStatusReport) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from QuoteStatusReport.
func (m QuoteStatusReport) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from QuoteStatusReport.
func (m QuoteStatusReport) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from QuoteStatusReport.
func (m QuoteStatusReport) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from QuoteStatusReport.
func (m QuoteStatusReport) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from QuoteStatusReport.
func (m QuoteStatusReport) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from QuoteStatusReport.
func (m QuoteStatusReport) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidPx() (*field.BidPx, errors.MessageRejectError) {
	f := new(field.BidPx)
	err := m.Body.Get(f)
	return f, err
}

//GetBidPx reads a BidPx from QuoteStatusReport.
func (m QuoteStatusReport) GetBidPx(f *field.BidPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferPx() (*field.OfferPx, errors.MessageRejectError) {
	f := new(field.OfferPx)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferPx reads a OfferPx from QuoteStatusReport.
func (m QuoteStatusReport) GetOfferPx(f *field.OfferPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MktBidPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MktBidPx() (*field.MktBidPx, errors.MessageRejectError) {
	f := new(field.MktBidPx)
	err := m.Body.Get(f)
	return f, err
}

//GetMktBidPx reads a MktBidPx from QuoteStatusReport.
func (m QuoteStatusReport) GetMktBidPx(f *field.MktBidPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MktOfferPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MktOfferPx() (*field.MktOfferPx, errors.MessageRejectError) {
	f := new(field.MktOfferPx)
	err := m.Body.Get(f)
	return f, err
}

//GetMktOfferPx reads a MktOfferPx from QuoteStatusReport.
func (m QuoteStatusReport) GetMktOfferPx(f *field.MktOfferPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinBidSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MinBidSize() (*field.MinBidSize, errors.MessageRejectError) {
	f := new(field.MinBidSize)
	err := m.Body.Get(f)
	return f, err
}

//GetMinBidSize reads a MinBidSize from QuoteStatusReport.
func (m QuoteStatusReport) GetMinBidSize(f *field.MinBidSize) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidSize() (*field.BidSize, errors.MessageRejectError) {
	f := new(field.BidSize)
	err := m.Body.Get(f)
	return f, err
}

//GetBidSize reads a BidSize from QuoteStatusReport.
func (m QuoteStatusReport) GetBidSize(f *field.BidSize) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinOfferSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MinOfferSize() (*field.MinOfferSize, errors.MessageRejectError) {
	f := new(field.MinOfferSize)
	err := m.Body.Get(f)
	return f, err
}

//GetMinOfferSize reads a MinOfferSize from QuoteStatusReport.
func (m QuoteStatusReport) GetMinOfferSize(f *field.MinOfferSize) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferSize() (*field.OfferSize, errors.MessageRejectError) {
	f := new(field.OfferSize)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferSize reads a OfferSize from QuoteStatusReport.
func (m QuoteStatusReport) GetOfferSize(f *field.OfferSize) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValidUntilTime is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ValidUntilTime() (*field.ValidUntilTime, errors.MessageRejectError) {
	f := new(field.ValidUntilTime)
	err := m.Body.Get(f)
	return f, err
}

//GetValidUntilTime reads a ValidUntilTime from QuoteStatusReport.
func (m QuoteStatusReport) GetValidUntilTime(f *field.ValidUntilTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidSpotRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidSpotRate() (*field.BidSpotRate, errors.MessageRejectError) {
	f := new(field.BidSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//GetBidSpotRate reads a BidSpotRate from QuoteStatusReport.
func (m QuoteStatusReport) GetBidSpotRate(f *field.BidSpotRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferSpotRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferSpotRate() (*field.OfferSpotRate, errors.MessageRejectError) {
	f := new(field.OfferSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferSpotRate reads a OfferSpotRate from QuoteStatusReport.
func (m QuoteStatusReport) GetOfferSpotRate(f *field.OfferSpotRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidForwardPoints is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidForwardPoints() (*field.BidForwardPoints, errors.MessageRejectError) {
	f := new(field.BidForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//GetBidForwardPoints reads a BidForwardPoints from QuoteStatusReport.
func (m QuoteStatusReport) GetBidForwardPoints(f *field.BidForwardPoints) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferForwardPoints is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferForwardPoints() (*field.OfferForwardPoints, errors.MessageRejectError) {
	f := new(field.OfferForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferForwardPoints reads a OfferForwardPoints from QuoteStatusReport.
func (m QuoteStatusReport) GetOfferForwardPoints(f *field.OfferForwardPoints) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MidPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MidPx() (*field.MidPx, errors.MessageRejectError) {
	f := new(field.MidPx)
	err := m.Body.Get(f)
	return f, err
}

//GetMidPx reads a MidPx from QuoteStatusReport.
func (m QuoteStatusReport) GetMidPx(f *field.MidPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidYield is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidYield() (*field.BidYield, errors.MessageRejectError) {
	f := new(field.BidYield)
	err := m.Body.Get(f)
	return f, err
}

//GetBidYield reads a BidYield from QuoteStatusReport.
func (m QuoteStatusReport) GetBidYield(f *field.BidYield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MidYield is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MidYield() (*field.MidYield, errors.MessageRejectError) {
	f := new(field.MidYield)
	err := m.Body.Get(f)
	return f, err
}

//GetMidYield reads a MidYield from QuoteStatusReport.
func (m QuoteStatusReport) GetMidYield(f *field.MidYield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferYield is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferYield() (*field.OfferYield, errors.MessageRejectError) {
	f := new(field.OfferYield)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferYield reads a OfferYield from QuoteStatusReport.
func (m QuoteStatusReport) GetOfferYield(f *field.OfferYield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from QuoteStatusReport.
func (m QuoteStatusReport) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) FutSettDate() (*field.FutSettDate, errors.MessageRejectError) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from QuoteStatusReport.
func (m QuoteStatusReport) GetFutSettDate(f *field.FutSettDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrdType() (*field.OrdType, errors.MessageRejectError) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from QuoteStatusReport.
func (m QuoteStatusReport) GetOrdType(f *field.OrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) FutSettDate2() (*field.FutSettDate2, errors.MessageRejectError) {
	f := new(field.FutSettDate2)
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate2 reads a FutSettDate2 from QuoteStatusReport.
func (m QuoteStatusReport) GetFutSettDate2(f *field.FutSettDate2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrderQty2() (*field.OrderQty2, errors.MessageRejectError) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty2 reads a OrderQty2 from QuoteStatusReport.
func (m QuoteStatusReport) GetOrderQty2(f *field.OrderQty2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidForwardPoints2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidForwardPoints2() (*field.BidForwardPoints2, errors.MessageRejectError) {
	f := new(field.BidForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}

//GetBidForwardPoints2 reads a BidForwardPoints2 from QuoteStatusReport.
func (m QuoteStatusReport) GetBidForwardPoints2(f *field.BidForwardPoints2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferForwardPoints2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferForwardPoints2() (*field.OfferForwardPoints2, errors.MessageRejectError) {
	f := new(field.OfferForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferForwardPoints2 reads a OfferForwardPoints2 from QuoteStatusReport.
func (m QuoteStatusReport) GetOfferForwardPoints2(f *field.OfferForwardPoints2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from QuoteStatusReport.
func (m QuoteStatusReport) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrBidFxRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlCurrBidFxRate() (*field.SettlCurrBidFxRate, errors.MessageRejectError) {
	f := new(field.SettlCurrBidFxRate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrBidFxRate reads a SettlCurrBidFxRate from QuoteStatusReport.
func (m QuoteStatusReport) GetSettlCurrBidFxRate(f *field.SettlCurrBidFxRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrOfferFxRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlCurrOfferFxRate() (*field.SettlCurrOfferFxRate, errors.MessageRejectError) {
	f := new(field.SettlCurrOfferFxRate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrOfferFxRate reads a SettlCurrOfferFxRate from QuoteStatusReport.
func (m QuoteStatusReport) GetSettlCurrOfferFxRate(f *field.SettlCurrOfferFxRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRateCalc is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalc, errors.MessageRejectError) {
	f := new(field.SettlCurrFxRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRateCalc reads a SettlCurrFxRateCalc from QuoteStatusReport.
func (m QuoteStatusReport) GetSettlCurrFxRateCalc(f *field.SettlCurrFxRateCalc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Commission() (*field.Commission, errors.MessageRejectError) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from QuoteStatusReport.
func (m QuoteStatusReport) GetCommission(f *field.Commission) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CommType() (*field.CommType, errors.MessageRejectError) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from QuoteStatusReport.
func (m QuoteStatusReport) GetCommType(f *field.CommType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CustOrderCapacity() (*field.CustOrderCapacity, errors.MessageRejectError) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from QuoteStatusReport.
func (m QuoteStatusReport) GetCustOrderCapacity(f *field.CustOrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ExDestination() (*field.ExDestination, errors.MessageRejectError) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from QuoteStatusReport.
func (m QuoteStatusReport) GetExDestination(f *field.ExDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteStatus is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteStatus() (*field.QuoteStatus, errors.MessageRejectError) {
	f := new(field.QuoteStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteStatus reads a QuoteStatus from QuoteStatusReport.
func (m QuoteStatusReport) GetQuoteStatus(f *field.QuoteStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}
