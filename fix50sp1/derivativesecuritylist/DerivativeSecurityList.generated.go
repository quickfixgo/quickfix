package derivativesecuritylist

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//DerivativeSecurityList is the fix50sp1 DerivativeSecurityList type, MsgType = AA
type DerivativeSecurityList struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

//FromMessage creates a DerivativeSecurityList from a quickfix.Message instance
func FromMessage(m *quickfix.Message) DerivativeSecurityList {
	return DerivativeSecurityList{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m DerivativeSecurityList) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a DerivativeSecurityList initialized with the required fields for DerivativeSecurityList
func New(securityresponseid field.SecurityResponseIDField) (m DerivativeSecurityList) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("AA"))
	m.Set(securityresponseid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg DerivativeSecurityList, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "8", "AA", r
}

//SetNoRelatedSym sets NoRelatedSym, Tag 146
func (m DerivativeSecurityList) SetNoRelatedSym(f NoRelatedSymRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingCouponPaymentDate sets UnderlyingCouponPaymentDate, Tag 241
func (m DerivativeSecurityList) SetUnderlyingCouponPaymentDate(v string) {
	m.Set(field.NewUnderlyingCouponPaymentDate(v))
}

//SetUnderlyingIssueDate sets UnderlyingIssueDate, Tag 242
func (m DerivativeSecurityList) SetUnderlyingIssueDate(v string) {
	m.Set(field.NewUnderlyingIssueDate(v))
}

//SetUnderlyingRepoCollateralSecurityType sets UnderlyingRepoCollateralSecurityType, Tag 243
func (m DerivativeSecurityList) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.Set(field.NewUnderlyingRepoCollateralSecurityType(v))
}

//SetUnderlyingRepurchaseTerm sets UnderlyingRepurchaseTerm, Tag 244
func (m DerivativeSecurityList) SetUnderlyingRepurchaseTerm(v int) {
	m.Set(field.NewUnderlyingRepurchaseTerm(v))
}

//SetUnderlyingRepurchaseRate sets UnderlyingRepurchaseRate, Tag 245
func (m DerivativeSecurityList) SetUnderlyingRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingRepurchaseRate(value, scale))
}

//SetUnderlyingFactor sets UnderlyingFactor, Tag 246
func (m DerivativeSecurityList) SetUnderlyingFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFactor(value, scale))
}

//SetUnderlyingRedemptionDate sets UnderlyingRedemptionDate, Tag 247
func (m DerivativeSecurityList) SetUnderlyingRedemptionDate(v string) {
	m.Set(field.NewUnderlyingRedemptionDate(v))
}

//SetUnderlyingCreditRating sets UnderlyingCreditRating, Tag 256
func (m DerivativeSecurityList) SetUnderlyingCreditRating(v string) {
	m.Set(field.NewUnderlyingCreditRating(v))
}

//SetUnderlyingSecurityIDSource sets UnderlyingSecurityIDSource, Tag 305
func (m DerivativeSecurityList) SetUnderlyingSecurityIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityIDSource(v))
}

//SetUnderlyingIssuer sets UnderlyingIssuer, Tag 306
func (m DerivativeSecurityList) SetUnderlyingIssuer(v string) {
	m.Set(field.NewUnderlyingIssuer(v))
}

//SetUnderlyingSecurityDesc sets UnderlyingSecurityDesc, Tag 307
func (m DerivativeSecurityList) SetUnderlyingSecurityDesc(v string) {
	m.Set(field.NewUnderlyingSecurityDesc(v))
}

//SetUnderlyingSecurityExchange sets UnderlyingSecurityExchange, Tag 308
func (m DerivativeSecurityList) SetUnderlyingSecurityExchange(v string) {
	m.Set(field.NewUnderlyingSecurityExchange(v))
}

//SetUnderlyingSecurityID sets UnderlyingSecurityID, Tag 309
func (m DerivativeSecurityList) SetUnderlyingSecurityID(v string) {
	m.Set(field.NewUnderlyingSecurityID(v))
}

//SetUnderlyingSecurityType sets UnderlyingSecurityType, Tag 310
func (m DerivativeSecurityList) SetUnderlyingSecurityType(v string) {
	m.Set(field.NewUnderlyingSecurityType(v))
}

//SetUnderlyingSymbol sets UnderlyingSymbol, Tag 311
func (m DerivativeSecurityList) SetUnderlyingSymbol(v string) {
	m.Set(field.NewUnderlyingSymbol(v))
}

//SetUnderlyingSymbolSfx sets UnderlyingSymbolSfx, Tag 312
func (m DerivativeSecurityList) SetUnderlyingSymbolSfx(v string) {
	m.Set(field.NewUnderlyingSymbolSfx(v))
}

//SetUnderlyingMaturityMonthYear sets UnderlyingMaturityMonthYear, Tag 313
func (m DerivativeSecurityList) SetUnderlyingMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingMaturityMonthYear(v))
}

//SetUnderlyingPutOrCall sets UnderlyingPutOrCall, Tag 315
func (m DerivativeSecurityList) SetUnderlyingPutOrCall(v int) {
	m.Set(field.NewUnderlyingPutOrCall(v))
}

//SetUnderlyingStrikePrice sets UnderlyingStrikePrice, Tag 316
func (m DerivativeSecurityList) SetUnderlyingStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStrikePrice(value, scale))
}

//SetUnderlyingOptAttribute sets UnderlyingOptAttribute, Tag 317
func (m DerivativeSecurityList) SetUnderlyingOptAttribute(v string) {
	m.Set(field.NewUnderlyingOptAttribute(v))
}

//SetUnderlyingCurrency sets UnderlyingCurrency, Tag 318
func (m DerivativeSecurityList) SetUnderlyingCurrency(v string) {
	m.Set(field.NewUnderlyingCurrency(v))
}

//SetSecurityReqID sets SecurityReqID, Tag 320
func (m DerivativeSecurityList) SetSecurityReqID(v string) {
	m.Set(field.NewSecurityReqID(v))
}

//SetSecurityResponseID sets SecurityResponseID, Tag 322
func (m DerivativeSecurityList) SetSecurityResponseID(v string) {
	m.Set(field.NewSecurityResponseID(v))
}

//SetEncodedUnderlyingIssuerLen sets EncodedUnderlyingIssuerLen, Tag 362
func (m DerivativeSecurityList) SetEncodedUnderlyingIssuerLen(v int) {
	m.Set(field.NewEncodedUnderlyingIssuerLen(v))
}

//SetEncodedUnderlyingIssuer sets EncodedUnderlyingIssuer, Tag 363
func (m DerivativeSecurityList) SetEncodedUnderlyingIssuer(v string) {
	m.Set(field.NewEncodedUnderlyingIssuer(v))
}

//SetEncodedUnderlyingSecurityDescLen sets EncodedUnderlyingSecurityDescLen, Tag 364
func (m DerivativeSecurityList) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.Set(field.NewEncodedUnderlyingSecurityDescLen(v))
}

//SetEncodedUnderlyingSecurityDesc sets EncodedUnderlyingSecurityDesc, Tag 365
func (m DerivativeSecurityList) SetEncodedUnderlyingSecurityDesc(v string) {
	m.Set(field.NewEncodedUnderlyingSecurityDesc(v))
}

//SetTotNoRelatedSym sets TotNoRelatedSym, Tag 393
func (m DerivativeSecurityList) SetTotNoRelatedSym(v int) {
	m.Set(field.NewTotNoRelatedSym(v))
}

//SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435
func (m DerivativeSecurityList) SetUnderlyingCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCouponRate(value, scale))
}

//SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436
func (m DerivativeSecurityList) SetUnderlyingContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingContractMultiplier(value, scale))
}

//SetNoUnderlyingSecurityAltID sets NoUnderlyingSecurityAltID, Tag 457
func (m DerivativeSecurityList) SetNoUnderlyingSecurityAltID(f NoUnderlyingSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingProduct sets UnderlyingProduct, Tag 462
func (m DerivativeSecurityList) SetUnderlyingProduct(v int) {
	m.Set(field.NewUnderlyingProduct(v))
}

//SetUnderlyingCFICode sets UnderlyingCFICode, Tag 463
func (m DerivativeSecurityList) SetUnderlyingCFICode(v string) {
	m.Set(field.NewUnderlyingCFICode(v))
}

//SetUnderlyingMaturityDate sets UnderlyingMaturityDate, Tag 542
func (m DerivativeSecurityList) SetUnderlyingMaturityDate(v string) {
	m.Set(field.NewUnderlyingMaturityDate(v))
}

//SetSecurityRequestResult sets SecurityRequestResult, Tag 560
func (m DerivativeSecurityList) SetSecurityRequestResult(v enum.SecurityRequestResult) {
	m.Set(field.NewSecurityRequestResult(v))
}

//SetUnderlyingCountryOfIssue sets UnderlyingCountryOfIssue, Tag 592
func (m DerivativeSecurityList) SetUnderlyingCountryOfIssue(v string) {
	m.Set(field.NewUnderlyingCountryOfIssue(v))
}

//SetUnderlyingStateOrProvinceOfIssue sets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m DerivativeSecurityList) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.Set(field.NewUnderlyingStateOrProvinceOfIssue(v))
}

//SetUnderlyingLocaleOfIssue sets UnderlyingLocaleOfIssue, Tag 594
func (m DerivativeSecurityList) SetUnderlyingLocaleOfIssue(v string) {
	m.Set(field.NewUnderlyingLocaleOfIssue(v))
}

//SetUnderlyingInstrRegistry sets UnderlyingInstrRegistry, Tag 595
func (m DerivativeSecurityList) SetUnderlyingInstrRegistry(v string) {
	m.Set(field.NewUnderlyingInstrRegistry(v))
}

//SetUnderlyingSecuritySubType sets UnderlyingSecuritySubType, Tag 763
func (m DerivativeSecurityList) SetUnderlyingSecuritySubType(v string) {
	m.Set(field.NewUnderlyingSecuritySubType(v))
}

//SetUnderlyingPx sets UnderlyingPx, Tag 810
func (m DerivativeSecurityList) SetUnderlyingPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPx(value, scale))
}

//SetUnderlyingCPProgram sets UnderlyingCPProgram, Tag 877
func (m DerivativeSecurityList) SetUnderlyingCPProgram(v string) {
	m.Set(field.NewUnderlyingCPProgram(v))
}

//SetUnderlyingCPRegType sets UnderlyingCPRegType, Tag 878
func (m DerivativeSecurityList) SetUnderlyingCPRegType(v string) {
	m.Set(field.NewUnderlyingCPRegType(v))
}

//SetUnderlyingQty sets UnderlyingQty, Tag 879
func (m DerivativeSecurityList) SetUnderlyingQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingQty(value, scale))
}

//SetUnderlyingDirtyPrice sets UnderlyingDirtyPrice, Tag 882
func (m DerivativeSecurityList) SetUnderlyingDirtyPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDirtyPrice(value, scale))
}

//SetUnderlyingEndPrice sets UnderlyingEndPrice, Tag 883
func (m DerivativeSecurityList) SetUnderlyingEndPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndPrice(value, scale))
}

//SetUnderlyingStartValue sets UnderlyingStartValue, Tag 884
func (m DerivativeSecurityList) SetUnderlyingStartValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStartValue(value, scale))
}

//SetUnderlyingCurrentValue sets UnderlyingCurrentValue, Tag 885
func (m DerivativeSecurityList) SetUnderlyingCurrentValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCurrentValue(value, scale))
}

//SetUnderlyingEndValue sets UnderlyingEndValue, Tag 886
func (m DerivativeSecurityList) SetUnderlyingEndValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndValue(value, scale))
}

//SetNoUnderlyingStips sets NoUnderlyingStips, Tag 887
func (m DerivativeSecurityList) SetNoUnderlyingStips(f NoUnderlyingStipsRepeatingGroup) {
	m.SetGroup(f)
}

//SetLastFragment sets LastFragment, Tag 893
func (m DerivativeSecurityList) SetLastFragment(v bool) {
	m.Set(field.NewLastFragment(v))
}

//SetUnderlyingStrikeCurrency sets UnderlyingStrikeCurrency, Tag 941
func (m DerivativeSecurityList) SetUnderlyingStrikeCurrency(v string) {
	m.Set(field.NewUnderlyingStrikeCurrency(v))
}

//SetUnderlyingAllocationPercent sets UnderlyingAllocationPercent, Tag 972
func (m DerivativeSecurityList) SetUnderlyingAllocationPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAllocationPercent(value, scale))
}

//SetUnderlyingCashAmount sets UnderlyingCashAmount, Tag 973
func (m DerivativeSecurityList) SetUnderlyingCashAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCashAmount(value, scale))
}

//SetUnderlyingCashType sets UnderlyingCashType, Tag 974
func (m DerivativeSecurityList) SetUnderlyingCashType(v enum.UnderlyingCashType) {
	m.Set(field.NewUnderlyingCashType(v))
}

//SetUnderlyingSettlementType sets UnderlyingSettlementType, Tag 975
func (m DerivativeSecurityList) SetUnderlyingSettlementType(v enum.UnderlyingSettlementType) {
	m.Set(field.NewUnderlyingSettlementType(v))
}

//SetUnderlyingUnitOfMeasure sets UnderlyingUnitOfMeasure, Tag 998
func (m DerivativeSecurityList) SetUnderlyingUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingUnitOfMeasure(v))
}

//SetUnderlyingTimeUnit sets UnderlyingTimeUnit, Tag 1000
func (m DerivativeSecurityList) SetUnderlyingTimeUnit(v string) {
	m.Set(field.NewUnderlyingTimeUnit(v))
}

//SetUnderlyingCapValue sets UnderlyingCapValue, Tag 1038
func (m DerivativeSecurityList) SetUnderlyingCapValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCapValue(value, scale))
}

//SetUnderlyingSettlMethod sets UnderlyingSettlMethod, Tag 1039
func (m DerivativeSecurityList) SetUnderlyingSettlMethod(v string) {
	m.Set(field.NewUnderlyingSettlMethod(v))
}

//SetUnderlyingAdjustedQuantity sets UnderlyingAdjustedQuantity, Tag 1044
func (m DerivativeSecurityList) SetUnderlyingAdjustedQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAdjustedQuantity(value, scale))
}

//SetUnderlyingFXRate sets UnderlyingFXRate, Tag 1045
func (m DerivativeSecurityList) SetUnderlyingFXRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFXRate(value, scale))
}

//SetUnderlyingFXRateCalc sets UnderlyingFXRateCalc, Tag 1046
func (m DerivativeSecurityList) SetUnderlyingFXRateCalc(v enum.UnderlyingFXRateCalc) {
	m.Set(field.NewUnderlyingFXRateCalc(v))
}

//SetNoUndlyInstrumentParties sets NoUndlyInstrumentParties, Tag 1058
func (m DerivativeSecurityList) SetNoUndlyInstrumentParties(f NoUndlyInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetApplID sets ApplID, Tag 1180
func (m DerivativeSecurityList) SetApplID(v string) {
	m.Set(field.NewApplID(v))
}

//SetApplSeqNum sets ApplSeqNum, Tag 1181
func (m DerivativeSecurityList) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

//SetUnderlyingMaturityTime sets UnderlyingMaturityTime, Tag 1213
func (m DerivativeSecurityList) SetUnderlyingMaturityTime(v string) {
	m.Set(field.NewUnderlyingMaturityTime(v))
}

//SetDerivativeSymbol sets DerivativeSymbol, Tag 1214
func (m DerivativeSecurityList) SetDerivativeSymbol(v string) {
	m.Set(field.NewDerivativeSymbol(v))
}

//SetDerivativeSymbolSfx sets DerivativeSymbolSfx, Tag 1215
func (m DerivativeSecurityList) SetDerivativeSymbolSfx(v string) {
	m.Set(field.NewDerivativeSymbolSfx(v))
}

//SetDerivativeSecurityID sets DerivativeSecurityID, Tag 1216
func (m DerivativeSecurityList) SetDerivativeSecurityID(v string) {
	m.Set(field.NewDerivativeSecurityID(v))
}

//SetDerivativeSecurityIDSource sets DerivativeSecurityIDSource, Tag 1217
func (m DerivativeSecurityList) SetDerivativeSecurityIDSource(v string) {
	m.Set(field.NewDerivativeSecurityIDSource(v))
}

//SetNoDerivativeSecurityAltID sets NoDerivativeSecurityAltID, Tag 1218
func (m DerivativeSecurityList) SetNoDerivativeSecurityAltID(f NoDerivativeSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetDerivativeOptPayAmount sets DerivativeOptPayAmount, Tag 1225
func (m DerivativeSecurityList) SetDerivativeOptPayAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeOptPayAmount(value, scale))
}

//SetDerivativeProductComplex sets DerivativeProductComplex, Tag 1228
func (m DerivativeSecurityList) SetDerivativeProductComplex(v string) {
	m.Set(field.NewDerivativeProductComplex(v))
}

//SetDerivFlexProductEligibilityIndicator sets DerivFlexProductEligibilityIndicator, Tag 1243
func (m DerivativeSecurityList) SetDerivFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewDerivFlexProductEligibilityIndicator(v))
}

//SetDerivativeProduct sets DerivativeProduct, Tag 1246
func (m DerivativeSecurityList) SetDerivativeProduct(v int) {
	m.Set(field.NewDerivativeProduct(v))
}

//SetDerivativeSecurityGroup sets DerivativeSecurityGroup, Tag 1247
func (m DerivativeSecurityList) SetDerivativeSecurityGroup(v string) {
	m.Set(field.NewDerivativeSecurityGroup(v))
}

//SetDerivativeCFICode sets DerivativeCFICode, Tag 1248
func (m DerivativeSecurityList) SetDerivativeCFICode(v string) {
	m.Set(field.NewDerivativeCFICode(v))
}

//SetDerivativeSecurityType sets DerivativeSecurityType, Tag 1249
func (m DerivativeSecurityList) SetDerivativeSecurityType(v string) {
	m.Set(field.NewDerivativeSecurityType(v))
}

//SetDerivativeSecuritySubType sets DerivativeSecuritySubType, Tag 1250
func (m DerivativeSecurityList) SetDerivativeSecuritySubType(v string) {
	m.Set(field.NewDerivativeSecuritySubType(v))
}

//SetDerivativeMaturityMonthYear sets DerivativeMaturityMonthYear, Tag 1251
func (m DerivativeSecurityList) SetDerivativeMaturityMonthYear(v string) {
	m.Set(field.NewDerivativeMaturityMonthYear(v))
}

//SetDerivativeMaturityDate sets DerivativeMaturityDate, Tag 1252
func (m DerivativeSecurityList) SetDerivativeMaturityDate(v string) {
	m.Set(field.NewDerivativeMaturityDate(v))
}

//SetDerivativeMaturityTime sets DerivativeMaturityTime, Tag 1253
func (m DerivativeSecurityList) SetDerivativeMaturityTime(v string) {
	m.Set(field.NewDerivativeMaturityTime(v))
}

//SetDerivativeSettleOnOpenFlag sets DerivativeSettleOnOpenFlag, Tag 1254
func (m DerivativeSecurityList) SetDerivativeSettleOnOpenFlag(v string) {
	m.Set(field.NewDerivativeSettleOnOpenFlag(v))
}

//SetDerivativeInstrmtAssignmentMethod sets DerivativeInstrmtAssignmentMethod, Tag 1255
func (m DerivativeSecurityList) SetDerivativeInstrmtAssignmentMethod(v string) {
	m.Set(field.NewDerivativeInstrmtAssignmentMethod(v))
}

//SetDerivativeSecurityStatus sets DerivativeSecurityStatus, Tag 1256
func (m DerivativeSecurityList) SetDerivativeSecurityStatus(v string) {
	m.Set(field.NewDerivativeSecurityStatus(v))
}

//SetDerivativeInstrRegistry sets DerivativeInstrRegistry, Tag 1257
func (m DerivativeSecurityList) SetDerivativeInstrRegistry(v string) {
	m.Set(field.NewDerivativeInstrRegistry(v))
}

//SetDerivativeCountryOfIssue sets DerivativeCountryOfIssue, Tag 1258
func (m DerivativeSecurityList) SetDerivativeCountryOfIssue(v string) {
	m.Set(field.NewDerivativeCountryOfIssue(v))
}

//SetDerivativeStateOrProvinceOfIssue sets DerivativeStateOrProvinceOfIssue, Tag 1259
func (m DerivativeSecurityList) SetDerivativeStateOrProvinceOfIssue(v string) {
	m.Set(field.NewDerivativeStateOrProvinceOfIssue(v))
}

//SetDerivativeLocaleOfIssue sets DerivativeLocaleOfIssue, Tag 1260
func (m DerivativeSecurityList) SetDerivativeLocaleOfIssue(v string) {
	m.Set(field.NewDerivativeLocaleOfIssue(v))
}

//SetDerivativeStrikePrice sets DerivativeStrikePrice, Tag 1261
func (m DerivativeSecurityList) SetDerivativeStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeStrikePrice(value, scale))
}

