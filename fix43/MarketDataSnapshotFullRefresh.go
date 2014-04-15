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

//CreateMarketDataSnapshotFullRefreshBuilder returns an initialized MarketDataSnapshotFullRefreshBuilder with specified required fields.
func CreateMarketDataSnapshotFullRefreshBuilder(
	nomdentries field.NoMDEntries) MarketDataSnapshotFullRefreshBuilder {
	var builder MarketDataSnapshotFullRefreshBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(nomdentries)
	return builder
}

//MDReqID is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MDReqID() (field.MDReqID, error) {
	var f field.MDReqID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//FinancialStatus is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) FinancialStatus() (field.FinancialStatus, error) {
	var f field.FinancialStatus
	err := m.Body.Get(&f)
	return f, err
}

//CorporateAction is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CorporateAction() (field.CorporateAction, error) {
	var f field.CorporateAction
	err := m.Body.Get(&f)
	return f, err
}

//TotalVolumeTraded is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) TotalVolumeTraded() (field.TotalVolumeTraded, error) {
	var f field.TotalVolumeTraded
	err := m.Body.Get(&f)
	return f, err
}

//TotalVolumeTradedDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) TotalVolumeTradedDate() (field.TotalVolumeTradedDate, error) {
	var f field.TotalVolumeTradedDate
	err := m.Body.Get(&f)
	return f, err
}

//TotalVolumeTradedTime is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) TotalVolumeTradedTime() (field.TotalVolumeTradedTime, error) {
	var f field.TotalVolumeTradedTime
	err := m.Body.Get(&f)
	return f, err
}

//NetChgPrevDay is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NetChgPrevDay() (field.NetChgPrevDay, error) {
	var f field.NetChgPrevDay
	err := m.Body.Get(&f)
	return f, err
}

//NoMDEntries is a required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoMDEntries() (field.NoMDEntries, error) {
	var f field.NoMDEntries
	err := m.Body.Get(&f)
	return f, err
}
