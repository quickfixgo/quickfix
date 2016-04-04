package instrument

func New() *Instrument {
	var m Instrument
	return &m
}

//NoSecurityAltID is a repeating group in Instrument
type NoSecurityAltID struct {
	//SecurityAltID is a non-required field for NoSecurityAltID.
	SecurityAltID *string `fix:"455"`
	//SecurityAltIDSource is a non-required field for NoSecurityAltID.
	SecurityAltIDSource *string `fix:"456"`
}

func (m *NoSecurityAltID) SetSecurityAltID(v string)       { m.SecurityAltID = &v }
func (m *NoSecurityAltID) SetSecurityAltIDSource(v string) { m.SecurityAltIDSource = &v }

//Instrument is a fix43 Component
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
}

func (m *Instrument) SetSymbol(v string)                     { m.Symbol = &v }
func (m *Instrument) SetSymbolSfx(v string)                  { m.SymbolSfx = &v }
func (m *Instrument) SetSecurityID(v string)                 { m.SecurityID = &v }
func (m *Instrument) SetSecurityIDSource(v string)           { m.SecurityIDSource = &v }
func (m *Instrument) SetNoSecurityAltID(v []NoSecurityAltID) { m.NoSecurityAltID = v }
func (m *Instrument) SetProduct(v int)                       { m.Product = &v }
func (m *Instrument) SetCFICode(v string)                    { m.CFICode = &v }
func (m *Instrument) SetSecurityType(v string)               { m.SecurityType = &v }
func (m *Instrument) SetMaturityMonthYear(v string)          { m.MaturityMonthYear = &v }
func (m *Instrument) SetMaturityDate(v string)               { m.MaturityDate = &v }
func (m *Instrument) SetCouponPaymentDate(v string)          { m.CouponPaymentDate = &v }
func (m *Instrument) SetIssueDate(v string)                  { m.IssueDate = &v }
func (m *Instrument) SetRepoCollateralSecurityType(v int)    { m.RepoCollateralSecurityType = &v }
func (m *Instrument) SetRepurchaseTerm(v int)                { m.RepurchaseTerm = &v }
func (m *Instrument) SetRepurchaseRate(v float64)            { m.RepurchaseRate = &v }
func (m *Instrument) SetFactor(v float64)                    { m.Factor = &v }
func (m *Instrument) SetCreditRating(v string)               { m.CreditRating = &v }
func (m *Instrument) SetInstrRegistry(v string)              { m.InstrRegistry = &v }
func (m *Instrument) SetCountryOfIssue(v string)             { m.CountryOfIssue = &v }
func (m *Instrument) SetStateOrProvinceOfIssue(v string)     { m.StateOrProvinceOfIssue = &v }
func (m *Instrument) SetLocaleOfIssue(v string)              { m.LocaleOfIssue = &v }
func (m *Instrument) SetRedemptionDate(v string)             { m.RedemptionDate = &v }
func (m *Instrument) SetStrikePrice(v float64)               { m.StrikePrice = &v }
func (m *Instrument) SetOptAttribute(v string)               { m.OptAttribute = &v }
func (m *Instrument) SetContractMultiplier(v float64)        { m.ContractMultiplier = &v }
func (m *Instrument) SetCouponRate(v float64)                { m.CouponRate = &v }
func (m *Instrument) SetSecurityExchange(v string)           { m.SecurityExchange = &v }
func (m *Instrument) SetIssuer(v string)                     { m.Issuer = &v }
func (m *Instrument) SetEncodedIssuerLen(v int)              { m.EncodedIssuerLen = &v }
func (m *Instrument) SetEncodedIssuer(v string)              { m.EncodedIssuer = &v }
func (m *Instrument) SetSecurityDesc(v string)               { m.SecurityDesc = &v }
func (m *Instrument) SetEncodedSecurityDescLen(v int)        { m.EncodedSecurityDescLen = &v }
func (m *Instrument) SetEncodedSecurityDesc(v string)        { m.EncodedSecurityDesc = &v }
