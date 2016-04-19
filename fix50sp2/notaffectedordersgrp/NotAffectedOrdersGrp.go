package notaffectedordersgrp

//New returns an initialized NotAffectedOrdersGrp instance
func New() *NotAffectedOrdersGrp {
	var m NotAffectedOrdersGrp
	return &m
}

//NoNotAffectedOrders is a repeating group in NotAffectedOrdersGrp
type NoNotAffectedOrders struct {
	//NotAffOrigClOrdID is a non-required field for NoNotAffectedOrders.
	NotAffOrigClOrdID *string `fix:"1372"`
	//NotAffectedOrderID is a non-required field for NoNotAffectedOrders.
	NotAffectedOrderID *string `fix:"1371"`
}

//NewNoNotAffectedOrders returns an initialized NoNotAffectedOrders instance
func NewNoNotAffectedOrders() *NoNotAffectedOrders {
	var m NoNotAffectedOrders
	return &m
}

func (m *NoNotAffectedOrders) SetNotAffOrigClOrdID(v string)  { m.NotAffOrigClOrdID = &v }
func (m *NoNotAffectedOrders) SetNotAffectedOrderID(v string) { m.NotAffectedOrderID = &v }

//NotAffectedOrdersGrp is a fix50sp2 Component
type NotAffectedOrdersGrp struct {
	//NoNotAffectedOrders is a non-required field for NotAffectedOrdersGrp.
	NoNotAffectedOrders []NoNotAffectedOrders `fix:"1370,omitempty"`
}

func (m *NotAffectedOrdersGrp) SetNoNotAffectedOrders(v []NoNotAffectedOrders) {
	m.NoNotAffectedOrders = v
}
