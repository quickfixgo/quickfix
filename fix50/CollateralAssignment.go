package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CollateralAssignment msg type = AY.
type CollateralAssignment struct {
	message.Message
}

//CollateralAssignmentBuilder builds CollateralAssignment messages.
type CollateralAssignmentBuilder struct {
	message.MessageBuilder
}

//CreateCollateralAssignmentBuilder returns an initialized CollateralAssignmentBuilder with specified required fields.
func CreateCollateralAssignmentBuilder(
	collasgnid field.CollAsgnID,
	collasgnreason field.CollAsgnReason,
	collasgntranstype field.CollAsgnTransType,
	transacttime field.TransactTime) CollateralAssignmentBuilder {
	var builder CollateralAssignmentBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(collasgnid)
	builder.Body.Set(collasgnreason)
	builder.Body.Set(collasgntranstype)
	builder.Body.Set(transacttime)
	return builder
}

//CollAsgnID is a required field for CollateralAssignment.
func (m CollateralAssignment) CollAsgnID() (field.CollAsgnID, error) {
	var f field.CollAsgnID
	err := m.Body.Get(&f)
	return f, err
}

//CollReqID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CollReqID() (field.CollReqID, error) {
	var f field.CollReqID
	err := m.Body.Get(&f)
	return f, err
}

//CollAsgnReason is a required field for CollateralAssignment.
func (m CollateralAssignment) CollAsgnReason() (field.CollAsgnReason, error) {
	var f field.CollAsgnReason
	err := m.Body.Get(&f)
	return f, err
}

//CollAsgnTransType is a required field for CollateralAssignment.
func (m CollateralAssignment) CollAsgnTransType() (field.CollAsgnTransType, error) {
	var f field.CollAsgnTransType
	err := m.Body.Get(&f)
	return f, err
}

//CollAsgnRefID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CollAsgnRefID() (field.CollAsgnRefID, error) {
	var f field.CollAsgnRefID
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for CollateralAssignment.
func (m CollateralAssignment) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ExpireTime() (field.ExpireTime, error) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrderID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryOrderID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecondaryOrderID() (field.SecondaryOrderID, error) {
	var f field.SecondaryOrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecondaryClOrdID() (field.SecondaryClOrdID, error) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//NoExecs is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoExecs() (field.NoExecs, error) {
	var f field.NoExecs
	err := m.Body.Get(&f)
	return f, err
}

//NoTrades is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoTrades() (field.NoTrades, error) {
	var f field.NoTrades
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for CollateralAssignment.
func (m CollateralAssignment) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for CollateralAssignment.
func (m CollateralAssignment) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for CollateralAssignment.
func (m CollateralAssignment) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for CollateralAssignment.
func (m CollateralAssignment) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for CollateralAssignment.
func (m CollateralAssignment) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for CollateralAssignment.
func (m CollateralAssignment) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for CollateralAssignment.
func (m CollateralAssignment) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for CollateralAssignment.
func (m CollateralAssignment) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AgreementDesc() (field.AgreementDesc, error) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AgreementID() (field.AgreementID, error) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AgreementDate() (field.AgreementDate, error) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AgreementCurrency() (field.AgreementCurrency, error) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) TerminationType() (field.TerminationType, error) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StartDate() (field.StartDate, error) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EndDate() (field.EndDate, error) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) DeliveryType() (field.DeliveryType, error) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MarginRatio() (field.MarginRatio, error) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//Quantity is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Quantity() (field.Quantity, error) {
	var f field.Quantity
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) QtyType() (field.QtyType, error) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//MarginExcess is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MarginExcess() (field.MarginExcess, error) {
	var f field.MarginExcess
	err := m.Body.Get(&f)
	return f, err
}

//TotalNetValue is a non-required field for CollateralAssignment.
func (m CollateralAssignment) TotalNetValue() (field.TotalNetValue, error) {
	var f field.TotalNetValue
	err := m.Body.Get(&f)
	return f, err
}

//CashOutstanding is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CashOutstanding() (field.CashOutstanding, error) {
	var f field.CashOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoTrdRegTimestamps() (field.NoTrdRegTimestamps, error) {
	var f field.NoTrdRegTimestamps
	err := m.Body.Get(&f)
	return f, err
}

//Side is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//NoMiscFees is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoMiscFees() (field.NoMiscFees, error) {
	var f field.NoMiscFees
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestAmt is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AccruedInterestAmt() (field.AccruedInterestAmt, error) {
	var f field.AccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//EndAccruedInterestAmt is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EndAccruedInterestAmt() (field.EndAccruedInterestAmt, error) {
	var f field.EndAccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//StartCash is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StartCash() (field.StartCash, error) {
	var f field.StartCash
	err := m.Body.Get(&f)
	return f, err
}

//EndCash is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EndCash() (field.EndCash, error) {
	var f field.EndCash
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkPrice() (field.BenchmarkPrice, error) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkPriceType() (field.BenchmarkPriceType, error) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkSecurityID() (field.BenchmarkSecurityID, error) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, error) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoStipulations() (field.NoStipulations, error) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//SettlDeliveryType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SettlDeliveryType() (field.SettlDeliveryType, error) {
	var f field.SettlDeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StandInstDbType() (field.StandInstDbType, error) {
	var f field.StandInstDbType
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbName is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StandInstDbName() (field.StandInstDbName, error) {
	var f field.StandInstDbName
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StandInstDbID() (field.StandInstDbID, error) {
	var f field.StandInstDbID
	err := m.Body.Get(&f)
	return f, err
}

//NoDlvyInst is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoDlvyInst() (field.NoDlvyInst, error) {
	var f field.NoDlvyInst
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SettlSessID() (field.SettlSessID, error) {
	var f field.SettlSessID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessSubID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SettlSessSubID() (field.SettlSessSubID, error) {
	var f field.SettlSessSubID
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
