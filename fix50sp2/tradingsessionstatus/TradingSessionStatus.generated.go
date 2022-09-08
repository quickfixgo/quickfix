package tradingsessionstatus

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/alpacahq/quickfix"
	"github.com/alpacahq/quickfix/enum"
	"github.com/alpacahq/quickfix/field"
	"github.com/alpacahq/quickfix/fixt11"
	"github.com/alpacahq/quickfix/tag"
)

// TradingSessionStatus is the fix50sp2 TradingSessionStatus type, MsgType = h
type TradingSessionStatus struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a TradingSessionStatus from a quickfix.Message instance
func FromMessage(m *quickfix.Message) TradingSessionStatus {
	return TradingSessionStatus{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance
func (m TradingSessionStatus) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a TradingSessionStatus initialized with the required fields for TradingSessionStatus
func New(tradingsessionid field.TradingSessionIDField, tradsesstatus field.TradSesStatusField) (m TradingSessionStatus) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("h"))
	m.Set(tradingsessionid)
	m.Set(tradsesstatus)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg TradingSessionStatus, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "h", r
}

// SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m TradingSessionStatus) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

// SetSecurityID sets SecurityID, Tag 48
func (m TradingSessionStatus) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetSymbol sets Symbol, Tag 55
func (m TradingSessionStatus) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetText sets Text, Tag 58
func (m TradingSessionStatus) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetSymbolSfx sets SymbolSfx, Tag 65
func (m TradingSessionStatus) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

// SetIssuer sets Issuer, Tag 106
func (m TradingSessionStatus) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

// SetSecurityDesc sets SecurityDesc, Tag 107
func (m TradingSessionStatus) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

// SetSecurityType sets SecurityType, Tag 167
func (m TradingSessionStatus) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

// SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m TradingSessionStatus) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

// SetPutOrCall sets PutOrCall, Tag 201
func (m TradingSessionStatus) SetPutOrCall(v enum.PutOrCall) {
	m.Set(field.NewPutOrCall(v))
}

// SetStrikePrice sets StrikePrice, Tag 202
func (m TradingSessionStatus) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

// SetOptAttribute sets OptAttribute, Tag 206
func (m TradingSessionStatus) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

// SetSecurityExchange sets SecurityExchange, Tag 207
func (m TradingSessionStatus) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

// SetCouponRate sets CouponRate, Tag 223
func (m TradingSessionStatus) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

// SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m TradingSessionStatus) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

// SetIssueDate sets IssueDate, Tag 225
func (m TradingSessionStatus) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

// SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m TradingSessionStatus) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

// SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m TradingSessionStatus) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

// SetFactor sets Factor, Tag 228
func (m TradingSessionStatus) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

// SetContractMultiplier sets ContractMultiplier, Tag 231
func (m TradingSessionStatus) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

// SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m TradingSessionStatus) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

// SetRedemptionDate sets RedemptionDate, Tag 240
func (m TradingSessionStatus) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

// SetCreditRating sets CreditRating, Tag 255
func (m TradingSessionStatus) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

// SetUnsolicitedIndicator sets UnsolicitedIndicator, Tag 325
func (m TradingSessionStatus) SetUnsolicitedIndicator(v bool) {
	m.Set(field.NewUnsolicitedIndicator(v))
}

// SetTradSesReqID sets TradSesReqID, Tag 335
func (m TradingSessionStatus) SetTradSesReqID(v string) {
	m.Set(field.NewTradSesReqID(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336
func (m TradingSessionStatus) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetTradSesMethod sets TradSesMethod, Tag 338
func (m TradingSessionStatus) SetTradSesMethod(v enum.TradSesMethod) {
	m.Set(field.NewTradSesMethod(v))
}

// SetTradSesMode sets TradSesMode, Tag 339
func (m TradingSessionStatus) SetTradSesMode(v enum.TradSesMode) {
	m.Set(field.NewTradSesMode(v))
}

// SetTradSesStatus sets TradSesStatus, Tag 340
func (m TradingSessionStatus) SetTradSesStatus(v enum.TradSesStatus) {
	m.Set(field.NewTradSesStatus(v))
}

// SetTradSesStartTime sets TradSesStartTime, Tag 341
func (m TradingSessionStatus) SetTradSesStartTime(v time.Time) {
	m.Set(field.NewTradSesStartTime(v))
}

// SetTradSesOpenTime sets TradSesOpenTime, Tag 342
func (m TradingSessionStatus) SetTradSesOpenTime(v time.Time) {
	m.Set(field.NewTradSesOpenTime(v))
}

// SetTradSesPreCloseTime sets TradSesPreCloseTime, Tag 343
func (m TradingSessionStatus) SetTradSesPreCloseTime(v time.Time) {
	m.Set(field.NewTradSesPreCloseTime(v))
}

// SetTradSesCloseTime sets TradSesCloseTime, Tag 344
func (m TradingSessionStatus) SetTradSesCloseTime(v time.Time) {
	m.Set(field.NewTradSesCloseTime(v))
}

// SetTradSesEndTime sets TradSesEndTime, Tag 345
func (m TradingSessionStatus) SetTradSesEndTime(v time.Time) {
	m.Set(field.NewTradSesEndTime(v))
}

// SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m TradingSessionStatus) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

// SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m TradingSessionStatus) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

// SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m TradingSessionStatus) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

// SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m TradingSessionStatus) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m TradingSessionStatus) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355
func (m TradingSessionStatus) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetTotalVolumeTraded sets TotalVolumeTraded, Tag 387
func (m TradingSessionStatus) SetTotalVolumeTraded(value decimal.Decimal, scale int32) {
	m.Set(field.NewTotalVolumeTraded(value, scale))
}

// SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m TradingSessionStatus) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetProduct sets Product, Tag 460
func (m TradingSessionStatus) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

// SetCFICode sets CFICode, Tag 461
func (m TradingSessionStatus) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

// SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m TradingSessionStatus) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

// SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m TradingSessionStatus) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

// SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m TradingSessionStatus) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

// SetMaturityDate sets MaturityDate, Tag 541
func (m TradingSessionStatus) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

// SetInstrRegistry sets InstrRegistry, Tag 543
func (m TradingSessionStatus) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

// SetTradSesStatusRejReason sets TradSesStatusRejReason, Tag 567
func (m TradingSessionStatus) SetTradSesStatusRejReason(v enum.TradSesStatusRejReason) {
	m.Set(field.NewTradSesStatusRejReason(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m TradingSessionStatus) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m TradingSessionStatus) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

// SetPool sets Pool, Tag 691
func (m TradingSessionStatus) SetPool(v string) {
	m.Set(field.NewPool(v))
}

// SetSecuritySubType sets SecuritySubType, Tag 762
func (m TradingSessionStatus) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

// SetNoEvents sets NoEvents, Tag 864
func (m TradingSessionStatus) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

// SetDatedDate sets DatedDate, Tag 873
func (m TradingSessionStatus) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

// SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m TradingSessionStatus) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

// SetCPProgram sets CPProgram, Tag 875
func (m TradingSessionStatus) SetCPProgram(v enum.CPProgram) {
	m.Set(field.NewCPProgram(v))
}

// SetCPRegType sets CPRegType, Tag 876
func (m TradingSessionStatus) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

// SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m TradingSessionStatus) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

// SetSecurityStatus sets SecurityStatus, Tag 965
func (m TradingSessionStatus) SetSecurityStatus(v enum.SecurityStatus) {
	m.Set(field.NewSecurityStatus(v))
}

// SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966
func (m TradingSessionStatus) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

// SetStrikeMultiplier sets StrikeMultiplier, Tag 967
func (m TradingSessionStatus) SetStrikeMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeMultiplier(value, scale))
}

// SetStrikeValue sets StrikeValue, Tag 968
func (m TradingSessionStatus) SetStrikeValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeValue(value, scale))
}

// SetMinPriceIncrement sets MinPriceIncrement, Tag 969
func (m TradingSessionStatus) SetMinPriceIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrement(value, scale))
}

// SetPositionLimit sets PositionLimit, Tag 970
func (m TradingSessionStatus) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

