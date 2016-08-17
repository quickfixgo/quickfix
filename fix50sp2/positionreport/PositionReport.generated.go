package positionreport

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//PositionReport is the fix50sp2 PositionReport type, MsgType = AP
type PositionReport struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a PositionReport from a quickfix.Message instance
func FromMessage(m quickfix.Message) PositionReport {
	return PositionReport{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m PositionReport) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a PositionReport initialized with the required fields for PositionReport
func New(posmaintrptid field.PosMaintRptIDField, clearingbusinessdate field.ClearingBusinessDateField) (m PositionReport) {
	m.Header = fixt11.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("AP"))
	m.Set(posmaintrptid)
	m.Set(clearingbusinessdate)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg PositionReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "AP", r
}

//SetAccount sets Account, Tag 1
func (m PositionReport) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetCurrency sets Currency, Tag 15
func (m PositionReport) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m PositionReport) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m PositionReport) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m PositionReport) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m PositionReport) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m PositionReport) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIssuer sets Issuer, Tag 106
func (m PositionReport) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m PositionReport) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m PositionReport) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m PositionReport) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m PositionReport) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m PositionReport) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m PositionReport) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m PositionReport) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m PositionReport) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m PositionReport) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m PositionReport) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m PositionReport) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m PositionReport) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m PositionReport) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

//SetFactor sets Factor, Tag 228
func (m PositionReport) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m PositionReport) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m PositionReport) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m PositionReport) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m PositionReport) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263
func (m PositionReport) SetSubscriptionRequestType(v string) {
	m.Set(field.NewSubscriptionRequestType(v))
}

//SetUnsolicitedIndicator sets UnsolicitedIndicator, Tag 325
func (m PositionReport) SetUnsolicitedIndicator(v bool) {
	m.Set(field.NewUnsolicitedIndicator(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m PositionReport) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m PositionReport) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m PositionReport) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m PositionReport) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m PositionReport) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m PositionReport) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetPriceType sets PriceType, Tag 423
func (m PositionReport) SetPriceType(v int) {
	m.Set(field.NewPriceType(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m PositionReport) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m PositionReport) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m PositionReport) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m PositionReport) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m PositionReport) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m PositionReport) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m PositionReport) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetRegistStatus sets RegistStatus, Tag 506
func (m PositionReport) SetRegistStatus(v string) {
	m.Set(field.NewRegistStatus(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m PositionReport) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m PositionReport) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetNoLegs sets NoLegs, Tag 555
func (m PositionReport) SetNoLegs(f NoLegsRepeatingGroup) {
	m.SetGroup(f)
}

//SetMatchStatus sets MatchStatus, Tag 573
func (m PositionReport) SetMatchStatus(v string) {
	m.Set(field.NewMatchStatus(v))
}

//SetAccountType sets AccountType, Tag 581
func (m PositionReport) SetAccountType(v int) {
	m.Set(field.NewAccountType(v))
}

//SetAcctIDSource sets AcctIDSource, Tag 660
func (m PositionReport) SetAcctIDSource(v int) {
	m.Set(field.NewAcctIDSource(v))
}

//SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m PositionReport) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

//SetPool sets Pool, Tag 691
func (m PositionReport) SetPool(v string) {
	m.Set(field.NewPool(v))
}

//SetNoPositions sets NoPositions, Tag 702
func (m PositionReport) SetNoPositions(f NoPositionsRepeatingGroup) {
	m.SetGroup(f)
}

//SetPosReqID sets PosReqID, Tag 710
func (m PositionReport) SetPosReqID(v string) {
	m.Set(field.NewPosReqID(v))
}

//SetNoUnderlyings sets NoUnderlyings, Tag 711
func (m PositionReport) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

//SetClearingBusinessDate sets ClearingBusinessDate, Tag 715
func (m PositionReport) SetClearingBusinessDate(v string) {
	m.Set(field.NewClearingBusinessDate(v))
}

//SetSettlSessID sets SettlSessID, Tag 716
func (m PositionReport) SetSettlSessID(v string) {
	m.Set(field.NewSettlSessID(v))
}

//SetSettlSessSubID sets SettlSessSubID, Tag 717
func (m PositionReport) SetSettlSessSubID(v string) {
	m.Set(field.NewSettlSessSubID(v))
}

//SetPosMaintRptID sets PosMaintRptID, Tag 721
func (m PositionReport) SetPosMaintRptID(v string) {
	m.Set(field.NewPosMaintRptID(v))
}

//SetPosReqType sets PosReqType, Tag 724
func (m PositionReport) SetPosReqType(v int) {
	m.Set(field.NewPosReqType(v))
}

//SetTotalNumPosReports sets TotalNumPosReports, Tag 727
func (m PositionReport) SetTotalNumPosReports(v int) {
	m.Set(field.NewTotalNumPosReports(v))
}

//SetPosReqResult sets PosReqResult, Tag 728
func (m PositionReport) SetPosReqResult(v int) {
	m.Set(field.NewPosReqResult(v))
}

//SetSettlPrice sets SettlPrice, Tag 730
func (m PositionReport) SetSettlPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewSettlPrice(value, scale))
}

//SetSettlPriceType sets SettlPriceType, Tag 731
func (m PositionReport) SetSettlPriceType(v int) {
	m.Set(field.NewSettlPriceType(v))
}

//SetPriorSettlPrice sets PriorSettlPrice, Tag 734
func (m PositionReport) SetPriorSettlPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPriorSettlPrice(value, scale))
}

//SetDeliveryDate sets DeliveryDate, Tag 743
func (m PositionReport) SetDeliveryDate(v string) {
	m.Set(field.NewDeliveryDate(v))
}

//SetNoPosAmt sets NoPosAmt, Tag 753
func (m PositionReport) SetNoPosAmt(f NoPosAmtRepeatingGroup) {
	m.SetGroup(f)
}

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m PositionReport) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

//SetPriceDelta sets PriceDelta, Tag 811
func (m PositionReport) SetPriceDelta(value decimal.Decimal, scale int32) {
	m.Set(field.NewPriceDelta(value, scale))
}

//SetNoEvents sets NoEvents, Tag 864
func (m PositionReport) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDatedDate sets DatedDate, Tag 873
func (m PositionReport) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

//SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m PositionReport) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

//SetCPProgram sets CPProgram, Tag 875
func (m PositionReport) SetCPProgram(v int) {
	m.Set(field.NewCPProgram(v))
}

//SetCPRegType sets CPRegType, Tag 876
func (m PositionReport) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

//SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m PositionReport) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

//SetSecurityStatus sets SecurityStatus, Tag 965
func (m PositionReport) SetSecurityStatus(v string) {
	m.Set(field.NewSecurityStatus(v))
}

//SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966
func (m PositionReport) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

//SetStrikeMultiplier sets StrikeMultiplier, Tag 967
func (m PositionReport) SetStrikeMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeMultiplier(value, scale))
}

//SetStrikeValue sets StrikeValue, Tag 968
func (m PositionReport) SetStrikeValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeValue(value, scale))
}

//SetMinPriceIncrement sets MinPriceIncrement, Tag 969
func (m PositionReport) SetMinPriceIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrement(value, scale))
}

//SetPositionLimit sets PositionLimit, Tag 970
func (m PositionReport) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

//SetNTPositionLimit sets NTPositionLimit, Tag 971
func (m PositionReport) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

//SetUnitOfMeasure sets UnitOfMeasure, Tag 996
func (m PositionReport) SetUnitOfMeasure(v string) {
	m.Set(field.NewUnitOfMeasure(v))
}

//SetTimeUnit sets TimeUnit, Tag 997
func (m PositionReport) SetTimeUnit(v string) {
	m.Set(field.NewTimeUnit(v))
}

//SetMessageEventSource sets MessageEventSource, Tag 1011
func (m PositionReport) SetMessageEventSource(v string) {
	m.Set(field.NewMessageEventSource(v))
}

//SetNoInstrumentParties sets NoInstrumentParties, Tag 1018
func (m PositionReport) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049
func (m PositionReport) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

//SetMaturityTime sets MaturityTime, Tag 1079
func (m PositionReport) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

