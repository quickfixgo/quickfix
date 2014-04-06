package fix44

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type QuoteStatusReport struct {
	quickfix.Message
}

func (m *QuoteStatusReport) YieldCalcDate() (*field.YieldCalcDate, error) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) DeliveryType() (*field.DeliveryType, error) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MinBidSize() (*field.MinBidSize, error) {
	f := new(field.MinBidSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) CashOrderQty() (*field.CashOrderQty, error) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) NoStipulations() (*field.NoStipulations, error) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) YieldRedemptionDate() (*field.YieldRedemptionDate, error) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, error) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OfferSize() (*field.OfferSize, error) {
	f := new(field.OfferSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OfferSpotRate() (*field.OfferSpotRate, error) {
	f := new(field.OfferSpotRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) QuoteType() (*field.QuoteType, error) {
	f := new(field.QuoteType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) CommType() (*field.CommType, error) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) ValidUntilTime() (*field.ValidUntilTime, error) {
	f := new(field.ValidUntilTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SettlCurrOfferFxRate() (*field.SettlCurrOfferFxRate, error) {
	f := new(field.SettlCurrOfferFxRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OrderQty2() (*field.OrderQty2, error) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BidSize() (*field.BidSize, error) {
	f := new(field.BidSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) QuoteStatusReqID() (*field.QuoteStatusReqID, error) {
	f := new(field.QuoteStatusReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SettlCurrBidFxRate() (*field.SettlCurrBidFxRate, error) {
	f := new(field.SettlCurrBidFxRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) QuoteRespID() (*field.QuoteRespID, error) {
	f := new(field.QuoteRespID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) YieldType() (*field.YieldType, error) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) CustOrderCapacity() (*field.CustOrderCapacity, error) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) TerminationType() (*field.TerminationType, error) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MktOfferPx() (*field.MktOfferPx, error) {
	f := new(field.MktOfferPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MarginRatio() (*field.MarginRatio, error) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BidPx() (*field.BidPx, error) {
	f := new(field.BidPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BenchmarkSecurityID() (*field.BenchmarkSecurityID, error) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OfferForwardPoints() (*field.OfferForwardPoints, error) {
	f := new(field.OfferForwardPoints)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OfferYield() (*field.OfferYield, error) {
	f := new(field.OfferYield)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) NoQuoteQualifiers() (*field.NoQuoteQualifiers, error) {
	f := new(field.NoQuoteQualifiers)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BidForwardPoints() (*field.BidForwardPoints, error) {
	f := new(field.BidForwardPoints)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BidYield() (*field.BidYield, error) {
	f := new(field.BidYield)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BidForwardPoints2() (*field.BidForwardPoints2, error) {
	f := new(field.BidForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OrderPercent() (*field.OrderPercent, error) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BenchmarkPriceType() (*field.BenchmarkPriceType, error) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MktBidPx() (*field.MktBidPx, error) {
	f := new(field.MktBidPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalc, error) {
	f := new(field.SettlCurrFxRateCalc)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) ExDestination() (*field.ExDestination, error) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) YieldRedemptionPrice() (*field.YieldRedemptionPrice, error) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BenchmarkPrice() (*field.BenchmarkPrice, error) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) AgreementID() (*field.AgreementID, error) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SettlDate2() (*field.SettlDate2, error) {
	f := new(field.SettlDate2)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OfferPx() (*field.OfferPx, error) {
	f := new(field.OfferPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OfferForwardPoints2() (*field.OfferForwardPoints2, error) {
	f := new(field.OfferForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) AgreementCurrency() (*field.AgreementCurrency, error) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MidYield() (*field.MidYield, error) {
	f := new(field.MidYield)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) QuoteStatus() (*field.QuoteStatus, error) {
	f := new(field.QuoteStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Yield() (*field.Yield, error) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Commission() (*field.Commission, error) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) EndDate() (*field.EndDate, error) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) RoundingModulus() (*field.RoundingModulus, error) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) AgreementDesc() (*field.AgreementDesc, error) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, error) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) RoundingDirection() (*field.RoundingDirection, error) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MinOfferSize() (*field.MinOfferSize, error) {
	f := new(field.MinOfferSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BidSpotRate() (*field.BidSpotRate, error) {
	f := new(field.BidSpotRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) StartDate() (*field.StartDate, error) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) AgreementDate() (*field.AgreementDate, error) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MidPx() (*field.MidPx, error) {
	f := new(field.MidPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
