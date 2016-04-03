package notaffectedordersgrp

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

//NotAffectedOrdersGrp is a fix50sp1 Component
type NotAffectedOrdersGrp struct {
	//NoNotAffectedOrders is a non-required field for NotAffectedOrdersGrp.
	NoNotAffectedOrders []NoNotAffectedOrders `fix:"1370,omitempty"`
}

func (m *NotAffectedOrdersGrp) SetNoNotAffectedOrders(v []NoNotAffectedOrders) {
	m.NoNotAffectedOrders = v
}
