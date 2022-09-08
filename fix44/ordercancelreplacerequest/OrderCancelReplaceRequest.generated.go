package ordercancelreplacerequest

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/alpacahq/quickfix"
	"github.com/alpacahq/quickfix/enum"
	"github.com/alpacahq/quickfix/field"
	"github.com/alpacahq/quickfix/fix44"
	"github.com/alpacahq/quickfix/tag"
)

// OrderCancelReplaceRequest is the fix44 OrderCancelReplaceRequest type, MsgType = G
type OrderCancelReplaceRequest struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a OrderCancelReplaceRequest from a quickfix.Message instance
func FromMessage(m *quickfix.Message) OrderCancelReplaceRequest {
	return OrderCancelReplaceRequest{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance
func (m OrderCancelReplaceRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a OrderCancelReplaceRequest initialized with the required fields for OrderCancelReplaceRequest
func New(origclordid field.OrigClOrdIDField, clordid field.ClOrdIDField, side field.SideField, transacttime field.TransactTimeField, ordtype field.OrdTypeField) (m OrderCancelReplaceRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("G"))
	m.Set(origclordid)
	m.Set(clordid)
	m.Set(side)
	m.Set(transacttime)
	m.Set(ordtype)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "G", r
}

// SetAccount sets Account, Tag 1
func (m OrderCancelReplaceRequest) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

// SetClOrdID sets ClOrdID, Tag 11
func (m OrderCancelReplaceRequest) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

// SetCommission sets Commission, Tag 12
func (m OrderCancelReplaceRequest) SetCommission(value decimal.Decimal, scale int32) {
	m.Set(field.NewCommission(value, scale))
}

// SetCommType sets CommType, Tag 13
func (m OrderCancelReplaceRequest) SetCommType(v enum.CommType) {
	m.Set(field.NewCommType(v))
}

// SetCurrency sets Currency, Tag 15
func (m OrderCancelReplaceRequest) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

// SetExecInst sets ExecInst, Tag 18
func (m OrderCancelReplaceRequest) SetExecInst(v enum.ExecInst) {
	m.Set(field.NewExecInst(v))
}

// SetHandlInst sets HandlInst, Tag 21
func (m OrderCancelReplaceRequest) SetHandlInst(v enum.HandlInst) {
	m.Set(field.NewHandlInst(v))
}

// SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m OrderCancelReplaceRequest) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

// SetOrderID sets OrderID, Tag 37
func (m OrderCancelReplaceRequest) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

// SetOrderQty sets OrderQty, Tag 38
func (m OrderCancelReplaceRequest) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

// SetOrdType sets OrdType, Tag 40
func (m OrderCancelReplaceRequest) SetOrdType(v enum.OrdType) {
	m.Set(field.NewOrdType(v))
}

// SetOrigClOrdID sets OrigClOrdID, Tag 41
func (m OrderCancelReplaceRequest) SetOrigClOrdID(v string) {
	m.Set(field.NewOrigClOrdID(v))
}

// SetPrice sets Price, Tag 44
func (m OrderCancelReplaceRequest) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

// SetSecurityID sets SecurityID, Tag 48
func (m OrderCancelReplaceRequest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetSide sets Side, Tag 54
func (m OrderCancelReplaceRequest) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

// SetSymbol sets Symbol, Tag 55
func (m OrderCancelReplaceRequest) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetText sets Text, Tag 58
func (m OrderCancelReplaceRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTimeInForce sets TimeInForce, Tag 59
func (m OrderCancelReplaceRequest) SetTimeInForce(v enum.TimeInForce) {
	m.Set(field.NewTimeInForce(v))
}

// SetTransactTime sets TransactTime, Tag 60
func (m OrderCancelReplaceRequest) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetSettlType sets SettlType, Tag 63
func (m OrderCancelReplaceRequest) SetSettlType(v enum.SettlType) {
	m.Set(field.NewSettlType(v))
}

// SetSettlDate sets SettlDate, Tag 64
func (m OrderCancelReplaceRequest) SetSettlDate(v string) {
	m.Set(field.NewSettlDate(v))
}

// SetSymbolSfx sets SymbolSfx, Tag 65
func (m OrderCancelReplaceRequest) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

// SetListID sets ListID, Tag 66
func (m OrderCancelReplaceRequest) SetListID(v string) {
	m.Set(field.NewListID(v))
}

// SetAllocID sets AllocID, Tag 70
func (m OrderCancelReplaceRequest) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

// SetTradeDate sets TradeDate, Tag 75
func (m OrderCancelReplaceRequest) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

// SetPositionEffect sets PositionEffect, Tag 77
func (m OrderCancelReplaceRequest) SetPositionEffect(v enum.PositionEffect) {
	m.Set(field.NewPositionEffect(v))
}

// SetNoAllocs sets NoAllocs, Tag 78
func (m OrderCancelReplaceRequest) SetNoAllocs(f NoAllocsRepeatingGroup) {
	m.SetGroup(f)
}

// SetStopPx sets StopPx, Tag 99
func (m OrderCancelReplaceRequest) SetStopPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewStopPx(value, scale))
}

// SetExDestination sets ExDestination, Tag 100
func (m OrderCancelReplaceRequest) SetExDestination(v enum.ExDestination) {
	m.Set(field.NewExDestination(v))
}

// SetIssuer sets Issuer, Tag 106
func (m OrderCancelReplaceRequest) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

// SetSecurityDesc sets SecurityDesc, Tag 107
func (m OrderCancelReplaceRequest) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

// SetMinQty sets MinQty, Tag 110
func (m OrderCancelReplaceRequest) SetMinQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinQty(value, scale))
}

// SetMaxFloor sets MaxFloor, Tag 111
func (m OrderCancelReplaceRequest) SetMaxFloor(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxFloor(value, scale))
}

// SetLocateReqd sets LocateReqd, Tag 114
func (m OrderCancelReplaceRequest) SetLocateReqd(v bool) {
	m.Set(field.NewLocateReqd(v))
}

// SetSettlCurrency sets SettlCurrency, Tag 120
func (m OrderCancelReplaceRequest) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

// SetForexReq sets ForexReq, Tag 121
func (m OrderCancelReplaceRequest) SetForexReq(v bool) {
	m.Set(field.NewForexReq(v))
}

// SetExpireTime sets ExpireTime, Tag 126
func (m OrderCancelReplaceRequest) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

// SetCashOrderQty sets CashOrderQty, Tag 152
func (m OrderCancelReplaceRequest) SetCashOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCashOrderQty(value, scale))
}

// SetSecurityType sets SecurityType, Tag 167
func (m OrderCancelReplaceRequest) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

// SetEffectiveTime sets EffectiveTime, Tag 168
func (m OrderCancelReplaceRequest) SetEffectiveTime(v time.Time) {
	m.Set(field.NewEffectiveTime(v))
}

// SetOrderQty2 sets OrderQty2, Tag 192
func (m OrderCancelReplaceRequest) SetOrderQty2(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty2(value, scale))
}

// SetSettlDate2 sets SettlDate2, Tag 193
func (m OrderCancelReplaceRequest) SetSettlDate2(v string) {
	m.Set(field.NewSettlDate2(v))
}

// SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m OrderCancelReplaceRequest) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

// SetStrikePrice sets StrikePrice, Tag 202
func (m OrderCancelReplaceRequest) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

// SetCoveredOrUncovered sets CoveredOrUncovered, Tag 203
func (m OrderCancelReplaceRequest) SetCoveredOrUncovered(v enum.CoveredOrUncovered) {
	m.Set(field.NewCoveredOrUncovered(v))
}

// SetOptAttribute sets OptAttribute, Tag 206
func (m OrderCancelReplaceRequest) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

// SetSecurityExchange sets SecurityExchange, Tag 207
func (m OrderCancelReplaceRequest) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

// SetMaxShow sets MaxShow, Tag 210
func (m OrderCancelReplaceRequest) SetMaxShow(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxShow(value, scale))
}

// SetPegOffsetValue sets PegOffsetValue, Tag 211
func (m OrderCancelReplaceRequest) SetPegOffsetValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewPegOffsetValue(value, scale))
}

// SetSpread sets Spread, Tag 218
func (m OrderCancelReplaceRequest) SetSpread(value decimal.Decimal, scale int32) {
	m.Set(field.NewSpread(value, scale))
}

// SetBenchmarkCurveCurrency sets BenchmarkCurveCurrency, Tag 220
func (m OrderCancelReplaceRequest) SetBenchmarkCurveCurrency(v string) {
	m.Set(field.NewBenchmarkCurveCurrency(v))
}

// SetBenchmarkCurveName sets BenchmarkCurveName, Tag 221
func (m OrderCancelReplaceRequest) SetBenchmarkCurveName(v enum.BenchmarkCurveName) {
	m.Set(field.NewBenchmarkCurveName(v))
}

// SetBenchmarkCurvePoint sets BenchmarkCurvePoint, Tag 222
func (m OrderCancelReplaceRequest) SetBenchmarkCurvePoint(v string) {
	m.Set(field.NewBenchmarkCurvePoint(v))
}

// SetCouponRate sets CouponRate, Tag 223
func (m OrderCancelReplaceRequest) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

// SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m OrderCancelReplaceRequest) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

// SetIssueDate sets IssueDate, Tag 225
func (m OrderCancelReplaceRequest) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

// SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m OrderCancelReplaceRequest) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

// SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m OrderCancelReplaceRequest) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

// SetFactor sets Factor, Tag 228
func (m OrderCancelReplaceRequest) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

// SetTradeOriginationDate sets TradeOriginationDate, Tag 229
func (m OrderCancelReplaceRequest) SetTradeOriginationDate(v string) {
	m.Set(field.NewTradeOriginationDate(v))
}

// SetContractMultiplier sets ContractMultiplier, Tag 231
func (m OrderCancelReplaceRequest) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

// SetYieldType sets YieldType, Tag 235
func (m OrderCancelReplaceRequest) SetYieldType(v enum.YieldType) {
	m.Set(field.NewYieldType(v))
}

// SetYield sets Yield, Tag 236
func (m OrderCancelReplaceRequest) SetYield(value decimal.Decimal, scale int32) {
	m.Set(field.NewYield(value, scale))
}

// SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m OrderCancelReplaceRequest) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

// SetRedemptionDate sets RedemptionDate, Tag 240
func (m OrderCancelReplaceRequest) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

// SetCreditRating sets CreditRating, Tag 255
func (m OrderCancelReplaceRequest) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

// SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m OrderCancelReplaceRequest) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

// SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m OrderCancelReplaceRequest) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

// SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m OrderCancelReplaceRequest) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

// SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m OrderCancelReplaceRequest) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m OrderCancelReplaceRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355
func (m OrderCancelReplaceRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetComplianceID sets ComplianceID, Tag 376
func (m OrderCancelReplaceRequest) SetComplianceID(v string) {
	m.Set(field.NewComplianceID(v))
}

// SetSolicitedFlag sets SolicitedFlag, Tag 377
func (m OrderCancelReplaceRequest) SetSolicitedFlag(v bool) {
	m.Set(field.NewSolicitedFlag(v))
}

// SetNoTradingSessions sets NoTradingSessions, Tag 386
func (m OrderCancelReplaceRequest) SetNoTradingSessions(f NoTradingSessionsRepeatingGroup) {
	m.SetGroup(f)
}

// SetDiscretionInst sets DiscretionInst, Tag 388
func (m OrderCancelReplaceRequest) SetDiscretionInst(v enum.DiscretionInst) {
	m.Set(field.NewDiscretionInst(v))
}

// SetDiscretionOffsetValue sets DiscretionOffsetValue, Tag 389
func (m OrderCancelReplaceRequest) SetDiscretionOffsetValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewDiscretionOffsetValue(value, scale))
}

// SetPriceType sets PriceType, Tag 423
func (m OrderCancelReplaceRequest) SetPriceType(v enum.PriceType) {
	m.Set(field.NewPriceType(v))
}

// SetGTBookingInst sets GTBookingInst, Tag 427
func (m OrderCancelReplaceRequest) SetGTBookingInst(v enum.GTBookingInst) {
	m.Set(field.NewGTBookingInst(v))
}

// SetExpireDate sets ExpireDate, Tag 432
func (m OrderCancelReplaceRequest) SetExpireDate(v string) {
	m.Set(field.NewExpireDate(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m OrderCancelReplaceRequest) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m OrderCancelReplaceRequest) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetProduct sets Product, Tag 460
func (m OrderCancelReplaceRequest) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

// SetCFICode sets CFICode, Tag 461
func (m OrderCancelReplaceRequest) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

// SetRoundingDirection sets RoundingDirection, Tag 468
func (m OrderCancelReplaceRequest) SetRoundingDirection(v enum.RoundingDirection) {
	m.Set(field.NewRoundingDirection(v))
}

// SetRoundingModulus sets RoundingModulus, Tag 469
func (m OrderCancelReplaceRequest) SetRoundingModulus(value decimal.Decimal, scale int32) {
	m.Set(field.NewRoundingModulus(value, scale))
}

// SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m OrderCancelReplaceRequest) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

// SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m OrderCancelReplaceRequest) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

// SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m OrderCancelReplaceRequest) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

// SetCommCurrency sets CommCurrency, Tag 479
func (m OrderCancelReplaceRequest) SetCommCurrency(v string) {
	m.Set(field.NewCommCurrency(v))
}

// SetCancellationRights sets CancellationRights, Tag 480
func (m OrderCancelReplaceRequest) SetCancellationRights(v enum.CancellationRights) {
	m.Set(field.NewCancellationRights(v))
}

// SetMoneyLaunderingStatus sets MoneyLaunderingStatus, Tag 481
func (m OrderCancelReplaceRequest) SetMoneyLaunderingStatus(v enum.MoneyLaunderingStatus) {
	m.Set(field.NewMoneyLaunderingStatus(v))
}

// SetDesignation sets Designation, Tag 494
func (m OrderCancelReplaceRequest) SetDesignation(v string) {
	m.Set(field.NewDesignation(v))
}

// SetFundRenewWaiv sets FundRenewWaiv, Tag 497
func (m OrderCancelReplaceRequest) SetFundRenewWaiv(v enum.FundRenewWaiv) {
	m.Set(field.NewFundRenewWaiv(v))
}

// SetRegistID sets RegistID, Tag 513
func (m OrderCancelReplaceRequest) SetRegistID(v string) {
	m.Set(field.NewRegistID(v))
}

// SetOrderPercent sets OrderPercent, Tag 516
func (m OrderCancelReplaceRequest) SetOrderPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderPercent(value, scale))
}

// SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526
func (m OrderCancelReplaceRequest) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

// SetOrderCapacity sets OrderCapacity, Tag 528
func (m OrderCancelReplaceRequest) SetOrderCapacity(v enum.OrderCapacity) {
	m.Set(field.NewOrderCapacity(v))
}

