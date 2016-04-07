package dlvyinstgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/settlparties"
)

func New() *DlvyInstGrp {
	var m DlvyInstGrp
	return &m
}

//NoDlvyInst is a repeating group in DlvyInstGrp
type NoDlvyInst struct {
	//SettlInstSource is a non-required field for NoDlvyInst.
	SettlInstSource *string `fix:"165"`
	//DlvyInstType is a non-required field for NoDlvyInst.
	DlvyInstType *string `fix:"787"`
	//SettlParties is a non-required component for NoDlvyInst.
	SettlParties *settlparties.SettlParties
}

func (m *NoDlvyInst) SetSettlInstSource(v string)                 { m.SettlInstSource = &v }
func (m *NoDlvyInst) SetDlvyInstType(v string)                    { m.DlvyInstType = &v }
func (m *NoDlvyInst) SetSettlParties(v settlparties.SettlParties) { m.SettlParties = &v }

//DlvyInstGrp is a fix50sp1 Component
type DlvyInstGrp struct {
	//NoDlvyInst is a non-required field for DlvyInstGrp.
	NoDlvyInst []NoDlvyInst `fix:"85,omitempty"`
}

func (m *DlvyInstGrp) SetNoDlvyInst(v []NoDlvyInst) { m.NoDlvyInst = v }
