package fix43

import (
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
	quoteid field.QuoteID) QuoteBuilder {
	var builder QuoteBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(quoteid)
	return builder
}

//QuoteReqID is a non-required field for Quote.
func (m Quote) QuoteReqID() (field.QuoteReqID, error) {
	var f field.QuoteReqID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a required field for Quote.
func (m Quote) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteType is a non-required field for Quote.
func (m Quote) QuoteType() (field.QuoteType, error) {
	var f field.QuoteType
	err := m.Body.Get(&f)
	return f, err
}

//QuoteResponseLevel is a non-required field for Quote.
func (m Quote) QuoteResponseLevel() (field.QuoteResponseLevel, error) {
	var f field.QuoteResponseLevel
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for Quote.
func (m Quote) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for Quote.
func (m Quote) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for Quote.
func (m Quote) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for Quote.
func (m Quote) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for Quote.
func (m Quote) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for Quote.
func (m Quote) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for Quote.
func (m Quote) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for Quote.
func (m Quote) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for Quote.
func (m Quote) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for Quote.
func (m Quote) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for Quote.
func (m Quote) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for Quote.
func (m Quote) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for Quote.
func (m Quote) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for Quote.
func (m Quote) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for Quote.
func (m Quote) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for Quote.
func (m Quote) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for Quote.
func (m Quote) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for Quote.
func (m Quote) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for Quote.
func (m Quote) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for Quote.
func (m Quote) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for Quote.
func (m Quote) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for Quote.
func (m Quote) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for Quote.
func (m Quote) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for Quote.
func (m Quote) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for Quote.
func (m Quote) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for Quote.
func (m Quote) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for Quote.
func (m Quote) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for Quote.
func (m Quote) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for Quote.
func (m Quote) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for Quote.
func (m Quote) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for Quote.
func (m Quote) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for Quote.
func (m Quote) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for Quote.
func (m Quote) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for Quote.
func (m Quote) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for Quote.
func (m Quote) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for Quote.
func (m Quote) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for Quote.
func (m Quote) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for Quote.
func (m Quote) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//BidPx is a non-required field for Quote.
func (m Quote) BidPx() (field.BidPx, error) {
	var f field.BidPx
	err := m.Body.Get(&f)
	return f, err
}

//OfferPx is a non-required field for Quote.
func (m Quote) OfferPx() (field.OfferPx, error) {
	var f field.OfferPx
	err := m.Body.Get(&f)
	return f, err
}

//MktBidPx is a non-required field for Quote.
func (m Quote) MktBidPx() (field.MktBidPx, error) {
	var f field.MktBidPx
	err := m.Body.Get(&f)
	return f, err
}

//MktOfferPx is a non-required field for Quote.
func (m Quote) MktOfferPx() (field.MktOfferPx, error) {
	var f field.MktOfferPx
	err := m.Body.Get(&f)
	return f, err
}

//MinBidSize is a non-required field for Quote.
func (m Quote) MinBidSize() (field.MinBidSize, error) {
	var f field.MinBidSize
	err := m.Body.Get(&f)
	return f, err
}

//BidSize is a non-required field for Quote.
func (m Quote) BidSize() (field.BidSize, error) {
	var f field.BidSize
	err := m.Body.Get(&f)
	return f, err
}

//MinOfferSize is a non-required field for Quote.
func (m Quote) MinOfferSize() (field.MinOfferSize, error) {
	var f field.MinOfferSize
	err := m.Body.Get(&f)
	return f, err
}

//OfferSize is a non-required field for Quote.
func (m Quote) OfferSize() (field.OfferSize, error) {
	var f field.OfferSize
	err := m.Body.Get(&f)
	return f, err
}

//ValidUntilTime is a non-required field for Quote.
func (m Quote) ValidUntilTime() (field.ValidUntilTime, error) {
	var f field.ValidUntilTime
	err := m.Body.Get(&f)
	return f, err
}

//BidSpotRate is a non-required field for Quote.
func (m Quote) BidSpotRate() (field.BidSpotRate, error) {
	var f field.BidSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//OfferSpotRate is a non-required field for Quote.
func (m Quote) OfferSpotRate() (field.OfferSpotRate, error) {
	var f field.OfferSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//BidForwardPoints is a non-required field for Quote.
func (m Quote) BidForwardPoints() (field.BidForwardPoints, error) {
	var f field.BidForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//OfferForwardPoints is a non-required field for Quote.
func (m Quote) OfferForwardPoints() (field.OfferForwardPoints, error) {
	var f field.OfferForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//MidPx is a non-required field for Quote.
func (m Quote) MidPx() (field.MidPx, error) {
	var f field.MidPx
	err := m.Body.Get(&f)
	return f, err
}

//BidYield is a non-required field for Quote.
func (m Quote) BidYield() (field.BidYield, error) {
	var f field.BidYield
	err := m.Body.Get(&f)
	return f, err
}

//MidYield is a non-required field for Quote.
func (m Quote) MidYield() (field.MidYield, error) {
	var f field.MidYield
	err := m.Body.Get(&f)
	return f, err
}

//OfferYield is a non-required field for Quote.
func (m Quote) OfferYield() (field.OfferYield, error) {
	var f field.OfferYield
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for Quote.
func (m Quote) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//SettlmntTyp is a non-required field for Quote.
func (m Quote) SettlmntTyp() (field.SettlmntTyp, error) {
	var f field.SettlmntTyp
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate is a non-required field for Quote.
func (m Quote) FutSettDate() (field.FutSettDate, error) {
	var f field.FutSettDate
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a non-required field for Quote.
func (m Quote) OrdType() (field.OrdType, error) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate2 is a non-required field for Quote.
func (m Quote) FutSettDate2() (field.FutSettDate2, error) {
	var f field.FutSettDate2
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty2 is a non-required field for Quote.
func (m Quote) OrderQty2() (field.OrderQty2, error) {
	var f field.OrderQty2
	err := m.Body.Get(&f)
	return f, err
}

//BidForwardPoints2 is a non-required field for Quote.
func (m Quote) BidForwardPoints2() (field.BidForwardPoints2, error) {
	var f field.BidForwardPoints2
	err := m.Body.Get(&f)
	return f, err
}

//OfferForwardPoints2 is a non-required field for Quote.
func (m Quote) OfferForwardPoints2() (field.OfferForwardPoints2, error) {
	var f field.OfferForwardPoints2
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for Quote.
func (m Quote) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrBidFxRate is a non-required field for Quote.
func (m Quote) SettlCurrBidFxRate() (field.SettlCurrBidFxRate, error) {
	var f field.SettlCurrBidFxRate
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrOfferFxRate is a non-required field for Quote.
func (m Quote) SettlCurrOfferFxRate() (field.SettlCurrOfferFxRate, error) {
	var f field.SettlCurrOfferFxRate
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrFxRateCalc is a non-required field for Quote.
func (m Quote) SettlCurrFxRateCalc() (field.SettlCurrFxRateCalc, error) {
	var f field.SettlCurrFxRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for Quote.
func (m Quote) Commission() (field.Commission, error) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for Quote.
func (m Quote) CommType() (field.CommType, error) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderCapacity is a non-required field for Quote.
func (m Quote) CustOrderCapacity() (field.CustOrderCapacity, error) {
	var f field.CustOrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//ExDestination is a non-required field for Quote.
func (m Quote) ExDestination() (field.ExDestination, error) {
	var f field.ExDestination
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for Quote.
func (m Quote) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for Quote.
func (m Quote) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for Quote.
func (m Quote) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
