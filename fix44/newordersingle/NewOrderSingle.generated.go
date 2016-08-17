package newordersingle

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/tag"
)

//NewOrderSingle is the fix44 NewOrderSingle type, MsgType = D
type NewOrderSingle struct {
	fix44.Header
	quickfix.Body
	fix44.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a NewOrderSingle from a quickfix.Message instance
func FromMessage(m quickfix.Message) NewOrderSingle {
	return NewOrderSingle{
		Header:      fix44.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix44.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m NewOrderSingle) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a NewOrderSingle initialized with the required fields for NewOrderSingle
func New(clordid field.ClOrdIDField, side field.SideField, transacttime field.TransactTimeField, ordtype field.OrdTypeField) (m NewOrderSingle) {
	m.Header = fix44.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("D"))
	m.Set(clordid)
	m.Set(side)
	m.Set(transacttime)
	m.Set(ordtype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg NewOrderSingle, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "D", r
}

//SetAccount sets Account, Tag 1
func (m NewOrderSingle) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m NewOrderSingle) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetCommission sets Commission, Tag 12
func (m NewOrderSingle) SetCommission(value decimal.Decimal, scale int32) {
	m.Set(field.NewCommission(value, scale))
}

//SetCommType sets CommType, Tag 13
func (m NewOrderSingle) SetCommType(v string) {
	m.Set(field.NewCommType(v))
}

//SetCurrency sets Currency, Tag 15
func (m NewOrderSingle) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetExecInst sets ExecInst, Tag 18
func (m NewOrderSingle) SetExecInst(v string) {
	m.Set(field.NewExecInst(v))
}

//SetHandlInst sets HandlInst, Tag 21
func (m NewOrderSingle) SetHandlInst(v string) {
	m.Set(field.NewHandlInst(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m NewOrderSingle) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetIOIID sets IOIID, Tag 23
func (m NewOrderSingle) SetIOIID(v string) {
	m.Set(field.NewIOIID(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m NewOrderSingle) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

//SetOrdType sets OrdType, Tag 40
func (m NewOrderSingle) SetOrdType(v string) {
	m.Set(field.NewOrdType(v))
}

//SetPrice sets Price, Tag 44
func (m NewOrderSingle) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

//SetSecurityID sets SecurityID, Tag 48
func (m NewOrderSingle) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m NewOrderSingle) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m NewOrderSingle) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m NewOrderSingle) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTimeInForce sets TimeInForce, Tag 59
func (m NewOrderSingle) SetTimeInForce(v string) {
	m.Set(field.NewTimeInForce(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m NewOrderSingle) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSettlType sets SettlType, Tag 63
func (m NewOrderSingle) SetSettlType(v string) {
	m.Set(field.NewSettlType(v))
}

//SetSettlDate sets SettlDate, Tag 64
func (m NewOrderSingle) SetSettlDate(v string) {
	m.Set(field.NewSettlDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m NewOrderSingle) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetAllocID sets AllocID, Tag 70
func (m NewOrderSingle) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m NewOrderSingle) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetPositionEffect sets PositionEffect, Tag 77
func (m NewOrderSingle) SetPositionEffect(v string) {
	m.Set(field.NewPositionEffect(v))
}

//SetNoAllocs sets NoAllocs, Tag 78
func (m NewOrderSingle) SetNoAllocs(f NoAllocsRepeatingGroup) {
	m.SetGroup(f)
}

//SetProcessCode sets ProcessCode, Tag 81
func (m NewOrderSingle) SetProcessCode(v string) {
	m.Set(field.NewProcessCode(v))
}

//SetStopPx sets StopPx, Tag 99
func (m NewOrderSingle) SetStopPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewStopPx(value, scale))
}

//SetExDestination sets ExDestination, Tag 100
func (m NewOrderSingle) SetExDestination(v string) {
	m.Set(field.NewExDestination(v))
}

//SetIssuer sets Issuer, Tag 106
func (m NewOrderSingle) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m NewOrderSingle) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetMinQty sets MinQty, Tag 110
func (m NewOrderSingle) SetMinQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinQty(value, scale))
}

//SetMaxFloor sets MaxFloor, Tag 111
func (m NewOrderSingle) SetMaxFloor(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxFloor(value, scale))
}

//SetLocateReqd sets LocateReqd, Tag 114
func (m NewOrderSingle) SetLocateReqd(v bool) {
	m.Set(field.NewLocateReqd(v))
}

//SetQuoteID sets QuoteID, Tag 117
func (m NewOrderSingle) SetQuoteID(v string) {
	m.Set(field.NewQuoteID(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m NewOrderSingle) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetForexReq sets ForexReq, Tag 121
func (m NewOrderSingle) SetForexReq(v bool) {
	m.Set(field.NewForexReq(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m NewOrderSingle) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetPrevClosePx sets PrevClosePx, Tag 140
func (m NewOrderSingle) SetPrevClosePx(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrevClosePx(value, scale))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m NewOrderSingle) SetCashOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCashOrderQty(value, scale))
}

//SetSecurityType sets SecurityType, Tag 167
func (m NewOrderSingle) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetEffectiveTime sets EffectiveTime, Tag 168
func (m NewOrderSingle) SetEffectiveTime(v time.Time) {
	m.Set(field.NewEffectiveTime(v))
}

//SetOrderQty2 sets OrderQty2, Tag 192
func (m NewOrderSingle) SetOrderQty2(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty2(value, scale))
}

//SetSettlDate2 sets SettlDate2, Tag 193
func (m NewOrderSingle) SetSettlDate2(v string) {
	m.Set(field.NewSettlDate2(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m NewOrderSingle) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m NewOrderSingle) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetCoveredOrUncovered sets CoveredOrUncovered, Tag 203
func (m NewOrderSingle) SetCoveredOrUncovered(v int) {
	m.Set(field.NewCoveredOrUncovered(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m NewOrderSingle) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m NewOrderSingle) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetMaxShow sets MaxShow, Tag 210
func (m NewOrderSingle) SetMaxShow(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxShow(value, scale))
}

//SetPegOffsetValue sets PegOffsetValue, Tag 211
func (m NewOrderSingle) SetPegOffsetValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewPegOffsetValue(value, scale))
}

//SetSpread sets Spread, Tag 218
func (m NewOrderSingle) SetSpread(value decimal.Decimal, scale int32) {
	m.Set(field.NewSpread(value, scale))
}

//SetBenchmarkCurveCurrency sets BenchmarkCurveCurrency, Tag 220
func (m NewOrderSingle) SetBenchmarkCurveCurrency(v string) {
	m.Set(field.NewBenchmarkCurveCurrency(v))
}

//SetBenchmarkCurveName sets BenchmarkCurveName, Tag 221
func (m NewOrderSingle) SetBenchmarkCurveName(v string) {
	m.Set(field.NewBenchmarkCurveName(v))
}

//SetBenchmarkCurvePoint sets BenchmarkCurvePoint, Tag 222
func (m NewOrderSingle) SetBenchmarkCurvePoint(v string) {
	m.Set(field.NewBenchmarkCurvePoint(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m NewOrderSingle) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m NewOrderSingle) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m NewOrderSingle) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m NewOrderSingle) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m NewOrderSingle) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

//SetFactor sets Factor, Tag 228
func (m NewOrderSingle) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

//SetTradeOriginationDate sets TradeOriginationDate, Tag 229
func (m NewOrderSingle) SetTradeOriginationDate(v string) {
	m.Set(field.NewTradeOriginationDate(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m NewOrderSingle) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetNoStipulations sets NoStipulations, Tag 232
func (m NewOrderSingle) SetNoStipulations(f NoStipulationsRepeatingGroup) {
	m.SetGroup(f)
}

//SetYieldType sets YieldType, Tag 235
func (m NewOrderSingle) SetYieldType(v string) {
	m.Set(field.NewYieldType(v))
}

//SetYield sets Yield, Tag 236
func (m NewOrderSingle) SetYield(value decimal.Decimal, scale int32) {
	m.Set(field.NewYield(value, scale))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m NewOrderSingle) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m NewOrderSingle) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m NewOrderSingle) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m NewOrderSingle) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m NewOrderSingle) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m NewOrderSingle) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m NewOrderSingle) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m NewOrderSingle) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m NewOrderSingle) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetComplianceID sets ComplianceID, Tag 376
func (m NewOrderSingle) SetComplianceID(v string) {
	m.Set(field.NewComplianceID(v))
}

//SetSolicitedFlag sets SolicitedFlag, Tag 377
func (m NewOrderSingle) SetSolicitedFlag(v bool) {
	m.Set(field.NewSolicitedFlag(v))
}

//SetNoTradingSessions sets NoTradingSessions, Tag 386
func (m NewOrderSingle) SetNoTradingSessions(f NoTradingSessionsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDiscretionInst sets DiscretionInst, Tag 388
func (m NewOrderSingle) SetDiscretionInst(v string) {
	m.Set(field.NewDiscretionInst(v))
}

//SetDiscretionOffsetValue sets DiscretionOffsetValue, Tag 389
func (m NewOrderSingle) SetDiscretionOffsetValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewDiscretionOffsetValue(value, scale))
}

//SetPriceType sets PriceType, Tag 423
func (m NewOrderSingle) SetPriceType(v int) {
	m.Set(field.NewPriceType(v))
}

//SetGTBookingInst sets GTBookingInst, Tag 427
func (m NewOrderSingle) SetGTBookingInst(v int) {
	m.Set(field.NewGTBookingInst(v))
}

//SetExpireDate sets ExpireDate, Tag 432
func (m NewOrderSingle) SetExpireDate(v string) {
	m.Set(field.NewExpireDate(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m NewOrderSingle) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m NewOrderSingle) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m NewOrderSingle) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m NewOrderSingle) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetRoundingDirection sets RoundingDirection, Tag 468
func (m NewOrderSingle) SetRoundingDirection(v string) {
	m.Set(field.NewRoundingDirection(v))
}

//SetRoundingModulus sets RoundingModulus, Tag 469
func (m NewOrderSingle) SetRoundingModulus(value decimal.Decimal, scale int32) {
	m.Set(field.NewRoundingModulus(value, scale))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m NewOrderSingle) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m NewOrderSingle) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m NewOrderSingle) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetCommCurrency sets CommCurrency, Tag 479
func (m NewOrderSingle) SetCommCurrency(v string) {
	m.Set(field.NewCommCurrency(v))
}

//SetCancellationRights sets CancellationRights, Tag 480
func (m NewOrderSingle) SetCancellationRights(v string) {
	m.Set(field.NewCancellationRights(v))
}

//SetMoneyLaunderingStatus sets MoneyLaunderingStatus, Tag 481
func (m NewOrderSingle) SetMoneyLaunderingStatus(v string) {
	m.Set(field.NewMoneyLaunderingStatus(v))
}

//SetDesignation sets Designation, Tag 494
func (m NewOrderSingle) SetDesignation(v string) {
	m.Set(field.NewDesignation(v))
}

//SetFundRenewWaiv sets FundRenewWaiv, Tag 497
func (m NewOrderSingle) SetFundRenewWaiv(v string) {
	m.Set(field.NewFundRenewWaiv(v))
}

//SetRegistID sets RegistID, Tag 513
func (m NewOrderSingle) SetRegistID(v string) {
	m.Set(field.NewRegistID(v))
}

//SetOrderPercent sets OrderPercent, Tag 516
func (m NewOrderSingle) SetOrderPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderPercent(value, scale))
}

//SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526
func (m NewOrderSingle) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

//SetOrderCapacity sets OrderCapacity, Tag 528
func (m NewOrderSingle) SetOrderCapacity(v string) {
	m.Set(field.NewOrderCapacity(v))
}

//SetOrderRestrictions sets OrderRestrictions, Tag 529
func (m NewOrderSingle) SetOrderRestrictions(v string) {
	m.Set(field.NewOrderRestrictions(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m NewOrderSingle) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m NewOrderSingle) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetCashMargin sets CashMargin, Tag 544
func (m NewOrderSingle) SetCashMargin(v string) {
	m.Set(field.NewCashMargin(v))
}

//SetAccountType sets AccountType, Tag 581
func (m NewOrderSingle) SetAccountType(v int) {
	m.Set(field.NewAccountType(v))
}

//SetCustOrderCapacity sets CustOrderCapacity, Tag 582
func (m NewOrderSingle) SetCustOrderCapacity(v int) {
	m.Set(field.NewCustOrderCapacity(v))
}

//SetClOrdLinkID sets ClOrdLinkID, Tag 583
func (m NewOrderSingle) SetClOrdLinkID(v string) {
	m.Set(field.NewClOrdLinkID(v))
}

//SetDayBookingInst sets DayBookingInst, Tag 589
func (m NewOrderSingle) SetDayBookingInst(v string) {
	m.Set(field.NewDayBookingInst(v))
}

//SetBookingUnit sets BookingUnit, Tag 590
func (m NewOrderSingle) SetBookingUnit(v string) {
	m.Set(field.NewBookingUnit(v))
}

//SetPreallocMethod sets PreallocMethod, Tag 591
func (m NewOrderSingle) SetPreallocMethod(v string) {
	m.Set(field.NewPreallocMethod(v))
}

//SetClearingFeeIndicator sets ClearingFeeIndicator, Tag 635
func (m NewOrderSingle) SetClearingFeeIndicator(v string) {
	m.Set(field.NewClearingFeeIndicator(v))
}

//SetPrice2 sets Price2, Tag 640
func (m NewOrderSingle) SetPrice2(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice2(value, scale))
}

