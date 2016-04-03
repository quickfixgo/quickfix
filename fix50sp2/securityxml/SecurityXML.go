package securityxml

func New() *SecurityXML {
	var m SecurityXML
	return &m
}

//SecurityXML is a fix50sp2 Component
type SecurityXML struct {
	//SecurityXMLLen is a non-required field for SecurityXML.
	SecurityXMLLen *int `fix:"1184"`
	//SecurityXML is a non-required field for SecurityXML.
	SecurityXML *string `fix:"1185"`
	//SecurityXMLSchema is a non-required field for SecurityXML.
	SecurityXMLSchema *string `fix:"1186"`
}

func (m *SecurityXML) SetSecurityXMLLen(v int)       { m.SecurityXMLLen = &v }
func (m *SecurityXML) SetSecurityXML(v string)       { m.SecurityXML = &v }
func (m *SecurityXML) SetSecurityXMLSchema(v string) { m.SecurityXMLSchema = &v }
