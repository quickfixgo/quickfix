package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//NewOrderSingle msg type = D.
type NewOrderSingle struct {
	message.Message
}

//NewOrderSingleBuilder builds NewOrderSingle messages.
type NewOrderSingleBuilder struct {
	message.MessageBuilder
}

//CreateNewOrderSingleBuilder returns an initialized NewOrderSingleBuilder with specified required fields.
func CreateNewOrderSingleBuilder(
	clordid field.ClOrdID,
	side field.Side,
	transacttime field.TransactTime,
	ordtype field.OrdType) NewOrderSingleBuilder {
	var builder NewOrderSingleBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(clordid)
	builder.Body.Set(side)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//ClOrdID is a required field for NewOrderSingle.
func (m NewOrderSingle) ClOrdID() (field.ClOrdID, errors.MessageRejectError) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecondaryClOrdID() (field.SecondaryClOrdID, errors.MessageRejectError) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdLinkID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ClOrdLinkID() (field.ClOrdLinkID, errors.MessageRejectError) {
	var f field.ClOrdLinkID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoPartyIDs() (field.NoPartyIDs, errors.MessageRejectError) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//TradeOriginationDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TradeOriginationDate() (field.TradeOriginationDate, errors.MessageRejectError) {
	var f field.TradeOriginationDate
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TradeDate() (field.TradeDate, errors.MessageRejectError) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Account() (field.Account, errors.MessageRejectError) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AcctIDSource() (field.AcctIDSource, errors.MessageRejectError) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AccountType() (field.AccountType, errors.MessageRejectError) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//DayBookingInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DayBookingInst() (field.DayBookingInst, errors.MessageRejectError) {
	var f field.DayBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//BookingUnit is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BookingUnit() (field.BookingUnit, errors.MessageRejectError) {
	var f field.BookingUnit
	err := m.Body.Get(&f)
	return f, err
}

//PreallocMethod is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PreallocMethod() (field.PreallocMethod, errors.MessageRejectError) {
	var f field.PreallocMethod
	err := m.Body.Get(&f)
	return f, err
}

//AllocID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AllocID() (field.AllocID, errors.MessageRejectError) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//NoAllocs is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoAllocs() (field.NoAllocs, errors.MessageRejectError) {
	var f field.NoAllocs
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SettlType() (field.SettlType, errors.MessageRejectError) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SettlDate() (field.SettlDate, errors.MessageRejectError) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//CashMargin is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CashMargin() (field.CashMargin, errors.MessageRejectError) {
	var f field.CashMargin
	err := m.Body.Get(&f)
	return f, err
}

//ClearingFeeIndicator is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ClearingFeeIndicator() (field.ClearingFeeIndicator, errors.MessageRejectError) {
	var f field.ClearingFeeIndicator
	err := m.Body.Get(&f)
	return f, err
}

//HandlInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) HandlInst() (field.HandlInst, errors.MessageRejectError) {
	var f field.HandlInst
	err := m.Body.Get(&f)
	return f, err
}

//ExecInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExecInst() (field.ExecInst, errors.MessageRejectError) {
	var f field.ExecInst
	err := m.Body.Get(&f)
	return f, err
}

//MinQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MinQty() (field.MinQty, errors.MessageRejectError) {
	var f field.MinQty
	err := m.Body.Get(&f)
	return f, err
}

//MaxFloor is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaxFloor() (field.MaxFloor, errors.MessageRejectError) {
	var f field.MaxFloor
	err := m.Body.Get(&f)
	return f, err
}

//ExDestination is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExDestination() (field.ExDestination, errors.MessageRejectError) {
	var f field.ExDestination
	err := m.Body.Get(&f)
	return f, err
}

//NoTradingSessions is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoTradingSessions() (field.NoTradingSessions, errors.MessageRejectError) {
	var f field.NoTradingSessions
	err := m.Body.Get(&f)
	return f, err
}

