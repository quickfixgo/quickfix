package derivativeinstrumentattribute

//NoDerivativeInstrAttrib is a repeating group in DerivativeInstrumentAttribute
type NoDerivativeInstrAttrib struct {
	//DerivativeInstrAttribType is a non-required field for NoDerivativeInstrAttrib.
	DerivativeInstrAttribType *int `fix:"1313"`
	//DerivativeInstrAttribValue is a non-required field for NoDerivativeInstrAttrib.
	DerivativeInstrAttribValue *string `fix:"1314"`
}

//DerivativeInstrumentAttribute is a fix50sp1 Component
type DerivativeInstrumentAttribute struct {
	//NoDerivativeInstrAttrib is a non-required field for DerivativeInstrumentAttribute.
	NoDerivativeInstrAttrib []NoDerivativeInstrAttrib `fix:"1311,omitempty"`
}

func (m *DerivativeInstrumentAttribute) SetNoDerivativeInstrAttrib(v []NoDerivativeInstrAttrib) {
	m.NoDerivativeInstrAttrib = v
}
