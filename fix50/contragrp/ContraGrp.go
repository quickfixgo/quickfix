package contragrp

import (
	"time"
)

func New() *ContraGrp {
	var m ContraGrp
	return &m
}

//NoContraBrokers is a repeating group in ContraGrp
type NoContraBrokers struct {
	//ContraBroker is a non-required field for NoContraBrokers.
	ContraBroker *string `fix:"375"`
	//ContraTrader is a non-required field for NoContraBrokers.
	ContraTrader *string `fix:"337"`
	//ContraTradeQty is a non-required field for NoContraBrokers.
	ContraTradeQty *float64 `fix:"437"`
	//ContraTradeTime is a non-required field for NoContraBrokers.
	ContraTradeTime *time.Time `fix:"438"`
	//ContraLegRefID is a non-required field for NoContraBrokers.
	ContraLegRefID *string `fix:"655"`
}

func (m *NoContraBrokers) SetContraBroker(v string)       { m.ContraBroker = &v }
func (m *NoContraBrokers) SetContraTrader(v string)       { m.ContraTrader = &v }
func (m *NoContraBrokers) SetContraTradeQty(v float64)    { m.ContraTradeQty = &v }
func (m *NoContraBrokers) SetContraTradeTime(v time.Time) { m.ContraTradeTime = &v }
func (m *NoContraBrokers) SetContraLegRefID(v string)     { m.ContraLegRefID = &v }

//ContraGrp is a fix50 Component
type ContraGrp struct {
	//NoContraBrokers is a non-required field for ContraGrp.
	NoContraBrokers []NoContraBrokers `fix:"382,omitempty"`
}

func (m *ContraGrp) SetNoContraBrokers(v []NoContraBrokers) { m.NoContraBrokers = v }
