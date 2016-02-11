package mdreqgrp

//NoMDEntryTypes is a repeating group in MDReqGrp
type NoMDEntryTypes struct {
	//MDEntryType is a required field for NoMDEntryTypes.
	MDEntryType string `fix:"269"`
}

//Component is a fix50sp2 MDReqGrp Component
type Component struct {
	//NoMDEntryTypes is a required field for MDReqGrp.
	NoMDEntryTypes []NoMDEntryTypes `fix:"267"`
}

func New() *Component { return new(Component) }