//SetMinPriceIncrementAmount sets MinPriceIncrementAmount, Tag 1146
func (m PositionReport) SetMinPriceIncrementAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrementAmount(value, scale))
}

//SetUnitOfMeasureQty sets UnitOfMeasureQty, Tag 1147
func (m PositionReport) SetUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnitOfMeasureQty(value, scale))
}

//SetSecurityGroup sets SecurityGroup, Tag 1151
func (m PositionReport) SetSecurityGroup(v string) {
	m.Set(field.NewSecurityGroup(v))
}

//SetApplID sets ApplID, Tag 1180
func (m PositionReport) SetApplID(v string) {
	m.Set(field.NewApplID(v))
}

//SetApplSeqNum sets ApplSeqNum, Tag 1181
func (m PositionReport) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

//SetSecurityXMLLen sets SecurityXMLLen, Tag 1184
func (m PositionReport) SetSecurityXMLLen(v int) {
	m.Set(field.NewSecurityXMLLen(v))
}

//SetSecurityXML sets SecurityXML, Tag 1185
func (m PositionReport) SetSecurityXML(v string) {
	m.Set(field.NewSecurityXML(v))
}

//SetSecurityXMLSchema sets SecurityXMLSchema, Tag 1186
func (m PositionReport) SetSecurityXMLSchema(v string) {
	m.Set(field.NewSecurityXMLSchema(v))
}

//SetPriceUnitOfMeasure sets PriceUnitOfMeasure, Tag 1191
func (m PositionReport) SetPriceUnitOfMeasure(v string) {
	m.Set(field.NewPriceUnitOfMeasure(v))
}

//SetPriceUnitOfMeasureQty sets PriceUnitOfMeasureQty, Tag 1192
func (m PositionReport) SetPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewPriceUnitOfMeasureQty(value, scale))
}

//SetSettlMethod sets SettlMethod, Tag 1193
func (m PositionReport) SetSettlMethod(v string) {
	m.Set(field.NewSettlMethod(v))
}

//SetExerciseStyle sets ExerciseStyle, Tag 1194
func (m PositionReport) SetExerciseStyle(v int) {
	m.Set(field.NewExerciseStyle(v))
}

//SetOptPayoutAmount sets OptPayoutAmount, Tag 1195
func (m PositionReport) SetOptPayoutAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewOptPayoutAmount(value, scale))
}

//SetPriceQuoteMethod sets PriceQuoteMethod, Tag 1196
func (m PositionReport) SetPriceQuoteMethod(v string) {
	m.Set(field.NewPriceQuoteMethod(v))
}

//SetValuationMethod sets ValuationMethod, Tag 1197
func (m PositionReport) SetValuationMethod(v string) {
	m.Set(field.NewValuationMethod(v))
}

//SetListMethod sets ListMethod, Tag 1198
func (m PositionReport) SetListMethod(v int) {
	m.Set(field.NewListMethod(v))
}

//SetCapPrice sets CapPrice, Tag 1199
func (m PositionReport) SetCapPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewCapPrice(value, scale))
}

//SetFloorPrice sets FloorPrice, Tag 1200
func (m PositionReport) SetFloorPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewFloorPrice(value, scale))
}

//SetProductComplex sets ProductComplex, Tag 1227
func (m PositionReport) SetProductComplex(v string) {
	m.Set(field.NewProductComplex(v))
}

//SetFlexProductEligibilityIndicator sets FlexProductEligibilityIndicator, Tag 1242
func (m PositionReport) SetFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewFlexProductEligibilityIndicator(v))
}

//SetFlexibleIndicator sets FlexibleIndicator, Tag 1244
func (m PositionReport) SetFlexibleIndicator(v bool) {
	m.Set(field.NewFlexibleIndicator(v))
}

//SetApplLastSeqNum sets ApplLastSeqNum, Tag 1350
func (m PositionReport) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

//SetApplResendFlag sets ApplResendFlag, Tag 1352
func (m PositionReport) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

//SetModelType sets ModelType, Tag 1434
func (m PositionReport) SetModelType(v int) {
	m.Set(field.NewModelType(v))
}

//SetContractMultiplierUnit sets ContractMultiplierUnit, Tag 1435
func (m PositionReport) SetContractMultiplierUnit(v int) {
	m.Set(field.NewContractMultiplierUnit(v))
}

//SetFlowScheduleType sets FlowScheduleType, Tag 1439
func (m PositionReport) SetFlowScheduleType(v int) {
	m.Set(field.NewFlowScheduleType(v))
}

//SetRestructuringType sets RestructuringType, Tag 1449
func (m PositionReport) SetRestructuringType(v string) {
	m.Set(field.NewRestructuringType(v))
}

//SetSeniority sets Seniority, Tag 1450
func (m PositionReport) SetSeniority(v string) {
	m.Set(field.NewSeniority(v))
}

//SetNotionalPercentageOutstanding sets NotionalPercentageOutstanding, Tag 1451
func (m PositionReport) SetNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewNotionalPercentageOutstanding(value, scale))
}

//SetOriginalNotionalPercentageOutstanding sets OriginalNotionalPercentageOutstanding, Tag 1452
func (m PositionReport) SetOriginalNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewOriginalNotionalPercentageOutstanding(value, scale))
}

//SetAttachmentPoint sets AttachmentPoint, Tag 1457
func (m PositionReport) SetAttachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewAttachmentPoint(value, scale))
}

//SetDetachmentPoint sets DetachmentPoint, Tag 1458
func (m PositionReport) SetDetachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewDetachmentPoint(value, scale))
}

//SetStrikePriceDeterminationMethod sets StrikePriceDeterminationMethod, Tag 1478
func (m PositionReport) SetStrikePriceDeterminationMethod(v int) {
	m.Set(field.NewStrikePriceDeterminationMethod(v))
}

//SetStrikePriceBoundaryMethod sets StrikePriceBoundaryMethod, Tag 1479
func (m PositionReport) SetStrikePriceBoundaryMethod(v int) {
	m.Set(field.NewStrikePriceBoundaryMethod(v))
}

//SetStrikePriceBoundaryPrecision sets StrikePriceBoundaryPrecision, Tag 1480
func (m PositionReport) SetStrikePriceBoundaryPrecision(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePriceBoundaryPrecision(value, scale))
}

//SetUnderlyingPriceDeterminationMethod sets UnderlyingPriceDeterminationMethod, Tag 1481
func (m PositionReport) SetUnderlyingPriceDeterminationMethod(v int) {
	m.Set(field.NewUnderlyingPriceDeterminationMethod(v))
}

//SetOptPayoutType sets OptPayoutType, Tag 1482
func (m PositionReport) SetOptPayoutType(v int) {
	m.Set(field.NewOptPayoutType(v))
}

//SetNoComplexEvents sets NoComplexEvents, Tag 1483
func (m PositionReport) SetNoComplexEvents(f NoComplexEventsRepeatingGroup) {
	m.SetGroup(f)
}

