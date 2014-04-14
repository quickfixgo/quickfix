package fix43

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//MarketDataSnapshotFullRefresh msg type = W.
type MarketDataSnapshotFullRefresh struct {
	message.Message
}

//MarketDataSnapshotFullRefreshBuilder builds MarketDataSnapshotFullRefresh messages.
type MarketDataSnapshotFullRefreshBuilder struct {
	message.MessageBuilder
}

//NewMarketDataSnapshotFullRefreshBuilder returns an initialized MarketDataSnapshotFullRefreshBuilder with specified required fields.
func NewMarketDataSnapshotFullRefreshBuilder(
	nomdentries field.NoMDEntries) *MarketDataSnapshotFullRefreshBuilder {
	builder := new(MarketDataSnapshotFullRefreshBuilder)
	builder.Body.Set(nomdentries)
	return builder
}

//MDReqID is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) MDReqID() (*field.MDReqID, error) {
	f := new(field.MDReqID)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//FinancialStatus is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) FinancialStatus() (*field.FinancialStatus, error) {
	f := new(field.FinancialStatus)
	err := m.Body.Get(f)
	return f, err
}

//CorporateAction is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) CorporateAction() (*field.CorporateAction, error) {
	f := new(field.CorporateAction)
	err := m.Body.Get(f)
	return f, err
}

//TotalVolumeTraded is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) TotalVolumeTraded() (*field.TotalVolumeTraded, error) {
	f := new(field.TotalVolumeTraded)
	err := m.Body.Get(f)
	return f, err
}

//TotalVolumeTradedDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) TotalVolumeTradedDate() (*field.TotalVolumeTradedDate, error) {
	f := new(field.TotalVolumeTradedDate)
	err := m.Body.Get(f)
	return f, err
}

//TotalVolumeTradedTime is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) TotalVolumeTradedTime() (*field.TotalVolumeTradedTime, error) {
	f := new(field.TotalVolumeTradedTime)
	err := m.Body.Get(f)
	return f, err
}

//NetChgPrevDay is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) NetChgPrevDay() (*field.NetChgPrevDay, error) {
	f := new(field.NetChgPrevDay)
	err := m.Body.Get(f)
	return f, err
}

//NoMDEntries is a required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) NoMDEntries() (*field.NoMDEntries, error) {
	f := new(field.NoMDEntries)
	err := m.Body.Get(f)
	return f, err
}