//ProcessCode is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ProcessCode() (field.ProcessCode, errors.MessageRejectError) {
	var f field.ProcessCode
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityIDSource() (field.SecurityIDSource, errors.MessageRejectError) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoSecurityAltID() (field.NoSecurityAltID, errors.MessageRejectError) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Product() (field.Product, errors.MessageRejectError) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CFICode() (field.CFICode, errors.MessageRejectError) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecuritySubType() (field.SecuritySubType, errors.MessageRejectError) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaturityDate() (field.MaturityDate, errors.MessageRejectError) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CouponPaymentDate() (field.CouponPaymentDate, errors.MessageRejectError) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) IssueDate() (field.IssueDate, errors.MessageRejectError) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, errors.MessageRejectError) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RepurchaseTerm() (field.RepurchaseTerm, errors.MessageRejectError) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RepurchaseRate() (field.RepurchaseRate, errors.MessageRejectError) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Factor() (field.Factor, errors.MessageRejectError) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CreditRating() (field.CreditRating, errors.MessageRejectError) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for NewOrderSingle.
func (m NewOrderSingle) InstrRegistry() (field.InstrRegistry, errors.MessageRejectError) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CountryOfIssue() (field.CountryOfIssue, errors.MessageRejectError) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for NewOrderSingle.
func (m NewOrderSingle) LocaleOfIssue() (field.LocaleOfIssue, errors.MessageRejectError) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RedemptionDate() (field.RedemptionDate, errors.MessageRejectError) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StrikeCurrency() (field.StrikeCurrency, errors.MessageRejectError) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ContractMultiplier() (field.ContractMultiplier, errors.MessageRejectError) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CouponRate() (field.CouponRate, errors.MessageRejectError) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedIssuerLen() (field.EncodedIssuerLen, errors.MessageRejectError) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedIssuer() (field.EncodedIssuer, errors.MessageRejectError) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedSecurityDesc() (field.EncodedSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Pool() (field.Pool, errors.MessageRejectError) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ContractSettlMonth() (field.ContractSettlMonth, errors.MessageRejectError) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CPProgram() (field.CPProgram, errors.MessageRejectError) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CPRegType() (field.CPRegType, errors.MessageRejectError) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoEvents() (field.NoEvents, errors.MessageRejectError) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DatedDate() (field.DatedDate, errors.MessageRejectError) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) InterestAccrualDate() (field.InterestAccrualDate, errors.MessageRejectError) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityStatus() (field.SecurityStatus, errors.MessageRejectError) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SettleOnOpenFlag() (field.SettleOnOpenFlag, errors.MessageRejectError) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for NewOrderSingle.
func (m NewOrderSingle) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StrikeMultiplier() (field.StrikeMultiplier, errors.MessageRejectError) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StrikeValue() (field.StrikeValue, errors.MessageRejectError) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MinPriceIncrement() (field.MinPriceIncrement, errors.MessageRejectError) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PositionLimit() (field.PositionLimit, errors.MessageRejectError) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NTPositionLimit() (field.NTPositionLimit, errors.MessageRejectError) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoInstrumentParties() (field.NoInstrumentParties, errors.MessageRejectError) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for NewOrderSingle.
func (m NewOrderSingle) UnitOfMeasure() (field.UnitOfMeasure, errors.MessageRejectError) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TimeUnit() (field.TimeUnit, errors.MessageRejectError) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaturityTime() (field.MaturityTime, errors.MessageRejectError) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AgreementDesc() (field.AgreementDesc, errors.MessageRejectError) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AgreementID() (field.AgreementID, errors.MessageRejectError) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AgreementDate() (field.AgreementDate, errors.MessageRejectError) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AgreementCurrency() (field.AgreementCurrency, errors.MessageRejectError) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TerminationType() (field.TerminationType, errors.MessageRejectError) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StartDate() (field.StartDate, errors.MessageRejectError) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EndDate() (field.EndDate, errors.MessageRejectError) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DeliveryType() (field.DeliveryType, errors.MessageRejectError) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MarginRatio() (field.MarginRatio, errors.MessageRejectError) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoUnderlyings() (field.NoUnderlyings, errors.MessageRejectError) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//PrevClosePx is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PrevClosePx() (field.PrevClosePx, errors.MessageRejectError) {
	var f field.PrevClosePx
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for NewOrderSingle.
func (m NewOrderSingle) Side() (field.Side, errors.MessageRejectError) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//LocateReqd is a non-required field for NewOrderSingle.
func (m NewOrderSingle) LocateReqd() (field.LocateReqd, errors.MessageRejectError) {
	var f field.LocateReqd
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for NewOrderSingle.
func (m NewOrderSingle) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoStipulations() (field.NoStipulations, errors.MessageRejectError) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) QtyType() (field.QtyType, errors.MessageRejectError) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderQty() (field.OrderQty, errors.MessageRejectError) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CashOrderQty() (field.CashOrderQty, errors.MessageRejectError) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrderPercent is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderPercent() (field.OrderPercent, errors.MessageRejectError) {
	var f field.OrderPercent
	err := m.Body.Get(&f)
	return f, err
}

