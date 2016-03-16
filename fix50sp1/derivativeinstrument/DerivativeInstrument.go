package derivativeinstrument

import (
	"github.com/quickfixgo/quickfix/fix50sp1/derivativeeventsgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/derivativeinstrumentparties"
	"github.com/quickfixgo/quickfix/fix50sp1/derivativesecurityaltidgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/derivativesecurityxml"
)

//DerivativeInstrument is a fix50sp1 Component
type DerivativeInstrument struct {
	//DerivativeSymbol is a non-required field for DerivativeInstrument.
	DerivativeSymbol *string `fix:"1214"`
	//DerivativeSymbolSfx is a non-required field for DerivativeInstrument.
	DerivativeSymbolSfx *string `fix:"1215"`
	//DerivativeSecurityID is a non-required field for DerivativeInstrument.
	DerivativeSecurityID *string `fix:"1216"`
	//DerivativeSecurityIDSource is a non-required field for DerivativeInstrument.
	DerivativeSecurityIDSource *string `fix:"1217"`
	//DerivativeSecurityAltIDGrp is a non-required component for DerivativeInstrument.
	DerivativeSecurityAltIDGrp *derivativesecurityaltidgrp.DerivativeSecurityAltIDGrp
	//DerivativeProduct is a non-required field for DerivativeInstrument.
	DerivativeProduct *int `fix:"1246"`
	//DerivativeProductComplex is a non-required field for DerivativeInstrument.
	DerivativeProductComplex *string `fix:"1228"`
	//DerivFlexProductEligibilityIndicator is a non-required field for DerivativeInstrument.
	DerivFlexProductEligibilityIndicator *bool `fix:"1243"`
	//DerivativeSecurityGroup is a non-required field for DerivativeInstrument.
	DerivativeSecurityGroup *string `fix:"1247"`
	//DerivativeCFICode is a non-required field for DerivativeInstrument.
	DerivativeCFICode *string `fix:"1248"`
	//DerivativeSecurityType is a non-required field for DerivativeInstrument.
	DerivativeSecurityType *string `fix:"1249"`
	//DerivativeSecuritySubType is a non-required field for DerivativeInstrument.
	DerivativeSecuritySubType *string `fix:"1250"`
	//DerivativeMaturityMonthYear is a non-required field for DerivativeInstrument.
	DerivativeMaturityMonthYear *string `fix:"1251"`
	//DerivativeMaturityDate is a non-required field for DerivativeInstrument.
	DerivativeMaturityDate *string `fix:"1252"`
	//DerivativeMaturityTime is a non-required field for DerivativeInstrument.
	DerivativeMaturityTime *string `fix:"1253"`
	//DerivativeSettleOnOpenFlag is a non-required field for DerivativeInstrument.
	DerivativeSettleOnOpenFlag *string `fix:"1254"`
	//DerivativeInstrmtAssignmentMethod is a non-required field for DerivativeInstrument.
	DerivativeInstrmtAssignmentMethod *string `fix:"1255"`
	//DerivativeSecurityStatus is a non-required field for DerivativeInstrument.
	DerivativeSecurityStatus *string `fix:"1256"`
	//DerivativeIssueDate is a non-required field for DerivativeInstrument.
	DerivativeIssueDate *string `fix:"1276"`
	//DerivativeInstrRegistry is a non-required field for DerivativeInstrument.
	DerivativeInstrRegistry *string `fix:"1257"`
	//DerivativeCountryOfIssue is a non-required field for DerivativeInstrument.
	DerivativeCountryOfIssue *string `fix:"1258"`
	//DerivativeStateOrProvinceOfIssue is a non-required field for DerivativeInstrument.
	DerivativeStateOrProvinceOfIssue *string `fix:"1259"`
	//DerivativeStrikePrice is a non-required field for DerivativeInstrument.
	DerivativeStrikePrice *float64 `fix:"1261"`
	//DerivativeLocaleOfIssue is a non-required field for DerivativeInstrument.
	DerivativeLocaleOfIssue *string `fix:"1260"`
	//DerivativeStrikeCurrency is a non-required field for DerivativeInstrument.
	DerivativeStrikeCurrency *string `fix:"1262"`
	//DerivativeStrikeMultiplier is a non-required field for DerivativeInstrument.
	DerivativeStrikeMultiplier *float64 `fix:"1263"`
	//DerivativeStrikeValue is a non-required field for DerivativeInstrument.
	DerivativeStrikeValue *float64 `fix:"1264"`
	//DerivativeOptAttribute is a non-required field for DerivativeInstrument.
	DerivativeOptAttribute *string `fix:"1265"`
	//DerivativeContractMultiplier is a non-required field for DerivativeInstrument.
	DerivativeContractMultiplier *float64 `fix:"1266"`
	//DerivativeMinPriceIncrement is a non-required field for DerivativeInstrument.
	DerivativeMinPriceIncrement *float64 `fix:"1267"`
	//DerivativeMinPriceIncrementAmount is a non-required field for DerivativeInstrument.
	DerivativeMinPriceIncrementAmount *float64 `fix:"1268"`
	//DerivativeUnitOfMeasure is a non-required field for DerivativeInstrument.
	DerivativeUnitOfMeasure *string `fix:"1269"`
	//DerivativeUnitOfMeasureQty is a non-required field for DerivativeInstrument.
	DerivativeUnitOfMeasureQty *float64 `fix:"1270"`
	//DerivativePriceUnitOfMeasure is a non-required field for DerivativeInstrument.
	DerivativePriceUnitOfMeasure *string `fix:"1315"`
	//DerivativePriceUnitOfMeasureQty is a non-required field for DerivativeInstrument.
	DerivativePriceUnitOfMeasureQty *float64 `fix:"1316"`
	//DerivativeExerciseStyle is a non-required field for DerivativeInstrument.
	DerivativeExerciseStyle *string `fix:"1299"`
	//DerivativeOptPayAmount is a non-required field for DerivativeInstrument.
	DerivativeOptPayAmount *float64 `fix:"1225"`
	//DerivativeTimeUnit is a non-required field for DerivativeInstrument.
	DerivativeTimeUnit *string `fix:"1271"`
	//DerivativeSecurityExchange is a non-required field for DerivativeInstrument.
	DerivativeSecurityExchange *string `fix:"1272"`
	//DerivativePositionLimit is a non-required field for DerivativeInstrument.
	DerivativePositionLimit *int `fix:"1273"`
	//DerivativeNTPositionLimit is a non-required field for DerivativeInstrument.
	DerivativeNTPositionLimit *int `fix:"1274"`
	//DerivativeIssuer is a non-required field for DerivativeInstrument.
	DerivativeIssuer *string `fix:"1275"`
	//DerivativeEncodedIssuerLen is a non-required field for DerivativeInstrument.
	DerivativeEncodedIssuerLen *int `fix:"1277"`
	//DerivativeEncodedIssuer is a non-required field for DerivativeInstrument.
	DerivativeEncodedIssuer *string `fix:"1278"`
	//DerivativeSecurityDesc is a non-required field for DerivativeInstrument.
	DerivativeSecurityDesc *string `fix:"1279"`
	//DerivativeEncodedSecurityDescLen is a non-required field for DerivativeInstrument.
	DerivativeEncodedSecurityDescLen *int `fix:"1280"`
	//DerivativeEncodedSecurityDesc is a non-required field for DerivativeInstrument.
	DerivativeEncodedSecurityDesc *string `fix:"1281"`
	//DerivativeContractSettlMonth is a non-required field for DerivativeInstrument.
	DerivativeContractSettlMonth *string `fix:"1285"`
	//DerivativeEventsGrp is a non-required component for DerivativeInstrument.
	DerivativeEventsGrp *derivativeeventsgrp.DerivativeEventsGrp
	//DerivativeInstrumentParties is a non-required component for DerivativeInstrument.
	DerivativeInstrumentParties *derivativeinstrumentparties.DerivativeInstrumentParties
	//DerivativeSettlMethod is a non-required field for DerivativeInstrument.
	DerivativeSettlMethod *string `fix:"1317"`
	//DerivativePriceQuoteMethod is a non-required field for DerivativeInstrument.
	DerivativePriceQuoteMethod *string `fix:"1318"`
	//DerivativeFuturesValuationMethod is a non-required field for DerivativeInstrument.
	DerivativeFuturesValuationMethod *string `fix:"1319"`
	//DerivativeListMethod is a non-required field for DerivativeInstrument.
	DerivativeListMethod *int `fix:"1320"`
	//DerivativeCapPrice is a non-required field for DerivativeInstrument.
	DerivativeCapPrice *float64 `fix:"1321"`
	//DerivativeFloorPrice is a non-required field for DerivativeInstrument.
	DerivativeFloorPrice *float64 `fix:"1322"`
	//DerivativePutOrCall is a non-required field for DerivativeInstrument.
	DerivativePutOrCall *int `fix:"1323"`
	//DerivativeSecurityXML is a non-required component for DerivativeInstrument.
	DerivativeSecurityXML *derivativesecurityxml.DerivativeSecurityXML
}