// SetOrderRestrictions sets OrderRestrictions, Tag 529
func (m OrderCancelReplaceRequest) SetOrderRestrictions(v enum.OrderRestrictions) {
	m.Set(field.NewOrderRestrictions(v))
}

// SetMaturityDate sets MaturityDate, Tag 541
func (m OrderCancelReplaceRequest) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

// SetInstrRegistry sets InstrRegistry, Tag 543
func (m OrderCancelReplaceRequest) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

// SetCashMargin sets CashMargin, Tag 544
func (m OrderCancelReplaceRequest) SetCashMargin(v enum.CashMargin) {
	m.Set(field.NewCashMargin(v))
}

// SetAccountType sets AccountType, Tag 581
func (m OrderCancelReplaceRequest) SetAccountType(v enum.AccountType) {
	m.Set(field.NewAccountType(v))
}

// SetCustOrderCapacity sets CustOrderCapacity, Tag 582
func (m OrderCancelReplaceRequest) SetCustOrderCapacity(v enum.CustOrderCapacity) {
	m.Set(field.NewCustOrderCapacity(v))
}

// SetClOrdLinkID sets ClOrdLinkID, Tag 583
func (m OrderCancelReplaceRequest) SetClOrdLinkID(v string) {
	m.Set(field.NewClOrdLinkID(v))
}

// SetOrigOrdModTime sets OrigOrdModTime, Tag 586
func (m OrderCancelReplaceRequest) SetOrigOrdModTime(v time.Time) {
	m.Set(field.NewOrigOrdModTime(v))
}

// SetDayBookingInst sets DayBookingInst, Tag 589
func (m OrderCancelReplaceRequest) SetDayBookingInst(v enum.DayBookingInst) {
	m.Set(field.NewDayBookingInst(v))
}

// SetBookingUnit sets BookingUnit, Tag 590
func (m OrderCancelReplaceRequest) SetBookingUnit(v enum.BookingUnit) {
	m.Set(field.NewBookingUnit(v))
}

// SetPreallocMethod sets PreallocMethod, Tag 591
func (m OrderCancelReplaceRequest) SetPreallocMethod(v enum.PreallocMethod) {
	m.Set(field.NewPreallocMethod(v))
}

// SetClearingFeeIndicator sets ClearingFeeIndicator, Tag 635
func (m OrderCancelReplaceRequest) SetClearingFeeIndicator(v enum.ClearingFeeIndicator) {
	m.Set(field.NewClearingFeeIndicator(v))
}

// SetPrice2 sets Price2, Tag 640
func (m OrderCancelReplaceRequest) SetPrice2(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice2(value, scale))
}

// SetAcctIDSource sets AcctIDSource, Tag 660
func (m OrderCancelReplaceRequest) SetAcctIDSource(v enum.AcctIDSource) {
	m.Set(field.NewAcctIDSource(v))
}

// SetBenchmarkPrice sets BenchmarkPrice, Tag 662
func (m OrderCancelReplaceRequest) SetBenchmarkPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewBenchmarkPrice(value, scale))
}

// SetBenchmarkPriceType sets BenchmarkPriceType, Tag 663
func (m OrderCancelReplaceRequest) SetBenchmarkPriceType(v int) {
	m.Set(field.NewBenchmarkPriceType(v))
}

// SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m OrderCancelReplaceRequest) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

// SetPool sets Pool, Tag 691
func (m OrderCancelReplaceRequest) SetPool(v string) {
	m.Set(field.NewPool(v))
}

// SetYieldRedemptionDate sets YieldRedemptionDate, Tag 696
func (m OrderCancelReplaceRequest) SetYieldRedemptionDate(v string) {
	m.Set(field.NewYieldRedemptionDate(v))
}

// SetYieldRedemptionPrice sets YieldRedemptionPrice, Tag 697
func (m OrderCancelReplaceRequest) SetYieldRedemptionPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewYieldRedemptionPrice(value, scale))
}

// SetYieldRedemptionPriceType sets YieldRedemptionPriceType, Tag 698
func (m OrderCancelReplaceRequest) SetYieldRedemptionPriceType(v int) {
	m.Set(field.NewYieldRedemptionPriceType(v))
}

// SetBenchmarkSecurityID sets BenchmarkSecurityID, Tag 699
func (m OrderCancelReplaceRequest) SetBenchmarkSecurityID(v string) {
	m.Set(field.NewBenchmarkSecurityID(v))
}

// SetYieldCalcDate sets YieldCalcDate, Tag 701
func (m OrderCancelReplaceRequest) SetYieldCalcDate(v string) {
	m.Set(field.NewYieldCalcDate(v))
}

// SetNoUnderlyings sets NoUnderlyings, Tag 711
func (m OrderCancelReplaceRequest) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

// SetBenchmarkSecurityIDSource sets BenchmarkSecurityIDSource, Tag 761
func (m OrderCancelReplaceRequest) SetBenchmarkSecurityIDSource(v string) {
	m.Set(field.NewBenchmarkSecurityIDSource(v))
}

// SetSecuritySubType sets SecuritySubType, Tag 762
func (m OrderCancelReplaceRequest) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

// SetBookingType sets BookingType, Tag 775
func (m OrderCancelReplaceRequest) SetBookingType(v enum.BookingType) {
	m.Set(field.NewBookingType(v))
}

// SetTerminationType sets TerminationType, Tag 788
func (m OrderCancelReplaceRequest) SetTerminationType(v enum.TerminationType) {
	m.Set(field.NewTerminationType(v))
}

// SetPegMoveType sets PegMoveType, Tag 835
func (m OrderCancelReplaceRequest) SetPegMoveType(v enum.PegMoveType) {
	m.Set(field.NewPegMoveType(v))
}

// SetPegOffsetType sets PegOffsetType, Tag 836
func (m OrderCancelReplaceRequest) SetPegOffsetType(v enum.PegOffsetType) {
	m.Set(field.NewPegOffsetType(v))
}

// SetPegLimitType sets PegLimitType, Tag 837
func (m OrderCancelReplaceRequest) SetPegLimitType(v enum.PegLimitType) {
	m.Set(field.NewPegLimitType(v))
}

// SetPegRoundDirection sets PegRoundDirection, Tag 838
func (m OrderCancelReplaceRequest) SetPegRoundDirection(v enum.PegRoundDirection) {
	m.Set(field.NewPegRoundDirection(v))
}

// SetPegScope sets PegScope, Tag 840
func (m OrderCancelReplaceRequest) SetPegScope(v enum.PegScope) {
	m.Set(field.NewPegScope(v))
}

// SetDiscretionMoveType sets DiscretionMoveType, Tag 841
func (m OrderCancelReplaceRequest) SetDiscretionMoveType(v enum.DiscretionMoveType) {
	m.Set(field.NewDiscretionMoveType(v))
}

// SetDiscretionOffsetType sets DiscretionOffsetType, Tag 842
func (m OrderCancelReplaceRequest) SetDiscretionOffsetType(v enum.DiscretionOffsetType) {
	m.Set(field.NewDiscretionOffsetType(v))
}

// SetDiscretionLimitType sets DiscretionLimitType, Tag 843
func (m OrderCancelReplaceRequest) SetDiscretionLimitType(v enum.DiscretionLimitType) {
	m.Set(field.NewDiscretionLimitType(v))
}

// SetDiscretionRoundDirection sets DiscretionRoundDirection, Tag 844
func (m OrderCancelReplaceRequest) SetDiscretionRoundDirection(v enum.DiscretionRoundDirection) {
	m.Set(field.NewDiscretionRoundDirection(v))
}

// SetDiscretionScope sets DiscretionScope, Tag 846
func (m OrderCancelReplaceRequest) SetDiscretionScope(v enum.DiscretionScope) {
	m.Set(field.NewDiscretionScope(v))
}

// SetTargetStrategy sets TargetStrategy, Tag 847
func (m OrderCancelReplaceRequest) SetTargetStrategy(v enum.TargetStrategy) {
	m.Set(field.NewTargetStrategy(v))
}

// SetTargetStrategyParameters sets TargetStrategyParameters, Tag 848
func (m OrderCancelReplaceRequest) SetTargetStrategyParameters(v string) {
	m.Set(field.NewTargetStrategyParameters(v))
}

// SetParticipationRate sets ParticipationRate, Tag 849
func (m OrderCancelReplaceRequest) SetParticipationRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewParticipationRate(value, scale))
}

// SetQtyType sets QtyType, Tag 854
func (m OrderCancelReplaceRequest) SetQtyType(v enum.QtyType) {
	m.Set(field.NewQtyType(v))
}

// SetNoEvents sets NoEvents, Tag 864
func (m OrderCancelReplaceRequest) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

// SetDatedDate sets DatedDate, Tag 873
func (m OrderCancelReplaceRequest) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

// SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m OrderCancelReplaceRequest) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

// SetCPProgram sets CPProgram, Tag 875
func (m OrderCancelReplaceRequest) SetCPProgram(v enum.CPProgram) {
	m.Set(field.NewCPProgram(v))
}

// SetCPRegType sets CPRegType, Tag 876
func (m OrderCancelReplaceRequest) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

// SetMarginRatio sets MarginRatio, Tag 898
func (m OrderCancelReplaceRequest) SetMarginRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewMarginRatio(value, scale))
}

// SetAgreementDesc sets AgreementDesc, Tag 913
func (m OrderCancelReplaceRequest) SetAgreementDesc(v string) {
	m.Set(field.NewAgreementDesc(v))
}

// SetAgreementID sets AgreementID, Tag 914
func (m OrderCancelReplaceRequest) SetAgreementID(v string) {
	m.Set(field.NewAgreementID(v))
}

// SetAgreementDate sets AgreementDate, Tag 915
func (m OrderCancelReplaceRequest) SetAgreementDate(v string) {
	m.Set(field.NewAgreementDate(v))
}

// SetStartDate sets StartDate, Tag 916
func (m OrderCancelReplaceRequest) SetStartDate(v string) {
	m.Set(field.NewStartDate(v))
}

// SetEndDate sets EndDate, Tag 917
func (m OrderCancelReplaceRequest) SetEndDate(v string) {
	m.Set(field.NewEndDate(v))
}

// SetAgreementCurrency sets AgreementCurrency, Tag 918
func (m OrderCancelReplaceRequest) SetAgreementCurrency(v string) {
	m.Set(field.NewAgreementCurrency(v))
}

// SetDeliveryType sets DeliveryType, Tag 919
func (m OrderCancelReplaceRequest) SetDeliveryType(v enum.DeliveryType) {
	m.Set(field.NewDeliveryType(v))
}

// SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m OrderCancelReplaceRequest) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

