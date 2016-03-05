package strmasgnreqgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/strmasgnreqinstrmtgrp"
)

//NoAsgnReqs is a repeating group in StrmAsgnReqGrp
type NoAsgnReqs struct {
	//Parties Component
	parties.Parties
	//StrmAsgnReqInstrmtGrp Component
	strmasgnreqinstrmtgrp.StrmAsgnReqInstrmtGrp
}

//StrmAsgnReqGrp is a fix50sp2 Component
type StrmAsgnReqGrp struct {
	//NoAsgnReqs is a non-required field for StrmAsgnReqGrp.
	NoAsgnReqs []NoAsgnReqs `fix:"1499,omitempty"`
}

func (m *StrmAsgnReqGrp) SetNoAsgnReqs(v []NoAsgnReqs) { m.NoAsgnReqs = v }
