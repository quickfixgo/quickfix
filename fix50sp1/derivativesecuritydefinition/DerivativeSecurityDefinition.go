package derivativesecuritydefinition

import (
	"github.com/quickfixgo/quickfix/fix50sp1/derivativeinstrumentpartysubidsgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/securitytradingrules"
	"github.com/quickfixgo/quickfix/fix50sp1/strikerules"
	"time"
)

//NoDerivativeSecurityAltID is a repeating group in DerivativeSecurityDefinition
type NoDerivativeSecurityAltID struct {
	//DerivativeSecurityAltID is a non-required field for NoDerivativeSecurityAltID.
	DerivativeSecurityAltID *string `fix:"1219"`
	//DerivativeSecurityAltIDSource is a non-required field for NoDerivativeSecurityAltID.
	DerivativeSecurityAltIDSource *string `fix:"1220"`
}

//NoDerivativeEvents is a repeating group in DerivativeSecurityDefinition
type NoDerivativeEvents struct {
	//DerivativeEventType is a non-required field for NoDerivativeEvents.
	DerivativeEventType *int `fix:"1287"`
	//DerivativeEventDate is a non-required field for NoDerivativeEvents.
	DerivativeEventDate *string `fix:"1288"`
	//DerivativeEventTime is a non-required field for NoDerivativeEvents.
	DerivativeEventTime *time.Time `fix:"1289"`
	//DerivativeEventPx is a non-required field for NoDerivativeEvents.
	DerivativeEventPx *float64 `fix:"1290"`
	//DerivativeEventText is a non-required field for NoDerivativeEvents.
	DerivativeEventText *string `fix:"1291"`
}

//NoDerivativeInstrumentParties is a repeating group in DerivativeSecurityDefinition
type NoDerivativeInstrumentParties struct {
	//DerivativeInstrumentPartyID is a non-required field for NoDerivativeInstrumentParties.
	DerivativeInstrumentPartyID *string `fix:"1293"`
	//DerivativeInstrumentPartyIDSource is a non-required field for NoDerivativeInstrumentParties.
	DerivativeInstrumentPartyIDSource *string `fix:"1294"`
	//DerivativeInstrumentPartyRole is a non-required field for NoDerivativeInstrumentParties.
	DerivativeInstrumentPartyRole *int `fix:"1295"`
	//DerivativeInstrumentPartySubIDsGrp Component
	derivativeinstrumentpartysubidsgrp.DerivativeInstrumentPartySubIDsGrp
}

//NoDerivativeInstrAttrib is a repeating group in DerivativeSecurityDefinition
type NoDerivativeInstrAttrib struct {
	//DerivativeInstrAttribType is a non-required field for NoDerivativeInstrAttrib.
	DerivativeInstrAttribType *int `fix:"1313"`
	//DerivativeInstrAttribValue is a non-required field for NoDerivativeInstrAttrib.
	DerivativeInstrAttribValue *string `fix:"1314"`
}

//NoMarketSegments is a repeating group in DerivativeSecurityDefinition
type NoMarketSegments struct {
	//MarketID is a non-required field for NoMarketSegments.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for NoMarketSegments.
	MarketSegmentID *string `fix:"1300"`
	//SecurityTradingRules Component
	securitytradingrules.SecurityTradingRules
	//StrikeRules Component
	strikerules.StrikeRules
}

