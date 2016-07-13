package multilegordercancelreplace

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//MultilegOrderCancelReplace is the fix50 MultilegOrderCancelReplace type, MsgType = AC
type MultilegOrderCancelReplace struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a MultilegOrderCancelReplace from a quickfix.Message instance
func FromMessage(m quickfix.Message) MultilegOrderCancelReplace {
	return MultilegOrderCancelReplace{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m MultilegOrderCancelReplace) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a MultilegOrderCancelReplace initialized with the required fields for MultilegOrderCancelReplace
func New(origclordid field.OrigClOrdIDField, clordid field.ClOrdIDField, side field.SideField, transacttime field.TransactTimeField, ordtype field.OrdTypeField) (m MultilegOrderCancelReplace) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("AC"))
	m.Header.Set(field.NewBeginString("7"))
	m.Set(origclordid)
	m.Set(clordid)
	m.Set(side)
	m.Set(transacttime)
	m.Set(ordtype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg MultilegOrderCancelReplace, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "7", "AC", r
}

//SetAccount sets Account, Tag 1
func (m MultilegOrderCancelReplace) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m MultilegOrderCancelReplace) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetCommission sets Commission, Tag 12
func (m MultilegOrderCancelReplace) SetCommission(v float64) {
	m.Set(field.NewCommission(v))
}

//SetCommType sets CommType, Tag 13
func (m MultilegOrderCancelReplace) SetCommType(v string) {
	m.Set(field.NewCommType(v))
}

//SetCurrency sets Currency, Tag 15
func (m MultilegOrderCancelReplace) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetExecInst sets ExecInst, Tag 18
func (m MultilegOrderCancelReplace) SetExecInst(v string) {
	m.Set(field.NewExecInst(v))
}

//SetHandlInst sets HandlInst, Tag 21
func (m MultilegOrderCancelReplace) SetHandlInst(v string) {
	m.Set(field.NewHandlInst(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m MultilegOrderCancelReplace) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetIOIID sets IOIID, Tag 23
func (m MultilegOrderCancelReplace) SetIOIID(v string) {
	m.Set(field.NewIOIID(v))
}

//SetOrderID sets OrderID, Tag 37
func (m MultilegOrderCancelReplace) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m MultilegOrderCancelReplace) SetOrderQty(v float64) {
	m.Set(field.NewOrderQty(v))
}

//SetOrdType sets OrdType, Tag 40
func (m MultilegOrderCancelReplace) SetOrdType(v string) {
	m.Set(field.NewOrdType(v))
}

//SetOrigClOrdID sets OrigClOrdID, Tag 41
func (m MultilegOrderCancelReplace) SetOrigClOrdID(v string) {
	m.Set(field.NewOrigClOrdID(v))
}

//SetPrice sets Price, Tag 44
func (m MultilegOrderCancelReplace) SetPrice(v float64) {
	m.Set(field.NewPrice(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m MultilegOrderCancelReplace) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m MultilegOrderCancelReplace) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m MultilegOrderCancelReplace) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m MultilegOrderCancelReplace) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTimeInForce sets TimeInForce, Tag 59