// SetNTPositionLimit sets NTPositionLimit, Tag 971
func (m TradingSessionStatus) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

// SetUnitOfMeasure sets UnitOfMeasure, Tag 996
func (m TradingSessionStatus) SetUnitOfMeasure(v enum.UnitOfMeasure) {
	m.Set(field.NewUnitOfMeasure(v))
}

// SetTimeUnit sets TimeUnit, Tag 997
func (m TradingSessionStatus) SetTimeUnit(v enum.TimeUnit) {
	m.Set(field.NewTimeUnit(v))
}

// SetNoInstrumentParties sets NoInstrumentParties, Tag 1018
func (m TradingSessionStatus) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

// SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049
func (m TradingSessionStatus) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

// SetMaturityTime sets MaturityTime, Tag 1079
func (m TradingSessionStatus) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

// SetMinPriceIncrementAmount sets MinPriceIncrementAmount, Tag 1146
func (m TradingSessionStatus) SetMinPriceIncrementAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrementAmount(value, scale))
}

// SetUnitOfMeasureQty sets UnitOfMeasureQty, Tag 1147
func (m TradingSessionStatus) SetUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnitOfMeasureQty(value, scale))
}

// SetSecurityGroup sets SecurityGroup, Tag 1151
func (m TradingSessionStatus) SetSecurityGroup(v string) {
	m.Set(field.NewSecurityGroup(v))
}

// SetApplID sets ApplID, Tag 1180
func (m TradingSessionStatus) SetApplID(v string) {
	m.Set(field.NewApplID(v))
}

// SetApplSeqNum sets ApplSeqNum, Tag 1181
func (m TradingSessionStatus) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

// SetSecurityXMLLen sets SecurityXMLLen, Tag 1184
func (m TradingSessionStatus) SetSecurityXMLLen(v int) {
	m.Set(field.NewSecurityXMLLen(v))
}

// SetSecurityXML sets SecurityXML, Tag 1185
func (m TradingSessionStatus) SetSecurityXML(v string) {
	m.Set(field.NewSecurityXML(v))
}

// SetSecurityXMLSchema sets SecurityXMLSchema, Tag 1186
func (m TradingSessionStatus) SetSecurityXMLSchema(v string) {
	m.Set(field.NewSecurityXMLSchema(v))
}

// SetPriceUnitOfMeasure sets PriceUnitOfMeasure, Tag 1191
func (m TradingSessionStatus) SetPriceUnitOfMeasure(v string) {
	m.Set(field.NewPriceUnitOfMeasure(v))
}

// SetPriceUnitOfMeasureQty sets PriceUnitOfMeasureQty, Tag 1192
func (m TradingSessionStatus) SetPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewPriceUnitOfMeasureQty(value, scale))
}

// SetSettlMethod sets SettlMethod, Tag 1193
func (m TradingSessionStatus) SetSettlMethod(v enum.SettlMethod) {
	m.Set(field.NewSettlMethod(v))
}

// SetExerciseStyle sets ExerciseStyle, Tag 1194
func (m TradingSessionStatus) SetExerciseStyle(v enum.ExerciseStyle) {
	m.Set(field.NewExerciseStyle(v))
}

// SetOptPayoutAmount sets OptPayoutAmount, Tag 1195
func (m TradingSessionStatus) SetOptPayoutAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewOptPayoutAmount(value, scale))
}

// SetPriceQuoteMethod sets PriceQuoteMethod, Tag 1196
func (m TradingSessionStatus) SetPriceQuoteMethod(v enum.PriceQuoteMethod) {
	m.Set(field.NewPriceQuoteMethod(v))
}

// SetValuationMethod sets ValuationMethod, Tag 1197
func (m TradingSessionStatus) SetValuationMethod(v enum.ValuationMethod) {
	m.Set(field.NewValuationMethod(v))
}

// SetListMethod sets ListMethod, Tag 1198
func (m TradingSessionStatus) SetListMethod(v enum.ListMethod) {
	m.Set(field.NewListMethod(v))
}

// SetCapPrice sets CapPrice, Tag 1199
func (m TradingSessionStatus) SetCapPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewCapPrice(value, scale))
}

// SetFloorPrice sets FloorPrice, Tag 1200
func (m TradingSessionStatus) SetFloorPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewFloorPrice(value, scale))
}

// SetProductComplex sets ProductComplex, Tag 1227
func (m TradingSessionStatus) SetProductComplex(v string) {
	m.Set(field.NewProductComplex(v))
}

// SetFlexProductEligibilityIndicator sets FlexProductEligibilityIndicator, Tag 1242
func (m TradingSessionStatus) SetFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewFlexProductEligibilityIndicator(v))
}

// SetFlexibleIndicator sets FlexibleIndicator, Tag 1244
func (m TradingSessionStatus) SetFlexibleIndicator(v bool) {
	m.Set(field.NewFlexibleIndicator(v))
}

// SetMarketSegmentID sets MarketSegmentID, Tag 1300
func (m TradingSessionStatus) SetMarketSegmentID(v string) {
	m.Set(field.NewMarketSegmentID(v))
}

// SetMarketID sets MarketID, Tag 1301
func (m TradingSessionStatus) SetMarketID(v string) {
	m.Set(field.NewMarketID(v))
}

// SetApplLastSeqNum sets ApplLastSeqNum, Tag 1350
func (m TradingSessionStatus) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

// SetApplResendFlag sets ApplResendFlag, Tag 1352
func (m TradingSessionStatus) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

// SetTradSesEvent sets TradSesEvent, Tag 1368
func (m TradingSessionStatus) SetTradSesEvent(v enum.TradSesEvent) {
	m.Set(field.NewTradSesEvent(v))
}

// SetContractMultiplierUnit sets ContractMultiplierUnit, Tag 1435
func (m TradingSessionStatus) SetContractMultiplierUnit(v enum.ContractMultiplierUnit) {
	m.Set(field.NewContractMultiplierUnit(v))
}

// SetFlowScheduleType sets FlowScheduleType, Tag 1439
func (m TradingSessionStatus) SetFlowScheduleType(v enum.FlowScheduleType) {
	m.Set(field.NewFlowScheduleType(v))
}

// SetRestructuringType sets RestructuringType, Tag 1449
func (m TradingSessionStatus) SetRestructuringType(v enum.RestructuringType) {
	m.Set(field.NewRestructuringType(v))
}

// SetSeniority sets Seniority, Tag 1450
func (m TradingSessionStatus) SetSeniority(v enum.Seniority) {
	m.Set(field.NewSeniority(v))
}

// SetNotionalPercentageOutstanding sets NotionalPercentageOutstanding, Tag 1451
func (m TradingSessionStatus) SetNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewNotionalPercentageOutstanding(value, scale))
}

// SetOriginalNotionalPercentageOutstanding sets OriginalNotionalPercentageOutstanding, Tag 1452
func (m TradingSessionStatus) SetOriginalNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewOriginalNotionalPercentageOutstanding(value, scale))
}

// SetAttachmentPoint sets AttachmentPoint, Tag 1457
func (m TradingSessionStatus) SetAttachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewAttachmentPoint(value, scale))
}

// SetDetachmentPoint sets DetachmentPoint, Tag 1458
func (m TradingSessionStatus) SetDetachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewDetachmentPoint(value, scale))
}

// SetStrikePriceDeterminationMethod sets StrikePriceDeterminationMethod, Tag 1478
func (m TradingSessionStatus) SetStrikePriceDeterminationMethod(v enum.StrikePriceDeterminationMethod) {
	m.Set(field.NewStrikePriceDeterminationMethod(v))
}

// SetStrikePriceBoundaryMethod sets StrikePriceBoundaryMethod, Tag 1479
func (m TradingSessionStatus) SetStrikePriceBoundaryMethod(v enum.StrikePriceBoundaryMethod) {
	m.Set(field.NewStrikePriceBoundaryMethod(v))
}

