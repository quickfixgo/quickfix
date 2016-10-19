package derivativesecuritylistrequest

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//DerivativeSecurityListRequest is the fix50sp2 DerivativeSecurityListRequest type, MsgType = z
type DerivativeSecurityListRequest struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

//FromMessage creates a DerivativeSecurityListRequest from a quickfix.Message instance
func FromMessage(m *quickfix.Message) DerivativeSecurityListRequest {
	return DerivativeSecurityListRequest{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m DerivativeSecurityListRequest) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a DerivativeSecurityListRequest initialized with the required fields for DerivativeSecurityListRequest
func New(securityreqid field.SecurityReqIDField, securitylistrequesttype field.SecurityListRequestTypeField) (m DerivativeSecurityListRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("z"))
	m.Set(securityreqid)
	m.Set(securitylistrequesttype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg DerivativeSecurityListRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "z", r
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
func (m DerivativeSecurityListRequest) SetUnderlyingRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingRepurchaseRate(value, scale))
}

//SetUnderlyingFactor sets UnderlyingFactor, Tag 246
func (m DerivativeSecurityListRequest) SetUnderlyingFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFactor(value, scale))
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
func (m DerivativeSecurityListRequest) SetSubscriptionRequestType(v enum.SubscriptionRequestType) {
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
func (m DerivativeSecurityListRequest) SetUnderlyingStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStrikePrice(value, scale))
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
func (m DerivativeSecurityListRequest) SetTradingSessionID(v enum.TradingSessionID) {
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
func (m DerivativeSecurityListRequest) SetUnderlyingCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCouponRate(value, scale))
}

//SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436
func (m DerivativeSecurityListRequest) SetUnderlyingContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingContractMultiplier(value, scale))
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
func (m DerivativeSecurityListRequest) SetSecurityListRequestType(v enum.SecurityListRequestType) {
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
func (m DerivativeSecurityListRequest) SetTradingSessionSubID(v enum.TradingSessionSubID) {
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
func (m DerivativeSecurityListRequest) SetUnderlyingPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPx(value, scale))
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
func (m DerivativeSecurityListRequest) SetUnderlyingQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingQty(value, scale))
}

//SetUnderlyingDirtyPrice sets UnderlyingDirtyPrice, Tag 882
func (m DerivativeSecurityListRequest) SetUnderlyingDirtyPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDirtyPrice(value, scale))
}

//SetUnderlyingEndPrice sets UnderlyingEndPrice, Tag 883
func (m DerivativeSecurityListRequest) SetUnderlyingEndPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndPrice(value, scale))
}

//SetUnderlyingStartValue sets UnderlyingStartValue, Tag 884
func (m DerivativeSecurityListRequest) SetUnderlyingStartValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStartValue(value, scale))
}

//SetUnderlyingCurrentValue sets UnderlyingCurrentValue, Tag 885
func (m DerivativeSecurityListRequest) SetUnderlyingCurrentValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCurrentValue(value, scale))
}

//SetUnderlyingEndValue sets UnderlyingEndValue, Tag 886
func (m DerivativeSecurityListRequest) SetUnderlyingEndValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndValue(value, scale))
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
func (m DerivativeSecurityListRequest) SetUnderlyingAllocationPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAllocationPercent(value, scale))
}

//SetUnderlyingCashAmount sets UnderlyingCashAmount, Tag 973
func (m DerivativeSecurityListRequest) SetUnderlyingCashAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCashAmount(value, scale))
}

//SetUnderlyingCashType sets UnderlyingCashType, Tag 974
func (m DerivativeSecurityListRequest) SetUnderlyingCashType(v enum.UnderlyingCashType) {
	m.Set(field.NewUnderlyingCashType(v))
}

//SetUnderlyingSettlementType sets UnderlyingSettlementType, Tag 975
func (m DerivativeSecurityListRequest) SetUnderlyingSettlementType(v enum.UnderlyingSettlementType) {
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
func (m DerivativeSecurityListRequest) SetUnderlyingCapValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCapValue(value, scale))
}

//SetUnderlyingSettlMethod sets UnderlyingSettlMethod, Tag 1039
func (m DerivativeSecurityListRequest) SetUnderlyingSettlMethod(v string) {
	m.Set(field.NewUnderlyingSettlMethod(v))
}

//SetUnderlyingAdjustedQuantity sets UnderlyingAdjustedQuantity, Tag 1044
func (m DerivativeSecurityListRequest) SetUnderlyingAdjustedQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAdjustedQuantity(value, scale))
}

//SetUnderlyingFXRate sets UnderlyingFXRate, Tag 1045
func (m DerivativeSecurityListRequest) SetUnderlyingFXRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFXRate(value, scale))
}

