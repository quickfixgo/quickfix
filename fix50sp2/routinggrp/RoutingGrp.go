package routinggrp

//NoRoutingIDs is a repeating group in RoutingGrp
type NoRoutingIDs struct {
	//RoutingType is a non-required field for NoRoutingIDs.
	RoutingType *int `fix:"216"`
	//RoutingID is a non-required field for NoRoutingIDs.
	RoutingID *string `fix:"217"`
}

//Component is a fix50sp2 RoutingGrp Component
type Component struct {
	//NoRoutingIDs is a non-required field for RoutingGrp.
	NoRoutingIDs []NoRoutingIDs `fix:"215,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRoutingIDs(v []NoRoutingIDs) { m.NoRoutingIDs = v }