// SetStrikePriceBoundaryPrecision sets StrikePriceBoundaryPrecision, Tag 1480
func (m TradingSessionStatus) SetStrikePriceBoundaryPrecision(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePriceBoundaryPrecision(value, scale))
}

// SetUnderlyingPriceDeterminationMethod sets UnderlyingPriceDeterminationMethod, Tag 1481
func (m TradingSessionStatus) SetUnderlyingPriceDeterminationMethod(v enum.UnderlyingPriceDeterminationMethod) {
	m.Set(field.NewUnderlyingPriceDeterminationMethod(v))
}

// SetOptPayoutType sets OptPayoutType, Tag 1482
func (m TradingSessionStatus) SetOptPayoutType(v enum.OptPayoutType) {
	m.Set(field.NewOptPayoutType(v))
}

// SetNoComplexEvents sets NoComplexEvents, Tag 1483
func (m TradingSessionStatus) SetNoComplexEvents(f NoComplexEventsRepeatingGroup) {
	m.SetGroup(f)
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m TradingSessionStatus) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48
func (m TradingSessionStatus) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55
func (m TradingSessionStatus) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetText gets Text, Tag 58
func (m TradingSessionStatus) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbolSfx gets SymbolSfx, Tag 65
func (m TradingSessionStatus) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssuer gets Issuer, Tag 106
func (m TradingSessionStatus) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityDesc gets SecurityDesc, Tag 107
func (m TradingSessionStatus) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityType gets SecurityType, Tag 167
func (m TradingSessionStatus) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m TradingSessionStatus) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPutOrCall gets PutOrCall, Tag 201
func (m TradingSessionStatus) GetPutOrCall() (v enum.PutOrCall, err quickfix.MessageRejectError) {
	var f field.PutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePrice gets StrikePrice, Tag 202
func (m TradingSessionStatus) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptAttribute gets OptAttribute, Tag 206
func (m TradingSessionStatus) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityExchange gets SecurityExchange, Tag 207
func (m TradingSessionStatus) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponRate gets CouponRate, Tag 223
func (m TradingSessionStatus) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m TradingSessionStatus) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssueDate gets IssueDate, Tag 225
func (m TradingSessionStatus) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m TradingSessionStatus) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m TradingSessionStatus) GetRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFactor gets Factor, Tag 228
func (m TradingSessionStatus) GetFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractMultiplier gets ContractMultiplier, Tag 231
func (m TradingSessionStatus) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m TradingSessionStatus) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRedemptionDate gets RedemptionDate, Tag 240
func (m TradingSessionStatus) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCreditRating gets CreditRating, Tag 255
func (m TradingSessionStatus) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnsolicitedIndicator gets UnsolicitedIndicator, Tag 325
func (m TradingSessionStatus) GetUnsolicitedIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.UnsolicitedIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesReqID gets TradSesReqID, Tag 335
func (m TradingSessionStatus) GetTradSesReqID() (v string, err quickfix.MessageRejectError) {
	var f field.TradSesReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336
func (m TradingSessionStatus) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesMethod gets TradSesMethod, Tag 338
func (m TradingSessionStatus) GetTradSesMethod() (v enum.TradSesMethod, err quickfix.MessageRejectError) {
	var f field.TradSesMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesMode gets TradSesMode, Tag 339
func (m TradingSessionStatus) GetTradSesMode() (v enum.TradSesMode, err quickfix.MessageRejectError) {
	var f field.TradSesModeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesStatus gets TradSesStatus, Tag 340
func (m TradingSessionStatus) GetTradSesStatus() (v enum.TradSesStatus, err quickfix.MessageRejectError) {
	var f field.TradSesStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesStartTime gets TradSesStartTime, Tag 341
func (m TradingSessionStatus) GetTradSesStartTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesStartTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesOpenTime gets TradSesOpenTime, Tag 342
func (m TradingSessionStatus) GetTradSesOpenTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesOpenTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesPreCloseTime gets TradSesPreCloseTime, Tag 343
func (m TradingSessionStatus) GetTradSesPreCloseTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesPreCloseTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesCloseTime gets TradSesCloseTime, Tag 344
func (m TradingSessionStatus) GetTradSesCloseTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesCloseTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesEndTime gets TradSesEndTime, Tag 345
func (m TradingSessionStatus) GetTradSesEndTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesEndTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m TradingSessionStatus) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m TradingSessionStatus) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m TradingSessionStatus) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m TradingSessionStatus) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m TradingSessionStatus) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355
func (m TradingSessionStatus) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotalVolumeTraded gets TotalVolumeTraded, Tag 387
func (m TradingSessionStatus) GetTotalVolumeTraded() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TotalVolumeTradedField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m TradingSessionStatus) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetProduct gets Product, Tag 460
func (m TradingSessionStatus) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCFICode gets CFICode, Tag 461
func (m TradingSessionStatus) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m TradingSessionStatus) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m TradingSessionStatus) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m TradingSessionStatus) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityDate gets MaturityDate, Tag 541
func (m TradingSessionStatus) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrRegistry gets InstrRegistry, Tag 543
func (m TradingSessionStatus) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesStatusRejReason gets TradSesStatusRejReason, Tag 567
func (m TradingSessionStatus) GetTradSesStatusRejReason() (v enum.TradSesStatusRejReason, err quickfix.MessageRejectError) {
	var f field.TradSesStatusRejReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m TradingSessionStatus) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m TradingSessionStatus) GetContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.ContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPool gets Pool, Tag 691
