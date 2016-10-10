package crossordercancelreplacerequest

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//CrossOrderCancelReplaceRequest is the fix50 CrossOrderCancelReplaceRequest type, MsgType = t
type CrossOrderCancelReplaceRequest struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

//FromMessage creates a CrossOrderCancelReplaceRequest from a quickfix.Message instance
func FromMessage(m *quickfix.Message) CrossOrderCancelReplaceRequest {
	return CrossOrderCancelReplaceRequest{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m CrossOrderCancelReplaceRequest) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a CrossOrderCancelReplaceRequest initialized with the required fields for CrossOrderCancelReplaceRequest
func New(crossid field.CrossIDField, origcrossid field.OrigCrossIDField, crosstype field.CrossTypeField, crossprioritization field.CrossPrioritizationField, transacttime field.TransactTimeField, ordtype field.OrdTypeField) (m CrossOrderCancelReplaceRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("t"))
	m.Set(crossid)
	m.Set(origcrossid)
	m.Set(crosstype)
	m.Set(crossprioritization)
	m.Set(transacttime)
	m.Set(ordtype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg CrossOrderCancelReplaceRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "7", "t", r
}

//SetCurrency sets Currency, Tag 15
func (m CrossOrderCancelReplaceRequest) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetExecInst sets ExecInst, Tag 18
func (m CrossOrderCancelReplaceRequest) SetExecInst(v enum.ExecInst) {
	m.Set(field.NewExecInst(v))
}

//SetHandlInst sets HandlInst, Tag 21
func (m CrossOrderCancelReplaceRequest) SetHandlInst(v enum.HandlInst) {
	m.Set(field.NewHandlInst(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m CrossOrderCancelReplaceRequest) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetIOIID sets IOIID, Tag 23
func (m CrossOrderCancelReplaceRequest) SetIOIID(v string) {
	m.Set(field.NewIOIID(v))
}

//SetOrderID sets OrderID, Tag 37
func (m CrossOrderCancelReplaceRequest) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetOrdType sets OrdType, Tag 40
func (m CrossOrderCancelReplaceRequest) SetOrdType(v enum.OrdType) {
	m.Set(field.NewOrdType(v))
}

//SetPrice sets Price, Tag 44
func (m CrossOrderCancelReplaceRequest) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

//SetSecurityID sets SecurityID, Tag 48
func (m CrossOrderCancelReplaceRequest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m CrossOrderCancelReplaceRequest) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetTimeInForce sets TimeInForce, Tag 59
func (m CrossOrderCancelReplaceRequest) SetTimeInForce(v enum.TimeInForce) {
	m.Set(field.NewTimeInForce(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m CrossOrderCancelReplaceRequest) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSettlType sets SettlType, Tag 63
func (m CrossOrderCancelReplaceRequest) SetSettlType(v enum.SettlType) {
	m.Set(field.NewSettlType(v))
}

//SetSettlDate sets SettlDate, Tag 64
func (m CrossOrderCancelReplaceRequest) SetSettlDate(v string) {
	m.Set(field.NewSettlDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m CrossOrderCancelReplaceRequest) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

//SetProcessCode sets ProcessCode, Tag 81
func (m CrossOrderCancelReplaceRequest) SetProcessCode(v enum.ProcessCode) {
	m.Set(field.NewProcessCode(v))
}

//SetStopPx sets StopPx, Tag 99
func (m CrossOrderCancelReplaceRequest) SetStopPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewStopPx(value, scale))
}

//SetExDestination sets ExDestination, Tag 100
func (m CrossOrderCancelReplaceRequest) SetExDestination(v enum.ExDestination) {
	m.Set(field.NewExDestination(v))
}

//SetIssuer sets Issuer, Tag 106
func (m CrossOrderCancelReplaceRequest) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m CrossOrderCancelReplaceRequest) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetMinQty sets MinQty, Tag 110
func (m CrossOrderCancelReplaceRequest) SetMinQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinQty(value, scale))
}

//SetMaxFloor sets MaxFloor, Tag 111
func (m CrossOrderCancelReplaceRequest) SetMaxFloor(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxFloor(value, scale))
}

//SetLocateReqd sets LocateReqd, Tag 114
func (m CrossOrderCancelReplaceRequest) SetLocateReqd(v bool) {
	m.Set(field.NewLocateReqd(v))
}

//SetQuoteID sets QuoteID, Tag 117
func (m CrossOrderCancelReplaceRequest) SetQuoteID(v string) {
	m.Set(field.NewQuoteID(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m CrossOrderCancelReplaceRequest) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetPrevClosePx sets PrevClosePx, Tag 140
func (m CrossOrderCancelReplaceRequest) SetPrevClosePx(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrevClosePx(value, scale))
}

//SetSecurityType sets SecurityType, Tag 167
func (m CrossOrderCancelReplaceRequest) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

//SetEffectiveTime sets EffectiveTime, Tag 168
func (m CrossOrderCancelReplaceRequest) SetEffectiveTime(v time.Time) {
	m.Set(field.NewEffectiveTime(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m CrossOrderCancelReplaceRequest) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m CrossOrderCancelReplaceRequest) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m CrossOrderCancelReplaceRequest) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m CrossOrderCancelReplaceRequest) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetMaxShow sets MaxShow, Tag 210
func (m CrossOrderCancelReplaceRequest) SetMaxShow(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxShow(value, scale))
}

//SetPegOffsetValue sets PegOffsetValue, Tag 211
func (m CrossOrderCancelReplaceRequest) SetPegOffsetValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewPegOffsetValue(value, scale))
}

//SetSpread sets Spread, Tag 218
func (m CrossOrderCancelReplaceRequest) SetSpread(value decimal.Decimal, scale int32) {
	m.Set(field.NewSpread(value, scale))
}

//SetBenchmarkCurveCurrency sets BenchmarkCurveCurrency, Tag 220
func (m CrossOrderCancelReplaceRequest) SetBenchmarkCurveCurrency(v string) {
	m.Set(field.NewBenchmarkCurveCurrency(v))
}

//SetBenchmarkCurveName sets BenchmarkCurveName, Tag 221
func (m CrossOrderCancelReplaceRequest) SetBenchmarkCurveName(v enum.BenchmarkCurveName) {
	m.Set(field.NewBenchmarkCurveName(v))
}

//SetBenchmarkCurvePoint sets BenchmarkCurvePoint, Tag 222
func (m CrossOrderCancelReplaceRequest) SetBenchmarkCurvePoint(v string) {
	m.Set(field.NewBenchmarkCurvePoint(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m CrossOrderCancelReplaceRequest) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m CrossOrderCancelReplaceRequest) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m CrossOrderCancelReplaceRequest) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m CrossOrderCancelReplaceRequest) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m CrossOrderCancelReplaceRequest) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

//SetFactor sets Factor, Tag 228
func (m CrossOrderCancelReplaceRequest) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m CrossOrderCancelReplaceRequest) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetNoStipulations sets NoStipulations, Tag 232
func (m CrossOrderCancelReplaceRequest) SetNoStipulations(f NoStipulationsRepeatingGroup) {
	m.SetGroup(f)
}

//SetYieldType sets YieldType, Tag 235
func (m CrossOrderCancelReplaceRequest) SetYieldType(v enum.YieldType) {
	m.Set(field.NewYieldType(v))
}

//SetYield sets Yield, Tag 236
func (m CrossOrderCancelReplaceRequest) SetYield(value decimal.Decimal, scale int32) {
	m.Set(field.NewYield(value, scale))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m CrossOrderCancelReplaceRequest) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m CrossOrderCancelReplaceRequest) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m CrossOrderCancelReplaceRequest) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m CrossOrderCancelReplaceRequest) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m CrossOrderCancelReplaceRequest) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m CrossOrderCancelReplaceRequest) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m CrossOrderCancelReplaceRequest) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetComplianceID sets ComplianceID, Tag 376
func (m CrossOrderCancelReplaceRequest) SetComplianceID(v string) {
	m.Set(field.NewComplianceID(v))
}

//SetNoTradingSessions sets NoTradingSessions, Tag 386
func (m CrossOrderCancelReplaceRequest) SetNoTradingSessions(f NoTradingSessionsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDiscretionInst sets DiscretionInst, Tag 388
func (m CrossOrderCancelReplaceRequest) SetDiscretionInst(v enum.DiscretionInst) {
	m.Set(field.NewDiscretionInst(v))
}

//SetDiscretionOffsetValue sets DiscretionOffsetValue, Tag 389
func (m CrossOrderCancelReplaceRequest) SetDiscretionOffsetValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewDiscretionOffsetValue(value, scale))
}

//SetPriceType sets PriceType, Tag 423
func (m CrossOrderCancelReplaceRequest) SetPriceType(v enum.PriceType) {
	m.Set(field.NewPriceType(v))
}

//SetGTBookingInst sets GTBookingInst, Tag 427
func (m CrossOrderCancelReplaceRequest) SetGTBookingInst(v enum.GTBookingInst) {
	m.Set(field.NewGTBookingInst(v))
}

//SetExpireDate sets ExpireDate, Tag 432
func (m CrossOrderCancelReplaceRequest) SetExpireDate(v string) {
	m.Set(field.NewExpireDate(v))
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m CrossOrderCancelReplaceRequest) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m CrossOrderCancelReplaceRequest) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m CrossOrderCancelReplaceRequest) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m CrossOrderCancelReplaceRequest) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m CrossOrderCancelReplaceRequest) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m CrossOrderCancelReplaceRequest) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetCancellationRights sets CancellationRights, Tag 480
func (m CrossOrderCancelReplaceRequest) SetCancellationRights(v enum.CancellationRights) {
	m.Set(field.NewCancellationRights(v))
}

//SetMoneyLaunderingStatus sets MoneyLaunderingStatus, Tag 481
func (m CrossOrderCancelReplaceRequest) SetMoneyLaunderingStatus(v enum.MoneyLaunderingStatus) {
	m.Set(field.NewMoneyLaunderingStatus(v))
}

//SetTransBkdTime sets TransBkdTime, Tag 483
func (m CrossOrderCancelReplaceRequest) SetTransBkdTime(v time.Time) {
	m.Set(field.NewTransBkdTime(v))
}

//SetDesignation sets Designation, Tag 494
func (m CrossOrderCancelReplaceRequest) SetDesignation(v string) {
	m.Set(field.NewDesignation(v))
}

//SetRegistID sets RegistID, Tag 513
func (m CrossOrderCancelReplaceRequest) SetRegistID(v string) {
	m.Set(field.NewRegistID(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m CrossOrderCancelReplaceRequest) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m CrossOrderCancelReplaceRequest) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

//SetCrossID sets CrossID, Tag 548
func (m CrossOrderCancelReplaceRequest) SetCrossID(v string) {
	m.Set(field.NewCrossID(v))
}

//SetCrossType sets CrossType, Tag 549
func (m CrossOrderCancelReplaceRequest) SetCrossType(v enum.CrossType) {
	m.Set(field.NewCrossType(v))
}

//SetCrossPrioritization sets CrossPrioritization, Tag 550
func (m CrossOrderCancelReplaceRequest) SetCrossPrioritization(v enum.CrossPrioritization) {
	m.Set(field.NewCrossPrioritization(v))
}

//SetOrigCrossID sets OrigCrossID, Tag 551
func (m CrossOrderCancelReplaceRequest) SetOrigCrossID(v string) {
	m.Set(field.NewOrigCrossID(v))
}

//SetNoSides sets NoSides, Tag 552
func (m CrossOrderCancelReplaceRequest) SetNoSides(f NoSidesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoLegs sets NoLegs, Tag 555
func (m CrossOrderCancelReplaceRequest) SetNoLegs(f NoLegsRepeatingGroup) {
	m.SetGroup(f)
}

//SetBenchmarkPrice sets BenchmarkPrice, Tag 662
func (m CrossOrderCancelReplaceRequest) SetBenchmarkPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewBenchmarkPrice(value, scale))
}

//SetBenchmarkPriceType sets BenchmarkPriceType, Tag 663
func (m CrossOrderCancelReplaceRequest) SetBenchmarkPriceType(v int) {
	m.Set(field.NewBenchmarkPriceType(v))
}

//SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m CrossOrderCancelReplaceRequest) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

//SetPool sets Pool, Tag 691
func (m CrossOrderCancelReplaceRequest) SetPool(v string) {
	m.Set(field.NewPool(v))
}

//SetYieldRedemptionDate sets YieldRedemptionDate, Tag 696
func (m CrossOrderCancelReplaceRequest) SetYieldRedemptionDate(v string) {
	m.Set(field.NewYieldRedemptionDate(v))
}

//SetYieldRedemptionPrice sets YieldRedemptionPrice, Tag 697
func (m CrossOrderCancelReplaceRequest) SetYieldRedemptionPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewYieldRedemptionPrice(value, scale))
}

//SetYieldRedemptionPriceType sets YieldRedemptionPriceType, Tag 698
func (m CrossOrderCancelReplaceRequest) SetYieldRedemptionPriceType(v int) {
	m.Set(field.NewYieldRedemptionPriceType(v))
}

//SetBenchmarkSecurityID sets BenchmarkSecurityID, Tag 699
func (m CrossOrderCancelReplaceRequest) SetBenchmarkSecurityID(v string) {
	m.Set(field.NewBenchmarkSecurityID(v))
}

//SetYieldCalcDate sets YieldCalcDate, Tag 701
func (m CrossOrderCancelReplaceRequest) SetYieldCalcDate(v string) {
	m.Set(field.NewYieldCalcDate(v))
}

//SetNoUnderlyings sets NoUnderlyings, Tag 711
func (m CrossOrderCancelReplaceRequest) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

//SetBenchmarkSecurityIDSource sets BenchmarkSecurityIDSource, Tag 761
func (m CrossOrderCancelReplaceRequest) SetBenchmarkSecurityIDSource(v string) {
	m.Set(field.NewBenchmarkSecurityIDSource(v))
}

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m CrossOrderCancelReplaceRequest) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

//SetPegMoveType sets PegMoveType, Tag 835
func (m CrossOrderCancelReplaceRequest) SetPegMoveType(v enum.PegMoveType) {
	m.Set(field.NewPegMoveType(v))
}

//SetPegOffsetType sets PegOffsetType, Tag 836
func (m CrossOrderCancelReplaceRequest) SetPegOffsetType(v enum.PegOffsetType) {
	m.Set(field.NewPegOffsetType(v))
}

//SetPegLimitType sets PegLimitType, Tag 837
func (m CrossOrderCancelReplaceRequest) SetPegLimitType(v enum.PegLimitType) {
	m.Set(field.NewPegLimitType(v))
}

//SetPegRoundDirection sets PegRoundDirection, Tag 838
func (m CrossOrderCancelReplaceRequest) SetPegRoundDirection(v enum.PegRoundDirection) {
	m.Set(field.NewPegRoundDirection(v))
}

//SetPegScope sets PegScope, Tag 840
func (m CrossOrderCancelReplaceRequest) SetPegScope(v enum.PegScope) {
	m.Set(field.NewPegScope(v))
}

//SetDiscretionMoveType sets DiscretionMoveType, Tag 841
func (m CrossOrderCancelReplaceRequest) SetDiscretionMoveType(v enum.DiscretionMoveType) {
	m.Set(field.NewDiscretionMoveType(v))
}

//SetDiscretionOffsetType sets DiscretionOffsetType, Tag 842
func (m CrossOrderCancelReplaceRequest) SetDiscretionOffsetType(v enum.DiscretionOffsetType) {
	m.Set(field.NewDiscretionOffsetType(v))
}

