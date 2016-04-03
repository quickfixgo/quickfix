package nestedinstrumentattribute

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

//NestedInstrumentAttribute is a fix50sp2 Component
type NestedInstrumentAttribute struct {
	//NoNestedInstrAttrib is a non-required field for NestedInstrumentAttribute.
	NoNestedInstrAttrib []NoNestedInstrAttrib `fix:"1312,omitempty"`
}

func (m *NestedInstrumentAttribute) SetNoNestedInstrAttrib(v []NoNestedInstrAttrib) {
	m.NoNestedInstrAttrib = v
}
