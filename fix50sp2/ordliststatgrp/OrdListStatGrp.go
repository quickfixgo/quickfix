package ordliststatgrp

func New(noorders []NoOrders) *OrdListStatGrp {
	var m OrdListStatGrp
	m.SetNoOrders(noorders)
	return &m
}

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

func (m *NoOrders) SetClOrdID(v string)          { m.ClOrdID = &v }
func (m *NoOrders) SetSecondaryClOrdID(v string) { m.SecondaryClOrdID = &v }
func (m *NoOrders) SetCumQty(v float64)          { m.CumQty = v }
func (m *NoOrders) SetOrdStatus(v string)        { m.OrdStatus = v }
func (m *NoOrders) SetWorkingIndicator(v bool)   { m.WorkingIndicator = &v }
func (m *NoOrders) SetLeavesQty(v float64)       { m.LeavesQty = v }
func (m *NoOrders) SetCxlQty(v float64)          { m.CxlQty = v }
func (m *NoOrders) SetAvgPx(v float64)           { m.AvgPx = v }
func (m *NoOrders) SetOrdRejReason(v int)        { m.OrdRejReason = &v }
func (m *NoOrders) SetText(v string)             { m.Text = &v }
func (m *NoOrders) SetEncodedTextLen(v int)      { m.EncodedTextLen = &v }
func (m *NoOrders) SetEncodedText(v string)      { m.EncodedText = &v }
func (m *NoOrders) SetOrderID(v string)          { m.OrderID = &v }

//OrdListStatGrp is a fix50sp2 Component
type OrdListStatGrp struct {
	//NoOrders is a required field for OrdListStatGrp.
	NoOrders []NoOrders `fix:"73"`
}

func (m *OrdListStatGrp) SetNoOrders(v []NoOrders) { m.NoOrders = v }
