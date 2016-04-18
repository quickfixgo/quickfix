package legpreallocgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/nestedparties2"
)

//New returns an initialized LegPreAllocGrp instance
func New() *LegPreAllocGrp {
	var m LegPreAllocGrp
	return &m
}

//NoLegAllocs is a repeating group in LegPreAllocGrp
type NoLegAllocs struct {
	//LegAllocAccount is a non-required field for NoLegAllocs.
	LegAllocAccount *string `fix:"671"`
	//LegIndividualAllocID is a non-required field for NoLegAllocs.
	LegIndividualAllocID *string `fix:"672"`
	//LegAllocQty is a non-required field for NoLegAllocs.
	LegAllocQty *float64 `fix:"673"`
	//LegAllocAcctIDSource is a non-required field for NoLegAllocs.
	LegAllocAcctIDSource *string `fix:"674"`
	//LegAllocSettlCurrency is a non-required field for NoLegAllocs.
	LegAllocSettlCurrency *string `fix:"1367"`
	//NestedParties2 is a non-required component for NoLegAllocs.
	NestedParties2 *nestedparties2.NestedParties2
}

//NewNoLegAllocs returns an initialized NoLegAllocs instance
func NewNoLegAllocs() *NoLegAllocs {
	var m NoLegAllocs
	return &m
}

func (m *NoLegAllocs) SetLegAllocAccount(v string)                       { m.LegAllocAccount = &v }
func (m *NoLegAllocs) SetLegIndividualAllocID(v string)                  { m.LegIndividualAllocID = &v }
func (m *NoLegAllocs) SetLegAllocQty(v float64)                          { m.LegAllocQty = &v }
func (m *NoLegAllocs) SetLegAllocAcctIDSource(v string)                  { m.LegAllocAcctIDSource = &v }
func (m *NoLegAllocs) SetLegAllocSettlCurrency(v string)                 { m.LegAllocSettlCurrency = &v }
func (m *NoLegAllocs) SetNestedParties2(v nestedparties2.NestedParties2) { m.NestedParties2 = &v }

//LegPreAllocGrp is a fix50sp2 Component
type LegPreAllocGrp struct {
	//NoLegAllocs is a non-required field for LegPreAllocGrp.
	NoLegAllocs []NoLegAllocs `fix:"670,omitempty"`
}

func (m *LegPreAllocGrp) SetNoLegAllocs(v []NoLegAllocs) { m.NoLegAllocs = v }
