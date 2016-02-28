package instrument

import (
	"github.com/quickfixgo/quickfix/fix50/instrumentptyssubgrp"
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
	InstrumentPtysSubGrp instrumentptyssubgrp.Component
}

//Component is a fix50 Instrument Component
type Component struct {
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
}

func New() *Component { return new(Component) }

func (m *Component) SetSymbol(v string)                             { m.Symbol = &v }
func (m *Component) SetSymbolSfx(v string)                          { m.SymbolSfx = &v }
func (m *Component) SetSecurityID(v string)                         { m.SecurityID = &v }
func (m *Component) SetSecurityIDSource(v string)                   { m.SecurityIDSource = &v }
func (m *Component) SetNoSecurityAltID(v []NoSecurityAltID)         { m.NoSecurityAltID = v }
func (m *Component) SetProduct(v int)                               { m.Product = &v }
func (m *Component) SetCFICode(v string)                            { m.CFICode = &v }
func (m *Component) SetSecurityType(v string)                       { m.SecurityType = &v }
func (m *Component) SetSecuritySubType(v string)                    { m.SecuritySubType = &v }
func (m *Component) SetMaturityMonthYear(v string)                  { m.MaturityMonthYear = &v }
func (m *Component) SetMaturityDate(v string)                       { m.MaturityDate = &v }
func (m *Component) SetCouponPaymentDate(v string)                  { m.CouponPaymentDate = &v }
func (m *Component) SetIssueDate(v string)                          { m.IssueDate = &v }
func (m *Component) SetRepoCollateralSecurityType(v int)            { m.RepoCollateralSecurityType = &v }
func (m *Component) SetRepurchaseTerm(v int)                        { m.RepurchaseTerm = &v }
func (m *Component) SetRepurchaseRate(v float64)                    { m.RepurchaseRate = &v }
func (m *Component) SetFactor(v float64)                            { m.Factor = &v }
func (m *Component) SetCreditRating(v string)                       { m.CreditRating = &v }
func (m *Component) SetInstrRegistry(v string)                      { m.InstrRegistry = &v }
func (m *Component) SetCountryOfIssue(v string)                     { m.CountryOfIssue = &v }
func (m *Component) SetStateOrProvinceOfIssue(v string)             { m.StateOrProvinceOfIssue = &v }
func (m *Component) SetLocaleOfIssue(v string)                      { m.LocaleOfIssue = &v }
func (m *Component) SetRedemptionDate(v string)                     { m.RedemptionDate = &v }
func (m *Component) SetStrikePrice(v float64)                       { m.StrikePrice = &v }
func (m *Component) SetStrikeCurrency(v string)                     { m.StrikeCurrency = &v }
func (m *Component) SetOptAttribute(v string)                       { m.OptAttribute = &v }
func (m *Component) SetContractMultiplier(v float64)                { m.ContractMultiplier = &v }
func (m *Component) SetCouponRate(v float64)                        { m.CouponRate = &v }
func (m *Component) SetSecurityExchange(v string)                   { m.SecurityExchange = &v }
func (m *Component) SetIssuer(v string)                             { m.Issuer = &v }
func (m *Component) SetEncodedIssuerLen(v int)                      { m.EncodedIssuerLen = &v }
func (m *Component) SetEncodedIssuer(v string)                      { m.EncodedIssuer = &v }
func (m *Component) SetSecurityDesc(v string)                       { m.SecurityDesc = &v }
func (m *Component) SetEncodedSecurityDescLen(v int)                { m.EncodedSecurityDescLen = &v }
func (m *Component) SetEncodedSecurityDesc(v string)                { m.EncodedSecurityDesc = &v }
func (m *Component) SetPool(v string)                               { m.Pool = &v }
func (m *Component) SetContractSettlMonth(v string)                 { m.ContractSettlMonth = &v }
func (m *Component) SetCPProgram(v int)                             { m.CPProgram = &v }
func (m *Component) SetCPRegType(v string)                          { m.CPRegType = &v }
func (m *Component) SetNoEvents(v []NoEvents)                       { m.NoEvents = v }
func (m *Component) SetDatedDate(v string)                          { m.DatedDate = &v }
func (m *Component) SetInterestAccrualDate(v string)                { m.InterestAccrualDate = &v }
func (m *Component) SetSecurityStatus(v string)                     { m.SecurityStatus = &v }
func (m *Component) SetSettleOnOpenFlag(v string)                   { m.SettleOnOpenFlag = &v }
func (m *Component) SetInstrmtAssignmentMethod(v string)            { m.InstrmtAssignmentMethod = &v }
func (m *Component) SetStrikeMultiplier(v float64)                  { m.StrikeMultiplier = &v }
func (m *Component) SetStrikeValue(v float64)                       { m.StrikeValue = &v }
func (m *Component) SetMinPriceIncrement(v float64)                 { m.MinPriceIncrement = &v }
func (m *Component) SetPositionLimit(v int)                         { m.PositionLimit = &v }
func (m *Component) SetNTPositionLimit(v int)                       { m.NTPositionLimit = &v }
func (m *Component) SetNoInstrumentParties(v []NoInstrumentParties) { m.NoInstrumentParties = v }
func (m *Component) SetUnitOfMeasure(v string)                      { m.UnitOfMeasure = &v }
func (m *Component) SetTimeUnit(v string)                           { m.TimeUnit = &v }
func (m *Component) SetMaturityTime(v string)                       { m.MaturityTime = &v }