func (m *DerivativeInstrument) SetDerivativeSymbol(v string)     { m.DerivativeSymbol = &v }
func (m *DerivativeInstrument) SetDerivativeSymbolSfx(v string)  { m.DerivativeSymbolSfx = &v }
func (m *DerivativeInstrument) SetDerivativeSecurityID(v string) { m.DerivativeSecurityID = &v }
func (m *DerivativeInstrument) SetDerivativeSecurityIDSource(v string) {
	m.DerivativeSecurityIDSource = &v
}
func (m *DerivativeInstrument) SetDerivativeSecurityAltIDGrp(v derivativesecurityaltidgrp.DerivativeSecurityAltIDGrp) {
	m.DerivativeSecurityAltIDGrp = &v
}
func (m *DerivativeInstrument) SetDerivativeProduct(v int)           { m.DerivativeProduct = &v }
func (m *DerivativeInstrument) SetDerivativeProductComplex(v string) { m.DerivativeProductComplex = &v }
func (m *DerivativeInstrument) SetDerivFlexProductEligibilityIndicator(v bool) {
	m.DerivFlexProductEligibilityIndicator = &v
}
func (m *DerivativeInstrument) SetDerivativeSecurityGroup(v string) { m.DerivativeSecurityGroup = &v }
func (m *DerivativeInstrument) SetDerivativeCFICode(v string)       { m.DerivativeCFICode = &v }
func (m *DerivativeInstrument) SetDerivativeSecurityType(v string)  { m.DerivativeSecurityType = &v }
func (m *DerivativeInstrument) SetDerivativeSecuritySubType(v string) {
	m.DerivativeSecuritySubType = &v
}
func (m *DerivativeInstrument) SetDerivativeMaturityMonthYear(v string) {
	m.DerivativeMaturityMonthYear = &v
}
func (m *DerivativeInstrument) SetDerivativeMaturityDate(v string) { m.DerivativeMaturityDate = &v }
func (m *DerivativeInstrument) SetDerivativeMaturityTime(v string) { m.DerivativeMaturityTime = &v }
func (m *DerivativeInstrument) SetDerivativeSettleOnOpenFlag(v string) {
	m.DerivativeSettleOnOpenFlag = &v
}
func (m *DerivativeInstrument) SetDerivativeInstrmtAssignmentMethod(v string) {
	m.DerivativeInstrmtAssignmentMethod = &v
}
func (m *DerivativeInstrument) SetDerivativeSecurityStatus(v string) { m.DerivativeSecurityStatus = &v }
func (m *DerivativeInstrument) SetDerivativeIssueDate(v string)      { m.DerivativeIssueDate = &v }
func (m *DerivativeInstrument) SetDerivativeInstrRegistry(v string)  { m.DerivativeInstrRegistry = &v }
func (m *DerivativeInstrument) SetDerivativeCountryOfIssue(v string) { m.DerivativeCountryOfIssue = &v }
func (m *DerivativeInstrument) SetDerivativeStateOrProvinceOfIssue(v string) {
	m.DerivativeStateOrProvinceOfIssue = &v
}
func (m *DerivativeInstrument) SetDerivativeStrikePrice(v float64)   { m.DerivativeStrikePrice = &v }
func (m *DerivativeInstrument) SetDerivativeLocaleOfIssue(v string)  { m.DerivativeLocaleOfIssue = &v }
func (m *DerivativeInstrument) SetDerivativeStrikeCurrency(v string) { m.DerivativeStrikeCurrency = &v }
func (m *DerivativeInstrument) SetDerivativeStrikeMultiplier(v float64) {
	m.DerivativeStrikeMultiplier = &v
}
func (m *DerivativeInstrument) SetDerivativeStrikeValue(v float64) { m.DerivativeStrikeValue = &v }
func (m *DerivativeInstrument) SetDerivativeOptAttribute(v string) { m.DerivativeOptAttribute = &v }
func (m *DerivativeInstrument) SetDerivativeContractMultiplier(v float64) {
	m.DerivativeContractMultiplier = &v
}
func (m *DerivativeInstrument) SetDerivativeMinPriceIncrement(v float64) {
	m.DerivativeMinPriceIncrement = &v
}
func (m *DerivativeInstrument) SetDerivativeMinPriceIncrementAmount(v float64) {
	m.DerivativeMinPriceIncrementAmount = &v
}
func (m *DerivativeInstrument) SetDerivativeUnitOfMeasure(v string) { m.DerivativeUnitOfMeasure = &v }
func (m *DerivativeInstrument) SetDerivativeUnitOfMeasureQty(v float64) {
	m.DerivativeUnitOfMeasureQty = &v
}
func (m *DerivativeInstrument) SetDerivativePriceUnitOfMeasure(v string) {
	m.DerivativePriceUnitOfMeasure = &v
}
func (m *DerivativeInstrument) SetDerivativePriceUnitOfMeasureQty(v float64) {
	m.DerivativePriceUnitOfMeasureQty = &v
}
func (m *DerivativeInstrument) SetDerivativeExerciseStyle(v string) { m.DerivativeExerciseStyle = &v }
func (m *DerivativeInstrument) SetDerivativeOptPayAmount(v float64) { m.DerivativeOptPayAmount = &v }
func (m *DerivativeInstrument) SetDerivativeTimeUnit(v string)      { m.DerivativeTimeUnit = &v }
func (m *DerivativeInstrument) SetDerivativeSecurityExchange(v string) {
	m.DerivativeSecurityExchange = &v
}
func (m *DerivativeInstrument) SetDerivativePositionLimit(v int)    { m.DerivativePositionLimit = &v }
func (m *DerivativeInstrument) SetDerivativeNTPositionLimit(v int)  { m.DerivativeNTPositionLimit = &v }
func (m *DerivativeInstrument) SetDerivativeIssuer(v string)        { m.DerivativeIssuer = &v }
func (m *DerivativeInstrument) SetDerivativeEncodedIssuerLen(v int) { m.DerivativeEncodedIssuerLen = &v }
func (m *DerivativeInstrument) SetDerivativeEncodedIssuer(v string) { m.DerivativeEncodedIssuer = &v }
func (m *DerivativeInstrument) SetDerivativeSecurityDesc(v string)  { m.DerivativeSecurityDesc = &v }
func (m *DerivativeInstrument) SetDerivativeEncodedSecurityDescLen(v int) {
	m.DerivativeEncodedSecurityDescLen = &v
}
func (m *DerivativeInstrument) SetDerivativeEncodedSecurityDesc(v string) {
	m.DerivativeEncodedSecurityDesc = &v
}
func (m *DerivativeInstrument) SetDerivativeContractSettlMonth(v string) {
	m.DerivativeContractSettlMonth = &v
}
func (m *DerivativeInstrument) SetDerivativeEventsGrp(v derivativeeventsgrp.DerivativeEventsGrp) {
	m.DerivativeEventsGrp = &v
}
func (m *DerivativeInstrument) SetDerivativeInstrumentParties(v derivativeinstrumentparties.DerivativeInstrumentParties) {
	m.DerivativeInstrumentParties = &v
}
func (m *DerivativeInstrument) SetDerivativeSettlMethod(v string) { m.DerivativeSettlMethod = &v }
func (m *DerivativeInstrument) SetDerivativePriceQuoteMethod(v string) {
	m.DerivativePriceQuoteMethod = &v
}
func (m *DerivativeInstrument) SetDerivativeFuturesValuationMethod(v string) {
	m.DerivativeFuturesValuationMethod = &v
}
func (m *DerivativeInstrument) SetDerivativeListMethod(v int)     { m.DerivativeListMethod = &v }
func (m *DerivativeInstrument) SetDerivativeCapPrice(v float64)   { m.DerivativeCapPrice = &v }
func (m *DerivativeInstrument) SetDerivativeFloorPrice(v float64) { m.DerivativeFloorPrice = &v }
func (m *DerivativeInstrument) SetDerivativePutOrCall(v int)      { m.DerivativePutOrCall = &v }
func (m *DerivativeInstrument) SetDerivativeSecurityXML(v derivativesecurityxml.DerivativeSecurityXML) {
	m.DerivativeSecurityXML = &v
}
