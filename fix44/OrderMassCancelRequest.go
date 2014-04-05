package fix44

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type OrderMassCancelRequest struct {
	quickfixgo.Message
}

func (m *OrderMassCancelRequest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingSecurityType() (*field.UnderlyingSecurityType, error) {
	f := new(field.UnderlyingSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLen, error) {
	f := new(field.EncodedUnderlyingIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityType, error) {
	f := new(field.UnderlyingRepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingOptAttribute() (*field.UnderlyingOptAttribute, error) {
	f := new(field.UnderlyingOptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingCFICode() (*field.UnderlyingCFICode, error) {
	f := new(field.UnderlyingCFICode)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingQty() (*field.UnderlyingQty, error) {
	f := new(field.UnderlyingQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingEndValue() (*field.UnderlyingEndValue, error) {
	f := new(field.UnderlyingEndValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistry, error) {
	f := new(field.UnderlyingInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssue, error) {
	f := new(field.UnderlyingLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrency, error) {
	f := new(field.UnderlyingStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuer, error) {
	f := new(field.EncodedUnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) NoUnderlyingStips() (*field.NoUnderlyingStips, error) {
	f := new(field.NoUnderlyingStips)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingIssuer() (*field.UnderlyingIssuer, error) {
	f := new(field.UnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingSymbol() (*field.UnderlyingSymbol, error) {
	f := new(field.UnderlyingSymbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssue, error) {
	f := new(field.UnderlyingCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchange, error) {
	f := new(field.UnderlyingSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingCurrency() (*field.UnderlyingCurrency, error) {
	f := new(field.UnderlyingCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPrice, error) {
	f := new(field.UnderlyingDirtyPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingSecurityID() (*field.UnderlyingSecurityID, error) {
	f := new(field.UnderlyingSecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingCPRegType() (*field.UnderlyingCPRegType, error) {
	f := new(field.UnderlyingCPRegType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) MassCancelRequestType() (*field.MassCancelRequestType, error) {
	f := new(field.MassCancelRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSource, error) {
	f := new(field.UnderlyingSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingProduct() (*field.UnderlyingProduct, error) {
	f := new(field.UnderlyingProduct)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDate, error) {
	f := new(field.UnderlyingCouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingCouponRate() (*field.UnderlyingCouponRate, error) {
	f := new(field.UnderlyingCouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingStartValue() (*field.UnderlyingStartValue, error) {
	f := new(field.UnderlyingStartValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplier, error) {
	f := new(field.UnderlyingContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingCPProgram() (*field.UnderlyingCPProgram, error) {
	f := new(field.UnderlyingCPProgram)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingCurrentValue() (*field.UnderlyingCurrentValue, error) {
	f := new(field.UnderlyingCurrentValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDesc, error) {
	f := new(field.EncodedUnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfx, error) {
	f := new(field.UnderlyingSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltID, error) {
	f := new(field.NoUnderlyingSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingCreditRating() (*field.UnderlyingCreditRating, error) {
	f := new(field.UnderlyingCreditRating)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssue, error) {
	f := new(field.UnderlyingStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDesc, error) {
	f := new(field.UnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingEndPrice() (*field.UnderlyingEndPrice, error) {
	f := new(field.UnderlyingEndPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYear, error) {
	f := new(field.UnderlyingMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTerm, error) {
	f := new(field.UnderlyingRepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingIssueDate() (*field.UnderlyingIssueDate, error) {
	f := new(field.UnderlyingIssueDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRate, error) {
	f := new(field.UnderlyingRepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingFactor() (*field.UnderlyingFactor, error) {
	f := new(field.UnderlyingFactor)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubType, error) {
	f := new(field.UnderlyingSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDate, error) {
	f := new(field.UnderlyingRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingStrikePrice() (*field.UnderlyingStrikePrice, error) {
	f := new(field.UnderlyingStrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingMaturityDate() (*field.UnderlyingMaturityDate, error) {
	f := new(field.UnderlyingMaturityDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLen, error) {
	f := new(field.EncodedUnderlyingSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelRequest) UnderlyingPx() (*field.UnderlyingPx, error) {
	f := new(field.UnderlyingPx)
	err := m.Body.Get(f)
	return f, err
}
