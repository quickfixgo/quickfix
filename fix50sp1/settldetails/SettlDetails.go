package settldetails

import (
	"github.com/quickfixgo/quickfix/fix50sp1/settlparties"
)

//NoSettlDetails is a repeating group in SettlDetails
type NoSettlDetails struct {
	//SettlObligSource is a non-required field for NoSettlDetails.
	SettlObligSource *string `fix:"1164"`
	//SettlParties Component
	settlparties.SettlParties
}

//SettlDetails is a fix50sp1 Component
type SettlDetails struct {
	//NoSettlDetails is a non-required field for SettlDetails.
	NoSettlDetails []NoSettlDetails `fix:"1158,omitempty"`
}

func (m *SettlDetails) SetNoSettlDetails(v []NoSettlDetails) { m.NoSettlDetails = v }