//DerivativeSecurityDefinition is a fix50sp1 Component
type DerivativeSecurityDefinition struct {
	//DerivativeSymbol is a non-required field for DerivativeSecurityDefinition.
	DerivativeSymbol *string `fix:"1214"`
	//DerivativeSymbolSfx is a non-required field for DerivativeSecurityDefinition.
	DerivativeSymbolSfx *string `fix:"1215"`
	//DerivativeSecurityID is a non-required field for DerivativeSecurityDefinition.
	DerivativeSecurityID *string `fix:"1216"`
	//DerivativeSecurityIDSource is a non-required field for DerivativeSecurityDefinition.
	DerivativeSecurityIDSource *string `fix:"1217"`
	//NoDerivativeSecurityAltID is a non-required field for DerivativeSecurityDefinition.
	NoDerivativeSecurityAltID []NoDerivativeSecurityAltID `fix:"1218,omitempty"`
	//DerivativeProduct is a non-required field for DerivativeSecurityDefinition.
	DerivativeProduct *int `fix:"1246"`
	//DerivativeProductComplex is a non-required field for DerivativeSecurityDefinition.
	DerivativeProductComplex *string `fix:"1228"`
	//DerivFlexProductEligibilityIndicator is a non-required field for DerivativeSecurityDefinition.
	DerivFlexProductEligibilityIndicator *bool `fix:"1243"`
	//DerivativeSecurityGroup is a non-required field for DerivativeSecurityDefinition.
	DerivativeSecurityGroup *string `fix:"1247"`
	//DerivativeCFICode is a non-required field for DerivativeSecurityDefinition.
	DerivativeCFICode *string `fix:"1248"`
	//DerivativeSecurityType is a non-required field for DerivativeSecurityDefinition.
	DerivativeSecurityType *string `fix:"1249"`
	//DerivativeSecuritySubType is a non-required field for DerivativeSecurityDefinition.
	DerivativeSecuritySubType *string `fix:"1250"`
	//DerivativeMaturityMonthYear is a non-required field for DerivativeSecurityDefinition.
	DerivativeMaturityMonthYear *string `fix:"1251"`
	//DerivativeMaturityDate is a non-required field for DerivativeSecurityDefinition.
	DerivativeMaturityDate *string `fix:"1252"`
	//DerivativeMaturityTime is a non-required field for DerivativeSecurityDefinition.
	DerivativeMaturityTime *string `fix:"1253"`
	//DerivativeSettleOnOpenFlag is a non-required field for DerivativeSecurityDefinition.
	DerivativeSettleOnOpenFlag *string `fix:"1254"`
	//DerivativeInstrmtAssignmentMethod is a non-required field for DerivativeSecurityDefinition.
	DerivativeInstrmtAssignmentMethod *string `fix:"1255"`
	//DerivativeSecurityStatus is a non-required field for DerivativeSecurityDefinition.
	DerivativeSecurityStatus *string `fix:"1256"`
	//DerivativeIssueDate is a non-required field for DerivativeSecurityDefinition.
	DerivativeIssueDate *string `fix:"1276"`
	//DerivativeInstrRegistry is a non-required field for DerivativeSecurityDefinition.
	DerivativeInstrRegistry *string `fix:"1257"`
	//DerivativeCountryOfIssue is a non-required field for DerivativeSecurityDefinition.
	DerivativeCountryOfIssue *string `fix:"1258"`
	//DerivativeStateOrProvinceOfIssue is a non-required field for DerivativeSecurityDefinition.
	DerivativeStateOrProvinceOfIssue *string `fix:"1259"`
	//DerivativeStrikePrice is a non-required field for DerivativeSecurityDefinition.
	DerivativeStrikePrice *float64 `fix:"1261"`
	//DerivativeLocaleOfIssue is a non-required field for DerivativeSecurityDefinition.
	DerivativeLocaleOfIssue *string `fix:"1260"`
	//DerivativeStrikeCurrency is a non-required field for DerivativeSecurityDefinition.
	DerivativeStrikeCurrency *string `fix:"1262"`
	//DerivativeStrikeMultiplier is a non-required field for DerivativeSecurityDefinition.
	DerivativeStrikeMultiplier *float64 `fix:"1263"`
	//DerivativeStrikeValue is a non-required field for DerivativeSecurityDefinition.
	DerivativeStrikeValue *float64 `fix:"1264"`
	//DerivativeOptAttribute is a non-required field for DerivativeSecurityDefinition.
	DerivativeOptAttribute *string `fix:"1265"`
	//DerivativeContractMultiplier is a non-required field for DerivativeSecurityDefinition.
	DerivativeContractMultiplier *float64 `fix:"1266"`
	//DerivativeMinPriceIncrement is a non-required field for DerivativeSecurityDefinition.
	DerivativeMinPriceIncrement *float64 `fix:"1267"`
	//DerivativeMinPriceIncrementAmount is a non-required field for DerivativeSecurityDefinition.
	DerivativeMinPriceIncrementAmount *float64 `fix:"1268"`
	//DerivativeUnitOfMeasure is a non-required field for DerivativeSecurityDefinition.
	DerivativeUnitOfMeasure *string `fix:"1269"`
	//DerivativeUnitOfMeasureQty is a non-required field for DerivativeSecurityDefinition.
	DerivativeUnitOfMeasureQty *float64 `fix:"1270"`
	//DerivativePriceUnitOfMeasure is a non-required field for DerivativeSecurityDefinition.
	DerivativePriceUnitOfMeasure *string `fix:"1315"`
	//DerivativePriceUnitOfMeasureQty is a non-required field for DerivativeSecurityDefinition.
	DerivativePriceUnitOfMeasureQty *float64 `fix:"1316"`
	//DerivativeExerciseStyle is a non-required field for DerivativeSecurityDefinition.
	DerivativeExerciseStyle *string `fix:"1299"`
	//DerivativeOptPayAmount is a non-required field for DerivativeSecurityDefinition.
	DerivativeOptPayAmount *float64 `fix:"1225"`
	//DerivativeTimeUnit is a non-required field for DerivativeSecurityDefinition.
	DerivativeTimeUnit *string `fix:"1271"`
	//DerivativeSecurityExchange is a non-required field for DerivativeSecurityDefinition.
	DerivativeSecurityExchange *string `fix:"1272"`
	//DerivativePositionLimit is a non-required field for DerivativeSecurityDefinition.
	DerivativePositionLimit *int `fix:"1273"`
	//DerivativeNTPositionLimit is a non-required field for DerivativeSecurityDefinition.
	DerivativeNTPositionLimit *int `fix:"1274"`
	//DerivativeIssuer is a non-required field for DerivativeSecurityDefinition.
	DerivativeIssuer *string `fix:"1275"`
	//DerivativeEncodedIssuerLen is a non-required field for DerivativeSecurityDefinition.
	DerivativeEncodedIssuerLen *int `fix:"1277"`
	//DerivativeEncodedIssuer is a non-required field for DerivativeSecurityDefinition.
	DerivativeEncodedIssuer *string `fix:"1278"`
	//DerivativeSecurityDesc is a non-required field for DerivativeSecurityDefinition.
	DerivativeSecurityDesc *string `fix:"1279"`
	//DerivativeEncodedSecurityDescLen is a non-required field for DerivativeSecurityDefinition.
	DerivativeEncodedSecurityDescLen *int `fix:"1280"`
	//DerivativeEncodedSecurityDesc is a non-required field for DerivativeSecurityDefinition.
	DerivativeEncodedSecurityDesc *string `fix:"1281"`
	//DerivativeContractSettlMonth is a non-required field for DerivativeSecurityDefinition.
	DerivativeContractSettlMonth *string `fix:"1285"`
	//NoDerivativeEvents is a non-required field for DerivativeSecurityDefinition.
	NoDerivativeEvents []NoDerivativeEvents `fix:"1286,omitempty"`
	//NoDerivativeInstrumentParties is a non-required field for DerivativeSecurityDefinition.
	NoDerivativeInstrumentParties []NoDerivativeInstrumentParties `fix:"1292,omitempty"`
	//DerivativeSettlMethod is a non-required field for DerivativeSecurityDefinition.
	DerivativeSettlMethod *string `fix:"1317"`
	//DerivativePriceQuoteMethod is a non-required field for DerivativeSecurityDefinition.
	DerivativePriceQuoteMethod *string `fix:"1318"`
	//DerivativeFuturesValuationMethod is a non-required field for DerivativeSecurityDefinition.
	DerivativeFuturesValuationMethod *string `fix:"1319"`
	//DerivativeListMethod is a non-required field for DerivativeSecurityDefinition.
	DerivativeListMethod *int `fix:"1320"`
	//DerivativeCapPrice is a non-required field for DerivativeSecurityDefinition.
	DerivativeCapPrice *float64 `fix:"1321"`
	//DerivativeFloorPrice is a non-required field for DerivativeSecurityDefinition.
	DerivativeFloorPrice *float64 `fix:"1322"`
	//DerivativePutOrCall is a non-required field for DerivativeSecurityDefinition.
	DerivativePutOrCall *int `fix:"1323"`
	//DerivativeSecurityXMLLen is a non-required field for DerivativeSecurityDefinition.
	DerivativeSecurityXMLLen *int `fix:"1282"`
	//DerivativeSecurityXML is a non-required field for DerivativeSecurityDefinition.
	DerivativeSecurityXML *string `fix:"1283"`
	//DerivativeSecurityXMLSchema is a non-required field for DerivativeSecurityDefinition.
	DerivativeSecurityXMLSchema *string `fix:"1284"`
	//NoDerivativeInstrAttrib is a non-required field for DerivativeSecurityDefinition.
	NoDerivativeInstrAttrib []NoDerivativeInstrAttrib `fix:"1311,omitempty"`
	//NoMarketSegments is a non-required field for DerivativeSecurityDefinition.
	NoMarketSegments []NoMarketSegments `fix:"1310,omitempty"`
}

