package derivativesecurityxml

//Component is a fix50sp2 DerivativeSecurityXML Component
type Component struct {
	//DerivativeSecurityXMLLen is a non-required field for DerivativeSecurityXML.
	DerivativeSecurityXMLLen *int `fix:"1282"`
	//DerivativeSecurityXML is a non-required field for DerivativeSecurityXML.
	DerivativeSecurityXML *string `fix:"1283"`
	//DerivativeSecurityXMLSchema is a non-required field for DerivativeSecurityXML.
	DerivativeSecurityXMLSchema *string `fix:"1284"`
}

func New() *Component { return new(Component) }

func (m *Component) SetDerivativeSecurityXMLLen(v int)       { m.DerivativeSecurityXMLLen = &v }
func (m *Component) SetDerivativeSecurityXML(v string)       { m.DerivativeSecurityXML = &v }
func (m *Component) SetDerivativeSecurityXMLSchema(v string) { m.DerivativeSecurityXMLSchema = &v }