//SetDerivativeStrikeCurrency sets DerivativeStrikeCurrency, Tag 1262
func (m DerivativeSecurityList) SetDerivativeStrikeCurrency(v string) {
	m.Set(field.NewDerivativeStrikeCurrency(v))
}

//SetDerivativeStrikeMultiplier sets DerivativeStrikeMultiplier, Tag 1263
func (m DerivativeSecurityList) SetDerivativeStrikeMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeStrikeMultiplier(value, scale))
}

//SetDerivativeStrikeValue sets DerivativeStrikeValue, Tag 1264
func (m DerivativeSecurityList) SetDerivativeStrikeValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeStrikeValue(value, scale))
}

//SetDerivativeOptAttribute sets DerivativeOptAttribute, Tag 1265
func (m DerivativeSecurityList) SetDerivativeOptAttribute(v string) {
	m.Set(field.NewDerivativeOptAttribute(v))
}

//SetDerivativeContractMultiplier sets DerivativeContractMultiplier, Tag 1266
func (m DerivativeSecurityList) SetDerivativeContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeContractMultiplier(value, scale))
}

//SetDerivativeMinPriceIncrement sets DerivativeMinPriceIncrement, Tag 1267
func (m DerivativeSecurityList) SetDerivativeMinPriceIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeMinPriceIncrement(value, scale))
}

//SetDerivativeMinPriceIncrementAmount sets DerivativeMinPriceIncrementAmount, Tag 1268
func (m DerivativeSecurityList) SetDerivativeMinPriceIncrementAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeMinPriceIncrementAmount(value, scale))
}

//SetDerivativeUnitOfMeasure sets DerivativeUnitOfMeasure, Tag 1269
func (m DerivativeSecurityList) SetDerivativeUnitOfMeasure(v string) {
	m.Set(field.NewDerivativeUnitOfMeasure(v))
}

//SetDerivativeUnitOfMeasureQty sets DerivativeUnitOfMeasureQty, Tag 1270
func (m DerivativeSecurityList) SetDerivativeUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeUnitOfMeasureQty(value, scale))
}

//SetDerivativeTimeUnit sets DerivativeTimeUnit, Tag 1271
func (m DerivativeSecurityList) SetDerivativeTimeUnit(v string) {
	m.Set(field.NewDerivativeTimeUnit(v))
}

//SetDerivativeSecurityExchange sets DerivativeSecurityExchange, Tag 1272
func (m DerivativeSecurityList) SetDerivativeSecurityExchange(v string) {
	m.Set(field.NewDerivativeSecurityExchange(v))
}

//SetDerivativePositionLimit sets DerivativePositionLimit, Tag 1273
func (m DerivativeSecurityList) SetDerivativePositionLimit(v int) {
	m.Set(field.NewDerivativePositionLimit(v))
}

//SetDerivativeNTPositionLimit sets DerivativeNTPositionLimit, Tag 1274
func (m DerivativeSecurityList) SetDerivativeNTPositionLimit(v int) {
	m.Set(field.NewDerivativeNTPositionLimit(v))
}

//SetDerivativeIssuer sets DerivativeIssuer, Tag 1275
func (m DerivativeSecurityList) SetDerivativeIssuer(v string) {
	m.Set(field.NewDerivativeIssuer(v))
}

//SetDerivativeIssueDate sets DerivativeIssueDate, Tag 1276
func (m DerivativeSecurityList) SetDerivativeIssueDate(v string) {
	m.Set(field.NewDerivativeIssueDate(v))
}

//SetDerivativeEncodedIssuerLen sets DerivativeEncodedIssuerLen, Tag 1277
func (m DerivativeSecurityList) SetDerivativeEncodedIssuerLen(v int) {
	m.Set(field.NewDerivativeEncodedIssuerLen(v))
}

//SetDerivativeEncodedIssuer sets DerivativeEncodedIssuer, Tag 1278
func (m DerivativeSecurityList) SetDerivativeEncodedIssuer(v string) {
	m.Set(field.NewDerivativeEncodedIssuer(v))
}

//SetDerivativeSecurityDesc sets DerivativeSecurityDesc, Tag 1279
func (m DerivativeSecurityList) SetDerivativeSecurityDesc(v string) {
	m.Set(field.NewDerivativeSecurityDesc(v))
}

//SetDerivativeEncodedSecurityDescLen sets DerivativeEncodedSecurityDescLen, Tag 1280
func (m DerivativeSecurityList) SetDerivativeEncodedSecurityDescLen(v int) {
	m.Set(field.NewDerivativeEncodedSecurityDescLen(v))
}

//SetDerivativeEncodedSecurityDesc sets DerivativeEncodedSecurityDesc, Tag 1281
func (m DerivativeSecurityList) SetDerivativeEncodedSecurityDesc(v string) {
	m.Set(field.NewDerivativeEncodedSecurityDesc(v))
}

//SetDerivativeSecurityXMLLen sets DerivativeSecurityXMLLen, Tag 1282
func (m DerivativeSecurityList) SetDerivativeSecurityXMLLen(v int) {
	m.Set(field.NewDerivativeSecurityXMLLen(v))
}

//SetDerivativeSecurityXML sets DerivativeSecurityXML, Tag 1283
func (m DerivativeSecurityList) SetDerivativeSecurityXML(v string) {
	m.Set(field.NewDerivativeSecurityXML(v))
}

//SetDerivativeSecurityXMLSchema sets DerivativeSecurityXMLSchema, Tag 1284
func (m DerivativeSecurityList) SetDerivativeSecurityXMLSchema(v string) {
	m.Set(field.NewDerivativeSecurityXMLSchema(v))
}

//SetDerivativeContractSettlMonth sets DerivativeContractSettlMonth, Tag 1285
func (m DerivativeSecurityList) SetDerivativeContractSettlMonth(v string) {
	m.Set(field.NewDerivativeContractSettlMonth(v))
}

//SetNoDerivativeEvents sets NoDerivativeEvents, Tag 1286
func (m DerivativeSecurityList) SetNoDerivativeEvents(f NoDerivativeEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoDerivativeInstrumentParties sets NoDerivativeInstrumentParties, Tag 1292
func (m DerivativeSecurityList) SetNoDerivativeInstrumentParties(f NoDerivativeInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetDerivativeExerciseStyle sets DerivativeExerciseStyle, Tag 1299
func (m DerivativeSecurityList) SetDerivativeExerciseStyle(v string) {
	m.Set(field.NewDerivativeExerciseStyle(v))
}

//SetNoMarketSegments sets NoMarketSegments, Tag 1310
func (m DerivativeSecurityList) SetNoMarketSegments(f NoMarketSegmentsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoDerivativeInstrAttrib sets NoDerivativeInstrAttrib, Tag 1311
func (m DerivativeSecurityList) SetNoDerivativeInstrAttrib(f NoDerivativeInstrAttribRepeatingGroup) {
	m.SetGroup(f)
}

//SetDerivativePriceUnitOfMeasure sets DerivativePriceUnitOfMeasure, Tag 1315
func (m DerivativeSecurityList) SetDerivativePriceUnitOfMeasure(v string) {
	m.Set(field.NewDerivativePriceUnitOfMeasure(v))
}

//SetDerivativePriceUnitOfMeasureQty sets DerivativePriceUnitOfMeasureQty, Tag 1316
func (m DerivativeSecurityList) SetDerivativePriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativePriceUnitOfMeasureQty(value, scale))
}

//SetDerivativeSettlMethod sets DerivativeSettlMethod, Tag 1317
func (m DerivativeSecurityList) SetDerivativeSettlMethod(v string) {
	m.Set(field.NewDerivativeSettlMethod(v))
}

//SetDerivativePriceQuoteMethod sets DerivativePriceQuoteMethod, Tag 1318
func (m DerivativeSecurityList) SetDerivativePriceQuoteMethod(v string) {
	m.Set(field.NewDerivativePriceQuoteMethod(v))
}

//SetDerivativeFuturesValuationMethod sets DerivativeFuturesValuationMethod, Tag 1319
func (m DerivativeSecurityList) SetDerivativeFuturesValuationMethod(v string) {
	m.Set(field.NewDerivativeFuturesValuationMethod(v))
}

//SetDerivativeListMethod sets DerivativeListMethod, Tag 1320
func (m DerivativeSecurityList) SetDerivativeListMethod(v int) {
	m.Set(field.NewDerivativeListMethod(v))
}

//SetDerivativeCapPrice sets DerivativeCapPrice, Tag 1321
func (m DerivativeSecurityList) SetDerivativeCapPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeCapPrice(value, scale))
}

//SetDerivativeFloorPrice sets DerivativeFloorPrice, Tag 1322
func (m DerivativeSecurityList) SetDerivativeFloorPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewDerivativeFloorPrice(value, scale))
}

//SetDerivativePutOrCall sets DerivativePutOrCall, Tag 1323
func (m DerivativeSecurityList) SetDerivativePutOrCall(v int) {
	m.Set(field.NewDerivativePutOrCall(v))
}

//SetApplLastSeqNum sets ApplLastSeqNum, Tag 1350
func (m DerivativeSecurityList) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

//SetApplResendFlag sets ApplResendFlag, Tag 1352
func (m DerivativeSecurityList) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

//SetUnderlyingExerciseStyle sets UnderlyingExerciseStyle, Tag 1419
func (m DerivativeSecurityList) SetUnderlyingExerciseStyle(v int) {
	m.Set(field.NewUnderlyingExerciseStyle(v))
}

//SetUnderlyingUnitOfMeasureQty sets UnderlyingUnitOfMeasureQty, Tag 1423
func (m DerivativeSecurityList) SetUnderlyingUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingUnitOfMeasureQty(value, scale))
}

//SetUnderlyingPriceUnitOfMeasure sets UnderlyingPriceUnitOfMeasure, Tag 1424
func (m DerivativeSecurityList) SetUnderlyingPriceUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasure(v))
}

//SetUnderlyingPriceUnitOfMeasureQty sets UnderlyingPriceUnitOfMeasureQty, Tag 1425
func (m DerivativeSecurityList) SetUnderlyingPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasureQty(value, scale))
}