//SetUnderlyingFXRateCalc sets UnderlyingFXRateCalc, Tag 1046
func (m DerivativeSecurityListRequest) SetUnderlyingFXRateCalc(v enum.UnderlyingFXRateCalc) {
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
func (m DerivativeSecurityListRequest) SetDerivativeOptPayAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeOptPayAmount(value, scale))
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
func (m DerivativeSecurityListRequest) SetDerivativeStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeStrikePrice(value, scale))
}

//SetDerivativeStrikeCurrency sets DerivativeStrikeCurrency, Tag 1262
func (m DerivativeSecurityListRequest) SetDerivativeStrikeCurrency(v string) {
	m.Set(field.NewDerivativeStrikeCurrency(v))
}

//SetDerivativeStrikeMultiplier sets DerivativeStrikeMultiplier, Tag 1263
func (m DerivativeSecurityListRequest) SetDerivativeStrikeMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeStrikeMultiplier(value, scale))
}

//SetDerivativeStrikeValue sets DerivativeStrikeValue, Tag 1264
func (m DerivativeSecurityListRequest) SetDerivativeStrikeValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeStrikeValue(value, scale))
}

//SetDerivativeOptAttribute sets DerivativeOptAttribute, Tag 1265
func (m DerivativeSecurityListRequest) SetDerivativeOptAttribute(v string) {
	m.Set(field.NewDerivativeOptAttribute(v))
}

//SetDerivativeContractMultiplier sets DerivativeContractMultiplier, Tag 1266
func (m DerivativeSecurityListRequest) SetDerivativeContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeContractMultiplier(value, scale))
}

//SetDerivativeMinPriceIncrement sets DerivativeMinPriceIncrement, Tag 1267
func (m DerivativeSecurityListRequest) SetDerivativeMinPriceIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeMinPriceIncrement(value, scale))
}

//SetDerivativeMinPriceIncrementAmount sets DerivativeMinPriceIncrementAmount, Tag 1268
func (m DerivativeSecurityListRequest) SetDerivativeMinPriceIncrementAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeMinPriceIncrementAmount(value, scale))
}

//SetDerivativeUnitOfMeasure sets DerivativeUnitOfMeasure, Tag 1269
func (m DerivativeSecurityListRequest) SetDerivativeUnitOfMeasure(v string) {
	m.Set(field.NewDerivativeUnitOfMeasure(v))
}

//SetDerivativeUnitOfMeasureQty sets DerivativeUnitOfMeasureQty, Tag 1270
func (m DerivativeSecurityListRequest) SetDerivativeUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeUnitOfMeasureQty(value, scale))
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
func (m DerivativeSecurityListRequest) SetDerivativePriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativePriceUnitOfMeasureQty(value, scale))
}

//SetDerivativeSettlMethod sets DerivativeSettlMethod, Tag 1317
func (m DerivativeSecurityListRequest) SetDerivativeSettlMethod(v string) {
	m.Set(field.NewDerivativeSettlMethod(v))
}

//SetDerivativePriceQuoteMethod sets DerivativePriceQuoteMethod, Tag 1318
func (m DerivativeSecurityListRequest) SetDerivativePriceQuoteMethod(v string) {
	m.Set(field.NewDerivativePriceQuoteMethod(v))
}

//SetDerivativeValuationMethod sets DerivativeValuationMethod, Tag 1319
func (m DerivativeSecurityListRequest) SetDerivativeValuationMethod(v string) {
	m.Set(field.NewDerivativeValuationMethod(v))
}

//SetDerivativeListMethod sets DerivativeListMethod, Tag 1320
func (m DerivativeSecurityListRequest) SetDerivativeListMethod(v int) {
	m.Set(field.NewDerivativeListMethod(v))
}

//SetDerivativeCapPrice sets DerivativeCapPrice, Tag 1321
func (m DerivativeSecurityListRequest) SetDerivativeCapPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeCapPrice(value, scale))
}

//SetDerivativeFloorPrice sets DerivativeFloorPrice, Tag 1322
func (m DerivativeSecurityListRequest) SetDerivativeFloorPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeFloorPrice(value, scale))
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
func (m DerivativeSecurityListRequest) SetUnderlyingUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingUnitOfMeasureQty(value, scale))
}

//SetUnderlyingPriceUnitOfMeasure sets UnderlyingPriceUnitOfMeasure, Tag 1424
func (m DerivativeSecurityListRequest) SetUnderlyingPriceUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasure(v))
}

//SetUnderlyingPriceUnitOfMeasureQty sets UnderlyingPriceUnitOfMeasureQty, Tag 1425
func (m DerivativeSecurityListRequest) SetUnderlyingPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasureQty(value, scale))
}

//SetUnderlyingContractMultiplierUnit sets UnderlyingContractMultiplierUnit, Tag 1437
func (m DerivativeSecurityListRequest) SetUnderlyingContractMultiplierUnit(v int) {
	m.Set(field.NewUnderlyingContractMultiplierUnit(v))
}

