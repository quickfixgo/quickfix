package instrument

import (
	"github.com/quickfixgo/quickfix/fix50sp2/complexeventdates"
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentptyssubgrp"
	"time"
)

//NoSecurityAltID is a repeating group in Instrument
type NoSecurityAltID struct {
	//SecurityAltID is a non-required field for NoSecurityAltID.
	SecurityAltID *string `fix:"455"`
	//SecurityAltIDSource is a non-required field for NoSecurityAltID.
	SecurityAltIDSource *string `fix:"456"`
}

//NoEvents is a repeating group in Instrument
type NoEvents struct {
	//EventType is a non-required field for NoEvents.
	EventType *int `fix:"865"`
	//EventDate is a non-required field for NoEvents.
	EventDate *string `fix:"866"`
	//EventPx is a non-required field for NoEvents.
	EventPx *float64 `fix:"867"`
	//EventText is a non-required field for NoEvents.
	EventText *string `fix:"868"`
	//EventTime is a non-required field for NoEvents.
	EventTime *time.Time `fix:"1145"`
}

//NoInstrumentParties is a repeating group in Instrument
type NoInstrumentParties struct {
	//InstrumentPartyID is a non-required field for NoInstrumentParties.
	InstrumentPartyID *string `fix:"1019"`
	//InstrumentPartyIDSource is a non-required field for NoInstrumentParties.
	InstrumentPartyIDSource *string `fix:"1050"`
	//InstrumentPartyRole is a non-required field for NoInstrumentParties.
	InstrumentPartyRole *int `fix:"1051"`
	//InstrumentPtysSubGrp Component
	instrumentptyssubgrp.InstrumentPtysSubGrp
}

//NoComplexEvents is a repeating group in Instrument
type NoComplexEvents struct {
	//ComplexEventType is a non-required field for NoComplexEvents.
	ComplexEventType *int `fix:"1484"`
	//ComplexOptPayoutAmount is a non-required field for NoComplexEvents.
	ComplexOptPayoutAmount *float64 `fix:"1485"`
	//ComplexEventPrice is a non-required field for NoComplexEvents.
	ComplexEventPrice *float64 `fix:"1486"`
	//ComplexEventPriceBoundaryMethod is a non-required field for NoComplexEvents.
	ComplexEventPriceBoundaryMethod *int `fix:"1487"`
	//ComplexEventPriceBoundaryPrecision is a non-required field for NoComplexEvents.
	ComplexEventPriceBoundaryPrecision *float64 `fix:"1488"`
	//ComplexEventPriceTimeType is a non-required field for NoComplexEvents.
	ComplexEventPriceTimeType *int `fix:"1489"`
	//ComplexEventCondition is a non-required field for NoComplexEvents.
	ComplexEventCondition *int `fix:"1490"`
	//ComplexEventDates Component
	complexeventdates.ComplexEventDates
}

