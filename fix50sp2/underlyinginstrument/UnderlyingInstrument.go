package underlyinginstrument

import (
	"github.com/quickfixgo/quickfix/fix50sp2/undlyinstrumentptyssubgrp"
)

//NoUnderlyingSecurityAltID is a repeating group in UnderlyingInstrument
type NoUnderlyingSecurityAltID struct {
	//UnderlyingSecurityAltID is a non-required field for NoUnderlyingSecurityAltID.
	UnderlyingSecurityAltID *string `fix:"458"`
	//UnderlyingSecurityAltIDSource is a non-required field for NoUnderlyingSecurityAltID.
	UnderlyingSecurityAltIDSource *string `fix:"459"`
}

//NoUnderlyingStips is a repeating group in UnderlyingInstrument
type NoUnderlyingStips struct {
	//UnderlyingStipType is a non-required field for NoUnderlyingStips.
	UnderlyingStipType *string `fix:"888"`
	//UnderlyingStipValue is a non-required field for NoUnderlyingStips.
	UnderlyingStipValue *string `fix:"889"`
}

//NoUndlyInstrumentParties is a repeating group in UnderlyingInstrument
type NoUndlyInstrumentParties struct {
	//UnderlyingInstrumentPartyID is a non-required field for NoUndlyInstrumentParties.
	UnderlyingInstrumentPartyID *string `fix:"1059"`
	//UnderlyingInstrumentPartyIDSource is a non-required field for NoUndlyInstrumentParties.
	UnderlyingInstrumentPartyIDSource *string `fix:"1060"`
	//UnderlyingInstrumentPartyRole is a non-required field for NoUndlyInstrumentParties.
	UnderlyingInstrumentPartyRole *int `fix:"1061"`
	//UndlyInstrumentPtysSubGrp Component
	undlyinstrumentptyssubgrp.UndlyInstrumentPtysSubGrp
}

