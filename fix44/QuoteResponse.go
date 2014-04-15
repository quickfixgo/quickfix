package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//QuoteResponse msg type = AJ.
type QuoteResponse struct {
	message.Message
}

//QuoteResponseBuilder builds QuoteResponse messages.
type QuoteResponseBuilder struct {
	message.MessageBuilder
}

//CreateQuoteResponseBuilder returns an initialized QuoteResponseBuilder with specified required fields.
func CreateQuoteResponseBuilder(
	quoterespid field.QuoteRespID,
	quoteresptype field.QuoteRespType) QuoteResponseBuilder {
	var builder QuoteResponseBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(quoterespid)
	builder.Body.Set(quoteresptype)
	return builder
}

//QuoteRespID is a required field for QuoteResponse.
func (m QuoteResponse) QuoteRespID() (field.QuoteRespID, error) {
	var f field.QuoteRespID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a non-required field for QuoteResponse.
func (m QuoteResponse) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteRespType is a required field for QuoteResponse.
func (m QuoteResponse) QuoteRespType() (field.QuoteRespType, error) {
	var f field.QuoteRespType
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for QuoteResponse.
func (m QuoteResponse) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrderCapacity is a non-required field for QuoteResponse.
func (m QuoteResponse) OrderCapacity() (field.OrderCapacity, error) {
	var f field.OrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//IOIID is a non-required field for QuoteResponse.
func (m QuoteResponse) IOIID() (field.IOIID, error) {
	var f field.IOIID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteType is a non-required field for QuoteResponse.
func (m QuoteResponse) QuoteType() (field.QuoteType, error) {
	var f field.QuoteType
	err := m.Body.Get(&f)
	return f, err
}

//NoQuoteQualifiers is a non-required field for QuoteResponse.
func (m QuoteResponse) NoQuoteQualifiers() (field.NoQuoteQualifiers, error) {
	var f field.NoQuoteQualifiers
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for QuoteResponse.
func (m QuoteResponse) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for QuoteResponse.
func (m QuoteResponse) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for QuoteResponse.
func (m QuoteResponse) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for QuoteResponse.
func (m QuoteResponse) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for QuoteResponse.
func (m QuoteResponse) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for QuoteResponse.
func (m QuoteResponse) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for QuoteResponse.
func (m QuoteResponse) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for QuoteResponse.
func (m QuoteResponse) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for QuoteResponse.
func (m QuoteResponse) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for QuoteResponse.
func (m QuoteResponse) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for QuoteResponse.
func (m QuoteResponse) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for QuoteResponse.
func (m QuoteResponse) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for QuoteResponse.
func (m QuoteResponse) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for QuoteResponse.
func (m QuoteResponse) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for QuoteResponse.
func (m QuoteResponse) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for QuoteResponse.
func (m QuoteResponse) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for QuoteResponse.
func (m QuoteResponse) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for QuoteResponse.
func (m QuoteResponse) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for QuoteResponse.
func (m QuoteResponse) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for QuoteResponse.
func (m QuoteResponse) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for QuoteResponse.
func (m QuoteResponse) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for QuoteResponse.
func (m QuoteResponse) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for QuoteResponse.
func (m QuoteResponse) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for QuoteResponse.
func (m QuoteResponse) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for QuoteResponse.
func (m QuoteResponse) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for QuoteResponse.
func (m QuoteResponse) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for QuoteResponse.
func (m QuoteResponse) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for QuoteResponse.
func (m QuoteResponse) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for QuoteResponse.
func (m QuoteResponse) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for QuoteResponse.
func (m QuoteResponse) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for QuoteResponse.
func (m QuoteResponse) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for QuoteResponse.
func (m QuoteResponse) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for QuoteResponse.
func (m QuoteResponse) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for QuoteResponse.
func (m QuoteResponse) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for QuoteResponse.
func (m QuoteResponse) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for QuoteResponse.
func (m QuoteResponse) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for QuoteResponse.
func (m QuoteResponse) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for QuoteResponse.
func (m QuoteResponse) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for QuoteResponse.
func (m QuoteResponse) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for QuoteResponse.
func (m QuoteResponse) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for QuoteResponse.
func (m QuoteResponse) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for QuoteResponse.
func (m QuoteResponse) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for QuoteResponse.
func (m QuoteResponse) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for QuoteResponse.
func (m QuoteResponse) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for QuoteResponse.
func (m QuoteResponse) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for QuoteResponse.
func (m QuoteResponse) AgreementDesc() (field.AgreementDesc, error) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for QuoteResponse.
func (m QuoteResponse) AgreementID() (field.AgreementID, error) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for QuoteResponse.
func (m QuoteResponse) AgreementDate() (field.AgreementDate, error) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for QuoteResponse.
func (m QuoteResponse) AgreementCurrency() (field.AgreementCurrency, error) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for QuoteResponse.
func (m QuoteResponse) TerminationType() (field.TerminationType, error) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for QuoteResponse.
func (m QuoteResponse) StartDate() (field.StartDate, error) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for QuoteResponse.
func (m QuoteResponse) EndDate() (field.EndDate, error) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for QuoteResponse.
func (m QuoteResponse) DeliveryType() (field.DeliveryType, error) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for QuoteResponse.
func (m QuoteResponse) MarginRatio() (field.MarginRatio, error) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for QuoteResponse.
func (m QuoteResponse) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//Side is a non-required field for QuoteResponse.
func (m QuoteResponse) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for QuoteResponse.
func (m QuoteResponse) OrderQty() (field.OrderQty, error) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for QuoteResponse.
func (m QuoteResponse) CashOrderQty() (field.CashOrderQty, error) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrderPercent is a non-required field for QuoteResponse.
func (m QuoteResponse) OrderPercent() (field.OrderPercent, error) {
	var f field.OrderPercent
	err := m.Body.Get(&f)
	return f, err
}

//RoundingDirection is a non-required field for QuoteResponse.
func (m QuoteResponse) RoundingDirection() (field.RoundingDirection, error) {
	var f field.RoundingDirection
	err := m.Body.Get(&f)
	return f, err
}

//RoundingModulus is a non-required field for QuoteResponse.
func (m QuoteResponse) RoundingModulus() (field.RoundingModulus, error) {
	var f field.RoundingModulus
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for QuoteResponse.
func (m QuoteResponse) SettlType() (field.SettlType, error) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for QuoteResponse.
func (m QuoteResponse) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate2 is a non-required field for QuoteResponse.
func (m QuoteResponse) SettlDate2() (field.SettlDate2, error) {
	var f field.SettlDate2
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty2 is a non-required field for QuoteResponse.
func (m QuoteResponse) OrderQty2() (field.OrderQty2, error) {
	var f field.OrderQty2
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for QuoteResponse.
func (m QuoteResponse) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for QuoteResponse.
func (m QuoteResponse) NoStipulations() (field.NoStipulations, error) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for QuoteResponse.
func (m QuoteResponse) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for QuoteResponse.
func (m QuoteResponse) AcctIDSource() (field.AcctIDSource, error) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for QuoteResponse.
func (m QuoteResponse) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for QuoteResponse.
func (m QuoteResponse) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//BidPx is a non-required field for QuoteResponse.
func (m QuoteResponse) BidPx() (field.BidPx, error) {
	var f field.BidPx
	err := m.Body.Get(&f)
	return f, err
}

//OfferPx is a non-required field for QuoteResponse.
func (m QuoteResponse) OfferPx() (field.OfferPx, error) {
	var f field.OfferPx
	err := m.Body.Get(&f)
	return f, err
}

//MktBidPx is a non-required field for QuoteResponse.
func (m QuoteResponse) MktBidPx() (field.MktBidPx, error) {
	var f field.MktBidPx
	err := m.Body.Get(&f)
	return f, err
}

//MktOfferPx is a non-required field for QuoteResponse.
func (m QuoteResponse) MktOfferPx() (field.MktOfferPx, error) {
	var f field.MktOfferPx
	err := m.Body.Get(&f)
	return f, err
}

//MinBidSize is a non-required field for QuoteResponse.
func (m QuoteResponse) MinBidSize() (field.MinBidSize, error) {
	var f field.MinBidSize
	err := m.Body.Get(&f)
	return f, err
}

//BidSize is a non-required field for QuoteResponse.
func (m QuoteResponse) BidSize() (field.BidSize, error) {
	var f field.BidSize
	err := m.Body.Get(&f)
	return f, err
}

//MinOfferSize is a non-required field for QuoteResponse.
func (m QuoteResponse) MinOfferSize() (field.MinOfferSize, error) {
	var f field.MinOfferSize
	err := m.Body.Get(&f)
	return f, err
}

//OfferSize is a non-required field for QuoteResponse.
func (m QuoteResponse) OfferSize() (field.OfferSize, error) {
	var f field.OfferSize
	err := m.Body.Get(&f)
	return f, err
}

//ValidUntilTime is a non-required field for QuoteResponse.
func (m QuoteResponse) ValidUntilTime() (field.ValidUntilTime, error) {
	var f field.ValidUntilTime
	err := m.Body.Get(&f)
	return f, err
}

//BidSpotRate is a non-required field for QuoteResponse.
func (m QuoteResponse) BidSpotRate() (field.BidSpotRate, error) {
	var f field.BidSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//OfferSpotRate is a non-required field for QuoteResponse.
func (m QuoteResponse) OfferSpotRate() (field.OfferSpotRate, error) {
	var f field.OfferSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//BidForwardPoints is a non-required field for QuoteResponse.
func (m QuoteResponse) BidForwardPoints() (field.BidForwardPoints, error) {
	var f field.BidForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//OfferForwardPoints is a non-required field for QuoteResponse.
func (m QuoteResponse) OfferForwardPoints() (field.OfferForwardPoints, error) {
	var f field.OfferForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//MidPx is a non-required field for QuoteResponse.
func (m QuoteResponse) MidPx() (field.MidPx, error) {
	var f field.MidPx
	err := m.Body.Get(&f)
	return f, err
}

//BidYield is a non-required field for QuoteResponse.
func (m QuoteResponse) BidYield() (field.BidYield, error) {
	var f field.BidYield
	err := m.Body.Get(&f)
	return f, err
}

//MidYield is a non-required field for QuoteResponse.
func (m QuoteResponse) MidYield() (field.MidYield, error) {
	var f field.MidYield
	err := m.Body.Get(&f)
	return f, err
}

//OfferYield is a non-required field for QuoteResponse.
func (m QuoteResponse) OfferYield() (field.OfferYield, error) {
	var f field.OfferYield
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for QuoteResponse.
func (m QuoteResponse) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a non-required field for QuoteResponse.
func (m QuoteResponse) OrdType() (field.OrdType, error) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//BidForwardPoints2 is a non-required field for QuoteResponse.
func (m QuoteResponse) BidForwardPoints2() (field.BidForwardPoints2, error) {
	var f field.BidForwardPoints2
	err := m.Body.Get(&f)
	return f, err
}

//OfferForwardPoints2 is a non-required field for QuoteResponse.
func (m QuoteResponse) OfferForwardPoints2() (field.OfferForwardPoints2, error) {
	var f field.OfferForwardPoints2
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrBidFxRate is a non-required field for QuoteResponse.
func (m QuoteResponse) SettlCurrBidFxRate() (field.SettlCurrBidFxRate, error) {
	var f field.SettlCurrBidFxRate
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrOfferFxRate is a non-required field for QuoteResponse.
func (m QuoteResponse) SettlCurrOfferFxRate() (field.SettlCurrOfferFxRate, error) {
	var f field.SettlCurrOfferFxRate
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrFxRateCalc is a non-required field for QuoteResponse.
func (m QuoteResponse) SettlCurrFxRateCalc() (field.SettlCurrFxRateCalc, error) {
	var f field.SettlCurrFxRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for QuoteResponse.
func (m QuoteResponse) Commission() (field.Commission, error) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for QuoteResponse.
func (m QuoteResponse) CommType() (field.CommType, error) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderCapacity is a non-required field for QuoteResponse.
func (m QuoteResponse) CustOrderCapacity() (field.CustOrderCapacity, error) {
	var f field.CustOrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//ExDestination is a non-required field for QuoteResponse.
func (m QuoteResponse) ExDestination() (field.ExDestination, error) {
	var f field.ExDestination
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for QuoteResponse.
func (m QuoteResponse) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for QuoteResponse.
func (m QuoteResponse) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for QuoteResponse.
func (m QuoteResponse) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for QuoteResponse.
func (m QuoteResponse) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for QuoteResponse.
func (m QuoteResponse) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for QuoteResponse.
func (m QuoteResponse) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for QuoteResponse.
func (m QuoteResponse) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for QuoteResponse.
func (m QuoteResponse) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for QuoteResponse.
func (m QuoteResponse) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for QuoteResponse.
func (m QuoteResponse) BenchmarkPrice() (field.BenchmarkPrice, error) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for QuoteResponse.
func (m QuoteResponse) BenchmarkPriceType() (field.BenchmarkPriceType, error) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for QuoteResponse.
func (m QuoteResponse) BenchmarkSecurityID() (field.BenchmarkSecurityID, error) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for QuoteResponse.
func (m QuoteResponse) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, error) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for QuoteResponse.
func (m QuoteResponse) YieldType() (field.YieldType, error) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for QuoteResponse.
func (m QuoteResponse) Yield() (field.Yield, error) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//YieldCalcDate is a non-required field for QuoteResponse.
func (m QuoteResponse) YieldCalcDate() (field.YieldCalcDate, error) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for QuoteResponse.
func (m QuoteResponse) YieldRedemptionDate() (field.YieldRedemptionDate, error) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for QuoteResponse.
func (m QuoteResponse) YieldRedemptionPrice() (field.YieldRedemptionPrice, error) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for QuoteResponse.
func (m QuoteResponse) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, error) {
	var f field.YieldRedemptionPriceType
	err := m.Body.Get(&f)
	return f, err
}
