package underlyinginstrument

//NoUnderlyingSecurityAltID is a repeating group in UnderlyingInstrument
type NoUnderlyingSecurityAltID struct {
	//UnderlyingSecurityAltID is a non-required field for NoUnderlyingSecurityAltID.
	UnderlyingSecurityAltID *string `fix:"458"`
	//UnderlyingSecurityAltIDSource is a non-required field for NoUnderlyingSecurityAltID.
	UnderlyingSecurityAltIDSource *string `fix:"459"`
}

//NoUnderlyingStips is a repeating group in UnderlyingInstrument
type NoUnderlyingStips struct {
	//UnderlyingStipType is a non-required field for NoUnderlyingStips.
	UnderlyingStipType *string `fix:"888"`
	//UnderlyingStipValue is a non-required field for NoUnderlyingStips.
	UnderlyingStipValue *string `fix:"889"`
}

//UnderlyingInstrument is a fix44 Component
type UnderlyingInstrument struct {
	//UnderlyingSymbol is a non-required field for UnderlyingInstrument.
	UnderlyingSymbol *string `fix:"311"`
	//UnderlyingSymbolSfx is a non-required field for UnderlyingInstrument.
	UnderlyingSymbolSfx *string `fix:"312"`
	//UnderlyingSecurityID is a non-required field for UnderlyingInstrument.
	UnderlyingSecurityID *string `fix:"309"`
	//UnderlyingSecurityIDSource is a non-required field for UnderlyingInstrument.
	UnderlyingSecurityIDSource *string `fix:"305"`
	//NoUnderlyingSecurityAltID is a non-required field for UnderlyingInstrument.
	NoUnderlyingSecurityAltID []NoUnderlyingSecurityAltID `fix:"457,omitempty"`
	//UnderlyingProduct is a non-required field for UnderlyingInstrument.
	UnderlyingProduct *int `fix:"462"`
	//UnderlyingCFICode is a non-required field for UnderlyingInstrument.
	UnderlyingCFICode *string `fix:"463"`
	//UnderlyingSecurityType is a non-required field for UnderlyingInstrument.
	UnderlyingSecurityType *string `fix:"310"`
	//UnderlyingSecuritySubType is a non-required field for UnderlyingInstrument.
	UnderlyingSecuritySubType *string `fix:"763"`
	//UnderlyingMaturityMonthYear is a non-required field for UnderlyingInstrument.
	UnderlyingMaturityMonthYear *string `fix:"313"`
	//UnderlyingMaturityDate is a non-required field for UnderlyingInstrument.
	UnderlyingMaturityDate *string `fix:"542"`
	//UnderlyingCouponPaymentDate is a non-required field for UnderlyingInstrument.
	UnderlyingCouponPaymentDate *string `fix:"241"`
	//UnderlyingIssueDate is a non-required field for UnderlyingInstrument.
	UnderlyingIssueDate *string `fix:"242"`
	//UnderlyingRepoCollateralSecurityType is a non-required field for UnderlyingInstrument.
	UnderlyingRepoCollateralSecurityType *int `fix:"243"`
	//UnderlyingRepurchaseTerm is a non-required field for UnderlyingInstrument.
	UnderlyingRepurchaseTerm *int `fix:"244"`
	//UnderlyingRepurchaseRate is a non-required field for UnderlyingInstrument.
	UnderlyingRepurchaseRate *float64 `fix:"245"`
	//UnderlyingFactor is a non-required field for UnderlyingInstrument.
	UnderlyingFactor *float64 `fix:"246"`
	//UnderlyingCreditRating is a non-required field for UnderlyingInstrument.
	UnderlyingCreditRating *string `fix:"256"`
	//UnderlyingInstrRegistry is a non-required field for UnderlyingInstrument.
	UnderlyingInstrRegistry *string `fix:"595"`
	//UnderlyingCountryOfIssue is a non-required field for UnderlyingInstrument.
	UnderlyingCountryOfIssue *string `fix:"592"`
	//UnderlyingStateOrProvinceOfIssue is a non-required field for UnderlyingInstrument.
	UnderlyingStateOrProvinceOfIssue *string `fix:"593"`
	//UnderlyingLocaleOfIssue is a non-required field for UnderlyingInstrument.
	UnderlyingLocaleOfIssue *string `fix:"594"`
	//UnderlyingRedemptionDate is a non-required field for UnderlyingInstrument.
	UnderlyingRedemptionDate *string `fix:"247"`
	//UnderlyingStrikePrice is a non-required field for UnderlyingInstrument.
	UnderlyingStrikePrice *float64 `fix:"316"`
	//UnderlyingStrikeCurrency is a non-required field for UnderlyingInstrument.
	UnderlyingStrikeCurrency *string `fix:"941"`
	//UnderlyingOptAttribute is a non-required field for UnderlyingInstrument.
	UnderlyingOptAttribute *string `fix:"317"`
	//UnderlyingContractMultiplier is a non-required field for UnderlyingInstrument.
	UnderlyingContractMultiplier *float64 `fix:"436"`
	//UnderlyingCouponRate is a non-required field for UnderlyingInstrument.
	UnderlyingCouponRate *float64 `fix:"435"`
	//UnderlyingSecurityExchange is a non-required field for UnderlyingInstrument.
	UnderlyingSecurityExchange *string `fix:"308"`
	//UnderlyingIssuer is a non-required field for UnderlyingInstrument.
	UnderlyingIssuer *string `fix:"306"`
	//EncodedUnderlyingIssuerLen is a non-required field for UnderlyingInstrument.
	EncodedUnderlyingIssuerLen *int `fix:"362"`
	//EncodedUnderlyingIssuer is a non-required field for UnderlyingInstrument.
	EncodedUnderlyingIssuer *string `fix:"363"`
	//UnderlyingSecurityDesc is a non-required field for UnderlyingInstrument.
	UnderlyingSecurityDesc *string `fix:"307"`
	//EncodedUnderlyingSecurityDescLen is a non-required field for UnderlyingInstrument.
	EncodedUnderlyingSecurityDescLen *int `fix:"364"`
	//EncodedUnderlyingSecurityDesc is a non-required field for UnderlyingInstrument.
	EncodedUnderlyingSecurityDesc *string `fix:"365"`
	//UnderlyingCPProgram is a non-required field for UnderlyingInstrument.
	UnderlyingCPProgram *string `fix:"877"`
	//UnderlyingCPRegType is a non-required field for UnderlyingInstrument.
	UnderlyingCPRegType *string `fix:"878"`
	//UnderlyingCurrency is a non-required field for UnderlyingInstrument.
	UnderlyingCurrency *string `fix:"318"`
	//UnderlyingQty is a non-required field for UnderlyingInstrument.
	UnderlyingQty *float64 `fix:"879"`
	//UnderlyingPx is a non-required field for UnderlyingInstrument.
	UnderlyingPx *float64 `fix:"810"`
	//UnderlyingDirtyPrice is a non-required field for UnderlyingInstrument.
	UnderlyingDirtyPrice *float64 `fix:"882"`
	//UnderlyingEndPrice is a non-required field for UnderlyingInstrument.
	UnderlyingEndPrice *float64 `fix:"883"`
	//UnderlyingStartValue is a non-required field for UnderlyingInstrument.
	UnderlyingStartValue *float64 `fix:"884"`
	//UnderlyingCurrentValue is a non-required field for UnderlyingInstrument.
	UnderlyingCurrentValue *float64 `fix:"885"`
	//UnderlyingEndValue is a non-required field for UnderlyingInstrument.
	UnderlyingEndValue *float64 `fix:"886"`
	//NoUnderlyingStips is a non-required field for UnderlyingInstrument.
	NoUnderlyingStips []NoUnderlyingStips `fix:"887,omitempty"`
}

