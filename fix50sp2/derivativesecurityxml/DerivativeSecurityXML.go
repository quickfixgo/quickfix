package derivativesecurityxml

//DerivativeSecurityXML is a fix50sp2 Component
type DerivativeSecurityXML struct {
	//DerivativeSecurityXMLLen is a non-required field for DerivativeSecurityXML.
	DerivativeSecurityXMLLen *int `fix:"1282"`
	//DerivativeSecurityXML is a non-required field for DerivativeSecurityXML.
	DerivativeSecurityXML *string `fix:"1283"`
	//DerivativeSecurityXMLSchema is a non-required field for DerivativeSecurityXML.
	DerivativeSecurityXMLSchema *string `fix:"1284"`
}

func (m *DerivativeSecurityXML) SetDerivativeSecurityXMLLen(v int) { m.DerivativeSecurityXMLLen = &v }
func (m *DerivativeSecurityXML) SetDerivativeSecurityXML(v string) { m.DerivativeSecurityXML = &v }
func (m *DerivativeSecurityXML) SetDerivativeSecurityXMLSchema(v string) {
	m.DerivativeSecurityXMLSchema = &v
}