func (m MultilegOrderCancelReplace) SetTimeInForce(v string) {
	m.Set(field.NewTimeInForce(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m MultilegOrderCancelReplace) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSettlType sets SettlType, Tag 63
func (m MultilegOrderCancelReplace) SetSettlType(v string) {
	m.Set(field.NewSettlType(v))
}

//SetSettlDate sets SettlDate, Tag 64
func (m MultilegOrderCancelReplace) SetSettlDate(v string) {
	m.Set(field.NewSettlDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m MultilegOrderCancelReplace) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetAllocID sets AllocID, Tag 70
func (m MultilegOrderCancelReplace) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m MultilegOrderCancelReplace) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetPositionEffect sets PositionEffect, Tag 77
func (m MultilegOrderCancelReplace) SetPositionEffect(v string) {
	m.Set(field.NewPositionEffect(v))
}

//SetNoAllocs sets NoAllocs, Tag 78
func (m MultilegOrderCancelReplace) SetNoAllocs(f NoAllocsRepeatingGroup) {
	m.SetGroup(f)
}

//SetProcessCode sets ProcessCode, Tag 81
func (m MultilegOrderCancelReplace) SetProcessCode(v string) {
	m.Set(field.NewProcessCode(v))
}

//SetStopPx sets StopPx, Tag 99
func (m MultilegOrderCancelReplace) SetStopPx(v float64) {
	m.Set(field.NewStopPx(v))
}

//SetExDestination sets ExDestination, Tag 100
func (m MultilegOrderCancelReplace) SetExDestination(v string) {
	m.Set(field.NewExDestination(v))
}

//SetIssuer sets Issuer, Tag 106
func (m MultilegOrderCancelReplace) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m MultilegOrderCancelReplace) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetMinQty sets MinQty, Tag 110
func (m MultilegOrderCancelReplace) SetMinQty(v float64) {
	m.Set(field.NewMinQty(v))
}

//SetMaxFloor sets MaxFloor, Tag 111
func (m MultilegOrderCancelReplace) SetMaxFloor(v float64) {
	m.Set(field.NewMaxFloor(v))
}

//SetLocateReqd sets LocateReqd, Tag 114
func (m MultilegOrderCancelReplace) SetLocateReqd(v bool) {
	m.Set(field.NewLocateReqd(v))
}

//SetQuoteID sets QuoteID, Tag 117
func (m MultilegOrderCancelReplace) SetQuoteID(v string) {
	m.Set(field.NewQuoteID(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m MultilegOrderCancelReplace) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetForexReq sets ForexReq, Tag 121
func (m MultilegOrderCancelReplace) SetForexReq(v bool) {
	m.Set(field.NewForexReq(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m MultilegOrderCancelReplace) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetPrevClosePx sets PrevClosePx, Tag 140
func (m MultilegOrderCancelReplace) SetPrevClosePx(v float64) {
	m.Set(field.NewPrevClosePx(v))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m MultilegOrderCancelReplace) SetCashOrderQty(v float64) {
	m.Set(field.NewCashOrderQty(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m MultilegOrderCancelReplace) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetEffectiveTime sets EffectiveTime, Tag 168
func (m MultilegOrderCancelReplace) SetEffectiveTime(v time.Time) {
	m.Set(field.NewEffectiveTime(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m MultilegOrderCancelReplace) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m MultilegOrderCancelReplace) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetCoveredOrUncovered sets CoveredOrUncovered, Tag 203
func (m MultilegOrderCancelReplace) SetCoveredOrUncovered(v int) {
	m.Set(field.NewCoveredOrUncovered(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m MultilegOrderCancelReplace) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m MultilegOrderCancelReplace) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetMaxShow sets MaxShow, Tag 210
func (m MultilegOrderCancelReplace) SetMaxShow(v float64) {
	m.Set(field.NewMaxShow(v))
}

//SetPegOffsetValue sets PegOffsetValue, Tag 211
func (m MultilegOrderCancelReplace) SetPegOffsetValue(v float64) {
	m.Set(field.NewPegOffsetValue(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m MultilegOrderCancelReplace) SetCouponRate(v float64) {
	m.Set(field.NewCouponRate(v))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m MultilegOrderCancelReplace) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m MultilegOrderCancelReplace) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m MultilegOrderCancelReplace) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m MultilegOrderCancelReplace) SetRepurchaseRate(v float64) {
	m.Set(field.NewRepurchaseRate(v))
}

//SetFactor sets Factor, Tag 228
func (m MultilegOrderCancelReplace) SetFactor(v float64) {
	m.Set(field.NewFactor(v))
}

//SetTradeOriginationDate sets TradeOriginationDate, Tag 229
func (m MultilegOrderCancelReplace) SetTradeOriginationDate(v string) {
	m.Set(field.NewTradeOriginationDate(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m MultilegOrderCancelReplace) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m MultilegOrderCancelReplace) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m MultilegOrderCancelReplace) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m MultilegOrderCancelReplace) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m MultilegOrderCancelReplace) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m MultilegOrderCancelReplace) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m MultilegOrderCancelReplace) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m MultilegOrderCancelReplace) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m MultilegOrderCancelReplace) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m MultilegOrderCancelReplace) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetComplianceID sets ComplianceID, Tag 376
func (m MultilegOrderCancelReplace) SetComplianceID(v string) {
	m.Set(field.NewComplianceID(v))
}

//SetSolicitedFlag sets SolicitedFlag, Tag 377
func (m MultilegOrderCancelReplace) SetSolicitedFlag(v bool) {
	m.Set(field.NewSolicitedFlag(v))
}

//SetNoTradingSessions sets NoTradingSessions, Tag 386
func (m MultilegOrderCancelReplace) SetNoTradingSessions(f NoTradingSessionsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDiscretionInst sets DiscretionInst, Tag 388
func (m MultilegOrderCancelReplace) SetDiscretionInst(v string) {
	m.Set(field.NewDiscretionInst(v))
}

//SetDiscretionOffsetValue sets DiscretionOffsetValue, Tag 389
func (m MultilegOrderCancelReplace) SetDiscretionOffsetValue(v float64) {
	m.Set(field.NewDiscretionOffsetValue(v))
}

//SetPriceType sets PriceType, Tag 423
func (m MultilegOrderCancelReplace) SetPriceType(v int) {
	m.Set(field.NewPriceType(v))
}

//SetGTBookingInst sets GTBookingInst, Tag 427
func (m MultilegOrderCancelReplace) SetGTBookingInst(v int) {
	m.Set(field.NewGTBookingInst(v))
}

//SetExpireDate sets ExpireDate, Tag 432
func (m MultilegOrderCancelReplace) SetExpireDate(v string) {
	m.Set(field.NewExpireDate(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m MultilegOrderCancelReplace) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m MultilegOrderCancelReplace) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m MultilegOrderCancelReplace) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m MultilegOrderCancelReplace) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetRoundingDirection sets RoundingDirection, Tag 468
func (m MultilegOrderCancelReplace) SetRoundingDirection(v string) {
	m.Set(field.NewRoundingDirection(v))
}

//SetRoundingModulus sets RoundingModulus, Tag 469
func (m MultilegOrderCancelReplace) SetRoundingModulus(v float64) {
	m.Set(field.NewRoundingModulus(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m MultilegOrderCancelReplace) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m MultilegOrderCancelReplace) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m MultilegOrderCancelReplace) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetCommCurrency sets CommCurrency, Tag 479
func (m MultilegOrderCancelReplace) SetCommCurrency(v string) {
	m.Set(field.NewCommCurrency(v))
}

//SetCancellationRights sets CancellationRights, Tag 480
func (m MultilegOrderCancelReplace) SetCancellationRights(v string) {
	m.Set(field.NewCancellationRights(v))
}

//SetMoneyLaunderingStatus sets MoneyLaunderingStatus, Tag 481
func (m MultilegOrderCancelReplace) SetMoneyLaunderingStatus(v string) {
	m.Set(field.NewMoneyLaunderingStatus(v))
}

//SetDesignation sets Designation, Tag 494
func (m MultilegOrderCancelReplace) SetDesignation(v string) {
	m.Set(field.NewDesignation(v))
}

//SetFundRenewWaiv sets FundRenewWaiv, Tag 497
func (m MultilegOrderCancelReplace) SetFundRenewWaiv(v string) {
	m.Set(field.NewFundRenewWaiv(v))
}

//SetRegistID sets RegistID, Tag 513
func (m MultilegOrderCancelReplace) SetRegistID(v string) {
	m.Set(field.NewRegistID(v))
}

//SetOrderPercent sets OrderPercent, Tag 516
func (m MultilegOrderCancelReplace) SetOrderPercent(v float64) {
	m.Set(field.NewOrderPercent(v))
}

//SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526
func (m MultilegOrderCancelReplace) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

//SetOrderCapacity sets OrderCapacity, Tag 528
func (m MultilegOrderCancelReplace) SetOrderCapacity(v string) {
	m.Set(field.NewOrderCapacity(v))
}

//SetOrderRestrictions sets OrderRestrictions, Tag 529
func (m MultilegOrderCancelReplace) SetOrderRestrictions(v string) {
	m.Set(field.NewOrderRestrictions(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m MultilegOrderCancelReplace) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m MultilegOrderCancelReplace) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetCashMargin sets CashMargin, Tag 544
func (m MultilegOrderCancelReplace) SetCashMargin(v string) {
	m.Set(field.NewCashMargin(v))
}

//SetNoLegs sets NoLegs, Tag 555
func (m MultilegOrderCancelReplace) SetNoLegs(f NoLegsRepeatingGroup) {
	m.SetGroup(f)
}

//SetMultiLegRptTypeReq sets MultiLegRptTypeReq, Tag 563
func (m MultilegOrderCancelReplace) SetMultiLegRptTypeReq(v int) {
	m.Set(field.NewMultiLegRptTypeReq(v))
}

//SetAccountType sets AccountType, Tag 581
func (m MultilegOrderCancelReplace) SetAccountType(v int) {
	m.Set(field.NewAccountType(v))
}

//SetCustOrderCapacity sets CustOrderCapacity, Tag 582
func (m MultilegOrderCancelReplace) SetCustOrderCapacity(v int) {
	m.Set(field.NewCustOrderCapacity(v))
}

//SetClOrdLinkID sets ClOrdLinkID, Tag 583
func (m MultilegOrderCancelReplace) SetClOrdLinkID(v string) {
	m.Set(field.NewClOrdLinkID(v))
}

//SetOrigOrdModTime sets OrigOrdModTime, Tag 586
func (m MultilegOrderCancelReplace) SetOrigOrdModTime(v time.Time) {
	m.Set(field.NewOrigOrdModTime(v))
}

//SetDayBookingInst sets DayBookingInst, Tag 589
func (m MultilegOrderCancelReplace) SetDayBookingInst(v string) {
	m.Set(field.NewDayBookingInst(v))
}

//SetBookingUnit sets BookingUnit, Tag 590
func (m MultilegOrderCancelReplace) SetBookingUnit(v string) {
	m.Set(field.NewBookingUnit(v))
}

//SetPreallocMethod sets PreallocMethod, Tag 591
func (m MultilegOrderCancelReplace) SetPreallocMethod(v string) {
	m.Set(field.NewPreallocMethod(v))
}

//SetClearingFeeIndicator sets ClearingFeeIndicator, Tag 635
func (m MultilegOrderCancelReplace) SetClearingFeeIndicator(v string) {
	m.Set(field.NewClearingFeeIndicator(v))
}

//SetAcctIDSource sets AcctIDSource, Tag 660
func (m MultilegOrderCancelReplace) SetAcctIDSource(v int) {
	m.Set(field.NewAcctIDSource(v))
}

//SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m MultilegOrderCancelReplace) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

//SetPool sets Pool, Tag 691
func (m MultilegOrderCancelReplace) SetPool(v string) {
	m.Set(field.NewPool(v))
}

//SetNoUnderlyings sets NoUnderlyings, Tag 711
func (m MultilegOrderCancelReplace) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m MultilegOrderCancelReplace) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

//SetBookingType sets BookingType, Tag 775
func (m MultilegOrderCancelReplace) SetBookingType(v int) {
	m.Set(field.NewBookingType(v))
}

//SetPegMoveType sets PegMoveType, Tag 835
func (m MultilegOrderCancelReplace) SetPegMoveType(v int) {
	m.Set(field.NewPegMoveType(v))
}

//SetPegOffsetType sets PegOffsetType, Tag 836
func (m MultilegOrderCancelReplace) SetPegOffsetType(v int) {
	m.Set(field.NewPegOffsetType(v))
}

//SetPegLimitType sets PegLimitType, Tag 837
func (m MultilegOrderCancelReplace) SetPegLimitType(v int) {
	m.Set(field.NewPegLimitType(v))
}

//SetPegRoundDirection sets PegRoundDirection, Tag 838
func (m MultilegOrderCancelReplace) SetPegRoundDirection(v int) {
	m.Set(field.NewPegRoundDirection(v))
}

//SetPegScope sets PegScope, Tag 840
func (m MultilegOrderCancelReplace) SetPegScope(v int) {
	m.Set(field.NewPegScope(v))
}

//SetDiscretionMoveType sets DiscretionMoveType, Tag 841
func (m MultilegOrderCancelReplace) SetDiscretionMoveType(v int) {
	m.Set(field.NewDiscretionMoveType(v))
}

//SetDiscretionOffsetType sets DiscretionOffsetType, Tag 842
func (m MultilegOrderCancelReplace) SetDiscretionOffsetType(v int) {
	m.Set(field.NewDiscretionOffsetType(v))
}

//SetDiscretionLimitType sets DiscretionLimitType, Tag 843
func (m MultilegOrderCancelReplace) SetDiscretionLimitType(v int) {
	m.Set(field.NewDiscretionLimitType(v))
}

//SetDiscretionRoundDirection sets DiscretionRoundDirection, Tag 844
func (m MultilegOrderCancelReplace) SetDiscretionRoundDirection(v int) {
	m.Set(field.NewDiscretionRoundDirection(v))
}

//SetDiscretionScope sets DiscretionScope, Tag 846
func (m MultilegOrderCancelReplace) SetDiscretionScope(v int) {
	m.Set(field.NewDiscretionScope(v))
}

//SetTargetStrategy sets TargetStrategy, Tag 847
func (m MultilegOrderCancelReplace) SetTargetStrategy(v int) {
	m.Set(field.NewTargetStrategy(v))
}

//SetTargetStrategyParameters sets TargetStrategyParameters, Tag 848
func (m MultilegOrderCancelReplace) SetTargetStrategyParameters(v string) {
	m.Set(field.NewTargetStrategyParameters(v))
}

//SetParticipationRate sets ParticipationRate, Tag 849
func (m MultilegOrderCancelReplace) SetParticipationRate(v float64) {
	m.Set(field.NewParticipationRate(v))
}

//SetQtyType sets QtyType, Tag 854
func (m MultilegOrderCancelReplace) SetQtyType(v int) {
	m.Set(field.NewQtyType(v))
}

//SetNoEvents sets NoEvents, Tag 864
func (m MultilegOrderCancelReplace) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDatedDate sets DatedDate, Tag 873
func (m MultilegOrderCancelReplace) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

//SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m MultilegOrderCancelReplace) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

//SetCPProgram sets CPProgram, Tag 875
func (m MultilegOrderCancelReplace) SetCPProgram(v int) {
	m.Set(field.NewCPProgram(v))
}

//SetCPRegType sets CPRegType, Tag 876
func (m MultilegOrderCancelReplace) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

//SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m MultilegOrderCancelReplace) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

//SetNoStrategyParameters sets NoStrategyParameters, Tag 957
func (m MultilegOrderCancelReplace) SetNoStrategyParameters(f NoStrategyParametersRepeatingGroup) {
	m.SetGroup(f)
}

//SetSecurityStatus sets SecurityStatus, Tag 965
func (m MultilegOrderCancelReplace) SetSecurityStatus(v string) {
	m.Set(field.NewSecurityStatus(v))
}

//SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966
func (m MultilegOrderCancelReplace) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

//SetStrikeMultiplier sets StrikeMultiplier, Tag 967
func (m MultilegOrderCancelReplace) SetStrikeMultiplier(v float64) {
	m.Set(field.NewStrikeMultiplier(v))
}

//SetStrikeValue sets StrikeValue, Tag 968
func (m MultilegOrderCancelReplace) SetStrikeValue(v float64) {
	m.Set(field.NewStrikeValue(v))
}

//SetMinPriceIncrement sets MinPriceIncrement, Tag 969
func (m MultilegOrderCancelReplace) SetMinPriceIncrement(v float64) {
	m.Set(field.NewMinPriceIncrement(v))
}

//SetPositionLimit sets PositionLimit, Tag 970
func (m MultilegOrderCancelReplace) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

//SetNTPositionLimit sets NTPositionLimit, Tag 971
func (m MultilegOrderCancelReplace) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

//SetUnitOfMeasure sets UnitOfMeasure, Tag 996
func (m MultilegOrderCancelReplace) SetUnitOfMeasure(v string) {
	m.Set(field.NewUnitOfMeasure(v))
}

//SetTimeUnit sets TimeUnit, Tag 997
func (m MultilegOrderCancelReplace) SetTimeUnit(v string) {
	m.Set(field.NewTimeUnit(v))
}

//SetNoInstrumentParties sets NoInstrumentParties, Tag 1018
func (m MultilegOrderCancelReplace) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049
func (m MultilegOrderCancelReplace) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

//SetSwapPoints sets SwapPoints, Tag 1069
func (m MultilegOrderCancelReplace) SetSwapPoints(v float64) {
	m.Set(field.NewSwapPoints(v))
}

//SetMaturityTime sets MaturityTime, Tag 1079
func (m MultilegOrderCancelReplace) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

//SetSecondaryDisplayQty sets SecondaryDisplayQty, Tag 1082
func (m MultilegOrderCancelReplace) SetSecondaryDisplayQty(v float64) {
	m.Set(field.NewSecondaryDisplayQty(v))
}

//SetDisplayWhen sets DisplayWhen, Tag 1083
func (m MultilegOrderCancelReplace) SetDisplayWhen(v string) {
	m.Set(field.NewDisplayWhen(v))
}

//SetDisplayMethod sets DisplayMethod, Tag 1084
func (m MultilegOrderCancelReplace) SetDisplayMethod(v string) {
	m.Set(field.NewDisplayMethod(v))
}

//SetDisplayLowQty sets DisplayLowQty, Tag 1085
func (m MultilegOrderCancelReplace) SetDisplayLowQty(v float64) {
	m.Set(field.NewDisplayLowQty(v))
}

//SetDisplayHighQty sets DisplayHighQty, Tag 1086
func (m MultilegOrderCancelReplace) SetDisplayHighQty(v float64) {
	m.Set(field.NewDisplayHighQty(v))
}

//SetDisplayMinIncr sets DisplayMinIncr, Tag 1087
func (m MultilegOrderCancelReplace) SetDisplayMinIncr(v float64) {
	m.Set(field.NewDisplayMinIncr(v))
}

//SetRefreshQty sets RefreshQty, Tag 1088
func (m MultilegOrderCancelReplace) SetRefreshQty(v float64) {
	m.Set(field.NewRefreshQty(v))
}

//SetMatchIncrement sets MatchIncrement, Tag 1089
func (m MultilegOrderCancelReplace) SetMatchIncrement(v float64) {
	m.Set(field.NewMatchIncrement(v))
}

//SetMaxPriceLevels sets MaxPriceLevels, Tag 1090
func (m MultilegOrderCancelReplace) SetMaxPriceLevels(v int) {
	m.Set(field.NewMaxPriceLevels(v))
}

//SetPreTradeAnonymity sets PreTradeAnonymity, Tag 1091
func (m MultilegOrderCancelReplace) SetPreTradeAnonymity(v bool) {
	m.Set(field.NewPreTradeAnonymity(v))
}

//SetPriceProtectionScope sets PriceProtectionScope, Tag 1092
func (m MultilegOrderCancelReplace) SetPriceProtectionScope(v string) {
	m.Set(field.NewPriceProtectionScope(v))
}

//SetPegPriceType sets PegPriceType, Tag 1094
func (m MultilegOrderCancelReplace) SetPegPriceType(v int) {
	m.Set(field.NewPegPriceType(v))
}

//SetPegSecurityIDSource sets PegSecurityIDSource, Tag 1096
func (m MultilegOrderCancelReplace) SetPegSecurityIDSource(v string) {
	m.Set(field.NewPegSecurityIDSource(v))
}

//SetPegSecurityID sets PegSecurityID, Tag 1097
func (m MultilegOrderCancelReplace) SetPegSecurityID(v string) {
	m.Set(field.NewPegSecurityID(v))
}

//SetPegSymbol sets PegSymbol, Tag 1098
func (m MultilegOrderCancelReplace) SetPegSymbol(v string) {
	m.Set(field.NewPegSymbol(v))
}

//SetPegSecurityDesc sets PegSecurityDesc, Tag 1099
func (m MultilegOrderCancelReplace) SetPegSecurityDesc(v string) {
	m.Set(field.NewPegSecurityDesc(v))
}

//SetTriggerType sets TriggerType, Tag 1100
func (m MultilegOrderCancelReplace) SetTriggerType(v string) {
	m.Set(field.NewTriggerType(v))
}

//SetTriggerAction sets TriggerAction, Tag 1101
func (m MultilegOrderCancelReplace) SetTriggerAction(v string) {
	m.Set(field.NewTriggerAction(v))
}

//SetTriggerPrice sets TriggerPrice, Tag 1102
func (m MultilegOrderCancelReplace) SetTriggerPrice(v float64) {
	m.Set(field.NewTriggerPrice(v))
}

//SetTriggerSymbol sets TriggerSymbol, Tag 1103
func (m MultilegOrderCancelReplace) SetTriggerSymbol(v string) {
	m.Set(field.NewTriggerSymbol(v))
}

//SetTriggerSecurityID sets TriggerSecurityID, Tag 1104
func (m MultilegOrderCancelReplace) SetTriggerSecurityID(v string) {
	m.Set(field.NewTriggerSecurityID(v))
}

//SetTriggerSecurityIDSource sets TriggerSecurityIDSource, Tag 1105
func (m MultilegOrderCancelReplace) SetTriggerSecurityIDSource(v string) {
	m.Set(field.NewTriggerSecurityIDSource(v))
}

//SetTriggerSecurityDesc sets TriggerSecurityDesc, Tag 1106
func (m MultilegOrderCancelReplace) SetTriggerSecurityDesc(v string) {
	m.Set(field.NewTriggerSecurityDesc(v))
}

//SetTriggerPriceType sets TriggerPriceType, Tag 1107
func (m MultilegOrderCancelReplace) SetTriggerPriceType(v string) {
	m.Set(field.NewTriggerPriceType(v))
}

//SetTriggerPriceTypeScope sets TriggerPriceTypeScope, Tag 1108
func (m MultilegOrderCancelReplace) SetTriggerPriceTypeScope(v string) {
	m.Set(field.NewTriggerPriceTypeScope(v))
}

//SetTriggerPriceDirection sets TriggerPriceDirection, Tag 1109
func (m MultilegOrderCancelReplace) SetTriggerPriceDirection(v string) {
	m.Set(field.NewTriggerPriceDirection(v))
}

//SetTriggerNewPrice sets TriggerNewPrice, Tag 1110
func (m MultilegOrderCancelReplace) SetTriggerNewPrice(v float64) {
	m.Set(field.NewTriggerNewPrice(v))
}

//SetTriggerOrderType sets TriggerOrderType, Tag 1111
func (m MultilegOrderCancelReplace) SetTriggerOrderType(v string) {
	m.Set(field.NewTriggerOrderType(v))
}

//SetTriggerNewQty sets TriggerNewQty, Tag 1112
func (m MultilegOrderCancelReplace) SetTriggerNewQty(v float64) {
	m.Set(field.NewTriggerNewQty(v))
}

//SetTriggerTradingSessionID sets TriggerTradingSessionID, Tag 1113
func (m MultilegOrderCancelReplace) SetTriggerTradingSessionID(v string) {
	m.Set(field.NewTriggerTradingSessionID(v))
}

//SetTriggerTradingSessionSubID sets TriggerTradingSessionSubID, Tag 1114
func (m MultilegOrderCancelReplace) SetTriggerTradingSessionSubID(v string) {
	m.Set(field.NewTriggerTradingSessionSubID(v))
}

//SetExDestinationIDSource sets ExDestinationIDSource, Tag 1133
func (m MultilegOrderCancelReplace) SetExDestinationIDSource(v string) {
	m.Set(field.NewExDestinationIDSource(v))
}

//SetDisplayQty sets DisplayQty, Tag 1138
func (m MultilegOrderCancelReplace) SetDisplayQty(v float64) {
	m.Set(field.NewDisplayQty(v))
}

//GetAccount gets Account, Tag 1
func (m MultilegOrderCancelReplace) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m MultilegOrderCancelReplace) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommission gets Commission, Tag 12
func (m MultilegOrderCancelReplace) GetCommission() (f field.CommissionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommType gets CommType, Tag 13
func (m MultilegOrderCancelReplace) GetCommType() (f field.CommTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m MultilegOrderCancelReplace) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecInst gets ExecInst, Tag 18
func (m MultilegOrderCancelReplace) GetExecInst() (f field.ExecInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetHandlInst gets HandlInst, Tag 21
func (m MultilegOrderCancelReplace) GetHandlInst() (f field.HandlInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m MultilegOrderCancelReplace) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIID gets IOIID, Tag 23
func (m MultilegOrderCancelReplace) GetIOIID() (f field.IOIIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m MultilegOrderCancelReplace) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m MultilegOrderCancelReplace) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdType gets OrdType, Tag 40
func (m MultilegOrderCancelReplace) GetOrdType() (f field.OrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigClOrdID gets OrigClOrdID, Tag 41
func (m MultilegOrderCancelReplace) GetOrigClOrdID() (f field.OrigClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m MultilegOrderCancelReplace) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m MultilegOrderCancelReplace) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m MultilegOrderCancelReplace) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m MultilegOrderCancelReplace) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m MultilegOrderCancelReplace) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m MultilegOrderCancelReplace) GetTimeInForce() (f field.TimeInForceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m MultilegOrderCancelReplace) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlType gets SettlType, Tag 63
func (m MultilegOrderCancelReplace) GetSettlType() (f field.SettlTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlDate gets SettlDate, Tag 64
func (m MultilegOrderCancelReplace) GetSettlDate() (f field.SettlDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m MultilegOrderCancelReplace) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocID gets AllocID, Tag 70
func (m MultilegOrderCancelReplace) GetAllocID() (f field.AllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m MultilegOrderCancelReplace) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionEffect gets PositionEffect, Tag 77
func (m MultilegOrderCancelReplace) GetPositionEffect() (f field.PositionEffectField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoAllocs gets NoAllocs, Tag 78
func (m MultilegOrderCancelReplace) GetNoAllocs() (NoAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProcessCode gets ProcessCode, Tag 81
func (m MultilegOrderCancelReplace) GetProcessCode() (f field.ProcessCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStopPx gets StopPx, Tag 99
func (m MultilegOrderCancelReplace) GetStopPx() (f field.StopPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExDestination gets ExDestination, Tag 100
func (m MultilegOrderCancelReplace) GetExDestination() (f field.ExDestinationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m MultilegOrderCancelReplace) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m MultilegOrderCancelReplace) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinQty gets MinQty, Tag 110
func (m MultilegOrderCancelReplace) GetMinQty() (f field.MinQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxFloor gets MaxFloor, Tag 111
func (m MultilegOrderCancelReplace) GetMaxFloor() (f field.MaxFloorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocateReqd gets LocateReqd, Tag 114
func (m MultilegOrderCancelReplace) GetLocateReqd() (f field.LocateReqdField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteID gets QuoteID, Tag 117
func (m MultilegOrderCancelReplace) GetQuoteID() (f field.QuoteIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m MultilegOrderCancelReplace) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetForexReq gets ForexReq, Tag 121
func (m MultilegOrderCancelReplace) GetForexReq() (f field.ForexReqField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m MultilegOrderCancelReplace) GetExpireTime() (f field.ExpireTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrevClosePx gets PrevClosePx, Tag 140
func (m MultilegOrderCancelReplace) GetPrevClosePx() (f field.PrevClosePxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m MultilegOrderCancelReplace) GetCashOrderQty() (f field.CashOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m MultilegOrderCancelReplace) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEffectiveTime gets EffectiveTime, Tag 168
func (m MultilegOrderCancelReplace) GetEffectiveTime() (f field.EffectiveTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m MultilegOrderCancelReplace) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m MultilegOrderCancelReplace) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCoveredOrUncovered gets CoveredOrUncovered, Tag 203
func (m MultilegOrderCancelReplace) GetCoveredOrUncovered() (f field.CoveredOrUncoveredField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m MultilegOrderCancelReplace) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m MultilegOrderCancelReplace) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxShow gets MaxShow, Tag 210
func (m MultilegOrderCancelReplace) GetMaxShow() (f field.MaxShowField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegOffsetValue gets PegOffsetValue, Tag 211
func (m MultilegOrderCancelReplace) GetPegOffsetValue() (f field.PegOffsetValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m MultilegOrderCancelReplace) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m MultilegOrderCancelReplace) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m MultilegOrderCancelReplace) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m MultilegOrderCancelReplace) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m MultilegOrderCancelReplace) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m MultilegOrderCancelReplace) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeOriginationDate gets TradeOriginationDate, Tag 229
func (m MultilegOrderCancelReplace) GetTradeOriginationDate() (f field.TradeOriginationDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m MultilegOrderCancelReplace) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m MultilegOrderCancelReplace) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m MultilegOrderCancelReplace) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m MultilegOrderCancelReplace) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m MultilegOrderCancelReplace) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m MultilegOrderCancelReplace) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m MultilegOrderCancelReplace) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m MultilegOrderCancelReplace) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m MultilegOrderCancelReplace) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m MultilegOrderCancelReplace) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplianceID gets ComplianceID, Tag 376
func (m MultilegOrderCancelReplace) GetComplianceID() (f field.ComplianceIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSolicitedFlag gets SolicitedFlag, Tag 377
func (m MultilegOrderCancelReplace) GetSolicitedFlag() (f field.SolicitedFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoTradingSessions gets NoTradingSessions, Tag 386
func (m MultilegOrderCancelReplace) GetNoTradingSessions() (NoTradingSessionsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTradingSessionsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDiscretionInst gets DiscretionInst, Tag 388
func (m MultilegOrderCancelReplace) GetDiscretionInst() (f field.DiscretionInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionOffsetValue gets DiscretionOffsetValue, Tag 389
func (m MultilegOrderCancelReplace) GetDiscretionOffsetValue() (f field.DiscretionOffsetValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceType gets PriceType, Tag 423
func (m MultilegOrderCancelReplace) GetPriceType() (f field.PriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetGTBookingInst gets GTBookingInst, Tag 427
func (m MultilegOrderCancelReplace) GetGTBookingInst() (f field.GTBookingInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireDate gets ExpireDate, Tag 432
func (m MultilegOrderCancelReplace) GetExpireDate() (f field.ExpireDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m MultilegOrderCancelReplace) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m MultilegOrderCancelReplace) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m MultilegOrderCancelReplace) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m MultilegOrderCancelReplace) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingDirection gets RoundingDirection, Tag 468
func (m MultilegOrderCancelReplace) GetRoundingDirection() (f field.RoundingDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingModulus gets RoundingModulus, Tag 469
func (m MultilegOrderCancelReplace) GetRoundingModulus() (f field.RoundingModulusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m MultilegOrderCancelReplace) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m MultilegOrderCancelReplace) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m MultilegOrderCancelReplace) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommCurrency gets CommCurrency, Tag 479
func (m MultilegOrderCancelReplace) GetCommCurrency() (f field.CommCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCancellationRights gets CancellationRights, Tag 480
func (m MultilegOrderCancelReplace) GetCancellationRights() (f field.CancellationRightsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMoneyLaunderingStatus gets MoneyLaunderingStatus, Tag 481
func (m MultilegOrderCancelReplace) GetMoneyLaunderingStatus() (f field.MoneyLaunderingStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDesignation gets Designation, Tag 494
func (m MultilegOrderCancelReplace) GetDesignation() (f field.DesignationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFundRenewWaiv gets FundRenewWaiv, Tag 497
func (m MultilegOrderCancelReplace) GetFundRenewWaiv() (f field.FundRenewWaivField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRegistID gets RegistID, Tag 513
func (m MultilegOrderCancelReplace) GetRegistID() (f field.RegistIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderPercent gets OrderPercent, Tag 516
func (m MultilegOrderCancelReplace) GetOrderPercent() (f field.OrderPercentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526
func (m MultilegOrderCancelReplace) GetSecondaryClOrdID() (f field.SecondaryClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderCapacity gets OrderCapacity, Tag 528
func (m MultilegOrderCancelReplace) GetOrderCapacity() (f field.OrderCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderRestrictions gets OrderRestrictions, Tag 529
func (m MultilegOrderCancelReplace) GetOrderRestrictions() (f field.OrderRestrictionsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m MultilegOrderCancelReplace) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m MultilegOrderCancelReplace) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashMargin gets CashMargin, Tag 544
func (m MultilegOrderCancelReplace) GetCashMargin() (f field.CashMarginField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoLegs gets NoLegs, Tag 555
func (m MultilegOrderCancelReplace) GetNoLegs() (NoLegsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetMultiLegRptTypeReq gets MultiLegRptTypeReq, Tag 563
func (m MultilegOrderCancelReplace) GetMultiLegRptTypeReq() (f field.MultiLegRptTypeReqField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccountType gets AccountType, Tag 581
func (m MultilegOrderCancelReplace) GetAccountType() (f field.AccountTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCustOrderCapacity gets CustOrderCapacity, Tag 582
func (m MultilegOrderCancelReplace) GetCustOrderCapacity() (f field.CustOrderCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdLinkID gets ClOrdLinkID, Tag 583
func (m MultilegOrderCancelReplace) GetClOrdLinkID() (f field.ClOrdLinkIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigOrdModTime gets OrigOrdModTime, Tag 586
func (m MultilegOrderCancelReplace) GetOrigOrdModTime() (f field.OrigOrdModTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDayBookingInst gets DayBookingInst, Tag 589
func (m MultilegOrderCancelReplace) GetDayBookingInst() (f field.DayBookingInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBookingUnit gets BookingUnit, Tag 590
func (m MultilegOrderCancelReplace) GetBookingUnit() (f field.BookingUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPreallocMethod gets PreallocMethod, Tag 591
func (m MultilegOrderCancelReplace) GetPreallocMethod() (f field.PreallocMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClearingFeeIndicator gets ClearingFeeIndicator, Tag 635
func (m MultilegOrderCancelReplace) GetClearingFeeIndicator() (f field.ClearingFeeIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAcctIDSource gets AcctIDSource, Tag 660
func (m MultilegOrderCancelReplace) GetAcctIDSource() (f field.AcctIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m MultilegOrderCancelReplace) GetContractSettlMonth() (f field.ContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPool gets Pool, Tag 691
func (m MultilegOrderCancelReplace) GetPool() (f field.PoolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyings gets NoUnderlyings, Tag 711
func (m MultilegOrderCancelReplace) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m MultilegOrderCancelReplace) GetSecuritySubType() (f field.SecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBookingType gets BookingType, Tag 775
func (m MultilegOrderCancelReplace) GetBookingType() (f field.BookingTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegMoveType gets PegMoveType, Tag 835
func (m MultilegOrderCancelReplace) GetPegMoveType() (f field.PegMoveTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegOffsetType gets PegOffsetType, Tag 836
func (m MultilegOrderCancelReplace) GetPegOffsetType() (f field.PegOffsetTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegLimitType gets PegLimitType, Tag 837
func (m MultilegOrderCancelReplace) GetPegLimitType() (f field.PegLimitTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegRoundDirection gets PegRoundDirection, Tag 838
func (m MultilegOrderCancelReplace) GetPegRoundDirection() (f field.PegRoundDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegScope gets PegScope, Tag 840
func (m MultilegOrderCancelReplace) GetPegScope() (f field.PegScopeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionMoveType gets DiscretionMoveType, Tag 841
func (m MultilegOrderCancelReplace) GetDiscretionMoveType() (f field.DiscretionMoveTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionOffsetType gets DiscretionOffsetType, Tag 842
func (m MultilegOrderCancelReplace) GetDiscretionOffsetType() (f field.DiscretionOffsetTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionLimitType gets DiscretionLimitType, Tag 843
func (m MultilegOrderCancelReplace) GetDiscretionLimitType() (f field.DiscretionLimitTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionRoundDirection gets DiscretionRoundDirection, Tag 844
func (m MultilegOrderCancelReplace) GetDiscretionRoundDirection() (f field.DiscretionRoundDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionScope gets DiscretionScope, Tag 846
func (m MultilegOrderCancelReplace) GetDiscretionScope() (f field.DiscretionScopeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTargetStrategy gets TargetStrategy, Tag 847
func (m MultilegOrderCancelReplace) GetTargetStrategy() (f field.TargetStrategyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTargetStrategyParameters gets TargetStrategyParameters, Tag 848
func (m MultilegOrderCancelReplace) GetTargetStrategyParameters() (f field.TargetStrategyParametersField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetParticipationRate gets ParticipationRate, Tag 849
func (m MultilegOrderCancelReplace) GetParticipationRate() (f field.ParticipationRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQtyType gets QtyType, Tag 854
func (m MultilegOrderCancelReplace) GetQtyType() (f field.QtyTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoEvents gets NoEvents, Tag 864
func (m MultilegOrderCancelReplace) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDatedDate gets DatedDate, Tag 873
func (m MultilegOrderCancelReplace) GetDatedDate() (f field.DatedDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m MultilegOrderCancelReplace) GetInterestAccrualDate() (f field.InterestAccrualDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPProgram gets CPProgram, Tag 875
func (m MultilegOrderCancelReplace) GetCPProgram() (f field.CPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPRegType gets CPRegType, Tag 876
func (m MultilegOrderCancelReplace) GetCPRegType() (f field.CPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m MultilegOrderCancelReplace) GetStrikeCurrency() (f field.StrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoStrategyParameters gets NoStrategyParameters, Tag 957
func (m MultilegOrderCancelReplace) GetNoStrategyParameters() (NoStrategyParametersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStrategyParametersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSecurityStatus gets SecurityStatus, Tag 965
func (m MultilegOrderCancelReplace) GetSecurityStatus() (f field.SecurityStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966
func (m MultilegOrderCancelReplace) GetSettleOnOpenFlag() (f field.SettleOnOpenFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeMultiplier gets StrikeMultiplier, Tag 967
func (m MultilegOrderCancelReplace) GetStrikeMultiplier() (f field.StrikeMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeValue gets StrikeValue, Tag 968
func (m MultilegOrderCancelReplace) GetStrikeValue() (f field.StrikeValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrement gets MinPriceIncrement, Tag 969
func (m MultilegOrderCancelReplace) GetMinPriceIncrement() (f field.MinPriceIncrementField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionLimit gets PositionLimit, Tag 970
func (m MultilegOrderCancelReplace) GetPositionLimit() (f field.PositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNTPositionLimit gets NTPositionLimit, Tag 971
func (m MultilegOrderCancelReplace) GetNTPositionLimit() (f field.NTPositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnitOfMeasure gets UnitOfMeasure, Tag 996
func (m MultilegOrderCancelReplace) GetUnitOfMeasure() (f field.UnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeUnit gets TimeUnit, Tag 997
func (m MultilegOrderCancelReplace) GetTimeUnit() (f field.TimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoInstrumentParties gets NoInstrumentParties, Tag 1018
func (m MultilegOrderCancelReplace) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049
func (m MultilegOrderCancelReplace) GetInstrmtAssignmentMethod() (f field.InstrmtAssignmentMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSwapPoints gets SwapPoints, Tag 1069
func (m MultilegOrderCancelReplace) GetSwapPoints() (f field.SwapPointsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityTime gets MaturityTime, Tag 1079
func (m MultilegOrderCancelReplace) GetMaturityTime() (f field.MaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryDisplayQty gets SecondaryDisplayQty, Tag 1082
func (m MultilegOrderCancelReplace) GetSecondaryDisplayQty() (f field.SecondaryDisplayQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDisplayWhen gets DisplayWhen, Tag 1083
func (m MultilegOrderCancelReplace) GetDisplayWhen() (f field.DisplayWhenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDisplayMethod gets DisplayMethod, Tag 1084
func (m MultilegOrderCancelReplace) GetDisplayMethod() (f field.DisplayMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDisplayLowQty gets DisplayLowQty, Tag 1085
func (m MultilegOrderCancelReplace) GetDisplayLowQty() (f field.DisplayLowQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDisplayHighQty gets DisplayHighQty, Tag 1086
func (m MultilegOrderCancelReplace) GetDisplayHighQty() (f field.DisplayHighQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDisplayMinIncr gets DisplayMinIncr, Tag 1087
func (m MultilegOrderCancelReplace) GetDisplayMinIncr() (f field.DisplayMinIncrField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRefreshQty gets RefreshQty, Tag 1088
func (m MultilegOrderCancelReplace) GetRefreshQty() (f field.RefreshQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMatchIncrement gets MatchIncrement, Tag 1089
func (m MultilegOrderCancelReplace) GetMatchIncrement() (f field.MatchIncrementField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxPriceLevels gets MaxPriceLevels, Tag 1090
func (m MultilegOrderCancelReplace) GetMaxPriceLevels() (f field.MaxPriceLevelsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPreTradeAnonymity gets PreTradeAnonymity, Tag 1091
func (m MultilegOrderCancelReplace) GetPreTradeAnonymity() (f field.PreTradeAnonymityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceProtectionScope gets PriceProtectionScope, Tag 1092
func (m MultilegOrderCancelReplace) GetPriceProtectionScope() (f field.PriceProtectionScopeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegPriceType gets PegPriceType, Tag 1094
func (m MultilegOrderCancelReplace) GetPegPriceType() (f field.PegPriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegSecurityIDSource gets PegSecurityIDSource, Tag 1096
func (m MultilegOrderCancelReplace) GetPegSecurityIDSource() (f field.PegSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegSecurityID gets PegSecurityID, Tag 1097
func (m MultilegOrderCancelReplace) GetPegSecurityID() (f field.PegSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegSymbol gets PegSymbol, Tag 1098
func (m MultilegOrderCancelReplace) GetPegSymbol() (f field.PegSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegSecurityDesc gets PegSecurityDesc, Tag 1099
func (m MultilegOrderCancelReplace) GetPegSecurityDesc() (f field.PegSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerType gets TriggerType, Tag 1100
func (m MultilegOrderCancelReplace) GetTriggerType() (f field.TriggerTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerAction gets TriggerAction, Tag 1101
func (m MultilegOrderCancelReplace) GetTriggerAction() (f field.TriggerActionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerPrice gets TriggerPrice, Tag 1102
func (m MultilegOrderCancelReplace) GetTriggerPrice() (f field.TriggerPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerSymbol gets TriggerSymbol, Tag 1103
func (m MultilegOrderCancelReplace) GetTriggerSymbol() (f field.TriggerSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerSecurityID gets TriggerSecurityID, Tag 1104
func (m MultilegOrderCancelReplace) GetTriggerSecurityID() (f field.TriggerSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerSecurityIDSource gets TriggerSecurityIDSource, Tag 1105
func (m MultilegOrderCancelReplace) GetTriggerSecurityIDSource() (f field.TriggerSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerSecurityDesc gets TriggerSecurityDesc, Tag 1106
func (m MultilegOrderCancelReplace) GetTriggerSecurityDesc() (f field.TriggerSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerPriceType gets TriggerPriceType, Tag 1107
func (m MultilegOrderCancelReplace) GetTriggerPriceType() (f field.TriggerPriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerPriceTypeScope gets TriggerPriceTypeScope, Tag 1108
func (m MultilegOrderCancelReplace) GetTriggerPriceTypeScope() (f field.TriggerPriceTypeScopeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerPriceDirection gets TriggerPriceDirection, Tag 1109
func (m MultilegOrderCancelReplace) GetTriggerPriceDirection() (f field.TriggerPriceDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerNewPrice gets TriggerNewPrice, Tag 1110
func (m MultilegOrderCancelReplace) GetTriggerNewPrice() (f field.TriggerNewPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerOrderType gets TriggerOrderType, Tag 1111
func (m MultilegOrderCancelReplace) GetTriggerOrderType() (f field.TriggerOrderTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerNewQty gets TriggerNewQty, Tag 1112
func (m MultilegOrderCancelReplace) GetTriggerNewQty() (f field.TriggerNewQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerTradingSessionID gets TriggerTradingSessionID, Tag 1113
func (m MultilegOrderCancelReplace) GetTriggerTradingSessionID() (f field.TriggerTradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTriggerTradingSessionSubID gets TriggerTradingSessionSubID, Tag 1114
func (m MultilegOrderCancelReplace) GetTriggerTradingSessionSubID() (f field.TriggerTradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExDestinationIDSource gets ExDestinationIDSource, Tag 1133
func (m MultilegOrderCancelReplace) GetExDestinationIDSource() (f field.ExDestinationIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDisplayQty gets DisplayQty, Tag 1138
func (m MultilegOrderCancelReplace) GetDisplayQty() (f field.DisplayQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAccount returns true if Account is present, Tag 1
func (m MultilegOrderCancelReplace) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m MultilegOrderCancelReplace) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasCommission returns true if Commission is present, Tag 12
func (m MultilegOrderCancelReplace) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m MultilegOrderCancelReplace) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m MultilegOrderCancelReplace) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasExecInst returns true if ExecInst is present, Tag 18
func (m MultilegOrderCancelReplace) HasExecInst() bool {
	return m.Has(tag.ExecInst)
}

//HasHandlInst returns true if HandlInst is present, Tag 21
func (m MultilegOrderCancelReplace) HasHandlInst() bool {
	return m.Has(tag.HandlInst)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m MultilegOrderCancelReplace) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasIOIID returns true if IOIID is present, Tag 23
func (m MultilegOrderCancelReplace) HasIOIID() bool {
	return m.Has(tag.IOIID)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m MultilegOrderCancelReplace) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m MultilegOrderCancelReplace) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m MultilegOrderCancelReplace) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasOrigClOrdID returns true if OrigClOrdID is present, Tag 41
func (m MultilegOrderCancelReplace) HasOrigClOrdID() bool {
	return m.Has(tag.OrigClOrdID)
}

//HasPrice returns true if Price is present, Tag 44
func (m MultilegOrderCancelReplace) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m MultilegOrderCancelReplace) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m MultilegOrderCancelReplace) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m MultilegOrderCancelReplace) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m MultilegOrderCancelReplace) HasText() bool {
	return m.Has(tag.Text)
}

//HasTimeInForce returns true if TimeInForce is present, Tag 59
func (m MultilegOrderCancelReplace) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m MultilegOrderCancelReplace) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSettlType returns true if SettlType is present, Tag 63
func (m MultilegOrderCancelReplace) HasSettlType() bool {
	return m.Has(tag.SettlType)
}

//HasSettlDate returns true if SettlDate is present, Tag 64
func (m MultilegOrderCancelReplace) HasSettlDate() bool {
	return m.Has(tag.SettlDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m MultilegOrderCancelReplace) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasAllocID returns true if AllocID is present, Tag 70
func (m MultilegOrderCancelReplace) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m MultilegOrderCancelReplace) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasPositionEffect returns true if PositionEffect is present, Tag 77
func (m MultilegOrderCancelReplace) HasPositionEffect() bool {
	return m.Has(tag.PositionEffect)
}

//HasNoAllocs returns true if NoAllocs is present, Tag 78
func (m MultilegOrderCancelReplace) HasNoAllocs() bool {
	return m.Has(tag.NoAllocs)
}

//HasProcessCode returns true if ProcessCode is present, Tag 81
func (m MultilegOrderCancelReplace) HasProcessCode() bool {
	return m.Has(tag.ProcessCode)
}

//HasStopPx returns true if StopPx is present, Tag 99
func (m MultilegOrderCancelReplace) HasStopPx() bool {
	return m.Has(tag.StopPx)
}

//HasExDestination returns true if ExDestination is present, Tag 100
func (m MultilegOrderCancelReplace) HasExDestination() bool {
	return m.Has(tag.ExDestination)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m MultilegOrderCancelReplace) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m MultilegOrderCancelReplace) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasMinQty returns true if MinQty is present, Tag 110
func (m MultilegOrderCancelReplace) HasMinQty() bool {
	return m.Has(tag.MinQty)
}

//HasMaxFloor returns true if MaxFloor is present, Tag 111
func (m MultilegOrderCancelReplace) HasMaxFloor() bool {
	return m.Has(tag.MaxFloor)
}

//HasLocateReqd returns true if LocateReqd is present, Tag 114
func (m MultilegOrderCancelReplace) HasLocateReqd() bool {
	return m.Has(tag.LocateReqd)
}

//HasQuoteID returns true if QuoteID is present, Tag 117
func (m MultilegOrderCancelReplace) HasQuoteID() bool {
	return m.Has(tag.QuoteID)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m MultilegOrderCancelReplace) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasForexReq returns true if ForexReq is present, Tag 121
func (m MultilegOrderCancelReplace) HasForexReq() bool {
	return m.Has(tag.ForexReq)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m MultilegOrderCancelReplace) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasPrevClosePx returns true if PrevClosePx is present, Tag 140
func (m MultilegOrderCancelReplace) HasPrevClosePx() bool {
	return m.Has(tag.PrevClosePx)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m MultilegOrderCancelReplace) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m MultilegOrderCancelReplace) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasEffectiveTime returns true if EffectiveTime is present, Tag 168
func (m MultilegOrderCancelReplace) HasEffectiveTime() bool {
	return m.Has(tag.EffectiveTime)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m MultilegOrderCancelReplace) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m MultilegOrderCancelReplace) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasCoveredOrUncovered returns true if CoveredOrUncovered is present, Tag 203
func (m MultilegOrderCancelReplace) HasCoveredOrUncovered() bool {
	return m.Has(tag.CoveredOrUncovered)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m MultilegOrderCancelReplace) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m MultilegOrderCancelReplace) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasMaxShow returns true if MaxShow is present, Tag 210
func (m MultilegOrderCancelReplace) HasMaxShow() bool {
	return m.Has(tag.MaxShow)
}

//HasPegOffsetValue returns true if PegOffsetValue is present, Tag 211
func (m MultilegOrderCancelReplace) HasPegOffsetValue() bool {
	return m.Has(tag.PegOffsetValue)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m MultilegOrderCancelReplace) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m MultilegOrderCancelReplace) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m MultilegOrderCancelReplace) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m MultilegOrderCancelReplace) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m MultilegOrderCancelReplace) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m MultilegOrderCancelReplace) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasTradeOriginationDate returns true if TradeOriginationDate is present, Tag 229
func (m MultilegOrderCancelReplace) HasTradeOriginationDate() bool {
	return m.Has(tag.TradeOriginationDate)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m MultilegOrderCancelReplace) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m MultilegOrderCancelReplace) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m MultilegOrderCancelReplace) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m MultilegOrderCancelReplace) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m MultilegOrderCancelReplace) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m MultilegOrderCancelReplace) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m MultilegOrderCancelReplace) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m MultilegOrderCancelReplace) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m MultilegOrderCancelReplace) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m MultilegOrderCancelReplace) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasComplianceID returns true if ComplianceID is present, Tag 376
func (m MultilegOrderCancelReplace) HasComplianceID() bool {
	return m.Has(tag.ComplianceID)
}

//HasSolicitedFlag returns true if SolicitedFlag is present, Tag 377
func (m MultilegOrderCancelReplace) HasSolicitedFlag() bool {
	return m.Has(tag.SolicitedFlag)
}

//HasNoTradingSessions returns true if NoTradingSessions is present, Tag 386
func (m MultilegOrderCancelReplace) HasNoTradingSessions() bool {
	return m.Has(tag.NoTradingSessions)
}

//HasDiscretionInst returns true if DiscretionInst is present, Tag 388
func (m MultilegOrderCancelReplace) HasDiscretionInst() bool {
	return m.Has(tag.DiscretionInst)
}

//HasDiscretionOffsetValue returns true if DiscretionOffsetValue is present, Tag 389
func (m MultilegOrderCancelReplace) HasDiscretionOffsetValue() bool {
	return m.Has(tag.DiscretionOffsetValue)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m MultilegOrderCancelReplace) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasGTBookingInst returns true if GTBookingInst is present, Tag 427
func (m MultilegOrderCancelReplace) HasGTBookingInst() bool {
	return m.Has(tag.GTBookingInst)
}

//HasExpireDate returns true if ExpireDate is present, Tag 432
func (m MultilegOrderCancelReplace) HasExpireDate() bool {
	return m.Has(tag.ExpireDate)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m MultilegOrderCancelReplace) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m MultilegOrderCancelReplace) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m MultilegOrderCancelReplace) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m MultilegOrderCancelReplace) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasRoundingDirection returns true if RoundingDirection is present, Tag 468
func (m MultilegOrderCancelReplace) HasRoundingDirection() bool {
	return m.Has(tag.RoundingDirection)
}

//HasRoundingModulus returns true if RoundingModulus is present, Tag 469
func (m MultilegOrderCancelReplace) HasRoundingModulus() bool {
	return m.Has(tag.RoundingModulus)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m MultilegOrderCancelReplace) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m MultilegOrderCancelReplace) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m MultilegOrderCancelReplace) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasCommCurrency returns true if CommCurrency is present, Tag 479
func (m MultilegOrderCancelReplace) HasCommCurrency() bool {
	return m.Has(tag.CommCurrency)
}

//HasCancellationRights returns true if CancellationRights is present, Tag 480
func (m MultilegOrderCancelReplace) HasCancellationRights() bool {
	return m.Has(tag.CancellationRights)
}

//HasMoneyLaunderingStatus returns true if MoneyLaunderingStatus is present, Tag 481
func (m MultilegOrderCancelReplace) HasMoneyLaunderingStatus() bool {
	return m.Has(tag.MoneyLaunderingStatus)
}

//HasDesignation returns true if Designation is present, Tag 494
func (m MultilegOrderCancelReplace) HasDesignation() bool {
	return m.Has(tag.Designation)
}

//HasFundRenewWaiv returns true if FundRenewWaiv is present, Tag 497
func (m MultilegOrderCancelReplace) HasFundRenewWaiv() bool {
	return m.Has(tag.FundRenewWaiv)
}

//HasRegistID returns true if RegistID is present, Tag 513
func (m MultilegOrderCancelReplace) HasRegistID() bool {
	return m.Has(tag.RegistID)
}

//HasOrderPercent returns true if OrderPercent is present, Tag 516
func (m MultilegOrderCancelReplace) HasOrderPercent() bool {
	return m.Has(tag.OrderPercent)
}

//HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526
func (m MultilegOrderCancelReplace) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

//HasOrderCapacity returns true if OrderCapacity is present, Tag 528
func (m MultilegOrderCancelReplace) HasOrderCapacity() bool {
	return m.Has(tag.OrderCapacity)
}

//HasOrderRestrictions returns true if OrderRestrictions is present, Tag 529
func (m MultilegOrderCancelReplace) HasOrderRestrictions() bool {
	return m.Has(tag.OrderRestrictions)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m MultilegOrderCancelReplace) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m MultilegOrderCancelReplace) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasCashMargin returns true if CashMargin is present, Tag 544
func (m MultilegOrderCancelReplace) HasCashMargin() bool {
	return m.Has(tag.CashMargin)
}

//HasNoLegs returns true if NoLegs is present, Tag 555
func (m MultilegOrderCancelReplace) HasNoLegs() bool {
	return m.Has(tag.NoLegs)
}

//HasMultiLegRptTypeReq returns true if MultiLegRptTypeReq is present, Tag 563
func (m MultilegOrderCancelReplace) HasMultiLegRptTypeReq() bool {
	return m.Has(tag.MultiLegRptTypeReq)
}

//HasAccountType returns true if AccountType is present, Tag 581
func (m MultilegOrderCancelReplace) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

//HasCustOrderCapacity returns true if CustOrderCapacity is present, Tag 582
func (m MultilegOrderCancelReplace) HasCustOrderCapacity() bool {
	return m.Has(tag.CustOrderCapacity)
}

//HasClOrdLinkID returns true if ClOrdLinkID is present, Tag 583
func (m MultilegOrderCancelReplace) HasClOrdLinkID() bool {
	return m.Has(tag.ClOrdLinkID)
}

//HasOrigOrdModTime returns true if OrigOrdModTime is present, Tag 586
func (m MultilegOrderCancelReplace) HasOrigOrdModTime() bool {
	return m.Has(tag.OrigOrdModTime)
}

//HasDayBookingInst returns true if DayBookingInst is present, Tag 589
func (m MultilegOrderCancelReplace) HasDayBookingInst() bool {
	return m.Has(tag.DayBookingInst)
}

//HasBookingUnit returns true if BookingUnit is present, Tag 590
func (m MultilegOrderCancelReplace) HasBookingUnit() bool {
	return m.Has(tag.BookingUnit)
}

//HasPreallocMethod returns true if PreallocMethod is present, Tag 591
func (m MultilegOrderCancelReplace) HasPreallocMethod() bool {
	return m.Has(tag.PreallocMethod)
}

//HasClearingFeeIndicator returns true if ClearingFeeIndicator is present, Tag 635
func (m MultilegOrderCancelReplace) HasClearingFeeIndicator() bool {
	return m.Has(tag.ClearingFeeIndicator)
}

//HasAcctIDSource returns true if AcctIDSource is present, Tag 660
func (m MultilegOrderCancelReplace) HasAcctIDSource() bool {
	return m.Has(tag.AcctIDSource)
}

//HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m MultilegOrderCancelReplace) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

//HasPool returns true if Pool is present, Tag 691
func (m MultilegOrderCancelReplace) HasPool() bool {
	return m.Has(tag.Pool)
}

//HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711
func (m MultilegOrderCancelReplace) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m MultilegOrderCancelReplace) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

//HasBookingType returns true if BookingType is present, Tag 775
func (m MultilegOrderCancelReplace) HasBookingType() bool {
	return m.Has(tag.BookingType)
}

//HasPegMoveType returns true if PegMoveType is present, Tag 835
func (m MultilegOrderCancelReplace) HasPegMoveType() bool {
	return m.Has(tag.PegMoveType)
}

//HasPegOffsetType returns true if PegOffsetType is present, Tag 836
func (m MultilegOrderCancelReplace) HasPegOffsetType() bool {
	return m.Has(tag.PegOffsetType)
}

//HasPegLimitType returns true if PegLimitType is present, Tag 837
func (m MultilegOrderCancelReplace) HasPegLimitType() bool {
	return m.Has(tag.PegLimitType)
}

//HasPegRoundDirection returns true if PegRoundDirection is present, Tag 838
func (m MultilegOrderCancelReplace) HasPegRoundDirection() bool {
	return m.Has(tag.PegRoundDirection)
}

//HasPegScope returns true if PegScope is present, Tag 840
func (m MultilegOrderCancelReplace) HasPegScope() bool {
	return m.Has(tag.PegScope)
}

//HasDiscretionMoveType returns true if DiscretionMoveType is present, Tag 841
func (m MultilegOrderCancelReplace) HasDiscretionMoveType() bool {
	return m.Has(tag.DiscretionMoveType)
}

//HasDiscretionOffsetType returns true if DiscretionOffsetType is present, Tag 842
func (m MultilegOrderCancelReplace) HasDiscretionOffsetType() bool {
	return m.Has(tag.DiscretionOffsetType)
}

//HasDiscretionLimitType returns true if DiscretionLimitType is present, Tag 843
func (m MultilegOrderCancelReplace) HasDiscretionLimitType() bool {
	return m.Has(tag.DiscretionLimitType)
}

//HasDiscretionRoundDirection returns true if DiscretionRoundDirection is present, Tag 844
func (m MultilegOrderCancelReplace) HasDiscretionRoundDirection() bool {
	return m.Has(tag.DiscretionRoundDirection)
}

//HasDiscretionScope returns true if DiscretionScope is present, Tag 846
func (m MultilegOrderCancelReplace) HasDiscretionScope() bool {
	return m.Has(tag.DiscretionScope)
}

//HasTargetStrategy returns true if TargetStrategy is present, Tag 847
func (m MultilegOrderCancelReplace) HasTargetStrategy() bool {
	return m.Has(tag.TargetStrategy)
}

//HasTargetStrategyParameters returns true if TargetStrategyParameters is present, Tag 848
func (m MultilegOrderCancelReplace) HasTargetStrategyParameters() bool {
	return m.Has(tag.TargetStrategyParameters)
}

//HasParticipationRate returns true if ParticipationRate is present, Tag 849
func (m MultilegOrderCancelReplace) HasParticipationRate() bool {
	return m.Has(tag.ParticipationRate)
}

//HasQtyType returns true if QtyType is present, Tag 854
func (m MultilegOrderCancelReplace) HasQtyType() bool {
	return m.Has(tag.QtyType)
}

//HasNoEvents returns true if NoEvents is present, Tag 864
func (m MultilegOrderCancelReplace) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

//HasDatedDate returns true if DatedDate is present, Tag 873
func (m MultilegOrderCancelReplace) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

//HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m MultilegOrderCancelReplace) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

//HasCPProgram returns true if CPProgram is present, Tag 875
func (m MultilegOrderCancelReplace) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

//HasCPRegType returns true if CPRegType is present, Tag 876
func (m MultilegOrderCancelReplace) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

//HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m MultilegOrderCancelReplace) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

//HasNoStrategyParameters returns true if NoStrategyParameters is present, Tag 957
func (m MultilegOrderCancelReplace) HasNoStrategyParameters() bool {
	return m.Has(tag.NoStrategyParameters)
}

//HasSecurityStatus returns true if SecurityStatus is present, Tag 965
func (m MultilegOrderCancelReplace) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

//HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966
func (m MultilegOrderCancelReplace) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

//HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967
func (m MultilegOrderCancelReplace) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

//HasStrikeValue returns true if StrikeValue is present, Tag 968
func (m MultilegOrderCancelReplace) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

//HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969
func (m MultilegOrderCancelReplace) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

//HasPositionLimit returns true if PositionLimit is present, Tag 970
func (m MultilegOrderCancelReplace) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

//HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971
func (m MultilegOrderCancelReplace) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

//HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996
func (m MultilegOrderCancelReplace) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

//HasTimeUnit returns true if TimeUnit is present, Tag 997
func (m MultilegOrderCancelReplace) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

//HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018
func (m MultilegOrderCancelReplace) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

//HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049
func (m MultilegOrderCancelReplace) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

//HasSwapPoints returns true if SwapPoints is present, Tag 1069
func (m MultilegOrderCancelReplace) HasSwapPoints() bool {
	return m.Has(tag.SwapPoints)
}

//HasMaturityTime returns true if MaturityTime is present, Tag 1079
func (m MultilegOrderCancelReplace) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

//HasSecondaryDisplayQty returns true if SecondaryDisplayQty is present, Tag 1082
func (m MultilegOrderCancelReplace) HasSecondaryDisplayQty() bool {
	return m.Has(tag.SecondaryDisplayQty)
}

//HasDisplayWhen returns true if DisplayWhen is present, Tag 1083
func (m MultilegOrderCancelReplace) HasDisplayWhen() bool {
	return m.Has(tag.DisplayWhen)
}

//HasDisplayMethod returns true if DisplayMethod is present, Tag 1084
func (m MultilegOrderCancelReplace) HasDisplayMethod() bool {
	return m.Has(tag.DisplayMethod)
}

//HasDisplayLowQty returns true if DisplayLowQty is present, Tag 1085
func (m MultilegOrderCancelReplace) HasDisplayLowQty() bool {
	return m.Has(tag.DisplayLowQty)
}

//HasDisplayHighQty returns true if DisplayHighQty is present, Tag 1086
func (m MultilegOrderCancelReplace) HasDisplayHighQty() bool {
	return m.Has(tag.DisplayHighQty)
}

//HasDisplayMinIncr returns true if DisplayMinIncr is present, Tag 1087
func (m MultilegOrderCancelReplace) HasDisplayMinIncr() bool {
	return m.Has(tag.DisplayMinIncr)
}

//HasRefreshQty returns true if RefreshQty is present, Tag 1088
func (m MultilegOrderCancelReplace) HasRefreshQty() bool {
	return m.Has(tag.RefreshQty)
}

//HasMatchIncrement returns true if MatchIncrement is present, Tag 1089
func (m MultilegOrderCancelReplace) HasMatchIncrement() bool {
	return m.Has(tag.MatchIncrement)
}

//HasMaxPriceLevels returns true if MaxPriceLevels is present, Tag 1090
func (m MultilegOrderCancelReplace) HasMaxPriceLevels() bool {
	return m.Has(tag.MaxPriceLevels)
}

//HasPreTradeAnonymity returns true if PreTradeAnonymity is present, Tag 1091
func (m MultilegOrderCancelReplace) HasPreTradeAnonymity() bool {
	return m.Has(tag.PreTradeAnonymity)
}

//HasPriceProtectionScope returns true if PriceProtectionScope is present, Tag 1092
func (m MultilegOrderCancelReplace) HasPriceProtectionScope() bool {
	return m.Has(tag.PriceProtectionScope)
}

//HasPegPriceType returns true if PegPriceType is present, Tag 1094
func (m MultilegOrderCancelReplace) HasPegPriceType() bool {
	return m.Has(tag.PegPriceType)
}

//HasPegSecurityIDSource returns true if PegSecurityIDSource is present, Tag 1096
func (m MultilegOrderCancelReplace) HasPegSecurityIDSource() bool {
	return m.Has(tag.PegSecurityIDSource)
}

//HasPegSecurityID returns true if PegSecurityID is present, Tag 1097
func (m MultilegOrderCancelReplace) HasPegSecurityID() bool {
	return m.Has(tag.PegSecurityID)
}

//HasPegSymbol returns true if PegSymbol is present, Tag 1098
func (m MultilegOrderCancelReplace) HasPegSymbol() bool {
	return m.Has(tag.PegSymbol)
}

//HasPegSecurityDesc returns true if PegSecurityDesc is present, Tag 1099
func (m MultilegOrderCancelReplace) HasPegSecurityDesc() bool {
	return m.Has(tag.PegSecurityDesc)
}

//HasTriggerType returns true if TriggerType is present, Tag 1100
func (m MultilegOrderCancelReplace) HasTriggerType() bool {
	return m.Has(tag.TriggerType)
}

//HasTriggerAction returns true if TriggerAction is present, Tag 1101
func (m MultilegOrderCancelReplace) HasTriggerAction() bool {
	return m.Has(tag.TriggerAction)
}

//HasTriggerPrice returns true if TriggerPrice is present, Tag 1102
func (m MultilegOrderCancelReplace) HasTriggerPrice() bool {
	return m.Has(tag.TriggerPrice)
}

//HasTriggerSymbol returns true if TriggerSymbol is present, Tag 1103
func (m MultilegOrderCancelReplace) HasTriggerSymbol() bool {
	return m.Has(tag.TriggerSymbol)
}

//HasTriggerSecurityID returns true if TriggerSecurityID is present, Tag 1104
func (m MultilegOrderCancelReplace) HasTriggerSecurityID() bool {
	return m.Has(tag.TriggerSecurityID)
}

//HasTriggerSecurityIDSource returns true if TriggerSecurityIDSource is present, Tag 1105
func (m MultilegOrderCancelReplace) HasTriggerSecurityIDSource() bool {
	return m.Has(tag.TriggerSecurityIDSource)
}

//HasTriggerSecurityDesc returns true if TriggerSecurityDesc is present, Tag 1106
func (m MultilegOrderCancelReplace) HasTriggerSecurityDesc() bool {
	return m.Has(tag.TriggerSecurityDesc)
}

//HasTriggerPriceType returns true if TriggerPriceType is present, Tag 1107
func (m MultilegOrderCancelReplace) HasTriggerPriceType() bool {
	return m.Has(tag.TriggerPriceType)
}

//HasTriggerPriceTypeScope returns true if TriggerPriceTypeScope is present, Tag 1108
func (m MultilegOrderCancelReplace) HasTriggerPriceTypeScope() bool {
	return m.Has(tag.TriggerPriceTypeScope)
}

//HasTriggerPriceDirection returns true if TriggerPriceDirection is present, Tag 1109
func (m MultilegOrderCancelReplace) HasTriggerPriceDirection() bool {
	return m.Has(tag.TriggerPriceDirection)
}

//HasTriggerNewPrice returns true if TriggerNewPrice is present, Tag 1110
func (m MultilegOrderCancelReplace) HasTriggerNewPrice() bool {
	return m.Has(tag.TriggerNewPrice)
}

//HasTriggerOrderType returns true if TriggerOrderType is present, Tag 1111
func (m MultilegOrderCancelReplace) HasTriggerOrderType() bool {
	return m.Has(tag.TriggerOrderType)
}

//HasTriggerNewQty returns true if TriggerNewQty is present, Tag 1112
func (m MultilegOrderCancelReplace) HasTriggerNewQty() bool {
	return m.Has(tag.TriggerNewQty)
}

//HasTriggerTradingSessionID returns true if TriggerTradingSessionID is present, Tag 1113
func (m MultilegOrderCancelReplace) HasTriggerTradingSessionID() bool {
	return m.Has(tag.TriggerTradingSessionID)
}

//HasTriggerTradingSessionSubID returns true if TriggerTradingSessionSubID is present, Tag 1114
func (m MultilegOrderCancelReplace) HasTriggerTradingSessionSubID() bool {
	return m.Has(tag.TriggerTradingSessionSubID)
}

//HasExDestinationIDSource returns true if ExDestinationIDSource is present, Tag 1133
func (m MultilegOrderCancelReplace) HasExDestinationIDSource() bool {
	return m.Has(tag.ExDestinationIDSource)
}

//HasDisplayQty returns true if DisplayQty is present, Tag 1138
func (m MultilegOrderCancelReplace) HasDisplayQty() bool {
	return m.Has(tag.DisplayQty)
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

//SetNoNested3PartyIDs sets NoNested3PartyIDs, Tag 948
func (m NoAllocs) SetNoNested3PartyIDs(f NoNested3PartyIDsRepeatingGroup) {
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

//GetNoNested3PartyIDs gets NoNested3PartyIDs, Tag 948
func (m NoAllocs) GetNoNested3PartyIDs() (NoNested3PartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNested3PartyIDsRepeatingGroup()
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

//HasNoNested3PartyIDs returns true if NoNested3PartyIDs is present, Tag 948
func (m NoAllocs) HasNoNested3PartyIDs() bool {
	return m.Has(tag.NoNested3PartyIDs)
}

//HasAllocQty returns true if AllocQty is present, Tag 80
func (m NoAllocs) HasAllocQty() bool {
	return m.Has(tag.AllocQty)
}

//NoNested3PartyIDs is a repeating group element, Tag 948
type NoNested3PartyIDs struct {
	quickfix.Group
}

//SetNested3PartyID sets Nested3PartyID, Tag 949
func (m NoNested3PartyIDs) SetNested3PartyID(v string) {
	m.Set(field.NewNested3PartyID(v))
}

//SetNested3PartyIDSource sets Nested3PartyIDSource, Tag 950
func (m NoNested3PartyIDs) SetNested3PartyIDSource(v string) {
	m.Set(field.NewNested3PartyIDSource(v))
}

//SetNested3PartyRole sets Nested3PartyRole, Tag 951
func (m NoNested3PartyIDs) SetNested3PartyRole(v int) {
	m.Set(field.NewNested3PartyRole(v))
}

//SetNoNested3PartySubIDs sets NoNested3PartySubIDs, Tag 952
func (m NoNested3PartyIDs) SetNoNested3PartySubIDs(f NoNested3PartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetNested3PartyID gets Nested3PartyID, Tag 949
func (m NoNested3PartyIDs) GetNested3PartyID() (f field.Nested3PartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNested3PartyIDSource gets Nested3PartyIDSource, Tag 950
func (m NoNested3PartyIDs) GetNested3PartyIDSource() (f field.Nested3PartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNested3PartyRole gets Nested3PartyRole, Tag 951
func (m NoNested3PartyIDs) GetNested3PartyRole() (f field.Nested3PartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoNested3PartySubIDs gets NoNested3PartySubIDs, Tag 952
func (m NoNested3PartyIDs) GetNoNested3PartySubIDs() (NoNested3PartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNested3PartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasNested3PartyID returns true if Nested3PartyID is present, Tag 949
func (m NoNested3PartyIDs) HasNested3PartyID() bool {
	return m.Has(tag.Nested3PartyID)
}

//HasNested3PartyIDSource returns true if Nested3PartyIDSource is present, Tag 950
func (m NoNested3PartyIDs) HasNested3PartyIDSource() bool {
	return m.Has(tag.Nested3PartyIDSource)
}

//HasNested3PartyRole returns true if Nested3PartyRole is present, Tag 951
func (m NoNested3PartyIDs) HasNested3PartyRole() bool {
	return m.Has(tag.Nested3PartyRole)
}

//HasNoNested3PartySubIDs returns true if NoNested3PartySubIDs is present, Tag 952
func (m NoNested3PartyIDs) HasNoNested3PartySubIDs() bool {
	return m.Has(tag.NoNested3PartySubIDs)
}

//NoNested3PartySubIDs is a repeating group element, Tag 952
type NoNested3PartySubIDs struct {
	quickfix.Group
}

//SetNested3PartySubID sets Nested3PartySubID, Tag 953
func (m NoNested3PartySubIDs) SetNested3PartySubID(v string) {
	m.Set(field.NewNested3PartySubID(v))
}

//SetNested3PartySubIDType sets Nested3PartySubIDType, Tag 954
func (m NoNested3PartySubIDs) SetNested3PartySubIDType(v int) {
	m.Set(field.NewNested3PartySubIDType(v))
}

//GetNested3PartySubID gets Nested3PartySubID, Tag 953
func (m NoNested3PartySubIDs) GetNested3PartySubID() (f field.Nested3PartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNested3PartySubIDType gets Nested3PartySubIDType, Tag 954
func (m NoNested3PartySubIDs) GetNested3PartySubIDType() (f field.Nested3PartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasNested3PartySubID returns true if Nested3PartySubID is present, Tag 953
func (m NoNested3PartySubIDs) HasNested3PartySubID() bool {
	return m.Has(tag.Nested3PartySubID)
}

//HasNested3PartySubIDType returns true if Nested3PartySubIDType is present, Tag 954
func (m NoNested3PartySubIDs) HasNested3PartySubIDType() bool {
	return m.Has(tag.Nested3PartySubIDType)
}

//NoNested3PartySubIDsRepeatingGroup is a repeating group, Tag 952
type NoNested3PartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNested3PartySubIDsRepeatingGroup returns an initialized, NoNested3PartySubIDsRepeatingGroup
func NewNoNested3PartySubIDsRepeatingGroup() NoNested3PartySubIDsRepeatingGroup {
	return NoNested3PartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNested3PartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Nested3PartySubID), quickfix.GroupElement(tag.Nested3PartySubIDType)})}
}

//Add create and append a new NoNested3PartySubIDs to this group
func (m NoNested3PartySubIDsRepeatingGroup) Add() NoNested3PartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoNested3PartySubIDs{g}
}

//Get returns the ith NoNested3PartySubIDs in the NoNested3PartySubIDsRepeatinGroup
func (m NoNested3PartySubIDsRepeatingGroup) Get(i int) NoNested3PartySubIDs {
	return NoNested3PartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoNested3PartyIDsRepeatingGroup is a repeating group, Tag 948
type NoNested3PartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNested3PartyIDsRepeatingGroup returns an initialized, NoNested3PartyIDsRepeatingGroup
func NewNoNested3PartyIDsRepeatingGroup() NoNested3PartyIDsRepeatingGroup {
	return NoNested3PartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNested3PartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Nested3PartyID), quickfix.GroupElement(tag.Nested3PartyIDSource), quickfix.GroupElement(tag.Nested3PartyRole), NewNoNested3PartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoNested3PartyIDs to this group
func (m NoNested3PartyIDsRepeatingGroup) Add() NoNested3PartyIDs {
	g := m.RepeatingGroup.Add()
	return NoNested3PartyIDs{g}
}

//Get returns the ith NoNested3PartyIDs in the NoNested3PartyIDsRepeatinGroup
func (m NoNested3PartyIDsRepeatingGroup) Get(i int) NoNested3PartyIDs {
	return NoNested3PartyIDs{m.RepeatingGroup.Get(i)}
}

//NoAllocsRepeatingGroup is a repeating group, Tag 78
type NoAllocsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoAllocsRepeatingGroup returns an initialized, NoAllocsRepeatingGroup
func NewNoAllocsRepeatingGroup() NoAllocsRepeatingGroup {
	return NoAllocsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAllocs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AllocAccount), quickfix.GroupElement(tag.AllocAcctIDSource), quickfix.GroupElement(tag.AllocSettlCurrency), quickfix.GroupElement(tag.IndividualAllocID), NewNoNested3PartyIDsRepeatingGroup(), quickfix.GroupElement(tag.AllocQty)})}
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

//NoLegs is a repeating group element, Tag 555
type NoLegs struct {
	quickfix.Group
}

//SetLegSymbol sets LegSymbol, Tag 600
func (m NoLegs) SetLegSymbol(v string) {
	m.Set(field.NewLegSymbol(v))
}

//SetLegSymbolSfx sets LegSymbolSfx, Tag 601
func (m NoLegs) SetLegSymbolSfx(v string) {
	m.Set(field.NewLegSymbolSfx(v))
}

//SetLegSecurityID sets LegSecurityID, Tag 602
func (m NoLegs) SetLegSecurityID(v string) {
	m.Set(field.NewLegSecurityID(v))
}

//SetLegSecurityIDSource sets LegSecurityIDSource, Tag 603
func (m NoLegs) SetLegSecurityIDSource(v string) {
	m.Set(field.NewLegSecurityIDSource(v))
}

//SetNoLegSecurityAltID sets NoLegSecurityAltID, Tag 604
func (m NoLegs) SetNoLegSecurityAltID(f NoLegSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetLegProduct sets LegProduct, Tag 607
func (m NoLegs) SetLegProduct(v int) {
	m.Set(field.NewLegProduct(v))
}

//SetLegCFICode sets LegCFICode, Tag 608
func (m NoLegs) SetLegCFICode(v string) {
	m.Set(field.NewLegCFICode(v))
}

//SetLegSecurityType sets LegSecurityType, Tag 609
func (m NoLegs) SetLegSecurityType(v string) {
	m.Set(field.NewLegSecurityType(v))
}

//SetLegSecuritySubType sets LegSecuritySubType, Tag 764
func (m NoLegs) SetLegSecuritySubType(v string) {
	m.Set(field.NewLegSecuritySubType(v))
}

//SetLegMaturityMonthYear sets LegMaturityMonthYear, Tag 610
func (m NoLegs) SetLegMaturityMonthYear(v string) {
	m.Set(field.NewLegMaturityMonthYear(v))
}

//SetLegMaturityDate sets LegMaturityDate, Tag 611
func (m NoLegs) SetLegMaturityDate(v string) {
	m.Set(field.NewLegMaturityDate(v))
}

//SetLegCouponPaymentDate sets LegCouponPaymentDate, Tag 248
func (m NoLegs) SetLegCouponPaymentDate(v string) {
	m.Set(field.NewLegCouponPaymentDate(v))
}

//SetLegIssueDate sets LegIssueDate, Tag 249
func (m NoLegs) SetLegIssueDate(v string) {
	m.Set(field.NewLegIssueDate(v))
}

//SetLegRepoCollateralSecurityType sets LegRepoCollateralSecurityType, Tag 250
func (m NoLegs) SetLegRepoCollateralSecurityType(v int) {
	m.Set(field.NewLegRepoCollateralSecurityType(v))
}

//SetLegRepurchaseTerm sets LegRepurchaseTerm, Tag 251
func (m NoLegs) SetLegRepurchaseTerm(v int) {
	m.Set(field.NewLegRepurchaseTerm(v))
}

//SetLegRepurchaseRate sets LegRepurchaseRate, Tag 252
func (m NoLegs) SetLegRepurchaseRate(v float64) {
	m.Set(field.NewLegRepurchaseRate(v))
}

//SetLegFactor sets LegFactor, Tag 253
func (m NoLegs) SetLegFactor(v float64) {
	m.Set(field.NewLegFactor(v))
}

//SetLegCreditRating sets LegCreditRating, Tag 257
func (m NoLegs) SetLegCreditRating(v string) {
	m.Set(field.NewLegCreditRating(v))
}

//SetLegInstrRegistry sets LegInstrRegistry, Tag 599
func (m NoLegs) SetLegInstrRegistry(v string) {
	m.Set(field.NewLegInstrRegistry(v))
}

//SetLegCountryOfIssue sets LegCountryOfIssue, Tag 596
func (m NoLegs) SetLegCountryOfIssue(v string) {
	m.Set(field.NewLegCountryOfIssue(v))
}

//SetLegStateOrProvinceOfIssue sets LegStateOrProvinceOfIssue, Tag 597
func (m NoLegs) SetLegStateOrProvinceOfIssue(v string) {
	m.Set(field.NewLegStateOrProvinceOfIssue(v))
}

//SetLegLocaleOfIssue sets LegLocaleOfIssue, Tag 598
func (m NoLegs) SetLegLocaleOfIssue(v string) {
	m.Set(field.NewLegLocaleOfIssue(v))
}

//SetLegRedemptionDate sets LegRedemptionDate, Tag 254
func (m NoLegs) SetLegRedemptionDate(v string) {
	m.Set(field.NewLegRedemptionDate(v))
}

//SetLegStrikePrice sets LegStrikePrice, Tag 612
func (m NoLegs) SetLegStrikePrice(v float64) {
	m.Set(field.NewLegStrikePrice(v))
}

//SetLegStrikeCurrency sets LegStrikeCurrency, Tag 942
func (m NoLegs) SetLegStrikeCurrency(v string) {
	m.Set(field.NewLegStrikeCurrency(v))
}

//SetLegOptAttribute sets LegOptAttribute, Tag 613
func (m NoLegs) SetLegOptAttribute(v string) {
	m.Set(field.NewLegOptAttribute(v))
}

//SetLegContractMultiplier sets LegContractMultiplier, Tag 614
func (m NoLegs) SetLegContractMultiplier(v float64) {
	m.Set(field.NewLegContractMultiplier(v))
}

//SetLegCouponRate sets LegCouponRate, Tag 615
func (m NoLegs) SetLegCouponRate(v float64) {
	m.Set(field.NewLegCouponRate(v))
}

//SetLegSecurityExchange sets LegSecurityExchange, Tag 616
func (m NoLegs) SetLegSecurityExchange(v string) {
	m.Set(field.NewLegSecurityExchange(v))
}

//SetLegIssuer sets LegIssuer, Tag 617
func (m NoLegs) SetLegIssuer(v string) {
	m.Set(field.NewLegIssuer(v))
}

//SetEncodedLegIssuerLen sets EncodedLegIssuerLen, Tag 618
func (m NoLegs) SetEncodedLegIssuerLen(v int) {
	m.Set(field.NewEncodedLegIssuerLen(v))
}

//SetEncodedLegIssuer sets EncodedLegIssuer, Tag 619
func (m NoLegs) SetEncodedLegIssuer(v string) {
	m.Set(field.NewEncodedLegIssuer(v))
}

//SetLegSecurityDesc sets LegSecurityDesc, Tag 620
func (m NoLegs) SetLegSecurityDesc(v string) {
	m.Set(field.NewLegSecurityDesc(v))
}

//SetEncodedLegSecurityDescLen sets EncodedLegSecurityDescLen, Tag 621
func (m NoLegs) SetEncodedLegSecurityDescLen(v int) {
	m.Set(field.NewEncodedLegSecurityDescLen(v))
}

//SetEncodedLegSecurityDesc sets EncodedLegSecurityDesc, Tag 622
func (m NoLegs) SetEncodedLegSecurityDesc(v string) {
	m.Set(field.NewEncodedLegSecurityDesc(v))
}

//SetLegRatioQty sets LegRatioQty, Tag 623
func (m NoLegs) SetLegRatioQty(v float64) {
	m.Set(field.NewLegRatioQty(v))
}

//SetLegSide sets LegSide, Tag 624
func (m NoLegs) SetLegSide(v string) {
	m.Set(field.NewLegSide(v))
}

//SetLegCurrency sets LegCurrency, Tag 556
func (m NoLegs) SetLegCurrency(v string) {
	m.Set(field.NewLegCurrency(v))
}

//SetLegPool sets LegPool, Tag 740
func (m NoLegs) SetLegPool(v string) {
	m.Set(field.NewLegPool(v))
}

//SetLegDatedDate sets LegDatedDate, Tag 739
func (m NoLegs) SetLegDatedDate(v string) {
	m.Set(field.NewLegDatedDate(v))
}

//SetLegContractSettlMonth sets LegContractSettlMonth, Tag 955
func (m NoLegs) SetLegContractSettlMonth(v string) {
	m.Set(field.NewLegContractSettlMonth(v))
}

//SetLegInterestAccrualDate sets LegInterestAccrualDate, Tag 956
func (m NoLegs) SetLegInterestAccrualDate(v string) {
	m.Set(field.NewLegInterestAccrualDate(v))
}

//SetLegUnitOfMeasure sets LegUnitOfMeasure, Tag 999
func (m NoLegs) SetLegUnitOfMeasure(v string) {
	m.Set(field.NewLegUnitOfMeasure(v))
}

//SetLegTimeUnit sets LegTimeUnit, Tag 1001
func (m NoLegs) SetLegTimeUnit(v string) {
	m.Set(field.NewLegTimeUnit(v))
}

//SetLegQty sets LegQty, Tag 687
func (m NoLegs) SetLegQty(v float64) {
	m.Set(field.NewLegQty(v))
}

//SetLegSwapType sets LegSwapType, Tag 690
func (m NoLegs) SetLegSwapType(v int) {
	m.Set(field.NewLegSwapType(v))
}

//SetNoLegStipulations sets NoLegStipulations, Tag 683
func (m NoLegs) SetNoLegStipulations(f NoLegStipulationsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoLegAllocs sets NoLegAllocs, Tag 670
func (m NoLegs) SetNoLegAllocs(f NoLegAllocsRepeatingGroup) {
	m.SetGroup(f)
}

//SetLegPositionEffect sets LegPositionEffect, Tag 564
func (m NoLegs) SetLegPositionEffect(v string) {
	m.Set(field.NewLegPositionEffect(v))
}

//SetLegCoveredOrUncovered sets LegCoveredOrUncovered, Tag 565
func (m NoLegs) SetLegCoveredOrUncovered(v int) {
	m.Set(field.NewLegCoveredOrUncovered(v))
}

//SetNoNestedPartyIDs sets NoNestedPartyIDs, Tag 539
func (m NoLegs) SetNoNestedPartyIDs(f NoNestedPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetLegRefID sets LegRefID, Tag 654
func (m NoLegs) SetLegRefID(v string) {
	m.Set(field.NewLegRefID(v))
}

//SetLegPrice sets LegPrice, Tag 566
func (m NoLegs) SetLegPrice(v float64) {
	m.Set(field.NewLegPrice(v))
}

//SetLegSettlType sets LegSettlType, Tag 587
func (m NoLegs) SetLegSettlType(v string) {
	m.Set(field.NewLegSettlType(v))
}

//SetLegSettlDate sets LegSettlDate, Tag 588
func (m NoLegs) SetLegSettlDate(v string) {
	m.Set(field.NewLegSettlDate(v))
}

//SetLegOptionRatio sets LegOptionRatio, Tag 1017
func (m NoLegs) SetLegOptionRatio(v float64) {
	m.Set(field.NewLegOptionRatio(v))
}

//SetLegOrderQty sets LegOrderQty, Tag 685
func (m NoLegs) SetLegOrderQty(v float64) {
	m.Set(field.NewLegOrderQty(v))
}

//GetLegSymbol gets LegSymbol, Tag 600
func (m NoLegs) GetLegSymbol() (f field.LegSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSymbolSfx gets LegSymbolSfx, Tag 601
func (m NoLegs) GetLegSymbolSfx() (f field.LegSymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityID gets LegSecurityID, Tag 602
func (m NoLegs) GetLegSecurityID() (f field.LegSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityIDSource gets LegSecurityIDSource, Tag 603
func (m NoLegs) GetLegSecurityIDSource() (f field.LegSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoLegSecurityAltID gets NoLegSecurityAltID, Tag 604
func (m NoLegs) GetNoLegSecurityAltID() (NoLegSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLegProduct gets LegProduct, Tag 607
func (m NoLegs) GetLegProduct() (f field.LegProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCFICode gets LegCFICode, Tag 608
func (m NoLegs) GetLegCFICode() (f field.LegCFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityType gets LegSecurityType, Tag 609
func (m NoLegs) GetLegSecurityType() (f field.LegSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecuritySubType gets LegSecuritySubType, Tag 764
func (m NoLegs) GetLegSecuritySubType() (f field.LegSecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegMaturityMonthYear gets LegMaturityMonthYear, Tag 610
func (m NoLegs) GetLegMaturityMonthYear() (f field.LegMaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegMaturityDate gets LegMaturityDate, Tag 611
func (m NoLegs) GetLegMaturityDate() (f field.LegMaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCouponPaymentDate gets LegCouponPaymentDate, Tag 248
func (m NoLegs) GetLegCouponPaymentDate() (f field.LegCouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegIssueDate gets LegIssueDate, Tag 249
func (m NoLegs) GetLegIssueDate() (f field.LegIssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRepoCollateralSecurityType gets LegRepoCollateralSecurityType, Tag 250
func (m NoLegs) GetLegRepoCollateralSecurityType() (f field.LegRepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRepurchaseTerm gets LegRepurchaseTerm, Tag 251
func (m NoLegs) GetLegRepurchaseTerm() (f field.LegRepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRepurchaseRate gets LegRepurchaseRate, Tag 252
func (m NoLegs) GetLegRepurchaseRate() (f field.LegRepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegFactor gets LegFactor, Tag 253
func (m NoLegs) GetLegFactor() (f field.LegFactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCreditRating gets LegCreditRating, Tag 257
func (m NoLegs) GetLegCreditRating() (f field.LegCreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegInstrRegistry gets LegInstrRegistry, Tag 599
func (m NoLegs) GetLegInstrRegistry() (f field.LegInstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCountryOfIssue gets LegCountryOfIssue, Tag 596
func (m NoLegs) GetLegCountryOfIssue() (f field.LegCountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegStateOrProvinceOfIssue gets LegStateOrProvinceOfIssue, Tag 597
func (m NoLegs) GetLegStateOrProvinceOfIssue() (f field.LegStateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegLocaleOfIssue gets LegLocaleOfIssue, Tag 598
func (m NoLegs) GetLegLocaleOfIssue() (f field.LegLocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRedemptionDate gets LegRedemptionDate, Tag 254
func (m NoLegs) GetLegRedemptionDate() (f field.LegRedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegStrikePrice gets LegStrikePrice, Tag 612
func (m NoLegs) GetLegStrikePrice() (f field.LegStrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegStrikeCurrency gets LegStrikeCurrency, Tag 942
func (m NoLegs) GetLegStrikeCurrency() (f field.LegStrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegOptAttribute gets LegOptAttribute, Tag 613
func (m NoLegs) GetLegOptAttribute() (f field.LegOptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegContractMultiplier gets LegContractMultiplier, Tag 614
func (m NoLegs) GetLegContractMultiplier() (f field.LegContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCouponRate gets LegCouponRate, Tag 615
func (m NoLegs) GetLegCouponRate() (f field.LegCouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityExchange gets LegSecurityExchange, Tag 616
func (m NoLegs) GetLegSecurityExchange() (f field.LegSecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegIssuer gets LegIssuer, Tag 617
func (m NoLegs) GetLegIssuer() (f field.LegIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegIssuerLen gets EncodedLegIssuerLen, Tag 618
func (m NoLegs) GetEncodedLegIssuerLen() (f field.EncodedLegIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegIssuer gets EncodedLegIssuer, Tag 619
func (m NoLegs) GetEncodedLegIssuer() (f field.EncodedLegIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityDesc gets LegSecurityDesc, Tag 620
func (m NoLegs) GetLegSecurityDesc() (f field.LegSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegSecurityDescLen gets EncodedLegSecurityDescLen, Tag 621
func (m NoLegs) GetEncodedLegSecurityDescLen() (f field.EncodedLegSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegSecurityDesc gets EncodedLegSecurityDesc, Tag 622
func (m NoLegs) GetEncodedLegSecurityDesc() (f field.EncodedLegSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRatioQty gets LegRatioQty, Tag 623
func (m NoLegs) GetLegRatioQty() (f field.LegRatioQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSide gets LegSide, Tag 624
func (m NoLegs) GetLegSide() (f field.LegSideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCurrency gets LegCurrency, Tag 556
func (m NoLegs) GetLegCurrency() (f field.LegCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPool gets LegPool, Tag 740
func (m NoLegs) GetLegPool() (f field.LegPoolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegDatedDate gets LegDatedDate, Tag 739
func (m NoLegs) GetLegDatedDate() (f field.LegDatedDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegContractSettlMonth gets LegContractSettlMonth, Tag 955
func (m NoLegs) GetLegContractSettlMonth() (f field.LegContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegInterestAccrualDate gets LegInterestAccrualDate, Tag 956
func (m NoLegs) GetLegInterestAccrualDate() (f field.LegInterestAccrualDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegUnitOfMeasure gets LegUnitOfMeasure, Tag 999
func (m NoLegs) GetLegUnitOfMeasure() (f field.LegUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegTimeUnit gets LegTimeUnit, Tag 1001
func (m NoLegs) GetLegTimeUnit() (f field.LegTimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegQty gets LegQty, Tag 687
func (m NoLegs) GetLegQty() (f field.LegQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSwapType gets LegSwapType, Tag 690
func (m NoLegs) GetLegSwapType() (f field.LegSwapTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoLegStipulations gets NoLegStipulations, Tag 683
func (m NoLegs) GetNoLegStipulations() (NoLegStipulationsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegStipulationsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoLegAllocs gets NoLegAllocs, Tag 670
func (m NoLegs) GetNoLegAllocs() (NoLegAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLegPositionEffect gets LegPositionEffect, Tag 564
func (m NoLegs) GetLegPositionEffect() (f field.LegPositionEffectField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCoveredOrUncovered gets LegCoveredOrUncovered, Tag 565
func (m NoLegs) GetLegCoveredOrUncovered() (f field.LegCoveredOrUncoveredField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoNestedPartyIDs gets NoNestedPartyIDs, Tag 539
func (m NoLegs) GetNoNestedPartyIDs() (NoNestedPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLegRefID gets LegRefID, Tag 654
func (m NoLegs) GetLegRefID() (f field.LegRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPrice gets LegPrice, Tag 566
func (m NoLegs) GetLegPrice() (f field.LegPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSettlType gets LegSettlType, Tag 587
func (m NoLegs) GetLegSettlType() (f field.LegSettlTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSettlDate gets LegSettlDate, Tag 588
func (m NoLegs) GetLegSettlDate() (f field.LegSettlDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegOptionRatio gets LegOptionRatio, Tag 1017
func (m NoLegs) GetLegOptionRatio() (f field.LegOptionRatioField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegOrderQty gets LegOrderQty, Tag 685
func (m NoLegs) GetLegOrderQty() (f field.LegOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasLegSymbol returns true if LegSymbol is present, Tag 600
func (m NoLegs) HasLegSymbol() bool {
	return m.Has(tag.LegSymbol)
}

//HasLegSymbolSfx returns true if LegSymbolSfx is present, Tag 601
func (m NoLegs) HasLegSymbolSfx() bool {
	return m.Has(tag.LegSymbolSfx)
}

//HasLegSecurityID returns true if LegSecurityID is present, Tag 602
func (m NoLegs) HasLegSecurityID() bool {
	return m.Has(tag.LegSecurityID)
}

//HasLegSecurityIDSource returns true if LegSecurityIDSource is present, Tag 603
func (m NoLegs) HasLegSecurityIDSource() bool {
	return m.Has(tag.LegSecurityIDSource)
}

//HasNoLegSecurityAltID returns true if NoLegSecurityAltID is present, Tag 604
func (m NoLegs) HasNoLegSecurityAltID() bool {
	return m.Has(tag.NoLegSecurityAltID)
}

//HasLegProduct returns true if LegProduct is present, Tag 607
func (m NoLegs) HasLegProduct() bool {
	return m.Has(tag.LegProduct)
}

//HasLegCFICode returns true if LegCFICode is present, Tag 608
func (m NoLegs) HasLegCFICode() bool {
	return m.Has(tag.LegCFICode)
}

//HasLegSecurityType returns true if LegSecurityType is present, Tag 609
func (m NoLegs) HasLegSecurityType() bool {
	return m.Has(tag.LegSecurityType)
}

//HasLegSecuritySubType returns true if LegSecuritySubType is present, Tag 764
func (m NoLegs) HasLegSecuritySubType() bool {
	return m.Has(tag.LegSecuritySubType)
}

//HasLegMaturityMonthYear returns true if LegMaturityMonthYear is present, Tag 610
func (m NoLegs) HasLegMaturityMonthYear() bool {
	return m.Has(tag.LegMaturityMonthYear)
}

//HasLegMaturityDate returns true if LegMaturityDate is present, Tag 611
func (m NoLegs) HasLegMaturityDate() bool {
	return m.Has(tag.LegMaturityDate)
}

//HasLegCouponPaymentDate returns true if LegCouponPaymentDate is present, Tag 248
func (m NoLegs) HasLegCouponPaymentDate() bool {
	return m.Has(tag.LegCouponPaymentDate)
}

//HasLegIssueDate returns true if LegIssueDate is present, Tag 249
func (m NoLegs) HasLegIssueDate() bool {
	return m.Has(tag.LegIssueDate)
}

//HasLegRepoCollateralSecurityType returns true if LegRepoCollateralSecurityType is present, Tag 250
func (m NoLegs) HasLegRepoCollateralSecurityType() bool {
	return m.Has(tag.LegRepoCollateralSecurityType)
}

//HasLegRepurchaseTerm returns true if LegRepurchaseTerm is present, Tag 251
func (m NoLegs) HasLegRepurchaseTerm() bool {
	return m.Has(tag.LegRepurchaseTerm)
}

//HasLegRepurchaseRate returns true if LegRepurchaseRate is present, Tag 252
func (m NoLegs) HasLegRepurchaseRate() bool {
	return m.Has(tag.LegRepurchaseRate)
}

//HasLegFactor returns true if LegFactor is present, Tag 253
func (m NoLegs) HasLegFactor() bool {
	return m.Has(tag.LegFactor)
}

//HasLegCreditRating returns true if LegCreditRating is present, Tag 257
func (m NoLegs) HasLegCreditRating() bool {
	return m.Has(tag.LegCreditRating)
}

//HasLegInstrRegistry returns true if LegInstrRegistry is present, Tag 599
func (m NoLegs) HasLegInstrRegistry() bool {
	return m.Has(tag.LegInstrRegistry)
}

//HasLegCountryOfIssue returns true if LegCountryOfIssue is present, Tag 596
func (m NoLegs) HasLegCountryOfIssue() bool {
	return m.Has(tag.LegCountryOfIssue)
}

//HasLegStateOrProvinceOfIssue returns true if LegStateOrProvinceOfIssue is present, Tag 597
func (m NoLegs) HasLegStateOrProvinceOfIssue() bool {
	return m.Has(tag.LegStateOrProvinceOfIssue)
}

//HasLegLocaleOfIssue returns true if LegLocaleOfIssue is present, Tag 598
func (m NoLegs) HasLegLocaleOfIssue() bool {
	return m.Has(tag.LegLocaleOfIssue)
}

//HasLegRedemptionDate returns true if LegRedemptionDate is present, Tag 254
func (m NoLegs) HasLegRedemptionDate() bool {
	return m.Has(tag.LegRedemptionDate)
}

//HasLegStrikePrice returns true if LegStrikePrice is present, Tag 612
func (m NoLegs) HasLegStrikePrice() bool {
	return m.Has(tag.LegStrikePrice)
}

//HasLegStrikeCurrency returns true if LegStrikeCurrency is present, Tag 942
func (m NoLegs) HasLegStrikeCurrency() bool {
	return m.Has(tag.LegStrikeCurrency)
}

//HasLegOptAttribute returns true if LegOptAttribute is present, Tag 613
func (m NoLegs) HasLegOptAttribute() bool {
	return m.Has(tag.LegOptAttribute)
}

//HasLegContractMultiplier returns true if LegContractMultiplier is present, Tag 614
func (m NoLegs) HasLegContractMultiplier() bool {
	return m.Has(tag.LegContractMultiplier)
}

//HasLegCouponRate returns true if LegCouponRate is present, Tag 615
func (m NoLegs) HasLegCouponRate() bool {
	return m.Has(tag.LegCouponRate)
}

//HasLegSecurityExchange returns true if LegSecurityExchange is present, Tag 616
func (m NoLegs) HasLegSecurityExchange() bool {
	return m.Has(tag.LegSecurityExchange)
}

//HasLegIssuer returns true if LegIssuer is present, Tag 617
func (m NoLegs) HasLegIssuer() bool {
	return m.Has(tag.LegIssuer)
}

//HasEncodedLegIssuerLen returns true if EncodedLegIssuerLen is present, Tag 618
func (m NoLegs) HasEncodedLegIssuerLen() bool {
	return m.Has(tag.EncodedLegIssuerLen)
}

//HasEncodedLegIssuer returns true if EncodedLegIssuer is present, Tag 619
func (m NoLegs) HasEncodedLegIssuer() bool {
	return m.Has(tag.EncodedLegIssuer)
}

//HasLegSecurityDesc returns true if LegSecurityDesc is present, Tag 620
func (m NoLegs) HasLegSecurityDesc() bool {
	return m.Has(tag.LegSecurityDesc)
}

//HasEncodedLegSecurityDescLen returns true if EncodedLegSecurityDescLen is present, Tag 621
func (m NoLegs) HasEncodedLegSecurityDescLen() bool {
	return m.Has(tag.EncodedLegSecurityDescLen)
}

//HasEncodedLegSecurityDesc returns true if EncodedLegSecurityDesc is present, Tag 622
func (m NoLegs) HasEncodedLegSecurityDesc() bool {
	return m.Has(tag.EncodedLegSecurityDesc)
}

//HasLegRatioQty returns true if LegRatioQty is present, Tag 623
func (m NoLegs) HasLegRatioQty() bool {
	return m.Has(tag.LegRatioQty)
}

//HasLegSide returns true if LegSide is present, Tag 624
func (m NoLegs) HasLegSide() bool {
	return m.Has(tag.LegSide)
}

//HasLegCurrency returns true if LegCurrency is present, Tag 556
func (m NoLegs) HasLegCurrency() bool {
	return m.Has(tag.LegCurrency)
}

//HasLegPool returns true if LegPool is present, Tag 740
func (m NoLegs) HasLegPool() bool {
	return m.Has(tag.LegPool)
}

//HasLegDatedDate returns true if LegDatedDate is present, Tag 739
func (m NoLegs) HasLegDatedDate() bool {
	return m.Has(tag.LegDatedDate)
}

//HasLegContractSettlMonth returns true if LegContractSettlMonth is present, Tag 955
func (m NoLegs) HasLegContractSettlMonth() bool {
	return m.Has(tag.LegContractSettlMonth)
}

//HasLegInterestAccrualDate returns true if LegInterestAccrualDate is present, Tag 956
func (m NoLegs) HasLegInterestAccrualDate() bool {
	return m.Has(tag.LegInterestAccrualDate)
}

//HasLegUnitOfMeasure returns true if LegUnitOfMeasure is present, Tag 999
func (m NoLegs) HasLegUnitOfMeasure() bool {
	return m.Has(tag.LegUnitOfMeasure)
}

//HasLegTimeUnit returns true if LegTimeUnit is present, Tag 1001
func (m NoLegs) HasLegTimeUnit() bool {
	return m.Has(tag.LegTimeUnit)
}

//HasLegQty returns true if LegQty is present, Tag 687
func (m NoLegs) HasLegQty() bool {
	return m.Has(tag.LegQty)
}

//HasLegSwapType returns true if LegSwapType is present, Tag 690
func (m NoLegs) HasLegSwapType() bool {
	return m.Has(tag.LegSwapType)
}

//HasNoLegStipulations returns true if NoLegStipulations is present, Tag 683
func (m NoLegs) HasNoLegStipulations() bool {
	return m.Has(tag.NoLegStipulations)
}

//HasNoLegAllocs returns true if NoLegAllocs is present, Tag 670
func (m NoLegs) HasNoLegAllocs() bool {
	return m.Has(tag.NoLegAllocs)
}

//HasLegPositionEffect returns true if LegPositionEffect is present, Tag 564
func (m NoLegs) HasLegPositionEffect() bool {
	return m.Has(tag.LegPositionEffect)
}

//HasLegCoveredOrUncovered returns true if LegCoveredOrUncovered is present, Tag 565
func (m NoLegs) HasLegCoveredOrUncovered() bool {
	return m.Has(tag.LegCoveredOrUncovered)
}

//HasNoNestedPartyIDs returns true if NoNestedPartyIDs is present, Tag 539
func (m NoLegs) HasNoNestedPartyIDs() bool {
	return m.Has(tag.NoNestedPartyIDs)
}

//HasLegRefID returns true if LegRefID is present, Tag 654
func (m NoLegs) HasLegRefID() bool {
	return m.Has(tag.LegRefID)
}

//HasLegPrice returns true if LegPrice is present, Tag 566
func (m NoLegs) HasLegPrice() bool {
	return m.Has(tag.LegPrice)
}

//HasLegSettlType returns true if LegSettlType is present, Tag 587
func (m NoLegs) HasLegSettlType() bool {
	return m.Has(tag.LegSettlType)
}

//HasLegSettlDate returns true if LegSettlDate is present, Tag 588
func (m NoLegs) HasLegSettlDate() bool {
	return m.Has(tag.LegSettlDate)
}

//HasLegOptionRatio returns true if LegOptionRatio is present, Tag 1017
func (m NoLegs) HasLegOptionRatio() bool {
	return m.Has(tag.LegOptionRatio)
}

//HasLegOrderQty returns true if LegOrderQty is present, Tag 685
func (m NoLegs) HasLegOrderQty() bool {
	return m.Has(tag.LegOrderQty)
}

//NoLegSecurityAltID is a repeating group element, Tag 604
type NoLegSecurityAltID struct {
	quickfix.Group
}

//SetLegSecurityAltID sets LegSecurityAltID, Tag 605
func (m NoLegSecurityAltID) SetLegSecurityAltID(v string) {
	m.Set(field.NewLegSecurityAltID(v))
}

//SetLegSecurityAltIDSource sets LegSecurityAltIDSource, Tag 606
func (m NoLegSecurityAltID) SetLegSecurityAltIDSource(v string) {
	m.Set(field.NewLegSecurityAltIDSource(v))
}

//GetLegSecurityAltID gets LegSecurityAltID, Tag 605
func (m NoLegSecurityAltID) GetLegSecurityAltID() (f field.LegSecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityAltIDSource gets LegSecurityAltIDSource, Tag 606
func (m NoLegSecurityAltID) GetLegSecurityAltIDSource() (f field.LegSecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasLegSecurityAltID returns true if LegSecurityAltID is present, Tag 605
func (m NoLegSecurityAltID) HasLegSecurityAltID() bool {
	return m.Has(tag.LegSecurityAltID)
}

//HasLegSecurityAltIDSource returns true if LegSecurityAltIDSource is present, Tag 606
func (m NoLegSecurityAltID) HasLegSecurityAltIDSource() bool {
	return m.Has(tag.LegSecurityAltIDSource)
}

//NoLegSecurityAltIDRepeatingGroup is a repeating group, Tag 604
type NoLegSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegSecurityAltIDRepeatingGroup returns an initialized, NoLegSecurityAltIDRepeatingGroup
func NewNoLegSecurityAltIDRepeatingGroup() NoLegSecurityAltIDRepeatingGroup {
	return NoLegSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSecurityAltID), quickfix.GroupElement(tag.LegSecurityAltIDSource)})}
}

//Add create and append a new NoLegSecurityAltID to this group
func (m NoLegSecurityAltIDRepeatingGroup) Add() NoLegSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoLegSecurityAltID{g}
}

//Get returns the ith NoLegSecurityAltID in the NoLegSecurityAltIDRepeatinGroup
func (m NoLegSecurityAltIDRepeatingGroup) Get(i int) NoLegSecurityAltID {
	return NoLegSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoLegStipulations is a repeating group element, Tag 683
type NoLegStipulations struct {
	quickfix.Group
}

//SetLegStipulationType sets LegStipulationType, Tag 688
func (m NoLegStipulations) SetLegStipulationType(v string) {
	m.Set(field.NewLegStipulationType(v))
}

//SetLegStipulationValue sets LegStipulationValue, Tag 689
func (m NoLegStipulations) SetLegStipulationValue(v string) {
	m.Set(field.NewLegStipulationValue(v))
}

//GetLegStipulationType gets LegStipulationType, Tag 688
func (m NoLegStipulations) GetLegStipulationType() (f field.LegStipulationTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegStipulationValue gets LegStipulationValue, Tag 689
func (m NoLegStipulations) GetLegStipulationValue() (f field.LegStipulationValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasLegStipulationType returns true if LegStipulationType is present, Tag 688
func (m NoLegStipulations) HasLegStipulationType() bool {
	return m.Has(tag.LegStipulationType)
}

//HasLegStipulationValue returns true if LegStipulationValue is present, Tag 689
func (m NoLegStipulations) HasLegStipulationValue() bool {
	return m.Has(tag.LegStipulationValue)
}

//NoLegStipulationsRepeatingGroup is a repeating group, Tag 683
type NoLegStipulationsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegStipulationsRepeatingGroup returns an initialized, NoLegStipulationsRepeatingGroup
func NewNoLegStipulationsRepeatingGroup() NoLegStipulationsRepeatingGroup {
	return NoLegStipulationsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegStipulations,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegStipulationType), quickfix.GroupElement(tag.LegStipulationValue)})}
}

//Add create and append a new NoLegStipulations to this group
func (m NoLegStipulationsRepeatingGroup) Add() NoLegStipulations {
	g := m.RepeatingGroup.Add()
	return NoLegStipulations{g}
}

//Get returns the ith NoLegStipulations in the NoLegStipulationsRepeatinGroup
func (m NoLegStipulationsRepeatingGroup) Get(i int) NoLegStipulations {
	return NoLegStipulations{m.RepeatingGroup.Get(i)}
}

//NoLegAllocs is a repeating group element, Tag 670
type NoLegAllocs struct {
	quickfix.Group
}

//SetLegAllocAccount sets LegAllocAccount, Tag 671
func (m NoLegAllocs) SetLegAllocAccount(v string) {
	m.Set(field.NewLegAllocAccount(v))
}

//SetLegIndividualAllocID sets LegIndividualAllocID, Tag 672
func (m NoLegAllocs) SetLegIndividualAllocID(v string) {
	m.Set(field.NewLegIndividualAllocID(v))
}

//SetNoNested2PartyIDs sets NoNested2PartyIDs, Tag 756
func (m NoLegAllocs) SetNoNested2PartyIDs(f NoNested2PartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetLegAllocQty sets LegAllocQty, Tag 673
func (m NoLegAllocs) SetLegAllocQty(v float64) {
	m.Set(field.NewLegAllocQty(v))
}

//SetLegAllocAcctIDSource sets LegAllocAcctIDSource, Tag 674
func (m NoLegAllocs) SetLegAllocAcctIDSource(v string) {
	m.Set(field.NewLegAllocAcctIDSource(v))
}

//SetLegSettlCurrency sets LegSettlCurrency, Tag 675
func (m NoLegAllocs) SetLegSettlCurrency(v string) {
	m.Set(field.NewLegSettlCurrency(v))
}

//GetLegAllocAccount gets LegAllocAccount, Tag 671
func (m NoLegAllocs) GetLegAllocAccount() (f field.LegAllocAccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegIndividualAllocID gets LegIndividualAllocID, Tag 672
func (m NoLegAllocs) GetLegIndividualAllocID() (f field.LegIndividualAllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoNested2PartyIDs gets NoNested2PartyIDs, Tag 756
func (m NoLegAllocs) GetNoNested2PartyIDs() (NoNested2PartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNested2PartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLegAllocQty gets LegAllocQty, Tag 673
func (m NoLegAllocs) GetLegAllocQty() (f field.LegAllocQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegAllocAcctIDSource gets LegAllocAcctIDSource, Tag 674
func (m NoLegAllocs) GetLegAllocAcctIDSource() (f field.LegAllocAcctIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSettlCurrency gets LegSettlCurrency, Tag 675
func (m NoLegAllocs) GetLegSettlCurrency() (f field.LegSettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasLegAllocAccount returns true if LegAllocAccount is present, Tag 671
func (m NoLegAllocs) HasLegAllocAccount() bool {
	return m.Has(tag.LegAllocAccount)
}

//HasLegIndividualAllocID returns true if LegIndividualAllocID is present, Tag 672
func (m NoLegAllocs) HasLegIndividualAllocID() bool {
	return m.Has(tag.LegIndividualAllocID)
}

//HasNoNested2PartyIDs returns true if NoNested2PartyIDs is present, Tag 756
func (m NoLegAllocs) HasNoNested2PartyIDs() bool {
	return m.Has(tag.NoNested2PartyIDs)
}

//HasLegAllocQty returns true if LegAllocQty is present, Tag 673
func (m NoLegAllocs) HasLegAllocQty() bool {
	return m.Has(tag.LegAllocQty)
}

//HasLegAllocAcctIDSource returns true if LegAllocAcctIDSource is present, Tag 674
func (m NoLegAllocs) HasLegAllocAcctIDSource() bool {
	return m.Has(tag.LegAllocAcctIDSource)
}

//HasLegSettlCurrency returns true if LegSettlCurrency is present, Tag 675
func (m NoLegAllocs) HasLegSettlCurrency() bool {
	return m.Has(tag.LegSettlCurrency)
}

//NoNested2PartyIDs is a repeating group element, Tag 756
type NoNested2PartyIDs struct {
	quickfix.Group
}

//SetNested2PartyID sets Nested2PartyID, Tag 757
func (m NoNested2PartyIDs) SetNested2PartyID(v string) {
	m.Set(field.NewNested2PartyID(v))
}

//SetNested2PartyIDSource sets Nested2PartyIDSource, Tag 758
func (m NoNested2PartyIDs) SetNested2PartyIDSource(v string) {
	m.Set(field.NewNested2PartyIDSource(v))
}

//SetNested2PartyRole sets Nested2PartyRole, Tag 759
func (m NoNested2PartyIDs) SetNested2PartyRole(v int) {
	m.Set(field.NewNested2PartyRole(v))
}

//SetNoNested2PartySubIDs sets NoNested2PartySubIDs, Tag 806
func (m NoNested2PartyIDs) SetNoNested2PartySubIDs(f NoNested2PartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetNested2PartyID gets Nested2PartyID, Tag 757
func (m NoNested2PartyIDs) GetNested2PartyID() (f field.Nested2PartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNested2PartyIDSource gets Nested2PartyIDSource, Tag 758
func (m NoNested2PartyIDs) GetNested2PartyIDSource() (f field.Nested2PartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNested2PartyRole gets Nested2PartyRole, Tag 759
func (m NoNested2PartyIDs) GetNested2PartyRole() (f field.Nested2PartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoNested2PartySubIDs gets NoNested2PartySubIDs, Tag 806
func (m NoNested2PartyIDs) GetNoNested2PartySubIDs() (NoNested2PartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNested2PartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasNested2PartyID returns true if Nested2PartyID is present, Tag 757
func (m NoNested2PartyIDs) HasNested2PartyID() bool {
	return m.Has(tag.Nested2PartyID)
}

//HasNested2PartyIDSource returns true if Nested2PartyIDSource is present, Tag 758
func (m NoNested2PartyIDs) HasNested2PartyIDSource() bool {
	return m.Has(tag.Nested2PartyIDSource)
}

//HasNested2PartyRole returns true if Nested2PartyRole is present, Tag 759
func (m NoNested2PartyIDs) HasNested2PartyRole() bool {
	return m.Has(tag.Nested2PartyRole)
}

//HasNoNested2PartySubIDs returns true if NoNested2PartySubIDs is present, Tag 806
func (m NoNested2PartyIDs) HasNoNested2PartySubIDs() bool {
	return m.Has(tag.NoNested2PartySubIDs)
}

//NoNested2PartySubIDs is a repeating group element, Tag 806
type NoNested2PartySubIDs struct {
	quickfix.Group
}

//SetNested2PartySubID sets Nested2PartySubID, Tag 760
func (m NoNested2PartySubIDs) SetNested2PartySubID(v string) {
	m.Set(field.NewNested2PartySubID(v))
}

//SetNested2PartySubIDType sets Nested2PartySubIDType, Tag 807
func (m NoNested2PartySubIDs) SetNested2PartySubIDType(v int) {
	m.Set(field.NewNested2PartySubIDType(v))
}

//GetNested2PartySubID gets Nested2PartySubID, Tag 760
func (m NoNested2PartySubIDs) GetNested2PartySubID() (f field.Nested2PartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNested2PartySubIDType gets Nested2PartySubIDType, Tag 807
func (m NoNested2PartySubIDs) GetNested2PartySubIDType() (f field.Nested2PartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasNested2PartySubID returns true if Nested2PartySubID is present, Tag 760
func (m NoNested2PartySubIDs) HasNested2PartySubID() bool {
	return m.Has(tag.Nested2PartySubID)
}

//HasNested2PartySubIDType returns true if Nested2PartySubIDType is present, Tag 807
func (m NoNested2PartySubIDs) HasNested2PartySubIDType() bool {
	return m.Has(tag.Nested2PartySubIDType)
}

//NoNested2PartySubIDsRepeatingGroup is a repeating group, Tag 806
type NoNested2PartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNested2PartySubIDsRepeatingGroup returns an initialized, NoNested2PartySubIDsRepeatingGroup
func NewNoNested2PartySubIDsRepeatingGroup() NoNested2PartySubIDsRepeatingGroup {
	return NoNested2PartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNested2PartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Nested2PartySubID), quickfix.GroupElement(tag.Nested2PartySubIDType)})}
}

//Add create and append a new NoNested2PartySubIDs to this group
func (m NoNested2PartySubIDsRepeatingGroup) Add() NoNested2PartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoNested2PartySubIDs{g}
}

//Get returns the ith NoNested2PartySubIDs in the NoNested2PartySubIDsRepeatinGroup
func (m NoNested2PartySubIDsRepeatingGroup) Get(i int) NoNested2PartySubIDs {
	return NoNested2PartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoNested2PartyIDsRepeatingGroup is a repeating group, Tag 756
type NoNested2PartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNested2PartyIDsRepeatingGroup returns an initialized, NoNested2PartyIDsRepeatingGroup
func NewNoNested2PartyIDsRepeatingGroup() NoNested2PartyIDsRepeatingGroup {
	return NoNested2PartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNested2PartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Nested2PartyID), quickfix.GroupElement(tag.Nested2PartyIDSource), quickfix.GroupElement(tag.Nested2PartyRole), NewNoNested2PartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoNested2PartyIDs to this group
func (m NoNested2PartyIDsRepeatingGroup) Add() NoNested2PartyIDs {
	g := m.RepeatingGroup.Add()
	return NoNested2PartyIDs{g}
}

//Get returns the ith NoNested2PartyIDs in the NoNested2PartyIDsRepeatinGroup
func (m NoNested2PartyIDsRepeatingGroup) Get(i int) NoNested2PartyIDs {
	return NoNested2PartyIDs{m.RepeatingGroup.Get(i)}
}

//NoLegAllocsRepeatingGroup is a repeating group, Tag 670
type NoLegAllocsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegAllocsRepeatingGroup returns an initialized, NoLegAllocsRepeatingGroup
func NewNoLegAllocsRepeatingGroup() NoLegAllocsRepeatingGroup {
	return NoLegAllocsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegAllocs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegAllocAccount), quickfix.GroupElement(tag.LegIndividualAllocID), NewNoNested2PartyIDsRepeatingGroup(), quickfix.GroupElement(tag.LegAllocQty), quickfix.GroupElement(tag.LegAllocAcctIDSource), quickfix.GroupElement(tag.LegSettlCurrency)})}
}

//Add create and append a new NoLegAllocs to this group
func (m NoLegAllocsRepeatingGroup) Add() NoLegAllocs {
	g := m.RepeatingGroup.Add()
	return NoLegAllocs{g}
}

//Get returns the ith NoLegAllocs in the NoLegAllocsRepeatinGroup
func (m NoLegAllocsRepeatingGroup) Get(i int) NoLegAllocs {
	return NoLegAllocs{m.RepeatingGroup.Get(i)}
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

//NoLegsRepeatingGroup is a repeating group, Tag 555
type NoLegsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegsRepeatingGroup returns an initialized, NoLegsRepeatingGroup
func NewNoLegsRepeatingGroup() NoLegsRepeatingGroup {
	return NoLegsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSymbol), quickfix.GroupElement(tag.LegSymbolSfx), quickfix.GroupElement(tag.LegSecurityID), quickfix.GroupElement(tag.LegSecurityIDSource), NewNoLegSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.LegProduct), quickfix.GroupElement(tag.LegCFICode), quickfix.GroupElement(tag.LegSecurityType), quickfix.GroupElement(tag.LegSecuritySubType), quickfix.GroupElement(tag.LegMaturityMonthYear), quickfix.GroupElement(tag.LegMaturityDate), quickfix.GroupElement(tag.LegCouponPaymentDate), quickfix.GroupElement(tag.LegIssueDate), quickfix.GroupElement(tag.LegRepoCollateralSecurityType), quickfix.GroupElement(tag.LegRepurchaseTerm), quickfix.GroupElement(tag.LegRepurchaseRate), quickfix.GroupElement(tag.LegFactor), quickfix.GroupElement(tag.LegCreditRating), quickfix.GroupElement(tag.LegInstrRegistry), quickfix.GroupElement(tag.LegCountryOfIssue), quickfix.GroupElement(tag.LegStateOrProvinceOfIssue), quickfix.GroupElement(tag.LegLocaleOfIssue), quickfix.GroupElement(tag.LegRedemptionDate), quickfix.GroupElement(tag.LegStrikePrice), quickfix.GroupElement(tag.LegStrikeCurrency), quickfix.GroupElement(tag.LegOptAttribute), quickfix.GroupElement(tag.LegContractMultiplier), quickfix.GroupElement(tag.LegCouponRate), quickfix.GroupElement(tag.LegSecurityExchange), quickfix.GroupElement(tag.LegIssuer), quickfix.GroupElement(tag.EncodedLegIssuerLen), quickfix.GroupElement(tag.EncodedLegIssuer), quickfix.GroupElement(tag.LegSecurityDesc), quickfix.GroupElement(tag.EncodedLegSecurityDescLen), quickfix.GroupElement(tag.EncodedLegSecurityDesc), quickfix.GroupElement(tag.LegRatioQty), quickfix.GroupElement(tag.LegSide), quickfix.GroupElement(tag.LegCurrency), quickfix.GroupElement(tag.LegPool), quickfix.GroupElement(tag.LegDatedDate), quickfix.GroupElement(tag.LegContractSettlMonth), quickfix.GroupElement(tag.LegInterestAccrualDate), quickfix.GroupElement(tag.LegUnitOfMeasure), quickfix.GroupElement(tag.LegTimeUnit), quickfix.GroupElement(tag.LegQty), quickfix.GroupElement(tag.LegSwapType), NewNoLegStipulationsRepeatingGroup(), NewNoLegAllocsRepeatingGroup(), quickfix.GroupElement(tag.LegPositionEffect), quickfix.GroupElement(tag.LegCoveredOrUncovered), NewNoNestedPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.LegRefID), quickfix.GroupElement(tag.LegPrice), quickfix.GroupElement(tag.LegSettlType), quickfix.GroupElement(tag.LegSettlDate), quickfix.GroupElement(tag.LegOptionRatio), quickfix.GroupElement(tag.LegOrderQty)})}
}

//Add create and append a new NoLegs to this group
func (m NoLegsRepeatingGroup) Add() NoLegs {
	g := m.RepeatingGroup.Add()
	return NoLegs{g}
}

//Get returns the ith NoLegs in the NoLegsRepeatinGroup
func (m NoLegsRepeatingGroup) Get(i int) NoLegs {
	return NoLegs{m.RepeatingGroup.Get(i)}
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

//SetUndlyInstrumentPartyID sets UndlyInstrumentPartyID, Tag 1059
func (m NoUndlyInstrumentParties) SetUndlyInstrumentPartyID(v string) {
	m.Set(field.NewUndlyInstrumentPartyID(v))
}

//SetUndlyInstrumentPartyIDSource sets UndlyInstrumentPartyIDSource, Tag 1060
func (m NoUndlyInstrumentParties) SetUndlyInstrumentPartyIDSource(v string) {
	m.Set(field.NewUndlyInstrumentPartyIDSource(v))
}

//SetUndlyInstrumentPartyRole sets UndlyInstrumentPartyRole, Tag 1061
func (m NoUndlyInstrumentParties) SetUndlyInstrumentPartyRole(v int) {
	m.Set(field.NewUndlyInstrumentPartyRole(v))
}

//SetNoUndlyInstrumentPartySubIDs sets NoUndlyInstrumentPartySubIDs, Tag 1062
func (m NoUndlyInstrumentParties) SetNoUndlyInstrumentPartySubIDs(f NoUndlyInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetUndlyInstrumentPartyID gets UndlyInstrumentPartyID, Tag 1059
func (m NoUndlyInstrumentParties) GetUndlyInstrumentPartyID() (f field.UndlyInstrumentPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUndlyInstrumentPartyIDSource gets UndlyInstrumentPartyIDSource, Tag 1060
func (m NoUndlyInstrumentParties) GetUndlyInstrumentPartyIDSource() (f field.UndlyInstrumentPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUndlyInstrumentPartyRole gets UndlyInstrumentPartyRole, Tag 1061
func (m NoUndlyInstrumentParties) GetUndlyInstrumentPartyRole() (f field.UndlyInstrumentPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUndlyInstrumentPartySubIDs gets NoUndlyInstrumentPartySubIDs, Tag 1062
func (m NoUndlyInstrumentParties) GetNoUndlyInstrumentPartySubIDs() (NoUndlyInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasUndlyInstrumentPartyID returns true if UndlyInstrumentPartyID is present, Tag 1059
func (m NoUndlyInstrumentParties) HasUndlyInstrumentPartyID() bool {
	return m.Has(tag.UndlyInstrumentPartyID)
}

//HasUndlyInstrumentPartyIDSource returns true if UndlyInstrumentPartyIDSource is present, Tag 1060
func (m NoUndlyInstrumentParties) HasUndlyInstrumentPartyIDSource() bool {
	return m.Has(tag.UndlyInstrumentPartyIDSource)
}

//HasUndlyInstrumentPartyRole returns true if UndlyInstrumentPartyRole is present, Tag 1061
func (m NoUndlyInstrumentParties) HasUndlyInstrumentPartyRole() bool {
	return m.Has(tag.UndlyInstrumentPartyRole)
}

//HasNoUndlyInstrumentPartySubIDs returns true if NoUndlyInstrumentPartySubIDs is present, Tag 1062
func (m NoUndlyInstrumentParties) HasNoUndlyInstrumentPartySubIDs() bool {
	return m.Has(tag.NoUndlyInstrumentPartySubIDs)
}

//NoUndlyInstrumentPartySubIDs is a repeating group element, Tag 1062
type NoUndlyInstrumentPartySubIDs struct {
	quickfix.Group
}

//SetUndlyInstrumentPartySubID sets UndlyInstrumentPartySubID, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) SetUndlyInstrumentPartySubID(v string) {
	m.Set(field.NewUndlyInstrumentPartySubID(v))
}

//SetUndlyInstrumentPartySubIDType sets UndlyInstrumentPartySubIDType, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) SetUndlyInstrumentPartySubIDType(v int) {
	m.Set(field.NewUndlyInstrumentPartySubIDType(v))
}

//GetUndlyInstrumentPartySubID gets UndlyInstrumentPartySubID, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) GetUndlyInstrumentPartySubID() (f field.UndlyInstrumentPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUndlyInstrumentPartySubIDType gets UndlyInstrumentPartySubIDType, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) GetUndlyInstrumentPartySubIDType() (f field.UndlyInstrumentPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUndlyInstrumentPartySubID returns true if UndlyInstrumentPartySubID is present, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) HasUndlyInstrumentPartySubID() bool {
	return m.Has(tag.UndlyInstrumentPartySubID)
}

//HasUndlyInstrumentPartySubIDType returns true if UndlyInstrumentPartySubIDType is present, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) HasUndlyInstrumentPartySubIDType() bool {
	return m.Has(tag.UndlyInstrumentPartySubIDType)
}

//NoUndlyInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1062
type NoUndlyInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUndlyInstrumentPartySubIDsRepeatingGroup returns an initialized, NoUndlyInstrumentPartySubIDsRepeatingGroup
func NewNoUndlyInstrumentPartySubIDsRepeatingGroup() NoUndlyInstrumentPartySubIDsRepeatingGroup {
	return NoUndlyInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UndlyInstrumentPartySubID), quickfix.GroupElement(tag.UndlyInstrumentPartySubIDType)})}
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UndlyInstrumentPartyID), quickfix.GroupElement(tag.UndlyInstrumentPartyIDSource), quickfix.GroupElement(tag.UndlyInstrumentPartyRole), NewNoUndlyInstrumentPartySubIDsRepeatingGroup()})}
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingSecurityIDSource), NewNoUnderlyingSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingProduct), quickfix.GroupElement(tag.UnderlyingCFICode), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingSecuritySubType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDate), quickfix.GroupElement(tag.UnderlyingCouponPaymentDate), quickfix.GroupElement(tag.UnderlyingIssueDate), quickfix.GroupElement(tag.UnderlyingRepoCollateralSecurityType), quickfix.GroupElement(tag.UnderlyingRepurchaseTerm), quickfix.GroupElement(tag.UnderlyingRepurchaseRate), quickfix.GroupElement(tag.UnderlyingFactor), quickfix.GroupElement(tag.UnderlyingCreditRating), quickfix.GroupElement(tag.UnderlyingInstrRegistry), quickfix.GroupElement(tag.UnderlyingCountryOfIssue), quickfix.GroupElement(tag.UnderlyingStateOrProvinceOfIssue), quickfix.GroupElement(tag.UnderlyingLocaleOfIssue), quickfix.GroupElement(tag.UnderlyingRedemptionDate), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingStrikeCurrency), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.UnderlyingCPProgram), quickfix.GroupElement(tag.UnderlyingCPRegType), quickfix.GroupElement(tag.UnderlyingCurrency), quickfix.GroupElement(tag.UnderlyingQty), quickfix.GroupElement(tag.UnderlyingPx), quickfix.GroupElement(tag.UnderlyingDirtyPrice), quickfix.GroupElement(tag.UnderlyingEndPrice), quickfix.GroupElement(tag.UnderlyingStartValue), quickfix.GroupElement(tag.UnderlyingCurrentValue), quickfix.GroupElement(tag.UnderlyingEndValue), NewNoUnderlyingStipsRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingAllocationPercent), quickfix.GroupElement(tag.UnderlyingSettlementType), quickfix.GroupElement(tag.UnderlyingCashAmount), quickfix.GroupElement(tag.UnderlyingCashType), quickfix.GroupElement(tag.UnderlyingUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingTimeUnit), quickfix.GroupElement(tag.UnderlyingCapValue), NewNoUndlyInstrumentPartiesRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingSettlMethod), quickfix.GroupElement(tag.UnderlyingAdjustedQuantity), quickfix.GroupElement(tag.UnderlyingFXRate), quickfix.GroupElement(tag.UnderlyingFXRateCalc)})}
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
func (m NoEvents) SetEventPx(v float64) {
	m.Set(field.NewEventPx(v))
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartyID), quickfix.GroupElement(tag.InstrumentPartyIDSource), quickfix.GroupElement(tag.InstrumentPartyRole), NewNoInstrumentPartySubIDsRepeatingGroup()})}
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
