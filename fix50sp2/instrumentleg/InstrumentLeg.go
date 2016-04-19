package instrumentleg

import (
	"github.com/quickfixgo/quickfix/fix50sp2/legsecaltidgrp"
)

//New returns an initialized InstrumentLeg instance
func New() *InstrumentLeg {
	var m InstrumentLeg
	return &m
}

//InstrumentLeg is a fix50sp2 Component
type InstrumentLeg struct {
	//LegSymbol is a non-required field for InstrumentLeg.
	LegSymbol *string `fix:"600"`
	//LegSymbolSfx is a non-required field for InstrumentLeg.
	LegSymbolSfx *string `fix:"601"`
	//LegSecurityID is a non-required field for InstrumentLeg.
	LegSecurityID *string `fix:"602"`
	//LegSecurityIDSource is a non-required field for InstrumentLeg.
	LegSecurityIDSource *string `fix:"603"`
	//LegSecAltIDGrp is a non-required component for InstrumentLeg.
	LegSecAltIDGrp *legsecaltidgrp.LegSecAltIDGrp
	//LegProduct is a non-required field for InstrumentLeg.
	LegProduct *int `fix:"607"`
	//LegCFICode is a non-required field for InstrumentLeg.
	LegCFICode *string `fix:"608"`
	//LegSecurityType is a non-required field for InstrumentLeg.
	LegSecurityType *string `fix:"609"`
	//LegSecuritySubType is a non-required field for InstrumentLeg.
	LegSecuritySubType *string `fix:"764"`
	//LegMaturityMonthYear is a non-required field for InstrumentLeg.
	LegMaturityMonthYear *string `fix:"610"`
	//LegMaturityDate is a non-required field for InstrumentLeg.
	LegMaturityDate *string `fix:"611"`
	//LegCouponPaymentDate is a non-required field for InstrumentLeg.
	LegCouponPaymentDate *string `fix:"248"`
	//LegIssueDate is a non-required field for InstrumentLeg.
	LegIssueDate *string `fix:"249"`
	//LegRepoCollateralSecurityType is a non-required field for InstrumentLeg.
	LegRepoCollateralSecurityType *int `fix:"250"`
	//LegRepurchaseTerm is a non-required field for InstrumentLeg.
	LegRepurchaseTerm *int `fix:"251"`
	//LegRepurchaseRate is a non-required field for InstrumentLeg.
	LegRepurchaseRate *float64 `fix:"252"`
	//LegFactor is a non-required field for InstrumentLeg.
	LegFactor *float64 `fix:"253"`
	//LegCreditRating is a non-required field for InstrumentLeg.
	LegCreditRating *string `fix:"257"`
	//LegInstrRegistry is a non-required field for InstrumentLeg.
	LegInstrRegistry *string `fix:"599"`
	//LegCountryOfIssue is a non-required field for InstrumentLeg.
	LegCountryOfIssue *string `fix:"596"`
	//LegStateOrProvinceOfIssue is a non-required field for InstrumentLeg.
	LegStateOrProvinceOfIssue *string `fix:"597"`
	//LegLocaleOfIssue is a non-required field for InstrumentLeg.
	LegLocaleOfIssue *string `fix:"598"`
	//LegRedemptionDate is a non-required field for InstrumentLeg.
	LegRedemptionDate *string `fix:"254"`
	//LegStrikePrice is a non-required field for InstrumentLeg.
	LegStrikePrice *float64 `fix:"612"`
	//LegStrikeCurrency is a non-required field for InstrumentLeg.
	LegStrikeCurrency *string `fix:"942"`
	//LegOptAttribute is a non-required field for InstrumentLeg.
	LegOptAttribute *string `fix:"613"`
	//LegContractMultiplier is a non-required field for InstrumentLeg.
	LegContractMultiplier *float64 `fix:"614"`
	//LegCouponRate is a non-required field for InstrumentLeg.
	LegCouponRate *float64 `fix:"615"`
	//LegSecurityExchange is a non-required field for InstrumentLeg.
	LegSecurityExchange *string `fix:"616"`
	//LegIssuer is a non-required field for InstrumentLeg.
	LegIssuer *string `fix:"617"`
	//EncodedLegIssuerLen is a non-required field for InstrumentLeg.
	EncodedLegIssuerLen *int `fix:"618"`
	//EncodedLegIssuer is a non-required field for InstrumentLeg.
	EncodedLegIssuer *string `fix:"619"`
	//LegSecurityDesc is a non-required field for InstrumentLeg.
	LegSecurityDesc *string `fix:"620"`
	//EncodedLegSecurityDescLen is a non-required field for InstrumentLeg.
	EncodedLegSecurityDescLen *int `fix:"621"`
	//EncodedLegSecurityDesc is a non-required field for InstrumentLeg.
	EncodedLegSecurityDesc *string `fix:"622"`
	//LegRatioQty is a non-required field for InstrumentLeg.
	LegRatioQty *float64 `fix:"623"`
	//LegSide is a non-required field for InstrumentLeg.
	LegSide *string `fix:"624"`
	//LegCurrency is a non-required field for InstrumentLeg.
	LegCurrency *string `fix:"556"`
	//LegPool is a non-required field for InstrumentLeg.
	LegPool *string `fix:"740"`
	//LegDatedDate is a non-required field for InstrumentLeg.
	LegDatedDate *string `fix:"739"`
	//LegContractSettlMonth is a non-required field for InstrumentLeg.
	LegContractSettlMonth *string `fix:"955"`
	//LegInterestAccrualDate is a non-required field for InstrumentLeg.
	LegInterestAccrualDate *string `fix:"956"`
	//LegUnitOfMeasure is a non-required field for InstrumentLeg.
	LegUnitOfMeasure *string `fix:"999"`
	//LegTimeUnit is a non-required field for InstrumentLeg.
	LegTimeUnit *string `fix:"1001"`
	//LegOptionRatio is a non-required field for InstrumentLeg.
	LegOptionRatio *float64 `fix:"1017"`
	//LegPrice is a non-required field for InstrumentLeg.
	LegPrice *float64 `fix:"566"`
	//LegMaturityTime is a non-required field for InstrumentLeg.
	LegMaturityTime *string `fix:"1212"`
	//LegPutOrCall is a non-required field for InstrumentLeg.
	LegPutOrCall *int `fix:"1358"`
	//LegExerciseStyle is a non-required field for InstrumentLeg.
	LegExerciseStyle *int `fix:"1420"`
	//LegUnitOfMeasureQty is a non-required field for InstrumentLeg.
	LegUnitOfMeasureQty *float64 `fix:"1224"`
	//LegPriceUnitOfMeasure is a non-required field for InstrumentLeg.
	LegPriceUnitOfMeasure *string `fix:"1421"`
	//LegPriceUnitOfMeasureQty is a non-required field for InstrumentLeg.
	LegPriceUnitOfMeasureQty *float64 `fix:"1422"`
	//LegContractMultiplierUnit is a non-required field for InstrumentLeg.
	LegContractMultiplierUnit *int `fix:"1436"`
	//LegFlowScheduleType is a non-required field for InstrumentLeg.
	LegFlowScheduleType *int `fix:"1440"`
}