//SetAcctIDSource sets AcctIDSource, Tag 660
func (m NewOrderSingle) SetAcctIDSource(v int) {
	m.Set(field.NewAcctIDSource(v))
}

//SetBenchmarkPrice sets BenchmarkPrice, Tag 662
func (m NewOrderSingle) SetBenchmarkPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewBenchmarkPrice(value, scale))
}

//SetBenchmarkPriceType sets BenchmarkPriceType, Tag 663
func (m NewOrderSingle) SetBenchmarkPriceType(v int) {
	m.Set(field.NewBenchmarkPriceType(v))
}

//SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m NewOrderSingle) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

//SetPool sets Pool, Tag 691
func (m NewOrderSingle) SetPool(v string) {
	m.Set(field.NewPool(v))
}

//SetYieldRedemptionDate sets YieldRedemptionDate, Tag 696
func (m NewOrderSingle) SetYieldRedemptionDate(v string) {
	m.Set(field.NewYieldRedemptionDate(v))
}

//SetYieldRedemptionPrice sets YieldRedemptionPrice, Tag 697
func (m NewOrderSingle) SetYieldRedemptionPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewYieldRedemptionPrice(value, scale))
}

//SetYieldRedemptionPriceType sets YieldRedemptionPriceType, Tag 698
func (m NewOrderSingle) SetYieldRedemptionPriceType(v int) {
	m.Set(field.NewYieldRedemptionPriceType(v))
}

//SetBenchmarkSecurityID sets BenchmarkSecurityID, Tag 699
func (m NewOrderSingle) SetBenchmarkSecurityID(v string) {
	m.Set(field.NewBenchmarkSecurityID(v))
}

//SetYieldCalcDate sets YieldCalcDate, Tag 701
func (m NewOrderSingle) SetYieldCalcDate(v string) {
	m.Set(field.NewYieldCalcDate(v))
}

//SetNoUnderlyings sets NoUnderlyings, Tag 711
func (m NewOrderSingle) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

//SetBenchmarkSecurityIDSource sets BenchmarkSecurityIDSource, Tag 761
func (m NewOrderSingle) SetBenchmarkSecurityIDSource(v string) {
	m.Set(field.NewBenchmarkSecurityIDSource(v))
}

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m NewOrderSingle) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

//SetBookingType sets BookingType, Tag 775
func (m NewOrderSingle) SetBookingType(v int) {
	m.Set(field.NewBookingType(v))
}

//SetTerminationType sets TerminationType, Tag 788
func (m NewOrderSingle) SetTerminationType(v int) {
	m.Set(field.NewTerminationType(v))
}

//SetPegMoveType sets PegMoveType, Tag 835
func (m NewOrderSingle) SetPegMoveType(v int) {
	m.Set(field.NewPegMoveType(v))
}

//SetPegOffsetType sets PegOffsetType, Tag 836
func (m NewOrderSingle) SetPegOffsetType(v int) {
	m.Set(field.NewPegOffsetType(v))
}

//SetPegLimitType sets PegLimitType, Tag 837
func (m NewOrderSingle) SetPegLimitType(v int) {
	m.Set(field.NewPegLimitType(v))
}

//SetPegRoundDirection sets PegRoundDirection, Tag 838
func (m NewOrderSingle) SetPegRoundDirection(v int) {
	m.Set(field.NewPegRoundDirection(v))
}

//SetPegScope sets PegScope, Tag 840
func (m NewOrderSingle) SetPegScope(v int) {
	m.Set(field.NewPegScope(v))
}

//SetDiscretionMoveType sets DiscretionMoveType, Tag 841
func (m NewOrderSingle) SetDiscretionMoveType(v int) {
	m.Set(field.NewDiscretionMoveType(v))
}

//SetDiscretionOffsetType sets DiscretionOffsetType, Tag 842
func (m NewOrderSingle) SetDiscretionOffsetType(v int) {
	m.Set(field.NewDiscretionOffsetType(v))
}

//SetDiscretionLimitType sets DiscretionLimitType, Tag 843
func (m NewOrderSingle) SetDiscretionLimitType(v int) {
	m.Set(field.NewDiscretionLimitType(v))
}

//SetDiscretionRoundDirection sets DiscretionRoundDirection, Tag 844
func (m NewOrderSingle) SetDiscretionRoundDirection(v int) {
	m.Set(field.NewDiscretionRoundDirection(v))
}

//SetDiscretionScope sets DiscretionScope, Tag 846
func (m NewOrderSingle) SetDiscretionScope(v int) {
	m.Set(field.NewDiscretionScope(v))
}

//SetTargetStrategy sets TargetStrategy, Tag 847
func (m NewOrderSingle) SetTargetStrategy(v int) {
	m.Set(field.NewTargetStrategy(v))
}

//SetTargetStrategyParameters sets TargetStrategyParameters, Tag 848
func (m NewOrderSingle) SetTargetStrategyParameters(v string) {
	m.Set(field.NewTargetStrategyParameters(v))
}

//SetParticipationRate sets ParticipationRate, Tag 849
func (m NewOrderSingle) SetParticipationRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewParticipationRate(value, scale))
}

//SetQtyType sets QtyType, Tag 854
func (m NewOrderSingle) SetQtyType(v int) {
	m.Set(field.NewQtyType(v))
}

//SetNoEvents sets NoEvents, Tag 864
func (m NewOrderSingle) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDatedDate sets DatedDate, Tag 873
func (m NewOrderSingle) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

//SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m NewOrderSingle) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

//SetCPProgram sets CPProgram, Tag 875
func (m NewOrderSingle) SetCPProgram(v int) {
	m.Set(field.NewCPProgram(v))
}

//SetCPRegType sets CPRegType, Tag 876
func (m NewOrderSingle) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

//SetMarginRatio sets MarginRatio, Tag 898
func (m NewOrderSingle) SetMarginRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewMarginRatio(value, scale))
}

//SetAgreementDesc sets AgreementDesc, Tag 913
func (m NewOrderSingle) SetAgreementDesc(v string) {
	m.Set(field.NewAgreementDesc(v))
}

//SetAgreementID sets AgreementID, Tag 914
func (m NewOrderSingle) SetAgreementID(v string) {
	m.Set(field.NewAgreementID(v))
}

//SetAgreementDate sets AgreementDate, Tag 915
func (m NewOrderSingle) SetAgreementDate(v string) {
	m.Set(field.NewAgreementDate(v))
}

//SetStartDate sets StartDate, Tag 916
func (m NewOrderSingle) SetStartDate(v string) {
	m.Set(field.NewStartDate(v))
}

//SetEndDate sets EndDate, Tag 917
func (m NewOrderSingle) SetEndDate(v string) {
	m.Set(field.NewEndDate(v))
}

//SetAgreementCurrency sets AgreementCurrency, Tag 918
func (m NewOrderSingle) SetAgreementCurrency(v string) {
	m.Set(field.NewAgreementCurrency(v))
}

//SetDeliveryType sets DeliveryType, Tag 919
func (m NewOrderSingle) SetDeliveryType(v int) {
	m.Set(field.NewDeliveryType(v))
}

//SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m NewOrderSingle) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

