package riskinstrumentscope

import (
	"github.com/quickfixgo/quickfix/fix50sp2/risksecaltidgrp"
)

func New() *RiskInstrumentScope {
	var m RiskInstrumentScope
	return &m
}

//NoRiskInstruments is a repeating group in RiskInstrumentScope
type NoRiskInstruments struct {
	//RiskInstrumentOperator is a non-required field for NoRiskInstruments.
	RiskInstrumentOperator *int `fix:"1535"`
	//RiskSymbol is a non-required field for NoRiskInstruments.
	RiskSymbol *string `fix:"1536"`
	//RiskSymbolSfx is a non-required field for NoRiskInstruments.
	RiskSymbolSfx *string `fix:"1537"`
	//RiskSecurityID is a non-required field for NoRiskInstruments.
	RiskSecurityID *string `fix:"1538"`
	//RiskSecurityIDSource is a non-required field for NoRiskInstruments.
	RiskSecurityIDSource *string `fix:"1539"`
	//RiskSecAltIDGrp is a non-required component for NoRiskInstruments.
	RiskSecAltIDGrp *risksecaltidgrp.RiskSecAltIDGrp
	//RiskProduct is a non-required field for NoRiskInstruments.
	RiskProduct *int `fix:"1543"`
	//RiskProductComplex is a non-required field for NoRiskInstruments.
	RiskProductComplex *string `fix:"1544"`
	//RiskSecurityGroup is a non-required field for NoRiskInstruments.
	RiskSecurityGroup *string `fix:"1545"`
	//RiskCFICode is a non-required field for NoRiskInstruments.
	RiskCFICode *string `fix:"1546"`
	//RiskSecurityType is a non-required field for NoRiskInstruments.
	RiskSecurityType *string `fix:"1547"`
	//RiskSecuritySubType is a non-required field for NoRiskInstruments.
	RiskSecuritySubType *string `fix:"1548"`
	//RiskMaturityMonthYear is a non-required field for NoRiskInstruments.
	RiskMaturityMonthYear *string `fix:"1549"`
	//RiskMaturityTime is a non-required field for NoRiskInstruments.
	RiskMaturityTime *string `fix:"1550"`
	//RiskRestructuringType is a non-required field for NoRiskInstruments.
	RiskRestructuringType *string `fix:"1551"`
	//RiskSeniority is a non-required field for NoRiskInstruments.
	RiskSeniority *string `fix:"1552"`
	//RiskPutOrCall is a non-required field for NoRiskInstruments.
	RiskPutOrCall *int `fix:"1553"`
	//RiskFlexibleIndicator is a non-required field for NoRiskInstruments.
	RiskFlexibleIndicator *bool `fix:"1554"`
	//RiskCouponRate is a non-required field for NoRiskInstruments.
	RiskCouponRate *float64 `fix:"1555"`
	//RiskSecurityExchange is a non-required field for NoRiskInstruments.
	RiskSecurityExchange *string `fix:"1616"`
	//RiskSecurityDesc is a non-required field for NoRiskInstruments.
	RiskSecurityDesc *string `fix:"1556"`
	//RiskEncodedSecurityDescLen is a non-required field for NoRiskInstruments.
	RiskEncodedSecurityDescLen *int `fix:"1620"`
	//RiskEncodedSecurityDesc is a non-required field for NoRiskInstruments.
	RiskEncodedSecurityDesc *string `fix:"1621"`
	//RiskInstrumentSettlType is a non-required field for NoRiskInstruments.
	RiskInstrumentSettlType *string `fix:"1557"`
	//RiskInstrumentMultiplier is a non-required field for NoRiskInstruments.
	RiskInstrumentMultiplier *float64 `fix:"1558"`
}

func (m *NoRiskInstruments) SetRiskInstrumentOperator(v int)  { m.RiskInstrumentOperator = &v }
func (m *NoRiskInstruments) SetRiskSymbol(v string)           { m.RiskSymbol = &v }
func (m *NoRiskInstruments) SetRiskSymbolSfx(v string)        { m.RiskSymbolSfx = &v }
func (m *NoRiskInstruments) SetRiskSecurityID(v string)       { m.RiskSecurityID = &v }
func (m *NoRiskInstruments) SetRiskSecurityIDSource(v string) { m.RiskSecurityIDSource = &v }
func (m *NoRiskInstruments) SetRiskSecAltIDGrp(v risksecaltidgrp.RiskSecAltIDGrp) {
	m.RiskSecAltIDGrp = &v
}
func (m *NoRiskInstruments) SetRiskProduct(v int)                  { m.RiskProduct = &v }
func (m *NoRiskInstruments) SetRiskProductComplex(v string)        { m.RiskProductComplex = &v }
func (m *NoRiskInstruments) SetRiskSecurityGroup(v string)         { m.RiskSecurityGroup = &v }
func (m *NoRiskInstruments) SetRiskCFICode(v string)               { m.RiskCFICode = &v }
func (m *NoRiskInstruments) SetRiskSecurityType(v string)          { m.RiskSecurityType = &v }
func (m *NoRiskInstruments) SetRiskSecuritySubType(v string)       { m.RiskSecuritySubType = &v }
func (m *NoRiskInstruments) SetRiskMaturityMonthYear(v string)     { m.RiskMaturityMonthYear = &v }
func (m *NoRiskInstruments) SetRiskMaturityTime(v string)          { m.RiskMaturityTime = &v }
func (m *NoRiskInstruments) SetRiskRestructuringType(v string)     { m.RiskRestructuringType = &v }
func (m *NoRiskInstruments) SetRiskSeniority(v string)             { m.RiskSeniority = &v }
func (m *NoRiskInstruments) SetRiskPutOrCall(v int)                { m.RiskPutOrCall = &v }
func (m *NoRiskInstruments) SetRiskFlexibleIndicator(v bool)       { m.RiskFlexibleIndicator = &v }
func (m *NoRiskInstruments) SetRiskCouponRate(v float64)           { m.RiskCouponRate = &v }
func (m *NoRiskInstruments) SetRiskSecurityExchange(v string)      { m.RiskSecurityExchange = &v }
func (m *NoRiskInstruments) SetRiskSecurityDesc(v string)          { m.RiskSecurityDesc = &v }
func (m *NoRiskInstruments) SetRiskEncodedSecurityDescLen(v int)   { m.RiskEncodedSecurityDescLen = &v }
func (m *NoRiskInstruments) SetRiskEncodedSecurityDesc(v string)   { m.RiskEncodedSecurityDesc = &v }
func (m *NoRiskInstruments) SetRiskInstrumentSettlType(v string)   { m.RiskInstrumentSettlType = &v }
func (m *NoRiskInstruments) SetRiskInstrumentMultiplier(v float64) { m.RiskInstrumentMultiplier = &v }

//RiskInstrumentScope is a fix50sp2 Component
type RiskInstrumentScope struct {
	//NoRiskInstruments is a non-required field for RiskInstrumentScope.
	NoRiskInstruments []NoRiskInstruments `fix:"1534,omitempty"`
}

func (m *RiskInstrumentScope) SetNoRiskInstruments(v []NoRiskInstruments) { m.NoRiskInstruments = v }