//GetNoRelatedSym gets NoRelatedSym, Tag 146
func (m DerivativeSecurityList) GetNoRelatedSym() (NoRelatedSymRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedSymRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingCouponPaymentDate gets UnderlyingCouponPaymentDate, Tag 241
func (m DerivativeSecurityList) GetUnderlyingCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingIssueDate gets UnderlyingIssueDate, Tag 242
func (m DerivativeSecurityList) GetUnderlyingIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRepoCollateralSecurityType gets UnderlyingRepoCollateralSecurityType, Tag 243
func (m DerivativeSecurityList) GetUnderlyingRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRepurchaseTerm gets UnderlyingRepurchaseTerm, Tag 244
func (m DerivativeSecurityList) GetUnderlyingRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRepurchaseRate gets UnderlyingRepurchaseRate, Tag 245
func (m DerivativeSecurityList) GetUnderlyingRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingFactor gets UnderlyingFactor, Tag 246
func (m DerivativeSecurityList) GetUnderlyingFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingFactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRedemptionDate gets UnderlyingRedemptionDate, Tag 247
func (m DerivativeSecurityList) GetUnderlyingRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCreditRating gets UnderlyingCreditRating, Tag 256
func (m DerivativeSecurityList) GetUnderlyingCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityIDSource gets UnderlyingSecurityIDSource, Tag 305
func (m DerivativeSecurityList) GetUnderlyingSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306
func (m DerivativeSecurityList) GetUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307
func (m DerivativeSecurityList) GetUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308
func (m DerivativeSecurityList) GetUnderlyingSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309
func (m DerivativeSecurityList) GetUnderlyingSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310
func (m DerivativeSecurityList) GetUnderlyingSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311
func (m DerivativeSecurityList) GetUnderlyingSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312
func (m DerivativeSecurityList) GetUnderlyingSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313
func (m DerivativeSecurityList) GetUnderlyingMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingPutOrCall gets UnderlyingPutOrCall, Tag 315
func (m DerivativeSecurityList) GetUnderlyingPutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingPutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316
func (m DerivativeSecurityList) GetUnderlyingStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317
func (m DerivativeSecurityList) GetUnderlyingOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCurrency gets UnderlyingCurrency, Tag 318
func (m DerivativeSecurityList) GetUnderlyingCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityReqID gets SecurityReqID, Tag 320
func (m DerivativeSecurityList) GetSecurityReqID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityResponseID gets SecurityResponseID, Tag 322
func (m DerivativeSecurityList) GetSecurityResponseID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityResponseIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362
func (m DerivativeSecurityList) GetEncodedUnderlyingIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363
func (m DerivativeSecurityList) GetEncodedUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364
func (m DerivativeSecurityList) GetEncodedUnderlyingSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365
func (m DerivativeSecurityList) GetEncodedUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTotNoRelatedSym gets TotNoRelatedSym, Tag 393
func (m DerivativeSecurityList) GetTotNoRelatedSym() (v int, err quickfix.MessageRejectError) {
	var f field.TotNoRelatedSymField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435
func (m DerivativeSecurityList) GetUnderlyingCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436
func (m DerivativeSecurityList) GetUnderlyingContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoUnderlyingSecurityAltID gets NoUnderlyingSecurityAltID, Tag 457
func (m DerivativeSecurityList) GetNoUnderlyingSecurityAltID() (NoUnderlyingSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingProduct gets UnderlyingProduct, Tag 462
func (m DerivativeSecurityList) GetUnderlyingProduct() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCFICode gets UnderlyingCFICode, Tag 463
func (m DerivativeSecurityList) GetUnderlyingCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingMaturityDate gets UnderlyingMaturityDate, Tag 542
func (m DerivativeSecurityList) GetUnderlyingMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityRequestResult gets SecurityRequestResult, Tag 560
func (m DerivativeSecurityList) GetSecurityRequestResult() (v enum.SecurityRequestResult, err quickfix.MessageRejectError) {
	var f field.SecurityRequestResultField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCountryOfIssue gets UnderlyingCountryOfIssue, Tag 592
func (m DerivativeSecurityList) GetUnderlyingCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStateOrProvinceOfIssue gets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m DerivativeSecurityList) GetUnderlyingStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingLocaleOfIssue gets UnderlyingLocaleOfIssue, Tag 594
func (m DerivativeSecurityList) GetUnderlyingLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingInstrRegistry gets UnderlyingInstrRegistry, Tag 595
func (m DerivativeSecurityList) GetUnderlyingInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecuritySubType gets UnderlyingSecuritySubType, Tag 763
func (m DerivativeSecurityList) GetUnderlyingSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingPx gets UnderlyingPx, Tag 810
func (m DerivativeSecurityList) GetUnderlyingPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCPProgram gets UnderlyingCPProgram, Tag 877
func (m DerivativeSecurityList) GetUnderlyingCPProgram() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCPRegType gets UnderlyingCPRegType, Tag 878
func (m DerivativeSecurityList) GetUnderlyingCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingQty gets UnderlyingQty, Tag 879
func (m DerivativeSecurityList) GetUnderlyingQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingDirtyPrice gets UnderlyingDirtyPrice, Tag 882
func (m DerivativeSecurityList) GetUnderlyingDirtyPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingDirtyPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingEndPrice gets UnderlyingEndPrice, Tag 883
func (m DerivativeSecurityList) GetUnderlyingEndPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStartValue gets UnderlyingStartValue, Tag 884
func (m DerivativeSecurityList) GetUnderlyingStartValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStartValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCurrentValue gets UnderlyingCurrentValue, Tag 885
func (m DerivativeSecurityList) GetUnderlyingCurrentValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrentValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingEndValue gets UnderlyingEndValue, Tag 886
func (m DerivativeSecurityList) GetUnderlyingEndValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoUnderlyingStips gets NoUnderlyingStips, Tag 887
func (m DerivativeSecurityList) GetNoUnderlyingStips() (NoUnderlyingStipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingStipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLastFragment gets LastFragment, Tag 893
func (m DerivativeSecurityList) GetLastFragment() (v bool, err quickfix.MessageRejectError) {
	var f field.LastFragmentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStrikeCurrency gets UnderlyingStrikeCurrency, Tag 941
func (m DerivativeSecurityList) GetUnderlyingStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingAllocationPercent gets UnderlyingAllocationPercent, Tag 972
func (m DerivativeSecurityList) GetUnderlyingAllocationPercent() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAllocationPercentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCashAmount gets UnderlyingCashAmount, Tag 973
func (m DerivativeSecurityList) GetUnderlyingCashAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCashAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCashType gets UnderlyingCashType, Tag 974
func (m DerivativeSecurityList) GetUnderlyingCashType() (v enum.UnderlyingCashType, err quickfix.MessageRejectError) {
	var f field.UnderlyingCashTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSettlementType gets UnderlyingSettlementType, Tag 975
func (m DerivativeSecurityList) GetUnderlyingSettlementType() (v enum.UnderlyingSettlementType, err quickfix.MessageRejectError) {
	var f field.UnderlyingSettlementTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingUnitOfMeasure gets UnderlyingUnitOfMeasure, Tag 998
func (m DerivativeSecurityList) GetUnderlyingUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingTimeUnit gets UnderlyingTimeUnit, Tag 1000
func (m DerivativeSecurityList) GetUnderlyingTimeUnit() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingTimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCapValue gets UnderlyingCapValue, Tag 1038
func (m DerivativeSecurityList) GetUnderlyingCapValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCapValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSettlMethod gets UnderlyingSettlMethod, Tag 1039
func (m DerivativeSecurityList) GetUnderlyingSettlMethod() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSettlMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingAdjustedQuantity gets UnderlyingAdjustedQuantity, Tag 1044
func (m DerivativeSecurityList) GetUnderlyingAdjustedQuantity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAdjustedQuantityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingFXRate gets UnderlyingFXRate, Tag 1045
func (m DerivativeSecurityList) GetUnderlyingFXRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingFXRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingFXRateCalc gets UnderlyingFXRateCalc, Tag 1046
func (m DerivativeSecurityList) GetUnderlyingFXRateCalc() (v enum.UnderlyingFXRateCalc, err quickfix.MessageRejectError) {
	var f field.UnderlyingFXRateCalcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoUndlyInstrumentParties gets NoUndlyInstrumentParties, Tag 1058
func (m DerivativeSecurityList) GetNoUndlyInstrumentParties() (NoUndlyInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetApplID gets ApplID, Tag 1180
func (m DerivativeSecurityList) GetApplID() (v string, err quickfix.MessageRejectError) {
	var f field.ApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplSeqNum gets ApplSeqNum, Tag 1181
func (m DerivativeSecurityList) GetApplSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingMaturityTime gets UnderlyingMaturityTime, Tag 1213
func (m DerivativeSecurityList) GetUnderlyingMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSymbol gets DerivativeSymbol, Tag 1214
func (m DerivativeSecurityList) GetDerivativeSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSymbolSfx gets DerivativeSymbolSfx, Tag 1215
func (m DerivativeSecurityList) GetDerivativeSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityID gets DerivativeSecurityID, Tag 1216
func (m DerivativeSecurityList) GetDerivativeSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityIDSource gets DerivativeSecurityIDSource, Tag 1217
func (m DerivativeSecurityList) GetDerivativeSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoDerivativeSecurityAltID gets NoDerivativeSecurityAltID, Tag 1218
func (m DerivativeSecurityList) GetNoDerivativeSecurityAltID() (NoDerivativeSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoDerivativeSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDerivativeOptPayAmount gets DerivativeOptPayAmount, Tag 1225
func (m DerivativeSecurityList) GetDerivativeOptPayAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeOptPayAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeProductComplex gets DerivativeProductComplex, Tag 1228
func (m DerivativeSecurityList) GetDerivativeProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivFlexProductEligibilityIndicator gets DerivFlexProductEligibilityIndicator, Tag 1243
func (m DerivativeSecurityList) GetDerivFlexProductEligibilityIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.DerivFlexProductEligibilityIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeProduct gets DerivativeProduct, Tag 1246
func (m DerivativeSecurityList) GetDerivativeProduct() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityGroup gets DerivativeSecurityGroup, Tag 1247
func (m DerivativeSecurityList) GetDerivativeSecurityGroup() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityGroupField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeCFICode gets DerivativeCFICode, Tag 1248
func (m DerivativeSecurityList) GetDerivativeCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityType gets DerivativeSecurityType, Tag 1249
func (m DerivativeSecurityList) GetDerivativeSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecuritySubType gets DerivativeSecuritySubType, Tag 1250
func (m DerivativeSecurityList) GetDerivativeSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeMaturityMonthYear gets DerivativeMaturityMonthYear, Tag 1251
func (m DerivativeSecurityList) GetDerivativeMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeMaturityDate gets DerivativeMaturityDate, Tag 1252
func (m DerivativeSecurityList) GetDerivativeMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeMaturityTime gets DerivativeMaturityTime, Tag 1253
func (m DerivativeSecurityList) GetDerivativeMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSettleOnOpenFlag gets DerivativeSettleOnOpenFlag, Tag 1254
func (m DerivativeSecurityList) GetDerivativeSettleOnOpenFlag() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSettleOnOpenFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeInstrmtAssignmentMethod gets DerivativeInstrmtAssignmentMethod, Tag 1255
func (m DerivativeSecurityList) GetDerivativeInstrmtAssignmentMethod() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeInstrmtAssignmentMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityStatus gets DerivativeSecurityStatus, Tag 1256
func (m DerivativeSecurityList) GetDerivativeSecurityStatus() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeInstrRegistry gets DerivativeInstrRegistry, Tag 1257
func (m DerivativeSecurityList) GetDerivativeInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeCountryOfIssue gets DerivativeCountryOfIssue, Tag 1258
func (m DerivativeSecurityList) GetDerivativeCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeStateOrProvinceOfIssue gets DerivativeStateOrProvinceOfIssue, Tag 1259
func (m DerivativeSecurityList) GetDerivativeStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeLocaleOfIssue gets DerivativeLocaleOfIssue, Tag 1260
func (m DerivativeSecurityList) GetDerivativeLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeStrikePrice gets DerivativeStrikePrice, Tag 1261
func (m DerivativeSecurityList) GetDerivativeStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeStrikeCurrency gets DerivativeStrikeCurrency, Tag 1262
func (m DerivativeSecurityList) GetDerivativeStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeStrikeMultiplier gets DerivativeStrikeMultiplier, Tag 1263
func (m DerivativeSecurityList) GetDerivativeStrikeMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeStrikeMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeStrikeValue gets DerivativeStrikeValue, Tag 1264
func (m DerivativeSecurityList) GetDerivativeStrikeValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeStrikeValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeOptAttribute gets DerivativeOptAttribute, Tag 1265
func (m DerivativeSecurityList) GetDerivativeOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeContractMultiplier gets DerivativeContractMultiplier, Tag 1266
func (m DerivativeSecurityList) GetDerivativeContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeMinPriceIncrement gets DerivativeMinPriceIncrement, Tag 1267
func (m DerivativeSecurityList) GetDerivativeMinPriceIncrement() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeMinPriceIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeMinPriceIncrementAmount gets DerivativeMinPriceIncrementAmount, Tag 1268
func (m DerivativeSecurityList) GetDerivativeMinPriceIncrementAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeMinPriceIncrementAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeUnitOfMeasure gets DerivativeUnitOfMeasure, Tag 1269
func (m DerivativeSecurityList) GetDerivativeUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeUnitOfMeasureQty gets DerivativeUnitOfMeasureQty, Tag 1270
func (m DerivativeSecurityList) GetDerivativeUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeTimeUnit gets DerivativeTimeUnit, Tag 1271
func (m DerivativeSecurityList) GetDerivativeTimeUnit() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeTimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityExchange gets DerivativeSecurityExchange, Tag 1272
func (m DerivativeSecurityList) GetDerivativeSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativePositionLimit gets DerivativePositionLimit, Tag 1273
func (m DerivativeSecurityList) GetDerivativePositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativePositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeNTPositionLimit gets DerivativeNTPositionLimit, Tag 1274
func (m DerivativeSecurityList) GetDerivativeNTPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeNTPositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeIssuer gets DerivativeIssuer, Tag 1275
func (m DerivativeSecurityList) GetDerivativeIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeIssueDate gets DerivativeIssueDate, Tag 1276
func (m DerivativeSecurityList) GetDerivativeIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeEncodedIssuerLen gets DerivativeEncodedIssuerLen, Tag 1277
func (m DerivativeSecurityList) GetDerivativeEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeEncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeEncodedIssuer gets DerivativeEncodedIssuer, Tag 1278
func (m DerivativeSecurityList) GetDerivativeEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeEncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityDesc gets DerivativeSecurityDesc, Tag 1279
func (m DerivativeSecurityList) GetDerivativeSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeEncodedSecurityDescLen gets DerivativeEncodedSecurityDescLen, Tag 1280
func (m DerivativeSecurityList) GetDerivativeEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeEncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeEncodedSecurityDesc gets DerivativeEncodedSecurityDesc, Tag 1281
func (m DerivativeSecurityList) GetDerivativeEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeEncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityXMLLen gets DerivativeSecurityXMLLen, Tag 1282
func (m DerivativeSecurityList) GetDerivativeSecurityXMLLen() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityXMLLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityXML gets DerivativeSecurityXML, Tag 1283
func (m DerivativeSecurityList) GetDerivativeSecurityXML() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityXMLField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSecurityXMLSchema gets DerivativeSecurityXMLSchema, Tag 1284
func (m DerivativeSecurityList) GetDerivativeSecurityXMLSchema() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSecurityXMLSchemaField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeContractSettlMonth gets DerivativeContractSettlMonth, Tag 1285
func (m DerivativeSecurityList) GetDerivativeContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoDerivativeEvents gets NoDerivativeEvents, Tag 1286
func (m DerivativeSecurityList) GetNoDerivativeEvents() (NoDerivativeEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoDerivativeEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoDerivativeInstrumentParties gets NoDerivativeInstrumentParties, Tag 1292
func (m DerivativeSecurityList) GetNoDerivativeInstrumentParties() (NoDerivativeInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoDerivativeInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDerivativeExerciseStyle gets DerivativeExerciseStyle, Tag 1299
func (m DerivativeSecurityList) GetDerivativeExerciseStyle() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoMarketSegments gets NoMarketSegments, Tag 1310
func (m DerivativeSecurityList) GetNoMarketSegments() (NoMarketSegmentsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMarketSegmentsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoDerivativeInstrAttrib gets NoDerivativeInstrAttrib, Tag 1311
func (m DerivativeSecurityList) GetNoDerivativeInstrAttrib() (NoDerivativeInstrAttribRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoDerivativeInstrAttribRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDerivativePriceUnitOfMeasure gets DerivativePriceUnitOfMeasure, Tag 1315
func (m DerivativeSecurityList) GetDerivativePriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativePriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativePriceUnitOfMeasureQty gets DerivativePriceUnitOfMeasureQty, Tag 1316
func (m DerivativeSecurityList) GetDerivativePriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativePriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeSettlMethod gets DerivativeSettlMethod, Tag 1317
func (m DerivativeSecurityList) GetDerivativeSettlMethod() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeSettlMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativePriceQuoteMethod gets DerivativePriceQuoteMethod, Tag 1318
func (m DerivativeSecurityList) GetDerivativePriceQuoteMethod() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativePriceQuoteMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeFuturesValuationMethod gets DerivativeFuturesValuationMethod, Tag 1319
func (m DerivativeSecurityList) GetDerivativeFuturesValuationMethod() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeFuturesValuationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeListMethod gets DerivativeListMethod, Tag 1320
func (m DerivativeSecurityList) GetDerivativeListMethod() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeListMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeCapPrice gets DerivativeCapPrice, Tag 1321
func (m DerivativeSecurityList) GetDerivativeCapPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeCapPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeFloorPrice gets DerivativeFloorPrice, Tag 1322
func (m DerivativeSecurityList) GetDerivativeFloorPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DerivativeFloorPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativePutOrCall gets DerivativePutOrCall, Tag 1323
func (m DerivativeSecurityList) GetDerivativePutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativePutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplLastSeqNum gets ApplLastSeqNum, Tag 1350
func (m DerivativeSecurityList) GetApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplResendFlag gets ApplResendFlag, Tag 1352
func (m DerivativeSecurityList) GetApplResendFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.ApplResendFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingExerciseStyle gets UnderlyingExerciseStyle, Tag 1419
func (m DerivativeSecurityList) GetUnderlyingExerciseStyle() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingUnitOfMeasureQty gets UnderlyingUnitOfMeasureQty, Tag 1423
func (m DerivativeSecurityList) GetUnderlyingUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingPriceUnitOfMeasure gets UnderlyingPriceUnitOfMeasure, Tag 1424
func (m DerivativeSecurityList) GetUnderlyingPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingPriceUnitOfMeasureQty gets UnderlyingPriceUnitOfMeasureQty, Tag 1425
func (m DerivativeSecurityList) GetUnderlyingPriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasNoRelatedSym returns true if NoRelatedSym is present, Tag 146
func (m DerivativeSecurityList) HasNoRelatedSym() bool {
	return m.Has(tag.NoRelatedSym)
}

//HasUnderlyingCouponPaymentDate returns true if UnderlyingCouponPaymentDate is present, Tag 241
func (m DerivativeSecurityList) HasUnderlyingCouponPaymentDate() bool {
	return m.Has(tag.UnderlyingCouponPaymentDate)
}

//HasUnderlyingIssueDate returns true if UnderlyingIssueDate is present, Tag 242
func (m DerivativeSecurityList) HasUnderlyingIssueDate() bool {
	return m.Has(tag.UnderlyingIssueDate)
}

//HasUnderlyingRepoCollateralSecurityType returns true if UnderlyingRepoCollateralSecurityType is present, Tag 243
func (m DerivativeSecurityList) HasUnderlyingRepoCollateralSecurityType() bool {
	return m.Has(tag.UnderlyingRepoCollateralSecurityType)
}

//HasUnderlyingRepurchaseTerm returns true if UnderlyingRepurchaseTerm is present, Tag 244
func (m DerivativeSecurityList) HasUnderlyingRepurchaseTerm() bool {
	return m.Has(tag.UnderlyingRepurchaseTerm)
}

//HasUnderlyingRepurchaseRate returns true if UnderlyingRepurchaseRate is present, Tag 245
func (m DerivativeSecurityList) HasUnderlyingRepurchaseRate() bool {
	return m.Has(tag.UnderlyingRepurchaseRate)
}

//HasUnderlyingFactor returns true if UnderlyingFactor is present, Tag 246
func (m DerivativeSecurityList) HasUnderlyingFactor() bool {
	return m.Has(tag.UnderlyingFactor)
}

//HasUnderlyingRedemptionDate returns true if UnderlyingRedemptionDate is present, Tag 247
func (m DerivativeSecurityList) HasUnderlyingRedemptionDate() bool {
	return m.Has(tag.UnderlyingRedemptionDate)
}

//HasUnderlyingCreditRating returns true if UnderlyingCreditRating is present, Tag 256
func (m DerivativeSecurityList) HasUnderlyingCreditRating() bool {
	return m.Has(tag.UnderlyingCreditRating)
}

//HasUnderlyingSecurityIDSource returns true if UnderlyingSecurityIDSource is present, Tag 305
func (m DerivativeSecurityList) HasUnderlyingSecurityIDSource() bool {
	return m.Has(tag.UnderlyingSecurityIDSource)
}

//HasUnderlyingIssuer returns true if UnderlyingIssuer is present, Tag 306
func (m DerivativeSecurityList) HasUnderlyingIssuer() bool {
	return m.Has(tag.UnderlyingIssuer)
}

//HasUnderlyingSecurityDesc returns true if UnderlyingSecurityDesc is present, Tag 307
func (m DerivativeSecurityList) HasUnderlyingSecurityDesc() bool {
	return m.Has(tag.UnderlyingSecurityDesc)
}

//HasUnderlyingSecurityExchange returns true if UnderlyingSecurityExchange is present, Tag 308
func (m DerivativeSecurityList) HasUnderlyingSecurityExchange() bool {
	return m.Has(tag.UnderlyingSecurityExchange)
}

//HasUnderlyingSecurityID returns true if UnderlyingSecurityID is present, Tag 309
func (m DerivativeSecurityList) HasUnderlyingSecurityID() bool {
	return m.Has(tag.UnderlyingSecurityID)
}

//HasUnderlyingSecurityType returns true if UnderlyingSecurityType is present, Tag 310
func (m DerivativeSecurityList) HasUnderlyingSecurityType() bool {
	return m.Has(tag.UnderlyingSecurityType)
}

//HasUnderlyingSymbol returns true if UnderlyingSymbol is present, Tag 311
func (m DerivativeSecurityList) HasUnderlyingSymbol() bool {
	return m.Has(tag.UnderlyingSymbol)
}

//HasUnderlyingSymbolSfx returns true if UnderlyingSymbolSfx is present, Tag 312
func (m DerivativeSecurityList) HasUnderlyingSymbolSfx() bool {
	return m.Has(tag.UnderlyingSymbolSfx)
}

//HasUnderlyingMaturityMonthYear returns true if UnderlyingMaturityMonthYear is present, Tag 313
func (m DerivativeSecurityList) HasUnderlyingMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingMaturityMonthYear)
}

//HasUnderlyingPutOrCall returns true if UnderlyingPutOrCall is present, Tag 315
func (m DerivativeSecurityList) HasUnderlyingPutOrCall() bool {
	return m.Has(tag.UnderlyingPutOrCall)
}

//HasUnderlyingStrikePrice returns true if UnderlyingStrikePrice is present, Tag 316
func (m DerivativeSecurityList) HasUnderlyingStrikePrice() bool {
	return m.Has(tag.UnderlyingStrikePrice)
}

//HasUnderlyingOptAttribute returns true if UnderlyingOptAttribute is present, Tag 317
func (m DerivativeSecurityList) HasUnderlyingOptAttribute() bool {
	return m.Has(tag.UnderlyingOptAttribute)
}

//HasUnderlyingCurrency returns true if UnderlyingCurrency is present, Tag 318
func (m DerivativeSecurityList) HasUnderlyingCurrency() bool {
	return m.Has(tag.UnderlyingCurrency)
}

//HasSecurityReqID returns true if SecurityReqID is present, Tag 320
func (m DerivativeSecurityList) HasSecurityReqID() bool {
	return m.Has(tag.SecurityReqID)
}

//HasSecurityResponseID returns true if SecurityResponseID is present, Tag 322
func (m DerivativeSecurityList) HasSecurityResponseID() bool {
	return m.Has(tag.SecurityResponseID)
}

//HasEncodedUnderlyingIssuerLen returns true if EncodedUnderlyingIssuerLen is present, Tag 362
func (m DerivativeSecurityList) HasEncodedUnderlyingIssuerLen() bool {
	return m.Has(tag.EncodedUnderlyingIssuerLen)
}

//HasEncodedUnderlyingIssuer returns true if EncodedUnderlyingIssuer is present, Tag 363
func (m DerivativeSecurityList) HasEncodedUnderlyingIssuer() bool {
	return m.Has(tag.EncodedUnderlyingIssuer)
}

//HasEncodedUnderlyingSecurityDescLen returns true if EncodedUnderlyingSecurityDescLen is present, Tag 364
func (m DerivativeSecurityList) HasEncodedUnderlyingSecurityDescLen() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDescLen)
}

//HasEncodedUnderlyingSecurityDesc returns true if EncodedUnderlyingSecurityDesc is present, Tag 365
func (m DerivativeSecurityList) HasEncodedUnderlyingSecurityDesc() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDesc)
}

//HasTotNoRelatedSym returns true if TotNoRelatedSym is present, Tag 393
func (m DerivativeSecurityList) HasTotNoRelatedSym() bool {
	return m.Has(tag.TotNoRelatedSym)
}

//HasUnderlyingCouponRate returns true if UnderlyingCouponRate is present, Tag 435
func (m DerivativeSecurityList) HasUnderlyingCouponRate() bool {
	return m.Has(tag.UnderlyingCouponRate)
}

//HasUnderlyingContractMultiplier returns true if UnderlyingContractMultiplier is present, Tag 436
func (m DerivativeSecurityList) HasUnderlyingContractMultiplier() bool {
	return m.Has(tag.UnderlyingContractMultiplier)
}

//HasNoUnderlyingSecurityAltID returns true if NoUnderlyingSecurityAltID is present, Tag 457
func (m DerivativeSecurityList) HasNoUnderlyingSecurityAltID() bool {
	return m.Has(tag.NoUnderlyingSecurityAltID)
}

//HasUnderlyingProduct returns true if UnderlyingProduct is present, Tag 462
func (m DerivativeSecurityList) HasUnderlyingProduct() bool {
	return m.Has(tag.UnderlyingProduct)
}

//HasUnderlyingCFICode returns true if UnderlyingCFICode is present, Tag 463
func (m DerivativeSecurityList) HasUnderlyingCFICode() bool {
	return m.Has(tag.UnderlyingCFICode)
}

//HasUnderlyingMaturityDate returns true if UnderlyingMaturityDate is present, Tag 542
func (m DerivativeSecurityList) HasUnderlyingMaturityDate() bool {
	return m.Has(tag.UnderlyingMaturityDate)
}

//HasSecurityRequestResult returns true if SecurityRequestResult is present, Tag 560
func (m DerivativeSecurityList) HasSecurityRequestResult() bool {
	return m.Has(tag.SecurityRequestResult)
}

//HasUnderlyingCountryOfIssue returns true if UnderlyingCountryOfIssue is present, Tag 592
func (m DerivativeSecurityList) HasUnderlyingCountryOfIssue() bool {
	return m.Has(tag.UnderlyingCountryOfIssue)
}

//HasUnderlyingStateOrProvinceOfIssue returns true if UnderlyingStateOrProvinceOfIssue is present, Tag 593
func (m DerivativeSecurityList) HasUnderlyingStateOrProvinceOfIssue() bool {
	return m.Has(tag.UnderlyingStateOrProvinceOfIssue)
}

//HasUnderlyingLocaleOfIssue returns true if UnderlyingLocaleOfIssue is present, Tag 594
func (m DerivativeSecurityList) HasUnderlyingLocaleOfIssue() bool {
	return m.Has(tag.UnderlyingLocaleOfIssue)
}

//HasUnderlyingInstrRegistry returns true if UnderlyingInstrRegistry is present, Tag 595
func (m DerivativeSecurityList) HasUnderlyingInstrRegistry() bool {
	return m.Has(tag.UnderlyingInstrRegistry)
}

//HasUnderlyingSecuritySubType returns true if UnderlyingSecuritySubType is present, Tag 763
func (m DerivativeSecurityList) HasUnderlyingSecuritySubType() bool {
	return m.Has(tag.UnderlyingSecuritySubType)
}

//HasUnderlyingPx returns true if UnderlyingPx is present, Tag 810
func (m DerivativeSecurityList) HasUnderlyingPx() bool {
	return m.Has(tag.UnderlyingPx)
}

//HasUnderlyingCPProgram returns true if UnderlyingCPProgram is present, Tag 877
func (m DerivativeSecurityList) HasUnderlyingCPProgram() bool {
	return m.Has(tag.UnderlyingCPProgram)
}

//HasUnderlyingCPRegType returns true if UnderlyingCPRegType is present, Tag 878
func (m DerivativeSecurityList) HasUnderlyingCPRegType() bool {
	return m.Has(tag.UnderlyingCPRegType)
}

//HasUnderlyingQty returns true if UnderlyingQty is present, Tag 879
func (m DerivativeSecurityList) HasUnderlyingQty() bool {
	return m.Has(tag.UnderlyingQty)
}

//HasUnderlyingDirtyPrice returns true if UnderlyingDirtyPrice is present, Tag 882
func (m DerivativeSecurityList) HasUnderlyingDirtyPrice() bool {
	return m.Has(tag.UnderlyingDirtyPrice)
}

//HasUnderlyingEndPrice returns true if UnderlyingEndPrice is present, Tag 883
func (m DerivativeSecurityList) HasUnderlyingEndPrice() bool {
	return m.Has(tag.UnderlyingEndPrice)
}

//HasUnderlyingStartValue returns true if UnderlyingStartValue is present, Tag 884
func (m DerivativeSecurityList) HasUnderlyingStartValue() bool {
	return m.Has(tag.UnderlyingStartValue)
}

//HasUnderlyingCurrentValue returns true if UnderlyingCurrentValue is present, Tag 885
func (m DerivativeSecurityList) HasUnderlyingCurrentValue() bool {
	return m.Has(tag.UnderlyingCurrentValue)
}

//HasUnderlyingEndValue returns true if UnderlyingEndValue is present, Tag 886
func (m DerivativeSecurityList) HasUnderlyingEndValue() bool {
	return m.Has(tag.UnderlyingEndValue)
}

//HasNoUnderlyingStips returns true if NoUnderlyingStips is present, Tag 887
func (m DerivativeSecurityList) HasNoUnderlyingStips() bool {
	return m.Has(tag.NoUnderlyingStips)
}

//HasLastFragment returns true if LastFragment is present, Tag 893
func (m DerivativeSecurityList) HasLastFragment() bool {
	return m.Has(tag.LastFragment)
}

//HasUnderlyingStrikeCurrency returns true if UnderlyingStrikeCurrency is present, Tag 941
func (m DerivativeSecurityList) HasUnderlyingStrikeCurrency() bool {
	return m.Has(tag.UnderlyingStrikeCurrency)
}

//HasUnderlyingAllocationPercent returns true if UnderlyingAllocationPercent is present, Tag 972
func (m DerivativeSecurityList) HasUnderlyingAllocationPercent() bool {
	return m.Has(tag.UnderlyingAllocationPercent)
}

//HasUnderlyingCashAmount returns true if UnderlyingCashAmount is present, Tag 973
func (m DerivativeSecurityList) HasUnderlyingCashAmount() bool {
	return m.Has(tag.UnderlyingCashAmount)
}

//HasUnderlyingCashType returns true if UnderlyingCashType is present, Tag 974
func (m DerivativeSecurityList) HasUnderlyingCashType() bool {
	return m.Has(tag.UnderlyingCashType)
}

//HasUnderlyingSettlementType returns true if UnderlyingSettlementType is present, Tag 975
func (m DerivativeSecurityList) HasUnderlyingSettlementType() bool {
	return m.Has(tag.UnderlyingSettlementType)
}

//HasUnderlyingUnitOfMeasure returns true if UnderlyingUnitOfMeasure is present, Tag 998
func (m DerivativeSecurityList) HasUnderlyingUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingUnitOfMeasure)
}

//HasUnderlyingTimeUnit returns true if UnderlyingTimeUnit is present, Tag 1000
func (m DerivativeSecurityList) HasUnderlyingTimeUnit() bool {
	return m.Has(tag.UnderlyingTimeUnit)
}

//HasUnderlyingCapValue returns true if UnderlyingCapValue is present, Tag 1038
func (m DerivativeSecurityList) HasUnderlyingCapValue() bool {
	return m.Has(tag.UnderlyingCapValue)
}

//HasUnderlyingSettlMethod returns true if UnderlyingSettlMethod is present, Tag 1039
func (m DerivativeSecurityList) HasUnderlyingSettlMethod() bool {
	return m.Has(tag.UnderlyingSettlMethod)
}

//HasUnderlyingAdjustedQuantity returns true if UnderlyingAdjustedQuantity is present, Tag 1044
func (m DerivativeSecurityList) HasUnderlyingAdjustedQuantity() bool {
	return m.Has(tag.UnderlyingAdjustedQuantity)
}

//HasUnderlyingFXRate returns true if UnderlyingFXRate is present, Tag 1045
func (m DerivativeSecurityList) HasUnderlyingFXRate() bool {
	return m.Has(tag.UnderlyingFXRate)
}

//HasUnderlyingFXRateCalc returns true if UnderlyingFXRateCalc is present, Tag 1046
func (m DerivativeSecurityList) HasUnderlyingFXRateCalc() bool {
	return m.Has(tag.UnderlyingFXRateCalc)
}

//HasNoUndlyInstrumentParties returns true if NoUndlyInstrumentParties is present, Tag 1058
func (m DerivativeSecurityList) HasNoUndlyInstrumentParties() bool {
	return m.Has(tag.NoUndlyInstrumentParties)
}

//HasApplID returns true if ApplID is present, Tag 1180
func (m DerivativeSecurityList) HasApplID() bool {
	return m.Has(tag.ApplID)
}

//HasApplSeqNum returns true if ApplSeqNum is present, Tag 1181
func (m DerivativeSecurityList) HasApplSeqNum() bool {
	return m.Has(tag.ApplSeqNum)
}

//HasUnderlyingMaturityTime returns true if UnderlyingMaturityTime is present, Tag 1213
func (m DerivativeSecurityList) HasUnderlyingMaturityTime() bool {
	return m.Has(tag.UnderlyingMaturityTime)
}

//HasDerivativeSymbol returns true if DerivativeSymbol is present, Tag 1214
func (m DerivativeSecurityList) HasDerivativeSymbol() bool {
	return m.Has(tag.DerivativeSymbol)
}

//HasDerivativeSymbolSfx returns true if DerivativeSymbolSfx is present, Tag 1215
func (m DerivativeSecurityList) HasDerivativeSymbolSfx() bool {
	return m.Has(tag.DerivativeSymbolSfx)
}

//HasDerivativeSecurityID returns true if DerivativeSecurityID is present, Tag 1216
func (m DerivativeSecurityList) HasDerivativeSecurityID() bool {
	return m.Has(tag.DerivativeSecurityID)
}

//HasDerivativeSecurityIDSource returns true if DerivativeSecurityIDSource is present, Tag 1217
func (m DerivativeSecurityList) HasDerivativeSecurityIDSource() bool {
	return m.Has(tag.DerivativeSecurityIDSource)
}

//HasNoDerivativeSecurityAltID returns true if NoDerivativeSecurityAltID is present, Tag 1218
func (m DerivativeSecurityList) HasNoDerivativeSecurityAltID() bool {
	return m.Has(tag.NoDerivativeSecurityAltID)
}

//HasDerivativeOptPayAmount returns true if DerivativeOptPayAmount is present, Tag 1225
func (m DerivativeSecurityList) HasDerivativeOptPayAmount() bool {
	return m.Has(tag.DerivativeOptPayAmount)
}

//HasDerivativeProductComplex returns true if DerivativeProductComplex is present, Tag 1228
func (m DerivativeSecurityList) HasDerivativeProductComplex() bool {
	return m.Has(tag.DerivativeProductComplex)
}

//HasDerivFlexProductEligibilityIndicator returns true if DerivFlexProductEligibilityIndicator is present, Tag 1243
func (m DerivativeSecurityList) HasDerivFlexProductEligibilityIndicator() bool {
	return m.Has(tag.DerivFlexProductEligibilityIndicator)
}

//HasDerivativeProduct returns true if DerivativeProduct is present, Tag 1246
func (m DerivativeSecurityList) HasDerivativeProduct() bool {
	return m.Has(tag.DerivativeProduct)
}

//HasDerivativeSecurityGroup returns true if DerivativeSecurityGroup is present, Tag 1247
func (m DerivativeSecurityList) HasDerivativeSecurityGroup() bool {
	return m.Has(tag.DerivativeSecurityGroup)
}

//HasDerivativeCFICode returns true if DerivativeCFICode is present, Tag 1248
func (m DerivativeSecurityList) HasDerivativeCFICode() bool {
	return m.Has(tag.DerivativeCFICode)
}

//HasDerivativeSecurityType returns true if DerivativeSecurityType is present, Tag 1249
func (m DerivativeSecurityList) HasDerivativeSecurityType() bool {
	return m.Has(tag.DerivativeSecurityType)
}

//HasDerivativeSecuritySubType returns true if DerivativeSecuritySubType is present, Tag 1250
func (m DerivativeSecurityList) HasDerivativeSecuritySubType() bool {
	return m.Has(tag.DerivativeSecuritySubType)
}

//HasDerivativeMaturityMonthYear returns true if DerivativeMaturityMonthYear is present, Tag 1251
func (m DerivativeSecurityList) HasDerivativeMaturityMonthYear() bool {
	return m.Has(tag.DerivativeMaturityMonthYear)
}

//HasDerivativeMaturityDate returns true if DerivativeMaturityDate is present, Tag 1252
func (m DerivativeSecurityList) HasDerivativeMaturityDate() bool {
	return m.Has(tag.DerivativeMaturityDate)
}

//HasDerivativeMaturityTime returns true if DerivativeMaturityTime is present, Tag 1253
func (m DerivativeSecurityList) HasDerivativeMaturityTime() bool {
	return m.Has(tag.DerivativeMaturityTime)
}

//HasDerivativeSettleOnOpenFlag returns true if DerivativeSettleOnOpenFlag is present, Tag 1254
func (m DerivativeSecurityList) HasDerivativeSettleOnOpenFlag() bool {
	return m.Has(tag.DerivativeSettleOnOpenFlag)
}

//HasDerivativeInstrmtAssignmentMethod returns true if DerivativeInstrmtAssignmentMethod is present, Tag 1255
func (m DerivativeSecurityList) HasDerivativeInstrmtAssignmentMethod() bool {
	return m.Has(tag.DerivativeInstrmtAssignmentMethod)
}

//HasDerivativeSecurityStatus returns true if DerivativeSecurityStatus is present, Tag 1256
func (m DerivativeSecurityList) HasDerivativeSecurityStatus() bool {
	return m.Has(tag.DerivativeSecurityStatus)
}

//HasDerivativeInstrRegistry returns true if DerivativeInstrRegistry is present, Tag 1257
func (m DerivativeSecurityList) HasDerivativeInstrRegistry() bool {
	return m.Has(tag.DerivativeInstrRegistry)
}

//HasDerivativeCountryOfIssue returns true if DerivativeCountryOfIssue is present, Tag 1258
func (m DerivativeSecurityList) HasDerivativeCountryOfIssue() bool {
	return m.Has(tag.DerivativeCountryOfIssue)
}

//HasDerivativeStateOrProvinceOfIssue returns true if DerivativeStateOrProvinceOfIssue is present, Tag 1259
func (m DerivativeSecurityList) HasDerivativeStateOrProvinceOfIssue() bool {
	return m.Has(tag.DerivativeStateOrProvinceOfIssue)
}

//HasDerivativeLocaleOfIssue returns true if DerivativeLocaleOfIssue is present, Tag 1260
func (m DerivativeSecurityList) HasDerivativeLocaleOfIssue() bool {
	return m.Has(tag.DerivativeLocaleOfIssue)
}

//HasDerivativeStrikePrice returns true if DerivativeStrikePrice is present, Tag 1261
func (m DerivativeSecurityList) HasDerivativeStrikePrice() bool {
	return m.Has(tag.DerivativeStrikePrice)
}

//HasDerivativeStrikeCurrency returns true if DerivativeStrikeCurrency is present, Tag 1262
func (m DerivativeSecurityList) HasDerivativeStrikeCurrency() bool {
	return m.Has(tag.DerivativeStrikeCurrency)
}

//HasDerivativeStrikeMultiplier returns true if DerivativeStrikeMultiplier is present, Tag 1263
func (m DerivativeSecurityList) HasDerivativeStrikeMultiplier() bool {
	return m.Has(tag.DerivativeStrikeMultiplier)
}

//HasDerivativeStrikeValue returns true if DerivativeStrikeValue is present, Tag 1264
func (m DerivativeSecurityList) HasDerivativeStrikeValue() bool {
	return m.Has(tag.DerivativeStrikeValue)
}

//HasDerivativeOptAttribute returns true if DerivativeOptAttribute is present, Tag 1265
func (m DerivativeSecurityList) HasDerivativeOptAttribute() bool {
	return m.Has(tag.DerivativeOptAttribute)
}

//HasDerivativeContractMultiplier returns true if DerivativeContractMultiplier is present, Tag 1266
func (m DerivativeSecurityList) HasDerivativeContractMultiplier() bool {
	return m.Has(tag.DerivativeContractMultiplier)
}

//HasDerivativeMinPriceIncrement returns true if DerivativeMinPriceIncrement is present, Tag 1267
func (m DerivativeSecurityList) HasDerivativeMinPriceIncrement() bool {
	return m.Has(tag.DerivativeMinPriceIncrement)
}

//HasDerivativeMinPriceIncrementAmount returns true if DerivativeMinPriceIncrementAmount is present, Tag 1268
func (m DerivativeSecurityList) HasDerivativeMinPriceIncrementAmount() bool {
	return m.Has(tag.DerivativeMinPriceIncrementAmount)
}

//HasDerivativeUnitOfMeasure returns true if DerivativeUnitOfMeasure is present, Tag 1269
func (m DerivativeSecurityList) HasDerivativeUnitOfMeasure() bool {
	return m.Has(tag.DerivativeUnitOfMeasure)
}

//HasDerivativeUnitOfMeasureQty returns true if DerivativeUnitOfMeasureQty is present, Tag 1270
func (m DerivativeSecurityList) HasDerivativeUnitOfMeasureQty() bool {
	return m.Has(tag.DerivativeUnitOfMeasureQty)
}

//HasDerivativeTimeUnit returns true if DerivativeTimeUnit is present, Tag 1271
func (m DerivativeSecurityList) HasDerivativeTimeUnit() bool {
	return m.Has(tag.DerivativeTimeUnit)
}

//HasDerivativeSecurityExchange returns true if DerivativeSecurityExchange is present, Tag 1272
func (m DerivativeSecurityList) HasDerivativeSecurityExchange() bool {
	return m.Has(tag.DerivativeSecurityExchange)
}

//HasDerivativePositionLimit returns true if DerivativePositionLimit is present, Tag 1273
func (m DerivativeSecurityList) HasDerivativePositionLimit() bool {
	return m.Has(tag.DerivativePositionLimit)
}

//HasDerivativeNTPositionLimit returns true if DerivativeNTPositionLimit is present, Tag 1274
func (m DerivativeSecurityList) HasDerivativeNTPositionLimit() bool {
	return m.Has(tag.DerivativeNTPositionLimit)
}

//HasDerivativeIssuer returns true if DerivativeIssuer is present, Tag 1275
func (m DerivativeSecurityList) HasDerivativeIssuer() bool {
	return m.Has(tag.DerivativeIssuer)
}

//HasDerivativeIssueDate returns true if DerivativeIssueDate is present, Tag 1276
func (m DerivativeSecurityList) HasDerivativeIssueDate() bool {
	return m.Has(tag.DerivativeIssueDate)
}

//HasDerivativeEncodedIssuerLen returns true if DerivativeEncodedIssuerLen is present, Tag 1277
func (m DerivativeSecurityList) HasDerivativeEncodedIssuerLen() bool {
	return m.Has(tag.DerivativeEncodedIssuerLen)
}

//HasDerivativeEncodedIssuer returns true if DerivativeEncodedIssuer is present, Tag 1278
func (m DerivativeSecurityList) HasDerivativeEncodedIssuer() bool {
	return m.Has(tag.DerivativeEncodedIssuer)
}

//HasDerivativeSecurityDesc returns true if DerivativeSecurityDesc is present, Tag 1279
func (m DerivativeSecurityList) HasDerivativeSecurityDesc() bool {
	return m.Has(tag.DerivativeSecurityDesc)
}

//HasDerivativeEncodedSecurityDescLen returns true if DerivativeEncodedSecurityDescLen is present, Tag 1280
func (m DerivativeSecurityList) HasDerivativeEncodedSecurityDescLen() bool {
	return m.Has(tag.DerivativeEncodedSecurityDescLen)
}

//HasDerivativeEncodedSecurityDesc returns true if DerivativeEncodedSecurityDesc is present, Tag 1281
func (m DerivativeSecurityList) HasDerivativeEncodedSecurityDesc() bool {
	return m.Has(tag.DerivativeEncodedSecurityDesc)
}

//HasDerivativeSecurityXMLLen returns true if DerivativeSecurityXMLLen is present, Tag 1282
func (m DerivativeSecurityList) HasDerivativeSecurityXMLLen() bool {
	return m.Has(tag.DerivativeSecurityXMLLen)
}

//HasDerivativeSecurityXML returns true if DerivativeSecurityXML is present, Tag 1283
func (m DerivativeSecurityList) HasDerivativeSecurityXML() bool {
	return m.Has(tag.DerivativeSecurityXML)
}

//HasDerivativeSecurityXMLSchema returns true if DerivativeSecurityXMLSchema is present, Tag 1284
func (m DerivativeSecurityList) HasDerivativeSecurityXMLSchema() bool {
	return m.Has(tag.DerivativeSecurityXMLSchema)
}

//HasDerivativeContractSettlMonth returns true if DerivativeContractSettlMonth is present, Tag 1285
func (m DerivativeSecurityList) HasDerivativeContractSettlMonth() bool {
	return m.Has(tag.DerivativeContractSettlMonth)
}

//HasNoDerivativeEvents returns true if NoDerivativeEvents is present, Tag 1286
func (m DerivativeSecurityList) HasNoDerivativeEvents() bool {
	return m.Has(tag.NoDerivativeEvents)
}

//HasNoDerivativeInstrumentParties returns true if NoDerivativeInstrumentParties is present, Tag 1292
func (m DerivativeSecurityList) HasNoDerivativeInstrumentParties() bool {
	return m.Has(tag.NoDerivativeInstrumentParties)
}

//HasDerivativeExerciseStyle returns true if DerivativeExerciseStyle is present, Tag 1299
func (m DerivativeSecurityList) HasDerivativeExerciseStyle() bool {
	return m.Has(tag.DerivativeExerciseStyle)
}

//HasNoMarketSegments returns true if NoMarketSegments is present, Tag 1310
func (m DerivativeSecurityList) HasNoMarketSegments() bool {
	return m.Has(tag.NoMarketSegments)
}

//HasNoDerivativeInstrAttrib returns true if NoDerivativeInstrAttrib is present, Tag 1311
func (m DerivativeSecurityList) HasNoDerivativeInstrAttrib() bool {
	return m.Has(tag.NoDerivativeInstrAttrib)
}

//HasDerivativePriceUnitOfMeasure returns true if DerivativePriceUnitOfMeasure is present, Tag 1315
func (m DerivativeSecurityList) HasDerivativePriceUnitOfMeasure() bool {
	return m.Has(tag.DerivativePriceUnitOfMeasure)
}

//HasDerivativePriceUnitOfMeasureQty returns true if DerivativePriceUnitOfMeasureQty is present, Tag 1316
func (m DerivativeSecurityList) HasDerivativePriceUnitOfMeasureQty() bool {
	return m.Has(tag.DerivativePriceUnitOfMeasureQty)
}

//HasDerivativeSettlMethod returns true if DerivativeSettlMethod is present, Tag 1317
func (m DerivativeSecurityList) HasDerivativeSettlMethod() bool {
	return m.Has(tag.DerivativeSettlMethod)
}

//HasDerivativePriceQuoteMethod returns true if DerivativePriceQuoteMethod is present, Tag 1318
func (m DerivativeSecurityList) HasDerivativePriceQuoteMethod() bool {
	return m.Has(tag.DerivativePriceQuoteMethod)
}

//HasDerivativeFuturesValuationMethod returns true if DerivativeFuturesValuationMethod is present, Tag 1319
func (m DerivativeSecurityList) HasDerivativeFuturesValuationMethod() bool {
	return m.Has(tag.DerivativeFuturesValuationMethod)
}

//HasDerivativeListMethod returns true if DerivativeListMethod is present, Tag 1320
func (m DerivativeSecurityList) HasDerivativeListMethod() bool {
	return m.Has(tag.DerivativeListMethod)
}

//HasDerivativeCapPrice returns true if DerivativeCapPrice is present, Tag 1321
func (m DerivativeSecurityList) HasDerivativeCapPrice() bool {
	return m.Has(tag.DerivativeCapPrice)
}

//HasDerivativeFloorPrice returns true if DerivativeFloorPrice is present, Tag 1322
func (m DerivativeSecurityList) HasDerivativeFloorPrice() bool {
	return m.Has(tag.DerivativeFloorPrice)
}

//HasDerivativePutOrCall returns true if DerivativePutOrCall is present, Tag 1323
func (m DerivativeSecurityList) HasDerivativePutOrCall() bool {
	return m.Has(tag.DerivativePutOrCall)
}

//HasApplLastSeqNum returns true if ApplLastSeqNum is present, Tag 1350
func (m DerivativeSecurityList) HasApplLastSeqNum() bool {
	return m.Has(tag.ApplLastSeqNum)
}

//HasApplResendFlag returns true if ApplResendFlag is present, Tag 1352
func (m DerivativeSecurityList) HasApplResendFlag() bool {
	return m.Has(tag.ApplResendFlag)
}

//HasUnderlyingExerciseStyle returns true if UnderlyingExerciseStyle is present, Tag 1419
func (m DerivativeSecurityList) HasUnderlyingExerciseStyle() bool {
	return m.Has(tag.UnderlyingExerciseStyle)
}

//HasUnderlyingUnitOfMeasureQty returns true if UnderlyingUnitOfMeasureQty is present, Tag 1423
func (m DerivativeSecurityList) HasUnderlyingUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingUnitOfMeasureQty)
}

//HasUnderlyingPriceUnitOfMeasure returns true if UnderlyingPriceUnitOfMeasure is present, Tag 1424
func (m DerivativeSecurityList) HasUnderlyingPriceUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasure)
}

//HasUnderlyingPriceUnitOfMeasureQty returns true if UnderlyingPriceUnitOfMeasureQty is present, Tag 1425
func (m DerivativeSecurityList) HasUnderlyingPriceUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasureQty)
}

//NoRelatedSym is a repeating group element, Tag 146
type NoRelatedSym struct {
	*quickfix.Group
}

//SetSymbol sets Symbol, Tag 55
func (m NoRelatedSym) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m NoRelatedSym) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m NoRelatedSym) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m NoRelatedSym) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m NoRelatedSym) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m NoRelatedSym) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m NoRelatedSym) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m NoRelatedSym) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m NoRelatedSym) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m NoRelatedSym) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m NoRelatedSym) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m NoRelatedSym) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m NoRelatedSym) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m NoRelatedSym) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m NoRelatedSym) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m NoRelatedSym) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

