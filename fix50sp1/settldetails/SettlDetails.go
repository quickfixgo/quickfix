package settldetails

import (
	"github.com/quickfixgo/quickfix/fix50sp1/settlparties"
)

//NoSettlDetails is a repeating group in SettlDetails
type NoSettlDetails struct {
	//SettlObligSource is a non-required field for NoSettlDetails.
	SettlObligSource *string `fix:"1164"`
	//SettlParties Component
	SettlParties settlparties.Component
}

//Component is a fix50sp1 SettlDetails Component
type Component struct {
	//NoSettlDetails is a non-required field for SettlDetails.
	NoSettlDetails []NoSettlDetails `fix:"1158,omitempty"`
}

func New() *Component { return new(Component) }
