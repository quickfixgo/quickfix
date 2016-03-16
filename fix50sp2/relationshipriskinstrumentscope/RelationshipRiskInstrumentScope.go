package relationshipriskinstrumentscope

import (
	"github.com/quickfixgo/quickfix/fix50sp2/relationshiprisksecaltidgrp"
)

//NoRelationshipRiskInstruments is a repeating group in RelationshipRiskInstrumentScope
type NoRelationshipRiskInstruments struct {
	//RelationshipRiskInstrumentOperator is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskInstrumentOperator *int `fix:"1588"`
	//RelationshipRiskSymbol is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskSymbol *string `fix:"1589"`
	//RelationshipRiskSymbolSfx is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskSymbolSfx *string `fix:"1590"`
	//RelationshipRiskSecurityID is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskSecurityID *string `fix:"1591"`
	//RelationshipRiskSecurityIDSource is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskSecurityIDSource *string `fix:"1592"`
	//RelationshipRiskSecAltIDGrp is a non-required component for NoRelationshipRiskInstruments.
	RelationshipRiskSecAltIDGrp *relationshiprisksecaltidgrp.RelationshipRiskSecAltIDGrp
	//RelationshipRiskProduct is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskProduct *int `fix:"1596"`
	//RelationshipRiskProductComplex is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskProductComplex *string `fix:"1597"`
	//RelationshipRiskSecurityGroup is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskSecurityGroup *string `fix:"1598"`
	//RelationshipRiskCFICode is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskCFICode *string `fix:"1599"`
	//RelationshipRiskSecurityType is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskSecurityType *string `fix:"1600"`
	//RelationshipRiskSecuritySubType is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskSecuritySubType *string `fix:"1601"`
	//RelationshipRiskMaturityMonthYear is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskMaturityMonthYear *string `fix:"1602"`
	//RelationshipRiskMaturityTime is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskMaturityTime *string `fix:"1603"`
	//RelationshipRiskRestructuringType is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskRestructuringType *string `fix:"1604"`
	//RelationshipRiskSeniority is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskSeniority *string `fix:"1605"`
	//RelationshipRiskPutOrCall is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskPutOrCall *int `fix:"1606"`
	//RelationshipRiskFlexibleIndicator is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskFlexibleIndicator *bool `fix:"1607"`
	//RelationshipRiskCouponRate is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskCouponRate *float64 `fix:"1608"`
	//RelationshipRiskSecurityExchange is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskSecurityExchange *string `fix:"1609"`
	//RelationshipRiskSecurityDesc is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskSecurityDesc *string `fix:"1610"`
	//RelationshipRiskEncodedSecurityDescLen is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskEncodedSecurityDescLen *int `fix:"1618"`
	//RelationshipRiskEncodedSecurityDesc is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskEncodedSecurityDesc *string `fix:"1619"`
	//RelationshipRiskInstrumentSettlType is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskInstrumentSettlType *string `fix:"1611"`
	//RelationshipRiskInstrumentMultiplier is a non-required field for NoRelationshipRiskInstruments.
	RelationshipRiskInstrumentMultiplier *float64 `fix:"1612"`
}

//RelationshipRiskInstrumentScope is a fix50sp2 Component
type RelationshipRiskInstrumentScope struct {
	//NoRelationshipRiskInstruments is a non-required field for RelationshipRiskInstrumentScope.
	NoRelationshipRiskInstruments []NoRelationshipRiskInstruments `fix:"1587,omitempty"`
}

func (m *RelationshipRiskInstrumentScope) SetNoRelationshipRiskInstruments(v []NoRelationshipRiskInstruments) {
	m.NoRelationshipRiskInstruments = v
}
