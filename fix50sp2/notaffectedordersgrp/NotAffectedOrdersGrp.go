package notaffectedordersgrp

//NoNotAffectedOrders is a repeating group in NotAffectedOrdersGrp
type NoNotAffectedOrders struct {
	//NotAffOrigClOrdID is a non-required field for NoNotAffectedOrders.
	NotAffOrigClOrdID *string `fix:"1372"`
	//NotAffectedOrderID is a non-required field for NoNotAffectedOrders.
	NotAffectedOrderID *string `fix:"1371"`
}

//Component is a fix50sp2 NotAffectedOrdersGrp Component
type Component struct {
	//NoNotAffectedOrders is a non-required field for NotAffectedOrdersGrp.
	NoNotAffectedOrders []NoNotAffectedOrders `fix:"1370,omitempty"`
}

func New() *Component { return new(Component) }
