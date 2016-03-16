package underlyingleginstrument

import (
	"github.com/quickfixgo/quickfix/fix50sp1/underlyinglegsecurityaltidgrp"
)

//UnderlyingLegInstrument is a fix50sp1 Component
type UnderlyingLegInstrument struct {
	//UnderlyingLegSymbol is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegSymbol *string `fix:"1330"`
	//UnderlyingLegSymbolSfx is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegSymbolSfx *string `fix:"1331"`
	//UnderlyingLegSecurityID is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegSecurityID *string `fix:"1332"`
	//UnderlyingLegSecurityIDSource is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegSecurityIDSource *string `fix:"1333"`
	//UnderlyingLegSecurityAltIDGrp is a non-required component for UnderlyingLegInstrument.
	UnderlyingLegSecurityAltIDGrp *underlyinglegsecurityaltidgrp.UnderlyingLegSecurityAltIDGrp
	//UnderlyingLegCFICode is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegCFICode *string `fix:"1344"`
	//UnderlyingLegSecurityType is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegSecurityType *string `fix:"1337"`
	//UnderlyingLegSecuritySubType is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegSecuritySubType *string `fix:"1338"`
	//UnderlyingLegMaturityMonthYear is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegMaturityMonthYear *string `fix:"1339"`
	//UnderlyingLegMaturityDate is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegMaturityDate *string `fix:"1345"`
	//UnderlyingLegMaturityTime is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegMaturityTime *string `fix:"1405"`
	//UnderlyingLegStrikePrice is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegStrikePrice *float64 `fix:"1340"`
	//UnderlyingLegOptAttribute is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegOptAttribute *string `fix:"1391"`
	//UnderlyingLegPutOrCall is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegPutOrCall *int `fix:"1343"`
	//UnderlyingLegSecurityExchange is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegSecurityExchange *string `fix:"1341"`
	//UnderlyingLegSecurityDesc is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegSecurityDesc *string `fix:"1392"`
}

func (m *UnderlyingLegInstrument) SetUnderlyingLegSymbol(v string)     { m.UnderlyingLegSymbol = &v }
func (m *UnderlyingLegInstrument) SetUnderlyingLegSymbolSfx(v string)  { m.UnderlyingLegSymbolSfx = &v }
func (m *UnderlyingLegInstrument) SetUnderlyingLegSecurityID(v string) { m.UnderlyingLegSecurityID = &v }
func (m *UnderlyingLegInstrument) SetUnderlyingLegSecurityIDSource(v string) {
	m.UnderlyingLegSecurityIDSource = &v
}
func (m *UnderlyingLegInstrument) SetUnderlyingLegSecurityAltIDGrp(v underlyinglegsecurityaltidgrp.UnderlyingLegSecurityAltIDGrp) {
	m.UnderlyingLegSecurityAltIDGrp = &v
}
func (m *UnderlyingLegInstrument) SetUnderlyingLegCFICode(v string) { m.UnderlyingLegCFICode = &v }
func (m *UnderlyingLegInstrument) SetUnderlyingLegSecurityType(v string) {
	m.UnderlyingLegSecurityType = &v
}
func (m *UnderlyingLegInstrument) SetUnderlyingLegSecuritySubType(v string) {
	m.UnderlyingLegSecuritySubType = &v
}
func (m *UnderlyingLegInstrument) SetUnderlyingLegMaturityMonthYear(v string) {
	m.UnderlyingLegMaturityMonthYear = &v
}
func (m *UnderlyingLegInstrument) SetUnderlyingLegMaturityDate(v string) {
	m.UnderlyingLegMaturityDate = &v
}
func (m *UnderlyingLegInstrument) SetUnderlyingLegMaturityTime(v string) {
	m.UnderlyingLegMaturityTime = &v
}
func (m *UnderlyingLegInstrument) SetUnderlyingLegStrikePrice(v float64) {
	m.UnderlyingLegStrikePrice = &v
}
func (m *UnderlyingLegInstrument) SetUnderlyingLegOptAttribute(v string) {
	m.UnderlyingLegOptAttribute = &v
}
func (m *UnderlyingLegInstrument) SetUnderlyingLegPutOrCall(v int) { m.UnderlyingLegPutOrCall = &v }
func (m *UnderlyingLegInstrument) SetUnderlyingLegSecurityExchange(v string) {
	m.UnderlyingLegSecurityExchange = &v
}
func (m *UnderlyingLegInstrument) SetUnderlyingLegSecurityDesc(v string) {
	m.UnderlyingLegSecurityDesc = &v
}