// GetAccount gets Account, Tag 1
func (m OrderCancelReplaceRequest) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetClOrdID gets ClOrdID, Tag 11
func (m OrderCancelReplaceRequest) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCommission gets Commission, Tag 12
func (m OrderCancelReplaceRequest) GetCommission() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CommissionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCommType gets CommType, Tag 13
func (m OrderCancelReplaceRequest) GetCommType() (v enum.CommType, err quickfix.MessageRejectError) {
	var f field.CommTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCurrency gets Currency, Tag 15
func (m OrderCancelReplaceRequest) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExecInst gets ExecInst, Tag 18
func (m OrderCancelReplaceRequest) GetExecInst() (v enum.ExecInst, err quickfix.MessageRejectError) {
	var f field.ExecInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetHandlInst gets HandlInst, Tag 21
func (m OrderCancelReplaceRequest) GetHandlInst() (v enum.HandlInst, err quickfix.MessageRejectError) {
	var f field.HandlInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m OrderCancelReplaceRequest) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderID gets OrderID, Tag 37
func (m OrderCancelReplaceRequest) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderQty gets OrderQty, Tag 38
func (m OrderCancelReplaceRequest) GetOrderQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrdType gets OrdType, Tag 40
func (m OrderCancelReplaceRequest) GetOrdType() (v enum.OrdType, err quickfix.MessageRejectError) {
	var f field.OrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrigClOrdID gets OrigClOrdID, Tag 41
func (m OrderCancelReplaceRequest) GetOrigClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.OrigClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPrice gets Price, Tag 44
func (m OrderCancelReplaceRequest) GetPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48
func (m OrderCancelReplaceRequest) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSide gets Side, Tag 54
func (m OrderCancelReplaceRequest) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55
func (m OrderCancelReplaceRequest) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetText gets Text, Tag 58
func (m OrderCancelReplaceRequest) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTimeInForce gets TimeInForce, Tag 59
func (m OrderCancelReplaceRequest) GetTimeInForce() (v enum.TimeInForce, err quickfix.MessageRejectError) {
	var f field.TimeInForceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60
func (m OrderCancelReplaceRequest) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlType gets SettlType, Tag 63
func (m OrderCancelReplaceRequest) GetSettlType() (v enum.SettlType, err quickfix.MessageRejectError) {
	var f field.SettlTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlDate gets SettlDate, Tag 64
func (m OrderCancelReplaceRequest) GetSettlDate() (v string, err quickfix.MessageRejectError) {
	var f field.SettlDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbolSfx gets SymbolSfx, Tag 65
func (m OrderCancelReplaceRequest) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetListID gets ListID, Tag 66
func (m OrderCancelReplaceRequest) GetListID() (v string, err quickfix.MessageRejectError) {
	var f field.ListIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocID gets AllocID, Tag 70
func (m OrderCancelReplaceRequest) GetAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.AllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradeDate gets TradeDate, Tag 75
func (m OrderCancelReplaceRequest) GetTradeDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPositionEffect gets PositionEffect, Tag 77
func (m OrderCancelReplaceRequest) GetPositionEffect() (v enum.PositionEffect, err quickfix.MessageRejectError) {
	var f field.PositionEffectField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoAllocs gets NoAllocs, Tag 78
func (m OrderCancelReplaceRequest) GetNoAllocs() (NoAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetStopPx gets StopPx, Tag 99
func (m OrderCancelReplaceRequest) GetStopPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StopPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExDestination gets ExDestination, Tag 100
func (m OrderCancelReplaceRequest) GetExDestination() (v enum.ExDestination, err quickfix.MessageRejectError) {
	var f field.ExDestinationField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssuer gets Issuer, Tag 106
func (m OrderCancelReplaceRequest) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityDesc gets SecurityDesc, Tag 107
func (m OrderCancelReplaceRequest) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMinQty gets MinQty, Tag 110
func (m OrderCancelReplaceRequest) GetMinQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaxFloor gets MaxFloor, Tag 111
func (m OrderCancelReplaceRequest) GetMaxFloor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MaxFloorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLocateReqd gets LocateReqd, Tag 114
func (m OrderCancelReplaceRequest) GetLocateReqd() (v bool, err quickfix.MessageRejectError) {
	var f field.LocateReqdField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlCurrency gets SettlCurrency, Tag 120
func (m OrderCancelReplaceRequest) GetSettlCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.SettlCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetForexReq gets ForexReq, Tag 121
func (m OrderCancelReplaceRequest) GetForexReq() (v bool, err quickfix.MessageRejectError) {
	var f field.ForexReqField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExpireTime gets ExpireTime, Tag 126
func (m OrderCancelReplaceRequest) GetExpireTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ExpireTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCashOrderQty gets CashOrderQty, Tag 152
func (m OrderCancelReplaceRequest) GetCashOrderQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CashOrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityType gets SecurityType, Tag 167
func (m OrderCancelReplaceRequest) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEffectiveTime gets EffectiveTime, Tag 168
func (m OrderCancelReplaceRequest) GetEffectiveTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.EffectiveTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderQty2 gets OrderQty2, Tag 192
func (m OrderCancelReplaceRequest) GetOrderQty2() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderQty2Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlDate2 gets SettlDate2, Tag 193
func (m OrderCancelReplaceRequest) GetSettlDate2() (v string, err quickfix.MessageRejectError) {
	var f field.SettlDate2Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m OrderCancelReplaceRequest) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePrice gets StrikePrice, Tag 202
func (m OrderCancelReplaceRequest) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCoveredOrUncovered gets CoveredOrUncovered, Tag 203
func (m OrderCancelReplaceRequest) GetCoveredOrUncovered() (v enum.CoveredOrUncovered, err quickfix.MessageRejectError) {
	var f field.CoveredOrUncoveredField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptAttribute gets OptAttribute, Tag 206
func (m OrderCancelReplaceRequest) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityExchange gets SecurityExchange, Tag 207
func (m OrderCancelReplaceRequest) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaxShow gets MaxShow, Tag 210
func (m OrderCancelReplaceRequest) GetMaxShow() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MaxShowField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPegOffsetValue gets PegOffsetValue, Tag 211
func (m OrderCancelReplaceRequest) GetPegOffsetValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PegOffsetValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSpread gets Spread, Tag 218
func (m OrderCancelReplaceRequest) GetSpread() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SpreadField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkCurveCurrency gets BenchmarkCurveCurrency, Tag 220
func (m OrderCancelReplaceRequest) GetBenchmarkCurveCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkCurveCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkCurveName gets BenchmarkCurveName, Tag 221
func (m OrderCancelReplaceRequest) GetBenchmarkCurveName() (v enum.BenchmarkCurveName, err quickfix.MessageRejectError) {
	var f field.BenchmarkCurveNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkCurvePoint gets BenchmarkCurvePoint, Tag 222
func (m OrderCancelReplaceRequest) GetBenchmarkCurvePoint() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkCurvePointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponRate gets CouponRate, Tag 223
func (m OrderCancelReplaceRequest) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m OrderCancelReplaceRequest) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssueDate gets IssueDate, Tag 225
func (m OrderCancelReplaceRequest) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m OrderCancelReplaceRequest) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m OrderCancelReplaceRequest) GetRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFactor gets Factor, Tag 228
func (m OrderCancelReplaceRequest) GetFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradeOriginationDate gets TradeOriginationDate, Tag 229
func (m OrderCancelReplaceRequest) GetTradeOriginationDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeOriginationDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractMultiplier gets ContractMultiplier, Tag 231
func (m OrderCancelReplaceRequest) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetYieldType gets YieldType, Tag 235
func (m OrderCancelReplaceRequest) GetYieldType() (v enum.YieldType, err quickfix.MessageRejectError) {
	var f field.YieldTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetYield gets Yield, Tag 236
func (m OrderCancelReplaceRequest) GetYield() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.YieldField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m OrderCancelReplaceRequest) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRedemptionDate gets RedemptionDate, Tag 240
func (m OrderCancelReplaceRequest) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCreditRating gets CreditRating, Tag 255
func (m OrderCancelReplaceRequest) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m OrderCancelReplaceRequest) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m OrderCancelReplaceRequest) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m OrderCancelReplaceRequest) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m OrderCancelReplaceRequest) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m OrderCancelReplaceRequest) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355
func (m OrderCancelReplaceRequest) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplianceID gets ComplianceID, Tag 376
func (m OrderCancelReplaceRequest) GetComplianceID() (v string, err quickfix.MessageRejectError) {
	var f field.ComplianceIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSolicitedFlag gets SolicitedFlag, Tag 377
func (m OrderCancelReplaceRequest) GetSolicitedFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.SolicitedFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoTradingSessions gets NoTradingSessions, Tag 386
func (m OrderCancelReplaceRequest) GetNoTradingSessions() (NoTradingSessionsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTradingSessionsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetDiscretionInst gets DiscretionInst, Tag 388
func (m OrderCancelReplaceRequest) GetDiscretionInst() (v enum.DiscretionInst, err quickfix.MessageRejectError) {
	var f field.DiscretionInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDiscretionOffsetValue gets DiscretionOffsetValue, Tag 389
func (m OrderCancelReplaceRequest) GetDiscretionOffsetValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DiscretionOffsetValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceType gets PriceType, Tag 423
func (m OrderCancelReplaceRequest) GetPriceType() (v enum.PriceType, err quickfix.MessageRejectError) {
	var f field.PriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetGTBookingInst gets GTBookingInst, Tag 427
func (m OrderCancelReplaceRequest) GetGTBookingInst() (v enum.GTBookingInst, err quickfix.MessageRejectError) {
	var f field.GTBookingInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExpireDate gets ExpireDate, Tag 432
func (m OrderCancelReplaceRequest) GetExpireDate() (v string, err quickfix.MessageRejectError) {
	var f field.ExpireDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m OrderCancelReplaceRequest) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m OrderCancelReplaceRequest) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetProduct gets Product, Tag 460
func (m OrderCancelReplaceRequest) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCFICode gets CFICode, Tag 461
func (m OrderCancelReplaceRequest) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRoundingDirection gets RoundingDirection, Tag 468
func (m OrderCancelReplaceRequest) GetRoundingDirection() (v enum.RoundingDirection, err quickfix.MessageRejectError) {
	var f field.RoundingDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRoundingModulus gets RoundingModulus, Tag 469
func (m OrderCancelReplaceRequest) GetRoundingModulus() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RoundingModulusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m OrderCancelReplaceRequest) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m OrderCancelReplaceRequest) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m OrderCancelReplaceRequest) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCommCurrency gets CommCurrency, Tag 479
func (m OrderCancelReplaceRequest) GetCommCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CommCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCancellationRights gets CancellationRights, Tag 480
func (m OrderCancelReplaceRequest) GetCancellationRights() (v enum.CancellationRights, err quickfix.MessageRejectError) {
	var f field.CancellationRightsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMoneyLaunderingStatus gets MoneyLaunderingStatus, Tag 481
func (m OrderCancelReplaceRequest) GetMoneyLaunderingStatus() (v enum.MoneyLaunderingStatus, err quickfix.MessageRejectError) {
	var f field.MoneyLaunderingStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDesignation gets Designation, Tag 494
func (m OrderCancelReplaceRequest) GetDesignation() (v string, err quickfix.MessageRejectError) {
	var f field.DesignationField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFundRenewWaiv gets FundRenewWaiv, Tag 497
func (m OrderCancelReplaceRequest) GetFundRenewWaiv() (v enum.FundRenewWaiv, err quickfix.MessageRejectError) {
	var f field.FundRenewWaivField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRegistID gets RegistID, Tag 513
func (m OrderCancelReplaceRequest) GetRegistID() (v string, err quickfix.MessageRejectError) {
	var f field.RegistIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderPercent gets OrderPercent, Tag 516
func (m OrderCancelReplaceRequest) GetOrderPercent() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderPercentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526
func (m OrderCancelReplaceRequest) GetSecondaryClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderCapacity gets OrderCapacity, Tag 528
func (m OrderCancelReplaceRequest) GetOrderCapacity() (v enum.OrderCapacity, err quickfix.MessageRejectError) {
	var f field.OrderCapacityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderRestrictions gets OrderRestrictions, Tag 529
func (m OrderCancelReplaceRequest) GetOrderRestrictions() (v enum.OrderRestrictions, err quickfix.MessageRejectError) {
	var f field.OrderRestrictionsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityDate gets MaturityDate, Tag 541
func (m OrderCancelReplaceRequest) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrRegistry gets InstrRegistry, Tag 543
func (m OrderCancelReplaceRequest) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCashMargin gets CashMargin, Tag 544
func (m OrderCancelReplaceRequest) GetCashMargin() (v enum.CashMargin, err quickfix.MessageRejectError) {
	var f field.CashMarginField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAccountType gets AccountType, Tag 581
func (m OrderCancelReplaceRequest) GetAccountType() (v enum.AccountType, err quickfix.MessageRejectError) {
	var f field.AccountTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCustOrderCapacity gets CustOrderCapacity, Tag 582
func (m OrderCancelReplaceRequest) GetCustOrderCapacity() (v enum.CustOrderCapacity, err quickfix.MessageRejectError) {
	var f field.CustOrderCapacityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetClOrdLinkID gets ClOrdLinkID, Tag 583
func (m OrderCancelReplaceRequest) GetClOrdLinkID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdLinkIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrigOrdModTime gets OrigOrdModTime, Tag 586
func (m OrderCancelReplaceRequest) GetOrigOrdModTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.OrigOrdModTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDayBookingInst gets DayBookingInst, Tag 589
func (m OrderCancelReplaceRequest) GetDayBookingInst() (v enum.DayBookingInst, err quickfix.MessageRejectError) {
	var f field.DayBookingInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBookingUnit gets BookingUnit, Tag 590
func (m OrderCancelReplaceRequest) GetBookingUnit() (v enum.BookingUnit, err quickfix.MessageRejectError) {
	var f field.BookingUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPreallocMethod gets PreallocMethod, Tag 591
func (m OrderCancelReplaceRequest) GetPreallocMethod() (v enum.PreallocMethod, err quickfix.MessageRejectError) {
	var f field.PreallocMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetClearingFeeIndicator gets ClearingFeeIndicator, Tag 635
func (m OrderCancelReplaceRequest) GetClearingFeeIndicator() (v enum.ClearingFeeIndicator, err quickfix.MessageRejectError) {
	var f field.ClearingFeeIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPrice2 gets Price2, Tag 640
func (m OrderCancelReplaceRequest) GetPrice2() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.Price2Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAcctIDSource gets AcctIDSource, Tag 660
func (m OrderCancelReplaceRequest) GetAcctIDSource() (v enum.AcctIDSource, err quickfix.MessageRejectError) {
	var f field.AcctIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkPrice gets BenchmarkPrice, Tag 662
func (m OrderCancelReplaceRequest) GetBenchmarkPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BenchmarkPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkPriceType gets BenchmarkPriceType, Tag 663
func (m OrderCancelReplaceRequest) GetBenchmarkPriceType() (v int, err quickfix.MessageRejectError) {
	var f field.BenchmarkPriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m OrderCancelReplaceRequest) GetContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.ContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPool gets Pool, Tag 691
func (m OrderCancelReplaceRequest) GetPool() (v string, err quickfix.MessageRejectError) {
	var f field.PoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetYieldRedemptionDate gets YieldRedemptionDate, Tag 696
func (m OrderCancelReplaceRequest) GetYieldRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.YieldRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetYieldRedemptionPrice gets YieldRedemptionPrice, Tag 697
func (m OrderCancelReplaceRequest) GetYieldRedemptionPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.YieldRedemptionPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetYieldRedemptionPriceType gets YieldRedemptionPriceType, Tag 698
func (m OrderCancelReplaceRequest) GetYieldRedemptionPriceType() (v int, err quickfix.MessageRejectError) {
	var f field.YieldRedemptionPriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkSecurityID gets BenchmarkSecurityID, Tag 699
func (m OrderCancelReplaceRequest) GetBenchmarkSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetYieldCalcDate gets YieldCalcDate, Tag 701
func (m OrderCancelReplaceRequest) GetYieldCalcDate() (v string, err quickfix.MessageRejectError) {
	var f field.YieldCalcDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUnderlyings gets NoUnderlyings, Tag 711
func (m OrderCancelReplaceRequest) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetBenchmarkSecurityIDSource gets BenchmarkSecurityIDSource, Tag 761
func (m OrderCancelReplaceRequest) GetBenchmarkSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecuritySubType gets SecuritySubType, Tag 762
func (m OrderCancelReplaceRequest) GetSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBookingType gets BookingType, Tag 775
func (m OrderCancelReplaceRequest) GetBookingType() (v enum.BookingType, err quickfix.MessageRejectError) {
	var f field.BookingTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTerminationType gets TerminationType, Tag 788
func (m OrderCancelReplaceRequest) GetTerminationType() (v enum.TerminationType, err quickfix.MessageRejectError) {
	var f field.TerminationTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPegMoveType gets PegMoveType, Tag 835
func (m OrderCancelReplaceRequest) GetPegMoveType() (v enum.PegMoveType, err quickfix.MessageRejectError) {
	var f field.PegMoveTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPegOffsetType gets PegOffsetType, Tag 836
func (m OrderCancelReplaceRequest) GetPegOffsetType() (v enum.PegOffsetType, err quickfix.MessageRejectError) {
	var f field.PegOffsetTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPegLimitType gets PegLimitType, Tag 837
func (m OrderCancelReplaceRequest) GetPegLimitType() (v enum.PegLimitType, err quickfix.MessageRejectError) {
	var f field.PegLimitTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPegRoundDirection gets PegRoundDirection, Tag 838
func (m OrderCancelReplaceRequest) GetPegRoundDirection() (v enum.PegRoundDirection, err quickfix.MessageRejectError) {
	var f field.PegRoundDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPegScope gets PegScope, Tag 840
func (m OrderCancelReplaceRequest) GetPegScope() (v enum.PegScope, err quickfix.MessageRejectError) {
	var f field.PegScopeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDiscretionMoveType gets DiscretionMoveType, Tag 841
func (m OrderCancelReplaceRequest) GetDiscretionMoveType() (v enum.DiscretionMoveType, err quickfix.MessageRejectError) {
	var f field.DiscretionMoveTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDiscretionOffsetType gets DiscretionOffsetType, Tag 842
func (m OrderCancelReplaceRequest) GetDiscretionOffsetType() (v enum.DiscretionOffsetType, err quickfix.MessageRejectError) {
	var f field.DiscretionOffsetTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDiscretionLimitType gets DiscretionLimitType, Tag 843
func (m OrderCancelReplaceRequest) GetDiscretionLimitType() (v enum.DiscretionLimitType, err quickfix.MessageRejectError) {
	var f field.DiscretionLimitTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDiscretionRoundDirection gets DiscretionRoundDirection, Tag 844
func (m OrderCancelReplaceRequest) GetDiscretionRoundDirection() (v enum.DiscretionRoundDirection, err quickfix.MessageRejectError) {
	var f field.DiscretionRoundDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDiscretionScope gets DiscretionScope, Tag 846
func (m OrderCancelReplaceRequest) GetDiscretionScope() (v enum.DiscretionScope, err quickfix.MessageRejectError) {
	var f field.DiscretionScopeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTargetStrategy gets TargetStrategy, Tag 847
func (m OrderCancelReplaceRequest) GetTargetStrategy() (v enum.TargetStrategy, err quickfix.MessageRejectError) {
	var f field.TargetStrategyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTargetStrategyParameters gets TargetStrategyParameters, Tag 848
func (m OrderCancelReplaceRequest) GetTargetStrategyParameters() (v string, err quickfix.MessageRejectError) {
	var f field.TargetStrategyParametersField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetParticipationRate gets ParticipationRate, Tag 849
func (m OrderCancelReplaceRequest) GetParticipationRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ParticipationRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQtyType gets QtyType, Tag 854
func (m OrderCancelReplaceRequest) GetQtyType() (v enum.QtyType, err quickfix.MessageRejectError) {
	var f field.QtyTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoEvents gets NoEvents, Tag 864
func (m OrderCancelReplaceRequest) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetDatedDate gets DatedDate, Tag 873
func (m OrderCancelReplaceRequest) GetDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.DatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m OrderCancelReplaceRequest) GetInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.InterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPProgram gets CPProgram, Tag 875
func (m OrderCancelReplaceRequest) GetCPProgram() (v enum.CPProgram, err quickfix.MessageRejectError) {
	var f field.CPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPRegType gets CPRegType, Tag 876
func (m OrderCancelReplaceRequest) GetCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.CPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarginRatio gets MarginRatio, Tag 898
func (m OrderCancelReplaceRequest) GetMarginRatio() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MarginRatioField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAgreementDesc gets AgreementDesc, Tag 913
func (m OrderCancelReplaceRequest) GetAgreementDesc() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAgreementID gets AgreementID, Tag 914
func (m OrderCancelReplaceRequest) GetAgreementID() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAgreementDate gets AgreementDate, Tag 915
func (m OrderCancelReplaceRequest) GetAgreementDate() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStartDate gets StartDate, Tag 916
func (m OrderCancelReplaceRequest) GetStartDate() (v string, err quickfix.MessageRejectError) {
	var f field.StartDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEndDate gets EndDate, Tag 917
func (m OrderCancelReplaceRequest) GetEndDate() (v string, err quickfix.MessageRejectError) {
	var f field.EndDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAgreementCurrency gets AgreementCurrency, Tag 918
func (m OrderCancelReplaceRequest) GetAgreementCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDeliveryType gets DeliveryType, Tag 919
func (m OrderCancelReplaceRequest) GetDeliveryType() (v enum.DeliveryType, err quickfix.MessageRejectError) {
	var f field.DeliveryTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m OrderCancelReplaceRequest) GetStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.StrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasAccount returns true if Account is present, Tag 1
func (m OrderCancelReplaceRequest) HasAccount() bool {
	return m.Has(tag.Account)
}

// HasClOrdID returns true if ClOrdID is present, Tag 11
func (m OrderCancelReplaceRequest) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

// HasCommission returns true if Commission is present, Tag 12
func (m OrderCancelReplaceRequest) HasCommission() bool {
	return m.Has(tag.Commission)
}

// HasCommType returns true if CommType is present, Tag 13
func (m OrderCancelReplaceRequest) HasCommType() bool {
	return m.Has(tag.CommType)
}

// HasCurrency returns true if Currency is present, Tag 15
func (m OrderCancelReplaceRequest) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// HasExecInst returns true if ExecInst is present, Tag 18
func (m OrderCancelReplaceRequest) HasExecInst() bool {
	return m.Has(tag.ExecInst)
}

// HasHandlInst returns true if HandlInst is present, Tag 21
func (m OrderCancelReplaceRequest) HasHandlInst() bool {
	return m.Has(tag.HandlInst)
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m OrderCancelReplaceRequest) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasOrderID returns true if OrderID is present, Tag 37
func (m OrderCancelReplaceRequest) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

// HasOrderQty returns true if OrderQty is present, Tag 38
func (m OrderCancelReplaceRequest) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

// HasOrdType returns true if OrdType is present, Tag 40
func (m OrderCancelReplaceRequest) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

// HasOrigClOrdID returns true if OrigClOrdID is present, Tag 41
func (m OrderCancelReplaceRequest) HasOrigClOrdID() bool {
	return m.Has(tag.OrigClOrdID)
}

// HasPrice returns true if Price is present, Tag 44
func (m OrderCancelReplaceRequest) HasPrice() bool {
	return m.Has(tag.Price)
}

// HasSecurityID returns true if SecurityID is present, Tag 48
func (m OrderCancelReplaceRequest) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSide returns true if Side is present, Tag 54
func (m OrderCancelReplaceRequest) HasSide() bool {
	return m.Has(tag.Side)
}

// HasSymbol returns true if Symbol is present, Tag 55
func (m OrderCancelReplaceRequest) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasText returns true if Text is present, Tag 58
func (m OrderCancelReplaceRequest) HasText() bool {
	return m.Has(tag.Text)
}

// HasTimeInForce returns true if TimeInForce is present, Tag 59
func (m OrderCancelReplaceRequest) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

// HasTransactTime returns true if TransactTime is present, Tag 60
func (m OrderCancelReplaceRequest) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasSettlType returns true if SettlType is present, Tag 63
func (m OrderCancelReplaceRequest) HasSettlType() bool {
	return m.Has(tag.SettlType)
}

// HasSettlDate returns true if SettlDate is present, Tag 64
func (m OrderCancelReplaceRequest) HasSettlDate() bool {
	return m.Has(tag.SettlDate)
}

// HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m OrderCancelReplaceRequest) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

// HasListID returns true if ListID is present, Tag 66
func (m OrderCancelReplaceRequest) HasListID() bool {
	return m.Has(tag.ListID)
}

// HasAllocID returns true if AllocID is present, Tag 70
func (m OrderCancelReplaceRequest) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

// HasTradeDate returns true if TradeDate is present, Tag 75
func (m OrderCancelReplaceRequest) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

// HasPositionEffect returns true if PositionEffect is present, Tag 77
func (m OrderCancelReplaceRequest) HasPositionEffect() bool {
	return m.Has(tag.PositionEffect)
}

// HasNoAllocs returns true if NoAllocs is present, Tag 78
func (m OrderCancelReplaceRequest) HasNoAllocs() bool {
	return m.Has(tag.NoAllocs)
}

// HasStopPx returns true if StopPx is present, Tag 99
func (m OrderCancelReplaceRequest) HasStopPx() bool {
	return m.Has(tag.StopPx)
}

// HasExDestination returns true if ExDestination is present, Tag 100
func (m OrderCancelReplaceRequest) HasExDestination() bool {
	return m.Has(tag.ExDestination)
}

// HasIssuer returns true if Issuer is present, Tag 106
func (m OrderCancelReplaceRequest) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

// HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m OrderCancelReplaceRequest) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

// HasMinQty returns true if MinQty is present, Tag 110
func (m OrderCancelReplaceRequest) HasMinQty() bool {
	return m.Has(tag.MinQty)
}

// HasMaxFloor returns true if MaxFloor is present, Tag 111
func (m OrderCancelReplaceRequest) HasMaxFloor() bool {
	return m.Has(tag.MaxFloor)
}

// HasLocateReqd returns true if LocateReqd is present, Tag 114
func (m OrderCancelReplaceRequest) HasLocateReqd() bool {
	return m.Has(tag.LocateReqd)
}

// HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m OrderCancelReplaceRequest) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

// HasForexReq returns true if ForexReq is present, Tag 121
func (m OrderCancelReplaceRequest) HasForexReq() bool {
	return m.Has(tag.ForexReq)
}

// HasExpireTime returns true if ExpireTime is present, Tag 126
func (m OrderCancelReplaceRequest) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

// HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m OrderCancelReplaceRequest) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

// HasSecurityType returns true if SecurityType is present, Tag 167
func (m OrderCancelReplaceRequest) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

// HasEffectiveTime returns true if EffectiveTime is present, Tag 168
func (m OrderCancelReplaceRequest) HasEffectiveTime() bool {
	return m.Has(tag.EffectiveTime)
}

// HasOrderQty2 returns true if OrderQty2 is present, Tag 192
func (m OrderCancelReplaceRequest) HasOrderQty2() bool {
	return m.Has(tag.OrderQty2)
}

// HasSettlDate2 returns true if SettlDate2 is present, Tag 193
func (m OrderCancelReplaceRequest) HasSettlDate2() bool {
	return m.Has(tag.SettlDate2)
}

// HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m OrderCancelReplaceRequest) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

// HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m OrderCancelReplaceRequest) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

// HasCoveredOrUncovered returns true if CoveredOrUncovered is present, Tag 203
func (m OrderCancelReplaceRequest) HasCoveredOrUncovered() bool {
	return m.Has(tag.CoveredOrUncovered)
}

// HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m OrderCancelReplaceRequest) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

// HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m OrderCancelReplaceRequest) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

// HasMaxShow returns true if MaxShow is present, Tag 210
func (m OrderCancelReplaceRequest) HasMaxShow() bool {
	return m.Has(tag.MaxShow)
}

// HasPegOffsetValue returns true if PegOffsetValue is present, Tag 211
func (m OrderCancelReplaceRequest) HasPegOffsetValue() bool {
	return m.Has(tag.PegOffsetValue)
}

// HasSpread returns true if Spread is present, Tag 218
func (m OrderCancelReplaceRequest) HasSpread() bool {
	return m.Has(tag.Spread)
}

// HasBenchmarkCurveCurrency returns true if BenchmarkCurveCurrency is present, Tag 220
func (m OrderCancelReplaceRequest) HasBenchmarkCurveCurrency() bool {
	return m.Has(tag.BenchmarkCurveCurrency)
}

// HasBenchmarkCurveName returns true if BenchmarkCurveName is present, Tag 221
func (m OrderCancelReplaceRequest) HasBenchmarkCurveName() bool {
	return m.Has(tag.BenchmarkCurveName)
}

// HasBenchmarkCurvePoint returns true if BenchmarkCurvePoint is present, Tag 222
func (m OrderCancelReplaceRequest) HasBenchmarkCurvePoint() bool {
	return m.Has(tag.BenchmarkCurvePoint)
}

// HasCouponRate returns true if CouponRate is present, Tag 223
func (m OrderCancelReplaceRequest) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

// HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m OrderCancelReplaceRequest) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

// HasIssueDate returns true if IssueDate is present, Tag 225
func (m OrderCancelReplaceRequest) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

// HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m OrderCancelReplaceRequest) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

// HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m OrderCancelReplaceRequest) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

// HasFactor returns true if Factor is present, Tag 228
func (m OrderCancelReplaceRequest) HasFactor() bool {
	return m.Has(tag.Factor)
}

// HasTradeOriginationDate returns true if TradeOriginationDate is present, Tag 229
func (m OrderCancelReplaceRequest) HasTradeOriginationDate() bool {
	return m.Has(tag.TradeOriginationDate)
}

// HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m OrderCancelReplaceRequest) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

// HasYieldType returns true if YieldType is present, Tag 235
func (m OrderCancelReplaceRequest) HasYieldType() bool {
	return m.Has(tag.YieldType)
}

// HasYield returns true if Yield is present, Tag 236
func (m OrderCancelReplaceRequest) HasYield() bool {
	return m.Has(tag.Yield)
}

// HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m OrderCancelReplaceRequest) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

// HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m OrderCancelReplaceRequest) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

// HasCreditRating returns true if CreditRating is present, Tag 255
func (m OrderCancelReplaceRequest) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

// HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m OrderCancelReplaceRequest) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

// HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m OrderCancelReplaceRequest) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

// HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m OrderCancelReplaceRequest) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

// HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m OrderCancelReplaceRequest) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m OrderCancelReplaceRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355
func (m OrderCancelReplaceRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasComplianceID returns true if ComplianceID is present, Tag 376
func (m OrderCancelReplaceRequest) HasComplianceID() bool {
	return m.Has(tag.ComplianceID)
}

// HasSolicitedFlag returns true if SolicitedFlag is present, Tag 377
func (m OrderCancelReplaceRequest) HasSolicitedFlag() bool {
	return m.Has(tag.SolicitedFlag)
}

// HasNoTradingSessions returns true if NoTradingSessions is present, Tag 386
func (m OrderCancelReplaceRequest) HasNoTradingSessions() bool {
	return m.Has(tag.NoTradingSessions)
}

// HasDiscretionInst returns true if DiscretionInst is present, Tag 388
func (m OrderCancelReplaceRequest) HasDiscretionInst() bool {
	return m.Has(tag.DiscretionInst)
}

// HasDiscretionOffsetValue returns true if DiscretionOffsetValue is present, Tag 389
func (m OrderCancelReplaceRequest) HasDiscretionOffsetValue() bool {
	return m.Has(tag.DiscretionOffsetValue)
}

// HasPriceType returns true if PriceType is present, Tag 423
func (m OrderCancelReplaceRequest) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

// HasGTBookingInst returns true if GTBookingInst is present, Tag 427
func (m OrderCancelReplaceRequest) HasGTBookingInst() bool {
	return m.Has(tag.GTBookingInst)
}

// HasExpireDate returns true if ExpireDate is present, Tag 432
func (m OrderCancelReplaceRequest) HasExpireDate() bool {
	return m.Has(tag.ExpireDate)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m OrderCancelReplaceRequest) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m OrderCancelReplaceRequest) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

// HasProduct returns true if Product is present, Tag 460
func (m OrderCancelReplaceRequest) HasProduct() bool {
	return m.Has(tag.Product)
}

// HasCFICode returns true if CFICode is present, Tag 461
func (m OrderCancelReplaceRequest) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

// HasRoundingDirection returns true if RoundingDirection is present, Tag 468
func (m OrderCancelReplaceRequest) HasRoundingDirection() bool {
	return m.Has(tag.RoundingDirection)
}

// HasRoundingModulus returns true if RoundingModulus is present, Tag 469
func (m OrderCancelReplaceRequest) HasRoundingModulus() bool {
	return m.Has(tag.RoundingModulus)
}

// HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m OrderCancelReplaceRequest) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

// HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m OrderCancelReplaceRequest) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

// HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m OrderCancelReplaceRequest) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

// HasCommCurrency returns true if CommCurrency is present, Tag 479
func (m OrderCancelReplaceRequest) HasCommCurrency() bool {
	return m.Has(tag.CommCurrency)
}

// HasCancellationRights returns true if CancellationRights is present, Tag 480
func (m OrderCancelReplaceRequest) HasCancellationRights() bool {
	return m.Has(tag.CancellationRights)
}

// HasMoneyLaunderingStatus returns true if MoneyLaunderingStatus is present, Tag 481
func (m OrderCancelReplaceRequest) HasMoneyLaunderingStatus() bool {
	return m.Has(tag.MoneyLaunderingStatus)
}

// HasDesignation returns true if Designation is present, Tag 494
func (m OrderCancelReplaceRequest) HasDesignation() bool {
	return m.Has(tag.Designation)
}

// HasFundRenewWaiv returns true if FundRenewWaiv is present, Tag 497
func (m OrderCancelReplaceRequest) HasFundRenewWaiv() bool {
	return m.Has(tag.FundRenewWaiv)
}

// HasRegistID returns true if RegistID is present, Tag 513
func (m OrderCancelReplaceRequest) HasRegistID() bool {
	return m.Has(tag.RegistID)
}

// HasOrderPercent returns true if OrderPercent is present, Tag 516
func (m OrderCancelReplaceRequest) HasOrderPercent() bool {
	return m.Has(tag.OrderPercent)
}

// HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526
func (m OrderCancelReplaceRequest) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

// HasOrderCapacity returns true if OrderCapacity is present, Tag 528
func (m OrderCancelReplaceRequest) HasOrderCapacity() bool {
	return m.Has(tag.OrderCapacity)
}

// HasOrderRestrictions returns true if OrderRestrictions is present, Tag 529
func (m OrderCancelReplaceRequest) HasOrderRestrictions() bool {
	return m.Has(tag.OrderRestrictions)
}

// HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m OrderCancelReplaceRequest) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

// HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m OrderCancelReplaceRequest) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

// HasCashMargin returns true if CashMargin is present, Tag 544
func (m OrderCancelReplaceRequest) HasCashMargin() bool {
	return m.Has(tag.CashMargin)
}

// HasAccountType returns true if AccountType is present, Tag 581
func (m OrderCancelReplaceRequest) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

// HasCustOrderCapacity returns true if CustOrderCapacity is present, Tag 582
func (m OrderCancelReplaceRequest) HasCustOrderCapacity() bool {
	return m.Has(tag.CustOrderCapacity)
}

// HasClOrdLinkID returns true if ClOrdLinkID is present, Tag 583
func (m OrderCancelReplaceRequest) HasClOrdLinkID() bool {
	return m.Has(tag.ClOrdLinkID)
}

// HasOrigOrdModTime returns true if OrigOrdModTime is present, Tag 586
func (m OrderCancelReplaceRequest) HasOrigOrdModTime() bool {
	return m.Has(tag.OrigOrdModTime)
}

// HasDayBookingInst returns true if DayBookingInst is present, Tag 589
func (m OrderCancelReplaceRequest) HasDayBookingInst() bool {
	return m.Has(tag.DayBookingInst)
}

// HasBookingUnit returns true if BookingUnit is present, Tag 590
func (m OrderCancelReplaceRequest) HasBookingUnit() bool {
	return m.Has(tag.BookingUnit)
}

// HasPreallocMethod returns true if PreallocMethod is present, Tag 591
func (m OrderCancelReplaceRequest) HasPreallocMethod() bool {
	return m.Has(tag.PreallocMethod)
}

// HasClearingFeeIndicator returns true if ClearingFeeIndicator is present, Tag 635
func (m OrderCancelReplaceRequest) HasClearingFeeIndicator() bool {
	return m.Has(tag.ClearingFeeIndicator)
}

// HasPrice2 returns true if Price2 is present, Tag 640
func (m OrderCancelReplaceRequest) HasPrice2() bool {
	return m.Has(tag.Price2)
}

// HasAcctIDSource returns true if AcctIDSource is present, Tag 660
func (m OrderCancelReplaceRequest) HasAcctIDSource() bool {
	return m.Has(tag.AcctIDSource)
}

// HasBenchmarkPrice returns true if BenchmarkPrice is present, Tag 662
func (m OrderCancelReplaceRequest) HasBenchmarkPrice() bool {
	return m.Has(tag.BenchmarkPrice)
}

// HasBenchmarkPriceType returns true if BenchmarkPriceType is present, Tag 663
func (m OrderCancelReplaceRequest) HasBenchmarkPriceType() bool {
	return m.Has(tag.BenchmarkPriceType)
}

// HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m OrderCancelReplaceRequest) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

// HasPool returns true if Pool is present, Tag 691
func (m OrderCancelReplaceRequest) HasPool() bool {
	return m.Has(tag.Pool)
}

// HasYieldRedemptionDate returns true if YieldRedemptionDate is present, Tag 696
func (m OrderCancelReplaceRequest) HasYieldRedemptionDate() bool {
	return m.Has(tag.YieldRedemptionDate)
}

// HasYieldRedemptionPrice returns true if YieldRedemptionPrice is present, Tag 697
func (m OrderCancelReplaceRequest) HasYieldRedemptionPrice() bool {
	return m.Has(tag.YieldRedemptionPrice)
}

// HasYieldRedemptionPriceType returns true if YieldRedemptionPriceType is present, Tag 698
func (m OrderCancelReplaceRequest) HasYieldRedemptionPriceType() bool {
	return m.Has(tag.YieldRedemptionPriceType)
}

// HasBenchmarkSecurityID returns true if BenchmarkSecurityID is present, Tag 699
func (m OrderCancelReplaceRequest) HasBenchmarkSecurityID() bool {
	return m.Has(tag.BenchmarkSecurityID)
}

// HasYieldCalcDate returns true if YieldCalcDate is present, Tag 701
func (m OrderCancelReplaceRequest) HasYieldCalcDate() bool {
	return m.Has(tag.YieldCalcDate)
}

// HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711
func (m OrderCancelReplaceRequest) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

// HasBenchmarkSecurityIDSource returns true if BenchmarkSecurityIDSource is present, Tag 761
func (m OrderCancelReplaceRequest) HasBenchmarkSecurityIDSource() bool {
	return m.Has(tag.BenchmarkSecurityIDSource)
}

// HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m OrderCancelReplaceRequest) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

// HasBookingType returns true if BookingType is present, Tag 775
func (m OrderCancelReplaceRequest) HasBookingType() bool {
	return m.Has(tag.BookingType)
}

// HasTerminationType returns true if TerminationType is present, Tag 788
func (m OrderCancelReplaceRequest) HasTerminationType() bool {
	return m.Has(tag.TerminationType)
}

// HasPegMoveType returns true if PegMoveType is present, Tag 835
func (m OrderCancelReplaceRequest) HasPegMoveType() bool {
	return m.Has(tag.PegMoveType)
}

// HasPegOffsetType returns true if PegOffsetType is present, Tag 836
func (m OrderCancelReplaceRequest) HasPegOffsetType() bool {
	return m.Has(tag.PegOffsetType)
}

// HasPegLimitType returns true if PegLimitType is present, Tag 837
func (m OrderCancelReplaceRequest) HasPegLimitType() bool {
	return m.Has(tag.PegLimitType)
}

// HasPegRoundDirection returns true if PegRoundDirection is present, Tag 838
func (m OrderCancelReplaceRequest) HasPegRoundDirection() bool {
	return m.Has(tag.PegRoundDirection)
}

// HasPegScope returns true if PegScope is present, Tag 840
func (m OrderCancelReplaceRequest) HasPegScope() bool {
	return m.Has(tag.PegScope)
}

// HasDiscretionMoveType returns true if DiscretionMoveType is present, Tag 841
func (m OrderCancelReplaceRequest) HasDiscretionMoveType() bool {
	return m.Has(tag.DiscretionMoveType)
}

// HasDiscretionOffsetType returns true if DiscretionOffsetType is present, Tag 842
func (m OrderCancelReplaceRequest) HasDiscretionOffsetType() bool {
	return m.Has(tag.DiscretionOffsetType)
}

// HasDiscretionLimitType returns true if DiscretionLimitType is present, Tag 843
func (m OrderCancelReplaceRequest) HasDiscretionLimitType() bool {
	return m.Has(tag.DiscretionLimitType)
}

// HasDiscretionRoundDirection returns true if DiscretionRoundDirection is present, Tag 844
func (m OrderCancelReplaceRequest) HasDiscretionRoundDirection() bool {
	return m.Has(tag.DiscretionRoundDirection)
}

// HasDiscretionScope returns true if DiscretionScope is present, Tag 846
func (m OrderCancelReplaceRequest) HasDiscretionScope() bool {
	return m.Has(tag.DiscretionScope)
}

// HasTargetStrategy returns true if TargetStrategy is present, Tag 847
func (m OrderCancelReplaceRequest) HasTargetStrategy() bool {
	return m.Has(tag.TargetStrategy)
}

// HasTargetStrategyParameters returns true if TargetStrategyParameters is present, Tag 848
func (m OrderCancelReplaceRequest) HasTargetStrategyParameters() bool {
	return m.Has(tag.TargetStrategyParameters)
}

// HasParticipationRate returns true if ParticipationRate is present, Tag 849
func (m OrderCancelReplaceRequest) HasParticipationRate() bool {
	return m.Has(tag.ParticipationRate)
}

// HasQtyType returns true if QtyType is present, Tag 854
func (m OrderCancelReplaceRequest) HasQtyType() bool {
	return m.Has(tag.QtyType)
}

// HasNoEvents returns true if NoEvents is present, Tag 864
func (m OrderCancelReplaceRequest) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

// HasDatedDate returns true if DatedDate is present, Tag 873
func (m OrderCancelReplaceRequest) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

// HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m OrderCancelReplaceRequest) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

// HasCPProgram returns true if CPProgram is present, Tag 875
func (m OrderCancelReplaceRequest) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

// HasCPRegType returns true if CPRegType is present, Tag 876
func (m OrderCancelReplaceRequest) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

// HasMarginRatio returns true if MarginRatio is present, Tag 898
func (m OrderCancelReplaceRequest) HasMarginRatio() bool {
	return m.Has(tag.MarginRatio)
}

// HasAgreementDesc returns true if AgreementDesc is present, Tag 913
func (m OrderCancelReplaceRequest) HasAgreementDesc() bool {
	return m.Has(tag.AgreementDesc)
}

// HasAgreementID returns true if AgreementID is present, Tag 914
func (m OrderCancelReplaceRequest) HasAgreementID() bool {
	return m.Has(tag.AgreementID)
}

// HasAgreementDate returns true if AgreementDate is present, Tag 915
func (m OrderCancelReplaceRequest) HasAgreementDate() bool {
	return m.Has(tag.AgreementDate)
}

// HasStartDate returns true if StartDate is present, Tag 916
func (m OrderCancelReplaceRequest) HasStartDate() bool {
	return m.Has(tag.StartDate)
}

// HasEndDate returns true if EndDate is present, Tag 917
func (m OrderCancelReplaceRequest) HasEndDate() bool {
	return m.Has(tag.EndDate)
}

// HasAgreementCurrency returns true if AgreementCurrency is present, Tag 918
func (m OrderCancelReplaceRequest) HasAgreementCurrency() bool {
	return m.Has(tag.AgreementCurrency)
}

// HasDeliveryType returns true if DeliveryType is present, Tag 919
func (m OrderCancelReplaceRequest) HasDeliveryType() bool {
	return m.Has(tag.DeliveryType)
}

// HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m OrderCancelReplaceRequest) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

// NoAllocs is a repeating group element, Tag 78
type NoAllocs struct {
	*quickfix.Group
}

// SetAllocAccount sets AllocAccount, Tag 79
func (m NoAllocs) SetAllocAccount(v string) {
	m.Set(field.NewAllocAccount(v))
}

// SetAllocAcctIDSource sets AllocAcctIDSource, Tag 661
func (m NoAllocs) SetAllocAcctIDSource(v int) {
	m.Set(field.NewAllocAcctIDSource(v))
}

// SetAllocSettlCurrency sets AllocSettlCurrency, Tag 736
func (m NoAllocs) SetAllocSettlCurrency(v string) {
	m.Set(field.NewAllocSettlCurrency(v))
}

// SetIndividualAllocID sets IndividualAllocID, Tag 467
func (m NoAllocs) SetIndividualAllocID(v string) {
	m.Set(field.NewIndividualAllocID(v))
}

// SetNoNestedPartyIDs sets NoNestedPartyIDs, Tag 539
func (m NoAllocs) SetNoNestedPartyIDs(f NoNestedPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetAllocQty sets AllocQty, Tag 80
func (m NoAllocs) SetAllocQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewAllocQty(value, scale))
}

// GetAllocAccount gets AllocAccount, Tag 79
func (m NoAllocs) GetAllocAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AllocAccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocAcctIDSource gets AllocAcctIDSource, Tag 661
func (m NoAllocs) GetAllocAcctIDSource() (v int, err quickfix.MessageRejectError) {
	var f field.AllocAcctIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocSettlCurrency gets AllocSettlCurrency, Tag 736
func (m NoAllocs) GetAllocSettlCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.AllocSettlCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIndividualAllocID gets IndividualAllocID, Tag 467
func (m NoAllocs) GetIndividualAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.IndividualAllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoNestedPartyIDs gets NoNestedPartyIDs, Tag 539
func (m NoAllocs) GetNoNestedPartyIDs() (NoNestedPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetAllocQty gets AllocQty, Tag 80
func (m NoAllocs) GetAllocQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AllocQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasAllocAccount returns true if AllocAccount is present, Tag 79
func (m NoAllocs) HasAllocAccount() bool {
	return m.Has(tag.AllocAccount)
}

// HasAllocAcctIDSource returns true if AllocAcctIDSource is present, Tag 661
func (m NoAllocs) HasAllocAcctIDSource() bool {
	return m.Has(tag.AllocAcctIDSource)
}

// HasAllocSettlCurrency returns true if AllocSettlCurrency is present, Tag 736
func (m NoAllocs) HasAllocSettlCurrency() bool {
	return m.Has(tag.AllocSettlCurrency)
}

// HasIndividualAllocID returns true if IndividualAllocID is present, Tag 467
func (m NoAllocs) HasIndividualAllocID() bool {
	return m.Has(tag.IndividualAllocID)
}

// HasNoNestedPartyIDs returns true if NoNestedPartyIDs is present, Tag 539
func (m NoAllocs) HasNoNestedPartyIDs() bool {
	return m.Has(tag.NoNestedPartyIDs)
}

// HasAllocQty returns true if AllocQty is present, Tag 80
func (m NoAllocs) HasAllocQty() bool {
	return m.Has(tag.AllocQty)
}

// NoNestedPartyIDs is a repeating group element, Tag 539
type NoNestedPartyIDs struct {
	*quickfix.Group
}

// SetNestedPartyID sets NestedPartyID, Tag 524
func (m NoNestedPartyIDs) SetNestedPartyID(v string) {
	m.Set(field.NewNestedPartyID(v))
}

// SetNestedPartyIDSource sets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) SetNestedPartyIDSource(v string) {
	m.Set(field.NewNestedPartyIDSource(v))
}

// SetNestedPartyRole sets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) SetNestedPartyRole(v int) {
	m.Set(field.NewNestedPartyRole(v))
}

// SetNoNestedPartySubIDs sets NoNestedPartySubIDs, Tag 804
func (m NoNestedPartyIDs) SetNoNestedPartySubIDs(f NoNestedPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetNestedPartyID gets NestedPartyID, Tag 524
func (m NoNestedPartyIDs) GetNestedPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNestedPartyIDSource gets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) GetNestedPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNestedPartyRole gets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) GetNestedPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.NestedPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoNestedPartySubIDs gets NoNestedPartySubIDs, Tag 804
func (m NoNestedPartyIDs) GetNoNestedPartySubIDs() (NoNestedPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasNestedPartyID returns true if NestedPartyID is present, Tag 524
func (m NoNestedPartyIDs) HasNestedPartyID() bool {
	return m.Has(tag.NestedPartyID)
}

// HasNestedPartyIDSource returns true if NestedPartyIDSource is present, Tag 525
func (m NoNestedPartyIDs) HasNestedPartyIDSource() bool {
	return m.Has(tag.NestedPartyIDSource)
}

// HasNestedPartyRole returns true if NestedPartyRole is present, Tag 538
func (m NoNestedPartyIDs) HasNestedPartyRole() bool {
	return m.Has(tag.NestedPartyRole)
}

// HasNoNestedPartySubIDs returns true if NoNestedPartySubIDs is present, Tag 804
func (m NoNestedPartyIDs) HasNoNestedPartySubIDs() bool {
	return m.Has(tag.NoNestedPartySubIDs)
}

// NoNestedPartySubIDs is a repeating group element, Tag 804
type NoNestedPartySubIDs struct {
	*quickfix.Group
}

// SetNestedPartySubID sets NestedPartySubID, Tag 545
func (m NoNestedPartySubIDs) SetNestedPartySubID(v string) {
	m.Set(field.NewNestedPartySubID(v))
}

// SetNestedPartySubIDType sets NestedPartySubIDType, Tag 805
func (m NoNestedPartySubIDs) SetNestedPartySubIDType(v int) {
	m.Set(field.NewNestedPartySubIDType(v))
}

// GetNestedPartySubID gets NestedPartySubID, Tag 545
func (m NoNestedPartySubIDs) GetNestedPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNestedPartySubIDType gets NestedPartySubIDType, Tag 805
func (m NoNestedPartySubIDs) GetNestedPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.NestedPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasNestedPartySubID returns true if NestedPartySubID is present, Tag 545
func (m NoNestedPartySubIDs) HasNestedPartySubID() bool {
	return m.Has(tag.NestedPartySubID)
}

// HasNestedPartySubIDType returns true if NestedPartySubIDType is present, Tag 805
func (m NoNestedPartySubIDs) HasNestedPartySubIDType() bool {
	return m.Has(tag.NestedPartySubIDType)
}

// NoNestedPartySubIDsRepeatingGroup is a repeating group, Tag 804
type NoNestedPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoNestedPartySubIDsRepeatingGroup returns an initialized, NoNestedPartySubIDsRepeatingGroup
func NewNoNestedPartySubIDsRepeatingGroup() NoNestedPartySubIDsRepeatingGroup {
	return NoNestedPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNestedPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedPartySubID), quickfix.GroupElement(tag.NestedPartySubIDType)}),
	}
}

// Add create and append a new NoNestedPartySubIDs to this group
func (m NoNestedPartySubIDsRepeatingGroup) Add() NoNestedPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoNestedPartySubIDs{g}
}

// Get returns the ith NoNestedPartySubIDs in the NoNestedPartySubIDsRepeatinGroup
func (m NoNestedPartySubIDsRepeatingGroup) Get(i int) NoNestedPartySubIDs {
	return NoNestedPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoNestedPartyIDsRepeatingGroup is a repeating group, Tag 539
type NoNestedPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoNestedPartyIDsRepeatingGroup returns an initialized, NoNestedPartyIDsRepeatingGroup
func NewNoNestedPartyIDsRepeatingGroup() NoNestedPartyIDsRepeatingGroup {
	return NoNestedPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNestedPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedPartyID), quickfix.GroupElement(tag.NestedPartyIDSource), quickfix.GroupElement(tag.NestedPartyRole), NewNoNestedPartySubIDsRepeatingGroup()}),
	}
}

// Add create and append a new NoNestedPartyIDs to this group
func (m NoNestedPartyIDsRepeatingGroup) Add() NoNestedPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoNestedPartyIDs{g}
}

// Get returns the ith NoNestedPartyIDs in the NoNestedPartyIDsRepeatinGroup
func (m NoNestedPartyIDsRepeatingGroup) Get(i int) NoNestedPartyIDs {
	return NoNestedPartyIDs{m.RepeatingGroup.Get(i)}
}

// NoAllocsRepeatingGroup is a repeating group, Tag 78
type NoAllocsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoAllocsRepeatingGroup returns an initialized, NoAllocsRepeatingGroup
func NewNoAllocsRepeatingGroup() NoAllocsRepeatingGroup {
	return NoAllocsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAllocs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AllocAccount), quickfix.GroupElement(tag.AllocAcctIDSource), quickfix.GroupElement(tag.AllocSettlCurrency), quickfix.GroupElement(tag.IndividualAllocID), NewNoNestedPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.AllocQty)}),
	}
}

// Add create and append a new NoAllocs to this group
func (m NoAllocsRepeatingGroup) Add() NoAllocs {
	g := m.RepeatingGroup.Add()
	return NoAllocs{g}
}

// Get returns the ith NoAllocs in the NoAllocsRepeatinGroup
func (m NoAllocsRepeatingGroup) Get(i int) NoAllocs {
	return NoAllocs{m.RepeatingGroup.Get(i)}
}

// NoTradingSessions is a repeating group element, Tag 386
type NoTradingSessions struct {
	*quickfix.Group
}

// SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoTradingSessions) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m NoTradingSessions) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoTradingSessions) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m NoTradingSessions) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m NoTradingSessions) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m NoTradingSessions) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

// NoTradingSessionsRepeatingGroup is a repeating group, Tag 386
type NoTradingSessionsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoTradingSessionsRepeatingGroup returns an initialized, NoTradingSessionsRepeatingGroup
func NewNoTradingSessionsRepeatingGroup() NoTradingSessionsRepeatingGroup {
	return NoTradingSessionsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTradingSessions,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.TradingSessionSubID)}),
	}
}

// Add create and append a new NoTradingSessions to this group
func (m NoTradingSessionsRepeatingGroup) Add() NoTradingSessions {
	g := m.RepeatingGroup.Add()
	return NoTradingSessions{g}
}

// Get returns the ith NoTradingSessions in the NoTradingSessionsRepeatinGroup
func (m NoTradingSessionsRepeatingGroup) Get(i int) NoTradingSessions {
	return NoTradingSessions{m.RepeatingGroup.Get(i)}
}

// NoPartyIDs is a repeating group element, Tag 453
type NoPartyIDs struct {
	*quickfix.Group
}

// SetPartyID sets PartyID, Tag 448
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

// SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

// SetPartyRole sets PartyRole, Tag 452
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

// SetNoPartySubIDs sets NoPartySubIDs, Tag 802
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetPartyID gets PartyID, Tag 448
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyRole gets PartyRole, Tag 452
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartySubIDs gets NoPartySubIDs, Tag 802
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasPartyID returns true if PartyID is present, Tag 448
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

// HasPartyIDSource returns true if PartyIDSource is present, Tag 447
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

// HasPartyRole returns true if PartyRole is present, Tag 452
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

// HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

// NoPartySubIDs is a repeating group element, Tag 802
type NoPartySubIDs struct {
	*quickfix.Group
}

// SetPartySubID sets PartySubID, Tag 523
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

// SetPartySubIDType sets PartySubIDType, Tag 803
func (m NoPartySubIDs) SetPartySubIDType(v enum.PartySubIDType) {
	m.Set(field.NewPartySubIDType(v))
}

// GetPartySubID gets PartySubID, Tag 523
func (m NoPartySubIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartySubIDType gets PartySubIDType, Tag 803
func (m NoPartySubIDs) GetPartySubIDType() (v enum.PartySubIDType, err quickfix.MessageRejectError) {
	var f field.PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartySubID returns true if PartySubID is present, Tag 523
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

// HasPartySubIDType returns true if PartySubIDType is present, Tag 803
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

// NoPartySubIDsRepeatingGroup is a repeating group, Tag 802
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartySubID), quickfix.GroupElement(tag.PartySubIDType)}),
	}
}

// Add create and append a new NoPartySubIDs to this group
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

// Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoPartyIDsRepeatingGroup is a repeating group, Tag 453
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), NewNoPartySubIDsRepeatingGroup()}),
	}
}

// Add create and append a new NoPartyIDs to this group
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

// Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

// NoSecurityAltID is a repeating group element, Tag 454
type NoSecurityAltID struct {
	*quickfix.Group
}

// SetSecurityAltID sets SecurityAltID, Tag 455
func (m NoSecurityAltID) SetSecurityAltID(v string) {
	m.Set(field.NewSecurityAltID(v))
}

// SetSecurityAltIDSource sets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) SetSecurityAltIDSource(v string) {
	m.Set(field.NewSecurityAltIDSource(v))
}

// GetSecurityAltID gets SecurityAltID, Tag 455
func (m NoSecurityAltID) GetSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) GetSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSecurityAltID returns true if SecurityAltID is present, Tag 455
func (m NoSecurityAltID) HasSecurityAltID() bool {
	return m.Has(tag.SecurityAltID)
}

// HasSecurityAltIDSource returns true if SecurityAltIDSource is present, Tag 456
func (m NoSecurityAltID) HasSecurityAltIDSource() bool {
	return m.Has(tag.SecurityAltIDSource)
}

// NoSecurityAltIDRepeatingGroup is a repeating group, Tag 454
type NoSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoSecurityAltIDRepeatingGroup returns an initialized, NoSecurityAltIDRepeatingGroup
func NewNoSecurityAltIDRepeatingGroup() NoSecurityAltIDRepeatingGroup {
	return NoSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SecurityAltID), quickfix.GroupElement(tag.SecurityAltIDSource)}),
	}
}

// Add create and append a new NoSecurityAltID to this group
func (m NoSecurityAltIDRepeatingGroup) Add() NoSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoSecurityAltID{g}
}

// Get returns the ith NoSecurityAltID in the NoSecurityAltIDRepeatinGroup
func (m NoSecurityAltIDRepeatingGroup) Get(i int) NoSecurityAltID {
	return NoSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoUnderlyings is a repeating group element, Tag 711
type NoUnderlyings struct {
	*quickfix.Group
}

// SetUnderlyingSymbol sets UnderlyingSymbol, Tag 311
func (m NoUnderlyings) SetUnderlyingSymbol(v string) {
	m.Set(field.NewUnderlyingSymbol(v))
}

// SetUnderlyingSymbolSfx sets UnderlyingSymbolSfx, Tag 312
func (m NoUnderlyings) SetUnderlyingSymbolSfx(v string) {
	m.Set(field.NewUnderlyingSymbolSfx(v))
}

// SetUnderlyingSecurityID sets UnderlyingSecurityID, Tag 309
func (m NoUnderlyings) SetUnderlyingSecurityID(v string) {
	m.Set(field.NewUnderlyingSecurityID(v))
}

// SetUnderlyingSecurityIDSource sets UnderlyingSecurityIDSource, Tag 305
func (m NoUnderlyings) SetUnderlyingSecurityIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityIDSource(v))
}

// SetNoUnderlyingSecurityAltID sets NoUnderlyingSecurityAltID, Tag 457
func (m NoUnderlyings) SetNoUnderlyingSecurityAltID(f NoUnderlyingSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetUnderlyingProduct sets UnderlyingProduct, Tag 462
func (m NoUnderlyings) SetUnderlyingProduct(v int) {
	m.Set(field.NewUnderlyingProduct(v))
}

// SetUnderlyingCFICode sets UnderlyingCFICode, Tag 463
func (m NoUnderlyings) SetUnderlyingCFICode(v string) {
	m.Set(field.NewUnderlyingCFICode(v))
}

// SetUnderlyingSecurityType sets UnderlyingSecurityType, Tag 310
func (m NoUnderlyings) SetUnderlyingSecurityType(v string) {
	m.Set(field.NewUnderlyingSecurityType(v))
}

// SetUnderlyingSecuritySubType sets UnderlyingSecuritySubType, Tag 763
func (m NoUnderlyings) SetUnderlyingSecuritySubType(v string) {
	m.Set(field.NewUnderlyingSecuritySubType(v))
}

// SetUnderlyingMaturityMonthYear sets UnderlyingMaturityMonthYear, Tag 313
func (m NoUnderlyings) SetUnderlyingMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingMaturityMonthYear(v))
}

// SetUnderlyingMaturityDate sets UnderlyingMaturityDate, Tag 542
func (m NoUnderlyings) SetUnderlyingMaturityDate(v string) {
	m.Set(field.NewUnderlyingMaturityDate(v))
}

// SetUnderlyingCouponPaymentDate sets UnderlyingCouponPaymentDate, Tag 241
func (m NoUnderlyings) SetUnderlyingCouponPaymentDate(v string) {
	m.Set(field.NewUnderlyingCouponPaymentDate(v))
}

// SetUnderlyingIssueDate sets UnderlyingIssueDate, Tag 242
func (m NoUnderlyings) SetUnderlyingIssueDate(v string) {
	m.Set(field.NewUnderlyingIssueDate(v))
}

// SetUnderlyingRepoCollateralSecurityType sets UnderlyingRepoCollateralSecurityType, Tag 243
func (m NoUnderlyings) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.Set(field.NewUnderlyingRepoCollateralSecurityType(v))
}

// SetUnderlyingRepurchaseTerm sets UnderlyingRepurchaseTerm, Tag 244
func (m NoUnderlyings) SetUnderlyingRepurchaseTerm(v int) {
	m.Set(field.NewUnderlyingRepurchaseTerm(v))
}

// SetUnderlyingRepurchaseRate sets UnderlyingRepurchaseRate, Tag 245
func (m NoUnderlyings) SetUnderlyingRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingRepurchaseRate(value, scale))
}

// SetUnderlyingFactor sets UnderlyingFactor, Tag 246
func (m NoUnderlyings) SetUnderlyingFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFactor(value, scale))
}

// SetUnderlyingCreditRating sets UnderlyingCreditRating, Tag 256
func (m NoUnderlyings) SetUnderlyingCreditRating(v string) {
	m.Set(field.NewUnderlyingCreditRating(v))
}

// SetUnderlyingInstrRegistry sets UnderlyingInstrRegistry, Tag 595
func (m NoUnderlyings) SetUnderlyingInstrRegistry(v string) {
	m.Set(field.NewUnderlyingInstrRegistry(v))
}

// SetUnderlyingCountryOfIssue sets UnderlyingCountryOfIssue, Tag 592
func (m NoUnderlyings) SetUnderlyingCountryOfIssue(v string) {
	m.Set(field.NewUnderlyingCountryOfIssue(v))
}

// SetUnderlyingStateOrProvinceOfIssue sets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m NoUnderlyings) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.Set(field.NewUnderlyingStateOrProvinceOfIssue(v))
}

// SetUnderlyingLocaleOfIssue sets UnderlyingLocaleOfIssue, Tag 594
func (m NoUnderlyings) SetUnderlyingLocaleOfIssue(v string) {
	m.Set(field.NewUnderlyingLocaleOfIssue(v))
}

// SetUnderlyingRedemptionDate sets UnderlyingRedemptionDate, Tag 247
func (m NoUnderlyings) SetUnderlyingRedemptionDate(v string) {
	m.Set(field.NewUnderlyingRedemptionDate(v))
}

// SetUnderlyingStrikePrice sets UnderlyingStrikePrice, Tag 316
func (m NoUnderlyings) SetUnderlyingStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStrikePrice(value, scale))
}

