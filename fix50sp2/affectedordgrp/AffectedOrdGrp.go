package affectedordgrp

//NoAffectedOrders is a repeating group in AffectedOrdGrp
type NoAffectedOrders struct {
	//OrigClOrdID is a non-required field for NoAffectedOrders.
	OrigClOrdID *string `fix:"41"`
	//AffectedOrderID is a non-required field for NoAffectedOrders.
	AffectedOrderID *string `fix:"535"`
	//AffectedSecondaryOrderID is a non-required field for NoAffectedOrders.
	AffectedSecondaryOrderID *string `fix:"536"`
}

//Component is a fix50sp2 AffectedOrdGrp Component
type Component struct {
	//NoAffectedOrders is a non-required field for AffectedOrdGrp.
	NoAffectedOrders []NoAffectedOrders `fix:"534,omitempty"`
}

func New() *Component { return new(Component) }
