package derivativeinstrumentparties

import (
	"github.com/quickfixgo/quickfix/fix50sp1/derivativeinstrumentpartysubidsgrp"
)

//New returns an initialized DerivativeInstrumentParties instance
func New() *DerivativeInstrumentParties {
	var m DerivativeInstrumentParties
	return &m
}

//NoDerivativeInstrumentParties is a repeating group in DerivativeInstrumentParties
type NoDerivativeInstrumentParties struct {
	//DerivativeInstrumentPartyID is a non-required field for NoDerivativeInstrumentParties.
	DerivativeInstrumentPartyID *string `fix:"1293"`
	//DerivativeInstrumentPartyIDSource is a non-required field for NoDerivativeInstrumentParties.
	DerivativeInstrumentPartyIDSource *string `fix:"1294"`
	//DerivativeInstrumentPartyRole is a non-required field for NoDerivativeInstrumentParties.
	DerivativeInstrumentPartyRole *int `fix:"1295"`
	//DerivativeInstrumentPartySubIDsGrp is a non-required component for NoDerivativeInstrumentParties.
	DerivativeInstrumentPartySubIDsGrp *derivativeinstrumentpartysubidsgrp.DerivativeInstrumentPartySubIDsGrp
}

//NewNoDerivativeInstrumentParties returns an initialized NoDerivativeInstrumentParties instance
func NewNoDerivativeInstrumentParties() *NoDerivativeInstrumentParties {
	var m NoDerivativeInstrumentParties
	return &m
}

func (m *NoDerivativeInstrumentParties) SetDerivativeInstrumentPartyID(v string) {
	m.DerivativeInstrumentPartyID = &v
}
func (m *NoDerivativeInstrumentParties) SetDerivativeInstrumentPartyIDSource(v string) {
	m.DerivativeInstrumentPartyIDSource = &v
}
func (m *NoDerivativeInstrumentParties) SetDerivativeInstrumentPartyRole(v int) {
	m.DerivativeInstrumentPartyRole = &v
}
func (m *NoDerivativeInstrumentParties) SetDerivativeInstrumentPartySubIDsGrp(v derivativeinstrumentpartysubidsgrp.DerivativeInstrumentPartySubIDsGrp) {
	m.DerivativeInstrumentPartySubIDsGrp = &v
}

//DerivativeInstrumentParties is a fix50sp1 Component
type DerivativeInstrumentParties struct {
	//NoDerivativeInstrumentParties is a non-required field for DerivativeInstrumentParties.
	NoDerivativeInstrumentParties []NoDerivativeInstrumentParties `fix:"1292,omitempty"`
}

func (m *DerivativeInstrumentParties) SetNoDerivativeInstrumentParties(v []NoDerivativeInstrumentParties) {
	m.NoDerivativeInstrumentParties = v
}