//UnderlyingInstrument is a fix50sp2 Component
type UnderlyingInstrument struct {
	//UnderlyingSymbol is a non-required field for UnderlyingInstrument.
	UnderlyingSymbol *string `fix:"311"`
	//UnderlyingSymbolSfx is a non-required field for UnderlyingInstrument.
	UnderlyingSymbolSfx *string `fix:"312"`
	//UnderlyingSecurityID is a non-required field for UnderlyingInstrument.
	UnderlyingSecurityID *string `fix:"309"`
	//UnderlyingSecurityIDSource is a non-required field for UnderlyingInstrument.
	UnderlyingSecurityIDSource *string `fix:"305"`
	//NoUnderlyingSecurityAltID is a non-required field for UnderlyingInstrument.
	NoUnderlyingSecurityAltID []NoUnderlyingSecurityAltID `fix:"457,omitempty"`
	//UnderlyingProduct is a non-required field for UnderlyingInstrument.
	UnderlyingProduct *int `fix:"462"`
	//UnderlyingCFICode is a non-required field for UnderlyingInstrument.
	UnderlyingCFICode *string `fix:"463"`
	//UnderlyingSecurityType is a non-required field for UnderlyingInstrument.
	UnderlyingSecurityType *string `fix:"310"`
	//UnderlyingSecuritySubType is a non-required field for UnderlyingInstrument.
	UnderlyingSecuritySubType *string `fix:"763"`
	//UnderlyingMaturityMonthYear is a non-required field for UnderlyingInstrument.
	UnderlyingMaturityMonthYear *string `fix:"313"`
	//UnderlyingMaturityDate is a non-required field for UnderlyingInstrument.
	UnderlyingMaturityDate *string `fix:"542"`
	//UnderlyingCouponPaymentDate is a non-required field for UnderlyingInstrument.
	UnderlyingCouponPaymentDate *string `fix:"241"`
	//UnderlyingIssueDate is a non-required field for UnderlyingInstrument.
	UnderlyingIssueDate *string `fix:"242"`
	//UnderlyingRepoCollateralSecurityType is a non-required field for UnderlyingInstrument.
	UnderlyingRepoCollateralSecurityType *int `fix:"243"`
	//UnderlyingRepurchaseTerm is a non-required field for UnderlyingInstrument.
	UnderlyingRepurchaseTerm *int `fix:"244"`
	//UnderlyingRepurchaseRate is a non-required field for UnderlyingInstrument.
	UnderlyingRepurchaseRate *float64 `fix:"245"`
	//UnderlyingFactor is a non-required field for UnderlyingInstrument.
	UnderlyingFactor *float64 `fix:"246"`
	//UnderlyingCreditRating is a non-required field for UnderlyingInstrument.
	UnderlyingCreditRating *string `fix:"256"`
	//UnderlyingInstrRegistry is a non-required field for UnderlyingInstrument.
	UnderlyingInstrRegistry *string `fix:"595"`
	//UnderlyingCountryOfIssue is a non-required field for UnderlyingInstrument.
	UnderlyingCountryOfIssue *string `fix:"592"`
	//UnderlyingStateOrProvinceOfIssue is a non-required field for UnderlyingInstrument.
	UnderlyingStateOrProvinceOfIssue *string `fix:"593"`
	//UnderlyingLocaleOfIssue is a non-required field for UnderlyingInstrument.
	UnderlyingLocaleOfIssue *string `fix:"594"`
	//UnderlyingRedemptionDate is a non-required field for UnderlyingInstrument.
	UnderlyingRedemptionDate *string `fix:"247"`
	//UnderlyingStrikePrice is a non-required field for UnderlyingInstrument.
	UnderlyingStrikePrice *float64 `fix:"316"`
	//UnderlyingStrikeCurrency is a non-required field for UnderlyingInstrument.
	UnderlyingStrikeCurrency *string `fix:"941"`
	//UnderlyingOptAttribute is a non-required field for UnderlyingInstrument.
	UnderlyingOptAttribute *string `fix:"317"`
	//UnderlyingContractMultiplier is a non-required field for UnderlyingInstrument.
	UnderlyingContractMultiplier *float64 `fix:"436"`
	//UnderlyingCouponRate is a non-required field for UnderlyingInstrument.
	UnderlyingCouponRate *float64 `fix:"435"`
	//UnderlyingSecurityExchange is a non-required field for UnderlyingInstrument.
	UnderlyingSecurityExchange *string `fix:"308"`
	//UnderlyingIssuer is a non-required field for UnderlyingInstrument.
	UnderlyingIssuer *string `fix:"306"`
	//EncodedUnderlyingIssuerLen is a non-required field for UnderlyingInstrument.
	EncodedUnderlyingIssuerLen *int `fix:"362"`
	//EncodedUnderlyingIssuer is a non-required field for UnderlyingInstrument.
	EncodedUnderlyingIssuer *string `fix:"363"`
	//UnderlyingSecurityDesc is a non-required field for UnderlyingInstrument.
	UnderlyingSecurityDesc *string `fix:"307"`
	//EncodedUnderlyingSecurityDescLen is a non-required field for UnderlyingInstrument.
	EncodedUnderlyingSecurityDescLen *int `fix:"364"`
	//EncodedUnderlyingSecurityDesc is a non-required field for UnderlyingInstrument.
	EncodedUnderlyingSecurityDesc *string `fix:"365"`
	//UnderlyingCPProgram is a non-required field for UnderlyingInstrument.
	UnderlyingCPProgram *string `fix:"877"`
	//UnderlyingCPRegType is a non-required field for UnderlyingInstrument.
	UnderlyingCPRegType *string `fix:"878"`
	//UnderlyingCurrency is a non-required field for UnderlyingInstrument.
	UnderlyingCurrency *string `fix:"318"`
	//UnderlyingQty is a non-required field for UnderlyingInstrument.
	UnderlyingQty *float64 `fix:"879"`
	//UnderlyingPx is a non-required field for UnderlyingInstrument.
	UnderlyingPx *float64 `fix:"810"`
	//UnderlyingDirtyPrice is a non-required field for UnderlyingInstrument.
	UnderlyingDirtyPrice *float64 `fix:"882"`
	//UnderlyingEndPrice is a non-required field for UnderlyingInstrument.
	UnderlyingEndPrice *float64 `fix:"883"`
	//UnderlyingStartValue is a non-required field for UnderlyingInstrument.
	UnderlyingStartValue *float64 `fix:"884"`
	//UnderlyingCurrentValue is a non-required field for UnderlyingInstrument.
	UnderlyingCurrentValue *float64 `fix:"885"`
	//UnderlyingEndValue is a non-required field for UnderlyingInstrument.
	UnderlyingEndValue *float64 `fix:"886"`
	//NoUnderlyingStips is a non-required field for UnderlyingInstrument.
	NoUnderlyingStips []NoUnderlyingStips `fix:"887,omitempty"`
	//UnderlyingAllocationPercent is a non-required field for UnderlyingInstrument.
	UnderlyingAllocationPercent *float64 `fix:"972"`
	//UnderlyingSettlementType is a non-required field for UnderlyingInstrument.
	UnderlyingSettlementType *int `fix:"975"`
	//UnderlyingCashAmount is a non-required field for UnderlyingInstrument.
	UnderlyingCashAmount *float64 `fix:"973"`
	//UnderlyingCashType is a non-required field for UnderlyingInstrument.
	UnderlyingCashType *string `fix:"974"`
	//UnderlyingUnitOfMeasure is a non-required field for UnderlyingInstrument.
	UnderlyingUnitOfMeasure *string `fix:"998"`
	//UnderlyingTimeUnit is a non-required field for UnderlyingInstrument.
	UnderlyingTimeUnit *string `fix:"1000"`
	//UnderlyingCapValue is a non-required field for UnderlyingInstrument.
	UnderlyingCapValue *float64 `fix:"1038"`
	//NoUndlyInstrumentParties is a non-required field for UnderlyingInstrument.
	NoUndlyInstrumentParties []NoUndlyInstrumentParties `fix:"1058,omitempty"`
	//UnderlyingSettlMethod is a non-required field for UnderlyingInstrument.
	UnderlyingSettlMethod *string `fix:"1039"`
	//UnderlyingAdjustedQuantity is a non-required field for UnderlyingInstrument.
	UnderlyingAdjustedQuantity *float64 `fix:"1044"`
	//UnderlyingFXRate is a non-required field for UnderlyingInstrument.
	UnderlyingFXRate *float64 `fix:"1045"`
	//UnderlyingFXRateCalc is a non-required field for UnderlyingInstrument.
	UnderlyingFXRateCalc *string `fix:"1046"`
	//UnderlyingMaturityTime is a non-required field for UnderlyingInstrument.
	UnderlyingMaturityTime *string `fix:"1213"`
	//UnderlyingPutOrCall is a non-required field for UnderlyingInstrument.
	UnderlyingPutOrCall *int `fix:"315"`
	//UnderlyingExerciseStyle is a non-required field for UnderlyingInstrument.
	UnderlyingExerciseStyle *int `fix:"1419"`
	//UnderlyingUnitOfMeasureQty is a non-required field for UnderlyingInstrument.
	UnderlyingUnitOfMeasureQty *float64 `fix:"1423"`
	//UnderlyingPriceUnitOfMeasure is a non-required field for UnderlyingInstrument.
	UnderlyingPriceUnitOfMeasure *string `fix:"1424"`
	//UnderlyingPriceUnitOfMeasureQty is a non-required field for UnderlyingInstrument.
	UnderlyingPriceUnitOfMeasureQty *float64 `fix:"1425"`
	//UnderlyingContractMultiplierUnit is a non-required field for UnderlyingInstrument.
	UnderlyingContractMultiplierUnit *int `fix:"1437"`
	//UnderlyingFlowScheduleType is a non-required field for UnderlyingInstrument.
	UnderlyingFlowScheduleType *int `fix:"1441"`
	//UnderlyingRestructuringType is a non-required field for UnderlyingInstrument.
	UnderlyingRestructuringType *string `fix:"1453"`
	//UnderlyingSeniority is a non-required field for UnderlyingInstrument.
	UnderlyingSeniority *string `fix:"1454"`
	//UnderlyingNotionalPercentageOutstanding is a non-required field for UnderlyingInstrument.
	UnderlyingNotionalPercentageOutstanding *float64 `fix:"1455"`
	//UnderlyingOriginalNotionalPercentageOutstanding is a non-required field for UnderlyingInstrument.
	UnderlyingOriginalNotionalPercentageOutstanding *float64 `fix:"1456"`
	//UnderlyingAttachmentPoint is a non-required field for UnderlyingInstrument.
	UnderlyingAttachmentPoint *float64 `fix:"1459"`
	//UnderlyingDetachmentPoint is a non-required field for UnderlyingInstrument.
	UnderlyingDetachmentPoint *float64 `fix:"1460"`
}

