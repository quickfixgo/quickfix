package securityxml

//Component is a fix50sp2 SecurityXML Component
type Component struct {
	//SecurityXMLLen is a non-required field for SecurityXML.
	SecurityXMLLen *int `fix:"1184"`
	//SecurityXML is a non-required field for SecurityXML.
	SecurityXML *string `fix:"1185"`
	//SecurityXMLSchema is a non-required field for SecurityXML.
	SecurityXMLSchema *string `fix:"1186"`
}

func New() *Component { return new(Component) }

func (m *Component) SetSecurityXMLLen(v int)       { m.SecurityXMLLen = &v }
func (m *Component) SetSecurityXML(v string)       { m.SecurityXML = &v }
func (m *Component) SetSecurityXMLSchema(v string) { m.SecurityXMLSchema = &v }
