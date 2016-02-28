package instrumentleg

//NoLegSecurityAltID is a repeating group in InstrumentLeg
type NoLegSecurityAltID struct {
	//LegSecurityAltID is a non-required field for NoLegSecurityAltID.
	LegSecurityAltID *string `fix:"605"`
	//LegSecurityAltIDSource is a non-required field for NoLegSecurityAltID.
	LegSecurityAltIDSource *string `fix:"606"`
}

//Component is a fix50sp1 InstrumentLeg Component
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
}

func New() *Component { return new(Component) }

func (m *Component) SetLegSymbol(v string)                        { m.LegSymbol = &v }
func (m *Component) SetLegSymbolSfx(v string)                     { m.LegSymbolSfx = &v }
func (m *Component) SetLegSecurityID(v string)                    { m.LegSecurityID = &v }
func (m *Component) SetLegSecurityIDSource(v string)              { m.LegSecurityIDSource = &v }
func (m *Component) SetNoLegSecurityAltID(v []NoLegSecurityAltID) { m.NoLegSecurityAltID = v }
func (m *Component) SetLegProduct(v int)                          { m.LegProduct = &v }
func (m *Component) SetLegCFICode(v string)                       { m.LegCFICode = &v }
func (m *Component) SetLegSecurityType(v string)                  { m.LegSecurityType = &v }
func (m *Component) SetLegSecuritySubType(v string)               { m.LegSecuritySubType = &v }
func (m *Component) SetLegMaturityMonthYear(v string)             { m.LegMaturityMonthYear = &v }
func (m *Component) SetLegMaturityDate(v string)                  { m.LegMaturityDate = &v }
func (m *Component) SetLegCouponPaymentDate(v string)             { m.LegCouponPaymentDate = &v }
func (m *Component) SetLegIssueDate(v string)                     { m.LegIssueDate = &v }
func (m *Component) SetLegRepoCollateralSecurityType(v int)       { m.LegRepoCollateralSecurityType = &v }
func (m *Component) SetLegRepurchaseTerm(v int)                   { m.LegRepurchaseTerm = &v }
func (m *Component) SetLegRepurchaseRate(v float64)               { m.LegRepurchaseRate = &v }
func (m *Component) SetLegFactor(v float64)                       { m.LegFactor = &v }
func (m *Component) SetLegCreditRating(v string)                  { m.LegCreditRating = &v }
func (m *Component) SetLegInstrRegistry(v string)                 { m.LegInstrRegistry = &v }
func (m *Component) SetLegCountryOfIssue(v string)                { m.LegCountryOfIssue = &v }
func (m *Component) SetLegStateOrProvinceOfIssue(v string)        { m.LegStateOrProvinceOfIssue = &v }
func (m *Component) SetLegLocaleOfIssue(v string)                 { m.LegLocaleOfIssue = &v }
func (m *Component) SetLegRedemptionDate(v string)                { m.LegRedemptionDate = &v }
func (m *Component) SetLegStrikePrice(v float64)                  { m.LegStrikePrice = &v }
func (m *Component) SetLegStrikeCurrency(v string)                { m.LegStrikeCurrency = &v }
func (m *Component) SetLegOptAttribute(v string)                  { m.LegOptAttribute = &v }
func (m *Component) SetLegContractMultiplier(v float64)           { m.LegContractMultiplier = &v }
func (m *Component) SetLegCouponRate(v float64)                   { m.LegCouponRate = &v }
func (m *Component) SetLegSecurityExchange(v string)              { m.LegSecurityExchange = &v }
func (m *Component) SetLegIssuer(v string)                        { m.LegIssuer = &v }
func (m *Component) SetEncodedLegIssuerLen(v int)                 { m.EncodedLegIssuerLen = &v }
func (m *Component) SetEncodedLegIssuer(v string)                 { m.EncodedLegIssuer = &v }
func (m *Component) SetLegSecurityDesc(v string)                  { m.LegSecurityDesc = &v }
func (m *Component) SetEncodedLegSecurityDescLen(v int)           { m.EncodedLegSecurityDescLen = &v }
func (m *Component) SetEncodedLegSecurityDesc(v string)           { m.EncodedLegSecurityDesc = &v }
func (m *Component) SetLegRatioQty(v float64)                     { m.LegRatioQty = &v }
func (m *Component) SetLegSide(v string)                          { m.LegSide = &v }
func (m *Component) SetLegCurrency(v string)                      { m.LegCurrency = &v }
func (m *Component) SetLegPool(v string)                          { m.LegPool = &v }
func (m *Component) SetLegDatedDate(v string)                     { m.LegDatedDate = &v }
func (m *Component) SetLegContractSettlMonth(v string)            { m.LegContractSettlMonth = &v }
func (m *Component) SetLegInterestAccrualDate(v string)           { m.LegInterestAccrualDate = &v }
func (m *Component) SetLegUnitOfMeasure(v string)                 { m.LegUnitOfMeasure = &v }
func (m *Component) SetLegTimeUnit(v string)                      { m.LegTimeUnit = &v }
func (m *Component) SetLegOptionRatio(v float64)                  { m.LegOptionRatio = &v }
func (m *Component) SetLegPrice(v float64)                        { m.LegPrice = &v }
func (m *Component) SetLegMaturityTime(v string)                  { m.LegMaturityTime = &v }
func (m *Component) SetLegPutOrCall(v int)                        { m.LegPutOrCall = &v }
func (m *Component) SetLegExerciseStyle(v int)                    { m.LegExerciseStyle = &v }
func (m *Component) SetLegUnitOfMeasureQty(v float64)             { m.LegUnitOfMeasureQty = &v }
func (m *Component) SetLegPriceUnitOfMeasure(v string)            { m.LegPriceUnitOfMeasure = &v }
func (m *Component) SetLegPriceUnitOfMeasureQty(v float64)        { m.LegPriceUnitOfMeasureQty = &v }
