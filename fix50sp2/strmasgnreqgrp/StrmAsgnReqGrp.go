package strmasgnreqgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/strmasgnreqinstrmtgrp"
)

//New returns an initialized StrmAsgnReqGrp instance
func New() *StrmAsgnReqGrp {
	var m StrmAsgnReqGrp
	return &m
}

//NoAsgnReqs is a repeating group in StrmAsgnReqGrp
type NoAsgnReqs struct {
	//Parties is a non-required component for NoAsgnReqs.
	Parties *parties.Parties
	//StrmAsgnReqInstrmtGrp is a non-required component for NoAsgnReqs.
	StrmAsgnReqInstrmtGrp *strmasgnreqinstrmtgrp.StrmAsgnReqInstrmtGrp
}

//NewNoAsgnReqs returns an initialized NoAsgnReqs instance
func NewNoAsgnReqs() *NoAsgnReqs {
	var m NoAsgnReqs
	return &m
}

func (m *NoAsgnReqs) SetParties(v parties.Parties) { m.Parties = &v }
func (m *NoAsgnReqs) SetStrmAsgnReqInstrmtGrp(v strmasgnreqinstrmtgrp.StrmAsgnReqInstrmtGrp) {
	m.StrmAsgnReqInstrmtGrp = &v
}

//StrmAsgnReqGrp is a fix50sp2 Component
type StrmAsgnReqGrp struct {
	//NoAsgnReqs is a non-required field for StrmAsgnReqGrp.
	NoAsgnReqs []NoAsgnReqs `fix:"1499,omitempty"`
}

func (m *StrmAsgnReqGrp) SetNoAsgnReqs(v []NoAsgnReqs) { m.NoAsgnReqs = v }
