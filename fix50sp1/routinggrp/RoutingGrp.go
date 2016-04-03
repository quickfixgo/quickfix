package routinggrp

func New() *RoutingGrp {
	var m RoutingGrp
	return &m
}

//NoRoutingIDs is a repeating group in RoutingGrp
type NoRoutingIDs struct {
	//RoutingType is a non-required field for NoRoutingIDs.
	RoutingType *int `fix:"216"`
	//RoutingID is a non-required field for NoRoutingIDs.
	RoutingID *string `fix:"217"`
}

//RoutingGrp is a fix50sp1 Component
type RoutingGrp struct {
	//NoRoutingIDs is a non-required field for RoutingGrp.
	NoRoutingIDs []NoRoutingIDs `fix:"215,omitempty"`
}

func (m *RoutingGrp) SetNoRoutingIDs(v []NoRoutingIDs) { m.NoRoutingIDs = v }