func (m *InstrumentLeg) SetLegSymbol(v string)                             { m.LegSymbol = &v }
func (m *InstrumentLeg) SetLegSymbolSfx(v string)                          { m.LegSymbolSfx = &v }
func (m *InstrumentLeg) SetLegSecurityID(v string)                         { m.LegSecurityID = &v }
func (m *InstrumentLeg) SetLegSecurityIDSource(v string)                   { m.LegSecurityIDSource = &v }
func (m *InstrumentLeg) SetLegSecAltIDGrp(v legsecaltidgrp.LegSecAltIDGrp) { m.LegSecAltIDGrp = &v }
func (m *InstrumentLeg) SetLegProduct(v int)                               { m.LegProduct = &v }
func (m *InstrumentLeg) SetLegCFICode(v string)                            { m.LegCFICode = &v }
func (m *InstrumentLeg) SetLegSecurityType(v string)                       { m.LegSecurityType = &v }
func (m *InstrumentLeg) SetLegSecuritySubType(v string)                    { m.LegSecuritySubType = &v }
func (m *InstrumentLeg) SetLegMaturityMonthYear(v string)                  { m.LegMaturityMonthYear = &v }
func (m *InstrumentLeg) SetLegMaturityDate(v string)                       { m.LegMaturityDate = &v }
func (m *InstrumentLeg) SetLegCouponPaymentDate(v string)                  { m.LegCouponPaymentDate = &v }
func (m *InstrumentLeg) SetLegIssueDate(v string)                          { m.LegIssueDate = &v }
func (m *InstrumentLeg) SetLegRepoCollateralSecurityType(v int)            { m.LegRepoCollateralSecurityType = &v }
func (m *InstrumentLeg) SetLegRepurchaseTerm(v int)                        { m.LegRepurchaseTerm = &v }
func (m *InstrumentLeg) SetLegRepurchaseRate(v float64)                    { m.LegRepurchaseRate = &v }
func (m *InstrumentLeg) SetLegFactor(v float64)                            { m.LegFactor = &v }
func (m *InstrumentLeg) SetLegCreditRating(v string)                       { m.LegCreditRating = &v }
func (m *InstrumentLeg) SetLegInstrRegistry(v string)                      { m.LegInstrRegistry = &v }
func (m *InstrumentLeg) SetLegCountryOfIssue(v string)                     { m.LegCountryOfIssue = &v }
func (m *InstrumentLeg) SetLegStateOrProvinceOfIssue(v string)             { m.LegStateOrProvinceOfIssue = &v }
func (m *InstrumentLeg) SetLegLocaleOfIssue(v string)                      { m.LegLocaleOfIssue = &v }
func (m *InstrumentLeg) SetLegRedemptionDate(v string)                     { m.LegRedemptionDate = &v }
func (m *InstrumentLeg) SetLegStrikePrice(v float64)                       { m.LegStrikePrice = &v }
func (m *InstrumentLeg) SetLegStrikeCurrency(v string)                     { m.LegStrikeCurrency = &v }
func (m *InstrumentLeg) SetLegOptAttribute(v string)                       { m.LegOptAttribute = &v }
func (m *InstrumentLeg) SetLegContractMultiplier(v float64)                { m.LegContractMultiplier = &v }
func (m *InstrumentLeg) SetLegCouponRate(v float64)                        { m.LegCouponRate = &v }
func (m *InstrumentLeg) SetLegSecurityExchange(v string)                   { m.LegSecurityExchange = &v }
func (m *InstrumentLeg) SetLegIssuer(v string)                             { m.LegIssuer = &v }
func (m *InstrumentLeg) SetEncodedLegIssuerLen(v int)                      { m.EncodedLegIssuerLen = &v }
func (m *InstrumentLeg) SetEncodedLegIssuer(v string)                      { m.EncodedLegIssuer = &v }
func (m *InstrumentLeg) SetLegSecurityDesc(v string)                       { m.LegSecurityDesc = &v }
func (m *InstrumentLeg) SetEncodedLegSecurityDescLen(v int)                { m.EncodedLegSecurityDescLen = &v }
func (m *InstrumentLeg) SetEncodedLegSecurityDesc(v string)                { m.EncodedLegSecurityDesc = &v }
func (m *InstrumentLeg) SetLegRatioQty(v float64)                          { m.LegRatioQty = &v }
func (m *InstrumentLeg) SetLegSide(v string)                               { m.LegSide = &v }
func (m *InstrumentLeg) SetLegCurrency(v string)                           { m.LegCurrency = &v }
func (m *InstrumentLeg) SetLegPool(v string)                               { m.LegPool = &v }
func (m *InstrumentLeg) SetLegDatedDate(v string)                          { m.LegDatedDate = &v }
func (m *InstrumentLeg) SetLegContractSettlMonth(v string)                 { m.LegContractSettlMonth = &v }
func (m *InstrumentLeg) SetLegInterestAccrualDate(v string)                { m.LegInterestAccrualDate = &v }
func (m *InstrumentLeg) SetLegUnitOfMeasure(v string)                      { m.LegUnitOfMeasure = &v }
func (m *InstrumentLeg) SetLegTimeUnit(v string)                           { m.LegTimeUnit = &v }
func (m *InstrumentLeg) SetLegOptionRatio(v float64)                       { m.LegOptionRatio = &v }
func (m *InstrumentLeg) SetLegPrice(v float64)                             { m.LegPrice = &v }
func (m *InstrumentLeg) SetLegMaturityTime(v string)                       { m.LegMaturityTime = &v }
func (m *InstrumentLeg) SetLegPutOrCall(v int)                             { m.LegPutOrCall = &v }
func (m *InstrumentLeg) SetLegExerciseStyle(v int)                         { m.LegExerciseStyle = &v }
func (m *InstrumentLeg) SetLegUnitOfMeasureQty(v float64)                  { m.LegUnitOfMeasureQty = &v }
func (m *InstrumentLeg) SetLegPriceUnitOfMeasure(v string)                 { m.LegPriceUnitOfMeasure = &v }
func (m *InstrumentLeg) SetLegPriceUnitOfMeasureQty(v float64)             { m.LegPriceUnitOfMeasureQty = &v }
func (m *InstrumentLeg) SetLegContractMultiplierUnit(v int)                { m.LegContractMultiplierUnit = &v }
func (m *InstrumentLeg) SetLegFlowScheduleType(v int)                      { m.LegFlowScheduleType = &v }