func (m *UnderlyingInstrument) SetUnderlyingSymbol(v string)     { m.UnderlyingSymbol = &v }
func (m *UnderlyingInstrument) SetUnderlyingSymbolSfx(v string)  { m.UnderlyingSymbolSfx = &v }
func (m *UnderlyingInstrument) SetUnderlyingSecurityID(v string) { m.UnderlyingSecurityID = &v }
func (m *UnderlyingInstrument) SetUnderlyingSecurityIDSource(v string) {
	m.UnderlyingSecurityIDSource = &v
}
func (m *UnderlyingInstrument) SetNoUnderlyingSecurityAltID(v []NoUnderlyingSecurityAltID) {
	m.NoUnderlyingSecurityAltID = v
}
func (m *UnderlyingInstrument) SetUnderlyingProduct(v int)         { m.UnderlyingProduct = &v }
func (m *UnderlyingInstrument) SetUnderlyingCFICode(v string)      { m.UnderlyingCFICode = &v }
func (m *UnderlyingInstrument) SetUnderlyingSecurityType(v string) { m.UnderlyingSecurityType = &v }
func (m *UnderlyingInstrument) SetUnderlyingSecuritySubType(v string) {
	m.UnderlyingSecuritySubType = &v
}
func (m *UnderlyingInstrument) SetUnderlyingMaturityMonthYear(v string) {
	m.UnderlyingMaturityMonthYear = &v
}
func (m *UnderlyingInstrument) SetUnderlyingMaturityDate(v string) { m.UnderlyingMaturityDate = &v }
func (m *UnderlyingInstrument) SetUnderlyingCouponPaymentDate(v string) {
	m.UnderlyingCouponPaymentDate = &v
}
func (m *UnderlyingInstrument) SetUnderlyingIssueDate(v string) { m.UnderlyingIssueDate = &v }
func (m *UnderlyingInstrument) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.UnderlyingRepoCollateralSecurityType = &v
}
func (m *UnderlyingInstrument) SetUnderlyingRepurchaseTerm(v int)     { m.UnderlyingRepurchaseTerm = &v }
func (m *UnderlyingInstrument) SetUnderlyingRepurchaseRate(v float64) { m.UnderlyingRepurchaseRate = &v }
func (m *UnderlyingInstrument) SetUnderlyingFactor(v float64)         { m.UnderlyingFactor = &v }
func (m *UnderlyingInstrument) SetUnderlyingCreditRating(v string)    { m.UnderlyingCreditRating = &v }
func (m *UnderlyingInstrument) SetUnderlyingInstrRegistry(v string)   { m.UnderlyingInstrRegistry = &v }
func (m *UnderlyingInstrument) SetUnderlyingCountryOfIssue(v string)  { m.UnderlyingCountryOfIssue = &v }
func (m *UnderlyingInstrument) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.UnderlyingStateOrProvinceOfIssue = &v
}
func (m *UnderlyingInstrument) SetUnderlyingLocaleOfIssue(v string)  { m.UnderlyingLocaleOfIssue = &v }
func (m *UnderlyingInstrument) SetUnderlyingRedemptionDate(v string) { m.UnderlyingRedemptionDate = &v }
func (m *UnderlyingInstrument) SetUnderlyingStrikePrice(v float64)   { m.UnderlyingStrikePrice = &v }
func (m *UnderlyingInstrument) SetUnderlyingStrikeCurrency(v string) { m.UnderlyingStrikeCurrency = &v }
func (m *UnderlyingInstrument) SetUnderlyingOptAttribute(v string)   { m.UnderlyingOptAttribute = &v }
func (m *UnderlyingInstrument) SetUnderlyingContractMultiplier(v float64) {
	m.UnderlyingContractMultiplier = &v
}
func (m *UnderlyingInstrument) SetUnderlyingCouponRate(v float64) { m.UnderlyingCouponRate = &v }
func (m *UnderlyingInstrument) SetUnderlyingSecurityExchange(v string) {
	m.UnderlyingSecurityExchange = &v
}
func (m *UnderlyingInstrument) SetUnderlyingIssuer(v string)        { m.UnderlyingIssuer = &v }
func (m *UnderlyingInstrument) SetEncodedUnderlyingIssuerLen(v int) { m.EncodedUnderlyingIssuerLen = &v }
func (m *UnderlyingInstrument) SetEncodedUnderlyingIssuer(v string) { m.EncodedUnderlyingIssuer = &v }
func (m *UnderlyingInstrument) SetUnderlyingSecurityDesc(v string)  { m.UnderlyingSecurityDesc = &v }
func (m *UnderlyingInstrument) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.EncodedUnderlyingSecurityDescLen = &v
}
func (m *UnderlyingInstrument) SetEncodedUnderlyingSecurityDesc(v string) {
	m.EncodedUnderlyingSecurityDesc = &v
}
func (m *UnderlyingInstrument) SetUnderlyingCPProgram(v string)            { m.UnderlyingCPProgram = &v }
func (m *UnderlyingInstrument) SetUnderlyingCPRegType(v string)            { m.UnderlyingCPRegType = &v }
func (m *UnderlyingInstrument) SetUnderlyingCurrency(v string)             { m.UnderlyingCurrency = &v }
func (m *UnderlyingInstrument) SetUnderlyingQty(v float64)                 { m.UnderlyingQty = &v }
func (m *UnderlyingInstrument) SetUnderlyingPx(v float64)                  { m.UnderlyingPx = &v }
func (m *UnderlyingInstrument) SetUnderlyingDirtyPrice(v float64)          { m.UnderlyingDirtyPrice = &v }
func (m *UnderlyingInstrument) SetUnderlyingEndPrice(v float64)            { m.UnderlyingEndPrice = &v }
func (m *UnderlyingInstrument) SetUnderlyingStartValue(v float64)          { m.UnderlyingStartValue = &v }
func (m *UnderlyingInstrument) SetUnderlyingCurrentValue(v float64)        { m.UnderlyingCurrentValue = &v }
func (m *UnderlyingInstrument) SetUnderlyingEndValue(v float64)            { m.UnderlyingEndValue = &v }
func (m *UnderlyingInstrument) SetNoUnderlyingStips(v []NoUnderlyingStips) { m.NoUnderlyingStips = v }
