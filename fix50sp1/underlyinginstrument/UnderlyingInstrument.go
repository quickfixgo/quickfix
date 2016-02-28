package underlyinginstrument

import (
	"github.com/quickfixgo/quickfix/fix50sp1/undlyinstrumentptyssubgrp"
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
	//UndlyInstrumentPartyID is a non-required field for NoUndlyInstrumentParties.
	UndlyInstrumentPartyID *string `fix:"1059"`
	//UndlyInstrumentPartyIDSource is a non-required field for NoUndlyInstrumentParties.
	UndlyInstrumentPartyIDSource *string `fix:"1060"`
	//UndlyInstrumentPartyRole is a non-required field for NoUndlyInstrumentParties.
	UndlyInstrumentPartyRole *int `fix:"1061"`
	//UndlyInstrumentPtysSubGrp Component
	UndlyInstrumentPtysSubGrp undlyinstrumentptyssubgrp.Component
}

//Component is a fix50sp1 UnderlyingInstrument Component
type Component struct {
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
}

func New() *Component { return new(Component) }

func (m *Component) SetUnderlyingSymbol(v string)           { m.UnderlyingSymbol = &v }
func (m *Component) SetUnderlyingSymbolSfx(v string)        { m.UnderlyingSymbolSfx = &v }
func (m *Component) SetUnderlyingSecurityID(v string)       { m.UnderlyingSecurityID = &v }
func (m *Component) SetUnderlyingSecurityIDSource(v string) { m.UnderlyingSecurityIDSource = &v }
func (m *Component) SetNoUnderlyingSecurityAltID(v []NoUnderlyingSecurityAltID) {
	m.NoUnderlyingSecurityAltID = v
}
func (m *Component) SetUnderlyingProduct(v int)              { m.UnderlyingProduct = &v }
func (m *Component) SetUnderlyingCFICode(v string)           { m.UnderlyingCFICode = &v }
func (m *Component) SetUnderlyingSecurityType(v string)      { m.UnderlyingSecurityType = &v }
func (m *Component) SetUnderlyingSecuritySubType(v string)   { m.UnderlyingSecuritySubType = &v }
func (m *Component) SetUnderlyingMaturityMonthYear(v string) { m.UnderlyingMaturityMonthYear = &v }
func (m *Component) SetUnderlyingMaturityDate(v string)      { m.UnderlyingMaturityDate = &v }
func (m *Component) SetUnderlyingCouponPaymentDate(v string) { m.UnderlyingCouponPaymentDate = &v }
func (m *Component) SetUnderlyingIssueDate(v string)         { m.UnderlyingIssueDate = &v }
func (m *Component) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.UnderlyingRepoCollateralSecurityType = &v
}
func (m *Component) SetUnderlyingRepurchaseTerm(v int)     { m.UnderlyingRepurchaseTerm = &v }
func (m *Component) SetUnderlyingRepurchaseRate(v float64) { m.UnderlyingRepurchaseRate = &v }
func (m *Component) SetUnderlyingFactor(v float64)         { m.UnderlyingFactor = &v }
func (m *Component) SetUnderlyingCreditRating(v string)    { m.UnderlyingCreditRating = &v }
func (m *Component) SetUnderlyingInstrRegistry(v string)   { m.UnderlyingInstrRegistry = &v }
func (m *Component) SetUnderlyingCountryOfIssue(v string)  { m.UnderlyingCountryOfIssue = &v }
func (m *Component) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.UnderlyingStateOrProvinceOfIssue = &v
}
func (m *Component) SetUnderlyingLocaleOfIssue(v string)       { m.UnderlyingLocaleOfIssue = &v }
func (m *Component) SetUnderlyingRedemptionDate(v string)      { m.UnderlyingRedemptionDate = &v }
func (m *Component) SetUnderlyingStrikePrice(v float64)        { m.UnderlyingStrikePrice = &v }
func (m *Component) SetUnderlyingStrikeCurrency(v string)      { m.UnderlyingStrikeCurrency = &v }
func (m *Component) SetUnderlyingOptAttribute(v string)        { m.UnderlyingOptAttribute = &v }
func (m *Component) SetUnderlyingContractMultiplier(v float64) { m.UnderlyingContractMultiplier = &v }
func (m *Component) SetUnderlyingCouponRate(v float64)         { m.UnderlyingCouponRate = &v }
func (m *Component) SetUnderlyingSecurityExchange(v string)    { m.UnderlyingSecurityExchange = &v }
func (m *Component) SetUnderlyingIssuer(v string)              { m.UnderlyingIssuer = &v }
func (m *Component) SetEncodedUnderlyingIssuerLen(v int)       { m.EncodedUnderlyingIssuerLen = &v }
func (m *Component) SetEncodedUnderlyingIssuer(v string)       { m.EncodedUnderlyingIssuer = &v }
func (m *Component) SetUnderlyingSecurityDesc(v string)        { m.UnderlyingSecurityDesc = &v }
func (m *Component) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.EncodedUnderlyingSecurityDescLen = &v
}
func (m *Component) SetEncodedUnderlyingSecurityDesc(v string)  { m.EncodedUnderlyingSecurityDesc = &v }
func (m *Component) SetUnderlyingCPProgram(v string)            { m.UnderlyingCPProgram = &v }
func (m *Component) SetUnderlyingCPRegType(v string)            { m.UnderlyingCPRegType = &v }
func (m *Component) SetUnderlyingCurrency(v string)             { m.UnderlyingCurrency = &v }
func (m *Component) SetUnderlyingQty(v float64)                 { m.UnderlyingQty = &v }
func (m *Component) SetUnderlyingPx(v float64)                  { m.UnderlyingPx = &v }
func (m *Component) SetUnderlyingDirtyPrice(v float64)          { m.UnderlyingDirtyPrice = &v }
func (m *Component) SetUnderlyingEndPrice(v float64)            { m.UnderlyingEndPrice = &v }
func (m *Component) SetUnderlyingStartValue(v float64)          { m.UnderlyingStartValue = &v }
func (m *Component) SetUnderlyingCurrentValue(v float64)        { m.UnderlyingCurrentValue = &v }
func (m *Component) SetUnderlyingEndValue(v float64)            { m.UnderlyingEndValue = &v }
func (m *Component) SetNoUnderlyingStips(v []NoUnderlyingStips) { m.NoUnderlyingStips = v }
func (m *Component) SetUnderlyingAllocationPercent(v float64)   { m.UnderlyingAllocationPercent = &v }
func (m *Component) SetUnderlyingSettlementType(v int)          { m.UnderlyingSettlementType = &v }
func (m *Component) SetUnderlyingCashAmount(v float64)          { m.UnderlyingCashAmount = &v }
func (m *Component) SetUnderlyingCashType(v string)             { m.UnderlyingCashType = &v }
func (m *Component) SetUnderlyingUnitOfMeasure(v string)        { m.UnderlyingUnitOfMeasure = &v }
func (m *Component) SetUnderlyingTimeUnit(v string)             { m.UnderlyingTimeUnit = &v }
func (m *Component) SetUnderlyingCapValue(v float64)            { m.UnderlyingCapValue = &v }
func (m *Component) SetNoUndlyInstrumentParties(v []NoUndlyInstrumentParties) {
	m.NoUndlyInstrumentParties = v
}
func (m *Component) SetUnderlyingSettlMethod(v string)        { m.UnderlyingSettlMethod = &v }
func (m *Component) SetUnderlyingAdjustedQuantity(v float64)  { m.UnderlyingAdjustedQuantity = &v }
func (m *Component) SetUnderlyingFXRate(v float64)            { m.UnderlyingFXRate = &v }
func (m *Component) SetUnderlyingFXRateCalc(v string)         { m.UnderlyingFXRateCalc = &v }
func (m *Component) SetUnderlyingMaturityTime(v string)       { m.UnderlyingMaturityTime = &v }
func (m *Component) SetUnderlyingPutOrCall(v int)             { m.UnderlyingPutOrCall = &v }
func (m *Component) SetUnderlyingExerciseStyle(v int)         { m.UnderlyingExerciseStyle = &v }
func (m *Component) SetUnderlyingUnitOfMeasureQty(v float64)  { m.UnderlyingUnitOfMeasureQty = &v }
func (m *Component) SetUnderlyingPriceUnitOfMeasure(v string) { m.UnderlyingPriceUnitOfMeasure = &v }
func (m *Component) SetUnderlyingPriceUnitOfMeasureQty(v float64) {
	m.UnderlyingPriceUnitOfMeasureQty = &v
}