//SetDerivativeContractMultiplierUnit sets DerivativeContractMultiplierUnit, Tag 1438
func (m DerivativeSecurityListRequest) SetDerivativeContractMultiplierUnit(v int) {
	m.Set(field.NewDerivativeContractMultiplierUnit(v))
}

//SetUnderlyingFlowScheduleType sets UnderlyingFlowScheduleType, Tag 1441
func (m DerivativeSecurityListRequest) SetUnderlyingFlowScheduleType(v int) {
	m.Set(field.NewUnderlyingFlowScheduleType(v))
}

//SetDerivativeFlowScheduleType sets DerivativeFlowScheduleType, Tag 1442
func (m DerivativeSecurityListRequest) SetDerivativeFlowScheduleType(v int) {
	m.Set(field.NewDerivativeFlowScheduleType(v))
}

//SetUnderlyingRestructuringType sets UnderlyingRestructuringType, Tag 1453
func (m DerivativeSecurityListRequest) SetUnderlyingRestructuringType(v string) {
	m.Set(field.NewUnderlyingRestructuringType(v))
}

//SetUnderlyingSeniority sets UnderlyingSeniority, Tag 1454
func (m DerivativeSecurityListRequest) SetUnderlyingSeniority(v string) {
	m.Set(field.NewUnderlyingSeniority(v))
}

//SetUnderlyingNotionalPercentageOutstanding sets UnderlyingNotionalPercentageOutstanding, Tag 1455
func (m DerivativeSecurityListRequest) SetUnderlyingNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingNotionalPercentageOutstanding(value, scale))
}

//SetUnderlyingOriginalNotionalPercentageOutstanding sets UnderlyingOriginalNotionalPercentageOutstanding, Tag 1456
func (m DerivativeSecurityListRequest) SetUnderlyingOriginalNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingOriginalNotionalPercentageOutstanding(value, scale))
}

//SetUnderlyingAttachmentPoint sets UnderlyingAttachmentPoint, Tag 1459
func (m DerivativeSecurityListRequest) SetUnderlyingAttachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAttachmentPoint(value, scale))
}

//SetUnderlyingDetachmentPoint sets UnderlyingDetachmentPoint, Tag 1460
func (m DerivativeSecurityListRequest) SetUnderlyingDetachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDetachmentPoint(value, scale))
}