//RoundingDirection is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RoundingDirection() (field.RoundingDirection, errors.MessageRejectError) {
	var f field.RoundingDirection
	err := m.Body.Get(&f)
	return f, err
}

//RoundingModulus is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RoundingModulus() (field.RoundingModulus, errors.MessageRejectError) {
	var f field.RoundingModulus
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a required field for NewOrderSingle.
func (m NewOrderSingle) OrdType() (field.OrdType, errors.MessageRejectError) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PriceType() (field.PriceType, errors.MessageRejectError) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Price() (field.Price, errors.MessageRejectError) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//StopPx is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StopPx() (field.StopPx, errors.MessageRejectError) {
	var f field.StopPx
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Spread() (field.Spread, errors.MessageRejectError) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkCurveName() (field.BenchmarkCurveName, errors.MessageRejectError) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, errors.MessageRejectError) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkPrice() (field.BenchmarkPrice, errors.MessageRejectError) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkPriceType() (field.BenchmarkPriceType, errors.MessageRejectError) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkSecurityID() (field.BenchmarkSecurityID, errors.MessageRejectError) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) YieldType() (field.YieldType, errors.MessageRejectError) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Yield() (field.Yield, errors.MessageRejectError) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//YieldCalcDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) YieldCalcDate() (field.YieldCalcDate, errors.MessageRejectError) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) YieldRedemptionDate() (field.YieldRedemptionDate, errors.MessageRejectError) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for NewOrderSingle.
func (m NewOrderSingle) YieldRedemptionPrice() (field.YieldRedemptionPrice, errors.MessageRejectError) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, errors.MessageRejectError) {
	var f field.YieldRedemptionPriceType
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Currency() (field.Currency, errors.MessageRejectError) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//ComplianceID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ComplianceID() (field.ComplianceID, errors.MessageRejectError) {
	var f field.ComplianceID
	err := m.Body.Get(&f)
	return f, err
}

//SolicitedFlag is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SolicitedFlag() (field.SolicitedFlag, errors.MessageRejectError) {
	var f field.SolicitedFlag
	err := m.Body.Get(&f)
	return f, err
}

//IOIID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) IOIID() (field.IOIID, errors.MessageRejectError) {
	var f field.IOIID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) QuoteID() (field.QuoteID, errors.MessageRejectError) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//TimeInForce is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TimeInForce() (field.TimeInForce, errors.MessageRejectError) {
	var f field.TimeInForce
	err := m.Body.Get(&f)
	return f, err
}

//EffectiveTime is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EffectiveTime() (field.EffectiveTime, errors.MessageRejectError) {
	var f field.EffectiveTime
	err := m.Body.Get(&f)
	return f, err
}

//ExpireDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExpireDate() (field.ExpireDate, errors.MessageRejectError) {
	var f field.ExpireDate
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExpireTime() (field.ExpireTime, errors.MessageRejectError) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//GTBookingInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) GTBookingInst() (field.GTBookingInst, errors.MessageRejectError) {
	var f field.GTBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Commission() (field.Commission, errors.MessageRejectError) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CommType() (field.CommType, errors.MessageRejectError) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//CommCurrency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CommCurrency() (field.CommCurrency, errors.MessageRejectError) {
	var f field.CommCurrency
	err := m.Body.Get(&f)
	return f, err
}

