package tradecapturereport

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/tag"
)

//TradeCaptureReport is the fix44 TradeCaptureReport type, MsgType = AE
type TradeCaptureReport struct {
	fix44.Header
	quickfix.Body
	fix44.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a TradeCaptureReport from a quickfix.Message instance
func FromMessage(m quickfix.Message) TradeCaptureReport {
	return TradeCaptureReport{
		Header:      fix44.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix44.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m TradeCaptureReport) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a TradeCaptureReport initialized with the required fields for TradeCaptureReport
func New(tradereportid field.TradeReportIDField, previouslyreported field.PreviouslyReportedField, lastqty field.LastQtyField, lastpx field.LastPxField, tradedate field.TradeDateField, transacttime field.TransactTimeField) (m TradeCaptureReport) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("AE"))
	m.Set(tradereportid)
	m.Set(previouslyreported)
	m.Set(lastqty)
	m.Set(lastpx)
	m.Set(tradedate)
	m.Set(transacttime)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg TradeCaptureReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "AE", r
}

//SetAvgPx sets AvgPx, Tag 6
func (m TradeCaptureReport) SetAvgPx(v float64) {
	m.Set(field.NewAvgPx(v))
}

//SetExecID sets ExecID, Tag 17
func (m TradeCaptureReport) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m TradeCaptureReport) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetLastMkt sets LastMkt, Tag 30
func (m TradeCaptureReport) SetLastMkt(v string) {
	m.Set(field.NewLastMkt(v))
}

//SetLastPx sets LastPx, Tag 31
func (m TradeCaptureReport) SetLastPx(v float64) {
	m.Set(field.NewLastPx(v))
}

//SetLastQty sets LastQty, Tag 32
func (m TradeCaptureReport) SetLastQty(v float64) {
	m.Set(field.NewLastQty(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m TradeCaptureReport) SetOrderQty(v float64) {
	m.Set(field.NewOrderQty(v))
}

//SetOrdStatus sets OrdStatus, Tag 39
func (m TradeCaptureReport) SetOrdStatus(v string) {
	m.Set(field.NewOrdStatus(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m TradeCaptureReport) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m TradeCaptureReport) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m TradeCaptureReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSettlType sets SettlType, Tag 63
func (m TradeCaptureReport) SetSettlType(v string) {
	m.Set(field.NewSettlType(v))
}

//SetSettlDate sets SettlDate, Tag 64
func (m TradeCaptureReport) SetSettlDate(v string) {
	m.Set(field.NewSettlDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m TradeCaptureReport) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m TradeCaptureReport) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetIssuer sets Issuer, Tag 106
func (m TradeCaptureReport) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m TradeCaptureReport) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetExecType sets ExecType, Tag 150
func (m TradeCaptureReport) SetExecType(v string) {
	m.Set(field.NewExecType(v))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m TradeCaptureReport) SetCashOrderQty(v float64) {
	m.Set(field.NewCashOrderQty(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m TradeCaptureReport) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetLastSpotRate sets LastSpotRate, Tag 194
func (m TradeCaptureReport) SetLastSpotRate(v float64) {
	m.Set(field.NewLastSpotRate(v))
}

//SetLastForwardPoints sets LastForwardPoints, Tag 195
func (m TradeCaptureReport) SetLastForwardPoints(v float64) {
	m.Set(field.NewLastForwardPoints(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m TradeCaptureReport) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m TradeCaptureReport) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m TradeCaptureReport) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m TradeCaptureReport) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetSpread sets Spread, Tag 218
func (m TradeCaptureReport) SetSpread(v float64) {
	m.Set(field.NewSpread(v))
}

//SetBenchmarkCurveCurrency sets BenchmarkCurveCurrency, Tag 220
func (m TradeCaptureReport) SetBenchmarkCurveCurrency(v string) {
	m.Set(field.NewBenchmarkCurveCurrency(v))
}

//SetBenchmarkCurveName sets BenchmarkCurveName, Tag 221
func (m TradeCaptureReport) SetBenchmarkCurveName(v string) {
	m.Set(field.NewBenchmarkCurveName(v))
}

//SetBenchmarkCurvePoint sets BenchmarkCurvePoint, Tag 222
func (m TradeCaptureReport) SetBenchmarkCurvePoint(v string) {
	m.Set(field.NewBenchmarkCurvePoint(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m TradeCaptureReport) SetCouponRate(v float64) {
	m.Set(field.NewCouponRate(v))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m TradeCaptureReport) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m TradeCaptureReport) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m TradeCaptureReport) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m TradeCaptureReport) SetRepurchaseRate(v float64) {
	m.Set(field.NewRepurchaseRate(v))
}

//SetFactor sets Factor, Tag 228
func (m TradeCaptureReport) SetFactor(v float64) {
	m.Set(field.NewFactor(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m TradeCaptureReport) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
}

//SetYieldType sets YieldType, Tag 235
func (m TradeCaptureReport) SetYieldType(v string) {
	m.Set(field.NewYieldType(v))
}

//SetYield sets Yield, Tag 236
func (m TradeCaptureReport) SetYield(v float64) {
	m.Set(field.NewYield(v))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m TradeCaptureReport) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m TradeCaptureReport) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m TradeCaptureReport) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263
func (m TradeCaptureReport) SetSubscriptionRequestType(v string) {
	m.Set(field.NewSubscriptionRequestType(v))
}

//SetUnsolicitedIndicator sets UnsolicitedIndicator, Tag 325
func (m TradeCaptureReport) SetUnsolicitedIndicator(v bool) {
	m.Set(field.NewUnsolicitedIndicator(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m TradeCaptureReport) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m TradeCaptureReport) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m TradeCaptureReport) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m TradeCaptureReport) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetExecRestatementReason sets ExecRestatementReason, Tag 378
func (m TradeCaptureReport) SetExecRestatementReason(v int) {
	m.Set(field.NewExecRestatementReason(v))
}

//SetPriceType sets PriceType, Tag 423
func (m TradeCaptureReport) SetPriceType(v int) {
	m.Set(field.NewPriceType(v))
}

//SetMultiLegReportingType sets MultiLegReportingType, Tag 442
func (m TradeCaptureReport) SetMultiLegReportingType(v string) {
	m.Set(field.NewMultiLegReportingType(v))
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m TradeCaptureReport) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m TradeCaptureReport) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m TradeCaptureReport) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetRoundingDirection sets RoundingDirection, Tag 468
func (m TradeCaptureReport) SetRoundingDirection(v string) {
	m.Set(field.NewRoundingDirection(v))
}

//SetRoundingModulus sets RoundingModulus, Tag 469
func (m TradeCaptureReport) SetRoundingModulus(v float64) {
	m.Set(field.NewRoundingModulus(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m TradeCaptureReport) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m TradeCaptureReport) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m TradeCaptureReport) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetTradeReportTransType sets TradeReportTransType, Tag 487
func (m TradeCaptureReport) SetTradeReportTransType(v int) {
	m.Set(field.NewTradeReportTransType(v))
}

//SetOrderPercent sets OrderPercent, Tag 516
func (m TradeCaptureReport) SetOrderPercent(v float64) {
	m.Set(field.NewOrderPercent(v))
}