//GetCurrency gets Currency, Tag 15
func (m DerivativeSecurityListRequest) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m DerivativeSecurityListRequest) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCouponPaymentDate gets UnderlyingCouponPaymentDate, Tag 241
func (m DerivativeSecurityListRequest) GetUnderlyingCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingIssueDate gets UnderlyingIssueDate, Tag 242
func (m DerivativeSecurityListRequest) GetUnderlyingIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRepoCollateralSecurityType gets UnderlyingRepoCollateralSecurityType, Tag 243
func (m DerivativeSecurityListRequest) GetUnderlyingRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRepurchaseTerm gets UnderlyingRepurchaseTerm, Tag 244
func (m DerivativeSecurityListRequest) GetUnderlyingRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRepurchaseRate gets UnderlyingRepurchaseRate, Tag 245
func (m DerivativeSecurityListRequest) GetUnderlyingRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingFactor gets UnderlyingFactor, Tag 246
func (m DerivativeSecurityListRequest) GetUnderlyingFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingFactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRedemptionDate gets UnderlyingRedemptionDate, Tag 247
func (m DerivativeSecurityListRequest) GetUnderlyingRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCreditRating gets UnderlyingCreditRating, Tag 256
func (m DerivativeSecurityListRequest) GetUnderlyingCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263
func (m DerivativeSecurityListRequest) GetSubscriptionRequestType() (v enum.SubscriptionRequestType, err quickfix.MessageRejectError) {
	var f field.SubscriptionRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityIDSource gets UnderlyingSecurityIDSource, Tag 305
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306
func (m DerivativeSecurityListRequest) GetUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311
func (m DerivativeSecurityListRequest) GetUnderlyingSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312
func (m DerivativeSecurityListRequest) GetUnderlyingSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313
func (m DerivativeSecurityListRequest) GetUnderlyingMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingPutOrCall gets UnderlyingPutOrCall, Tag 315
func (m DerivativeSecurityListRequest) GetUnderlyingPutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingPutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316
func (m DerivativeSecurityListRequest) GetUnderlyingStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317
func (m DerivativeSecurityListRequest) GetUnderlyingOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCurrency gets UnderlyingCurrency, Tag 318
func (m DerivativeSecurityListRequest) GetUnderlyingCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityReqID gets SecurityReqID, Tag 320
func (m DerivativeSecurityListRequest) GetSecurityReqID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m DerivativeSecurityListRequest) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m DerivativeSecurityListRequest) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m DerivativeSecurityListRequest) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435
func (m DerivativeSecurityListRequest) GetUnderlyingCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436
func (m DerivativeSecurityListRequest) GetUnderlyingContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoUnderlyingSecurityAltID gets NoUnderlyingSecurityAltID, Tag 457
func (m DerivativeSecurityListRequest) GetNoUnderlyingSecurityAltID() (NoUnderlyingSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingProduct gets UnderlyingProduct, Tag 462
func (m DerivativeSecurityListRequest) GetUnderlyingProduct() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCFICode gets UnderlyingCFICode, Tag 463
func (m DerivativeSecurityListRequest) GetUnderlyingCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingMaturityDate gets UnderlyingMaturityDate, Tag 542
func (m DerivativeSecurityListRequest) GetUnderlyingMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityListRequestType gets SecurityListRequestType, Tag 559
func (m DerivativeSecurityListRequest) GetSecurityListRequestType() (v enum.SecurityListRequestType, err quickfix.MessageRejectError) {
	var f field.SecurityListRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCountryOfIssue gets UnderlyingCountryOfIssue, Tag 592
func (m DerivativeSecurityListRequest) GetUnderlyingCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStateOrProvinceOfIssue gets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m DerivativeSecurityListRequest) GetUnderlyingStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingLocaleOfIssue gets UnderlyingLocaleOfIssue, Tag 594
func (m DerivativeSecurityListRequest) GetUnderlyingLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingInstrRegistry gets UnderlyingInstrRegistry, Tag 595
func (m DerivativeSecurityListRequest) GetUnderlyingInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m DerivativeSecurityListRequest) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m DerivativeSecurityListRequest) GetSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecuritySubType gets UnderlyingSecuritySubType, Tag 763
func (m DerivativeSecurityListRequest) GetUnderlyingSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingPx gets UnderlyingPx, Tag 810
func (m DerivativeSecurityListRequest) GetUnderlyingPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCPProgram gets UnderlyingCPProgram, Tag 877
func (m DerivativeSecurityListRequest) GetUnderlyingCPProgram() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCPRegType gets UnderlyingCPRegType, Tag 878
func (m DerivativeSecurityListRequest) GetUnderlyingCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingQty gets UnderlyingQty, Tag 879
func (m DerivativeSecurityListRequest) GetUnderlyingQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingDirtyPrice gets UnderlyingDirtyPrice, Tag 882
func (m DerivativeSecurityListRequest) GetUnderlyingDirtyPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingDirtyPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingEndPrice gets UnderlyingEndPrice, Tag 883
func (m DerivativeSecurityListRequest) GetUnderlyingEndPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStartValue gets UnderlyingStartValue, Tag 884
func (m DerivativeSecurityListRequest) GetUnderlyingStartValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStartValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCurrentValue gets UnderlyingCurrentValue, Tag 885
func (m DerivativeSecurityListRequest) GetUnderlyingCurrentValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrentValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingEndValue gets UnderlyingEndValue, Tag 886
func (m DerivativeSecurityListRequest) GetUnderlyingEndValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoUnderlyingStips gets NoUnderlyingStips, Tag 887
func (m DerivativeSecurityListRequest) GetNoUnderlyingStips() (NoUnderlyingStipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingStipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingStrikeCurrency gets UnderlyingStrikeCurrency, Tag 941
func (m DerivativeSecurityListRequest) GetUnderlyingStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingAllocationPercent gets UnderlyingAllocationPercent, Tag 972
func (m DerivativeSecurityListRequest) GetUnderlyingAllocationPercent() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAllocationPercentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCashAmount gets UnderlyingCashAmount, Tag 973
func (m DerivativeSecurityListRequest) GetUnderlyingCashAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCashAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCashType gets UnderlyingCashType, Tag 974
func (m DerivativeSecurityListRequest) GetUnderlyingCashType() (v enum.UnderlyingCashType, err quickfix.MessageRejectError) {
	var f field.UnderlyingCashTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSettlementType gets UnderlyingSettlementType, Tag 975
func (m DerivativeSecurityListRequest) GetUnderlyingSettlementType() (v enum.UnderlyingSettlementType, err quickfix.MessageRejectError) {
	var f field.UnderlyingSettlementTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingUnitOfMeasure gets UnderlyingUnitOfMeasure, Tag 998
func (m DerivativeSecurityListRequest) GetUnderlyingUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingTimeUnit gets UnderlyingTimeUnit, Tag 1000
func (m DerivativeSecurityListRequest) GetUnderlyingTimeUnit() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingTimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCapValue gets UnderlyingCapValue, Tag 1038
func (m DerivativeSecurityListRequest) GetUnderlyingCapValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCapValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSettlMethod gets UnderlyingSettlMethod, Tag 1039
func (m DerivativeSecurityListRequest) GetUnderlyingSettlMethod() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSettlMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingAdjustedQuantity gets UnderlyingAdjustedQuantity, Tag 1044
func (m DerivativeSecurityListRequest) GetUnderlyingAdjustedQuantity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAdjustedQuantityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingFXRate gets UnderlyingFXRate, Tag 1045
func (m DerivativeSecurityListRequest) GetUnderlyingFXRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingFXRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingFXRateCalc gets UnderlyingFXRateCalc, Tag 1046
func (m DerivativeSecurityListRequest) GetUnderlyingFXRateCalc() (v enum.UnderlyingFXRateCalc, err quickfix.MessageRejectError) {
	var f field.UnderlyingFXRateCalcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoUndlyInstrumentParties gets NoUndlyInstrumentParties, Tag 1058
func (m DerivativeSecurityListRequest) GetNoUndlyInstrumentParties() (NoUndlyInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingMaturityTime gets UnderlyingMaturityTime, Tag 1213
func (m DerivativeSecurityListRequest) GetUnderlyingMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSymbol gets DerivativeSymbol, Tag 1214
func (m DerivativeSecurityListRequest) GetDerivativeSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSymbolSfx gets DerivativeSymbolSfx, Tag 1215
func (m DerivativeSecurityListRequest) GetDerivativeSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityID gets DerivativeSecurityID, Tag 1216
func (m DerivativeSecurityListRequest) GetDerivativeSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityIDSource gets DerivativeSecurityIDSource, Tag 1217
func (m DerivativeSecurityListRequest) GetDerivativeSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoDerivativeSecurityAltID gets NoDerivativeSecurityAltID, Tag 1218
func (m DerivativeSecurityListRequest) GetNoDerivativeSecurityAltID() (NoDerivativeSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoDerivativeSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDerivativeOptPayAmount gets DerivativeOptPayAmount, Tag 1225
func (m DerivativeSecurityListRequest) GetDerivativeOptPayAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeOptPayAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeProductComplex gets DerivativeProductComplex, Tag 1228
func (m DerivativeSecurityListRequest) GetDerivativeProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivFlexProductEligibilityIndicator gets DerivFlexProductEligibilityIndicator, Tag 1243
func (m DerivativeSecurityListRequest) GetDerivFlexProductEligibilityIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.DerivFlexProductEligibilityIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeProduct gets DerivativeProduct, Tag 1246
func (m DerivativeSecurityListRequest) GetDerivativeProduct() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityGroup gets DerivativeSecurityGroup, Tag 1247
func (m DerivativeSecurityListRequest) GetDerivativeSecurityGroup() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityGroupField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeCFICode gets DerivativeCFICode, Tag 1248
func (m DerivativeSecurityListRequest) GetDerivativeCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityType gets DerivativeSecurityType, Tag 1249
func (m DerivativeSecurityListRequest) GetDerivativeSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecuritySubType gets DerivativeSecuritySubType, Tag 1250
func (m DerivativeSecurityListRequest) GetDerivativeSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeMaturityMonthYear gets DerivativeMaturityMonthYear, Tag 1251
func (m DerivativeSecurityListRequest) GetDerivativeMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeMaturityDate gets DerivativeMaturityDate, Tag 1252
func (m DerivativeSecurityListRequest) GetDerivativeMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeMaturityTime gets DerivativeMaturityTime, Tag 1253
func (m DerivativeSecurityListRequest) GetDerivativeMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSettleOnOpenFlag gets DerivativeSettleOnOpenFlag, Tag 1254
func (m DerivativeSecurityListRequest) GetDerivativeSettleOnOpenFlag() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSettleOnOpenFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeInstrmtAssignmentMethod gets DerivativeInstrmtAssignmentMethod, Tag 1255
func (m DerivativeSecurityListRequest) GetDerivativeInstrmtAssignmentMethod() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeInstrmtAssignmentMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityStatus gets DerivativeSecurityStatus, Tag 1256
func (m DerivativeSecurityListRequest) GetDerivativeSecurityStatus() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeInstrRegistry gets DerivativeInstrRegistry, Tag 1257
func (m DerivativeSecurityListRequest) GetDerivativeInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeCountryOfIssue gets DerivativeCountryOfIssue, Tag 1258
func (m DerivativeSecurityListRequest) GetDerivativeCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeStateOrProvinceOfIssue gets DerivativeStateOrProvinceOfIssue, Tag 1259
func (m DerivativeSecurityListRequest) GetDerivativeStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeLocaleOfIssue gets DerivativeLocaleOfIssue, Tag 1260
func (m DerivativeSecurityListRequest) GetDerivativeLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeStrikePrice gets DerivativeStrikePrice, Tag 1261
func (m DerivativeSecurityListRequest) GetDerivativeStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeStrikeCurrency gets DerivativeStrikeCurrency, Tag 1262
func (m DerivativeSecurityListRequest) GetDerivativeStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeStrikeMultiplier gets DerivativeStrikeMultiplier, Tag 1263
func (m DerivativeSecurityListRequest) GetDerivativeStrikeMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeStrikeMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeStrikeValue gets DerivativeStrikeValue, Tag 1264
func (m DerivativeSecurityListRequest) GetDerivativeStrikeValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeStrikeValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeOptAttribute gets DerivativeOptAttribute, Tag 1265
func (m DerivativeSecurityListRequest) GetDerivativeOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeContractMultiplier gets DerivativeContractMultiplier, Tag 1266
func (m DerivativeSecurityListRequest) GetDerivativeContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeMinPriceIncrement gets DerivativeMinPriceIncrement, Tag 1267
func (m DerivativeSecurityListRequest) GetDerivativeMinPriceIncrement() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeMinPriceIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeMinPriceIncrementAmount gets DerivativeMinPriceIncrementAmount, Tag 1268
func (m DerivativeSecurityListRequest) GetDerivativeMinPriceIncrementAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeMinPriceIncrementAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeUnitOfMeasure gets DerivativeUnitOfMeasure, Tag 1269
func (m DerivativeSecurityListRequest) GetDerivativeUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeUnitOfMeasureQty gets DerivativeUnitOfMeasureQty, Tag 1270
func (m DerivativeSecurityListRequest) GetDerivativeUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeTimeUnit gets DerivativeTimeUnit, Tag 1271
func (m DerivativeSecurityListRequest) GetDerivativeTimeUnit() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeTimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityExchange gets DerivativeSecurityExchange, Tag 1272
func (m DerivativeSecurityListRequest) GetDerivativeSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativePositionLimit gets DerivativePositionLimit, Tag 1273
func (m DerivativeSecurityListRequest) GetDerivativePositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativePositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeNTPositionLimit gets DerivativeNTPositionLimit, Tag 1274
func (m DerivativeSecurityListRequest) GetDerivativeNTPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeNTPositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeIssuer gets DerivativeIssuer, Tag 1275
func (m DerivativeSecurityListRequest) GetDerivativeIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeIssueDate gets DerivativeIssueDate, Tag 1276
func (m DerivativeSecurityListRequest) GetDerivativeIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeEncodedIssuerLen gets DerivativeEncodedIssuerLen, Tag 1277
func (m DerivativeSecurityListRequest) GetDerivativeEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeEncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeEncodedIssuer gets DerivativeEncodedIssuer, Tag 1278
func (m DerivativeSecurityListRequest) GetDerivativeEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeEncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityDesc gets DerivativeSecurityDesc, Tag 1279
func (m DerivativeSecurityListRequest) GetDerivativeSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeEncodedSecurityDescLen gets DerivativeEncodedSecurityDescLen, Tag 1280
func (m DerivativeSecurityListRequest) GetDerivativeEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeEncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeEncodedSecurityDesc gets DerivativeEncodedSecurityDesc, Tag 1281
func (m DerivativeSecurityListRequest) GetDerivativeEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeEncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityXMLLen gets DerivativeSecurityXMLLen, Tag 1282
func (m DerivativeSecurityListRequest) GetDerivativeSecurityXMLLen() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityXMLLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityXML gets DerivativeSecurityXML, Tag 1283
func (m DerivativeSecurityListRequest) GetDerivativeSecurityXML() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityXMLField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityXMLSchema gets DerivativeSecurityXMLSchema, Tag 1284
func (m DerivativeSecurityListRequest) GetDerivativeSecurityXMLSchema() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityXMLSchemaField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeContractSettlMonth gets DerivativeContractSettlMonth, Tag 1285
func (m DerivativeSecurityListRequest) GetDerivativeContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
func (m DerivativeSecurityListRequest) GetDerivativeExerciseStyle() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMarketSegmentID gets MarketSegmentID, Tag 1300
func (m DerivativeSecurityListRequest) GetMarketSegmentID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketSegmentIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMarketID gets MarketID, Tag 1301
func (m DerivativeSecurityListRequest) GetMarketID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativePriceUnitOfMeasure gets DerivativePriceUnitOfMeasure, Tag 1315
func (m DerivativeSecurityListRequest) GetDerivativePriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativePriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativePriceUnitOfMeasureQty gets DerivativePriceUnitOfMeasureQty, Tag 1316
func (m DerivativeSecurityListRequest) GetDerivativePriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativePriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSettlMethod gets DerivativeSettlMethod, Tag 1317
func (m DerivativeSecurityListRequest) GetDerivativeSettlMethod() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSettlMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativePriceQuoteMethod gets DerivativePriceQuoteMethod, Tag 1318
func (m DerivativeSecurityListRequest) GetDerivativePriceQuoteMethod() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativePriceQuoteMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeValuationMethod gets DerivativeValuationMethod, Tag 1319
func (m DerivativeSecurityListRequest) GetDerivativeValuationMethod() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeValuationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeListMethod gets DerivativeListMethod, Tag 1320
func (m DerivativeSecurityListRequest) GetDerivativeListMethod() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeListMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeCapPrice gets DerivativeCapPrice, Tag 1321
func (m DerivativeSecurityListRequest) GetDerivativeCapPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeCapPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeFloorPrice gets DerivativeFloorPrice, Tag 1322
func (m DerivativeSecurityListRequest) GetDerivativeFloorPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeFloorPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativePutOrCall gets DerivativePutOrCall, Tag 1323
func (m DerivativeSecurityListRequest) GetDerivativePutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativePutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingExerciseStyle gets UnderlyingExerciseStyle, Tag 1419
func (m DerivativeSecurityListRequest) GetUnderlyingExerciseStyle() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingUnitOfMeasureQty gets UnderlyingUnitOfMeasureQty, Tag 1423
func (m DerivativeSecurityListRequest) GetUnderlyingUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingPriceUnitOfMeasure gets UnderlyingPriceUnitOfMeasure, Tag 1424
func (m DerivativeSecurityListRequest) GetUnderlyingPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingPriceUnitOfMeasureQty gets UnderlyingPriceUnitOfMeasureQty, Tag 1425
func (m DerivativeSecurityListRequest) GetUnderlyingPriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingContractMultiplierUnit gets UnderlyingContractMultiplierUnit, Tag 1437
func (m DerivativeSecurityListRequest) GetUnderlyingContractMultiplierUnit() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingContractMultiplierUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeContractMultiplierUnit gets DerivativeContractMultiplierUnit, Tag 1438
func (m DerivativeSecurityListRequest) GetDerivativeContractMultiplierUnit() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeContractMultiplierUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingFlowScheduleType gets UnderlyingFlowScheduleType, Tag 1441
func (m DerivativeSecurityListRequest) GetUnderlyingFlowScheduleType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingFlowScheduleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeFlowScheduleType gets DerivativeFlowScheduleType, Tag 1442
func (m DerivativeSecurityListRequest) GetDerivativeFlowScheduleType() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeFlowScheduleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRestructuringType gets UnderlyingRestructuringType, Tag 1453
func (m DerivativeSecurityListRequest) GetUnderlyingRestructuringType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingRestructuringTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSeniority gets UnderlyingSeniority, Tag 1454
func (m DerivativeSecurityListRequest) GetUnderlyingSeniority() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSeniorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingNotionalPercentageOutstanding gets UnderlyingNotionalPercentageOutstanding, Tag 1455
func (m DerivativeSecurityListRequest) GetUnderlyingNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingNotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingOriginalNotionalPercentageOutstanding gets UnderlyingOriginalNotionalPercentageOutstanding, Tag 1456
func (m DerivativeSecurityListRequest) GetUnderlyingOriginalNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingOriginalNotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingAttachmentPoint gets UnderlyingAttachmentPoint, Tag 1459
func (m DerivativeSecurityListRequest) GetUnderlyingAttachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAttachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingDetachmentPoint gets UnderlyingDetachmentPoint, Tag 1460
func (m DerivativeSecurityListRequest) GetUnderlyingDetachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingDetachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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

//HasDerivativeValuationMethod returns true if DerivativeValuationMethod is present, Tag 1319
func (m DerivativeSecurityListRequest) HasDerivativeValuationMethod() bool {
	return m.Has(tag.DerivativeValuationMethod)
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

//HasUnderlyingContractMultiplierUnit returns true if UnderlyingContractMultiplierUnit is present, Tag 1437
func (m DerivativeSecurityListRequest) HasUnderlyingContractMultiplierUnit() bool {
	return m.Has(tag.UnderlyingContractMultiplierUnit)
}

//HasDerivativeContractMultiplierUnit returns true if DerivativeContractMultiplierUnit is present, Tag 1438
func (m DerivativeSecurityListRequest) HasDerivativeContractMultiplierUnit() bool {
	return m.Has(tag.DerivativeContractMultiplierUnit)
}

//HasUnderlyingFlowScheduleType returns true if UnderlyingFlowScheduleType is present, Tag 1441
func (m DerivativeSecurityListRequest) HasUnderlyingFlowScheduleType() bool {
	return m.Has(tag.UnderlyingFlowScheduleType)
}

//HasDerivativeFlowScheduleType returns true if DerivativeFlowScheduleType is present, Tag 1442
func (m DerivativeSecurityListRequest) HasDerivativeFlowScheduleType() bool {
	return m.Has(tag.DerivativeFlowScheduleType)
}

//HasUnderlyingRestructuringType returns true if UnderlyingRestructuringType is present, Tag 1453
func (m DerivativeSecurityListRequest) HasUnderlyingRestructuringType() bool {
	return m.Has(tag.UnderlyingRestructuringType)
}

//HasUnderlyingSeniority returns true if UnderlyingSeniority is present, Tag 1454
func (m DerivativeSecurityListRequest) HasUnderlyingSeniority() bool {
	return m.Has(tag.UnderlyingSeniority)
}

//HasUnderlyingNotionalPercentageOutstanding returns true if UnderlyingNotionalPercentageOutstanding is present, Tag 1455
func (m DerivativeSecurityListRequest) HasUnderlyingNotionalPercentageOutstanding() bool {
	return m.Has(tag.UnderlyingNotionalPercentageOutstanding)
}

//HasUnderlyingOriginalNotionalPercentageOutstanding returns true if UnderlyingOriginalNotionalPercentageOutstanding is present, Tag 1456
func (m DerivativeSecurityListRequest) HasUnderlyingOriginalNotionalPercentageOutstanding() bool {
	return m.Has(tag.UnderlyingOriginalNotionalPercentageOutstanding)
}

//HasUnderlyingAttachmentPoint returns true if UnderlyingAttachmentPoint is present, Tag 1459
func (m DerivativeSecurityListRequest) HasUnderlyingAttachmentPoint() bool {
	return m.Has(tag.UnderlyingAttachmentPoint)
}

//HasUnderlyingDetachmentPoint returns true if UnderlyingDetachmentPoint is present, Tag 1460
func (m DerivativeSecurityListRequest) HasUnderlyingDetachmentPoint() bool {
	return m.Has(tag.UnderlyingDetachmentPoint)
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
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingInstrumentPartyIDSource gets UnderlyingInstrumentPartyIDSource, Tag 1060
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingInstrumentPartyRole gets UnderlyingInstrumentPartyRole, Tag 1061
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartyRoleField
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
	*quickfix.Group
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
func (m NoUndlyInstrumentPartySubIDs) GetUnderlyingInstrumentPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingInstrumentPartySubIDType gets UnderlyingInstrumentPartySubIDType, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) GetUnderlyingInstrumentPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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

//NoDerivativeSecurityAltID is a repeating group element, Tag 1218
type NoDerivativeSecurityAltID struct {
	*quickfix.Group
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
func (m NoDerivativeSecurityAltID) GetDerivativeSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityAltIDSource gets DerivativeSecurityAltIDSource, Tag 1220
func (m NoDerivativeSecurityAltID) GetDerivativeSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
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
func (m NoDerivativeEvents) SetDerivativeEventPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeEventPx(value, scale))
}

//SetDerivativeEventText sets DerivativeEventText, Tag 1291
func (m NoDerivativeEvents) SetDerivativeEventText(v string) {
	m.Set(field.NewDerivativeEventText(v))
}

//GetDerivativeEventType gets DerivativeEventType, Tag 1287
func (m NoDerivativeEvents) GetDerivativeEventType() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeEventTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeEventDate gets DerivativeEventDate, Tag 1288
func (m NoDerivativeEvents) GetDerivativeEventDate() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeEventDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeEventTime gets DerivativeEventTime, Tag 1289
func (m NoDerivativeEvents) GetDerivativeEventTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.DerivativeEventTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeEventPx gets DerivativeEventPx, Tag 1290
func (m NoDerivativeEvents) GetDerivativeEventPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeEventPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeEventText gets DerivativeEventText, Tag 1291
func (m NoDerivativeEvents) GetDerivativeEventText() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeEventTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
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
func (m NoDerivativeInstrumentParties) GetDerivativeInstrumentPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeInstrumentPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeInstrumentPartyIDSource gets DerivativeInstrumentPartyIDSource, Tag 1294
func (m NoDerivativeInstrumentParties) GetDerivativeInstrumentPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeInstrumentPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeInstrumentPartyRole gets DerivativeInstrumentPartyRole, Tag 1295
func (m NoDerivativeInstrumentParties) GetDerivativeInstrumentPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeInstrumentPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
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
func (m NoDerivativeInstrumentPartySubIDs) GetDerivativeInstrumentPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeInstrumentPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeInstrumentPartySubIDType gets DerivativeInstrumentPartySubIDType, Tag 1298
func (m NoDerivativeInstrumentPartySubIDs) GetDerivativeInstrumentPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeInstrumentPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