//SetDiscretionLimitType sets DiscretionLimitType, Tag 843
func (m CrossOrderCancelReplaceRequest) SetDiscretionLimitType(v enum.DiscretionLimitType) {
	m.Set(field.NewDiscretionLimitType(v))
}

//SetDiscretionRoundDirection sets DiscretionRoundDirection, Tag 844
func (m CrossOrderCancelReplaceRequest) SetDiscretionRoundDirection(v enum.DiscretionRoundDirection) {
	m.Set(field.NewDiscretionRoundDirection(v))
}

//SetDiscretionScope sets DiscretionScope, Tag 846
func (m CrossOrderCancelReplaceRequest) SetDiscretionScope(v enum.DiscretionScope) {
	m.Set(field.NewDiscretionScope(v))
}

//SetTargetStrategy sets TargetStrategy, Tag 847
func (m CrossOrderCancelReplaceRequest) SetTargetStrategy(v enum.TargetStrategy) {
	m.Set(field.NewTargetStrategy(v))
}

//SetTargetStrategyParameters sets TargetStrategyParameters, Tag 848
func (m CrossOrderCancelReplaceRequest) SetTargetStrategyParameters(v string) {
	m.Set(field.NewTargetStrategyParameters(v))
}

//SetParticipationRate sets ParticipationRate, Tag 849
func (m CrossOrderCancelReplaceRequest) SetParticipationRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewParticipationRate(value, scale))
}

//SetNoEvents sets NoEvents, Tag 864
func (m CrossOrderCancelReplaceRequest) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDatedDate sets DatedDate, Tag 873
func (m CrossOrderCancelReplaceRequest) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

//SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m CrossOrderCancelReplaceRequest) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

//SetCPProgram sets CPProgram, Tag 875
func (m CrossOrderCancelReplaceRequest) SetCPProgram(v enum.CPProgram) {
	m.Set(field.NewCPProgram(v))
}

//SetCPRegType sets CPRegType, Tag 876
func (m CrossOrderCancelReplaceRequest) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

//SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m CrossOrderCancelReplaceRequest) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

//SetNoStrategyParameters sets NoStrategyParameters, Tag 957
func (m CrossOrderCancelReplaceRequest) SetNoStrategyParameters(f NoStrategyParametersRepeatingGroup) {
	m.SetGroup(f)
}

//SetHostCrossID sets HostCrossID, Tag 961
func (m CrossOrderCancelReplaceRequest) SetHostCrossID(v string) {
	m.Set(field.NewHostCrossID(v))
}

//SetSecurityStatus sets SecurityStatus, Tag 965
func (m CrossOrderCancelReplaceRequest) SetSecurityStatus(v enum.SecurityStatus) {
	m.Set(field.NewSecurityStatus(v))
}

//SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966
func (m CrossOrderCancelReplaceRequest) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

//SetStrikeMultiplier sets StrikeMultiplier, Tag 967
func (m CrossOrderCancelReplaceRequest) SetStrikeMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeMultiplier(value, scale))
}

//SetStrikeValue sets StrikeValue, Tag 968
func (m CrossOrderCancelReplaceRequest) SetStrikeValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeValue(value, scale))
}

//SetMinPriceIncrement sets MinPriceIncrement, Tag 969
func (m CrossOrderCancelReplaceRequest) SetMinPriceIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrement(value, scale))
}

//SetPositionLimit sets PositionLimit, Tag 970
func (m CrossOrderCancelReplaceRequest) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

//SetNTPositionLimit sets NTPositionLimit, Tag 971
func (m CrossOrderCancelReplaceRequest) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

//SetUnitOfMeasure sets UnitOfMeasure, Tag 996
func (m CrossOrderCancelReplaceRequest) SetUnitOfMeasure(v enum.UnitOfMeasure) {
	m.Set(field.NewUnitOfMeasure(v))
}

//SetTimeUnit sets TimeUnit, Tag 997
func (m CrossOrderCancelReplaceRequest) SetTimeUnit(v enum.TimeUnit) {
	m.Set(field.NewTimeUnit(v))
}

//SetNoInstrumentParties sets NoInstrumentParties, Tag 1018
func (m CrossOrderCancelReplaceRequest) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049
func (m CrossOrderCancelReplaceRequest) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

//SetMaturityTime sets MaturityTime, Tag 1079
func (m CrossOrderCancelReplaceRequest) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

//SetSecondaryDisplayQty sets SecondaryDisplayQty, Tag 1082
func (m CrossOrderCancelReplaceRequest) SetSecondaryDisplayQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewSecondaryDisplayQty(value, scale))
}

//SetDisplayWhen sets DisplayWhen, Tag 1083
func (m CrossOrderCancelReplaceRequest) SetDisplayWhen(v enum.DisplayWhen) {
	m.Set(field.NewDisplayWhen(v))
}

//SetDisplayMethod sets DisplayMethod, Tag 1084
func (m CrossOrderCancelReplaceRequest) SetDisplayMethod(v enum.DisplayMethod) {
	m.Set(field.NewDisplayMethod(v))
}

//SetDisplayLowQty sets DisplayLowQty, Tag 1085
func (m CrossOrderCancelReplaceRequest) SetDisplayLowQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewDisplayLowQty(value, scale))
}

//SetDisplayHighQty sets DisplayHighQty, Tag 1086
func (m CrossOrderCancelReplaceRequest) SetDisplayHighQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewDisplayHighQty(value, scale))
}

//SetDisplayMinIncr sets DisplayMinIncr, Tag 1087
func (m CrossOrderCancelReplaceRequest) SetDisplayMinIncr(value decimal.Decimal, scale int32) {
	m.Set(field.NewDisplayMinIncr(value, scale))
}

//SetRefreshQty sets RefreshQty, Tag 1088
func (m CrossOrderCancelReplaceRequest) SetRefreshQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewRefreshQty(value, scale))
}

//SetMatchIncrement sets MatchIncrement, Tag 1089
func (m CrossOrderCancelReplaceRequest) SetMatchIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewMatchIncrement(value, scale))
}

//SetMaxPriceLevels sets MaxPriceLevels, Tag 1090
func (m CrossOrderCancelReplaceRequest) SetMaxPriceLevels(v int) {
	m.Set(field.NewMaxPriceLevels(v))
}

//SetPriceProtectionScope sets PriceProtectionScope, Tag 1092
func (m CrossOrderCancelReplaceRequest) SetPriceProtectionScope(v enum.PriceProtectionScope) {
	m.Set(field.NewPriceProtectionScope(v))
}

//SetPegPriceType sets PegPriceType, Tag 1094
func (m CrossOrderCancelReplaceRequest) SetPegPriceType(v enum.PegPriceType) {
	m.Set(field.NewPegPriceType(v))
}

//SetPegSecurityIDSource sets PegSecurityIDSource, Tag 1096
func (m CrossOrderCancelReplaceRequest) SetPegSecurityIDSource(v string) {
	m.Set(field.NewPegSecurityIDSource(v))
}

//SetPegSecurityID sets PegSecurityID, Tag 1097
func (m CrossOrderCancelReplaceRequest) SetPegSecurityID(v string) {
	m.Set(field.NewPegSecurityID(v))
}

//SetPegSymbol sets PegSymbol, Tag 1098
func (m CrossOrderCancelReplaceRequest) SetPegSymbol(v string) {
	m.Set(field.NewPegSymbol(v))
}

//SetPegSecurityDesc sets PegSecurityDesc, Tag 1099
func (m CrossOrderCancelReplaceRequest) SetPegSecurityDesc(v string) {
	m.Set(field.NewPegSecurityDesc(v))
}

//SetTriggerType sets TriggerType, Tag 1100
func (m CrossOrderCancelReplaceRequest) SetTriggerType(v enum.TriggerType) {
	m.Set(field.NewTriggerType(v))
}

//SetTriggerAction sets TriggerAction, Tag 1101
func (m CrossOrderCancelReplaceRequest) SetTriggerAction(v enum.TriggerAction) {
	m.Set(field.NewTriggerAction(v))
}

//SetTriggerPrice sets TriggerPrice, Tag 1102
func (m CrossOrderCancelReplaceRequest) SetTriggerPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewTriggerPrice(value, scale))
}

//SetTriggerSymbol sets TriggerSymbol, Tag 1103
func (m CrossOrderCancelReplaceRequest) SetTriggerSymbol(v string) {
	m.Set(field.NewTriggerSymbol(v))
}

//SetTriggerSecurityID sets TriggerSecurityID, Tag 1104
func (m CrossOrderCancelReplaceRequest) SetTriggerSecurityID(v string) {
	m.Set(field.NewTriggerSecurityID(v))
}

//SetTriggerSecurityIDSource sets TriggerSecurityIDSource, Tag 1105
func (m CrossOrderCancelReplaceRequest) SetTriggerSecurityIDSource(v string) {
	m.Set(field.NewTriggerSecurityIDSource(v))
}

//SetTriggerSecurityDesc sets TriggerSecurityDesc, Tag 1106
func (m CrossOrderCancelReplaceRequest) SetTriggerSecurityDesc(v string) {
	m.Set(field.NewTriggerSecurityDesc(v))
}

//SetTriggerPriceType sets TriggerPriceType, Tag 1107
func (m CrossOrderCancelReplaceRequest) SetTriggerPriceType(v enum.TriggerPriceType) {
	m.Set(field.NewTriggerPriceType(v))
}

//SetTriggerPriceTypeScope sets TriggerPriceTypeScope, Tag 1108
func (m CrossOrderCancelReplaceRequest) SetTriggerPriceTypeScope(v enum.TriggerPriceTypeScope) {
	m.Set(field.NewTriggerPriceTypeScope(v))
}

//SetTriggerPriceDirection sets TriggerPriceDirection, Tag 1109
func (m CrossOrderCancelReplaceRequest) SetTriggerPriceDirection(v enum.TriggerPriceDirection) {
	m.Set(field.NewTriggerPriceDirection(v))
}

//SetTriggerNewPrice sets TriggerNewPrice, Tag 1110
func (m CrossOrderCancelReplaceRequest) SetTriggerNewPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewTriggerNewPrice(value, scale))
}

//SetTriggerOrderType sets TriggerOrderType, Tag 1111
func (m CrossOrderCancelReplaceRequest) SetTriggerOrderType(v enum.TriggerOrderType) {
	m.Set(field.NewTriggerOrderType(v))
}

//SetTriggerNewQty sets TriggerNewQty, Tag 1112
func (m CrossOrderCancelReplaceRequest) SetTriggerNewQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewTriggerNewQty(value, scale))
}

//SetTriggerTradingSessionID sets TriggerTradingSessionID, Tag 1113
func (m CrossOrderCancelReplaceRequest) SetTriggerTradingSessionID(v string) {
	m.Set(field.NewTriggerTradingSessionID(v))
}

//SetTriggerTradingSessionSubID sets TriggerTradingSessionSubID, Tag 1114
func (m CrossOrderCancelReplaceRequest) SetTriggerTradingSessionSubID(v string) {
	m.Set(field.NewTriggerTradingSessionSubID(v))
}

//SetNoRootPartyIDs sets NoRootPartyIDs, Tag 1116
func (m CrossOrderCancelReplaceRequest) SetNoRootPartyIDs(f NoRootPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetExDestinationIDSource sets ExDestinationIDSource, Tag 1133
func (m CrossOrderCancelReplaceRequest) SetExDestinationIDSource(v enum.ExDestinationIDSource) {
	m.Set(field.NewExDestinationIDSource(v))
}

//SetDisplayQty sets DisplayQty, Tag 1138
func (m CrossOrderCancelReplaceRequest) SetDisplayQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewDisplayQty(value, scale))
}