func (m TradingSessionStatus) GetPool() (v string, err quickfix.MessageRejectError) {
	var f field.PoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecuritySubType gets SecuritySubType, Tag 762
func (m TradingSessionStatus) GetSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoEvents gets NoEvents, Tag 864
func (m TradingSessionStatus) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetDatedDate gets DatedDate, Tag 873
func (m TradingSessionStatus) GetDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.DatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m TradingSessionStatus) GetInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.InterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPProgram gets CPProgram, Tag 875
func (m TradingSessionStatus) GetCPProgram() (v enum.CPProgram, err quickfix.MessageRejectError) {
	var f field.CPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPRegType gets CPRegType, Tag 876
func (m TradingSessionStatus) GetCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.CPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m TradingSessionStatus) GetStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.StrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityStatus gets SecurityStatus, Tag 965
func (m TradingSessionStatus) GetSecurityStatus() (v enum.SecurityStatus, err quickfix.MessageRejectError) {
	var f field.SecurityStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966
func (m TradingSessionStatus) GetSettleOnOpenFlag() (v string, err quickfix.MessageRejectError) {
	var f field.SettleOnOpenFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeMultiplier gets StrikeMultiplier, Tag 967
func (m TradingSessionStatus) GetStrikeMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikeMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeValue gets StrikeValue, Tag 968
func (m TradingSessionStatus) GetStrikeValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikeValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMinPriceIncrement gets MinPriceIncrement, Tag 969
func (m TradingSessionStatus) GetMinPriceIncrement() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinPriceIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPositionLimit gets PositionLimit, Tag 970
func (m TradingSessionStatus) GetPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.PositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNTPositionLimit gets NTPositionLimit, Tag 971
func (m TradingSessionStatus) GetNTPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.NTPositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnitOfMeasure gets UnitOfMeasure, Tag 996
func (m TradingSessionStatus) GetUnitOfMeasure() (v enum.UnitOfMeasure, err quickfix.MessageRejectError) {
	var f field.UnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTimeUnit gets TimeUnit, Tag 997
func (m TradingSessionStatus) GetTimeUnit() (v enum.TimeUnit, err quickfix.MessageRejectError) {
	var f field.TimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoInstrumentParties gets NoInstrumentParties, Tag 1018
func (m TradingSessionStatus) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049
func (m TradingSessionStatus) GetInstrmtAssignmentMethod() (v string, err quickfix.MessageRejectError) {
	var f field.InstrmtAssignmentMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityTime gets MaturityTime, Tag 1079
func (m TradingSessionStatus) GetMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMinPriceIncrementAmount gets MinPriceIncrementAmount, Tag 1146
func (m TradingSessionStatus) GetMinPriceIncrementAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinPriceIncrementAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnitOfMeasureQty gets UnitOfMeasureQty, Tag 1147
func (m TradingSessionStatus) GetUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityGroup gets SecurityGroup, Tag 1151
func (m TradingSessionStatus) GetSecurityGroup() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityGroupField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplID gets ApplID, Tag 1180
func (m TradingSessionStatus) GetApplID() (v string, err quickfix.MessageRejectError) {
	var f field.ApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplSeqNum gets ApplSeqNum, Tag 1181
func (m TradingSessionStatus) GetApplSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityXMLLen gets SecurityXMLLen, Tag 1184
func (m TradingSessionStatus) GetSecurityXMLLen() (v int, err quickfix.MessageRejectError) {
	var f field.SecurityXMLLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityXML gets SecurityXML, Tag 1185
func (m TradingSessionStatus) GetSecurityXML() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityXMLField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityXMLSchema gets SecurityXMLSchema, Tag 1186
func (m TradingSessionStatus) GetSecurityXMLSchema() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityXMLSchemaField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceUnitOfMeasure gets PriceUnitOfMeasure, Tag 1191
func (m TradingSessionStatus) GetPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.PriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceUnitOfMeasureQty gets PriceUnitOfMeasureQty, Tag 1192
func (m TradingSessionStatus) GetPriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlMethod gets SettlMethod, Tag 1193
func (m TradingSessionStatus) GetSettlMethod() (v enum.SettlMethod, err quickfix.MessageRejectError) {
	var f field.SettlMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExerciseStyle gets ExerciseStyle, Tag 1194
func (m TradingSessionStatus) GetExerciseStyle() (v enum.ExerciseStyle, err quickfix.MessageRejectError) {
	var f field.ExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptPayoutAmount gets OptPayoutAmount, Tag 1195
func (m TradingSessionStatus) GetOptPayoutAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OptPayoutAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceQuoteMethod gets PriceQuoteMethod, Tag 1196
func (m TradingSessionStatus) GetPriceQuoteMethod() (v enum.PriceQuoteMethod, err quickfix.MessageRejectError) {
	var f field.PriceQuoteMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetValuationMethod gets ValuationMethod, Tag 1197
func (m TradingSessionStatus) GetValuationMethod() (v enum.ValuationMethod, err quickfix.MessageRejectError) {
	var f field.ValuationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetListMethod gets ListMethod, Tag 1198
func (m TradingSessionStatus) GetListMethod() (v enum.ListMethod, err quickfix.MessageRejectError) {
	var f field.ListMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCapPrice gets CapPrice, Tag 1199
func (m TradingSessionStatus) GetCapPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CapPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFloorPrice gets FloorPrice, Tag 1200
func (m TradingSessionStatus) GetFloorPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FloorPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetProductComplex gets ProductComplex, Tag 1227
func (m TradingSessionStatus) GetProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.ProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFlexProductEligibilityIndicator gets FlexProductEligibilityIndicator, Tag 1242
func (m TradingSessionStatus) GetFlexProductEligibilityIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.FlexProductEligibilityIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFlexibleIndicator gets FlexibleIndicator, Tag 1244
func (m TradingSessionStatus) GetFlexibleIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.FlexibleIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketSegmentID gets MarketSegmentID, Tag 1300
func (m TradingSessionStatus) GetMarketSegmentID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketSegmentIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketID gets MarketID, Tag 1301
func (m TradingSessionStatus) GetMarketID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplLastSeqNum gets ApplLastSeqNum, Tag 1350
func (m TradingSessionStatus) GetApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplResendFlag gets ApplResendFlag, Tag 1352
func (m TradingSessionStatus) GetApplResendFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.ApplResendFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesEvent gets TradSesEvent, Tag 1368
func (m TradingSessionStatus) GetTradSesEvent() (v enum.TradSesEvent, err quickfix.MessageRejectError) {
	var f field.TradSesEventField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractMultiplierUnit gets ContractMultiplierUnit, Tag 1435
func (m TradingSessionStatus) GetContractMultiplierUnit() (v enum.ContractMultiplierUnit, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFlowScheduleType gets FlowScheduleType, Tag 1439
func (m TradingSessionStatus) GetFlowScheduleType() (v enum.FlowScheduleType, err quickfix.MessageRejectError) {
	var f field.FlowScheduleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRestructuringType gets RestructuringType, Tag 1449
func (m TradingSessionStatus) GetRestructuringType() (v enum.RestructuringType, err quickfix.MessageRejectError) {
	var f field.RestructuringTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSeniority gets Seniority, Tag 1450
func (m TradingSessionStatus) GetSeniority() (v enum.Seniority, err quickfix.MessageRejectError) {
	var f field.SeniorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNotionalPercentageOutstanding gets NotionalPercentageOutstanding, Tag 1451
func (m TradingSessionStatus) GetNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.NotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOriginalNotionalPercentageOutstanding gets OriginalNotionalPercentageOutstanding, Tag 1452
func (m TradingSessionStatus) GetOriginalNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OriginalNotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAttachmentPoint gets AttachmentPoint, Tag 1457
func (m TradingSessionStatus) GetAttachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AttachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDetachmentPoint gets DetachmentPoint, Tag 1458
func (m TradingSessionStatus) GetDetachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DetachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePriceDeterminationMethod gets StrikePriceDeterminationMethod, Tag 1478
func (m TradingSessionStatus) GetStrikePriceDeterminationMethod() (v enum.StrikePriceDeterminationMethod, err quickfix.MessageRejectError) {
	var f field.StrikePriceDeterminationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePriceBoundaryMethod gets StrikePriceBoundaryMethod, Tag 1479
func (m TradingSessionStatus) GetStrikePriceBoundaryMethod() (v enum.StrikePriceBoundaryMethod, err quickfix.MessageRejectError) {
	var f field.StrikePriceBoundaryMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePriceBoundaryPrecision gets StrikePriceBoundaryPrecision, Tag 1480
func (m TradingSessionStatus) GetStrikePriceBoundaryPrecision() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceBoundaryPrecisionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPriceDeterminationMethod gets UnderlyingPriceDeterminationMethod, Tag 1481
func (m TradingSessionStatus) GetUnderlyingPriceDeterminationMethod() (v enum.UnderlyingPriceDeterminationMethod, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceDeterminationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptPayoutType gets OptPayoutType, Tag 1482
func (m TradingSessionStatus) GetOptPayoutType() (v enum.OptPayoutType, err quickfix.MessageRejectError) {
	var f field.OptPayoutTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoComplexEvents gets NoComplexEvents, Tag 1483
func (m TradingSessionStatus) GetNoComplexEvents() (NoComplexEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m TradingSessionStatus) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasSecurityID returns true if SecurityID is present, Tag 48
func (m TradingSessionStatus) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSymbol returns true if Symbol is present, Tag 55
func (m TradingSessionStatus) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasText returns true if Text is present, Tag 58
func (m TradingSessionStatus) HasText() bool {
	return m.Has(tag.Text)
}

// HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m TradingSessionStatus) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

// HasIssuer returns true if Issuer is present, Tag 106
func (m TradingSessionStatus) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

// HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m TradingSessionStatus) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

// HasSecurityType returns true if SecurityType is present, Tag 167
func (m TradingSessionStatus) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

// HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m TradingSessionStatus) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

// HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m TradingSessionStatus) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

// HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m TradingSessionStatus) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

// HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m TradingSessionStatus) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

// HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m TradingSessionStatus) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

// HasCouponRate returns true if CouponRate is present, Tag 223
func (m TradingSessionStatus) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

// HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m TradingSessionStatus) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

// HasIssueDate returns true if IssueDate is present, Tag 225
func (m TradingSessionStatus) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

// HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m TradingSessionStatus) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

// HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m TradingSessionStatus) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

// HasFactor returns true if Factor is present, Tag 228
func (m TradingSessionStatus) HasFactor() bool {
	return m.Has(tag.Factor)
}

// HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m TradingSessionStatus) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

// HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m TradingSessionStatus) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

// HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m TradingSessionStatus) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

// HasCreditRating returns true if CreditRating is present, Tag 255
func (m TradingSessionStatus) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

// HasUnsolicitedIndicator returns true if UnsolicitedIndicator is present, Tag 325
func (m TradingSessionStatus) HasUnsolicitedIndicator() bool {
	return m.Has(tag.UnsolicitedIndicator)
}

// HasTradSesReqID returns true if TradSesReqID is present, Tag 335
func (m TradingSessionStatus) HasTradSesReqID() bool {
	return m.Has(tag.TradSesReqID)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m TradingSessionStatus) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasTradSesMethod returns true if TradSesMethod is present, Tag 338
func (m TradingSessionStatus) HasTradSesMethod() bool {
	return m.Has(tag.TradSesMethod)
}

// HasTradSesMode returns true if TradSesMode is present, Tag 339
func (m TradingSessionStatus) HasTradSesMode() bool {
	return m.Has(tag.TradSesMode)
}

// HasTradSesStatus returns true if TradSesStatus is present, Tag 340
func (m TradingSessionStatus) HasTradSesStatus() bool {
	return m.Has(tag.TradSesStatus)
}

// HasTradSesStartTime returns true if TradSesStartTime is present, Tag 341
func (m TradingSessionStatus) HasTradSesStartTime() bool {
	return m.Has(tag.TradSesStartTime)
}

// HasTradSesOpenTime returns true if TradSesOpenTime is present, Tag 342
func (m TradingSessionStatus) HasTradSesOpenTime() bool {
	return m.Has(tag.TradSesOpenTime)
}

// HasTradSesPreCloseTime returns true if TradSesPreCloseTime is present, Tag 343
func (m TradingSessionStatus) HasTradSesPreCloseTime() bool {
	return m.Has(tag.TradSesPreCloseTime)
}

// HasTradSesCloseTime returns true if TradSesCloseTime is present, Tag 344
func (m TradingSessionStatus) HasTradSesCloseTime() bool {
	return m.Has(tag.TradSesCloseTime)
}

// HasTradSesEndTime returns true if TradSesEndTime is present, Tag 345
func (m TradingSessionStatus) HasTradSesEndTime() bool {
	return m.Has(tag.TradSesEndTime)
}

// HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m TradingSessionStatus) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

// HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m TradingSessionStatus) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

// HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m TradingSessionStatus) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

// HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m TradingSessionStatus) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m TradingSessionStatus) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355
func (m TradingSessionStatus) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasTotalVolumeTraded returns true if TotalVolumeTraded is present, Tag 387
func (m TradingSessionStatus) HasTotalVolumeTraded() bool {
	return m.Has(tag.TotalVolumeTraded)
}

// HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m TradingSessionStatus) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

// HasProduct returns true if Product is present, Tag 460
func (m TradingSessionStatus) HasProduct() bool {
	return m.Has(tag.Product)
}

// HasCFICode returns true if CFICode is present, Tag 461
func (m TradingSessionStatus) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

// HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m TradingSessionStatus) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

// HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m TradingSessionStatus) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

// HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m TradingSessionStatus) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

// HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m TradingSessionStatus) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

// HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m TradingSessionStatus) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

// HasTradSesStatusRejReason returns true if TradSesStatusRejReason is present, Tag 567
func (m TradingSessionStatus) HasTradSesStatusRejReason() bool {
	return m.Has(tag.TradSesStatusRejReason)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m TradingSessionStatus) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

// HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m TradingSessionStatus) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

// HasPool returns true if Pool is present, Tag 691
func (m TradingSessionStatus) HasPool() bool {
	return m.Has(tag.Pool)
}

// HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m TradingSessionStatus) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

// HasNoEvents returns true if NoEvents is present, Tag 864
func (m TradingSessionStatus) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

// HasDatedDate returns true if DatedDate is present, Tag 873
func (m TradingSessionStatus) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

// HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m TradingSessionStatus) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

// HasCPProgram returns true if CPProgram is present, Tag 875
func (m TradingSessionStatus) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

// HasCPRegType returns true if CPRegType is present, Tag 876
func (m TradingSessionStatus) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

// HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m TradingSessionStatus) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

// HasSecurityStatus returns true if SecurityStatus is present, Tag 965
func (m TradingSessionStatus) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

// HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966
func (m TradingSessionStatus) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

// HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967
func (m TradingSessionStatus) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

// HasStrikeValue returns true if StrikeValue is present, Tag 968
func (m TradingSessionStatus) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

// HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969
func (m TradingSessionStatus) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

// HasPositionLimit returns true if PositionLimit is present, Tag 970
func (m TradingSessionStatus) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

// HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971
func (m TradingSessionStatus) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

// HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996
func (m TradingSessionStatus) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

// HasTimeUnit returns true if TimeUnit is present, Tag 997
func (m TradingSessionStatus) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

// HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018
func (m TradingSessionStatus) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

// HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049
func (m TradingSessionStatus) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

// HasMaturityTime returns true if MaturityTime is present, Tag 1079
func (m TradingSessionStatus) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

// HasMinPriceIncrementAmount returns true if MinPriceIncrementAmount is present, Tag 1146
func (m TradingSessionStatus) HasMinPriceIncrementAmount() bool {
	return m.Has(tag.MinPriceIncrementAmount)
}

// HasUnitOfMeasureQty returns true if UnitOfMeasureQty is present, Tag 1147
func (m TradingSessionStatus) HasUnitOfMeasureQty() bool {
	return m.Has(tag.UnitOfMeasureQty)
}

// HasSecurityGroup returns true if SecurityGroup is present, Tag 1151
func (m TradingSessionStatus) HasSecurityGroup() bool {
	return m.Has(tag.SecurityGroup)
}

// HasApplID returns true if ApplID is present, Tag 1180
func (m TradingSessionStatus) HasApplID() bool {
	return m.Has(tag.ApplID)
}

// HasApplSeqNum returns true if ApplSeqNum is present, Tag 1181
func (m TradingSessionStatus) HasApplSeqNum() bool {
	return m.Has(tag.ApplSeqNum)
}

// HasSecurityXMLLen returns true if SecurityXMLLen is present, Tag 1184
func (m TradingSessionStatus) HasSecurityXMLLen() bool {
	return m.Has(tag.SecurityXMLLen)
}

// HasSecurityXML returns true if SecurityXML is present, Tag 1185
func (m TradingSessionStatus) HasSecurityXML() bool {
	return m.Has(tag.SecurityXML)
}

// HasSecurityXMLSchema returns true if SecurityXMLSchema is present, Tag 1186
func (m TradingSessionStatus) HasSecurityXMLSchema() bool {
	return m.Has(tag.SecurityXMLSchema)
}

// HasPriceUnitOfMeasure returns true if PriceUnitOfMeasure is present, Tag 1191
func (m TradingSessionStatus) HasPriceUnitOfMeasure() bool {
	return m.Has(tag.PriceUnitOfMeasure)
}

// HasPriceUnitOfMeasureQty returns true if PriceUnitOfMeasureQty is present, Tag 1192
func (m TradingSessionStatus) HasPriceUnitOfMeasureQty() bool {
	return m.Has(tag.PriceUnitOfMeasureQty)
}

// HasSettlMethod returns true if SettlMethod is present, Tag 1193
func (m TradingSessionStatus) HasSettlMethod() bool {
	return m.Has(tag.SettlMethod)
}

// HasExerciseStyle returns true if ExerciseStyle is present, Tag 1194
func (m TradingSessionStatus) HasExerciseStyle() bool {
	return m.Has(tag.ExerciseStyle)
}

// HasOptPayoutAmount returns true if OptPayoutAmount is present, Tag 1195
func (m TradingSessionStatus) HasOptPayoutAmount() bool {
	return m.Has(tag.OptPayoutAmount)
}

// HasPriceQuoteMethod returns true if PriceQuoteMethod is present, Tag 1196
func (m TradingSessionStatus) HasPriceQuoteMethod() bool {
	return m.Has(tag.PriceQuoteMethod)
}

// HasValuationMethod returns true if ValuationMethod is present, Tag 1197
func (m TradingSessionStatus) HasValuationMethod() bool {
	return m.Has(tag.ValuationMethod)
}

// HasListMethod returns true if ListMethod is present, Tag 1198
func (m TradingSessionStatus) HasListMethod() bool {
	return m.Has(tag.ListMethod)
}

// HasCapPrice returns true if CapPrice is present, Tag 1199
func (m TradingSessionStatus) HasCapPrice() bool {
	return m.Has(tag.CapPrice)
}

// HasFloorPrice returns true if FloorPrice is present, Tag 1200
func (m TradingSessionStatus) HasFloorPrice() bool {
	return m.Has(tag.FloorPrice)
}

// HasProductComplex returns true if ProductComplex is present, Tag 1227
func (m TradingSessionStatus) HasProductComplex() bool {
	return m.Has(tag.ProductComplex)
}

// HasFlexProductEligibilityIndicator returns true if FlexProductEligibilityIndicator is present, Tag 1242
func (m TradingSessionStatus) HasFlexProductEligibilityIndicator() bool {
	return m.Has(tag.FlexProductEligibilityIndicator)
}

// HasFlexibleIndicator returns true if FlexibleIndicator is present, Tag 1244
func (m TradingSessionStatus) HasFlexibleIndicator() bool {
	return m.Has(tag.FlexibleIndicator)
}

// HasMarketSegmentID returns true if MarketSegmentID is present, Tag 1300
func (m TradingSessionStatus) HasMarketSegmentID() bool {
	return m.Has(tag.MarketSegmentID)
}

// HasMarketID returns true if MarketID is present, Tag 1301
func (m TradingSessionStatus) HasMarketID() bool {
	return m.Has(tag.MarketID)
}

// HasApplLastSeqNum returns true if ApplLastSeqNum is present, Tag 1350
func (m TradingSessionStatus) HasApplLastSeqNum() bool {
	return m.Has(tag.ApplLastSeqNum)
}

// HasApplResendFlag returns true if ApplResendFlag is present, Tag 1352
func (m TradingSessionStatus) HasApplResendFlag() bool {
	return m.Has(tag.ApplResendFlag)
}

// HasTradSesEvent returns true if TradSesEvent is present, Tag 1368
func (m TradingSessionStatus) HasTradSesEvent() bool {
	return m.Has(tag.TradSesEvent)
}

// HasContractMultiplierUnit returns true if ContractMultiplierUnit is present, Tag 1435
func (m TradingSessionStatus) HasContractMultiplierUnit() bool {
	return m.Has(tag.ContractMultiplierUnit)
}

// HasFlowScheduleType returns true if FlowScheduleType is present, Tag 1439
func (m TradingSessionStatus) HasFlowScheduleType() bool {
	return m.Has(tag.FlowScheduleType)
}

// HasRestructuringType returns true if RestructuringType is present, Tag 1449
func (m TradingSessionStatus) HasRestructuringType() bool {
	return m.Has(tag.RestructuringType)
}

// HasSeniority returns true if Seniority is present, Tag 1450
func (m TradingSessionStatus) HasSeniority() bool {
	return m.Has(tag.Seniority)
}

// HasNotionalPercentageOutstanding returns true if NotionalPercentageOutstanding is present, Tag 1451
func (m TradingSessionStatus) HasNotionalPercentageOutstanding() bool {
	return m.Has(tag.NotionalPercentageOutstanding)
}

// HasOriginalNotionalPercentageOutstanding returns true if OriginalNotionalPercentageOutstanding is present, Tag 1452
func (m TradingSessionStatus) HasOriginalNotionalPercentageOutstanding() bool {
	return m.Has(tag.OriginalNotionalPercentageOutstanding)
}

// HasAttachmentPoint returns true if AttachmentPoint is present, Tag 1457
func (m TradingSessionStatus) HasAttachmentPoint() bool {
	return m.Has(tag.AttachmentPoint)
}

// HasDetachmentPoint returns true if DetachmentPoint is present, Tag 1458
func (m TradingSessionStatus) HasDetachmentPoint() bool {
	return m.Has(tag.DetachmentPoint)
}

// HasStrikePriceDeterminationMethod returns true if StrikePriceDeterminationMethod is present, Tag 1478
func (m TradingSessionStatus) HasStrikePriceDeterminationMethod() bool {
	return m.Has(tag.StrikePriceDeterminationMethod)
}

// HasStrikePriceBoundaryMethod returns true if StrikePriceBoundaryMethod is present, Tag 1479
func (m TradingSessionStatus) HasStrikePriceBoundaryMethod() bool {
	return m.Has(tag.StrikePriceBoundaryMethod)
}

// HasStrikePriceBoundaryPrecision returns true if StrikePriceBoundaryPrecision is present, Tag 1480
func (m TradingSessionStatus) HasStrikePriceBoundaryPrecision() bool {
	return m.Has(tag.StrikePriceBoundaryPrecision)
}

// HasUnderlyingPriceDeterminationMethod returns true if UnderlyingPriceDeterminationMethod is present, Tag 1481
func (m TradingSessionStatus) HasUnderlyingPriceDeterminationMethod() bool {
	return m.Has(tag.UnderlyingPriceDeterminationMethod)
}

// HasOptPayoutType returns true if OptPayoutType is present, Tag 1482
func (m TradingSessionStatus) HasOptPayoutType() bool {
	return m.Has(tag.OptPayoutType)
}

// HasNoComplexEvents returns true if NoComplexEvents is present, Tag 1483
func (m TradingSessionStatus) HasNoComplexEvents() bool {
	return m.Has(tag.NoComplexEvents)
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

// SetEventTime sets EventTime, Tag 1145
func (m NoEvents) SetEventTime(v time.Time) {
	m.Set(field.NewEventTime(v))
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

// GetEventTime gets EventTime, Tag 1145
func (m NoEvents) GetEventTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.EventTimeField
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

// HasEventTime returns true if EventTime is present, Tag 1145
func (m NoEvents) HasEventTime() bool {
	return m.Has(tag.EventTime)
}

// NoEventsRepeatingGroup is a repeating group, Tag 864
type NoEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoEventsRepeatingGroup returns an initialized, NoEventsRepeatingGroup
func NewNoEventsRepeatingGroup() NoEventsRepeatingGroup {
	return NoEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.EventType), quickfix.GroupElement(tag.EventDate), quickfix.GroupElement(tag.EventPx), quickfix.GroupElement(tag.EventText), quickfix.GroupElement(tag.EventTime)}),
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

// NoInstrumentParties is a repeating group element, Tag 1018
type NoInstrumentParties struct {
	*quickfix.Group
}

// SetInstrumentPartyID sets InstrumentPartyID, Tag 1019
func (m NoInstrumentParties) SetInstrumentPartyID(v string) {
	m.Set(field.NewInstrumentPartyID(v))
}

// SetInstrumentPartyIDSource sets InstrumentPartyIDSource, Tag 1050
func (m NoInstrumentParties) SetInstrumentPartyIDSource(v string) {
	m.Set(field.NewInstrumentPartyIDSource(v))
}

// SetInstrumentPartyRole sets InstrumentPartyRole, Tag 1051
func (m NoInstrumentParties) SetInstrumentPartyRole(v int) {
	m.Set(field.NewInstrumentPartyRole(v))
}

// SetNoInstrumentPartySubIDs sets NoInstrumentPartySubIDs, Tag 1052
func (m NoInstrumentParties) SetNoInstrumentPartySubIDs(f NoInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetInstrumentPartyID gets InstrumentPartyID, Tag 1019
func (m NoInstrumentParties) GetInstrumentPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentPartyIDSource gets InstrumentPartyIDSource, Tag 1050
func (m NoInstrumentParties) GetInstrumentPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentPartyRole gets InstrumentPartyRole, Tag 1051
func (m NoInstrumentParties) GetInstrumentPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoInstrumentPartySubIDs gets NoInstrumentPartySubIDs, Tag 1052
func (m NoInstrumentParties) GetNoInstrumentPartySubIDs() (NoInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasInstrumentPartyID returns true if InstrumentPartyID is present, Tag 1019
func (m NoInstrumentParties) HasInstrumentPartyID() bool {
	return m.Has(tag.InstrumentPartyID)
}

// HasInstrumentPartyIDSource returns true if InstrumentPartyIDSource is present, Tag 1050
func (m NoInstrumentParties) HasInstrumentPartyIDSource() bool {
	return m.Has(tag.InstrumentPartyIDSource)
}

// HasInstrumentPartyRole returns true if InstrumentPartyRole is present, Tag 1051
func (m NoInstrumentParties) HasInstrumentPartyRole() bool {
	return m.Has(tag.InstrumentPartyRole)
}

// HasNoInstrumentPartySubIDs returns true if NoInstrumentPartySubIDs is present, Tag 1052
func (m NoInstrumentParties) HasNoInstrumentPartySubIDs() bool {
	return m.Has(tag.NoInstrumentPartySubIDs)
}

// NoInstrumentPartySubIDs is a repeating group element, Tag 1052
type NoInstrumentPartySubIDs struct {
	*quickfix.Group
}

// SetInstrumentPartySubID sets InstrumentPartySubID, Tag 1053
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubID(v string) {
	m.Set(field.NewInstrumentPartySubID(v))
}

// SetInstrumentPartySubIDType sets InstrumentPartySubIDType, Tag 1054
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubIDType(v int) {
	m.Set(field.NewInstrumentPartySubIDType(v))
}

// GetInstrumentPartySubID gets InstrumentPartySubID, Tag 1053
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentPartySubIDType gets InstrumentPartySubIDType, Tag 1054
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasInstrumentPartySubID returns true if InstrumentPartySubID is present, Tag 1053
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubID() bool {
	return m.Has(tag.InstrumentPartySubID)
}

// HasInstrumentPartySubIDType returns true if InstrumentPartySubIDType is present, Tag 1054
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubIDType() bool {
	return m.Has(tag.InstrumentPartySubIDType)
}

// NoInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1052
type NoInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoInstrumentPartySubIDsRepeatingGroup returns an initialized, NoInstrumentPartySubIDsRepeatingGroup
func NewNoInstrumentPartySubIDsRepeatingGroup() NoInstrumentPartySubIDsRepeatingGroup {
	return NoInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartySubID), quickfix.GroupElement(tag.InstrumentPartySubIDType)}),
	}
}

// Add create and append a new NoInstrumentPartySubIDs to this group
func (m NoInstrumentPartySubIDsRepeatingGroup) Add() NoInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoInstrumentPartySubIDs{g}
}

// Get returns the ith NoInstrumentPartySubIDs in the NoInstrumentPartySubIDsRepeatinGroup
func (m NoInstrumentPartySubIDsRepeatingGroup) Get(i int) NoInstrumentPartySubIDs {
	return NoInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoInstrumentPartiesRepeatingGroup is a repeating group, Tag 1018
type NoInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoInstrumentPartiesRepeatingGroup returns an initialized, NoInstrumentPartiesRepeatingGroup
func NewNoInstrumentPartiesRepeatingGroup() NoInstrumentPartiesRepeatingGroup {
	return NoInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartyID), quickfix.GroupElement(tag.InstrumentPartyIDSource), quickfix.GroupElement(tag.InstrumentPartyRole), NewNoInstrumentPartySubIDsRepeatingGroup()}),
	}
}

// Add create and append a new NoInstrumentParties to this group
func (m NoInstrumentPartiesRepeatingGroup) Add() NoInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoInstrumentParties{g}
}

// Get returns the ith NoInstrumentParties in the NoInstrumentPartiesRepeatinGroup
func (m NoInstrumentPartiesRepeatingGroup) Get(i int) NoInstrumentParties {
	return NoInstrumentParties{m.RepeatingGroup.Get(i)}
}

// NoComplexEvents is a repeating group element, Tag 1483
type NoComplexEvents struct {
	*quickfix.Group
}

// SetComplexEventType sets ComplexEventType, Tag 1484
func (m NoComplexEvents) SetComplexEventType(v enum.ComplexEventType) {
	m.Set(field.NewComplexEventType(v))
}

// SetComplexOptPayoutAmount sets ComplexOptPayoutAmount, Tag 1485
func (m NoComplexEvents) SetComplexOptPayoutAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexOptPayoutAmount(value, scale))
}

// SetComplexEventPrice sets ComplexEventPrice, Tag 1486
func (m NoComplexEvents) SetComplexEventPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexEventPrice(value, scale))
}

// SetComplexEventPriceBoundaryMethod sets ComplexEventPriceBoundaryMethod, Tag 1487
func (m NoComplexEvents) SetComplexEventPriceBoundaryMethod(v enum.ComplexEventPriceBoundaryMethod) {
	m.Set(field.NewComplexEventPriceBoundaryMethod(v))
}

// SetComplexEventPriceBoundaryPrecision sets ComplexEventPriceBoundaryPrecision, Tag 1488
func (m NoComplexEvents) SetComplexEventPriceBoundaryPrecision(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexEventPriceBoundaryPrecision(value, scale))
}

// SetComplexEventPriceTimeType sets ComplexEventPriceTimeType, Tag 1489
func (m NoComplexEvents) SetComplexEventPriceTimeType(v enum.ComplexEventPriceTimeType) {
	m.Set(field.NewComplexEventPriceTimeType(v))
}

// SetComplexEventCondition sets ComplexEventCondition, Tag 1490
func (m NoComplexEvents) SetComplexEventCondition(v enum.ComplexEventCondition) {
	m.Set(field.NewComplexEventCondition(v))
}

// SetNoComplexEventDates sets NoComplexEventDates, Tag 1491
func (m NoComplexEvents) SetNoComplexEventDates(f NoComplexEventDatesRepeatingGroup) {
	m.SetGroup(f)
}

// GetComplexEventType gets ComplexEventType, Tag 1484
func (m NoComplexEvents) GetComplexEventType() (v enum.ComplexEventType, err quickfix.MessageRejectError) {
	var f field.ComplexEventTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexOptPayoutAmount gets ComplexOptPayoutAmount, Tag 1485
func (m NoComplexEvents) GetComplexOptPayoutAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ComplexOptPayoutAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPrice gets ComplexEventPrice, Tag 1486
func (m NoComplexEvents) GetComplexEventPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPriceBoundaryMethod gets ComplexEventPriceBoundaryMethod, Tag 1487
func (m NoComplexEvents) GetComplexEventPriceBoundaryMethod() (v enum.ComplexEventPriceBoundaryMethod, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceBoundaryMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPriceBoundaryPrecision gets ComplexEventPriceBoundaryPrecision, Tag 1488
func (m NoComplexEvents) GetComplexEventPriceBoundaryPrecision() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceBoundaryPrecisionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPriceTimeType gets ComplexEventPriceTimeType, Tag 1489
func (m NoComplexEvents) GetComplexEventPriceTimeType() (v enum.ComplexEventPriceTimeType, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceTimeTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventCondition gets ComplexEventCondition, Tag 1490
func (m NoComplexEvents) GetComplexEventCondition() (v enum.ComplexEventCondition, err quickfix.MessageRejectError) {
	var f field.ComplexEventConditionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoComplexEventDates gets NoComplexEventDates, Tag 1491
func (m NoComplexEvents) GetNoComplexEventDates() (NoComplexEventDatesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventDatesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasComplexEventType returns true if ComplexEventType is present, Tag 1484
func (m NoComplexEvents) HasComplexEventType() bool {
	return m.Has(tag.ComplexEventType)
}

// HasComplexOptPayoutAmount returns true if ComplexOptPayoutAmount is present, Tag 1485
func (m NoComplexEvents) HasComplexOptPayoutAmount() bool {
	return m.Has(tag.ComplexOptPayoutAmount)
}

// HasComplexEventPrice returns true if ComplexEventPrice is present, Tag 1486
func (m NoComplexEvents) HasComplexEventPrice() bool {
	return m.Has(tag.ComplexEventPrice)
}

// HasComplexEventPriceBoundaryMethod returns true if ComplexEventPriceBoundaryMethod is present, Tag 1487
func (m NoComplexEvents) HasComplexEventPriceBoundaryMethod() bool {
	return m.Has(tag.ComplexEventPriceBoundaryMethod)
}

// HasComplexEventPriceBoundaryPrecision returns true if ComplexEventPriceBoundaryPrecision is present, Tag 1488
func (m NoComplexEvents) HasComplexEventPriceBoundaryPrecision() bool {
	return m.Has(tag.ComplexEventPriceBoundaryPrecision)
}

// HasComplexEventPriceTimeType returns true if ComplexEventPriceTimeType is present, Tag 1489
func (m NoComplexEvents) HasComplexEventPriceTimeType() bool {
	return m.Has(tag.ComplexEventPriceTimeType)
}

// HasComplexEventCondition returns true if ComplexEventCondition is present, Tag 1490
func (m NoComplexEvents) HasComplexEventCondition() bool {
	return m.Has(tag.ComplexEventCondition)
}

// HasNoComplexEventDates returns true if NoComplexEventDates is present, Tag 1491
func (m NoComplexEvents) HasNoComplexEventDates() bool {
	return m.Has(tag.NoComplexEventDates)
}

// NoComplexEventDates is a repeating group element, Tag 1491
type NoComplexEventDates struct {
	*quickfix.Group
}

// SetComplexEventStartDate sets ComplexEventStartDate, Tag 1492
func (m NoComplexEventDates) SetComplexEventStartDate(v time.Time) {
	m.Set(field.NewComplexEventStartDate(v))
}

// SetComplexEventEndDate sets ComplexEventEndDate, Tag 1493
func (m NoComplexEventDates) SetComplexEventEndDate(v time.Time) {
	m.Set(field.NewComplexEventEndDate(v))
}

// SetNoComplexEventTimes sets NoComplexEventTimes, Tag 1494
func (m NoComplexEventDates) SetNoComplexEventTimes(f NoComplexEventTimesRepeatingGroup) {
	m.SetGroup(f)
}

// GetComplexEventStartDate gets ComplexEventStartDate, Tag 1492
func (m NoComplexEventDates) GetComplexEventStartDate() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ComplexEventStartDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventEndDate gets ComplexEventEndDate, Tag 1493
func (m NoComplexEventDates) GetComplexEventEndDate() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ComplexEventEndDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoComplexEventTimes gets NoComplexEventTimes, Tag 1494
func (m NoComplexEventDates) GetNoComplexEventTimes() (NoComplexEventTimesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventTimesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasComplexEventStartDate returns true if ComplexEventStartDate is present, Tag 1492
func (m NoComplexEventDates) HasComplexEventStartDate() bool {
	return m.Has(tag.ComplexEventStartDate)
}

// HasComplexEventEndDate returns true if ComplexEventEndDate is present, Tag 1493
func (m NoComplexEventDates) HasComplexEventEndDate() bool {
	return m.Has(tag.ComplexEventEndDate)
}

// HasNoComplexEventTimes returns true if NoComplexEventTimes is present, Tag 1494
func (m NoComplexEventDates) HasNoComplexEventTimes() bool {
	return m.Has(tag.NoComplexEventTimes)
}

// NoComplexEventTimes is a repeating group element, Tag 1494
type NoComplexEventTimes struct {
	*quickfix.Group
}

// SetComplexEventStartTime sets ComplexEventStartTime, Tag 1495
func (m NoComplexEventTimes) SetComplexEventStartTime(v string) {
	m.Set(field.NewComplexEventStartTime(v))
}

// SetComplexEventEndTime sets ComplexEventEndTime, Tag 1496
func (m NoComplexEventTimes) SetComplexEventEndTime(v string) {
	m.Set(field.NewComplexEventEndTime(v))
}

// GetComplexEventStartTime gets ComplexEventStartTime, Tag 1495
func (m NoComplexEventTimes) GetComplexEventStartTime() (v string, err quickfix.MessageRejectError) {
	var f field.ComplexEventStartTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventEndTime gets ComplexEventEndTime, Tag 1496
func (m NoComplexEventTimes) GetComplexEventEndTime() (v string, err quickfix.MessageRejectError) {
	var f field.ComplexEventEndTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasComplexEventStartTime returns true if ComplexEventStartTime is present, Tag 1495
func (m NoComplexEventTimes) HasComplexEventStartTime() bool {
	return m.Has(tag.ComplexEventStartTime)
}

// HasComplexEventEndTime returns true if ComplexEventEndTime is present, Tag 1496
func (m NoComplexEventTimes) HasComplexEventEndTime() bool {
	return m.Has(tag.ComplexEventEndTime)
}

// NoComplexEventTimesRepeatingGroup is a repeating group, Tag 1494
type NoComplexEventTimesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoComplexEventTimesRepeatingGroup returns an initialized, NoComplexEventTimesRepeatingGroup
func NewNoComplexEventTimesRepeatingGroup() NoComplexEventTimesRepeatingGroup {
	return NoComplexEventTimesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEventTimes,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventStartTime), quickfix.GroupElement(tag.ComplexEventEndTime)}),
	}
}

// Add create and append a new NoComplexEventTimes to this group
func (m NoComplexEventTimesRepeatingGroup) Add() NoComplexEventTimes {
	g := m.RepeatingGroup.Add()
	return NoComplexEventTimes{g}
}

// Get returns the ith NoComplexEventTimes in the NoComplexEventTimesRepeatinGroup
func (m NoComplexEventTimesRepeatingGroup) Get(i int) NoComplexEventTimes {
	return NoComplexEventTimes{m.RepeatingGroup.Get(i)}
}

// NoComplexEventDatesRepeatingGroup is a repeating group, Tag 1491
type NoComplexEventDatesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoComplexEventDatesRepeatingGroup returns an initialized, NoComplexEventDatesRepeatingGroup
func NewNoComplexEventDatesRepeatingGroup() NoComplexEventDatesRepeatingGroup {
	return NoComplexEventDatesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEventDates,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventStartDate), quickfix.GroupElement(tag.ComplexEventEndDate), NewNoComplexEventTimesRepeatingGroup()}),
	}
}

// Add create and append a new NoComplexEventDates to this group
func (m NoComplexEventDatesRepeatingGroup) Add() NoComplexEventDates {
	g := m.RepeatingGroup.Add()
	return NoComplexEventDates{g}
}

// Get returns the ith NoComplexEventDates in the NoComplexEventDatesRepeatinGroup
func (m NoComplexEventDatesRepeatingGroup) Get(i int) NoComplexEventDates {
	return NoComplexEventDates{m.RepeatingGroup.Get(i)}
}

// NoComplexEventsRepeatingGroup is a repeating group, Tag 1483
type NoComplexEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoComplexEventsRepeatingGroup returns an initialized, NoComplexEventsRepeatingGroup
func NewNoComplexEventsRepeatingGroup() NoComplexEventsRepeatingGroup {
	return NoComplexEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventType), quickfix.GroupElement(tag.ComplexOptPayoutAmount), quickfix.GroupElement(tag.ComplexEventPrice), quickfix.GroupElement(tag.ComplexEventPriceBoundaryMethod), quickfix.GroupElement(tag.ComplexEventPriceBoundaryPrecision), quickfix.GroupElement(tag.ComplexEventPriceTimeType), quickfix.GroupElement(tag.ComplexEventCondition), NewNoComplexEventDatesRepeatingGroup()}),
	}
}

// Add create and append a new NoComplexEvents to this group
func (m NoComplexEventsRepeatingGroup) Add() NoComplexEvents {
	g := m.RepeatingGroup.Add()
	return NoComplexEvents{g}
}

// Get returns the ith NoComplexEvents in the NoComplexEventsRepeatinGroup
func (m NoComplexEventsRepeatingGroup) Get(i int) NoComplexEvents {
	return NoComplexEvents{m.RepeatingGroup.Get(i)}
}
