package ordercancelreplacerequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//OrderCancelReplaceRequest is the fix50sp2 OrderCancelReplaceRequest type, MsgType = G
type OrderCancelReplaceRequest struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a OrderCancelReplaceRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) OrderCancelReplaceRequest {
	return OrderCancelReplaceRequest{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m OrderCancelReplaceRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a OrderCancelReplaceRequest initialized with the required fields for OrderCancelReplaceRequest
func New(clordid field.ClOrdIDField, side field.SideField, transacttime field.TransactTimeField, ordtype field.OrdTypeField) (m OrderCancelReplaceRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("G"))
	m.Set(clordid)
	m.Set(side)
	m.Set(transacttime)
	m.Set(ordtype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "G", r
}

//SetAccount sets Account, Tag 1
func (m OrderCancelReplaceRequest) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m OrderCancelReplaceRequest) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetCommission sets Commission, Tag 12
func (m OrderCancelReplaceRequest) SetCommission(v float64) {
	m.Set(field.NewCommission(v))
}

//SetCommType sets CommType, Tag 13
func (m OrderCancelReplaceRequest) SetCommType(v string) {
	m.Set(field.NewCommType(v))
}

//SetCurrency sets Currency, Tag 15
func (m OrderCancelReplaceRequest) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetExecInst sets ExecInst, Tag 18
func (m OrderCancelReplaceRequest) SetExecInst(v string) {
	m.Set(field.NewExecInst(v))
}

//SetHandlInst sets HandlInst, Tag 21
func (m OrderCancelReplaceRequest) SetHandlInst(v string) {
	m.Set(field.NewHandlInst(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m OrderCancelReplaceRequest) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetOrderID sets OrderID, Tag 37
func (m OrderCancelReplaceRequest) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m OrderCancelReplaceRequest) SetOrderQty(v float64) {
	m.Set(field.NewOrderQty(v))
}

//SetOrdType sets OrdType, Tag 40
func (m OrderCancelReplaceRequest) SetOrdType(v string) {
	m.Set(field.NewOrdType(v))
}

//SetOrigClOrdID sets OrigClOrdID, Tag 41
func (m OrderCancelReplaceRequest) SetOrigClOrdID(v string) {
	m.Set(field.NewOrigClOrdID(v))
}

//SetPrice sets Price, Tag 44
func (m OrderCancelReplaceRequest) SetPrice(v float64) {
	m.Set(field.NewPrice(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m OrderCancelReplaceRequest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m OrderCancelReplaceRequest) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m OrderCancelReplaceRequest) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m OrderCancelReplaceRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTimeInForce sets TimeInForce, Tag 59
func (m OrderCancelReplaceRequest) SetTimeInForce(v string) {
	m.Set(field.NewTimeInForce(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m OrderCancelReplaceRequest) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSettlType sets SettlType, Tag 63
func (m OrderCancelReplaceRequest) SetSettlType(v string) {
	m.Set(field.NewSettlType(v))
}

//SetSettlDate sets SettlDate, Tag 64
func (m OrderCancelReplaceRequest) SetSettlDate(v string) {
	m.Set(field.NewSettlDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m OrderCancelReplaceRequest) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetListID sets ListID, Tag 66
func (m OrderCancelReplaceRequest) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetAllocID sets AllocID, Tag 70
func (m OrderCancelReplaceRequest) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m OrderCancelReplaceRequest) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetPositionEffect sets PositionEffect, Tag 77
func (m OrderCancelReplaceRequest) SetPositionEffect(v string) {
	m.Set(field.NewPositionEffect(v))
}

//SetNoAllocs sets NoAllocs, Tag 78
func (m OrderCancelReplaceRequest) SetNoAllocs(f NoAllocsRepeatingGroup) {
	m.SetGroup(f)
}

//SetStopPx sets StopPx, Tag 99
func (m OrderCancelReplaceRequest) SetStopPx(v float64) {
	m.Set(field.NewStopPx(v))
}

//SetExDestination sets ExDestination, Tag 100
func (m OrderCancelReplaceRequest) SetExDestination(v string) {
	m.Set(field.NewExDestination(v))
}

//SetIssuer sets Issuer, Tag 106
func (m OrderCancelReplaceRequest) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m OrderCancelReplaceRequest) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetMinQty sets MinQty, Tag 110
func (m OrderCancelReplaceRequest) SetMinQty(v float64) {
	m.Set(field.NewMinQty(v))
}

//SetMaxFloor sets MaxFloor, Tag 111
func (m OrderCancelReplaceRequest) SetMaxFloor(v float64) {
	m.Set(field.NewMaxFloor(v))
}

//SetLocateReqd sets LocateReqd, Tag 114
func (m OrderCancelReplaceRequest) SetLocateReqd(v bool) {
	m.Set(field.NewLocateReqd(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m OrderCancelReplaceRequest) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetForexReq sets ForexReq, Tag 121
func (m OrderCancelReplaceRequest) SetForexReq(v bool) {
	m.Set(field.NewForexReq(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m OrderCancelReplaceRequest) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m OrderCancelReplaceRequest) SetCashOrderQty(v float64) {
	m.Set(field.NewCashOrderQty(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m OrderCancelReplaceRequest) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetEffectiveTime sets EffectiveTime, Tag 168
func (m OrderCancelReplaceRequest) SetEffectiveTime(v time.Time) {
	m.Set(field.NewEffectiveTime(v))
}

//SetOrderQty2 sets OrderQty2, Tag 192
func (m OrderCancelReplaceRequest) SetOrderQty2(v float64) {
	m.Set(field.NewOrderQty2(v))
}

//SetSettlDate2 sets SettlDate2, Tag 193
func (m OrderCancelReplaceRequest) SetSettlDate2(v string) {
	m.Set(field.NewSettlDate2(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m OrderCancelReplaceRequest) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m OrderCancelReplaceRequest) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m OrderCancelReplaceRequest) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetCoveredOrUncovered sets CoveredOrUncovered, Tag 203
func (m OrderCancelReplaceRequest) SetCoveredOrUncovered(v int) {
	m.Set(field.NewCoveredOrUncovered(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m OrderCancelReplaceRequest) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m OrderCancelReplaceRequest) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetMaxShow sets MaxShow, Tag 210
func (m OrderCancelReplaceRequest) SetMaxShow(v float64) {
	m.Set(field.NewMaxShow(v))
}

//SetPegOffsetValue sets PegOffsetValue, Tag 211
func (m OrderCancelReplaceRequest) SetPegOffsetValue(v float64) {
	m.Set(field.NewPegOffsetValue(v))
}

//SetSpread sets Spread, Tag 218
func (m OrderCancelReplaceRequest) SetSpread(v float64) {
	m.Set(field.NewSpread(v))
}

//SetBenchmarkCurveCurrency sets BenchmarkCurveCurrency, Tag 220
func (m OrderCancelReplaceRequest) SetBenchmarkCurveCurrency(v string) {
	m.Set(field.NewBenchmarkCurveCurrency(v))
}

//SetBenchmarkCurveName sets BenchmarkCurveName, Tag 221
func (m OrderCancelReplaceRequest) SetBenchmarkCurveName(v string) {
	m.Set(field.NewBenchmarkCurveName(v))
}

//SetBenchmarkCurvePoint sets BenchmarkCurvePoint, Tag 222
func (m OrderCancelReplaceRequest) SetBenchmarkCurvePoint(v string) {
	m.Set(field.NewBenchmarkCurvePoint(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m OrderCancelReplaceRequest) SetCouponRate(v float64) {
	m.Set(field.NewCouponRate(v))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m OrderCancelReplaceRequest) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m OrderCancelReplaceRequest) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m OrderCancelReplaceRequest) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m OrderCancelReplaceRequest) SetRepurchaseRate(v float64) {
	m.Set(field.NewRepurchaseRate(v))
}

//SetFactor sets Factor, Tag 228
func (m OrderCancelReplaceRequest) SetFactor(v float64) {
	m.Set(field.NewFactor(v))
}

//SetTradeOriginationDate sets TradeOriginationDate, Tag 229
func (m OrderCancelReplaceRequest) SetTradeOriginationDate(v string) {
	m.Set(field.NewTradeOriginationDate(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m OrderCancelReplaceRequest) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
}

//SetYieldType sets YieldType, Tag 235
func (m OrderCancelReplaceRequest) SetYieldType(v string) {
	m.Set(field.NewYieldType(v))
}

//SetYield sets Yield, Tag 236
func (m OrderCancelReplaceRequest) SetYield(v float64) {
	m.Set(field.NewYield(v))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m OrderCancelReplaceRequest) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m OrderCancelReplaceRequest) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m OrderCancelReplaceRequest) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m OrderCancelReplaceRequest) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m OrderCancelReplaceRequest) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m OrderCancelReplaceRequest) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m OrderCancelReplaceRequest) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m OrderCancelReplaceRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m OrderCancelReplaceRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetComplianceID sets ComplianceID, Tag 376
func (m OrderCancelReplaceRequest) SetComplianceID(v string) {
	m.Set(field.NewComplianceID(v))
}

//SetSolicitedFlag sets SolicitedFlag, Tag 377
func (m OrderCancelReplaceRequest) SetSolicitedFlag(v bool) {
	m.Set(field.NewSolicitedFlag(v))
}

//SetNoTradingSessions sets NoTradingSessions, Tag 386
func (m OrderCancelReplaceRequest) SetNoTradingSessions(f NoTradingSessionsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDiscretionInst sets DiscretionInst, Tag 388
func (m OrderCancelReplaceRequest) SetDiscretionInst(v string) {
	m.Set(field.NewDiscretionInst(v))
}

//SetDiscretionOffsetValue sets DiscretionOffsetValue, Tag 389
func (m OrderCancelReplaceRequest) SetDiscretionOffsetValue(v float64) {
	m.Set(field.NewDiscretionOffsetValue(v))
}

//SetPriceType sets PriceType, Tag 423
func (m OrderCancelReplaceRequest) SetPriceType(v int) {
	m.Set(field.NewPriceType(v))
}

//SetGTBookingInst sets GTBookingInst, Tag 427
func (m OrderCancelReplaceRequest) SetGTBookingInst(v int) {
	m.Set(field.NewGTBookingInst(v))
}

//SetExpireDate sets ExpireDate, Tag 432
func (m OrderCancelReplaceRequest) SetExpireDate(v string) {
	m.Set(field.NewExpireDate(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m OrderCancelReplaceRequest) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m OrderCancelReplaceRequest) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m OrderCancelReplaceRequest) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m OrderCancelReplaceRequest) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetRoundingDirection sets RoundingDirection, Tag 468
func (m OrderCancelReplaceRequest) SetRoundingDirection(v string) {
	m.Set(field.NewRoundingDirection(v))
}

//SetRoundingModulus sets RoundingModulus, Tag 469
func (m OrderCancelReplaceRequest) SetRoundingModulus(v float64) {
	m.Set(field.NewRoundingModulus(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m OrderCancelReplaceRequest) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m OrderCancelReplaceRequest) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m OrderCancelReplaceRequest) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetCommCurrency sets CommCurrency, Tag 479
func (m OrderCancelReplaceRequest) SetCommCurrency(v string) {
	m.Set(field.NewCommCurrency(v))
}

//SetCancellationRights sets CancellationRights, Tag 480
func (m OrderCancelReplaceRequest) SetCancellationRights(v string) {
	m.Set(field.NewCancellationRights(v))
}

//SetMoneyLaunderingStatus sets MoneyLaunderingStatus, Tag 481
func (m OrderCancelReplaceRequest) SetMoneyLaunderingStatus(v string) {
	m.Set(field.NewMoneyLaunderingStatus(v))
}

//SetDesignation sets Designation, Tag 494
func (m OrderCancelReplaceRequest) SetDesignation(v string) {
	m.Set(field.NewDesignation(v))
}

//SetFundRenewWaiv sets FundRenewWaiv, Tag 497
func (m OrderCancelReplaceRequest) SetFundRenewWaiv(v string) {
	m.Set(field.NewFundRenewWaiv(v))
}

//SetRegistID sets RegistID, Tag 513
func (m OrderCancelReplaceRequest) SetRegistID(v string) {
	m.Set(field.NewRegistID(v))
}

//SetOrderPercent sets OrderPercent, Tag 516
func (m OrderCancelReplaceRequest) SetOrderPercent(v float64) {
	m.Set(field.NewOrderPercent(v))
}

//SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526
func (m OrderCancelReplaceRequest) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

//SetOrderCapacity sets OrderCapacity, Tag 528
func (m OrderCancelReplaceRequest) SetOrderCapacity(v string) {
	m.Set(field.NewOrderCapacity(v))
}

//SetOrderRestrictions sets OrderRestrictions, Tag 529
func (m OrderCancelReplaceRequest) SetOrderRestrictions(v string) {
	m.Set(field.NewOrderRestrictions(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m OrderCancelReplaceRequest) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m OrderCancelReplaceRequest) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetCashMargin sets CashMargin, Tag 544
func (m OrderCancelReplaceRequest) SetCashMargin(v string) {
	m.Set(field.NewCashMargin(v))
}

//SetAccountType sets AccountType, Tag 581
func (m OrderCancelReplaceRequest) SetAccountType(v int) {
	m.Set(field.NewAccountType(v))
}

//SetCustOrderCapacity sets CustOrderCapacity, Tag 582
func (m OrderCancelReplaceRequest) SetCustOrderCapacity(v int) {
	m.Set(field.NewCustOrderCapacity(v))
}

//SetClOrdLinkID sets ClOrdLinkID, Tag 583
func (m OrderCancelReplaceRequest) SetClOrdLinkID(v string) {
	m.Set(field.NewClOrdLinkID(v))
}

//SetOrigOrdModTime sets OrigOrdModTime, Tag 586
func (m OrderCancelReplaceRequest) SetOrigOrdModTime(v time.Time) {
	m.Set(field.NewOrigOrdModTime(v))
}

//SetDayBookingInst sets DayBookingInst, Tag 589
func (m OrderCancelReplaceRequest) SetDayBookingInst(v string) {
	m.Set(field.NewDayBookingInst(v))
}

//SetBookingUnit sets BookingUnit, Tag 590
func (m OrderCancelReplaceRequest) SetBookingUnit(v string) {
	m.Set(field.NewBookingUnit(v))
}

//SetPreallocMethod sets PreallocMethod, Tag 591
func (m OrderCancelReplaceRequest) SetPreallocMethod(v string) {
	m.Set(field.NewPreallocMethod(v))
}

//SetClearingFeeIndicator sets ClearingFeeIndicator, Tag 635
func (m OrderCancelReplaceRequest) SetClearingFeeIndicator(v string) {
	m.Set(field.NewClearingFeeIndicator(v))
}

//SetPrice2 sets Price2, Tag 640
func (m OrderCancelReplaceRequest) SetPrice2(v float64) {
	m.Set(field.NewPrice2(v))
}

//SetAcctIDSource sets AcctIDSource, Tag 660
func (m OrderCancelReplaceRequest) SetAcctIDSource(v int) {
	m.Set(field.NewAcctIDSource(v))
}

//SetBenchmarkPrice sets BenchmarkPrice, Tag 662
func (m OrderCancelReplaceRequest) SetBenchmarkPrice(v float64) {
	m.Set(field.NewBenchmarkPrice(v))
}

//SetBenchmarkPriceType sets BenchmarkPriceType, Tag 663
func (m OrderCancelReplaceRequest) SetBenchmarkPriceType(v int) {
	m.Set(field.NewBenchmarkPriceType(v))
}

//SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m OrderCancelReplaceRequest) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

//SetPool sets Pool, Tag 691
func (m OrderCancelReplaceRequest) SetPool(v string) {
	m.Set(field.NewPool(v))
}

//SetYieldRedemptionDate sets YieldRedemptionDate, Tag 696
func (m OrderCancelReplaceRequest) SetYieldRedemptionDate(v string) {
	m.Set(field.NewYieldRedemptionDate(v))
}

//SetYieldRedemptionPrice sets YieldRedemptionPrice, Tag 697
func (m OrderCancelReplaceRequest) SetYieldRedemptionPrice(v float64) {
	m.Set(field.NewYieldRedemptionPrice(v))
}

//SetYieldRedemptionPriceType sets YieldRedemptionPriceType, Tag 698
func (m OrderCancelReplaceRequest) SetYieldRedemptionPriceType(v int) {
	m.Set(field.NewYieldRedemptionPriceType(v))
}

//SetBenchmarkSecurityID sets BenchmarkSecurityID, Tag 699
func (m OrderCancelReplaceRequest) SetBenchmarkSecurityID(v string) {
	m.Set(field.NewBenchmarkSecurityID(v))
}

//SetYieldCalcDate sets YieldCalcDate, Tag 701
func (m OrderCancelReplaceRequest) SetYieldCalcDate(v string) {
	m.Set(field.NewYieldCalcDate(v))
}

//SetNoUnderlyings sets NoUnderlyings, Tag 711
func (m OrderCancelReplaceRequest) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

//SetBenchmarkSecurityIDSource sets BenchmarkSecurityIDSource, Tag 761
func (m OrderCancelReplaceRequest) SetBenchmarkSecurityIDSource(v string) {
	m.Set(field.NewBenchmarkSecurityIDSource(v))
}

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m OrderCancelReplaceRequest) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

//SetNoTrdRegTimestamps sets NoTrdRegTimestamps, Tag 768
func (m OrderCancelReplaceRequest) SetNoTrdRegTimestamps(f NoTrdRegTimestampsRepeatingGroup) {
	m.SetGroup(f)
}

//SetBookingType sets BookingType, Tag 775
func (m OrderCancelReplaceRequest) SetBookingType(v int) {
	m.Set(field.NewBookingType(v))
}

//SetTerminationType sets TerminationType, Tag 788
func (m OrderCancelReplaceRequest) SetTerminationType(v int) {
	m.Set(field.NewTerminationType(v))
}

//SetPegMoveType sets PegMoveType, Tag 835
func (m OrderCancelReplaceRequest) SetPegMoveType(v int) {
	m.Set(field.NewPegMoveType(v))
}

//SetPegOffsetType sets PegOffsetType, Tag 836
func (m OrderCancelReplaceRequest) SetPegOffsetType(v int) {
	m.Set(field.NewPegOffsetType(v))
}

//SetPegLimitType sets PegLimitType, Tag 837
func (m OrderCancelReplaceRequest) SetPegLimitType(v int) {
	m.Set(field.NewPegLimitType(v))
}

//SetPegRoundDirection sets PegRoundDirection, Tag 838
func (m OrderCancelReplaceRequest) SetPegRoundDirection(v int) {
	m.Set(field.NewPegRoundDirection(v))
}

//SetPegScope sets PegScope, Tag 840
func (m OrderCancelReplaceRequest) SetPegScope(v int) {
	m.Set(field.NewPegScope(v))
}

//SetDiscretionMoveType sets DiscretionMoveType, Tag 841
func (m OrderCancelReplaceRequest) SetDiscretionMoveType(v int) {
	m.Set(field.NewDiscretionMoveType(v))
}

//SetDiscretionOffsetType sets DiscretionOffsetType, Tag 842
func (m OrderCancelReplaceRequest) SetDiscretionOffsetType(v int) {
	m.Set(field.NewDiscretionOffsetType(v))
}

//SetDiscretionLimitType sets DiscretionLimitType, Tag 843
func (m OrderCancelReplaceRequest) SetDiscretionLimitType(v int) {
	m.Set(field.NewDiscretionLimitType(v))
}

//SetDiscretionRoundDirection sets DiscretionRoundDirection, Tag 844
func (m OrderCancelReplaceRequest) SetDiscretionRoundDirection(v int) {
	m.Set(field.NewDiscretionRoundDirection(v))
}

//SetDiscretionScope sets DiscretionScope, Tag 846
func (m OrderCancelReplaceRequest) SetDiscretionScope(v int) {
	m.Set(field.NewDiscretionScope(v))
}

//SetTargetStrategy sets TargetStrategy, Tag 847
func (m OrderCancelReplaceRequest) SetTargetStrategy(v int) {
	m.Set(field.NewTargetStrategy(v))
}

//SetTargetStrategyParameters sets TargetStrategyParameters, Tag 848
func (m OrderCancelReplaceRequest) SetTargetStrategyParameters(v string) {
	m.Set(field.NewTargetStrategyParameters(v))
}

//SetParticipationRate sets ParticipationRate, Tag 849
func (m OrderCancelReplaceRequest) SetParticipationRate(v float64) {
	m.Set(field.NewParticipationRate(v))
}

//SetQtyType sets QtyType, Tag 854
func (m OrderCancelReplaceRequest) SetQtyType(v int) {
	m.Set(field.NewQtyType(v))
}

//SetNoEvents sets NoEvents, Tag 864
func (m OrderCancelReplaceRequest) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDatedDate sets DatedDate, Tag 873
func (m OrderCancelReplaceRequest) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

//SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m OrderCancelReplaceRequest) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

//SetCPProgram sets CPProgram, Tag 875
func (m OrderCancelReplaceRequest) SetCPProgram(v int) {
	m.Set(field.NewCPProgram(v))
}

//SetCPRegType sets CPRegType, Tag 876
func (m OrderCancelReplaceRequest) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

//SetMarginRatio sets MarginRatio, Tag 898
func (m OrderCancelReplaceRequest) SetMarginRatio(v float64) {
	m.Set(field.NewMarginRatio(v))
}

//SetAgreementDesc sets AgreementDesc, Tag 913
func (m OrderCancelReplaceRequest) SetAgreementDesc(v string) {
	m.Set(field.NewAgreementDesc(v))
}

//SetAgreementID sets AgreementID, Tag 914
func (m OrderCancelReplaceRequest) SetAgreementID(v string) {
	m.Set(field.NewAgreementID(v))
}

//SetAgreementDate sets AgreementDate, Tag 915
func (m OrderCancelReplaceRequest) SetAgreementDate(v string) {
	m.Set(field.NewAgreementDate(v))
}

//SetStartDate sets StartDate, Tag 916
func (m OrderCancelReplaceRequest) SetStartDate(v string) {
	m.Set(field.NewStartDate(v))
}

//SetEndDate sets EndDate, Tag 917
func (m OrderCancelReplaceRequest) SetEndDate(v string) {
	m.Set(field.NewEndDate(v))
}

//SetAgreementCurrency sets AgreementCurrency, Tag 918
func (m OrderCancelReplaceRequest) SetAgreementCurrency(v string) {
	m.Set(field.NewAgreementCurrency(v))
}

//SetDeliveryType sets DeliveryType, Tag 919
func (m OrderCancelReplaceRequest) SetDeliveryType(v int) {
	m.Set(field.NewDeliveryType(v))
}

//SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m OrderCancelReplaceRequest) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

//SetNoStrategyParameters sets NoStrategyParameters, Tag 957
func (m OrderCancelReplaceRequest) SetNoStrategyParameters(f NoStrategyParametersRepeatingGroup) {
	m.SetGroup(f)
}

//SetSecurityStatus sets SecurityStatus, Tag 965
func (m OrderCancelReplaceRequest) SetSecurityStatus(v string) {
	m.Set(field.NewSecurityStatus(v))
}

//SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966
func (m OrderCancelReplaceRequest) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

//SetStrikeMultiplier sets StrikeMultiplier, Tag 967
func (m OrderCancelReplaceRequest) SetStrikeMultiplier(v float64) {
	m.Set(field.NewStrikeMultiplier(v))
}

//SetStrikeValue sets StrikeValue, Tag 968
func (m OrderCancelReplaceRequest) SetStrikeValue(v float64) {
	m.Set(field.NewStrikeValue(v))
}

//SetMinPriceIncrement sets MinPriceIncrement, Tag 969
func (m OrderCancelReplaceRequest) SetMinPriceIncrement(v float64) {
	m.Set(field.NewMinPriceIncrement(v))
}

//SetPositionLimit sets PositionLimit, Tag 970
func (m OrderCancelReplaceRequest) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

//SetNTPositionLimit sets NTPositionLimit, Tag 971
func (m OrderCancelReplaceRequest) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

//SetUnitOfMeasure sets UnitOfMeasure, Tag 996
func (m OrderCancelReplaceRequest) SetUnitOfMeasure(v string) {
	m.Set(field.NewUnitOfMeasure(v))
}

//SetTimeUnit sets TimeUnit, Tag 997
func (m OrderCancelReplaceRequest) SetTimeUnit(v string) {
	m.Set(field.NewTimeUnit(v))
}

//SetNoInstrumentParties sets NoInstrumentParties, Tag 1018
func (m OrderCancelReplaceRequest) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetManualOrderIndicator sets ManualOrderIndicator, Tag 1028
func (m OrderCancelReplaceRequest) SetManualOrderIndicator(v bool) {
	m.Set(field.NewManualOrderIndicator(v))
}

//SetCustDirectedOrder sets CustDirectedOrder, Tag 1029
func (m OrderCancelReplaceRequest) SetCustDirectedOrder(v bool) {
	m.Set(field.NewCustDirectedOrder(v))
}

//SetReceivedDeptID sets ReceivedDeptID, Tag 1030
func (m OrderCancelReplaceRequest) SetReceivedDeptID(v string) {
	m.Set(field.NewReceivedDeptID(v))
}

//SetCustOrderHandlingInst sets CustOrderHandlingInst, Tag 1031
func (m OrderCancelReplaceRequest) SetCustOrderHandlingInst(v string) {
	m.Set(field.NewCustOrderHandlingInst(v))
}

//SetOrderHandlingInstSource sets OrderHandlingInstSource, Tag 1032
func (m OrderCancelReplaceRequest) SetOrderHandlingInstSource(v int) {
	m.Set(field.NewOrderHandlingInstSource(v))
}

//SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049
func (m OrderCancelReplaceRequest) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

//SetMaturityTime sets MaturityTime, Tag 1079
func (m OrderCancelReplaceRequest) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

//SetSecondaryDisplayQty sets SecondaryDisplayQty, Tag 1082
func (m OrderCancelReplaceRequest) SetSecondaryDisplayQty(v float64) {
	m.Set(field.NewSecondaryDisplayQty(v))
}

//SetDisplayWhen sets DisplayWhen, Tag 1083
func (m OrderCancelReplaceRequest) SetDisplayWhen(v string) {
	m.Set(field.NewDisplayWhen(v))
}

//SetDisplayMethod sets DisplayMethod, Tag 1084
func (m OrderCancelReplaceRequest) SetDisplayMethod(v string) {
	m.Set(field.NewDisplayMethod(v))
}

//SetDisplayLowQty sets DisplayLowQty, Tag 1085
func (m OrderCancelReplaceRequest) SetDisplayLowQty(v float64) {
	m.Set(field.NewDisplayLowQty(v))
}

//SetDisplayHighQty sets DisplayHighQty, Tag 1086
func (m OrderCancelReplaceRequest) SetDisplayHighQty(v float64) {
	m.Set(field.NewDisplayHighQty(v))
}

//SetDisplayMinIncr sets DisplayMinIncr, Tag 1087
func (m OrderCancelReplaceRequest) SetDisplayMinIncr(v float64) {
	m.Set(field.NewDisplayMinIncr(v))
}

//SetRefreshQty sets RefreshQty, Tag 1088
func (m OrderCancelReplaceRequest) SetRefreshQty(v float64) {
	m.Set(field.NewRefreshQty(v))
}

//SetMatchIncrement sets MatchIncrement, Tag 1089
func (m OrderCancelReplaceRequest) SetMatchIncrement(v float64) {
	m.Set(field.NewMatchIncrement(v))
}

//SetMaxPriceLevels sets MaxPriceLevels, Tag 1090
func (m OrderCancelReplaceRequest) SetMaxPriceLevels(v int) {
	m.Set(field.NewMaxPriceLevels(v))
}

//SetPreTradeAnonymity sets PreTradeAnonymity, Tag 1091
func (m OrderCancelReplaceRequest) SetPreTradeAnonymity(v bool) {
	m.Set(field.NewPreTradeAnonymity(v))
}

//SetPriceProtectionScope sets PriceProtectionScope, Tag 1092
func (m OrderCancelReplaceRequest) SetPriceProtectionScope(v string) {
	m.Set(field.NewPriceProtectionScope(v))
}

//SetPegPriceType sets PegPriceType, Tag 1094
func (m OrderCancelReplaceRequest) SetPegPriceType(v int) {
	m.Set(field.NewPegPriceType(v))
}

//SetPegSecurityIDSource sets PegSecurityIDSource, Tag 1096
func (m OrderCancelReplaceRequest) SetPegSecurityIDSource(v string) {
	m.Set(field.NewPegSecurityIDSource(v))
}

//SetPegSecurityID sets PegSecurityID, Tag 1097
func (m OrderCancelReplaceRequest) SetPegSecurityID(v string) {
	m.Set(field.NewPegSecurityID(v))
}

//SetPegSymbol sets PegSymbol, Tag 1098
func (m OrderCancelReplaceRequest) SetPegSymbol(v string) {
	m.Set(field.NewPegSymbol(v))
}

//SetPegSecurityDesc sets PegSecurityDesc, Tag 1099
func (m OrderCancelReplaceRequest) SetPegSecurityDesc(v string) {
	m.Set(field.NewPegSecurityDesc(v))
}

//SetTriggerType sets TriggerType, Tag 1100
func (m OrderCancelReplaceRequest) SetTriggerType(v string) {
	m.Set(field.NewTriggerType(v))
}

//SetTriggerAction sets TriggerAction, Tag 1101
func (m OrderCancelReplaceRequest) SetTriggerAction(v string) {
	m.Set(field.NewTriggerAction(v))
}

//SetTriggerPrice sets TriggerPrice, Tag 1102
func (m OrderCancelReplaceRequest) SetTriggerPrice(v float64) {
	m.Set(field.NewTriggerPrice(v))
}

//SetTriggerSymbol sets TriggerSymbol, Tag 1103
func (m OrderCancelReplaceRequest) SetTriggerSymbol(v string) {
	m.Set(field.NewTriggerSymbol(v))
}

//SetTriggerSecurityID sets TriggerSecurityID, Tag 1104
func (m OrderCancelReplaceRequest) SetTriggerSecurityID(v string) {
	m.Set(field.NewTriggerSecurityID(v))
}

//SetTriggerSecurityIDSource sets TriggerSecurityIDSource, Tag 1105
func (m OrderCancelReplaceRequest) SetTriggerSecurityIDSource(v string) {
	m.Set(field.NewTriggerSecurityIDSource(v))
}

//SetTriggerSecurityDesc sets TriggerSecurityDesc, Tag 1106
func (m OrderCancelReplaceRequest) SetTriggerSecurityDesc(v string) {
	m.Set(field.NewTriggerSecurityDesc(v))
}

//SetTriggerPriceType sets TriggerPriceType, Tag 1107
func (m OrderCancelReplaceRequest) SetTriggerPriceType(v string) {
	m.Set(field.NewTriggerPriceType(v))
}

//SetTriggerPriceTypeScope sets TriggerPriceTypeScope, Tag 1108
func (m OrderCancelReplaceRequest) SetTriggerPriceTypeScope(v string) {
	m.Set(field.NewTriggerPriceTypeScope(v))
}

//SetTriggerPriceDirection sets TriggerPriceDirection, Tag 1109
func (m OrderCancelReplaceRequest) SetTriggerPriceDirection(v string) {
	m.Set(field.NewTriggerPriceDirection(v))
}

//SetTriggerNewPrice sets TriggerNewPrice, Tag 1110
func (m OrderCancelReplaceRequest) SetTriggerNewPrice(v float64) {
	m.Set(field.NewTriggerNewPrice(v))
}

//SetTriggerOrderType sets TriggerOrderType, Tag 1111
func (m OrderCancelReplaceRequest) SetTriggerOrderType(v string) {
	m.Set(field.NewTriggerOrderType(v))
}

//SetTriggerNewQty sets TriggerNewQty, Tag 1112
func (m OrderCancelReplaceRequest) SetTriggerNewQty(v float64) {
	m.Set(field.NewTriggerNewQty(v))
}

//SetTriggerTradingSessionID sets TriggerTradingSessionID, Tag 1113
func (m OrderCancelReplaceRequest) SetTriggerTradingSessionID(v string) {
	m.Set(field.NewTriggerTradingSessionID(v))
}

//SetTriggerTradingSessionSubID sets TriggerTradingSessionSubID, Tag 1114
func (m OrderCancelReplaceRequest) SetTriggerTradingSessionSubID(v string) {
	m.Set(field.NewTriggerTradingSessionSubID(v))
}

//SetExDestinationIDSource sets ExDestinationIDSource, Tag 1133
func (m OrderCancelReplaceRequest) SetExDestinationIDSource(v string) {
	m.Set(field.NewExDestinationIDSource(v))
}

//SetDisplayQty sets DisplayQty, Tag 1138
func (m OrderCancelReplaceRequest) SetDisplayQty(v float64) {
	m.Set(field.NewDisplayQty(v))
}

//SetMinPriceIncrementAmount sets MinPriceIncrementAmount, Tag 1146
func (m OrderCancelReplaceRequest) SetMinPriceIncrementAmount(v float64) {
	m.Set(field.NewMinPriceIncrementAmount(v))
}

//SetUnitOfMeasureQty sets UnitOfMeasureQty, Tag 1147
func (m OrderCancelReplaceRequest) SetUnitOfMeasureQty(v float64) {
	m.Set(field.NewUnitOfMeasureQty(v))
}

//SetSecurityGroup sets SecurityGroup, Tag 1151
func (m OrderCancelReplaceRequest) SetSecurityGroup(v string) {
	m.Set(field.NewSecurityGroup(v))
}

//SetSecurityXMLLen sets SecurityXMLLen, Tag 1184
func (m OrderCancelReplaceRequest) SetSecurityXMLLen(v int) {
	m.Set(field.NewSecurityXMLLen(v))
}

//SetSecurityXML sets SecurityXML, Tag 1185
func (m OrderCancelReplaceRequest) SetSecurityXML(v string) {
	m.Set(field.NewSecurityXML(v))
}

//SetSecurityXMLSchema sets SecurityXMLSchema, Tag 1186
func (m OrderCancelReplaceRequest) SetSecurityXMLSchema(v string) {
	m.Set(field.NewSecurityXMLSchema(v))
}

//SetPriceUnitOfMeasure sets PriceUnitOfMeasure, Tag 1191
func (m OrderCancelReplaceRequest) SetPriceUnitOfMeasure(v string) {
	m.Set(field.NewPriceUnitOfMeasure(v))
}

//SetPriceUnitOfMeasureQty sets PriceUnitOfMeasureQty, Tag 1192
func (m OrderCancelReplaceRequest) SetPriceUnitOfMeasureQty(v float64) {
	m.Set(field.NewPriceUnitOfMeasureQty(v))
}

//SetSettlMethod sets SettlMethod, Tag 1193
func (m OrderCancelReplaceRequest) SetSettlMethod(v string) {
	m.Set(field.NewSettlMethod(v))
}

//SetExerciseStyle sets ExerciseStyle, Tag 1194
func (m OrderCancelReplaceRequest) SetExerciseStyle(v int) {
	m.Set(field.NewExerciseStyle(v))
}

//SetOptPayoutAmount sets OptPayoutAmount, Tag 1195
func (m OrderCancelReplaceRequest) SetOptPayoutAmount(v float64) {
	m.Set(field.NewOptPayoutAmount(v))
}

//SetPriceQuoteMethod sets PriceQuoteMethod, Tag 1196
func (m OrderCancelReplaceRequest) SetPriceQuoteMethod(v string) {
	m.Set(field.NewPriceQuoteMethod(v))
}

//SetValuationMethod sets ValuationMethod, Tag 1197
func (m OrderCancelReplaceRequest) SetValuationMethod(v string) {
	m.Set(field.NewValuationMethod(v))
}

//SetListMethod sets ListMethod, Tag 1198
func (m OrderCancelReplaceRequest) SetListMethod(v int) {
	m.Set(field.NewListMethod(v))
}

//SetCapPrice sets CapPrice, Tag 1199
func (m OrderCancelReplaceRequest) SetCapPrice(v float64) {
	m.Set(field.NewCapPrice(v))
}

//SetFloorPrice sets FloorPrice, Tag 1200
func (m OrderCancelReplaceRequest) SetFloorPrice(v float64) {
	m.Set(field.NewFloorPrice(v))
}

//SetProductComplex sets ProductComplex, Tag 1227
func (m OrderCancelReplaceRequest) SetProductComplex(v string) {
	m.Set(field.NewProductComplex(v))
}

//SetFlexProductEligibilityIndicator sets FlexProductEligibilityIndicator, Tag 1242
func (m OrderCancelReplaceRequest) SetFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewFlexProductEligibilityIndicator(v))
}

//SetFlexibleIndicator sets FlexibleIndicator, Tag 1244
func (m OrderCancelReplaceRequest) SetFlexibleIndicator(v bool) {
	m.Set(field.NewFlexibleIndicator(v))
}

//SetContractMultiplierUnit sets ContractMultiplierUnit, Tag 1435
func (m OrderCancelReplaceRequest) SetContractMultiplierUnit(v int) {
	m.Set(field.NewContractMultiplierUnit(v))
}

//SetFlowScheduleType sets FlowScheduleType, Tag 1439
func (m OrderCancelReplaceRequest) SetFlowScheduleType(v int) {
	m.Set(field.NewFlowScheduleType(v))
}

//SetRestructuringType sets RestructuringType, Tag 1449
func (m OrderCancelReplaceRequest) SetRestructuringType(v string) {
	m.Set(field.NewRestructuringType(v))
}

//SetSeniority sets Seniority, Tag 1450
func (m OrderCancelReplaceRequest) SetSeniority(v string) {
	m.Set(field.NewSeniority(v))
}

//SetNotionalPercentageOutstanding sets NotionalPercentageOutstanding, Tag 1451
func (m OrderCancelReplaceRequest) SetNotionalPercentageOutstanding(v float64) {
	m.Set(field.NewNotionalPercentageOutstanding(v))
}

//SetOriginalNotionalPercentageOutstanding sets OriginalNotionalPercentageOutstanding, Tag 1452
func (m OrderCancelReplaceRequest) SetOriginalNotionalPercentageOutstanding(v float64) {
	m.Set(field.NewOriginalNotionalPercentageOutstanding(v))
}

//SetAttachmentPoint sets AttachmentPoint, Tag 1457
func (m OrderCancelReplaceRequest) SetAttachmentPoint(v float64) {
	m.Set(field.NewAttachmentPoint(v))
}

//SetDetachmentPoint sets DetachmentPoint, Tag 1458
func (m OrderCancelReplaceRequest) SetDetachmentPoint(v float64) {
	m.Set(field.NewDetachmentPoint(v))
}

//SetStrikePriceDeterminationMethod sets StrikePriceDeterminationMethod, Tag 1478
func (m OrderCancelReplaceRequest) SetStrikePriceDeterminationMethod(v int) {
	m.Set(field.NewStrikePriceDeterminationMethod(v))
}

//SetStrikePriceBoundaryMethod sets StrikePriceBoundaryMethod, Tag 1479
func (m OrderCancelReplaceRequest) SetStrikePriceBoundaryMethod(v int) {
	m.Set(field.NewStrikePriceBoundaryMethod(v))
}

//SetStrikePriceBoundaryPrecision sets StrikePriceBoundaryPrecision, Tag 1480
func (m OrderCancelReplaceRequest) SetStrikePriceBoundaryPrecision(v float64) {
	m.Set(field.NewStrikePriceBoundaryPrecision(v))
}

//SetUnderlyingPriceDeterminationMethod sets UnderlyingPriceDeterminationMethod, Tag 1481
func (m OrderCancelReplaceRequest) SetUnderlyingPriceDeterminationMethod(v int) {
	m.Set(field.NewUnderlyingPriceDeterminationMethod(v))
}

//SetOptPayoutType sets OptPayoutType, Tag 1482
func (m OrderCancelReplaceRequest) SetOptPayoutType(v int) {
	m.Set(field.NewOptPayoutType(v))
}

//SetNoComplexEvents sets NoComplexEvents, Tag 1483
func (m OrderCancelReplaceRequest) SetNoComplexEvents(f NoComplexEventsRepeatingGroup) {
	m.SetGroup(f)
}

//GetAccount gets Account, Tag 1
func (m OrderCancelReplaceRequest) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m OrderCancelReplaceRequest) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommission gets Commission, Tag 12
func (m OrderCancelReplaceRequest) GetCommission() (f field.CommissionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommType gets CommType, Tag 13
func (m OrderCancelReplaceRequest) GetCommType() (f field.CommTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m OrderCancelReplaceRequest) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecInst gets ExecInst, Tag 18
func (m OrderCancelReplaceRequest) GetExecInst() (f field.ExecInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetHandlInst gets HandlInst, Tag 21
func (m OrderCancelReplaceRequest) GetHandlInst() (f field.HandlInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m OrderCancelReplaceRequest) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m OrderCancelReplaceRequest) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m OrderCancelReplaceRequest) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdType gets OrdType, Tag 40
func (m OrderCancelReplaceRequest) GetOrdType() (f field.OrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigClOrdID gets OrigClOrdID, Tag 41
func (m OrderCancelReplaceRequest) GetOrigClOrdID() (f field.OrigClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m OrderCancelReplaceRequest) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m OrderCancelReplaceRequest) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m OrderCancelReplaceRequest) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m OrderCancelReplaceRequest) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m OrderCancelReplaceRequest) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m OrderCancelReplaceRequest) GetTimeInForce() (f field.TimeInForceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m OrderCancelReplaceRequest) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlType gets SettlType, Tag 63
func (m OrderCancelReplaceRequest) GetSettlType() (f field.SettlTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlDate gets SettlDate, Tag 64
func (m OrderCancelReplaceRequest) GetSettlDate() (f field.SettlDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m OrderCancelReplaceRequest) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListID gets ListID, Tag 66
func (m OrderCancelReplaceRequest) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocID gets AllocID, Tag 70
func (m OrderCancelReplaceRequest) GetAllocID() (f field.AllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m OrderCancelReplaceRequest) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionEffect gets PositionEffect, Tag 77
func (m OrderCancelReplaceRequest) GetPositionEffect() (f field.PositionEffectField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoAllocs gets NoAllocs, Tag 78
func (m OrderCancelReplaceRequest) GetNoAllocs() (NoAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetStopPx gets StopPx, Tag 99
func (m OrderCancelReplaceRequest) GetStopPx() (f field.StopPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExDestination gets ExDestination, Tag 100
func (m OrderCancelReplaceRequest) GetExDestination() (f field.ExDestinationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m OrderCancelReplaceRequest) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m OrderCancelReplaceRequest) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinQty gets MinQty, Tag 110
func (m OrderCancelReplaceRequest) GetMinQty() (f field.MinQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxFloor gets MaxFloor, Tag 111
func (m OrderCancelReplaceRequest) GetMaxFloor() (f field.MaxFloorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocateReqd gets LocateReqd, Tag 114
func (m OrderCancelReplaceRequest) GetLocateReqd() (f field.LocateReqdField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m OrderCancelReplaceRequest) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetForexReq gets ForexReq, Tag 121
func (m OrderCancelReplaceRequest) GetForexReq() (f field.ForexReqField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m OrderCancelReplaceRequest) GetExpireTime() (f field.ExpireTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m OrderCancelReplaceRequest) GetCashOrderQty() (f field.CashOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m OrderCancelReplaceRequest) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEffectiveTime gets EffectiveTime, Tag 168
func (m OrderCancelReplaceRequest) GetEffectiveTime() (f field.EffectiveTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty2 gets OrderQty2, Tag 192
func (m OrderCancelReplaceRequest) GetOrderQty2() (f field.OrderQty2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlDate2 gets SettlDate2, Tag 193
func (m OrderCancelReplaceRequest) GetSettlDate2() (f field.SettlDate2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m OrderCancelReplaceRequest) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m OrderCancelReplaceRequest) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m OrderCancelReplaceRequest) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCoveredOrUncovered gets CoveredOrUncovered, Tag 203
func (m OrderCancelReplaceRequest) GetCoveredOrUncovered() (f field.CoveredOrUncoveredField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m OrderCancelReplaceRequest) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m OrderCancelReplaceRequest) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxShow gets MaxShow, Tag 210
func (m OrderCancelReplaceRequest) GetMaxShow() (f field.MaxShowField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegOffsetValue gets PegOffsetValue, Tag 211
func (m OrderCancelReplaceRequest) GetPegOffsetValue() (f field.PegOffsetValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSpread gets Spread, Tag 218
func (m OrderCancelReplaceRequest) GetSpread() (f field.SpreadField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveCurrency gets BenchmarkCurveCurrency, Tag 220
func (m OrderCancelReplaceRequest) GetBenchmarkCurveCurrency() (f field.BenchmarkCurveCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveName gets BenchmarkCurveName, Tag 221
func (m OrderCancelReplaceRequest) GetBenchmarkCurveName() (f field.BenchmarkCurveNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurvePoint gets BenchmarkCurvePoint, Tag 222
func (m OrderCancelReplaceRequest) GetBenchmarkCurvePoint() (f field.BenchmarkCurvePointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m OrderCancelReplaceRequest) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m OrderCancelReplaceRequest) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m OrderCancelReplaceRequest) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m OrderCancelReplaceRequest) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m OrderCancelReplaceRequest) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m OrderCancelReplaceRequest) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeOriginationDate gets TradeOriginationDate, Tag 229
func (m OrderCancelReplaceRequest) GetTradeOriginationDate() (f field.TradeOriginationDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m OrderCancelReplaceRequest) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldType gets YieldType, Tag 235
func (m OrderCancelReplaceRequest) GetYieldType() (f field.YieldTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYield gets Yield, Tag 236
func (m OrderCancelReplaceRequest) GetYield() (f field.YieldField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m OrderCancelReplaceRequest) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m OrderCancelReplaceRequest) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m OrderCancelReplaceRequest) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m OrderCancelReplaceRequest) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m OrderCancelReplaceRequest) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m OrderCancelReplaceRequest) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m OrderCancelReplaceRequest) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m OrderCancelReplaceRequest) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m OrderCancelReplaceRequest) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplianceID gets ComplianceID, Tag 376
func (m OrderCancelReplaceRequest) GetComplianceID() (f field.ComplianceIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSolicitedFlag gets SolicitedFlag, Tag 377
func (m OrderCancelReplaceRequest) GetSolicitedFlag() (f field.SolicitedFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoTradingSessions gets NoTradingSessions, Tag 386
func (m OrderCancelReplaceRequest) GetNoTradingSessions() (NoTradingSessionsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTradingSessionsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDiscretionInst gets DiscretionInst, Tag 388
func (m OrderCancelReplaceRequest) GetDiscretionInst() (f field.DiscretionInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionOffsetValue gets DiscretionOffsetValue, Tag 389
func (m OrderCancelReplaceRequest) GetDiscretionOffsetValue() (f field.DiscretionOffsetValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceType gets PriceType, Tag 423
func (m OrderCancelReplaceRequest) GetPriceType() (f field.PriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetGTBookingInst gets GTBookingInst, Tag 427
func (m OrderCancelReplaceRequest) GetGTBookingInst() (f field.GTBookingInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireDate gets ExpireDate, Tag 432
func (m OrderCancelReplaceRequest) GetExpireDate() (f field.ExpireDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m OrderCancelReplaceRequest) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m OrderCancelReplaceRequest) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m OrderCancelReplaceRequest) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m OrderCancelReplaceRequest) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingDirection gets RoundingDirection, Tag 468
func (m OrderCancelReplaceRequest) GetRoundingDirection() (f field.RoundingDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingModulus gets RoundingModulus, Tag 469
func (m OrderCancelReplaceRequest) GetRoundingModulus() (f field.RoundingModulusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m OrderCancelReplaceRequest) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m OrderCancelReplaceRequest) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m OrderCancelReplaceRequest) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommCurrency gets CommCurrency, Tag 479
func (m OrderCancelReplaceRequest) GetCommCurrency() (f field.CommCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCancellationRights gets CancellationRights, Tag 480
func (m OrderCancelReplaceRequest) GetCancellationRights() (f field.CancellationRightsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMoneyLaunderingStatus gets MoneyLaunderingStatus, Tag 481
func (m OrderCancelReplaceRequest) GetMoneyLaunderingStatus() (f field.MoneyLaunderingStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDesignation gets Designation, Tag 494
func (m OrderCancelReplaceRequest) GetDesignation() (f field.DesignationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFundRenewWaiv gets FundRenewWaiv, Tag 497
func (m OrderCancelReplaceRequest) GetFundRenewWaiv() (f field.FundRenewWaivField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRegistID gets RegistID, Tag 513
func (m OrderCancelReplaceRequest) GetRegistID() (f field.RegistIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderPercent gets OrderPercent, Tag 516
func (m OrderCancelReplaceRequest) GetOrderPercent() (f field.OrderPercentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526
func (m OrderCancelReplaceRequest) GetSecondaryClOrdID() (f field.SecondaryClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderCapacity gets OrderCapacity, Tag 528
func (m OrderCancelReplaceRequest) GetOrderCapacity() (f field.OrderCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderRestrictions gets OrderRestrictions, Tag 529
func (m OrderCancelReplaceRequest) GetOrderRestrictions() (f field.OrderRestrictionsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m OrderCancelReplaceRequest) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m OrderCancelReplaceRequest) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashMargin gets CashMargin, Tag 544
func (m OrderCancelReplaceRequest) GetCashMargin() (f field.CashMarginField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccountType gets AccountType, Tag 581
func (m OrderCancelReplaceRequest) GetAccountType() (f field.AccountTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCustOrderCapacity gets CustOrderCapacity, Tag 582
func (m OrderCancelReplaceRequest) GetCustOrderCapacity() (f field.CustOrderCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdLinkID gets ClOrdLinkID, Tag 583
func (m OrderCancelReplaceRequest) GetClOrdLinkID() (f field.ClOrdLinkIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigOrdModTime gets OrigOrdModTime, Tag 586
func (m OrderCancelReplaceRequest) GetOrigOrdModTime() (f field.OrigOrdModTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDayBookingInst gets DayBookingInst, Tag 589
func (m OrderCancelReplaceRequest) GetDayBookingInst() (f field.DayBookingInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBookingUnit gets BookingUnit, Tag 590
func (m OrderCancelReplaceRequest) GetBookingUnit() (f field.BookingUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPreallocMethod gets PreallocMethod, Tag 591
func (m OrderCancelReplaceRequest) GetPreallocMethod() (f field.PreallocMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClearingFeeIndicator gets ClearingFeeIndicator, Tag 635
func (m OrderCancelReplaceRequest) GetClearingFeeIndicator() (f field.ClearingFeeIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice2 gets Price2, Tag 640
func (m OrderCancelReplaceRequest) GetPrice2() (f field.Price2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAcctIDSource gets AcctIDSource, Tag 660
func (m OrderCancelReplaceRequest) GetAcctIDSource() (f field.AcctIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkPrice gets BenchmarkPrice, Tag 662
func (m OrderCancelReplaceRequest) GetBenchmarkPrice() (f field.BenchmarkPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkPriceType gets BenchmarkPriceType, Tag 663
func (m OrderCancelReplaceRequest) GetBenchmarkPriceType() (f field.BenchmarkPriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m OrderCancelReplaceRequest) GetContractSettlMonth() (f field.ContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPool gets Pool, Tag 691
func (m OrderCancelReplaceRequest) GetPool() (f field.PoolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldRedemptionDate gets YieldRedemptionDate, Tag 696
func (m OrderCancelReplaceRequest) GetYieldRedemptionDate() (f field.YieldRedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldRedemptionPrice gets YieldRedemptionPrice, Tag 697
func (m OrderCancelReplaceRequest) GetYieldRedemptionPrice() (f field.YieldRedemptionPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldRedemptionPriceType gets YieldRedemptionPriceType, Tag 698
func (m OrderCancelReplaceRequest) GetYieldRedemptionPriceType() (f field.YieldRedemptionPriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkSecurityID gets BenchmarkSecurityID, Tag 699
func (m OrderCancelReplaceRequest) GetBenchmarkSecurityID() (f field.BenchmarkSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldCalcDate gets YieldCalcDate, Tag 701
func (m OrderCancelReplaceRequest) GetYieldCalcDate() (f field.YieldCalcDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyings gets NoUnderlyings, Tag 711
func (m OrderCancelReplaceRequest) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetBenchmarkSecurityIDSource gets BenchmarkSecurityIDSource, Tag 761
func (m OrderCancelReplaceRequest) GetBenchmarkSecurityIDSource() (f field.BenchmarkSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m OrderCancelReplaceRequest) GetSecuritySubType() (f field.SecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoTrdRegTimestamps gets NoTrdRegTimestamps, Tag 768
func (m OrderCancelReplaceRequest) GetNoTrdRegTimestamps() (NoTrdRegTimestampsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTrdRegTimestampsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetBookingType gets BookingType, Tag 775
func (m OrderCancelReplaceRequest) GetBookingType() (f field.BookingTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTerminationType gets TerminationType, Tag 788
func (m OrderCancelReplaceRequest) GetTerminationType() (f field.TerminationTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegMoveType gets PegMoveType, Tag 835
func (m OrderCancelReplaceRequest) GetPegMoveType() (f field.PegMoveTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegOffsetType gets PegOffsetType, Tag 836
func (m OrderCancelReplaceRequest) GetPegOffsetType() (f field.PegOffsetTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegLimitType gets PegLimitType, Tag 837
func (m OrderCancelReplaceRequest) GetPegLimitType() (f field.PegLimitTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegRoundDirection gets PegRoundDirection, Tag 838
func (m OrderCancelReplaceRequest) GetPegRoundDirection() (f field.PegRoundDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegScope gets PegScope, Tag 840
func (m OrderCancelReplaceRequest) GetPegScope() (f field.PegScopeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionMoveType gets DiscretionMoveType, Tag 841
func (m OrderCancelReplaceRequest) GetDiscretionMoveType() (f field.DiscretionMoveTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionOffsetType gets DiscretionOffsetType, Tag 842
func (m OrderCancelReplaceRequest) GetDiscretionOffsetType() (f field.DiscretionOffsetTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionLimitType gets DiscretionLimitType, Tag 843
func (m OrderCancelReplaceRequest) GetDiscretionLimitType() (f field.DiscretionLimitTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionRoundDirection gets DiscretionRoundDirection, Tag 844
func (m OrderCancelReplaceRequest) GetDiscretionRoundDirection() (f field.DiscretionRoundDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionScope gets DiscretionScope, Tag 846
func (m OrderCancelReplaceRequest) GetDiscretionScope() (f field.DiscretionScopeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTargetStrategy gets TargetStrategy, Tag 847
func (m OrderCancelReplaceRequest) GetTargetStrategy() (f field.TargetStrategyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTargetStrategyParameters gets TargetStrategyParameters, Tag 848
func (m OrderCancelReplaceRequest) GetTargetStrategyParameters() (f field.TargetStrategyParametersField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetParticipationRate gets ParticipationRate, Tag 849
func (m OrderCancelReplaceRequest) GetParticipationRate() (f field.ParticipationRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQtyType gets QtyType, Tag 854
func (m OrderCancelReplaceRequest) GetQtyType() (f field.QtyTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoEvents gets NoEvents, Tag 864
func (m OrderCancelReplaceRequest) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDatedDate gets DatedDate, Tag 873
func (m OrderCancelReplaceRequest) GetDatedDate() (f field.DatedDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m OrderCancelReplaceRequest) GetInterestAccrualDate() (f field.InterestAccrualDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPProgram gets CPProgram, Tag 875
func (m OrderCancelReplaceRequest) GetCPProgram() (f field.CPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPRegType gets CPRegType, Tag 876
func (m OrderCancelReplaceRequest) GetCPRegType() (f field.CPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMarginRatio gets MarginRatio, Tag 898
func (m OrderCancelReplaceRequest) GetMarginRatio() (f field.MarginRatioField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementDesc gets AgreementDesc, Tag 913
func (m OrderCancelReplaceRequest) GetAgreementDesc() (f field.AgreementDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementID gets AgreementID, Tag 914
func (m OrderCancelReplaceRequest) GetAgreementID() (f field.AgreementIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementDate gets AgreementDate, Tag 915
func (m OrderCancelReplaceRequest) GetAgreementDate() (f field.AgreementDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStartDate gets StartDate, Tag 916
func (m OrderCancelReplaceRequest) GetStartDate() (f field.StartDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEndDate gets EndDate, Tag 917
func (m OrderCancelReplaceRequest) GetEndDate() (f field.EndDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementCurrency gets AgreementCurrency, Tag 918
func (m OrderCancelReplaceRequest) GetAgreementCurrency() (f field.AgreementCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDeliveryType gets DeliveryType, Tag 919
func (m OrderCancelReplaceRequest) GetDeliveryType() (f field.DeliveryTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m OrderCancelReplaceRequest) GetStrikeCurrency() (f field.StrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoStrategyParameters gets NoStrategyParameters, Tag 957
func (m OrderCancelReplaceRequest) GetNoStrategyParameters() (NoStrategyParametersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStrategyParametersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSecurityStatus gets SecurityStatus, Tag 965
func (m OrderCancelReplaceRequest) GetSecurityStatus() (f field.SecurityStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966
func (m OrderCancelReplaceRequest) GetSettleOnOpenFlag() (f field.SettleOnOpenFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeMultiplier gets StrikeMultiplier, Tag 967
func (m OrderCancelReplaceRequest) GetStrikeMultiplier() (f field.StrikeMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeValue gets StrikeValue, Tag 968
func (m OrderCancelReplaceRequest) GetStrikeValue() (f field.StrikeValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrement gets MinPriceIncrement, Tag 969
func (m OrderCancelReplaceRequest) GetMinPriceIncrement() (f field.MinPriceIncrementField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionLimit gets PositionLimit, Tag 970
func (m OrderCancelReplaceRequest) GetPositionLimit() (f field.PositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNTPositionLimit gets NTPositionLimit, Tag 971
func (m OrderCancelReplaceRequest) GetNTPositionLimit() (f field.NTPositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnitOfMeasure gets UnitOfMeasure, Tag 996
func (m OrderCancelReplaceRequest) GetUnitOfMeasure() (f field.UnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeUnit gets TimeUnit, Tag 997
func (m OrderCancelReplaceRequest) GetTimeUnit() (f field.TimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoInstrumentParties gets NoInstrumentParties, Tag 1018
func (m OrderCancelReplaceRequest) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetManualOrderIndicator gets ManualOrderIndicator, Tag 1028
func (m OrderCancelReplaceRequest) GetManualOrderIndicator() (f field.ManualOrderIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCustDirectedOrder gets CustDirectedOrder, Tag 1029
func (m OrderCancelReplaceRequest) GetCustDirectedOrder() (f field.CustDirectedOrderField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetReceivedDeptID gets ReceivedDeptID, Tag 1030
func (m OrderCancelReplaceRequest) GetReceivedDeptID() (f field.ReceivedDeptIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCustOrderHandlingInst gets CustOrderHandlingInst, Tag 1031
func (m OrderCancelReplaceRequest) GetCustOrderHandlingInst() (f field.CustOrderHandlingInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderHandlingInstSource gets OrderHandlingInstSource, Tag 1032
func (m OrderCancelReplaceRequest) GetOrderHandlingInstSource() (f field.OrderHandlingInstSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049
func (m OrderCancelReplaceRequest) GetInstrmtAssignmentMethod() (f field.InstrmtAssignmentMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityTime gets MaturityTime, Tag 1079
func (m OrderCancelReplaceRequest) GetMaturityTime() (f field.MaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryDisplayQty gets SecondaryDisplayQty, Tag 1082
func (m OrderCancelReplaceRequest) GetSecondaryDisplayQty() (f field.SecondaryDisplayQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDisplayWhen gets DisplayWhen, Tag 1083
func (m OrderCancelReplaceRequest) GetDisplayWhen() (f field.DisplayWhenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDisplayMethod gets DisplayMethod, Tag 1084
func (m OrderCancelReplaceRequest) GetDisplayMethod() (f field.DisplayMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDisplayLowQty gets DisplayLowQty, Tag 1085
func (m OrderCancelReplaceRequest) GetDisplayLowQty() (f field.DisplayLowQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDisplayHighQty gets DisplayHighQty, Tag 1086
func (m OrderCancelReplaceRequest) GetDisplayHighQty() (f field.DisplayHighQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDisplayMinIncr gets DisplayMinIncr, Tag 1087
func (m OrderCancelReplaceRequest) GetDisplayMinIncr() (f field.DisplayMinIncrField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRefreshQty gets RefreshQty, Tag 1088
func (m OrderCancelReplaceRequest) GetRefreshQty() (f field.RefreshQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMatchIncrement gets MatchIncrement, Tag 1089
func (m OrderCancelReplaceRequest) GetMatchIncrement() (f field.MatchIncrementField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxPriceLevels gets MaxPriceLevels, Tag 1090
func (m OrderCancelReplaceRequest) GetMaxPriceLevels() (f field.MaxPriceLevelsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPreTradeAnonymity gets PreTradeAnonymity, Tag 1091
func (m OrderCancelReplaceRequest) GetPreTradeAnonymity() (f field.PreTradeAnonymityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceProtectionScope gets PriceProtectionScope, Tag 1092
func (m OrderCancelReplaceRequest) GetPriceProtectionScope() (f field.PriceProtectionScopeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegPriceType gets PegPriceType, Tag 1094
func (m OrderCancelReplaceRequest) GetPegPriceType() (f field.PegPriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegSecurityIDSource gets PegSecurityIDSource, Tag 1096
func (m OrderCancelReplaceRequest) GetPegSecurityIDSource() (f field.PegSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegSecurityID gets PegSecurityID, Tag 1097
func (m OrderCancelReplaceRequest) GetPegSecurityID() (f field.PegSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegSymbol gets PegSymbol, Tag 1098
func (m OrderCancelReplaceRequest) GetPegSymbol() (f field.PegSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegSecurityDesc gets PegSecurityDesc, Tag 1099
func (m OrderCancelReplaceRequest) GetPegSecurityDesc() (f field.PegSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerType gets TriggerType, Tag 1100
func (m OrderCancelReplaceRequest) GetTriggerType() (f field.TriggerTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerAction gets TriggerAction, Tag 1101
func (m OrderCancelReplaceRequest) GetTriggerAction() (f field.TriggerActionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerPrice gets TriggerPrice, Tag 1102
func (m OrderCancelReplaceRequest) GetTriggerPrice() (f field.TriggerPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerSymbol gets TriggerSymbol, Tag 1103
func (m OrderCancelReplaceRequest) GetTriggerSymbol() (f field.TriggerSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerSecurityID gets TriggerSecurityID, Tag 1104
func (m OrderCancelReplaceRequest) GetTriggerSecurityID() (f field.TriggerSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerSecurityIDSource gets TriggerSecurityIDSource, Tag 1105
func (m OrderCancelReplaceRequest) GetTriggerSecurityIDSource() (f field.TriggerSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerSecurityDesc gets TriggerSecurityDesc, Tag 1106
func (m OrderCancelReplaceRequest) GetTriggerSecurityDesc() (f field.TriggerSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerPriceType gets TriggerPriceType, Tag 1107
func (m OrderCancelReplaceRequest) GetTriggerPriceType() (f field.TriggerPriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerPriceTypeScope gets TriggerPriceTypeScope, Tag 1108
func (m OrderCancelReplaceRequest) GetTriggerPriceTypeScope() (f field.TriggerPriceTypeScopeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerPriceDirection gets TriggerPriceDirection, Tag 1109
func (m OrderCancelReplaceRequest) GetTriggerPriceDirection() (f field.TriggerPriceDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerNewPrice gets TriggerNewPrice, Tag 1110
func (m OrderCancelReplaceRequest) GetTriggerNewPrice() (f field.TriggerNewPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerOrderType gets TriggerOrderType, Tag 1111
func (m OrderCancelReplaceRequest) GetTriggerOrderType() (f field.TriggerOrderTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerNewQty gets TriggerNewQty, Tag 1112
func (m OrderCancelReplaceRequest) GetTriggerNewQty() (f field.TriggerNewQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerTradingSessionID gets TriggerTradingSessionID, Tag 1113
func (m OrderCancelReplaceRequest) GetTriggerTradingSessionID() (f field.TriggerTradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerTradingSessionSubID gets TriggerTradingSessionSubID, Tag 1114
func (m OrderCancelReplaceRequest) GetTriggerTradingSessionSubID() (f field.TriggerTradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExDestinationIDSource gets ExDestinationIDSource, Tag 1133
func (m OrderCancelReplaceRequest) GetExDestinationIDSource() (f field.ExDestinationIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDisplayQty gets DisplayQty, Tag 1138
func (m OrderCancelReplaceRequest) GetDisplayQty() (f field.DisplayQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrementAmount gets MinPriceIncrementAmount, Tag 1146
func (m OrderCancelReplaceRequest) GetMinPriceIncrementAmount() (f field.MinPriceIncrementAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnitOfMeasureQty gets UnitOfMeasureQty, Tag 1147
func (m OrderCancelReplaceRequest) GetUnitOfMeasureQty() (f field.UnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityGroup gets SecurityGroup, Tag 1151
func (m OrderCancelReplaceRequest) GetSecurityGroup() (f field.SecurityGroupField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXMLLen gets SecurityXMLLen, Tag 1184
func (m OrderCancelReplaceRequest) GetSecurityXMLLen() (f field.SecurityXMLLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXML gets SecurityXML, Tag 1185
func (m OrderCancelReplaceRequest) GetSecurityXML() (f field.SecurityXMLField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXMLSchema gets SecurityXMLSchema, Tag 1186
func (m OrderCancelReplaceRequest) GetSecurityXMLSchema() (f field.SecurityXMLSchemaField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceUnitOfMeasure gets PriceUnitOfMeasure, Tag 1191
func (m OrderCancelReplaceRequest) GetPriceUnitOfMeasure() (f field.PriceUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceUnitOfMeasureQty gets PriceUnitOfMeasureQty, Tag 1192
func (m OrderCancelReplaceRequest) GetPriceUnitOfMeasureQty() (f field.PriceUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlMethod gets SettlMethod, Tag 1193
func (m OrderCancelReplaceRequest) GetSettlMethod() (f field.SettlMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExerciseStyle gets ExerciseStyle, Tag 1194
func (m OrderCancelReplaceRequest) GetExerciseStyle() (f field.ExerciseStyleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptPayoutAmount gets OptPayoutAmount, Tag 1195
func (m OrderCancelReplaceRequest) GetOptPayoutAmount() (f field.OptPayoutAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceQuoteMethod gets PriceQuoteMethod, Tag 1196
func (m OrderCancelReplaceRequest) GetPriceQuoteMethod() (f field.PriceQuoteMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetValuationMethod gets ValuationMethod, Tag 1197
func (m OrderCancelReplaceRequest) GetValuationMethod() (f field.ValuationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListMethod gets ListMethod, Tag 1198
func (m OrderCancelReplaceRequest) GetListMethod() (f field.ListMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCapPrice gets CapPrice, Tag 1199
func (m OrderCancelReplaceRequest) GetCapPrice() (f field.CapPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFloorPrice gets FloorPrice, Tag 1200
func (m OrderCancelReplaceRequest) GetFloorPrice() (f field.FloorPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProductComplex gets ProductComplex, Tag 1227
func (m OrderCancelReplaceRequest) GetProductComplex() (f field.ProductComplexField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlexProductEligibilityIndicator gets FlexProductEligibilityIndicator, Tag 1242
func (m OrderCancelReplaceRequest) GetFlexProductEligibilityIndicator() (f field.FlexProductEligibilityIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlexibleIndicator gets FlexibleIndicator, Tag 1244
func (m OrderCancelReplaceRequest) GetFlexibleIndicator() (f field.FlexibleIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplierUnit gets ContractMultiplierUnit, Tag 1435
func (m OrderCancelReplaceRequest) GetContractMultiplierUnit() (f field.ContractMultiplierUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlowScheduleType gets FlowScheduleType, Tag 1439
func (m OrderCancelReplaceRequest) GetFlowScheduleType() (f field.FlowScheduleTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRestructuringType gets RestructuringType, Tag 1449
func (m OrderCancelReplaceRequest) GetRestructuringType() (f field.RestructuringTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSeniority gets Seniority, Tag 1450
func (m OrderCancelReplaceRequest) GetSeniority() (f field.SeniorityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNotionalPercentageOutstanding gets NotionalPercentageOutstanding, Tag 1451
func (m OrderCancelReplaceRequest) GetNotionalPercentageOutstanding() (f field.NotionalPercentageOutstandingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOriginalNotionalPercentageOutstanding gets OriginalNotionalPercentageOutstanding, Tag 1452
func (m OrderCancelReplaceRequest) GetOriginalNotionalPercentageOutstanding() (f field.OriginalNotionalPercentageOutstandingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAttachmentPoint gets AttachmentPoint, Tag 1457
func (m OrderCancelReplaceRequest) GetAttachmentPoint() (f field.AttachmentPointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDetachmentPoint gets DetachmentPoint, Tag 1458
func (m OrderCancelReplaceRequest) GetDetachmentPoint() (f field.DetachmentPointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePriceDeterminationMethod gets StrikePriceDeterminationMethod, Tag 1478
func (m OrderCancelReplaceRequest) GetStrikePriceDeterminationMethod() (f field.StrikePriceDeterminationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePriceBoundaryMethod gets StrikePriceBoundaryMethod, Tag 1479
func (m OrderCancelReplaceRequest) GetStrikePriceBoundaryMethod() (f field.StrikePriceBoundaryMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePriceBoundaryPrecision gets StrikePriceBoundaryPrecision, Tag 1480
func (m OrderCancelReplaceRequest) GetStrikePriceBoundaryPrecision() (f field.StrikePriceBoundaryPrecisionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPriceDeterminationMethod gets UnderlyingPriceDeterminationMethod, Tag 1481
func (m OrderCancelReplaceRequest) GetUnderlyingPriceDeterminationMethod() (f field.UnderlyingPriceDeterminationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptPayoutType gets OptPayoutType, Tag 1482
func (m OrderCancelReplaceRequest) GetOptPayoutType() (f field.OptPayoutTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoComplexEvents gets NoComplexEvents, Tag 1483
func (m OrderCancelReplaceRequest) GetNoComplexEvents() (NoComplexEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasAccount returns true if Account is present, Tag 1
func (m OrderCancelReplaceRequest) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m OrderCancelReplaceRequest) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasCommission returns true if Commission is present, Tag 12
func (m OrderCancelReplaceRequest) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m OrderCancelReplaceRequest) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m OrderCancelReplaceRequest) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasExecInst returns true if ExecInst is present, Tag 18
func (m OrderCancelReplaceRequest) HasExecInst() bool {
	return m.Has(tag.ExecInst)
}

//HasHandlInst returns true if HandlInst is present, Tag 21
func (m OrderCancelReplaceRequest) HasHandlInst() bool {
	return m.Has(tag.HandlInst)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m OrderCancelReplaceRequest) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m OrderCancelReplaceRequest) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m OrderCancelReplaceRequest) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m OrderCancelReplaceRequest) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasOrigClOrdID returns true if OrigClOrdID is present, Tag 41
func (m OrderCancelReplaceRequest) HasOrigClOrdID() bool {
	return m.Has(tag.OrigClOrdID)
}

//HasPrice returns true if Price is present, Tag 44
func (m OrderCancelReplaceRequest) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m OrderCancelReplaceRequest) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m OrderCancelReplaceRequest) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m OrderCancelReplaceRequest) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m OrderCancelReplaceRequest) HasText() bool {
	return m.Has(tag.Text)
}

//HasTimeInForce returns true if TimeInForce is present, Tag 59
func (m OrderCancelReplaceRequest) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m OrderCancelReplaceRequest) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSettlType returns true if SettlType is present, Tag 63
func (m OrderCancelReplaceRequest) HasSettlType() bool {
	return m.Has(tag.SettlType)
}

//HasSettlDate returns true if SettlDate is present, Tag 64
func (m OrderCancelReplaceRequest) HasSettlDate() bool {
	return m.Has(tag.SettlDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m OrderCancelReplaceRequest) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasListID returns true if ListID is present, Tag 66
func (m OrderCancelReplaceRequest) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasAllocID returns true if AllocID is present, Tag 70
func (m OrderCancelReplaceRequest) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m OrderCancelReplaceRequest) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasPositionEffect returns true if PositionEffect is present, Tag 77
func (m OrderCancelReplaceRequest) HasPositionEffect() bool {
	return m.Has(tag.PositionEffect)
}

//HasNoAllocs returns true if NoAllocs is present, Tag 78
func (m OrderCancelReplaceRequest) HasNoAllocs() bool {
	return m.Has(tag.NoAllocs)
}

//HasStopPx returns true if StopPx is present, Tag 99
func (m OrderCancelReplaceRequest) HasStopPx() bool {
	return m.Has(tag.StopPx)
}

//HasExDestination returns true if ExDestination is present, Tag 100
func (m OrderCancelReplaceRequest) HasExDestination() bool {
	return m.Has(tag.ExDestination)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m OrderCancelReplaceRequest) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m OrderCancelReplaceRequest) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasMinQty returns true if MinQty is present, Tag 110
func (m OrderCancelReplaceRequest) HasMinQty() bool {
	return m.Has(tag.MinQty)
}

//HasMaxFloor returns true if MaxFloor is present, Tag 111
func (m OrderCancelReplaceRequest) HasMaxFloor() bool {
	return m.Has(tag.MaxFloor)
}

//HasLocateReqd returns true if LocateReqd is present, Tag 114
func (m OrderCancelReplaceRequest) HasLocateReqd() bool {
	return m.Has(tag.LocateReqd)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m OrderCancelReplaceRequest) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasForexReq returns true if ForexReq is present, Tag 121
func (m OrderCancelReplaceRequest) HasForexReq() bool {
	return m.Has(tag.ForexReq)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m OrderCancelReplaceRequest) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m OrderCancelReplaceRequest) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m OrderCancelReplaceRequest) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasEffectiveTime returns true if EffectiveTime is present, Tag 168
func (m OrderCancelReplaceRequest) HasEffectiveTime() bool {
	return m.Has(tag.EffectiveTime)
}

//HasOrderQty2 returns true if OrderQty2 is present, Tag 192
func (m OrderCancelReplaceRequest) HasOrderQty2() bool {
	return m.Has(tag.OrderQty2)
}

//HasSettlDate2 returns true if SettlDate2 is present, Tag 193
func (m OrderCancelReplaceRequest) HasSettlDate2() bool {
	return m.Has(tag.SettlDate2)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m OrderCancelReplaceRequest) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m OrderCancelReplaceRequest) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m OrderCancelReplaceRequest) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasCoveredOrUncovered returns true if CoveredOrUncovered is present, Tag 203
func (m OrderCancelReplaceRequest) HasCoveredOrUncovered() bool {
	return m.Has(tag.CoveredOrUncovered)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m OrderCancelReplaceRequest) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m OrderCancelReplaceRequest) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasMaxShow returns true if MaxShow is present, Tag 210
func (m OrderCancelReplaceRequest) HasMaxShow() bool {
	return m.Has(tag.MaxShow)
}

//HasPegOffsetValue returns true if PegOffsetValue is present, Tag 211
func (m OrderCancelReplaceRequest) HasPegOffsetValue() bool {
	return m.Has(tag.PegOffsetValue)
}

//HasSpread returns true if Spread is present, Tag 218
func (m OrderCancelReplaceRequest) HasSpread() bool {
	return m.Has(tag.Spread)
}

//HasBenchmarkCurveCurrency returns true if BenchmarkCurveCurrency is present, Tag 220
func (m OrderCancelReplaceRequest) HasBenchmarkCurveCurrency() bool {
	return m.Has(tag.BenchmarkCurveCurrency)
}

//HasBenchmarkCurveName returns true if BenchmarkCurveName is present, Tag 221
func (m OrderCancelReplaceRequest) HasBenchmarkCurveName() bool {
	return m.Has(tag.BenchmarkCurveName)
}

//HasBenchmarkCurvePoint returns true if BenchmarkCurvePoint is present, Tag 222
func (m OrderCancelReplaceRequest) HasBenchmarkCurvePoint() bool {
	return m.Has(tag.BenchmarkCurvePoint)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m OrderCancelReplaceRequest) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m OrderCancelReplaceRequest) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m OrderCancelReplaceRequest) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m OrderCancelReplaceRequest) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m OrderCancelReplaceRequest) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m OrderCancelReplaceRequest) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasTradeOriginationDate returns true if TradeOriginationDate is present, Tag 229
func (m OrderCancelReplaceRequest) HasTradeOriginationDate() bool {
	return m.Has(tag.TradeOriginationDate)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m OrderCancelReplaceRequest) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasYieldType returns true if YieldType is present, Tag 235
func (m OrderCancelReplaceRequest) HasYieldType() bool {
	return m.Has(tag.YieldType)
}

//HasYield returns true if Yield is present, Tag 236
func (m OrderCancelReplaceRequest) HasYield() bool {
	return m.Has(tag.Yield)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m OrderCancelReplaceRequest) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m OrderCancelReplaceRequest) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m OrderCancelReplaceRequest) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m OrderCancelReplaceRequest) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m OrderCancelReplaceRequest) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m OrderCancelReplaceRequest) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m OrderCancelReplaceRequest) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m OrderCancelReplaceRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m OrderCancelReplaceRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasComplianceID returns true if ComplianceID is present, Tag 376
func (m OrderCancelReplaceRequest) HasComplianceID() bool {
	return m.Has(tag.ComplianceID)
}

//HasSolicitedFlag returns true if SolicitedFlag is present, Tag 377
func (m OrderCancelReplaceRequest) HasSolicitedFlag() bool {
	return m.Has(tag.SolicitedFlag)
}

//HasNoTradingSessions returns true if NoTradingSessions is present, Tag 386
func (m OrderCancelReplaceRequest) HasNoTradingSessions() bool {
	return m.Has(tag.NoTradingSessions)
}

//HasDiscretionInst returns true if DiscretionInst is present, Tag 388
func (m OrderCancelReplaceRequest) HasDiscretionInst() bool {
	return m.Has(tag.DiscretionInst)
}

//HasDiscretionOffsetValue returns true if DiscretionOffsetValue is present, Tag 389
func (m OrderCancelReplaceRequest) HasDiscretionOffsetValue() bool {
	return m.Has(tag.DiscretionOffsetValue)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m OrderCancelReplaceRequest) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasGTBookingInst returns true if GTBookingInst is present, Tag 427
func (m OrderCancelReplaceRequest) HasGTBookingInst() bool {
	return m.Has(tag.GTBookingInst)
}

//HasExpireDate returns true if ExpireDate is present, Tag 432
func (m OrderCancelReplaceRequest) HasExpireDate() bool {
	return m.Has(tag.ExpireDate)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m OrderCancelReplaceRequest) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m OrderCancelReplaceRequest) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m OrderCancelReplaceRequest) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m OrderCancelReplaceRequest) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasRoundingDirection returns true if RoundingDirection is present, Tag 468
func (m OrderCancelReplaceRequest) HasRoundingDirection() bool {
	return m.Has(tag.RoundingDirection)
}

//HasRoundingModulus returns true if RoundingModulus is present, Tag 469
func (m OrderCancelReplaceRequest) HasRoundingModulus() bool {
	return m.Has(tag.RoundingModulus)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m OrderCancelReplaceRequest) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m OrderCancelReplaceRequest) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m OrderCancelReplaceRequest) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasCommCurrency returns true if CommCurrency is present, Tag 479
func (m OrderCancelReplaceRequest) HasCommCurrency() bool {
	return m.Has(tag.CommCurrency)
}

//HasCancellationRights returns true if CancellationRights is present, Tag 480
func (m OrderCancelReplaceRequest) HasCancellationRights() bool {
	return m.Has(tag.CancellationRights)
}

//HasMoneyLaunderingStatus returns true if MoneyLaunderingStatus is present, Tag 481
func (m OrderCancelReplaceRequest) HasMoneyLaunderingStatus() bool {
	return m.Has(tag.MoneyLaunderingStatus)
}

//HasDesignation returns true if Designation is present, Tag 494
func (m OrderCancelReplaceRequest) HasDesignation() bool {
	return m.Has(tag.Designation)
}

//HasFundRenewWaiv returns true if FundRenewWaiv is present, Tag 497
func (m OrderCancelReplaceRequest) HasFundRenewWaiv() bool {
	return m.Has(tag.FundRenewWaiv)
}

//HasRegistID returns true if RegistID is present, Tag 513
func (m OrderCancelReplaceRequest) HasRegistID() bool {
	return m.Has(tag.RegistID)
}

//HasOrderPercent returns true if OrderPercent is present, Tag 516
func (m OrderCancelReplaceRequest) HasOrderPercent() bool {
	return m.Has(tag.OrderPercent)
}

//HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526
func (m OrderCancelReplaceRequest) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

//HasOrderCapacity returns true if OrderCapacity is present, Tag 528
func (m OrderCancelReplaceRequest) HasOrderCapacity() bool {
	return m.Has(tag.OrderCapacity)
}

//HasOrderRestrictions returns true if OrderRestrictions is present, Tag 529
func (m OrderCancelReplaceRequest) HasOrderRestrictions() bool {
	return m.Has(tag.OrderRestrictions)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m OrderCancelReplaceRequest) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m OrderCancelReplaceRequest) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasCashMargin returns true if CashMargin is present, Tag 544
func (m OrderCancelReplaceRequest) HasCashMargin() bool {
	return m.Has(tag.CashMargin)
}

//HasAccountType returns true if AccountType is present, Tag 581
func (m OrderCancelReplaceRequest) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

//HasCustOrderCapacity returns true if CustOrderCapacity is present, Tag 582
func (m OrderCancelReplaceRequest) HasCustOrderCapacity() bool {
	return m.Has(tag.CustOrderCapacity)
}

//HasClOrdLinkID returns true if ClOrdLinkID is present, Tag 583
func (m OrderCancelReplaceRequest) HasClOrdLinkID() bool {
	return m.Has(tag.ClOrdLinkID)
}

//HasOrigOrdModTime returns true if OrigOrdModTime is present, Tag 586
func (m OrderCancelReplaceRequest) HasOrigOrdModTime() bool {
	return m.Has(tag.OrigOrdModTime)
}

//HasDayBookingInst returns true if DayBookingInst is present, Tag 589
func (m OrderCancelReplaceRequest) HasDayBookingInst() bool {
	return m.Has(tag.DayBookingInst)
}

//HasBookingUnit returns true if BookingUnit is present, Tag 590
func (m OrderCancelReplaceRequest) HasBookingUnit() bool {
	return m.Has(tag.BookingUnit)
}

//HasPreallocMethod returns true if PreallocMethod is present, Tag 591
func (m OrderCancelReplaceRequest) HasPreallocMethod() bool {
	return m.Has(tag.PreallocMethod)
}

//HasClearingFeeIndicator returns true if ClearingFeeIndicator is present, Tag 635
func (m OrderCancelReplaceRequest) HasClearingFeeIndicator() bool {
	return m.Has(tag.ClearingFeeIndicator)
}

//HasPrice2 returns true if Price2 is present, Tag 640
func (m OrderCancelReplaceRequest) HasPrice2() bool {
	return m.Has(tag.Price2)
}

//HasAcctIDSource returns true if AcctIDSource is present, Tag 660
func (m OrderCancelReplaceRequest) HasAcctIDSource() bool {
	return m.Has(tag.AcctIDSource)
}

//HasBenchmarkPrice returns true if BenchmarkPrice is present, Tag 662
func (m OrderCancelReplaceRequest) HasBenchmarkPrice() bool {
	return m.Has(tag.BenchmarkPrice)
}

//HasBenchmarkPriceType returns true if BenchmarkPriceType is present, Tag 663
func (m OrderCancelReplaceRequest) HasBenchmarkPriceType() bool {
	return m.Has(tag.BenchmarkPriceType)
}

//HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m OrderCancelReplaceRequest) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

//HasPool returns true if Pool is present, Tag 691
func (m OrderCancelReplaceRequest) HasPool() bool {
	return m.Has(tag.Pool)
}

//HasYieldRedemptionDate returns true if YieldRedemptionDate is present, Tag 696
func (m OrderCancelReplaceRequest) HasYieldRedemptionDate() bool {
	return m.Has(tag.YieldRedemptionDate)
}

//HasYieldRedemptionPrice returns true if YieldRedemptionPrice is present, Tag 697
func (m OrderCancelReplaceRequest) HasYieldRedemptionPrice() bool {
	return m.Has(tag.YieldRedemptionPrice)
}

//HasYieldRedemptionPriceType returns true if YieldRedemptionPriceType is present, Tag 698
func (m OrderCancelReplaceRequest) HasYieldRedemptionPriceType() bool {
	return m.Has(tag.YieldRedemptionPriceType)
}

//HasBenchmarkSecurityID returns true if BenchmarkSecurityID is present, Tag 699
func (m OrderCancelReplaceRequest) HasBenchmarkSecurityID() bool {
	return m.Has(tag.BenchmarkSecurityID)
}

//HasYieldCalcDate returns true if YieldCalcDate is present, Tag 701
func (m OrderCancelReplaceRequest) HasYieldCalcDate() bool {
	return m.Has(tag.YieldCalcDate)
}

//HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711
func (m OrderCancelReplaceRequest) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

//HasBenchmarkSecurityIDSource returns true if BenchmarkSecurityIDSource is present, Tag 761
func (m OrderCancelReplaceRequest) HasBenchmarkSecurityIDSource() bool {
	return m.Has(tag.BenchmarkSecurityIDSource)
}

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m OrderCancelReplaceRequest) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

//HasNoTrdRegTimestamps returns true if NoTrdRegTimestamps is present, Tag 768
func (m OrderCancelReplaceRequest) HasNoTrdRegTimestamps() bool {
	return m.Has(tag.NoTrdRegTimestamps)
}

//HasBookingType returns true if BookingType is present, Tag 775
func (m OrderCancelReplaceRequest) HasBookingType() bool {
	return m.Has(tag.BookingType)
}

//HasTerminationType returns true if TerminationType is present, Tag 788
func (m OrderCancelReplaceRequest) HasTerminationType() bool {
	return m.Has(tag.TerminationType)
}

//HasPegMoveType returns true if PegMoveType is present, Tag 835
func (m OrderCancelReplaceRequest) HasPegMoveType() bool {
	return m.Has(tag.PegMoveType)
}

//HasPegOffsetType returns true if PegOffsetType is present, Tag 836
func (m OrderCancelReplaceRequest) HasPegOffsetType() bool {
	return m.Has(tag.PegOffsetType)
}

//HasPegLimitType returns true if PegLimitType is present, Tag 837
func (m OrderCancelReplaceRequest) HasPegLimitType() bool {
	return m.Has(tag.PegLimitType)
}

//HasPegRoundDirection returns true if PegRoundDirection is present, Tag 838
func (m OrderCancelReplaceRequest) HasPegRoundDirection() bool {
	return m.Has(tag.PegRoundDirection)
}

//HasPegScope returns true if PegScope is present, Tag 840
func (m OrderCancelReplaceRequest) HasPegScope() bool {
	return m.Has(tag.PegScope)
}

//HasDiscretionMoveType returns true if DiscretionMoveType is present, Tag 841
func (m OrderCancelReplaceRequest) HasDiscretionMoveType() bool {
	return m.Has(tag.DiscretionMoveType)
}

//HasDiscretionOffsetType returns true if DiscretionOffsetType is present, Tag 842
func (m OrderCancelReplaceRequest) HasDiscretionOffsetType() bool {
	return m.Has(tag.DiscretionOffsetType)
}

//HasDiscretionLimitType returns true if DiscretionLimitType is present, Tag 843
func (m OrderCancelReplaceRequest) HasDiscretionLimitType() bool {
	return m.Has(tag.DiscretionLimitType)
}

//HasDiscretionRoundDirection returns true if DiscretionRoundDirection is present, Tag 844
func (m OrderCancelReplaceRequest) HasDiscretionRoundDirection() bool {
	return m.Has(tag.DiscretionRoundDirection)
}

//HasDiscretionScope returns true if DiscretionScope is present, Tag 846
func (m OrderCancelReplaceRequest) HasDiscretionScope() bool {
	return m.Has(tag.DiscretionScope)
}

//HasTargetStrategy returns true if TargetStrategy is present, Tag 847
func (m OrderCancelReplaceRequest) HasTargetStrategy() bool {
	return m.Has(tag.TargetStrategy)
}

//HasTargetStrategyParameters returns true if TargetStrategyParameters is present, Tag 848
func (m OrderCancelReplaceRequest) HasTargetStrategyParameters() bool {
	return m.Has(tag.TargetStrategyParameters)
}

//HasParticipationRate returns true if ParticipationRate is present, Tag 849
func (m OrderCancelReplaceRequest) HasParticipationRate() bool {
	return m.Has(tag.ParticipationRate)
}

//HasQtyType returns true if QtyType is present, Tag 854
func (m OrderCancelReplaceRequest) HasQtyType() bool {
	return m.Has(tag.QtyType)
}

//HasNoEvents returns true if NoEvents is present, Tag 864
func (m OrderCancelReplaceRequest) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

//HasDatedDate returns true if DatedDate is present, Tag 873
func (m OrderCancelReplaceRequest) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

//HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m OrderCancelReplaceRequest) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

//HasCPProgram returns true if CPProgram is present, Tag 875
func (m OrderCancelReplaceRequest) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

//HasCPRegType returns true if CPRegType is present, Tag 876
func (m OrderCancelReplaceRequest) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

//HasMarginRatio returns true if MarginRatio is present, Tag 898
func (m OrderCancelReplaceRequest) HasMarginRatio() bool {
	return m.Has(tag.MarginRatio)
}

//HasAgreementDesc returns true if AgreementDesc is present, Tag 913
func (m OrderCancelReplaceRequest) HasAgreementDesc() bool {
	return m.Has(tag.AgreementDesc)
}

//HasAgreementID returns true if AgreementID is present, Tag 914
func (m OrderCancelReplaceRequest) HasAgreementID() bool {
	return m.Has(tag.AgreementID)
}

//HasAgreementDate returns true if AgreementDate is present, Tag 915
func (m OrderCancelReplaceRequest) HasAgreementDate() bool {
	return m.Has(tag.AgreementDate)
}

//HasStartDate returns true if StartDate is present, Tag 916
func (m OrderCancelReplaceRequest) HasStartDate() bool {
	return m.Has(tag.StartDate)
}

//HasEndDate returns true if EndDate is present, Tag 917
func (m OrderCancelReplaceRequest) HasEndDate() bool {
	return m.Has(tag.EndDate)
}

//HasAgreementCurrency returns true if AgreementCurrency is present, Tag 918
func (m OrderCancelReplaceRequest) HasAgreementCurrency() bool {
	return m.Has(tag.AgreementCurrency)
}

//HasDeliveryType returns true if DeliveryType is present, Tag 919
func (m OrderCancelReplaceRequest) HasDeliveryType() bool {
	return m.Has(tag.DeliveryType)
}

//HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m OrderCancelReplaceRequest) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

//HasNoStrategyParameters returns true if NoStrategyParameters is present, Tag 957
func (m OrderCancelReplaceRequest) HasNoStrategyParameters() bool {
	return m.Has(tag.NoStrategyParameters)
}

//HasSecurityStatus returns true if SecurityStatus is present, Tag 965
func (m OrderCancelReplaceRequest) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

//HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966
func (m OrderCancelReplaceRequest) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

//HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967
func (m OrderCancelReplaceRequest) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

//HasStrikeValue returns true if StrikeValue is present, Tag 968
func (m OrderCancelReplaceRequest) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

//HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969
func (m OrderCancelReplaceRequest) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

//HasPositionLimit returns true if PositionLimit is present, Tag 970
func (m OrderCancelReplaceRequest) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

//HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971
func (m OrderCancelReplaceRequest) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

//HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996
func (m OrderCancelReplaceRequest) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

//HasTimeUnit returns true if TimeUnit is present, Tag 997
func (m OrderCancelReplaceRequest) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

//HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018
func (m OrderCancelReplaceRequest) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

//HasManualOrderIndicator returns true if ManualOrderIndicator is present, Tag 1028
func (m OrderCancelReplaceRequest) HasManualOrderIndicator() bool {
	return m.Has(tag.ManualOrderIndicator)
}

//HasCustDirectedOrder returns true if CustDirectedOrder is present, Tag 1029
func (m OrderCancelReplaceRequest) HasCustDirectedOrder() bool {
	return m.Has(tag.CustDirectedOrder)
}

//HasReceivedDeptID returns true if ReceivedDeptID is present, Tag 1030
func (m OrderCancelReplaceRequest) HasReceivedDeptID() bool {
	return m.Has(tag.ReceivedDeptID)
}

//HasCustOrderHandlingInst returns true if CustOrderHandlingInst is present, Tag 1031
func (m OrderCancelReplaceRequest) HasCustOrderHandlingInst() bool {
	return m.Has(tag.CustOrderHandlingInst)
}

//HasOrderHandlingInstSource returns true if OrderHandlingInstSource is present, Tag 1032
func (m OrderCancelReplaceRequest) HasOrderHandlingInstSource() bool {
	return m.Has(tag.OrderHandlingInstSource)
}

//HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049
func (m OrderCancelReplaceRequest) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

//HasMaturityTime returns true if MaturityTime is present, Tag 1079
func (m OrderCancelReplaceRequest) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

//HasSecondaryDisplayQty returns true if SecondaryDisplayQty is present, Tag 1082
func (m OrderCancelReplaceRequest) HasSecondaryDisplayQty() bool {
	return m.Has(tag.SecondaryDisplayQty)
}

//HasDisplayWhen returns true if DisplayWhen is present, Tag 1083
func (m OrderCancelReplaceRequest) HasDisplayWhen() bool {
	return m.Has(tag.DisplayWhen)
}

//HasDisplayMethod returns true if DisplayMethod is present, Tag 1084
func (m OrderCancelReplaceRequest) HasDisplayMethod() bool {
	return m.Has(tag.DisplayMethod)
}

//HasDisplayLowQty returns true if DisplayLowQty is present, Tag 1085
func (m OrderCancelReplaceRequest) HasDisplayLowQty() bool {
	return m.Has(tag.DisplayLowQty)
}

//HasDisplayHighQty returns true if DisplayHighQty is present, Tag 1086
func (m OrderCancelReplaceRequest) HasDisplayHighQty() bool {
	return m.Has(tag.DisplayHighQty)
}

//HasDisplayMinIncr returns true if DisplayMinIncr is present, Tag 1087
func (m OrderCancelReplaceRequest) HasDisplayMinIncr() bool {
	return m.Has(tag.DisplayMinIncr)
}

//HasRefreshQty returns true if RefreshQty is present, Tag 1088
func (m OrderCancelReplaceRequest) HasRefreshQty() bool {
	return m.Has(tag.RefreshQty)
}

//HasMatchIncrement returns true if MatchIncrement is present, Tag 1089
func (m OrderCancelReplaceRequest) HasMatchIncrement() bool {
	return m.Has(tag.MatchIncrement)
}

//HasMaxPriceLevels returns true if MaxPriceLevels is present, Tag 1090
func (m OrderCancelReplaceRequest) HasMaxPriceLevels() bool {
	return m.Has(tag.MaxPriceLevels)
}

//HasPreTradeAnonymity returns true if PreTradeAnonymity is present, Tag 1091
func (m OrderCancelReplaceRequest) HasPreTradeAnonymity() bool {
	return m.Has(tag.PreTradeAnonymity)
}

//HasPriceProtectionScope returns true if PriceProtectionScope is present, Tag 1092
func (m OrderCancelReplaceRequest) HasPriceProtectionScope() bool {
	return m.Has(tag.PriceProtectionScope)
}

//HasPegPriceType returns true if PegPriceType is present, Tag 1094
func (m OrderCancelReplaceRequest) HasPegPriceType() bool {
	return m.Has(tag.PegPriceType)
}

//HasPegSecurityIDSource returns true if PegSecurityIDSource is present, Tag 1096
func (m OrderCancelReplaceRequest) HasPegSecurityIDSource() bool {
	return m.Has(tag.PegSecurityIDSource)
}

//HasPegSecurityID returns true if PegSecurityID is present, Tag 1097
func (m OrderCancelReplaceRequest) HasPegSecurityID() bool {
	return m.Has(tag.PegSecurityID)
}

//HasPegSymbol returns true if PegSymbol is present, Tag 1098
func (m OrderCancelReplaceRequest) HasPegSymbol() bool {
	return m.Has(tag.PegSymbol)
}

//HasPegSecurityDesc returns true if PegSecurityDesc is present, Tag 1099
func (m OrderCancelReplaceRequest) HasPegSecurityDesc() bool {
	return m.Has(tag.PegSecurityDesc)
}

//HasTriggerType returns true if TriggerType is present, Tag 1100
func (m OrderCancelReplaceRequest) HasTriggerType() bool {
	return m.Has(tag.TriggerType)
}

//HasTriggerAction returns true if TriggerAction is present, Tag 1101
func (m OrderCancelReplaceRequest) HasTriggerAction() bool {
	return m.Has(tag.TriggerAction)
}

//HasTriggerPrice returns true if TriggerPrice is present, Tag 1102
func (m OrderCancelReplaceRequest) HasTriggerPrice() bool {
	return m.Has(tag.TriggerPrice)
}

//HasTriggerSymbol returns true if TriggerSymbol is present, Tag 1103
func (m OrderCancelReplaceRequest) HasTriggerSymbol() bool {
	return m.Has(tag.TriggerSymbol)
}

//HasTriggerSecurityID returns true if TriggerSecurityID is present, Tag 1104
func (m OrderCancelReplaceRequest) HasTriggerSecurityID() bool {
	return m.Has(tag.TriggerSecurityID)
}

//HasTriggerSecurityIDSource returns true if TriggerSecurityIDSource is present, Tag 1105
func (m OrderCancelReplaceRequest) HasTriggerSecurityIDSource() bool {
	return m.Has(tag.TriggerSecurityIDSource)
}

//HasTriggerSecurityDesc returns true if TriggerSecurityDesc is present, Tag 1106
func (m OrderCancelReplaceRequest) HasTriggerSecurityDesc() bool {
	return m.Has(tag.TriggerSecurityDesc)
}

//HasTriggerPriceType returns true if TriggerPriceType is present, Tag 1107
func (m OrderCancelReplaceRequest) HasTriggerPriceType() bool {
	return m.Has(tag.TriggerPriceType)
}

//HasTriggerPriceTypeScope returns true if TriggerPriceTypeScope is present, Tag 1108
func (m OrderCancelReplaceRequest) HasTriggerPriceTypeScope() bool {
	return m.Has(tag.TriggerPriceTypeScope)
}

//HasTriggerPriceDirection returns true if TriggerPriceDirection is present, Tag 1109
func (m OrderCancelReplaceRequest) HasTriggerPriceDirection() bool {
	return m.Has(tag.TriggerPriceDirection)
}

//HasTriggerNewPrice returns true if TriggerNewPrice is present, Tag 1110
func (m OrderCancelReplaceRequest) HasTriggerNewPrice() bool {
	return m.Has(tag.TriggerNewPrice)
}

//HasTriggerOrderType returns true if TriggerOrderType is present, Tag 1111
func (m OrderCancelReplaceRequest) HasTriggerOrderType() bool {
	return m.Has(tag.TriggerOrderType)
}

//HasTriggerNewQty returns true if TriggerNewQty is present, Tag 1112
func (m OrderCancelReplaceRequest) HasTriggerNewQty() bool {
	return m.Has(tag.TriggerNewQty)
}

//HasTriggerTradingSessionID returns true if TriggerTradingSessionID is present, Tag 1113
func (m OrderCancelReplaceRequest) HasTriggerTradingSessionID() bool {
	return m.Has(tag.TriggerTradingSessionID)
}

//HasTriggerTradingSessionSubID returns true if TriggerTradingSessionSubID is present, Tag 1114
func (m OrderCancelReplaceRequest) HasTriggerTradingSessionSubID() bool {
	return m.Has(tag.TriggerTradingSessionSubID)
}

//HasExDestinationIDSource returns true if ExDestinationIDSource is present, Tag 1133
func (m OrderCancelReplaceRequest) HasExDestinationIDSource() bool {
	return m.Has(tag.ExDestinationIDSource)
}

//HasDisplayQty returns true if DisplayQty is present, Tag 1138
func (m OrderCancelReplaceRequest) HasDisplayQty() bool {
	return m.Has(tag.DisplayQty)
}

//HasMinPriceIncrementAmount returns true if MinPriceIncrementAmount is present, Tag 1146
func (m OrderCancelReplaceRequest) HasMinPriceIncrementAmount() bool {
	return m.Has(tag.MinPriceIncrementAmount)
}

//HasUnitOfMeasureQty returns true if UnitOfMeasureQty is present, Tag 1147
func (m OrderCancelReplaceRequest) HasUnitOfMeasureQty() bool {
	return m.Has(tag.UnitOfMeasureQty)
}

//HasSecurityGroup returns true if SecurityGroup is present, Tag 1151
func (m OrderCancelReplaceRequest) HasSecurityGroup() bool {
	return m.Has(tag.SecurityGroup)
}

//HasSecurityXMLLen returns true if SecurityXMLLen is present, Tag 1184
func (m OrderCancelReplaceRequest) HasSecurityXMLLen() bool {
	return m.Has(tag.SecurityXMLLen)
}

//HasSecurityXML returns true if SecurityXML is present, Tag 1185
func (m OrderCancelReplaceRequest) HasSecurityXML() bool {
	return m.Has(tag.SecurityXML)
}

//HasSecurityXMLSchema returns true if SecurityXMLSchema is present, Tag 1186
func (m OrderCancelReplaceRequest) HasSecurityXMLSchema() bool {
	return m.Has(tag.SecurityXMLSchema)
}

//HasPriceUnitOfMeasure returns true if PriceUnitOfMeasure is present, Tag 1191
func (m OrderCancelReplaceRequest) HasPriceUnitOfMeasure() bool {
	return m.Has(tag.PriceUnitOfMeasure)
}

//HasPriceUnitOfMeasureQty returns true if PriceUnitOfMeasureQty is present, Tag 1192
func (m OrderCancelReplaceRequest) HasPriceUnitOfMeasureQty() bool {
	return m.Has(tag.PriceUnitOfMeasureQty)
}

//HasSettlMethod returns true if SettlMethod is present, Tag 1193
func (m OrderCancelReplaceRequest) HasSettlMethod() bool {
	return m.Has(tag.SettlMethod)
}

//HasExerciseStyle returns true if ExerciseStyle is present, Tag 1194
func (m OrderCancelReplaceRequest) HasExerciseStyle() bool {
	return m.Has(tag.ExerciseStyle)
}

//HasOptPayoutAmount returns true if OptPayoutAmount is present, Tag 1195
func (m OrderCancelReplaceRequest) HasOptPayoutAmount() bool {
	return m.Has(tag.OptPayoutAmount)
}

//HasPriceQuoteMethod returns true if PriceQuoteMethod is present, Tag 1196
func (m OrderCancelReplaceRequest) HasPriceQuoteMethod() bool {
	return m.Has(tag.PriceQuoteMethod)
}

//HasValuationMethod returns true if ValuationMethod is present, Tag 1197
func (m OrderCancelReplaceRequest) HasValuationMethod() bool {
	return m.Has(tag.ValuationMethod)
}

//HasListMethod returns true if ListMethod is present, Tag 1198
func (m OrderCancelReplaceRequest) HasListMethod() bool {
	return m.Has(tag.ListMethod)
}

//HasCapPrice returns true if CapPrice is present, Tag 1199
func (m OrderCancelReplaceRequest) HasCapPrice() bool {
	return m.Has(tag.CapPrice)
}

//HasFloorPrice returns true if FloorPrice is present, Tag 1200
func (m OrderCancelReplaceRequest) HasFloorPrice() bool {
	return m.Has(tag.FloorPrice)
}

//HasProductComplex returns true if ProductComplex is present, Tag 1227
func (m OrderCancelReplaceRequest) HasProductComplex() bool {
	return m.Has(tag.ProductComplex)
}

//HasFlexProductEligibilityIndicator returns true if FlexProductEligibilityIndicator is present, Tag 1242
func (m OrderCancelReplaceRequest) HasFlexProductEligibilityIndicator() bool {
	return m.Has(tag.FlexProductEligibilityIndicator)
}

//HasFlexibleIndicator returns true if FlexibleIndicator is present, Tag 1244
func (m OrderCancelReplaceRequest) HasFlexibleIndicator() bool {
	return m.Has(tag.FlexibleIndicator)
}

//HasContractMultiplierUnit returns true if ContractMultiplierUnit is present, Tag 1435
func (m OrderCancelReplaceRequest) HasContractMultiplierUnit() bool {
	return m.Has(tag.ContractMultiplierUnit)
}

//HasFlowScheduleType returns true if FlowScheduleType is present, Tag 1439
func (m OrderCancelReplaceRequest) HasFlowScheduleType() bool {
	return m.Has(tag.FlowScheduleType)
}

//HasRestructuringType returns true if RestructuringType is present, Tag 1449
func (m OrderCancelReplaceRequest) HasRestructuringType() bool {
	return m.Has(tag.RestructuringType)
}

//HasSeniority returns true if Seniority is present, Tag 1450
func (m OrderCancelReplaceRequest) HasSeniority() bool {
	return m.Has(tag.Seniority)
}

//HasNotionalPercentageOutstanding returns true if NotionalPercentageOutstanding is present, Tag 1451
func (m OrderCancelReplaceRequest) HasNotionalPercentageOutstanding() bool {
	return m.Has(tag.NotionalPercentageOutstanding)
}

//HasOriginalNotionalPercentageOutstanding returns true if OriginalNotionalPercentageOutstanding is present, Tag 1452
func (m OrderCancelReplaceRequest) HasOriginalNotionalPercentageOutstanding() bool {
	return m.Has(tag.OriginalNotionalPercentageOutstanding)
}

//HasAttachmentPoint returns true if AttachmentPoint is present, Tag 1457
func (m OrderCancelReplaceRequest) HasAttachmentPoint() bool {
	return m.Has(tag.AttachmentPoint)
}

//HasDetachmentPoint returns true if DetachmentPoint is present, Tag 1458
func (m OrderCancelReplaceRequest) HasDetachmentPoint() bool {
	return m.Has(tag.DetachmentPoint)
}

//HasStrikePriceDeterminationMethod returns true if StrikePriceDeterminationMethod is present, Tag 1478
func (m OrderCancelReplaceRequest) HasStrikePriceDeterminationMethod() bool {
	return m.Has(tag.StrikePriceDeterminationMethod)
}

//HasStrikePriceBoundaryMethod returns true if StrikePriceBoundaryMethod is present, Tag 1479
func (m OrderCancelReplaceRequest) HasStrikePriceBoundaryMethod() bool {
	return m.Has(tag.StrikePriceBoundaryMethod)
}

//HasStrikePriceBoundaryPrecision returns true if StrikePriceBoundaryPrecision is present, Tag 1480
func (m OrderCancelReplaceRequest) HasStrikePriceBoundaryPrecision() bool {
	return m.Has(tag.StrikePriceBoundaryPrecision)
}

//HasUnderlyingPriceDeterminationMethod returns true if UnderlyingPriceDeterminationMethod is present, Tag 1481
func (m OrderCancelReplaceRequest) HasUnderlyingPriceDeterminationMethod() bool {
	return m.Has(tag.UnderlyingPriceDeterminationMethod)
}

//HasOptPayoutType returns true if OptPayoutType is present, Tag 1482
func (m OrderCancelReplaceRequest) HasOptPayoutType() bool {
	return m.Has(tag.OptPayoutType)
}

//HasNoComplexEvents returns true if NoComplexEvents is present, Tag 1483
func (m OrderCancelReplaceRequest) HasNoComplexEvents() bool {
	return m.Has(tag.NoComplexEvents)
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
func (m NoAllocs) SetAllocQty(v float64) {
	m.Set(field.NewAllocQty(v))
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedPartyID), quickfix.GroupElement(tag.NestedPartyIDSource), quickfix.GroupElement(tag.NestedPartyRole), quickfix.GroupElement(tag.NoNestedPartySubIDs)})}
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AllocAccount), quickfix.GroupElement(tag.AllocAcctIDSource), quickfix.GroupElement(tag.AllocSettlCurrency), quickfix.GroupElement(tag.IndividualAllocID), quickfix.GroupElement(tag.NoNestedPartyIDs), quickfix.GroupElement(tag.AllocQty)})}
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), quickfix.GroupElement(tag.NoPartySubIDs)})}
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
func (m NoUnderlyings) SetUnderlyingRepurchaseRate(v float64) {
	m.Set(field.NewUnderlyingRepurchaseRate(v))
}

//SetUnderlyingFactor sets UnderlyingFactor, Tag 246
func (m NoUnderlyings) SetUnderlyingFactor(v float64) {
	m.Set(field.NewUnderlyingFactor(v))
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
func (m NoUnderlyings) SetUnderlyingStrikePrice(v float64) {
	m.Set(field.NewUnderlyingStrikePrice(v))
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
func (m NoUnderlyings) SetUnderlyingContractMultiplier(v float64) {
	m.Set(field.NewUnderlyingContractMultiplier(v))
}

//SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435
func (m NoUnderlyings) SetUnderlyingCouponRate(v float64) {
	m.Set(field.NewUnderlyingCouponRate(v))
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
func (m NoUnderlyings) SetUnderlyingQty(v float64) {
	m.Set(field.NewUnderlyingQty(v))
}

//SetUnderlyingPx sets UnderlyingPx, Tag 810
func (m NoUnderlyings) SetUnderlyingPx(v float64) {
	m.Set(field.NewUnderlyingPx(v))
}

//SetUnderlyingDirtyPrice sets UnderlyingDirtyPrice, Tag 882
func (m NoUnderlyings) SetUnderlyingDirtyPrice(v float64) {
	m.Set(field.NewUnderlyingDirtyPrice(v))
}

//SetUnderlyingEndPrice sets UnderlyingEndPrice, Tag 883
func (m NoUnderlyings) SetUnderlyingEndPrice(v float64) {
	m.Set(field.NewUnderlyingEndPrice(v))
}

//SetUnderlyingStartValue sets UnderlyingStartValue, Tag 884
func (m NoUnderlyings) SetUnderlyingStartValue(v float64) {
	m.Set(field.NewUnderlyingStartValue(v))
}

//SetUnderlyingCurrentValue sets UnderlyingCurrentValue, Tag 885
func (m NoUnderlyings) SetUnderlyingCurrentValue(v float64) {
	m.Set(field.NewUnderlyingCurrentValue(v))
}

//SetUnderlyingEndValue sets UnderlyingEndValue, Tag 886
func (m NoUnderlyings) SetUnderlyingEndValue(v float64) {
	m.Set(field.NewUnderlyingEndValue(v))
}

//SetNoUnderlyingStips sets NoUnderlyingStips, Tag 887
func (m NoUnderlyings) SetNoUnderlyingStips(f NoUnderlyingStipsRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingAllocationPercent sets UnderlyingAllocationPercent, Tag 972
func (m NoUnderlyings) SetUnderlyingAllocationPercent(v float64) {
	m.Set(field.NewUnderlyingAllocationPercent(v))
}

//SetUnderlyingSettlementType sets UnderlyingSettlementType, Tag 975
func (m NoUnderlyings) SetUnderlyingSettlementType(v int) {
	m.Set(field.NewUnderlyingSettlementType(v))
}

//SetUnderlyingCashAmount sets UnderlyingCashAmount, Tag 973
func (m NoUnderlyings) SetUnderlyingCashAmount(v float64) {
	m.Set(field.NewUnderlyingCashAmount(v))
}

//SetUnderlyingCashType sets UnderlyingCashType, Tag 974
func (m NoUnderlyings) SetUnderlyingCashType(v string) {
	m.Set(field.NewUnderlyingCashType(v))
}

//SetUnderlyingUnitOfMeasure sets UnderlyingUnitOfMeasure, Tag 998
func (m NoUnderlyings) SetUnderlyingUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingUnitOfMeasure(v))
}

//SetUnderlyingTimeUnit sets UnderlyingTimeUnit, Tag 1000
func (m NoUnderlyings) SetUnderlyingTimeUnit(v string) {
	m.Set(field.NewUnderlyingTimeUnit(v))
}

//SetUnderlyingCapValue sets UnderlyingCapValue, Tag 1038
func (m NoUnderlyings) SetUnderlyingCapValue(v float64) {
	m.Set(field.NewUnderlyingCapValue(v))
}

//SetNoUndlyInstrumentParties sets NoUndlyInstrumentParties, Tag 1058
func (m NoUnderlyings) SetNoUndlyInstrumentParties(f NoUndlyInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingSettlMethod sets UnderlyingSettlMethod, Tag 1039
func (m NoUnderlyings) SetUnderlyingSettlMethod(v string) {
	m.Set(field.NewUnderlyingSettlMethod(v))
}

//SetUnderlyingAdjustedQuantity sets UnderlyingAdjustedQuantity, Tag 1044
func (m NoUnderlyings) SetUnderlyingAdjustedQuantity(v float64) {
	m.Set(field.NewUnderlyingAdjustedQuantity(v))
}

//SetUnderlyingFXRate sets UnderlyingFXRate, Tag 1045
func (m NoUnderlyings) SetUnderlyingFXRate(v float64) {
	m.Set(field.NewUnderlyingFXRate(v))
}

//SetUnderlyingFXRateCalc sets UnderlyingFXRateCalc, Tag 1046
func (m NoUnderlyings) SetUnderlyingFXRateCalc(v string) {
	m.Set(field.NewUnderlyingFXRateCalc(v))
}

//SetUnderlyingMaturityTime sets UnderlyingMaturityTime, Tag 1213
func (m NoUnderlyings) SetUnderlyingMaturityTime(v string) {
	m.Set(field.NewUnderlyingMaturityTime(v))
}

//SetUnderlyingPutOrCall sets UnderlyingPutOrCall, Tag 315
func (m NoUnderlyings) SetUnderlyingPutOrCall(v int) {
	m.Set(field.NewUnderlyingPutOrCall(v))
}

//SetUnderlyingExerciseStyle sets UnderlyingExerciseStyle, Tag 1419
func (m NoUnderlyings) SetUnderlyingExerciseStyle(v int) {
	m.Set(field.NewUnderlyingExerciseStyle(v))
}

//SetUnderlyingUnitOfMeasureQty sets UnderlyingUnitOfMeasureQty, Tag 1423
func (m NoUnderlyings) SetUnderlyingUnitOfMeasureQty(v float64) {
	m.Set(field.NewUnderlyingUnitOfMeasureQty(v))
}

//SetUnderlyingPriceUnitOfMeasure sets UnderlyingPriceUnitOfMeasure, Tag 1424
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasure(v))
}

//SetUnderlyingPriceUnitOfMeasureQty sets UnderlyingPriceUnitOfMeasureQty, Tag 1425
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasureQty(v float64) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasureQty(v))
}

//SetUnderlyingContractMultiplierUnit sets UnderlyingContractMultiplierUnit, Tag 1437
func (m NoUnderlyings) SetUnderlyingContractMultiplierUnit(v int) {
	m.Set(field.NewUnderlyingContractMultiplierUnit(v))
}

//SetUnderlyingFlowScheduleType sets UnderlyingFlowScheduleType, Tag 1441
func (m NoUnderlyings) SetUnderlyingFlowScheduleType(v int) {
	m.Set(field.NewUnderlyingFlowScheduleType(v))
}

//SetUnderlyingRestructuringType sets UnderlyingRestructuringType, Tag 1453
func (m NoUnderlyings) SetUnderlyingRestructuringType(v string) {
	m.Set(field.NewUnderlyingRestructuringType(v))
}

//SetUnderlyingSeniority sets UnderlyingSeniority, Tag 1454
func (m NoUnderlyings) SetUnderlyingSeniority(v string) {
	m.Set(field.NewUnderlyingSeniority(v))
}

//SetUnderlyingNotionalPercentageOutstanding sets UnderlyingNotionalPercentageOutstanding, Tag 1455
func (m NoUnderlyings) SetUnderlyingNotionalPercentageOutstanding(v float64) {
	m.Set(field.NewUnderlyingNotionalPercentageOutstanding(v))
}

//SetUnderlyingOriginalNotionalPercentageOutstanding sets UnderlyingOriginalNotionalPercentageOutstanding, Tag 1456
func (m NoUnderlyings) SetUnderlyingOriginalNotionalPercentageOutstanding(v float64) {
	m.Set(field.NewUnderlyingOriginalNotionalPercentageOutstanding(v))
}

//SetUnderlyingAttachmentPoint sets UnderlyingAttachmentPoint, Tag 1459
func (m NoUnderlyings) SetUnderlyingAttachmentPoint(v float64) {
	m.Set(field.NewUnderlyingAttachmentPoint(v))
}

//SetUnderlyingDetachmentPoint sets UnderlyingDetachmentPoint, Tag 1460
func (m NoUnderlyings) SetUnderlyingDetachmentPoint(v float64) {
	m.Set(field.NewUnderlyingDetachmentPoint(v))
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

//GetUnderlyingAllocationPercent gets UnderlyingAllocationPercent, Tag 972
func (m NoUnderlyings) GetUnderlyingAllocationPercent() (f field.UnderlyingAllocationPercentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSettlementType gets UnderlyingSettlementType, Tag 975
func (m NoUnderlyings) GetUnderlyingSettlementType() (f field.UnderlyingSettlementTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCashAmount gets UnderlyingCashAmount, Tag 973
func (m NoUnderlyings) GetUnderlyingCashAmount() (f field.UnderlyingCashAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCashType gets UnderlyingCashType, Tag 974
func (m NoUnderlyings) GetUnderlyingCashType() (f field.UnderlyingCashTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingUnitOfMeasure gets UnderlyingUnitOfMeasure, Tag 998
func (m NoUnderlyings) GetUnderlyingUnitOfMeasure() (f field.UnderlyingUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingTimeUnit gets UnderlyingTimeUnit, Tag 1000
func (m NoUnderlyings) GetUnderlyingTimeUnit() (f field.UnderlyingTimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCapValue gets UnderlyingCapValue, Tag 1038
func (m NoUnderlyings) GetUnderlyingCapValue() (f field.UnderlyingCapValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUndlyInstrumentParties gets NoUndlyInstrumentParties, Tag 1058
func (m NoUnderlyings) GetNoUndlyInstrumentParties() (NoUndlyInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingSettlMethod gets UnderlyingSettlMethod, Tag 1039
func (m NoUnderlyings) GetUnderlyingSettlMethod() (f field.UnderlyingSettlMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingAdjustedQuantity gets UnderlyingAdjustedQuantity, Tag 1044
func (m NoUnderlyings) GetUnderlyingAdjustedQuantity() (f field.UnderlyingAdjustedQuantityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFXRate gets UnderlyingFXRate, Tag 1045
func (m NoUnderlyings) GetUnderlyingFXRate() (f field.UnderlyingFXRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFXRateCalc gets UnderlyingFXRateCalc, Tag 1046
func (m NoUnderlyings) GetUnderlyingFXRateCalc() (f field.UnderlyingFXRateCalcField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingMaturityTime gets UnderlyingMaturityTime, Tag 1213
func (m NoUnderlyings) GetUnderlyingMaturityTime() (f field.UnderlyingMaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPutOrCall gets UnderlyingPutOrCall, Tag 315
func (m NoUnderlyings) GetUnderlyingPutOrCall() (f field.UnderlyingPutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingExerciseStyle gets UnderlyingExerciseStyle, Tag 1419
func (m NoUnderlyings) GetUnderlyingExerciseStyle() (f field.UnderlyingExerciseStyleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingUnitOfMeasureQty gets UnderlyingUnitOfMeasureQty, Tag 1423
func (m NoUnderlyings) GetUnderlyingUnitOfMeasureQty() (f field.UnderlyingUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPriceUnitOfMeasure gets UnderlyingPriceUnitOfMeasure, Tag 1424
func (m NoUnderlyings) GetUnderlyingPriceUnitOfMeasure() (f field.UnderlyingPriceUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPriceUnitOfMeasureQty gets UnderlyingPriceUnitOfMeasureQty, Tag 1425
func (m NoUnderlyings) GetUnderlyingPriceUnitOfMeasureQty() (f field.UnderlyingPriceUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingContractMultiplierUnit gets UnderlyingContractMultiplierUnit, Tag 1437
func (m NoUnderlyings) GetUnderlyingContractMultiplierUnit() (f field.UnderlyingContractMultiplierUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFlowScheduleType gets UnderlyingFlowScheduleType, Tag 1441
func (m NoUnderlyings) GetUnderlyingFlowScheduleType() (f field.UnderlyingFlowScheduleTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRestructuringType gets UnderlyingRestructuringType, Tag 1453
func (m NoUnderlyings) GetUnderlyingRestructuringType() (f field.UnderlyingRestructuringTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSeniority gets UnderlyingSeniority, Tag 1454
func (m NoUnderlyings) GetUnderlyingSeniority() (f field.UnderlyingSeniorityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingNotionalPercentageOutstanding gets UnderlyingNotionalPercentageOutstanding, Tag 1455
func (m NoUnderlyings) GetUnderlyingNotionalPercentageOutstanding() (f field.UnderlyingNotionalPercentageOutstandingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingOriginalNotionalPercentageOutstanding gets UnderlyingOriginalNotionalPercentageOutstanding, Tag 1456
func (m NoUnderlyings) GetUnderlyingOriginalNotionalPercentageOutstanding() (f field.UnderlyingOriginalNotionalPercentageOutstandingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingAttachmentPoint gets UnderlyingAttachmentPoint, Tag 1459
func (m NoUnderlyings) GetUnderlyingAttachmentPoint() (f field.UnderlyingAttachmentPointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingDetachmentPoint gets UnderlyingDetachmentPoint, Tag 1460
func (m NoUnderlyings) GetUnderlyingDetachmentPoint() (f field.UnderlyingDetachmentPointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
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

//HasUnderlyingAllocationPercent returns true if UnderlyingAllocationPercent is present, Tag 972
func (m NoUnderlyings) HasUnderlyingAllocationPercent() bool {
	return m.Has(tag.UnderlyingAllocationPercent)
}

//HasUnderlyingSettlementType returns true if UnderlyingSettlementType is present, Tag 975
func (m NoUnderlyings) HasUnderlyingSettlementType() bool {
	return m.Has(tag.UnderlyingSettlementType)
}

//HasUnderlyingCashAmount returns true if UnderlyingCashAmount is present, Tag 973
func (m NoUnderlyings) HasUnderlyingCashAmount() bool {
	return m.Has(tag.UnderlyingCashAmount)
}

//HasUnderlyingCashType returns true if UnderlyingCashType is present, Tag 974
func (m NoUnderlyings) HasUnderlyingCashType() bool {
	return m.Has(tag.UnderlyingCashType)
}

//HasUnderlyingUnitOfMeasure returns true if UnderlyingUnitOfMeasure is present, Tag 998
func (m NoUnderlyings) HasUnderlyingUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingUnitOfMeasure)
}

//HasUnderlyingTimeUnit returns true if UnderlyingTimeUnit is present, Tag 1000
func (m NoUnderlyings) HasUnderlyingTimeUnit() bool {
	return m.Has(tag.UnderlyingTimeUnit)
}

//HasUnderlyingCapValue returns true if UnderlyingCapValue is present, Tag 1038
func (m NoUnderlyings) HasUnderlyingCapValue() bool {
	return m.Has(tag.UnderlyingCapValue)
}

//HasNoUndlyInstrumentParties returns true if NoUndlyInstrumentParties is present, Tag 1058
func (m NoUnderlyings) HasNoUndlyInstrumentParties() bool {
	return m.Has(tag.NoUndlyInstrumentParties)
}

//HasUnderlyingSettlMethod returns true if UnderlyingSettlMethod is present, Tag 1039
func (m NoUnderlyings) HasUnderlyingSettlMethod() bool {
	return m.Has(tag.UnderlyingSettlMethod)
}

//HasUnderlyingAdjustedQuantity returns true if UnderlyingAdjustedQuantity is present, Tag 1044
func (m NoUnderlyings) HasUnderlyingAdjustedQuantity() bool {
	return m.Has(tag.UnderlyingAdjustedQuantity)
}

//HasUnderlyingFXRate returns true if UnderlyingFXRate is present, Tag 1045
func (m NoUnderlyings) HasUnderlyingFXRate() bool {
	return m.Has(tag.UnderlyingFXRate)
}

//HasUnderlyingFXRateCalc returns true if UnderlyingFXRateCalc is present, Tag 1046
func (m NoUnderlyings) HasUnderlyingFXRateCalc() bool {
	return m.Has(tag.UnderlyingFXRateCalc)
}

//HasUnderlyingMaturityTime returns true if UnderlyingMaturityTime is present, Tag 1213
func (m NoUnderlyings) HasUnderlyingMaturityTime() bool {
	return m.Has(tag.UnderlyingMaturityTime)
}

//HasUnderlyingPutOrCall returns true if UnderlyingPutOrCall is present, Tag 315
func (m NoUnderlyings) HasUnderlyingPutOrCall() bool {
	return m.Has(tag.UnderlyingPutOrCall)
}

//HasUnderlyingExerciseStyle returns true if UnderlyingExerciseStyle is present, Tag 1419
func (m NoUnderlyings) HasUnderlyingExerciseStyle() bool {
	return m.Has(tag.UnderlyingExerciseStyle)
}

//HasUnderlyingUnitOfMeasureQty returns true if UnderlyingUnitOfMeasureQty is present, Tag 1423
func (m NoUnderlyings) HasUnderlyingUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingUnitOfMeasureQty)
}

//HasUnderlyingPriceUnitOfMeasure returns true if UnderlyingPriceUnitOfMeasure is present, Tag 1424
func (m NoUnderlyings) HasUnderlyingPriceUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasure)
}

//HasUnderlyingPriceUnitOfMeasureQty returns true if UnderlyingPriceUnitOfMeasureQty is present, Tag 1425
func (m NoUnderlyings) HasUnderlyingPriceUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasureQty)
}

//HasUnderlyingContractMultiplierUnit returns true if UnderlyingContractMultiplierUnit is present, Tag 1437
func (m NoUnderlyings) HasUnderlyingContractMultiplierUnit() bool {
	return m.Has(tag.UnderlyingContractMultiplierUnit)
}

//HasUnderlyingFlowScheduleType returns true if UnderlyingFlowScheduleType is present, Tag 1441
func (m NoUnderlyings) HasUnderlyingFlowScheduleType() bool {
	return m.Has(tag.UnderlyingFlowScheduleType)
}

//HasUnderlyingRestructuringType returns true if UnderlyingRestructuringType is present, Tag 1453
func (m NoUnderlyings) HasUnderlyingRestructuringType() bool {
	return m.Has(tag.UnderlyingRestructuringType)
}

//HasUnderlyingSeniority returns true if UnderlyingSeniority is present, Tag 1454
func (m NoUnderlyings) HasUnderlyingSeniority() bool {
	return m.Has(tag.UnderlyingSeniority)
}

//HasUnderlyingNotionalPercentageOutstanding returns true if UnderlyingNotionalPercentageOutstanding is present, Tag 1455
func (m NoUnderlyings) HasUnderlyingNotionalPercentageOutstanding() bool {
	return m.Has(tag.UnderlyingNotionalPercentageOutstanding)
}

//HasUnderlyingOriginalNotionalPercentageOutstanding returns true if UnderlyingOriginalNotionalPercentageOutstanding is present, Tag 1456
func (m NoUnderlyings) HasUnderlyingOriginalNotionalPercentageOutstanding() bool {
	return m.Has(tag.UnderlyingOriginalNotionalPercentageOutstanding)
}

//HasUnderlyingAttachmentPoint returns true if UnderlyingAttachmentPoint is present, Tag 1459
func (m NoUnderlyings) HasUnderlyingAttachmentPoint() bool {
	return m.Has(tag.UnderlyingAttachmentPoint)
}

//HasUnderlyingDetachmentPoint returns true if UnderlyingDetachmentPoint is present, Tag 1460
func (m NoUnderlyings) HasUnderlyingDetachmentPoint() bool {
	return m.Has(tag.UnderlyingDetachmentPoint)
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

//NoUndlyInstrumentParties is a repeating group element, Tag 1058
type NoUndlyInstrumentParties struct {
	quickfix.Group
}

//SetUnderlyingInstrumentPartyID sets UnderlyingInstrumentPartyID, Tag 1059
func (m NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyID(v string) {
	m.Set(field.NewUnderlyingInstrumentPartyID(v))
}

//SetUnderlyingInstrumentPartyIDSource sets UnderlyingInstrumentPartyIDSource, Tag 1060
func (m NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyIDSource(v string) {
	m.Set(field.NewUnderlyingInstrumentPartyIDSource(v))
}

//SetUnderlyingInstrumentPartyRole sets UnderlyingInstrumentPartyRole, Tag 1061
func (m NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyRole(v int) {
	m.Set(field.NewUnderlyingInstrumentPartyRole(v))
}

//SetNoUndlyInstrumentPartySubIDs sets NoUndlyInstrumentPartySubIDs, Tag 1062
func (m NoUndlyInstrumentParties) SetNoUndlyInstrumentPartySubIDs(f NoUndlyInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetUnderlyingInstrumentPartyID gets UnderlyingInstrumentPartyID, Tag 1059
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyID() (f field.UnderlyingInstrumentPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingInstrumentPartyIDSource gets UnderlyingInstrumentPartyIDSource, Tag 1060
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyIDSource() (f field.UnderlyingInstrumentPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingInstrumentPartyRole gets UnderlyingInstrumentPartyRole, Tag 1061
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyRole() (f field.UnderlyingInstrumentPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUndlyInstrumentPartySubIDs gets NoUndlyInstrumentPartySubIDs, Tag 1062
func (m NoUndlyInstrumentParties) GetNoUndlyInstrumentPartySubIDs() (NoUndlyInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasUnderlyingInstrumentPartyID returns true if UnderlyingInstrumentPartyID is present, Tag 1059
func (m NoUndlyInstrumentParties) HasUnderlyingInstrumentPartyID() bool {
	return m.Has(tag.UnderlyingInstrumentPartyID)
}

//HasUnderlyingInstrumentPartyIDSource returns true if UnderlyingInstrumentPartyIDSource is present, Tag 1060
func (m NoUndlyInstrumentParties) HasUnderlyingInstrumentPartyIDSource() bool {
	return m.Has(tag.UnderlyingInstrumentPartyIDSource)
}

//HasUnderlyingInstrumentPartyRole returns true if UnderlyingInstrumentPartyRole is present, Tag 1061
func (m NoUndlyInstrumentParties) HasUnderlyingInstrumentPartyRole() bool {
	return m.Has(tag.UnderlyingInstrumentPartyRole)
}

//HasNoUndlyInstrumentPartySubIDs returns true if NoUndlyInstrumentPartySubIDs is present, Tag 1062
func (m NoUndlyInstrumentParties) HasNoUndlyInstrumentPartySubIDs() bool {
	return m.Has(tag.NoUndlyInstrumentPartySubIDs)
}

//NoUndlyInstrumentPartySubIDs is a repeating group element, Tag 1062
type NoUndlyInstrumentPartySubIDs struct {
	quickfix.Group
}

//SetUnderlyingInstrumentPartySubID sets UnderlyingInstrumentPartySubID, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) SetUnderlyingInstrumentPartySubID(v string) {
	m.Set(field.NewUnderlyingInstrumentPartySubID(v))
}

//SetUnderlyingInstrumentPartySubIDType sets UnderlyingInstrumentPartySubIDType, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) SetUnderlyingInstrumentPartySubIDType(v int) {
	m.Set(field.NewUnderlyingInstrumentPartySubIDType(v))
}

//GetUnderlyingInstrumentPartySubID gets UnderlyingInstrumentPartySubID, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) GetUnderlyingInstrumentPartySubID() (f field.UnderlyingInstrumentPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingInstrumentPartySubIDType gets UnderlyingInstrumentPartySubIDType, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) GetUnderlyingInstrumentPartySubIDType() (f field.UnderlyingInstrumentPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingInstrumentPartySubID returns true if UnderlyingInstrumentPartySubID is present, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) HasUnderlyingInstrumentPartySubID() bool {
	return m.Has(tag.UnderlyingInstrumentPartySubID)
}

//HasUnderlyingInstrumentPartySubIDType returns true if UnderlyingInstrumentPartySubIDType is present, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) HasUnderlyingInstrumentPartySubIDType() bool {
	return m.Has(tag.UnderlyingInstrumentPartySubIDType)
}

//NoUndlyInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1062
type NoUndlyInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUndlyInstrumentPartySubIDsRepeatingGroup returns an initialized, NoUndlyInstrumentPartySubIDsRepeatingGroup
func NewNoUndlyInstrumentPartySubIDsRepeatingGroup() NoUndlyInstrumentPartySubIDsRepeatingGroup {
	return NoUndlyInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingInstrumentPartySubID), quickfix.GroupElement(tag.UnderlyingInstrumentPartySubIDType)})}
}

//Add create and append a new NoUndlyInstrumentPartySubIDs to this group
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Add() NoUndlyInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentPartySubIDs{g}
}

//Get returns the ith NoUndlyInstrumentPartySubIDs in the NoUndlyInstrumentPartySubIDsRepeatinGroup
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Get(i int) NoUndlyInstrumentPartySubIDs {
	return NoUndlyInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoUndlyInstrumentPartiesRepeatingGroup is a repeating group, Tag 1058
type NoUndlyInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUndlyInstrumentPartiesRepeatingGroup returns an initialized, NoUndlyInstrumentPartiesRepeatingGroup
func NewNoUndlyInstrumentPartiesRepeatingGroup() NoUndlyInstrumentPartiesRepeatingGroup {
	return NoUndlyInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingInstrumentPartyID), quickfix.GroupElement(tag.UnderlyingInstrumentPartyIDSource), quickfix.GroupElement(tag.UnderlyingInstrumentPartyRole), quickfix.GroupElement(tag.NoUndlyInstrumentPartySubIDs)})}
}

//Add create and append a new NoUndlyInstrumentParties to this group
func (m NoUndlyInstrumentPartiesRepeatingGroup) Add() NoUndlyInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentParties{g}
}

//Get returns the ith NoUndlyInstrumentParties in the NoUndlyInstrumentPartiesRepeatinGroup
func (m NoUndlyInstrumentPartiesRepeatingGroup) Get(i int) NoUndlyInstrumentParties {
	return NoUndlyInstrumentParties{m.RepeatingGroup.Get(i)}
}

//NoUnderlyingsRepeatingGroup is a repeating group, Tag 711
type NoUnderlyingsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingsRepeatingGroup returns an initialized, NoUnderlyingsRepeatingGroup
func NewNoUnderlyingsRepeatingGroup() NoUnderlyingsRepeatingGroup {
	return NoUnderlyingsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyings,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingSecurityIDSource), quickfix.GroupElement(tag.NoUnderlyingSecurityAltID), quickfix.GroupElement(tag.UnderlyingProduct), quickfix.GroupElement(tag.UnderlyingCFICode), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingSecuritySubType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDate), quickfix.GroupElement(tag.UnderlyingCouponPaymentDate), quickfix.GroupElement(tag.UnderlyingIssueDate), quickfix.GroupElement(tag.UnderlyingRepoCollateralSecurityType), quickfix.GroupElement(tag.UnderlyingRepurchaseTerm), quickfix.GroupElement(tag.UnderlyingRepurchaseRate), quickfix.GroupElement(tag.UnderlyingFactor), quickfix.GroupElement(tag.UnderlyingCreditRating), quickfix.GroupElement(tag.UnderlyingInstrRegistry), quickfix.GroupElement(tag.UnderlyingCountryOfIssue), quickfix.GroupElement(tag.UnderlyingStateOrProvinceOfIssue), quickfix.GroupElement(tag.UnderlyingLocaleOfIssue), quickfix.GroupElement(tag.UnderlyingRedemptionDate), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingStrikeCurrency), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.UnderlyingCPProgram), quickfix.GroupElement(tag.UnderlyingCPRegType), quickfix.GroupElement(tag.UnderlyingCurrency), quickfix.GroupElement(tag.UnderlyingQty), quickfix.GroupElement(tag.UnderlyingPx), quickfix.GroupElement(tag.UnderlyingDirtyPrice), quickfix.GroupElement(tag.UnderlyingEndPrice), quickfix.GroupElement(tag.UnderlyingStartValue), quickfix.GroupElement(tag.UnderlyingCurrentValue), quickfix.GroupElement(tag.UnderlyingEndValue), quickfix.GroupElement(tag.NoUnderlyingStips), quickfix.GroupElement(tag.UnderlyingAllocationPercent), quickfix.GroupElement(tag.UnderlyingSettlementType), quickfix.GroupElement(tag.UnderlyingCashAmount), quickfix.GroupElement(tag.UnderlyingCashType), quickfix.GroupElement(tag.UnderlyingUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingTimeUnit), quickfix.GroupElement(tag.UnderlyingCapValue), quickfix.GroupElement(tag.NoUndlyInstrumentParties), quickfix.GroupElement(tag.UnderlyingSettlMethod), quickfix.GroupElement(tag.UnderlyingAdjustedQuantity), quickfix.GroupElement(tag.UnderlyingFXRate), quickfix.GroupElement(tag.UnderlyingFXRateCalc), quickfix.GroupElement(tag.UnderlyingMaturityTime), quickfix.GroupElement(tag.UnderlyingPutOrCall), quickfix.GroupElement(tag.UnderlyingExerciseStyle), quickfix.GroupElement(tag.UnderlyingUnitOfMeasureQty), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasureQty), quickfix.GroupElement(tag.UnderlyingContractMultiplierUnit), quickfix.GroupElement(tag.UnderlyingFlowScheduleType), quickfix.GroupElement(tag.UnderlyingRestructuringType), quickfix.GroupElement(tag.UnderlyingSeniority), quickfix.GroupElement(tag.UnderlyingNotionalPercentageOutstanding), quickfix.GroupElement(tag.UnderlyingOriginalNotionalPercentageOutstanding), quickfix.GroupElement(tag.UnderlyingAttachmentPoint), quickfix.GroupElement(tag.UnderlyingDetachmentPoint)})}
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

//NoTrdRegTimestamps is a repeating group element, Tag 768
type NoTrdRegTimestamps struct {
	quickfix.Group
}

//SetTrdRegTimestamp sets TrdRegTimestamp, Tag 769
func (m NoTrdRegTimestamps) SetTrdRegTimestamp(v time.Time) {
	m.Set(field.NewTrdRegTimestamp(v))
}

//SetTrdRegTimestampType sets TrdRegTimestampType, Tag 770
func (m NoTrdRegTimestamps) SetTrdRegTimestampType(v int) {
	m.Set(field.NewTrdRegTimestampType(v))
}

//SetTrdRegTimestampOrigin sets TrdRegTimestampOrigin, Tag 771
func (m NoTrdRegTimestamps) SetTrdRegTimestampOrigin(v string) {
	m.Set(field.NewTrdRegTimestampOrigin(v))
}

//SetDeskType sets DeskType, Tag 1033
func (m NoTrdRegTimestamps) SetDeskType(v string) {
	m.Set(field.NewDeskType(v))
}

//SetDeskTypeSource sets DeskTypeSource, Tag 1034
func (m NoTrdRegTimestamps) SetDeskTypeSource(v int) {
	m.Set(field.NewDeskTypeSource(v))
}

//SetDeskOrderHandlingInst sets DeskOrderHandlingInst, Tag 1035
func (m NoTrdRegTimestamps) SetDeskOrderHandlingInst(v string) {
	m.Set(field.NewDeskOrderHandlingInst(v))
}

//GetTrdRegTimestamp gets TrdRegTimestamp, Tag 769
func (m NoTrdRegTimestamps) GetTrdRegTimestamp() (f field.TrdRegTimestampField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTrdRegTimestampType gets TrdRegTimestampType, Tag 770
func (m NoTrdRegTimestamps) GetTrdRegTimestampType() (f field.TrdRegTimestampTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTrdRegTimestampOrigin gets TrdRegTimestampOrigin, Tag 771
func (m NoTrdRegTimestamps) GetTrdRegTimestampOrigin() (f field.TrdRegTimestampOriginField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDeskType gets DeskType, Tag 1033
func (m NoTrdRegTimestamps) GetDeskType() (f field.DeskTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDeskTypeSource gets DeskTypeSource, Tag 1034
func (m NoTrdRegTimestamps) GetDeskTypeSource() (f field.DeskTypeSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDeskOrderHandlingInst gets DeskOrderHandlingInst, Tag 1035
func (m NoTrdRegTimestamps) GetDeskOrderHandlingInst() (f field.DeskOrderHandlingInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasTrdRegTimestamp returns true if TrdRegTimestamp is present, Tag 769
func (m NoTrdRegTimestamps) HasTrdRegTimestamp() bool {
	return m.Has(tag.TrdRegTimestamp)
}

//HasTrdRegTimestampType returns true if TrdRegTimestampType is present, Tag 770
func (m NoTrdRegTimestamps) HasTrdRegTimestampType() bool {
	return m.Has(tag.TrdRegTimestampType)
}

//HasTrdRegTimestampOrigin returns true if TrdRegTimestampOrigin is present, Tag 771
func (m NoTrdRegTimestamps) HasTrdRegTimestampOrigin() bool {
	return m.Has(tag.TrdRegTimestampOrigin)
}

//HasDeskType returns true if DeskType is present, Tag 1033
func (m NoTrdRegTimestamps) HasDeskType() bool {
	return m.Has(tag.DeskType)
}

//HasDeskTypeSource returns true if DeskTypeSource is present, Tag 1034
func (m NoTrdRegTimestamps) HasDeskTypeSource() bool {
	return m.Has(tag.DeskTypeSource)
}

//HasDeskOrderHandlingInst returns true if DeskOrderHandlingInst is present, Tag 1035
func (m NoTrdRegTimestamps) HasDeskOrderHandlingInst() bool {
	return m.Has(tag.DeskOrderHandlingInst)
}

//NoTrdRegTimestampsRepeatingGroup is a repeating group, Tag 768
type NoTrdRegTimestampsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTrdRegTimestampsRepeatingGroup returns an initialized, NoTrdRegTimestampsRepeatingGroup
func NewNoTrdRegTimestampsRepeatingGroup() NoTrdRegTimestampsRepeatingGroup {
	return NoTrdRegTimestampsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTrdRegTimestamps,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TrdRegTimestamp), quickfix.GroupElement(tag.TrdRegTimestampType), quickfix.GroupElement(tag.TrdRegTimestampOrigin), quickfix.GroupElement(tag.DeskType), quickfix.GroupElement(tag.DeskTypeSource), quickfix.GroupElement(tag.DeskOrderHandlingInst)})}
}

//Add create and append a new NoTrdRegTimestamps to this group
func (m NoTrdRegTimestampsRepeatingGroup) Add() NoTrdRegTimestamps {
	g := m.RepeatingGroup.Add()
	return NoTrdRegTimestamps{g}
}

//Get returns the ith NoTrdRegTimestamps in the NoTrdRegTimestampsRepeatinGroup
func (m NoTrdRegTimestampsRepeatingGroup) Get(i int) NoTrdRegTimestamps {
	return NoTrdRegTimestamps{m.RepeatingGroup.Get(i)}
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
func (m NoEvents) SetEventPx(v float64) {
	m.Set(field.NewEventPx(v))
}

//SetEventText sets EventText, Tag 868
func (m NoEvents) SetEventText(v string) {
	m.Set(field.NewEventText(v))
}

//SetEventTime sets EventTime, Tag 1145
func (m NoEvents) SetEventTime(v time.Time) {
	m.Set(field.NewEventTime(v))
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

//GetEventTime gets EventTime, Tag 1145
func (m NoEvents) GetEventTime() (f field.EventTimeField, err quickfix.MessageRejectError) {
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

//HasEventTime returns true if EventTime is present, Tag 1145
func (m NoEvents) HasEventTime() bool {
	return m.Has(tag.EventTime)
}

//NoEventsRepeatingGroup is a repeating group, Tag 864
type NoEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoEventsRepeatingGroup returns an initialized, NoEventsRepeatingGroup
func NewNoEventsRepeatingGroup() NoEventsRepeatingGroup {
	return NoEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.EventType), quickfix.GroupElement(tag.EventDate), quickfix.GroupElement(tag.EventPx), quickfix.GroupElement(tag.EventText), quickfix.GroupElement(tag.EventTime)})}
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

//NoStrategyParameters is a repeating group element, Tag 957
type NoStrategyParameters struct {
	quickfix.Group
}

//SetStrategyParameterName sets StrategyParameterName, Tag 958
func (m NoStrategyParameters) SetStrategyParameterName(v string) {
	m.Set(field.NewStrategyParameterName(v))
}

//SetStrategyParameterType sets StrategyParameterType, Tag 959
func (m NoStrategyParameters) SetStrategyParameterType(v int) {
	m.Set(field.NewStrategyParameterType(v))
}

//SetStrategyParameterValue sets StrategyParameterValue, Tag 960
func (m NoStrategyParameters) SetStrategyParameterValue(v string) {
	m.Set(field.NewStrategyParameterValue(v))
}

//GetStrategyParameterName gets StrategyParameterName, Tag 958
func (m NoStrategyParameters) GetStrategyParameterName() (f field.StrategyParameterNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrategyParameterType gets StrategyParameterType, Tag 959
func (m NoStrategyParameters) GetStrategyParameterType() (f field.StrategyParameterTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrategyParameterValue gets StrategyParameterValue, Tag 960
func (m NoStrategyParameters) GetStrategyParameterValue() (f field.StrategyParameterValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasStrategyParameterName returns true if StrategyParameterName is present, Tag 958
func (m NoStrategyParameters) HasStrategyParameterName() bool {
	return m.Has(tag.StrategyParameterName)
}

//HasStrategyParameterType returns true if StrategyParameterType is present, Tag 959
func (m NoStrategyParameters) HasStrategyParameterType() bool {
	return m.Has(tag.StrategyParameterType)
}

//HasStrategyParameterValue returns true if StrategyParameterValue is present, Tag 960
func (m NoStrategyParameters) HasStrategyParameterValue() bool {
	return m.Has(tag.StrategyParameterValue)
}

//NoStrategyParametersRepeatingGroup is a repeating group, Tag 957
type NoStrategyParametersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoStrategyParametersRepeatingGroup returns an initialized, NoStrategyParametersRepeatingGroup
func NewNoStrategyParametersRepeatingGroup() NoStrategyParametersRepeatingGroup {
	return NoStrategyParametersRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoStrategyParameters,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.StrategyParameterName), quickfix.GroupElement(tag.StrategyParameterType), quickfix.GroupElement(tag.StrategyParameterValue)})}
}

//Add create and append a new NoStrategyParameters to this group
func (m NoStrategyParametersRepeatingGroup) Add() NoStrategyParameters {
	g := m.RepeatingGroup.Add()
	return NoStrategyParameters{g}
}

//Get returns the ith NoStrategyParameters in the NoStrategyParametersRepeatinGroup
func (m NoStrategyParametersRepeatingGroup) Get(i int) NoStrategyParameters {
	return NoStrategyParameters{m.RepeatingGroup.Get(i)}
}

//NoInstrumentParties is a repeating group element, Tag 1018
type NoInstrumentParties struct {
	quickfix.Group
}

//SetInstrumentPartyID sets InstrumentPartyID, Tag 1019
func (m NoInstrumentParties) SetInstrumentPartyID(v string) {
	m.Set(field.NewInstrumentPartyID(v))
}

//SetInstrumentPartyIDSource sets InstrumentPartyIDSource, Tag 1050
func (m NoInstrumentParties) SetInstrumentPartyIDSource(v string) {
	m.Set(field.NewInstrumentPartyIDSource(v))
}

//SetInstrumentPartyRole sets InstrumentPartyRole, Tag 1051
func (m NoInstrumentParties) SetInstrumentPartyRole(v int) {
	m.Set(field.NewInstrumentPartyRole(v))
}

//SetNoInstrumentPartySubIDs sets NoInstrumentPartySubIDs, Tag 1052
func (m NoInstrumentParties) SetNoInstrumentPartySubIDs(f NoInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetInstrumentPartyID gets InstrumentPartyID, Tag 1019
func (m NoInstrumentParties) GetInstrumentPartyID() (f field.InstrumentPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrumentPartyIDSource gets InstrumentPartyIDSource, Tag 1050
func (m NoInstrumentParties) GetInstrumentPartyIDSource() (f field.InstrumentPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrumentPartyRole gets InstrumentPartyRole, Tag 1051
func (m NoInstrumentParties) GetInstrumentPartyRole() (f field.InstrumentPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoInstrumentPartySubIDs gets NoInstrumentPartySubIDs, Tag 1052
func (m NoInstrumentParties) GetNoInstrumentPartySubIDs() (NoInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasInstrumentPartyID returns true if InstrumentPartyID is present, Tag 1019
func (m NoInstrumentParties) HasInstrumentPartyID() bool {
	return m.Has(tag.InstrumentPartyID)
}

//HasInstrumentPartyIDSource returns true if InstrumentPartyIDSource is present, Tag 1050
func (m NoInstrumentParties) HasInstrumentPartyIDSource() bool {
	return m.Has(tag.InstrumentPartyIDSource)
}

//HasInstrumentPartyRole returns true if InstrumentPartyRole is present, Tag 1051
func (m NoInstrumentParties) HasInstrumentPartyRole() bool {
	return m.Has(tag.InstrumentPartyRole)
}

//HasNoInstrumentPartySubIDs returns true if NoInstrumentPartySubIDs is present, Tag 1052
func (m NoInstrumentParties) HasNoInstrumentPartySubIDs() bool {
	return m.Has(tag.NoInstrumentPartySubIDs)
}

//NoInstrumentPartySubIDs is a repeating group element, Tag 1052
type NoInstrumentPartySubIDs struct {
	quickfix.Group
}

//SetInstrumentPartySubID sets InstrumentPartySubID, Tag 1053
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubID(v string) {
	m.Set(field.NewInstrumentPartySubID(v))
}

//SetInstrumentPartySubIDType sets InstrumentPartySubIDType, Tag 1054
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubIDType(v int) {
	m.Set(field.NewInstrumentPartySubIDType(v))
}

//GetInstrumentPartySubID gets InstrumentPartySubID, Tag 1053
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubID() (f field.InstrumentPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrumentPartySubIDType gets InstrumentPartySubIDType, Tag 1054
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubIDType() (f field.InstrumentPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasInstrumentPartySubID returns true if InstrumentPartySubID is present, Tag 1053
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubID() bool {
	return m.Has(tag.InstrumentPartySubID)
}

//HasInstrumentPartySubIDType returns true if InstrumentPartySubIDType is present, Tag 1054
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubIDType() bool {
	return m.Has(tag.InstrumentPartySubIDType)
}

//NoInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1052
type NoInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoInstrumentPartySubIDsRepeatingGroup returns an initialized, NoInstrumentPartySubIDsRepeatingGroup
func NewNoInstrumentPartySubIDsRepeatingGroup() NoInstrumentPartySubIDsRepeatingGroup {
	return NoInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartySubID), quickfix.GroupElement(tag.InstrumentPartySubIDType)})}
}

//Add create and append a new NoInstrumentPartySubIDs to this group
func (m NoInstrumentPartySubIDsRepeatingGroup) Add() NoInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoInstrumentPartySubIDs{g}
}

//Get returns the ith NoInstrumentPartySubIDs in the NoInstrumentPartySubIDsRepeatinGroup
func (m NoInstrumentPartySubIDsRepeatingGroup) Get(i int) NoInstrumentPartySubIDs {
	return NoInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoInstrumentPartiesRepeatingGroup is a repeating group, Tag 1018
type NoInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoInstrumentPartiesRepeatingGroup returns an initialized, NoInstrumentPartiesRepeatingGroup
func NewNoInstrumentPartiesRepeatingGroup() NoInstrumentPartiesRepeatingGroup {
	return NoInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartyID), quickfix.GroupElement(tag.InstrumentPartyIDSource), quickfix.GroupElement(tag.InstrumentPartyRole), quickfix.GroupElement(tag.NoInstrumentPartySubIDs)})}
}

//Add create and append a new NoInstrumentParties to this group
func (m NoInstrumentPartiesRepeatingGroup) Add() NoInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoInstrumentParties{g}
}

//Get returns the ith NoInstrumentParties in the NoInstrumentPartiesRepeatinGroup
func (m NoInstrumentPartiesRepeatingGroup) Get(i int) NoInstrumentParties {
	return NoInstrumentParties{m.RepeatingGroup.Get(i)}
}

//NoComplexEvents is a repeating group element, Tag 1483
type NoComplexEvents struct {
	quickfix.Group
}

//SetComplexEventType sets ComplexEventType, Tag 1484
func (m NoComplexEvents) SetComplexEventType(v int) {
	m.Set(field.NewComplexEventType(v))
}

//SetComplexOptPayoutAmount sets ComplexOptPayoutAmount, Tag 1485
func (m NoComplexEvents) SetComplexOptPayoutAmount(v float64) {
	m.Set(field.NewComplexOptPayoutAmount(v))
}

//SetComplexEventPrice sets ComplexEventPrice, Tag 1486
func (m NoComplexEvents) SetComplexEventPrice(v float64) {
	m.Set(field.NewComplexEventPrice(v))
}

//SetComplexEventPriceBoundaryMethod sets ComplexEventPriceBoundaryMethod, Tag 1487
func (m NoComplexEvents) SetComplexEventPriceBoundaryMethod(v int) {
	m.Set(field.NewComplexEventPriceBoundaryMethod(v))
}

//SetComplexEventPriceBoundaryPrecision sets ComplexEventPriceBoundaryPrecision, Tag 1488
func (m NoComplexEvents) SetComplexEventPriceBoundaryPrecision(v float64) {
	m.Set(field.NewComplexEventPriceBoundaryPrecision(v))
}

//SetComplexEventPriceTimeType sets ComplexEventPriceTimeType, Tag 1489
func (m NoComplexEvents) SetComplexEventPriceTimeType(v int) {
	m.Set(field.NewComplexEventPriceTimeType(v))
}

//SetComplexEventCondition sets ComplexEventCondition, Tag 1490
func (m NoComplexEvents) SetComplexEventCondition(v int) {
	m.Set(field.NewComplexEventCondition(v))
}

//SetNoComplexEventDates sets NoComplexEventDates, Tag 1491
func (m NoComplexEvents) SetNoComplexEventDates(f NoComplexEventDatesRepeatingGroup) {
	m.SetGroup(f)
}

//GetComplexEventType gets ComplexEventType, Tag 1484
func (m NoComplexEvents) GetComplexEventType() (f field.ComplexEventTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexOptPayoutAmount gets ComplexOptPayoutAmount, Tag 1485
func (m NoComplexEvents) GetComplexOptPayoutAmount() (f field.ComplexOptPayoutAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventPrice gets ComplexEventPrice, Tag 1486
func (m NoComplexEvents) GetComplexEventPrice() (f field.ComplexEventPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventPriceBoundaryMethod gets ComplexEventPriceBoundaryMethod, Tag 1487
func (m NoComplexEvents) GetComplexEventPriceBoundaryMethod() (f field.ComplexEventPriceBoundaryMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventPriceBoundaryPrecision gets ComplexEventPriceBoundaryPrecision, Tag 1488
func (m NoComplexEvents) GetComplexEventPriceBoundaryPrecision() (f field.ComplexEventPriceBoundaryPrecisionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventPriceTimeType gets ComplexEventPriceTimeType, Tag 1489
func (m NoComplexEvents) GetComplexEventPriceTimeType() (f field.ComplexEventPriceTimeTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventCondition gets ComplexEventCondition, Tag 1490
func (m NoComplexEvents) GetComplexEventCondition() (f field.ComplexEventConditionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoComplexEventDates gets NoComplexEventDates, Tag 1491
func (m NoComplexEvents) GetNoComplexEventDates() (NoComplexEventDatesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventDatesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasComplexEventType returns true if ComplexEventType is present, Tag 1484
func (m NoComplexEvents) HasComplexEventType() bool {
	return m.Has(tag.ComplexEventType)
}

//HasComplexOptPayoutAmount returns true if ComplexOptPayoutAmount is present, Tag 1485
func (m NoComplexEvents) HasComplexOptPayoutAmount() bool {
	return m.Has(tag.ComplexOptPayoutAmount)
}

//HasComplexEventPrice returns true if ComplexEventPrice is present, Tag 1486
func (m NoComplexEvents) HasComplexEventPrice() bool {
	return m.Has(tag.ComplexEventPrice)
}

//HasComplexEventPriceBoundaryMethod returns true if ComplexEventPriceBoundaryMethod is present, Tag 1487
func (m NoComplexEvents) HasComplexEventPriceBoundaryMethod() bool {
	return m.Has(tag.ComplexEventPriceBoundaryMethod)
}

//HasComplexEventPriceBoundaryPrecision returns true if ComplexEventPriceBoundaryPrecision is present, Tag 1488
func (m NoComplexEvents) HasComplexEventPriceBoundaryPrecision() bool {
	return m.Has(tag.ComplexEventPriceBoundaryPrecision)
}

//HasComplexEventPriceTimeType returns true if ComplexEventPriceTimeType is present, Tag 1489
func (m NoComplexEvents) HasComplexEventPriceTimeType() bool {
	return m.Has(tag.ComplexEventPriceTimeType)
}

//HasComplexEventCondition returns true if ComplexEventCondition is present, Tag 1490
func (m NoComplexEvents) HasComplexEventCondition() bool {
	return m.Has(tag.ComplexEventCondition)
}

//HasNoComplexEventDates returns true if NoComplexEventDates is present, Tag 1491
func (m NoComplexEvents) HasNoComplexEventDates() bool {
	return m.Has(tag.NoComplexEventDates)
}

//NoComplexEventDates is a repeating group element, Tag 1491
type NoComplexEventDates struct {
	quickfix.Group
}

//SetComplexEventStartDate sets ComplexEventStartDate, Tag 1492
func (m NoComplexEventDates) SetComplexEventStartDate(v time.Time) {
	m.Set(field.NewComplexEventStartDate(v))
}

//SetComplexEventEndDate sets ComplexEventEndDate, Tag 1493
func (m NoComplexEventDates) SetComplexEventEndDate(v time.Time) {
	m.Set(field.NewComplexEventEndDate(v))
}

//SetNoComplexEventTimes sets NoComplexEventTimes, Tag 1494
func (m NoComplexEventDates) SetNoComplexEventTimes(f NoComplexEventTimesRepeatingGroup) {
	m.SetGroup(f)
}

//GetComplexEventStartDate gets ComplexEventStartDate, Tag 1492
func (m NoComplexEventDates) GetComplexEventStartDate() (f field.ComplexEventStartDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventEndDate gets ComplexEventEndDate, Tag 1493
func (m NoComplexEventDates) GetComplexEventEndDate() (f field.ComplexEventEndDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoComplexEventTimes gets NoComplexEventTimes, Tag 1494
func (m NoComplexEventDates) GetNoComplexEventTimes() (NoComplexEventTimesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventTimesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasComplexEventStartDate returns true if ComplexEventStartDate is present, Tag 1492
func (m NoComplexEventDates) HasComplexEventStartDate() bool {
	return m.Has(tag.ComplexEventStartDate)
}

//HasComplexEventEndDate returns true if ComplexEventEndDate is present, Tag 1493
func (m NoComplexEventDates) HasComplexEventEndDate() bool {
	return m.Has(tag.ComplexEventEndDate)
}

//HasNoComplexEventTimes returns true if NoComplexEventTimes is present, Tag 1494
func (m NoComplexEventDates) HasNoComplexEventTimes() bool {
	return m.Has(tag.NoComplexEventTimes)
}

//NoComplexEventTimes is a repeating group element, Tag 1494
type NoComplexEventTimes struct {
	quickfix.Group
}

//SetComplexEventStartTime sets ComplexEventStartTime, Tag 1495
func (m NoComplexEventTimes) SetComplexEventStartTime(v string) {
	m.Set(field.NewComplexEventStartTime(v))
}

//SetComplexEventEndTime sets ComplexEventEndTime, Tag 1496
func (m NoComplexEventTimes) SetComplexEventEndTime(v string) {
	m.Set(field.NewComplexEventEndTime(v))
}

//GetComplexEventStartTime gets ComplexEventStartTime, Tag 1495
func (m NoComplexEventTimes) GetComplexEventStartTime() (f field.ComplexEventStartTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventEndTime gets ComplexEventEndTime, Tag 1496
func (m NoComplexEventTimes) GetComplexEventEndTime() (f field.ComplexEventEndTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasComplexEventStartTime returns true if ComplexEventStartTime is present, Tag 1495
func (m NoComplexEventTimes) HasComplexEventStartTime() bool {
	return m.Has(tag.ComplexEventStartTime)
}

//HasComplexEventEndTime returns true if ComplexEventEndTime is present, Tag 1496
func (m NoComplexEventTimes) HasComplexEventEndTime() bool {
	return m.Has(tag.ComplexEventEndTime)
}

//NoComplexEventTimesRepeatingGroup is a repeating group, Tag 1494
type NoComplexEventTimesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoComplexEventTimesRepeatingGroup returns an initialized, NoComplexEventTimesRepeatingGroup
func NewNoComplexEventTimesRepeatingGroup() NoComplexEventTimesRepeatingGroup {
	return NoComplexEventTimesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEventTimes,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventStartTime), quickfix.GroupElement(tag.ComplexEventEndTime)})}
}

//Add create and append a new NoComplexEventTimes to this group
func (m NoComplexEventTimesRepeatingGroup) Add() NoComplexEventTimes {
	g := m.RepeatingGroup.Add()
	return NoComplexEventTimes{g}
}

//Get returns the ith NoComplexEventTimes in the NoComplexEventTimesRepeatinGroup
func (m NoComplexEventTimesRepeatingGroup) Get(i int) NoComplexEventTimes {
	return NoComplexEventTimes{m.RepeatingGroup.Get(i)}
}

//NoComplexEventDatesRepeatingGroup is a repeating group, Tag 1491
type NoComplexEventDatesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoComplexEventDatesRepeatingGroup returns an initialized, NoComplexEventDatesRepeatingGroup
func NewNoComplexEventDatesRepeatingGroup() NoComplexEventDatesRepeatingGroup {
	return NoComplexEventDatesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEventDates,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventStartDate), quickfix.GroupElement(tag.ComplexEventEndDate), quickfix.GroupElement(tag.NoComplexEventTimes)})}
}

//Add create and append a new NoComplexEventDates to this group
func (m NoComplexEventDatesRepeatingGroup) Add() NoComplexEventDates {
	g := m.RepeatingGroup.Add()
	return NoComplexEventDates{g}
}

//Get returns the ith NoComplexEventDates in the NoComplexEventDatesRepeatinGroup
func (m NoComplexEventDatesRepeatingGroup) Get(i int) NoComplexEventDates {
	return NoComplexEventDates{m.RepeatingGroup.Get(i)}
}

//NoComplexEventsRepeatingGroup is a repeating group, Tag 1483
type NoComplexEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoComplexEventsRepeatingGroup returns an initialized, NoComplexEventsRepeatingGroup
func NewNoComplexEventsRepeatingGroup() NoComplexEventsRepeatingGroup {
	return NoComplexEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventType), quickfix.GroupElement(tag.ComplexOptPayoutAmount), quickfix.GroupElement(tag.ComplexEventPrice), quickfix.GroupElement(tag.ComplexEventPriceBoundaryMethod), quickfix.GroupElement(tag.ComplexEventPriceBoundaryPrecision), quickfix.GroupElement(tag.ComplexEventPriceTimeType), quickfix.GroupElement(tag.ComplexEventCondition), quickfix.GroupElement(tag.NoComplexEventDates)})}
}

//Add create and append a new NoComplexEvents to this group
func (m NoComplexEventsRepeatingGroup) Add() NoComplexEvents {
	g := m.RepeatingGroup.Add()
	return NoComplexEvents{g}
}

//Get returns the ith NoComplexEvents in the NoComplexEventsRepeatinGroup
func (m NoComplexEventsRepeatingGroup) Get(i int) NoComplexEvents {
	return NoComplexEvents{m.RepeatingGroup.Get(i)}
}
