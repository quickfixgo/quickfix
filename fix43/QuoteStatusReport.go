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

//NewQuoteStatusReportBuilder returns an initialized QuoteStatusReportBuilder with specified required fields.
func NewQuoteStatusReportBuilder(
	quoteid field.QuoteID) *QuoteStatusReportBuilder {
	builder := new(QuoteStatusReportBuilder)
	builder.Body.Set(quoteid)
	return builder
}

//QuoteStatusReqID is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) QuoteStatusReqID() (*field.QuoteStatusReqID, error) {
	f := new(field.QuoteStatusReqID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteReqID is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteID is a required field for QuoteStatusReport.
func (m *QuoteStatusReport) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) QuoteType() (*field.QuoteType, error) {
	f := new(field.QuoteType)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//AccountType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//BidPx is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BidPx() (*field.BidPx, error) {
	f := new(field.BidPx)
	err := m.Body.Get(f)
	return f, err
}

//OfferPx is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OfferPx() (*field.OfferPx, error) {
	f := new(field.OfferPx)
	err := m.Body.Get(f)
	return f, err
}

//MktBidPx is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MktBidPx() (*field.MktBidPx, error) {
	f := new(field.MktBidPx)
	err := m.Body.Get(f)
	return f, err
}

//MktOfferPx is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MktOfferPx() (*field.MktOfferPx, error) {
	f := new(field.MktOfferPx)
	err := m.Body.Get(f)
	return f, err
}

//MinBidSize is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MinBidSize() (*field.MinBidSize, error) {
	f := new(field.MinBidSize)
	err := m.Body.Get(f)
	return f, err
}

//BidSize is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BidSize() (*field.BidSize, error) {
	f := new(field.BidSize)
	err := m.Body.Get(f)
	return f, err
}

//MinOfferSize is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MinOfferSize() (*field.MinOfferSize, error) {
	f := new(field.MinOfferSize)
	err := m.Body.Get(f)
	return f, err
}

//OfferSize is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OfferSize() (*field.OfferSize, error) {
	f := new(field.OfferSize)
	err := m.Body.Get(f)
	return f, err
}

//ValidUntilTime is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) ValidUntilTime() (*field.ValidUntilTime, error) {
	f := new(field.ValidUntilTime)
	err := m.Body.Get(f)
	return f, err
}

//BidSpotRate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BidSpotRate() (*field.BidSpotRate, error) {
	f := new(field.BidSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//OfferSpotRate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OfferSpotRate() (*field.OfferSpotRate, error) {
	f := new(field.OfferSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//BidForwardPoints is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BidForwardPoints() (*field.BidForwardPoints, error) {
	f := new(field.BidForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//OfferForwardPoints is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OfferForwardPoints() (*field.OfferForwardPoints, error) {
	f := new(field.OfferForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//MidPx is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MidPx() (*field.MidPx, error) {
	f := new(field.MidPx)
	err := m.Body.Get(f)
	return f, err
}

//BidYield is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BidYield() (*field.BidYield, error) {
	f := new(field.BidYield)
	err := m.Body.Get(f)
	return f, err
}

//MidYield is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MidYield() (*field.MidYield, error) {
	f := new(field.MidYield)
	err := m.Body.Get(f)
	return f, err
}

//OfferYield is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OfferYield() (*field.OfferYield, error) {
	f := new(field.OfferYield)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//FutSettDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) FutSettDate() (*field.FutSettDate, error) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//OrdType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//FutSettDate2 is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) FutSettDate2() (*field.FutSettDate2, error) {
	f := new(field.FutSettDate2)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty2 is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OrderQty2() (*field.OrderQty2, error) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}

//BidForwardPoints2 is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BidForwardPoints2() (*field.BidForwardPoints2, error) {
	f := new(field.BidForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}

//OfferForwardPoints2 is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OfferForwardPoints2() (*field.OfferForwardPoints2, error) {
	f := new(field.OfferForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrBidFxRate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SettlCurrBidFxRate() (*field.SettlCurrBidFxRate, error) {
	f := new(field.SettlCurrBidFxRate)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrOfferFxRate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SettlCurrOfferFxRate() (*field.SettlCurrOfferFxRate, error) {
	f := new(field.SettlCurrOfferFxRate)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrFxRateCalc is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalc, error) {
	f := new(field.SettlCurrFxRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//Commission is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Commission() (*field.Commission, error) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//CommType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CommType() (*field.CommType, error) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//CustOrderCapacity is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CustOrderCapacity() (*field.CustOrderCapacity, error) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//ExDestination is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) ExDestination() (*field.ExDestination, error) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//QuoteStatus is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) QuoteStatus() (*field.QuoteStatus, error) {
	f := new(field.QuoteStatus)
	err := m.Body.Get(f)
	return f, err
}