//SetFactor sets Factor, Tag 228
func (m NoRelatedSym) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

//SetCreditRating sets CreditRating, Tag 255
func (m NoRelatedSym) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m NoRelatedSym) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m NoRelatedSym) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m NoRelatedSym) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m NoRelatedSym) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m NoRelatedSym) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m NoRelatedSym) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m NoRelatedSym) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m NoRelatedSym) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m NoRelatedSym) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetCouponRate sets CouponRate, Tag 223
func (m NoRelatedSym) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m NoRelatedSym) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetIssuer sets Issuer, Tag 106
func (m NoRelatedSym) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m NoRelatedSym) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m NoRelatedSym) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m NoRelatedSym) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m NoRelatedSym) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m NoRelatedSym) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetPool sets Pool, Tag 691
func (m NoRelatedSym) SetPool(v string) {
	m.Set(field.NewPool(v))
}

//SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m NoRelatedSym) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

//SetCPProgram sets CPProgram, Tag 875
func (m NoRelatedSym) SetCPProgram(v enum.CPProgram) {
	m.Set(field.NewCPProgram(v))
}

//SetCPRegType sets CPRegType, Tag 876
func (m NoRelatedSym) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

//SetNoEvents sets NoEvents, Tag 864
func (m NoRelatedSym) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDatedDate sets DatedDate, Tag 873
func (m NoRelatedSym) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

//SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m NoRelatedSym) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

//SetSecurityStatus sets SecurityStatus, Tag 965
func (m NoRelatedSym) SetSecurityStatus(v enum.SecurityStatus) {
	m.Set(field.NewSecurityStatus(v))
}

//SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966
func (m NoRelatedSym) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

//SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049
func (m NoRelatedSym) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

//SetStrikeMultiplier sets StrikeMultiplier, Tag 967
func (m NoRelatedSym) SetStrikeMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeMultiplier(value, scale))
}

//SetStrikeValue sets StrikeValue, Tag 968
func (m NoRelatedSym) SetStrikeValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeValue(value, scale))
}

//SetMinPriceIncrement sets MinPriceIncrement, Tag 969
func (m NoRelatedSym) SetMinPriceIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrement(value, scale))
}

//SetPositionLimit sets PositionLimit, Tag 970
func (m NoRelatedSym) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

//SetNTPositionLimit sets NTPositionLimit, Tag 971
func (m NoRelatedSym) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

//SetNoInstrumentParties sets NoInstrumentParties, Tag 1018
func (m NoRelatedSym) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnitOfMeasure sets UnitOfMeasure, Tag 996
func (m NoRelatedSym) SetUnitOfMeasure(v enum.UnitOfMeasure) {
	m.Set(field.NewUnitOfMeasure(v))
}

//SetTimeUnit sets TimeUnit, Tag 997
func (m NoRelatedSym) SetTimeUnit(v enum.TimeUnit) {
	m.Set(field.NewTimeUnit(v))
}

//SetMaturityTime sets MaturityTime, Tag 1079
func (m NoRelatedSym) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

//SetSecurityGroup sets SecurityGroup, Tag 1151
func (m NoRelatedSym) SetSecurityGroup(v string) {
	m.Set(field.NewSecurityGroup(v))
}

//SetMinPriceIncrementAmount sets MinPriceIncrementAmount, Tag 1146
func (m NoRelatedSym) SetMinPriceIncrementAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrementAmount(value, scale))
}

//SetUnitOfMeasureQty sets UnitOfMeasureQty, Tag 1147
func (m NoRelatedSym) SetUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnitOfMeasureQty(value, scale))
}

//SetSecurityXMLLen sets SecurityXMLLen, Tag 1184
func (m NoRelatedSym) SetSecurityXMLLen(v int) {
	m.Set(field.NewSecurityXMLLen(v))
}

//SetSecurityXML sets SecurityXML, Tag 1185
func (m NoRelatedSym) SetSecurityXML(v string) {
	m.Set(field.NewSecurityXML(v))
}

//SetSecurityXMLSchema sets SecurityXMLSchema, Tag 1186
func (m NoRelatedSym) SetSecurityXMLSchema(v string) {
	m.Set(field.NewSecurityXMLSchema(v))
}

//SetProductComplex sets ProductComplex, Tag 1227
func (m NoRelatedSym) SetProductComplex(v string) {
	m.Set(field.NewProductComplex(v))
}

//SetPriceUnitOfMeasure sets PriceUnitOfMeasure, Tag 1191
func (m NoRelatedSym) SetPriceUnitOfMeasure(v string) {
	m.Set(field.NewPriceUnitOfMeasure(v))
}

//SetPriceUnitOfMeasureQty sets PriceUnitOfMeasureQty, Tag 1192
func (m NoRelatedSym) SetPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewPriceUnitOfMeasureQty(value, scale))
}

//SetSettlMethod sets SettlMethod, Tag 1193
func (m NoRelatedSym) SetSettlMethod(v enum.SettlMethod) {
	m.Set(field.NewSettlMethod(v))
}

//SetExerciseStyle sets ExerciseStyle, Tag 1194
func (m NoRelatedSym) SetExerciseStyle(v enum.ExerciseStyle) {
	m.Set(field.NewExerciseStyle(v))
}

//SetOptPayAmount sets OptPayAmount, Tag 1195
func (m NoRelatedSym) SetOptPayAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewOptPayAmount(value, scale))
}

//SetPriceQuoteMethod sets PriceQuoteMethod, Tag 1196
func (m NoRelatedSym) SetPriceQuoteMethod(v enum.PriceQuoteMethod) {
	m.Set(field.NewPriceQuoteMethod(v))
}

//SetListMethod sets ListMethod, Tag 1198
func (m NoRelatedSym) SetListMethod(v enum.ListMethod) {
	m.Set(field.NewListMethod(v))
}

//SetCapPrice sets CapPrice, Tag 1199
func (m NoRelatedSym) SetCapPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewCapPrice(value, scale))
}

//SetFloorPrice sets FloorPrice, Tag 1200
func (m NoRelatedSym) SetFloorPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewFloorPrice(value, scale))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m NoRelatedSym) SetPutOrCall(v enum.PutOrCall) {
	m.Set(field.NewPutOrCall(v))
}

//SetFlexibleIndicator sets FlexibleIndicator, Tag 1244
func (m NoRelatedSym) SetFlexibleIndicator(v bool) {
	m.Set(field.NewFlexibleIndicator(v))
}

//SetFlexProductEligibilityIndicator sets FlexProductEligibilityIndicator, Tag 1242
func (m NoRelatedSym) SetFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewFlexProductEligibilityIndicator(v))
}

//SetFuturesValuationMethod sets FuturesValuationMethod, Tag 1197
func (m NoRelatedSym) SetFuturesValuationMethod(v enum.FuturesValuationMethod) {
	m.Set(field.NewFuturesValuationMethod(v))
}

//SetCurrency sets Currency, Tag 15
func (m NoRelatedSym) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetDeliveryForm sets DeliveryForm, Tag 668
func (m NoRelatedSym) SetDeliveryForm(v enum.DeliveryForm) {
	m.Set(field.NewDeliveryForm(v))
}

//SetPctAtRisk sets PctAtRisk, Tag 869
func (m NoRelatedSym) SetPctAtRisk(value decimal.Decimal, scale int32) {
	m.Set(field.NewPctAtRisk(value, scale))
}

//SetNoInstrAttrib sets NoInstrAttrib, Tag 870
func (m NoRelatedSym) SetNoInstrAttrib(f NoInstrAttribRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoLegs sets NoLegs, Tag 555
func (m NoRelatedSym) SetNoLegs(f NoLegsRepeatingGroup) {
	m.SetGroup(f)
}

//SetText sets Text, Tag 58
func (m NoRelatedSym) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m NoRelatedSym) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m NoRelatedSym) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetSecondaryPriceLimitType sets SecondaryPriceLimitType, Tag 1305
func (m NoRelatedSym) SetSecondaryPriceLimitType(v int) {
	m.Set(field.NewSecondaryPriceLimitType(v))
}

//SetSecondaryLowLimitPrice sets SecondaryLowLimitPrice, Tag 1221
func (m NoRelatedSym) SetSecondaryLowLimitPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewSecondaryLowLimitPrice(value, scale))
}

//SetSecondaryHighLimitPrice sets SecondaryHighLimitPrice, Tag 1230
func (m NoRelatedSym) SetSecondaryHighLimitPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewSecondaryHighLimitPrice(value, scale))
}

//SetSecondaryTradingReferencePrice sets SecondaryTradingReferencePrice, Tag 1240
func (m NoRelatedSym) SetSecondaryTradingReferencePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewSecondaryTradingReferencePrice(value, scale))
}

//SetCorporateAction sets CorporateAction, Tag 292
func (m NoRelatedSym) SetCorporateAction(v enum.CorporateAction) {
	m.Set(field.NewCorporateAction(v))
}

//GetSymbol gets Symbol, Tag 55
func (m NoRelatedSym) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m NoRelatedSym) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m NoRelatedSym) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m NoRelatedSym) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m NoRelatedSym) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m NoRelatedSym) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCFICode gets CFICode, Tag 461
func (m NoRelatedSym) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m NoRelatedSym) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m NoRelatedSym) GetSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m NoRelatedSym) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m NoRelatedSym) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m NoRelatedSym) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m NoRelatedSym) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m NoRelatedSym) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m NoRelatedSym) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m NoRelatedSym) GetRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFactor gets Factor, Tag 228
func (m NoRelatedSym) GetFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m NoRelatedSym) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m NoRelatedSym) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m NoRelatedSym) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m NoRelatedSym) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m NoRelatedSym) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m NoRelatedSym) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m NoRelatedSym) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m NoRelatedSym) GetStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.StrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m NoRelatedSym) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m NoRelatedSym) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m NoRelatedSym) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m NoRelatedSym) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssuer gets Issuer, Tag 106
func (m NoRelatedSym) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m NoRelatedSym) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m NoRelatedSym) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m NoRelatedSym) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m NoRelatedSym) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m NoRelatedSym) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPool gets Pool, Tag 691
func (m NoRelatedSym) GetPool() (v string, err quickfix.MessageRejectError) {
	var f field.PoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m NoRelatedSym) GetContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.ContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCPProgram gets CPProgram, Tag 875
func (m NoRelatedSym) GetCPProgram() (v enum.CPProgram, err quickfix.MessageRejectError) {
	var f field.CPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCPRegType gets CPRegType, Tag 876
func (m NoRelatedSym) GetCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.CPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoEvents gets NoEvents, Tag 864
func (m NoRelatedSym) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDatedDate gets DatedDate, Tag 873
func (m NoRelatedSym) GetDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.DatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m NoRelatedSym) GetInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.InterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityStatus gets SecurityStatus, Tag 965
func (m NoRelatedSym) GetSecurityStatus() (v enum.SecurityStatus, err quickfix.MessageRejectError) {
	var f field.SecurityStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966
func (m NoRelatedSym) GetSettleOnOpenFlag() (v string, err quickfix.MessageRejectError) {
	var f field.SettleOnOpenFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049
func (m NoRelatedSym) GetInstrmtAssignmentMethod() (v string, err quickfix.MessageRejectError) {
	var f field.InstrmtAssignmentMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikeMultiplier gets StrikeMultiplier, Tag 967
func (m NoRelatedSym) GetStrikeMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikeMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikeValue gets StrikeValue, Tag 968
func (m NoRelatedSym) GetStrikeValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikeValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMinPriceIncrement gets MinPriceIncrement, Tag 969
func (m NoRelatedSym) GetMinPriceIncrement() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinPriceIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPositionLimit gets PositionLimit, Tag 970
func (m NoRelatedSym) GetPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.PositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNTPositionLimit gets NTPositionLimit, Tag 971
func (m NoRelatedSym) GetNTPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.NTPositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoInstrumentParties gets NoInstrumentParties, Tag 1018
func (m NoRelatedSym) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnitOfMeasure gets UnitOfMeasure, Tag 996
func (m NoRelatedSym) GetUnitOfMeasure() (v enum.UnitOfMeasure, err quickfix.MessageRejectError) {
	var f field.UnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTimeUnit gets TimeUnit, Tag 997
func (m NoRelatedSym) GetTimeUnit() (v enum.TimeUnit, err quickfix.MessageRejectError) {
	var f field.TimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityTime gets MaturityTime, Tag 1079
func (m NoRelatedSym) GetMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityGroup gets SecurityGroup, Tag 1151
func (m NoRelatedSym) GetSecurityGroup() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityGroupField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMinPriceIncrementAmount gets MinPriceIncrementAmount, Tag 1146
func (m NoRelatedSym) GetMinPriceIncrementAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinPriceIncrementAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnitOfMeasureQty gets UnitOfMeasureQty, Tag 1147
func (m NoRelatedSym) GetUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityXMLLen gets SecurityXMLLen, Tag 1184
func (m NoRelatedSym) GetSecurityXMLLen() (v int, err quickfix.MessageRejectError) {
	var f field.SecurityXMLLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityXML gets SecurityXML, Tag 1185
func (m NoRelatedSym) GetSecurityXML() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityXMLField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityXMLSchema gets SecurityXMLSchema, Tag 1186
func (m NoRelatedSym) GetSecurityXMLSchema() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityXMLSchemaField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetProductComplex gets ProductComplex, Tag 1227
func (m NoRelatedSym) GetProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.ProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPriceUnitOfMeasure gets PriceUnitOfMeasure, Tag 1191
func (m NoRelatedSym) GetPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.PriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPriceUnitOfMeasureQty gets PriceUnitOfMeasureQty, Tag 1192
func (m NoRelatedSym) GetPriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlMethod gets SettlMethod, Tag 1193
func (m NoRelatedSym) GetSettlMethod() (v enum.SettlMethod, err quickfix.MessageRejectError) {
	var f field.SettlMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExerciseStyle gets ExerciseStyle, Tag 1194
func (m NoRelatedSym) GetExerciseStyle() (v enum.ExerciseStyle, err quickfix.MessageRejectError) {
	var f field.ExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOptPayAmount gets OptPayAmount, Tag 1195
func (m NoRelatedSym) GetOptPayAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OptPayAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPriceQuoteMethod gets PriceQuoteMethod, Tag 1196
func (m NoRelatedSym) GetPriceQuoteMethod() (v enum.PriceQuoteMethod, err quickfix.MessageRejectError) {
	var f field.PriceQuoteMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetListMethod gets ListMethod, Tag 1198
func (m NoRelatedSym) GetListMethod() (v enum.ListMethod, err quickfix.MessageRejectError) {
	var f field.ListMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCapPrice gets CapPrice, Tag 1199
func (m NoRelatedSym) GetCapPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CapPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFloorPrice gets FloorPrice, Tag 1200
func (m NoRelatedSym) GetFloorPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FloorPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m NoRelatedSym) GetPutOrCall() (v enum.PutOrCall, err quickfix.MessageRejectError) {
	var f field.PutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFlexibleIndicator gets FlexibleIndicator, Tag 1244
func (m NoRelatedSym) GetFlexibleIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.FlexibleIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFlexProductEligibilityIndicator gets FlexProductEligibilityIndicator, Tag 1242
func (m NoRelatedSym) GetFlexProductEligibilityIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.FlexProductEligibilityIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFuturesValuationMethod gets FuturesValuationMethod, Tag 1197
func (m NoRelatedSym) GetFuturesValuationMethod() (v enum.FuturesValuationMethod, err quickfix.MessageRejectError) {
	var f field.FuturesValuationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCurrency gets Currency, Tag 15
func (m NoRelatedSym) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDeliveryForm gets DeliveryForm, Tag 668
func (m NoRelatedSym) GetDeliveryForm() (v enum.DeliveryForm, err quickfix.MessageRejectError) {
	var f field.DeliveryFormField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPctAtRisk gets PctAtRisk, Tag 869
func (m NoRelatedSym) GetPctAtRisk() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PctAtRiskField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoInstrAttrib gets NoInstrAttrib, Tag 870
func (m NoRelatedSym) GetNoInstrAttrib() (NoInstrAttribRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrAttribRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoLegs gets NoLegs, Tag 555
func (m NoRelatedSym) GetNoLegs() (NoLegsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetText gets Text, Tag 58
func (m NoRelatedSym) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m NoRelatedSym) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m NoRelatedSym) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecondaryPriceLimitType gets SecondaryPriceLimitType, Tag 1305
func (m NoRelatedSym) GetSecondaryPriceLimitType() (v int, err quickfix.MessageRejectError) {
	var f field.SecondaryPriceLimitTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecondaryLowLimitPrice gets SecondaryLowLimitPrice, Tag 1221
func (m NoRelatedSym) GetSecondaryLowLimitPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SecondaryLowLimitPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecondaryHighLimitPrice gets SecondaryHighLimitPrice, Tag 1230
func (m NoRelatedSym) GetSecondaryHighLimitPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SecondaryHighLimitPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecondaryTradingReferencePrice gets SecondaryTradingReferencePrice, Tag 1240
func (m NoRelatedSym) GetSecondaryTradingReferencePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SecondaryTradingReferencePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCorporateAction gets CorporateAction, Tag 292
func (m NoRelatedSym) GetCorporateAction() (v enum.CorporateAction, err quickfix.MessageRejectError) {
	var f field.CorporateActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m NoRelatedSym) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m NoRelatedSym) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m NoRelatedSym) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m NoRelatedSym) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m NoRelatedSym) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m NoRelatedSym) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m NoRelatedSym) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m NoRelatedSym) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m NoRelatedSym) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m NoRelatedSym) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m NoRelatedSym) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m NoRelatedSym) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m NoRelatedSym) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m NoRelatedSym) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m NoRelatedSym) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m NoRelatedSym) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m NoRelatedSym) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m NoRelatedSym) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m NoRelatedSym) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m NoRelatedSym) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m NoRelatedSym) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m NoRelatedSym) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m NoRelatedSym) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m NoRelatedSym) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m NoRelatedSym) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m NoRelatedSym) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m NoRelatedSym) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m NoRelatedSym) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m NoRelatedSym) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m NoRelatedSym) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m NoRelatedSym) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m NoRelatedSym) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m NoRelatedSym) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m NoRelatedSym) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m NoRelatedSym) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasPool returns true if Pool is present, Tag 691
func (m NoRelatedSym) HasPool() bool {
	return m.Has(tag.Pool)
}

//HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m NoRelatedSym) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

//HasCPProgram returns true if CPProgram is present, Tag 875
func (m NoRelatedSym) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

//HasCPRegType returns true if CPRegType is present, Tag 876
func (m NoRelatedSym) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

//HasNoEvents returns true if NoEvents is present, Tag 864
func (m NoRelatedSym) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

//HasDatedDate returns true if DatedDate is present, Tag 873
func (m NoRelatedSym) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

//HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m NoRelatedSym) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

//HasSecurityStatus returns true if SecurityStatus is present, Tag 965
func (m NoRelatedSym) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

//HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966
func (m NoRelatedSym) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

//HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049
func (m NoRelatedSym) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

//HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967
func (m NoRelatedSym) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

//HasStrikeValue returns true if StrikeValue is present, Tag 968
func (m NoRelatedSym) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

//HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969
func (m NoRelatedSym) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

//HasPositionLimit returns true if PositionLimit is present, Tag 970
func (m NoRelatedSym) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

//HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971
func (m NoRelatedSym) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

//HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018
func (m NoRelatedSym) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

//HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996
func (m NoRelatedSym) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

//HasTimeUnit returns true if TimeUnit is present, Tag 997
func (m NoRelatedSym) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

//HasMaturityTime returns true if MaturityTime is present, Tag 1079
func (m NoRelatedSym) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

//HasSecurityGroup returns true if SecurityGroup is present, Tag 1151
func (m NoRelatedSym) HasSecurityGroup() bool {
	return m.Has(tag.SecurityGroup)
}

//HasMinPriceIncrementAmount returns true if MinPriceIncrementAmount is present, Tag 1146
func (m NoRelatedSym) HasMinPriceIncrementAmount() bool {
	return m.Has(tag.MinPriceIncrementAmount)
}

//HasUnitOfMeasureQty returns true if UnitOfMeasureQty is present, Tag 1147
func (m NoRelatedSym) HasUnitOfMeasureQty() bool {
	return m.Has(tag.UnitOfMeasureQty)
}

//HasSecurityXMLLen returns true if SecurityXMLLen is present, Tag 1184
func (m NoRelatedSym) HasSecurityXMLLen() bool {
	return m.Has(tag.SecurityXMLLen)
}

//HasSecurityXML returns true if SecurityXML is present, Tag 1185
func (m NoRelatedSym) HasSecurityXML() bool {
	return m.Has(tag.SecurityXML)
}

//HasSecurityXMLSchema returns true if SecurityXMLSchema is present, Tag 1186
func (m NoRelatedSym) HasSecurityXMLSchema() bool {
	return m.Has(tag.SecurityXMLSchema)
}

//HasProductComplex returns true if ProductComplex is present, Tag 1227
func (m NoRelatedSym) HasProductComplex() bool {
	return m.Has(tag.ProductComplex)
}

//HasPriceUnitOfMeasure returns true if PriceUnitOfMeasure is present, Tag 1191
func (m NoRelatedSym) HasPriceUnitOfMeasure() bool {
	return m.Has(tag.PriceUnitOfMeasure)
}

//HasPriceUnitOfMeasureQty returns true if PriceUnitOfMeasureQty is present, Tag 1192
func (m NoRelatedSym) HasPriceUnitOfMeasureQty() bool {
	return m.Has(tag.PriceUnitOfMeasureQty)
}

