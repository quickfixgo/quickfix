package dlvyinstgrp

import (
	"github.com/quickfixgo/quickfix/fix50/settlparties"
)

//NoDlvyInst is a repeating group in DlvyInstGrp
type NoDlvyInst struct {
	//SettlInstSource is a non-required field for NoDlvyInst.
	SettlInstSource *string `fix:"165"`
	//DlvyInstType is a non-required field for NoDlvyInst.
	DlvyInstType *string `fix:"787"`
	//SettlParties Component
	SettlParties settlparties.Component
}

//Component is a fix50 DlvyInstGrp Component
type Component struct {
	//NoDlvyInst is a non-required field for DlvyInstGrp.
	NoDlvyInst []NoDlvyInst `fix:"85,omitempty"`
}

func New() *Component { return new(Component) }