//GetAccount gets Account, Tag 1
func (m PositionReport) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m PositionReport) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m PositionReport) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m PositionReport) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m PositionReport) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m PositionReport) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m PositionReport) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m PositionReport) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m PositionReport) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m PositionReport) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m PositionReport) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m PositionReport) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m PositionReport) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m PositionReport) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m PositionReport) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m PositionReport) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m PositionReport) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m PositionReport) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m PositionReport) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m PositionReport) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m PositionReport) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m PositionReport) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m PositionReport) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m PositionReport) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m PositionReport) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m PositionReport) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263
func (m PositionReport) GetSubscriptionRequestType() (f field.SubscriptionRequestTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnsolicitedIndicator gets UnsolicitedIndicator, Tag 325
func (m PositionReport) GetUnsolicitedIndicator() (f field.UnsolicitedIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m PositionReport) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m PositionReport) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m PositionReport) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m PositionReport) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m PositionReport) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m PositionReport) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceType gets PriceType, Tag 423
func (m PositionReport) GetPriceType() (f field.PriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m PositionReport) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m PositionReport) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m PositionReport) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m PositionReport) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m PositionReport) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m PositionReport) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m PositionReport) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRegistStatus gets RegistStatus, Tag 506
func (m PositionReport) GetRegistStatus() (f field.RegistStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m PositionReport) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m PositionReport) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoLegs gets NoLegs, Tag 555
func (m PositionReport) GetNoLegs() (NoLegsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetMatchStatus gets MatchStatus, Tag 573
func (m PositionReport) GetMatchStatus() (f field.MatchStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccountType gets AccountType, Tag 581
func (m PositionReport) GetAccountType() (f field.AccountTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAcctIDSource gets AcctIDSource, Tag 660
func (m PositionReport) GetAcctIDSource() (f field.AcctIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m PositionReport) GetContractSettlMonth() (f field.ContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPool gets Pool, Tag 691
func (m PositionReport) GetPool() (f field.PoolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPositions gets NoPositions, Tag 702
func (m PositionReport) GetNoPositions() (NoPositionsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPositionsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetPosReqID gets PosReqID, Tag 710
func (m PositionReport) GetPosReqID() (f field.PosReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyings gets NoUnderlyings, Tag 711
func (m PositionReport) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetClearingBusinessDate gets ClearingBusinessDate, Tag 715
func (m PositionReport) GetClearingBusinessDate() (f field.ClearingBusinessDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlSessID gets SettlSessID, Tag 716
func (m PositionReport) GetSettlSessID() (f field.SettlSessIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlSessSubID gets SettlSessSubID, Tag 717
func (m PositionReport) GetSettlSessSubID() (f field.SettlSessSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPosMaintRptID gets PosMaintRptID, Tag 721
func (m PositionReport) GetPosMaintRptID() (f field.PosMaintRptIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPosReqType gets PosReqType, Tag 724
func (m PositionReport) GetPosReqType() (f field.PosReqTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTotalNumPosReports gets TotalNumPosReports, Tag 727
func (m PositionReport) GetTotalNumPosReports() (f field.TotalNumPosReportsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPosReqResult gets PosReqResult, Tag 728
func (m PositionReport) GetPosReqResult() (f field.PosReqResultField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlPrice gets SettlPrice, Tag 730
func (m PositionReport) GetSettlPrice() (f field.SettlPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlPriceType gets SettlPriceType, Tag 731
func (m PositionReport) GetSettlPriceType() (f field.SettlPriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriorSettlPrice gets PriorSettlPrice, Tag 734
func (m PositionReport) GetPriorSettlPrice() (f field.PriorSettlPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDeliveryDate gets DeliveryDate, Tag 743
func (m PositionReport) GetDeliveryDate() (f field.DeliveryDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPosAmt gets NoPosAmt, Tag 753
func (m PositionReport) GetNoPosAmt() (NoPosAmtRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPosAmtRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m PositionReport) GetSecuritySubType() (f field.SecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceDelta gets PriceDelta, Tag 811
func (m PositionReport) GetPriceDelta() (f field.PriceDeltaField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoEvents gets NoEvents, Tag 864
func (m PositionReport) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDatedDate gets DatedDate, Tag 873
func (m PositionReport) GetDatedDate() (f field.DatedDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m PositionReport) GetInterestAccrualDate() (f field.InterestAccrualDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPProgram gets CPProgram, Tag 875
func (m PositionReport) GetCPProgram() (f field.CPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPRegType gets CPRegType, Tag 876
func (m PositionReport) GetCPRegType() (f field.CPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m PositionReport) GetStrikeCurrency() (f field.StrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityStatus gets SecurityStatus, Tag 965
func (m PositionReport) GetSecurityStatus() (f field.SecurityStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966
func (m PositionReport) GetSettleOnOpenFlag() (f field.SettleOnOpenFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeMultiplier gets StrikeMultiplier, Tag 967
func (m PositionReport) GetStrikeMultiplier() (f field.StrikeMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeValue gets StrikeValue, Tag 968
func (m PositionReport) GetStrikeValue() (f field.StrikeValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrement gets MinPriceIncrement, Tag 969
func (m PositionReport) GetMinPriceIncrement() (f field.MinPriceIncrementField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionLimit gets PositionLimit, Tag 970
func (m PositionReport) GetPositionLimit() (f field.PositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNTPositionLimit gets NTPositionLimit, Tag 971
func (m PositionReport) GetNTPositionLimit() (f field.NTPositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnitOfMeasure gets UnitOfMeasure, Tag 996
func (m PositionReport) GetUnitOfMeasure() (f field.UnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeUnit gets TimeUnit, Tag 997
func (m PositionReport) GetTimeUnit() (f field.TimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMessageEventSource gets MessageEventSource, Tag 1011
func (m PositionReport) GetMessageEventSource() (f field.MessageEventSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoInstrumentParties gets NoInstrumentParties, Tag 1018
func (m PositionReport) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049
func (m PositionReport) GetInstrmtAssignmentMethod() (f field.InstrmtAssignmentMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityTime gets MaturityTime, Tag 1079
func (m PositionReport) GetMaturityTime() (f field.MaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrementAmount gets MinPriceIncrementAmount, Tag 1146
func (m PositionReport) GetMinPriceIncrementAmount() (f field.MinPriceIncrementAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnitOfMeasureQty gets UnitOfMeasureQty, Tag 1147
func (m PositionReport) GetUnitOfMeasureQty() (f field.UnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityGroup gets SecurityGroup, Tag 1151
func (m PositionReport) GetSecurityGroup() (f field.SecurityGroupField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetApplID gets ApplID, Tag 1180
func (m PositionReport) GetApplID() (f field.ApplIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetApplSeqNum gets ApplSeqNum, Tag 1181
func (m PositionReport) GetApplSeqNum() (f field.ApplSeqNumField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXMLLen gets SecurityXMLLen, Tag 1184
func (m PositionReport) GetSecurityXMLLen() (f field.SecurityXMLLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXML gets SecurityXML, Tag 1185
func (m PositionReport) GetSecurityXML() (f field.SecurityXMLField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXMLSchema gets SecurityXMLSchema, Tag 1186
func (m PositionReport) GetSecurityXMLSchema() (f field.SecurityXMLSchemaField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceUnitOfMeasure gets PriceUnitOfMeasure, Tag 1191
func (m PositionReport) GetPriceUnitOfMeasure() (f field.PriceUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceUnitOfMeasureQty gets PriceUnitOfMeasureQty, Tag 1192
func (m PositionReport) GetPriceUnitOfMeasureQty() (f field.PriceUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlMethod gets SettlMethod, Tag 1193
func (m PositionReport) GetSettlMethod() (f field.SettlMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExerciseStyle gets ExerciseStyle, Tag 1194
func (m PositionReport) GetExerciseStyle() (f field.ExerciseStyleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptPayoutAmount gets OptPayoutAmount, Tag 1195
func (m PositionReport) GetOptPayoutAmount() (f field.OptPayoutAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceQuoteMethod gets PriceQuoteMethod, Tag 1196
func (m PositionReport) GetPriceQuoteMethod() (f field.PriceQuoteMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetValuationMethod gets ValuationMethod, Tag 1197
func (m PositionReport) GetValuationMethod() (f field.ValuationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListMethod gets ListMethod, Tag 1198
func (m PositionReport) GetListMethod() (f field.ListMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCapPrice gets CapPrice, Tag 1199
func (m PositionReport) GetCapPrice() (f field.CapPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFloorPrice gets FloorPrice, Tag 1200
func (m PositionReport) GetFloorPrice() (f field.FloorPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProductComplex gets ProductComplex, Tag 1227
func (m PositionReport) GetProductComplex() (f field.ProductComplexField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlexProductEligibilityIndicator gets FlexProductEligibilityIndicator, Tag 1242
func (m PositionReport) GetFlexProductEligibilityIndicator() (f field.FlexProductEligibilityIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlexibleIndicator gets FlexibleIndicator, Tag 1244
func (m PositionReport) GetFlexibleIndicator() (f field.FlexibleIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetApplLastSeqNum gets ApplLastSeqNum, Tag 1350
func (m PositionReport) GetApplLastSeqNum() (f field.ApplLastSeqNumField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetApplResendFlag gets ApplResendFlag, Tag 1352
func (m PositionReport) GetApplResendFlag() (f field.ApplResendFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetModelType gets ModelType, Tag 1434
func (m PositionReport) GetModelType() (f field.ModelTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplierUnit gets ContractMultiplierUnit, Tag 1435
func (m PositionReport) GetContractMultiplierUnit() (f field.ContractMultiplierUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlowScheduleType gets FlowScheduleType, Tag 1439
func (m PositionReport) GetFlowScheduleType() (f field.FlowScheduleTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRestructuringType gets RestructuringType, Tag 1449
func (m PositionReport) GetRestructuringType() (f field.RestructuringTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSeniority gets Seniority, Tag 1450
func (m PositionReport) GetSeniority() (f field.SeniorityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNotionalPercentageOutstanding gets NotionalPercentageOutstanding, Tag 1451
func (m PositionReport) GetNotionalPercentageOutstanding() (f field.NotionalPercentageOutstandingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOriginalNotionalPercentageOutstanding gets OriginalNotionalPercentageOutstanding, Tag 1452
func (m PositionReport) GetOriginalNotionalPercentageOutstanding() (f field.OriginalNotionalPercentageOutstandingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAttachmentPoint gets AttachmentPoint, Tag 1457
func (m PositionReport) GetAttachmentPoint() (f field.AttachmentPointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDetachmentPoint gets DetachmentPoint, Tag 1458
func (m PositionReport) GetDetachmentPoint() (f field.DetachmentPointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePriceDeterminationMethod gets StrikePriceDeterminationMethod, Tag 1478
func (m PositionReport) GetStrikePriceDeterminationMethod() (f field.StrikePriceDeterminationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePriceBoundaryMethod gets StrikePriceBoundaryMethod, Tag 1479
func (m PositionReport) GetStrikePriceBoundaryMethod() (f field.StrikePriceBoundaryMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePriceBoundaryPrecision gets StrikePriceBoundaryPrecision, Tag 1480
func (m PositionReport) GetStrikePriceBoundaryPrecision() (f field.StrikePriceBoundaryPrecisionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPriceDeterminationMethod gets UnderlyingPriceDeterminationMethod, Tag 1481
func (m PositionReport) GetUnderlyingPriceDeterminationMethod() (f field.UnderlyingPriceDeterminationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptPayoutType gets OptPayoutType, Tag 1482
func (m PositionReport) GetOptPayoutType() (f field.OptPayoutTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoComplexEvents gets NoComplexEvents, Tag 1483
func (m PositionReport) GetNoComplexEvents() (NoComplexEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasAccount returns true if Account is present, Tag 1
func (m PositionReport) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m PositionReport) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m PositionReport) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m PositionReport) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m PositionReport) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m PositionReport) HasText() bool {
	return m.Has(tag.Text)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m PositionReport) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m PositionReport) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m PositionReport) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m PositionReport) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m PositionReport) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m PositionReport) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m PositionReport) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m PositionReport) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m PositionReport) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m PositionReport) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m PositionReport) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m PositionReport) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m PositionReport) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m PositionReport) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m PositionReport) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m PositionReport) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m PositionReport) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m PositionReport) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m PositionReport) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m PositionReport) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263
func (m PositionReport) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

//HasUnsolicitedIndicator returns true if UnsolicitedIndicator is present, Tag 325
func (m PositionReport) HasUnsolicitedIndicator() bool {
	return m.Has(tag.UnsolicitedIndicator)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m PositionReport) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m PositionReport) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m PositionReport) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m PositionReport) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m PositionReport) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m PositionReport) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m PositionReport) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m PositionReport) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m PositionReport) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m PositionReport) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m PositionReport) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m PositionReport) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m PositionReport) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m PositionReport) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasRegistStatus returns true if RegistStatus is present, Tag 506
func (m PositionReport) HasRegistStatus() bool {
	return m.Has(tag.RegistStatus)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m PositionReport) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m PositionReport) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasNoLegs returns true if NoLegs is present, Tag 555
func (m PositionReport) HasNoLegs() bool {
	return m.Has(tag.NoLegs)
}

//HasMatchStatus returns true if MatchStatus is present, Tag 573
func (m PositionReport) HasMatchStatus() bool {
	return m.Has(tag.MatchStatus)
}

//HasAccountType returns true if AccountType is present, Tag 581
func (m PositionReport) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

//HasAcctIDSource returns true if AcctIDSource is present, Tag 660
func (m PositionReport) HasAcctIDSource() bool {
	return m.Has(tag.AcctIDSource)
}

//HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m PositionReport) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

//HasPool returns true if Pool is present, Tag 691
func (m PositionReport) HasPool() bool {
	return m.Has(tag.Pool)
}

//HasNoPositions returns true if NoPositions is present, Tag 702
func (m PositionReport) HasNoPositions() bool {
	return m.Has(tag.NoPositions)
}

//HasPosReqID returns true if PosReqID is present, Tag 710
func (m PositionReport) HasPosReqID() bool {
	return m.Has(tag.PosReqID)
}

//HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711
func (m PositionReport) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

//HasClearingBusinessDate returns true if ClearingBusinessDate is present, Tag 715
func (m PositionReport) HasClearingBusinessDate() bool {
	return m.Has(tag.ClearingBusinessDate)
}

//HasSettlSessID returns true if SettlSessID is present, Tag 716
func (m PositionReport) HasSettlSessID() bool {
	return m.Has(tag.SettlSessID)
}

//HasSettlSessSubID returns true if SettlSessSubID is present, Tag 717
func (m PositionReport) HasSettlSessSubID() bool {
	return m.Has(tag.SettlSessSubID)
}

//HasPosMaintRptID returns true if PosMaintRptID is present, Tag 721
func (m PositionReport) HasPosMaintRptID() bool {
	return m.Has(tag.PosMaintRptID)
}

//HasPosReqType returns true if PosReqType is present, Tag 724
func (m PositionReport) HasPosReqType() bool {
	return m.Has(tag.PosReqType)
}

//HasTotalNumPosReports returns true if TotalNumPosReports is present, Tag 727
func (m PositionReport) HasTotalNumPosReports() bool {
	return m.Has(tag.TotalNumPosReports)
}

//HasPosReqResult returns true if PosReqResult is present, Tag 728
func (m PositionReport) HasPosReqResult() bool {
	return m.Has(tag.PosReqResult)
}

//HasSettlPrice returns true if SettlPrice is present, Tag 730
func (m PositionReport) HasSettlPrice() bool {
	return m.Has(tag.SettlPrice)
}

//HasSettlPriceType returns true if SettlPriceType is present, Tag 731
func (m PositionReport) HasSettlPriceType() bool {
	return m.Has(tag.SettlPriceType)
}

//HasPriorSettlPrice returns true if PriorSettlPrice is present, Tag 734
func (m PositionReport) HasPriorSettlPrice() bool {
	return m.Has(tag.PriorSettlPrice)
}

//HasDeliveryDate returns true if DeliveryDate is present, Tag 743
func (m PositionReport) HasDeliveryDate() bool {
	return m.Has(tag.DeliveryDate)
}

//HasNoPosAmt returns true if NoPosAmt is present, Tag 753
func (m PositionReport) HasNoPosAmt() bool {
	return m.Has(tag.NoPosAmt)
}

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m PositionReport) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

//HasPriceDelta returns true if PriceDelta is present, Tag 811
func (m PositionReport) HasPriceDelta() bool {
	return m.Has(tag.PriceDelta)
}

//HasNoEvents returns true if NoEvents is present, Tag 864
func (m PositionReport) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

//HasDatedDate returns true if DatedDate is present, Tag 873
func (m PositionReport) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

//HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m PositionReport) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

//HasCPProgram returns true if CPProgram is present, Tag 875
func (m PositionReport) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

//HasCPRegType returns true if CPRegType is present, Tag 876
func (m PositionReport) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

//HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m PositionReport) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

//HasSecurityStatus returns true if SecurityStatus is present, Tag 965
func (m PositionReport) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

//HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966
func (m PositionReport) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

//HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967
func (m PositionReport) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

//HasStrikeValue returns true if StrikeValue is present, Tag 968
func (m PositionReport) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

//HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969
func (m PositionReport) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

//HasPositionLimit returns true if PositionLimit is present, Tag 970
func (m PositionReport) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

//HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971
func (m PositionReport) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

//HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996
func (m PositionReport) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

//HasTimeUnit returns true if TimeUnit is present, Tag 997
func (m PositionReport) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

//HasMessageEventSource returns true if MessageEventSource is present, Tag 1011
func (m PositionReport) HasMessageEventSource() bool {
	return m.Has(tag.MessageEventSource)
}

//HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018
func (m PositionReport) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

//HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049
func (m PositionReport) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

//HasMaturityTime returns true if MaturityTime is present, Tag 1079
func (m PositionReport) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

//HasMinPriceIncrementAmount returns true if MinPriceIncrementAmount is present, Tag 1146
func (m PositionReport) HasMinPriceIncrementAmount() bool {
	return m.Has(tag.MinPriceIncrementAmount)
}

//HasUnitOfMeasureQty returns true if UnitOfMeasureQty is present, Tag 1147
func (m PositionReport) HasUnitOfMeasureQty() bool {
	return m.Has(tag.UnitOfMeasureQty)
}

//HasSecurityGroup returns true if SecurityGroup is present, Tag 1151
func (m PositionReport) HasSecurityGroup() bool {
	return m.Has(tag.SecurityGroup)
}

//HasApplID returns true if ApplID is present, Tag 1180
func (m PositionReport) HasApplID() bool {
	return m.Has(tag.ApplID)
}

//HasApplSeqNum returns true if ApplSeqNum is present, Tag 1181
func (m PositionReport) HasApplSeqNum() bool {
	return m.Has(tag.ApplSeqNum)
}

//HasSecurityXMLLen returns true if SecurityXMLLen is present, Tag 1184
func (m PositionReport) HasSecurityXMLLen() bool {
	return m.Has(tag.SecurityXMLLen)
}

//HasSecurityXML returns true if SecurityXML is present, Tag 1185
func (m PositionReport) HasSecurityXML() bool {
	return m.Has(tag.SecurityXML)
}

//HasSecurityXMLSchema returns true if SecurityXMLSchema is present, Tag 1186
func (m PositionReport) HasSecurityXMLSchema() bool {
	return m.Has(tag.SecurityXMLSchema)
}

//HasPriceUnitOfMeasure returns true if PriceUnitOfMeasure is present, Tag 1191
func (m PositionReport) HasPriceUnitOfMeasure() bool {
	return m.Has(tag.PriceUnitOfMeasure)
}

//HasPriceUnitOfMeasureQty returns true if PriceUnitOfMeasureQty is present, Tag 1192
func (m PositionReport) HasPriceUnitOfMeasureQty() bool {
	return m.Has(tag.PriceUnitOfMeasureQty)
}

//HasSettlMethod returns true if SettlMethod is present, Tag 1193
func (m PositionReport) HasSettlMethod() bool {
	return m.Has(tag.SettlMethod)
}

//HasExerciseStyle returns true if ExerciseStyle is present, Tag 1194
func (m PositionReport) HasExerciseStyle() bool {
	return m.Has(tag.ExerciseStyle)
}

//HasOptPayoutAmount returns true if OptPayoutAmount is present, Tag 1195
func (m PositionReport) HasOptPayoutAmount() bool {
	return m.Has(tag.OptPayoutAmount)
}

//HasPriceQuoteMethod returns true if PriceQuoteMethod is present, Tag 1196
func (m PositionReport) HasPriceQuoteMethod() bool {
	return m.Has(tag.PriceQuoteMethod)
}

//HasValuationMethod returns true if ValuationMethod is present, Tag 1197
func (m PositionReport) HasValuationMethod() bool {
	return m.Has(tag.ValuationMethod)
}

//HasListMethod returns true if ListMethod is present, Tag 1198
func (m PositionReport) HasListMethod() bool {
	return m.Has(tag.ListMethod)
}

//HasCapPrice returns true if CapPrice is present, Tag 1199
func (m PositionReport) HasCapPrice() bool {
	return m.Has(tag.CapPrice)
}

//HasFloorPrice returns true if FloorPrice is present, Tag 1200
func (m PositionReport) HasFloorPrice() bool {
	return m.Has(tag.FloorPrice)
}

//HasProductComplex returns true if ProductComplex is present, Tag 1227
func (m PositionReport) HasProductComplex() bool {
	return m.Has(tag.ProductComplex)
}

//HasFlexProductEligibilityIndicator returns true if FlexProductEligibilityIndicator is present, Tag 1242
func (m PositionReport) HasFlexProductEligibilityIndicator() bool {
	return m.Has(tag.FlexProductEligibilityIndicator)
}

//HasFlexibleIndicator returns true if FlexibleIndicator is present, Tag 1244
func (m PositionReport) HasFlexibleIndicator() bool {
	return m.Has(tag.FlexibleIndicator)
}

//HasApplLastSeqNum returns true if ApplLastSeqNum is present, Tag 1350
func (m PositionReport) HasApplLastSeqNum() bool {
	return m.Has(tag.ApplLastSeqNum)
}

//HasApplResendFlag returns true if ApplResendFlag is present, Tag 1352
func (m PositionReport) HasApplResendFlag() bool {
	return m.Has(tag.ApplResendFlag)
}

//HasModelType returns true if ModelType is present, Tag 1434
func (m PositionReport) HasModelType() bool {
	return m.Has(tag.ModelType)
}

//HasContractMultiplierUnit returns true if ContractMultiplierUnit is present, Tag 1435
func (m PositionReport) HasContractMultiplierUnit() bool {
	return m.Has(tag.ContractMultiplierUnit)
}

//HasFlowScheduleType returns true if FlowScheduleType is present, Tag 1439
func (m PositionReport) HasFlowScheduleType() bool {
	return m.Has(tag.FlowScheduleType)
}

//HasRestructuringType returns true if RestructuringType is present, Tag 1449
func (m PositionReport) HasRestructuringType() bool {
	return m.Has(tag.RestructuringType)
}

//HasSeniority returns true if Seniority is present, Tag 1450
func (m PositionReport) HasSeniority() bool {
	return m.Has(tag.Seniority)
}

//HasNotionalPercentageOutstanding returns true if NotionalPercentageOutstanding is present, Tag 1451
func (m PositionReport) HasNotionalPercentageOutstanding() bool {
	return m.Has(tag.NotionalPercentageOutstanding)
}

//HasOriginalNotionalPercentageOutstanding returns true if OriginalNotionalPercentageOutstanding is present, Tag 1452
func (m PositionReport) HasOriginalNotionalPercentageOutstanding() bool {
	return m.Has(tag.OriginalNotionalPercentageOutstanding)
}

//HasAttachmentPoint returns true if AttachmentPoint is present, Tag 1457
func (m PositionReport) HasAttachmentPoint() bool {
	return m.Has(tag.AttachmentPoint)
}

//HasDetachmentPoint returns true if DetachmentPoint is present, Tag 1458
func (m PositionReport) HasDetachmentPoint() bool {
	return m.Has(tag.DetachmentPoint)
}

//HasStrikePriceDeterminationMethod returns true if StrikePriceDeterminationMethod is present, Tag 1478
func (m PositionReport) HasStrikePriceDeterminationMethod() bool {
	return m.Has(tag.StrikePriceDeterminationMethod)
}

//HasStrikePriceBoundaryMethod returns true if StrikePriceBoundaryMethod is present, Tag 1479
func (m PositionReport) HasStrikePriceBoundaryMethod() bool {
	return m.Has(tag.StrikePriceBoundaryMethod)
}

//HasStrikePriceBoundaryPrecision returns true if StrikePriceBoundaryPrecision is present, Tag 1480
func (m PositionReport) HasStrikePriceBoundaryPrecision() bool {
	return m.Has(tag.StrikePriceBoundaryPrecision)
}

//HasUnderlyingPriceDeterminationMethod returns true if UnderlyingPriceDeterminationMethod is present, Tag 1481
func (m PositionReport) HasUnderlyingPriceDeterminationMethod() bool {
	return m.Has(tag.UnderlyingPriceDeterminationMethod)
}

//HasOptPayoutType returns true if OptPayoutType is present, Tag 1482
func (m PositionReport) HasOptPayoutType() bool {
	return m.Has(tag.OptPayoutType)
}

//HasNoComplexEvents returns true if NoComplexEvents is present, Tag 1483
func (m PositionReport) HasNoComplexEvents() bool {
	return m.Has(tag.NoComplexEvents)
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

//SetLegOptionRatio sets LegOptionRatio, Tag 1017
func (m NoLegs) SetLegOptionRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegOptionRatio(value, scale))
}

//SetLegPrice sets LegPrice, Tag 566
func (m NoLegs) SetLegPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegPrice(value, scale))
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
func (m NoLegs) SetLegUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegUnitOfMeasureQty(value, scale))
}

//SetLegPriceUnitOfMeasure sets LegPriceUnitOfMeasure, Tag 1421
func (m NoLegs) SetLegPriceUnitOfMeasure(v string) {
	m.Set(field.NewLegPriceUnitOfMeasure(v))
}

//SetLegPriceUnitOfMeasureQty sets LegPriceUnitOfMeasureQty, Tag 1422
func (m NoLegs) SetLegPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegPriceUnitOfMeasureQty(value, scale))
}

//SetLegContractMultiplierUnit sets LegContractMultiplierUnit, Tag 1436
func (m NoLegs) SetLegContractMultiplierUnit(v int) {
	m.Set(field.NewLegContractMultiplierUnit(v))
}

//SetLegFlowScheduleType sets LegFlowScheduleType, Tag 1440
func (m NoLegs) SetLegFlowScheduleType(v int) {
	m.Set(field.NewLegFlowScheduleType(v))
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

//GetLegContractMultiplierUnit gets LegContractMultiplierUnit, Tag 1436
func (m NoLegs) GetLegContractMultiplierUnit() (f field.LegContractMultiplierUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegFlowScheduleType gets LegFlowScheduleType, Tag 1440
func (m NoLegs) GetLegFlowScheduleType() (f field.LegFlowScheduleTypeField, err quickfix.MessageRejectError) {
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

//HasLegContractMultiplierUnit returns true if LegContractMultiplierUnit is present, Tag 1436
func (m NoLegs) HasLegContractMultiplierUnit() bool {
	return m.Has(tag.LegContractMultiplierUnit)
}

//HasLegFlowScheduleType returns true if LegFlowScheduleType is present, Tag 1440
func (m NoLegs) HasLegFlowScheduleType() bool {
	return m.Has(tag.LegFlowScheduleType)
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

//NoLegsRepeatingGroup is a repeating group, Tag 555
type NoLegsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegsRepeatingGroup returns an initialized, NoLegsRepeatingGroup
func NewNoLegsRepeatingGroup() NoLegsRepeatingGroup {
	return NoLegsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSymbol), quickfix.GroupElement(tag.LegSymbolSfx), quickfix.GroupElement(tag.LegSecurityID), quickfix.GroupElement(tag.LegSecurityIDSource), NewNoLegSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.LegProduct), quickfix.GroupElement(tag.LegCFICode), quickfix.GroupElement(tag.LegSecurityType), quickfix.GroupElement(tag.LegSecuritySubType), quickfix.GroupElement(tag.LegMaturityMonthYear), quickfix.GroupElement(tag.LegMaturityDate), quickfix.GroupElement(tag.LegCouponPaymentDate), quickfix.GroupElement(tag.LegIssueDate), quickfix.GroupElement(tag.LegRepoCollateralSecurityType), quickfix.GroupElement(tag.LegRepurchaseTerm), quickfix.GroupElement(tag.LegRepurchaseRate), quickfix.GroupElement(tag.LegFactor), quickfix.GroupElement(tag.LegCreditRating), quickfix.GroupElement(tag.LegInstrRegistry), quickfix.GroupElement(tag.LegCountryOfIssue), quickfix.GroupElement(tag.LegStateOrProvinceOfIssue), quickfix.GroupElement(tag.LegLocaleOfIssue), quickfix.GroupElement(tag.LegRedemptionDate), quickfix.GroupElement(tag.LegStrikePrice), quickfix.GroupElement(tag.LegStrikeCurrency), quickfix.GroupElement(tag.LegOptAttribute), quickfix.GroupElement(tag.LegContractMultiplier), quickfix.GroupElement(tag.LegCouponRate), quickfix.GroupElement(tag.LegSecurityExchange), quickfix.GroupElement(tag.LegIssuer), quickfix.GroupElement(tag.EncodedLegIssuerLen), quickfix.GroupElement(tag.EncodedLegIssuer), quickfix.GroupElement(tag.LegSecurityDesc), quickfix.GroupElement(tag.EncodedLegSecurityDescLen), quickfix.GroupElement(tag.EncodedLegSecurityDesc), quickfix.GroupElement(tag.LegRatioQty), quickfix.GroupElement(tag.LegSide), quickfix.GroupElement(tag.LegCurrency), quickfix.GroupElement(tag.LegPool), quickfix.GroupElement(tag.LegDatedDate), quickfix.GroupElement(tag.LegContractSettlMonth), quickfix.GroupElement(tag.LegInterestAccrualDate), quickfix.GroupElement(tag.LegUnitOfMeasure), quickfix.GroupElement(tag.LegTimeUnit), quickfix.GroupElement(tag.LegOptionRatio), quickfix.GroupElement(tag.LegPrice), quickfix.GroupElement(tag.LegMaturityTime), quickfix.GroupElement(tag.LegPutOrCall), quickfix.GroupElement(tag.LegExerciseStyle), quickfix.GroupElement(tag.LegUnitOfMeasureQty), quickfix.GroupElement(tag.LegPriceUnitOfMeasure), quickfix.GroupElement(tag.LegPriceUnitOfMeasureQty), quickfix.GroupElement(tag.LegContractMultiplierUnit), quickfix.GroupElement(tag.LegFlowScheduleType)})}
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

//NoPositions is a repeating group element, Tag 702
type NoPositions struct {
	quickfix.Group
}

//SetPosType sets PosType, Tag 703
func (m NoPositions) SetPosType(v string) {
	m.Set(field.NewPosType(v))
}

//SetLongQty sets LongQty, Tag 704
func (m NoPositions) SetLongQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLongQty(value, scale))
}

//SetShortQty sets ShortQty, Tag 705
func (m NoPositions) SetShortQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewShortQty(value, scale))
}

//SetPosQtyStatus sets PosQtyStatus, Tag 706
func (m NoPositions) SetPosQtyStatus(v int) {
	m.Set(field.NewPosQtyStatus(v))
}

//SetNoNestedPartyIDs sets NoNestedPartyIDs, Tag 539
func (m NoPositions) SetNoNestedPartyIDs(f NoNestedPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetQuantityDate sets QuantityDate, Tag 976
func (m NoPositions) SetQuantityDate(v string) {
	m.Set(field.NewQuantityDate(v))
}

//GetPosType gets PosType, Tag 703
func (m NoPositions) GetPosType() (f field.PosTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLongQty gets LongQty, Tag 704
func (m NoPositions) GetLongQty() (f field.LongQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetShortQty gets ShortQty, Tag 705
func (m NoPositions) GetShortQty() (f field.ShortQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPosQtyStatus gets PosQtyStatus, Tag 706
func (m NoPositions) GetPosQtyStatus() (f field.PosQtyStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoNestedPartyIDs gets NoNestedPartyIDs, Tag 539
func (m NoPositions) GetNoNestedPartyIDs() (NoNestedPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetQuantityDate gets QuantityDate, Tag 976
func (m NoPositions) GetQuantityDate() (f field.QuantityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasPosType returns true if PosType is present, Tag 703
func (m NoPositions) HasPosType() bool {
	return m.Has(tag.PosType)
}

//HasLongQty returns true if LongQty is present, Tag 704
func (m NoPositions) HasLongQty() bool {
	return m.Has(tag.LongQty)
}

//HasShortQty returns true if ShortQty is present, Tag 705
func (m NoPositions) HasShortQty() bool {
	return m.Has(tag.ShortQty)
}

//HasPosQtyStatus returns true if PosQtyStatus is present, Tag 706
func (m NoPositions) HasPosQtyStatus() bool {
	return m.Has(tag.PosQtyStatus)
}

//HasNoNestedPartyIDs returns true if NoNestedPartyIDs is present, Tag 539
func (m NoPositions) HasNoNestedPartyIDs() bool {
	return m.Has(tag.NoNestedPartyIDs)
}

//HasQuantityDate returns true if QuantityDate is present, Tag 976
func (m NoPositions) HasQuantityDate() bool {
	return m.Has(tag.QuantityDate)
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

//NoPositionsRepeatingGroup is a repeating group, Tag 702
type NoPositionsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPositionsRepeatingGroup returns an initialized, NoPositionsRepeatingGroup
func NewNoPositionsRepeatingGroup() NoPositionsRepeatingGroup {
	return NoPositionsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPositions,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PosType), quickfix.GroupElement(tag.LongQty), quickfix.GroupElement(tag.ShortQty), quickfix.GroupElement(tag.PosQtyStatus), NewNoNestedPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.QuantityDate)})}
}

//Add create and append a new NoPositions to this group
func (m NoPositionsRepeatingGroup) Add() NoPositions {
	g := m.RepeatingGroup.Add()
	return NoPositions{g}
}

//Get returns the ith NoPositions in the NoPositionsRepeatinGroup
func (m NoPositionsRepeatingGroup) Get(i int) NoPositions {
	return NoPositions{m.RepeatingGroup.Get(i)}
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

//SetUnderlyingAllocationPercent sets UnderlyingAllocationPercent, Tag 972
func (m NoUnderlyings) SetUnderlyingAllocationPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAllocationPercent(value, scale))
}

//SetUnderlyingSettlementType sets UnderlyingSettlementType, Tag 975
func (m NoUnderlyings) SetUnderlyingSettlementType(v int) {
	m.Set(field.NewUnderlyingSettlementType(v))
}

//SetUnderlyingCashAmount sets UnderlyingCashAmount, Tag 973
func (m NoUnderlyings) SetUnderlyingCashAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCashAmount(value, scale))
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
func (m NoUnderlyings) SetUnderlyingUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingUnitOfMeasureQty(value, scale))
}

//SetUnderlyingPriceUnitOfMeasure sets UnderlyingPriceUnitOfMeasure, Tag 1424
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasure(v))
}

//SetUnderlyingPriceUnitOfMeasureQty sets UnderlyingPriceUnitOfMeasureQty, Tag 1425
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasureQty(value, scale))
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
func (m NoUnderlyings) SetUnderlyingNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingNotionalPercentageOutstanding(value, scale))
}

//SetUnderlyingOriginalNotionalPercentageOutstanding sets UnderlyingOriginalNotionalPercentageOutstanding, Tag 1456
func (m NoUnderlyings) SetUnderlyingOriginalNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingOriginalNotionalPercentageOutstanding(value, scale))
}

//SetUnderlyingAttachmentPoint sets UnderlyingAttachmentPoint, Tag 1459
func (m NoUnderlyings) SetUnderlyingAttachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAttachmentPoint(value, scale))
}

//SetUnderlyingDetachmentPoint sets UnderlyingDetachmentPoint, Tag 1460
func (m NoUnderlyings) SetUnderlyingDetachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDetachmentPoint(value, scale))
}

//SetUnderlyingSettlPrice sets UnderlyingSettlPrice, Tag 732
func (m NoUnderlyings) SetUnderlyingSettlPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingSettlPrice(value, scale))
}

//SetUnderlyingSettlPriceType sets UnderlyingSettlPriceType, Tag 733
func (m NoUnderlyings) SetUnderlyingSettlPriceType(v int) {
	m.Set(field.NewUnderlyingSettlPriceType(v))
}

//SetNoUnderlyingAmounts sets NoUnderlyingAmounts, Tag 984
func (m NoUnderlyings) SetNoUnderlyingAmounts(f NoUnderlyingAmountsRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingDeliveryAmount sets UnderlyingDeliveryAmount, Tag 1037
func (m NoUnderlyings) SetUnderlyingDeliveryAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDeliveryAmount(value, scale))
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

//GetUnderlyingSettlPrice gets UnderlyingSettlPrice, Tag 732
func (m NoUnderlyings) GetUnderlyingSettlPrice() (f field.UnderlyingSettlPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSettlPriceType gets UnderlyingSettlPriceType, Tag 733
func (m NoUnderlyings) GetUnderlyingSettlPriceType() (f field.UnderlyingSettlPriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyingAmounts gets NoUnderlyingAmounts, Tag 984
func (m NoUnderlyings) GetNoUnderlyingAmounts() (NoUnderlyingAmountsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingAmountsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingDeliveryAmount gets UnderlyingDeliveryAmount, Tag 1037
func (m NoUnderlyings) GetUnderlyingDeliveryAmount() (f field.UnderlyingDeliveryAmountField, err quickfix.MessageRejectError) {
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

//HasUnderlyingSettlPrice returns true if UnderlyingSettlPrice is present, Tag 732
func (m NoUnderlyings) HasUnderlyingSettlPrice() bool {
	return m.Has(tag.UnderlyingSettlPrice)
}

//HasUnderlyingSettlPriceType returns true if UnderlyingSettlPriceType is present, Tag 733
func (m NoUnderlyings) HasUnderlyingSettlPriceType() bool {
	return m.Has(tag.UnderlyingSettlPriceType)
}

//HasNoUnderlyingAmounts returns true if NoUnderlyingAmounts is present, Tag 984
func (m NoUnderlyings) HasNoUnderlyingAmounts() bool {
	return m.Has(tag.NoUnderlyingAmounts)
}

//HasUnderlyingDeliveryAmount returns true if UnderlyingDeliveryAmount is present, Tag 1037
func (m NoUnderlyings) HasUnderlyingDeliveryAmount() bool {
	return m.Has(tag.UnderlyingDeliveryAmount)
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingInstrumentPartyID), quickfix.GroupElement(tag.UnderlyingInstrumentPartyIDSource), quickfix.GroupElement(tag.UnderlyingInstrumentPartyRole), NewNoUndlyInstrumentPartySubIDsRepeatingGroup()})}
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

//NoUnderlyingAmounts is a repeating group element, Tag 984
type NoUnderlyingAmounts struct {
	quickfix.Group
}

//SetUnderlyingPayAmount sets UnderlyingPayAmount, Tag 985
func (m NoUnderlyingAmounts) SetUnderlyingPayAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPayAmount(value, scale))
}

//SetUnderlyingCollectAmount sets UnderlyingCollectAmount, Tag 986
func (m NoUnderlyingAmounts) SetUnderlyingCollectAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCollectAmount(value, scale))
}

//SetUnderlyingSettlementDate sets UnderlyingSettlementDate, Tag 987
func (m NoUnderlyingAmounts) SetUnderlyingSettlementDate(v string) {
	m.Set(field.NewUnderlyingSettlementDate(v))
}

//SetUnderlyingSettlementStatus sets UnderlyingSettlementStatus, Tag 988
func (m NoUnderlyingAmounts) SetUnderlyingSettlementStatus(v string) {
	m.Set(field.NewUnderlyingSettlementStatus(v))
}

//GetUnderlyingPayAmount gets UnderlyingPayAmount, Tag 985
func (m NoUnderlyingAmounts) GetUnderlyingPayAmount() (f field.UnderlyingPayAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCollectAmount gets UnderlyingCollectAmount, Tag 986
func (m NoUnderlyingAmounts) GetUnderlyingCollectAmount() (f field.UnderlyingCollectAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSettlementDate gets UnderlyingSettlementDate, Tag 987
func (m NoUnderlyingAmounts) GetUnderlyingSettlementDate() (f field.UnderlyingSettlementDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSettlementStatus gets UnderlyingSettlementStatus, Tag 988
func (m NoUnderlyingAmounts) GetUnderlyingSettlementStatus() (f field.UnderlyingSettlementStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingPayAmount returns true if UnderlyingPayAmount is present, Tag 985
func (m NoUnderlyingAmounts) HasUnderlyingPayAmount() bool {
	return m.Has(tag.UnderlyingPayAmount)
}

//HasUnderlyingCollectAmount returns true if UnderlyingCollectAmount is present, Tag 986
func (m NoUnderlyingAmounts) HasUnderlyingCollectAmount() bool {
	return m.Has(tag.UnderlyingCollectAmount)
}

//HasUnderlyingSettlementDate returns true if UnderlyingSettlementDate is present, Tag 987
func (m NoUnderlyingAmounts) HasUnderlyingSettlementDate() bool {
	return m.Has(tag.UnderlyingSettlementDate)
}

//HasUnderlyingSettlementStatus returns true if UnderlyingSettlementStatus is present, Tag 988
func (m NoUnderlyingAmounts) HasUnderlyingSettlementStatus() bool {
	return m.Has(tag.UnderlyingSettlementStatus)
}

//NoUnderlyingAmountsRepeatingGroup is a repeating group, Tag 984
type NoUnderlyingAmountsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingAmountsRepeatingGroup returns an initialized, NoUnderlyingAmountsRepeatingGroup
func NewNoUnderlyingAmountsRepeatingGroup() NoUnderlyingAmountsRepeatingGroup {
	return NoUnderlyingAmountsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingAmounts,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingPayAmount), quickfix.GroupElement(tag.UnderlyingCollectAmount), quickfix.GroupElement(tag.UnderlyingSettlementDate), quickfix.GroupElement(tag.UnderlyingSettlementStatus)})}
}

//Add create and append a new NoUnderlyingAmounts to this group
func (m NoUnderlyingAmountsRepeatingGroup) Add() NoUnderlyingAmounts {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingAmounts{g}
}

//Get returns the ith NoUnderlyingAmounts in the NoUnderlyingAmountsRepeatinGroup
func (m NoUnderlyingAmountsRepeatingGroup) Get(i int) NoUnderlyingAmounts {
	return NoUnderlyingAmounts{m.RepeatingGroup.Get(i)}
}

//NoUnderlyingsRepeatingGroup is a repeating group, Tag 711
type NoUnderlyingsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingsRepeatingGroup returns an initialized, NoUnderlyingsRepeatingGroup
func NewNoUnderlyingsRepeatingGroup() NoUnderlyingsRepeatingGroup {
	return NoUnderlyingsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyings,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingSecurityIDSource), NewNoUnderlyingSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingProduct), quickfix.GroupElement(tag.UnderlyingCFICode), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingSecuritySubType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDate), quickfix.GroupElement(tag.UnderlyingCouponPaymentDate), quickfix.GroupElement(tag.UnderlyingIssueDate), quickfix.GroupElement(tag.UnderlyingRepoCollateralSecurityType), quickfix.GroupElement(tag.UnderlyingRepurchaseTerm), quickfix.GroupElement(tag.UnderlyingRepurchaseRate), quickfix.GroupElement(tag.UnderlyingFactor), quickfix.GroupElement(tag.UnderlyingCreditRating), quickfix.GroupElement(tag.UnderlyingInstrRegistry), quickfix.GroupElement(tag.UnderlyingCountryOfIssue), quickfix.GroupElement(tag.UnderlyingStateOrProvinceOfIssue), quickfix.GroupElement(tag.UnderlyingLocaleOfIssue), quickfix.GroupElement(tag.UnderlyingRedemptionDate), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingStrikeCurrency), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.UnderlyingCPProgram), quickfix.GroupElement(tag.UnderlyingCPRegType), quickfix.GroupElement(tag.UnderlyingCurrency), quickfix.GroupElement(tag.UnderlyingQty), quickfix.GroupElement(tag.UnderlyingPx), quickfix.GroupElement(tag.UnderlyingDirtyPrice), quickfix.GroupElement(tag.UnderlyingEndPrice), quickfix.GroupElement(tag.UnderlyingStartValue), quickfix.GroupElement(tag.UnderlyingCurrentValue), quickfix.GroupElement(tag.UnderlyingEndValue), NewNoUnderlyingStipsRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingAllocationPercent), quickfix.GroupElement(tag.UnderlyingSettlementType), quickfix.GroupElement(tag.UnderlyingCashAmount), quickfix.GroupElement(tag.UnderlyingCashType), quickfix.GroupElement(tag.UnderlyingUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingTimeUnit), quickfix.GroupElement(tag.UnderlyingCapValue), NewNoUndlyInstrumentPartiesRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingSettlMethod), quickfix.GroupElement(tag.UnderlyingAdjustedQuantity), quickfix.GroupElement(tag.UnderlyingFXRate), quickfix.GroupElement(tag.UnderlyingFXRateCalc), quickfix.GroupElement(tag.UnderlyingMaturityTime), quickfix.GroupElement(tag.UnderlyingPutOrCall), quickfix.GroupElement(tag.UnderlyingExerciseStyle), quickfix.GroupElement(tag.UnderlyingUnitOfMeasureQty), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasureQty), quickfix.GroupElement(tag.UnderlyingContractMultiplierUnit), quickfix.GroupElement(tag.UnderlyingFlowScheduleType), quickfix.GroupElement(tag.UnderlyingRestructuringType), quickfix.GroupElement(tag.UnderlyingSeniority), quickfix.GroupElement(tag.UnderlyingNotionalPercentageOutstanding), quickfix.GroupElement(tag.UnderlyingOriginalNotionalPercentageOutstanding), quickfix.GroupElement(tag.UnderlyingAttachmentPoint), quickfix.GroupElement(tag.UnderlyingDetachmentPoint), quickfix.GroupElement(tag.UnderlyingSettlPrice), quickfix.GroupElement(tag.UnderlyingSettlPriceType), NewNoUnderlyingAmountsRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingDeliveryAmount)})}
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
func (m NoPosAmt) SetPosAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewPosAmt(value, scale))
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

