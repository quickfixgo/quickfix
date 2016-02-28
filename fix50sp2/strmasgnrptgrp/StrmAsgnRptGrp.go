package strmasgnrptgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/strmasgnrptinstrmtgrp"
)

//NoAsgnReqs is a repeating group in StrmAsgnRptGrp
type NoAsgnReqs struct {
	//Parties Component
	Parties parties.Component
	//StrmAsgnRptInstrmtGrp Component
	StrmAsgnRptInstrmtGrp strmasgnrptinstrmtgrp.Component
}

//Component is a fix50sp2 StrmAsgnRptGrp Component
type Component struct {
	//NoAsgnReqs is a non-required field for StrmAsgnRptGrp.
	NoAsgnReqs []NoAsgnReqs `fix:"1499,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoAsgnReqs(v []NoAsgnReqs) { m.NoAsgnReqs = v }