func (m *UnderlyingInstrument) SetUnderlyingSymbol(v string)     { m.UnderlyingSymbol = &v }
func (m *UnderlyingInstrument) SetUnderlyingSymbolSfx(v string)  { m.UnderlyingSymbolSfx = &v }
func (m *UnderlyingInstrument) SetUnderlyingSecurityID(v string) { m.UnderlyingSecurityID = &v }
func (m *UnderlyingInstrument) SetUnderlyingSecurityIDSource(v string) {
	m.UnderlyingSecurityIDSource = &v
}
func (m *UnderlyingInstrument) SetNoUnderlyingSecurityAltID(v []NoUnderlyingSecurityAltID) {
	m.NoUnderlyingSecurityAltID = v
}
func (m *UnderlyingInstrument) SetUnderlyingProduct(v int)         { m.UnderlyingProduct = &v }
func (m *UnderlyingInstrument) SetUnderlyingCFICode(v string)      { m.UnderlyingCFICode = &v }
func (m *UnderlyingInstrument) SetUnderlyingSecurityType(v string) { m.UnderlyingSecurityType = &v }
func (m *UnderlyingInstrument) SetUnderlyingSecuritySubType(v string) {
	m.UnderlyingSecuritySubType = &v
}
func (m *UnderlyingInstrument) SetUnderlyingMaturityMonthYear(v string) {
	m.UnderlyingMaturityMonthYear = &v
}
func (m *UnderlyingInstrument) SetUnderlyingMaturityDate(v string) { m.UnderlyingMaturityDate = &v }
func (m *UnderlyingInstrument) SetUnderlyingCouponPaymentDate(v string) {
	m.UnderlyingCouponPaymentDate = &v
}
func (m *UnderlyingInstrument) SetUnderlyingIssueDate(v string) { m.UnderlyingIssueDate = &v }
func (m *UnderlyingInstrument) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.UnderlyingRepoCollateralSecurityType = &v
}
func (m *UnderlyingInstrument) SetUnderlyingRepurchaseTerm(v int)     { m.UnderlyingRepurchaseTerm = &v }
func (m *UnderlyingInstrument) SetUnderlyingRepurchaseRate(v float64) { m.UnderlyingRepurchaseRate = &v }
func (m *UnderlyingInstrument) SetUnderlyingFactor(v float64)         { m.UnderlyingFactor = &v }
func (m *UnderlyingInstrument) SetUnderlyingCreditRating(v string)    { m.UnderlyingCreditRating = &v }
func (m *UnderlyingInstrument) SetUnderlyingInstrRegistry(v string)   { m.UnderlyingInstrRegistry = &v }
func (m *UnderlyingInstrument) SetUnderlyingCountryOfIssue(v string)  { m.UnderlyingCountryOfIssue = &v }
func (m *UnderlyingInstrument) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.UnderlyingStateOrProvinceOfIssue = &v
}
func (m *UnderlyingInstrument) SetUnderlyingLocaleOfIssue(v string)  { m.UnderlyingLocaleOfIssue = &v }
func (m *UnderlyingInstrument) SetUnderlyingRedemptionDate(v string) { m.UnderlyingRedemptionDate = &v }
func (m *UnderlyingInstrument) SetUnderlyingStrikePrice(v float64)   { m.UnderlyingStrikePrice = &v }
func (m *UnderlyingInstrument) SetUnderlyingStrikeCurrency(v string) { m.UnderlyingStrikeCurrency = &v }
func (m *UnderlyingInstrument) SetUnderlyingOptAttribute(v string)   { m.UnderlyingOptAttribute = &v }
func (m *UnderlyingInstrument) SetUnderlyingContractMultiplier(v float64) {
	m.UnderlyingContractMultiplier = &v
}
func (m *UnderlyingInstrument) SetUnderlyingCouponRate(v float64) { m.UnderlyingCouponRate = &v }
func (m *UnderlyingInstrument) SetUnderlyingSecurityExchange(v string) {
	m.UnderlyingSecurityExchange = &v
}
func (m *UnderlyingInstrument) SetUnderlyingIssuer(v string)        { m.UnderlyingIssuer = &v }
func (m *UnderlyingInstrument) SetEncodedUnderlyingIssuerLen(v int) { m.EncodedUnderlyingIssuerLen = &v }
func (m *UnderlyingInstrument) SetEncodedUnderlyingIssuer(v string) { m.EncodedUnderlyingIssuer = &v }
func (m *UnderlyingInstrument) SetUnderlyingSecurityDesc(v string)  { m.UnderlyingSecurityDesc = &v }
func (m *UnderlyingInstrument) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.EncodedUnderlyingSecurityDescLen = &v
}
func (m *UnderlyingInstrument) SetEncodedUnderlyingSecurityDesc(v string) {
	m.EncodedUnderlyingSecurityDesc = &v
}
func (m *UnderlyingInstrument) SetUnderlyingCPProgram(v string)            { m.UnderlyingCPProgram = &v }
func (m *UnderlyingInstrument) SetUnderlyingCPRegType(v string)            { m.UnderlyingCPRegType = &v }
func (m *UnderlyingInstrument) SetUnderlyingCurrency(v string)             { m.UnderlyingCurrency = &v }
func (m *UnderlyingInstrument) SetUnderlyingQty(v float64)                 { m.UnderlyingQty = &v }
func (m *UnderlyingInstrument) SetUnderlyingPx(v float64)                  { m.UnderlyingPx = &v }
func (m *UnderlyingInstrument) SetUnderlyingDirtyPrice(v float64)          { m.UnderlyingDirtyPrice = &v }
func (m *UnderlyingInstrument) SetUnderlyingEndPrice(v float64)            { m.UnderlyingEndPrice = &v }
func (m *UnderlyingInstrument) SetUnderlyingStartValue(v float64)          { m.UnderlyingStartValue = &v }
func (m *UnderlyingInstrument) SetUnderlyingCurrentValue(v float64)        { m.UnderlyingCurrentValue = &v }
func (m *UnderlyingInstrument) SetUnderlyingEndValue(v float64)            { m.UnderlyingEndValue = &v }
func (m *UnderlyingInstrument) SetNoUnderlyingStips(v []NoUnderlyingStips) { m.NoUnderlyingStips = v }
func (m *UnderlyingInstrument) SetUnderlyingAllocationPercent(v float64) {
	m.UnderlyingAllocationPercent = &v
}
func (m *UnderlyingInstrument) SetUnderlyingSettlementType(v int)   { m.UnderlyingSettlementType = &v }
func (m *UnderlyingInstrument) SetUnderlyingCashAmount(v float64)   { m.UnderlyingCashAmount = &v }
func (m *UnderlyingInstrument) SetUnderlyingCashType(v string)      { m.UnderlyingCashType = &v }
func (m *UnderlyingInstrument) SetUnderlyingUnitOfMeasure(v string) { m.UnderlyingUnitOfMeasure = &v }
func (m *UnderlyingInstrument) SetUnderlyingTimeUnit(v string)      { m.UnderlyingTimeUnit = &v }
func (m *UnderlyingInstrument) SetUnderlyingCapValue(v float64)     { m.UnderlyingCapValue = &v }
func (m *UnderlyingInstrument) SetNoUndlyInstrumentParties(v []NoUndlyInstrumentParties) {
	m.NoUndlyInstrumentParties = v
}
func (m *UnderlyingInstrument) SetUnderlyingSettlMethod(v string) { m.UnderlyingSettlMethod = &v }
func (m *UnderlyingInstrument) SetUnderlyingAdjustedQuantity(v float64) {
	m.UnderlyingAdjustedQuantity = &v
}
func (m *UnderlyingInstrument) SetUnderlyingFXRate(v float64)      { m.UnderlyingFXRate = &v }
func (m *UnderlyingInstrument) SetUnderlyingFXRateCalc(v string)   { m.UnderlyingFXRateCalc = &v }
func (m *UnderlyingInstrument) SetUnderlyingMaturityTime(v string) { m.UnderlyingMaturityTime = &v }
func (m *UnderlyingInstrument) SetUnderlyingPutOrCall(v int)       { m.UnderlyingPutOrCall = &v }
func (m *UnderlyingInstrument) SetUnderlyingExerciseStyle(v int)   { m.UnderlyingExerciseStyle = &v }
func (m *UnderlyingInstrument) SetUnderlyingUnitOfMeasureQty(v float64) {
	m.UnderlyingUnitOfMeasureQty = &v
}
func (m *UnderlyingInstrument) SetUnderlyingPriceUnitOfMeasure(v string) {
	m.UnderlyingPriceUnitOfMeasure = &v
}
func (m *UnderlyingInstrument) SetUnderlyingPriceUnitOfMeasureQty(v float64) {
	m.UnderlyingPriceUnitOfMeasureQty = &v
}
func (m *UnderlyingInstrument) SetUnderlyingContractMultiplierUnit(v int) {
	m.UnderlyingContractMultiplierUnit = &v
}
func (m *UnderlyingInstrument) SetUnderlyingFlowScheduleType(v int) { m.UnderlyingFlowScheduleType = &v }
func (m *UnderlyingInstrument) SetUnderlyingRestructuringType(v string) {
	m.UnderlyingRestructuringType = &v
}
func (m *UnderlyingInstrument) SetUnderlyingSeniority(v string) { m.UnderlyingSeniority = &v }
func (m *UnderlyingInstrument) SetUnderlyingNotionalPercentageOutstanding(v float64) {
	m.UnderlyingNotionalPercentageOutstanding = &v
}
func (m *UnderlyingInstrument) SetUnderlyingOriginalNotionalPercentageOutstanding(v float64) {
	m.UnderlyingOriginalNotionalPercentageOutstanding = &v
}
func (m *UnderlyingInstrument) SetUnderlyingAttachmentPoint(v float64) {
	m.UnderlyingAttachmentPoint = &v
}
func (m *UnderlyingInstrument) SetUnderlyingDetachmentPoint(v float64) {
	m.UnderlyingDetachmentPoint = &v
}