//GetAccount gets Account, Tag 1
func (m NewOrderSingle) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m NewOrderSingle) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommission gets Commission, Tag 12
func (m NewOrderSingle) GetCommission() (f field.CommissionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommType gets CommType, Tag 13
func (m NewOrderSingle) GetCommType() (f field.CommTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m NewOrderSingle) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecInst gets ExecInst, Tag 18
func (m NewOrderSingle) GetExecInst() (f field.ExecInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetHandlInst gets HandlInst, Tag 21
func (m NewOrderSingle) GetHandlInst() (f field.HandlInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m NewOrderSingle) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIID gets IOIID, Tag 23
func (m NewOrderSingle) GetIOIID() (f field.IOIIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m NewOrderSingle) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdType gets OrdType, Tag 40
func (m NewOrderSingle) GetOrdType() (f field.OrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m NewOrderSingle) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m NewOrderSingle) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m NewOrderSingle) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m NewOrderSingle) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m NewOrderSingle) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m NewOrderSingle) GetTimeInForce() (f field.TimeInForceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m NewOrderSingle) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlType gets SettlType, Tag 63
func (m NewOrderSingle) GetSettlType() (f field.SettlTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlDate gets SettlDate, Tag 64
func (m NewOrderSingle) GetSettlDate() (f field.SettlDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m NewOrderSingle) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocID gets AllocID, Tag 70
func (m NewOrderSingle) GetAllocID() (f field.AllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m NewOrderSingle) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionEffect gets PositionEffect, Tag 77
func (m NewOrderSingle) GetPositionEffect() (f field.PositionEffectField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoAllocs gets NoAllocs, Tag 78
func (m NewOrderSingle) GetNoAllocs() (NoAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProcessCode gets ProcessCode, Tag 81
func (m NewOrderSingle) GetProcessCode() (f field.ProcessCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStopPx gets StopPx, Tag 99
func (m NewOrderSingle) GetStopPx() (f field.StopPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExDestination gets ExDestination, Tag 100
func (m NewOrderSingle) GetExDestination() (f field.ExDestinationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m NewOrderSingle) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m NewOrderSingle) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinQty gets MinQty, Tag 110
func (m NewOrderSingle) GetMinQty() (f field.MinQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxFloor gets MaxFloor, Tag 111
func (m NewOrderSingle) GetMaxFloor() (f field.MaxFloorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocateReqd gets LocateReqd, Tag 114
func (m NewOrderSingle) GetLocateReqd() (f field.LocateReqdField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteID gets QuoteID, Tag 117
func (m NewOrderSingle) GetQuoteID() (f field.QuoteIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m NewOrderSingle) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetForexReq gets ForexReq, Tag 121
func (m NewOrderSingle) GetForexReq() (f field.ForexReqField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m NewOrderSingle) GetExpireTime() (f field.ExpireTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrevClosePx gets PrevClosePx, Tag 140
func (m NewOrderSingle) GetPrevClosePx() (f field.PrevClosePxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m NewOrderSingle) GetCashOrderQty() (f field.CashOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m NewOrderSingle) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEffectiveTime gets EffectiveTime, Tag 168
func (m NewOrderSingle) GetEffectiveTime() (f field.EffectiveTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty2 gets OrderQty2, Tag 192
func (m NewOrderSingle) GetOrderQty2() (f field.OrderQty2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlDate2 gets SettlDate2, Tag 193
func (m NewOrderSingle) GetSettlDate2() (f field.SettlDate2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m NewOrderSingle) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m NewOrderSingle) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCoveredOrUncovered gets CoveredOrUncovered, Tag 203
func (m NewOrderSingle) GetCoveredOrUncovered() (f field.CoveredOrUncoveredField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m NewOrderSingle) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m NewOrderSingle) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxShow gets MaxShow, Tag 210
func (m NewOrderSingle) GetMaxShow() (f field.MaxShowField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegOffsetValue gets PegOffsetValue, Tag 211
func (m NewOrderSingle) GetPegOffsetValue() (f field.PegOffsetValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSpread gets Spread, Tag 218
func (m NewOrderSingle) GetSpread() (f field.SpreadField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveCurrency gets BenchmarkCurveCurrency, Tag 220
func (m NewOrderSingle) GetBenchmarkCurveCurrency() (f field.BenchmarkCurveCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveName gets BenchmarkCurveName, Tag 221
func (m NewOrderSingle) GetBenchmarkCurveName() (f field.BenchmarkCurveNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurvePoint gets BenchmarkCurvePoint, Tag 222
func (m NewOrderSingle) GetBenchmarkCurvePoint() (f field.BenchmarkCurvePointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m NewOrderSingle) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m NewOrderSingle) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m NewOrderSingle) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m NewOrderSingle) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m NewOrderSingle) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m NewOrderSingle) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeOriginationDate gets TradeOriginationDate, Tag 229
func (m NewOrderSingle) GetTradeOriginationDate() (f field.TradeOriginationDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m NewOrderSingle) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoStipulations gets NoStipulations, Tag 232
func (m NewOrderSingle) GetNoStipulations() (NoStipulationsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStipulationsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetYieldType gets YieldType, Tag 235
func (m NewOrderSingle) GetYieldType() (f field.YieldTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYield gets Yield, Tag 236
func (m NewOrderSingle) GetYield() (f field.YieldField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m NewOrderSingle) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m NewOrderSingle) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m NewOrderSingle) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m NewOrderSingle) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m NewOrderSingle) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m NewOrderSingle) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m NewOrderSingle) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m NewOrderSingle) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m NewOrderSingle) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplianceID gets ComplianceID, Tag 376
func (m NewOrderSingle) GetComplianceID() (f field.ComplianceIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSolicitedFlag gets SolicitedFlag, Tag 377
func (m NewOrderSingle) GetSolicitedFlag() (f field.SolicitedFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoTradingSessions gets NoTradingSessions, Tag 386
func (m NewOrderSingle) GetNoTradingSessions() (NoTradingSessionsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTradingSessionsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDiscretionInst gets DiscretionInst, Tag 388
func (m NewOrderSingle) GetDiscretionInst() (f field.DiscretionInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionOffsetValue gets DiscretionOffsetValue, Tag 389
func (m NewOrderSingle) GetDiscretionOffsetValue() (f field.DiscretionOffsetValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceType gets PriceType, Tag 423
func (m NewOrderSingle) GetPriceType() (f field.PriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetGTBookingInst gets GTBookingInst, Tag 427
func (m NewOrderSingle) GetGTBookingInst() (f field.GTBookingInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireDate gets ExpireDate, Tag 432
func (m NewOrderSingle) GetExpireDate() (f field.ExpireDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m NewOrderSingle) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m NewOrderSingle) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m NewOrderSingle) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m NewOrderSingle) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingDirection gets RoundingDirection, Tag 468
func (m NewOrderSingle) GetRoundingDirection() (f field.RoundingDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingModulus gets RoundingModulus, Tag 469
func (m NewOrderSingle) GetRoundingModulus() (f field.RoundingModulusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m NewOrderSingle) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m NewOrderSingle) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m NewOrderSingle) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommCurrency gets CommCurrency, Tag 479
func (m NewOrderSingle) GetCommCurrency() (f field.CommCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCancellationRights gets CancellationRights, Tag 480
func (m NewOrderSingle) GetCancellationRights() (f field.CancellationRightsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMoneyLaunderingStatus gets MoneyLaunderingStatus, Tag 481
func (m NewOrderSingle) GetMoneyLaunderingStatus() (f field.MoneyLaunderingStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDesignation gets Designation, Tag 494
func (m NewOrderSingle) GetDesignation() (f field.DesignationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFundRenewWaiv gets FundRenewWaiv, Tag 497
func (m NewOrderSingle) GetFundRenewWaiv() (f field.FundRenewWaivField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRegistID gets RegistID, Tag 513
func (m NewOrderSingle) GetRegistID() (f field.RegistIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderPercent gets OrderPercent, Tag 516
func (m NewOrderSingle) GetOrderPercent() (f field.OrderPercentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526
func (m NewOrderSingle) GetSecondaryClOrdID() (f field.SecondaryClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderCapacity gets OrderCapacity, Tag 528
func (m NewOrderSingle) GetOrderCapacity() (f field.OrderCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderRestrictions gets OrderRestrictions, Tag 529
func (m NewOrderSingle) GetOrderRestrictions() (f field.OrderRestrictionsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m NewOrderSingle) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m NewOrderSingle) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashMargin gets CashMargin, Tag 544
func (m NewOrderSingle) GetCashMargin() (f field.CashMarginField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccountType gets AccountType, Tag 581
func (m NewOrderSingle) GetAccountType() (f field.AccountTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCustOrderCapacity gets CustOrderCapacity, Tag 582
func (m NewOrderSingle) GetCustOrderCapacity() (f field.CustOrderCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdLinkID gets ClOrdLinkID, Tag 583
func (m NewOrderSingle) GetClOrdLinkID() (f field.ClOrdLinkIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDayBookingInst gets DayBookingInst, Tag 589
func (m NewOrderSingle) GetDayBookingInst() (f field.DayBookingInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBookingUnit gets BookingUnit, Tag 590
func (m NewOrderSingle) GetBookingUnit() (f field.BookingUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPreallocMethod gets PreallocMethod, Tag 591
func (m NewOrderSingle) GetPreallocMethod() (f field.PreallocMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClearingFeeIndicator gets ClearingFeeIndicator, Tag 635
func (m NewOrderSingle) GetClearingFeeIndicator() (f field.ClearingFeeIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice2 gets Price2, Tag 640
func (m NewOrderSingle) GetPrice2() (f field.Price2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAcctIDSource gets AcctIDSource, Tag 660
func (m NewOrderSingle) GetAcctIDSource() (f field.AcctIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkPrice gets BenchmarkPrice, Tag 662
func (m NewOrderSingle) GetBenchmarkPrice() (f field.BenchmarkPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkPriceType gets BenchmarkPriceType, Tag 663
func (m NewOrderSingle) GetBenchmarkPriceType() (f field.BenchmarkPriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m NewOrderSingle) GetContractSettlMonth() (f field.ContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPool gets Pool, Tag 691
func (m NewOrderSingle) GetPool() (f field.PoolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldRedemptionDate gets YieldRedemptionDate, Tag 696
func (m NewOrderSingle) GetYieldRedemptionDate() (f field.YieldRedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldRedemptionPrice gets YieldRedemptionPrice, Tag 697
func (m NewOrderSingle) GetYieldRedemptionPrice() (f field.YieldRedemptionPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldRedemptionPriceType gets YieldRedemptionPriceType, Tag 698
func (m NewOrderSingle) GetYieldRedemptionPriceType() (f field.YieldRedemptionPriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkSecurityID gets BenchmarkSecurityID, Tag 699
func (m NewOrderSingle) GetBenchmarkSecurityID() (f field.BenchmarkSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldCalcDate gets YieldCalcDate, Tag 701
func (m NewOrderSingle) GetYieldCalcDate() (f field.YieldCalcDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyings gets NoUnderlyings, Tag 711
func (m NewOrderSingle) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetBenchmarkSecurityIDSource gets BenchmarkSecurityIDSource, Tag 761
func (m NewOrderSingle) GetBenchmarkSecurityIDSource() (f field.BenchmarkSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m NewOrderSingle) GetSecuritySubType() (f field.SecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBookingType gets BookingType, Tag 775
func (m NewOrderSingle) GetBookingType() (f field.BookingTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTerminationType gets TerminationType, Tag 788
func (m NewOrderSingle) GetTerminationType() (f field.TerminationTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegMoveType gets PegMoveType, Tag 835
func (m NewOrderSingle) GetPegMoveType() (f field.PegMoveTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegOffsetType gets PegOffsetType, Tag 836
func (m NewOrderSingle) GetPegOffsetType() (f field.PegOffsetTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegLimitType gets PegLimitType, Tag 837
func (m NewOrderSingle) GetPegLimitType() (f field.PegLimitTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegRoundDirection gets PegRoundDirection, Tag 838
func (m NewOrderSingle) GetPegRoundDirection() (f field.PegRoundDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegScope gets PegScope, Tag 840
func (m NewOrderSingle) GetPegScope() (f field.PegScopeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionMoveType gets DiscretionMoveType, Tag 841
func (m NewOrderSingle) GetDiscretionMoveType() (f field.DiscretionMoveTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionOffsetType gets DiscretionOffsetType, Tag 842
func (m NewOrderSingle) GetDiscretionOffsetType() (f field.DiscretionOffsetTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionLimitType gets DiscretionLimitType, Tag 843
func (m NewOrderSingle) GetDiscretionLimitType() (f field.DiscretionLimitTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionRoundDirection gets DiscretionRoundDirection, Tag 844
func (m NewOrderSingle) GetDiscretionRoundDirection() (f field.DiscretionRoundDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionScope gets DiscretionScope, Tag 846
func (m NewOrderSingle) GetDiscretionScope() (f field.DiscretionScopeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTargetStrategy gets TargetStrategy, Tag 847
func (m NewOrderSingle) GetTargetStrategy() (f field.TargetStrategyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTargetStrategyParameters gets TargetStrategyParameters, Tag 848
func (m NewOrderSingle) GetTargetStrategyParameters() (f field.TargetStrategyParametersField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetParticipationRate gets ParticipationRate, Tag 849
func (m NewOrderSingle) GetParticipationRate() (f field.ParticipationRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQtyType gets QtyType, Tag 854
func (m NewOrderSingle) GetQtyType() (f field.QtyTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoEvents gets NoEvents, Tag 864
func (m NewOrderSingle) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDatedDate gets DatedDate, Tag 873
func (m NewOrderSingle) GetDatedDate() (f field.DatedDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m NewOrderSingle) GetInterestAccrualDate() (f field.InterestAccrualDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPProgram gets CPProgram, Tag 875
func (m NewOrderSingle) GetCPProgram() (f field.CPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPRegType gets CPRegType, Tag 876
func (m NewOrderSingle) GetCPRegType() (f field.CPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMarginRatio gets MarginRatio, Tag 898
func (m NewOrderSingle) GetMarginRatio() (f field.MarginRatioField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementDesc gets AgreementDesc, Tag 913
func (m NewOrderSingle) GetAgreementDesc() (f field.AgreementDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementID gets AgreementID, Tag 914
func (m NewOrderSingle) GetAgreementID() (f field.AgreementIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementDate gets AgreementDate, Tag 915
func (m NewOrderSingle) GetAgreementDate() (f field.AgreementDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStartDate gets StartDate, Tag 916
func (m NewOrderSingle) GetStartDate() (f field.StartDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEndDate gets EndDate, Tag 917
func (m NewOrderSingle) GetEndDate() (f field.EndDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementCurrency gets AgreementCurrency, Tag 918
func (m NewOrderSingle) GetAgreementCurrency() (f field.AgreementCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDeliveryType gets DeliveryType, Tag 919
func (m NewOrderSingle) GetDeliveryType() (f field.DeliveryTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m NewOrderSingle) GetStrikeCurrency() (f field.StrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAccount returns true if Account is present, Tag 1
func (m NewOrderSingle) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m NewOrderSingle) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasCommission returns true if Commission is present, Tag 12
func (m NewOrderSingle) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m NewOrderSingle) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m NewOrderSingle) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasExecInst returns true if ExecInst is present, Tag 18
func (m NewOrderSingle) HasExecInst() bool {
	return m.Has(tag.ExecInst)
}

//HasHandlInst returns true if HandlInst is present, Tag 21
func (m NewOrderSingle) HasHandlInst() bool {
	return m.Has(tag.HandlInst)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m NewOrderSingle) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasIOIID returns true if IOIID is present, Tag 23
func (m NewOrderSingle) HasIOIID() bool {
	return m.Has(tag.IOIID)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m NewOrderSingle) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m NewOrderSingle) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasPrice returns true if Price is present, Tag 44
func (m NewOrderSingle) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m NewOrderSingle) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m NewOrderSingle) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m NewOrderSingle) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m NewOrderSingle) HasText() bool {
	return m.Has(tag.Text)
}

//HasTimeInForce returns true if TimeInForce is present, Tag 59
func (m NewOrderSingle) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m NewOrderSingle) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSettlType returns true if SettlType is present, Tag 63
func (m NewOrderSingle) HasSettlType() bool {
	return m.Has(tag.SettlType)
}

//HasSettlDate returns true if SettlDate is present, Tag 64
func (m NewOrderSingle) HasSettlDate() bool {
	return m.Has(tag.SettlDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m NewOrderSingle) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasAllocID returns true if AllocID is present, Tag 70
func (m NewOrderSingle) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m NewOrderSingle) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasPositionEffect returns true if PositionEffect is present, Tag 77
func (m NewOrderSingle) HasPositionEffect() bool {
	return m.Has(tag.PositionEffect)
}

//HasNoAllocs returns true if NoAllocs is present, Tag 78
func (m NewOrderSingle) HasNoAllocs() bool {
	return m.Has(tag.NoAllocs)
}

//HasProcessCode returns true if ProcessCode is present, Tag 81
func (m NewOrderSingle) HasProcessCode() bool {
	return m.Has(tag.ProcessCode)
}

//HasStopPx returns true if StopPx is present, Tag 99
func (m NewOrderSingle) HasStopPx() bool {
	return m.Has(tag.StopPx)
}

//HasExDestination returns true if ExDestination is present, Tag 100
func (m NewOrderSingle) HasExDestination() bool {
	return m.Has(tag.ExDestination)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m NewOrderSingle) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m NewOrderSingle) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasMinQty returns true if MinQty is present, Tag 110
func (m NewOrderSingle) HasMinQty() bool {
	return m.Has(tag.MinQty)
}

//HasMaxFloor returns true if MaxFloor is present, Tag 111
func (m NewOrderSingle) HasMaxFloor() bool {
	return m.Has(tag.MaxFloor)
}

//HasLocateReqd returns true if LocateReqd is present, Tag 114
func (m NewOrderSingle) HasLocateReqd() bool {
	return m.Has(tag.LocateReqd)
}

//HasQuoteID returns true if QuoteID is present, Tag 117
func (m NewOrderSingle) HasQuoteID() bool {
	return m.Has(tag.QuoteID)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m NewOrderSingle) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasForexReq returns true if ForexReq is present, Tag 121
func (m NewOrderSingle) HasForexReq() bool {
	return m.Has(tag.ForexReq)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m NewOrderSingle) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasPrevClosePx returns true if PrevClosePx is present, Tag 140
func (m NewOrderSingle) HasPrevClosePx() bool {
	return m.Has(tag.PrevClosePx)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m NewOrderSingle) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m NewOrderSingle) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasEffectiveTime returns true if EffectiveTime is present, Tag 168
func (m NewOrderSingle) HasEffectiveTime() bool {
	return m.Has(tag.EffectiveTime)
}

//HasOrderQty2 returns true if OrderQty2 is present, Tag 192
func (m NewOrderSingle) HasOrderQty2() bool {
	return m.Has(tag.OrderQty2)
}

//HasSettlDate2 returns true if SettlDate2 is present, Tag 193
func (m NewOrderSingle) HasSettlDate2() bool {
	return m.Has(tag.SettlDate2)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m NewOrderSingle) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m NewOrderSingle) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasCoveredOrUncovered returns true if CoveredOrUncovered is present, Tag 203
func (m NewOrderSingle) HasCoveredOrUncovered() bool {
	return m.Has(tag.CoveredOrUncovered)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m NewOrderSingle) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m NewOrderSingle) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasMaxShow returns true if MaxShow is present, Tag 210
func (m NewOrderSingle) HasMaxShow() bool {
	return m.Has(tag.MaxShow)
}

//HasPegOffsetValue returns true if PegOffsetValue is present, Tag 211
func (m NewOrderSingle) HasPegOffsetValue() bool {
	return m.Has(tag.PegOffsetValue)
}

//HasSpread returns true if Spread is present, Tag 218
func (m NewOrderSingle) HasSpread() bool {
	return m.Has(tag.Spread)
}

//HasBenchmarkCurveCurrency returns true if BenchmarkCurveCurrency is present, Tag 220
func (m NewOrderSingle) HasBenchmarkCurveCurrency() bool {
	return m.Has(tag.BenchmarkCurveCurrency)
}

//HasBenchmarkCurveName returns true if BenchmarkCurveName is present, Tag 221
func (m NewOrderSingle) HasBenchmarkCurveName() bool {
	return m.Has(tag.BenchmarkCurveName)
}

//HasBenchmarkCurvePoint returns true if BenchmarkCurvePoint is present, Tag 222
func (m NewOrderSingle) HasBenchmarkCurvePoint() bool {
	return m.Has(tag.BenchmarkCurvePoint)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m NewOrderSingle) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m NewOrderSingle) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m NewOrderSingle) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m NewOrderSingle) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m NewOrderSingle) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m NewOrderSingle) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasTradeOriginationDate returns true if TradeOriginationDate is present, Tag 229
func (m NewOrderSingle) HasTradeOriginationDate() bool {
	return m.Has(tag.TradeOriginationDate)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m NewOrderSingle) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasNoStipulations returns true if NoStipulations is present, Tag 232
func (m NewOrderSingle) HasNoStipulations() bool {
	return m.Has(tag.NoStipulations)
}

//HasYieldType returns true if YieldType is present, Tag 235
func (m NewOrderSingle) HasYieldType() bool {
	return m.Has(tag.YieldType)
}

//HasYield returns true if Yield is present, Tag 236
func (m NewOrderSingle) HasYield() bool {
	return m.Has(tag.Yield)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m NewOrderSingle) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m NewOrderSingle) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m NewOrderSingle) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m NewOrderSingle) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m NewOrderSingle) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m NewOrderSingle) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m NewOrderSingle) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m NewOrderSingle) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m NewOrderSingle) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasComplianceID returns true if ComplianceID is present, Tag 376
func (m NewOrderSingle) HasComplianceID() bool {
	return m.Has(tag.ComplianceID)
}

//HasSolicitedFlag returns true if SolicitedFlag is present, Tag 377
func (m NewOrderSingle) HasSolicitedFlag() bool {
	return m.Has(tag.SolicitedFlag)
}

//HasNoTradingSessions returns true if NoTradingSessions is present, Tag 386
func (m NewOrderSingle) HasNoTradingSessions() bool {
	return m.Has(tag.NoTradingSessions)
}

//HasDiscretionInst returns true if DiscretionInst is present, Tag 388
func (m NewOrderSingle) HasDiscretionInst() bool {
	return m.Has(tag.DiscretionInst)
}

//HasDiscretionOffsetValue returns true if DiscretionOffsetValue is present, Tag 389
func (m NewOrderSingle) HasDiscretionOffsetValue() bool {
	return m.Has(tag.DiscretionOffsetValue)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m NewOrderSingle) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasGTBookingInst returns true if GTBookingInst is present, Tag 427
func (m NewOrderSingle) HasGTBookingInst() bool {
	return m.Has(tag.GTBookingInst)
}

//HasExpireDate returns true if ExpireDate is present, Tag 432
func (m NewOrderSingle) HasExpireDate() bool {
	return m.Has(tag.ExpireDate)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m NewOrderSingle) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m NewOrderSingle) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m NewOrderSingle) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m NewOrderSingle) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasRoundingDirection returns true if RoundingDirection is present, Tag 468
func (m NewOrderSingle) HasRoundingDirection() bool {
	return m.Has(tag.RoundingDirection)
}

//HasRoundingModulus returns true if RoundingModulus is present, Tag 469
func (m NewOrderSingle) HasRoundingModulus() bool {
	return m.Has(tag.RoundingModulus)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m NewOrderSingle) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m NewOrderSingle) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m NewOrderSingle) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasCommCurrency returns true if CommCurrency is present, Tag 479
func (m NewOrderSingle) HasCommCurrency() bool {
	return m.Has(tag.CommCurrency)
}

//HasCancellationRights returns true if CancellationRights is present, Tag 480
func (m NewOrderSingle) HasCancellationRights() bool {
	return m.Has(tag.CancellationRights)
}

//HasMoneyLaunderingStatus returns true if MoneyLaunderingStatus is present, Tag 481
func (m NewOrderSingle) HasMoneyLaunderingStatus() bool {
	return m.Has(tag.MoneyLaunderingStatus)
}

//HasDesignation returns true if Designation is present, Tag 494
func (m NewOrderSingle) HasDesignation() bool {
	return m.Has(tag.Designation)
}

//HasFundRenewWaiv returns true if FundRenewWaiv is present, Tag 497
func (m NewOrderSingle) HasFundRenewWaiv() bool {
	return m.Has(tag.FundRenewWaiv)
}

//HasRegistID returns true if RegistID is present, Tag 513
func (m NewOrderSingle) HasRegistID() bool {
	return m.Has(tag.RegistID)
}

//HasOrderPercent returns true if OrderPercent is present, Tag 516
func (m NewOrderSingle) HasOrderPercent() bool {
	return m.Has(tag.OrderPercent)
}

//HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526
func (m NewOrderSingle) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

//HasOrderCapacity returns true if OrderCapacity is present, Tag 528
func (m NewOrderSingle) HasOrderCapacity() bool {
	return m.Has(tag.OrderCapacity)
}

//HasOrderRestrictions returns true if OrderRestrictions is present, Tag 529
func (m NewOrderSingle) HasOrderRestrictions() bool {
	return m.Has(tag.OrderRestrictions)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m NewOrderSingle) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m NewOrderSingle) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasCashMargin returns true if CashMargin is present, Tag 544
func (m NewOrderSingle) HasCashMargin() bool {
	return m.Has(tag.CashMargin)
}

//HasAccountType returns true if AccountType is present, Tag 581
func (m NewOrderSingle) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

//HasCustOrderCapacity returns true if CustOrderCapacity is present, Tag 582
func (m NewOrderSingle) HasCustOrderCapacity() bool {
	return m.Has(tag.CustOrderCapacity)
}

//HasClOrdLinkID returns true if ClOrdLinkID is present, Tag 583
func (m NewOrderSingle) HasClOrdLinkID() bool {
	return m.Has(tag.ClOrdLinkID)
}

//HasDayBookingInst returns true if DayBookingInst is present, Tag 589
func (m NewOrderSingle) HasDayBookingInst() bool {
	return m.Has(tag.DayBookingInst)
}

//HasBookingUnit returns true if BookingUnit is present, Tag 590
func (m NewOrderSingle) HasBookingUnit() bool {
	return m.Has(tag.BookingUnit)
}

//HasPreallocMethod returns true if PreallocMethod is present, Tag 591
func (m NewOrderSingle) HasPreallocMethod() bool {
	return m.Has(tag.PreallocMethod)
}

//HasClearingFeeIndicator returns true if ClearingFeeIndicator is present, Tag 635
func (m NewOrderSingle) HasClearingFeeIndicator() bool {
	return m.Has(tag.ClearingFeeIndicator)
}

//HasPrice2 returns true if Price2 is present, Tag 640
func (m NewOrderSingle) HasPrice2() bool {
	return m.Has(tag.Price2)
}

//HasAcctIDSource returns true if AcctIDSource is present, Tag 660
func (m NewOrderSingle) HasAcctIDSource() bool {
	return m.Has(tag.AcctIDSource)
}

//HasBenchmarkPrice returns true if BenchmarkPrice is present, Tag 662
func (m NewOrderSingle) HasBenchmarkPrice() bool {
	return m.Has(tag.BenchmarkPrice)
}

//HasBenchmarkPriceType returns true if BenchmarkPriceType is present, Tag 663
func (m NewOrderSingle) HasBenchmarkPriceType() bool {
	return m.Has(tag.BenchmarkPriceType)
}

//HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m NewOrderSingle) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

//HasPool returns true if Pool is present, Tag 691
func (m NewOrderSingle) HasPool() bool {
	return m.Has(tag.Pool)
}

//HasYieldRedemptionDate returns true if YieldRedemptionDate is present, Tag 696
func (m NewOrderSingle) HasYieldRedemptionDate() bool {
	return m.Has(tag.YieldRedemptionDate)
}

//HasYieldRedemptionPrice returns true if YieldRedemptionPrice is present, Tag 697
func (m NewOrderSingle) HasYieldRedemptionPrice() bool {
	return m.Has(tag.YieldRedemptionPrice)
}

//HasYieldRedemptionPriceType returns true if YieldRedemptionPriceType is present, Tag 698
func (m NewOrderSingle) HasYieldRedemptionPriceType() bool {
	return m.Has(tag.YieldRedemptionPriceType)
}

//HasBenchmarkSecurityID returns true if BenchmarkSecurityID is present, Tag 699
func (m NewOrderSingle) HasBenchmarkSecurityID() bool {
	return m.Has(tag.BenchmarkSecurityID)
}

//HasYieldCalcDate returns true if YieldCalcDate is present, Tag 701
func (m NewOrderSingle) HasYieldCalcDate() bool {
	return m.Has(tag.YieldCalcDate)
}

//HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711
func (m NewOrderSingle) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

//HasBenchmarkSecurityIDSource returns true if BenchmarkSecurityIDSource is present, Tag 761
func (m NewOrderSingle) HasBenchmarkSecurityIDSource() bool {
	return m.Has(tag.BenchmarkSecurityIDSource)
}

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m NewOrderSingle) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

//HasBookingType returns true if BookingType is present, Tag 775
func (m NewOrderSingle) HasBookingType() bool {
	return m.Has(tag.BookingType)
}

//HasTerminationType returns true if TerminationType is present, Tag 788
func (m NewOrderSingle) HasTerminationType() bool {
	return m.Has(tag.TerminationType)
}

//HasPegMoveType returns true if PegMoveType is present, Tag 835
func (m NewOrderSingle) HasPegMoveType() bool {
	return m.Has(tag.PegMoveType)
}

//HasPegOffsetType returns true if PegOffsetType is present, Tag 836
func (m NewOrderSingle) HasPegOffsetType() bool {
	return m.Has(tag.PegOffsetType)
}

//HasPegLimitType returns true if PegLimitType is present, Tag 837
func (m NewOrderSingle) HasPegLimitType() bool {
	return m.Has(tag.PegLimitType)
}

//HasPegRoundDirection returns true if PegRoundDirection is present, Tag 838
func (m NewOrderSingle) HasPegRoundDirection() bool {
	return m.Has(tag.PegRoundDirection)
}

//HasPegScope returns true if PegScope is present, Tag 840
func (m NewOrderSingle) HasPegScope() bool {
	return m.Has(tag.PegScope)
}

//HasDiscretionMoveType returns true if DiscretionMoveType is present, Tag 841
func (m NewOrderSingle) HasDiscretionMoveType() bool {
	return m.Has(tag.DiscretionMoveType)
}

//HasDiscretionOffsetType returns true if DiscretionOffsetType is present, Tag 842
func (m NewOrderSingle) HasDiscretionOffsetType() bool {
	return m.Has(tag.DiscretionOffsetType)
}

//HasDiscretionLimitType returns true if DiscretionLimitType is present, Tag 843
func (m NewOrderSingle) HasDiscretionLimitType() bool {
	return m.Has(tag.DiscretionLimitType)
}

//HasDiscretionRoundDirection returns true if DiscretionRoundDirection is present, Tag 844
func (m NewOrderSingle) HasDiscretionRoundDirection() bool {
	return m.Has(tag.DiscretionRoundDirection)
}

//HasDiscretionScope returns true if DiscretionScope is present, Tag 846
func (m NewOrderSingle) HasDiscretionScope() bool {
	return m.Has(tag.DiscretionScope)
}

//HasTargetStrategy returns true if TargetStrategy is present, Tag 847
func (m NewOrderSingle) HasTargetStrategy() bool {
	return m.Has(tag.TargetStrategy)
}

//HasTargetStrategyParameters returns true if TargetStrategyParameters is present, Tag 848
func (m NewOrderSingle) HasTargetStrategyParameters() bool {
	return m.Has(tag.TargetStrategyParameters)
}

//HasParticipationRate returns true if ParticipationRate is present, Tag 849
func (m NewOrderSingle) HasParticipationRate() bool {
	return m.Has(tag.ParticipationRate)
}

//HasQtyType returns true if QtyType is present, Tag 854
func (m NewOrderSingle) HasQtyType() bool {
	return m.Has(tag.QtyType)
}

//HasNoEvents returns true if NoEvents is present, Tag 864
func (m NewOrderSingle) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

//HasDatedDate returns true if DatedDate is present, Tag 873
func (m NewOrderSingle) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

//HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m NewOrderSingle) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

//HasCPProgram returns true if CPProgram is present, Tag 875
func (m NewOrderSingle) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

//HasCPRegType returns true if CPRegType is present, Tag 876
func (m NewOrderSingle) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

//HasMarginRatio returns true if MarginRatio is present, Tag 898
func (m NewOrderSingle) HasMarginRatio() bool {
	return m.Has(tag.MarginRatio)
}

//HasAgreementDesc returns true if AgreementDesc is present, Tag 913
func (m NewOrderSingle) HasAgreementDesc() bool {
	return m.Has(tag.AgreementDesc)
}

//HasAgreementID returns true if AgreementID is present, Tag 914
func (m NewOrderSingle) HasAgreementID() bool {
	return m.Has(tag.AgreementID)
}

//HasAgreementDate returns true if AgreementDate is present, Tag 915
func (m NewOrderSingle) HasAgreementDate() bool {
	return m.Has(tag.AgreementDate)
}

//HasStartDate returns true if StartDate is present, Tag 916
func (m NewOrderSingle) HasStartDate() bool {
	return m.Has(tag.StartDate)
}

//HasEndDate returns true if EndDate is present, Tag 917
func (m NewOrderSingle) HasEndDate() bool {
	return m.Has(tag.EndDate)
}

//HasAgreementCurrency returns true if AgreementCurrency is present, Tag 918
func (m NewOrderSingle) HasAgreementCurrency() bool {
	return m.Has(tag.AgreementCurrency)
}

//HasDeliveryType returns true if DeliveryType is present, Tag 919
func (m NewOrderSingle) HasDeliveryType() bool {
	return m.Has(tag.DeliveryType)
}

//HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m NewOrderSingle) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

//NoAllocs is a repeating group element, Tag 78
type NoAllocs struct {
	quickfix.Group
}

//SetAllocAccount sets AllocAccount, Tag 79
func (m NoAllocs) SetAllocAccount(v string) {
	m.Set(field.NewAllocAccount(v))
}

//SetAllocAcctIDSource sets AllocAcctIDSource, Tag 661
func (m NoAllocs) SetAllocAcctIDSource(v int) {
	m.Set(field.NewAllocAcctIDSource(v))
}

//SetAllocSettlCurrency sets AllocSettlCurrency, Tag 736
func (m NoAllocs) SetAllocSettlCurrency(v string) {
	m.Set(field.NewAllocSettlCurrency(v))
}

//SetIndividualAllocID sets IndividualAllocID, Tag 467
func (m NoAllocs) SetIndividualAllocID(v string) {
	m.Set(field.NewIndividualAllocID(v))
}

//SetNoNestedPartyIDs sets NoNestedPartyIDs, Tag 539
func (m NoAllocs) SetNoNestedPartyIDs(f NoNestedPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetAllocQty sets AllocQty, Tag 80
func (m NoAllocs) SetAllocQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewAllocQty(value, scale))
}

//GetAllocAccount gets AllocAccount, Tag 79
func (m NoAllocs) GetAllocAccount() (f field.AllocAccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocAcctIDSource gets AllocAcctIDSource, Tag 661
func (m NoAllocs) GetAllocAcctIDSource() (f field.AllocAcctIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocSettlCurrency gets AllocSettlCurrency, Tag 736
func (m NoAllocs) GetAllocSettlCurrency() (f field.AllocSettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIndividualAllocID gets IndividualAllocID, Tag 467
func (m NoAllocs) GetIndividualAllocID() (f field.IndividualAllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoNestedPartyIDs gets NoNestedPartyIDs, Tag 539
func (m NoAllocs) GetNoNestedPartyIDs() (NoNestedPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAllocQty gets AllocQty, Tag 80
func (m NoAllocs) GetAllocQty() (f field.AllocQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAllocAccount returns true if AllocAccount is present, Tag 79
func (m NoAllocs) HasAllocAccount() bool {
	return m.Has(tag.AllocAccount)
}

//HasAllocAcctIDSource returns true if AllocAcctIDSource is present, Tag 661
func (m NoAllocs) HasAllocAcctIDSource() bool {
	return m.Has(tag.AllocAcctIDSource)
}

//HasAllocSettlCurrency returns true if AllocSettlCurrency is present, Tag 736
func (m NoAllocs) HasAllocSettlCurrency() bool {
	return m.Has(tag.AllocSettlCurrency)
}

//HasIndividualAllocID returns true if IndividualAllocID is present, Tag 467
func (m NoAllocs) HasIndividualAllocID() bool {
	return m.Has(tag.IndividualAllocID)
}

//HasNoNestedPartyIDs returns true if NoNestedPartyIDs is present, Tag 539
func (m NoAllocs) HasNoNestedPartyIDs() bool {
	return m.Has(tag.NoNestedPartyIDs)
}

//HasAllocQty returns true if AllocQty is present, Tag 80
func (m NoAllocs) HasAllocQty() bool {
	return m.Has(tag.AllocQty)
}

//NoNestedPartyIDs is a repeating group element, Tag 539
type NoNestedPartyIDs struct {
	quickfix.Group
}

//SetNestedPartyID sets NestedPartyID, Tag 524
func (m NoNestedPartyIDs) SetNestedPartyID(v string) {
	m.Set(field.NewNestedPartyID(v))
}

//SetNestedPartyIDSource sets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) SetNestedPartyIDSource(v string) {
	m.Set(field.NewNestedPartyIDSource(v))
}

//SetNestedPartyRole sets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) SetNestedPartyRole(v int) {
	m.Set(field.NewNestedPartyRole(v))
}

//SetNoNestedPartySubIDs sets NoNestedPartySubIDs, Tag 804
func (m NoNestedPartyIDs) SetNoNestedPartySubIDs(f NoNestedPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetNestedPartyID gets NestedPartyID, Tag 524
func (m NoNestedPartyIDs) GetNestedPartyID() (f field.NestedPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNestedPartyIDSource gets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) GetNestedPartyIDSource() (f field.NestedPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNestedPartyRole gets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) GetNestedPartyRole() (f field.NestedPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoNestedPartySubIDs gets NoNestedPartySubIDs, Tag 804
func (m NoNestedPartyIDs) GetNoNestedPartySubIDs() (NoNestedPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasNestedPartyID returns true if NestedPartyID is present, Tag 524
func (m NoNestedPartyIDs) HasNestedPartyID() bool {
	return m.Has(tag.NestedPartyID)
}

//HasNestedPartyIDSource returns true if NestedPartyIDSource is present, Tag 525
func (m NoNestedPartyIDs) HasNestedPartyIDSource() bool {
	return m.Has(tag.NestedPartyIDSource)
}

//HasNestedPartyRole returns true if NestedPartyRole is present, Tag 538
func (m NoNestedPartyIDs) HasNestedPartyRole() bool {
	return m.Has(tag.NestedPartyRole)
}

//HasNoNestedPartySubIDs returns true if NoNestedPartySubIDs is present, Tag 804
func (m NoNestedPartyIDs) HasNoNestedPartySubIDs() bool {
	return m.Has(tag.NoNestedPartySubIDs)
}

//NoNestedPartySubIDs is a repeating group element, Tag 804
type NoNestedPartySubIDs struct {
	quickfix.Group
}

//SetNestedPartySubID sets NestedPartySubID, Tag 545
func (m NoNestedPartySubIDs) SetNestedPartySubID(v string) {
	m.Set(field.NewNestedPartySubID(v))
}

//SetNestedPartySubIDType sets NestedPartySubIDType, Tag 805
func (m NoNestedPartySubIDs) SetNestedPartySubIDType(v int) {
	m.Set(field.NewNestedPartySubIDType(v))
}

//GetNestedPartySubID gets NestedPartySubID, Tag 545
func (m NoNestedPartySubIDs) GetNestedPartySubID() (f field.NestedPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNestedPartySubIDType gets NestedPartySubIDType, Tag 805
func (m NoNestedPartySubIDs) GetNestedPartySubIDType() (f field.NestedPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasNestedPartySubID returns true if NestedPartySubID is present, Tag 545
func (m NoNestedPartySubIDs) HasNestedPartySubID() bool {
	return m.Has(tag.NestedPartySubID)
}

//HasNestedPartySubIDType returns true if NestedPartySubIDType is present, Tag 805
func (m NoNestedPartySubIDs) HasNestedPartySubIDType() bool {
	return m.Has(tag.NestedPartySubIDType)
}

//NoNestedPartySubIDsRepeatingGroup is a repeating group, Tag 804
type NoNestedPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNestedPartySubIDsRepeatingGroup returns an initialized, NoNestedPartySubIDsRepeatingGroup
func NewNoNestedPartySubIDsRepeatingGroup() NoNestedPartySubIDsRepeatingGroup {
	return NoNestedPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNestedPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedPartySubID), quickfix.GroupElement(tag.NestedPartySubIDType)})}
}

//Add create and append a new NoNestedPartySubIDs to this group
func (m NoNestedPartySubIDsRepeatingGroup) Add() NoNestedPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoNestedPartySubIDs{g}
}

//Get returns the ith NoNestedPartySubIDs in the NoNestedPartySubIDsRepeatinGroup
func (m NoNestedPartySubIDsRepeatingGroup) Get(i int) NoNestedPartySubIDs {
	return NoNestedPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoNestedPartyIDsRepeatingGroup is a repeating group, Tag 539
type NoNestedPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNestedPartyIDsRepeatingGroup returns an initialized, NoNestedPartyIDsRepeatingGroup
func NewNoNestedPartyIDsRepeatingGroup() NoNestedPartyIDsRepeatingGroup {
	return NoNestedPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNestedPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedPartyID), quickfix.GroupElement(tag.NestedPartyIDSource), quickfix.GroupElement(tag.NestedPartyRole), NewNoNestedPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoNestedPartyIDs to this group
func (m NoNestedPartyIDsRepeatingGroup) Add() NoNestedPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoNestedPartyIDs{g}
}

//Get returns the ith NoNestedPartyIDs in the NoNestedPartyIDsRepeatinGroup
func (m NoNestedPartyIDsRepeatingGroup) Get(i int) NoNestedPartyIDs {
	return NoNestedPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoAllocsRepeatingGroup is a repeating group, Tag 78
type NoAllocsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoAllocsRepeatingGroup returns an initialized, NoAllocsRepeatingGroup
func NewNoAllocsRepeatingGroup() NoAllocsRepeatingGroup {
	return NoAllocsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAllocs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AllocAccount), quickfix.GroupElement(tag.AllocAcctIDSource), quickfix.GroupElement(tag.AllocSettlCurrency), quickfix.GroupElement(tag.IndividualAllocID), NewNoNestedPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.AllocQty)})}
}

//Add create and append a new NoAllocs to this group
func (m NoAllocsRepeatingGroup) Add() NoAllocs {
	g := m.RepeatingGroup.Add()
	return NoAllocs{g}
}

//Get returns the ith NoAllocs in the NoAllocsRepeatinGroup
func (m NoAllocsRepeatingGroup) Get(i int) NoAllocs {
	return NoAllocs{m.RepeatingGroup.Get(i)}
}

//NoStipulations is a repeating group element, Tag 232
type NoStipulations struct {
	quickfix.Group
}

//SetStipulationType sets StipulationType, Tag 233
func (m NoStipulations) SetStipulationType(v string) {
	m.Set(field.NewStipulationType(v))
}

//SetStipulationValue sets StipulationValue, Tag 234
func (m NoStipulations) SetStipulationValue(v string) {
	m.Set(field.NewStipulationValue(v))
}

//GetStipulationType gets StipulationType, Tag 233
func (m NoStipulations) GetStipulationType() (f field.StipulationTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStipulationValue gets StipulationValue, Tag 234
func (m NoStipulations) GetStipulationValue() (f field.StipulationValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasStipulationType returns true if StipulationType is present, Tag 233
func (m NoStipulations) HasStipulationType() bool {
	return m.Has(tag.StipulationType)
}

//HasStipulationValue returns true if StipulationValue is present, Tag 234
func (m NoStipulations) HasStipulationValue() bool {
	return m.Has(tag.StipulationValue)
}

//NoStipulationsRepeatingGroup is a repeating group, Tag 232
type NoStipulationsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoStipulationsRepeatingGroup returns an initialized, NoStipulationsRepeatingGroup
func NewNoStipulationsRepeatingGroup() NoStipulationsRepeatingGroup {
	return NoStipulationsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoStipulations,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.StipulationType), quickfix.GroupElement(tag.StipulationValue)})}
}

//Add create and append a new NoStipulations to this group
func (m NoStipulationsRepeatingGroup) Add() NoStipulations {
	g := m.RepeatingGroup.Add()
	return NoStipulations{g}
}

//Get returns the ith NoStipulations in the NoStipulationsRepeatinGroup
func (m NoStipulationsRepeatingGroup) Get(i int) NoStipulations {
	return NoStipulations{m.RepeatingGroup.Get(i)}
}

//NoTradingSessions is a repeating group element, Tag 386
type NoTradingSessions struct {
	quickfix.Group
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoTradingSessions) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m NoTradingSessions) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoTradingSessions) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m NoTradingSessions) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m NoTradingSessions) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m NoTradingSessions) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//NoTradingSessionsRepeatingGroup is a repeating group, Tag 386
type NoTradingSessionsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTradingSessionsRepeatingGroup returns an initialized, NoTradingSessionsRepeatingGroup
func NewNoTradingSessionsRepeatingGroup() NoTradingSessionsRepeatingGroup {
	return NoTradingSessionsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTradingSessions,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.TradingSessionSubID)})}
}

//Add create and append a new NoTradingSessions to this group
func (m NoTradingSessionsRepeatingGroup) Add() NoTradingSessions {
	g := m.RepeatingGroup.Add()
	return NoTradingSessions{g}
}

//Get returns the ith NoTradingSessions in the NoTradingSessionsRepeatinGroup
func (m NoTradingSessionsRepeatingGroup) Get(i int) NoTradingSessions {
	return NoTradingSessions{m.RepeatingGroup.Get(i)}
}

//NoPartyIDs is a repeating group element, Tag 453
type NoPartyIDs struct {
	quickfix.Group
}

//SetPartyID sets PartyID, Tag 448
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

//SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyIDs) SetPartyIDSource(v string) {
	m.Set(field.NewPartyIDSource(v))
}

//SetPartyRole sets PartyRole, Tag 452
func (m NoPartyIDs) SetPartyRole(v int) {
	m.Set(field.NewPartyRole(v))
}

//SetNoPartySubIDs sets NoPartySubIDs, Tag 802
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetPartyID gets PartyID, Tag 448
func (m NoPartyIDs) GetPartyID() (f field.PartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyIDs) GetPartyIDSource() (f field.PartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartyRole gets PartyRole, Tag 452
func (m NoPartyIDs) GetPartyRole() (f field.PartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartySubIDs gets NoPartySubIDs, Tag 802
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasPartyID returns true if PartyID is present, Tag 448
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

//HasPartyIDSource returns true if PartyIDSource is present, Tag 447
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

//HasPartyRole returns true if PartyRole is present, Tag 452
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

//HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

//NoPartySubIDs is a repeating group element, Tag 802
type NoPartySubIDs struct {
	quickfix.Group
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//SetPartySubIDType sets PartySubIDType, Tag 803
func (m NoPartySubIDs) SetPartySubIDType(v int) {
	m.Set(field.NewPartySubIDType(v))
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartySubIDs) GetPartySubID() (f field.PartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartySubIDType gets PartySubIDType, Tag 803
func (m NoPartySubIDs) GetPartySubIDType() (f field.PartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasPartySubID returns true if PartySubID is present, Tag 523
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

//HasPartySubIDType returns true if PartySubIDType is present, Tag 803
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

//NoPartySubIDsRepeatingGroup is a repeating group, Tag 802
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartySubID), quickfix.GroupElement(tag.PartySubIDType)})}
}

//Add create and append a new NoPartySubIDs to this group
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

//Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoPartyIDsRepeatingGroup is a repeating group, Tag 453
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), NewNoPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoPartyIDs to this group
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

//Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoSecurityAltID is a repeating group element, Tag 454
type NoSecurityAltID struct {
	quickfix.Group
}

//SetSecurityAltID sets SecurityAltID, Tag 455
func (m NoSecurityAltID) SetSecurityAltID(v string) {
	m.Set(field.NewSecurityAltID(v))
}

//SetSecurityAltIDSource sets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) SetSecurityAltIDSource(v string) {
	m.Set(field.NewSecurityAltIDSource(v))
}

//GetSecurityAltID gets SecurityAltID, Tag 455
func (m NoSecurityAltID) GetSecurityAltID() (f field.SecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) GetSecurityAltIDSource() (f field.SecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasSecurityAltID returns true if SecurityAltID is present, Tag 455
func (m NoSecurityAltID) HasSecurityAltID() bool {
	return m.Has(tag.SecurityAltID)
}

//HasSecurityAltIDSource returns true if SecurityAltIDSource is present, Tag 456
func (m NoSecurityAltID) HasSecurityAltIDSource() bool {
	return m.Has(tag.SecurityAltIDSource)
}

//NoSecurityAltIDRepeatingGroup is a repeating group, Tag 454
type NoSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSecurityAltIDRepeatingGroup returns an initialized, NoSecurityAltIDRepeatingGroup
func NewNoSecurityAltIDRepeatingGroup() NoSecurityAltIDRepeatingGroup {
	return NoSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SecurityAltID), quickfix.GroupElement(tag.SecurityAltIDSource)})}
}

//Add create and append a new NoSecurityAltID to this group
func (m NoSecurityAltIDRepeatingGroup) Add() NoSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoSecurityAltID{g}
}

//Get returns the ith NoSecurityAltID in the NoSecurityAltIDRepeatinGroup
func (m NoSecurityAltIDRepeatingGroup) Get(i int) NoSecurityAltID {
	return NoSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoUnderlyings is a repeating group element, Tag 711
type NoUnderlyings struct {
	quickfix.Group
}

//SetUnderlyingSymbol sets UnderlyingSymbol, Tag 311
func (m NoUnderlyings) SetUnderlyingSymbol(v string) {
	m.Set(field.NewUnderlyingSymbol(v))
}

//SetUnderlyingSymbolSfx sets UnderlyingSymbolSfx, Tag 312
func (m NoUnderlyings) SetUnderlyingSymbolSfx(v string) {
	m.Set(field.NewUnderlyingSymbolSfx(v))
}

//SetUnderlyingSecurityID sets UnderlyingSecurityID, Tag 309
func (m NoUnderlyings) SetUnderlyingSecurityID(v string) {
	m.Set(field.NewUnderlyingSecurityID(v))
}

//SetUnderlyingSecurityIDSource sets UnderlyingSecurityIDSource, Tag 305
func (m NoUnderlyings) SetUnderlyingSecurityIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityIDSource(v))
}

//SetNoUnderlyingSecurityAltID sets NoUnderlyingSecurityAltID, Tag 457
func (m NoUnderlyings) SetNoUnderlyingSecurityAltID(f NoUnderlyingSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingProduct sets UnderlyingProduct, Tag 462
func (m NoUnderlyings) SetUnderlyingProduct(v int) {
	m.Set(field.NewUnderlyingProduct(v))
}

//SetUnderlyingCFICode sets UnderlyingCFICode, Tag 463
func (m NoUnderlyings) SetUnderlyingCFICode(v string) {
	m.Set(field.NewUnderlyingCFICode(v))
}

//SetUnderlyingSecurityType sets UnderlyingSecurityType, Tag 310
func (m NoUnderlyings) SetUnderlyingSecurityType(v string) {
	m.Set(field.NewUnderlyingSecurityType(v))
}

//SetUnderlyingSecuritySubType sets UnderlyingSecuritySubType, Tag 763
func (m NoUnderlyings) SetUnderlyingSecuritySubType(v string) {
	m.Set(field.NewUnderlyingSecuritySubType(v))
}

//SetUnderlyingMaturityMonthYear sets UnderlyingMaturityMonthYear, Tag 313
func (m NoUnderlyings) SetUnderlyingMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingMaturityMonthYear(v))
}

//SetUnderlyingMaturityDate sets UnderlyingMaturityDate, Tag 542
func (m NoUnderlyings) SetUnderlyingMaturityDate(v string) {
	m.Set(field.NewUnderlyingMaturityDate(v))
}

//SetUnderlyingCouponPaymentDate sets UnderlyingCouponPaymentDate, Tag 241
func (m NoUnderlyings) SetUnderlyingCouponPaymentDate(v string) {
	m.Set(field.NewUnderlyingCouponPaymentDate(v))
}

//SetUnderlyingIssueDate sets UnderlyingIssueDate, Tag 242
func (m NoUnderlyings) SetUnderlyingIssueDate(v string) {
	m.Set(field.NewUnderlyingIssueDate(v))
}

//SetUnderlyingRepoCollateralSecurityType sets UnderlyingRepoCollateralSecurityType, Tag 243
func (m NoUnderlyings) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.Set(field.NewUnderlyingRepoCollateralSecurityType(v))
}

//SetUnderlyingRepurchaseTerm sets UnderlyingRepurchaseTerm, Tag 244
func (m NoUnderlyings) SetUnderlyingRepurchaseTerm(v int) {
	m.Set(field.NewUnderlyingRepurchaseTerm(v))
}

//SetUnderlyingRepurchaseRate sets UnderlyingRepurchaseRate, Tag 245
func (m NoUnderlyings) SetUnderlyingRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingRepurchaseRate(value, scale))
}

//SetUnderlyingFactor sets UnderlyingFactor, Tag 246
func (m NoUnderlyings) SetUnderlyingFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFactor(value, scale))
}

