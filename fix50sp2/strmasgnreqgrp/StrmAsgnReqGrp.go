package strmasgnreqgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/strmasgnreqinstrmtgrp"
)

//NoAsgnReqs is a repeating group in StrmAsgnReqGrp
type NoAsgnReqs struct {
	//Parties Component
	Parties parties.Component
	//StrmAsgnReqInstrmtGrp Component
	StrmAsgnReqInstrmtGrp strmasgnreqinstrmtgrp.Component
}

//Component is a fix50sp2 StrmAsgnReqGrp Component
type Component struct {
	//NoAsgnReqs is a non-required field for StrmAsgnReqGrp.
	NoAsgnReqs []NoAsgnReqs `fix:"1499,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoAsgnReqs(v []NoAsgnReqs) { m.NoAsgnReqs = v }