//Instrument is a fix50sp2 Component
type Instrument struct {
	//Symbol is a non-required field for Instrument.
	Symbol *string `fix:"55"`
	//SymbolSfx is a non-required field for Instrument.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for Instrument.
	SecurityID *string `fix:"48"`
	//SecurityIDSource is a non-required field for Instrument.
	SecurityIDSource *string `fix:"22"`
	//NoSecurityAltID is a non-required field for Instrument.
	NoSecurityAltID []NoSecurityAltID `fix:"454,omitempty"`
	//Product is a non-required field for Instrument.
	Product *int `fix:"460"`
	//CFICode is a non-required field for Instrument.
	CFICode *string `fix:"461"`
	//SecurityType is a non-required field for Instrument.
	SecurityType *string `fix:"167"`
	//SecuritySubType is a non-required field for Instrument.
	SecuritySubType *string `fix:"762"`
	//MaturityMonthYear is a non-required field for Instrument.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDate is a non-required field for Instrument.
	MaturityDate *string `fix:"541"`
	//CouponPaymentDate is a non-required field for Instrument.
	CouponPaymentDate *string `fix:"224"`
	//IssueDate is a non-required field for Instrument.
	IssueDate *string `fix:"225"`
	//RepoCollateralSecurityType is a non-required field for Instrument.
	RepoCollateralSecurityType *int `fix:"239"`
	//RepurchaseTerm is a non-required field for Instrument.
	RepurchaseTerm *int `fix:"226"`
	//RepurchaseRate is a non-required field for Instrument.
	RepurchaseRate *float64 `fix:"227"`
	//Factor is a non-required field for Instrument.
	Factor *float64 `fix:"228"`
	//CreditRating is a non-required field for Instrument.
	CreditRating *string `fix:"255"`
	//InstrRegistry is a non-required field for Instrument.
	InstrRegistry *string `fix:"543"`
	//CountryOfIssue is a non-required field for Instrument.
	CountryOfIssue *string `fix:"470"`
	//StateOrProvinceOfIssue is a non-required field for Instrument.
	StateOrProvinceOfIssue *string `fix:"471"`
	//LocaleOfIssue is a non-required field for Instrument.
	LocaleOfIssue *string `fix:"472"`
	//RedemptionDate is a non-required field for Instrument.
	RedemptionDate *string `fix:"240"`
	//StrikePrice is a non-required field for Instrument.
	StrikePrice *float64 `fix:"202"`
	//StrikeCurrency is a non-required field for Instrument.
	StrikeCurrency *string `fix:"947"`
	//OptAttribute is a non-required field for Instrument.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for Instrument.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for Instrument.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for Instrument.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for Instrument.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for Instrument.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for Instrument.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for Instrument.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for Instrument.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for Instrument.
	EncodedSecurityDesc *string `fix:"351"`
	//Pool is a non-required field for Instrument.
	Pool *string `fix:"691"`
	//ContractSettlMonth is a non-required field for Instrument.
	ContractSettlMonth *string `fix:"667"`
	//CPProgram is a non-required field for Instrument.
	CPProgram *int `fix:"875"`
	//CPRegType is a non-required field for Instrument.
	CPRegType *string `fix:"876"`
	//NoEvents is a non-required field for Instrument.
	NoEvents []NoEvents `fix:"864,omitempty"`
	//DatedDate is a non-required field for Instrument.
	DatedDate *string `fix:"873"`
	//InterestAccrualDate is a non-required field for Instrument.
	InterestAccrualDate *string `fix:"874"`
	//SecurityStatus is a non-required field for Instrument.
	SecurityStatus *string `fix:"965"`
	//SettleOnOpenFlag is a non-required field for Instrument.
	SettleOnOpenFlag *string `fix:"966"`
	//InstrmtAssignmentMethod is a non-required field for Instrument.
	InstrmtAssignmentMethod *string `fix:"1049"`
	//StrikeMultiplier is a non-required field for Instrument.
	StrikeMultiplier *float64 `fix:"967"`
	//StrikeValue is a non-required field for Instrument.
	StrikeValue *float64 `fix:"968"`
	//MinPriceIncrement is a non-required field for Instrument.
	MinPriceIncrement *float64 `fix:"969"`
	//PositionLimit is a non-required field for Instrument.
	PositionLimit *int `fix:"970"`
	//NTPositionLimit is a non-required field for Instrument.
	NTPositionLimit *int `fix:"971"`
	//NoInstrumentParties is a non-required field for Instrument.
	NoInstrumentParties []NoInstrumentParties `fix:"1018,omitempty"`
	//UnitOfMeasure is a non-required field for Instrument.
	UnitOfMeasure *string `fix:"996"`
	//TimeUnit is a non-required field for Instrument.
	TimeUnit *string `fix:"997"`
	//MaturityTime is a non-required field for Instrument.
	MaturityTime *string `fix:"1079"`
	//SecurityGroup is a non-required field for Instrument.
	SecurityGroup *string `fix:"1151"`
	//MinPriceIncrementAmount is a non-required field for Instrument.
	MinPriceIncrementAmount *float64 `fix:"1146"`
	//UnitOfMeasureQty is a non-required field for Instrument.
	UnitOfMeasureQty *float64 `fix:"1147"`
	//SecurityXMLLen is a non-required field for Instrument.
	SecurityXMLLen *int `fix:"1184"`
	//SecurityXML is a non-required field for Instrument.
	SecurityXML *string `fix:"1185"`
	//SecurityXMLSchema is a non-required field for Instrument.
	SecurityXMLSchema *string `fix:"1186"`
	//ProductComplex is a non-required field for Instrument.
	ProductComplex *string `fix:"1227"`
	//PriceUnitOfMeasure is a non-required field for Instrument.
	PriceUnitOfMeasure *string `fix:"1191"`
	//PriceUnitOfMeasureQty is a non-required field for Instrument.
	PriceUnitOfMeasureQty *float64 `fix:"1192"`
	//SettlMethod is a non-required field for Instrument.
	SettlMethod *string `fix:"1193"`
	//ExerciseStyle is a non-required field for Instrument.
	ExerciseStyle *int `fix:"1194"`
	//OptPayoutAmount is a non-required field for Instrument.
	OptPayoutAmount *float64 `fix:"1195"`
	//PriceQuoteMethod is a non-required field for Instrument.
	PriceQuoteMethod *string `fix:"1196"`
	//ListMethod is a non-required field for Instrument.
	ListMethod *int `fix:"1198"`
	//CapPrice is a non-required field for Instrument.
	CapPrice *float64 `fix:"1199"`
	//FloorPrice is a non-required field for Instrument.
	FloorPrice *float64 `fix:"1200"`
	//PutOrCall is a non-required field for Instrument.
	PutOrCall *int `fix:"201"`
	//FlexibleIndicator is a non-required field for Instrument.
	FlexibleIndicator *bool `fix:"1244"`
	//FlexProductEligibilityIndicator is a non-required field for Instrument.
	FlexProductEligibilityIndicator *bool `fix:"1242"`
	//ValuationMethod is a non-required field for Instrument.
	ValuationMethod *string `fix:"1197"`
	//ContractMultiplierUnit is a non-required field for Instrument.
	ContractMultiplierUnit *int `fix:"1435"`
	//FlowScheduleType is a non-required field for Instrument.
	FlowScheduleType *int `fix:"1439"`
	//RestructuringType is a non-required field for Instrument.
	RestructuringType *string `fix:"1449"`
	//Seniority is a non-required field for Instrument.
	Seniority *string `fix:"1450"`
	//NotionalPercentageOutstanding is a non-required field for Instrument.
	NotionalPercentageOutstanding *float64 `fix:"1451"`
	//OriginalNotionalPercentageOutstanding is a non-required field for Instrument.
	OriginalNotionalPercentageOutstanding *float64 `fix:"1452"`
	//AttachmentPoint is a non-required field for Instrument.
	AttachmentPoint *float64 `fix:"1457"`
	//DetachmentPoint is a non-required field for Instrument.
	DetachmentPoint *float64 `fix:"1458"`
	//StrikePriceDeterminationMethod is a non-required field for Instrument.
	StrikePriceDeterminationMethod *int `fix:"1478"`
	//StrikePriceBoundaryMethod is a non-required field for Instrument.
	StrikePriceBoundaryMethod *int `fix:"1479"`
	//StrikePriceBoundaryPrecision is a non-required field for Instrument.
	StrikePriceBoundaryPrecision *float64 `fix:"1480"`
	//UnderlyingPriceDeterminationMethod is a non-required field for Instrument.
	UnderlyingPriceDeterminationMethod *int `fix:"1481"`
	//OptPayoutType is a non-required field for Instrument.
	OptPayoutType *int `fix:"1482"`
	//NoComplexEvents is a non-required field for Instrument.
	NoComplexEvents []NoComplexEvents `fix:"1483,omitempty"`
}

