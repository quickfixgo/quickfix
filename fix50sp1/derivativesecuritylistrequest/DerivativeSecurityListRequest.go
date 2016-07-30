package derivativesecuritylistrequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//DerivativeSecurityListRequest is the fix50sp1 DerivativeSecurityListRequest type, MsgType = z
type DerivativeSecurityListRequest struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a DerivativeSecurityListRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) DerivativeSecurityListRequest {
	return DerivativeSecurityListRequest{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m DerivativeSecurityListRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a DerivativeSecurityListRequest initialized with the required fields for DerivativeSecurityListRequest
func New(securityreqid field.SecurityReqIDField, securitylistrequesttype field.SecurityListRequestTypeField) (m DerivativeSecurityListRequest) {
	m.Header = fixt11.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("z"))
	m.Header.Set(field.NewBeginString("8"))
	m.Set(securityreqid)
	m.Set(securitylistrequesttype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg DerivativeSecurityListRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "8", "z", r
}

//SetCurrency sets Currency, Tag 15
func (m DerivativeSecurityListRequest) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetText sets Text, Tag 58
func (m DerivativeSecurityListRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetUnderlyingCouponPaymentDate sets UnderlyingCouponPaymentDate, Tag 241
func (m DerivativeSecurityListRequest) SetUnderlyingCouponPaymentDate(v string) {
	m.Set(field.NewUnderlyingCouponPaymentDate(v))
}

//SetUnderlyingIssueDate sets UnderlyingIssueDate, Tag 242
func (m DerivativeSecurityListRequest) SetUnderlyingIssueDate(v string) {
	m.Set(field.NewUnderlyingIssueDate(v))
}

//SetUnderlyingRepoCollateralSecurityType sets UnderlyingRepoCollateralSecurityType, Tag 243
func (m DerivativeSecurityListRequest) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.Set(field.NewUnderlyingRepoCollateralSecurityType(v))
}

//SetUnderlyingRepurchaseTerm sets UnderlyingRepurchaseTerm, Tag 244
func (m DerivativeSecurityListRequest) SetUnderlyingRepurchaseTerm(v int) {
	m.Set(field.NewUnderlyingRepurchaseTerm(v))
}

//SetUnderlyingRepurchaseRate sets UnderlyingRepurchaseRate, Tag 245
func (m DerivativeSecurityListRequest) SetUnderlyingRepurchaseRate(v float64) {
	m.Set(field.NewUnderlyingRepurchaseRate(v))
}

//SetUnderlyingFactor sets UnderlyingFactor, Tag 246
func (m DerivativeSecurityListRequest) SetUnderlyingFactor(v float64) {
	m.Set(field.NewUnderlyingFactor(v))
}

//SetUnderlyingRedemptionDate sets UnderlyingRedemptionDate, Tag 247
func (m DerivativeSecurityListRequest) SetUnderlyingRedemptionDate(v string) {
	m.Set(field.NewUnderlyingRedemptionDate(v))
}

//SetUnderlyingCreditRating sets UnderlyingCreditRating, Tag 256
func (m DerivativeSecurityListRequest) SetUnderlyingCreditRating(v string) {
	m.Set(field.NewUnderlyingCreditRating(v))
}

//SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263
func (m DerivativeSecurityListRequest) SetSubscriptionRequestType(v string) {
	m.Set(field.NewSubscriptionRequestType(v))
}

//SetUnderlyingSecurityIDSource sets UnderlyingSecurityIDSource, Tag 305
func (m DerivativeSecurityListRequest) SetUnderlyingSecurityIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityIDSource(v))
}

//SetUnderlyingIssuer sets UnderlyingIssuer, Tag 306
func (m DerivativeSecurityListRequest) SetUnderlyingIssuer(v string) {
	m.Set(field.NewUnderlyingIssuer(v))
}

//SetUnderlyingSecurityDesc sets UnderlyingSecurityDesc, Tag 307
func (m DerivativeSecurityListRequest) SetUnderlyingSecurityDesc(v string) {
	m.Set(field.NewUnderlyingSecurityDesc(v))
}

//SetUnderlyingSecurityExchange sets UnderlyingSecurityExchange, Tag 308
func (m DerivativeSecurityListRequest) SetUnderlyingSecurityExchange(v string) {
	m.Set(field.NewUnderlyingSecurityExchange(v))
}

//SetUnderlyingSecurityID sets UnderlyingSecurityID, Tag 309
func (m DerivativeSecurityListRequest) SetUnderlyingSecurityID(v string) {
	m.Set(field.NewUnderlyingSecurityID(v))
}

//SetUnderlyingSecurityType sets UnderlyingSecurityType, Tag 310
func (m DerivativeSecurityListRequest) SetUnderlyingSecurityType(v string) {
	m.Set(field.NewUnderlyingSecurityType(v))
}

//SetUnderlyingSymbol sets UnderlyingSymbol, Tag 311
func (m DerivativeSecurityListRequest) SetUnderlyingSymbol(v string) {
	m.Set(field.NewUnderlyingSymbol(v))
}

//SetUnderlyingSymbolSfx sets UnderlyingSymbolSfx, Tag 312
func (m DerivativeSecurityListRequest) SetUnderlyingSymbolSfx(v string) {
	m.Set(field.NewUnderlyingSymbolSfx(v))
}

//SetUnderlyingMaturityMonthYear sets UnderlyingMaturityMonthYear, Tag 313
func (m DerivativeSecurityListRequest) SetUnderlyingMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingMaturityMonthYear(v))
}

//SetUnderlyingPutOrCall sets UnderlyingPutOrCall, Tag 315
func (m DerivativeSecurityListRequest) SetUnderlyingPutOrCall(v int) {
	m.Set(field.NewUnderlyingPutOrCall(v))
}

//SetUnderlyingStrikePrice sets UnderlyingStrikePrice, Tag 316
func (m DerivativeSecurityListRequest) SetUnderlyingStrikePrice(v float64) {
	m.Set(field.NewUnderlyingStrikePrice(v))
}

//SetUnderlyingOptAttribute sets UnderlyingOptAttribute, Tag 317
func (m DerivativeSecurityListRequest) SetUnderlyingOptAttribute(v string) {
	m.Set(field.NewUnderlyingOptAttribute(v))
}

//SetUnderlyingCurrency sets UnderlyingCurrency, Tag 318
func (m DerivativeSecurityListRequest) SetUnderlyingCurrency(v string) {
	m.Set(field.NewUnderlyingCurrency(v))
}

