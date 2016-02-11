package underlyingleginstrument

//NoUnderlyingLegSecurityAltID is a repeating group in UnderlyingLegInstrument
type NoUnderlyingLegSecurityAltID struct {
	//UnderlyingLegSecurityAltID is a non-required field for NoUnderlyingLegSecurityAltID.
	UnderlyingLegSecurityAltID *string `fix:"1335"`
	//UnderlyingLegSecurityAltIDSource is a non-required field for NoUnderlyingLegSecurityAltID.
	UnderlyingLegSecurityAltIDSource *string `fix:"1336"`
}

//Component is a fix50sp2 UnderlyingLegInstrument Component
type Component struct {
	//UnderlyingLegSymbol is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegSymbol *string `fix:"1330"`
	//UnderlyingLegSymbolSfx is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegSymbolSfx *string `fix:"1331"`
	//UnderlyingLegSecurityID is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegSecurityID *string `fix:"1332"`
	//UnderlyingLegSecurityIDSource is a non-required field for UnderlyingLegInstrument.
	UnderlyingLegSecurityIDSource *string `fix:"1333"`
	//NoUnderlyingLegSecurityAltID is a non-required field for UnderlyingLegInstrument.
	NoUnderlyingLegSecurityAltID []NoUnderlyingLegSecurityAltID `fix:"1334,omitempty"`
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

func New() *Component { return new(Component) }