// SetUnderlyingStrikeCurrency sets UnderlyingStrikeCurrency, Tag 941
func (m NoUnderlyings) SetUnderlyingStrikeCurrency(v string) {
	m.Set(field.NewUnderlyingStrikeCurrency(v))
}

// SetUnderlyingOptAttribute sets UnderlyingOptAttribute, Tag 317
func (m NoUnderlyings) SetUnderlyingOptAttribute(v string) {
	m.Set(field.NewUnderlyingOptAttribute(v))
}

// SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436
func (m NoUnderlyings) SetUnderlyingContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingContractMultiplier(value, scale))
}

// SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435
func (m NoUnderlyings) SetUnderlyingCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCouponRate(value, scale))
}

// SetUnderlyingSecurityExchange sets UnderlyingSecurityExchange, Tag 308
func (m NoUnderlyings) SetUnderlyingSecurityExchange(v string) {
	m.Set(field.NewUnderlyingSecurityExchange(v))
}

// SetUnderlyingIssuer sets UnderlyingIssuer, Tag 306
func (m NoUnderlyings) SetUnderlyingIssuer(v string) {
	m.Set(field.NewUnderlyingIssuer(v))
}

// SetEncodedUnderlyingIssuerLen sets EncodedUnderlyingIssuerLen, Tag 362
func (m NoUnderlyings) SetEncodedUnderlyingIssuerLen(v int) {
	m.Set(field.NewEncodedUnderlyingIssuerLen(v))
}

