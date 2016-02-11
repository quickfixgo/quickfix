package derivativesecurityaltidgrp

//NoDerivativeSecurityAltID is a repeating group in DerivativeSecurityAltIDGrp
type NoDerivativeSecurityAltID struct {
	//DerivativeSecurityAltID is a non-required field for NoDerivativeSecurityAltID.
	DerivativeSecurityAltID *string `fix:"1219"`
	//DerivativeSecurityAltIDSource is a non-required field for NoDerivativeSecurityAltID.
	DerivativeSecurityAltIDSource *string `fix:"1220"`
}

//Component is a fix50sp1 DerivativeSecurityAltIDGrp Component
type Component struct {
	//NoDerivativeSecurityAltID is a non-required field for DerivativeSecurityAltIDGrp.
	NoDerivativeSecurityAltID []NoDerivativeSecurityAltID `fix:"1218,omitempty"`
}

func New() *Component { return new(Component) }
