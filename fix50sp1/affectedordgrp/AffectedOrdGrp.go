package affectedordgrp

//New returns an initialized AffectedOrdGrp instance
func New() *AffectedOrdGrp {
	var m AffectedOrdGrp
	return &m
}

//NoAffectedOrders is a repeating group in AffectedOrdGrp
type NoAffectedOrders struct {
	//OrigClOrdID is a non-required field for NoAffectedOrders.
	OrigClOrdID *string `fix:"41"`
	//AffectedOrderID is a non-required field for NoAffectedOrders.
	AffectedOrderID *string `fix:"535"`
	//AffectedSecondaryOrderID is a non-required field for NoAffectedOrders.
	AffectedSecondaryOrderID *string `fix:"536"`
}

//NewNoAffectedOrders returns an initialized NoAffectedOrders instance
func NewNoAffectedOrders() *NoAffectedOrders {
	var m NoAffectedOrders
	return &m
}

func (m *NoAffectedOrders) SetOrigClOrdID(v string)              { m.OrigClOrdID = &v }
func (m *NoAffectedOrders) SetAffectedOrderID(v string)          { m.AffectedOrderID = &v }
func (m *NoAffectedOrders) SetAffectedSecondaryOrderID(v string) { m.AffectedSecondaryOrderID = &v }

//AffectedOrdGrp is a fix50sp1 Component
type AffectedOrdGrp struct {
	//NoAffectedOrders is a non-required field for AffectedOrdGrp.
	NoAffectedOrders []NoAffectedOrders `fix:"534,omitempty"`
}

func (m *AffectedOrdGrp) SetNoAffectedOrders(v []NoAffectedOrders) { m.NoAffectedOrders = v }