//FundRenewWaiv is a non-required field for NewOrderSingle.
func (m NewOrderSingle) FundRenewWaiv() (field.FundRenewWaiv, errors.MessageRejectError) {
	var f field.FundRenewWaiv
	err := m.Body.Get(&f)
	return f, err
}

//OrderCapacity is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderCapacity() (field.OrderCapacity, errors.MessageRejectError) {
	var f field.OrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//OrderRestrictions is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderRestrictions() (field.OrderRestrictions, errors.MessageRejectError) {
	var f field.OrderRestrictions
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderCapacity is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CustOrderCapacity() (field.CustOrderCapacity, errors.MessageRejectError) {
	var f field.CustOrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//ForexReq is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ForexReq() (field.ForexReq, errors.MessageRejectError) {
	var f field.ForexReq
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SettlCurrency() (field.SettlCurrency, errors.MessageRejectError) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BookingType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BookingType() (field.BookingType, errors.MessageRejectError) {
	var f field.BookingType
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate2 is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SettlDate2() (field.SettlDate2, errors.MessageRejectError) {
	var f field.SettlDate2
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty2 is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderQty2() (field.OrderQty2, errors.MessageRejectError) {
	var f field.OrderQty2
	err := m.Body.Get(&f)
	return f, err
}

//Price2 is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Price2() (field.Price2, errors.MessageRejectError) {
	var f field.Price2
	err := m.Body.Get(&f)
	return f, err
}

//PositionEffect is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PositionEffect() (field.PositionEffect, errors.MessageRejectError) {
	var f field.PositionEffect
	err := m.Body.Get(&f)
	return f, err
}

//CoveredOrUncovered is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CoveredOrUncovered() (field.CoveredOrUncovered, errors.MessageRejectError) {
	var f field.CoveredOrUncovered
	err := m.Body.Get(&f)
	return f, err
}

//MaxShow is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaxShow() (field.MaxShow, errors.MessageRejectError) {
	var f field.MaxShow
	err := m.Body.Get(&f)
	return f, err
}

//PegOffsetValue is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegOffsetValue() (field.PegOffsetValue, errors.MessageRejectError) {
	var f field.PegOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//PegMoveType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegMoveType() (field.PegMoveType, errors.MessageRejectError) {
	var f field.PegMoveType
	err := m.Body.Get(&f)
	return f, err
}

//PegOffsetType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegOffsetType() (field.PegOffsetType, errors.MessageRejectError) {
	var f field.PegOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//PegLimitType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegLimitType() (field.PegLimitType, errors.MessageRejectError) {
	var f field.PegLimitType
	err := m.Body.Get(&f)
	return f, err
}

//PegRoundDirection is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegRoundDirection() (field.PegRoundDirection, errors.MessageRejectError) {
	var f field.PegRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//PegScope is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegScope() (field.PegScope, errors.MessageRejectError) {
	var f field.PegScope
	err := m.Body.Get(&f)
	return f, err
}

