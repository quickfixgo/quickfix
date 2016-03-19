package settldetails

import (
	"github.com/quickfixgo/quickfix/fix50sp2/settlparties"
)

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

//SettlDetails is a fix50sp2 Component
type SettlDetails struct {
	//NoSettlDetails is a non-required field for SettlDetails.
	NoSettlDetails []NoSettlDetails `fix:"1158,omitempty"`
}

func (m *SettlDetails) SetNoSettlDetails(v []NoSettlDetails) { m.NoSettlDetails = v }
