package strmasgnrptgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/strmasgnrptinstrmtgrp"
)

//NoAsgnReqs is a repeating group in StrmAsgnRptGrp
type NoAsgnReqs struct {
	//Parties Component
	parties.Parties
	//StrmAsgnRptInstrmtGrp Component
	strmasgnrptinstrmtgrp.StrmAsgnRptInstrmtGrp
}

//StrmAsgnRptGrp is a fix50sp2 Component
type StrmAsgnRptGrp struct {
	//NoAsgnReqs is a non-required field for StrmAsgnRptGrp.
	NoAsgnReqs []NoAsgnReqs `fix:"1499,omitempty"`
}

func (m *StrmAsgnRptGrp) SetNoAsgnReqs(v []NoAsgnReqs) { m.NoAsgnReqs = v }