//PegPriceType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegPriceType() (field.PegPriceType, errors.MessageRejectError) {
	var f field.PegPriceType
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityIDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegSecurityIDSource() (field.PegSecurityIDSource, errors.MessageRejectError) {
	var f field.PegSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegSecurityID() (field.PegSecurityID, errors.MessageRejectError) {
	var f field.PegSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//PegSymbol is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegSymbol() (field.PegSymbol, errors.MessageRejectError) {
	var f field.PegSymbol
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityDesc is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegSecurityDesc() (field.PegSecurityDesc, errors.MessageRejectError) {
	var f field.PegSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionInst() (field.DiscretionInst, errors.MessageRejectError) {
	var f field.DiscretionInst
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetValue is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionOffsetValue() (field.DiscretionOffsetValue, errors.MessageRejectError) {
	var f field.DiscretionOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionMoveType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionMoveType() (field.DiscretionMoveType, errors.MessageRejectError) {
	var f field.DiscretionMoveType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionOffsetType() (field.DiscretionOffsetType, errors.MessageRejectError) {
	var f field.DiscretionOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionLimitType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionLimitType() (field.DiscretionLimitType, errors.MessageRejectError) {
	var f field.DiscretionLimitType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionRoundDirection is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionRoundDirection() (field.DiscretionRoundDirection, errors.MessageRejectError) {
	var f field.DiscretionRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionScope is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionScope() (field.DiscretionScope, errors.MessageRejectError) {
	var f field.DiscretionScope
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategy is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TargetStrategy() (field.TargetStrategy, errors.MessageRejectError) {
	var f field.TargetStrategy
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategyParameters is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TargetStrategyParameters() (field.TargetStrategyParameters, errors.MessageRejectError) {
	var f field.TargetStrategyParameters
	err := m.Body.Get(&f)
	return f, err
}

//ParticipationRate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ParticipationRate() (field.ParticipationRate, errors.MessageRejectError) {
	var f field.ParticipationRate
	err := m.Body.Get(&f)
	return f, err
}

//CancellationRights is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CancellationRights() (field.CancellationRights, errors.MessageRejectError) {
	var f field.CancellationRights
	err := m.Body.Get(&f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MoneyLaunderingStatus() (field.MoneyLaunderingStatus, errors.MessageRejectError) {
	var f field.MoneyLaunderingStatus
	err := m.Body.Get(&f)
	return f, err
}

//RegistID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RegistID() (field.RegistID, errors.MessageRejectError) {
	var f field.RegistID
	err := m.Body.Get(&f)
	return f, err
}

//Designation is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Designation() (field.Designation, errors.MessageRejectError) {
	var f field.Designation
	err := m.Body.Get(&f)
	return f, err
}

//NoStrategyParameters is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoStrategyParameters() (field.NoStrategyParameters, errors.MessageRejectError) {
	var f field.NoStrategyParameters
	err := m.Body.Get(&f)
	return f, err
}

//ManualOrderIndicator is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ManualOrderIndicator() (field.ManualOrderIndicator, errors.MessageRejectError) {
	var f field.ManualOrderIndicator
	err := m.Body.Get(&f)
	return f, err
}

//CustDirectedOrder is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CustDirectedOrder() (field.CustDirectedOrder, errors.MessageRejectError) {
	var f field.CustDirectedOrder
	err := m.Body.Get(&f)
	return f, err
}

//ReceivedDeptID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ReceivedDeptID() (field.ReceivedDeptID, errors.MessageRejectError) {
	var f field.ReceivedDeptID
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderHandlingInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CustOrderHandlingInst() (field.CustOrderHandlingInst, errors.MessageRejectError) {
	var f field.CustOrderHandlingInst
	err := m.Body.Get(&f)
	return f, err
}

//OrderHandlingInstSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderHandlingInstSource() (field.OrderHandlingInstSource, errors.MessageRejectError) {
	var f field.OrderHandlingInstSource
	err := m.Body.Get(&f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoTrdRegTimestamps() (field.NoTrdRegTimestamps, errors.MessageRejectError) {
	var f field.NoTrdRegTimestamps
	err := m.Body.Get(&f)
	return f, err
}

//MatchIncrement is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MatchIncrement() (field.MatchIncrement, errors.MessageRejectError) {
	var f field.MatchIncrement
	err := m.Body.Get(&f)
	return f, err
}

//MaxPriceLevels is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaxPriceLevels() (field.MaxPriceLevels, errors.MessageRejectError) {
	var f field.MaxPriceLevels
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryDisplayQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecondaryDisplayQty() (field.SecondaryDisplayQty, errors.MessageRejectError) {
	var f field.SecondaryDisplayQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayWhen is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DisplayWhen() (field.DisplayWhen, errors.MessageRejectError) {
	var f field.DisplayWhen
	err := m.Body.Get(&f)
	return f, err
}

//DisplayMethod is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DisplayMethod() (field.DisplayMethod, errors.MessageRejectError) {
	var f field.DisplayMethod
	err := m.Body.Get(&f)
	return f, err
}

//DisplayLowQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DisplayLowQty() (field.DisplayLowQty, errors.MessageRejectError) {
	var f field.DisplayLowQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayHighQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DisplayHighQty() (field.DisplayHighQty, errors.MessageRejectError) {
	var f field.DisplayHighQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayMinIncr is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DisplayMinIncr() (field.DisplayMinIncr, errors.MessageRejectError) {
	var f field.DisplayMinIncr
	err := m.Body.Get(&f)
	return f, err
}

//RefreshQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RefreshQty() (field.RefreshQty, errors.MessageRejectError) {
	var f field.RefreshQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DisplayQty() (field.DisplayQty, errors.MessageRejectError) {
	var f field.DisplayQty
	err := m.Body.Get(&f)
	return f, err
}

//PriceProtectionScope is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PriceProtectionScope() (field.PriceProtectionScope, errors.MessageRejectError) {
	var f field.PriceProtectionScope
	err := m.Body.Get(&f)
	return f, err
}

//TriggerType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerType() (field.TriggerType, errors.MessageRejectError) {
	var f field.TriggerType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerAction is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerAction() (field.TriggerAction, errors.MessageRejectError) {
	var f field.TriggerAction
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPrice is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerPrice() (field.TriggerPrice, errors.MessageRejectError) {
	var f field.TriggerPrice
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSymbol is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerSymbol() (field.TriggerSymbol, errors.MessageRejectError) {
	var f field.TriggerSymbol
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerSecurityID() (field.TriggerSecurityID, errors.MessageRejectError) {
	var f field.TriggerSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityIDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerSecurityIDSource() (field.TriggerSecurityIDSource, errors.MessageRejectError) {
	var f field.TriggerSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityDesc is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerSecurityDesc() (field.TriggerSecurityDesc, errors.MessageRejectError) {
	var f field.TriggerSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerPriceType() (field.TriggerPriceType, errors.MessageRejectError) {
	var f field.TriggerPriceType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceTypeScope is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerPriceTypeScope() (field.TriggerPriceTypeScope, errors.MessageRejectError) {
	var f field.TriggerPriceTypeScope
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceDirection is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerPriceDirection() (field.TriggerPriceDirection, errors.MessageRejectError) {
	var f field.TriggerPriceDirection
	err := m.Body.Get(&f)
	return f, err
}

//TriggerNewPrice is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerNewPrice() (field.TriggerNewPrice, errors.MessageRejectError) {
	var f field.TriggerNewPrice
	err := m.Body.Get(&f)
	return f, err
}

//TriggerOrderType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerOrderType() (field.TriggerOrderType, errors.MessageRejectError) {
	var f field.TriggerOrderType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerNewQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerNewQty() (field.TriggerNewQty, errors.MessageRejectError) {
	var f field.TriggerNewQty
	err := m.Body.Get(&f)
	return f, err
}

//TriggerTradingSessionID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerTradingSessionID() (field.TriggerTradingSessionID, errors.MessageRejectError) {
	var f field.TriggerTradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TriggerTradingSessionSubID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerTradingSessionSubID() (field.TriggerTradingSessionSubID, errors.MessageRejectError) {
	var f field.TriggerTradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//PreTradeAnonymity is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PreTradeAnonymity() (field.PreTradeAnonymity, errors.MessageRejectError) {
	var f field.PreTradeAnonymity
	err := m.Body.Get(&f)
	return f, err
}

//RefOrderID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RefOrderID() (field.RefOrderID, errors.MessageRejectError) {
	var f field.RefOrderID
	err := m.Body.Get(&f)
	return f, err
}

//RefOrderIDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RefOrderIDSource() (field.RefOrderIDSource, errors.MessageRejectError) {
	var f field.RefOrderIDSource
	err := m.Body.Get(&f)
	return f, err
}

//ExDestinationIDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExDestinationIDSource() (field.ExDestinationIDSource, errors.MessageRejectError) {
	var f field.ExDestinationIDSource
	err := m.Body.Get(&f)
	return f, err
}
