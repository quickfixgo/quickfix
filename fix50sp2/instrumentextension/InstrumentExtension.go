package instrumentextension

import (
	"github.com/quickfixgo/quickfix/fix50sp2/attrbgrp"
)

func New() *InstrumentExtension {
	var m InstrumentExtension
	return &m
}

//InstrumentExtension is a fix50sp2 Component
type InstrumentExtension struct {
	//DeliveryForm is a non-required field for InstrumentExtension.
	DeliveryForm *int `fix:"668"`
	//PctAtRisk is a non-required field for InstrumentExtension.
	PctAtRisk *float64 `fix:"869"`
	//AttrbGrp is a non-required component for InstrumentExtension.
	AttrbGrp *attrbgrp.AttrbGrp
}

func (m *InstrumentExtension) SetDeliveryForm(v int)           { m.DeliveryForm = &v }
func (m *InstrumentExtension) SetPctAtRisk(v float64)          { m.PctAtRisk = &v }
func (m *InstrumentExtension) SetAttrbGrp(v attrbgrp.AttrbGrp) { m.AttrbGrp = &v }
