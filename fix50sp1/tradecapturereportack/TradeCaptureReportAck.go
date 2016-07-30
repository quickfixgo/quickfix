package tradecapturereportack

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//TradeCaptureReportAck is the fix50sp1 TradeCaptureReportAck type, MsgType = AR
type TradeCaptureReportAck struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a TradeCaptureReportAck from a quickfix.Message instance
func FromMessage(m quickfix.Message) TradeCaptureReportAck {
	return TradeCaptureReportAck{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m TradeCaptureReportAck) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a TradeCaptureReportAck initialized with the required fields for TradeCaptureReportAck
func New() (m TradeCaptureReportAck) {
	m.Header = fixt11.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("AR"))
	m.Header.Set(field.NewBeginString("8"))

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg TradeCaptureReportAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "8", "AR", r
}

//SetAvgPx sets AvgPx, Tag 6
func (m TradeCaptureReportAck) SetAvgPx(v float64) {
	m.Set(field.NewAvgPx(v))
}

//SetCurrency sets Currency, Tag 15
func (m TradeCaptureReportAck) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetExecID sets ExecID, Tag 17
func (m TradeCaptureReportAck) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m TradeCaptureReportAck) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetLastMkt sets LastMkt, Tag 30
func (m TradeCaptureReportAck) SetLastMkt(v string) {
	m.Set(field.NewLastMkt(v))
}

//SetLastPx sets LastPx, Tag 31
func (m TradeCaptureReportAck) SetLastPx(v float64) {
	m.Set(field.NewLastPx(v))
}

//SetLastQty sets LastQty, Tag 32
func (m TradeCaptureReportAck) SetLastQty(v float64) {
	m.Set(field.NewLastQty(v))
}

//SetOrdStatus sets OrdStatus, Tag 39
func (m TradeCaptureReportAck) SetOrdStatus(v string) {
	m.Set(field.NewOrdStatus(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m TradeCaptureReportAck) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m TradeCaptureReportAck) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m TradeCaptureReportAck) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m TradeCaptureReportAck) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSettlType sets SettlType, Tag 63
func (m TradeCaptureReportAck) SetSettlType(v string) {
	m.Set(field.NewSettlType(v))
}

//SetSettlDate sets SettlDate, Tag 64
func (m TradeCaptureReportAck) SetSettlDate(v string) {
	m.Set(field.NewSettlDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m TradeCaptureReportAck) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m TradeCaptureReportAck) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetIssuer sets Issuer, Tag 106
func (m TradeCaptureReportAck) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m TradeCaptureReportAck) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m TradeCaptureReportAck) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetExecType sets ExecType, Tag 150
func (m TradeCaptureReportAck) SetExecType(v string) {
	m.Set(field.NewExecType(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m TradeCaptureReportAck) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetLastSpotRate sets LastSpotRate, Tag 194
func (m TradeCaptureReportAck) SetLastSpotRate(v float64) {
	m.Set(field.NewLastSpotRate(v))
}

//SetLastForwardPoints sets LastForwardPoints, Tag 195
func (m TradeCaptureReportAck) SetLastForwardPoints(v float64) {
	m.Set(field.NewLastForwardPoints(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m TradeCaptureReportAck) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m TradeCaptureReportAck) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m TradeCaptureReportAck) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m TradeCaptureReportAck) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m TradeCaptureReportAck) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m TradeCaptureReportAck) SetCouponRate(v float64) {
	m.Set(field.NewCouponRate(v))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m TradeCaptureReportAck) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m TradeCaptureReportAck) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m TradeCaptureReportAck) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m TradeCaptureReportAck) SetRepurchaseRate(v float64) {
	m.Set(field.NewRepurchaseRate(v))
}

//SetFactor sets Factor, Tag 228
func (m TradeCaptureReportAck) SetFactor(v float64) {
	m.Set(field.NewFactor(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m TradeCaptureReportAck) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m TradeCaptureReportAck) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m TradeCaptureReportAck) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m TradeCaptureReportAck) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263
func (m TradeCaptureReportAck) SetSubscriptionRequestType(v string) {
	m.Set(field.NewSubscriptionRequestType(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m TradeCaptureReportAck) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m TradeCaptureReportAck) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m TradeCaptureReportAck) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m TradeCaptureReportAck) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m TradeCaptureReportAck) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m TradeCaptureReportAck) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetExecRestatementReason sets ExecRestatementReason, Tag 378
func (m TradeCaptureReportAck) SetExecRestatementReason(v int) {
	m.Set(field.NewExecRestatementReason(v))
}

//SetGrossTradeAmt sets GrossTradeAmt, Tag 381
func (m TradeCaptureReportAck) SetGrossTradeAmt(v float64) {
	m.Set(field.NewGrossTradeAmt(v))
}

//SetPriceType sets PriceType, Tag 423
func (m TradeCaptureReportAck) SetPriceType(v int) {
	m.Set(field.NewPriceType(v))
}

//SetMultiLegReportingType sets MultiLegReportingType, Tag 442
func (m TradeCaptureReportAck) SetMultiLegReportingType(v string) {
	m.Set(field.NewMultiLegReportingType(v))
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m TradeCaptureReportAck) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m TradeCaptureReportAck) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m TradeCaptureReportAck) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m TradeCaptureReportAck) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m TradeCaptureReportAck) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m TradeCaptureReportAck) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetTradeReportTransType sets TradeReportTransType, Tag 487
func (m TradeCaptureReportAck) SetTradeReportTransType(v int) {
	m.Set(field.NewTradeReportTransType(v))
}

//SetSecondaryExecID sets SecondaryExecID, Tag 527
func (m TradeCaptureReportAck) SetSecondaryExecID(v string) {
	m.Set(field.NewSecondaryExecID(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m TradeCaptureReportAck) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m TradeCaptureReportAck) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetNoSides sets NoSides, Tag 552
func (m TradeCaptureReportAck) SetNoSides(f NoSidesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoLegs sets NoLegs, Tag 555
func (m TradeCaptureReportAck) SetNoLegs(f NoLegsRepeatingGroup) {
	m.SetGroup(f)
}

//SetPreviouslyReported sets PreviouslyReported, Tag 570
func (m TradeCaptureReportAck) SetPreviouslyReported(v bool) {
	m.Set(field.NewPreviouslyReported(v))
}

//SetTradeReportID sets TradeReportID, Tag 571
func (m TradeCaptureReportAck) SetTradeReportID(v string) {
	m.Set(field.NewTradeReportID(v))
}

//SetTradeReportRefID sets TradeReportRefID, Tag 572
func (m TradeCaptureReportAck) SetTradeReportRefID(v string) {
	m.Set(field.NewTradeReportRefID(v))
}

//SetMatchStatus sets MatchStatus, Tag 573
func (m TradeCaptureReportAck) SetMatchStatus(v string) {
	m.Set(field.NewMatchStatus(v))
}

//SetMatchType sets MatchType, Tag 574
func (m TradeCaptureReportAck) SetMatchType(v string) {
	m.Set(field.NewMatchType(v))
}

//SetClearingFeeIndicator sets ClearingFeeIndicator, Tag 635
func (m TradeCaptureReportAck) SetClearingFeeIndicator(v string) {
	m.Set(field.NewClearingFeeIndicator(v))
}

//SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m TradeCaptureReportAck) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

//SetLastParPx sets LastParPx, Tag 669
func (m TradeCaptureReportAck) SetLastParPx(v float64) {
	m.Set(field.NewLastParPx(v))
}

//SetPool sets Pool, Tag 691
func (m TradeCaptureReportAck) SetPool(v string) {
	m.Set(field.NewPool(v))
}

//SetNoUnderlyings sets NoUnderlyings, Tag 711
func (m TradeCaptureReportAck) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

//SetClearingBusinessDate sets ClearingBusinessDate, Tag 715
func (m TradeCaptureReportAck) SetClearingBusinessDate(v string) {
	m.Set(field.NewClearingBusinessDate(v))
}

//SetSettlSessID sets SettlSessID, Tag 716
func (m TradeCaptureReportAck) SetSettlSessID(v string) {
	m.Set(field.NewSettlSessID(v))
}

//SetSettlSessSubID sets SettlSessSubID, Tag 717
func (m TradeCaptureReportAck) SetSettlSessSubID(v string) {
	m.Set(field.NewSettlSessSubID(v))
}

//SetResponseTransportType sets ResponseTransportType, Tag 725
func (m TradeCaptureReportAck) SetResponseTransportType(v int) {
	m.Set(field.NewResponseTransportType(v))
}

//SetResponseDestination sets ResponseDestination, Tag 726
func (m TradeCaptureReportAck) SetResponseDestination(v string) {
	m.Set(field.NewResponseDestination(v))
}

//SetTradeReportRejectReason sets TradeReportRejectReason, Tag 751
func (m TradeCaptureReportAck) SetTradeReportRejectReason(v int) {
	m.Set(field.NewTradeReportRejectReason(v))
}

//SetNoPosAmt sets NoPosAmt, Tag 753
func (m TradeCaptureReportAck) SetNoPosAmt(f NoPosAmtRepeatingGroup) {
	m.SetGroup(f)
}

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m TradeCaptureReportAck) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

//SetNoTrdRegTimestamps sets NoTrdRegTimestamps, Tag 768
func (m TradeCaptureReportAck) SetNoTrdRegTimestamps(f NoTrdRegTimestampsRepeatingGroup) {
	m.SetGroup(f)
}

//SetLastUpdateTime sets LastUpdateTime, Tag 779
func (m TradeCaptureReportAck) SetLastUpdateTime(v time.Time) {
	m.Set(field.NewLastUpdateTime(v))
}

//SetCopyMsgIndicator sets CopyMsgIndicator, Tag 797
func (m TradeCaptureReportAck) SetCopyMsgIndicator(v bool) {
	m.Set(field.NewCopyMsgIndicator(v))
}

//SetSecondaryTradeReportID sets SecondaryTradeReportID, Tag 818
func (m TradeCaptureReportAck) SetSecondaryTradeReportID(v string) {
	m.Set(field.NewSecondaryTradeReportID(v))
}

//SetAvgPxIndicator sets AvgPxIndicator, Tag 819
func (m TradeCaptureReportAck) SetAvgPxIndicator(v int) {
	m.Set(field.NewAvgPxIndicator(v))
}

//SetTradeLinkID sets TradeLinkID, Tag 820
func (m TradeCaptureReportAck) SetTradeLinkID(v string) {
	m.Set(field.NewTradeLinkID(v))
}

//SetUnderlyingTradingSessionID sets UnderlyingTradingSessionID, Tag 822
func (m TradeCaptureReportAck) SetUnderlyingTradingSessionID(v string) {
	m.Set(field.NewUnderlyingTradingSessionID(v))
}

//SetUnderlyingTradingSessionSubID sets UnderlyingTradingSessionSubID, Tag 823
func (m TradeCaptureReportAck) SetUnderlyingTradingSessionSubID(v string) {
	m.Set(field.NewUnderlyingTradingSessionSubID(v))
}

//SetTradeLegRefID sets TradeLegRefID, Tag 824
func (m TradeCaptureReportAck) SetTradeLegRefID(v string) {
	m.Set(field.NewTradeLegRefID(v))
}

//SetTrdType sets TrdType, Tag 828
func (m TradeCaptureReportAck) SetTrdType(v int) {
	m.Set(field.NewTrdType(v))
}

//SetTrdSubType sets TrdSubType, Tag 829
func (m TradeCaptureReportAck) SetTrdSubType(v int) {
	m.Set(field.NewTrdSubType(v))
}

//SetTransferReason sets TransferReason, Tag 830
func (m TradeCaptureReportAck) SetTransferReason(v string) {
	m.Set(field.NewTransferReason(v))
}

//SetPublishTrdIndicator sets PublishTrdIndicator, Tag 852
func (m TradeCaptureReportAck) SetPublishTrdIndicator(v bool) {
	m.Set(field.NewPublishTrdIndicator(v))
}

//SetShortSaleReason sets ShortSaleReason, Tag 853
func (m TradeCaptureReportAck) SetShortSaleReason(v int) {
	m.Set(field.NewShortSaleReason(v))
}

//SetQtyType sets QtyType, Tag 854
func (m TradeCaptureReportAck) SetQtyType(v int) {
	m.Set(field.NewQtyType(v))
}

//SetSecondaryTrdType sets SecondaryTrdType, Tag 855
func (m TradeCaptureReportAck) SetSecondaryTrdType(v int) {
	m.Set(field.NewSecondaryTrdType(v))
}

//SetTradeReportType sets TradeReportType, Tag 856
func (m TradeCaptureReportAck) SetTradeReportType(v int) {
	m.Set(field.NewTradeReportType(v))
}

//SetNoEvents sets NoEvents, Tag 864
func (m TradeCaptureReportAck) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDatedDate sets DatedDate, Tag 873
func (m TradeCaptureReportAck) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

//SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m TradeCaptureReportAck) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

//SetCPProgram sets CPProgram, Tag 875
func (m TradeCaptureReportAck) SetCPProgram(v int) {
	m.Set(field.NewCPProgram(v))
}

//SetCPRegType sets CPRegType, Tag 876
func (m TradeCaptureReportAck) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

//SetTrdMatchID sets TrdMatchID, Tag 880
func (m TradeCaptureReportAck) SetTrdMatchID(v string) {
	m.Set(field.NewTrdMatchID(v))
}

//SetSecondaryTradeReportRefID sets SecondaryTradeReportRefID, Tag 881
func (m TradeCaptureReportAck) SetSecondaryTradeReportRefID(v string) {
	m.Set(field.NewSecondaryTradeReportRefID(v))
}

//SetTrdRptStatus sets TrdRptStatus, Tag 939
func (m TradeCaptureReportAck) SetTrdRptStatus(v int) {
	m.Set(field.NewTrdRptStatus(v))
}

//SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m TradeCaptureReportAck) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

//SetSecurityStatus sets SecurityStatus, Tag 965
func (m TradeCaptureReportAck) SetSecurityStatus(v string) {
	m.Set(field.NewSecurityStatus(v))
}

//SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966
func (m TradeCaptureReportAck) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

//SetStrikeMultiplier sets StrikeMultiplier, Tag 967
func (m TradeCaptureReportAck) SetStrikeMultiplier(v float64) {
	m.Set(field.NewStrikeMultiplier(v))
}

//SetStrikeValue sets StrikeValue, Tag 968
func (m TradeCaptureReportAck) SetStrikeValue(v float64) {
	m.Set(field.NewStrikeValue(v))
}

//SetMinPriceIncrement sets MinPriceIncrement, Tag 969
func (m TradeCaptureReportAck) SetMinPriceIncrement(v float64) {
	m.Set(field.NewMinPriceIncrement(v))
}

//SetPositionLimit sets PositionLimit, Tag 970
func (m TradeCaptureReportAck) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

//SetNTPositionLimit sets NTPositionLimit, Tag 971
func (m TradeCaptureReportAck) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

//SetRndPx sets RndPx, Tag 991
func (m TradeCaptureReportAck) SetRndPx(v float64) {
	m.Set(field.NewRndPx(v))
}

//SetTierCode sets TierCode, Tag 994
func (m TradeCaptureReportAck) SetTierCode(v string) {
	m.Set(field.NewTierCode(v))
}

//SetUnitOfMeasure sets UnitOfMeasure, Tag 996
func (m TradeCaptureReportAck) SetUnitOfMeasure(v string) {
	m.Set(field.NewUnitOfMeasure(v))
}

//SetTimeUnit sets TimeUnit, Tag 997
func (m TradeCaptureReportAck) SetTimeUnit(v string) {
	m.Set(field.NewTimeUnit(v))
}

//SetTradeID sets TradeID, Tag 1003
func (m TradeCaptureReportAck) SetTradeID(v string) {
	m.Set(field.NewTradeID(v))
}

//SetMessageEventSource sets MessageEventSource, Tag 1011
func (m TradeCaptureReportAck) SetMessageEventSource(v string) {
	m.Set(field.NewMessageEventSource(v))
}

//SetAsOfIndicator sets AsOfIndicator, Tag 1015
func (m TradeCaptureReportAck) SetAsOfIndicator(v string) {
	m.Set(field.NewAsOfIndicator(v))
}

//SetNoInstrumentParties sets NoInstrumentParties, Tag 1018
func (m TradeCaptureReportAck) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetSecondaryTradeID sets SecondaryTradeID, Tag 1040
func (m TradeCaptureReportAck) SetSecondaryTradeID(v string) {
	m.Set(field.NewSecondaryTradeID(v))
}

//SetFirmTradeID sets FirmTradeID, Tag 1041
func (m TradeCaptureReportAck) SetFirmTradeID(v string) {
	m.Set(field.NewFirmTradeID(v))
}

//SetSecondaryFirmTradeID sets SecondaryFirmTradeID, Tag 1042
func (m TradeCaptureReportAck) SetSecondaryFirmTradeID(v string) {
	m.Set(field.NewSecondaryFirmTradeID(v))
}

//SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049
func (m TradeCaptureReportAck) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

//SetCalculatedCcyLastQty sets CalculatedCcyLastQty, Tag 1056
func (m TradeCaptureReportAck) SetCalculatedCcyLastQty(v float64) {
	m.Set(field.NewCalculatedCcyLastQty(v))
}

//SetLastSwapPoints sets LastSwapPoints, Tag 1071
func (m TradeCaptureReportAck) SetLastSwapPoints(v float64) {
	m.Set(field.NewLastSwapPoints(v))
}

//SetMaturityTime sets MaturityTime, Tag 1079
func (m TradeCaptureReportAck) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

//SetNoRootPartyIDs sets NoRootPartyIDs, Tag 1116
func (m TradeCaptureReportAck) SetNoRootPartyIDs(f NoRootPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetTradeHandlingInstr sets TradeHandlingInstr, Tag 1123
func (m TradeCaptureReportAck) SetTradeHandlingInstr(v string) {
	m.Set(field.NewTradeHandlingInstr(v))
}

//SetOrigTradeHandlingInstr sets OrigTradeHandlingInstr, Tag 1124
func (m TradeCaptureReportAck) SetOrigTradeHandlingInstr(v string) {
	m.Set(field.NewOrigTradeHandlingInstr(v))
}

//SetOrigTradeDate sets OrigTradeDate, Tag 1125
func (m TradeCaptureReportAck) SetOrigTradeDate(v string) {
	m.Set(field.NewOrigTradeDate(v))
}

//SetOrigTradeID sets OrigTradeID, Tag 1126
func (m TradeCaptureReportAck) SetOrigTradeID(v string) {
	m.Set(field.NewOrigTradeID(v))
}

//SetOrigSecondaryTradeID sets OrigSecondaryTradeID, Tag 1127
func (m TradeCaptureReportAck) SetOrigSecondaryTradeID(v string) {
	m.Set(field.NewOrigSecondaryTradeID(v))
}

//SetRptSys sets RptSys, Tag 1135
func (m TradeCaptureReportAck) SetRptSys(v string) {
	m.Set(field.NewRptSys(v))
}

//SetMinPriceIncrementAmount sets MinPriceIncrementAmount, Tag 1146
func (m TradeCaptureReportAck) SetMinPriceIncrementAmount(v float64) {
	m.Set(field.NewMinPriceIncrementAmount(v))
}

//SetUnitOfMeasureQty sets UnitOfMeasureQty, Tag 1147
func (m TradeCaptureReportAck) SetUnitOfMeasureQty(v float64) {
	m.Set(field.NewUnitOfMeasureQty(v))
}

//SetSecurityGroup sets SecurityGroup, Tag 1151
func (m TradeCaptureReportAck) SetSecurityGroup(v string) {
	m.Set(field.NewSecurityGroup(v))
}

//SetSecurityXMLLen sets SecurityXMLLen, Tag 1184
func (m TradeCaptureReportAck) SetSecurityXMLLen(v int) {
	m.Set(field.NewSecurityXMLLen(v))
}

//SetSecurityXML sets SecurityXML, Tag 1185
func (m TradeCaptureReportAck) SetSecurityXML(v string) {
	m.Set(field.NewSecurityXML(v))
}

//SetSecurityXMLSchema sets SecurityXMLSchema, Tag 1186
func (m TradeCaptureReportAck) SetSecurityXMLSchema(v string) {
	m.Set(field.NewSecurityXMLSchema(v))
}

//SetPriceUnitOfMeasure sets PriceUnitOfMeasure, Tag 1191
func (m TradeCaptureReportAck) SetPriceUnitOfMeasure(v string) {
	m.Set(field.NewPriceUnitOfMeasure(v))
}

//SetPriceUnitOfMeasureQty sets PriceUnitOfMeasureQty, Tag 1192
func (m TradeCaptureReportAck) SetPriceUnitOfMeasureQty(v float64) {
	m.Set(field.NewPriceUnitOfMeasureQty(v))
}

//SetSettlMethod sets SettlMethod, Tag 1193
func (m TradeCaptureReportAck) SetSettlMethod(v string) {
	m.Set(field.NewSettlMethod(v))
}

//SetExerciseStyle sets ExerciseStyle, Tag 1194
func (m TradeCaptureReportAck) SetExerciseStyle(v int) {
	m.Set(field.NewExerciseStyle(v))
}

//SetOptPayAmount sets OptPayAmount, Tag 1195
func (m TradeCaptureReportAck) SetOptPayAmount(v float64) {
	m.Set(field.NewOptPayAmount(v))
}

//SetPriceQuoteMethod sets PriceQuoteMethod, Tag 1196
func (m TradeCaptureReportAck) SetPriceQuoteMethod(v string) {
	m.Set(field.NewPriceQuoteMethod(v))
}

//SetFuturesValuationMethod sets FuturesValuationMethod, Tag 1197
func (m TradeCaptureReportAck) SetFuturesValuationMethod(v string) {
	m.Set(field.NewFuturesValuationMethod(v))
}

//SetListMethod sets ListMethod, Tag 1198
func (m TradeCaptureReportAck) SetListMethod(v int) {
	m.Set(field.NewListMethod(v))
}

//SetCapPrice sets CapPrice, Tag 1199
func (m TradeCaptureReportAck) SetCapPrice(v float64) {
	m.Set(field.NewCapPrice(v))
}

//SetFloorPrice sets FloorPrice, Tag 1200
func (m TradeCaptureReportAck) SetFloorPrice(v float64) {
	m.Set(field.NewFloorPrice(v))
}

//SetProductComplex sets ProductComplex, Tag 1227
func (m TradeCaptureReportAck) SetProductComplex(v string) {
	m.Set(field.NewProductComplex(v))
}

//SetFlexProductEligibilityIndicator sets FlexProductEligibilityIndicator, Tag 1242
func (m TradeCaptureReportAck) SetFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewFlexProductEligibilityIndicator(v))
}