// SetEncodedUnderlyingIssuer sets EncodedUnderlyingIssuer, Tag 363
func (m NoUnderlyings) SetEncodedUnderlyingIssuer(v string) {
	m.Set(field.NewEncodedUnderlyingIssuer(v))
}

// SetUnderlyingSecurityDesc sets UnderlyingSecurityDesc, Tag 307
func (m NoUnderlyings) SetUnderlyingSecurityDesc(v string) {
	m.Set(field.NewUnderlyingSecurityDesc(v))
}

// SetEncodedUnderlyingSecurityDescLen sets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.Set(field.NewEncodedUnderlyingSecurityDescLen(v))
}

// SetEncodedUnderlyingSecurityDesc sets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDesc(v string) {
	m.Set(field.NewEncodedUnderlyingSecurityDesc(v))
}

// SetUnderlyingCPProgram sets UnderlyingCPProgram, Tag 877
func (m NoUnderlyings) SetUnderlyingCPProgram(v string) {
	m.Set(field.NewUnderlyingCPProgram(v))
}

// SetUnderlyingCPRegType sets UnderlyingCPRegType, Tag 878
func (m NoUnderlyings) SetUnderlyingCPRegType(v string) {
	m.Set(field.NewUnderlyingCPRegType(v))
}

// SetUnderlyingCurrency sets UnderlyingCurrency, Tag 318
func (m NoUnderlyings) SetUnderlyingCurrency(v string) {
	m.Set(field.NewUnderlyingCurrency(v))
}

