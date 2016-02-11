package ordliststatgrp

//NoOrders is a repeating group in OrdListStatGrp
type NoOrders struct {
	//ClOrdID is a non-required field for NoOrders.
	ClOrdID *string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoOrders.
	SecondaryClOrdID *string `fix:"526"`
	//CumQty is a required field for NoOrders.
	CumQty float64 `fix:"14"`
	//OrdStatus is a required field for NoOrders.
	OrdStatus string `fix:"39"`
	//WorkingIndicator is a non-required field for NoOrders.
	WorkingIndicator *bool `fix:"636"`
	//LeavesQty is a required field for NoOrders.
	LeavesQty float64 `fix:"151"`
	//CxlQty is a required field for NoOrders.
	CxlQty float64 `fix:"84"`
	//AvgPx is a required field for NoOrders.
	AvgPx float64 `fix:"6"`
	//OrdRejReason is a non-required field for NoOrders.
	OrdRejReason *int `fix:"103"`
	//Text is a non-required field for NoOrders.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoOrders.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoOrders.
	EncodedText *string `fix:"355"`
	//OrderID is a non-required field for NoOrders.
	OrderID *string `fix:"37"`
}

//Component is a fix50sp1 OrdListStatGrp Component
type Component struct {
	//NoOrders is a required field for OrdListStatGrp.
	NoOrders []NoOrders `fix:"73"`
}

func New() *Component { return new(Component) }
