package settldetails

import (
	"github.com/quickfixgo/quickfix/fix50sp1/settlparties"
)

//New returns an initialized SettlDetails instance
func New() *SettlDetails {
	var m SettlDetails
	return &m
}

//NoSettlDetails is a repeating group in SettlDetails
type NoSettlDetails struct {
	//SettlObligSource is a non-required field for NoSettlDetails.
	SettlObligSource *string `fix:"1164"`
	//SettlParties is a non-required component for NoSettlDetails.
	SettlParties *settlparties.SettlParties
}

//NewNoSettlDetails returns an initialized NoSettlDetails instance
func NewNoSettlDetails() *NoSettlDetails {
	var m NoSettlDetails
	return &m
}

func (m *NoSettlDetails) SetSettlObligSource(v string)                { m.SettlObligSource = &v }
func (m *NoSettlDetails) SetSettlParties(v settlparties.SettlParties) { m.SettlParties = &v }

//SettlDetails is a fix50sp1 Component
type SettlDetails struct {
	//NoSettlDetails is a non-required field for SettlDetails.
	NoSettlDetails []NoSettlDetails `fix:"1158,omitempty"`
}

func (m *SettlDetails) SetNoSettlDetails(v []NoSettlDetails) { m.NoSettlDetails = v }