// SetUnderlyingQty sets UnderlyingQty, Tag 879
func (m NoUnderlyings) SetUnderlyingQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingQty(value, scale))
}

// SetUnderlyingPx sets UnderlyingPx, Tag 810
func (m NoUnderlyings) SetUnderlyingPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPx(value, scale))
}

// SetUnderlyingDirtyPrice sets UnderlyingDirtyPrice, Tag 882
func (m NoUnderlyings) SetUnderlyingDirtyPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDirtyPrice(value, scale))
}

// SetUnderlyingEndPrice sets UnderlyingEndPrice, Tag 883
func (m NoUnderlyings) SetUnderlyingEndPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndPrice(value, scale))
}

// SetUnderlyingStartValue sets UnderlyingStartValue, Tag 884
func (m NoUnderlyings) SetUnderlyingStartValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStartValue(value, scale))
}

// SetUnderlyingCurrentValue sets UnderlyingCurrentValue, Tag 885
func (m NoUnderlyings) SetUnderlyingCurrentValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCurrentValue(value, scale))
}

// SetUnderlyingEndValue sets UnderlyingEndValue, Tag 886
func (m NoUnderlyings) SetUnderlyingEndValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndValue(value, scale))
}

// SetNoUnderlyingStips sets NoUnderlyingStips, Tag 887
func (m NoUnderlyings) SetNoUnderlyingStips(f NoUnderlyingStipsRepeatingGroup) {
	m.SetGroup(f)
}

// GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311
func (m NoUnderlyings) GetUnderlyingSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312
func (m NoUnderlyings) GetUnderlyingSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309
func (m NoUnderlyings) GetUnderlyingSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityIDSource gets UnderlyingSecurityIDSource, Tag 305
func (m NoUnderlyings) GetUnderlyingSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUnderlyingSecurityAltID gets NoUnderlyingSecurityAltID, Tag 457
func (m NoUnderlyings) GetNoUnderlyingSecurityAltID() (NoUnderlyingSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUnderlyingProduct gets UnderlyingProduct, Tag 462
func (m NoUnderlyings) GetUnderlyingProduct() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCFICode gets UnderlyingCFICode, Tag 463
func (m NoUnderlyings) GetUnderlyingCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310
func (m NoUnderlyings) GetUnderlyingSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecuritySubType gets UnderlyingSecuritySubType, Tag 763
func (m NoUnderlyings) GetUnderlyingSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313
func (m NoUnderlyings) GetUnderlyingMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingMaturityDate gets UnderlyingMaturityDate, Tag 542
func (m NoUnderlyings) GetUnderlyingMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCouponPaymentDate gets UnderlyingCouponPaymentDate, Tag 241
func (m NoUnderlyings) GetUnderlyingCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingIssueDate gets UnderlyingIssueDate, Tag 242
func (m NoUnderlyings) GetUnderlyingIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepoCollateralSecurityType gets UnderlyingRepoCollateralSecurityType, Tag 243
func (m NoUnderlyings) GetUnderlyingRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepurchaseTerm gets UnderlyingRepurchaseTerm, Tag 244
func (m NoUnderlyings) GetUnderlyingRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepurchaseRate gets UnderlyingRepurchaseRate, Tag 245
func (m NoUnderlyings) GetUnderlyingRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingFactor gets UnderlyingFactor, Tag 246
func (m NoUnderlyings) GetUnderlyingFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingFactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCreditRating gets UnderlyingCreditRating, Tag 256
func (m NoUnderlyings) GetUnderlyingCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingInstrRegistry gets UnderlyingInstrRegistry, Tag 595
func (m NoUnderlyings) GetUnderlyingInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCountryOfIssue gets UnderlyingCountryOfIssue, Tag 592
func (m NoUnderlyings) GetUnderlyingCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStateOrProvinceOfIssue gets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m NoUnderlyings) GetUnderlyingStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingLocaleOfIssue gets UnderlyingLocaleOfIssue, Tag 594
func (m NoUnderlyings) GetUnderlyingLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRedemptionDate gets UnderlyingRedemptionDate, Tag 247
func (m NoUnderlyings) GetUnderlyingRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316
func (m NoUnderlyings) GetUnderlyingStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStrikeCurrency gets UnderlyingStrikeCurrency, Tag 941
func (m NoUnderlyings) GetUnderlyingStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317
func (m NoUnderlyings) GetUnderlyingOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436
func (m NoUnderlyings) GetUnderlyingContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435
func (m NoUnderlyings) GetUnderlyingCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308
func (m NoUnderlyings) GetUnderlyingSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306
func (m NoUnderlyings) GetUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362
func (m NoUnderlyings) GetEncodedUnderlyingIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363
func (m NoUnderlyings) GetEncodedUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307
func (m NoUnderlyings) GetUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCPProgram gets UnderlyingCPProgram, Tag 877
func (m NoUnderlyings) GetUnderlyingCPProgram() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCPRegType gets UnderlyingCPRegType, Tag 878
func (m NoUnderlyings) GetUnderlyingCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCurrency gets UnderlyingCurrency, Tag 318
func (m NoUnderlyings) GetUnderlyingCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingQty gets UnderlyingQty, Tag 879
func (m NoUnderlyings) GetUnderlyingQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPx gets UnderlyingPx, Tag 810
func (m NoUnderlyings) GetUnderlyingPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingDirtyPrice gets UnderlyingDirtyPrice, Tag 882
func (m NoUnderlyings) GetUnderlyingDirtyPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingDirtyPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingEndPrice gets UnderlyingEndPrice, Tag 883
func (m NoUnderlyings) GetUnderlyingEndPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStartValue gets UnderlyingStartValue, Tag 884
func (m NoUnderlyings) GetUnderlyingStartValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStartValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCurrentValue gets UnderlyingCurrentValue, Tag 885
func (m NoUnderlyings) GetUnderlyingCurrentValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrentValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingEndValue gets UnderlyingEndValue, Tag 886
func (m NoUnderlyings) GetUnderlyingEndValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUnderlyingStips gets NoUnderlyingStips, Tag 887
func (m NoUnderlyings) GetNoUnderlyingStips() (NoUnderlyingStipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingStipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasUnderlyingSymbol returns true if UnderlyingSymbol is present, Tag 311
func (m NoUnderlyings) HasUnderlyingSymbol() bool {
	return m.Has(tag.UnderlyingSymbol)
}

// HasUnderlyingSymbolSfx returns true if UnderlyingSymbolSfx is present, Tag 312
func (m NoUnderlyings) HasUnderlyingSymbolSfx() bool {
	return m.Has(tag.UnderlyingSymbolSfx)
}

// HasUnderlyingSecurityID returns true if UnderlyingSecurityID is present, Tag 309
func (m NoUnderlyings) HasUnderlyingSecurityID() bool {
	return m.Has(tag.UnderlyingSecurityID)
}

// HasUnderlyingSecurityIDSource returns true if UnderlyingSecurityIDSource is present, Tag 305
func (m NoUnderlyings) HasUnderlyingSecurityIDSource() bool {
	return m.Has(tag.UnderlyingSecurityIDSource)
}

// HasNoUnderlyingSecurityAltID returns true if NoUnderlyingSecurityAltID is present, Tag 457
func (m NoUnderlyings) HasNoUnderlyingSecurityAltID() bool {
	return m.Has(tag.NoUnderlyingSecurityAltID)
}

// HasUnderlyingProduct returns true if UnderlyingProduct is present, Tag 462
func (m NoUnderlyings) HasUnderlyingProduct() bool {
	return m.Has(tag.UnderlyingProduct)
}

// HasUnderlyingCFICode returns true if UnderlyingCFICode is present, Tag 463
func (m NoUnderlyings) HasUnderlyingCFICode() bool {
	return m.Has(tag.UnderlyingCFICode)
}

// HasUnderlyingSecurityType returns true if UnderlyingSecurityType is present, Tag 310
func (m NoUnderlyings) HasUnderlyingSecurityType() bool {
	return m.Has(tag.UnderlyingSecurityType)
}

// HasUnderlyingSecuritySubType returns true if UnderlyingSecuritySubType is present, Tag 763
func (m NoUnderlyings) HasUnderlyingSecuritySubType() bool {
	return m.Has(tag.UnderlyingSecuritySubType)
}

// HasUnderlyingMaturityMonthYear returns true if UnderlyingMaturityMonthYear is present, Tag 313
func (m NoUnderlyings) HasUnderlyingMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingMaturityMonthYear)
}

// HasUnderlyingMaturityDate returns true if UnderlyingMaturityDate is present, Tag 542
func (m NoUnderlyings) HasUnderlyingMaturityDate() bool {
	return m.Has(tag.UnderlyingMaturityDate)
}

// HasUnderlyingCouponPaymentDate returns true if UnderlyingCouponPaymentDate is present, Tag 241
func (m NoUnderlyings) HasUnderlyingCouponPaymentDate() bool {
	return m.Has(tag.UnderlyingCouponPaymentDate)
}

// HasUnderlyingIssueDate returns true if UnderlyingIssueDate is present, Tag 242
func (m NoUnderlyings) HasUnderlyingIssueDate() bool {
	return m.Has(tag.UnderlyingIssueDate)
}

// HasUnderlyingRepoCollateralSecurityType returns true if UnderlyingRepoCollateralSecurityType is present, Tag 243
func (m NoUnderlyings) HasUnderlyingRepoCollateralSecurityType() bool {
	return m.Has(tag.UnderlyingRepoCollateralSecurityType)
}

// HasUnderlyingRepurchaseTerm returns true if UnderlyingRepurchaseTerm is present, Tag 244
func (m NoUnderlyings) HasUnderlyingRepurchaseTerm() bool {
	return m.Has(tag.UnderlyingRepurchaseTerm)
}

// HasUnderlyingRepurchaseRate returns true if UnderlyingRepurchaseRate is present, Tag 245
func (m NoUnderlyings) HasUnderlyingRepurchaseRate() bool {
	return m.Has(tag.UnderlyingRepurchaseRate)
}

// HasUnderlyingFactor returns true if UnderlyingFactor is present, Tag 246
func (m NoUnderlyings) HasUnderlyingFactor() bool {
	return m.Has(tag.UnderlyingFactor)
}

// HasUnderlyingCreditRating returns true if UnderlyingCreditRating is present, Tag 256
func (m NoUnderlyings) HasUnderlyingCreditRating() bool {
	return m.Has(tag.UnderlyingCreditRating)
}

// HasUnderlyingInstrRegistry returns true if UnderlyingInstrRegistry is present, Tag 595
func (m NoUnderlyings) HasUnderlyingInstrRegistry() bool {
	return m.Has(tag.UnderlyingInstrRegistry)
}

// HasUnderlyingCountryOfIssue returns true if UnderlyingCountryOfIssue is present, Tag 592
func (m NoUnderlyings) HasUnderlyingCountryOfIssue() bool {
	return m.Has(tag.UnderlyingCountryOfIssue)
}

// HasUnderlyingStateOrProvinceOfIssue returns true if UnderlyingStateOrProvinceOfIssue is present, Tag 593
func (m NoUnderlyings) HasUnderlyingStateOrProvinceOfIssue() bool {
	return m.Has(tag.UnderlyingStateOrProvinceOfIssue)
}