func (m *DerivativeSecurityDefinition) SetDerivativeSymbol(v string)     { m.DerivativeSymbol = &v }
func (m *DerivativeSecurityDefinition) SetDerivativeSymbolSfx(v string)  { m.DerivativeSymbolSfx = &v }
func (m *DerivativeSecurityDefinition) SetDerivativeSecurityID(v string) { m.DerivativeSecurityID = &v }
func (m *DerivativeSecurityDefinition) SetDerivativeSecurityIDSource(v string) {
	m.DerivativeSecurityIDSource = &v
}
func (m *DerivativeSecurityDefinition) SetNoDerivativeSecurityAltID(v []NoDerivativeSecurityAltID) {
	m.NoDerivativeSecurityAltID = v
}
func (m *DerivativeSecurityDefinition) SetDerivativeProduct(v int) { m.DerivativeProduct = &v }
func (m *DerivativeSecurityDefinition) SetDerivativeProductComplex(v string) {
	m.DerivativeProductComplex = &v
}
func (m *DerivativeSecurityDefinition) SetDerivFlexProductEligibilityIndicator(v bool) {
	m.DerivFlexProductEligibilityIndicator = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeSecurityGroup(v string) {
	m.DerivativeSecurityGroup = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeCFICode(v string) { m.DerivativeCFICode = &v }
func (m *DerivativeSecurityDefinition) SetDerivativeSecurityType(v string) {
	m.DerivativeSecurityType = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeSecuritySubType(v string) {
	m.DerivativeSecuritySubType = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeMaturityMonthYear(v string) {
	m.DerivativeMaturityMonthYear = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeMaturityDate(v string) {
	m.DerivativeMaturityDate = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeMaturityTime(v string) {
	m.DerivativeMaturityTime = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeSettleOnOpenFlag(v string) {
	m.DerivativeSettleOnOpenFlag = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeInstrmtAssignmentMethod(v string) {
	m.DerivativeInstrmtAssignmentMethod = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeSecurityStatus(v string) {
	m.DerivativeSecurityStatus = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeIssueDate(v string) { m.DerivativeIssueDate = &v }
func (m *DerivativeSecurityDefinition) SetDerivativeInstrRegistry(v string) {
	m.DerivativeInstrRegistry = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeCountryOfIssue(v string) {
	m.DerivativeCountryOfIssue = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeStateOrProvinceOfIssue(v string) {
	m.DerivativeStateOrProvinceOfIssue = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeStrikePrice(v float64) {
	m.DerivativeStrikePrice = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeLocaleOfIssue(v string) {
	m.DerivativeLocaleOfIssue = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeStrikeCurrency(v string) {
	m.DerivativeStrikeCurrency = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeStrikeMultiplier(v float64) {
	m.DerivativeStrikeMultiplier = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeStrikeValue(v float64) {
	m.DerivativeStrikeValue = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeOptAttribute(v string) {
	m.DerivativeOptAttribute = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeContractMultiplier(v float64) {
	m.DerivativeContractMultiplier = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeMinPriceIncrement(v float64) {
	m.DerivativeMinPriceIncrement = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeMinPriceIncrementAmount(v float64) {
	m.DerivativeMinPriceIncrementAmount = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeUnitOfMeasure(v string) {
	m.DerivativeUnitOfMeasure = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeUnitOfMeasureQty(v float64) {
	m.DerivativeUnitOfMeasureQty = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativePriceUnitOfMeasure(v string) {
	m.DerivativePriceUnitOfMeasure = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativePriceUnitOfMeasureQty(v float64) {
	m.DerivativePriceUnitOfMeasureQty = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeExerciseStyle(v string) {
	m.DerivativeExerciseStyle = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeOptPayAmount(v float64) {
	m.DerivativeOptPayAmount = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeTimeUnit(v string) { m.DerivativeTimeUnit = &v }
func (m *DerivativeSecurityDefinition) SetDerivativeSecurityExchange(v string) {
	m.DerivativeSecurityExchange = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativePositionLimit(v int) {
	m.DerivativePositionLimit = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeNTPositionLimit(v int) {
	m.DerivativeNTPositionLimit = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeIssuer(v string) { m.DerivativeIssuer = &v }
func (m *DerivativeSecurityDefinition) SetDerivativeEncodedIssuerLen(v int) {
	m.DerivativeEncodedIssuerLen = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeEncodedIssuer(v string) {
	m.DerivativeEncodedIssuer = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeSecurityDesc(v string) {
	m.DerivativeSecurityDesc = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeEncodedSecurityDescLen(v int) {
	m.DerivativeEncodedSecurityDescLen = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeEncodedSecurityDesc(v string) {
	m.DerivativeEncodedSecurityDesc = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeContractSettlMonth(v string) {
	m.DerivativeContractSettlMonth = &v
}
func (m *DerivativeSecurityDefinition) SetNoDerivativeEvents(v []NoDerivativeEvents) {
	m.NoDerivativeEvents = v
}
func (m *DerivativeSecurityDefinition) SetNoDerivativeInstrumentParties(v []NoDerivativeInstrumentParties) {
	m.NoDerivativeInstrumentParties = v
}
func (m *DerivativeSecurityDefinition) SetDerivativeSettlMethod(v string) {
	m.DerivativeSettlMethod = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativePriceQuoteMethod(v string) {
	m.DerivativePriceQuoteMethod = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeFuturesValuationMethod(v string) {
	m.DerivativeFuturesValuationMethod = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeListMethod(v int)     { m.DerivativeListMethod = &v }
func (m *DerivativeSecurityDefinition) SetDerivativeCapPrice(v float64)   { m.DerivativeCapPrice = &v }
func (m *DerivativeSecurityDefinition) SetDerivativeFloorPrice(v float64) { m.DerivativeFloorPrice = &v }
func (m *DerivativeSecurityDefinition) SetDerivativePutOrCall(v int)      { m.DerivativePutOrCall = &v }
func (m *DerivativeSecurityDefinition) SetDerivativeSecurityXMLLen(v int) {
	m.DerivativeSecurityXMLLen = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeSecurityXML(v string) {
	m.DerivativeSecurityXML = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeSecurityXMLSchema(v string) {
	m.DerivativeSecurityXMLSchema = &v
}
func (m *DerivativeSecurityDefinition) SetNoDerivativeInstrAttrib(v []NoDerivativeInstrAttrib) {
	m.NoDerivativeInstrAttrib = v
}
func (m *DerivativeSecurityDefinition) SetNoMarketSegments(v []NoMarketSegments) {
	m.NoMarketSegments = v
}