//SetUnderlyingCreditRating sets UnderlyingCreditRating, Tag 256
func (m NoUnderlyings) SetUnderlyingCreditRating(v string) {
	m.Set(field.NewUnderlyingCreditRating(v))
}

//SetUnderlyingInstrRegistry sets UnderlyingInstrRegistry, Tag 595
func (m NoUnderlyings) SetUnderlyingInstrRegistry(v string) {
	m.Set(field.NewUnderlyingInstrRegistry(v))
}

//SetUnderlyingCountryOfIssue sets UnderlyingCountryOfIssue, Tag 592
func (m NoUnderlyings) SetUnderlyingCountryOfIssue(v string) {
	m.Set(field.NewUnderlyingCountryOfIssue(v))
}

//SetUnderlyingStateOrProvinceOfIssue sets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m NoUnderlyings) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.Set(field.NewUnderlyingStateOrProvinceOfIssue(v))
}

//SetUnderlyingLocaleOfIssue sets UnderlyingLocaleOfIssue, Tag 594
func (m NoUnderlyings) SetUnderlyingLocaleOfIssue(v string) {
	m.Set(field.NewUnderlyingLocaleOfIssue(v))
}

//SetUnderlyingRedemptionDate sets UnderlyingRedemptionDate, Tag 247
func (m NoUnderlyings) SetUnderlyingRedemptionDate(v string) {
	m.Set(field.NewUnderlyingRedemptionDate(v))
}

