package nestedinstrumentattribute

//NoNestedInstrAttrib is a repeating group in NestedInstrumentAttribute
type NoNestedInstrAttrib struct {
	//NestedInstrAttribType is a non-required field for NoNestedInstrAttrib.
	NestedInstrAttribType *int `fix:"1210"`
	//NestedInstrAttribValue is a non-required field for NoNestedInstrAttrib.
	NestedInstrAttribValue *string `fix:"1211"`
}

//Component is a fix50sp1 NestedInstrumentAttribute Component
type Component struct {
	//NoNestedInstrAttrib is a non-required field for NestedInstrumentAttribute.
	NoNestedInstrAttrib []NoNestedInstrAttrib `fix:"1312,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoNestedInstrAttrib(v []NoNestedInstrAttrib) { m.NoNestedInstrAttrib = v }
