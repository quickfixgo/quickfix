package nestedinstrumentattribute

//New returns an initialized NestedInstrumentAttribute instance
func New() *NestedInstrumentAttribute {
	var m NestedInstrumentAttribute
	return &m
}

//NoNestedInstrAttrib is a repeating group in NestedInstrumentAttribute
type NoNestedInstrAttrib struct {
	//NestedInstrAttribType is a non-required field for NoNestedInstrAttrib.
	NestedInstrAttribType *int `fix:"1210"`
	//NestedInstrAttribValue is a non-required field for NoNestedInstrAttrib.
	NestedInstrAttribValue *string `fix:"1211"`
}

//NewNoNestedInstrAttrib returns an initialized NoNestedInstrAttrib instance
func NewNoNestedInstrAttrib() *NoNestedInstrAttrib {
	var m NoNestedInstrAttrib
	return &m
}

func (m *NoNestedInstrAttrib) SetNestedInstrAttribType(v int)     { m.NestedInstrAttribType = &v }
func (m *NoNestedInstrAttrib) SetNestedInstrAttribValue(v string) { m.NestedInstrAttribValue = &v }

//NestedInstrumentAttribute is a fix50sp1 Component
type NestedInstrumentAttribute struct {
	//NoNestedInstrAttrib is a non-required field for NestedInstrumentAttribute.
	NoNestedInstrAttrib []NoNestedInstrAttrib `fix:"1312,omitempty"`
}

func (m *NestedInstrumentAttribute) SetNoNestedInstrAttrib(v []NoNestedInstrAttrib) {
	m.NoNestedInstrAttrib = v
}
