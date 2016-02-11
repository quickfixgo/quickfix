package ordallocgrp

import (
	"github.com/quickfixgo/quickfix/fix50/nestedparties2"
)

//NoOrders is a repeating group in OrdAllocGrp
type NoOrders struct {
	//ClOrdID is a non-required field for NoOrders.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for NoOrders.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for NoOrders.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for NoOrders.
	SecondaryClOrdID *string `fix:"526"`
	//ListID is a non-required field for NoOrders.
	ListID *string `fix:"66"`
	//NestedParties2 Component
	NestedParties2 nestedparties2.Component
	//OrderQty is a non-required field for NoOrders.
	OrderQty *float64 `fix:"38"`
	//OrderAvgPx is a non-required field for NoOrders.
	OrderAvgPx *float64 `fix:"799"`
	//OrderBookingQty is a non-required field for NoOrders.
	OrderBookingQty *float64 `fix:"800"`
}

//Component is a fix50 OrdAllocGrp Component
type Component struct {
	//NoOrders is a non-required field for OrdAllocGrp.
	NoOrders []NoOrders `fix:"73,omitempty"`
}

func New() *Component { return new(Component) }