//SetSecondaryExecID sets SecondaryExecID, Tag 527
func (m TradeCaptureReport) SetSecondaryExecID(v string) {
	m.Set(field.NewSecondaryExecID(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m TradeCaptureReport) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m TradeCaptureReport) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetNoSides sets NoSides, Tag 552
func (m TradeCaptureReport) SetNoSides(f NoSidesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoLegs sets NoLegs, Tag 555
func (m TradeCaptureReport) SetNoLegs(f NoLegsRepeatingGroup) {
	m.SetGroup(f)
}

//SetTradeRequestID sets TradeRequestID, Tag 568
func (m TradeCaptureReport) SetTradeRequestID(v string) {
	m.Set(field.NewTradeRequestID(v))
}

//SetPreviouslyReported sets PreviouslyReported, Tag 570
func (m TradeCaptureReport) SetPreviouslyReported(v bool) {
	m.Set(field.NewPreviouslyReported(v))
}

//SetTradeReportID sets TradeReportID, Tag 571
func (m TradeCaptureReport) SetTradeReportID(v string) {
	m.Set(field.NewTradeReportID(v))
}

//SetTradeReportRefID sets TradeReportRefID, Tag 572
func (m TradeCaptureReport) SetTradeReportRefID(v string) {
	m.Set(field.NewTradeReportRefID(v))
}

//SetMatchStatus sets MatchStatus, Tag 573
func (m TradeCaptureReport) SetMatchStatus(v string) {
	m.Set(field.NewMatchStatus(v))
}

//SetMatchType sets MatchType, Tag 574
func (m TradeCaptureReport) SetMatchType(v string) {
	m.Set(field.NewMatchType(v))
}

//SetBenchmarkPrice sets BenchmarkPrice, Tag 662
func (m TradeCaptureReport) SetBenchmarkPrice(v float64) {
	m.Set(field.NewBenchmarkPrice(v))
}

//SetBenchmarkPriceType sets BenchmarkPriceType, Tag 663
func (m TradeCaptureReport) SetBenchmarkPriceType(v int) {
	m.Set(field.NewBenchmarkPriceType(v))
}

//SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m TradeCaptureReport) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

//SetLastParPx sets LastParPx, Tag 669
func (m TradeCaptureReport) SetLastParPx(v float64) {
	m.Set(field.NewLastParPx(v))
}

//SetPool sets Pool, Tag 691
func (m TradeCaptureReport) SetPool(v string) {
	m.Set(field.NewPool(v))
}

//SetYieldRedemptionDate sets YieldRedemptionDate, Tag 696
func (m TradeCaptureReport) SetYieldRedemptionDate(v string) {
	m.Set(field.NewYieldRedemptionDate(v))
}

//SetYieldRedemptionPrice sets YieldRedemptionPrice, Tag 697
func (m TradeCaptureReport) SetYieldRedemptionPrice(v float64) {
	m.Set(field.NewYieldRedemptionPrice(v))
}

//SetYieldRedemptionPriceType sets YieldRedemptionPriceType, Tag 698
func (m TradeCaptureReport) SetYieldRedemptionPriceType(v int) {
	m.Set(field.NewYieldRedemptionPriceType(v))
}

//SetBenchmarkSecurityID sets BenchmarkSecurityID, Tag 699
func (m TradeCaptureReport) SetBenchmarkSecurityID(v string) {
	m.Set(field.NewBenchmarkSecurityID(v))
}

//SetYieldCalcDate sets YieldCalcDate, Tag 701
func (m TradeCaptureReport) SetYieldCalcDate(v string) {
	m.Set(field.NewYieldCalcDate(v))
}

//SetNoUnderlyings sets NoUnderlyings, Tag 711
func (m TradeCaptureReport) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

//SetClearingBusinessDate sets ClearingBusinessDate, Tag 715
func (m TradeCaptureReport) SetClearingBusinessDate(v string) {
	m.Set(field.NewClearingBusinessDate(v))
}

//SetTotNumTradeReports sets TotNumTradeReports, Tag 748
func (m TradeCaptureReport) SetTotNumTradeReports(v int) {
	m.Set(field.NewTotNumTradeReports(v))
}

//SetNoPosAmt sets NoPosAmt, Tag 753
func (m TradeCaptureReport) SetNoPosAmt(f NoPosAmtRepeatingGroup) {
	m.SetGroup(f)
}

//SetBenchmarkSecurityIDSource sets BenchmarkSecurityIDSource, Tag 761
func (m TradeCaptureReport) SetBenchmarkSecurityIDSource(v string) {
	m.Set(field.NewBenchmarkSecurityIDSource(v))
}

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m TradeCaptureReport) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

//SetNoTrdRegTimestamps sets NoTrdRegTimestamps, Tag 768
func (m TradeCaptureReport) SetNoTrdRegTimestamps(f NoTrdRegTimestampsRepeatingGroup) {
	m.SetGroup(f)
}

//SetTerminationType sets TerminationType, Tag 788
func (m TradeCaptureReport) SetTerminationType(v int) {
	m.Set(field.NewTerminationType(v))
}

//SetSecondaryTradeReportID sets SecondaryTradeReportID, Tag 818
func (m TradeCaptureReport) SetSecondaryTradeReportID(v string) {
	m.Set(field.NewSecondaryTradeReportID(v))
}

//SetAvgPxIndicator sets AvgPxIndicator, Tag 819
func (m TradeCaptureReport) SetAvgPxIndicator(v int) {
	m.Set(field.NewAvgPxIndicator(v))
}

//SetTradeLinkID sets TradeLinkID, Tag 820
func (m TradeCaptureReport) SetTradeLinkID(v string) {
	m.Set(field.NewTradeLinkID(v))
}

//SetUnderlyingTradingSessionID sets UnderlyingTradingSessionID, Tag 822
func (m TradeCaptureReport) SetUnderlyingTradingSessionID(v string) {
	m.Set(field.NewUnderlyingTradingSessionID(v))
}

//SetUnderlyingTradingSessionSubID sets UnderlyingTradingSessionSubID, Tag 823
func (m TradeCaptureReport) SetUnderlyingTradingSessionSubID(v string) {
	m.Set(field.NewUnderlyingTradingSessionSubID(v))
}

//SetTradeLegRefID sets TradeLegRefID, Tag 824
func (m TradeCaptureReport) SetTradeLegRefID(v string) {
	m.Set(field.NewTradeLegRefID(v))
}

//SetTrdType sets TrdType, Tag 828
func (m TradeCaptureReport) SetTrdType(v int) {
	m.Set(field.NewTrdType(v))
}

//SetTrdSubType sets TrdSubType, Tag 829
func (m TradeCaptureReport) SetTrdSubType(v int) {
	m.Set(field.NewTrdSubType(v))
}

//SetTransferReason sets TransferReason, Tag 830
func (m TradeCaptureReport) SetTransferReason(v string) {
	m.Set(field.NewTransferReason(v))
}

//SetQtyType sets QtyType, Tag 854
func (m TradeCaptureReport) SetQtyType(v int) {
	m.Set(field.NewQtyType(v))
}

//SetSecondaryTrdType sets SecondaryTrdType, Tag 855
func (m TradeCaptureReport) SetSecondaryTrdType(v int) {
	m.Set(field.NewSecondaryTrdType(v))
}

//SetTradeReportType sets TradeReportType, Tag 856
func (m TradeCaptureReport) SetTradeReportType(v int) {
	m.Set(field.NewTradeReportType(v))
}

//SetNoEvents sets NoEvents, Tag 864
func (m TradeCaptureReport) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDatedDate sets DatedDate, Tag 873
func (m TradeCaptureReport) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

//SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m TradeCaptureReport) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

//SetCPProgram sets CPProgram, Tag 875
func (m TradeCaptureReport) SetCPProgram(v int) {
	m.Set(field.NewCPProgram(v))
}

//SetCPRegType sets CPRegType, Tag 876
func (m TradeCaptureReport) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

//SetTrdMatchID sets TrdMatchID, Tag 880
func (m TradeCaptureReport) SetTrdMatchID(v string) {
	m.Set(field.NewTrdMatchID(v))
}

//SetSecondaryTradeReportRefID sets SecondaryTradeReportRefID, Tag 881
func (m TradeCaptureReport) SetSecondaryTradeReportRefID(v string) {
	m.Set(field.NewSecondaryTradeReportRefID(v))
}

//SetMarginRatio sets MarginRatio, Tag 898
func (m TradeCaptureReport) SetMarginRatio(v float64) {
	m.Set(field.NewMarginRatio(v))
}

//SetLastRptRequested sets LastRptRequested, Tag 912
func (m TradeCaptureReport) SetLastRptRequested(v bool) {
	m.Set(field.NewLastRptRequested(v))
}

//SetAgreementDesc sets AgreementDesc, Tag 913
func (m TradeCaptureReport) SetAgreementDesc(v string) {
	m.Set(field.NewAgreementDesc(v))
}

//SetAgreementID sets AgreementID, Tag 914
func (m TradeCaptureReport) SetAgreementID(v string) {
	m.Set(field.NewAgreementID(v))
}

//SetAgreementDate sets AgreementDate, Tag 915
func (m TradeCaptureReport) SetAgreementDate(v string) {
	m.Set(field.NewAgreementDate(v))
}

//SetStartDate sets StartDate, Tag 916
func (m TradeCaptureReport) SetStartDate(v string) {
	m.Set(field.NewStartDate(v))
}

//SetEndDate sets EndDate, Tag 917
func (m TradeCaptureReport) SetEndDate(v string) {
	m.Set(field.NewEndDate(v))
}

//SetAgreementCurrency sets AgreementCurrency, Tag 918
func (m TradeCaptureReport) SetAgreementCurrency(v string) {
	m.Set(field.NewAgreementCurrency(v))
}

//SetDeliveryType sets DeliveryType, Tag 919
func (m TradeCaptureReport) SetDeliveryType(v int) {
	m.Set(field.NewDeliveryType(v))
}

//SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m TradeCaptureReport) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

//GetAvgPx gets AvgPx, Tag 6
func (m TradeCaptureReport) GetAvgPx() (f field.AvgPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecID gets ExecID, Tag 17
func (m TradeCaptureReport) GetExecID() (f field.ExecIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m TradeCaptureReport) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastMkt gets LastMkt, Tag 30
func (m TradeCaptureReport) GetLastMkt() (f field.LastMktField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastPx gets LastPx, Tag 31
func (m TradeCaptureReport) GetLastPx() (f field.LastPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastQty gets LastQty, Tag 32
func (m TradeCaptureReport) GetLastQty() (f field.LastQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m TradeCaptureReport) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdStatus gets OrdStatus, Tag 39
func (m TradeCaptureReport) GetOrdStatus() (f field.OrdStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m TradeCaptureReport) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m TradeCaptureReport) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m TradeCaptureReport) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlType gets SettlType, Tag 63
func (m TradeCaptureReport) GetSettlType() (f field.SettlTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlDate gets SettlDate, Tag 64
func (m TradeCaptureReport) GetSettlDate() (f field.SettlDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m TradeCaptureReport) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m TradeCaptureReport) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m TradeCaptureReport) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m TradeCaptureReport) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecType gets ExecType, Tag 150
func (m TradeCaptureReport) GetExecType() (f field.ExecTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m TradeCaptureReport) GetCashOrderQty() (f field.CashOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m TradeCaptureReport) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastSpotRate gets LastSpotRate, Tag 194
func (m TradeCaptureReport) GetLastSpotRate() (f field.LastSpotRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastForwardPoints gets LastForwardPoints, Tag 195
func (m TradeCaptureReport) GetLastForwardPoints() (f field.LastForwardPointsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m TradeCaptureReport) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m TradeCaptureReport) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m TradeCaptureReport) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m TradeCaptureReport) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSpread gets Spread, Tag 218
func (m TradeCaptureReport) GetSpread() (f field.SpreadField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveCurrency gets BenchmarkCurveCurrency, Tag 220
func (m TradeCaptureReport) GetBenchmarkCurveCurrency() (f field.BenchmarkCurveCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveName gets BenchmarkCurveName, Tag 221
func (m TradeCaptureReport) GetBenchmarkCurveName() (f field.BenchmarkCurveNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurvePoint gets BenchmarkCurvePoint, Tag 222
func (m TradeCaptureReport) GetBenchmarkCurvePoint() (f field.BenchmarkCurvePointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m TradeCaptureReport) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m TradeCaptureReport) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m TradeCaptureReport) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m TradeCaptureReport) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m TradeCaptureReport) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m TradeCaptureReport) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m TradeCaptureReport) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldType gets YieldType, Tag 235
func (m TradeCaptureReport) GetYieldType() (f field.YieldTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYield gets Yield, Tag 236
func (m TradeCaptureReport) GetYield() (f field.YieldField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m TradeCaptureReport) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m TradeCaptureReport) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m TradeCaptureReport) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263
func (m TradeCaptureReport) GetSubscriptionRequestType() (f field.SubscriptionRequestTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnsolicitedIndicator gets UnsolicitedIndicator, Tag 325
func (m TradeCaptureReport) GetUnsolicitedIndicator() (f field.UnsolicitedIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m TradeCaptureReport) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m TradeCaptureReport) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m TradeCaptureReport) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m TradeCaptureReport) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecRestatementReason gets ExecRestatementReason, Tag 378
func (m TradeCaptureReport) GetExecRestatementReason() (f field.ExecRestatementReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceType gets PriceType, Tag 423
func (m TradeCaptureReport) GetPriceType() (f field.PriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMultiLegReportingType gets MultiLegReportingType, Tag 442
func (m TradeCaptureReport) GetMultiLegReportingType() (f field.MultiLegReportingTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m TradeCaptureReport) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m TradeCaptureReport) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m TradeCaptureReport) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingDirection gets RoundingDirection, Tag 468
func (m TradeCaptureReport) GetRoundingDirection() (f field.RoundingDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingModulus gets RoundingModulus, Tag 469
func (m TradeCaptureReport) GetRoundingModulus() (f field.RoundingModulusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m TradeCaptureReport) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m TradeCaptureReport) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m TradeCaptureReport) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeReportTransType gets TradeReportTransType, Tag 487
func (m TradeCaptureReport) GetTradeReportTransType() (f field.TradeReportTransTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderPercent gets OrderPercent, Tag 516
func (m TradeCaptureReport) GetOrderPercent() (f field.OrderPercentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryExecID gets SecondaryExecID, Tag 527
func (m TradeCaptureReport) GetSecondaryExecID() (f field.SecondaryExecIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m TradeCaptureReport) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m TradeCaptureReport) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSides gets NoSides, Tag 552
func (m TradeCaptureReport) GetNoSides() (NoSidesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSidesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoLegs gets NoLegs, Tag 555
func (m TradeCaptureReport) GetNoLegs() (NoLegsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetTradeRequestID gets TradeRequestID, Tag 568
func (m TradeCaptureReport) GetTradeRequestID() (f field.TradeRequestIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPreviouslyReported gets PreviouslyReported, Tag 570
func (m TradeCaptureReport) GetPreviouslyReported() (f field.PreviouslyReportedField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeReportID gets TradeReportID, Tag 571
func (m TradeCaptureReport) GetTradeReportID() (f field.TradeReportIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeReportRefID gets TradeReportRefID, Tag 572
func (m TradeCaptureReport) GetTradeReportRefID() (f field.TradeReportRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMatchStatus gets MatchStatus, Tag 573
func (m TradeCaptureReport) GetMatchStatus() (f field.MatchStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMatchType gets MatchType, Tag 574
func (m TradeCaptureReport) GetMatchType() (f field.MatchTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkPrice gets BenchmarkPrice, Tag 662
func (m TradeCaptureReport) GetBenchmarkPrice() (f field.BenchmarkPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkPriceType gets BenchmarkPriceType, Tag 663
func (m TradeCaptureReport) GetBenchmarkPriceType() (f field.BenchmarkPriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m TradeCaptureReport) GetContractSettlMonth() (f field.ContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastParPx gets LastParPx, Tag 669
func (m TradeCaptureReport) GetLastParPx() (f field.LastParPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPool gets Pool, Tag 691
func (m TradeCaptureReport) GetPool() (f field.PoolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldRedemptionDate gets YieldRedemptionDate, Tag 696
func (m TradeCaptureReport) GetYieldRedemptionDate() (f field.YieldRedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldRedemptionPrice gets YieldRedemptionPrice, Tag 697
func (m TradeCaptureReport) GetYieldRedemptionPrice() (f field.YieldRedemptionPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldRedemptionPriceType gets YieldRedemptionPriceType, Tag 698
func (m TradeCaptureReport) GetYieldRedemptionPriceType() (f field.YieldRedemptionPriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkSecurityID gets BenchmarkSecurityID, Tag 699
func (m TradeCaptureReport) GetBenchmarkSecurityID() (f field.BenchmarkSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldCalcDate gets YieldCalcDate, Tag 701
func (m TradeCaptureReport) GetYieldCalcDate() (f field.YieldCalcDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyings gets NoUnderlyings, Tag 711
func (m TradeCaptureReport) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetClearingBusinessDate gets ClearingBusinessDate, Tag 715
func (m TradeCaptureReport) GetClearingBusinessDate() (f field.ClearingBusinessDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTotNumTradeReports gets TotNumTradeReports, Tag 748
func (m TradeCaptureReport) GetTotNumTradeReports() (f field.TotNumTradeReportsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPosAmt gets NoPosAmt, Tag 753
func (m TradeCaptureReport) GetNoPosAmt() (NoPosAmtRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPosAmtRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetBenchmarkSecurityIDSource gets BenchmarkSecurityIDSource, Tag 761
func (m TradeCaptureReport) GetBenchmarkSecurityIDSource() (f field.BenchmarkSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m TradeCaptureReport) GetSecuritySubType() (f field.SecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoTrdRegTimestamps gets NoTrdRegTimestamps, Tag 768
func (m TradeCaptureReport) GetNoTrdRegTimestamps() (NoTrdRegTimestampsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTrdRegTimestampsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetTerminationType gets TerminationType, Tag 788
func (m TradeCaptureReport) GetTerminationType() (f field.TerminationTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryTradeReportID gets SecondaryTradeReportID, Tag 818
func (m TradeCaptureReport) GetSecondaryTradeReportID() (f field.SecondaryTradeReportIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAvgPxIndicator gets AvgPxIndicator, Tag 819
func (m TradeCaptureReport) GetAvgPxIndicator() (f field.AvgPxIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeLinkID gets TradeLinkID, Tag 820
func (m TradeCaptureReport) GetTradeLinkID() (f field.TradeLinkIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingTradingSessionID gets UnderlyingTradingSessionID, Tag 822
func (m TradeCaptureReport) GetUnderlyingTradingSessionID() (f field.UnderlyingTradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingTradingSessionSubID gets UnderlyingTradingSessionSubID, Tag 823
func (m TradeCaptureReport) GetUnderlyingTradingSessionSubID() (f field.UnderlyingTradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeLegRefID gets TradeLegRefID, Tag 824
func (m TradeCaptureReport) GetTradeLegRefID() (f field.TradeLegRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTrdType gets TrdType, Tag 828
func (m TradeCaptureReport) GetTrdType() (f field.TrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTrdSubType gets TrdSubType, Tag 829
func (m TradeCaptureReport) GetTrdSubType() (f field.TrdSubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransferReason gets TransferReason, Tag 830
func (m TradeCaptureReport) GetTransferReason() (f field.TransferReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQtyType gets QtyType, Tag 854
func (m TradeCaptureReport) GetQtyType() (f field.QtyTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryTrdType gets SecondaryTrdType, Tag 855
func (m TradeCaptureReport) GetSecondaryTrdType() (f field.SecondaryTrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeReportType gets TradeReportType, Tag 856
func (m TradeCaptureReport) GetTradeReportType() (f field.TradeReportTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoEvents gets NoEvents, Tag 864
func (m TradeCaptureReport) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDatedDate gets DatedDate, Tag 873
func (m TradeCaptureReport) GetDatedDate() (f field.DatedDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m TradeCaptureReport) GetInterestAccrualDate() (f field.InterestAccrualDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPProgram gets CPProgram, Tag 875
func (m TradeCaptureReport) GetCPProgram() (f field.CPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPRegType gets CPRegType, Tag 876
func (m TradeCaptureReport) GetCPRegType() (f field.CPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTrdMatchID gets TrdMatchID, Tag 880
func (m TradeCaptureReport) GetTrdMatchID() (f field.TrdMatchIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryTradeReportRefID gets SecondaryTradeReportRefID, Tag 881
func (m TradeCaptureReport) GetSecondaryTradeReportRefID() (f field.SecondaryTradeReportRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMarginRatio gets MarginRatio, Tag 898
func (m TradeCaptureReport) GetMarginRatio() (f field.MarginRatioField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastRptRequested gets LastRptRequested, Tag 912
func (m TradeCaptureReport) GetLastRptRequested() (f field.LastRptRequestedField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementDesc gets AgreementDesc, Tag 913
func (m TradeCaptureReport) GetAgreementDesc() (f field.AgreementDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementID gets AgreementID, Tag 914
func (m TradeCaptureReport) GetAgreementID() (f field.AgreementIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementDate gets AgreementDate, Tag 915
func (m TradeCaptureReport) GetAgreementDate() (f field.AgreementDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStartDate gets StartDate, Tag 916
func (m TradeCaptureReport) GetStartDate() (f field.StartDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEndDate gets EndDate, Tag 917
func (m TradeCaptureReport) GetEndDate() (f field.EndDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementCurrency gets AgreementCurrency, Tag 918
func (m TradeCaptureReport) GetAgreementCurrency() (f field.AgreementCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDeliveryType gets DeliveryType, Tag 919
func (m TradeCaptureReport) GetDeliveryType() (f field.DeliveryTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m TradeCaptureReport) GetStrikeCurrency() (f field.StrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAvgPx returns true if AvgPx is present, Tag 6
func (m TradeCaptureReport) HasAvgPx() bool {
	return m.Has(tag.AvgPx)
}

//HasExecID returns true if ExecID is present, Tag 17
func (m TradeCaptureReport) HasExecID() bool {
	return m.Has(tag.ExecID)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m TradeCaptureReport) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasLastMkt returns true if LastMkt is present, Tag 30
func (m TradeCaptureReport) HasLastMkt() bool {
	return m.Has(tag.LastMkt)
}

//HasLastPx returns true if LastPx is present, Tag 31
func (m TradeCaptureReport) HasLastPx() bool {
	return m.Has(tag.LastPx)
}

//HasLastQty returns true if LastQty is present, Tag 32
func (m TradeCaptureReport) HasLastQty() bool {
	return m.Has(tag.LastQty)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m TradeCaptureReport) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasOrdStatus returns true if OrdStatus is present, Tag 39
func (m TradeCaptureReport) HasOrdStatus() bool {
	return m.Has(tag.OrdStatus)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m TradeCaptureReport) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m TradeCaptureReport) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m TradeCaptureReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSettlType returns true if SettlType is present, Tag 63
func (m TradeCaptureReport) HasSettlType() bool {
	return m.Has(tag.SettlType)
}

//HasSettlDate returns true if SettlDate is present, Tag 64
func (m TradeCaptureReport) HasSettlDate() bool {
	return m.Has(tag.SettlDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m TradeCaptureReport) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m TradeCaptureReport) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m TradeCaptureReport) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m TradeCaptureReport) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasExecType returns true if ExecType is present, Tag 150
func (m TradeCaptureReport) HasExecType() bool {
	return m.Has(tag.ExecType)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m TradeCaptureReport) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m TradeCaptureReport) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasLastSpotRate returns true if LastSpotRate is present, Tag 194
func (m TradeCaptureReport) HasLastSpotRate() bool {
	return m.Has(tag.LastSpotRate)
}

//HasLastForwardPoints returns true if LastForwardPoints is present, Tag 195
func (m TradeCaptureReport) HasLastForwardPoints() bool {
	return m.Has(tag.LastForwardPoints)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m TradeCaptureReport) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m TradeCaptureReport) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m TradeCaptureReport) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m TradeCaptureReport) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasSpread returns true if Spread is present, Tag 218
func (m TradeCaptureReport) HasSpread() bool {
	return m.Has(tag.Spread)
}

//HasBenchmarkCurveCurrency returns true if BenchmarkCurveCurrency is present, Tag 220
func (m TradeCaptureReport) HasBenchmarkCurveCurrency() bool {
	return m.Has(tag.BenchmarkCurveCurrency)
}

//HasBenchmarkCurveName returns true if BenchmarkCurveName is present, Tag 221
func (m TradeCaptureReport) HasBenchmarkCurveName() bool {
	return m.Has(tag.BenchmarkCurveName)
}

//HasBenchmarkCurvePoint returns true if BenchmarkCurvePoint is present, Tag 222
func (m TradeCaptureReport) HasBenchmarkCurvePoint() bool {
	return m.Has(tag.BenchmarkCurvePoint)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m TradeCaptureReport) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m TradeCaptureReport) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m TradeCaptureReport) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m TradeCaptureReport) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m TradeCaptureReport) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m TradeCaptureReport) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m TradeCaptureReport) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasYieldType returns true if YieldType is present, Tag 235
func (m TradeCaptureReport) HasYieldType() bool {
	return m.Has(tag.YieldType)
}

//HasYield returns true if Yield is present, Tag 236
func (m TradeCaptureReport) HasYield() bool {
	return m.Has(tag.Yield)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m TradeCaptureReport) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m TradeCaptureReport) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m TradeCaptureReport) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263
func (m TradeCaptureReport) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

//HasUnsolicitedIndicator returns true if UnsolicitedIndicator is present, Tag 325
func (m TradeCaptureReport) HasUnsolicitedIndicator() bool {
	return m.Has(tag.UnsolicitedIndicator)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m TradeCaptureReport) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m TradeCaptureReport) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m TradeCaptureReport) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m TradeCaptureReport) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasExecRestatementReason returns true if ExecRestatementReason is present, Tag 378
func (m TradeCaptureReport) HasExecRestatementReason() bool {
	return m.Has(tag.ExecRestatementReason)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m TradeCaptureReport) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasMultiLegReportingType returns true if MultiLegReportingType is present, Tag 442
func (m TradeCaptureReport) HasMultiLegReportingType() bool {
	return m.Has(tag.MultiLegReportingType)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m TradeCaptureReport) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m TradeCaptureReport) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m TradeCaptureReport) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasRoundingDirection returns true if RoundingDirection is present, Tag 468
func (m TradeCaptureReport) HasRoundingDirection() bool {
	return m.Has(tag.RoundingDirection)
}

//HasRoundingModulus returns true if RoundingModulus is present, Tag 469
func (m TradeCaptureReport) HasRoundingModulus() bool {
	return m.Has(tag.RoundingModulus)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m TradeCaptureReport) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m TradeCaptureReport) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m TradeCaptureReport) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasTradeReportTransType returns true if TradeReportTransType is present, Tag 487
func (m TradeCaptureReport) HasTradeReportTransType() bool {
	return m.Has(tag.TradeReportTransType)
}

//HasOrderPercent returns true if OrderPercent is present, Tag 516
func (m TradeCaptureReport) HasOrderPercent() bool {
	return m.Has(tag.OrderPercent)
}

//HasSecondaryExecID returns true if SecondaryExecID is present, Tag 527
func (m TradeCaptureReport) HasSecondaryExecID() bool {
	return m.Has(tag.SecondaryExecID)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m TradeCaptureReport) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m TradeCaptureReport) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasNoSides returns true if NoSides is present, Tag 552
func (m TradeCaptureReport) HasNoSides() bool {
	return m.Has(tag.NoSides)
}

//HasNoLegs returns true if NoLegs is present, Tag 555
func (m TradeCaptureReport) HasNoLegs() bool {
	return m.Has(tag.NoLegs)
}

//HasTradeRequestID returns true if TradeRequestID is present, Tag 568
func (m TradeCaptureReport) HasTradeRequestID() bool {
	return m.Has(tag.TradeRequestID)
}

//HasPreviouslyReported returns true if PreviouslyReported is present, Tag 570
func (m TradeCaptureReport) HasPreviouslyReported() bool {
	return m.Has(tag.PreviouslyReported)
}

//HasTradeReportID returns true if TradeReportID is present, Tag 571
func (m TradeCaptureReport) HasTradeReportID() bool {
	return m.Has(tag.TradeReportID)
}

//HasTradeReportRefID returns true if TradeReportRefID is present, Tag 572
func (m TradeCaptureReport) HasTradeReportRefID() bool {
	return m.Has(tag.TradeReportRefID)
}

//HasMatchStatus returns true if MatchStatus is present, Tag 573
func (m TradeCaptureReport) HasMatchStatus() bool {
	return m.Has(tag.MatchStatus)
}

//HasMatchType returns true if MatchType is present, Tag 574
func (m TradeCaptureReport) HasMatchType() bool {
	return m.Has(tag.MatchType)
}

//HasBenchmarkPrice returns true if BenchmarkPrice is present, Tag 662
func (m TradeCaptureReport) HasBenchmarkPrice() bool {
	return m.Has(tag.BenchmarkPrice)
}

//HasBenchmarkPriceType returns true if BenchmarkPriceType is present, Tag 663
func (m TradeCaptureReport) HasBenchmarkPriceType() bool {
	return m.Has(tag.BenchmarkPriceType)
}

//HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m TradeCaptureReport) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

//HasLastParPx returns true if LastParPx is present, Tag 669
func (m TradeCaptureReport) HasLastParPx() bool {
	return m.Has(tag.LastParPx)
}

//HasPool returns true if Pool is present, Tag 691
func (m TradeCaptureReport) HasPool() bool {
	return m.Has(tag.Pool)
}

//HasYieldRedemptionDate returns true if YieldRedemptionDate is present, Tag 696
func (m TradeCaptureReport) HasYieldRedemptionDate() bool {
	return m.Has(tag.YieldRedemptionDate)
}

//HasYieldRedemptionPrice returns true if YieldRedemptionPrice is present, Tag 697
func (m TradeCaptureReport) HasYieldRedemptionPrice() bool {
	return m.Has(tag.YieldRedemptionPrice)
}

//HasYieldRedemptionPriceType returns true if YieldRedemptionPriceType is present, Tag 698
func (m TradeCaptureReport) HasYieldRedemptionPriceType() bool {
	return m.Has(tag.YieldRedemptionPriceType)
}

//HasBenchmarkSecurityID returns true if BenchmarkSecurityID is present, Tag 699
func (m TradeCaptureReport) HasBenchmarkSecurityID() bool {
	return m.Has(tag.BenchmarkSecurityID)
}

//HasYieldCalcDate returns true if YieldCalcDate is present, Tag 701
func (m TradeCaptureReport) HasYieldCalcDate() bool {
	return m.Has(tag.YieldCalcDate)
}

//HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711
func (m TradeCaptureReport) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

//HasClearingBusinessDate returns true if ClearingBusinessDate is present, Tag 715
func (m TradeCaptureReport) HasClearingBusinessDate() bool {
	return m.Has(tag.ClearingBusinessDate)
}

//HasTotNumTradeReports returns true if TotNumTradeReports is present, Tag 748
func (m TradeCaptureReport) HasTotNumTradeReports() bool {
	return m.Has(tag.TotNumTradeReports)
}

//HasNoPosAmt returns true if NoPosAmt is present, Tag 753
func (m TradeCaptureReport) HasNoPosAmt() bool {
	return m.Has(tag.NoPosAmt)
}

//HasBenchmarkSecurityIDSource returns true if BenchmarkSecurityIDSource is present, Tag 761
func (m TradeCaptureReport) HasBenchmarkSecurityIDSource() bool {
	return m.Has(tag.BenchmarkSecurityIDSource)
}

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m TradeCaptureReport) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

//HasNoTrdRegTimestamps returns true if NoTrdRegTimestamps is present, Tag 768
func (m TradeCaptureReport) HasNoTrdRegTimestamps() bool {
	return m.Has(tag.NoTrdRegTimestamps)
}

//HasTerminationType returns true if TerminationType is present, Tag 788
func (m TradeCaptureReport) HasTerminationType() bool {
	return m.Has(tag.TerminationType)
}

//HasSecondaryTradeReportID returns true if SecondaryTradeReportID is present, Tag 818
func (m TradeCaptureReport) HasSecondaryTradeReportID() bool {
	return m.Has(tag.SecondaryTradeReportID)
}

//HasAvgPxIndicator returns true if AvgPxIndicator is present, Tag 819
func (m TradeCaptureReport) HasAvgPxIndicator() bool {
	return m.Has(tag.AvgPxIndicator)
}

//HasTradeLinkID returns true if TradeLinkID is present, Tag 820
func (m TradeCaptureReport) HasTradeLinkID() bool {
	return m.Has(tag.TradeLinkID)
}

//HasUnderlyingTradingSessionID returns true if UnderlyingTradingSessionID is present, Tag 822
func (m TradeCaptureReport) HasUnderlyingTradingSessionID() bool {
	return m.Has(tag.UnderlyingTradingSessionID)
}

//HasUnderlyingTradingSessionSubID returns true if UnderlyingTradingSessionSubID is present, Tag 823
func (m TradeCaptureReport) HasUnderlyingTradingSessionSubID() bool {
	return m.Has(tag.UnderlyingTradingSessionSubID)
}

//HasTradeLegRefID returns true if TradeLegRefID is present, Tag 824
func (m TradeCaptureReport) HasTradeLegRefID() bool {
	return m.Has(tag.TradeLegRefID)
}

//HasTrdType returns true if TrdType is present, Tag 828
func (m TradeCaptureReport) HasTrdType() bool {
	return m.Has(tag.TrdType)
}

//HasTrdSubType returns true if TrdSubType is present, Tag 829
func (m TradeCaptureReport) HasTrdSubType() bool {
	return m.Has(tag.TrdSubType)
}

//HasTransferReason returns true if TransferReason is present, Tag 830
func (m TradeCaptureReport) HasTransferReason() bool {
	return m.Has(tag.TransferReason)
}

//HasQtyType returns true if QtyType is present, Tag 854
func (m TradeCaptureReport) HasQtyType() bool {
	return m.Has(tag.QtyType)
}

//HasSecondaryTrdType returns true if SecondaryTrdType is present, Tag 855
func (m TradeCaptureReport) HasSecondaryTrdType() bool {
	return m.Has(tag.SecondaryTrdType)
}

//HasTradeReportType returns true if TradeReportType is present, Tag 856
func (m TradeCaptureReport) HasTradeReportType() bool {
	return m.Has(tag.TradeReportType)
}

//HasNoEvents returns true if NoEvents is present, Tag 864
func (m TradeCaptureReport) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

//HasDatedDate returns true if DatedDate is present, Tag 873
func (m TradeCaptureReport) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

//HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m TradeCaptureReport) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

//HasCPProgram returns true if CPProgram is present, Tag 875
func (m TradeCaptureReport) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

//HasCPRegType returns true if CPRegType is present, Tag 876
func (m TradeCaptureReport) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

//HasTrdMatchID returns true if TrdMatchID is present, Tag 880
func (m TradeCaptureReport) HasTrdMatchID() bool {
	return m.Has(tag.TrdMatchID)
}

//HasSecondaryTradeReportRefID returns true if SecondaryTradeReportRefID is present, Tag 881
func (m TradeCaptureReport) HasSecondaryTradeReportRefID() bool {
	return m.Has(tag.SecondaryTradeReportRefID)
}

//HasMarginRatio returns true if MarginRatio is present, Tag 898
func (m TradeCaptureReport) HasMarginRatio() bool {
	return m.Has(tag.MarginRatio)
}

//HasLastRptRequested returns true if LastRptRequested is present, Tag 912
func (m TradeCaptureReport) HasLastRptRequested() bool {
	return m.Has(tag.LastRptRequested)
}

//HasAgreementDesc returns true if AgreementDesc is present, Tag 913
func (m TradeCaptureReport) HasAgreementDesc() bool {
	return m.Has(tag.AgreementDesc)
}

//HasAgreementID returns true if AgreementID is present, Tag 914
func (m TradeCaptureReport) HasAgreementID() bool {
	return m.Has(tag.AgreementID)
}

//HasAgreementDate returns true if AgreementDate is present, Tag 915
func (m TradeCaptureReport) HasAgreementDate() bool {
	return m.Has(tag.AgreementDate)
}

//HasStartDate returns true if StartDate is present, Tag 916
func (m TradeCaptureReport) HasStartDate() bool {
	return m.Has(tag.StartDate)
}

//HasEndDate returns true if EndDate is present, Tag 917
func (m TradeCaptureReport) HasEndDate() bool {
	return m.Has(tag.EndDate)
}

//HasAgreementCurrency returns true if AgreementCurrency is present, Tag 918
func (m TradeCaptureReport) HasAgreementCurrency() bool {
	return m.Has(tag.AgreementCurrency)
}

//HasDeliveryType returns true if DeliveryType is present, Tag 919
func (m TradeCaptureReport) HasDeliveryType() bool {
	return m.Has(tag.DeliveryType)
}

//HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m TradeCaptureReport) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
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

//NoSides is a repeating group element, Tag 552
type NoSides struct {
	quickfix.Group
}

//SetSide sets Side, Tag 54
func (m NoSides) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetOrderID sets OrderID, Tag 37
func (m NoSides) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetSecondaryOrderID sets SecondaryOrderID, Tag 198
func (m NoSides) SetSecondaryOrderID(v string) {
	m.Set(field.NewSecondaryOrderID(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m NoSides) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526
func (m NoSides) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

//SetListID sets ListID, Tag 66
func (m NoSides) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m NoSides) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetAccount sets Account, Tag 1
func (m NoSides) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetAcctIDSource sets AcctIDSource, Tag 660
func (m NoSides) SetAcctIDSource(v int) {
	m.Set(field.NewAcctIDSource(v))
}

//SetAccountType sets AccountType, Tag 581
func (m NoSides) SetAccountType(v int) {
	m.Set(field.NewAccountType(v))
}

//SetProcessCode sets ProcessCode, Tag 81
func (m NoSides) SetProcessCode(v string) {
	m.Set(field.NewProcessCode(v))
}

//SetOddLot sets OddLot, Tag 575
func (m NoSides) SetOddLot(v bool) {
	m.Set(field.NewOddLot(v))
}

//SetNoClearingInstructions sets NoClearingInstructions, Tag 576
func (m NoSides) SetNoClearingInstructions(f NoClearingInstructionsRepeatingGroup) {
	m.SetGroup(f)
}

//SetClearingFeeIndicator sets ClearingFeeIndicator, Tag 635
func (m NoSides) SetClearingFeeIndicator(v string) {
	m.Set(field.NewClearingFeeIndicator(v))
}

//SetTradeInputSource sets TradeInputSource, Tag 578
func (m NoSides) SetTradeInputSource(v string) {
	m.Set(field.NewTradeInputSource(v))
}

//SetTradeInputDevice sets TradeInputDevice, Tag 579
func (m NoSides) SetTradeInputDevice(v string) {
	m.Set(field.NewTradeInputDevice(v))
}

//SetOrderInputDevice sets OrderInputDevice, Tag 821
func (m NoSides) SetOrderInputDevice(v string) {
	m.Set(field.NewOrderInputDevice(v))
}

//SetCurrency sets Currency, Tag 15
func (m NoSides) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetComplianceID sets ComplianceID, Tag 376
func (m NoSides) SetComplianceID(v string) {
	m.Set(field.NewComplianceID(v))
}

//SetSolicitedFlag sets SolicitedFlag, Tag 377
func (m NoSides) SetSolicitedFlag(v bool) {
	m.Set(field.NewSolicitedFlag(v))
}

//SetOrderCapacity sets OrderCapacity, Tag 528
func (m NoSides) SetOrderCapacity(v string) {
	m.Set(field.NewOrderCapacity(v))
}

//SetOrderRestrictions sets OrderRestrictions, Tag 529
func (m NoSides) SetOrderRestrictions(v string) {
	m.Set(field.NewOrderRestrictions(v))
}

//SetCustOrderCapacity sets CustOrderCapacity, Tag 582
func (m NoSides) SetCustOrderCapacity(v int) {
	m.Set(field.NewCustOrderCapacity(v))
}

//SetOrdType sets OrdType, Tag 40
func (m NoSides) SetOrdType(v string) {
	m.Set(field.NewOrdType(v))
}

//SetExecInst sets ExecInst, Tag 18
func (m NoSides) SetExecInst(v string) {
	m.Set(field.NewExecInst(v))
}

//SetTransBkdTime sets TransBkdTime, Tag 483
func (m NoSides) SetTransBkdTime(v time.Time) {
	m.Set(field.NewTransBkdTime(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoSides) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m NoSides) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetTimeBracket sets TimeBracket, Tag 943
func (m NoSides) SetTimeBracket(v string) {
	m.Set(field.NewTimeBracket(v))
}

//SetCommission sets Commission, Tag 12
func (m NoSides) SetCommission(v float64) {
	m.Set(field.NewCommission(v))
}

//SetCommType sets CommType, Tag 13
func (m NoSides) SetCommType(v string) {
	m.Set(field.NewCommType(v))
}

//SetCommCurrency sets CommCurrency, Tag 479
func (m NoSides) SetCommCurrency(v string) {
	m.Set(field.NewCommCurrency(v))
}

//SetFundRenewWaiv sets FundRenewWaiv, Tag 497
func (m NoSides) SetFundRenewWaiv(v string) {
	m.Set(field.NewFundRenewWaiv(v))
}

//SetGrossTradeAmt sets GrossTradeAmt, Tag 381
func (m NoSides) SetGrossTradeAmt(v float64) {
	m.Set(field.NewGrossTradeAmt(v))
}

//SetNumDaysInterest sets NumDaysInterest, Tag 157
func (m NoSides) SetNumDaysInterest(v int) {
	m.Set(field.NewNumDaysInterest(v))
}

//SetExDate sets ExDate, Tag 230
func (m NoSides) SetExDate(v string) {
	m.Set(field.NewExDate(v))
}

//SetAccruedInterestRate sets AccruedInterestRate, Tag 158
func (m NoSides) SetAccruedInterestRate(v float64) {
	m.Set(field.NewAccruedInterestRate(v))
}

//SetAccruedInterestAmt sets AccruedInterestAmt, Tag 159
func (m NoSides) SetAccruedInterestAmt(v float64) {
	m.Set(field.NewAccruedInterestAmt(v))
}

//SetInterestAtMaturity sets InterestAtMaturity, Tag 738
func (m NoSides) SetInterestAtMaturity(v float64) {
	m.Set(field.NewInterestAtMaturity(v))
}

//SetEndAccruedInterestAmt sets EndAccruedInterestAmt, Tag 920
func (m NoSides) SetEndAccruedInterestAmt(v float64) {
	m.Set(field.NewEndAccruedInterestAmt(v))
}

//SetStartCash sets StartCash, Tag 921
func (m NoSides) SetStartCash(v float64) {
	m.Set(field.NewStartCash(v))
}

//SetEndCash sets EndCash, Tag 922
func (m NoSides) SetEndCash(v float64) {
	m.Set(field.NewEndCash(v))
}

//SetConcession sets Concession, Tag 238
func (m NoSides) SetConcession(v float64) {
	m.Set(field.NewConcession(v))
}

//SetTotalTakedown sets TotalTakedown, Tag 237
func (m NoSides) SetTotalTakedown(v float64) {
	m.Set(field.NewTotalTakedown(v))
}

//SetNetMoney sets NetMoney, Tag 118
func (m NoSides) SetNetMoney(v float64) {
	m.Set(field.NewNetMoney(v))
}

//SetSettlCurrAmt sets SettlCurrAmt, Tag 119
func (m NoSides) SetSettlCurrAmt(v float64) {
	m.Set(field.NewSettlCurrAmt(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m NoSides) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetSettlCurrFxRate sets SettlCurrFxRate, Tag 155
func (m NoSides) SetSettlCurrFxRate(v float64) {
	m.Set(field.NewSettlCurrFxRate(v))
}

//SetSettlCurrFxRateCalc sets SettlCurrFxRateCalc, Tag 156
func (m NoSides) SetSettlCurrFxRateCalc(v string) {
	m.Set(field.NewSettlCurrFxRateCalc(v))
}

//SetPositionEffect sets PositionEffect, Tag 77
func (m NoSides) SetPositionEffect(v string) {
	m.Set(field.NewPositionEffect(v))
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

//SetSideMultiLegReportingType sets SideMultiLegReportingType, Tag 752
func (m NoSides) SetSideMultiLegReportingType(v int) {
	m.Set(field.NewSideMultiLegReportingType(v))
}

//SetNoContAmts sets NoContAmts, Tag 518
func (m NoSides) SetNoContAmts(f NoContAmtsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoStipulations sets NoStipulations, Tag 232
func (m NoSides) SetNoStipulations(f NoStipulationsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoMiscFees sets NoMiscFees, Tag 136
func (m NoSides) SetNoMiscFees(f NoMiscFeesRepeatingGroup) {
	m.SetGroup(f)
}

//SetExchangeRule sets ExchangeRule, Tag 825
func (m NoSides) SetExchangeRule(v string) {
	m.Set(field.NewExchangeRule(v))
}

//SetTradeAllocIndicator sets TradeAllocIndicator, Tag 826
func (m NoSides) SetTradeAllocIndicator(v int) {
	m.Set(field.NewTradeAllocIndicator(v))
}

//SetPreallocMethod sets PreallocMethod, Tag 591
func (m NoSides) SetPreallocMethod(v string) {
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

//SetCopyMsgIndicator sets CopyMsgIndicator, Tag 797
func (m NoSides) SetCopyMsgIndicator(v bool) {
	m.Set(field.NewCopyMsgIndicator(v))
}

//SetPublishTrdIndicator sets PublishTrdIndicator, Tag 852
func (m NoSides) SetPublishTrdIndicator(v bool) {
	m.Set(field.NewPublishTrdIndicator(v))
}

//SetShortSaleReason sets ShortSaleReason, Tag 853
func (m NoSides) SetShortSaleReason(v int) {
	m.Set(field.NewShortSaleReason(v))
}

//GetSide gets Side, Tag 54
func (m NoSides) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m NoSides) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryOrderID gets SecondaryOrderID, Tag 198
func (m NoSides) GetSecondaryOrderID() (f field.SecondaryOrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m NoSides) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526
func (m NoSides) GetSecondaryClOrdID() (f field.SecondaryClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListID gets ListID, Tag 66
func (m NoSides) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m NoSides) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAccount gets Account, Tag 1
func (m NoSides) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAcctIDSource gets AcctIDSource, Tag 660
func (m NoSides) GetAcctIDSource() (f field.AcctIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccountType gets AccountType, Tag 581
func (m NoSides) GetAccountType() (f field.AccountTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProcessCode gets ProcessCode, Tag 81
func (m NoSides) GetProcessCode() (f field.ProcessCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOddLot gets OddLot, Tag 575
func (m NoSides) GetOddLot() (f field.OddLotField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoClearingInstructions gets NoClearingInstructions, Tag 576
func (m NoSides) GetNoClearingInstructions() (NoClearingInstructionsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoClearingInstructionsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetClearingFeeIndicator gets ClearingFeeIndicator, Tag 635
func (m NoSides) GetClearingFeeIndicator() (f field.ClearingFeeIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeInputSource gets TradeInputSource, Tag 578
func (m NoSides) GetTradeInputSource() (f field.TradeInputSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeInputDevice gets TradeInputDevice, Tag 579
func (m NoSides) GetTradeInputDevice() (f field.TradeInputDeviceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderInputDevice gets OrderInputDevice, Tag 821
func (m NoSides) GetOrderInputDevice() (f field.OrderInputDeviceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m NoSides) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplianceID gets ComplianceID, Tag 376
func (m NoSides) GetComplianceID() (f field.ComplianceIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSolicitedFlag gets SolicitedFlag, Tag 377
func (m NoSides) GetSolicitedFlag() (f field.SolicitedFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderCapacity gets OrderCapacity, Tag 528
func (m NoSides) GetOrderCapacity() (f field.OrderCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderRestrictions gets OrderRestrictions, Tag 529
func (m NoSides) GetOrderRestrictions() (f field.OrderRestrictionsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCustOrderCapacity gets CustOrderCapacity, Tag 582
func (m NoSides) GetCustOrderCapacity() (f field.CustOrderCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdType gets OrdType, Tag 40
func (m NoSides) GetOrdType() (f field.OrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecInst gets ExecInst, Tag 18
func (m NoSides) GetExecInst() (f field.ExecInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransBkdTime gets TransBkdTime, Tag 483
func (m NoSides) GetTransBkdTime() (f field.TransBkdTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoSides) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m NoSides) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeBracket gets TimeBracket, Tag 943
func (m NoSides) GetTimeBracket() (f field.TimeBracketField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommission gets Commission, Tag 12
func (m NoSides) GetCommission() (f field.CommissionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommType gets CommType, Tag 13
func (m NoSides) GetCommType() (f field.CommTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommCurrency gets CommCurrency, Tag 479
func (m NoSides) GetCommCurrency() (f field.CommCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFundRenewWaiv gets FundRenewWaiv, Tag 497
func (m NoSides) GetFundRenewWaiv() (f field.FundRenewWaivField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetGrossTradeAmt gets GrossTradeAmt, Tag 381
func (m NoSides) GetGrossTradeAmt() (f field.GrossTradeAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNumDaysInterest gets NumDaysInterest, Tag 157
func (m NoSides) GetNumDaysInterest() (f field.NumDaysInterestField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExDate gets ExDate, Tag 230
func (m NoSides) GetExDate() (f field.ExDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccruedInterestRate gets AccruedInterestRate, Tag 158
func (m NoSides) GetAccruedInterestRate() (f field.AccruedInterestRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccruedInterestAmt gets AccruedInterestAmt, Tag 159
func (m NoSides) GetAccruedInterestAmt() (f field.AccruedInterestAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInterestAtMaturity gets InterestAtMaturity, Tag 738
func (m NoSides) GetInterestAtMaturity() (f field.InterestAtMaturityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEndAccruedInterestAmt gets EndAccruedInterestAmt, Tag 920
func (m NoSides) GetEndAccruedInterestAmt() (f field.EndAccruedInterestAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStartCash gets StartCash, Tag 921
func (m NoSides) GetStartCash() (f field.StartCashField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEndCash gets EndCash, Tag 922
func (m NoSides) GetEndCash() (f field.EndCashField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetConcession gets Concession, Tag 238
func (m NoSides) GetConcession() (f field.ConcessionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTotalTakedown gets TotalTakedown, Tag 237
func (m NoSides) GetTotalTakedown() (f field.TotalTakedownField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNetMoney gets NetMoney, Tag 118
func (m NoSides) GetNetMoney() (f field.NetMoneyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrAmt gets SettlCurrAmt, Tag 119
func (m NoSides) GetSettlCurrAmt() (f field.SettlCurrAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m NoSides) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrFxRate gets SettlCurrFxRate, Tag 155
func (m NoSides) GetSettlCurrFxRate() (f field.SettlCurrFxRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrFxRateCalc gets SettlCurrFxRateCalc, Tag 156
func (m NoSides) GetSettlCurrFxRateCalc() (f field.SettlCurrFxRateCalcField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionEffect gets PositionEffect, Tag 77
func (m NoSides) GetPositionEffect() (f field.PositionEffectField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m NoSides) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m NoSides) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m NoSides) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSideMultiLegReportingType gets SideMultiLegReportingType, Tag 752
func (m NoSides) GetSideMultiLegReportingType() (f field.SideMultiLegReportingTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoContAmts gets NoContAmts, Tag 518
func (m NoSides) GetNoContAmts() (NoContAmtsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoContAmtsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoStipulations gets NoStipulations, Tag 232
func (m NoSides) GetNoStipulations() (NoStipulationsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStipulationsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoMiscFees gets NoMiscFees, Tag 136
func (m NoSides) GetNoMiscFees() (NoMiscFeesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMiscFeesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetExchangeRule gets ExchangeRule, Tag 825
func (m NoSides) GetExchangeRule() (f field.ExchangeRuleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeAllocIndicator gets TradeAllocIndicator, Tag 826
func (m NoSides) GetTradeAllocIndicator() (f field.TradeAllocIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPreallocMethod gets PreallocMethod, Tag 591
func (m NoSides) GetPreallocMethod() (f field.PreallocMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocID gets AllocID, Tag 70
func (m NoSides) GetAllocID() (f field.AllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoAllocs gets NoAllocs, Tag 78
func (m NoSides) GetNoAllocs() (NoAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetCopyMsgIndicator gets CopyMsgIndicator, Tag 797
func (m NoSides) GetCopyMsgIndicator() (f field.CopyMsgIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPublishTrdIndicator gets PublishTrdIndicator, Tag 852
func (m NoSides) GetPublishTrdIndicator() (f field.PublishTrdIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetShortSaleReason gets ShortSaleReason, Tag 853
func (m NoSides) GetShortSaleReason() (f field.ShortSaleReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasSide returns true if Side is present, Tag 54
func (m NoSides) HasSide() bool {
	return m.Has(tag.Side)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m NoSides) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasSecondaryOrderID returns true if SecondaryOrderID is present, Tag 198
func (m NoSides) HasSecondaryOrderID() bool {
	return m.Has(tag.SecondaryOrderID)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m NoSides) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526
func (m NoSides) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

//HasListID returns true if ListID is present, Tag 66
func (m NoSides) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m NoSides) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
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

//HasProcessCode returns true if ProcessCode is present, Tag 81
func (m NoSides) HasProcessCode() bool {
	return m.Has(tag.ProcessCode)
}

//HasOddLot returns true if OddLot is present, Tag 575
func (m NoSides) HasOddLot() bool {
	return m.Has(tag.OddLot)
}

//HasNoClearingInstructions returns true if NoClearingInstructions is present, Tag 576
func (m NoSides) HasNoClearingInstructions() bool {
	return m.Has(tag.NoClearingInstructions)
}

//HasClearingFeeIndicator returns true if ClearingFeeIndicator is present, Tag 635
func (m NoSides) HasClearingFeeIndicator() bool {
	return m.Has(tag.ClearingFeeIndicator)
}

//HasTradeInputSource returns true if TradeInputSource is present, Tag 578
func (m NoSides) HasTradeInputSource() bool {
	return m.Has(tag.TradeInputSource)
}

//HasTradeInputDevice returns true if TradeInputDevice is present, Tag 579
func (m NoSides) HasTradeInputDevice() bool {
	return m.Has(tag.TradeInputDevice)
}

//HasOrderInputDevice returns true if OrderInputDevice is present, Tag 821
func (m NoSides) HasOrderInputDevice() bool {
	return m.Has(tag.OrderInputDevice)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m NoSides) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasComplianceID returns true if ComplianceID is present, Tag 376
func (m NoSides) HasComplianceID() bool {
	return m.Has(tag.ComplianceID)
}

//HasSolicitedFlag returns true if SolicitedFlag is present, Tag 377
func (m NoSides) HasSolicitedFlag() bool {
	return m.Has(tag.SolicitedFlag)
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

//HasOrdType returns true if OrdType is present, Tag 40
func (m NoSides) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasExecInst returns true if ExecInst is present, Tag 18
func (m NoSides) HasExecInst() bool {
	return m.Has(tag.ExecInst)
}

//HasTransBkdTime returns true if TransBkdTime is present, Tag 483
func (m NoSides) HasTransBkdTime() bool {
	return m.Has(tag.TransBkdTime)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m NoSides) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m NoSides) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//HasTimeBracket returns true if TimeBracket is present, Tag 943
func (m NoSides) HasTimeBracket() bool {
	return m.Has(tag.TimeBracket)
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

//HasGrossTradeAmt returns true if GrossTradeAmt is present, Tag 381
func (m NoSides) HasGrossTradeAmt() bool {
	return m.Has(tag.GrossTradeAmt)
}

//HasNumDaysInterest returns true if NumDaysInterest is present, Tag 157
func (m NoSides) HasNumDaysInterest() bool {
	return m.Has(tag.NumDaysInterest)
}

//HasExDate returns true if ExDate is present, Tag 230
func (m NoSides) HasExDate() bool {
	return m.Has(tag.ExDate)
}

//HasAccruedInterestRate returns true if AccruedInterestRate is present, Tag 158
func (m NoSides) HasAccruedInterestRate() bool {
	return m.Has(tag.AccruedInterestRate)
}

//HasAccruedInterestAmt returns true if AccruedInterestAmt is present, Tag 159
func (m NoSides) HasAccruedInterestAmt() bool {
	return m.Has(tag.AccruedInterestAmt)
}

//HasInterestAtMaturity returns true if InterestAtMaturity is present, Tag 738
func (m NoSides) HasInterestAtMaturity() bool {
	return m.Has(tag.InterestAtMaturity)
}

//HasEndAccruedInterestAmt returns true if EndAccruedInterestAmt is present, Tag 920
func (m NoSides) HasEndAccruedInterestAmt() bool {
	return m.Has(tag.EndAccruedInterestAmt)
}

//HasStartCash returns true if StartCash is present, Tag 921
func (m NoSides) HasStartCash() bool {
	return m.Has(tag.StartCash)
}

//HasEndCash returns true if EndCash is present, Tag 922
func (m NoSides) HasEndCash() bool {
	return m.Has(tag.EndCash)
}

//HasConcession returns true if Concession is present, Tag 238
func (m NoSides) HasConcession() bool {
	return m.Has(tag.Concession)
}

//HasTotalTakedown returns true if TotalTakedown is present, Tag 237
func (m NoSides) HasTotalTakedown() bool {
	return m.Has(tag.TotalTakedown)
}

//HasNetMoney returns true if NetMoney is present, Tag 118
func (m NoSides) HasNetMoney() bool {
	return m.Has(tag.NetMoney)
}

//HasSettlCurrAmt returns true if SettlCurrAmt is present, Tag 119
func (m NoSides) HasSettlCurrAmt() bool {
	return m.Has(tag.SettlCurrAmt)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m NoSides) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasSettlCurrFxRate returns true if SettlCurrFxRate is present, Tag 155
func (m NoSides) HasSettlCurrFxRate() bool {
	return m.Has(tag.SettlCurrFxRate)
}

//HasSettlCurrFxRateCalc returns true if SettlCurrFxRateCalc is present, Tag 156
func (m NoSides) HasSettlCurrFxRateCalc() bool {
	return m.Has(tag.SettlCurrFxRateCalc)
}

//HasPositionEffect returns true if PositionEffect is present, Tag 77
func (m NoSides) HasPositionEffect() bool {
	return m.Has(tag.PositionEffect)
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

//HasSideMultiLegReportingType returns true if SideMultiLegReportingType is present, Tag 752
func (m NoSides) HasSideMultiLegReportingType() bool {
	return m.Has(tag.SideMultiLegReportingType)
}

//HasNoContAmts returns true if NoContAmts is present, Tag 518
func (m NoSides) HasNoContAmts() bool {
	return m.Has(tag.NoContAmts)
}

//HasNoStipulations returns true if NoStipulations is present, Tag 232
func (m NoSides) HasNoStipulations() bool {
	return m.Has(tag.NoStipulations)
}

//HasNoMiscFees returns true if NoMiscFees is present, Tag 136
func (m NoSides) HasNoMiscFees() bool {
	return m.Has(tag.NoMiscFees)
}

//HasExchangeRule returns true if ExchangeRule is present, Tag 825
func (m NoSides) HasExchangeRule() bool {
	return m.Has(tag.ExchangeRule)
}

//HasTradeAllocIndicator returns true if TradeAllocIndicator is present, Tag 826
func (m NoSides) HasTradeAllocIndicator() bool {
	return m.Has(tag.TradeAllocIndicator)
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

//HasCopyMsgIndicator returns true if CopyMsgIndicator is present, Tag 797
func (m NoSides) HasCopyMsgIndicator() bool {
	return m.Has(tag.CopyMsgIndicator)
}

//HasPublishTrdIndicator returns true if PublishTrdIndicator is present, Tag 852
func (m NoSides) HasPublishTrdIndicator() bool {
	return m.Has(tag.PublishTrdIndicator)
}

//HasShortSaleReason returns true if ShortSaleReason is present, Tag 853
func (m NoSides) HasShortSaleReason() bool {
	return m.Has(tag.ShortSaleReason)
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

//NoClearingInstructions is a repeating group element, Tag 576
type NoClearingInstructions struct {
	quickfix.Group
}

//SetClearingInstruction sets ClearingInstruction, Tag 577
func (m NoClearingInstructions) SetClearingInstruction(v int) {
	m.Set(field.NewClearingInstruction(v))
}

//GetClearingInstruction gets ClearingInstruction, Tag 577
func (m NoClearingInstructions) GetClearingInstruction() (f field.ClearingInstructionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasClearingInstruction returns true if ClearingInstruction is present, Tag 577
func (m NoClearingInstructions) HasClearingInstruction() bool {
	return m.Has(tag.ClearingInstruction)
}

//NoClearingInstructionsRepeatingGroup is a repeating group, Tag 576
type NoClearingInstructionsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoClearingInstructionsRepeatingGroup returns an initialized, NoClearingInstructionsRepeatingGroup
func NewNoClearingInstructionsRepeatingGroup() NoClearingInstructionsRepeatingGroup {
	return NoClearingInstructionsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoClearingInstructions,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ClearingInstruction)})}
}

//Add create and append a new NoClearingInstructions to this group
func (m NoClearingInstructionsRepeatingGroup) Add() NoClearingInstructions {
	g := m.RepeatingGroup.Add()
	return NoClearingInstructions{g}
}

//Get returns the ith NoClearingInstructions in the NoClearingInstructionsRepeatinGroup
func (m NoClearingInstructionsRepeatingGroup) Get(i int) NoClearingInstructions {
	return NoClearingInstructions{m.RepeatingGroup.Get(i)}
}

//NoContAmts is a repeating group element, Tag 518
type NoContAmts struct {
	quickfix.Group
}

//SetContAmtType sets ContAmtType, Tag 519
func (m NoContAmts) SetContAmtType(v int) {
	m.Set(field.NewContAmtType(v))
}

//SetContAmtValue sets ContAmtValue, Tag 520
func (m NoContAmts) SetContAmtValue(v float64) {
	m.Set(field.NewContAmtValue(v))
}

//SetContAmtCurr sets ContAmtCurr, Tag 521
func (m NoContAmts) SetContAmtCurr(v string) {
	m.Set(field.NewContAmtCurr(v))
}

//GetContAmtType gets ContAmtType, Tag 519
func (m NoContAmts) GetContAmtType() (f field.ContAmtTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContAmtValue gets ContAmtValue, Tag 520
func (m NoContAmts) GetContAmtValue() (f field.ContAmtValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContAmtCurr gets ContAmtCurr, Tag 521
func (m NoContAmts) GetContAmtCurr() (f field.ContAmtCurrField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasContAmtType returns true if ContAmtType is present, Tag 519
func (m NoContAmts) HasContAmtType() bool {
	return m.Has(tag.ContAmtType)
}

//HasContAmtValue returns true if ContAmtValue is present, Tag 520
func (m NoContAmts) HasContAmtValue() bool {
	return m.Has(tag.ContAmtValue)
}

//HasContAmtCurr returns true if ContAmtCurr is present, Tag 521
func (m NoContAmts) HasContAmtCurr() bool {
	return m.Has(tag.ContAmtCurr)
}

//NoContAmtsRepeatingGroup is a repeating group, Tag 518
type NoContAmtsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoContAmtsRepeatingGroup returns an initialized, NoContAmtsRepeatingGroup
func NewNoContAmtsRepeatingGroup() NoContAmtsRepeatingGroup {
	return NoContAmtsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoContAmts,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ContAmtType), quickfix.GroupElement(tag.ContAmtValue), quickfix.GroupElement(tag.ContAmtCurr)})}
}

//Add create and append a new NoContAmts to this group
func (m NoContAmtsRepeatingGroup) Add() NoContAmts {
	g := m.RepeatingGroup.Add()
	return NoContAmts{g}
}

//Get returns the ith NoContAmts in the NoContAmtsRepeatinGroup
func (m NoContAmtsRepeatingGroup) Get(i int) NoContAmts {
	return NoContAmts{m.RepeatingGroup.Get(i)}
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

//NoMiscFees is a repeating group element, Tag 136
type NoMiscFees struct {
	quickfix.Group
}

//SetMiscFeeAmt sets MiscFeeAmt, Tag 137
func (m NoMiscFees) SetMiscFeeAmt(v float64) {
	m.Set(field.NewMiscFeeAmt(v))
}

//SetMiscFeeCurr sets MiscFeeCurr, Tag 138
func (m NoMiscFees) SetMiscFeeCurr(v string) {
	m.Set(field.NewMiscFeeCurr(v))
}

//SetMiscFeeType sets MiscFeeType, Tag 139
func (m NoMiscFees) SetMiscFeeType(v string) {
	m.Set(field.NewMiscFeeType(v))
}

//SetMiscFeeBasis sets MiscFeeBasis, Tag 891
func (m NoMiscFees) SetMiscFeeBasis(v int) {
	m.Set(field.NewMiscFeeBasis(v))
}

//GetMiscFeeAmt gets MiscFeeAmt, Tag 137
func (m NoMiscFees) GetMiscFeeAmt() (f field.MiscFeeAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMiscFeeCurr gets MiscFeeCurr, Tag 138
func (m NoMiscFees) GetMiscFeeCurr() (f field.MiscFeeCurrField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMiscFeeType gets MiscFeeType, Tag 139
func (m NoMiscFees) GetMiscFeeType() (f field.MiscFeeTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMiscFeeBasis gets MiscFeeBasis, Tag 891
func (m NoMiscFees) GetMiscFeeBasis() (f field.MiscFeeBasisField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasMiscFeeAmt returns true if MiscFeeAmt is present, Tag 137
func (m NoMiscFees) HasMiscFeeAmt() bool {
	return m.Has(tag.MiscFeeAmt)
}

//HasMiscFeeCurr returns true if MiscFeeCurr is present, Tag 138
func (m NoMiscFees) HasMiscFeeCurr() bool {
	return m.Has(tag.MiscFeeCurr)
}

//HasMiscFeeType returns true if MiscFeeType is present, Tag 139
func (m NoMiscFees) HasMiscFeeType() bool {
	return m.Has(tag.MiscFeeType)
}

//HasMiscFeeBasis returns true if MiscFeeBasis is present, Tag 891
func (m NoMiscFees) HasMiscFeeBasis() bool {
	return m.Has(tag.MiscFeeBasis)
}

//NoMiscFeesRepeatingGroup is a repeating group, Tag 136
type NoMiscFeesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMiscFeesRepeatingGroup returns an initialized, NoMiscFeesRepeatingGroup
func NewNoMiscFeesRepeatingGroup() NoMiscFeesRepeatingGroup {
	return NoMiscFeesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMiscFees,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MiscFeeAmt), quickfix.GroupElement(tag.MiscFeeCurr), quickfix.GroupElement(tag.MiscFeeType), quickfix.GroupElement(tag.MiscFeeBasis)})}
}

//Add create and append a new NoMiscFees to this group
func (m NoMiscFeesRepeatingGroup) Add() NoMiscFees {
	g := m.RepeatingGroup.Add()
	return NoMiscFees{g}
}

//Get returns the ith NoMiscFees in the NoMiscFeesRepeatinGroup
func (m NoMiscFeesRepeatingGroup) Get(i int) NoMiscFees {
	return NoMiscFees{m.RepeatingGroup.Get(i)}
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

//SetNoNested2PartyIDs sets NoNested2PartyIDs, Tag 756
func (m NoAllocs) SetNoNested2PartyIDs(f NoNested2PartyIDsRepeatingGroup) {
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

//GetNoNested2PartyIDs gets NoNested2PartyIDs, Tag 756
func (m NoAllocs) GetNoNested2PartyIDs() (NoNested2PartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNested2PartyIDsRepeatingGroup()
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

//HasNoNested2PartyIDs returns true if NoNested2PartyIDs is present, Tag 756
func (m NoAllocs) HasNoNested2PartyIDs() bool {
	return m.Has(tag.NoNested2PartyIDs)
}

//HasAllocQty returns true if AllocQty is present, Tag 80
func (m NoAllocs) HasAllocQty() bool {
	return m.Has(tag.AllocQty)
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Nested2PartyID), quickfix.GroupElement(tag.Nested2PartyIDSource), quickfix.GroupElement(tag.Nested2PartyRole), quickfix.GroupElement(tag.NoNested2PartySubIDs)})}
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

//NoAllocsRepeatingGroup is a repeating group, Tag 78
type NoAllocsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoAllocsRepeatingGroup returns an initialized, NoAllocsRepeatingGroup
func NewNoAllocsRepeatingGroup() NoAllocsRepeatingGroup {
	return NoAllocsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAllocs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AllocAccount), quickfix.GroupElement(tag.AllocAcctIDSource), quickfix.GroupElement(tag.AllocSettlCurrency), quickfix.GroupElement(tag.IndividualAllocID), quickfix.GroupElement(tag.NoNested2PartyIDs), quickfix.GroupElement(tag.AllocQty)})}
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Side), quickfix.GroupElement(tag.OrderID), quickfix.GroupElement(tag.SecondaryOrderID), quickfix.GroupElement(tag.ClOrdID), quickfix.GroupElement(tag.SecondaryClOrdID), quickfix.GroupElement(tag.ListID), quickfix.GroupElement(tag.NoPartyIDs), quickfix.GroupElement(tag.Account), quickfix.GroupElement(tag.AcctIDSource), quickfix.GroupElement(tag.AccountType), quickfix.GroupElement(tag.ProcessCode), quickfix.GroupElement(tag.OddLot), quickfix.GroupElement(tag.NoClearingInstructions), quickfix.GroupElement(tag.ClearingFeeIndicator), quickfix.GroupElement(tag.TradeInputSource), quickfix.GroupElement(tag.TradeInputDevice), quickfix.GroupElement(tag.OrderInputDevice), quickfix.GroupElement(tag.Currency), quickfix.GroupElement(tag.ComplianceID), quickfix.GroupElement(tag.SolicitedFlag), quickfix.GroupElement(tag.OrderCapacity), quickfix.GroupElement(tag.OrderRestrictions), quickfix.GroupElement(tag.CustOrderCapacity), quickfix.GroupElement(tag.OrdType), quickfix.GroupElement(tag.ExecInst), quickfix.GroupElement(tag.TransBkdTime), quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.TradingSessionSubID), quickfix.GroupElement(tag.TimeBracket), quickfix.GroupElement(tag.Commission), quickfix.GroupElement(tag.CommType), quickfix.GroupElement(tag.CommCurrency), quickfix.GroupElement(tag.FundRenewWaiv), quickfix.GroupElement(tag.GrossTradeAmt), quickfix.GroupElement(tag.NumDaysInterest), quickfix.GroupElement(tag.ExDate), quickfix.GroupElement(tag.AccruedInterestRate), quickfix.GroupElement(tag.AccruedInterestAmt), quickfix.GroupElement(tag.InterestAtMaturity), quickfix.GroupElement(tag.EndAccruedInterestAmt), quickfix.GroupElement(tag.StartCash), quickfix.GroupElement(tag.EndCash), quickfix.GroupElement(tag.Concession), quickfix.GroupElement(tag.TotalTakedown), quickfix.GroupElement(tag.NetMoney), quickfix.GroupElement(tag.SettlCurrAmt), quickfix.GroupElement(tag.SettlCurrency), quickfix.GroupElement(tag.SettlCurrFxRate), quickfix.GroupElement(tag.SettlCurrFxRateCalc), quickfix.GroupElement(tag.PositionEffect), quickfix.GroupElement(tag.Text), quickfix.GroupElement(tag.EncodedTextLen), quickfix.GroupElement(tag.EncodedText), quickfix.GroupElement(tag.SideMultiLegReportingType), quickfix.GroupElement(tag.NoContAmts), quickfix.GroupElement(tag.NoStipulations), quickfix.GroupElement(tag.NoMiscFees), quickfix.GroupElement(tag.ExchangeRule), quickfix.GroupElement(tag.TradeAllocIndicator), quickfix.GroupElement(tag.PreallocMethod), quickfix.GroupElement(tag.AllocID), quickfix.GroupElement(tag.NoAllocs), quickfix.GroupElement(tag.CopyMsgIndicator), quickfix.GroupElement(tag.PublishTrdIndicator), quickfix.GroupElement(tag.ShortSaleReason)})}
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

//SetLegLastPx sets LegLastPx, Tag 637
func (m NoLegs) SetLegLastPx(v float64) {
	m.Set(field.NewLegLastPx(v))
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

//GetLegLastPx gets LegLastPx, Tag 637
func (m NoLegs) GetLegLastPx() (f field.LegLastPxField, err quickfix.MessageRejectError) {
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

//HasLegLastPx returns true if LegLastPx is present, Tag 637
func (m NoLegs) HasLegLastPx() bool {
	return m.Has(tag.LegLastPx)
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

//NoLegsRepeatingGroup is a repeating group, Tag 555
type NoLegsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegsRepeatingGroup returns an initialized, NoLegsRepeatingGroup
func NewNoLegsRepeatingGroup() NoLegsRepeatingGroup {
	return NoLegsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSymbol), quickfix.GroupElement(tag.LegSymbolSfx), quickfix.GroupElement(tag.LegSecurityID), quickfix.GroupElement(tag.LegSecurityIDSource), quickfix.GroupElement(tag.NoLegSecurityAltID), quickfix.GroupElement(tag.LegProduct), quickfix.GroupElement(tag.LegCFICode), quickfix.GroupElement(tag.LegSecurityType), quickfix.GroupElement(tag.LegSecuritySubType), quickfix.GroupElement(tag.LegMaturityMonthYear), quickfix.GroupElement(tag.LegMaturityDate), quickfix.GroupElement(tag.LegCouponPaymentDate), quickfix.GroupElement(tag.LegIssueDate), quickfix.GroupElement(tag.LegRepoCollateralSecurityType), quickfix.GroupElement(tag.LegRepurchaseTerm), quickfix.GroupElement(tag.LegRepurchaseRate), quickfix.GroupElement(tag.LegFactor), quickfix.GroupElement(tag.LegCreditRating), quickfix.GroupElement(tag.LegInstrRegistry), quickfix.GroupElement(tag.LegCountryOfIssue), quickfix.GroupElement(tag.LegStateOrProvinceOfIssue), quickfix.GroupElement(tag.LegLocaleOfIssue), quickfix.GroupElement(tag.LegRedemptionDate), quickfix.GroupElement(tag.LegStrikePrice), quickfix.GroupElement(tag.LegStrikeCurrency), quickfix.GroupElement(tag.LegOptAttribute), quickfix.GroupElement(tag.LegContractMultiplier), quickfix.GroupElement(tag.LegCouponRate), quickfix.GroupElement(tag.LegSecurityExchange), quickfix.GroupElement(tag.LegIssuer), quickfix.GroupElement(tag.EncodedLegIssuerLen), quickfix.GroupElement(tag.EncodedLegIssuer), quickfix.GroupElement(tag.LegSecurityDesc), quickfix.GroupElement(tag.EncodedLegSecurityDescLen), quickfix.GroupElement(tag.EncodedLegSecurityDesc), quickfix.GroupElement(tag.LegRatioQty), quickfix.GroupElement(tag.LegSide), quickfix.GroupElement(tag.LegCurrency), quickfix.GroupElement(tag.LegPool), quickfix.GroupElement(tag.LegDatedDate), quickfix.GroupElement(tag.LegContractSettlMonth), quickfix.GroupElement(tag.LegInterestAccrualDate), quickfix.GroupElement(tag.LegQty), quickfix.GroupElement(tag.LegSwapType), quickfix.GroupElement(tag.NoLegStipulations), quickfix.GroupElement(tag.LegPositionEffect), quickfix.GroupElement(tag.LegCoveredOrUncovered), quickfix.GroupElement(tag.NoNestedPartyIDs), quickfix.GroupElement(tag.LegRefID), quickfix.GroupElement(tag.LegPrice), quickfix.GroupElement(tag.LegSettlType), quickfix.GroupElement(tag.LegSettlDate), quickfix.GroupElement(tag.LegLastPx)})}
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingSecurityIDSource), quickfix.GroupElement(tag.NoUnderlyingSecurityAltID), quickfix.GroupElement(tag.UnderlyingProduct), quickfix.GroupElement(tag.UnderlyingCFICode), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingSecuritySubType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDate), quickfix.GroupElement(tag.UnderlyingCouponPaymentDate), quickfix.GroupElement(tag.UnderlyingIssueDate), quickfix.GroupElement(tag.UnderlyingRepoCollateralSecurityType), quickfix.GroupElement(tag.UnderlyingRepurchaseTerm), quickfix.GroupElement(tag.UnderlyingRepurchaseRate), quickfix.GroupElement(tag.UnderlyingFactor), quickfix.GroupElement(tag.UnderlyingCreditRating), quickfix.GroupElement(tag.UnderlyingInstrRegistry), quickfix.GroupElement(tag.UnderlyingCountryOfIssue), quickfix.GroupElement(tag.UnderlyingStateOrProvinceOfIssue), quickfix.GroupElement(tag.UnderlyingLocaleOfIssue), quickfix.GroupElement(tag.UnderlyingRedemptionDate), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingStrikeCurrency), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.UnderlyingCPProgram), quickfix.GroupElement(tag.UnderlyingCPRegType), quickfix.GroupElement(tag.UnderlyingCurrency), quickfix.GroupElement(tag.UnderlyingQty), quickfix.GroupElement(tag.UnderlyingPx), quickfix.GroupElement(tag.UnderlyingDirtyPrice), quickfix.GroupElement(tag.UnderlyingEndPrice), quickfix.GroupElement(tag.UnderlyingStartValue), quickfix.GroupElement(tag.UnderlyingCurrentValue), quickfix.GroupElement(tag.UnderlyingEndValue), quickfix.GroupElement(tag.NoUnderlyingStips)})}
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

//NoPosAmt is a repeating group element, Tag 753
type NoPosAmt struct {
	quickfix.Group
}

//SetPosAmtType sets PosAmtType, Tag 707
func (m NoPosAmt) SetPosAmtType(v string) {
	m.Set(field.NewPosAmtType(v))
}

//SetPosAmt sets PosAmt, Tag 708
func (m NoPosAmt) SetPosAmt(v float64) {
	m.Set(field.NewPosAmt(v))
}

//GetPosAmtType gets PosAmtType, Tag 707
func (m NoPosAmt) GetPosAmtType() (f field.PosAmtTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPosAmt gets PosAmt, Tag 708
func (m NoPosAmt) GetPosAmt() (f field.PosAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasPosAmtType returns true if PosAmtType is present, Tag 707
func (m NoPosAmt) HasPosAmtType() bool {
	return m.Has(tag.PosAmtType)
}

//HasPosAmt returns true if PosAmt is present, Tag 708
func (m NoPosAmt) HasPosAmt() bool {
	return m.Has(tag.PosAmt)
}

//NoPosAmtRepeatingGroup is a repeating group, Tag 753
type NoPosAmtRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPosAmtRepeatingGroup returns an initialized, NoPosAmtRepeatingGroup
func NewNoPosAmtRepeatingGroup() NoPosAmtRepeatingGroup {
	return NoPosAmtRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPosAmt,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PosAmtType), quickfix.GroupElement(tag.PosAmt)})}
}

//Add create and append a new NoPosAmt to this group
func (m NoPosAmtRepeatingGroup) Add() NoPosAmt {
	g := m.RepeatingGroup.Add()
	return NoPosAmt{g}
}

//Get returns the ith NoPosAmt in the NoPosAmtRepeatinGroup
func (m NoPosAmtRepeatingGroup) Get(i int) NoPosAmt {
	return NoPosAmt{m.RepeatingGroup.Get(i)}
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

//NoTrdRegTimestampsRepeatingGroup is a repeating group, Tag 768
type NoTrdRegTimestampsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTrdRegTimestampsRepeatingGroup returns an initialized, NoTrdRegTimestampsRepeatingGroup
func NewNoTrdRegTimestampsRepeatingGroup() NoTrdRegTimestampsRepeatingGroup {
	return NoTrdRegTimestampsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTrdRegTimestamps,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TrdRegTimestamp), quickfix.GroupElement(tag.TrdRegTimestampType), quickfix.GroupElement(tag.TrdRegTimestampOrigin)})}
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