//SetUnderlyingStrikePrice sets UnderlyingStrikePrice, Tag 316
func (m NoUnderlyings) SetUnderlyingStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStrikePrice(value, scale))
}

//SetUnderlyingStrikeCurrency sets UnderlyingStrikeCurrency, Tag 941
func (m NoUnderlyings) SetUnderlyingStrikeCurrency(v string) {
	m.Set(field.NewUnderlyingStrikeCurrency(v))
}

//SetUnderlyingOptAttribute sets UnderlyingOptAttribute, Tag 317
func (m NoUnderlyings) SetUnderlyingOptAttribute(v string) {
	m.Set(field.NewUnderlyingOptAttribute(v))
}

//SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436
func (m NoUnderlyings) SetUnderlyingContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingContractMultiplier(value, scale))
}

//SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435
func (m NoUnderlyings) SetUnderlyingCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCouponRate(value, scale))
}

//SetUnderlyingSecurityExchange sets UnderlyingSecurityExchange, Tag 308
func (m NoUnderlyings) SetUnderlyingSecurityExchange(v string) {
	m.Set(field.NewUnderlyingSecurityExchange(v))
}

//SetUnderlyingIssuer sets UnderlyingIssuer, Tag 306
func (m NoUnderlyings) SetUnderlyingIssuer(v string) {
	m.Set(field.NewUnderlyingIssuer(v))
}