//SetFlexibleIndicator sets FlexibleIndicator, Tag 1244
func (m TradeCaptureReportAck) SetFlexibleIndicator(v bool) {
	m.Set(field.NewFlexibleIndicator(v))
}

//SetFeeMultiplier sets FeeMultiplier, Tag 1329
func (m TradeCaptureReportAck) SetFeeMultiplier(v float64) {
	m.Set(field.NewFeeMultiplier(v))
}

//SetNoTrdRepIndicators sets NoTrdRepIndicators, Tag 1387
func (m TradeCaptureReportAck) SetNoTrdRepIndicators(f NoTrdRepIndicatorsRepeatingGroup) {
	m.SetGroup(f)
}

//SetTradePublishIndicator sets TradePublishIndicator, Tag 1390
func (m TradeCaptureReportAck) SetTradePublishIndicator(v int) {
	m.Set(field.NewTradePublishIndicator(v))
}

//GetAvgPx gets AvgPx, Tag 6
func (m TradeCaptureReportAck) GetAvgPx() (f field.AvgPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m TradeCaptureReportAck) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecID gets ExecID, Tag 17
func (m TradeCaptureReportAck) GetExecID() (f field.ExecIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m TradeCaptureReportAck) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastMkt gets LastMkt, Tag 30
func (m TradeCaptureReportAck) GetLastMkt() (f field.LastMktField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastPx gets LastPx, Tag 31
func (m TradeCaptureReportAck) GetLastPx() (f field.LastPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastQty gets LastQty, Tag 32
func (m TradeCaptureReportAck) GetLastQty() (f field.LastQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdStatus gets OrdStatus, Tag 39
func (m TradeCaptureReportAck) GetOrdStatus() (f field.OrdStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m TradeCaptureReportAck) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m TradeCaptureReportAck) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m TradeCaptureReportAck) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m TradeCaptureReportAck) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlType gets SettlType, Tag 63
func (m TradeCaptureReportAck) GetSettlType() (f field.SettlTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlDate gets SettlDate, Tag 64
func (m TradeCaptureReportAck) GetSettlDate() (f field.SettlDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m TradeCaptureReportAck) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m TradeCaptureReportAck) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m TradeCaptureReportAck) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m TradeCaptureReportAck) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m TradeCaptureReportAck) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecType gets ExecType, Tag 150
func (m TradeCaptureReportAck) GetExecType() (f field.ExecTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m TradeCaptureReportAck) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastSpotRate gets LastSpotRate, Tag 194
func (m TradeCaptureReportAck) GetLastSpotRate() (f field.LastSpotRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastForwardPoints gets LastForwardPoints, Tag 195
func (m TradeCaptureReportAck) GetLastForwardPoints() (f field.LastForwardPointsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m TradeCaptureReportAck) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m TradeCaptureReportAck) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m TradeCaptureReportAck) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m TradeCaptureReportAck) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m TradeCaptureReportAck) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m TradeCaptureReportAck) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m TradeCaptureReportAck) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m TradeCaptureReportAck) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m TradeCaptureReportAck) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m TradeCaptureReportAck) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m TradeCaptureReportAck) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m TradeCaptureReportAck) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m TradeCaptureReportAck) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m TradeCaptureReportAck) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m TradeCaptureReportAck) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263
func (m TradeCaptureReportAck) GetSubscriptionRequestType() (f field.SubscriptionRequestTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m TradeCaptureReportAck) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m TradeCaptureReportAck) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m TradeCaptureReportAck) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m TradeCaptureReportAck) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m TradeCaptureReportAck) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m TradeCaptureReportAck) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecRestatementReason gets ExecRestatementReason, Tag 378
func (m TradeCaptureReportAck) GetExecRestatementReason() (f field.ExecRestatementReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetGrossTradeAmt gets GrossTradeAmt, Tag 381
func (m TradeCaptureReportAck) GetGrossTradeAmt() (f field.GrossTradeAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceType gets PriceType, Tag 423
func (m TradeCaptureReportAck) GetPriceType() (f field.PriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMultiLegReportingType gets MultiLegReportingType, Tag 442
func (m TradeCaptureReportAck) GetMultiLegReportingType() (f field.MultiLegReportingTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m TradeCaptureReportAck) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m TradeCaptureReportAck) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m TradeCaptureReportAck) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m TradeCaptureReportAck) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m TradeCaptureReportAck) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m TradeCaptureReportAck) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeReportTransType gets TradeReportTransType, Tag 487
func (m TradeCaptureReportAck) GetTradeReportTransType() (f field.TradeReportTransTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryExecID gets SecondaryExecID, Tag 527
func (m TradeCaptureReportAck) GetSecondaryExecID() (f field.SecondaryExecIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m TradeCaptureReportAck) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m TradeCaptureReportAck) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSides gets NoSides, Tag 552
func (m TradeCaptureReportAck) GetNoSides() (NoSidesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSidesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoLegs gets NoLegs, Tag 555
func (m TradeCaptureReportAck) GetNoLegs() (NoLegsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetPreviouslyReported gets PreviouslyReported, Tag 570
func (m TradeCaptureReportAck) GetPreviouslyReported() (f field.PreviouslyReportedField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeReportID gets TradeReportID, Tag 571
func (m TradeCaptureReportAck) GetTradeReportID() (f field.TradeReportIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeReportRefID gets TradeReportRefID, Tag 572
func (m TradeCaptureReportAck) GetTradeReportRefID() (f field.TradeReportRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMatchStatus gets MatchStatus, Tag 573
func (m TradeCaptureReportAck) GetMatchStatus() (f field.MatchStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMatchType gets MatchType, Tag 574
func (m TradeCaptureReportAck) GetMatchType() (f field.MatchTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClearingFeeIndicator gets ClearingFeeIndicator, Tag 635
func (m TradeCaptureReportAck) GetClearingFeeIndicator() (f field.ClearingFeeIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m TradeCaptureReportAck) GetContractSettlMonth() (f field.ContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastParPx gets LastParPx, Tag 669
func (m TradeCaptureReportAck) GetLastParPx() (f field.LastParPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPool gets Pool, Tag 691
func (m TradeCaptureReportAck) GetPool() (f field.PoolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyings gets NoUnderlyings, Tag 711
func (m TradeCaptureReportAck) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetClearingBusinessDate gets ClearingBusinessDate, Tag 715
func (m TradeCaptureReportAck) GetClearingBusinessDate() (f field.ClearingBusinessDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlSessID gets SettlSessID, Tag 716
func (m TradeCaptureReportAck) GetSettlSessID() (f field.SettlSessIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlSessSubID gets SettlSessSubID, Tag 717
func (m TradeCaptureReportAck) GetSettlSessSubID() (f field.SettlSessSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetResponseTransportType gets ResponseTransportType, Tag 725
func (m TradeCaptureReportAck) GetResponseTransportType() (f field.ResponseTransportTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetResponseDestination gets ResponseDestination, Tag 726
func (m TradeCaptureReportAck) GetResponseDestination() (f field.ResponseDestinationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeReportRejectReason gets TradeReportRejectReason, Tag 751
func (m TradeCaptureReportAck) GetTradeReportRejectReason() (f field.TradeReportRejectReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPosAmt gets NoPosAmt, Tag 753
func (m TradeCaptureReportAck) GetNoPosAmt() (NoPosAmtRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPosAmtRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m TradeCaptureReportAck) GetSecuritySubType() (f field.SecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoTrdRegTimestamps gets NoTrdRegTimestamps, Tag 768
func (m TradeCaptureReportAck) GetNoTrdRegTimestamps() (NoTrdRegTimestampsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTrdRegTimestampsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLastUpdateTime gets LastUpdateTime, Tag 779
func (m TradeCaptureReportAck) GetLastUpdateTime() (f field.LastUpdateTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCopyMsgIndicator gets CopyMsgIndicator, Tag 797
func (m TradeCaptureReportAck) GetCopyMsgIndicator() (f field.CopyMsgIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryTradeReportID gets SecondaryTradeReportID, Tag 818
func (m TradeCaptureReportAck) GetSecondaryTradeReportID() (f field.SecondaryTradeReportIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAvgPxIndicator gets AvgPxIndicator, Tag 819
func (m TradeCaptureReportAck) GetAvgPxIndicator() (f field.AvgPxIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeLinkID gets TradeLinkID, Tag 820
func (m TradeCaptureReportAck) GetTradeLinkID() (f field.TradeLinkIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingTradingSessionID gets UnderlyingTradingSessionID, Tag 822
func (m TradeCaptureReportAck) GetUnderlyingTradingSessionID() (f field.UnderlyingTradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingTradingSessionSubID gets UnderlyingTradingSessionSubID, Tag 823
func (m TradeCaptureReportAck) GetUnderlyingTradingSessionSubID() (f field.UnderlyingTradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeLegRefID gets TradeLegRefID, Tag 824
func (m TradeCaptureReportAck) GetTradeLegRefID() (f field.TradeLegRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTrdType gets TrdType, Tag 828
func (m TradeCaptureReportAck) GetTrdType() (f field.TrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTrdSubType gets TrdSubType, Tag 829
func (m TradeCaptureReportAck) GetTrdSubType() (f field.TrdSubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransferReason gets TransferReason, Tag 830
func (m TradeCaptureReportAck) GetTransferReason() (f field.TransferReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPublishTrdIndicator gets PublishTrdIndicator, Tag 852
func (m TradeCaptureReportAck) GetPublishTrdIndicator() (f field.PublishTrdIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetShortSaleReason gets ShortSaleReason, Tag 853
func (m TradeCaptureReportAck) GetShortSaleReason() (f field.ShortSaleReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQtyType gets QtyType, Tag 854
func (m TradeCaptureReportAck) GetQtyType() (f field.QtyTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryTrdType gets SecondaryTrdType, Tag 855
func (m TradeCaptureReportAck) GetSecondaryTrdType() (f field.SecondaryTrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeReportType gets TradeReportType, Tag 856
func (m TradeCaptureReportAck) GetTradeReportType() (f field.TradeReportTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoEvents gets NoEvents, Tag 864
func (m TradeCaptureReportAck) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDatedDate gets DatedDate, Tag 873
func (m TradeCaptureReportAck) GetDatedDate() (f field.DatedDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m TradeCaptureReportAck) GetInterestAccrualDate() (f field.InterestAccrualDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPProgram gets CPProgram, Tag 875
func (m TradeCaptureReportAck) GetCPProgram() (f field.CPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPRegType gets CPRegType, Tag 876
func (m TradeCaptureReportAck) GetCPRegType() (f field.CPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTrdMatchID gets TrdMatchID, Tag 880
func (m TradeCaptureReportAck) GetTrdMatchID() (f field.TrdMatchIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryTradeReportRefID gets SecondaryTradeReportRefID, Tag 881
func (m TradeCaptureReportAck) GetSecondaryTradeReportRefID() (f field.SecondaryTradeReportRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTrdRptStatus gets TrdRptStatus, Tag 939
func (m TradeCaptureReportAck) GetTrdRptStatus() (f field.TrdRptStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m TradeCaptureReportAck) GetStrikeCurrency() (f field.StrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityStatus gets SecurityStatus, Tag 965
func (m TradeCaptureReportAck) GetSecurityStatus() (f field.SecurityStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966
func (m TradeCaptureReportAck) GetSettleOnOpenFlag() (f field.SettleOnOpenFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeMultiplier gets StrikeMultiplier, Tag 967
func (m TradeCaptureReportAck) GetStrikeMultiplier() (f field.StrikeMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeValue gets StrikeValue, Tag 968
func (m TradeCaptureReportAck) GetStrikeValue() (f field.StrikeValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrement gets MinPriceIncrement, Tag 969
func (m TradeCaptureReportAck) GetMinPriceIncrement() (f field.MinPriceIncrementField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionLimit gets PositionLimit, Tag 970
func (m TradeCaptureReportAck) GetPositionLimit() (f field.PositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNTPositionLimit gets NTPositionLimit, Tag 971
func (m TradeCaptureReportAck) GetNTPositionLimit() (f field.NTPositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRndPx gets RndPx, Tag 991
func (m TradeCaptureReportAck) GetRndPx() (f field.RndPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTierCode gets TierCode, Tag 994
func (m TradeCaptureReportAck) GetTierCode() (f field.TierCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnitOfMeasure gets UnitOfMeasure, Tag 996
func (m TradeCaptureReportAck) GetUnitOfMeasure() (f field.UnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeUnit gets TimeUnit, Tag 997
func (m TradeCaptureReportAck) GetTimeUnit() (f field.TimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeID gets TradeID, Tag 1003
func (m TradeCaptureReportAck) GetTradeID() (f field.TradeIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMessageEventSource gets MessageEventSource, Tag 1011
func (m TradeCaptureReportAck) GetMessageEventSource() (f field.MessageEventSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAsOfIndicator gets AsOfIndicator, Tag 1015
func (m TradeCaptureReportAck) GetAsOfIndicator() (f field.AsOfIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoInstrumentParties gets NoInstrumentParties, Tag 1018
func (m TradeCaptureReportAck) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSecondaryTradeID gets SecondaryTradeID, Tag 1040
func (m TradeCaptureReportAck) GetSecondaryTradeID() (f field.SecondaryTradeIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFirmTradeID gets FirmTradeID, Tag 1041
func (m TradeCaptureReportAck) GetFirmTradeID() (f field.FirmTradeIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryFirmTradeID gets SecondaryFirmTradeID, Tag 1042
func (m TradeCaptureReportAck) GetSecondaryFirmTradeID() (f field.SecondaryFirmTradeIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049
func (m TradeCaptureReportAck) GetInstrmtAssignmentMethod() (f field.InstrmtAssignmentMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCalculatedCcyLastQty gets CalculatedCcyLastQty, Tag 1056
func (m TradeCaptureReportAck) GetCalculatedCcyLastQty() (f field.CalculatedCcyLastQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastSwapPoints gets LastSwapPoints, Tag 1071
func (m TradeCaptureReportAck) GetLastSwapPoints() (f field.LastSwapPointsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityTime gets MaturityTime, Tag 1079
func (m TradeCaptureReportAck) GetMaturityTime() (f field.MaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoRootPartyIDs gets NoRootPartyIDs, Tag 1116
func (m TradeCaptureReportAck) GetNoRootPartyIDs() (NoRootPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRootPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetTradeHandlingInstr gets TradeHandlingInstr, Tag 1123
func (m TradeCaptureReportAck) GetTradeHandlingInstr() (f field.TradeHandlingInstrField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigTradeHandlingInstr gets OrigTradeHandlingInstr, Tag 1124
func (m TradeCaptureReportAck) GetOrigTradeHandlingInstr() (f field.OrigTradeHandlingInstrField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigTradeDate gets OrigTradeDate, Tag 1125
func (m TradeCaptureReportAck) GetOrigTradeDate() (f field.OrigTradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigTradeID gets OrigTradeID, Tag 1126
func (m TradeCaptureReportAck) GetOrigTradeID() (f field.OrigTradeIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigSecondaryTradeID gets OrigSecondaryTradeID, Tag 1127
func (m TradeCaptureReportAck) GetOrigSecondaryTradeID() (f field.OrigSecondaryTradeIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRptSys gets RptSys, Tag 1135
func (m TradeCaptureReportAck) GetRptSys() (f field.RptSysField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrementAmount gets MinPriceIncrementAmount, Tag 1146
func (m TradeCaptureReportAck) GetMinPriceIncrementAmount() (f field.MinPriceIncrementAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnitOfMeasureQty gets UnitOfMeasureQty, Tag 1147
func (m TradeCaptureReportAck) GetUnitOfMeasureQty() (f field.UnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityGroup gets SecurityGroup, Tag 1151
func (m TradeCaptureReportAck) GetSecurityGroup() (f field.SecurityGroupField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXMLLen gets SecurityXMLLen, Tag 1184
func (m TradeCaptureReportAck) GetSecurityXMLLen() (f field.SecurityXMLLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXML gets SecurityXML, Tag 1185
func (m TradeCaptureReportAck) GetSecurityXML() (f field.SecurityXMLField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXMLSchema gets SecurityXMLSchema, Tag 1186
func (m TradeCaptureReportAck) GetSecurityXMLSchema() (f field.SecurityXMLSchemaField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceUnitOfMeasure gets PriceUnitOfMeasure, Tag 1191
func (m TradeCaptureReportAck) GetPriceUnitOfMeasure() (f field.PriceUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceUnitOfMeasureQty gets PriceUnitOfMeasureQty, Tag 1192
func (m TradeCaptureReportAck) GetPriceUnitOfMeasureQty() (f field.PriceUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlMethod gets SettlMethod, Tag 1193
func (m TradeCaptureReportAck) GetSettlMethod() (f field.SettlMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExerciseStyle gets ExerciseStyle, Tag 1194
func (m TradeCaptureReportAck) GetExerciseStyle() (f field.ExerciseStyleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptPayAmount gets OptPayAmount, Tag 1195
func (m TradeCaptureReportAck) GetOptPayAmount() (f field.OptPayAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceQuoteMethod gets PriceQuoteMethod, Tag 1196
func (m TradeCaptureReportAck) GetPriceQuoteMethod() (f field.PriceQuoteMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFuturesValuationMethod gets FuturesValuationMethod, Tag 1197
func (m TradeCaptureReportAck) GetFuturesValuationMethod() (f field.FuturesValuationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListMethod gets ListMethod, Tag 1198
func (m TradeCaptureReportAck) GetListMethod() (f field.ListMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCapPrice gets CapPrice, Tag 1199
func (m TradeCaptureReportAck) GetCapPrice() (f field.CapPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFloorPrice gets FloorPrice, Tag 1200
func (m TradeCaptureReportAck) GetFloorPrice() (f field.FloorPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProductComplex gets ProductComplex, Tag 1227
func (m TradeCaptureReportAck) GetProductComplex() (f field.ProductComplexField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlexProductEligibilityIndicator gets FlexProductEligibilityIndicator, Tag 1242
func (m TradeCaptureReportAck) GetFlexProductEligibilityIndicator() (f field.FlexProductEligibilityIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlexibleIndicator gets FlexibleIndicator, Tag 1244
func (m TradeCaptureReportAck) GetFlexibleIndicator() (f field.FlexibleIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFeeMultiplier gets FeeMultiplier, Tag 1329
func (m TradeCaptureReportAck) GetFeeMultiplier() (f field.FeeMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoTrdRepIndicators gets NoTrdRepIndicators, Tag 1387
func (m TradeCaptureReportAck) GetNoTrdRepIndicators() (NoTrdRepIndicatorsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTrdRepIndicatorsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetTradePublishIndicator gets TradePublishIndicator, Tag 1390
func (m TradeCaptureReportAck) GetTradePublishIndicator() (f field.TradePublishIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAvgPx returns true if AvgPx is present, Tag 6
func (m TradeCaptureReportAck) HasAvgPx() bool {
	return m.Has(tag.AvgPx)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m TradeCaptureReportAck) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasExecID returns true if ExecID is present, Tag 17
func (m TradeCaptureReportAck) HasExecID() bool {
	return m.Has(tag.ExecID)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m TradeCaptureReportAck) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasLastMkt returns true if LastMkt is present, Tag 30
func (m TradeCaptureReportAck) HasLastMkt() bool {
	return m.Has(tag.LastMkt)
}

//HasLastPx returns true if LastPx is present, Tag 31
func (m TradeCaptureReportAck) HasLastPx() bool {
	return m.Has(tag.LastPx)
}

//HasLastQty returns true if LastQty is present, Tag 32
func (m TradeCaptureReportAck) HasLastQty() bool {
	return m.Has(tag.LastQty)
}

//HasOrdStatus returns true if OrdStatus is present, Tag 39
func (m TradeCaptureReportAck) HasOrdStatus() bool {
	return m.Has(tag.OrdStatus)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m TradeCaptureReportAck) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m TradeCaptureReportAck) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m TradeCaptureReportAck) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m TradeCaptureReportAck) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSettlType returns true if SettlType is present, Tag 63
func (m TradeCaptureReportAck) HasSettlType() bool {
	return m.Has(tag.SettlType)
}

//HasSettlDate returns true if SettlDate is present, Tag 64
func (m TradeCaptureReportAck) HasSettlDate() bool {
	return m.Has(tag.SettlDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m TradeCaptureReportAck) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m TradeCaptureReportAck) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m TradeCaptureReportAck) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m TradeCaptureReportAck) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m TradeCaptureReportAck) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasExecType returns true if ExecType is present, Tag 150
func (m TradeCaptureReportAck) HasExecType() bool {
	return m.Has(tag.ExecType)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m TradeCaptureReportAck) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasLastSpotRate returns true if LastSpotRate is present, Tag 194
func (m TradeCaptureReportAck) HasLastSpotRate() bool {
	return m.Has(tag.LastSpotRate)
}

//HasLastForwardPoints returns true if LastForwardPoints is present, Tag 195
func (m TradeCaptureReportAck) HasLastForwardPoints() bool {
	return m.Has(tag.LastForwardPoints)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m TradeCaptureReportAck) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m TradeCaptureReportAck) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m TradeCaptureReportAck) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m TradeCaptureReportAck) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m TradeCaptureReportAck) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m TradeCaptureReportAck) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m TradeCaptureReportAck) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m TradeCaptureReportAck) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m TradeCaptureReportAck) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m TradeCaptureReportAck) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m TradeCaptureReportAck) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m TradeCaptureReportAck) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m TradeCaptureReportAck) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m TradeCaptureReportAck) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m TradeCaptureReportAck) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263
func (m TradeCaptureReportAck) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m TradeCaptureReportAck) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m TradeCaptureReportAck) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m TradeCaptureReportAck) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m TradeCaptureReportAck) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m TradeCaptureReportAck) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m TradeCaptureReportAck) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasExecRestatementReason returns true if ExecRestatementReason is present, Tag 378
func (m TradeCaptureReportAck) HasExecRestatementReason() bool {
	return m.Has(tag.ExecRestatementReason)
}

//HasGrossTradeAmt returns true if GrossTradeAmt is present, Tag 381
func (m TradeCaptureReportAck) HasGrossTradeAmt() bool {
	return m.Has(tag.GrossTradeAmt)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m TradeCaptureReportAck) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasMultiLegReportingType returns true if MultiLegReportingType is present, Tag 442
func (m TradeCaptureReportAck) HasMultiLegReportingType() bool {
	return m.Has(tag.MultiLegReportingType)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m TradeCaptureReportAck) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m TradeCaptureReportAck) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m TradeCaptureReportAck) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m TradeCaptureReportAck) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m TradeCaptureReportAck) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m TradeCaptureReportAck) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasTradeReportTransType returns true if TradeReportTransType is present, Tag 487
func (m TradeCaptureReportAck) HasTradeReportTransType() bool {
	return m.Has(tag.TradeReportTransType)
}

//HasSecondaryExecID returns true if SecondaryExecID is present, Tag 527
func (m TradeCaptureReportAck) HasSecondaryExecID() bool {
	return m.Has(tag.SecondaryExecID)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m TradeCaptureReportAck) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m TradeCaptureReportAck) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasNoSides returns true if NoSides is present, Tag 552
func (m TradeCaptureReportAck) HasNoSides() bool {
	return m.Has(tag.NoSides)
}

//HasNoLegs returns true if NoLegs is present, Tag 555
func (m TradeCaptureReportAck) HasNoLegs() bool {
	return m.Has(tag.NoLegs)
}

//HasPreviouslyReported returns true if PreviouslyReported is present, Tag 570
func (m TradeCaptureReportAck) HasPreviouslyReported() bool {
	return m.Has(tag.PreviouslyReported)
}

//HasTradeReportID returns true if TradeReportID is present, Tag 571
func (m TradeCaptureReportAck) HasTradeReportID() bool {
	return m.Has(tag.TradeReportID)
}

//HasTradeReportRefID returns true if TradeReportRefID is present, Tag 572
func (m TradeCaptureReportAck) HasTradeReportRefID() bool {
	return m.Has(tag.TradeReportRefID)
}

//HasMatchStatus returns true if MatchStatus is present, Tag 573
func (m TradeCaptureReportAck) HasMatchStatus() bool {
	return m.Has(tag.MatchStatus)
}

//HasMatchType returns true if MatchType is present, Tag 574
func (m TradeCaptureReportAck) HasMatchType() bool {
	return m.Has(tag.MatchType)
}

//HasClearingFeeIndicator returns true if ClearingFeeIndicator is present, Tag 635
func (m TradeCaptureReportAck) HasClearingFeeIndicator() bool {
	return m.Has(tag.ClearingFeeIndicator)
}

//HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m TradeCaptureReportAck) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

//HasLastParPx returns true if LastParPx is present, Tag 669
func (m TradeCaptureReportAck) HasLastParPx() bool {
	return m.Has(tag.LastParPx)
}

//HasPool returns true if Pool is present, Tag 691
func (m TradeCaptureReportAck) HasPool() bool {
	return m.Has(tag.Pool)
}

//HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711
func (m TradeCaptureReportAck) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

//HasClearingBusinessDate returns true if ClearingBusinessDate is present, Tag 715
func (m TradeCaptureReportAck) HasClearingBusinessDate() bool {
	return m.Has(tag.ClearingBusinessDate)
}

//HasSettlSessID returns true if SettlSessID is present, Tag 716
func (m TradeCaptureReportAck) HasSettlSessID() bool {
	return m.Has(tag.SettlSessID)
}

//HasSettlSessSubID returns true if SettlSessSubID is present, Tag 717
func (m TradeCaptureReportAck) HasSettlSessSubID() bool {
	return m.Has(tag.SettlSessSubID)
}

//HasResponseTransportType returns true if ResponseTransportType is present, Tag 725
func (m TradeCaptureReportAck) HasResponseTransportType() bool {
	return m.Has(tag.ResponseTransportType)
}

//HasResponseDestination returns true if ResponseDestination is present, Tag 726
func (m TradeCaptureReportAck) HasResponseDestination() bool {
	return m.Has(tag.ResponseDestination)
}

//HasTradeReportRejectReason returns true if TradeReportRejectReason is present, Tag 751
func (m TradeCaptureReportAck) HasTradeReportRejectReason() bool {
	return m.Has(tag.TradeReportRejectReason)
}

//HasNoPosAmt returns true if NoPosAmt is present, Tag 753
func (m TradeCaptureReportAck) HasNoPosAmt() bool {
	return m.Has(tag.NoPosAmt)
}

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m TradeCaptureReportAck) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

//HasNoTrdRegTimestamps returns true if NoTrdRegTimestamps is present, Tag 768
func (m TradeCaptureReportAck) HasNoTrdRegTimestamps() bool {
	return m.Has(tag.NoTrdRegTimestamps)
}

//HasLastUpdateTime returns true if LastUpdateTime is present, Tag 779
func (m TradeCaptureReportAck) HasLastUpdateTime() bool {
	return m.Has(tag.LastUpdateTime)
}

//HasCopyMsgIndicator returns true if CopyMsgIndicator is present, Tag 797
func (m TradeCaptureReportAck) HasCopyMsgIndicator() bool {
	return m.Has(tag.CopyMsgIndicator)
}

//HasSecondaryTradeReportID returns true if SecondaryTradeReportID is present, Tag 818
func (m TradeCaptureReportAck) HasSecondaryTradeReportID() bool {
	return m.Has(tag.SecondaryTradeReportID)
}

//HasAvgPxIndicator returns true if AvgPxIndicator is present, Tag 819
func (m TradeCaptureReportAck) HasAvgPxIndicator() bool {
	return m.Has(tag.AvgPxIndicator)
}

//HasTradeLinkID returns true if TradeLinkID is present, Tag 820
func (m TradeCaptureReportAck) HasTradeLinkID() bool {
	return m.Has(tag.TradeLinkID)
}

//HasUnderlyingTradingSessionID returns true if UnderlyingTradingSessionID is present, Tag 822
func (m TradeCaptureReportAck) HasUnderlyingTradingSessionID() bool {
	return m.Has(tag.UnderlyingTradingSessionID)
}

//HasUnderlyingTradingSessionSubID returns true if UnderlyingTradingSessionSubID is present, Tag 823
func (m TradeCaptureReportAck) HasUnderlyingTradingSessionSubID() bool {
	return m.Has(tag.UnderlyingTradingSessionSubID)
}

//HasTradeLegRefID returns true if TradeLegRefID is present, Tag 824
func (m TradeCaptureReportAck) HasTradeLegRefID() bool {
	return m.Has(tag.TradeLegRefID)
}

//HasTrdType returns true if TrdType is present, Tag 828
func (m TradeCaptureReportAck) HasTrdType() bool {
	return m.Has(tag.TrdType)
}

//HasTrdSubType returns true if TrdSubType is present, Tag 829
func (m TradeCaptureReportAck) HasTrdSubType() bool {
	return m.Has(tag.TrdSubType)
}

//HasTransferReason returns true if TransferReason is present, Tag 830
func (m TradeCaptureReportAck) HasTransferReason() bool {
	return m.Has(tag.TransferReason)
}

//HasPublishTrdIndicator returns true if PublishTrdIndicator is present, Tag 852
func (m TradeCaptureReportAck) HasPublishTrdIndicator() bool {
	return m.Has(tag.PublishTrdIndicator)
}

//HasShortSaleReason returns true if ShortSaleReason is present, Tag 853
func (m TradeCaptureReportAck) HasShortSaleReason() bool {
	return m.Has(tag.ShortSaleReason)
}

//HasQtyType returns true if QtyType is present, Tag 854
func (m TradeCaptureReportAck) HasQtyType() bool {
	return m.Has(tag.QtyType)
}

//HasSecondaryTrdType returns true if SecondaryTrdType is present, Tag 855
func (m TradeCaptureReportAck) HasSecondaryTrdType() bool {
	return m.Has(tag.SecondaryTrdType)
}

//HasTradeReportType returns true if TradeReportType is present, Tag 856
func (m TradeCaptureReportAck) HasTradeReportType() bool {
	return m.Has(tag.TradeReportType)
}

//HasNoEvents returns true if NoEvents is present, Tag 864
func (m TradeCaptureReportAck) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

//HasDatedDate returns true if DatedDate is present, Tag 873
func (m TradeCaptureReportAck) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

//HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m TradeCaptureReportAck) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

//HasCPProgram returns true if CPProgram is present, Tag 875
func (m TradeCaptureReportAck) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

//HasCPRegType returns true if CPRegType is present, Tag 876
func (m TradeCaptureReportAck) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

//HasTrdMatchID returns true if TrdMatchID is present, Tag 880
func (m TradeCaptureReportAck) HasTrdMatchID() bool {
	return m.Has(tag.TrdMatchID)
}

//HasSecondaryTradeReportRefID returns true if SecondaryTradeReportRefID is present, Tag 881
func (m TradeCaptureReportAck) HasSecondaryTradeReportRefID() bool {
	return m.Has(tag.SecondaryTradeReportRefID)
}

//HasTrdRptStatus returns true if TrdRptStatus is present, Tag 939
func (m TradeCaptureReportAck) HasTrdRptStatus() bool {
	return m.Has(tag.TrdRptStatus)
}

//HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m TradeCaptureReportAck) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

//HasSecurityStatus returns true if SecurityStatus is present, Tag 965
func (m TradeCaptureReportAck) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

//HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966
func (m TradeCaptureReportAck) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

//HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967
func (m TradeCaptureReportAck) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

//HasStrikeValue returns true if StrikeValue is present, Tag 968
func (m TradeCaptureReportAck) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

//HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969
func (m TradeCaptureReportAck) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

//HasPositionLimit returns true if PositionLimit is present, Tag 970
func (m TradeCaptureReportAck) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

//HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971
func (m TradeCaptureReportAck) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

//HasRndPx returns true if RndPx is present, Tag 991
func (m TradeCaptureReportAck) HasRndPx() bool {
	return m.Has(tag.RndPx)
}

//HasTierCode returns true if TierCode is present, Tag 994
func (m TradeCaptureReportAck) HasTierCode() bool {
	return m.Has(tag.TierCode)
}

//HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996
func (m TradeCaptureReportAck) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

//HasTimeUnit returns true if TimeUnit is present, Tag 997
func (m TradeCaptureReportAck) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

//HasTradeID returns true if TradeID is present, Tag 1003
func (m TradeCaptureReportAck) HasTradeID() bool {
	return m.Has(tag.TradeID)
}

//HasMessageEventSource returns true if MessageEventSource is present, Tag 1011
func (m TradeCaptureReportAck) HasMessageEventSource() bool {
	return m.Has(tag.MessageEventSource)
}

//HasAsOfIndicator returns true if AsOfIndicator is present, Tag 1015
func (m TradeCaptureReportAck) HasAsOfIndicator() bool {
	return m.Has(tag.AsOfIndicator)
}

//HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018
func (m TradeCaptureReportAck) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

//HasSecondaryTradeID returns true if SecondaryTradeID is present, Tag 1040
func (m TradeCaptureReportAck) HasSecondaryTradeID() bool {
	return m.Has(tag.SecondaryTradeID)
}

//HasFirmTradeID returns true if FirmTradeID is present, Tag 1041
func (m TradeCaptureReportAck) HasFirmTradeID() bool {
	return m.Has(tag.FirmTradeID)
}

//HasSecondaryFirmTradeID returns true if SecondaryFirmTradeID is present, Tag 1042
func (m TradeCaptureReportAck) HasSecondaryFirmTradeID() bool {
	return m.Has(tag.SecondaryFirmTradeID)
}

//HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049
func (m TradeCaptureReportAck) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

//HasCalculatedCcyLastQty returns true if CalculatedCcyLastQty is present, Tag 1056
func (m TradeCaptureReportAck) HasCalculatedCcyLastQty() bool {
	return m.Has(tag.CalculatedCcyLastQty)
}

//HasLastSwapPoints returns true if LastSwapPoints is present, Tag 1071
func (m TradeCaptureReportAck) HasLastSwapPoints() bool {
	return m.Has(tag.LastSwapPoints)
}

//HasMaturityTime returns true if MaturityTime is present, Tag 1079
func (m TradeCaptureReportAck) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

//HasNoRootPartyIDs returns true if NoRootPartyIDs is present, Tag 1116
func (m TradeCaptureReportAck) HasNoRootPartyIDs() bool {
	return m.Has(tag.NoRootPartyIDs)
}

//HasTradeHandlingInstr returns true if TradeHandlingInstr is present, Tag 1123
func (m TradeCaptureReportAck) HasTradeHandlingInstr() bool {
	return m.Has(tag.TradeHandlingInstr)
}

//HasOrigTradeHandlingInstr returns true if OrigTradeHandlingInstr is present, Tag 1124
func (m TradeCaptureReportAck) HasOrigTradeHandlingInstr() bool {
	return m.Has(tag.OrigTradeHandlingInstr)
}

//HasOrigTradeDate returns true if OrigTradeDate is present, Tag 1125
func (m TradeCaptureReportAck) HasOrigTradeDate() bool {
	return m.Has(tag.OrigTradeDate)
}

//HasOrigTradeID returns true if OrigTradeID is present, Tag 1126
func (m TradeCaptureReportAck) HasOrigTradeID() bool {
	return m.Has(tag.OrigTradeID)
}

//HasOrigSecondaryTradeID returns true if OrigSecondaryTradeID is present, Tag 1127
func (m TradeCaptureReportAck) HasOrigSecondaryTradeID() bool {
	return m.Has(tag.OrigSecondaryTradeID)
}

//HasRptSys returns true if RptSys is present, Tag 1135
func (m TradeCaptureReportAck) HasRptSys() bool {
	return m.Has(tag.RptSys)
}

//HasMinPriceIncrementAmount returns true if MinPriceIncrementAmount is present, Tag 1146
func (m TradeCaptureReportAck) HasMinPriceIncrementAmount() bool {
	return m.Has(tag.MinPriceIncrementAmount)
}

//HasUnitOfMeasureQty returns true if UnitOfMeasureQty is present, Tag 1147
func (m TradeCaptureReportAck) HasUnitOfMeasureQty() bool {
	return m.Has(tag.UnitOfMeasureQty)
}

//HasSecurityGroup returns true if SecurityGroup is present, Tag 1151
func (m TradeCaptureReportAck) HasSecurityGroup() bool {
	return m.Has(tag.SecurityGroup)
}

//HasSecurityXMLLen returns true if SecurityXMLLen is present, Tag 1184
func (m TradeCaptureReportAck) HasSecurityXMLLen() bool {
	return m.Has(tag.SecurityXMLLen)
}

//HasSecurityXML returns true if SecurityXML is present, Tag 1185
func (m TradeCaptureReportAck) HasSecurityXML() bool {
	return m.Has(tag.SecurityXML)
}

//HasSecurityXMLSchema returns true if SecurityXMLSchema is present, Tag 1186
func (m TradeCaptureReportAck) HasSecurityXMLSchema() bool {
	return m.Has(tag.SecurityXMLSchema)
}

//HasPriceUnitOfMeasure returns true if PriceUnitOfMeasure is present, Tag 1191
func (m TradeCaptureReportAck) HasPriceUnitOfMeasure() bool {
	return m.Has(tag.PriceUnitOfMeasure)
}

//HasPriceUnitOfMeasureQty returns true if PriceUnitOfMeasureQty is present, Tag 1192
func (m TradeCaptureReportAck) HasPriceUnitOfMeasureQty() bool {
	return m.Has(tag.PriceUnitOfMeasureQty)
}

//HasSettlMethod returns true if SettlMethod is present, Tag 1193
func (m TradeCaptureReportAck) HasSettlMethod() bool {
	return m.Has(tag.SettlMethod)
}

//HasExerciseStyle returns true if ExerciseStyle is present, Tag 1194
func (m TradeCaptureReportAck) HasExerciseStyle() bool {
	return m.Has(tag.ExerciseStyle)
}

//HasOptPayAmount returns true if OptPayAmount is present, Tag 1195
func (m TradeCaptureReportAck) HasOptPayAmount() bool {
	return m.Has(tag.OptPayAmount)
}

//HasPriceQuoteMethod returns true if PriceQuoteMethod is present, Tag 1196
func (m TradeCaptureReportAck) HasPriceQuoteMethod() bool {
	return m.Has(tag.PriceQuoteMethod)
}

//HasFuturesValuationMethod returns true if FuturesValuationMethod is present, Tag 1197
func (m TradeCaptureReportAck) HasFuturesValuationMethod() bool {
	return m.Has(tag.FuturesValuationMethod)
}

//HasListMethod returns true if ListMethod is present, Tag 1198
func (m TradeCaptureReportAck) HasListMethod() bool {
	return m.Has(tag.ListMethod)
}

//HasCapPrice returns true if CapPrice is present, Tag 1199
func (m TradeCaptureReportAck) HasCapPrice() bool {
	return m.Has(tag.CapPrice)
}

//HasFloorPrice returns true if FloorPrice is present, Tag 1200
func (m TradeCaptureReportAck) HasFloorPrice() bool {
	return m.Has(tag.FloorPrice)
}

//HasProductComplex returns true if ProductComplex is present, Tag 1227
func (m TradeCaptureReportAck) HasProductComplex() bool {
	return m.Has(tag.ProductComplex)
}

//HasFlexProductEligibilityIndicator returns true if FlexProductEligibilityIndicator is present, Tag 1242
func (m TradeCaptureReportAck) HasFlexProductEligibilityIndicator() bool {
	return m.Has(tag.FlexProductEligibilityIndicator)
}

//HasFlexibleIndicator returns true if FlexibleIndicator is present, Tag 1244
func (m TradeCaptureReportAck) HasFlexibleIndicator() bool {
	return m.Has(tag.FlexibleIndicator)
}

//HasFeeMultiplier returns true if FeeMultiplier is present, Tag 1329
func (m TradeCaptureReportAck) HasFeeMultiplier() bool {
	return m.Has(tag.FeeMultiplier)
}

//HasNoTrdRepIndicators returns true if NoTrdRepIndicators is present, Tag 1387
func (m TradeCaptureReportAck) HasNoTrdRepIndicators() bool {
	return m.Has(tag.NoTrdRepIndicators)
}

//HasTradePublishIndicator returns true if TradePublishIndicator is present, Tag 1390
func (m TradeCaptureReportAck) HasTradePublishIndicator() bool {
	return m.Has(tag.TradePublishIndicator)
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

//SetLotType sets LotType, Tag 1093
func (m NoSides) SetLotType(v string) {
	m.Set(field.NewLotType(v))
}

//SetSideGrossTradeAmt sets SideGrossTradeAmt, Tag 1072
func (m NoSides) SetSideGrossTradeAmt(v float64) {
	m.Set(field.NewSideGrossTradeAmt(v))
}

//SetAggressorIndicator sets AggressorIndicator, Tag 1057
func (m NoSides) SetAggressorIndicator(v bool) {
	m.Set(field.NewAggressorIndicator(v))
}

//SetSideQty sets SideQty, Tag 1009
func (m NoSides) SetSideQty(v int) {
	m.Set(field.NewSideQty(v))
}

//SetSideTradeReportID sets SideTradeReportID, Tag 1005
func (m NoSides) SetSideTradeReportID(v string) {
	m.Set(field.NewSideTradeReportID(v))
}

//SetSideFillStationCd sets SideFillStationCd, Tag 1006
func (m NoSides) SetSideFillStationCd(v string) {
	m.Set(field.NewSideFillStationCd(v))
}

//SetSideReasonCd sets SideReasonCd, Tag 1007
func (m NoSides) SetSideReasonCd(v string) {
	m.Set(field.NewSideReasonCd(v))
}

//SetRptSeq sets RptSeq, Tag 83
func (m NoSides) SetRptSeq(v int) {
	m.Set(field.NewRptSeq(v))
}

//SetSideTrdSubTyp sets SideTrdSubTyp, Tag 1008
func (m NoSides) SetSideTrdSubTyp(v int) {
	m.Set(field.NewSideTrdSubTyp(v))
}

//SetNoSideTrdRegTS sets NoSideTrdRegTS, Tag 1016
func (m NoSides) SetNoSideTrdRegTS(f NoSideTrdRegTSRepeatingGroup) {
	m.SetGroup(f)
}

//SetNetGrossInd sets NetGrossInd, Tag 430
func (m NoSides) SetNetGrossInd(v int) {
	m.Set(field.NewNetGrossInd(v))
}

//SetSideCurrency sets SideCurrency, Tag 1154
func (m NoSides) SetSideCurrency(v string) {
	m.Set(field.NewSideCurrency(v))
}

//SetSideSettlCurrency sets SideSettlCurrency, Tag 1155
func (m NoSides) SetSideSettlCurrency(v string) {
	m.Set(field.NewSideSettlCurrency(v))
}

//SetNoSettlDetails sets NoSettlDetails, Tag 1158
func (m NoSides) SetNoSettlDetails(f NoSettlDetailsRepeatingGroup) {
	m.SetGroup(f)
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

//GetLotType gets LotType, Tag 1093
func (m NoSides) GetLotType() (f field.LotTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSideGrossTradeAmt gets SideGrossTradeAmt, Tag 1072
func (m NoSides) GetSideGrossTradeAmt() (f field.SideGrossTradeAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAggressorIndicator gets AggressorIndicator, Tag 1057
func (m NoSides) GetAggressorIndicator() (f field.AggressorIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSideQty gets SideQty, Tag 1009
func (m NoSides) GetSideQty() (f field.SideQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSideTradeReportID gets SideTradeReportID, Tag 1005
func (m NoSides) GetSideTradeReportID() (f field.SideTradeReportIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSideFillStationCd gets SideFillStationCd, Tag 1006
func (m NoSides) GetSideFillStationCd() (f field.SideFillStationCdField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSideReasonCd gets SideReasonCd, Tag 1007
func (m NoSides) GetSideReasonCd() (f field.SideReasonCdField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRptSeq gets RptSeq, Tag 83
func (m NoSides) GetRptSeq() (f field.RptSeqField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSideTrdSubTyp gets SideTrdSubTyp, Tag 1008
func (m NoSides) GetSideTrdSubTyp() (f field.SideTrdSubTypField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSideTrdRegTS gets NoSideTrdRegTS, Tag 1016
func (m NoSides) GetNoSideTrdRegTS() (NoSideTrdRegTSRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSideTrdRegTSRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNetGrossInd gets NetGrossInd, Tag 430
func (m NoSides) GetNetGrossInd() (f field.NetGrossIndField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSideCurrency gets SideCurrency, Tag 1154
func (m NoSides) GetSideCurrency() (f field.SideCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSideSettlCurrency gets SideSettlCurrency, Tag 1155
func (m NoSides) GetSideSettlCurrency() (f field.SideSettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSettlDetails gets NoSettlDetails, Tag 1158
func (m NoSides) GetNoSettlDetails() (NoSettlDetailsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSettlDetailsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
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

//HasLotType returns true if LotType is present, Tag 1093
func (m NoSides) HasLotType() bool {
	return m.Has(tag.LotType)
}

//HasSideGrossTradeAmt returns true if SideGrossTradeAmt is present, Tag 1072
func (m NoSides) HasSideGrossTradeAmt() bool {
	return m.Has(tag.SideGrossTradeAmt)
}

//HasAggressorIndicator returns true if AggressorIndicator is present, Tag 1057
func (m NoSides) HasAggressorIndicator() bool {
	return m.Has(tag.AggressorIndicator)
}

//HasSideQty returns true if SideQty is present, Tag 1009
func (m NoSides) HasSideQty() bool {
	return m.Has(tag.SideQty)
}

//HasSideTradeReportID returns true if SideTradeReportID is present, Tag 1005
func (m NoSides) HasSideTradeReportID() bool {
	return m.Has(tag.SideTradeReportID)
}

//HasSideFillStationCd returns true if SideFillStationCd is present, Tag 1006
func (m NoSides) HasSideFillStationCd() bool {
	return m.Has(tag.SideFillStationCd)
}

//HasSideReasonCd returns true if SideReasonCd is present, Tag 1007
func (m NoSides) HasSideReasonCd() bool {
	return m.Has(tag.SideReasonCd)
}

//HasRptSeq returns true if RptSeq is present, Tag 83
func (m NoSides) HasRptSeq() bool {
	return m.Has(tag.RptSeq)
}

//HasSideTrdSubTyp returns true if SideTrdSubTyp is present, Tag 1008
func (m NoSides) HasSideTrdSubTyp() bool {
	return m.Has(tag.SideTrdSubTyp)
}

//HasNoSideTrdRegTS returns true if NoSideTrdRegTS is present, Tag 1016
func (m NoSides) HasNoSideTrdRegTS() bool {
	return m.Has(tag.NoSideTrdRegTS)
}

//HasNetGrossInd returns true if NetGrossInd is present, Tag 430
func (m NoSides) HasNetGrossInd() bool {
	return m.Has(tag.NetGrossInd)
}

//HasSideCurrency returns true if SideCurrency is present, Tag 1154
func (m NoSides) HasSideCurrency() bool {
	return m.Has(tag.SideCurrency)
}

//HasSideSettlCurrency returns true if SideSettlCurrency is present, Tag 1155
func (m NoSides) HasSideSettlCurrency() bool {
	return m.Has(tag.SideSettlCurrency)
}

//HasNoSettlDetails returns true if NoSettlDetails is present, Tag 1158
func (m NoSides) HasNoSettlDetails() bool {
	return m.Has(tag.NoSettlDetails)
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

//SetAllocCustomerCapacity sets AllocCustomerCapacity, Tag 993
func (m NoAllocs) SetAllocCustomerCapacity(v string) {
	m.Set(field.NewAllocCustomerCapacity(v))
}

//SetAllocMethod sets AllocMethod, Tag 1002
func (m NoAllocs) SetAllocMethod(v int) {
	m.Set(field.NewAllocMethod(v))
}

//SetSecondaryIndividualAllocID sets SecondaryIndividualAllocID, Tag 989
func (m NoAllocs) SetSecondaryIndividualAllocID(v string) {
	m.Set(field.NewSecondaryIndividualAllocID(v))
}

//SetAllocClearingFeeIndicator sets AllocClearingFeeIndicator, Tag 1136
func (m NoAllocs) SetAllocClearingFeeIndicator(v string) {
	m.Set(field.NewAllocClearingFeeIndicator(v))
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

//GetAllocCustomerCapacity gets AllocCustomerCapacity, Tag 993
func (m NoAllocs) GetAllocCustomerCapacity() (f field.AllocCustomerCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocMethod gets AllocMethod, Tag 1002
func (m NoAllocs) GetAllocMethod() (f field.AllocMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryIndividualAllocID gets SecondaryIndividualAllocID, Tag 989
func (m NoAllocs) GetSecondaryIndividualAllocID() (f field.SecondaryIndividualAllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocClearingFeeIndicator gets AllocClearingFeeIndicator, Tag 1136
func (m NoAllocs) GetAllocClearingFeeIndicator() (f field.AllocClearingFeeIndicatorField, err quickfix.MessageRejectError) {
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

//HasAllocCustomerCapacity returns true if AllocCustomerCapacity is present, Tag 993
func (m NoAllocs) HasAllocCustomerCapacity() bool {
	return m.Has(tag.AllocCustomerCapacity)
}

//HasAllocMethod returns true if AllocMethod is present, Tag 1002
func (m NoAllocs) HasAllocMethod() bool {
	return m.Has(tag.AllocMethod)
}

//HasSecondaryIndividualAllocID returns true if SecondaryIndividualAllocID is present, Tag 989
func (m NoAllocs) HasSecondaryIndividualAllocID() bool {
	return m.Has(tag.SecondaryIndividualAllocID)
}

//HasAllocClearingFeeIndicator returns true if AllocClearingFeeIndicator is present, Tag 1136
func (m NoAllocs) HasAllocClearingFeeIndicator() bool {
	return m.Has(tag.AllocClearingFeeIndicator)
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

//NoAllocsRepeatingGroup is a repeating group, Tag 78
type NoAllocsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoAllocsRepeatingGroup returns an initialized, NoAllocsRepeatingGroup
func NewNoAllocsRepeatingGroup() NoAllocsRepeatingGroup {
	return NoAllocsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAllocs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AllocAccount), quickfix.GroupElement(tag.AllocAcctIDSource), quickfix.GroupElement(tag.AllocSettlCurrency), quickfix.GroupElement(tag.IndividualAllocID), NewNoNested2PartyIDsRepeatingGroup(), quickfix.GroupElement(tag.AllocQty), quickfix.GroupElement(tag.AllocCustomerCapacity), quickfix.GroupElement(tag.AllocMethod), quickfix.GroupElement(tag.SecondaryIndividualAllocID), quickfix.GroupElement(tag.AllocClearingFeeIndicator)})}
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

//NoSideTrdRegTS is a repeating group element, Tag 1016
type NoSideTrdRegTS struct {
	quickfix.Group
}

//SetSideTrdRegTimestamp sets SideTrdRegTimestamp, Tag 1012
func (m NoSideTrdRegTS) SetSideTrdRegTimestamp(v time.Time) {
	m.Set(field.NewSideTrdRegTimestamp(v))
}

//SetSideTrdRegTimestampType sets SideTrdRegTimestampType, Tag 1013
func (m NoSideTrdRegTS) SetSideTrdRegTimestampType(v int) {
	m.Set(field.NewSideTrdRegTimestampType(v))
}

//SetSideTrdRegTimestampSrc sets SideTrdRegTimestampSrc, Tag 1014
func (m NoSideTrdRegTS) SetSideTrdRegTimestampSrc(v string) {
	m.Set(field.NewSideTrdRegTimestampSrc(v))
}

//GetSideTrdRegTimestamp gets SideTrdRegTimestamp, Tag 1012
func (m NoSideTrdRegTS) GetSideTrdRegTimestamp() (f field.SideTrdRegTimestampField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSideTrdRegTimestampType gets SideTrdRegTimestampType, Tag 1013
func (m NoSideTrdRegTS) GetSideTrdRegTimestampType() (f field.SideTrdRegTimestampTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSideTrdRegTimestampSrc gets SideTrdRegTimestampSrc, Tag 1014
func (m NoSideTrdRegTS) GetSideTrdRegTimestampSrc() (f field.SideTrdRegTimestampSrcField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasSideTrdRegTimestamp returns true if SideTrdRegTimestamp is present, Tag 1012
func (m NoSideTrdRegTS) HasSideTrdRegTimestamp() bool {
	return m.Has(tag.SideTrdRegTimestamp)
}

//HasSideTrdRegTimestampType returns true if SideTrdRegTimestampType is present, Tag 1013
func (m NoSideTrdRegTS) HasSideTrdRegTimestampType() bool {
	return m.Has(tag.SideTrdRegTimestampType)
}

//HasSideTrdRegTimestampSrc returns true if SideTrdRegTimestampSrc is present, Tag 1014
func (m NoSideTrdRegTS) HasSideTrdRegTimestampSrc() bool {
	return m.Has(tag.SideTrdRegTimestampSrc)
}

//NoSideTrdRegTSRepeatingGroup is a repeating group, Tag 1016
type NoSideTrdRegTSRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSideTrdRegTSRepeatingGroup returns an initialized, NoSideTrdRegTSRepeatingGroup
func NewNoSideTrdRegTSRepeatingGroup() NoSideTrdRegTSRepeatingGroup {
	return NoSideTrdRegTSRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSideTrdRegTS,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SideTrdRegTimestamp), quickfix.GroupElement(tag.SideTrdRegTimestampType), quickfix.GroupElement(tag.SideTrdRegTimestampSrc)})}
}

//Add create and append a new NoSideTrdRegTS to this group
func (m NoSideTrdRegTSRepeatingGroup) Add() NoSideTrdRegTS {
	g := m.RepeatingGroup.Add()
	return NoSideTrdRegTS{g}
}

//Get returns the ith NoSideTrdRegTS in the NoSideTrdRegTSRepeatinGroup
func (m NoSideTrdRegTSRepeatingGroup) Get(i int) NoSideTrdRegTS {
	return NoSideTrdRegTS{m.RepeatingGroup.Get(i)}
}

//NoSettlDetails is a repeating group element, Tag 1158
type NoSettlDetails struct {
	quickfix.Group
}

//SetSettlObligSource sets SettlObligSource, Tag 1164
func (m NoSettlDetails) SetSettlObligSource(v string) {
	m.Set(field.NewSettlObligSource(v))
}

//SetNoSettlPartyIDs sets NoSettlPartyIDs, Tag 781
func (m NoSettlDetails) SetNoSettlPartyIDs(f NoSettlPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetSettlObligSource gets SettlObligSource, Tag 1164
func (m NoSettlDetails) GetSettlObligSource() (f field.SettlObligSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSettlPartyIDs gets NoSettlPartyIDs, Tag 781
func (m NoSettlDetails) GetNoSettlPartyIDs() (NoSettlPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSettlPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasSettlObligSource returns true if SettlObligSource is present, Tag 1164
func (m NoSettlDetails) HasSettlObligSource() bool {
	return m.Has(tag.SettlObligSource)
}

//HasNoSettlPartyIDs returns true if NoSettlPartyIDs is present, Tag 781
func (m NoSettlDetails) HasNoSettlPartyIDs() bool {
	return m.Has(tag.NoSettlPartyIDs)
}

//NoSettlPartyIDs is a repeating group element, Tag 781
type NoSettlPartyIDs struct {
	quickfix.Group
}

//SetSettlPartyID sets SettlPartyID, Tag 782
func (m NoSettlPartyIDs) SetSettlPartyID(v string) {
	m.Set(field.NewSettlPartyID(v))
}

//SetSettlPartyIDSource sets SettlPartyIDSource, Tag 783
func (m NoSettlPartyIDs) SetSettlPartyIDSource(v string) {
	m.Set(field.NewSettlPartyIDSource(v))
}

//SetSettlPartyRole sets SettlPartyRole, Tag 784
func (m NoSettlPartyIDs) SetSettlPartyRole(v int) {
	m.Set(field.NewSettlPartyRole(v))
}

//SetNoSettlPartySubIDs sets NoSettlPartySubIDs, Tag 801
func (m NoSettlPartyIDs) SetNoSettlPartySubIDs(f NoSettlPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetSettlPartyID gets SettlPartyID, Tag 782
func (m NoSettlPartyIDs) GetSettlPartyID() (f field.SettlPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlPartyIDSource gets SettlPartyIDSource, Tag 783
func (m NoSettlPartyIDs) GetSettlPartyIDSource() (f field.SettlPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlPartyRole gets SettlPartyRole, Tag 784
func (m NoSettlPartyIDs) GetSettlPartyRole() (f field.SettlPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSettlPartySubIDs gets NoSettlPartySubIDs, Tag 801
func (m NoSettlPartyIDs) GetNoSettlPartySubIDs() (NoSettlPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSettlPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasSettlPartyID returns true if SettlPartyID is present, Tag 782
func (m NoSettlPartyIDs) HasSettlPartyID() bool {
	return m.Has(tag.SettlPartyID)
}

//HasSettlPartyIDSource returns true if SettlPartyIDSource is present, Tag 783
func (m NoSettlPartyIDs) HasSettlPartyIDSource() bool {
	return m.Has(tag.SettlPartyIDSource)
}

//HasSettlPartyRole returns true if SettlPartyRole is present, Tag 784
func (m NoSettlPartyIDs) HasSettlPartyRole() bool {
	return m.Has(tag.SettlPartyRole)
}

//HasNoSettlPartySubIDs returns true if NoSettlPartySubIDs is present, Tag 801
func (m NoSettlPartyIDs) HasNoSettlPartySubIDs() bool {
	return m.Has(tag.NoSettlPartySubIDs)
}

//NoSettlPartySubIDs is a repeating group element, Tag 801
type NoSettlPartySubIDs struct {
	quickfix.Group
}

//SetSettlPartySubID sets SettlPartySubID, Tag 785
func (m NoSettlPartySubIDs) SetSettlPartySubID(v string) {
	m.Set(field.NewSettlPartySubID(v))
}

//SetSettlPartySubIDType sets SettlPartySubIDType, Tag 786
func (m NoSettlPartySubIDs) SetSettlPartySubIDType(v int) {
	m.Set(field.NewSettlPartySubIDType(v))
}

//GetSettlPartySubID gets SettlPartySubID, Tag 785
func (m NoSettlPartySubIDs) GetSettlPartySubID() (f field.SettlPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlPartySubIDType gets SettlPartySubIDType, Tag 786
func (m NoSettlPartySubIDs) GetSettlPartySubIDType() (f field.SettlPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasSettlPartySubID returns true if SettlPartySubID is present, Tag 785
func (m NoSettlPartySubIDs) HasSettlPartySubID() bool {
	return m.Has(tag.SettlPartySubID)
}

//HasSettlPartySubIDType returns true if SettlPartySubIDType is present, Tag 786
func (m NoSettlPartySubIDs) HasSettlPartySubIDType() bool {
	return m.Has(tag.SettlPartySubIDType)
}

//NoSettlPartySubIDsRepeatingGroup is a repeating group, Tag 801
type NoSettlPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSettlPartySubIDsRepeatingGroup returns an initialized, NoSettlPartySubIDsRepeatingGroup
func NewNoSettlPartySubIDsRepeatingGroup() NoSettlPartySubIDsRepeatingGroup {
	return NoSettlPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSettlPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SettlPartySubID), quickfix.GroupElement(tag.SettlPartySubIDType)})}
}

//Add create and append a new NoSettlPartySubIDs to this group
func (m NoSettlPartySubIDsRepeatingGroup) Add() NoSettlPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoSettlPartySubIDs{g}
}

//Get returns the ith NoSettlPartySubIDs in the NoSettlPartySubIDsRepeatinGroup
func (m NoSettlPartySubIDsRepeatingGroup) Get(i int) NoSettlPartySubIDs {
	return NoSettlPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoSettlPartyIDsRepeatingGroup is a repeating group, Tag 781
type NoSettlPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSettlPartyIDsRepeatingGroup returns an initialized, NoSettlPartyIDsRepeatingGroup
func NewNoSettlPartyIDsRepeatingGroup() NoSettlPartyIDsRepeatingGroup {
	return NoSettlPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSettlPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SettlPartyID), quickfix.GroupElement(tag.SettlPartyIDSource), quickfix.GroupElement(tag.SettlPartyRole), NewNoSettlPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoSettlPartyIDs to this group
func (m NoSettlPartyIDsRepeatingGroup) Add() NoSettlPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoSettlPartyIDs{g}
}

//Get returns the ith NoSettlPartyIDs in the NoSettlPartyIDsRepeatinGroup
func (m NoSettlPartyIDsRepeatingGroup) Get(i int) NoSettlPartyIDs {
	return NoSettlPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoSettlDetailsRepeatingGroup is a repeating group, Tag 1158
type NoSettlDetailsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSettlDetailsRepeatingGroup returns an initialized, NoSettlDetailsRepeatingGroup
func NewNoSettlDetailsRepeatingGroup() NoSettlDetailsRepeatingGroup {
	return NoSettlDetailsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSettlDetails,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SettlObligSource), NewNoSettlPartyIDsRepeatingGroup()})}
}

//Add create and append a new NoSettlDetails to this group
func (m NoSettlDetailsRepeatingGroup) Add() NoSettlDetails {
	g := m.RepeatingGroup.Add()
	return NoSettlDetails{g}
}

//Get returns the ith NoSettlDetails in the NoSettlDetailsRepeatinGroup
func (m NoSettlDetailsRepeatingGroup) Get(i int) NoSettlDetails {
	return NoSettlDetails{m.RepeatingGroup.Get(i)}
}

//NoSidesRepeatingGroup is a repeating group, Tag 552
type NoSidesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSidesRepeatingGroup returns an initialized, NoSidesRepeatingGroup
func NewNoSidesRepeatingGroup() NoSidesRepeatingGroup {
	return NoSidesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSides,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Side), quickfix.GroupElement(tag.OrderID), quickfix.GroupElement(tag.SecondaryOrderID), quickfix.GroupElement(tag.ClOrdID), quickfix.GroupElement(tag.SecondaryClOrdID), quickfix.GroupElement(tag.ListID), NewNoPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.Account), quickfix.GroupElement(tag.AcctIDSource), quickfix.GroupElement(tag.AccountType), quickfix.GroupElement(tag.ProcessCode), quickfix.GroupElement(tag.OddLot), NewNoClearingInstructionsRepeatingGroup(), quickfix.GroupElement(tag.TradeInputSource), quickfix.GroupElement(tag.TradeInputDevice), quickfix.GroupElement(tag.OrderInputDevice), quickfix.GroupElement(tag.ComplianceID), quickfix.GroupElement(tag.SolicitedFlag), quickfix.GroupElement(tag.OrderCapacity), quickfix.GroupElement(tag.OrderRestrictions), quickfix.GroupElement(tag.CustOrderCapacity), quickfix.GroupElement(tag.OrdType), quickfix.GroupElement(tag.ExecInst), quickfix.GroupElement(tag.TransBkdTime), quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.TradingSessionSubID), quickfix.GroupElement(tag.TimeBracket), quickfix.GroupElement(tag.Commission), quickfix.GroupElement(tag.CommType), quickfix.GroupElement(tag.CommCurrency), quickfix.GroupElement(tag.FundRenewWaiv), quickfix.GroupElement(tag.NumDaysInterest), quickfix.GroupElement(tag.ExDate), quickfix.GroupElement(tag.AccruedInterestRate), quickfix.GroupElement(tag.AccruedInterestAmt), quickfix.GroupElement(tag.InterestAtMaturity), quickfix.GroupElement(tag.EndAccruedInterestAmt), quickfix.GroupElement(tag.StartCash), quickfix.GroupElement(tag.EndCash), quickfix.GroupElement(tag.Concession), quickfix.GroupElement(tag.TotalTakedown), quickfix.GroupElement(tag.NetMoney), quickfix.GroupElement(tag.SettlCurrAmt), quickfix.GroupElement(tag.SettlCurrFxRate), quickfix.GroupElement(tag.SettlCurrFxRateCalc), quickfix.GroupElement(tag.PositionEffect), quickfix.GroupElement(tag.SideMultiLegReportingType), NewNoContAmtsRepeatingGroup(), NewNoStipulationsRepeatingGroup(), NewNoMiscFeesRepeatingGroup(), quickfix.GroupElement(tag.ExchangeRule), quickfix.GroupElement(tag.TradeAllocIndicator), quickfix.GroupElement(tag.PreallocMethod), quickfix.GroupElement(tag.AllocID), NewNoAllocsRepeatingGroup(), quickfix.GroupElement(tag.LotType), quickfix.GroupElement(tag.SideGrossTradeAmt), quickfix.GroupElement(tag.AggressorIndicator), quickfix.GroupElement(tag.SideQty), quickfix.GroupElement(tag.SideTradeReportID), quickfix.GroupElement(tag.SideFillStationCd), quickfix.GroupElement(tag.SideReasonCd), quickfix.GroupElement(tag.RptSeq), quickfix.GroupElement(tag.SideTrdSubTyp), NewNoSideTrdRegTSRepeatingGroup(), quickfix.GroupElement(tag.NetGrossInd), quickfix.GroupElement(tag.SideCurrency), quickfix.GroupElement(tag.SideSettlCurrency), NewNoSettlDetailsRepeatingGroup()})}
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

//SetLegUnitOfMeasure sets LegUnitOfMeasure, Tag 999
func (m NoLegs) SetLegUnitOfMeasure(v string) {
	m.Set(field.NewLegUnitOfMeasure(v))
}

//SetLegTimeUnit sets LegTimeUnit, Tag 1001
func (m NoLegs) SetLegTimeUnit(v string) {
	m.Set(field.NewLegTimeUnit(v))
}

//SetLegOptionRatio sets LegOptionRatio, Tag 1017
func (m NoLegs) SetLegOptionRatio(v float64) {
	m.Set(field.NewLegOptionRatio(v))
}

//SetLegPrice sets LegPrice, Tag 566
func (m NoLegs) SetLegPrice(v float64) {
	m.Set(field.NewLegPrice(v))
}

//SetLegMaturityTime sets LegMaturityTime, Tag 1212
func (m NoLegs) SetLegMaturityTime(v string) {
	m.Set(field.NewLegMaturityTime(v))
}

//SetLegPutOrCall sets LegPutOrCall, Tag 1358
func (m NoLegs) SetLegPutOrCall(v int) {
	m.Set(field.NewLegPutOrCall(v))
}

//SetLegExerciseStyle sets LegExerciseStyle, Tag 1420
func (m NoLegs) SetLegExerciseStyle(v int) {
	m.Set(field.NewLegExerciseStyle(v))
}

//SetLegUnitOfMeasureQty sets LegUnitOfMeasureQty, Tag 1224
func (m NoLegs) SetLegUnitOfMeasureQty(v float64) {
	m.Set(field.NewLegUnitOfMeasureQty(v))
}

//SetLegPriceUnitOfMeasure sets LegPriceUnitOfMeasure, Tag 1421
func (m NoLegs) SetLegPriceUnitOfMeasure(v string) {
	m.Set(field.NewLegPriceUnitOfMeasure(v))
}

//SetLegPriceUnitOfMeasureQty sets LegPriceUnitOfMeasureQty, Tag 1422
func (m NoLegs) SetLegPriceUnitOfMeasureQty(v float64) {
	m.Set(field.NewLegPriceUnitOfMeasureQty(v))
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

//SetLegReportID sets LegReportID, Tag 990
func (m NoLegs) SetLegReportID(v string) {
	m.Set(field.NewLegReportID(v))
}

//SetLegSettlCurrency sets LegSettlCurrency, Tag 675
func (m NoLegs) SetLegSettlCurrency(v string) {
	m.Set(field.NewLegSettlCurrency(v))
}

//SetLegLastForwardPoints sets LegLastForwardPoints, Tag 1073
func (m NoLegs) SetLegLastForwardPoints(v float64) {
	m.Set(field.NewLegLastForwardPoints(v))
}

//SetLegCalculatedCcyLastQty sets LegCalculatedCcyLastQty, Tag 1074
func (m NoLegs) SetLegCalculatedCcyLastQty(v float64) {
	m.Set(field.NewLegCalculatedCcyLastQty(v))
}

//SetLegGrossTradeAmt sets LegGrossTradeAmt, Tag 1075
func (m NoLegs) SetLegGrossTradeAmt(v float64) {
	m.Set(field.NewLegGrossTradeAmt(v))
}

//SetLegNumber sets LegNumber, Tag 1152
func (m NoLegs) SetLegNumber(v int) {
	m.Set(field.NewLegNumber(v))
}

//SetNoOfLegUnderlyings sets NoOfLegUnderlyings, Tag 1342
func (m NoLegs) SetNoOfLegUnderlyings(f NoOfLegUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

//SetLegVolatility sets LegVolatility, Tag 1379
func (m NoLegs) SetLegVolatility(v float64) {
	m.Set(field.NewLegVolatility(v))
}

//SetLegDividendYield sets LegDividendYield, Tag 1381
func (m NoLegs) SetLegDividendYield(v float64) {
	m.Set(field.NewLegDividendYield(v))
}

//SetLegCurrencyRatio sets LegCurrencyRatio, Tag 1383
func (m NoLegs) SetLegCurrencyRatio(v float64) {
	m.Set(field.NewLegCurrencyRatio(v))
}

//SetLegExecInst sets LegExecInst, Tag 1384
func (m NoLegs) SetLegExecInst(v string) {
	m.Set(field.NewLegExecInst(v))
}

//SetLegLastQty sets LegLastQty, Tag 1418
func (m NoLegs) SetLegLastQty(v float64) {
	m.Set(field.NewLegLastQty(v))
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

//GetLegOptionRatio gets LegOptionRatio, Tag 1017
func (m NoLegs) GetLegOptionRatio() (f field.LegOptionRatioField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPrice gets LegPrice, Tag 566
func (m NoLegs) GetLegPrice() (f field.LegPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegMaturityTime gets LegMaturityTime, Tag 1212
func (m NoLegs) GetLegMaturityTime() (f field.LegMaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPutOrCall gets LegPutOrCall, Tag 1358
func (m NoLegs) GetLegPutOrCall() (f field.LegPutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegExerciseStyle gets LegExerciseStyle, Tag 1420
func (m NoLegs) GetLegExerciseStyle() (f field.LegExerciseStyleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegUnitOfMeasureQty gets LegUnitOfMeasureQty, Tag 1224
func (m NoLegs) GetLegUnitOfMeasureQty() (f field.LegUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPriceUnitOfMeasure gets LegPriceUnitOfMeasure, Tag 1421
func (m NoLegs) GetLegPriceUnitOfMeasure() (f field.LegPriceUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPriceUnitOfMeasureQty gets LegPriceUnitOfMeasureQty, Tag 1422
func (m NoLegs) GetLegPriceUnitOfMeasureQty() (f field.LegPriceUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
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

//GetLegReportID gets LegReportID, Tag 990
func (m NoLegs) GetLegReportID() (f field.LegReportIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSettlCurrency gets LegSettlCurrency, Tag 675
func (m NoLegs) GetLegSettlCurrency() (f field.LegSettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegLastForwardPoints gets LegLastForwardPoints, Tag 1073
func (m NoLegs) GetLegLastForwardPoints() (f field.LegLastForwardPointsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCalculatedCcyLastQty gets LegCalculatedCcyLastQty, Tag 1074
func (m NoLegs) GetLegCalculatedCcyLastQty() (f field.LegCalculatedCcyLastQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegGrossTradeAmt gets LegGrossTradeAmt, Tag 1075
func (m NoLegs) GetLegGrossTradeAmt() (f field.LegGrossTradeAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegNumber gets LegNumber, Tag 1152
func (m NoLegs) GetLegNumber() (f field.LegNumberField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoOfLegUnderlyings gets NoOfLegUnderlyings, Tag 1342
func (m NoLegs) GetNoOfLegUnderlyings() (NoOfLegUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoOfLegUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLegVolatility gets LegVolatility, Tag 1379
func (m NoLegs) GetLegVolatility() (f field.LegVolatilityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegDividendYield gets LegDividendYield, Tag 1381
func (m NoLegs) GetLegDividendYield() (f field.LegDividendYieldField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCurrencyRatio gets LegCurrencyRatio, Tag 1383
func (m NoLegs) GetLegCurrencyRatio() (f field.LegCurrencyRatioField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegExecInst gets LegExecInst, Tag 1384
func (m NoLegs) GetLegExecInst() (f field.LegExecInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegLastQty gets LegLastQty, Tag 1418
func (m NoLegs) GetLegLastQty() (f field.LegLastQtyField, err quickfix.MessageRejectError) {
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

//HasLegOptionRatio returns true if LegOptionRatio is present, Tag 1017
func (m NoLegs) HasLegOptionRatio() bool {
	return m.Has(tag.LegOptionRatio)
}

//HasLegPrice returns true if LegPrice is present, Tag 566
func (m NoLegs) HasLegPrice() bool {
	return m.Has(tag.LegPrice)
}

//HasLegMaturityTime returns true if LegMaturityTime is present, Tag 1212
func (m NoLegs) HasLegMaturityTime() bool {
	return m.Has(tag.LegMaturityTime)
}

//HasLegPutOrCall returns true if LegPutOrCall is present, Tag 1358
func (m NoLegs) HasLegPutOrCall() bool {
	return m.Has(tag.LegPutOrCall)
}

//HasLegExerciseStyle returns true if LegExerciseStyle is present, Tag 1420
func (m NoLegs) HasLegExerciseStyle() bool {
	return m.Has(tag.LegExerciseStyle)
}

//HasLegUnitOfMeasureQty returns true if LegUnitOfMeasureQty is present, Tag 1224
func (m NoLegs) HasLegUnitOfMeasureQty() bool {
	return m.Has(tag.LegUnitOfMeasureQty)
}

//HasLegPriceUnitOfMeasure returns true if LegPriceUnitOfMeasure is present, Tag 1421
func (m NoLegs) HasLegPriceUnitOfMeasure() bool {
	return m.Has(tag.LegPriceUnitOfMeasure)
}

//HasLegPriceUnitOfMeasureQty returns true if LegPriceUnitOfMeasureQty is present, Tag 1422
func (m NoLegs) HasLegPriceUnitOfMeasureQty() bool {
	return m.Has(tag.LegPriceUnitOfMeasureQty)
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

//HasLegReportID returns true if LegReportID is present, Tag 990
func (m NoLegs) HasLegReportID() bool {
	return m.Has(tag.LegReportID)
}

//HasLegSettlCurrency returns true if LegSettlCurrency is present, Tag 675
func (m NoLegs) HasLegSettlCurrency() bool {
	return m.Has(tag.LegSettlCurrency)
}

//HasLegLastForwardPoints returns true if LegLastForwardPoints is present, Tag 1073
func (m NoLegs) HasLegLastForwardPoints() bool {
	return m.Has(tag.LegLastForwardPoints)
}

//HasLegCalculatedCcyLastQty returns true if LegCalculatedCcyLastQty is present, Tag 1074
func (m NoLegs) HasLegCalculatedCcyLastQty() bool {
	return m.Has(tag.LegCalculatedCcyLastQty)
}

//HasLegGrossTradeAmt returns true if LegGrossTradeAmt is present, Tag 1075
func (m NoLegs) HasLegGrossTradeAmt() bool {
	return m.Has(tag.LegGrossTradeAmt)
}

//HasLegNumber returns true if LegNumber is present, Tag 1152
func (m NoLegs) HasLegNumber() bool {
	return m.Has(tag.LegNumber)
}

//HasNoOfLegUnderlyings returns true if NoOfLegUnderlyings is present, Tag 1342
func (m NoLegs) HasNoOfLegUnderlyings() bool {
	return m.Has(tag.NoOfLegUnderlyings)
}

//HasLegVolatility returns true if LegVolatility is present, Tag 1379
func (m NoLegs) HasLegVolatility() bool {
	return m.Has(tag.LegVolatility)
}

//HasLegDividendYield returns true if LegDividendYield is present, Tag 1381
func (m NoLegs) HasLegDividendYield() bool {
	return m.Has(tag.LegDividendYield)
}

//HasLegCurrencyRatio returns true if LegCurrencyRatio is present, Tag 1383
func (m NoLegs) HasLegCurrencyRatio() bool {
	return m.Has(tag.LegCurrencyRatio)
}

//HasLegExecInst returns true if LegExecInst is present, Tag 1384
func (m NoLegs) HasLegExecInst() bool {
	return m.Has(tag.LegExecInst)
}

//HasLegLastQty returns true if LegLastQty is present, Tag 1418
func (m NoLegs) HasLegLastQty() bool {
	return m.Has(tag.LegLastQty)
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

//NoOfLegUnderlyings is a repeating group element, Tag 1342
type NoOfLegUnderlyings struct {
	quickfix.Group
}

//SetUnderlyingLegSymbol sets UnderlyingLegSymbol, Tag 1330
func (m NoOfLegUnderlyings) SetUnderlyingLegSymbol(v string) {
	m.Set(field.NewUnderlyingLegSymbol(v))
}

//SetUnderlyingLegSymbolSfx sets UnderlyingLegSymbolSfx, Tag 1331
func (m NoOfLegUnderlyings) SetUnderlyingLegSymbolSfx(v string) {
	m.Set(field.NewUnderlyingLegSymbolSfx(v))
}

//SetUnderlyingLegSecurityID sets UnderlyingLegSecurityID, Tag 1332
func (m NoOfLegUnderlyings) SetUnderlyingLegSecurityID(v string) {
	m.Set(field.NewUnderlyingLegSecurityID(v))
}

//SetUnderlyingLegSecurityIDSource sets UnderlyingLegSecurityIDSource, Tag 1333
func (m NoOfLegUnderlyings) SetUnderlyingLegSecurityIDSource(v string) {
	m.Set(field.NewUnderlyingLegSecurityIDSource(v))
}

//SetNoUnderlyingLegSecurityAltID sets NoUnderlyingLegSecurityAltID, Tag 1334
func (m NoOfLegUnderlyings) SetNoUnderlyingLegSecurityAltID(f NoUnderlyingLegSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingLegCFICode sets UnderlyingLegCFICode, Tag 1344
func (m NoOfLegUnderlyings) SetUnderlyingLegCFICode(v string) {
	m.Set(field.NewUnderlyingLegCFICode(v))
}

//SetUnderlyingLegSecurityType sets UnderlyingLegSecurityType, Tag 1337
func (m NoOfLegUnderlyings) SetUnderlyingLegSecurityType(v string) {
	m.Set(field.NewUnderlyingLegSecurityType(v))
}

//SetUnderlyingLegSecuritySubType sets UnderlyingLegSecuritySubType, Tag 1338
func (m NoOfLegUnderlyings) SetUnderlyingLegSecuritySubType(v string) {
	m.Set(field.NewUnderlyingLegSecuritySubType(v))
}

//SetUnderlyingLegMaturityMonthYear sets UnderlyingLegMaturityMonthYear, Tag 1339
func (m NoOfLegUnderlyings) SetUnderlyingLegMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingLegMaturityMonthYear(v))
}

//SetUnderlyingLegMaturityDate sets UnderlyingLegMaturityDate, Tag 1345
func (m NoOfLegUnderlyings) SetUnderlyingLegMaturityDate(v string) {
	m.Set(field.NewUnderlyingLegMaturityDate(v))
}

//SetUnderlyingLegMaturityTime sets UnderlyingLegMaturityTime, Tag 1405
func (m NoOfLegUnderlyings) SetUnderlyingLegMaturityTime(v string) {
	m.Set(field.NewUnderlyingLegMaturityTime(v))
}

//SetUnderlyingLegStrikePrice sets UnderlyingLegStrikePrice, Tag 1340
func (m NoOfLegUnderlyings) SetUnderlyingLegStrikePrice(v float64) {
	m.Set(field.NewUnderlyingLegStrikePrice(v))
}

//SetUnderlyingLegOptAttribute sets UnderlyingLegOptAttribute, Tag 1391
func (m NoOfLegUnderlyings) SetUnderlyingLegOptAttribute(v string) {
	m.Set(field.NewUnderlyingLegOptAttribute(v))
}

//SetUnderlyingLegPutOrCall sets UnderlyingLegPutOrCall, Tag 1343
func (m NoOfLegUnderlyings) SetUnderlyingLegPutOrCall(v int) {
	m.Set(field.NewUnderlyingLegPutOrCall(v))
}

//SetUnderlyingLegSecurityExchange sets UnderlyingLegSecurityExchange, Tag 1341
func (m NoOfLegUnderlyings) SetUnderlyingLegSecurityExchange(v string) {
	m.Set(field.NewUnderlyingLegSecurityExchange(v))
}

//SetUnderlyingLegSecurityDesc sets UnderlyingLegSecurityDesc, Tag 1392
func (m NoOfLegUnderlyings) SetUnderlyingLegSecurityDesc(v string) {
	m.Set(field.NewUnderlyingLegSecurityDesc(v))
}

//GetUnderlyingLegSymbol gets UnderlyingLegSymbol, Tag 1330
func (m NoOfLegUnderlyings) GetUnderlyingLegSymbol() (f field.UnderlyingLegSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLegSymbolSfx gets UnderlyingLegSymbolSfx, Tag 1331
func (m NoOfLegUnderlyings) GetUnderlyingLegSymbolSfx() (f field.UnderlyingLegSymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLegSecurityID gets UnderlyingLegSecurityID, Tag 1332
func (m NoOfLegUnderlyings) GetUnderlyingLegSecurityID() (f field.UnderlyingLegSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLegSecurityIDSource gets UnderlyingLegSecurityIDSource, Tag 1333
func (m NoOfLegUnderlyings) GetUnderlyingLegSecurityIDSource() (f field.UnderlyingLegSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyingLegSecurityAltID gets NoUnderlyingLegSecurityAltID, Tag 1334
func (m NoOfLegUnderlyings) GetNoUnderlyingLegSecurityAltID() (NoUnderlyingLegSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingLegSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingLegCFICode gets UnderlyingLegCFICode, Tag 1344
func (m NoOfLegUnderlyings) GetUnderlyingLegCFICode() (f field.UnderlyingLegCFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLegSecurityType gets UnderlyingLegSecurityType, Tag 1337
func (m NoOfLegUnderlyings) GetUnderlyingLegSecurityType() (f field.UnderlyingLegSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLegSecuritySubType gets UnderlyingLegSecuritySubType, Tag 1338
func (m NoOfLegUnderlyings) GetUnderlyingLegSecuritySubType() (f field.UnderlyingLegSecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLegMaturityMonthYear gets UnderlyingLegMaturityMonthYear, Tag 1339
func (m NoOfLegUnderlyings) GetUnderlyingLegMaturityMonthYear() (f field.UnderlyingLegMaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLegMaturityDate gets UnderlyingLegMaturityDate, Tag 1345
func (m NoOfLegUnderlyings) GetUnderlyingLegMaturityDate() (f field.UnderlyingLegMaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLegMaturityTime gets UnderlyingLegMaturityTime, Tag 1405
func (m NoOfLegUnderlyings) GetUnderlyingLegMaturityTime() (f field.UnderlyingLegMaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLegStrikePrice gets UnderlyingLegStrikePrice, Tag 1340
func (m NoOfLegUnderlyings) GetUnderlyingLegStrikePrice() (f field.UnderlyingLegStrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLegOptAttribute gets UnderlyingLegOptAttribute, Tag 1391
func (m NoOfLegUnderlyings) GetUnderlyingLegOptAttribute() (f field.UnderlyingLegOptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLegPutOrCall gets UnderlyingLegPutOrCall, Tag 1343
func (m NoOfLegUnderlyings) GetUnderlyingLegPutOrCall() (f field.UnderlyingLegPutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLegSecurityExchange gets UnderlyingLegSecurityExchange, Tag 1341
func (m NoOfLegUnderlyings) GetUnderlyingLegSecurityExchange() (f field.UnderlyingLegSecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLegSecurityDesc gets UnderlyingLegSecurityDesc, Tag 1392
func (m NoOfLegUnderlyings) GetUnderlyingLegSecurityDesc() (f field.UnderlyingLegSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingLegSymbol returns true if UnderlyingLegSymbol is present, Tag 1330
func (m NoOfLegUnderlyings) HasUnderlyingLegSymbol() bool {
	return m.Has(tag.UnderlyingLegSymbol)
}

//HasUnderlyingLegSymbolSfx returns true if UnderlyingLegSymbolSfx is present, Tag 1331
func (m NoOfLegUnderlyings) HasUnderlyingLegSymbolSfx() bool {
	return m.Has(tag.UnderlyingLegSymbolSfx)
}

//HasUnderlyingLegSecurityID returns true if UnderlyingLegSecurityID is present, Tag 1332
func (m NoOfLegUnderlyings) HasUnderlyingLegSecurityID() bool {
	return m.Has(tag.UnderlyingLegSecurityID)
}

//HasUnderlyingLegSecurityIDSource returns true if UnderlyingLegSecurityIDSource is present, Tag 1333
func (m NoOfLegUnderlyings) HasUnderlyingLegSecurityIDSource() bool {
	return m.Has(tag.UnderlyingLegSecurityIDSource)
}

//HasNoUnderlyingLegSecurityAltID returns true if NoUnderlyingLegSecurityAltID is present, Tag 1334
func (m NoOfLegUnderlyings) HasNoUnderlyingLegSecurityAltID() bool {
	return m.Has(tag.NoUnderlyingLegSecurityAltID)
}

//HasUnderlyingLegCFICode returns true if UnderlyingLegCFICode is present, Tag 1344
func (m NoOfLegUnderlyings) HasUnderlyingLegCFICode() bool {
	return m.Has(tag.UnderlyingLegCFICode)
}

//HasUnderlyingLegSecurityType returns true if UnderlyingLegSecurityType is present, Tag 1337
func (m NoOfLegUnderlyings) HasUnderlyingLegSecurityType() bool {
	return m.Has(tag.UnderlyingLegSecurityType)
}

//HasUnderlyingLegSecuritySubType returns true if UnderlyingLegSecuritySubType is present, Tag 1338
func (m NoOfLegUnderlyings) HasUnderlyingLegSecuritySubType() bool {
	return m.Has(tag.UnderlyingLegSecuritySubType)
}

//HasUnderlyingLegMaturityMonthYear returns true if UnderlyingLegMaturityMonthYear is present, Tag 1339
func (m NoOfLegUnderlyings) HasUnderlyingLegMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingLegMaturityMonthYear)
}

//HasUnderlyingLegMaturityDate returns true if UnderlyingLegMaturityDate is present, Tag 1345
func (m NoOfLegUnderlyings) HasUnderlyingLegMaturityDate() bool {
	return m.Has(tag.UnderlyingLegMaturityDate)
}

//HasUnderlyingLegMaturityTime returns true if UnderlyingLegMaturityTime is present, Tag 1405
func (m NoOfLegUnderlyings) HasUnderlyingLegMaturityTime() bool {
	return m.Has(tag.UnderlyingLegMaturityTime)
}

//HasUnderlyingLegStrikePrice returns true if UnderlyingLegStrikePrice is present, Tag 1340
func (m NoOfLegUnderlyings) HasUnderlyingLegStrikePrice() bool {
	return m.Has(tag.UnderlyingLegStrikePrice)
}

//HasUnderlyingLegOptAttribute returns true if UnderlyingLegOptAttribute is present, Tag 1391
func (m NoOfLegUnderlyings) HasUnderlyingLegOptAttribute() bool {
	return m.Has(tag.UnderlyingLegOptAttribute)
}

//HasUnderlyingLegPutOrCall returns true if UnderlyingLegPutOrCall is present, Tag 1343
func (m NoOfLegUnderlyings) HasUnderlyingLegPutOrCall() bool {
	return m.Has(tag.UnderlyingLegPutOrCall)
}

//HasUnderlyingLegSecurityExchange returns true if UnderlyingLegSecurityExchange is present, Tag 1341
func (m NoOfLegUnderlyings) HasUnderlyingLegSecurityExchange() bool {
	return m.Has(tag.UnderlyingLegSecurityExchange)
}

//HasUnderlyingLegSecurityDesc returns true if UnderlyingLegSecurityDesc is present, Tag 1392
func (m NoOfLegUnderlyings) HasUnderlyingLegSecurityDesc() bool {
	return m.Has(tag.UnderlyingLegSecurityDesc)
}

//NoUnderlyingLegSecurityAltID is a repeating group element, Tag 1334
type NoUnderlyingLegSecurityAltID struct {
	quickfix.Group
}

//SetUnderlyingLegSecurityAltID sets UnderlyingLegSecurityAltID, Tag 1335
func (m NoUnderlyingLegSecurityAltID) SetUnderlyingLegSecurityAltID(v string) {
	m.Set(field.NewUnderlyingLegSecurityAltID(v))
}

//SetUnderlyingLegSecurityAltIDSource sets UnderlyingLegSecurityAltIDSource, Tag 1336
func (m NoUnderlyingLegSecurityAltID) SetUnderlyingLegSecurityAltIDSource(v string) {
	m.Set(field.NewUnderlyingLegSecurityAltIDSource(v))
}

//GetUnderlyingLegSecurityAltID gets UnderlyingLegSecurityAltID, Tag 1335
func (m NoUnderlyingLegSecurityAltID) GetUnderlyingLegSecurityAltID() (f field.UnderlyingLegSecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLegSecurityAltIDSource gets UnderlyingLegSecurityAltIDSource, Tag 1336
func (m NoUnderlyingLegSecurityAltID) GetUnderlyingLegSecurityAltIDSource() (f field.UnderlyingLegSecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingLegSecurityAltID returns true if UnderlyingLegSecurityAltID is present, Tag 1335
func (m NoUnderlyingLegSecurityAltID) HasUnderlyingLegSecurityAltID() bool {
	return m.Has(tag.UnderlyingLegSecurityAltID)
}

//HasUnderlyingLegSecurityAltIDSource returns true if UnderlyingLegSecurityAltIDSource is present, Tag 1336
func (m NoUnderlyingLegSecurityAltID) HasUnderlyingLegSecurityAltIDSource() bool {
	return m.Has(tag.UnderlyingLegSecurityAltIDSource)
}

//NoUnderlyingLegSecurityAltIDRepeatingGroup is a repeating group, Tag 1334
type NoUnderlyingLegSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingLegSecurityAltIDRepeatingGroup returns an initialized, NoUnderlyingLegSecurityAltIDRepeatingGroup
func NewNoUnderlyingLegSecurityAltIDRepeatingGroup() NoUnderlyingLegSecurityAltIDRepeatingGroup {
	return NoUnderlyingLegSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingLegSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingLegSecurityAltID), quickfix.GroupElement(tag.UnderlyingLegSecurityAltIDSource)})}
}

//Add create and append a new NoUnderlyingLegSecurityAltID to this group
func (m NoUnderlyingLegSecurityAltIDRepeatingGroup) Add() NoUnderlyingLegSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingLegSecurityAltID{g}
}

//Get returns the ith NoUnderlyingLegSecurityAltID in the NoUnderlyingLegSecurityAltIDRepeatinGroup
func (m NoUnderlyingLegSecurityAltIDRepeatingGroup) Get(i int) NoUnderlyingLegSecurityAltID {
	return NoUnderlyingLegSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoOfLegUnderlyingsRepeatingGroup is a repeating group, Tag 1342
type NoOfLegUnderlyingsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoOfLegUnderlyingsRepeatingGroup returns an initialized, NoOfLegUnderlyingsRepeatingGroup
func NewNoOfLegUnderlyingsRepeatingGroup() NoOfLegUnderlyingsRepeatingGroup {
	return NoOfLegUnderlyingsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoOfLegUnderlyings,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingLegSymbol), quickfix.GroupElement(tag.UnderlyingLegSymbolSfx), quickfix.GroupElement(tag.UnderlyingLegSecurityID), quickfix.GroupElement(tag.UnderlyingLegSecurityIDSource), NewNoUnderlyingLegSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingLegCFICode), quickfix.GroupElement(tag.UnderlyingLegSecurityType), quickfix.GroupElement(tag.UnderlyingLegSecuritySubType), quickfix.GroupElement(tag.UnderlyingLegMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingLegMaturityDate), quickfix.GroupElement(tag.UnderlyingLegMaturityTime), quickfix.GroupElement(tag.UnderlyingLegStrikePrice), quickfix.GroupElement(tag.UnderlyingLegOptAttribute), quickfix.GroupElement(tag.UnderlyingLegPutOrCall), quickfix.GroupElement(tag.UnderlyingLegSecurityExchange), quickfix.GroupElement(tag.UnderlyingLegSecurityDesc)})}
}

//Add create and append a new NoOfLegUnderlyings to this group
func (m NoOfLegUnderlyingsRepeatingGroup) Add() NoOfLegUnderlyings {
	g := m.RepeatingGroup.Add()
	return NoOfLegUnderlyings{g}
}

//Get returns the ith NoOfLegUnderlyings in the NoOfLegUnderlyingsRepeatinGroup
func (m NoOfLegUnderlyingsRepeatingGroup) Get(i int) NoOfLegUnderlyings {
	return NoOfLegUnderlyings{m.RepeatingGroup.Get(i)}
}

//NoLegsRepeatingGroup is a repeating group, Tag 555
type NoLegsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegsRepeatingGroup returns an initialized, NoLegsRepeatingGroup
func NewNoLegsRepeatingGroup() NoLegsRepeatingGroup {
	return NoLegsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSymbol), quickfix.GroupElement(tag.LegSymbolSfx), quickfix.GroupElement(tag.LegSecurityID), quickfix.GroupElement(tag.LegSecurityIDSource), NewNoLegSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.LegProduct), quickfix.GroupElement(tag.LegCFICode), quickfix.GroupElement(tag.LegSecurityType), quickfix.GroupElement(tag.LegSecuritySubType), quickfix.GroupElement(tag.LegMaturityMonthYear), quickfix.GroupElement(tag.LegMaturityDate), quickfix.GroupElement(tag.LegCouponPaymentDate), quickfix.GroupElement(tag.LegIssueDate), quickfix.GroupElement(tag.LegRepoCollateralSecurityType), quickfix.GroupElement(tag.LegRepurchaseTerm), quickfix.GroupElement(tag.LegRepurchaseRate), quickfix.GroupElement(tag.LegFactor), quickfix.GroupElement(tag.LegCreditRating), quickfix.GroupElement(tag.LegInstrRegistry), quickfix.GroupElement(tag.LegCountryOfIssue), quickfix.GroupElement(tag.LegStateOrProvinceOfIssue), quickfix.GroupElement(tag.LegLocaleOfIssue), quickfix.GroupElement(tag.LegRedemptionDate), quickfix.GroupElement(tag.LegStrikePrice), quickfix.GroupElement(tag.LegStrikeCurrency), quickfix.GroupElement(tag.LegOptAttribute), quickfix.GroupElement(tag.LegContractMultiplier), quickfix.GroupElement(tag.LegCouponRate), quickfix.GroupElement(tag.LegSecurityExchange), quickfix.GroupElement(tag.LegIssuer), quickfix.GroupElement(tag.EncodedLegIssuerLen), quickfix.GroupElement(tag.EncodedLegIssuer), quickfix.GroupElement(tag.LegSecurityDesc), quickfix.GroupElement(tag.EncodedLegSecurityDescLen), quickfix.GroupElement(tag.EncodedLegSecurityDesc), quickfix.GroupElement(tag.LegRatioQty), quickfix.GroupElement(tag.LegSide), quickfix.GroupElement(tag.LegCurrency), quickfix.GroupElement(tag.LegPool), quickfix.GroupElement(tag.LegDatedDate), quickfix.GroupElement(tag.LegContractSettlMonth), quickfix.GroupElement(tag.LegInterestAccrualDate), quickfix.GroupElement(tag.LegUnitOfMeasure), quickfix.GroupElement(tag.LegTimeUnit), quickfix.GroupElement(tag.LegOptionRatio), quickfix.GroupElement(tag.LegPrice), quickfix.GroupElement(tag.LegMaturityTime), quickfix.GroupElement(tag.LegPutOrCall), quickfix.GroupElement(tag.LegExerciseStyle), quickfix.GroupElement(tag.LegUnitOfMeasureQty), quickfix.GroupElement(tag.LegPriceUnitOfMeasure), quickfix.GroupElement(tag.LegPriceUnitOfMeasureQty), quickfix.GroupElement(tag.LegQty), quickfix.GroupElement(tag.LegSwapType), NewNoLegStipulationsRepeatingGroup(), quickfix.GroupElement(tag.LegPositionEffect), quickfix.GroupElement(tag.LegCoveredOrUncovered), NewNoNestedPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.LegRefID), quickfix.GroupElement(tag.LegSettlType), quickfix.GroupElement(tag.LegSettlDate), quickfix.GroupElement(tag.LegLastPx), quickfix.GroupElement(tag.LegReportID), quickfix.GroupElement(tag.LegSettlCurrency), quickfix.GroupElement(tag.LegLastForwardPoints), quickfix.GroupElement(tag.LegCalculatedCcyLastQty), quickfix.GroupElement(tag.LegGrossTradeAmt), quickfix.GroupElement(tag.LegNumber), NewNoOfLegUnderlyingsRepeatingGroup(), quickfix.GroupElement(tag.LegVolatility), quickfix.GroupElement(tag.LegDividendYield), quickfix.GroupElement(tag.LegCurrencyRatio), quickfix.GroupElement(tag.LegExecInst), quickfix.GroupElement(tag.LegLastQty)})}
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingSecurityIDSource), NewNoUnderlyingSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingProduct), quickfix.GroupElement(tag.UnderlyingCFICode), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingSecuritySubType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDate), quickfix.GroupElement(tag.UnderlyingCouponPaymentDate), quickfix.GroupElement(tag.UnderlyingIssueDate), quickfix.GroupElement(tag.UnderlyingRepoCollateralSecurityType), quickfix.GroupElement(tag.UnderlyingRepurchaseTerm), quickfix.GroupElement(tag.UnderlyingRepurchaseRate), quickfix.GroupElement(tag.UnderlyingFactor), quickfix.GroupElement(tag.UnderlyingCreditRating), quickfix.GroupElement(tag.UnderlyingInstrRegistry), quickfix.GroupElement(tag.UnderlyingCountryOfIssue), quickfix.GroupElement(tag.UnderlyingStateOrProvinceOfIssue), quickfix.GroupElement(tag.UnderlyingLocaleOfIssue), quickfix.GroupElement(tag.UnderlyingRedemptionDate), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingStrikeCurrency), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.UnderlyingCPProgram), quickfix.GroupElement(tag.UnderlyingCPRegType), quickfix.GroupElement(tag.UnderlyingCurrency), quickfix.GroupElement(tag.UnderlyingQty), quickfix.GroupElement(tag.UnderlyingPx), quickfix.GroupElement(tag.UnderlyingDirtyPrice), quickfix.GroupElement(tag.UnderlyingEndPrice), quickfix.GroupElement(tag.UnderlyingStartValue), quickfix.GroupElement(tag.UnderlyingCurrentValue), quickfix.GroupElement(tag.UnderlyingEndValue), NewNoUnderlyingStipsRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingAllocationPercent), quickfix.GroupElement(tag.UnderlyingSettlementType), quickfix.GroupElement(tag.UnderlyingCashAmount), quickfix.GroupElement(tag.UnderlyingCashType), quickfix.GroupElement(tag.UnderlyingUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingTimeUnit), quickfix.GroupElement(tag.UnderlyingCapValue), NewNoUndlyInstrumentPartiesRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingSettlMethod), quickfix.GroupElement(tag.UnderlyingAdjustedQuantity), quickfix.GroupElement(tag.UnderlyingFXRate), quickfix.GroupElement(tag.UnderlyingFXRateCalc), quickfix.GroupElement(tag.UnderlyingMaturityTime), quickfix.GroupElement(tag.UnderlyingPutOrCall), quickfix.GroupElement(tag.UnderlyingExerciseStyle), quickfix.GroupElement(tag.UnderlyingUnitOfMeasureQty), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasureQty)})}
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

//SetPositionCurrency sets PositionCurrency, Tag 1055
func (m NoPosAmt) SetPositionCurrency(v string) {
	m.Set(field.NewPositionCurrency(v))
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

//GetPositionCurrency gets PositionCurrency, Tag 1055
func (m NoPosAmt) GetPositionCurrency() (f field.PositionCurrencyField, err quickfix.MessageRejectError) {
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

//HasPositionCurrency returns true if PositionCurrency is present, Tag 1055
func (m NoPosAmt) HasPositionCurrency() bool {
	return m.Has(tag.PositionCurrency)
}

//NoPosAmtRepeatingGroup is a repeating group, Tag 753
type NoPosAmtRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPosAmtRepeatingGroup returns an initialized, NoPosAmtRepeatingGroup
func NewNoPosAmtRepeatingGroup() NoPosAmtRepeatingGroup {
	return NoPosAmtRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPosAmt,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PosAmtType), quickfix.GroupElement(tag.PosAmt), quickfix.GroupElement(tag.PositionCurrency)})}
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

//NoRootPartyIDs is a repeating group element, Tag 1116
type NoRootPartyIDs struct {
	quickfix.Group
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
func (m NoRootPartyIDs) GetRootPartyID() (f field.RootPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRootPartyIDSource gets RootPartyIDSource, Tag 1118
func (m NoRootPartyIDs) GetRootPartyIDSource() (f field.RootPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRootPartyRole gets RootPartyRole, Tag 1119
func (m NoRootPartyIDs) GetRootPartyRole() (f field.RootPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
	quickfix.Group
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
func (m NoRootPartySubIDs) GetRootPartySubID() (f field.RootPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRootPartySubIDType gets RootPartySubIDType, Tag 1122
func (m NoRootPartySubIDs) GetRootPartySubIDType() (f field.RootPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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

//NoTrdRepIndicators is a repeating group element, Tag 1387
type NoTrdRepIndicators struct {
	quickfix.Group
}

//SetTrdRepPartyRole sets TrdRepPartyRole, Tag 1388
func (m NoTrdRepIndicators) SetTrdRepPartyRole(v int) {
	m.Set(field.NewTrdRepPartyRole(v))
}

//SetTrdRepIndicator sets TrdRepIndicator, Tag 1389
func (m NoTrdRepIndicators) SetTrdRepIndicator(v bool) {
	m.Set(field.NewTrdRepIndicator(v))
}

//GetTrdRepPartyRole gets TrdRepPartyRole, Tag 1388
func (m NoTrdRepIndicators) GetTrdRepPartyRole() (f field.TrdRepPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTrdRepIndicator gets TrdRepIndicator, Tag 1389
func (m NoTrdRepIndicators) GetTrdRepIndicator() (f field.TrdRepIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasTrdRepPartyRole returns true if TrdRepPartyRole is present, Tag 1388
func (m NoTrdRepIndicators) HasTrdRepPartyRole() bool {
	return m.Has(tag.TrdRepPartyRole)
}

//HasTrdRepIndicator returns true if TrdRepIndicator is present, Tag 1389
func (m NoTrdRepIndicators) HasTrdRepIndicator() bool {
	return m.Has(tag.TrdRepIndicator)
}

//NoTrdRepIndicatorsRepeatingGroup is a repeating group, Tag 1387
type NoTrdRepIndicatorsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTrdRepIndicatorsRepeatingGroup returns an initialized, NoTrdRepIndicatorsRepeatingGroup
func NewNoTrdRepIndicatorsRepeatingGroup() NoTrdRepIndicatorsRepeatingGroup {
	return NoTrdRepIndicatorsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTrdRepIndicators,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TrdRepPartyRole), quickfix.GroupElement(tag.TrdRepIndicator)})}
}

//Add create and append a new NoTrdRepIndicators to this group
func (m NoTrdRepIndicatorsRepeatingGroup) Add() NoTrdRepIndicators {
	g := m.RepeatingGroup.Add()
	return NoTrdRepIndicators{g}
}

//Get returns the ith NoTrdRepIndicators in the NoTrdRepIndicatorsRepeatinGroup
func (m NoTrdRepIndicatorsRepeatingGroup) Get(i int) NoTrdRepIndicators {
	return NoTrdRepIndicators{m.RepeatingGroup.Get(i)}
}