//SetSecurityReqID sets SecurityReqID, Tag 320
func (m DerivativeSecurityListRequest) SetSecurityReqID(v string) {
	m.Set(field.NewSecurityReqID(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m DerivativeSecurityListRequest) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m DerivativeSecurityListRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m DerivativeSecurityListRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetEncodedUnderlyingIssuerLen sets EncodedUnderlyingIssuerLen, Tag 362
func (m DerivativeSecurityListRequest) SetEncodedUnderlyingIssuerLen(v int) {
	m.Set(field.NewEncodedUnderlyingIssuerLen(v))
}

//SetEncodedUnderlyingIssuer sets EncodedUnderlyingIssuer, Tag 363
func (m DerivativeSecurityListRequest) SetEncodedUnderlyingIssuer(v string) {
	m.Set(field.NewEncodedUnderlyingIssuer(v))
}

//SetEncodedUnderlyingSecurityDescLen sets EncodedUnderlyingSecurityDescLen, Tag 364
func (m DerivativeSecurityListRequest) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.Set(field.NewEncodedUnderlyingSecurityDescLen(v))
}

//SetEncodedUnderlyingSecurityDesc sets EncodedUnderlyingSecurityDesc, Tag 365
func (m DerivativeSecurityListRequest) SetEncodedUnderlyingSecurityDesc(v string) {
	m.Set(field.NewEncodedUnderlyingSecurityDesc(v))
}

//SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435
func (m DerivativeSecurityListRequest) SetUnderlyingCouponRate(v float64) {
	m.Set(field.NewUnderlyingCouponRate(v))
}

//SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436
func (m DerivativeSecurityListRequest) SetUnderlyingContractMultiplier(v float64) {
	m.Set(field.NewUnderlyingContractMultiplier(v))
}

//SetNoUnderlyingSecurityAltID sets NoUnderlyingSecurityAltID, Tag 457
func (m DerivativeSecurityListRequest) SetNoUnderlyingSecurityAltID(f NoUnderlyingSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingProduct sets UnderlyingProduct, Tag 462
func (m DerivativeSecurityListRequest) SetUnderlyingProduct(v int) {
	m.Set(field.NewUnderlyingProduct(v))
}

//SetUnderlyingCFICode sets UnderlyingCFICode, Tag 463
func (m DerivativeSecurityListRequest) SetUnderlyingCFICode(v string) {
	m.Set(field.NewUnderlyingCFICode(v))
}

//SetUnderlyingMaturityDate sets UnderlyingMaturityDate, Tag 542
func (m DerivativeSecurityListRequest) SetUnderlyingMaturityDate(v string) {
	m.Set(field.NewUnderlyingMaturityDate(v))
}

//SetSecurityListRequestType sets SecurityListRequestType, Tag 559
func (m DerivativeSecurityListRequest) SetSecurityListRequestType(v int) {
	m.Set(field.NewSecurityListRequestType(v))
}

//SetUnderlyingCountryOfIssue sets UnderlyingCountryOfIssue, Tag 592
func (m DerivativeSecurityListRequest) SetUnderlyingCountryOfIssue(v string) {
	m.Set(field.NewUnderlyingCountryOfIssue(v))
}

//SetUnderlyingStateOrProvinceOfIssue sets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m DerivativeSecurityListRequest) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.Set(field.NewUnderlyingStateOrProvinceOfIssue(v))
}

//SetUnderlyingLocaleOfIssue sets UnderlyingLocaleOfIssue, Tag 594
func (m DerivativeSecurityListRequest) SetUnderlyingLocaleOfIssue(v string) {
	m.Set(field.NewUnderlyingLocaleOfIssue(v))
}

//SetUnderlyingInstrRegistry sets UnderlyingInstrRegistry, Tag 595
func (m DerivativeSecurityListRequest) SetUnderlyingInstrRegistry(v string) {
	m.Set(field.NewUnderlyingInstrRegistry(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m DerivativeSecurityListRequest) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m DerivativeSecurityListRequest) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

//SetUnderlyingSecuritySubType sets UnderlyingSecuritySubType, Tag 763
func (m DerivativeSecurityListRequest) SetUnderlyingSecuritySubType(v string) {
	m.Set(field.NewUnderlyingSecuritySubType(v))
}

//SetUnderlyingPx sets UnderlyingPx, Tag 810
func (m DerivativeSecurityListRequest) SetUnderlyingPx(v float64) {
	m.Set(field.NewUnderlyingPx(v))
}

//SetUnderlyingCPProgram sets UnderlyingCPProgram, Tag 877
func (m DerivativeSecurityListRequest) SetUnderlyingCPProgram(v string) {
	m.Set(field.NewUnderlyingCPProgram(v))
}

//SetUnderlyingCPRegType sets UnderlyingCPRegType, Tag 878
func (m DerivativeSecurityListRequest) SetUnderlyingCPRegType(v string) {
	m.Set(field.NewUnderlyingCPRegType(v))
}

//SetUnderlyingQty sets UnderlyingQty, Tag 879
func (m DerivativeSecurityListRequest) SetUnderlyingQty(v float64) {
	m.Set(field.NewUnderlyingQty(v))
}

//SetUnderlyingDirtyPrice sets UnderlyingDirtyPrice, Tag 882
func (m DerivativeSecurityListRequest) SetUnderlyingDirtyPrice(v float64) {
	m.Set(field.NewUnderlyingDirtyPrice(v))
}

//SetUnderlyingEndPrice sets UnderlyingEndPrice, Tag 883
func (m DerivativeSecurityListRequest) SetUnderlyingEndPrice(v float64) {
	m.Set(field.NewUnderlyingEndPrice(v))
}

//SetUnderlyingStartValue sets UnderlyingStartValue, Tag 884
func (m DerivativeSecurityListRequest) SetUnderlyingStartValue(v float64) {
	m.Set(field.NewUnderlyingStartValue(v))
}

//SetUnderlyingCurrentValue sets UnderlyingCurrentValue, Tag 885
func (m DerivativeSecurityListRequest) SetUnderlyingCurrentValue(v float64) {
	m.Set(field.NewUnderlyingCurrentValue(v))
}

//SetUnderlyingEndValue sets UnderlyingEndValue, Tag 886
func (m DerivativeSecurityListRequest) SetUnderlyingEndValue(v float64) {
	m.Set(field.NewUnderlyingEndValue(v))
}

//SetNoUnderlyingStips sets NoUnderlyingStips, Tag 887
func (m DerivativeSecurityListRequest) SetNoUnderlyingStips(f NoUnderlyingStipsRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingStrikeCurrency sets UnderlyingStrikeCurrency, Tag 941
func (m DerivativeSecurityListRequest) SetUnderlyingStrikeCurrency(v string) {
	m.Set(field.NewUnderlyingStrikeCurrency(v))
}

//SetUnderlyingAllocationPercent sets UnderlyingAllocationPercent, Tag 972
func (m DerivativeSecurityListRequest) SetUnderlyingAllocationPercent(v float64) {
	m.Set(field.NewUnderlyingAllocationPercent(v))
}

//SetUnderlyingCashAmount sets UnderlyingCashAmount, Tag 973
func (m DerivativeSecurityListRequest) SetUnderlyingCashAmount(v float64) {
	m.Set(field.NewUnderlyingCashAmount(v))
}

//SetUnderlyingCashType sets UnderlyingCashType, Tag 974
func (m DerivativeSecurityListRequest) SetUnderlyingCashType(v string) {
	m.Set(field.NewUnderlyingCashType(v))
}

//SetUnderlyingSettlementType sets UnderlyingSettlementType, Tag 975
func (m DerivativeSecurityListRequest) SetUnderlyingSettlementType(v int) {
	m.Set(field.NewUnderlyingSettlementType(v))
}

//SetUnderlyingUnitOfMeasure sets UnderlyingUnitOfMeasure, Tag 998
func (m DerivativeSecurityListRequest) SetUnderlyingUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingUnitOfMeasure(v))
}

//SetUnderlyingTimeUnit sets UnderlyingTimeUnit, Tag 1000
func (m DerivativeSecurityListRequest) SetUnderlyingTimeUnit(v string) {
	m.Set(field.NewUnderlyingTimeUnit(v))
}

//SetUnderlyingCapValue sets UnderlyingCapValue, Tag 1038
func (m DerivativeSecurityListRequest) SetUnderlyingCapValue(v float64) {
	m.Set(field.NewUnderlyingCapValue(v))
}

//SetUnderlyingSettlMethod sets UnderlyingSettlMethod, Tag 1039
func (m DerivativeSecurityListRequest) SetUnderlyingSettlMethod(v string) {
	m.Set(field.NewUnderlyingSettlMethod(v))
}

//SetUnderlyingAdjustedQuantity sets UnderlyingAdjustedQuantity, Tag 1044
func (m DerivativeSecurityListRequest) SetUnderlyingAdjustedQuantity(v float64) {
	m.Set(field.NewUnderlyingAdjustedQuantity(v))
}

//SetUnderlyingFXRate sets UnderlyingFXRate, Tag 1045
func (m DerivativeSecurityListRequest) SetUnderlyingFXRate(v float64) {
	m.Set(field.NewUnderlyingFXRate(v))
}

//SetUnderlyingFXRateCalc sets UnderlyingFXRateCalc, Tag 1046
func (m DerivativeSecurityListRequest) SetUnderlyingFXRateCalc(v string) {
	m.Set(field.NewUnderlyingFXRateCalc(v))
}

//SetNoUndlyInstrumentParties sets NoUndlyInstrumentParties, Tag 1058
func (m DerivativeSecurityListRequest) SetNoUndlyInstrumentParties(f NoUndlyInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingMaturityTime sets UnderlyingMaturityTime, Tag 1213
func (m DerivativeSecurityListRequest) SetUnderlyingMaturityTime(v string) {
	m.Set(field.NewUnderlyingMaturityTime(v))
}

//SetDerivativeSymbol sets DerivativeSymbol, Tag 1214
func (m DerivativeSecurityListRequest) SetDerivativeSymbol(v string) {
	m.Set(field.NewDerivativeSymbol(v))
}

//SetDerivativeSymbolSfx sets DerivativeSymbolSfx, Tag 1215
func (m DerivativeSecurityListRequest) SetDerivativeSymbolSfx(v string) {
	m.Set(field.NewDerivativeSymbolSfx(v))
}

//SetDerivativeSecurityID sets DerivativeSecurityID, Tag 1216
func (m DerivativeSecurityListRequest) SetDerivativeSecurityID(v string) {
	m.Set(field.NewDerivativeSecurityID(v))
}

//SetDerivativeSecurityIDSource sets DerivativeSecurityIDSource, Tag 1217
func (m DerivativeSecurityListRequest) SetDerivativeSecurityIDSource(v string) {
	m.Set(field.NewDerivativeSecurityIDSource(v))
}

//SetNoDerivativeSecurityAltID sets NoDerivativeSecurityAltID, Tag 1218
func (m DerivativeSecurityListRequest) SetNoDerivativeSecurityAltID(f NoDerivativeSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetDerivativeOptPayAmount sets DerivativeOptPayAmount, Tag 1225
func (m DerivativeSecurityListRequest) SetDerivativeOptPayAmount(v float64) {
	m.Set(field.NewDerivativeOptPayAmount(v))
}

//SetDerivativeProductComplex sets DerivativeProductComplex, Tag 1228
func (m DerivativeSecurityListRequest) SetDerivativeProductComplex(v string) {
	m.Set(field.NewDerivativeProductComplex(v))
}

//SetDerivFlexProductEligibilityIndicator sets DerivFlexProductEligibilityIndicator, Tag 1243
func (m DerivativeSecurityListRequest) SetDerivFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewDerivFlexProductEligibilityIndicator(v))
}

//SetDerivativeProduct sets DerivativeProduct, Tag 1246
func (m DerivativeSecurityListRequest) SetDerivativeProduct(v int) {
	m.Set(field.NewDerivativeProduct(v))
}

//SetDerivativeSecurityGroup sets DerivativeSecurityGroup, Tag 1247
func (m DerivativeSecurityListRequest) SetDerivativeSecurityGroup(v string) {
	m.Set(field.NewDerivativeSecurityGroup(v))
}

//SetDerivativeCFICode sets DerivativeCFICode, Tag 1248
func (m DerivativeSecurityListRequest) SetDerivativeCFICode(v string) {
	m.Set(field.NewDerivativeCFICode(v))
}

//SetDerivativeSecurityType sets DerivativeSecurityType, Tag 1249
func (m DerivativeSecurityListRequest) SetDerivativeSecurityType(v string) {
	m.Set(field.NewDerivativeSecurityType(v))
}

//SetDerivativeSecuritySubType sets DerivativeSecuritySubType, Tag 1250
func (m DerivativeSecurityListRequest) SetDerivativeSecuritySubType(v string) {
	m.Set(field.NewDerivativeSecuritySubType(v))
}

//SetDerivativeMaturityMonthYear sets DerivativeMaturityMonthYear, Tag 1251
func (m DerivativeSecurityListRequest) SetDerivativeMaturityMonthYear(v string) {
	m.Set(field.NewDerivativeMaturityMonthYear(v))
}

//SetDerivativeMaturityDate sets DerivativeMaturityDate, Tag 1252
func (m DerivativeSecurityListRequest) SetDerivativeMaturityDate(v string) {
	m.Set(field.NewDerivativeMaturityDate(v))
}

//SetDerivativeMaturityTime sets DerivativeMaturityTime, Tag 1253
func (m DerivativeSecurityListRequest) SetDerivativeMaturityTime(v string) {
	m.Set(field.NewDerivativeMaturityTime(v))
}

//SetDerivativeSettleOnOpenFlag sets DerivativeSettleOnOpenFlag, Tag 1254
func (m DerivativeSecurityListRequest) SetDerivativeSettleOnOpenFlag(v string) {
	m.Set(field.NewDerivativeSettleOnOpenFlag(v))
}

//SetDerivativeInstrmtAssignmentMethod sets DerivativeInstrmtAssignmentMethod, Tag 1255
func (m DerivativeSecurityListRequest) SetDerivativeInstrmtAssignmentMethod(v string) {
	m.Set(field.NewDerivativeInstrmtAssignmentMethod(v))
}

//SetDerivativeSecurityStatus sets DerivativeSecurityStatus, Tag 1256
func (m DerivativeSecurityListRequest) SetDerivativeSecurityStatus(v string) {
	m.Set(field.NewDerivativeSecurityStatus(v))
}

//SetDerivativeInstrRegistry sets DerivativeInstrRegistry, Tag 1257
func (m DerivativeSecurityListRequest) SetDerivativeInstrRegistry(v string) {
	m.Set(field.NewDerivativeInstrRegistry(v))
}

//SetDerivativeCountryOfIssue sets DerivativeCountryOfIssue, Tag 1258
func (m DerivativeSecurityListRequest) SetDerivativeCountryOfIssue(v string) {
	m.Set(field.NewDerivativeCountryOfIssue(v))
}

//SetDerivativeStateOrProvinceOfIssue sets DerivativeStateOrProvinceOfIssue, Tag 1259
func (m DerivativeSecurityListRequest) SetDerivativeStateOrProvinceOfIssue(v string) {
	m.Set(field.NewDerivativeStateOrProvinceOfIssue(v))
}

//SetDerivativeLocaleOfIssue sets DerivativeLocaleOfIssue, Tag 1260
func (m DerivativeSecurityListRequest) SetDerivativeLocaleOfIssue(v string) {
	m.Set(field.NewDerivativeLocaleOfIssue(v))
}

//SetDerivativeStrikePrice sets DerivativeStrikePrice, Tag 1261
func (m DerivativeSecurityListRequest) SetDerivativeStrikePrice(v float64) {
	m.Set(field.NewDerivativeStrikePrice(v))
}

//SetDerivativeStrikeCurrency sets DerivativeStrikeCurrency, Tag 1262
func (m DerivativeSecurityListRequest) SetDerivativeStrikeCurrency(v string) {
	m.Set(field.NewDerivativeStrikeCurrency(v))
}

//SetDerivativeStrikeMultiplier sets DerivativeStrikeMultiplier, Tag 1263
func (m DerivativeSecurityListRequest) SetDerivativeStrikeMultiplier(v float64) {
	m.Set(field.NewDerivativeStrikeMultiplier(v))
}

//SetDerivativeStrikeValue sets DerivativeStrikeValue, Tag 1264
func (m DerivativeSecurityListRequest) SetDerivativeStrikeValue(v float64) {
	m.Set(field.NewDerivativeStrikeValue(v))
}

//SetDerivativeOptAttribute sets DerivativeOptAttribute, Tag 1265
func (m DerivativeSecurityListRequest) SetDerivativeOptAttribute(v string) {
	m.Set(field.NewDerivativeOptAttribute(v))
}

//SetDerivativeContractMultiplier sets DerivativeContractMultiplier, Tag 1266
func (m DerivativeSecurityListRequest) SetDerivativeContractMultiplier(v float64) {
	m.Set(field.NewDerivativeContractMultiplier(v))
}

//SetDerivativeMinPriceIncrement sets DerivativeMinPriceIncrement, Tag 1267
func (m DerivativeSecurityListRequest) SetDerivativeMinPriceIncrement(v float64) {
	m.Set(field.NewDerivativeMinPriceIncrement(v))
}

//SetDerivativeMinPriceIncrementAmount sets DerivativeMinPriceIncrementAmount, Tag 1268
func (m DerivativeSecurityListRequest) SetDerivativeMinPriceIncrementAmount(v float64) {
	m.Set(field.NewDerivativeMinPriceIncrementAmount(v))
}

//SetDerivativeUnitOfMeasure sets DerivativeUnitOfMeasure, Tag 1269
func (m DerivativeSecurityListRequest) SetDerivativeUnitOfMeasure(v string) {
	m.Set(field.NewDerivativeUnitOfMeasure(v))
}

//SetDerivativeUnitOfMeasureQty sets DerivativeUnitOfMeasureQty, Tag 1270
func (m DerivativeSecurityListRequest) SetDerivativeUnitOfMeasureQty(v float64) {
	m.Set(field.NewDerivativeUnitOfMeasureQty(v))
}

//SetDerivativeTimeUnit sets DerivativeTimeUnit, Tag 1271
func (m DerivativeSecurityListRequest) SetDerivativeTimeUnit(v string) {
	m.Set(field.NewDerivativeTimeUnit(v))
}

//SetDerivativeSecurityExchange sets DerivativeSecurityExchange, Tag 1272
func (m DerivativeSecurityListRequest) SetDerivativeSecurityExchange(v string) {
	m.Set(field.NewDerivativeSecurityExchange(v))
}

//SetDerivativePositionLimit sets DerivativePositionLimit, Tag 1273
func (m DerivativeSecurityListRequest) SetDerivativePositionLimit(v int) {
	m.Set(field.NewDerivativePositionLimit(v))
}

//SetDerivativeNTPositionLimit sets DerivativeNTPositionLimit, Tag 1274
func (m DerivativeSecurityListRequest) SetDerivativeNTPositionLimit(v int) {
	m.Set(field.NewDerivativeNTPositionLimit(v))
}

//SetDerivativeIssuer sets DerivativeIssuer, Tag 1275
func (m DerivativeSecurityListRequest) SetDerivativeIssuer(v string) {
	m.Set(field.NewDerivativeIssuer(v))
}

//SetDerivativeIssueDate sets DerivativeIssueDate, Tag 1276
func (m DerivativeSecurityListRequest) SetDerivativeIssueDate(v string) {
	m.Set(field.NewDerivativeIssueDate(v))
}

//SetDerivativeEncodedIssuerLen sets DerivativeEncodedIssuerLen, Tag 1277
func (m DerivativeSecurityListRequest) SetDerivativeEncodedIssuerLen(v int) {
	m.Set(field.NewDerivativeEncodedIssuerLen(v))
}

//SetDerivativeEncodedIssuer sets DerivativeEncodedIssuer, Tag 1278
func (m DerivativeSecurityListRequest) SetDerivativeEncodedIssuer(v string) {
	m.Set(field.NewDerivativeEncodedIssuer(v))
}

//SetDerivativeSecurityDesc sets DerivativeSecurityDesc, Tag 1279
func (m DerivativeSecurityListRequest) SetDerivativeSecurityDesc(v string) {
	m.Set(field.NewDerivativeSecurityDesc(v))
}

//SetDerivativeEncodedSecurityDescLen sets DerivativeEncodedSecurityDescLen, Tag 1280
func (m DerivativeSecurityListRequest) SetDerivativeEncodedSecurityDescLen(v int) {
	m.Set(field.NewDerivativeEncodedSecurityDescLen(v))
}

//SetDerivativeEncodedSecurityDesc sets DerivativeEncodedSecurityDesc, Tag 1281
func (m DerivativeSecurityListRequest) SetDerivativeEncodedSecurityDesc(v string) {
	m.Set(field.NewDerivativeEncodedSecurityDesc(v))
}

//SetDerivativeSecurityXMLLen sets DerivativeSecurityXMLLen, Tag 1282
func (m DerivativeSecurityListRequest) SetDerivativeSecurityXMLLen(v int) {
	m.Set(field.NewDerivativeSecurityXMLLen(v))
}

//SetDerivativeSecurityXML sets DerivativeSecurityXML, Tag 1283
func (m DerivativeSecurityListRequest) SetDerivativeSecurityXML(v string) {
	m.Set(field.NewDerivativeSecurityXML(v))
}

//SetDerivativeSecurityXMLSchema sets DerivativeSecurityXMLSchema, Tag 1284
func (m DerivativeSecurityListRequest) SetDerivativeSecurityXMLSchema(v string) {
	m.Set(field.NewDerivativeSecurityXMLSchema(v))
}

//SetDerivativeContractSettlMonth sets DerivativeContractSettlMonth, Tag 1285
func (m DerivativeSecurityListRequest) SetDerivativeContractSettlMonth(v string) {
	m.Set(field.NewDerivativeContractSettlMonth(v))
}

//SetNoDerivativeEvents sets NoDerivativeEvents, Tag 1286
func (m DerivativeSecurityListRequest) SetNoDerivativeEvents(f NoDerivativeEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoDerivativeInstrumentParties sets NoDerivativeInstrumentParties, Tag 1292
func (m DerivativeSecurityListRequest) SetNoDerivativeInstrumentParties(f NoDerivativeInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetDerivativeExerciseStyle sets DerivativeExerciseStyle, Tag 1299
func (m DerivativeSecurityListRequest) SetDerivativeExerciseStyle(v string) {
	m.Set(field.NewDerivativeExerciseStyle(v))
}

//SetMarketSegmentID sets MarketSegmentID, Tag 1300
func (m DerivativeSecurityListRequest) SetMarketSegmentID(v string) {
	m.Set(field.NewMarketSegmentID(v))
}

//SetMarketID sets MarketID, Tag 1301
func (m DerivativeSecurityListRequest) SetMarketID(v string) {
	m.Set(field.NewMarketID(v))
}

//SetDerivativePriceUnitOfMeasure sets DerivativePriceUnitOfMeasure, Tag 1315
func (m DerivativeSecurityListRequest) SetDerivativePriceUnitOfMeasure(v string) {
	m.Set(field.NewDerivativePriceUnitOfMeasure(v))
}

//SetDerivativePriceUnitOfMeasureQty sets DerivativePriceUnitOfMeasureQty, Tag 1316
func (m DerivativeSecurityListRequest) SetDerivativePriceUnitOfMeasureQty(v float64) {
	m.Set(field.NewDerivativePriceUnitOfMeasureQty(v))
}

//SetDerivativeSettlMethod sets DerivativeSettlMethod, Tag 1317
func (m DerivativeSecurityListRequest) SetDerivativeSettlMethod(v string) {
	m.Set(field.NewDerivativeSettlMethod(v))
}

//SetDerivativePriceQuoteMethod sets DerivativePriceQuoteMethod, Tag 1318
func (m DerivativeSecurityListRequest) SetDerivativePriceQuoteMethod(v string) {
	m.Set(field.NewDerivativePriceQuoteMethod(v))
}

//SetDerivativeFuturesValuationMethod sets DerivativeFuturesValuationMethod, Tag 1319
func (m DerivativeSecurityListRequest) SetDerivativeFuturesValuationMethod(v string) {
	m.Set(field.NewDerivativeFuturesValuationMethod(v))
}

//SetDerivativeListMethod sets DerivativeListMethod, Tag 1320
func (m DerivativeSecurityListRequest) SetDerivativeListMethod(v int) {
	m.Set(field.NewDerivativeListMethod(v))
}

//SetDerivativeCapPrice sets DerivativeCapPrice, Tag 1321
func (m DerivativeSecurityListRequest) SetDerivativeCapPrice(v float64) {
	m.Set(field.NewDerivativeCapPrice(v))
}

//SetDerivativeFloorPrice sets DerivativeFloorPrice, Tag 1322
func (m DerivativeSecurityListRequest) SetDerivativeFloorPrice(v float64) {
	m.Set(field.NewDerivativeFloorPrice(v))
}

//SetDerivativePutOrCall sets DerivativePutOrCall, Tag 1323
func (m DerivativeSecurityListRequest) SetDerivativePutOrCall(v int) {
	m.Set(field.NewDerivativePutOrCall(v))
}

//SetUnderlyingExerciseStyle sets UnderlyingExerciseStyle, Tag 1419
func (m DerivativeSecurityListRequest) SetUnderlyingExerciseStyle(v int) {
	m.Set(field.NewUnderlyingExerciseStyle(v))
}

//SetUnderlyingUnitOfMeasureQty sets UnderlyingUnitOfMeasureQty, Tag 1423
func (m DerivativeSecurityListRequest) SetUnderlyingUnitOfMeasureQty(v float64) {
	m.Set(field.NewUnderlyingUnitOfMeasureQty(v))
}

//SetUnderlyingPriceUnitOfMeasure sets UnderlyingPriceUnitOfMeasure, Tag 1424
func (m DerivativeSecurityListRequest) SetUnderlyingPriceUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasure(v))
}

//SetUnderlyingPriceUnitOfMeasureQty sets UnderlyingPriceUnitOfMeasureQty, Tag 1425
func (m DerivativeSecurityListRequest) SetUnderlyingPriceUnitOfMeasureQty(v float64) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasureQty(v))
}

