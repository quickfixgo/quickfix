package derivativeinstrumentattribute

//NoDerivativeInstrAttrib is a repeating group in DerivativeInstrumentAttribute
type NoDerivativeInstrAttrib struct {
	//DerivativeInstrAttribType is a non-required field for NoDerivativeInstrAttrib.
	DerivativeInstrAttribType *int `fix:"1313"`
	//DerivativeInstrAttribValue is a non-required field for NoDerivativeInstrAttrib.
	DerivativeInstrAttribValue *string `fix:"1314"`
}

//Component is a fix50sp2 DerivativeInstrumentAttribute Component
type Component struct {
	//NoDerivativeInstrAttrib is a non-required field for DerivativeInstrumentAttribute.
	NoDerivativeInstrAttrib []NoDerivativeInstrAttrib `fix:"1311,omitempty"`
}

func New() *Component { return new(Component) }
