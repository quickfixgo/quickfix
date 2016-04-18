package strmasgnrptgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/strmasgnrptinstrmtgrp"
)

//New returns an initialized StrmAsgnRptGrp instance
func New() *StrmAsgnRptGrp {
	var m StrmAsgnRptGrp
	return &m
}

//NoAsgnReqs is a repeating group in StrmAsgnRptGrp
type NoAsgnReqs struct {
	//Parties is a non-required component for NoAsgnReqs.
	Parties *parties.Parties
	//StrmAsgnRptInstrmtGrp is a non-required component for NoAsgnReqs.
	StrmAsgnRptInstrmtGrp *strmasgnrptinstrmtgrp.StrmAsgnRptInstrmtGrp
}

//NewNoAsgnReqs returns an initialized NoAsgnReqs instance
func NewNoAsgnReqs() *NoAsgnReqs {
	var m NoAsgnReqs
	return &m
}

func (m *NoAsgnReqs) SetParties(v parties.Parties) { m.Parties = &v }
func (m *NoAsgnReqs) SetStrmAsgnRptInstrmtGrp(v strmasgnrptinstrmtgrp.StrmAsgnRptInstrmtGrp) {
	m.StrmAsgnRptInstrmtGrp = &v
}

//StrmAsgnRptGrp is a fix50sp2 Component
type StrmAsgnRptGrp struct {
	//NoAsgnReqs is a non-required field for StrmAsgnRptGrp.
	NoAsgnReqs []NoAsgnReqs `fix:"1499,omitempty"`
}

func (m *StrmAsgnRptGrp) SetNoAsgnReqs(v []NoAsgnReqs) { m.NoAsgnReqs = v }