func (m *Instrument) SetSymbol(v string)                             { m.Symbol = &v }
func (m *Instrument) SetSymbolSfx(v string)                          { m.SymbolSfx = &v }
func (m *Instrument) SetSecurityID(v string)                         { m.SecurityID = &v }
func (m *Instrument) SetSecurityIDSource(v string)                   { m.SecurityIDSource = &v }
func (m *Instrument) SetNoSecurityAltID(v []NoSecurityAltID)         { m.NoSecurityAltID = v }
func (m *Instrument) SetProduct(v int)                               { m.Product = &v }
func (m *Instrument) SetCFICode(v string)                            { m.CFICode = &v }
func (m *Instrument) SetSecurityType(v string)                       { m.SecurityType = &v }
func (m *Instrument) SetSecuritySubType(v string)                    { m.SecuritySubType = &v }
func (m *Instrument) SetMaturityMonthYear(v string)                  { m.MaturityMonthYear = &v }
func (m *Instrument) SetMaturityDate(v string)                       { m.MaturityDate = &v }
func (m *Instrument) SetCouponPaymentDate(v string)                  { m.CouponPaymentDate = &v }
func (m *Instrument) SetIssueDate(v string)                          { m.IssueDate = &v }
func (m *Instrument) SetRepoCollateralSecurityType(v int)            { m.RepoCollateralSecurityType = &v }
func (m *Instrument) SetRepurchaseTerm(v int)                        { m.RepurchaseTerm = &v }
func (m *Instrument) SetRepurchaseRate(v float64)                    { m.RepurchaseRate = &v }
func (m *Instrument) SetFactor(v float64)                            { m.Factor = &v }
func (m *Instrument) SetCreditRating(v string)                       { m.CreditRating = &v }
func (m *Instrument) SetInstrRegistry(v string)                      { m.InstrRegistry = &v }
func (m *Instrument) SetCountryOfIssue(v string)                     { m.CountryOfIssue = &v }
func (m *Instrument) SetStateOrProvinceOfIssue(v string)             { m.StateOrProvinceOfIssue = &v }
func (m *Instrument) SetLocaleOfIssue(v string)                      { m.LocaleOfIssue = &v }
func (m *Instrument) SetRedemptionDate(v string)                     { m.RedemptionDate = &v }
func (m *Instrument) SetStrikePrice(v float64)                       { m.StrikePrice = &v }
func (m *Instrument) SetStrikeCurrency(v string)                     { m.StrikeCurrency = &v }
func (m *Instrument) SetOptAttribute(v string)                       { m.OptAttribute = &v }
func (m *Instrument) SetContractMultiplier(v float64)                { m.ContractMultiplier = &v }
func (m *Instrument) SetCouponRate(v float64)                        { m.CouponRate = &v }
func (m *Instrument) SetSecurityExchange(v string)                   { m.SecurityExchange = &v }
func (m *Instrument) SetIssuer(v string)                             { m.Issuer = &v }
func (m *Instrument) SetEncodedIssuerLen(v int)                      { m.EncodedIssuerLen = &v }
func (m *Instrument) SetEncodedIssuer(v string)                      { m.EncodedIssuer = &v }
func (m *Instrument) SetSecurityDesc(v string)                       { m.SecurityDesc = &v }
func (m *Instrument) SetEncodedSecurityDescLen(v int)                { m.EncodedSecurityDescLen = &v }
func (m *Instrument) SetEncodedSecurityDesc(v string)                { m.EncodedSecurityDesc = &v }
func (m *Instrument) SetPool(v string)                               { m.Pool = &v }
func (m *Instrument) SetContractSettlMonth(v string)                 { m.ContractSettlMonth = &v }
func (m *Instrument) SetCPProgram(v int)                             { m.CPProgram = &v }
func (m *Instrument) SetCPRegType(v string)                          { m.CPRegType = &v }
func (m *Instrument) SetNoEvents(v []NoEvents)                       { m.NoEvents = v }
func (m *Instrument) SetDatedDate(v string)                          { m.DatedDate = &v }
func (m *Instrument) SetInterestAccrualDate(v string)                { m.InterestAccrualDate = &v }
func (m *Instrument) SetSecurityStatus(v string)                     { m.SecurityStatus = &v }
func (m *Instrument) SetSettleOnOpenFlag(v string)                   { m.SettleOnOpenFlag = &v }
func (m *Instrument) SetInstrmtAssignmentMethod(v string)            { m.InstrmtAssignmentMethod = &v }
func (m *Instrument) SetStrikeMultiplier(v float64)                  { m.StrikeMultiplier = &v }
func (m *Instrument) SetStrikeValue(v float64)                       { m.StrikeValue = &v }
func (m *Instrument) SetMinPriceIncrement(v float64)                 { m.MinPriceIncrement = &v }
func (m *Instrument) SetPositionLimit(v int)                         { m.PositionLimit = &v }
func (m *Instrument) SetNTPositionLimit(v int)                       { m.NTPositionLimit = &v }
func (m *Instrument) SetNoInstrumentParties(v []NoInstrumentParties) { m.NoInstrumentParties = v }
func (m *Instrument) SetUnitOfMeasure(v string)                      { m.UnitOfMeasure = &v }
func (m *Instrument) SetTimeUnit(v string)                           { m.TimeUnit = &v }
func (m *Instrument) SetMaturityTime(v string)                       { m.MaturityTime = &v }
func (m *Instrument) SetSecurityGroup(v string)                      { m.SecurityGroup = &v }
func (m *Instrument) SetMinPriceIncrementAmount(v float64)           { m.MinPriceIncrementAmount = &v }
func (m *Instrument) SetUnitOfMeasureQty(v float64)                  { m.UnitOfMeasureQty = &v }
func (m *Instrument) SetSecurityXMLLen(v int)                        { m.SecurityXMLLen = &v }
func (m *Instrument) SetSecurityXML(v string)                        { m.SecurityXML = &v }
func (m *Instrument) SetSecurityXMLSchema(v string)                  { m.SecurityXMLSchema = &v }
func (m *Instrument) SetProductComplex(v string)                     { m.ProductComplex = &v }
func (m *Instrument) SetPriceUnitOfMeasure(v string)                 { m.PriceUnitOfMeasure = &v }
func (m *Instrument) SetPriceUnitOfMeasureQty(v float64)             { m.PriceUnitOfMeasureQty = &v }
func (m *Instrument) SetSettlMethod(v string)                        { m.SettlMethod = &v }
func (m *Instrument) SetExerciseStyle(v int)                         { m.ExerciseStyle = &v }
func (m *Instrument) SetOptPayoutAmount(v float64)                   { m.OptPayoutAmount = &v }
func (m *Instrument) SetPriceQuoteMethod(v string)                   { m.PriceQuoteMethod = &v }
func (m *Instrument) SetListMethod(v int)                            { m.ListMethod = &v }
func (m *Instrument) SetCapPrice(v float64)                          { m.CapPrice = &v }
func (m *Instrument) SetFloorPrice(v float64)                        { m.FloorPrice = &v }
func (m *Instrument) SetPutOrCall(v int)                             { m.PutOrCall = &v }
func (m *Instrument) SetFlexibleIndicator(v bool)                    { m.FlexibleIndicator = &v }
func (m *Instrument) SetFlexProductEligibilityIndicator(v bool) {
	m.FlexProductEligibilityIndicator = &v
}
func (m *Instrument) SetValuationMethod(v string)                { m.ValuationMethod = &v }
func (m *Instrument) SetContractMultiplierUnit(v int)            { m.ContractMultiplierUnit = &v }
func (m *Instrument) SetFlowScheduleType(v int)                  { m.FlowScheduleType = &v }
func (m *Instrument) SetRestructuringType(v string)              { m.RestructuringType = &v }
func (m *Instrument) SetSeniority(v string)                      { m.Seniority = &v }
func (m *Instrument) SetNotionalPercentageOutstanding(v float64) { m.NotionalPercentageOutstanding = &v }
func (m *Instrument) SetOriginalNotionalPercentageOutstanding(v float64) {
	m.OriginalNotionalPercentageOutstanding = &v
}
func (m *Instrument) SetAttachmentPoint(v float64)              { m.AttachmentPoint = &v }
func (m *Instrument) SetDetachmentPoint(v float64)              { m.DetachmentPoint = &v }
func (m *Instrument) SetStrikePriceDeterminationMethod(v int)   { m.StrikePriceDeterminationMethod = &v }
func (m *Instrument) SetStrikePriceBoundaryMethod(v int)        { m.StrikePriceBoundaryMethod = &v }
func (m *Instrument) SetStrikePriceBoundaryPrecision(v float64) { m.StrikePriceBoundaryPrecision = &v }
func (m *Instrument) SetUnderlyingPriceDeterminationMethod(v int) {
	m.UnderlyingPriceDeterminationMethod = &v
}
func (m *Instrument) SetOptPayoutType(v int)                 { m.OptPayoutType = &v }
func (m *Instrument) SetNoComplexEvents(v []NoComplexEvents) { m.NoComplexEvents = v }
