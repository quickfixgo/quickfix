package applidrequestgrp

//NoApplIDs is a repeating group in ApplIDRequestGrp
type NoApplIDs struct {
	//RefApplID is a non-required field for NoApplIDs.
	RefApplID *string `fix:"1355"`
	//ApplBegSeqNum is a non-required field for NoApplIDs.
	ApplBegSeqNum *int `fix:"1182"`
	//ApplEndSeqNum is a non-required field for NoApplIDs.
	ApplEndSeqNum *int `fix:"1183"`
}

//Component is a fix50sp1 ApplIDRequestGrp Component
type Component struct {
	//NoApplIDs is a non-required field for ApplIDRequestGrp.
	NoApplIDs []NoApplIDs `fix:"1351,omitempty"`
}

func New() *Component { return new(Component) }
