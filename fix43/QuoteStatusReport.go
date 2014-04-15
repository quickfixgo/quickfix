package fix43

import (
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
	builder.Body.Set(quoteid)
	return builder
}

//QuoteStatusReqID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteStatusReqID() (field.QuoteStatusReqID, error) {
	var f field.QuoteStatusReqID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteReqID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteReqID() (field.QuoteReqID, error) {
	var f field.QuoteReqID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteType() (field.QuoteType, error) {
	var f field.QuoteType
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//BidPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidPx() (field.BidPx, error) {
	var f field.BidPx
	err := m.Body.Get(&f)
	return f, err
}

//OfferPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferPx() (field.OfferPx, error) {
	var f field.OfferPx
	err := m.Body.Get(&f)
	return f, err
}

//MktBidPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MktBidPx() (field.MktBidPx, error) {
	var f field.MktBidPx
	err := m.Body.Get(&f)
	return f, err
}

//MktOfferPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MktOfferPx() (field.MktOfferPx, error) {
	var f field.MktOfferPx
	err := m.Body.Get(&f)
	return f, err
}

//MinBidSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MinBidSize() (field.MinBidSize, error) {
	var f field.MinBidSize
	err := m.Body.Get(&f)
	return f, err
}

//BidSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidSize() (field.BidSize, error) {
	var f field.BidSize
	err := m.Body.Get(&f)
	return f, err
}

//MinOfferSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MinOfferSize() (field.MinOfferSize, error) {
	var f field.MinOfferSize
	err := m.Body.Get(&f)
	return f, err
}

//OfferSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferSize() (field.OfferSize, error) {
	var f field.OfferSize
	err := m.Body.Get(&f)
	return f, err
}

//ValidUntilTime is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ValidUntilTime() (field.ValidUntilTime, error) {
	var f field.ValidUntilTime
	err := m.Body.Get(&f)
	return f, err
}

//BidSpotRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidSpotRate() (field.BidSpotRate, error) {
	var f field.BidSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//OfferSpotRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferSpotRate() (field.OfferSpotRate, error) {
	var f field.OfferSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//BidForwardPoints is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidForwardPoints() (field.BidForwardPoints, error) {
	var f field.BidForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//OfferForwardPoints is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferForwardPoints() (field.OfferForwardPoints, error) {
	var f field.OfferForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//MidPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MidPx() (field.MidPx, error) {
	var f field.MidPx
	err := m.Body.Get(&f)
	return f, err
}

//BidYield is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidYield() (field.BidYield, error) {
	var f field.BidYield
	err := m.Body.Get(&f)
	return f, err
}

//MidYield is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MidYield() (field.MidYield, error) {
	var f field.MidYield
	err := m.Body.Get(&f)
	return f, err
}

//OfferYield is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferYield() (field.OfferYield, error) {
	var f field.OfferYield
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) FutSettDate() (field.FutSettDate, error) {
	var f field.FutSettDate
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrdType() (field.OrdType, error) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) FutSettDate2() (field.FutSettDate2, error) {
	var f field.FutSettDate2
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrderQty2() (field.OrderQty2, error) {
	var f field.OrderQty2
	err := m.Body.Get(&f)
	return f, err
}

//BidForwardPoints2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidForwardPoints2() (field.BidForwardPoints2, error) {
	var f field.BidForwardPoints2
	err := m.Body.Get(&f)
	return f, err
}

//OfferForwardPoints2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferForwardPoints2() (field.OfferForwardPoints2, error) {
	var f field.OfferForwardPoints2
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrBidFxRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlCurrBidFxRate() (field.SettlCurrBidFxRate, error) {
	var f field.SettlCurrBidFxRate
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrOfferFxRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlCurrOfferFxRate() (field.SettlCurrOfferFxRate, error) {
	var f field.SettlCurrOfferFxRate
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrFxRateCalc is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlCurrFxRateCalc() (field.SettlCurrFxRateCalc, error) {
	var f field.SettlCurrFxRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Commission() (field.Commission, error) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CommType() (field.CommType, error) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderCapacity is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CustOrderCapacity() (field.CustOrderCapacity, error) {
	var f field.CustOrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//ExDestination is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ExDestination() (field.ExDestination, error) {
	var f field.ExDestination
	err := m.Body.Get(&f)
	return f, err
}

//QuoteStatus is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteStatus() (field.QuoteStatus, error) {
	var f field.QuoteStatus
	err := m.Body.Get(&f)
	return f, err
}