//SetEncodedUnderlyingIssuerLen sets EncodedUnderlyingIssuerLen, Tag 362
func (m NoUnderlyings) SetEncodedUnderlyingIssuerLen(v int) {
	m.Set(field.NewEncodedUnderlyingIssuerLen(v))
}

//SetEncodedUnderlyingIssuer sets EncodedUnderlyingIssuer, Tag 363
func (m NoUnderlyings) SetEncodedUnderlyingIssuer(v string) {
	m.Set(field.NewEncodedUnderlyingIssuer(v))
}

//SetUnderlyingSecurityDesc sets UnderlyingSecurityDesc, Tag 307
func (m NoUnderlyings) SetUnderlyingSecurityDesc(v string) {
	m.Set(field.NewUnderlyingSecurityDesc(v))
}

//SetEncodedUnderlyingSecurityDescLen sets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.Set(field.NewEncodedUnderlyingSecurityDescLen(v))
}

//SetEncodedUnderlyingSecurityDesc sets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDesc(v string) {
	m.Set(field.NewEncodedUnderlyingSecurityDesc(v))
}

//SetUnderlyingCPProgram sets UnderlyingCPProgram, Tag 877
func (m NoUnderlyings) SetUnderlyingCPProgram(v string) {
	m.Set(field.NewUnderlyingCPProgram(v))
}

//SetUnderlyingCPRegType sets UnderlyingCPRegType, Tag 878
func (m NoUnderlyings) SetUnderlyingCPRegType(v string) {
	m.Set(field.NewUnderlyingCPRegType(v))
}

//SetUnderlyingCurrency sets UnderlyingCurrency, Tag 318
func (m NoUnderlyings) SetUnderlyingCurrency(v string) {
	m.Set(field.NewUnderlyingCurrency(v))
}

//SetUnderlyingQty sets UnderlyingQty, Tag 879
func (m NoUnderlyings) SetUnderlyingQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingQty(value, scale))
}

//SetUnderlyingPx sets UnderlyingPx, Tag 810
func (m NoUnderlyings) SetUnderlyingPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPx(value, scale))
}

//SetUnderlyingDirtyPrice sets UnderlyingDirtyPrice, Tag 882
func (m NoUnderlyings) SetUnderlyingDirtyPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDirtyPrice(value, scale))
}

//SetUnderlyingEndPrice sets UnderlyingEndPrice, Tag 883
func (m NoUnderlyings) SetUnderlyingEndPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndPrice(value, scale))
}

//SetUnderlyingStartValue sets UnderlyingStartValue, Tag 884
func (m NoUnderlyings) SetUnderlyingStartValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStartValue(value, scale))
}

//SetUnderlyingCurrentValue sets UnderlyingCurrentValue, Tag 885
func (m NoUnderlyings) SetUnderlyingCurrentValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCurrentValue(value, scale))
}

//SetUnderlyingEndValue sets UnderlyingEndValue, Tag 886
func (m NoUnderlyings) SetUnderlyingEndValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndValue(value, scale))
}