// HasUnderlyingLocaleOfIssue returns true if UnderlyingLocaleOfIssue is present, Tag 594
func (m NoUnderlyings) HasUnderlyingLocaleOfIssue() bool {
	return m.Has(tag.UnderlyingLocaleOfIssue)
}

// HasUnderlyingRedemptionDate returns true if UnderlyingRedemptionDate is present, Tag 247
func (m NoUnderlyings) HasUnderlyingRedemptionDate() bool {
	return m.Has(tag.UnderlyingRedemptionDate)
}

// HasUnderlyingStrikePrice returns true if UnderlyingStrikePrice is present, Tag 316
func (m NoUnderlyings) HasUnderlyingStrikePrice() bool {
	return m.Has(tag.UnderlyingStrikePrice)
}

// HasUnderlyingStrikeCurrency returns true if UnderlyingStrikeCurrency is present, Tag 941
func (m NoUnderlyings) HasUnderlyingStrikeCurrency() bool {
	return m.Has(tag.UnderlyingStrikeCurrency)
}

// HasUnderlyingOptAttribute returns true if UnderlyingOptAttribute is present, Tag 317
func (m NoUnderlyings) HasUnderlyingOptAttribute() bool {
	return m.Has(tag.UnderlyingOptAttribute)
}

// HasUnderlyingContractMultiplier returns true if UnderlyingContractMultiplier is present, Tag 436
func (m NoUnderlyings) HasUnderlyingContractMultiplier() bool {
	return m.Has(tag.UnderlyingContractMultiplier)
}

// HasUnderlyingCouponRate returns true if UnderlyingCouponRate is present, Tag 435
func (m NoUnderlyings) HasUnderlyingCouponRate() bool {
	return m.Has(tag.UnderlyingCouponRate)
}

// HasUnderlyingSecurityExchange returns true if UnderlyingSecurityExchange is present, Tag 308
func (m NoUnderlyings) HasUnderlyingSecurityExchange() bool {
	return m.Has(tag.UnderlyingSecurityExchange)
}

// HasUnderlyingIssuer returns true if UnderlyingIssuer is present, Tag 306
func (m NoUnderlyings) HasUnderlyingIssuer() bool {
	return m.Has(tag.UnderlyingIssuer)
}

// HasEncodedUnderlyingIssuerLen returns true if EncodedUnderlyingIssuerLen is present, Tag 362
func (m NoUnderlyings) HasEncodedUnderlyingIssuerLen() bool {
	return m.Has(tag.EncodedUnderlyingIssuerLen)
}

// HasEncodedUnderlyingIssuer returns true if EncodedUnderlyingIssuer is present, Tag 363
func (m NoUnderlyings) HasEncodedUnderlyingIssuer() bool {
	return m.Has(tag.EncodedUnderlyingIssuer)
}

// HasUnderlyingSecurityDesc returns true if UnderlyingSecurityDesc is present, Tag 307
func (m NoUnderlyings) HasUnderlyingSecurityDesc() bool {
	return m.Has(tag.UnderlyingSecurityDesc)
}

// HasEncodedUnderlyingSecurityDescLen returns true if EncodedUnderlyingSecurityDescLen is present, Tag 364
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDescLen() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDescLen)
}

// HasEncodedUnderlyingSecurityDesc returns true if EncodedUnderlyingSecurityDesc is present, Tag 365
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDesc() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDesc)
}

// HasUnderlyingCPProgram returns true if UnderlyingCPProgram is present, Tag 877
func (m NoUnderlyings) HasUnderlyingCPProgram() bool {
	return m.Has(tag.UnderlyingCPProgram)
}

// HasUnderlyingCPRegType returns true if UnderlyingCPRegType is present, Tag 878
func (m NoUnderlyings) HasUnderlyingCPRegType() bool {
	return m.Has(tag.UnderlyingCPRegType)
}

// HasUnderlyingCurrency returns true if UnderlyingCurrency is present, Tag 318
func (m NoUnderlyings) HasUnderlyingCurrency() bool {
	return m.Has(tag.UnderlyingCurrency)
}

// HasUnderlyingQty returns true if UnderlyingQty is present, Tag 879
func (m NoUnderlyings) HasUnderlyingQty() bool {
	return m.Has(tag.UnderlyingQty)
}

// HasUnderlyingPx returns true if UnderlyingPx is present, Tag 810
func (m NoUnderlyings) HasUnderlyingPx() bool {
	return m.Has(tag.UnderlyingPx)
}

// HasUnderlyingDirtyPrice returns true if UnderlyingDirtyPrice is present, Tag 882
func (m NoUnderlyings) HasUnderlyingDirtyPrice() bool {
	return m.Has(tag.UnderlyingDirtyPrice)
}

// HasUnderlyingEndPrice returns true if UnderlyingEndPrice is present, Tag 883
func (m NoUnderlyings) HasUnderlyingEndPrice() bool {
	return m.Has(tag.UnderlyingEndPrice)
}

// HasUnderlyingStartValue returns true if UnderlyingStartValue is present, Tag 884
func (m NoUnderlyings) HasUnderlyingStartValue() bool {
	return m.Has(tag.UnderlyingStartValue)
}

// HasUnderlyingCurrentValue returns true if UnderlyingCurrentValue is present, Tag 885
func (m NoUnderlyings) HasUnderlyingCurrentValue() bool {
	return m.Has(tag.UnderlyingCurrentValue)
}

// HasUnderlyingEndValue returns true if UnderlyingEndValue is present, Tag 886
func (m NoUnderlyings) HasUnderlyingEndValue() bool {
	return m.Has(tag.UnderlyingEndValue)
}

// HasNoUnderlyingStips returns true if NoUnderlyingStips is present, Tag 887
func (m NoUnderlyings) HasNoUnderlyingStips() bool {
	return m.Has(tag.NoUnderlyingStips)
}

// NoUnderlyingSecurityAltID is a repeating group element, Tag 457
type NoUnderlyingSecurityAltID struct {
	*quickfix.Group
}

// SetUnderlyingSecurityAltID sets UnderlyingSecurityAltID, Tag 458
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltID(v string) {
	m.Set(field.NewUnderlyingSecurityAltID(v))
}

// SetUnderlyingSecurityAltIDSource sets UnderlyingSecurityAltIDSource, Tag 459
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityAltIDSource(v))
}

// GetUnderlyingSecurityAltID gets UnderlyingSecurityAltID, Tag 458
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityAltIDSource gets UnderlyingSecurityAltIDSource, Tag 459
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingSecurityAltID returns true if UnderlyingSecurityAltID is present, Tag 458
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltID() bool {
	return m.Has(tag.UnderlyingSecurityAltID)
}

// HasUnderlyingSecurityAltIDSource returns true if UnderlyingSecurityAltIDSource is present, Tag 459
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltIDSource() bool {
	return m.Has(tag.UnderlyingSecurityAltIDSource)
}

// NoUnderlyingSecurityAltIDRepeatingGroup is a repeating group, Tag 457
type NoUnderlyingSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingSecurityAltIDRepeatingGroup returns an initialized, NoUnderlyingSecurityAltIDRepeatingGroup
func NewNoUnderlyingSecurityAltIDRepeatingGroup() NoUnderlyingSecurityAltIDRepeatingGroup {
	return NoUnderlyingSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSecurityAltID), quickfix.GroupElement(tag.UnderlyingSecurityAltIDSource)}),
	}
}

// Add create and append a new NoUnderlyingSecurityAltID to this group
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Add() NoUnderlyingSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingSecurityAltID{g}
}

// Get returns the ith NoUnderlyingSecurityAltID in the NoUnderlyingSecurityAltIDRepeatinGroup
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Get(i int) NoUnderlyingSecurityAltID {
	return NoUnderlyingSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoUnderlyingStips is a repeating group element, Tag 887
type NoUnderlyingStips struct {
	*quickfix.Group
}

// SetUnderlyingStipType sets UnderlyingStipType, Tag 888
func (m NoUnderlyingStips) SetUnderlyingStipType(v string) {
	m.Set(field.NewUnderlyingStipType(v))
}

// SetUnderlyingStipValue sets UnderlyingStipValue, Tag 889
func (m NoUnderlyingStips) SetUnderlyingStipValue(v string) {
	m.Set(field.NewUnderlyingStipValue(v))
}

// GetUnderlyingStipType gets UnderlyingStipType, Tag 888
func (m NoUnderlyingStips) GetUnderlyingStipType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStipTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStipValue gets UnderlyingStipValue, Tag 889
func (m NoUnderlyingStips) GetUnderlyingStipValue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStipValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingStipType returns true if UnderlyingStipType is present, Tag 888
func (m NoUnderlyingStips) HasUnderlyingStipType() bool {
	return m.Has(tag.UnderlyingStipType)
}

// HasUnderlyingStipValue returns true if UnderlyingStipValue is present, Tag 889
func (m NoUnderlyingStips) HasUnderlyingStipValue() bool {
	return m.Has(tag.UnderlyingStipValue)
}

// NoUnderlyingStipsRepeatingGroup is a repeating group, Tag 887
type NoUnderlyingStipsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingStipsRepeatingGroup returns an initialized, NoUnderlyingStipsRepeatingGroup
func NewNoUnderlyingStipsRepeatingGroup() NoUnderlyingStipsRepeatingGroup {
	return NoUnderlyingStipsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingStips,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingStipType), quickfix.GroupElement(tag.UnderlyingStipValue)}),
	}
}

// Add create and append a new NoUnderlyingStips to this group
func (m NoUnderlyingStipsRepeatingGroup) Add() NoUnderlyingStips {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingStips{g}
}

// Get returns the ith NoUnderlyingStips in the NoUnderlyingStipsRepeatinGroup
func (m NoUnderlyingStipsRepeatingGroup) Get(i int) NoUnderlyingStips {
	return NoUnderlyingStips{m.RepeatingGroup.Get(i)}
}

// NoUnderlyingsRepeatingGroup is a repeating group, Tag 711
type NoUnderlyingsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingsRepeatingGroup returns an initialized, NoUnderlyingsRepeatingGroup
func NewNoUnderlyingsRepeatingGroup() NoUnderlyingsRepeatingGroup {
	return NoUnderlyingsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyings,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingSecurityIDSource), NewNoUnderlyingSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingProduct), quickfix.GroupElement(tag.UnderlyingCFICode), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingSecuritySubType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDate), quickfix.GroupElement(tag.UnderlyingCouponPaymentDate), quickfix.GroupElement(tag.UnderlyingIssueDate), quickfix.GroupElement(tag.UnderlyingRepoCollateralSecurityType), quickfix.GroupElement(tag.UnderlyingRepurchaseTerm), quickfix.GroupElement(tag.UnderlyingRepurchaseRate), quickfix.GroupElement(tag.UnderlyingFactor), quickfix.GroupElement(tag.UnderlyingCreditRating), quickfix.GroupElement(tag.UnderlyingInstrRegistry), quickfix.GroupElement(tag.UnderlyingCountryOfIssue), quickfix.GroupElement(tag.UnderlyingStateOrProvinceOfIssue), quickfix.GroupElement(tag.UnderlyingLocaleOfIssue), quickfix.GroupElement(tag.UnderlyingRedemptionDate), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingStrikeCurrency), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.UnderlyingCPProgram), quickfix.GroupElement(tag.UnderlyingCPRegType), quickfix.GroupElement(tag.UnderlyingCurrency), quickfix.GroupElement(tag.UnderlyingQty), quickfix.GroupElement(tag.UnderlyingPx), quickfix.GroupElement(tag.UnderlyingDirtyPrice), quickfix.GroupElement(tag.UnderlyingEndPrice), quickfix.GroupElement(tag.UnderlyingStartValue), quickfix.GroupElement(tag.UnderlyingCurrentValue), quickfix.GroupElement(tag.UnderlyingEndValue), NewNoUnderlyingStipsRepeatingGroup()}),
	}
}

// Add create and append a new NoUnderlyings to this group
func (m NoUnderlyingsRepeatingGroup) Add() NoUnderlyings {
	g := m.RepeatingGroup.Add()
	return NoUnderlyings{g}
}

// Get returns the ith NoUnderlyings in the NoUnderlyingsRepeatinGroup
func (m NoUnderlyingsRepeatingGroup) Get(i int) NoUnderlyings {
	return NoUnderlyings{m.RepeatingGroup.Get(i)}
}

// NoEvents is a repeating group element, Tag 864
type NoEvents struct {
	*quickfix.Group
}

// SetEventType sets EventType, Tag 865
func (m NoEvents) SetEventType(v enum.EventType) {
	m.Set(field.NewEventType(v))
}

// SetEventDate sets EventDate, Tag 866
func (m NoEvents) SetEventDate(v string) {
	m.Set(field.NewEventDate(v))
}

// SetEventPx sets EventPx, Tag 867
func (m NoEvents) SetEventPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewEventPx(value, scale))
}

// SetEventText sets EventText, Tag 868
func (m NoEvents) SetEventText(v string) {
	m.Set(field.NewEventText(v))
}

// GetEventType gets EventType, Tag 865
func (m NoEvents) GetEventType() (v enum.EventType, err quickfix.MessageRejectError) {
	var f field.EventTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventDate gets EventDate, Tag 866
func (m NoEvents) GetEventDate() (v string, err quickfix.MessageRejectError) {
	var f field.EventDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventPx gets EventPx, Tag 867
func (m NoEvents) GetEventPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.EventPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventText gets EventText, Tag 868
func (m NoEvents) GetEventText() (v string, err quickfix.MessageRejectError) {
	var f field.EventTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasEventType returns true if EventType is present, Tag 865
func (m NoEvents) HasEventType() bool {
	return m.Has(tag.EventType)
}

// HasEventDate returns true if EventDate is present, Tag 866
func (m NoEvents) HasEventDate() bool {
	return m.Has(tag.EventDate)
}

// HasEventPx returns true if EventPx is present, Tag 867
func (m NoEvents) HasEventPx() bool {
	return m.Has(tag.EventPx)
}

// HasEventText returns true if EventText is present, Tag 868
func (m NoEvents) HasEventText() bool {
	return m.Has(tag.EventText)
}

// NoEventsRepeatingGroup is a repeating group, Tag 864
type NoEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoEventsRepeatingGroup returns an initialized, NoEventsRepeatingGroup
func NewNoEventsRepeatingGroup() NoEventsRepeatingGroup {
	return NoEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.EventType), quickfix.GroupElement(tag.EventDate), quickfix.GroupElement(tag.EventPx), quickfix.GroupElement(tag.EventText)}),
	}
}

// Add create and append a new NoEvents to this group
func (m NoEventsRepeatingGroup) Add() NoEvents {
	g := m.RepeatingGroup.Add()
	return NoEvents{g}
}

// Get returns the ith NoEvents in the NoEventsRepeatinGroup
func (m NoEventsRepeatingGroup) Get(i int) NoEvents {
	return NoEvents{m.RepeatingGroup.Get(i)}
}
