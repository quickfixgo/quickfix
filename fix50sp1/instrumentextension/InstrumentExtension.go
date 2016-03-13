package instrumentextension

import (
	"github.com/quickfixgo/quickfix/fix50sp1/attrbgrp"
)

//InstrumentExtension is a fix50sp1 Component
type InstrumentExtension struct {
	//DeliveryForm is a non-required field for InstrumentExtension.
	DeliveryForm *int `fix:"668"`
	//PctAtRisk is a non-required field for InstrumentExtension.
	PctAtRisk *float64 `fix:"869"`
	//AttrbGrp Component
	attrbgrp.AttrbGrp
}

func (m *InstrumentExtension) SetDeliveryForm(v int)  { m.DeliveryForm = &v }
func (m *InstrumentExtension) SetPctAtRisk(v float64) { m.PctAtRisk = &v }