//SetNoUnderlyingStips sets NoUnderlyingStips, Tag 887
func (m NoUnderlyings) SetNoUnderlyingStips(f NoUnderlyingStipsRepeatingGroup) {
	m.SetGroup(f)
}

//GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311
func (m NoUnderlyings) GetUnderlyingSymbol() (f field.UnderlyingSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312
func (m NoUnderlyings) GetUnderlyingSymbolSfx() (f field.UnderlyingSymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309
func (m NoUnderlyings) GetUnderlyingSecurityID() (f field.UnderlyingSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityIDSource gets UnderlyingSecurityIDSource, Tag 305
func (m NoUnderlyings) GetUnderlyingSecurityIDSource() (f field.UnderlyingSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyingSecurityAltID gets NoUnderlyingSecurityAltID, Tag 457
func (m NoUnderlyings) GetNoUnderlyingSecurityAltID() (NoUnderlyingSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingProduct gets UnderlyingProduct, Tag 462
func (m NoUnderlyings) GetUnderlyingProduct() (f field.UnderlyingProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCFICode gets UnderlyingCFICode, Tag 463
func (m NoUnderlyings) GetUnderlyingCFICode() (f field.UnderlyingCFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310
func (m NoUnderlyings) GetUnderlyingSecurityType() (f field.UnderlyingSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecuritySubType gets UnderlyingSecuritySubType, Tag 763
func (m NoUnderlyings) GetUnderlyingSecuritySubType() (f field.UnderlyingSecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313
func (m NoUnderlyings) GetUnderlyingMaturityMonthYear() (f field.UnderlyingMaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingMaturityDate gets UnderlyingMaturityDate, Tag 542
func (m NoUnderlyings) GetUnderlyingMaturityDate() (f field.UnderlyingMaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCouponPaymentDate gets UnderlyingCouponPaymentDate, Tag 241
func (m NoUnderlyings) GetUnderlyingCouponPaymentDate() (f field.UnderlyingCouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingIssueDate gets UnderlyingIssueDate, Tag 242
func (m NoUnderlyings) GetUnderlyingIssueDate() (f field.UnderlyingIssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRepoCollateralSecurityType gets UnderlyingRepoCollateralSecurityType, Tag 243
func (m NoUnderlyings) GetUnderlyingRepoCollateralSecurityType() (f field.UnderlyingRepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRepurchaseTerm gets UnderlyingRepurchaseTerm, Tag 244
func (m NoUnderlyings) GetUnderlyingRepurchaseTerm() (f field.UnderlyingRepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRepurchaseRate gets UnderlyingRepurchaseRate, Tag 245
func (m NoUnderlyings) GetUnderlyingRepurchaseRate() (f field.UnderlyingRepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFactor gets UnderlyingFactor, Tag 246
func (m NoUnderlyings) GetUnderlyingFactor() (f field.UnderlyingFactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCreditRating gets UnderlyingCreditRating, Tag 256
func (m NoUnderlyings) GetUnderlyingCreditRating() (f field.UnderlyingCreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingInstrRegistry gets UnderlyingInstrRegistry, Tag 595
func (m NoUnderlyings) GetUnderlyingInstrRegistry() (f field.UnderlyingInstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCountryOfIssue gets UnderlyingCountryOfIssue, Tag 592
func (m NoUnderlyings) GetUnderlyingCountryOfIssue() (f field.UnderlyingCountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStateOrProvinceOfIssue gets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m NoUnderlyings) GetUnderlyingStateOrProvinceOfIssue() (f field.UnderlyingStateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLocaleOfIssue gets UnderlyingLocaleOfIssue, Tag 594
func (m NoUnderlyings) GetUnderlyingLocaleOfIssue() (f field.UnderlyingLocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRedemptionDate gets UnderlyingRedemptionDate, Tag 247
func (m NoUnderlyings) GetUnderlyingRedemptionDate() (f field.UnderlyingRedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316
func (m NoUnderlyings) GetUnderlyingStrikePrice() (f field.UnderlyingStrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStrikeCurrency gets UnderlyingStrikeCurrency, Tag 941
func (m NoUnderlyings) GetUnderlyingStrikeCurrency() (f field.UnderlyingStrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317
func (m NoUnderlyings) GetUnderlyingOptAttribute() (f field.UnderlyingOptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436
func (m NoUnderlyings) GetUnderlyingContractMultiplier() (f field.UnderlyingContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435
func (m NoUnderlyings) GetUnderlyingCouponRate() (f field.UnderlyingCouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308
func (m NoUnderlyings) GetUnderlyingSecurityExchange() (f field.UnderlyingSecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306
func (m NoUnderlyings) GetUnderlyingIssuer() (f field.UnderlyingIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362
func (m NoUnderlyings) GetEncodedUnderlyingIssuerLen() (f field.EncodedUnderlyingIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363
func (m NoUnderlyings) GetEncodedUnderlyingIssuer() (f field.EncodedUnderlyingIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307
func (m NoUnderlyings) GetUnderlyingSecurityDesc() (f field.UnderlyingSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDescLen() (f field.EncodedUnderlyingSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDesc() (f field.EncodedUnderlyingSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCPProgram gets UnderlyingCPProgram, Tag 877
func (m NoUnderlyings) GetUnderlyingCPProgram() (f field.UnderlyingCPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCPRegType gets UnderlyingCPRegType, Tag 878
func (m NoUnderlyings) GetUnderlyingCPRegType() (f field.UnderlyingCPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCurrency gets UnderlyingCurrency, Tag 318
func (m NoUnderlyings) GetUnderlyingCurrency() (f field.UnderlyingCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingQty gets UnderlyingQty, Tag 879
func (m NoUnderlyings) GetUnderlyingQty() (f field.UnderlyingQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPx gets UnderlyingPx, Tag 810
func (m NoUnderlyings) GetUnderlyingPx() (f field.UnderlyingPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingDirtyPrice gets UnderlyingDirtyPrice, Tag 882
func (m NoUnderlyings) GetUnderlyingDirtyPrice() (f field.UnderlyingDirtyPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingEndPrice gets UnderlyingEndPrice, Tag 883
func (m NoUnderlyings) GetUnderlyingEndPrice() (f field.UnderlyingEndPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStartValue gets UnderlyingStartValue, Tag 884
func (m NoUnderlyings) GetUnderlyingStartValue() (f field.UnderlyingStartValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCurrentValue gets UnderlyingCurrentValue, Tag 885
func (m NoUnderlyings) GetUnderlyingCurrentValue() (f field.UnderlyingCurrentValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingEndValue gets UnderlyingEndValue, Tag 886
func (m NoUnderlyings) GetUnderlyingEndValue() (f field.UnderlyingEndValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyingStips gets NoUnderlyingStips, Tag 887
func (m NoUnderlyings) GetNoUnderlyingStips() (NoUnderlyingStipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingStipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasUnderlyingSymbol returns true if UnderlyingSymbol is present, Tag 311
func (m NoUnderlyings) HasUnderlyingSymbol() bool {
	return m.Has(tag.UnderlyingSymbol)
}

//HasUnderlyingSymbolSfx returns true if UnderlyingSymbolSfx is present, Tag 312
func (m NoUnderlyings) HasUnderlyingSymbolSfx() bool {
	return m.Has(tag.UnderlyingSymbolSfx)
}

//HasUnderlyingSecurityID returns true if UnderlyingSecurityID is present, Tag 309
func (m NoUnderlyings) HasUnderlyingSecurityID() bool {
	return m.Has(tag.UnderlyingSecurityID)
}

//HasUnderlyingSecurityIDSource returns true if UnderlyingSecurityIDSource is present, Tag 305
func (m NoUnderlyings) HasUnderlyingSecurityIDSource() bool {
	return m.Has(tag.UnderlyingSecurityIDSource)
}

//HasNoUnderlyingSecurityAltID returns true if NoUnderlyingSecurityAltID is present, Tag 457
func (m NoUnderlyings) HasNoUnderlyingSecurityAltID() bool {
	return m.Has(tag.NoUnderlyingSecurityAltID)
}

//HasUnderlyingProduct returns true if UnderlyingProduct is present, Tag 462
func (m NoUnderlyings) HasUnderlyingProduct() bool {
	return m.Has(tag.UnderlyingProduct)
}

//HasUnderlyingCFICode returns true if UnderlyingCFICode is present, Tag 463
func (m NoUnderlyings) HasUnderlyingCFICode() bool {
	return m.Has(tag.UnderlyingCFICode)
}

//HasUnderlyingSecurityType returns true if UnderlyingSecurityType is present, Tag 310
func (m NoUnderlyings) HasUnderlyingSecurityType() bool {
	return m.Has(tag.UnderlyingSecurityType)
}

//HasUnderlyingSecuritySubType returns true if UnderlyingSecuritySubType is present, Tag 763
func (m NoUnderlyings) HasUnderlyingSecuritySubType() bool {
	return m.Has(tag.UnderlyingSecuritySubType)
}

//HasUnderlyingMaturityMonthYear returns true if UnderlyingMaturityMonthYear is present, Tag 313
func (m NoUnderlyings) HasUnderlyingMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingMaturityMonthYear)
}

//HasUnderlyingMaturityDate returns true if UnderlyingMaturityDate is present, Tag 542
func (m NoUnderlyings) HasUnderlyingMaturityDate() bool {
	return m.Has(tag.UnderlyingMaturityDate)
}

//HasUnderlyingCouponPaymentDate returns true if UnderlyingCouponPaymentDate is present, Tag 241
func (m NoUnderlyings) HasUnderlyingCouponPaymentDate() bool {
	return m.Has(tag.UnderlyingCouponPaymentDate)
}

//HasUnderlyingIssueDate returns true if UnderlyingIssueDate is present, Tag 242
func (m NoUnderlyings) HasUnderlyingIssueDate() bool {
	return m.Has(tag.UnderlyingIssueDate)
}

//HasUnderlyingRepoCollateralSecurityType returns true if UnderlyingRepoCollateralSecurityType is present, Tag 243
func (m NoUnderlyings) HasUnderlyingRepoCollateralSecurityType() bool {
	return m.Has(tag.UnderlyingRepoCollateralSecurityType)
}

//HasUnderlyingRepurchaseTerm returns true if UnderlyingRepurchaseTerm is present, Tag 244
func (m NoUnderlyings) HasUnderlyingRepurchaseTerm() bool {
	return m.Has(tag.UnderlyingRepurchaseTerm)
}

//HasUnderlyingRepurchaseRate returns true if UnderlyingRepurchaseRate is present, Tag 245
func (m NoUnderlyings) HasUnderlyingRepurchaseRate() bool {
	return m.Has(tag.UnderlyingRepurchaseRate)
}

//HasUnderlyingFactor returns true if UnderlyingFactor is present, Tag 246
func (m NoUnderlyings) HasUnderlyingFactor() bool {
	return m.Has(tag.UnderlyingFactor)
}

//HasUnderlyingCreditRating returns true if UnderlyingCreditRating is present, Tag 256
func (m NoUnderlyings) HasUnderlyingCreditRating() bool {
	return m.Has(tag.UnderlyingCreditRating)
}

//HasUnderlyingInstrRegistry returns true if UnderlyingInstrRegistry is present, Tag 595
func (m NoUnderlyings) HasUnderlyingInstrRegistry() bool {
	return m.Has(tag.UnderlyingInstrRegistry)
}

//HasUnderlyingCountryOfIssue returns true if UnderlyingCountryOfIssue is present, Tag 592
func (m NoUnderlyings) HasUnderlyingCountryOfIssue() bool {
	return m.Has(tag.UnderlyingCountryOfIssue)
}

//HasUnderlyingStateOrProvinceOfIssue returns true if UnderlyingStateOrProvinceOfIssue is present, Tag 593
func (m NoUnderlyings) HasUnderlyingStateOrProvinceOfIssue() bool {
	return m.Has(tag.UnderlyingStateOrProvinceOfIssue)
}

//HasUnderlyingLocaleOfIssue returns true if UnderlyingLocaleOfIssue is present, Tag 594
func (m NoUnderlyings) HasUnderlyingLocaleOfIssue() bool {
	return m.Has(tag.UnderlyingLocaleOfIssue)
}

//HasUnderlyingRedemptionDate returns true if UnderlyingRedemptionDate is present, Tag 247
func (m NoUnderlyings) HasUnderlyingRedemptionDate() bool {
	return m.Has(tag.UnderlyingRedemptionDate)
}

//HasUnderlyingStrikePrice returns true if UnderlyingStrikePrice is present, Tag 316
func (m NoUnderlyings) HasUnderlyingStrikePrice() bool {
	return m.Has(tag.UnderlyingStrikePrice)
}

//HasUnderlyingStrikeCurrency returns true if UnderlyingStrikeCurrency is present, Tag 941
func (m NoUnderlyings) HasUnderlyingStrikeCurrency() bool {
	return m.Has(tag.UnderlyingStrikeCurrency)
}

//HasUnderlyingOptAttribute returns true if UnderlyingOptAttribute is present, Tag 317
func (m NoUnderlyings) HasUnderlyingOptAttribute() bool {
	return m.Has(tag.UnderlyingOptAttribute)
}

//HasUnderlyingContractMultiplier returns true if UnderlyingContractMultiplier is present, Tag 436
func (m NoUnderlyings) HasUnderlyingContractMultiplier() bool {
	return m.Has(tag.UnderlyingContractMultiplier)
}

//HasUnderlyingCouponRate returns true if UnderlyingCouponRate is present, Tag 435
func (m NoUnderlyings) HasUnderlyingCouponRate() bool {
	return m.Has(tag.UnderlyingCouponRate)
}

//HasUnderlyingSecurityExchange returns true if UnderlyingSecurityExchange is present, Tag 308
func (m NoUnderlyings) HasUnderlyingSecurityExchange() bool {
	return m.Has(tag.UnderlyingSecurityExchange)
}

//HasUnderlyingIssuer returns true if UnderlyingIssuer is present, Tag 306
func (m NoUnderlyings) HasUnderlyingIssuer() bool {
	return m.Has(tag.UnderlyingIssuer)
}

//HasEncodedUnderlyingIssuerLen returns true if EncodedUnderlyingIssuerLen is present, Tag 362
func (m NoUnderlyings) HasEncodedUnderlyingIssuerLen() bool {
	return m.Has(tag.EncodedUnderlyingIssuerLen)
}

//HasEncodedUnderlyingIssuer returns true if EncodedUnderlyingIssuer is present, Tag 363
func (m NoUnderlyings) HasEncodedUnderlyingIssuer() bool {
	return m.Has(tag.EncodedUnderlyingIssuer)
}

//HasUnderlyingSecurityDesc returns true if UnderlyingSecurityDesc is present, Tag 307
func (m NoUnderlyings) HasUnderlyingSecurityDesc() bool {
	return m.Has(tag.UnderlyingSecurityDesc)
}

//HasEncodedUnderlyingSecurityDescLen returns true if EncodedUnderlyingSecurityDescLen is present, Tag 364
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDescLen() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDescLen)
}

//HasEncodedUnderlyingSecurityDesc returns true if EncodedUnderlyingSecurityDesc is present, Tag 365
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDesc() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDesc)
}

//HasUnderlyingCPProgram returns true if UnderlyingCPProgram is present, Tag 877
func (m NoUnderlyings) HasUnderlyingCPProgram() bool {
	return m.Has(tag.UnderlyingCPProgram)
}

//HasUnderlyingCPRegType returns true if UnderlyingCPRegType is present, Tag 878
func (m NoUnderlyings) HasUnderlyingCPRegType() bool {
	return m.Has(tag.UnderlyingCPRegType)
}

//HasUnderlyingCurrency returns true if UnderlyingCurrency is present, Tag 318
func (m NoUnderlyings) HasUnderlyingCurrency() bool {
	return m.Has(tag.UnderlyingCurrency)
}

//HasUnderlyingQty returns true if UnderlyingQty is present, Tag 879
func (m NoUnderlyings) HasUnderlyingQty() bool {
	return m.Has(tag.UnderlyingQty)
}

//HasUnderlyingPx returns true if UnderlyingPx is present, Tag 810
func (m NoUnderlyings) HasUnderlyingPx() bool {
	return m.Has(tag.UnderlyingPx)
}

//HasUnderlyingDirtyPrice returns true if UnderlyingDirtyPrice is present, Tag 882
func (m NoUnderlyings) HasUnderlyingDirtyPrice() bool {
	return m.Has(tag.UnderlyingDirtyPrice)
}

//HasUnderlyingEndPrice returns true if UnderlyingEndPrice is present, Tag 883
func (m NoUnderlyings) HasUnderlyingEndPrice() bool {
	return m.Has(tag.UnderlyingEndPrice)
}

//HasUnderlyingStartValue returns true if UnderlyingStartValue is present, Tag 884
func (m NoUnderlyings) HasUnderlyingStartValue() bool {
	return m.Has(tag.UnderlyingStartValue)
}

//HasUnderlyingCurrentValue returns true if UnderlyingCurrentValue is present, Tag 885
func (m NoUnderlyings) HasUnderlyingCurrentValue() bool {
	return m.Has(tag.UnderlyingCurrentValue)
}

//HasUnderlyingEndValue returns true if UnderlyingEndValue is present, Tag 886
func (m NoUnderlyings) HasUnderlyingEndValue() bool {
	return m.Has(tag.UnderlyingEndValue)
}

//HasNoUnderlyingStips returns true if NoUnderlyingStips is present, Tag 887
func (m NoUnderlyings) HasNoUnderlyingStips() bool {
	return m.Has(tag.NoUnderlyingStips)
}

//NoUnderlyingSecurityAltID is a repeating group element, Tag 457
type NoUnderlyingSecurityAltID struct {
	quickfix.Group
}

//SetUnderlyingSecurityAltID sets UnderlyingSecurityAltID, Tag 458
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltID(v string) {
	m.Set(field.NewUnderlyingSecurityAltID(v))
}

//SetUnderlyingSecurityAltIDSource sets UnderlyingSecurityAltIDSource, Tag 459
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityAltIDSource(v))
}

//GetUnderlyingSecurityAltID gets UnderlyingSecurityAltID, Tag 458
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltID() (f field.UnderlyingSecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityAltIDSource gets UnderlyingSecurityAltIDSource, Tag 459
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltIDSource() (f field.UnderlyingSecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingSecurityAltID returns true if UnderlyingSecurityAltID is present, Tag 458
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltID() bool {
	return m.Has(tag.UnderlyingSecurityAltID)
}

//HasUnderlyingSecurityAltIDSource returns true if UnderlyingSecurityAltIDSource is present, Tag 459
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltIDSource() bool {
	return m.Has(tag.UnderlyingSecurityAltIDSource)
}

//NoUnderlyingSecurityAltIDRepeatingGroup is a repeating group, Tag 457
type NoUnderlyingSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingSecurityAltIDRepeatingGroup returns an initialized, NoUnderlyingSecurityAltIDRepeatingGroup
func NewNoUnderlyingSecurityAltIDRepeatingGroup() NoUnderlyingSecurityAltIDRepeatingGroup {
	return NoUnderlyingSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSecurityAltID), quickfix.GroupElement(tag.UnderlyingSecurityAltIDSource)})}
}

//Add create and append a new NoUnderlyingSecurityAltID to this group
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Add() NoUnderlyingSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingSecurityAltID{g}
}

//Get returns the ith NoUnderlyingSecurityAltID in the NoUnderlyingSecurityAltIDRepeatinGroup
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Get(i int) NoUnderlyingSecurityAltID {
	return NoUnderlyingSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoUnderlyingStips is a repeating group element, Tag 887
type NoUnderlyingStips struct {
	quickfix.Group
}

//SetUnderlyingStipType sets UnderlyingStipType, Tag 888
func (m NoUnderlyingStips) SetUnderlyingStipType(v string) {
	m.Set(field.NewUnderlyingStipType(v))
}

//SetUnderlyingStipValue sets UnderlyingStipValue, Tag 889
func (m NoUnderlyingStips) SetUnderlyingStipValue(v string) {
	m.Set(field.NewUnderlyingStipValue(v))
}

//GetUnderlyingStipType gets UnderlyingStipType, Tag 888
func (m NoUnderlyingStips) GetUnderlyingStipType() (f field.UnderlyingStipTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStipValue gets UnderlyingStipValue, Tag 889
func (m NoUnderlyingStips) GetUnderlyingStipValue() (f field.UnderlyingStipValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingStipType returns true if UnderlyingStipType is present, Tag 888
func (m NoUnderlyingStips) HasUnderlyingStipType() bool {
	return m.Has(tag.UnderlyingStipType)
}

//HasUnderlyingStipValue returns true if UnderlyingStipValue is present, Tag 889
func (m NoUnderlyingStips) HasUnderlyingStipValue() bool {
	return m.Has(tag.UnderlyingStipValue)
}

//NoUnderlyingStipsRepeatingGroup is a repeating group, Tag 887
type NoUnderlyingStipsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingStipsRepeatingGroup returns an initialized, NoUnderlyingStipsRepeatingGroup
func NewNoUnderlyingStipsRepeatingGroup() NoUnderlyingStipsRepeatingGroup {
	return NoUnderlyingStipsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingStips,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingStipType), quickfix.GroupElement(tag.UnderlyingStipValue)})}
}

//Add create and append a new NoUnderlyingStips to this group
func (m NoUnderlyingStipsRepeatingGroup) Add() NoUnderlyingStips {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingStips{g}
}

//Get returns the ith NoUnderlyingStips in the NoUnderlyingStipsRepeatinGroup
func (m NoUnderlyingStipsRepeatingGroup) Get(i int) NoUnderlyingStips {
	return NoUnderlyingStips{m.RepeatingGroup.Get(i)}
}

//NoUnderlyingsRepeatingGroup is a repeating group, Tag 711
type NoUnderlyingsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingsRepeatingGroup returns an initialized, NoUnderlyingsRepeatingGroup
func NewNoUnderlyingsRepeatingGroup() NoUnderlyingsRepeatingGroup {
	return NoUnderlyingsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyings,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingSecurityIDSource), NewNoUnderlyingSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingProduct), quickfix.GroupElement(tag.UnderlyingCFICode), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingSecuritySubType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDate), quickfix.GroupElement(tag.UnderlyingCouponPaymentDate), quickfix.GroupElement(tag.UnderlyingIssueDate), quickfix.GroupElement(tag.UnderlyingRepoCollateralSecurityType), quickfix.GroupElement(tag.UnderlyingRepurchaseTerm), quickfix.GroupElement(tag.UnderlyingRepurchaseRate), quickfix.GroupElement(tag.UnderlyingFactor), quickfix.GroupElement(tag.UnderlyingCreditRating), quickfix.GroupElement(tag.UnderlyingInstrRegistry), quickfix.GroupElement(tag.UnderlyingCountryOfIssue), quickfix.GroupElement(tag.UnderlyingStateOrProvinceOfIssue), quickfix.GroupElement(tag.UnderlyingLocaleOfIssue), quickfix.GroupElement(tag.UnderlyingRedemptionDate), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingStrikeCurrency), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.UnderlyingCPProgram), quickfix.GroupElement(tag.UnderlyingCPRegType), quickfix.GroupElement(tag.UnderlyingCurrency), quickfix.GroupElement(tag.UnderlyingQty), quickfix.GroupElement(tag.UnderlyingPx), quickfix.GroupElement(tag.UnderlyingDirtyPrice), quickfix.GroupElement(tag.UnderlyingEndPrice), quickfix.GroupElement(tag.UnderlyingStartValue), quickfix.GroupElement(tag.UnderlyingCurrentValue), quickfix.GroupElement(tag.UnderlyingEndValue), NewNoUnderlyingStipsRepeatingGroup()})}
}

//Add create and append a new NoUnderlyings to this group
func (m NoUnderlyingsRepeatingGroup) Add() NoUnderlyings {
	g := m.RepeatingGroup.Add()
	return NoUnderlyings{g}
}

//Get returns the ith NoUnderlyings in the NoUnderlyingsRepeatinGroup
func (m NoUnderlyingsRepeatingGroup) Get(i int) NoUnderlyings {
	return NoUnderlyings{m.RepeatingGroup.Get(i)}
}

//NoEvents is a repeating group element, Tag 864
type NoEvents struct {
	quickfix.Group
}

//SetEventType sets EventType, Tag 865
func (m NoEvents) SetEventType(v int) {
	m.Set(field.NewEventType(v))
}

//SetEventDate sets EventDate, Tag 866
func (m NoEvents) SetEventDate(v string) {
	m.Set(field.NewEventDate(v))
}

//SetEventPx sets EventPx, Tag 867
func (m NoEvents) SetEventPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewEventPx(value, scale))
}

//SetEventText sets EventText, Tag 868
func (m NoEvents) SetEventText(v string) {
	m.Set(field.NewEventText(v))
}

//GetEventType gets EventType, Tag 865
func (m NoEvents) GetEventType() (f field.EventTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEventDate gets EventDate, Tag 866
func (m NoEvents) GetEventDate() (f field.EventDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEventPx gets EventPx, Tag 867
func (m NoEvents) GetEventPx() (f field.EventPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEventText gets EventText, Tag 868
func (m NoEvents) GetEventText() (f field.EventTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasEventType returns true if EventType is present, Tag 865
func (m NoEvents) HasEventType() bool {
	return m.Has(tag.EventType)
}

//HasEventDate returns true if EventDate is present, Tag 866
func (m NoEvents) HasEventDate() bool {
	return m.Has(tag.EventDate)
}

//HasEventPx returns true if EventPx is present, Tag 867
func (m NoEvents) HasEventPx() bool {
	return m.Has(tag.EventPx)
}

//HasEventText returns true if EventText is present, Tag 868
func (m NoEvents) HasEventText() bool {
	return m.Has(tag.EventText)
}

//NoEventsRepeatingGroup is a repeating group, Tag 864
type NoEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoEventsRepeatingGroup returns an initialized, NoEventsRepeatingGroup
func NewNoEventsRepeatingGroup() NoEventsRepeatingGroup {
	return NoEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.EventType), quickfix.GroupElement(tag.EventDate), quickfix.GroupElement(tag.EventPx), quickfix.GroupElement(tag.EventText)})}
}

//Add create and append a new NoEvents to this group
func (m NoEventsRepeatingGroup) Add() NoEvents {
	g := m.RepeatingGroup.Add()
	return NoEvents{g}
}

//Get returns the ith NoEvents in the NoEventsRepeatinGroup
func (m NoEventsRepeatingGroup) Get(i int) NoEvents {
	return NoEvents{m.RepeatingGroup.Get(i)}
}
