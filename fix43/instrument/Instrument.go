package instrument

//NoSecurityAltID is a repeating group in Instrument
type NoSecurityAltID struct {
	//SecurityAltID is a non-required field for NoSecurityAltID.
	SecurityAltID *string `fix:"455"`
	//SecurityAltIDSource is a non-required field for NoSecurityAltID.
	SecurityAltIDSource *string `fix:"456"`
}

//Component is a fix43 Instrument Component
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

func New() *Component { return new(Component) }

func (m *Component) SetSymbol(v string)                     { m.Symbol = &v }
func (m *Component) SetSymbolSfx(v string)                  { m.SymbolSfx = &v }
func (m *Component) SetSecurityID(v string)                 { m.SecurityID = &v }
func (m *Component) SetSecurityIDSource(v string)           { m.SecurityIDSource = &v }
func (m *Component) SetNoSecurityAltID(v []NoSecurityAltID) { m.NoSecurityAltID = v }
func (m *Component) SetProduct(v int)                       { m.Product = &v }
func (m *Component) SetCFICode(v string)                    { m.CFICode = &v }
func (m *Component) SetSecurityType(v string)               { m.SecurityType = &v }
func (m *Component) SetMaturityMonthYear(v string)          { m.MaturityMonthYear = &v }
func (m *Component) SetMaturityDate(v string)               { m.MaturityDate = &v }
func (m *Component) SetCouponPaymentDate(v string)          { m.CouponPaymentDate = &v }
func (m *Component) SetIssueDate(v string)                  { m.IssueDate = &v }
func (m *Component) SetRepoCollateralSecurityType(v int)    { m.RepoCollateralSecurityType = &v }
func (m *Component) SetRepurchaseTerm(v int)                { m.RepurchaseTerm = &v }
func (m *Component) SetRepurchaseRate(v float64)            { m.RepurchaseRate = &v }
func (m *Component) SetFactor(v float64)                    { m.Factor = &v }
func (m *Component) SetCreditRating(v string)               { m.CreditRating = &v }
func (m *Component) SetInstrRegistry(v string)              { m.InstrRegistry = &v }
func (m *Component) SetCountryOfIssue(v string)             { m.CountryOfIssue = &v }
func (m *Component) SetStateOrProvinceOfIssue(v string)     { m.StateOrProvinceOfIssue = &v }
func (m *Component) SetLocaleOfIssue(v string)              { m.LocaleOfIssue = &v }
func (m *Component) SetRedemptionDate(v string)             { m.RedemptionDate = &v }
func (m *Component) SetStrikePrice(v float64)               { m.StrikePrice = &v }
func (m *Component) SetOptAttribute(v string)               { m.OptAttribute = &v }
func (m *Component) SetContractMultiplier(v float64)        { m.ContractMultiplier = &v }
func (m *Component) SetCouponRate(v float64)                { m.CouponRate = &v }
func (m *Component) SetSecurityExchange(v string)           { m.SecurityExchange = &v }
func (m *Component) SetIssuer(v string)                     { m.Issuer = &v }
func (m *Component) SetEncodedIssuerLen(v int)              { m.EncodedIssuerLen = &v }
func (m *Component) SetEncodedIssuer(v string)              { m.EncodedIssuer = &v }
func (m *Component) SetSecurityDesc(v string)               { m.SecurityDesc = &v }
func (m *Component) SetEncodedSecurityDescLen(v int)        { m.EncodedSecurityDescLen = &v }
func (m *Component) SetEncodedSecurityDesc(v string)        { m.EncodedSecurityDesc = &v }