//HasSettlMethod returns true if SettlMethod is present, Tag 1193
func (m NoRelatedSym) HasSettlMethod() bool {
	return m.Has(tag.SettlMethod)
}

//HasExerciseStyle returns true if ExerciseStyle is present, Tag 1194
func (m NoRelatedSym) HasExerciseStyle() bool {
	return m.Has(tag.ExerciseStyle)
}

//HasOptPayAmount returns true if OptPayAmount is present, Tag 1195
func (m NoRelatedSym) HasOptPayAmount() bool {
	return m.Has(tag.OptPayAmount)
}

//HasPriceQuoteMethod returns true if PriceQuoteMethod is present, Tag 1196
func (m NoRelatedSym) HasPriceQuoteMethod() bool {
	return m.Has(tag.PriceQuoteMethod)
}

//HasListMethod returns true if ListMethod is present, Tag 1198
func (m NoRelatedSym) HasListMethod() bool {
	return m.Has(tag.ListMethod)
}

//HasCapPrice returns true if CapPrice is present, Tag 1199
func (m NoRelatedSym) HasCapPrice() bool {
	return m.Has(tag.CapPrice)
}

//HasFloorPrice returns true if FloorPrice is present, Tag 1200
func (m NoRelatedSym) HasFloorPrice() bool {
	return m.Has(tag.FloorPrice)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m NoRelatedSym) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasFlexibleIndicator returns true if FlexibleIndicator is present, Tag 1244
func (m NoRelatedSym) HasFlexibleIndicator() bool {
	return m.Has(tag.FlexibleIndicator)
}

//HasFlexProductEligibilityIndicator returns true if FlexProductEligibilityIndicator is present, Tag 1242
func (m NoRelatedSym) HasFlexProductEligibilityIndicator() bool {
	return m.Has(tag.FlexProductEligibilityIndicator)
}

//HasFuturesValuationMethod returns true if FuturesValuationMethod is present, Tag 1197
func (m NoRelatedSym) HasFuturesValuationMethod() bool {
	return m.Has(tag.FuturesValuationMethod)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m NoRelatedSym) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasDeliveryForm returns true if DeliveryForm is present, Tag 668
func (m NoRelatedSym) HasDeliveryForm() bool {
	return m.Has(tag.DeliveryForm)
}

//HasPctAtRisk returns true if PctAtRisk is present, Tag 869
func (m NoRelatedSym) HasPctAtRisk() bool {
	return m.Has(tag.PctAtRisk)
}

//HasNoInstrAttrib returns true if NoInstrAttrib is present, Tag 870
func (m NoRelatedSym) HasNoInstrAttrib() bool {
	return m.Has(tag.NoInstrAttrib)
}

//HasNoLegs returns true if NoLegs is present, Tag 555
func (m NoRelatedSym) HasNoLegs() bool {
	return m.Has(tag.NoLegs)
}

//HasText returns true if Text is present, Tag 58
func (m NoRelatedSym) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m NoRelatedSym) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m NoRelatedSym) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasSecondaryPriceLimitType returns true if SecondaryPriceLimitType is present, Tag 1305
func (m NoRelatedSym) HasSecondaryPriceLimitType() bool {
	return m.Has(tag.SecondaryPriceLimitType)
}

//HasSecondaryLowLimitPrice returns true if SecondaryLowLimitPrice is present, Tag 1221
func (m NoRelatedSym) HasSecondaryLowLimitPrice() bool {
	return m.Has(tag.SecondaryLowLimitPrice)
}

//HasSecondaryHighLimitPrice returns true if SecondaryHighLimitPrice is present, Tag 1230
func (m NoRelatedSym) HasSecondaryHighLimitPrice() bool {
	return m.Has(tag.SecondaryHighLimitPrice)
}

//HasSecondaryTradingReferencePrice returns true if SecondaryTradingReferencePrice is present, Tag 1240
func (m NoRelatedSym) HasSecondaryTradingReferencePrice() bool {
	return m.Has(tag.SecondaryTradingReferencePrice)
}

//HasCorporateAction returns true if CorporateAction is present, Tag 292
func (m NoRelatedSym) HasCorporateAction() bool {
	return m.Has(tag.CorporateAction)
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

//SetEventTime sets EventTime, Tag 1145
func (m NoEvents) SetEventTime(v time.Time) {
	m.Set(field.NewEventTime(v))
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

//GetEventTime gets EventTime, Tag 1145
func (m NoEvents) GetEventTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.EventTimeField
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

//NoInstrAttrib is a repeating group element, Tag 870
type NoInstrAttrib struct {
	*quickfix.Group
}

//SetInstrAttribType sets InstrAttribType, Tag 871
func (m NoInstrAttrib) SetInstrAttribType(v enum.InstrAttribType) {
	m.Set(field.NewInstrAttribType(v))
}

//SetInstrAttribValue sets InstrAttribValue, Tag 872
func (m NoInstrAttrib) SetInstrAttribValue(v string) {
	m.Set(field.NewInstrAttribValue(v))
}

//GetInstrAttribType gets InstrAttribType, Tag 871
func (m NoInstrAttrib) GetInstrAttribType() (v enum.InstrAttribType, err quickfix.MessageRejectError) {
	var f field.InstrAttribTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrAttribValue gets InstrAttribValue, Tag 872
func (m NoInstrAttrib) GetInstrAttribValue() (v string, err quickfix.MessageRejectError) {
	var f field.InstrAttribValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasInstrAttribType returns true if InstrAttribType is present, Tag 871
func (m NoInstrAttrib) HasInstrAttribType() bool {
	return m.Has(tag.InstrAttribType)
}

//HasInstrAttribValue returns true if InstrAttribValue is present, Tag 872
func (m NoInstrAttrib) HasInstrAttribValue() bool {
	return m.Has(tag.InstrAttribValue)
}

//NoInstrAttribRepeatingGroup is a repeating group, Tag 870
type NoInstrAttribRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoInstrAttribRepeatingGroup returns an initialized, NoInstrAttribRepeatingGroup
func NewNoInstrAttribRepeatingGroup() NoInstrAttribRepeatingGroup {
	return NoInstrAttribRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrAttrib,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrAttribType), quickfix.GroupElement(tag.InstrAttribValue)})}
}

//Add create and append a new NoInstrAttrib to this group
func (m NoInstrAttribRepeatingGroup) Add() NoInstrAttrib {
	g := m.RepeatingGroup.Add()
	return NoInstrAttrib{g}
}

//Get returns the ith NoInstrAttrib in the NoInstrAttribRepeatinGroup
func (m NoInstrAttribRepeatingGroup) Get(i int) NoInstrAttrib {
	return NoInstrAttrib{m.RepeatingGroup.Get(i)}
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

//GetLegOptionRatio gets LegOptionRatio, Tag 1017
func (m NoLegs) GetLegOptionRatio() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegOptionRatioField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegPrice gets LegPrice, Tag 566
func (m NoLegs) GetLegPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegMaturityTime gets LegMaturityTime, Tag 1212
func (m NoLegs) GetLegMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegPutOrCall gets LegPutOrCall, Tag 1358
func (m NoLegs) GetLegPutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.LegPutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegExerciseStyle gets LegExerciseStyle, Tag 1420
func (m NoLegs) GetLegExerciseStyle() (v int, err quickfix.MessageRejectError) {
	var f field.LegExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegUnitOfMeasureQty gets LegUnitOfMeasureQty, Tag 1224
func (m NoLegs) GetLegUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegPriceUnitOfMeasure gets LegPriceUnitOfMeasure, Tag 1421
func (m NoLegs) GetLegPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.LegPriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegPriceUnitOfMeasureQty gets LegPriceUnitOfMeasureQty, Tag 1422
func (m NoLegs) GetLegPriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegPriceUnitOfMeasureQtyField
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSymbol), quickfix.GroupElement(tag.LegSymbolSfx), quickfix.GroupElement(tag.LegSecurityID), quickfix.GroupElement(tag.LegSecurityIDSource), NewNoLegSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.LegProduct), quickfix.GroupElement(tag.LegCFICode), quickfix.GroupElement(tag.LegSecurityType), quickfix.GroupElement(tag.LegSecuritySubType), quickfix.GroupElement(tag.LegMaturityMonthYear), quickfix.GroupElement(tag.LegMaturityDate), quickfix.GroupElement(tag.LegCouponPaymentDate), quickfix.GroupElement(tag.LegIssueDate), quickfix.GroupElement(tag.LegRepoCollateralSecurityType), quickfix.GroupElement(tag.LegRepurchaseTerm), quickfix.GroupElement(tag.LegRepurchaseRate), quickfix.GroupElement(tag.LegFactor), quickfix.GroupElement(tag.LegCreditRating), quickfix.GroupElement(tag.LegInstrRegistry), quickfix.GroupElement(tag.LegCountryOfIssue), quickfix.GroupElement(tag.LegStateOrProvinceOfIssue), quickfix.GroupElement(tag.LegLocaleOfIssue), quickfix.GroupElement(tag.LegRedemptionDate), quickfix.GroupElement(tag.LegStrikePrice), quickfix.GroupElement(tag.LegStrikeCurrency), quickfix.GroupElement(tag.LegOptAttribute), quickfix.GroupElement(tag.LegContractMultiplier), quickfix.GroupElement(tag.LegCouponRate), quickfix.GroupElement(tag.LegSecurityExchange), quickfix.GroupElement(tag.LegIssuer), quickfix.GroupElement(tag.EncodedLegIssuerLen), quickfix.GroupElement(tag.EncodedLegIssuer), quickfix.GroupElement(tag.LegSecurityDesc), quickfix.GroupElement(tag.EncodedLegSecurityDescLen), quickfix.GroupElement(tag.EncodedLegSecurityDesc), quickfix.GroupElement(tag.LegRatioQty), quickfix.GroupElement(tag.LegSide), quickfix.GroupElement(tag.LegCurrency), quickfix.GroupElement(tag.LegPool), quickfix.GroupElement(tag.LegDatedDate), quickfix.GroupElement(tag.LegContractSettlMonth), quickfix.GroupElement(tag.LegInterestAccrualDate), quickfix.GroupElement(tag.LegUnitOfMeasure), quickfix.GroupElement(tag.LegTimeUnit), quickfix.GroupElement(tag.LegOptionRatio), quickfix.GroupElement(tag.LegPrice), quickfix.GroupElement(tag.LegMaturityTime), quickfix.GroupElement(tag.LegPutOrCall), quickfix.GroupElement(tag.LegExerciseStyle), quickfix.GroupElement(tag.LegUnitOfMeasureQty), quickfix.GroupElement(tag.LegPriceUnitOfMeasure), quickfix.GroupElement(tag.LegPriceUnitOfMeasureQty)})}
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

//NoRelatedSymRepeatingGroup is a repeating group, Tag 146
type NoRelatedSymRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelatedSymRepeatingGroup returns an initialized, NoRelatedSymRepeatingGroup
func NewNoRelatedSymRepeatingGroup() NoRelatedSymRepeatingGroup {
	return NoRelatedSymRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedSym,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.SecurityIDSource), NewNoSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.Product), quickfix.GroupElement(tag.CFICode), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.SecuritySubType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDate), quickfix.GroupElement(tag.CouponPaymentDate), quickfix.GroupElement(tag.IssueDate), quickfix.GroupElement(tag.RepoCollateralSecurityType), quickfix.GroupElement(tag.RepurchaseTerm), quickfix.GroupElement(tag.RepurchaseRate), quickfix.GroupElement(tag.Factor), quickfix.GroupElement(tag.CreditRating), quickfix.GroupElement(tag.InstrRegistry), quickfix.GroupElement(tag.CountryOfIssue), quickfix.GroupElement(tag.StateOrProvinceOfIssue), quickfix.GroupElement(tag.LocaleOfIssue), quickfix.GroupElement(tag.RedemptionDate), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.StrikeCurrency), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.ContractMultiplier), quickfix.GroupElement(tag.CouponRate), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.EncodedIssuerLen), quickfix.GroupElement(tag.EncodedIssuer), quickfix.GroupElement(tag.SecurityDesc), quickfix.GroupElement(tag.EncodedSecurityDescLen), quickfix.GroupElement(tag.EncodedSecurityDesc), quickfix.GroupElement(tag.Pool), quickfix.GroupElement(tag.ContractSettlMonth), quickfix.GroupElement(tag.CPProgram), quickfix.GroupElement(tag.CPRegType), NewNoEventsRepeatingGroup(), quickfix.GroupElement(tag.DatedDate), quickfix.GroupElement(tag.InterestAccrualDate), quickfix.GroupElement(tag.SecurityStatus), quickfix.GroupElement(tag.SettleOnOpenFlag), quickfix.GroupElement(tag.InstrmtAssignmentMethod), quickfix.GroupElement(tag.StrikeMultiplier), quickfix.GroupElement(tag.StrikeValue), quickfix.GroupElement(tag.MinPriceIncrement), quickfix.GroupElement(tag.PositionLimit), quickfix.GroupElement(tag.NTPositionLimit), NewNoInstrumentPartiesRepeatingGroup(), quickfix.GroupElement(tag.UnitOfMeasure), quickfix.GroupElement(tag.TimeUnit), quickfix.GroupElement(tag.MaturityTime), quickfix.GroupElement(tag.SecurityGroup), quickfix.GroupElement(tag.MinPriceIncrementAmount), quickfix.GroupElement(tag.UnitOfMeasureQty), quickfix.GroupElement(tag.SecurityXMLLen), quickfix.GroupElement(tag.SecurityXML), quickfix.GroupElement(tag.SecurityXMLSchema), quickfix.GroupElement(tag.ProductComplex), quickfix.GroupElement(tag.PriceUnitOfMeasure), quickfix.GroupElement(tag.PriceUnitOfMeasureQty), quickfix.GroupElement(tag.SettlMethod), quickfix.GroupElement(tag.ExerciseStyle), quickfix.GroupElement(tag.OptPayAmount), quickfix.GroupElement(tag.PriceQuoteMethod), quickfix.GroupElement(tag.ListMethod), quickfix.GroupElement(tag.CapPrice), quickfix.GroupElement(tag.FloorPrice), quickfix.GroupElement(tag.PutOrCall), quickfix.GroupElement(tag.FlexibleIndicator), quickfix.GroupElement(tag.FlexProductEligibilityIndicator), quickfix.GroupElement(tag.FuturesValuationMethod), quickfix.GroupElement(tag.Currency), quickfix.GroupElement(tag.DeliveryForm), quickfix.GroupElement(tag.PctAtRisk), NewNoInstrAttribRepeatingGroup(), NewNoLegsRepeatingGroup(), quickfix.GroupElement(tag.Text), quickfix.GroupElement(tag.EncodedTextLen), quickfix.GroupElement(tag.EncodedText), quickfix.GroupElement(tag.SecondaryPriceLimitType), quickfix.GroupElement(tag.SecondaryLowLimitPrice), quickfix.GroupElement(tag.SecondaryHighLimitPrice), quickfix.GroupElement(tag.SecondaryTradingReferencePrice), quickfix.GroupElement(tag.CorporateAction)})}
}

//Add create and append a new NoRelatedSym to this group
func (m NoRelatedSymRepeatingGroup) Add() NoRelatedSym {
	g := m.RepeatingGroup.Add()
	return NoRelatedSym{g}
}

//Get returns the ith NoRelatedSym in the NoRelatedSymRepeatinGroup
func (m NoRelatedSymRepeatingGroup) Get(i int) NoRelatedSym {
	return NoRelatedSym{m.RepeatingGroup.Get(i)}
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

//NoMarketSegments is a repeating group element, Tag 1310
type NoMarketSegments struct {
	*quickfix.Group
}

//SetMarketID sets MarketID, Tag 1301
func (m NoMarketSegments) SetMarketID(v string) {
	m.Set(field.NewMarketID(v))
}

//SetMarketSegmentID sets MarketSegmentID, Tag 1300
func (m NoMarketSegments) SetMarketSegmentID(v string) {
	m.Set(field.NewMarketSegmentID(v))
}

//SetNoTickRules sets NoTickRules, Tag 1205
func (m NoMarketSegments) SetNoTickRules(f NoTickRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoLotTypeRules sets NoLotTypeRules, Tag 1234
func (m NoMarketSegments) SetNoLotTypeRules(f NoLotTypeRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetPriceLimitType sets PriceLimitType, Tag 1306
func (m NoMarketSegments) SetPriceLimitType(v enum.PriceLimitType) {
	m.Set(field.NewPriceLimitType(v))
}

//SetLowLimitPrice sets LowLimitPrice, Tag 1148
func (m NoMarketSegments) SetLowLimitPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLowLimitPrice(value, scale))
}

//SetHighLimitPrice sets HighLimitPrice, Tag 1149
func (m NoMarketSegments) SetHighLimitPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewHighLimitPrice(value, scale))
}

//SetTradingReferencePrice sets TradingReferencePrice, Tag 1150
func (m NoMarketSegments) SetTradingReferencePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewTradingReferencePrice(value, scale))
}

//SetExpirationCycle sets ExpirationCycle, Tag 827
func (m NoMarketSegments) SetExpirationCycle(v enum.ExpirationCycle) {
	m.Set(field.NewExpirationCycle(v))
}

//SetMinTradeVol sets MinTradeVol, Tag 562
func (m NoMarketSegments) SetMinTradeVol(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinTradeVol(value, scale))
}

//SetMaxTradeVol sets MaxTradeVol, Tag 1140
func (m NoMarketSegments) SetMaxTradeVol(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxTradeVol(value, scale))
}

//SetMaxPriceVariation sets MaxPriceVariation, Tag 1143
func (m NoMarketSegments) SetMaxPriceVariation(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxPriceVariation(value, scale))
}

//SetImpliedMarketIndicator sets ImpliedMarketIndicator, Tag 1144
func (m NoMarketSegments) SetImpliedMarketIndicator(v enum.ImpliedMarketIndicator) {
	m.Set(field.NewImpliedMarketIndicator(v))
}

//SetTradingCurrency sets TradingCurrency, Tag 1245
func (m NoMarketSegments) SetTradingCurrency(v string) {
	m.Set(field.NewTradingCurrency(v))
}

//SetRoundLot sets RoundLot, Tag 561
func (m NoMarketSegments) SetRoundLot(value decimal.Decimal, scale int32) {
	m.Set(field.NewRoundLot(value, scale))
}

//SetMultilegModel sets MultilegModel, Tag 1377
func (m NoMarketSegments) SetMultilegModel(v enum.MultilegModel) {
	m.Set(field.NewMultilegModel(v))
}

//SetMultilegPriceMethod sets MultilegPriceMethod, Tag 1378
func (m NoMarketSegments) SetMultilegPriceMethod(v enum.MultilegPriceMethod) {
	m.Set(field.NewMultilegPriceMethod(v))
}

//SetPriceType sets PriceType, Tag 423
func (m NoMarketSegments) SetPriceType(v enum.PriceType) {
	m.Set(field.NewPriceType(v))
}

