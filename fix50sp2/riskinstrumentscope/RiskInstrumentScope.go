package riskinstrumentscope

import (
	"github.com/quickfixgo/quickfix/fix50sp2/risksecaltidgrp"
)

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
	//RiskSecAltIDGrp Component
	RiskSecAltIDGrp risksecaltidgrp.Component
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

//Component is a fix50sp2 RiskInstrumentScope Component
type Component struct {
	//NoRiskInstruments is a non-required field for RiskInstrumentScope.
	NoRiskInstruments []NoRiskInstruments `fix:"1534,omitempty"`
}

func New() *Component { return new(Component) }