//GetCurrency gets Currency, Tag 15
func (m CrossOrderCancelReplaceRequest) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecInst gets ExecInst, Tag 18
func (m CrossOrderCancelReplaceRequest) GetExecInst() (v enum.ExecInst, err quickfix.MessageRejectError) {
	var f field.ExecInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetHandlInst gets HandlInst, Tag 21
func (m CrossOrderCancelReplaceRequest) GetHandlInst() (v enum.HandlInst, err quickfix.MessageRejectError) {
	var f field.HandlInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m CrossOrderCancelReplaceRequest) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIOIID gets IOIID, Tag 23
func (m CrossOrderCancelReplaceRequest) GetIOIID() (v string, err quickfix.MessageRejectError) {
	var f field.IOIIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderID gets OrderID, Tag 37
func (m CrossOrderCancelReplaceRequest) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrdType gets OrdType, Tag 40
func (m CrossOrderCancelReplaceRequest) GetOrdType() (v enum.OrdType, err quickfix.MessageRejectError) {
	var f field.OrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPrice gets Price, Tag 44
func (m CrossOrderCancelReplaceRequest) GetPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m CrossOrderCancelReplaceRequest) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbol gets Symbol, Tag 55
func (m CrossOrderCancelReplaceRequest) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m CrossOrderCancelReplaceRequest) GetTimeInForce() (v enum.TimeInForce, err quickfix.MessageRejectError) {
	var f field.TimeInForceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m CrossOrderCancelReplaceRequest) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlType gets SettlType, Tag 63
func (m CrossOrderCancelReplaceRequest) GetSettlType() (v enum.SettlType, err quickfix.MessageRejectError) {
	var f field.SettlTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlDate gets SettlDate, Tag 64
func (m CrossOrderCancelReplaceRequest) GetSettlDate() (v string, err quickfix.MessageRejectError) {
	var f field.SettlDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m CrossOrderCancelReplaceRequest) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetProcessCode gets ProcessCode, Tag 81
func (m CrossOrderCancelReplaceRequest) GetProcessCode() (v enum.ProcessCode, err quickfix.MessageRejectError) {
	var f field.ProcessCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStopPx gets StopPx, Tag 99
func (m CrossOrderCancelReplaceRequest) GetStopPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StopPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExDestination gets ExDestination, Tag 100
func (m CrossOrderCancelReplaceRequest) GetExDestination() (v enum.ExDestination, err quickfix.MessageRejectError) {
	var f field.ExDestinationField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssuer gets Issuer, Tag 106
func (m CrossOrderCancelReplaceRequest) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m CrossOrderCancelReplaceRequest) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMinQty gets MinQty, Tag 110
func (m CrossOrderCancelReplaceRequest) GetMinQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaxFloor gets MaxFloor, Tag 111
func (m CrossOrderCancelReplaceRequest) GetMaxFloor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MaxFloorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLocateReqd gets LocateReqd, Tag 114
func (m CrossOrderCancelReplaceRequest) GetLocateReqd() (v bool, err quickfix.MessageRejectError) {
	var f field.LocateReqdField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetQuoteID gets QuoteID, Tag 117
func (m CrossOrderCancelReplaceRequest) GetQuoteID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m CrossOrderCancelReplaceRequest) GetExpireTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ExpireTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPrevClosePx gets PrevClosePx, Tag 140
func (m CrossOrderCancelReplaceRequest) GetPrevClosePx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PrevClosePxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m CrossOrderCancelReplaceRequest) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEffectiveTime gets EffectiveTime, Tag 168
func (m CrossOrderCancelReplaceRequest) GetEffectiveTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.EffectiveTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m CrossOrderCancelReplaceRequest) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m CrossOrderCancelReplaceRequest) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m CrossOrderCancelReplaceRequest) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m CrossOrderCancelReplaceRequest) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaxShow gets MaxShow, Tag 210
func (m CrossOrderCancelReplaceRequest) GetMaxShow() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MaxShowField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegOffsetValue gets PegOffsetValue, Tag 211
func (m CrossOrderCancelReplaceRequest) GetPegOffsetValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PegOffsetValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSpread gets Spread, Tag 218
func (m CrossOrderCancelReplaceRequest) GetSpread() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SpreadField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBenchmarkCurveCurrency gets BenchmarkCurveCurrency, Tag 220
func (m CrossOrderCancelReplaceRequest) GetBenchmarkCurveCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkCurveCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBenchmarkCurveName gets BenchmarkCurveName, Tag 221
func (m CrossOrderCancelReplaceRequest) GetBenchmarkCurveName() (v enum.BenchmarkCurveName, err quickfix.MessageRejectError) {
	var f field.BenchmarkCurveNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBenchmarkCurvePoint gets BenchmarkCurvePoint, Tag 222
func (m CrossOrderCancelReplaceRequest) GetBenchmarkCurvePoint() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkCurvePointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m CrossOrderCancelReplaceRequest) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m CrossOrderCancelReplaceRequest) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m CrossOrderCancelReplaceRequest) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m CrossOrderCancelReplaceRequest) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m CrossOrderCancelReplaceRequest) GetRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFactor gets Factor, Tag 228
func (m CrossOrderCancelReplaceRequest) GetFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m CrossOrderCancelReplaceRequest) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoStipulations gets NoStipulations, Tag 232
func (m CrossOrderCancelReplaceRequest) GetNoStipulations() (NoStipulationsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStipulationsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetYieldType gets YieldType, Tag 235
func (m CrossOrderCancelReplaceRequest) GetYieldType() (v enum.YieldType, err quickfix.MessageRejectError) {
	var f field.YieldTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetYield gets Yield, Tag 236
func (m CrossOrderCancelReplaceRequest) GetYield() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.YieldField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m CrossOrderCancelReplaceRequest) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m CrossOrderCancelReplaceRequest) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m CrossOrderCancelReplaceRequest) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m CrossOrderCancelReplaceRequest) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m CrossOrderCancelReplaceRequest) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m CrossOrderCancelReplaceRequest) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m CrossOrderCancelReplaceRequest) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetComplianceID gets ComplianceID, Tag 376
func (m CrossOrderCancelReplaceRequest) GetComplianceID() (v string, err quickfix.MessageRejectError) {
	var f field.ComplianceIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoTradingSessions gets NoTradingSessions, Tag 386
func (m CrossOrderCancelReplaceRequest) GetNoTradingSessions() (NoTradingSessionsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTradingSessionsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDiscretionInst gets DiscretionInst, Tag 388
func (m CrossOrderCancelReplaceRequest) GetDiscretionInst() (v enum.DiscretionInst, err quickfix.MessageRejectError) {
	var f field.DiscretionInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDiscretionOffsetValue gets DiscretionOffsetValue, Tag 389
func (m CrossOrderCancelReplaceRequest) GetDiscretionOffsetValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DiscretionOffsetValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPriceType gets PriceType, Tag 423
func (m CrossOrderCancelReplaceRequest) GetPriceType() (v enum.PriceType, err quickfix.MessageRejectError) {
	var f field.PriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetGTBookingInst gets GTBookingInst, Tag 427
func (m CrossOrderCancelReplaceRequest) GetGTBookingInst() (v enum.GTBookingInst, err quickfix.MessageRejectError) {
	var f field.GTBookingInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExpireDate gets ExpireDate, Tag 432
func (m CrossOrderCancelReplaceRequest) GetExpireDate() (v string, err quickfix.MessageRejectError) {
	var f field.ExpireDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m CrossOrderCancelReplaceRequest) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m CrossOrderCancelReplaceRequest) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCFICode gets CFICode, Tag 461
func (m CrossOrderCancelReplaceRequest) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m CrossOrderCancelReplaceRequest) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m CrossOrderCancelReplaceRequest) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m CrossOrderCancelReplaceRequest) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCancellationRights gets CancellationRights, Tag 480
func (m CrossOrderCancelReplaceRequest) GetCancellationRights() (v enum.CancellationRights, err quickfix.MessageRejectError) {
	var f field.CancellationRightsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMoneyLaunderingStatus gets MoneyLaunderingStatus, Tag 481
func (m CrossOrderCancelReplaceRequest) GetMoneyLaunderingStatus() (v enum.MoneyLaunderingStatus, err quickfix.MessageRejectError) {
	var f field.MoneyLaunderingStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransBkdTime gets TransBkdTime, Tag 483
func (m CrossOrderCancelReplaceRequest) GetTransBkdTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransBkdTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDesignation gets Designation, Tag 494
func (m CrossOrderCancelReplaceRequest) GetDesignation() (v string, err quickfix.MessageRejectError) {
	var f field.DesignationField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRegistID gets RegistID, Tag 513
func (m CrossOrderCancelReplaceRequest) GetRegistID() (v string, err quickfix.MessageRejectError) {
	var f field.RegistIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m CrossOrderCancelReplaceRequest) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m CrossOrderCancelReplaceRequest) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCrossID gets CrossID, Tag 548
func (m CrossOrderCancelReplaceRequest) GetCrossID() (v string, err quickfix.MessageRejectError) {
	var f field.CrossIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCrossType gets CrossType, Tag 549
func (m CrossOrderCancelReplaceRequest) GetCrossType() (v enum.CrossType, err quickfix.MessageRejectError) {
	var f field.CrossTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCrossPrioritization gets CrossPrioritization, Tag 550
func (m CrossOrderCancelReplaceRequest) GetCrossPrioritization() (v enum.CrossPrioritization, err quickfix.MessageRejectError) {
	var f field.CrossPrioritizationField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrigCrossID gets OrigCrossID, Tag 551
func (m CrossOrderCancelReplaceRequest) GetOrigCrossID() (v string, err quickfix.MessageRejectError) {
	var f field.OrigCrossIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoSides gets NoSides, Tag 552
func (m CrossOrderCancelReplaceRequest) GetNoSides() (NoSidesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSidesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoLegs gets NoLegs, Tag 555
func (m CrossOrderCancelReplaceRequest) GetNoLegs() (NoLegsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetBenchmarkPrice gets BenchmarkPrice, Tag 662
func (m CrossOrderCancelReplaceRequest) GetBenchmarkPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BenchmarkPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBenchmarkPriceType gets BenchmarkPriceType, Tag 663
func (m CrossOrderCancelReplaceRequest) GetBenchmarkPriceType() (v int, err quickfix.MessageRejectError) {
	var f field.BenchmarkPriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m CrossOrderCancelReplaceRequest) GetContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.ContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPool gets Pool, Tag 691
func (m CrossOrderCancelReplaceRequest) GetPool() (v string, err quickfix.MessageRejectError) {
	var f field.PoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetYieldRedemptionDate gets YieldRedemptionDate, Tag 696
func (m CrossOrderCancelReplaceRequest) GetYieldRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.YieldRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetYieldRedemptionPrice gets YieldRedemptionPrice, Tag 697
func (m CrossOrderCancelReplaceRequest) GetYieldRedemptionPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.YieldRedemptionPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetYieldRedemptionPriceType gets YieldRedemptionPriceType, Tag 698
func (m CrossOrderCancelReplaceRequest) GetYieldRedemptionPriceType() (v int, err quickfix.MessageRejectError) {
	var f field.YieldRedemptionPriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBenchmarkSecurityID gets BenchmarkSecurityID, Tag 699
func (m CrossOrderCancelReplaceRequest) GetBenchmarkSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetYieldCalcDate gets YieldCalcDate, Tag 701
func (m CrossOrderCancelReplaceRequest) GetYieldCalcDate() (v string, err quickfix.MessageRejectError) {
	var f field.YieldCalcDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoUnderlyings gets NoUnderlyings, Tag 711
func (m CrossOrderCancelReplaceRequest) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetBenchmarkSecurityIDSource gets BenchmarkSecurityIDSource, Tag 761
func (m CrossOrderCancelReplaceRequest) GetBenchmarkSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m CrossOrderCancelReplaceRequest) GetSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegMoveType gets PegMoveType, Tag 835
func (m CrossOrderCancelReplaceRequest) GetPegMoveType() (v enum.PegMoveType, err quickfix.MessageRejectError) {
	var f field.PegMoveTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegOffsetType gets PegOffsetType, Tag 836
func (m CrossOrderCancelReplaceRequest) GetPegOffsetType() (v enum.PegOffsetType, err quickfix.MessageRejectError) {
	var f field.PegOffsetTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegLimitType gets PegLimitType, Tag 837
func (m CrossOrderCancelReplaceRequest) GetPegLimitType() (v enum.PegLimitType, err quickfix.MessageRejectError) {
	var f field.PegLimitTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegRoundDirection gets PegRoundDirection, Tag 838
func (m CrossOrderCancelReplaceRequest) GetPegRoundDirection() (v enum.PegRoundDirection, err quickfix.MessageRejectError) {
	var f field.PegRoundDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegScope gets PegScope, Tag 840
func (m CrossOrderCancelReplaceRequest) GetPegScope() (v enum.PegScope, err quickfix.MessageRejectError) {
	var f field.PegScopeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDiscretionMoveType gets DiscretionMoveType, Tag 841
func (m CrossOrderCancelReplaceRequest) GetDiscretionMoveType() (v enum.DiscretionMoveType, err quickfix.MessageRejectError) {
	var f field.DiscretionMoveTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDiscretionOffsetType gets DiscretionOffsetType, Tag 842
func (m CrossOrderCancelReplaceRequest) GetDiscretionOffsetType() (v enum.DiscretionOffsetType, err quickfix.MessageRejectError) {
	var f field.DiscretionOffsetTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDiscretionLimitType gets DiscretionLimitType, Tag 843
func (m CrossOrderCancelReplaceRequest) GetDiscretionLimitType() (v enum.DiscretionLimitType, err quickfix.MessageRejectError) {
	var f field.DiscretionLimitTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDiscretionRoundDirection gets DiscretionRoundDirection, Tag 844
func (m CrossOrderCancelReplaceRequest) GetDiscretionRoundDirection() (v enum.DiscretionRoundDirection, err quickfix.MessageRejectError) {
	var f field.DiscretionRoundDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDiscretionScope gets DiscretionScope, Tag 846
func (m CrossOrderCancelReplaceRequest) GetDiscretionScope() (v enum.DiscretionScope, err quickfix.MessageRejectError) {
	var f field.DiscretionScopeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTargetStrategy gets TargetStrategy, Tag 847
func (m CrossOrderCancelReplaceRequest) GetTargetStrategy() (v enum.TargetStrategy, err quickfix.MessageRejectError) {
	var f field.TargetStrategyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTargetStrategyParameters gets TargetStrategyParameters, Tag 848
func (m CrossOrderCancelReplaceRequest) GetTargetStrategyParameters() (v string, err quickfix.MessageRejectError) {
	var f field.TargetStrategyParametersField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetParticipationRate gets ParticipationRate, Tag 849
func (m CrossOrderCancelReplaceRequest) GetParticipationRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ParticipationRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoEvents gets NoEvents, Tag 864
func (m CrossOrderCancelReplaceRequest) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDatedDate gets DatedDate, Tag 873
func (m CrossOrderCancelReplaceRequest) GetDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.DatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m CrossOrderCancelReplaceRequest) GetInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.InterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCPProgram gets CPProgram, Tag 875
func (m CrossOrderCancelReplaceRequest) GetCPProgram() (v enum.CPProgram, err quickfix.MessageRejectError) {
	var f field.CPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCPRegType gets CPRegType, Tag 876
func (m CrossOrderCancelReplaceRequest) GetCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.CPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m CrossOrderCancelReplaceRequest) GetStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.StrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoStrategyParameters gets NoStrategyParameters, Tag 957
func (m CrossOrderCancelReplaceRequest) GetNoStrategyParameters() (NoStrategyParametersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStrategyParametersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetHostCrossID gets HostCrossID, Tag 961
func (m CrossOrderCancelReplaceRequest) GetHostCrossID() (v string, err quickfix.MessageRejectError) {
	var f field.HostCrossIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityStatus gets SecurityStatus, Tag 965
func (m CrossOrderCancelReplaceRequest) GetSecurityStatus() (v enum.SecurityStatus, err quickfix.MessageRejectError) {
	var f field.SecurityStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966
func (m CrossOrderCancelReplaceRequest) GetSettleOnOpenFlag() (v string, err quickfix.MessageRejectError) {
	var f field.SettleOnOpenFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikeMultiplier gets StrikeMultiplier, Tag 967
func (m CrossOrderCancelReplaceRequest) GetStrikeMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikeMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikeValue gets StrikeValue, Tag 968
func (m CrossOrderCancelReplaceRequest) GetStrikeValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikeValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMinPriceIncrement gets MinPriceIncrement, Tag 969
func (m CrossOrderCancelReplaceRequest) GetMinPriceIncrement() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinPriceIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPositionLimit gets PositionLimit, Tag 970
func (m CrossOrderCancelReplaceRequest) GetPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.PositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNTPositionLimit gets NTPositionLimit, Tag 971
func (m CrossOrderCancelReplaceRequest) GetNTPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.NTPositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnitOfMeasure gets UnitOfMeasure, Tag 996
func (m CrossOrderCancelReplaceRequest) GetUnitOfMeasure() (v enum.UnitOfMeasure, err quickfix.MessageRejectError) {
	var f field.UnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTimeUnit gets TimeUnit, Tag 997
func (m CrossOrderCancelReplaceRequest) GetTimeUnit() (v enum.TimeUnit, err quickfix.MessageRejectError) {
	var f field.TimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoInstrumentParties gets NoInstrumentParties, Tag 1018
func (m CrossOrderCancelReplaceRequest) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049
func (m CrossOrderCancelReplaceRequest) GetInstrmtAssignmentMethod() (v string, err quickfix.MessageRejectError) {
	var f field.InstrmtAssignmentMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityTime gets MaturityTime, Tag 1079
func (m CrossOrderCancelReplaceRequest) GetMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecondaryDisplayQty gets SecondaryDisplayQty, Tag 1082
func (m CrossOrderCancelReplaceRequest) GetSecondaryDisplayQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SecondaryDisplayQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDisplayWhen gets DisplayWhen, Tag 1083
func (m CrossOrderCancelReplaceRequest) GetDisplayWhen() (v enum.DisplayWhen, err quickfix.MessageRejectError) {
	var f field.DisplayWhenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDisplayMethod gets DisplayMethod, Tag 1084
func (m CrossOrderCancelReplaceRequest) GetDisplayMethod() (v enum.DisplayMethod, err quickfix.MessageRejectError) {
	var f field.DisplayMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDisplayLowQty gets DisplayLowQty, Tag 1085
func (m CrossOrderCancelReplaceRequest) GetDisplayLowQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DisplayLowQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDisplayHighQty gets DisplayHighQty, Tag 1086
func (m CrossOrderCancelReplaceRequest) GetDisplayHighQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DisplayHighQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDisplayMinIncr gets DisplayMinIncr, Tag 1087
func (m CrossOrderCancelReplaceRequest) GetDisplayMinIncr() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DisplayMinIncrField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRefreshQty gets RefreshQty, Tag 1088
func (m CrossOrderCancelReplaceRequest) GetRefreshQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RefreshQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMatchIncrement gets MatchIncrement, Tag 1089
func (m CrossOrderCancelReplaceRequest) GetMatchIncrement() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MatchIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaxPriceLevels gets MaxPriceLevels, Tag 1090
func (m CrossOrderCancelReplaceRequest) GetMaxPriceLevels() (v int, err quickfix.MessageRejectError) {
	var f field.MaxPriceLevelsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPriceProtectionScope gets PriceProtectionScope, Tag 1092
func (m CrossOrderCancelReplaceRequest) GetPriceProtectionScope() (v enum.PriceProtectionScope, err quickfix.MessageRejectError) {
	var f field.PriceProtectionScopeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegPriceType gets PegPriceType, Tag 1094
func (m CrossOrderCancelReplaceRequest) GetPegPriceType() (v enum.PegPriceType, err quickfix.MessageRejectError) {
	var f field.PegPriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegSecurityIDSource gets PegSecurityIDSource, Tag 1096
func (m CrossOrderCancelReplaceRequest) GetPegSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.PegSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegSecurityID gets PegSecurityID, Tag 1097
func (m CrossOrderCancelReplaceRequest) GetPegSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.PegSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegSymbol gets PegSymbol, Tag 1098
func (m CrossOrderCancelReplaceRequest) GetPegSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.PegSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegSecurityDesc gets PegSecurityDesc, Tag 1099
func (m CrossOrderCancelReplaceRequest) GetPegSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.PegSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerType gets TriggerType, Tag 1100
func (m CrossOrderCancelReplaceRequest) GetTriggerType() (v enum.TriggerType, err quickfix.MessageRejectError) {
	var f field.TriggerTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerAction gets TriggerAction, Tag 1101
func (m CrossOrderCancelReplaceRequest) GetTriggerAction() (v enum.TriggerAction, err quickfix.MessageRejectError) {
	var f field.TriggerActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerPrice gets TriggerPrice, Tag 1102
func (m CrossOrderCancelReplaceRequest) GetTriggerPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TriggerPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerSymbol gets TriggerSymbol, Tag 1103
func (m CrossOrderCancelReplaceRequest) GetTriggerSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.TriggerSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerSecurityID gets TriggerSecurityID, Tag 1104
func (m CrossOrderCancelReplaceRequest) GetTriggerSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.TriggerSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerSecurityIDSource gets TriggerSecurityIDSource, Tag 1105
func (m CrossOrderCancelReplaceRequest) GetTriggerSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.TriggerSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerSecurityDesc gets TriggerSecurityDesc, Tag 1106
func (m CrossOrderCancelReplaceRequest) GetTriggerSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.TriggerSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerPriceType gets TriggerPriceType, Tag 1107
func (m CrossOrderCancelReplaceRequest) GetTriggerPriceType() (v enum.TriggerPriceType, err quickfix.MessageRejectError) {
	var f field.TriggerPriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerPriceTypeScope gets TriggerPriceTypeScope, Tag 1108
func (m CrossOrderCancelReplaceRequest) GetTriggerPriceTypeScope() (v enum.TriggerPriceTypeScope, err quickfix.MessageRejectError) {
	var f field.TriggerPriceTypeScopeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerPriceDirection gets TriggerPriceDirection, Tag 1109
func (m CrossOrderCancelReplaceRequest) GetTriggerPriceDirection() (v enum.TriggerPriceDirection, err quickfix.MessageRejectError) {
	var f field.TriggerPriceDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerNewPrice gets TriggerNewPrice, Tag 1110
func (m CrossOrderCancelReplaceRequest) GetTriggerNewPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TriggerNewPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerOrderType gets TriggerOrderType, Tag 1111
func (m CrossOrderCancelReplaceRequest) GetTriggerOrderType() (v enum.TriggerOrderType, err quickfix.MessageRejectError) {
	var f field.TriggerOrderTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerNewQty gets TriggerNewQty, Tag 1112
func (m CrossOrderCancelReplaceRequest) GetTriggerNewQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TriggerNewQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerTradingSessionID gets TriggerTradingSessionID, Tag 1113
func (m CrossOrderCancelReplaceRequest) GetTriggerTradingSessionID() (v string, err quickfix.MessageRejectError) {
	var f field.TriggerTradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerTradingSessionSubID gets TriggerTradingSessionSubID, Tag 1114
func (m CrossOrderCancelReplaceRequest) GetTriggerTradingSessionSubID() (v string, err quickfix.MessageRejectError) {
	var f field.TriggerTradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoRootPartyIDs gets NoRootPartyIDs, Tag 1116
func (m CrossOrderCancelReplaceRequest) GetNoRootPartyIDs() (NoRootPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRootPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetExDestinationIDSource gets ExDestinationIDSource, Tag 1133
func (m CrossOrderCancelReplaceRequest) GetExDestinationIDSource() (v enum.ExDestinationIDSource, err quickfix.MessageRejectError) {
	var f field.ExDestinationIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDisplayQty gets DisplayQty, Tag 1138
func (m CrossOrderCancelReplaceRequest) GetDisplayQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DisplayQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasCurrency returns true if Currency is present, Tag 15
func (m CrossOrderCancelReplaceRequest) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasExecInst returns true if ExecInst is present, Tag 18
func (m CrossOrderCancelReplaceRequest) HasExecInst() bool {
	return m.Has(tag.ExecInst)
}

//HasHandlInst returns true if HandlInst is present, Tag 21
func (m CrossOrderCancelReplaceRequest) HasHandlInst() bool {
	return m.Has(tag.HandlInst)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m CrossOrderCancelReplaceRequest) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasIOIID returns true if IOIID is present, Tag 23
func (m CrossOrderCancelReplaceRequest) HasIOIID() bool {
	return m.Has(tag.IOIID)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m CrossOrderCancelReplaceRequest) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m CrossOrderCancelReplaceRequest) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasPrice returns true if Price is present, Tag 44
func (m CrossOrderCancelReplaceRequest) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m CrossOrderCancelReplaceRequest) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m CrossOrderCancelReplaceRequest) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasTimeInForce returns true if TimeInForce is present, Tag 59
func (m CrossOrderCancelReplaceRequest) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m CrossOrderCancelReplaceRequest) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSettlType returns true if SettlType is present, Tag 63
func (m CrossOrderCancelReplaceRequest) HasSettlType() bool {
	return m.Has(tag.SettlType)
}

//HasSettlDate returns true if SettlDate is present, Tag 64
func (m CrossOrderCancelReplaceRequest) HasSettlDate() bool {
	return m.Has(tag.SettlDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m CrossOrderCancelReplaceRequest) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasProcessCode returns true if ProcessCode is present, Tag 81
func (m CrossOrderCancelReplaceRequest) HasProcessCode() bool {
	return m.Has(tag.ProcessCode)
}

//HasStopPx returns true if StopPx is present, Tag 99
func (m CrossOrderCancelReplaceRequest) HasStopPx() bool {
	return m.Has(tag.StopPx)
}

//HasExDestination returns true if ExDestination is present, Tag 100
func (m CrossOrderCancelReplaceRequest) HasExDestination() bool {
	return m.Has(tag.ExDestination)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m CrossOrderCancelReplaceRequest) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m CrossOrderCancelReplaceRequest) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasMinQty returns true if MinQty is present, Tag 110
func (m CrossOrderCancelReplaceRequest) HasMinQty() bool {
	return m.Has(tag.MinQty)
}

//HasMaxFloor returns true if MaxFloor is present, Tag 111
func (m CrossOrderCancelReplaceRequest) HasMaxFloor() bool {
	return m.Has(tag.MaxFloor)
}

//HasLocateReqd returns true if LocateReqd is present, Tag 114
func (m CrossOrderCancelReplaceRequest) HasLocateReqd() bool {
	return m.Has(tag.LocateReqd)
}

//HasQuoteID returns true if QuoteID is present, Tag 117
func (m CrossOrderCancelReplaceRequest) HasQuoteID() bool {
	return m.Has(tag.QuoteID)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m CrossOrderCancelReplaceRequest) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasPrevClosePx returns true if PrevClosePx is present, Tag 140
func (m CrossOrderCancelReplaceRequest) HasPrevClosePx() bool {
	return m.Has(tag.PrevClosePx)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m CrossOrderCancelReplaceRequest) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasEffectiveTime returns true if EffectiveTime is present, Tag 168
func (m CrossOrderCancelReplaceRequest) HasEffectiveTime() bool {
	return m.Has(tag.EffectiveTime)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m CrossOrderCancelReplaceRequest) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m CrossOrderCancelReplaceRequest) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m CrossOrderCancelReplaceRequest) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m CrossOrderCancelReplaceRequest) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasMaxShow returns true if MaxShow is present, Tag 210
func (m CrossOrderCancelReplaceRequest) HasMaxShow() bool {
	return m.Has(tag.MaxShow)
}

//HasPegOffsetValue returns true if PegOffsetValue is present, Tag 211
func (m CrossOrderCancelReplaceRequest) HasPegOffsetValue() bool {
	return m.Has(tag.PegOffsetValue)
}

//HasSpread returns true if Spread is present, Tag 218
func (m CrossOrderCancelReplaceRequest) HasSpread() bool {
	return m.Has(tag.Spread)
}

//HasBenchmarkCurveCurrency returns true if BenchmarkCurveCurrency is present, Tag 220
func (m CrossOrderCancelReplaceRequest) HasBenchmarkCurveCurrency() bool {
	return m.Has(tag.BenchmarkCurveCurrency)
}

//HasBenchmarkCurveName returns true if BenchmarkCurveName is present, Tag 221
func (m CrossOrderCancelReplaceRequest) HasBenchmarkCurveName() bool {
	return m.Has(tag.BenchmarkCurveName)
}

//HasBenchmarkCurvePoint returns true if BenchmarkCurvePoint is present, Tag 222
func (m CrossOrderCancelReplaceRequest) HasBenchmarkCurvePoint() bool {
	return m.Has(tag.BenchmarkCurvePoint)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m CrossOrderCancelReplaceRequest) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m CrossOrderCancelReplaceRequest) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m CrossOrderCancelReplaceRequest) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m CrossOrderCancelReplaceRequest) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m CrossOrderCancelReplaceRequest) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m CrossOrderCancelReplaceRequest) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m CrossOrderCancelReplaceRequest) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasNoStipulations returns true if NoStipulations is present, Tag 232
func (m CrossOrderCancelReplaceRequest) HasNoStipulations() bool {
	return m.Has(tag.NoStipulations)
}

//HasYieldType returns true if YieldType is present, Tag 235
func (m CrossOrderCancelReplaceRequest) HasYieldType() bool {
	return m.Has(tag.YieldType)
}

//HasYield returns true if Yield is present, Tag 236
func (m CrossOrderCancelReplaceRequest) HasYield() bool {
	return m.Has(tag.Yield)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m CrossOrderCancelReplaceRequest) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m CrossOrderCancelReplaceRequest) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m CrossOrderCancelReplaceRequest) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m CrossOrderCancelReplaceRequest) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m CrossOrderCancelReplaceRequest) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m CrossOrderCancelReplaceRequest) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m CrossOrderCancelReplaceRequest) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasComplianceID returns true if ComplianceID is present, Tag 376
func (m CrossOrderCancelReplaceRequest) HasComplianceID() bool {
	return m.Has(tag.ComplianceID)
}

//HasNoTradingSessions returns true if NoTradingSessions is present, Tag 386
func (m CrossOrderCancelReplaceRequest) HasNoTradingSessions() bool {
	return m.Has(tag.NoTradingSessions)
}

//HasDiscretionInst returns true if DiscretionInst is present, Tag 388
func (m CrossOrderCancelReplaceRequest) HasDiscretionInst() bool {
	return m.Has(tag.DiscretionInst)
}

//HasDiscretionOffsetValue returns true if DiscretionOffsetValue is present, Tag 389
func (m CrossOrderCancelReplaceRequest) HasDiscretionOffsetValue() bool {
	return m.Has(tag.DiscretionOffsetValue)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m CrossOrderCancelReplaceRequest) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasGTBookingInst returns true if GTBookingInst is present, Tag 427
func (m CrossOrderCancelReplaceRequest) HasGTBookingInst() bool {
	return m.Has(tag.GTBookingInst)
}

//HasExpireDate returns true if ExpireDate is present, Tag 432
func (m CrossOrderCancelReplaceRequest) HasExpireDate() bool {
	return m.Has(tag.ExpireDate)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m CrossOrderCancelReplaceRequest) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m CrossOrderCancelReplaceRequest) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m CrossOrderCancelReplaceRequest) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m CrossOrderCancelReplaceRequest) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m CrossOrderCancelReplaceRequest) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m CrossOrderCancelReplaceRequest) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasCancellationRights returns true if CancellationRights is present, Tag 480
func (m CrossOrderCancelReplaceRequest) HasCancellationRights() bool {
	return m.Has(tag.CancellationRights)
}

//HasMoneyLaunderingStatus returns true if MoneyLaunderingStatus is present, Tag 481
func (m CrossOrderCancelReplaceRequest) HasMoneyLaunderingStatus() bool {
	return m.Has(tag.MoneyLaunderingStatus)
}

//HasTransBkdTime returns true if TransBkdTime is present, Tag 483
func (m CrossOrderCancelReplaceRequest) HasTransBkdTime() bool {
	return m.Has(tag.TransBkdTime)
}

//HasDesignation returns true if Designation is present, Tag 494
func (m CrossOrderCancelReplaceRequest) HasDesignation() bool {
	return m.Has(tag.Designation)
}

//HasRegistID returns true if RegistID is present, Tag 513
func (m CrossOrderCancelReplaceRequest) HasRegistID() bool {
	return m.Has(tag.RegistID)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m CrossOrderCancelReplaceRequest) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m CrossOrderCancelReplaceRequest) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasCrossID returns true if CrossID is present, Tag 548
func (m CrossOrderCancelReplaceRequest) HasCrossID() bool {
	return m.Has(tag.CrossID)
}

//HasCrossType returns true if CrossType is present, Tag 549
func (m CrossOrderCancelReplaceRequest) HasCrossType() bool {
	return m.Has(tag.CrossType)
}

//HasCrossPrioritization returns true if CrossPrioritization is present, Tag 550
func (m CrossOrderCancelReplaceRequest) HasCrossPrioritization() bool {
	return m.Has(tag.CrossPrioritization)
}

//HasOrigCrossID returns true if OrigCrossID is present, Tag 551
func (m CrossOrderCancelReplaceRequest) HasOrigCrossID() bool {
	return m.Has(tag.OrigCrossID)
}

//HasNoSides returns true if NoSides is present, Tag 552
func (m CrossOrderCancelReplaceRequest) HasNoSides() bool {
	return m.Has(tag.NoSides)
}

//HasNoLegs returns true if NoLegs is present, Tag 555
func (m CrossOrderCancelReplaceRequest) HasNoLegs() bool {
	return m.Has(tag.NoLegs)
}

//HasBenchmarkPrice returns true if BenchmarkPrice is present, Tag 662
func (m CrossOrderCancelReplaceRequest) HasBenchmarkPrice() bool {
	return m.Has(tag.BenchmarkPrice)
}

//HasBenchmarkPriceType returns true if BenchmarkPriceType is present, Tag 663
func (m CrossOrderCancelReplaceRequest) HasBenchmarkPriceType() bool {
	return m.Has(tag.BenchmarkPriceType)
}

//HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m CrossOrderCancelReplaceRequest) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

//HasPool returns true if Pool is present, Tag 691
func (m CrossOrderCancelReplaceRequest) HasPool() bool {
	return m.Has(tag.Pool)
}

//HasYieldRedemptionDate returns true if YieldRedemptionDate is present, Tag 696
func (m CrossOrderCancelReplaceRequest) HasYieldRedemptionDate() bool {
	return m.Has(tag.YieldRedemptionDate)
}

//HasYieldRedemptionPrice returns true if YieldRedemptionPrice is present, Tag 697
func (m CrossOrderCancelReplaceRequest) HasYieldRedemptionPrice() bool {
	return m.Has(tag.YieldRedemptionPrice)
}

//HasYieldRedemptionPriceType returns true if YieldRedemptionPriceType is present, Tag 698
func (m CrossOrderCancelReplaceRequest) HasYieldRedemptionPriceType() bool {
	return m.Has(tag.YieldRedemptionPriceType)
}

//HasBenchmarkSecurityID returns true if BenchmarkSecurityID is present, Tag 699
func (m CrossOrderCancelReplaceRequest) HasBenchmarkSecurityID() bool {
	return m.Has(tag.BenchmarkSecurityID)
}

//HasYieldCalcDate returns true if YieldCalcDate is present, Tag 701
func (m CrossOrderCancelReplaceRequest) HasYieldCalcDate() bool {
	return m.Has(tag.YieldCalcDate)
}

//HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711
func (m CrossOrderCancelReplaceRequest) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

//HasBenchmarkSecurityIDSource returns true if BenchmarkSecurityIDSource is present, Tag 761
func (m CrossOrderCancelReplaceRequest) HasBenchmarkSecurityIDSource() bool {
	return m.Has(tag.BenchmarkSecurityIDSource)
}

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m CrossOrderCancelReplaceRequest) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

//HasPegMoveType returns true if PegMoveType is present, Tag 835
func (m CrossOrderCancelReplaceRequest) HasPegMoveType() bool {
	return m.Has(tag.PegMoveType)
}

//HasPegOffsetType returns true if PegOffsetType is present, Tag 836
func (m CrossOrderCancelReplaceRequest) HasPegOffsetType() bool {
	return m.Has(tag.PegOffsetType)
}

//HasPegLimitType returns true if PegLimitType is present, Tag 837
func (m CrossOrderCancelReplaceRequest) HasPegLimitType() bool {
	return m.Has(tag.PegLimitType)
}

//HasPegRoundDirection returns true if PegRoundDirection is present, Tag 838
func (m CrossOrderCancelReplaceRequest) HasPegRoundDirection() bool {
	return m.Has(tag.PegRoundDirection)
}

//HasPegScope returns true if PegScope is present, Tag 840
func (m CrossOrderCancelReplaceRequest) HasPegScope() bool {
	return m.Has(tag.PegScope)
}

//HasDiscretionMoveType returns true if DiscretionMoveType is present, Tag 841
func (m CrossOrderCancelReplaceRequest) HasDiscretionMoveType() bool {
	return m.Has(tag.DiscretionMoveType)
}

//HasDiscretionOffsetType returns true if DiscretionOffsetType is present, Tag 842
func (m CrossOrderCancelReplaceRequest) HasDiscretionOffsetType() bool {
	return m.Has(tag.DiscretionOffsetType)
}

//HasDiscretionLimitType returns true if DiscretionLimitType is present, Tag 843
func (m CrossOrderCancelReplaceRequest) HasDiscretionLimitType() bool {
	return m.Has(tag.DiscretionLimitType)
}

//HasDiscretionRoundDirection returns true if DiscretionRoundDirection is present, Tag 844
func (m CrossOrderCancelReplaceRequest) HasDiscretionRoundDirection() bool {
	return m.Has(tag.DiscretionRoundDirection)
}

//HasDiscretionScope returns true if DiscretionScope is present, Tag 846
func (m CrossOrderCancelReplaceRequest) HasDiscretionScope() bool {
	return m.Has(tag.DiscretionScope)
}

//HasTargetStrategy returns true if TargetStrategy is present, Tag 847
func (m CrossOrderCancelReplaceRequest) HasTargetStrategy() bool {
	return m.Has(tag.TargetStrategy)
}

//HasTargetStrategyParameters returns true if TargetStrategyParameters is present, Tag 848
func (m CrossOrderCancelReplaceRequest) HasTargetStrategyParameters() bool {
	return m.Has(tag.TargetStrategyParameters)
}

//HasParticipationRate returns true if ParticipationRate is present, Tag 849
func (m CrossOrderCancelReplaceRequest) HasParticipationRate() bool {
	return m.Has(tag.ParticipationRate)
}

//HasNoEvents returns true if NoEvents is present, Tag 864
func (m CrossOrderCancelReplaceRequest) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

//HasDatedDate returns true if DatedDate is present, Tag 873
func (m CrossOrderCancelReplaceRequest) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

//HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m CrossOrderCancelReplaceRequest) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

//HasCPProgram returns true if CPProgram is present, Tag 875
func (m CrossOrderCancelReplaceRequest) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

//HasCPRegType returns true if CPRegType is present, Tag 876
func (m CrossOrderCancelReplaceRequest) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

//HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m CrossOrderCancelReplaceRequest) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

//HasNoStrategyParameters returns true if NoStrategyParameters is present, Tag 957
func (m CrossOrderCancelReplaceRequest) HasNoStrategyParameters() bool {
	return m.Has(tag.NoStrategyParameters)
}

//HasHostCrossID returns true if HostCrossID is present, Tag 961
func (m CrossOrderCancelReplaceRequest) HasHostCrossID() bool {
	return m.Has(tag.HostCrossID)
}

//HasSecurityStatus returns true if SecurityStatus is present, Tag 965
func (m CrossOrderCancelReplaceRequest) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

//HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966
func (m CrossOrderCancelReplaceRequest) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

//HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967
func (m CrossOrderCancelReplaceRequest) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

//HasStrikeValue returns true if StrikeValue is present, Tag 968
func (m CrossOrderCancelReplaceRequest) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

//HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969
func (m CrossOrderCancelReplaceRequest) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

//HasPositionLimit returns true if PositionLimit is present, Tag 970
func (m CrossOrderCancelReplaceRequest) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

//HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971
func (m CrossOrderCancelReplaceRequest) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

//HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996
func (m CrossOrderCancelReplaceRequest) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

//HasTimeUnit returns true if TimeUnit is present, Tag 997
func (m CrossOrderCancelReplaceRequest) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

//HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018
func (m CrossOrderCancelReplaceRequest) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

//HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049
func (m CrossOrderCancelReplaceRequest) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

//HasMaturityTime returns true if MaturityTime is present, Tag 1079
func (m CrossOrderCancelReplaceRequest) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

//HasSecondaryDisplayQty returns true if SecondaryDisplayQty is present, Tag 1082
func (m CrossOrderCancelReplaceRequest) HasSecondaryDisplayQty() bool {
	return m.Has(tag.SecondaryDisplayQty)
}

//HasDisplayWhen returns true if DisplayWhen is present, Tag 1083
func (m CrossOrderCancelReplaceRequest) HasDisplayWhen() bool {
	return m.Has(tag.DisplayWhen)
}

//HasDisplayMethod returns true if DisplayMethod is present, Tag 1084
func (m CrossOrderCancelReplaceRequest) HasDisplayMethod() bool {
	return m.Has(tag.DisplayMethod)
}

//HasDisplayLowQty returns true if DisplayLowQty is present, Tag 1085
func (m CrossOrderCancelReplaceRequest) HasDisplayLowQty() bool {
	return m.Has(tag.DisplayLowQty)
}

//HasDisplayHighQty returns true if DisplayHighQty is present, Tag 1086
func (m CrossOrderCancelReplaceRequest) HasDisplayHighQty() bool {
	return m.Has(tag.DisplayHighQty)
}

//HasDisplayMinIncr returns true if DisplayMinIncr is present, Tag 1087
func (m CrossOrderCancelReplaceRequest) HasDisplayMinIncr() bool {
	return m.Has(tag.DisplayMinIncr)
}

//HasRefreshQty returns true if RefreshQty is present, Tag 1088
func (m CrossOrderCancelReplaceRequest) HasRefreshQty() bool {
	return m.Has(tag.RefreshQty)
}

//HasMatchIncrement returns true if MatchIncrement is present, Tag 1089
func (m CrossOrderCancelReplaceRequest) HasMatchIncrement() bool {
	return m.Has(tag.MatchIncrement)
}

//HasMaxPriceLevels returns true if MaxPriceLevels is present, Tag 1090
func (m CrossOrderCancelReplaceRequest) HasMaxPriceLevels() bool {
	return m.Has(tag.MaxPriceLevels)
}

//HasPriceProtectionScope returns true if PriceProtectionScope is present, Tag 1092
func (m CrossOrderCancelReplaceRequest) HasPriceProtectionScope() bool {
	return m.Has(tag.PriceProtectionScope)
}

//HasPegPriceType returns true if PegPriceType is present, Tag 1094
func (m CrossOrderCancelReplaceRequest) HasPegPriceType() bool {
	return m.Has(tag.PegPriceType)
}

//HasPegSecurityIDSource returns true if PegSecurityIDSource is present, Tag 1096
func (m CrossOrderCancelReplaceRequest) HasPegSecurityIDSource() bool {
	return m.Has(tag.PegSecurityIDSource)
}

//HasPegSecurityID returns true if PegSecurityID is present, Tag 1097
func (m CrossOrderCancelReplaceRequest) HasPegSecurityID() bool {
	return m.Has(tag.PegSecurityID)
}

//HasPegSymbol returns true if PegSymbol is present, Tag 1098
func (m CrossOrderCancelReplaceRequest) HasPegSymbol() bool {
	return m.Has(tag.PegSymbol)
}

//HasPegSecurityDesc returns true if PegSecurityDesc is present, Tag 1099
func (m CrossOrderCancelReplaceRequest) HasPegSecurityDesc() bool {
	return m.Has(tag.PegSecurityDesc)
}

//HasTriggerType returns true if TriggerType is present, Tag 1100
func (m CrossOrderCancelReplaceRequest) HasTriggerType() bool {
	return m.Has(tag.TriggerType)
}

//HasTriggerAction returns true if TriggerAction is present, Tag 1101
func (m CrossOrderCancelReplaceRequest) HasTriggerAction() bool {
	return m.Has(tag.TriggerAction)
}

//HasTriggerPrice returns true if TriggerPrice is present, Tag 1102
func (m CrossOrderCancelReplaceRequest) HasTriggerPrice() bool {
	return m.Has(tag.TriggerPrice)
}

//HasTriggerSymbol returns true if TriggerSymbol is present, Tag 1103
func (m CrossOrderCancelReplaceRequest) HasTriggerSymbol() bool {
	return m.Has(tag.TriggerSymbol)
}

//HasTriggerSecurityID returns true if TriggerSecurityID is present, Tag 1104
func (m CrossOrderCancelReplaceRequest) HasTriggerSecurityID() bool {
	return m.Has(tag.TriggerSecurityID)
}

//HasTriggerSecurityIDSource returns true if TriggerSecurityIDSource is present, Tag 1105
func (m CrossOrderCancelReplaceRequest) HasTriggerSecurityIDSource() bool {
	return m.Has(tag.TriggerSecurityIDSource)
}

//HasTriggerSecurityDesc returns true if TriggerSecurityDesc is present, Tag 1106
func (m CrossOrderCancelReplaceRequest) HasTriggerSecurityDesc() bool {
	return m.Has(tag.TriggerSecurityDesc)
}

//HasTriggerPriceType returns true if TriggerPriceType is present, Tag 1107
func (m CrossOrderCancelReplaceRequest) HasTriggerPriceType() bool {
	return m.Has(tag.TriggerPriceType)
}

//HasTriggerPriceTypeScope returns true if TriggerPriceTypeScope is present, Tag 1108
func (m CrossOrderCancelReplaceRequest) HasTriggerPriceTypeScope() bool {
	return m.Has(tag.TriggerPriceTypeScope)
}

//HasTriggerPriceDirection returns true if TriggerPriceDirection is present, Tag 1109
func (m CrossOrderCancelReplaceRequest) HasTriggerPriceDirection() bool {
	return m.Has(tag.TriggerPriceDirection)
}

//HasTriggerNewPrice returns true if TriggerNewPrice is present, Tag 1110
func (m CrossOrderCancelReplaceRequest) HasTriggerNewPrice() bool {
	return m.Has(tag.TriggerNewPrice)
}

//HasTriggerOrderType returns true if TriggerOrderType is present, Tag 1111
func (m CrossOrderCancelReplaceRequest) HasTriggerOrderType() bool {
	return m.Has(tag.TriggerOrderType)
}

//HasTriggerNewQty returns true if TriggerNewQty is present, Tag 1112
func (m CrossOrderCancelReplaceRequest) HasTriggerNewQty() bool {
	return m.Has(tag.TriggerNewQty)
}

//HasTriggerTradingSessionID returns true if TriggerTradingSessionID is present, Tag 1113
func (m CrossOrderCancelReplaceRequest) HasTriggerTradingSessionID() bool {
	return m.Has(tag.TriggerTradingSessionID)
}

//HasTriggerTradingSessionSubID returns true if TriggerTradingSessionSubID is present, Tag 1114
func (m CrossOrderCancelReplaceRequest) HasTriggerTradingSessionSubID() bool {
	return m.Has(tag.TriggerTradingSessionSubID)
}

//HasNoRootPartyIDs returns true if NoRootPartyIDs is present, Tag 1116
func (m CrossOrderCancelReplaceRequest) HasNoRootPartyIDs() bool {
	return m.Has(tag.NoRootPartyIDs)
}

//HasExDestinationIDSource returns true if ExDestinationIDSource is present, Tag 1133
func (m CrossOrderCancelReplaceRequest) HasExDestinationIDSource() bool {
	return m.Has(tag.ExDestinationIDSource)
}

//HasDisplayQty returns true if DisplayQty is present, Tag 1138
func (m CrossOrderCancelReplaceRequest) HasDisplayQty() bool {
	return m.Has(tag.DisplayQty)
}

//NoStipulations is a repeating group element, Tag 232
type NoStipulations struct {
	*quickfix.Group
}

//SetStipulationType sets StipulationType, Tag 233
func (m NoStipulations) SetStipulationType(v enum.StipulationType) {
	m.Set(field.NewStipulationType(v))
}

//SetStipulationValue sets StipulationValue, Tag 234
func (m NoStipulations) SetStipulationValue(v string) {
	m.Set(field.NewStipulationValue(v))
}

//GetStipulationType gets StipulationType, Tag 233
func (m NoStipulations) GetStipulationType() (v enum.StipulationType, err quickfix.MessageRejectError) {
	var f field.StipulationTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStipulationValue gets StipulationValue, Tag 234
func (m NoStipulations) GetStipulationValue() (v string, err quickfix.MessageRejectError) {
	var f field.StipulationValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoTradingSessions) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m NoTradingSessions) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoTradingSessions) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m NoTradingSessions) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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

//NoSecurityAltID is a repeating group element, Tag 454
type NoSecurityAltID struct {
	*quickfix.Group
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
func (m NoSecurityAltID) GetSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) GetSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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

//NoSides is a repeating group element, Tag 552
type NoSides struct {
	*quickfix.Group
}

//SetSide sets Side, Tag 54
func (m NoSides) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m NoSides) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526
func (m NoSides) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

//SetClOrdLinkID sets ClOrdLinkID, Tag 583
func (m NoSides) SetClOrdLinkID(v string) {
	m.Set(field.NewClOrdLinkID(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m NoSides) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetTradeOriginationDate sets TradeOriginationDate, Tag 229
func (m NoSides) SetTradeOriginationDate(v string) {
	m.Set(field.NewTradeOriginationDate(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m NoSides) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetAccount sets Account, Tag 1
func (m NoSides) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetAcctIDSource sets AcctIDSource, Tag 660
func (m NoSides) SetAcctIDSource(v enum.AcctIDSource) {
	m.Set(field.NewAcctIDSource(v))
}

//SetAccountType sets AccountType, Tag 581
func (m NoSides) SetAccountType(v enum.AccountType) {
	m.Set(field.NewAccountType(v))
}

//SetDayBookingInst sets DayBookingInst, Tag 589
func (m NoSides) SetDayBookingInst(v enum.DayBookingInst) {
	m.Set(field.NewDayBookingInst(v))
}

//SetBookingUnit sets BookingUnit, Tag 590
func (m NoSides) SetBookingUnit(v enum.BookingUnit) {
	m.Set(field.NewBookingUnit(v))
}

//SetPreallocMethod sets PreallocMethod, Tag 591
func (m NoSides) SetPreallocMethod(v enum.PreallocMethod) {
	m.Set(field.NewPreallocMethod(v))
}

//SetAllocID sets AllocID, Tag 70
func (m NoSides) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

//SetNoAllocs sets NoAllocs, Tag 78
func (m NoSides) SetNoAllocs(f NoAllocsRepeatingGroup) {
	m.SetGroup(f)
}

//SetQtyType sets QtyType, Tag 854
func (m NoSides) SetQtyType(v enum.QtyType) {
	m.Set(field.NewQtyType(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m NoSides) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m NoSides) SetCashOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCashOrderQty(value, scale))
}

//SetOrderPercent sets OrderPercent, Tag 516
func (m NoSides) SetOrderPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderPercent(value, scale))
}

//SetRoundingDirection sets RoundingDirection, Tag 468
func (m NoSides) SetRoundingDirection(v enum.RoundingDirection) {
	m.Set(field.NewRoundingDirection(v))
}

//SetRoundingModulus sets RoundingModulus, Tag 469
func (m NoSides) SetRoundingModulus(value decimal.Decimal, scale int32) {
	m.Set(field.NewRoundingModulus(value, scale))
}

//SetCommission sets Commission, Tag 12
func (m NoSides) SetCommission(value decimal.Decimal, scale int32) {
	m.Set(field.NewCommission(value, scale))
}

//SetCommType sets CommType, Tag 13
func (m NoSides) SetCommType(v enum.CommType) {
	m.Set(field.NewCommType(v))
}

//SetCommCurrency sets CommCurrency, Tag 479
func (m NoSides) SetCommCurrency(v string) {
	m.Set(field.NewCommCurrency(v))
}

//SetFundRenewWaiv sets FundRenewWaiv, Tag 497
func (m NoSides) SetFundRenewWaiv(v enum.FundRenewWaiv) {
	m.Set(field.NewFundRenewWaiv(v))
}

//SetOrderCapacity sets OrderCapacity, Tag 528
func (m NoSides) SetOrderCapacity(v enum.OrderCapacity) {
	m.Set(field.NewOrderCapacity(v))
}

//SetOrderRestrictions sets OrderRestrictions, Tag 529
func (m NoSides) SetOrderRestrictions(v enum.OrderRestrictions) {
	m.Set(field.NewOrderRestrictions(v))
}

//SetCustOrderCapacity sets CustOrderCapacity, Tag 582
func (m NoSides) SetCustOrderCapacity(v enum.CustOrderCapacity) {
	m.Set(field.NewCustOrderCapacity(v))
}

//SetForexReq sets ForexReq, Tag 121
func (m NoSides) SetForexReq(v bool) {
	m.Set(field.NewForexReq(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m NoSides) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetBookingType sets BookingType, Tag 775
func (m NoSides) SetBookingType(v enum.BookingType) {
	m.Set(field.NewBookingType(v))
}

//SetText sets Text, Tag 58
func (m NoSides) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m NoSides) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m NoSides) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetPositionEffect sets PositionEffect, Tag 77
func (m NoSides) SetPositionEffect(v enum.PositionEffect) {
	m.Set(field.NewPositionEffect(v))
}

//SetCoveredOrUncovered sets CoveredOrUncovered, Tag 203
func (m NoSides) SetCoveredOrUncovered(v enum.CoveredOrUncovered) {
	m.Set(field.NewCoveredOrUncovered(v))
}

//SetCashMargin sets CashMargin, Tag 544
func (m NoSides) SetCashMargin(v enum.CashMargin) {
	m.Set(field.NewCashMargin(v))
}

//SetClearingFeeIndicator sets ClearingFeeIndicator, Tag 635
func (m NoSides) SetClearingFeeIndicator(v enum.ClearingFeeIndicator) {
	m.Set(field.NewClearingFeeIndicator(v))
}

//SetSolicitedFlag sets SolicitedFlag, Tag 377
func (m NoSides) SetSolicitedFlag(v bool) {
	m.Set(field.NewSolicitedFlag(v))
}

//SetSideComplianceID sets SideComplianceID, Tag 659
func (m NoSides) SetSideComplianceID(v string) {
	m.Set(field.NewSideComplianceID(v))
}

//SetSideTimeInForce sets SideTimeInForce, Tag 962
func (m NoSides) SetSideTimeInForce(v time.Time) {
	m.Set(field.NewSideTimeInForce(v))
}

//SetPreTradeAnonymity sets PreTradeAnonymity, Tag 1091
func (m NoSides) SetPreTradeAnonymity(v bool) {
	m.Set(field.NewPreTradeAnonymity(v))
}

//GetSide gets Side, Tag 54
func (m NoSides) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m NoSides) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526
func (m NoSides) GetSecondaryClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetClOrdLinkID gets ClOrdLinkID, Tag 583
func (m NoSides) GetClOrdLinkID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdLinkIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m NoSides) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetTradeOriginationDate gets TradeOriginationDate, Tag 229
func (m NoSides) GetTradeOriginationDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeOriginationDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m NoSides) GetTradeDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAccount gets Account, Tag 1
func (m NoSides) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAcctIDSource gets AcctIDSource, Tag 660
func (m NoSides) GetAcctIDSource() (v enum.AcctIDSource, err quickfix.MessageRejectError) {
	var f field.AcctIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAccountType gets AccountType, Tag 581
func (m NoSides) GetAccountType() (v enum.AccountType, err quickfix.MessageRejectError) {
	var f field.AccountTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDayBookingInst gets DayBookingInst, Tag 589
func (m NoSides) GetDayBookingInst() (v enum.DayBookingInst, err quickfix.MessageRejectError) {
	var f field.DayBookingInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBookingUnit gets BookingUnit, Tag 590
func (m NoSides) GetBookingUnit() (v enum.BookingUnit, err quickfix.MessageRejectError) {
	var f field.BookingUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPreallocMethod gets PreallocMethod, Tag 591
func (m NoSides) GetPreallocMethod() (v enum.PreallocMethod, err quickfix.MessageRejectError) {
	var f field.PreallocMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocID gets AllocID, Tag 70
func (m NoSides) GetAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.AllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoAllocs gets NoAllocs, Tag 78
func (m NoSides) GetNoAllocs() (NoAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetQtyType gets QtyType, Tag 854
func (m NoSides) GetQtyType() (v enum.QtyType, err quickfix.MessageRejectError) {
	var f field.QtyTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m NoSides) GetOrderQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m NoSides) GetCashOrderQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CashOrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderPercent gets OrderPercent, Tag 516
func (m NoSides) GetOrderPercent() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderPercentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRoundingDirection gets RoundingDirection, Tag 468
func (m NoSides) GetRoundingDirection() (v enum.RoundingDirection, err quickfix.MessageRejectError) {
	var f field.RoundingDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRoundingModulus gets RoundingModulus, Tag 469
func (m NoSides) GetRoundingModulus() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RoundingModulusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommission gets Commission, Tag 12
func (m NoSides) GetCommission() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CommissionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommType gets CommType, Tag 13
func (m NoSides) GetCommType() (v enum.CommType, err quickfix.MessageRejectError) {
	var f field.CommTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommCurrency gets CommCurrency, Tag 479
func (m NoSides) GetCommCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CommCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFundRenewWaiv gets FundRenewWaiv, Tag 497
func (m NoSides) GetFundRenewWaiv() (v enum.FundRenewWaiv, err quickfix.MessageRejectError) {
	var f field.FundRenewWaivField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderCapacity gets OrderCapacity, Tag 528
func (m NoSides) GetOrderCapacity() (v enum.OrderCapacity, err quickfix.MessageRejectError) {
	var f field.OrderCapacityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderRestrictions gets OrderRestrictions, Tag 529
func (m NoSides) GetOrderRestrictions() (v enum.OrderRestrictions, err quickfix.MessageRejectError) {
	var f field.OrderRestrictionsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCustOrderCapacity gets CustOrderCapacity, Tag 582
func (m NoSides) GetCustOrderCapacity() (v enum.CustOrderCapacity, err quickfix.MessageRejectError) {
	var f field.CustOrderCapacityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetForexReq gets ForexReq, Tag 121
func (m NoSides) GetForexReq() (v bool, err quickfix.MessageRejectError) {
	var f field.ForexReqField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m NoSides) GetSettlCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.SettlCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBookingType gets BookingType, Tag 775
func (m NoSides) GetBookingType() (v enum.BookingType, err quickfix.MessageRejectError) {
	var f field.BookingTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m NoSides) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m NoSides) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m NoSides) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPositionEffect gets PositionEffect, Tag 77
func (m NoSides) GetPositionEffect() (v enum.PositionEffect, err quickfix.MessageRejectError) {
	var f field.PositionEffectField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCoveredOrUncovered gets CoveredOrUncovered, Tag 203
func (m NoSides) GetCoveredOrUncovered() (v enum.CoveredOrUncovered, err quickfix.MessageRejectError) {
	var f field.CoveredOrUncoveredField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashMargin gets CashMargin, Tag 544
func (m NoSides) GetCashMargin() (v enum.CashMargin, err quickfix.MessageRejectError) {
	var f field.CashMarginField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetClearingFeeIndicator gets ClearingFeeIndicator, Tag 635
func (m NoSides) GetClearingFeeIndicator() (v enum.ClearingFeeIndicator, err quickfix.MessageRejectError) {
	var f field.ClearingFeeIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSolicitedFlag gets SolicitedFlag, Tag 377
func (m NoSides) GetSolicitedFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.SolicitedFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSideComplianceID gets SideComplianceID, Tag 659
func (m NoSides) GetSideComplianceID() (v string, err quickfix.MessageRejectError) {
	var f field.SideComplianceIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSideTimeInForce gets SideTimeInForce, Tag 962
func (m NoSides) GetSideTimeInForce() (v time.Time, err quickfix.MessageRejectError) {
	var f field.SideTimeInForceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPreTradeAnonymity gets PreTradeAnonymity, Tag 1091
func (m NoSides) GetPreTradeAnonymity() (v bool, err quickfix.MessageRejectError) {
	var f field.PreTradeAnonymityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasSide returns true if Side is present, Tag 54
func (m NoSides) HasSide() bool {
	return m.Has(tag.Side)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m NoSides) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526
func (m NoSides) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

//HasClOrdLinkID returns true if ClOrdLinkID is present, Tag 583
func (m NoSides) HasClOrdLinkID() bool {
	return m.Has(tag.ClOrdLinkID)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m NoSides) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasTradeOriginationDate returns true if TradeOriginationDate is present, Tag 229
func (m NoSides) HasTradeOriginationDate() bool {
	return m.Has(tag.TradeOriginationDate)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m NoSides) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasAccount returns true if Account is present, Tag 1
func (m NoSides) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasAcctIDSource returns true if AcctIDSource is present, Tag 660
func (m NoSides) HasAcctIDSource() bool {
	return m.Has(tag.AcctIDSource)
}

//HasAccountType returns true if AccountType is present, Tag 581
func (m NoSides) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

//HasDayBookingInst returns true if DayBookingInst is present, Tag 589
func (m NoSides) HasDayBookingInst() bool {
	return m.Has(tag.DayBookingInst)
}

//HasBookingUnit returns true if BookingUnit is present, Tag 590
func (m NoSides) HasBookingUnit() bool {
	return m.Has(tag.BookingUnit)
}

//HasPreallocMethod returns true if PreallocMethod is present, Tag 591
func (m NoSides) HasPreallocMethod() bool {
	return m.Has(tag.PreallocMethod)
}

//HasAllocID returns true if AllocID is present, Tag 70
func (m NoSides) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

//HasNoAllocs returns true if NoAllocs is present, Tag 78
func (m NoSides) HasNoAllocs() bool {
	return m.Has(tag.NoAllocs)
}

//HasQtyType returns true if QtyType is present, Tag 854
func (m NoSides) HasQtyType() bool {
	return m.Has(tag.QtyType)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m NoSides) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m NoSides) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasOrderPercent returns true if OrderPercent is present, Tag 516
func (m NoSides) HasOrderPercent() bool {
	return m.Has(tag.OrderPercent)
}

//HasRoundingDirection returns true if RoundingDirection is present, Tag 468
func (m NoSides) HasRoundingDirection() bool {
	return m.Has(tag.RoundingDirection)
}

//HasRoundingModulus returns true if RoundingModulus is present, Tag 469
func (m NoSides) HasRoundingModulus() bool {
	return m.Has(tag.RoundingModulus)
}

//HasCommission returns true if Commission is present, Tag 12
func (m NoSides) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m NoSides) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasCommCurrency returns true if CommCurrency is present, Tag 479
func (m NoSides) HasCommCurrency() bool {
	return m.Has(tag.CommCurrency)
}

//HasFundRenewWaiv returns true if FundRenewWaiv is present, Tag 497
func (m NoSides) HasFundRenewWaiv() bool {
	return m.Has(tag.FundRenewWaiv)
}

//HasOrderCapacity returns true if OrderCapacity is present, Tag 528
func (m NoSides) HasOrderCapacity() bool {
	return m.Has(tag.OrderCapacity)
}

//HasOrderRestrictions returns true if OrderRestrictions is present, Tag 529
func (m NoSides) HasOrderRestrictions() bool {
	return m.Has(tag.OrderRestrictions)
}

//HasCustOrderCapacity returns true if CustOrderCapacity is present, Tag 582
func (m NoSides) HasCustOrderCapacity() bool {
	return m.Has(tag.CustOrderCapacity)
}

//HasForexReq returns true if ForexReq is present, Tag 121
func (m NoSides) HasForexReq() bool {
	return m.Has(tag.ForexReq)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m NoSides) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasBookingType returns true if BookingType is present, Tag 775
func (m NoSides) HasBookingType() bool {
	return m.Has(tag.BookingType)
}

//HasText returns true if Text is present, Tag 58
func (m NoSides) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m NoSides) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m NoSides) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasPositionEffect returns true if PositionEffect is present, Tag 77
func (m NoSides) HasPositionEffect() bool {
	return m.Has(tag.PositionEffect)
}

//HasCoveredOrUncovered returns true if CoveredOrUncovered is present, Tag 203
func (m NoSides) HasCoveredOrUncovered() bool {
	return m.Has(tag.CoveredOrUncovered)
}

//HasCashMargin returns true if CashMargin is present, Tag 544
func (m NoSides) HasCashMargin() bool {
	return m.Has(tag.CashMargin)
}

//HasClearingFeeIndicator returns true if ClearingFeeIndicator is present, Tag 635
func (m NoSides) HasClearingFeeIndicator() bool {
	return m.Has(tag.ClearingFeeIndicator)
}

//HasSolicitedFlag returns true if SolicitedFlag is present, Tag 377
func (m NoSides) HasSolicitedFlag() bool {
	return m.Has(tag.SolicitedFlag)
}

//HasSideComplianceID returns true if SideComplianceID is present, Tag 659
func (m NoSides) HasSideComplianceID() bool {
	return m.Has(tag.SideComplianceID)
}

//HasSideTimeInForce returns true if SideTimeInForce is present, Tag 962
func (m NoSides) HasSideTimeInForce() bool {
	return m.Has(tag.SideTimeInForce)
}

//HasPreTradeAnonymity returns true if PreTradeAnonymity is present, Tag 1091
func (m NoSides) HasPreTradeAnonymity() bool {
	return m.Has(tag.PreTradeAnonymity)
}

//NoPartyIDs is a repeating group element, Tag 453
type NoPartyIDs struct {
	*quickfix.Group
}

//SetPartyID sets PartyID, Tag 448
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

//SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

//SetPartyRole sets PartyRole, Tag 452
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

//SetNoPartySubIDs sets NoPartySubIDs, Tag 802
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetPartyID gets PartyID, Tag 448
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyRole gets PartyRole, Tag 452
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//SetPartySubIDType sets PartySubIDType, Tag 803
func (m NoPartySubIDs) SetPartySubIDType(v enum.PartySubIDType) {
	m.Set(field.NewPartySubIDType(v))
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartySubIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartySubIDType gets PartySubIDType, Tag 803
func (m NoPartySubIDs) GetPartySubIDType() (v enum.PartySubIDType, err quickfix.MessageRejectError) {
	var f field.PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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

//NoAllocs is a repeating group element, Tag 78
type NoAllocs struct {
	*quickfix.Group
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
func (m NoAllocs) GetAllocAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AllocAccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocAcctIDSource gets AllocAcctIDSource, Tag 661
func (m NoAllocs) GetAllocAcctIDSource() (v int, err quickfix.MessageRejectError) {
	var f field.AllocAcctIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocSettlCurrency gets AllocSettlCurrency, Tag 736
func (m NoAllocs) GetAllocSettlCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.AllocSettlCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIndividualAllocID gets IndividualAllocID, Tag 467
func (m NoAllocs) GetIndividualAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.IndividualAllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoNestedPartyIDs gets NoNestedPartyIDs, Tag 539
func (m NoAllocs) GetNoNestedPartyIDs() (NoNestedPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAllocQty gets AllocQty, Tag 80
func (m NoAllocs) GetAllocQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AllocQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
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
func (m NoNestedPartyIDs) GetNestedPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNestedPartyIDSource gets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) GetNestedPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNestedPartyRole gets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) GetNestedPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.NestedPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
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
func (m NoNestedPartySubIDs) GetNestedPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNestedPartySubIDType gets NestedPartySubIDType, Tag 805
func (m NoNestedPartySubIDs) GetNestedPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.NestedPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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

//NoSidesRepeatingGroup is a repeating group, Tag 552
type NoSidesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSidesRepeatingGroup returns an initialized, NoSidesRepeatingGroup
func NewNoSidesRepeatingGroup() NoSidesRepeatingGroup {
	return NoSidesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSides,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Side), quickfix.GroupElement(tag.ClOrdID), quickfix.GroupElement(tag.SecondaryClOrdID), quickfix.GroupElement(tag.ClOrdLinkID), NewNoPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.TradeOriginationDate), quickfix.GroupElement(tag.TradeDate), quickfix.GroupElement(tag.Account), quickfix.GroupElement(tag.AcctIDSource), quickfix.GroupElement(tag.AccountType), quickfix.GroupElement(tag.DayBookingInst), quickfix.GroupElement(tag.BookingUnit), quickfix.GroupElement(tag.PreallocMethod), quickfix.GroupElement(tag.AllocID), NewNoAllocsRepeatingGroup(), quickfix.GroupElement(tag.QtyType), quickfix.GroupElement(tag.OrderQty), quickfix.GroupElement(tag.CashOrderQty), quickfix.GroupElement(tag.OrderPercent), quickfix.GroupElement(tag.RoundingDirection), quickfix.GroupElement(tag.RoundingModulus), quickfix.GroupElement(tag.Commission), quickfix.GroupElement(tag.CommType), quickfix.GroupElement(tag.CommCurrency), quickfix.GroupElement(tag.FundRenewWaiv), quickfix.GroupElement(tag.OrderCapacity), quickfix.GroupElement(tag.OrderRestrictions), quickfix.GroupElement(tag.CustOrderCapacity), quickfix.GroupElement(tag.ForexReq), quickfix.GroupElement(tag.SettlCurrency), quickfix.GroupElement(tag.BookingType), quickfix.GroupElement(tag.Text), quickfix.GroupElement(tag.EncodedTextLen), quickfix.GroupElement(tag.EncodedText), quickfix.GroupElement(tag.PositionEffect), quickfix.GroupElement(tag.CoveredOrUncovered), quickfix.GroupElement(tag.CashMargin), quickfix.GroupElement(tag.ClearingFeeIndicator), quickfix.GroupElement(tag.SolicitedFlag), quickfix.GroupElement(tag.SideComplianceID), quickfix.GroupElement(tag.SideTimeInForce), quickfix.GroupElement(tag.PreTradeAnonymity)})}
}

//Add create and append a new NoSides to this group
func (m NoSidesRepeatingGroup) Add() NoSides {
	g := m.RepeatingGroup.Add()
	return NoSides{g}
}

//Get returns the ith NoSides in the NoSidesRepeatinGroup
func (m NoSidesRepeatingGroup) Get(i int) NoSides {
	return NoSides{m.RepeatingGroup.Get(i)}
}

//NoLegs is a repeating group element, Tag 555
type NoLegs struct {
	*quickfix.Group
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
func (m NoLegs) SetLegRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegRepurchaseRate(value, scale))
}

//SetLegFactor sets LegFactor, Tag 253
func (m NoLegs) SetLegFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegFactor(value, scale))
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
func (m NoLegs) SetLegStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegStrikePrice(value, scale))
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
func (m NoLegs) SetLegContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegContractMultiplier(value, scale))
}

//SetLegCouponRate sets LegCouponRate, Tag 615
func (m NoLegs) SetLegCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegCouponRate(value, scale))
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
func (m NoLegs) SetLegRatioQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegRatioQty(value, scale))
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

//GetLegSymbol gets LegSymbol, Tag 600
func (m NoLegs) GetLegSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.LegSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSymbolSfx gets LegSymbolSfx, Tag 601
func (m NoLegs) GetLegSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.LegSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSecurityID gets LegSecurityID, Tag 602
func (m NoLegs) GetLegSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSecurityIDSource gets LegSecurityIDSource, Tag 603
func (m NoLegs) GetLegSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoLegSecurityAltID gets NoLegSecurityAltID, Tag 604
func (m NoLegs) GetNoLegSecurityAltID() (NoLegSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLegProduct gets LegProduct, Tag 607
func (m NoLegs) GetLegProduct() (v int, err quickfix.MessageRejectError) {
	var f field.LegProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegCFICode gets LegCFICode, Tag 608
func (m NoLegs) GetLegCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.LegCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSecurityType gets LegSecurityType, Tag 609
func (m NoLegs) GetLegSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSecuritySubType gets LegSecuritySubType, Tag 764
func (m NoLegs) GetLegSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegMaturityMonthYear gets LegMaturityMonthYear, Tag 610
func (m NoLegs) GetLegMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegMaturityDate gets LegMaturityDate, Tag 611
func (m NoLegs) GetLegMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegCouponPaymentDate gets LegCouponPaymentDate, Tag 248
func (m NoLegs) GetLegCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegIssueDate gets LegIssueDate, Tag 249
func (m NoLegs) GetLegIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegRepoCollateralSecurityType gets LegRepoCollateralSecurityType, Tag 250
func (m NoLegs) GetLegRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.LegRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegRepurchaseTerm gets LegRepurchaseTerm, Tag 251
func (m NoLegs) GetLegRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.LegRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegRepurchaseRate gets LegRepurchaseRate, Tag 252
func (m NoLegs) GetLegRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegFactor gets LegFactor, Tag 253
func (m NoLegs) GetLegFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegFactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegCreditRating gets LegCreditRating, Tag 257
func (m NoLegs) GetLegCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.LegCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegInstrRegistry gets LegInstrRegistry, Tag 599
func (m NoLegs) GetLegInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.LegInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegCountryOfIssue gets LegCountryOfIssue, Tag 596
func (m NoLegs) GetLegCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegStateOrProvinceOfIssue gets LegStateOrProvinceOfIssue, Tag 597
func (m NoLegs) GetLegStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegLocaleOfIssue gets LegLocaleOfIssue, Tag 598
func (m NoLegs) GetLegLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegRedemptionDate gets LegRedemptionDate, Tag 254
func (m NoLegs) GetLegRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegStrikePrice gets LegStrikePrice, Tag 612
func (m NoLegs) GetLegStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegStrikeCurrency gets LegStrikeCurrency, Tag 942
func (m NoLegs) GetLegStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.LegStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegOptAttribute gets LegOptAttribute, Tag 613
func (m NoLegs) GetLegOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.LegOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegContractMultiplier gets LegContractMultiplier, Tag 614
func (m NoLegs) GetLegContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegCouponRate gets LegCouponRate, Tag 615
func (m NoLegs) GetLegCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSecurityExchange gets LegSecurityExchange, Tag 616
func (m NoLegs) GetLegSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegIssuer gets LegIssuer, Tag 617
func (m NoLegs) GetLegIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.LegIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedLegIssuerLen gets EncodedLegIssuerLen, Tag 618
func (m NoLegs) GetEncodedLegIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedLegIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedLegIssuer gets EncodedLegIssuer, Tag 619
func (m NoLegs) GetEncodedLegIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedLegIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSecurityDesc gets LegSecurityDesc, Tag 620
func (m NoLegs) GetLegSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedLegSecurityDescLen gets EncodedLegSecurityDescLen, Tag 621
func (m NoLegs) GetEncodedLegSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedLegSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedLegSecurityDesc gets EncodedLegSecurityDesc, Tag 622
func (m NoLegs) GetEncodedLegSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedLegSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegRatioQty gets LegRatioQty, Tag 623
func (m NoLegs) GetLegRatioQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegRatioQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSide gets LegSide, Tag 624
func (m NoLegs) GetLegSide() (v string, err quickfix.MessageRejectError) {
	var f field.LegSideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegCurrency gets LegCurrency, Tag 556
func (m NoLegs) GetLegCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.LegCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegPool gets LegPool, Tag 740
func (m NoLegs) GetLegPool() (v string, err quickfix.MessageRejectError) {
	var f field.LegPoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegDatedDate gets LegDatedDate, Tag 739
func (m NoLegs) GetLegDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegDatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegContractSettlMonth gets LegContractSettlMonth, Tag 955
func (m NoLegs) GetLegContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.LegContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegInterestAccrualDate gets LegInterestAccrualDate, Tag 956
func (m NoLegs) GetLegInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegInterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegUnitOfMeasure gets LegUnitOfMeasure, Tag 999
func (m NoLegs) GetLegUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.LegUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegTimeUnit gets LegTimeUnit, Tag 1001
func (m NoLegs) GetLegTimeUnit() (v string, err quickfix.MessageRejectError) {
	var f field.LegTimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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

//NoLegSecurityAltID is a repeating group element, Tag 604
type NoLegSecurityAltID struct {
	*quickfix.Group
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
func (m NoLegSecurityAltID) GetLegSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSecurityAltIDSource gets LegSecurityAltIDSource, Tag 606
func (m NoLegSecurityAltID) GetLegSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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

//NoLegsRepeatingGroup is a repeating group, Tag 555
type NoLegsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegsRepeatingGroup returns an initialized, NoLegsRepeatingGroup
func NewNoLegsRepeatingGroup() NoLegsRepeatingGroup {
	return NoLegsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSymbol), quickfix.GroupElement(tag.LegSymbolSfx), quickfix.GroupElement(tag.LegSecurityID), quickfix.GroupElement(tag.LegSecurityIDSource), NewNoLegSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.LegProduct), quickfix.GroupElement(tag.LegCFICode), quickfix.GroupElement(tag.LegSecurityType), quickfix.GroupElement(tag.LegSecuritySubType), quickfix.GroupElement(tag.LegMaturityMonthYear), quickfix.GroupElement(tag.LegMaturityDate), quickfix.GroupElement(tag.LegCouponPaymentDate), quickfix.GroupElement(tag.LegIssueDate), quickfix.GroupElement(tag.LegRepoCollateralSecurityType), quickfix.GroupElement(tag.LegRepurchaseTerm), quickfix.GroupElement(tag.LegRepurchaseRate), quickfix.GroupElement(tag.LegFactor), quickfix.GroupElement(tag.LegCreditRating), quickfix.GroupElement(tag.LegInstrRegistry), quickfix.GroupElement(tag.LegCountryOfIssue), quickfix.GroupElement(tag.LegStateOrProvinceOfIssue), quickfix.GroupElement(tag.LegLocaleOfIssue), quickfix.GroupElement(tag.LegRedemptionDate), quickfix.GroupElement(tag.LegStrikePrice), quickfix.GroupElement(tag.LegStrikeCurrency), quickfix.GroupElement(tag.LegOptAttribute), quickfix.GroupElement(tag.LegContractMultiplier), quickfix.GroupElement(tag.LegCouponRate), quickfix.GroupElement(tag.LegSecurityExchange), quickfix.GroupElement(tag.LegIssuer), quickfix.GroupElement(tag.EncodedLegIssuerLen), quickfix.GroupElement(tag.EncodedLegIssuer), quickfix.GroupElement(tag.LegSecurityDesc), quickfix.GroupElement(tag.EncodedLegSecurityDescLen), quickfix.GroupElement(tag.EncodedLegSecurityDesc), quickfix.GroupElement(tag.LegRatioQty), quickfix.GroupElement(tag.LegSide), quickfix.GroupElement(tag.LegCurrency), quickfix.GroupElement(tag.LegPool), quickfix.GroupElement(tag.LegDatedDate), quickfix.GroupElement(tag.LegContractSettlMonth), quickfix.GroupElement(tag.LegInterestAccrualDate), quickfix.GroupElement(tag.LegUnitOfMeasure), quickfix.GroupElement(tag.LegTimeUnit)})}
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
	*quickfix.Group
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

//SetUnderlyingAllocationPercent sets UnderlyingAllocationPercent, Tag 972
func (m NoUnderlyings) SetUnderlyingAllocationPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAllocationPercent(value, scale))
}

//SetUnderlyingSettlementType sets UnderlyingSettlementType, Tag 975
func (m NoUnderlyings) SetUnderlyingSettlementType(v enum.UnderlyingSettlementType) {
	m.Set(field.NewUnderlyingSettlementType(v))
}

//SetUnderlyingCashAmount sets UnderlyingCashAmount, Tag 973
func (m NoUnderlyings) SetUnderlyingCashAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCashAmount(value, scale))
}

//SetUnderlyingCashType sets UnderlyingCashType, Tag 974
func (m NoUnderlyings) SetUnderlyingCashType(v enum.UnderlyingCashType) {
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
func (m NoUnderlyings) SetUnderlyingCapValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCapValue(value, scale))
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
func (m NoUnderlyings) SetUnderlyingAdjustedQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAdjustedQuantity(value, scale))
}

//SetUnderlyingFXRate sets UnderlyingFXRate, Tag 1045
func (m NoUnderlyings) SetUnderlyingFXRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFXRate(value, scale))
}

//SetUnderlyingFXRateCalc sets UnderlyingFXRateCalc, Tag 1046
func (m NoUnderlyings) SetUnderlyingFXRateCalc(v enum.UnderlyingFXRateCalc) {
	m.Set(field.NewUnderlyingFXRateCalc(v))
}

//GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311
func (m NoUnderlyings) GetUnderlyingSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312
func (m NoUnderlyings) GetUnderlyingSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309
func (m NoUnderlyings) GetUnderlyingSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityIDSource gets UnderlyingSecurityIDSource, Tag 305
func (m NoUnderlyings) GetUnderlyingSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoUnderlyingSecurityAltID gets NoUnderlyingSecurityAltID, Tag 457
func (m NoUnderlyings) GetNoUnderlyingSecurityAltID() (NoUnderlyingSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingProduct gets UnderlyingProduct, Tag 462
func (m NoUnderlyings) GetUnderlyingProduct() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCFICode gets UnderlyingCFICode, Tag 463
func (m NoUnderlyings) GetUnderlyingCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310
func (m NoUnderlyings) GetUnderlyingSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecuritySubType gets UnderlyingSecuritySubType, Tag 763
func (m NoUnderlyings) GetUnderlyingSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313
func (m NoUnderlyings) GetUnderlyingMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingMaturityDate gets UnderlyingMaturityDate, Tag 542
func (m NoUnderlyings) GetUnderlyingMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCouponPaymentDate gets UnderlyingCouponPaymentDate, Tag 241
func (m NoUnderlyings) GetUnderlyingCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingIssueDate gets UnderlyingIssueDate, Tag 242
func (m NoUnderlyings) GetUnderlyingIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRepoCollateralSecurityType gets UnderlyingRepoCollateralSecurityType, Tag 243
func (m NoUnderlyings) GetUnderlyingRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRepurchaseTerm gets UnderlyingRepurchaseTerm, Tag 244
func (m NoUnderlyings) GetUnderlyingRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRepurchaseRate gets UnderlyingRepurchaseRate, Tag 245
func (m NoUnderlyings) GetUnderlyingRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingFactor gets UnderlyingFactor, Tag 246
func (m NoUnderlyings) GetUnderlyingFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingFactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCreditRating gets UnderlyingCreditRating, Tag 256
func (m NoUnderlyings) GetUnderlyingCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingInstrRegistry gets UnderlyingInstrRegistry, Tag 595
func (m NoUnderlyings) GetUnderlyingInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCountryOfIssue gets UnderlyingCountryOfIssue, Tag 592
func (m NoUnderlyings) GetUnderlyingCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStateOrProvinceOfIssue gets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m NoUnderlyings) GetUnderlyingStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingLocaleOfIssue gets UnderlyingLocaleOfIssue, Tag 594
func (m NoUnderlyings) GetUnderlyingLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRedemptionDate gets UnderlyingRedemptionDate, Tag 247
func (m NoUnderlyings) GetUnderlyingRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316
func (m NoUnderlyings) GetUnderlyingStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStrikeCurrency gets UnderlyingStrikeCurrency, Tag 941
func (m NoUnderlyings) GetUnderlyingStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317
func (m NoUnderlyings) GetUnderlyingOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436
func (m NoUnderlyings) GetUnderlyingContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435
func (m NoUnderlyings) GetUnderlyingCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308
func (m NoUnderlyings) GetUnderlyingSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306
func (m NoUnderlyings) GetUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362
func (m NoUnderlyings) GetEncodedUnderlyingIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363
func (m NoUnderlyings) GetEncodedUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307
func (m NoUnderlyings) GetUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCPProgram gets UnderlyingCPProgram, Tag 877
func (m NoUnderlyings) GetUnderlyingCPProgram() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCPRegType gets UnderlyingCPRegType, Tag 878
func (m NoUnderlyings) GetUnderlyingCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCurrency gets UnderlyingCurrency, Tag 318
func (m NoUnderlyings) GetUnderlyingCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingQty gets UnderlyingQty, Tag 879
func (m NoUnderlyings) GetUnderlyingQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingPx gets UnderlyingPx, Tag 810
func (m NoUnderlyings) GetUnderlyingPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingDirtyPrice gets UnderlyingDirtyPrice, Tag 882
func (m NoUnderlyings) GetUnderlyingDirtyPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingDirtyPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingEndPrice gets UnderlyingEndPrice, Tag 883
func (m NoUnderlyings) GetUnderlyingEndPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStartValue gets UnderlyingStartValue, Tag 884
func (m NoUnderlyings) GetUnderlyingStartValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStartValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCurrentValue gets UnderlyingCurrentValue, Tag 885
func (m NoUnderlyings) GetUnderlyingCurrentValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrentValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingEndValue gets UnderlyingEndValue, Tag 886
func (m NoUnderlyings) GetUnderlyingEndValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoUnderlyingStips gets NoUnderlyingStips, Tag 887
func (m NoUnderlyings) GetNoUnderlyingStips() (NoUnderlyingStipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingStipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingAllocationPercent gets UnderlyingAllocationPercent, Tag 972
func (m NoUnderlyings) GetUnderlyingAllocationPercent() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAllocationPercentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSettlementType gets UnderlyingSettlementType, Tag 975
func (m NoUnderlyings) GetUnderlyingSettlementType() (v enum.UnderlyingSettlementType, err quickfix.MessageRejectError) {
	var f field.UnderlyingSettlementTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCashAmount gets UnderlyingCashAmount, Tag 973
func (m NoUnderlyings) GetUnderlyingCashAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCashAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCashType gets UnderlyingCashType, Tag 974
func (m NoUnderlyings) GetUnderlyingCashType() (v enum.UnderlyingCashType, err quickfix.MessageRejectError) {
	var f field.UnderlyingCashTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingUnitOfMeasure gets UnderlyingUnitOfMeasure, Tag 998
func (m NoUnderlyings) GetUnderlyingUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingTimeUnit gets UnderlyingTimeUnit, Tag 1000
func (m NoUnderlyings) GetUnderlyingTimeUnit() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingTimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCapValue gets UnderlyingCapValue, Tag 1038
func (m NoUnderlyings) GetUnderlyingCapValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCapValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoUndlyInstrumentParties gets NoUndlyInstrumentParties, Tag 1058
func (m NoUnderlyings) GetNoUndlyInstrumentParties() (NoUndlyInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingSettlMethod gets UnderlyingSettlMethod, Tag 1039
func (m NoUnderlyings) GetUnderlyingSettlMethod() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSettlMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingAdjustedQuantity gets UnderlyingAdjustedQuantity, Tag 1044
func (m NoUnderlyings) GetUnderlyingAdjustedQuantity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAdjustedQuantityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingFXRate gets UnderlyingFXRate, Tag 1045
func (m NoUnderlyings) GetUnderlyingFXRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingFXRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingFXRateCalc gets UnderlyingFXRateCalc, Tag 1046
func (m NoUnderlyings) GetUnderlyingFXRateCalc() (v enum.UnderlyingFXRateCalc, err quickfix.MessageRejectError) {
	var f field.UnderlyingFXRateCalcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
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
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityAltIDSource gets UnderlyingSecurityAltIDSource, Tag 459
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
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
func (m NoUnderlyingStips) GetUnderlyingStipType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStipTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStipValue gets UnderlyingStipValue, Tag 889
func (m NoUnderlyingStips) GetUnderlyingStipValue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStipValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
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
func (m NoUndlyInstrumentParties) GetUndlyInstrumentPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.UndlyInstrumentPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUndlyInstrumentPartyIDSource gets UndlyInstrumentPartyIDSource, Tag 1060
func (m NoUndlyInstrumentParties) GetUndlyInstrumentPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UndlyInstrumentPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUndlyInstrumentPartyRole gets UndlyInstrumentPartyRole, Tag 1061
func (m NoUndlyInstrumentParties) GetUndlyInstrumentPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.UndlyInstrumentPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
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
func (m NoUndlyInstrumentPartySubIDs) GetUndlyInstrumentPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.UndlyInstrumentPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUndlyInstrumentPartySubIDType gets UndlyInstrumentPartySubIDType, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) GetUndlyInstrumentPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.UndlyInstrumentPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
}

//SetEventType sets EventType, Tag 865
func (m NoEvents) SetEventType(v enum.EventType) {
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
func (m NoEvents) GetEventType() (v enum.EventType, err quickfix.MessageRejectError) {
	var f field.EventTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEventDate gets EventDate, Tag 866
func (m NoEvents) GetEventDate() (v string, err quickfix.MessageRejectError) {
	var f field.EventDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEventPx gets EventPx, Tag 867
func (m NoEvents) GetEventPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.EventPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEventText gets EventText, Tag 868
func (m NoEvents) GetEventText() (v string, err quickfix.MessageRejectError) {
	var f field.EventTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
}

//SetStrategyParameterName sets StrategyParameterName, Tag 958
func (m NoStrategyParameters) SetStrategyParameterName(v string) {
	m.Set(field.NewStrategyParameterName(v))
}

//SetStrategyParameterType sets StrategyParameterType, Tag 959
func (m NoStrategyParameters) SetStrategyParameterType(v enum.StrategyParameterType) {
	m.Set(field.NewStrategyParameterType(v))
}

//SetStrategyParameterValue sets StrategyParameterValue, Tag 960
func (m NoStrategyParameters) SetStrategyParameterValue(v string) {
	m.Set(field.NewStrategyParameterValue(v))
}

//GetStrategyParameterName gets StrategyParameterName, Tag 958
func (m NoStrategyParameters) GetStrategyParameterName() (v string, err quickfix.MessageRejectError) {
	var f field.StrategyParameterNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrategyParameterType gets StrategyParameterType, Tag 959
func (m NoStrategyParameters) GetStrategyParameterType() (v enum.StrategyParameterType, err quickfix.MessageRejectError) {
	var f field.StrategyParameterTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrategyParameterValue gets StrategyParameterValue, Tag 960
func (m NoStrategyParameters) GetStrategyParameterValue() (v string, err quickfix.MessageRejectError) {
	var f field.StrategyParameterValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
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
func (m NoInstrumentParties) GetInstrumentPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrumentPartyIDSource gets InstrumentPartyIDSource, Tag 1050
func (m NoInstrumentParties) GetInstrumentPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrumentPartyRole gets InstrumentPartyRole, Tag 1051
func (m NoInstrumentParties) GetInstrumentPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
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
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrumentPartySubIDType gets InstrumentPartySubIDType, Tag 1054
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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

//NoRootPartyIDs is a repeating group element, Tag 1116
type NoRootPartyIDs struct {
	*quickfix.Group
}

//SetRootPartyID sets RootPartyID, Tag 1117
func (m NoRootPartyIDs) SetRootPartyID(v string) {
	m.Set(field.NewRootPartyID(v))
}

//SetRootPartyIDSource sets RootPartyIDSource, Tag 1118
func (m NoRootPartyIDs) SetRootPartyIDSource(v string) {
	m.Set(field.NewRootPartyIDSource(v))
}

//SetRootPartyRole sets RootPartyRole, Tag 1119
func (m NoRootPartyIDs) SetRootPartyRole(v int) {
	m.Set(field.NewRootPartyRole(v))
}

//SetNoRootPartySubIDs sets NoRootPartySubIDs, Tag 1120
func (m NoRootPartyIDs) SetNoRootPartySubIDs(f NoRootPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetRootPartyID gets RootPartyID, Tag 1117
func (m NoRootPartyIDs) GetRootPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.RootPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRootPartyIDSource gets RootPartyIDSource, Tag 1118
func (m NoRootPartyIDs) GetRootPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.RootPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRootPartyRole gets RootPartyRole, Tag 1119
func (m NoRootPartyIDs) GetRootPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.RootPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoRootPartySubIDs gets NoRootPartySubIDs, Tag 1120
func (m NoRootPartyIDs) GetNoRootPartySubIDs() (NoRootPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRootPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasRootPartyID returns true if RootPartyID is present, Tag 1117
func (m NoRootPartyIDs) HasRootPartyID() bool {
	return m.Has(tag.RootPartyID)
}

//HasRootPartyIDSource returns true if RootPartyIDSource is present, Tag 1118
func (m NoRootPartyIDs) HasRootPartyIDSource() bool {
	return m.Has(tag.RootPartyIDSource)
}

//HasRootPartyRole returns true if RootPartyRole is present, Tag 1119
func (m NoRootPartyIDs) HasRootPartyRole() bool {
	return m.Has(tag.RootPartyRole)
}

//HasNoRootPartySubIDs returns true if NoRootPartySubIDs is present, Tag 1120
func (m NoRootPartyIDs) HasNoRootPartySubIDs() bool {
	return m.Has(tag.NoRootPartySubIDs)
}

//NoRootPartySubIDs is a repeating group element, Tag 1120
type NoRootPartySubIDs struct {
	*quickfix.Group
}

//SetRootPartySubID sets RootPartySubID, Tag 1121
func (m NoRootPartySubIDs) SetRootPartySubID(v string) {
	m.Set(field.NewRootPartySubID(v))
}

//SetRootPartySubIDType sets RootPartySubIDType, Tag 1122
func (m NoRootPartySubIDs) SetRootPartySubIDType(v int) {
	m.Set(field.NewRootPartySubIDType(v))
}

//GetRootPartySubID gets RootPartySubID, Tag 1121
func (m NoRootPartySubIDs) GetRootPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.RootPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRootPartySubIDType gets RootPartySubIDType, Tag 1122
func (m NoRootPartySubIDs) GetRootPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.RootPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasRootPartySubID returns true if RootPartySubID is present, Tag 1121
func (m NoRootPartySubIDs) HasRootPartySubID() bool {
	return m.Has(tag.RootPartySubID)
}

//HasRootPartySubIDType returns true if RootPartySubIDType is present, Tag 1122
func (m NoRootPartySubIDs) HasRootPartySubIDType() bool {
	return m.Has(tag.RootPartySubIDType)
}

//NoRootPartySubIDsRepeatingGroup is a repeating group, Tag 1120
type NoRootPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRootPartySubIDsRepeatingGroup returns an initialized, NoRootPartySubIDsRepeatingGroup
func NewNoRootPartySubIDsRepeatingGroup() NoRootPartySubIDsRepeatingGroup {
	return NoRootPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRootPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RootPartySubID), quickfix.GroupElement(tag.RootPartySubIDType)})}
}

//Add create and append a new NoRootPartySubIDs to this group
func (m NoRootPartySubIDsRepeatingGroup) Add() NoRootPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoRootPartySubIDs{g}
}

//Get returns the ith NoRootPartySubIDs in the NoRootPartySubIDsRepeatinGroup
func (m NoRootPartySubIDsRepeatingGroup) Get(i int) NoRootPartySubIDs {
	return NoRootPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoRootPartyIDsRepeatingGroup is a repeating group, Tag 1116
type NoRootPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRootPartyIDsRepeatingGroup returns an initialized, NoRootPartyIDsRepeatingGroup
func NewNoRootPartyIDsRepeatingGroup() NoRootPartyIDsRepeatingGroup {
	return NoRootPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRootPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RootPartyID), quickfix.GroupElement(tag.RootPartyIDSource), quickfix.GroupElement(tag.RootPartyRole), NewNoRootPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoRootPartyIDs to this group
func (m NoRootPartyIDsRepeatingGroup) Add() NoRootPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoRootPartyIDs{g}
}

//Get returns the ith NoRootPartyIDs in the NoRootPartyIDsRepeatinGroup
func (m NoRootPartyIDsRepeatingGroup) Get(i int) NoRootPartyIDs {
	return NoRootPartyIDs{m.RepeatingGroup.Get(i)}
}
