package instrumentleg

//NoLegSecurityAltID is a repeating group in InstrumentLeg
type NoLegSecurityAltID struct {
	//LegSecurityAltID is a non-required field for NoLegSecurityAltID.
	LegSecurityAltID *string `fix:"605"`
	//LegSecurityAltIDSource is a non-required field for NoLegSecurityAltID.
	LegSecurityAltIDSource *string `fix:"606"`
}

//Component is a fix43 InstrumentLeg Component
type Component struct {
	//LegSymbol is a non-required field for InstrumentLeg.
	LegSymbol *string `fix:"600"`
	//LegSymbolSfx is a non-required field for InstrumentLeg.
	LegSymbolSfx *string `fix:"601"`
	//LegSecurityID is a non-required field for InstrumentLeg.
	LegSecurityID *string `fix:"602"`
	//LegSecurityIDSource is a non-required field for InstrumentLeg.
	LegSecurityIDSource *string `fix:"603"`
	//NoLegSecurityAltID is a non-required field for InstrumentLeg.
	NoLegSecurityAltID []NoLegSecurityAltID `fix:"604,omitempty"`
	//LegProduct is a non-required field for InstrumentLeg.
	LegProduct *int `fix:"607"`
	//LegCFICode is a non-required field for InstrumentLeg.
	LegCFICode *string `fix:"608"`
	//LegSecurityType is a non-required field for InstrumentLeg.
	LegSecurityType *string `fix:"609"`
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
}

func New() *Component { return new(Component) }
