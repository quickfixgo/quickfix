package fix50

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

//SecuritySubType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
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

//StrikeCurrency is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
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

//Pool is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
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

//ApplQueueDepth is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ApplQueueDepth() (field.ApplQueueDepth, error) {
	var f field.ApplQueueDepth
	err := m.Body.Get(&f)
	return f, err
}

//ApplQueueResolution is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ApplQueueResolution() (field.ApplQueueResolution, error) {
	var f field.ApplQueueResolution
	err := m.Body.Get(&f)
	return f, err
}

//MDReportID is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MDReportID() (field.MDReportID, error) {
	var f field.MDReportID
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//MDBookType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MDBookType() (field.MDBookType, error) {
	var f field.MDBookType
	err := m.Body.Get(&f)
	return f, err
}

//MDFeedType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MDFeedType() (field.MDFeedType, error) {
	var f field.MDFeedType
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//NoRoutingIDs is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoRoutingIDs() (field.NoRoutingIDs, error) {
	var f field.NoRoutingIDs
	err := m.Body.Get(&f)
	return f, err
}