//SetNoTradingSessionRules sets NoTradingSessionRules, Tag 1309
func (m NoMarketSegments) SetNoTradingSessionRules(f NoTradingSessionRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoNestedInstrAttrib sets NoNestedInstrAttrib, Tag 1312
func (m NoMarketSegments) SetNoNestedInstrAttrib(f NoNestedInstrAttribRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoStrikeRules sets NoStrikeRules, Tag 1201
func (m NoMarketSegments) SetNoStrikeRules(f NoStrikeRulesRepeatingGroup) {
	m.SetGroup(f)
}

//GetMarketID gets MarketID, Tag 1301
func (m NoMarketSegments) GetMarketID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMarketSegmentID gets MarketSegmentID, Tag 1300
func (m NoMarketSegments) GetMarketSegmentID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketSegmentIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoTickRules gets NoTickRules, Tag 1205
func (m NoMarketSegments) GetNoTickRules() (NoTickRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTickRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoLotTypeRules gets NoLotTypeRules, Tag 1234
func (m NoMarketSegments) GetNoLotTypeRules() (NoLotTypeRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLotTypeRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetPriceLimitType gets PriceLimitType, Tag 1306
func (m NoMarketSegments) GetPriceLimitType() (v enum.PriceLimitType, err quickfix.MessageRejectError) {
	var f field.PriceLimitTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLowLimitPrice gets LowLimitPrice, Tag 1148
func (m NoMarketSegments) GetLowLimitPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LowLimitPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetHighLimitPrice gets HighLimitPrice, Tag 1149
func (m NoMarketSegments) GetHighLimitPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.HighLimitPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingReferencePrice gets TradingReferencePrice, Tag 1150
func (m NoMarketSegments) GetTradingReferencePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TradingReferencePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExpirationCycle gets ExpirationCycle, Tag 827
func (m NoMarketSegments) GetExpirationCycle() (v enum.ExpirationCycle, err quickfix.MessageRejectError) {
	var f field.ExpirationCycleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMinTradeVol gets MinTradeVol, Tag 562
func (m NoMarketSegments) GetMinTradeVol() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinTradeVolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaxTradeVol gets MaxTradeVol, Tag 1140
func (m NoMarketSegments) GetMaxTradeVol() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MaxTradeVolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaxPriceVariation gets MaxPriceVariation, Tag 1143
func (m NoMarketSegments) GetMaxPriceVariation() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MaxPriceVariationField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetImpliedMarketIndicator gets ImpliedMarketIndicator, Tag 1144
func (m NoMarketSegments) GetImpliedMarketIndicator() (v enum.ImpliedMarketIndicator, err quickfix.MessageRejectError) {
	var f field.ImpliedMarketIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingCurrency gets TradingCurrency, Tag 1245
func (m NoMarketSegments) GetTradingCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.TradingCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRoundLot gets RoundLot, Tag 561
func (m NoMarketSegments) GetRoundLot() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RoundLotField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMultilegModel gets MultilegModel, Tag 1377
func (m NoMarketSegments) GetMultilegModel() (v enum.MultilegModel, err quickfix.MessageRejectError) {
	var f field.MultilegModelField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMultilegPriceMethod gets MultilegPriceMethod, Tag 1378
func (m NoMarketSegments) GetMultilegPriceMethod() (v enum.MultilegPriceMethod, err quickfix.MessageRejectError) {
	var f field.MultilegPriceMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPriceType gets PriceType, Tag 423
func (m NoMarketSegments) GetPriceType() (v enum.PriceType, err quickfix.MessageRejectError) {
	var f field.PriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoTradingSessionRules gets NoTradingSessionRules, Tag 1309
func (m NoMarketSegments) GetNoTradingSessionRules() (NoTradingSessionRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTradingSessionRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoNestedInstrAttrib gets NoNestedInstrAttrib, Tag 1312
func (m NoMarketSegments) GetNoNestedInstrAttrib() (NoNestedInstrAttribRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedInstrAttribRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoStrikeRules gets NoStrikeRules, Tag 1201
func (m NoMarketSegments) GetNoStrikeRules() (NoStrikeRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStrikeRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasMarketID returns true if MarketID is present, Tag 1301
func (m NoMarketSegments) HasMarketID() bool {
	return m.Has(tag.MarketID)
}

//HasMarketSegmentID returns true if MarketSegmentID is present, Tag 1300
func (m NoMarketSegments) HasMarketSegmentID() bool {
	return m.Has(tag.MarketSegmentID)
}

//HasNoTickRules returns true if NoTickRules is present, Tag 1205
func (m NoMarketSegments) HasNoTickRules() bool {
	return m.Has(tag.NoTickRules)
}

//HasNoLotTypeRules returns true if NoLotTypeRules is present, Tag 1234
func (m NoMarketSegments) HasNoLotTypeRules() bool {
	return m.Has(tag.NoLotTypeRules)
}

//HasPriceLimitType returns true if PriceLimitType is present, Tag 1306
func (m NoMarketSegments) HasPriceLimitType() bool {
	return m.Has(tag.PriceLimitType)
}

//HasLowLimitPrice returns true if LowLimitPrice is present, Tag 1148
func (m NoMarketSegments) HasLowLimitPrice() bool {
	return m.Has(tag.LowLimitPrice)
}

//HasHighLimitPrice returns true if HighLimitPrice is present, Tag 1149
func (m NoMarketSegments) HasHighLimitPrice() bool {
	return m.Has(tag.HighLimitPrice)
}

//HasTradingReferencePrice returns true if TradingReferencePrice is present, Tag 1150
func (m NoMarketSegments) HasTradingReferencePrice() bool {
	return m.Has(tag.TradingReferencePrice)
}

//HasExpirationCycle returns true if ExpirationCycle is present, Tag 827
func (m NoMarketSegments) HasExpirationCycle() bool {
	return m.Has(tag.ExpirationCycle)
}

//HasMinTradeVol returns true if MinTradeVol is present, Tag 562
func (m NoMarketSegments) HasMinTradeVol() bool {
	return m.Has(tag.MinTradeVol)
}

//HasMaxTradeVol returns true if MaxTradeVol is present, Tag 1140
func (m NoMarketSegments) HasMaxTradeVol() bool {
	return m.Has(tag.MaxTradeVol)
}

//HasMaxPriceVariation returns true if MaxPriceVariation is present, Tag 1143
func (m NoMarketSegments) HasMaxPriceVariation() bool {
	return m.Has(tag.MaxPriceVariation)
}

//HasImpliedMarketIndicator returns true if ImpliedMarketIndicator is present, Tag 1144
func (m NoMarketSegments) HasImpliedMarketIndicator() bool {
	return m.Has(tag.ImpliedMarketIndicator)
}

//HasTradingCurrency returns true if TradingCurrency is present, Tag 1245
func (m NoMarketSegments) HasTradingCurrency() bool {
	return m.Has(tag.TradingCurrency)
}

//HasRoundLot returns true if RoundLot is present, Tag 561
func (m NoMarketSegments) HasRoundLot() bool {
	return m.Has(tag.RoundLot)
}

//HasMultilegModel returns true if MultilegModel is present, Tag 1377
func (m NoMarketSegments) HasMultilegModel() bool {
	return m.Has(tag.MultilegModel)
}

//HasMultilegPriceMethod returns true if MultilegPriceMethod is present, Tag 1378
func (m NoMarketSegments) HasMultilegPriceMethod() bool {
	return m.Has(tag.MultilegPriceMethod)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m NoMarketSegments) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasNoTradingSessionRules returns true if NoTradingSessionRules is present, Tag 1309
func (m NoMarketSegments) HasNoTradingSessionRules() bool {
	return m.Has(tag.NoTradingSessionRules)
}

//HasNoNestedInstrAttrib returns true if NoNestedInstrAttrib is present, Tag 1312
func (m NoMarketSegments) HasNoNestedInstrAttrib() bool {
	return m.Has(tag.NoNestedInstrAttrib)
}

//HasNoStrikeRules returns true if NoStrikeRules is present, Tag 1201
func (m NoMarketSegments) HasNoStrikeRules() bool {
	return m.Has(tag.NoStrikeRules)
}

//NoTickRules is a repeating group element, Tag 1205
type NoTickRules struct {
	*quickfix.Group
}

//SetStartTickPriceRange sets StartTickPriceRange, Tag 1206
func (m NoTickRules) SetStartTickPriceRange(value decimal.Decimal, scale int32) {
	m.Set(field.NewStartTickPriceRange(value, scale))
}

//SetEndTickPriceRange sets EndTickPriceRange, Tag 1207
func (m NoTickRules) SetEndTickPriceRange(value decimal.Decimal, scale int32) {
	m.Set(field.NewEndTickPriceRange(value, scale))
}

//SetTickIncrement sets TickIncrement, Tag 1208
func (m NoTickRules) SetTickIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewTickIncrement(value, scale))
}

//SetTickRuleType sets TickRuleType, Tag 1209
func (m NoTickRules) SetTickRuleType(v enum.TickRuleType) {
	m.Set(field.NewTickRuleType(v))
}

//GetStartTickPriceRange gets StartTickPriceRange, Tag 1206
func (m NoTickRules) GetStartTickPriceRange() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StartTickPriceRangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEndTickPriceRange gets EndTickPriceRange, Tag 1207
func (m NoTickRules) GetEndTickPriceRange() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.EndTickPriceRangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTickIncrement gets TickIncrement, Tag 1208
func (m NoTickRules) GetTickIncrement() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TickIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTickRuleType gets TickRuleType, Tag 1209
func (m NoTickRules) GetTickRuleType() (v enum.TickRuleType, err quickfix.MessageRejectError) {
	var f field.TickRuleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasStartTickPriceRange returns true if StartTickPriceRange is present, Tag 1206
func (m NoTickRules) HasStartTickPriceRange() bool {
	return m.Has(tag.StartTickPriceRange)
}

//HasEndTickPriceRange returns true if EndTickPriceRange is present, Tag 1207
func (m NoTickRules) HasEndTickPriceRange() bool {
	return m.Has(tag.EndTickPriceRange)
}

//HasTickIncrement returns true if TickIncrement is present, Tag 1208
func (m NoTickRules) HasTickIncrement() bool {
	return m.Has(tag.TickIncrement)
}

//HasTickRuleType returns true if TickRuleType is present, Tag 1209
func (m NoTickRules) HasTickRuleType() bool {
	return m.Has(tag.TickRuleType)
}

//NoTickRulesRepeatingGroup is a repeating group, Tag 1205
type NoTickRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTickRulesRepeatingGroup returns an initialized, NoTickRulesRepeatingGroup
func NewNoTickRulesRepeatingGroup() NoTickRulesRepeatingGroup {
	return NoTickRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTickRules,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.StartTickPriceRange), quickfix.GroupElement(tag.EndTickPriceRange), quickfix.GroupElement(tag.TickIncrement), quickfix.GroupElement(tag.TickRuleType)})}
}

//Add create and append a new NoTickRules to this group
func (m NoTickRulesRepeatingGroup) Add() NoTickRules {
	g := m.RepeatingGroup.Add()
	return NoTickRules{g}
}

//Get returns the ith NoTickRules in the NoTickRulesRepeatinGroup
func (m NoTickRulesRepeatingGroup) Get(i int) NoTickRules {
	return NoTickRules{m.RepeatingGroup.Get(i)}
}

//NoLotTypeRules is a repeating group element, Tag 1234
type NoLotTypeRules struct {
	*quickfix.Group
}

//SetLotType sets LotType, Tag 1093
func (m NoLotTypeRules) SetLotType(v enum.LotType) {
	m.Set(field.NewLotType(v))
}

//SetMinLotSize sets MinLotSize, Tag 1231
func (m NoLotTypeRules) SetMinLotSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinLotSize(value, scale))
}

//GetLotType gets LotType, Tag 1093
func (m NoLotTypeRules) GetLotType() (v enum.LotType, err quickfix.MessageRejectError) {
	var f field.LotTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMinLotSize gets MinLotSize, Tag 1231
func (m NoLotTypeRules) GetMinLotSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinLotSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasLotType returns true if LotType is present, Tag 1093
func (m NoLotTypeRules) HasLotType() bool {
	return m.Has(tag.LotType)
}

//HasMinLotSize returns true if MinLotSize is present, Tag 1231
func (m NoLotTypeRules) HasMinLotSize() bool {
	return m.Has(tag.MinLotSize)
}

//NoLotTypeRulesRepeatingGroup is a repeating group, Tag 1234
type NoLotTypeRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLotTypeRulesRepeatingGroup returns an initialized, NoLotTypeRulesRepeatingGroup
func NewNoLotTypeRulesRepeatingGroup() NoLotTypeRulesRepeatingGroup {
	return NoLotTypeRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLotTypeRules,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LotType), quickfix.GroupElement(tag.MinLotSize)})}
}

//Add create and append a new NoLotTypeRules to this group
func (m NoLotTypeRulesRepeatingGroup) Add() NoLotTypeRules {
	g := m.RepeatingGroup.Add()
	return NoLotTypeRules{g}
}

//Get returns the ith NoLotTypeRules in the NoLotTypeRulesRepeatinGroup
func (m NoLotTypeRulesRepeatingGroup) Get(i int) NoLotTypeRules {
	return NoLotTypeRules{m.RepeatingGroup.Get(i)}
}

//NoTradingSessionRules is a repeating group element, Tag 1309
type NoTradingSessionRules struct {
	*quickfix.Group
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoTradingSessionRules) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m NoTradingSessionRules) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetNoOrdTypeRules sets NoOrdTypeRules, Tag 1237
func (m NoTradingSessionRules) SetNoOrdTypeRules(f NoOrdTypeRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoTimeInForceRules sets NoTimeInForceRules, Tag 1239
func (m NoTradingSessionRules) SetNoTimeInForceRules(f NoTimeInForceRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoExecInstRules sets NoExecInstRules, Tag 1232
func (m NoTradingSessionRules) SetNoExecInstRules(f NoExecInstRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoMatchRules sets NoMatchRules, Tag 1235
func (m NoTradingSessionRules) SetNoMatchRules(f NoMatchRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoMDFeedTypes sets NoMDFeedTypes, Tag 1141
func (m NoTradingSessionRules) SetNoMDFeedTypes(f NoMDFeedTypesRepeatingGroup) {
	m.SetGroup(f)
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoTradingSessionRules) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m NoTradingSessionRules) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoOrdTypeRules gets NoOrdTypeRules, Tag 1237
func (m NoTradingSessionRules) GetNoOrdTypeRules() (NoOrdTypeRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoOrdTypeRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoTimeInForceRules gets NoTimeInForceRules, Tag 1239
func (m NoTradingSessionRules) GetNoTimeInForceRules() (NoTimeInForceRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTimeInForceRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoExecInstRules gets NoExecInstRules, Tag 1232
func (m NoTradingSessionRules) GetNoExecInstRules() (NoExecInstRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoExecInstRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoMatchRules gets NoMatchRules, Tag 1235
func (m NoTradingSessionRules) GetNoMatchRules() (NoMatchRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMatchRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoMDFeedTypes gets NoMDFeedTypes, Tag 1141
func (m NoTradingSessionRules) GetNoMDFeedTypes() (NoMDFeedTypesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMDFeedTypesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m NoTradingSessionRules) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m NoTradingSessionRules) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//HasNoOrdTypeRules returns true if NoOrdTypeRules is present, Tag 1237
func (m NoTradingSessionRules) HasNoOrdTypeRules() bool {
	return m.Has(tag.NoOrdTypeRules)
}

//HasNoTimeInForceRules returns true if NoTimeInForceRules is present, Tag 1239
func (m NoTradingSessionRules) HasNoTimeInForceRules() bool {
	return m.Has(tag.NoTimeInForceRules)
}

//HasNoExecInstRules returns true if NoExecInstRules is present, Tag 1232
func (m NoTradingSessionRules) HasNoExecInstRules() bool {
	return m.Has(tag.NoExecInstRules)
}

//HasNoMatchRules returns true if NoMatchRules is present, Tag 1235
func (m NoTradingSessionRules) HasNoMatchRules() bool {
	return m.Has(tag.NoMatchRules)
}

//HasNoMDFeedTypes returns true if NoMDFeedTypes is present, Tag 1141
func (m NoTradingSessionRules) HasNoMDFeedTypes() bool {
	return m.Has(tag.NoMDFeedTypes)
}

//NoOrdTypeRules is a repeating group element, Tag 1237
type NoOrdTypeRules struct {
	*quickfix.Group
}

//SetOrdType sets OrdType, Tag 40
func (m NoOrdTypeRules) SetOrdType(v enum.OrdType) {
	m.Set(field.NewOrdType(v))
}

//GetOrdType gets OrdType, Tag 40
func (m NoOrdTypeRules) GetOrdType() (v enum.OrdType, err quickfix.MessageRejectError) {
	var f field.OrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m NoOrdTypeRules) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//NoOrdTypeRulesRepeatingGroup is a repeating group, Tag 1237
type NoOrdTypeRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoOrdTypeRulesRepeatingGroup returns an initialized, NoOrdTypeRulesRepeatingGroup
func NewNoOrdTypeRulesRepeatingGroup() NoOrdTypeRulesRepeatingGroup {
	return NoOrdTypeRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoOrdTypeRules,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.OrdType)})}
}

//Add create and append a new NoOrdTypeRules to this group
func (m NoOrdTypeRulesRepeatingGroup) Add() NoOrdTypeRules {
	g := m.RepeatingGroup.Add()
	return NoOrdTypeRules{g}
}

//Get returns the ith NoOrdTypeRules in the NoOrdTypeRulesRepeatinGroup
func (m NoOrdTypeRulesRepeatingGroup) Get(i int) NoOrdTypeRules {
	return NoOrdTypeRules{m.RepeatingGroup.Get(i)}
}

//NoTimeInForceRules is a repeating group element, Tag 1239
type NoTimeInForceRules struct {
	*quickfix.Group
}

//SetTimeInForce sets TimeInForce, Tag 59
func (m NoTimeInForceRules) SetTimeInForce(v enum.TimeInForce) {
	m.Set(field.NewTimeInForce(v))
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m NoTimeInForceRules) GetTimeInForce() (v enum.TimeInForce, err quickfix.MessageRejectError) {
	var f field.TimeInForceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasTimeInForce returns true if TimeInForce is present, Tag 59
func (m NoTimeInForceRules) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

//NoTimeInForceRulesRepeatingGroup is a repeating group, Tag 1239
type NoTimeInForceRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTimeInForceRulesRepeatingGroup returns an initialized, NoTimeInForceRulesRepeatingGroup
func NewNoTimeInForceRulesRepeatingGroup() NoTimeInForceRulesRepeatingGroup {
	return NoTimeInForceRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTimeInForceRules,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TimeInForce)})}
}

//Add create and append a new NoTimeInForceRules to this group
func (m NoTimeInForceRulesRepeatingGroup) Add() NoTimeInForceRules {
	g := m.RepeatingGroup.Add()
	return NoTimeInForceRules{g}
}

//Get returns the ith NoTimeInForceRules in the NoTimeInForceRulesRepeatinGroup
func (m NoTimeInForceRulesRepeatingGroup) Get(i int) NoTimeInForceRules {
	return NoTimeInForceRules{m.RepeatingGroup.Get(i)}
}

//NoExecInstRules is a repeating group element, Tag 1232
type NoExecInstRules struct {
	*quickfix.Group
}

//SetExecInstValue sets ExecInstValue, Tag 1308
func (m NoExecInstRules) SetExecInstValue(v string) {
	m.Set(field.NewExecInstValue(v))
}

//GetExecInstValue gets ExecInstValue, Tag 1308
func (m NoExecInstRules) GetExecInstValue() (v string, err quickfix.MessageRejectError) {
	var f field.ExecInstValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasExecInstValue returns true if ExecInstValue is present, Tag 1308
func (m NoExecInstRules) HasExecInstValue() bool {
	return m.Has(tag.ExecInstValue)
}

//NoExecInstRulesRepeatingGroup is a repeating group, Tag 1232
type NoExecInstRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoExecInstRulesRepeatingGroup returns an initialized, NoExecInstRulesRepeatingGroup
func NewNoExecInstRulesRepeatingGroup() NoExecInstRulesRepeatingGroup {
	return NoExecInstRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoExecInstRules,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ExecInstValue)})}
}

//Add create and append a new NoExecInstRules to this group
func (m NoExecInstRulesRepeatingGroup) Add() NoExecInstRules {
	g := m.RepeatingGroup.Add()
	return NoExecInstRules{g}
}

//Get returns the ith NoExecInstRules in the NoExecInstRulesRepeatinGroup
func (m NoExecInstRulesRepeatingGroup) Get(i int) NoExecInstRules {
	return NoExecInstRules{m.RepeatingGroup.Get(i)}
}

//NoMatchRules is a repeating group element, Tag 1235
type NoMatchRules struct {
	*quickfix.Group
}

//SetMatchAlgorithm sets MatchAlgorithm, Tag 1142
func (m NoMatchRules) SetMatchAlgorithm(v string) {
	m.Set(field.NewMatchAlgorithm(v))
}

//SetMatchType sets MatchType, Tag 574
func (m NoMatchRules) SetMatchType(v enum.MatchType) {
	m.Set(field.NewMatchType(v))
}

//GetMatchAlgorithm gets MatchAlgorithm, Tag 1142
func (m NoMatchRules) GetMatchAlgorithm() (v string, err quickfix.MessageRejectError) {
	var f field.MatchAlgorithmField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMatchType gets MatchType, Tag 574
func (m NoMatchRules) GetMatchType() (v enum.MatchType, err quickfix.MessageRejectError) {
	var f field.MatchTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasMatchAlgorithm returns true if MatchAlgorithm is present, Tag 1142
func (m NoMatchRules) HasMatchAlgorithm() bool {
	return m.Has(tag.MatchAlgorithm)
}

//HasMatchType returns true if MatchType is present, Tag 574
func (m NoMatchRules) HasMatchType() bool {
	return m.Has(tag.MatchType)
}

//NoMatchRulesRepeatingGroup is a repeating group, Tag 1235
type NoMatchRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMatchRulesRepeatingGroup returns an initialized, NoMatchRulesRepeatingGroup
func NewNoMatchRulesRepeatingGroup() NoMatchRulesRepeatingGroup {
	return NoMatchRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMatchRules,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MatchAlgorithm), quickfix.GroupElement(tag.MatchType)})}
}

//Add create and append a new NoMatchRules to this group
func (m NoMatchRulesRepeatingGroup) Add() NoMatchRules {
	g := m.RepeatingGroup.Add()
	return NoMatchRules{g}
}

//Get returns the ith NoMatchRules in the NoMatchRulesRepeatinGroup
func (m NoMatchRulesRepeatingGroup) Get(i int) NoMatchRules {
	return NoMatchRules{m.RepeatingGroup.Get(i)}
}

//NoMDFeedTypes is a repeating group element, Tag 1141
type NoMDFeedTypes struct {
	*quickfix.Group
}

//SetMDFeedType sets MDFeedType, Tag 1022
func (m NoMDFeedTypes) SetMDFeedType(v string) {
	m.Set(field.NewMDFeedType(v))
}