//NoComplexEvents is a repeating group element, Tag 1483
type NoComplexEvents struct {
	quickfix.Group
}

//SetComplexEventType sets ComplexEventType, Tag 1484
func (m NoComplexEvents) SetComplexEventType(v int) {
	m.Set(field.NewComplexEventType(v))
}

//SetComplexOptPayoutAmount sets ComplexOptPayoutAmount, Tag 1485
func (m NoComplexEvents) SetComplexOptPayoutAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexOptPayoutAmount(value, scale))
}

//SetComplexEventPrice sets ComplexEventPrice, Tag 1486
func (m NoComplexEvents) SetComplexEventPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexEventPrice(value, scale))
}

//SetComplexEventPriceBoundaryMethod sets ComplexEventPriceBoundaryMethod, Tag 1487
func (m NoComplexEvents) SetComplexEventPriceBoundaryMethod(v int) {
	m.Set(field.NewComplexEventPriceBoundaryMethod(v))
}

//SetComplexEventPriceBoundaryPrecision sets ComplexEventPriceBoundaryPrecision, Tag 1488
func (m NoComplexEvents) SetComplexEventPriceBoundaryPrecision(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexEventPriceBoundaryPrecision(value, scale))
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventStartDate), quickfix.GroupElement(tag.ComplexEventEndDate), NewNoComplexEventTimesRepeatingGroup()})}
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventType), quickfix.GroupElement(tag.ComplexOptPayoutAmount), quickfix.GroupElement(tag.ComplexEventPrice), quickfix.GroupElement(tag.ComplexEventPriceBoundaryMethod), quickfix.GroupElement(tag.ComplexEventPriceBoundaryPrecision), quickfix.GroupElement(tag.ComplexEventPriceTimeType), quickfix.GroupElement(tag.ComplexEventCondition), NewNoComplexEventDatesRepeatingGroup()})}
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