//GetCurrency gets Currency, Tag 15
func (m DerivativeSecurityListRequest) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m DerivativeSecurityListRequest) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCouponPaymentDate gets UnderlyingCouponPaymentDate, Tag 241
func (m DerivativeSecurityListRequest) GetUnderlyingCouponPaymentDate() (f field.UnderlyingCouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingIssueDate gets UnderlyingIssueDate, Tag 242
func (m DerivativeSecurityListRequest) GetUnderlyingIssueDate() (f field.UnderlyingIssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRepoCollateralSecurityType gets UnderlyingRepoCollateralSecurityType, Tag 243
func (m DerivativeSecurityListRequest) GetUnderlyingRepoCollateralSecurityType() (f field.UnderlyingRepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRepurchaseTerm gets UnderlyingRepurchaseTerm, Tag 244
func (m DerivativeSecurityListRequest) GetUnderlyingRepurchaseTerm() (f field.UnderlyingRepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRepurchaseRate gets UnderlyingRepurchaseRate, Tag 245
func (m DerivativeSecurityListRequest) GetUnderlyingRepurchaseRate() (f field.UnderlyingRepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFactor gets UnderlyingFactor, Tag 246
func (m DerivativeSecurityListRequest) GetUnderlyingFactor() (f field.UnderlyingFactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRedemptionDate gets UnderlyingRedemptionDate, Tag 247
func (m DerivativeSecurityListRequest) GetUnderlyingRedemptionDate() (f field.UnderlyingRedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCreditRating gets UnderlyingCreditRating, Tag 256
func (m DerivativeSecurityListRequest) GetUnderlyingCreditRating() (f field.UnderlyingCreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263
func (m DerivativeSecurityListRequest) GetSubscriptionRequestType() (f field.SubscriptionRequestTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityIDSource gets UnderlyingSecurityIDSource, Tag 305
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityIDSource() (f field.UnderlyingSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306
func (m DerivativeSecurityListRequest) GetUnderlyingIssuer() (f field.UnderlyingIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityDesc() (f field.UnderlyingSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityExchange() (f field.UnderlyingSecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityID() (f field.UnderlyingSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityType() (f field.UnderlyingSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311
func (m DerivativeSecurityListRequest) GetUnderlyingSymbol() (f field.UnderlyingSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312
func (m DerivativeSecurityListRequest) GetUnderlyingSymbolSfx() (f field.UnderlyingSymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313
func (m DerivativeSecurityListRequest) GetUnderlyingMaturityMonthYear() (f field.UnderlyingMaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPutOrCall gets UnderlyingPutOrCall, Tag 315
func (m DerivativeSecurityListRequest) GetUnderlyingPutOrCall() (f field.UnderlyingPutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316
func (m DerivativeSecurityListRequest) GetUnderlyingStrikePrice() (f field.UnderlyingStrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317
func (m DerivativeSecurityListRequest) GetUnderlyingOptAttribute() (f field.UnderlyingOptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCurrency gets UnderlyingCurrency, Tag 318
func (m DerivativeSecurityListRequest) GetUnderlyingCurrency() (f field.UnderlyingCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityReqID gets SecurityReqID, Tag 320
func (m DerivativeSecurityListRequest) GetSecurityReqID() (f field.SecurityReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m DerivativeSecurityListRequest) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m DerivativeSecurityListRequest) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m DerivativeSecurityListRequest) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingIssuerLen() (f field.EncodedUnderlyingIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingIssuer() (f field.EncodedUnderlyingIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingSecurityDescLen() (f field.EncodedUnderlyingSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingSecurityDesc() (f field.EncodedUnderlyingSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435
func (m DerivativeSecurityListRequest) GetUnderlyingCouponRate() (f field.UnderlyingCouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436
func (m DerivativeSecurityListRequest) GetUnderlyingContractMultiplier() (f field.UnderlyingContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyingSecurityAltID gets NoUnderlyingSecurityAltID, Tag 457
func (m DerivativeSecurityListRequest) GetNoUnderlyingSecurityAltID() (NoUnderlyingSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingProduct gets UnderlyingProduct, Tag 462
func (m DerivativeSecurityListRequest) GetUnderlyingProduct() (f field.UnderlyingProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCFICode gets UnderlyingCFICode, Tag 463
func (m DerivativeSecurityListRequest) GetUnderlyingCFICode() (f field.UnderlyingCFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingMaturityDate gets UnderlyingMaturityDate, Tag 542
func (m DerivativeSecurityListRequest) GetUnderlyingMaturityDate() (f field.UnderlyingMaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityListRequestType gets SecurityListRequestType, Tag 559
func (m DerivativeSecurityListRequest) GetSecurityListRequestType() (f field.SecurityListRequestTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCountryOfIssue gets UnderlyingCountryOfIssue, Tag 592
func (m DerivativeSecurityListRequest) GetUnderlyingCountryOfIssue() (f field.UnderlyingCountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStateOrProvinceOfIssue gets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m DerivativeSecurityListRequest) GetUnderlyingStateOrProvinceOfIssue() (f field.UnderlyingStateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLocaleOfIssue gets UnderlyingLocaleOfIssue, Tag 594
func (m DerivativeSecurityListRequest) GetUnderlyingLocaleOfIssue() (f field.UnderlyingLocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingInstrRegistry gets UnderlyingInstrRegistry, Tag 595
func (m DerivativeSecurityListRequest) GetUnderlyingInstrRegistry() (f field.UnderlyingInstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m DerivativeSecurityListRequest) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m DerivativeSecurityListRequest) GetSecuritySubType() (f field.SecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecuritySubType gets UnderlyingSecuritySubType, Tag 763
func (m DerivativeSecurityListRequest) GetUnderlyingSecuritySubType() (f field.UnderlyingSecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPx gets UnderlyingPx, Tag 810
func (m DerivativeSecurityListRequest) GetUnderlyingPx() (f field.UnderlyingPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCPProgram gets UnderlyingCPProgram, Tag 877
func (m DerivativeSecurityListRequest) GetUnderlyingCPProgram() (f field.UnderlyingCPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCPRegType gets UnderlyingCPRegType, Tag 878
func (m DerivativeSecurityListRequest) GetUnderlyingCPRegType() (f field.UnderlyingCPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingQty gets UnderlyingQty, Tag 879
func (m DerivativeSecurityListRequest) GetUnderlyingQty() (f field.UnderlyingQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingDirtyPrice gets UnderlyingDirtyPrice, Tag 882
func (m DerivativeSecurityListRequest) GetUnderlyingDirtyPrice() (f field.UnderlyingDirtyPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingEndPrice gets UnderlyingEndPrice, Tag 883
func (m DerivativeSecurityListRequest) GetUnderlyingEndPrice() (f field.UnderlyingEndPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStartValue gets UnderlyingStartValue, Tag 884
func (m DerivativeSecurityListRequest) GetUnderlyingStartValue() (f field.UnderlyingStartValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCurrentValue gets UnderlyingCurrentValue, Tag 885
func (m DerivativeSecurityListRequest) GetUnderlyingCurrentValue() (f field.UnderlyingCurrentValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingEndValue gets UnderlyingEndValue, Tag 886
func (m DerivativeSecurityListRequest) GetUnderlyingEndValue() (f field.UnderlyingEndValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyingStips gets NoUnderlyingStips, Tag 887
func (m DerivativeSecurityListRequest) GetNoUnderlyingStips() (NoUnderlyingStipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingStipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingStrikeCurrency gets UnderlyingStrikeCurrency, Tag 941
func (m DerivativeSecurityListRequest) GetUnderlyingStrikeCurrency() (f field.UnderlyingStrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingAllocationPercent gets UnderlyingAllocationPercent, Tag 972
func (m DerivativeSecurityListRequest) GetUnderlyingAllocationPercent() (f field.UnderlyingAllocationPercentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCashAmount gets UnderlyingCashAmount, Tag 973
func (m DerivativeSecurityListRequest) GetUnderlyingCashAmount() (f field.UnderlyingCashAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCashType gets UnderlyingCashType, Tag 974
func (m DerivativeSecurityListRequest) GetUnderlyingCashType() (f field.UnderlyingCashTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSettlementType gets UnderlyingSettlementType, Tag 975
func (m DerivativeSecurityListRequest) GetUnderlyingSettlementType() (f field.UnderlyingSettlementTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingUnitOfMeasure gets UnderlyingUnitOfMeasure, Tag 998
func (m DerivativeSecurityListRequest) GetUnderlyingUnitOfMeasure() (f field.UnderlyingUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingTimeUnit gets UnderlyingTimeUnit, Tag 1000
func (m DerivativeSecurityListRequest) GetUnderlyingTimeUnit() (f field.UnderlyingTimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCapValue gets UnderlyingCapValue, Tag 1038
func (m DerivativeSecurityListRequest) GetUnderlyingCapValue() (f field.UnderlyingCapValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSettlMethod gets UnderlyingSettlMethod, Tag 1039
func (m DerivativeSecurityListRequest) GetUnderlyingSettlMethod() (f field.UnderlyingSettlMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingAdjustedQuantity gets UnderlyingAdjustedQuantity, Tag 1044
func (m DerivativeSecurityListRequest) GetUnderlyingAdjustedQuantity() (f field.UnderlyingAdjustedQuantityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFXRate gets UnderlyingFXRate, Tag 1045
func (m DerivativeSecurityListRequest) GetUnderlyingFXRate() (f field.UnderlyingFXRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFXRateCalc gets UnderlyingFXRateCalc, Tag 1046
func (m DerivativeSecurityListRequest) GetUnderlyingFXRateCalc() (f field.UnderlyingFXRateCalcField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUndlyInstrumentParties gets NoUndlyInstrumentParties, Tag 1058
func (m DerivativeSecurityListRequest) GetNoUndlyInstrumentParties() (NoUndlyInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingMaturityTime gets UnderlyingMaturityTime, Tag 1213
func (m DerivativeSecurityListRequest) GetUnderlyingMaturityTime() (f field.UnderlyingMaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSymbol gets DerivativeSymbol, Tag 1214
func (m DerivativeSecurityListRequest) GetDerivativeSymbol() (f field.DerivativeSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSymbolSfx gets DerivativeSymbolSfx, Tag 1215
func (m DerivativeSecurityListRequest) GetDerivativeSymbolSfx() (f field.DerivativeSymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSecurityID gets DerivativeSecurityID, Tag 1216
func (m DerivativeSecurityListRequest) GetDerivativeSecurityID() (f field.DerivativeSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSecurityIDSource gets DerivativeSecurityIDSource, Tag 1217
func (m DerivativeSecurityListRequest) GetDerivativeSecurityIDSource() (f field.DerivativeSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoDerivativeSecurityAltID gets NoDerivativeSecurityAltID, Tag 1218
func (m DerivativeSecurityListRequest) GetNoDerivativeSecurityAltID() (NoDerivativeSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoDerivativeSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDerivativeOptPayAmount gets DerivativeOptPayAmount, Tag 1225
func (m DerivativeSecurityListRequest) GetDerivativeOptPayAmount() (f field.DerivativeOptPayAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeProductComplex gets DerivativeProductComplex, Tag 1228
func (m DerivativeSecurityListRequest) GetDerivativeProductComplex() (f field.DerivativeProductComplexField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivFlexProductEligibilityIndicator gets DerivFlexProductEligibilityIndicator, Tag 1243
func (m DerivativeSecurityListRequest) GetDerivFlexProductEligibilityIndicator() (f field.DerivFlexProductEligibilityIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeProduct gets DerivativeProduct, Tag 1246
func (m DerivativeSecurityListRequest) GetDerivativeProduct() (f field.DerivativeProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSecurityGroup gets DerivativeSecurityGroup, Tag 1247
func (m DerivativeSecurityListRequest) GetDerivativeSecurityGroup() (f field.DerivativeSecurityGroupField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeCFICode gets DerivativeCFICode, Tag 1248
func (m DerivativeSecurityListRequest) GetDerivativeCFICode() (f field.DerivativeCFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSecurityType gets DerivativeSecurityType, Tag 1249
func (m DerivativeSecurityListRequest) GetDerivativeSecurityType() (f field.DerivativeSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSecuritySubType gets DerivativeSecuritySubType, Tag 1250
func (m DerivativeSecurityListRequest) GetDerivativeSecuritySubType() (f field.DerivativeSecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeMaturityMonthYear gets DerivativeMaturityMonthYear, Tag 1251
func (m DerivativeSecurityListRequest) GetDerivativeMaturityMonthYear() (f field.DerivativeMaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeMaturityDate gets DerivativeMaturityDate, Tag 1252
func (m DerivativeSecurityListRequest) GetDerivativeMaturityDate() (f field.DerivativeMaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeMaturityTime gets DerivativeMaturityTime, Tag 1253
func (m DerivativeSecurityListRequest) GetDerivativeMaturityTime() (f field.DerivativeMaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSettleOnOpenFlag gets DerivativeSettleOnOpenFlag, Tag 1254
func (m DerivativeSecurityListRequest) GetDerivativeSettleOnOpenFlag() (f field.DerivativeSettleOnOpenFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeInstrmtAssignmentMethod gets DerivativeInstrmtAssignmentMethod, Tag 1255
func (m DerivativeSecurityListRequest) GetDerivativeInstrmtAssignmentMethod() (f field.DerivativeInstrmtAssignmentMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSecurityStatus gets DerivativeSecurityStatus, Tag 1256
func (m DerivativeSecurityListRequest) GetDerivativeSecurityStatus() (f field.DerivativeSecurityStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeInstrRegistry gets DerivativeInstrRegistry, Tag 1257
func (m DerivativeSecurityListRequest) GetDerivativeInstrRegistry() (f field.DerivativeInstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeCountryOfIssue gets DerivativeCountryOfIssue, Tag 1258
func (m DerivativeSecurityListRequest) GetDerivativeCountryOfIssue() (f field.DerivativeCountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeStateOrProvinceOfIssue gets DerivativeStateOrProvinceOfIssue, Tag 1259
func (m DerivativeSecurityListRequest) GetDerivativeStateOrProvinceOfIssue() (f field.DerivativeStateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeLocaleOfIssue gets DerivativeLocaleOfIssue, Tag 1260
func (m DerivativeSecurityListRequest) GetDerivativeLocaleOfIssue() (f field.DerivativeLocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeStrikePrice gets DerivativeStrikePrice, Tag 1261
func (m DerivativeSecurityListRequest) GetDerivativeStrikePrice() (f field.DerivativeStrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeStrikeCurrency gets DerivativeStrikeCurrency, Tag 1262
func (m DerivativeSecurityListRequest) GetDerivativeStrikeCurrency() (f field.DerivativeStrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeStrikeMultiplier gets DerivativeStrikeMultiplier, Tag 1263
func (m DerivativeSecurityListRequest) GetDerivativeStrikeMultiplier() (f field.DerivativeStrikeMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeStrikeValue gets DerivativeStrikeValue, Tag 1264
func (m DerivativeSecurityListRequest) GetDerivativeStrikeValue() (f field.DerivativeStrikeValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeOptAttribute gets DerivativeOptAttribute, Tag 1265
func (m DerivativeSecurityListRequest) GetDerivativeOptAttribute() (f field.DerivativeOptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeContractMultiplier gets DerivativeContractMultiplier, Tag 1266
func (m DerivativeSecurityListRequest) GetDerivativeContractMultiplier() (f field.DerivativeContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeMinPriceIncrement gets DerivativeMinPriceIncrement, Tag 1267
func (m DerivativeSecurityListRequest) GetDerivativeMinPriceIncrement() (f field.DerivativeMinPriceIncrementField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeMinPriceIncrementAmount gets DerivativeMinPriceIncrementAmount, Tag 1268
func (m DerivativeSecurityListRequest) GetDerivativeMinPriceIncrementAmount() (f field.DerivativeMinPriceIncrementAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeUnitOfMeasure gets DerivativeUnitOfMeasure, Tag 1269
func (m DerivativeSecurityListRequest) GetDerivativeUnitOfMeasure() (f field.DerivativeUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeUnitOfMeasureQty gets DerivativeUnitOfMeasureQty, Tag 1270
func (m DerivativeSecurityListRequest) GetDerivativeUnitOfMeasureQty() (f field.DerivativeUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeTimeUnit gets DerivativeTimeUnit, Tag 1271
func (m DerivativeSecurityListRequest) GetDerivativeTimeUnit() (f field.DerivativeTimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSecurityExchange gets DerivativeSecurityExchange, Tag 1272
func (m DerivativeSecurityListRequest) GetDerivativeSecurityExchange() (f field.DerivativeSecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativePositionLimit gets DerivativePositionLimit, Tag 1273
func (m DerivativeSecurityListRequest) GetDerivativePositionLimit() (f field.DerivativePositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeNTPositionLimit gets DerivativeNTPositionLimit, Tag 1274
func (m DerivativeSecurityListRequest) GetDerivativeNTPositionLimit() (f field.DerivativeNTPositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeIssuer gets DerivativeIssuer, Tag 1275
func (m DerivativeSecurityListRequest) GetDerivativeIssuer() (f field.DerivativeIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeIssueDate gets DerivativeIssueDate, Tag 1276
func (m DerivativeSecurityListRequest) GetDerivativeIssueDate() (f field.DerivativeIssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeEncodedIssuerLen gets DerivativeEncodedIssuerLen, Tag 1277
func (m DerivativeSecurityListRequest) GetDerivativeEncodedIssuerLen() (f field.DerivativeEncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeEncodedIssuer gets DerivativeEncodedIssuer, Tag 1278
func (m DerivativeSecurityListRequest) GetDerivativeEncodedIssuer() (f field.DerivativeEncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSecurityDesc gets DerivativeSecurityDesc, Tag 1279
func (m DerivativeSecurityListRequest) GetDerivativeSecurityDesc() (f field.DerivativeSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeEncodedSecurityDescLen gets DerivativeEncodedSecurityDescLen, Tag 1280
func (m DerivativeSecurityListRequest) GetDerivativeEncodedSecurityDescLen() (f field.DerivativeEncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeEncodedSecurityDesc gets DerivativeEncodedSecurityDesc, Tag 1281
func (m DerivativeSecurityListRequest) GetDerivativeEncodedSecurityDesc() (f field.DerivativeEncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSecurityXMLLen gets DerivativeSecurityXMLLen, Tag 1282
func (m DerivativeSecurityListRequest) GetDerivativeSecurityXMLLen() (f field.DerivativeSecurityXMLLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSecurityXML gets DerivativeSecurityXML, Tag 1283
func (m DerivativeSecurityListRequest) GetDerivativeSecurityXML() (f field.DerivativeSecurityXMLField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSecurityXMLSchema gets DerivativeSecurityXMLSchema, Tag 1284
func (m DerivativeSecurityListRequest) GetDerivativeSecurityXMLSchema() (f field.DerivativeSecurityXMLSchemaField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeContractSettlMonth gets DerivativeContractSettlMonth, Tag 1285
func (m DerivativeSecurityListRequest) GetDerivativeContractSettlMonth() (f field.DerivativeContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoDerivativeEvents gets NoDerivativeEvents, Tag 1286
func (m DerivativeSecurityListRequest) GetNoDerivativeEvents() (NoDerivativeEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoDerivativeEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoDerivativeInstrumentParties gets NoDerivativeInstrumentParties, Tag 1292
func (m DerivativeSecurityListRequest) GetNoDerivativeInstrumentParties() (NoDerivativeInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoDerivativeInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDerivativeExerciseStyle gets DerivativeExerciseStyle, Tag 1299
func (m DerivativeSecurityListRequest) GetDerivativeExerciseStyle() (f field.DerivativeExerciseStyleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMarketSegmentID gets MarketSegmentID, Tag 1300
func (m DerivativeSecurityListRequest) GetMarketSegmentID() (f field.MarketSegmentIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMarketID gets MarketID, Tag 1301
func (m DerivativeSecurityListRequest) GetMarketID() (f field.MarketIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativePriceUnitOfMeasure gets DerivativePriceUnitOfMeasure, Tag 1315
func (m DerivativeSecurityListRequest) GetDerivativePriceUnitOfMeasure() (f field.DerivativePriceUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativePriceUnitOfMeasureQty gets DerivativePriceUnitOfMeasureQty, Tag 1316
func (m DerivativeSecurityListRequest) GetDerivativePriceUnitOfMeasureQty() (f field.DerivativePriceUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSettlMethod gets DerivativeSettlMethod, Tag 1317
func (m DerivativeSecurityListRequest) GetDerivativeSettlMethod() (f field.DerivativeSettlMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativePriceQuoteMethod gets DerivativePriceQuoteMethod, Tag 1318
func (m DerivativeSecurityListRequest) GetDerivativePriceQuoteMethod() (f field.DerivativePriceQuoteMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeFuturesValuationMethod gets DerivativeFuturesValuationMethod, Tag 1319
func (m DerivativeSecurityListRequest) GetDerivativeFuturesValuationMethod() (f field.DerivativeFuturesValuationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeListMethod gets DerivativeListMethod, Tag 1320
func (m DerivativeSecurityListRequest) GetDerivativeListMethod() (f field.DerivativeListMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeCapPrice gets DerivativeCapPrice, Tag 1321
func (m DerivativeSecurityListRequest) GetDerivativeCapPrice() (f field.DerivativeCapPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeFloorPrice gets DerivativeFloorPrice, Tag 1322
func (m DerivativeSecurityListRequest) GetDerivativeFloorPrice() (f field.DerivativeFloorPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativePutOrCall gets DerivativePutOrCall, Tag 1323
func (m DerivativeSecurityListRequest) GetDerivativePutOrCall() (f field.DerivativePutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingExerciseStyle gets UnderlyingExerciseStyle, Tag 1419
func (m DerivativeSecurityListRequest) GetUnderlyingExerciseStyle() (f field.UnderlyingExerciseStyleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingUnitOfMeasureQty gets UnderlyingUnitOfMeasureQty, Tag 1423
func (m DerivativeSecurityListRequest) GetUnderlyingUnitOfMeasureQty() (f field.UnderlyingUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPriceUnitOfMeasure gets UnderlyingPriceUnitOfMeasure, Tag 1424
func (m DerivativeSecurityListRequest) GetUnderlyingPriceUnitOfMeasure() (f field.UnderlyingPriceUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPriceUnitOfMeasureQty gets UnderlyingPriceUnitOfMeasureQty, Tag 1425
func (m DerivativeSecurityListRequest) GetUnderlyingPriceUnitOfMeasureQty() (f field.UnderlyingPriceUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasCurrency returns true if Currency is present, Tag 15
func (m DerivativeSecurityListRequest) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasText returns true if Text is present, Tag 58
func (m DerivativeSecurityListRequest) HasText() bool {
	return m.Has(tag.Text)
}

//HasUnderlyingCouponPaymentDate returns true if UnderlyingCouponPaymentDate is present, Tag 241
func (m DerivativeSecurityListRequest) HasUnderlyingCouponPaymentDate() bool {
	return m.Has(tag.UnderlyingCouponPaymentDate)
}

//HasUnderlyingIssueDate returns true if UnderlyingIssueDate is present, Tag 242
func (m DerivativeSecurityListRequest) HasUnderlyingIssueDate() bool {
	return m.Has(tag.UnderlyingIssueDate)
}

//HasUnderlyingRepoCollateralSecurityType returns true if UnderlyingRepoCollateralSecurityType is present, Tag 243
func (m DerivativeSecurityListRequest) HasUnderlyingRepoCollateralSecurityType() bool {
	return m.Has(tag.UnderlyingRepoCollateralSecurityType)
}

//HasUnderlyingRepurchaseTerm returns true if UnderlyingRepurchaseTerm is present, Tag 244
func (m DerivativeSecurityListRequest) HasUnderlyingRepurchaseTerm() bool {
	return m.Has(tag.UnderlyingRepurchaseTerm)
}

//HasUnderlyingRepurchaseRate returns true if UnderlyingRepurchaseRate is present, Tag 245
func (m DerivativeSecurityListRequest) HasUnderlyingRepurchaseRate() bool {
	return m.Has(tag.UnderlyingRepurchaseRate)
}

//HasUnderlyingFactor returns true if UnderlyingFactor is present, Tag 246
func (m DerivativeSecurityListRequest) HasUnderlyingFactor() bool {
	return m.Has(tag.UnderlyingFactor)
}

//HasUnderlyingRedemptionDate returns true if UnderlyingRedemptionDate is present, Tag 247
func (m DerivativeSecurityListRequest) HasUnderlyingRedemptionDate() bool {
	return m.Has(tag.UnderlyingRedemptionDate)
}

//HasUnderlyingCreditRating returns true if UnderlyingCreditRating is present, Tag 256
func (m DerivativeSecurityListRequest) HasUnderlyingCreditRating() bool {
	return m.Has(tag.UnderlyingCreditRating)
}

//HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263
func (m DerivativeSecurityListRequest) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

//HasUnderlyingSecurityIDSource returns true if UnderlyingSecurityIDSource is present, Tag 305
func (m DerivativeSecurityListRequest) HasUnderlyingSecurityIDSource() bool {
	return m.Has(tag.UnderlyingSecurityIDSource)
}

//HasUnderlyingIssuer returns true if UnderlyingIssuer is present, Tag 306
func (m DerivativeSecurityListRequest) HasUnderlyingIssuer() bool {
	return m.Has(tag.UnderlyingIssuer)
}

//HasUnderlyingSecurityDesc returns true if UnderlyingSecurityDesc is present, Tag 307
func (m DerivativeSecurityListRequest) HasUnderlyingSecurityDesc() bool {
	return m.Has(tag.UnderlyingSecurityDesc)
}

//HasUnderlyingSecurityExchange returns true if UnderlyingSecurityExchange is present, Tag 308
func (m DerivativeSecurityListRequest) HasUnderlyingSecurityExchange() bool {
	return m.Has(tag.UnderlyingSecurityExchange)
}

//HasUnderlyingSecurityID returns true if UnderlyingSecurityID is present, Tag 309
func (m DerivativeSecurityListRequest) HasUnderlyingSecurityID() bool {
	return m.Has(tag.UnderlyingSecurityID)
}

//HasUnderlyingSecurityType returns true if UnderlyingSecurityType is present, Tag 310
func (m DerivativeSecurityListRequest) HasUnderlyingSecurityType() bool {
	return m.Has(tag.UnderlyingSecurityType)
}

//HasUnderlyingSymbol returns true if UnderlyingSymbol is present, Tag 311
func (m DerivativeSecurityListRequest) HasUnderlyingSymbol() bool {
	return m.Has(tag.UnderlyingSymbol)
}

//HasUnderlyingSymbolSfx returns true if UnderlyingSymbolSfx is present, Tag 312
func (m DerivativeSecurityListRequest) HasUnderlyingSymbolSfx() bool {
	return m.Has(tag.UnderlyingSymbolSfx)
}

//HasUnderlyingMaturityMonthYear returns true if UnderlyingMaturityMonthYear is present, Tag 313
func (m DerivativeSecurityListRequest) HasUnderlyingMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingMaturityMonthYear)
}

//HasUnderlyingPutOrCall returns true if UnderlyingPutOrCall is present, Tag 315
func (m DerivativeSecurityListRequest) HasUnderlyingPutOrCall() bool {
	return m.Has(tag.UnderlyingPutOrCall)
}

//HasUnderlyingStrikePrice returns true if UnderlyingStrikePrice is present, Tag 316
func (m DerivativeSecurityListRequest) HasUnderlyingStrikePrice() bool {
	return m.Has(tag.UnderlyingStrikePrice)
}

//HasUnderlyingOptAttribute returns true if UnderlyingOptAttribute is present, Tag 317
func (m DerivativeSecurityListRequest) HasUnderlyingOptAttribute() bool {
	return m.Has(tag.UnderlyingOptAttribute)
}

//HasUnderlyingCurrency returns true if UnderlyingCurrency is present, Tag 318
func (m DerivativeSecurityListRequest) HasUnderlyingCurrency() bool {
	return m.Has(tag.UnderlyingCurrency)
}

//HasSecurityReqID returns true if SecurityReqID is present, Tag 320
func (m DerivativeSecurityListRequest) HasSecurityReqID() bool {
	return m.Has(tag.SecurityReqID)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m DerivativeSecurityListRequest) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m DerivativeSecurityListRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m DerivativeSecurityListRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasEncodedUnderlyingIssuerLen returns true if EncodedUnderlyingIssuerLen is present, Tag 362
func (m DerivativeSecurityListRequest) HasEncodedUnderlyingIssuerLen() bool {
	return m.Has(tag.EncodedUnderlyingIssuerLen)
}

//HasEncodedUnderlyingIssuer returns true if EncodedUnderlyingIssuer is present, Tag 363
func (m DerivativeSecurityListRequest) HasEncodedUnderlyingIssuer() bool {
	return m.Has(tag.EncodedUnderlyingIssuer)
}

//HasEncodedUnderlyingSecurityDescLen returns true if EncodedUnderlyingSecurityDescLen is present, Tag 364
func (m DerivativeSecurityListRequest) HasEncodedUnderlyingSecurityDescLen() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDescLen)
}

//HasEncodedUnderlyingSecurityDesc returns true if EncodedUnderlyingSecurityDesc is present, Tag 365
func (m DerivativeSecurityListRequest) HasEncodedUnderlyingSecurityDesc() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDesc)
}

//HasUnderlyingCouponRate returns true if UnderlyingCouponRate is present, Tag 435
func (m DerivativeSecurityListRequest) HasUnderlyingCouponRate() bool {
	return m.Has(tag.UnderlyingCouponRate)
}

//HasUnderlyingContractMultiplier returns true if UnderlyingContractMultiplier is present, Tag 436
func (m DerivativeSecurityListRequest) HasUnderlyingContractMultiplier() bool {
	return m.Has(tag.UnderlyingContractMultiplier)
}

//HasNoUnderlyingSecurityAltID returns true if NoUnderlyingSecurityAltID is present, Tag 457
func (m DerivativeSecurityListRequest) HasNoUnderlyingSecurityAltID() bool {
	return m.Has(tag.NoUnderlyingSecurityAltID)
}

//HasUnderlyingProduct returns true if UnderlyingProduct is present, Tag 462
func (m DerivativeSecurityListRequest) HasUnderlyingProduct() bool {
	return m.Has(tag.UnderlyingProduct)
}

//HasUnderlyingCFICode returns true if UnderlyingCFICode is present, Tag 463
func (m DerivativeSecurityListRequest) HasUnderlyingCFICode() bool {
	return m.Has(tag.UnderlyingCFICode)
}

//HasUnderlyingMaturityDate returns true if UnderlyingMaturityDate is present, Tag 542
func (m DerivativeSecurityListRequest) HasUnderlyingMaturityDate() bool {
	return m.Has(tag.UnderlyingMaturityDate)
}

//HasSecurityListRequestType returns true if SecurityListRequestType is present, Tag 559
func (m DerivativeSecurityListRequest) HasSecurityListRequestType() bool {
	return m.Has(tag.SecurityListRequestType)
}

//HasUnderlyingCountryOfIssue returns true if UnderlyingCountryOfIssue is present, Tag 592
func (m DerivativeSecurityListRequest) HasUnderlyingCountryOfIssue() bool {
	return m.Has(tag.UnderlyingCountryOfIssue)
}

//HasUnderlyingStateOrProvinceOfIssue returns true if UnderlyingStateOrProvinceOfIssue is present, Tag 593
func (m DerivativeSecurityListRequest) HasUnderlyingStateOrProvinceOfIssue() bool {
	return m.Has(tag.UnderlyingStateOrProvinceOfIssue)
}

//HasUnderlyingLocaleOfIssue returns true if UnderlyingLocaleOfIssue is present, Tag 594
func (m DerivativeSecurityListRequest) HasUnderlyingLocaleOfIssue() bool {
	return m.Has(tag.UnderlyingLocaleOfIssue)
}

//HasUnderlyingInstrRegistry returns true if UnderlyingInstrRegistry is present, Tag 595
func (m DerivativeSecurityListRequest) HasUnderlyingInstrRegistry() bool {
	return m.Has(tag.UnderlyingInstrRegistry)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m DerivativeSecurityListRequest) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m DerivativeSecurityListRequest) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

//HasUnderlyingSecuritySubType returns true if UnderlyingSecuritySubType is present, Tag 763
func (m DerivativeSecurityListRequest) HasUnderlyingSecuritySubType() bool {
	return m.Has(tag.UnderlyingSecuritySubType)
}

//HasUnderlyingPx returns true if UnderlyingPx is present, Tag 810
func (m DerivativeSecurityListRequest) HasUnderlyingPx() bool {
	return m.Has(tag.UnderlyingPx)
}

//HasUnderlyingCPProgram returns true if UnderlyingCPProgram is present, Tag 877
func (m DerivativeSecurityListRequest) HasUnderlyingCPProgram() bool {
	return m.Has(tag.UnderlyingCPProgram)
}

//HasUnderlyingCPRegType returns true if UnderlyingCPRegType is present, Tag 878
func (m DerivativeSecurityListRequest) HasUnderlyingCPRegType() bool {
	return m.Has(tag.UnderlyingCPRegType)
}

//HasUnderlyingQty returns true if UnderlyingQty is present, Tag 879
func (m DerivativeSecurityListRequest) HasUnderlyingQty() bool {
	return m.Has(tag.UnderlyingQty)
}

//HasUnderlyingDirtyPrice returns true if UnderlyingDirtyPrice is present, Tag 882
func (m DerivativeSecurityListRequest) HasUnderlyingDirtyPrice() bool {
	return m.Has(tag.UnderlyingDirtyPrice)
}

//HasUnderlyingEndPrice returns true if UnderlyingEndPrice is present, Tag 883
func (m DerivativeSecurityListRequest) HasUnderlyingEndPrice() bool {
	return m.Has(tag.UnderlyingEndPrice)
}

//HasUnderlyingStartValue returns true if UnderlyingStartValue is present, Tag 884
func (m DerivativeSecurityListRequest) HasUnderlyingStartValue() bool {
	return m.Has(tag.UnderlyingStartValue)
}

//HasUnderlyingCurrentValue returns true if UnderlyingCurrentValue is present, Tag 885
func (m DerivativeSecurityListRequest) HasUnderlyingCurrentValue() bool {
	return m.Has(tag.UnderlyingCurrentValue)
}

//HasUnderlyingEndValue returns true if UnderlyingEndValue is present, Tag 886
func (m DerivativeSecurityListRequest) HasUnderlyingEndValue() bool {
	return m.Has(tag.UnderlyingEndValue)
}

//HasNoUnderlyingStips returns true if NoUnderlyingStips is present, Tag 887
func (m DerivativeSecurityListRequest) HasNoUnderlyingStips() bool {
	return m.Has(tag.NoUnderlyingStips)
}

//HasUnderlyingStrikeCurrency returns true if UnderlyingStrikeCurrency is present, Tag 941
func (m DerivativeSecurityListRequest) HasUnderlyingStrikeCurrency() bool {
	return m.Has(tag.UnderlyingStrikeCurrency)
}

//HasUnderlyingAllocationPercent returns true if UnderlyingAllocationPercent is present, Tag 972
func (m DerivativeSecurityListRequest) HasUnderlyingAllocationPercent() bool {
	return m.Has(tag.UnderlyingAllocationPercent)
}

//HasUnderlyingCashAmount returns true if UnderlyingCashAmount is present, Tag 973
func (m DerivativeSecurityListRequest) HasUnderlyingCashAmount() bool {
	return m.Has(tag.UnderlyingCashAmount)
}

//HasUnderlyingCashType returns true if UnderlyingCashType is present, Tag 974
func (m DerivativeSecurityListRequest) HasUnderlyingCashType() bool {
	return m.Has(tag.UnderlyingCashType)
}

//HasUnderlyingSettlementType returns true if UnderlyingSettlementType is present, Tag 975
func (m DerivativeSecurityListRequest) HasUnderlyingSettlementType() bool {
	return m.Has(tag.UnderlyingSettlementType)
}

//HasUnderlyingUnitOfMeasure returns true if UnderlyingUnitOfMeasure is present, Tag 998
func (m DerivativeSecurityListRequest) HasUnderlyingUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingUnitOfMeasure)
}

//HasUnderlyingTimeUnit returns true if UnderlyingTimeUnit is present, Tag 1000
func (m DerivativeSecurityListRequest) HasUnderlyingTimeUnit() bool {
	return m.Has(tag.UnderlyingTimeUnit)
}

//HasUnderlyingCapValue returns true if UnderlyingCapValue is present, Tag 1038
func (m DerivativeSecurityListRequest) HasUnderlyingCapValue() bool {
	return m.Has(tag.UnderlyingCapValue)
}

//HasUnderlyingSettlMethod returns true if UnderlyingSettlMethod is present, Tag 1039
func (m DerivativeSecurityListRequest) HasUnderlyingSettlMethod() bool {
	return m.Has(tag.UnderlyingSettlMethod)
}

//HasUnderlyingAdjustedQuantity returns true if UnderlyingAdjustedQuantity is present, Tag 1044
func (m DerivativeSecurityListRequest) HasUnderlyingAdjustedQuantity() bool {
	return m.Has(tag.UnderlyingAdjustedQuantity)
}

//HasUnderlyingFXRate returns true if UnderlyingFXRate is present, Tag 1045
func (m DerivativeSecurityListRequest) HasUnderlyingFXRate() bool {
	return m.Has(tag.UnderlyingFXRate)
}

//HasUnderlyingFXRateCalc returns true if UnderlyingFXRateCalc is present, Tag 1046
func (m DerivativeSecurityListRequest) HasUnderlyingFXRateCalc() bool {
	return m.Has(tag.UnderlyingFXRateCalc)
}

//HasNoUndlyInstrumentParties returns true if NoUndlyInstrumentParties is present, Tag 1058
func (m DerivativeSecurityListRequest) HasNoUndlyInstrumentParties() bool {
	return m.Has(tag.NoUndlyInstrumentParties)
}

//HasUnderlyingMaturityTime returns true if UnderlyingMaturityTime is present, Tag 1213
func (m DerivativeSecurityListRequest) HasUnderlyingMaturityTime() bool {
	return m.Has(tag.UnderlyingMaturityTime)
}

//HasDerivativeSymbol returns true if DerivativeSymbol is present, Tag 1214
func (m DerivativeSecurityListRequest) HasDerivativeSymbol() bool {
	return m.Has(tag.DerivativeSymbol)
}

//HasDerivativeSymbolSfx returns true if DerivativeSymbolSfx is present, Tag 1215
func (m DerivativeSecurityListRequest) HasDerivativeSymbolSfx() bool {
	return m.Has(tag.DerivativeSymbolSfx)
}

//HasDerivativeSecurityID returns true if DerivativeSecurityID is present, Tag 1216
func (m DerivativeSecurityListRequest) HasDerivativeSecurityID() bool {
	return m.Has(tag.DerivativeSecurityID)
}

//HasDerivativeSecurityIDSource returns true if DerivativeSecurityIDSource is present, Tag 1217
func (m DerivativeSecurityListRequest) HasDerivativeSecurityIDSource() bool {
	return m.Has(tag.DerivativeSecurityIDSource)
}

//HasNoDerivativeSecurityAltID returns true if NoDerivativeSecurityAltID is present, Tag 1218
func (m DerivativeSecurityListRequest) HasNoDerivativeSecurityAltID() bool {
	return m.Has(tag.NoDerivativeSecurityAltID)
}

//HasDerivativeOptPayAmount returns true if DerivativeOptPayAmount is present, Tag 1225
func (m DerivativeSecurityListRequest) HasDerivativeOptPayAmount() bool {
	return m.Has(tag.DerivativeOptPayAmount)
}

//HasDerivativeProductComplex returns true if DerivativeProductComplex is present, Tag 1228
func (m DerivativeSecurityListRequest) HasDerivativeProductComplex() bool {
	return m.Has(tag.DerivativeProductComplex)
}

//HasDerivFlexProductEligibilityIndicator returns true if DerivFlexProductEligibilityIndicator is present, Tag 1243
func (m DerivativeSecurityListRequest) HasDerivFlexProductEligibilityIndicator() bool {
	return m.Has(tag.DerivFlexProductEligibilityIndicator)
}

//HasDerivativeProduct returns true if DerivativeProduct is present, Tag 1246
func (m DerivativeSecurityListRequest) HasDerivativeProduct() bool {
	return m.Has(tag.DerivativeProduct)
}

//HasDerivativeSecurityGroup returns true if DerivativeSecurityGroup is present, Tag 1247
func (m DerivativeSecurityListRequest) HasDerivativeSecurityGroup() bool {
	return m.Has(tag.DerivativeSecurityGroup)
}

//HasDerivativeCFICode returns true if DerivativeCFICode is present, Tag 1248
func (m DerivativeSecurityListRequest) HasDerivativeCFICode() bool {
	return m.Has(tag.DerivativeCFICode)
}

//HasDerivativeSecurityType returns true if DerivativeSecurityType is present, Tag 1249
func (m DerivativeSecurityListRequest) HasDerivativeSecurityType() bool {
	return m.Has(tag.DerivativeSecurityType)
}

//HasDerivativeSecuritySubType returns true if DerivativeSecuritySubType is present, Tag 1250
func (m DerivativeSecurityListRequest) HasDerivativeSecuritySubType() bool {
	return m.Has(tag.DerivativeSecuritySubType)
}

//HasDerivativeMaturityMonthYear returns true if DerivativeMaturityMonthYear is present, Tag 1251
func (m DerivativeSecurityListRequest) HasDerivativeMaturityMonthYear() bool {
	return m.Has(tag.DerivativeMaturityMonthYear)
}

//HasDerivativeMaturityDate returns true if DerivativeMaturityDate is present, Tag 1252
func (m DerivativeSecurityListRequest) HasDerivativeMaturityDate() bool {
	return m.Has(tag.DerivativeMaturityDate)
}

//HasDerivativeMaturityTime returns true if DerivativeMaturityTime is present, Tag 1253
func (m DerivativeSecurityListRequest) HasDerivativeMaturityTime() bool {
	return m.Has(tag.DerivativeMaturityTime)
}

//HasDerivativeSettleOnOpenFlag returns true if DerivativeSettleOnOpenFlag is present, Tag 1254
func (m DerivativeSecurityListRequest) HasDerivativeSettleOnOpenFlag() bool {
	return m.Has(tag.DerivativeSettleOnOpenFlag)
}

//HasDerivativeInstrmtAssignmentMethod returns true if DerivativeInstrmtAssignmentMethod is present, Tag 1255
func (m DerivativeSecurityListRequest) HasDerivativeInstrmtAssignmentMethod() bool {
	return m.Has(tag.DerivativeInstrmtAssignmentMethod)
}

//HasDerivativeSecurityStatus returns true if DerivativeSecurityStatus is present, Tag 1256
func (m DerivativeSecurityListRequest) HasDerivativeSecurityStatus() bool {
	return m.Has(tag.DerivativeSecurityStatus)
}

//HasDerivativeInstrRegistry returns true if DerivativeInstrRegistry is present, Tag 1257
func (m DerivativeSecurityListRequest) HasDerivativeInstrRegistry() bool {
	return m.Has(tag.DerivativeInstrRegistry)
}

//HasDerivativeCountryOfIssue returns true if DerivativeCountryOfIssue is present, Tag 1258
func (m DerivativeSecurityListRequest) HasDerivativeCountryOfIssue() bool {
	return m.Has(tag.DerivativeCountryOfIssue)
}

//HasDerivativeStateOrProvinceOfIssue returns true if DerivativeStateOrProvinceOfIssue is present, Tag 1259
func (m DerivativeSecurityListRequest) HasDerivativeStateOrProvinceOfIssue() bool {
	return m.Has(tag.DerivativeStateOrProvinceOfIssue)
}

//HasDerivativeLocaleOfIssue returns true if DerivativeLocaleOfIssue is present, Tag 1260
func (m DerivativeSecurityListRequest) HasDerivativeLocaleOfIssue() bool {
	return m.Has(tag.DerivativeLocaleOfIssue)
}

//HasDerivativeStrikePrice returns true if DerivativeStrikePrice is present, Tag 1261
func (m DerivativeSecurityListRequest) HasDerivativeStrikePrice() bool {
	return m.Has(tag.DerivativeStrikePrice)
}

//HasDerivativeStrikeCurrency returns true if DerivativeStrikeCurrency is present, Tag 1262
func (m DerivativeSecurityListRequest) HasDerivativeStrikeCurrency() bool {
	return m.Has(tag.DerivativeStrikeCurrency)
}

//HasDerivativeStrikeMultiplier returns true if DerivativeStrikeMultiplier is present, Tag 1263
func (m DerivativeSecurityListRequest) HasDerivativeStrikeMultiplier() bool {
	return m.Has(tag.DerivativeStrikeMultiplier)
}

//HasDerivativeStrikeValue returns true if DerivativeStrikeValue is present, Tag 1264
func (m DerivativeSecurityListRequest) HasDerivativeStrikeValue() bool {
	return m.Has(tag.DerivativeStrikeValue)
}

//HasDerivativeOptAttribute returns true if DerivativeOptAttribute is present, Tag 1265
func (m DerivativeSecurityListRequest) HasDerivativeOptAttribute() bool {
	return m.Has(tag.DerivativeOptAttribute)
}

//HasDerivativeContractMultiplier returns true if DerivativeContractMultiplier is present, Tag 1266
func (m DerivativeSecurityListRequest) HasDerivativeContractMultiplier() bool {
	return m.Has(tag.DerivativeContractMultiplier)
}

//HasDerivativeMinPriceIncrement returns true if DerivativeMinPriceIncrement is present, Tag 1267
func (m DerivativeSecurityListRequest) HasDerivativeMinPriceIncrement() bool {
	return m.Has(tag.DerivativeMinPriceIncrement)
}

//HasDerivativeMinPriceIncrementAmount returns true if DerivativeMinPriceIncrementAmount is present, Tag 1268
func (m DerivativeSecurityListRequest) HasDerivativeMinPriceIncrementAmount() bool {
	return m.Has(tag.DerivativeMinPriceIncrementAmount)
}

//HasDerivativeUnitOfMeasure returns true if DerivativeUnitOfMeasure is present, Tag 1269
func (m DerivativeSecurityListRequest) HasDerivativeUnitOfMeasure() bool {
	return m.Has(tag.DerivativeUnitOfMeasure)
}

//HasDerivativeUnitOfMeasureQty returns true if DerivativeUnitOfMeasureQty is present, Tag 1270
func (m DerivativeSecurityListRequest) HasDerivativeUnitOfMeasureQty() bool {
	return m.Has(tag.DerivativeUnitOfMeasureQty)
}

//HasDerivativeTimeUnit returns true if DerivativeTimeUnit is present, Tag 1271
func (m DerivativeSecurityListRequest) HasDerivativeTimeUnit() bool {
	return m.Has(tag.DerivativeTimeUnit)
}

//HasDerivativeSecurityExchange returns true if DerivativeSecurityExchange is present, Tag 1272
func (m DerivativeSecurityListRequest) HasDerivativeSecurityExchange() bool {
	return m.Has(tag.DerivativeSecurityExchange)
}

//HasDerivativePositionLimit returns true if DerivativePositionLimit is present, Tag 1273
func (m DerivativeSecurityListRequest) HasDerivativePositionLimit() bool {
	return m.Has(tag.DerivativePositionLimit)
}

//HasDerivativeNTPositionLimit returns true if DerivativeNTPositionLimit is present, Tag 1274
func (m DerivativeSecurityListRequest) HasDerivativeNTPositionLimit() bool {
	return m.Has(tag.DerivativeNTPositionLimit)
}

//HasDerivativeIssuer returns true if DerivativeIssuer is present, Tag 1275
func (m DerivativeSecurityListRequest) HasDerivativeIssuer() bool {
	return m.Has(tag.DerivativeIssuer)
}

//HasDerivativeIssueDate returns true if DerivativeIssueDate is present, Tag 1276
func (m DerivativeSecurityListRequest) HasDerivativeIssueDate() bool {
	return m.Has(tag.DerivativeIssueDate)
}

//HasDerivativeEncodedIssuerLen returns true if DerivativeEncodedIssuerLen is present, Tag 1277
func (m DerivativeSecurityListRequest) HasDerivativeEncodedIssuerLen() bool {
	return m.Has(tag.DerivativeEncodedIssuerLen)
}

//HasDerivativeEncodedIssuer returns true if DerivativeEncodedIssuer is present, Tag 1278
func (m DerivativeSecurityListRequest) HasDerivativeEncodedIssuer() bool {
	return m.Has(tag.DerivativeEncodedIssuer)
}

//HasDerivativeSecurityDesc returns true if DerivativeSecurityDesc is present, Tag 1279
func (m DerivativeSecurityListRequest) HasDerivativeSecurityDesc() bool {
	return m.Has(tag.DerivativeSecurityDesc)
}

//HasDerivativeEncodedSecurityDescLen returns true if DerivativeEncodedSecurityDescLen is present, Tag 1280
func (m DerivativeSecurityListRequest) HasDerivativeEncodedSecurityDescLen() bool {
	return m.Has(tag.DerivativeEncodedSecurityDescLen)
}

//HasDerivativeEncodedSecurityDesc returns true if DerivativeEncodedSecurityDesc is present, Tag 1281
func (m DerivativeSecurityListRequest) HasDerivativeEncodedSecurityDesc() bool {
	return m.Has(tag.DerivativeEncodedSecurityDesc)
}

//HasDerivativeSecurityXMLLen returns true if DerivativeSecurityXMLLen is present, Tag 1282
func (m DerivativeSecurityListRequest) HasDerivativeSecurityXMLLen() bool {
	return m.Has(tag.DerivativeSecurityXMLLen)
}

//HasDerivativeSecurityXML returns true if DerivativeSecurityXML is present, Tag 1283
func (m DerivativeSecurityListRequest) HasDerivativeSecurityXML() bool {
	return m.Has(tag.DerivativeSecurityXML)
}

//HasDerivativeSecurityXMLSchema returns true if DerivativeSecurityXMLSchema is present, Tag 1284
func (m DerivativeSecurityListRequest) HasDerivativeSecurityXMLSchema() bool {
	return m.Has(tag.DerivativeSecurityXMLSchema)
}

//HasDerivativeContractSettlMonth returns true if DerivativeContractSettlMonth is present, Tag 1285
func (m DerivativeSecurityListRequest) HasDerivativeContractSettlMonth() bool {
	return m.Has(tag.DerivativeContractSettlMonth)
}

//HasNoDerivativeEvents returns true if NoDerivativeEvents is present, Tag 1286
func (m DerivativeSecurityListRequest) HasNoDerivativeEvents() bool {
	return m.Has(tag.NoDerivativeEvents)
}

//HasNoDerivativeInstrumentParties returns true if NoDerivativeInstrumentParties is present, Tag 1292
func (m DerivativeSecurityListRequest) HasNoDerivativeInstrumentParties() bool {
	return m.Has(tag.NoDerivativeInstrumentParties)
}

//HasDerivativeExerciseStyle returns true if DerivativeExerciseStyle is present, Tag 1299
func (m DerivativeSecurityListRequest) HasDerivativeExerciseStyle() bool {
	return m.Has(tag.DerivativeExerciseStyle)
}

//HasMarketSegmentID returns true if MarketSegmentID is present, Tag 1300
func (m DerivativeSecurityListRequest) HasMarketSegmentID() bool {
	return m.Has(tag.MarketSegmentID)
}

//HasMarketID returns true if MarketID is present, Tag 1301
func (m DerivativeSecurityListRequest) HasMarketID() bool {
	return m.Has(tag.MarketID)
}

//HasDerivativePriceUnitOfMeasure returns true if DerivativePriceUnitOfMeasure is present, Tag 1315
func (m DerivativeSecurityListRequest) HasDerivativePriceUnitOfMeasure() bool {
	return m.Has(tag.DerivativePriceUnitOfMeasure)
}

//HasDerivativePriceUnitOfMeasureQty returns true if DerivativePriceUnitOfMeasureQty is present, Tag 1316
func (m DerivativeSecurityListRequest) HasDerivativePriceUnitOfMeasureQty() bool {
	return m.Has(tag.DerivativePriceUnitOfMeasureQty)
}

//HasDerivativeSettlMethod returns true if DerivativeSettlMethod is present, Tag 1317
func (m DerivativeSecurityListRequest) HasDerivativeSettlMethod() bool {
	return m.Has(tag.DerivativeSettlMethod)
}

//HasDerivativePriceQuoteMethod returns true if DerivativePriceQuoteMethod is present, Tag 1318
func (m DerivativeSecurityListRequest) HasDerivativePriceQuoteMethod() bool {
	return m.Has(tag.DerivativePriceQuoteMethod)
}

//HasDerivativeFuturesValuationMethod returns true if DerivativeFuturesValuationMethod is present, Tag 1319
func (m DerivativeSecurityListRequest) HasDerivativeFuturesValuationMethod() bool {
	return m.Has(tag.DerivativeFuturesValuationMethod)
}

//HasDerivativeListMethod returns true if DerivativeListMethod is present, Tag 1320
func (m DerivativeSecurityListRequest) HasDerivativeListMethod() bool {
	return m.Has(tag.DerivativeListMethod)
}

//HasDerivativeCapPrice returns true if DerivativeCapPrice is present, Tag 1321
func (m DerivativeSecurityListRequest) HasDerivativeCapPrice() bool {
	return m.Has(tag.DerivativeCapPrice)
}

//HasDerivativeFloorPrice returns true if DerivativeFloorPrice is present, Tag 1322
func (m DerivativeSecurityListRequest) HasDerivativeFloorPrice() bool {
	return m.Has(tag.DerivativeFloorPrice)
}

//HasDerivativePutOrCall returns true if DerivativePutOrCall is present, Tag 1323
func (m DerivativeSecurityListRequest) HasDerivativePutOrCall() bool {
	return m.Has(tag.DerivativePutOrCall)
}

//HasUnderlyingExerciseStyle returns true if UnderlyingExerciseStyle is present, Tag 1419
func (m DerivativeSecurityListRequest) HasUnderlyingExerciseStyle() bool {
	return m.Has(tag.UnderlyingExerciseStyle)
}

//HasUnderlyingUnitOfMeasureQty returns true if UnderlyingUnitOfMeasureQty is present, Tag 1423
func (m DerivativeSecurityListRequest) HasUnderlyingUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingUnitOfMeasureQty)
}

//HasUnderlyingPriceUnitOfMeasure returns true if UnderlyingPriceUnitOfMeasure is present, Tag 1424
func (m DerivativeSecurityListRequest) HasUnderlyingPriceUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasure)
}

//HasUnderlyingPriceUnitOfMeasureQty returns true if UnderlyingPriceUnitOfMeasureQty is present, Tag 1425
func (m DerivativeSecurityListRequest) HasUnderlyingPriceUnitOfMeasureQty() bool {
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

//NoDerivativeSecurityAltID is a repeating group element, Tag 1218
type NoDerivativeSecurityAltID struct {
	quickfix.Group
}

//SetDerivativeSecurityAltID sets DerivativeSecurityAltID, Tag 1219
func (m NoDerivativeSecurityAltID) SetDerivativeSecurityAltID(v string) {
	m.Set(field.NewDerivativeSecurityAltID(v))
}

//SetDerivativeSecurityAltIDSource sets DerivativeSecurityAltIDSource, Tag 1220
func (m NoDerivativeSecurityAltID) SetDerivativeSecurityAltIDSource(v string) {
	m.Set(field.NewDerivativeSecurityAltIDSource(v))
}

//GetDerivativeSecurityAltID gets DerivativeSecurityAltID, Tag 1219
func (m NoDerivativeSecurityAltID) GetDerivativeSecurityAltID() (f field.DerivativeSecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeSecurityAltIDSource gets DerivativeSecurityAltIDSource, Tag 1220
func (m NoDerivativeSecurityAltID) GetDerivativeSecurityAltIDSource() (f field.DerivativeSecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasDerivativeSecurityAltID returns true if DerivativeSecurityAltID is present, Tag 1219
func (m NoDerivativeSecurityAltID) HasDerivativeSecurityAltID() bool {
	return m.Has(tag.DerivativeSecurityAltID)
}

//HasDerivativeSecurityAltIDSource returns true if DerivativeSecurityAltIDSource is present, Tag 1220
func (m NoDerivativeSecurityAltID) HasDerivativeSecurityAltIDSource() bool {
	return m.Has(tag.DerivativeSecurityAltIDSource)
}

//NoDerivativeSecurityAltIDRepeatingGroup is a repeating group, Tag 1218
type NoDerivativeSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoDerivativeSecurityAltIDRepeatingGroup returns an initialized, NoDerivativeSecurityAltIDRepeatingGroup
func NewNoDerivativeSecurityAltIDRepeatingGroup() NoDerivativeSecurityAltIDRepeatingGroup {
	return NoDerivativeSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoDerivativeSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.DerivativeSecurityAltID), quickfix.GroupElement(tag.DerivativeSecurityAltIDSource)})}
}

//Add create and append a new NoDerivativeSecurityAltID to this group
func (m NoDerivativeSecurityAltIDRepeatingGroup) Add() NoDerivativeSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoDerivativeSecurityAltID{g}
}

//Get returns the ith NoDerivativeSecurityAltID in the NoDerivativeSecurityAltIDRepeatinGroup
func (m NoDerivativeSecurityAltIDRepeatingGroup) Get(i int) NoDerivativeSecurityAltID {
	return NoDerivativeSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoDerivativeEvents is a repeating group element, Tag 1286
type NoDerivativeEvents struct {
	quickfix.Group
}

//SetDerivativeEventType sets DerivativeEventType, Tag 1287
func (m NoDerivativeEvents) SetDerivativeEventType(v int) {
	m.Set(field.NewDerivativeEventType(v))
}

//SetDerivativeEventDate sets DerivativeEventDate, Tag 1288
func (m NoDerivativeEvents) SetDerivativeEventDate(v string) {
	m.Set(field.NewDerivativeEventDate(v))
}

//SetDerivativeEventTime sets DerivativeEventTime, Tag 1289
func (m NoDerivativeEvents) SetDerivativeEventTime(v time.Time) {
	m.Set(field.NewDerivativeEventTime(v))
}

//SetDerivativeEventPx sets DerivativeEventPx, Tag 1290
func (m NoDerivativeEvents) SetDerivativeEventPx(v float64) {
	m.Set(field.NewDerivativeEventPx(v))
}

//SetDerivativeEventText sets DerivativeEventText, Tag 1291
func (m NoDerivativeEvents) SetDerivativeEventText(v string) {
	m.Set(field.NewDerivativeEventText(v))
}

//GetDerivativeEventType gets DerivativeEventType, Tag 1287
func (m NoDerivativeEvents) GetDerivativeEventType() (f field.DerivativeEventTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeEventDate gets DerivativeEventDate, Tag 1288
func (m NoDerivativeEvents) GetDerivativeEventDate() (f field.DerivativeEventDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeEventTime gets DerivativeEventTime, Tag 1289
func (m NoDerivativeEvents) GetDerivativeEventTime() (f field.DerivativeEventTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeEventPx gets DerivativeEventPx, Tag 1290
func (m NoDerivativeEvents) GetDerivativeEventPx() (f field.DerivativeEventPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeEventText gets DerivativeEventText, Tag 1291
func (m NoDerivativeEvents) GetDerivativeEventText() (f field.DerivativeEventTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasDerivativeEventType returns true if DerivativeEventType is present, Tag 1287
func (m NoDerivativeEvents) HasDerivativeEventType() bool {
	return m.Has(tag.DerivativeEventType)
}

//HasDerivativeEventDate returns true if DerivativeEventDate is present, Tag 1288
func (m NoDerivativeEvents) HasDerivativeEventDate() bool {
	return m.Has(tag.DerivativeEventDate)
}

//HasDerivativeEventTime returns true if DerivativeEventTime is present, Tag 1289
func (m NoDerivativeEvents) HasDerivativeEventTime() bool {
	return m.Has(tag.DerivativeEventTime)
}

//HasDerivativeEventPx returns true if DerivativeEventPx is present, Tag 1290
func (m NoDerivativeEvents) HasDerivativeEventPx() bool {
	return m.Has(tag.DerivativeEventPx)
}

//HasDerivativeEventText returns true if DerivativeEventText is present, Tag 1291
func (m NoDerivativeEvents) HasDerivativeEventText() bool {
	return m.Has(tag.DerivativeEventText)
}

//NoDerivativeEventsRepeatingGroup is a repeating group, Tag 1286
type NoDerivativeEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoDerivativeEventsRepeatingGroup returns an initialized, NoDerivativeEventsRepeatingGroup
func NewNoDerivativeEventsRepeatingGroup() NoDerivativeEventsRepeatingGroup {
	return NoDerivativeEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoDerivativeEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.DerivativeEventType), quickfix.GroupElement(tag.DerivativeEventDate), quickfix.GroupElement(tag.DerivativeEventTime), quickfix.GroupElement(tag.DerivativeEventPx), quickfix.GroupElement(tag.DerivativeEventText)})}
}

//Add create and append a new NoDerivativeEvents to this group
func (m NoDerivativeEventsRepeatingGroup) Add() NoDerivativeEvents {
	g := m.RepeatingGroup.Add()
	return NoDerivativeEvents{g}
}

//Get returns the ith NoDerivativeEvents in the NoDerivativeEventsRepeatinGroup
func (m NoDerivativeEventsRepeatingGroup) Get(i int) NoDerivativeEvents {
	return NoDerivativeEvents{m.RepeatingGroup.Get(i)}
}

//NoDerivativeInstrumentParties is a repeating group element, Tag 1292
type NoDerivativeInstrumentParties struct {
	quickfix.Group
}

//SetDerivativeInstrumentPartyID sets DerivativeInstrumentPartyID, Tag 1293
func (m NoDerivativeInstrumentParties) SetDerivativeInstrumentPartyID(v string) {
	m.Set(field.NewDerivativeInstrumentPartyID(v))
}

//SetDerivativeInstrumentPartyIDSource sets DerivativeInstrumentPartyIDSource, Tag 1294
func (m NoDerivativeInstrumentParties) SetDerivativeInstrumentPartyIDSource(v string) {
	m.Set(field.NewDerivativeInstrumentPartyIDSource(v))
}

//SetDerivativeInstrumentPartyRole sets DerivativeInstrumentPartyRole, Tag 1295
func (m NoDerivativeInstrumentParties) SetDerivativeInstrumentPartyRole(v int) {
	m.Set(field.NewDerivativeInstrumentPartyRole(v))
}

//SetNoDerivativeInstrumentPartySubIDs sets NoDerivativeInstrumentPartySubIDs, Tag 1296
func (m NoDerivativeInstrumentParties) SetNoDerivativeInstrumentPartySubIDs(f NoDerivativeInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetDerivativeInstrumentPartyID gets DerivativeInstrumentPartyID, Tag 1293
func (m NoDerivativeInstrumentParties) GetDerivativeInstrumentPartyID() (f field.DerivativeInstrumentPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeInstrumentPartyIDSource gets DerivativeInstrumentPartyIDSource, Tag 1294
func (m NoDerivativeInstrumentParties) GetDerivativeInstrumentPartyIDSource() (f field.DerivativeInstrumentPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeInstrumentPartyRole gets DerivativeInstrumentPartyRole, Tag 1295
func (m NoDerivativeInstrumentParties) GetDerivativeInstrumentPartyRole() (f field.DerivativeInstrumentPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoDerivativeInstrumentPartySubIDs gets NoDerivativeInstrumentPartySubIDs, Tag 1296
func (m NoDerivativeInstrumentParties) GetNoDerivativeInstrumentPartySubIDs() (NoDerivativeInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoDerivativeInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasDerivativeInstrumentPartyID returns true if DerivativeInstrumentPartyID is present, Tag 1293
func (m NoDerivativeInstrumentParties) HasDerivativeInstrumentPartyID() bool {
	return m.Has(tag.DerivativeInstrumentPartyID)
}

//HasDerivativeInstrumentPartyIDSource returns true if DerivativeInstrumentPartyIDSource is present, Tag 1294
func (m NoDerivativeInstrumentParties) HasDerivativeInstrumentPartyIDSource() bool {
	return m.Has(tag.DerivativeInstrumentPartyIDSource)
}

//HasDerivativeInstrumentPartyRole returns true if DerivativeInstrumentPartyRole is present, Tag 1295
func (m NoDerivativeInstrumentParties) HasDerivativeInstrumentPartyRole() bool {
	return m.Has(tag.DerivativeInstrumentPartyRole)
}

//HasNoDerivativeInstrumentPartySubIDs returns true if NoDerivativeInstrumentPartySubIDs is present, Tag 1296
func (m NoDerivativeInstrumentParties) HasNoDerivativeInstrumentPartySubIDs() bool {
	return m.Has(tag.NoDerivativeInstrumentPartySubIDs)
}

//NoDerivativeInstrumentPartySubIDs is a repeating group element, Tag 1296
type NoDerivativeInstrumentPartySubIDs struct {
	quickfix.Group
}

//SetDerivativeInstrumentPartySubID sets DerivativeInstrumentPartySubID, Tag 1297
func (m NoDerivativeInstrumentPartySubIDs) SetDerivativeInstrumentPartySubID(v string) {
	m.Set(field.NewDerivativeInstrumentPartySubID(v))
}

//SetDerivativeInstrumentPartySubIDType sets DerivativeInstrumentPartySubIDType, Tag 1298
func (m NoDerivativeInstrumentPartySubIDs) SetDerivativeInstrumentPartySubIDType(v int) {
	m.Set(field.NewDerivativeInstrumentPartySubIDType(v))
}

//GetDerivativeInstrumentPartySubID gets DerivativeInstrumentPartySubID, Tag 1297
func (m NoDerivativeInstrumentPartySubIDs) GetDerivativeInstrumentPartySubID() (f field.DerivativeInstrumentPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDerivativeInstrumentPartySubIDType gets DerivativeInstrumentPartySubIDType, Tag 1298
func (m NoDerivativeInstrumentPartySubIDs) GetDerivativeInstrumentPartySubIDType() (f field.DerivativeInstrumentPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasDerivativeInstrumentPartySubID returns true if DerivativeInstrumentPartySubID is present, Tag 1297
func (m NoDerivativeInstrumentPartySubIDs) HasDerivativeInstrumentPartySubID() bool {
	return m.Has(tag.DerivativeInstrumentPartySubID)
}

//HasDerivativeInstrumentPartySubIDType returns true if DerivativeInstrumentPartySubIDType is present, Tag 1298
func (m NoDerivativeInstrumentPartySubIDs) HasDerivativeInstrumentPartySubIDType() bool {
	return m.Has(tag.DerivativeInstrumentPartySubIDType)
}

//NoDerivativeInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1296
type NoDerivativeInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoDerivativeInstrumentPartySubIDsRepeatingGroup returns an initialized, NoDerivativeInstrumentPartySubIDsRepeatingGroup
func NewNoDerivativeInstrumentPartySubIDsRepeatingGroup() NoDerivativeInstrumentPartySubIDsRepeatingGroup {
	return NoDerivativeInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoDerivativeInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.DerivativeInstrumentPartySubID), quickfix.GroupElement(tag.DerivativeInstrumentPartySubIDType)})}
}

//Add create and append a new NoDerivativeInstrumentPartySubIDs to this group
func (m NoDerivativeInstrumentPartySubIDsRepeatingGroup) Add() NoDerivativeInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoDerivativeInstrumentPartySubIDs{g}
}

//Get returns the ith NoDerivativeInstrumentPartySubIDs in the NoDerivativeInstrumentPartySubIDsRepeatinGroup
func (m NoDerivativeInstrumentPartySubIDsRepeatingGroup) Get(i int) NoDerivativeInstrumentPartySubIDs {
	return NoDerivativeInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoDerivativeInstrumentPartiesRepeatingGroup is a repeating group, Tag 1292
type NoDerivativeInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoDerivativeInstrumentPartiesRepeatingGroup returns an initialized, NoDerivativeInstrumentPartiesRepeatingGroup
func NewNoDerivativeInstrumentPartiesRepeatingGroup() NoDerivativeInstrumentPartiesRepeatingGroup {
	return NoDerivativeInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoDerivativeInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.DerivativeInstrumentPartyID), quickfix.GroupElement(tag.DerivativeInstrumentPartyIDSource), quickfix.GroupElement(tag.DerivativeInstrumentPartyRole), NewNoDerivativeInstrumentPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoDerivativeInstrumentParties to this group
func (m NoDerivativeInstrumentPartiesRepeatingGroup) Add() NoDerivativeInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoDerivativeInstrumentParties{g}
}

//Get returns the ith NoDerivativeInstrumentParties in the NoDerivativeInstrumentPartiesRepeatinGroup
func (m NoDerivativeInstrumentPartiesRepeatingGroup) Get(i int) NoDerivativeInstrumentParties {
	return NoDerivativeInstrumentParties{m.RepeatingGroup.Get(i)}
}