//SetMarketDepth sets MarketDepth, Tag 264
func (m NoMDFeedTypes) SetMarketDepth(v int) {
	m.Set(field.NewMarketDepth(v))
}

//SetMDBookType sets MDBookType, Tag 1021
func (m NoMDFeedTypes) SetMDBookType(v enum.MDBookType) {
	m.Set(field.NewMDBookType(v))
}

//GetMDFeedType gets MDFeedType, Tag 1022
func (m NoMDFeedTypes) GetMDFeedType() (v string, err quickfix.MessageRejectError) {
	var f field.MDFeedTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMarketDepth gets MarketDepth, Tag 264
func (m NoMDFeedTypes) GetMarketDepth() (v int, err quickfix.MessageRejectError) {
	var f field.MarketDepthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDBookType gets MDBookType, Tag 1021
func (m NoMDFeedTypes) GetMDBookType() (v enum.MDBookType, err quickfix.MessageRejectError) {
	var f field.MDBookTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasMDFeedType returns true if MDFeedType is present, Tag 1022
func (m NoMDFeedTypes) HasMDFeedType() bool {
	return m.Has(tag.MDFeedType)
}

//HasMarketDepth returns true if MarketDepth is present, Tag 264
func (m NoMDFeedTypes) HasMarketDepth() bool {
	return m.Has(tag.MarketDepth)
}

//HasMDBookType returns true if MDBookType is present, Tag 1021
func (m NoMDFeedTypes) HasMDBookType() bool {
	return m.Has(tag.MDBookType)
}

//NoMDFeedTypesRepeatingGroup is a repeating group, Tag 1141
type NoMDFeedTypesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMDFeedTypesRepeatingGroup returns an initialized, NoMDFeedTypesRepeatingGroup
func NewNoMDFeedTypesRepeatingGroup() NoMDFeedTypesRepeatingGroup {
	return NoMDFeedTypesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMDFeedTypes,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MDFeedType), quickfix.GroupElement(tag.MarketDepth), quickfix.GroupElement(tag.MDBookType)})}
}

//Add create and append a new NoMDFeedTypes to this group
func (m NoMDFeedTypesRepeatingGroup) Add() NoMDFeedTypes {
	g := m.RepeatingGroup.Add()
	return NoMDFeedTypes{g}
}

//Get returns the ith NoMDFeedTypes in the NoMDFeedTypesRepeatinGroup
func (m NoMDFeedTypesRepeatingGroup) Get(i int) NoMDFeedTypes {
	return NoMDFeedTypes{m.RepeatingGroup.Get(i)}
}

//NoTradingSessionRulesRepeatingGroup is a repeating group, Tag 1309
type NoTradingSessionRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTradingSessionRulesRepeatingGroup returns an initialized, NoTradingSessionRulesRepeatingGroup
func NewNoTradingSessionRulesRepeatingGroup() NoTradingSessionRulesRepeatingGroup {
	return NoTradingSessionRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTradingSessionRules,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.TradingSessionSubID), NewNoOrdTypeRulesRepeatingGroup(), NewNoTimeInForceRulesRepeatingGroup(), NewNoExecInstRulesRepeatingGroup(), NewNoMatchRulesRepeatingGroup(), NewNoMDFeedTypesRepeatingGroup()})}
}

//Add create and append a new NoTradingSessionRules to this group
func (m NoTradingSessionRulesRepeatingGroup) Add() NoTradingSessionRules {
	g := m.RepeatingGroup.Add()
	return NoTradingSessionRules{g}
}

//Get returns the ith NoTradingSessionRules in the NoTradingSessionRulesRepeatinGroup
func (m NoTradingSessionRulesRepeatingGroup) Get(i int) NoTradingSessionRules {
	return NoTradingSessionRules{m.RepeatingGroup.Get(i)}
}

//NoNestedInstrAttrib is a repeating group element, Tag 1312
type NoNestedInstrAttrib struct {
	*quickfix.Group
}

//SetNestedInstrAttribType sets NestedInstrAttribType, Tag 1210
func (m NoNestedInstrAttrib) SetNestedInstrAttribType(v int) {
	m.Set(field.NewNestedInstrAttribType(v))
}

//SetNestedInstrAttribValue sets NestedInstrAttribValue, Tag 1211
func (m NoNestedInstrAttrib) SetNestedInstrAttribValue(v string) {
	m.Set(field.NewNestedInstrAttribValue(v))
}

//GetNestedInstrAttribType gets NestedInstrAttribType, Tag 1210
func (m NoNestedInstrAttrib) GetNestedInstrAttribType() (v int, err quickfix.MessageRejectError) {
	var f field.NestedInstrAttribTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNestedInstrAttribValue gets NestedInstrAttribValue, Tag 1211
func (m NoNestedInstrAttrib) GetNestedInstrAttribValue() (v string, err quickfix.MessageRejectError) {
	var f field.NestedInstrAttribValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasNestedInstrAttribType returns true if NestedInstrAttribType is present, Tag 1210
func (m NoNestedInstrAttrib) HasNestedInstrAttribType() bool {
	return m.Has(tag.NestedInstrAttribType)
}

//HasNestedInstrAttribValue returns true if NestedInstrAttribValue is present, Tag 1211
func (m NoNestedInstrAttrib) HasNestedInstrAttribValue() bool {
	return m.Has(tag.NestedInstrAttribValue)
}

//NoNestedInstrAttribRepeatingGroup is a repeating group, Tag 1312
type NoNestedInstrAttribRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNestedInstrAttribRepeatingGroup returns an initialized, NoNestedInstrAttribRepeatingGroup
func NewNoNestedInstrAttribRepeatingGroup() NoNestedInstrAttribRepeatingGroup {
	return NoNestedInstrAttribRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNestedInstrAttrib,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedInstrAttribType), quickfix.GroupElement(tag.NestedInstrAttribValue)})}
}

//Add create and append a new NoNestedInstrAttrib to this group
func (m NoNestedInstrAttribRepeatingGroup) Add() NoNestedInstrAttrib {
	g := m.RepeatingGroup.Add()
	return NoNestedInstrAttrib{g}
}

//Get returns the ith NoNestedInstrAttrib in the NoNestedInstrAttribRepeatinGroup
func (m NoNestedInstrAttribRepeatingGroup) Get(i int) NoNestedInstrAttrib {
	return NoNestedInstrAttrib{m.RepeatingGroup.Get(i)}
}

//NoStrikeRules is a repeating group element, Tag 1201
type NoStrikeRules struct {
	*quickfix.Group
}

//SetStrikeRuleID sets StrikeRuleID, Tag 1223
func (m NoStrikeRules) SetStrikeRuleID(v string) {
	m.Set(field.NewStrikeRuleID(v))
}

//SetStartStrikePxRange sets StartStrikePxRange, Tag 1202
func (m NoStrikeRules) SetStartStrikePxRange(value decimal.Decimal, scale int32) {
	m.Set(field.NewStartStrikePxRange(value, scale))
}

//SetEndStrikePxRange sets EndStrikePxRange, Tag 1203
func (m NoStrikeRules) SetEndStrikePxRange(value decimal.Decimal, scale int32) {
	m.Set(field.NewEndStrikePxRange(value, scale))
}

//SetStrikeIncrement sets StrikeIncrement, Tag 1204
func (m NoStrikeRules) SetStrikeIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeIncrement(value, scale))
}

//SetStrikeExerciseStyle sets StrikeExerciseStyle, Tag 1304
func (m NoStrikeRules) SetStrikeExerciseStyle(v int) {
	m.Set(field.NewStrikeExerciseStyle(v))
}

//SetNoMaturityRules sets NoMaturityRules, Tag 1236
func (m NoStrikeRules) SetNoMaturityRules(f NoMaturityRulesRepeatingGroup) {
	m.SetGroup(f)
}

//GetStrikeRuleID gets StrikeRuleID, Tag 1223
func (m NoStrikeRules) GetStrikeRuleID() (v string, err quickfix.MessageRejectError) {
	var f field.StrikeRuleIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStartStrikePxRange gets StartStrikePxRange, Tag 1202
func (m NoStrikeRules) GetStartStrikePxRange() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StartStrikePxRangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEndStrikePxRange gets EndStrikePxRange, Tag 1203
func (m NoStrikeRules) GetEndStrikePxRange() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.EndStrikePxRangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikeIncrement gets StrikeIncrement, Tag 1204
func (m NoStrikeRules) GetStrikeIncrement() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikeIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikeExerciseStyle gets StrikeExerciseStyle, Tag 1304
func (m NoStrikeRules) GetStrikeExerciseStyle() (v int, err quickfix.MessageRejectError) {
	var f field.StrikeExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoMaturityRules gets NoMaturityRules, Tag 1236
func (m NoStrikeRules) GetNoMaturityRules() (NoMaturityRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMaturityRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasStrikeRuleID returns true if StrikeRuleID is present, Tag 1223
func (m NoStrikeRules) HasStrikeRuleID() bool {
	return m.Has(tag.StrikeRuleID)
}

//HasStartStrikePxRange returns true if StartStrikePxRange is present, Tag 1202
func (m NoStrikeRules) HasStartStrikePxRange() bool {
	return m.Has(tag.StartStrikePxRange)
}

//HasEndStrikePxRange returns true if EndStrikePxRange is present, Tag 1203
func (m NoStrikeRules) HasEndStrikePxRange() bool {
	return m.Has(tag.EndStrikePxRange)
}

//HasStrikeIncrement returns true if StrikeIncrement is present, Tag 1204
func (m NoStrikeRules) HasStrikeIncrement() bool {
	return m.Has(tag.StrikeIncrement)
}

//HasStrikeExerciseStyle returns true if StrikeExerciseStyle is present, Tag 1304
func (m NoStrikeRules) HasStrikeExerciseStyle() bool {
	return m.Has(tag.StrikeExerciseStyle)
}

//HasNoMaturityRules returns true if NoMaturityRules is present, Tag 1236
func (m NoStrikeRules) HasNoMaturityRules() bool {
	return m.Has(tag.NoMaturityRules)
}

//NoMaturityRules is a repeating group element, Tag 1236
type NoMaturityRules struct {
	*quickfix.Group
}

//SetMaturityRuleID sets MaturityRuleID, Tag 1222
func (m NoMaturityRules) SetMaturityRuleID(v string) {
	m.Set(field.NewMaturityRuleID(v))
}

//SetMaturityMonthYearFormat sets MaturityMonthYearFormat, Tag 1303
func (m NoMaturityRules) SetMaturityMonthYearFormat(v enum.MaturityMonthYearFormat) {
	m.Set(field.NewMaturityMonthYearFormat(v))
}

//SetMaturityMonthYearIncrementUnits sets MaturityMonthYearIncrementUnits, Tag 1302
func (m NoMaturityRules) SetMaturityMonthYearIncrementUnits(v enum.MaturityMonthYearIncrementUnits) {
	m.Set(field.NewMaturityMonthYearIncrementUnits(v))
}

//SetStartMaturityMonthYear sets StartMaturityMonthYear, Tag 1241
func (m NoMaturityRules) SetStartMaturityMonthYear(v string) {
	m.Set(field.NewStartMaturityMonthYear(v))
}

//SetEndMaturityMonthYear sets EndMaturityMonthYear, Tag 1226
func (m NoMaturityRules) SetEndMaturityMonthYear(v string) {
	m.Set(field.NewEndMaturityMonthYear(v))
}

//SetMaturityMonthYearIncrement sets MaturityMonthYearIncrement, Tag 1229
func (m NoMaturityRules) SetMaturityMonthYearIncrement(v int) {
	m.Set(field.NewMaturityMonthYearIncrement(v))
}

//GetMaturityRuleID gets MaturityRuleID, Tag 1222
func (m NoMaturityRules) GetMaturityRuleID() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityRuleIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityMonthYearFormat gets MaturityMonthYearFormat, Tag 1303
func (m NoMaturityRules) GetMaturityMonthYearFormat() (v enum.MaturityMonthYearFormat, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearFormatField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityMonthYearIncrementUnits gets MaturityMonthYearIncrementUnits, Tag 1302
func (m NoMaturityRules) GetMaturityMonthYearIncrementUnits() (v enum.MaturityMonthYearIncrementUnits, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearIncrementUnitsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStartMaturityMonthYear gets StartMaturityMonthYear, Tag 1241
func (m NoMaturityRules) GetStartMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.StartMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEndMaturityMonthYear gets EndMaturityMonthYear, Tag 1226
func (m NoMaturityRules) GetEndMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.EndMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityMonthYearIncrement gets MaturityMonthYearIncrement, Tag 1229
func (m NoMaturityRules) GetMaturityMonthYearIncrement() (v int, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasMaturityRuleID returns true if MaturityRuleID is present, Tag 1222
func (m NoMaturityRules) HasMaturityRuleID() bool {
	return m.Has(tag.MaturityRuleID)
}

//HasMaturityMonthYearFormat returns true if MaturityMonthYearFormat is present, Tag 1303
func (m NoMaturityRules) HasMaturityMonthYearFormat() bool {
	return m.Has(tag.MaturityMonthYearFormat)
}

//HasMaturityMonthYearIncrementUnits returns true if MaturityMonthYearIncrementUnits is present, Tag 1302
func (m NoMaturityRules) HasMaturityMonthYearIncrementUnits() bool {
	return m.Has(tag.MaturityMonthYearIncrementUnits)
}

//HasStartMaturityMonthYear returns true if StartMaturityMonthYear is present, Tag 1241
func (m NoMaturityRules) HasStartMaturityMonthYear() bool {
	return m.Has(tag.StartMaturityMonthYear)
}

//HasEndMaturityMonthYear returns true if EndMaturityMonthYear is present, Tag 1226
func (m NoMaturityRules) HasEndMaturityMonthYear() bool {
	return m.Has(tag.EndMaturityMonthYear)
}

//HasMaturityMonthYearIncrement returns true if MaturityMonthYearIncrement is present, Tag 1229
func (m NoMaturityRules) HasMaturityMonthYearIncrement() bool {
	return m.Has(tag.MaturityMonthYearIncrement)
}

//NoMaturityRulesRepeatingGroup is a repeating group, Tag 1236
type NoMaturityRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMaturityRulesRepeatingGroup returns an initialized, NoMaturityRulesRepeatingGroup
func NewNoMaturityRulesRepeatingGroup() NoMaturityRulesRepeatingGroup {
	return NoMaturityRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMaturityRules,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MaturityRuleID), quickfix.GroupElement(tag.MaturityMonthYearFormat), quickfix.GroupElement(tag.MaturityMonthYearIncrementUnits), quickfix.GroupElement(tag.StartMaturityMonthYear), quickfix.GroupElement(tag.EndMaturityMonthYear), quickfix.GroupElement(tag.MaturityMonthYearIncrement)})}
}

//Add create and append a new NoMaturityRules to this group
func (m NoMaturityRulesRepeatingGroup) Add() NoMaturityRules {
	g := m.RepeatingGroup.Add()
	return NoMaturityRules{g}
}

//Get returns the ith NoMaturityRules in the NoMaturityRulesRepeatinGroup
func (m NoMaturityRulesRepeatingGroup) Get(i int) NoMaturityRules {
	return NoMaturityRules{m.RepeatingGroup.Get(i)}
}

//NoStrikeRulesRepeatingGroup is a repeating group, Tag 1201
type NoStrikeRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoStrikeRulesRepeatingGroup returns an initialized, NoStrikeRulesRepeatingGroup
func NewNoStrikeRulesRepeatingGroup() NoStrikeRulesRepeatingGroup {
	return NoStrikeRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoStrikeRules,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.StrikeRuleID), quickfix.GroupElement(tag.StartStrikePxRange), quickfix.GroupElement(tag.EndStrikePxRange), quickfix.GroupElement(tag.StrikeIncrement), quickfix.GroupElement(tag.StrikeExerciseStyle), NewNoMaturityRulesRepeatingGroup()})}
}

//Add create and append a new NoStrikeRules to this group
func (m NoStrikeRulesRepeatingGroup) Add() NoStrikeRules {
	g := m.RepeatingGroup.Add()
	return NoStrikeRules{g}
}

//Get returns the ith NoStrikeRules in the NoStrikeRulesRepeatinGroup
func (m NoStrikeRulesRepeatingGroup) Get(i int) NoStrikeRules {
	return NoStrikeRules{m.RepeatingGroup.Get(i)}
}

//NoMarketSegmentsRepeatingGroup is a repeating group, Tag 1310
type NoMarketSegmentsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMarketSegmentsRepeatingGroup returns an initialized, NoMarketSegmentsRepeatingGroup
func NewNoMarketSegmentsRepeatingGroup() NoMarketSegmentsRepeatingGroup {
	return NoMarketSegmentsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMarketSegments,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MarketID), quickfix.GroupElement(tag.MarketSegmentID), NewNoTickRulesRepeatingGroup(), NewNoLotTypeRulesRepeatingGroup(), quickfix.GroupElement(tag.PriceLimitType), quickfix.GroupElement(tag.LowLimitPrice), quickfix.GroupElement(tag.HighLimitPrice), quickfix.GroupElement(tag.TradingReferencePrice), quickfix.GroupElement(tag.ExpirationCycle), quickfix.GroupElement(tag.MinTradeVol), quickfix.GroupElement(tag.MaxTradeVol), quickfix.GroupElement(tag.MaxPriceVariation), quickfix.GroupElement(tag.ImpliedMarketIndicator), quickfix.GroupElement(tag.TradingCurrency), quickfix.GroupElement(tag.RoundLot), quickfix.GroupElement(tag.MultilegModel), quickfix.GroupElement(tag.MultilegPriceMethod), quickfix.GroupElement(tag.PriceType), NewNoTradingSessionRulesRepeatingGroup(), NewNoNestedInstrAttribRepeatingGroup(), NewNoStrikeRulesRepeatingGroup()})}
}

//Add create and append a new NoMarketSegments to this group
func (m NoMarketSegmentsRepeatingGroup) Add() NoMarketSegments {
	g := m.RepeatingGroup.Add()
	return NoMarketSegments{g}
}

//Get returns the ith NoMarketSegments in the NoMarketSegmentsRepeatinGroup
func (m NoMarketSegmentsRepeatingGroup) Get(i int) NoMarketSegments {
	return NoMarketSegments{m.RepeatingGroup.Get(i)}
}

//NoDerivativeInstrAttrib is a repeating group element, Tag 1311
type NoDerivativeInstrAttrib struct {
	*quickfix.Group
}

//SetDerivativeInstrAttribType sets DerivativeInstrAttribType, Tag 1313
func (m NoDerivativeInstrAttrib) SetDerivativeInstrAttribType(v int) {
	m.Set(field.NewDerivativeInstrAttribType(v))
}

//SetDerivativeInstrAttribValue sets DerivativeInstrAttribValue, Tag 1314
func (m NoDerivativeInstrAttrib) SetDerivativeInstrAttribValue(v string) {
	m.Set(field.NewDerivativeInstrAttribValue(v))
}

//GetDerivativeInstrAttribType gets DerivativeInstrAttribType, Tag 1313
func (m NoDerivativeInstrAttrib) GetDerivativeInstrAttribType() (v int, err quickfix.MessageRejectError) {
	var f field.DerivativeInstrAttribTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDerivativeInstrAttribValue gets DerivativeInstrAttribValue, Tag 1314
func (m NoDerivativeInstrAttrib) GetDerivativeInstrAttribValue() (v string, err quickfix.MessageRejectError) {
	var f field.DerivativeInstrAttribValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasDerivativeInstrAttribType returns true if DerivativeInstrAttribType is present, Tag 1313
func (m NoDerivativeInstrAttrib) HasDerivativeInstrAttribType() bool {
	return m.Has(tag.DerivativeInstrAttribType)
}

//HasDerivativeInstrAttribValue returns true if DerivativeInstrAttribValue is present, Tag 1314
func (m NoDerivativeInstrAttrib) HasDerivativeInstrAttribValue() bool {
	return m.Has(tag.DerivativeInstrAttribValue)
}

//NoDerivativeInstrAttribRepeatingGroup is a repeating group, Tag 1311
type NoDerivativeInstrAttribRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoDerivativeInstrAttribRepeatingGroup returns an initialized, NoDerivativeInstrAttribRepeatingGroup
func NewNoDerivativeInstrAttribRepeatingGroup() NoDerivativeInstrAttribRepeatingGroup {
	return NoDerivativeInstrAttribRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoDerivativeInstrAttrib,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.DerivativeInstrAttribType), quickfix.GroupElement(tag.DerivativeInstrAttribValue)})}
}

//Add create and append a new NoDerivativeInstrAttrib to this group
func (m NoDerivativeInstrAttribRepeatingGroup) Add() NoDerivativeInstrAttrib {
	g := m.RepeatingGroup.Add()
	return NoDerivativeInstrAttrib{g}
}

//Get returns the ith NoDerivativeInstrAttrib in the NoDerivativeInstrAttribRepeatinGroup
func (m NoDerivativeInstrAttribRepeatingGroup) Get(i int) NoDerivativeInstrAttrib {
	return NoDerivativeInstrAttrib{m.RepeatingGroup.Get(i)}
}
