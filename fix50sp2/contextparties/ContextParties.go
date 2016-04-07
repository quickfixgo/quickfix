package contextparties

import (
	"github.com/quickfixgo/quickfix/fix50sp2/contextptyssubgrp"
)

func New() *ContextParties {
	var m ContextParties
	return &m
}

//NoContextPartyIDs is a repeating group in ContextParties
type NoContextPartyIDs struct {
	//ContextPartyID is a non-required field for NoContextPartyIDs.
	ContextPartyID *string `fix:"1523"`
	//ContextPartyIDSource is a non-required field for NoContextPartyIDs.
	ContextPartyIDSource *string `fix:"1524"`
	//ContextPartyRole is a non-required field for NoContextPartyIDs.
	ContextPartyRole *int `fix:"1525"`
	//ContextPtysSubGrp is a non-required component for NoContextPartyIDs.
	ContextPtysSubGrp *contextptyssubgrp.ContextPtysSubGrp
}

func (m *NoContextPartyIDs) SetContextPartyID(v string)       { m.ContextPartyID = &v }
func (m *NoContextPartyIDs) SetContextPartyIDSource(v string) { m.ContextPartyIDSource = &v }
func (m *NoContextPartyIDs) SetContextPartyRole(v int)        { m.ContextPartyRole = &v }
func (m *NoContextPartyIDs) SetContextPtysSubGrp(v contextptyssubgrp.ContextPtysSubGrp) {
	m.ContextPtysSubGrp = &v
}

//ContextParties is a fix50sp2 Component
type ContextParties struct {
	//NoContextPartyIDs is a non-required field for ContextParties.
	NoContextPartyIDs []NoContextPartyIDs `fix:"1522,omitempty"`
}

func (m *ContextParties) SetNoContextPartyIDs(v []NoContextPartyIDs) { m.NoContextPartyIDs = v }
