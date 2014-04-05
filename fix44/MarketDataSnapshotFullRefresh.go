package fix44

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type MarketDataSnapshotFullRefresh struct {
	quickfixgo.Message
}

func (m *MarketDataSnapshotFullRefresh) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) MDReqID() (*field.MDReqID, error) {
	f := new(field.MDReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) FinancialStatus() (*field.FinancialStatus, error) {
	f := new(field.FinancialStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) CorporateAction() (*field.CorporateAction, error) {
	f := new(field.CorporateAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) NoMDEntries() (*field.NoMDEntries, error) {
	f := new(field.NoMDEntries)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) ApplQueueDepth() (*field.ApplQueueDepth, error) {
	f := new(field.ApplQueueDepth)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) ApplQueueResolution() (*field.ApplQueueResolution, error) {
	f := new(field.ApplQueueResolution)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) NetChgPrevDay() (*field.NetChgPrevDay, error) {
	f := new(field.NetChgPrevDay)
	err := m.Body.Get(f)
	return f, err
}
