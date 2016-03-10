package mdreqgrp

//NoMDEntryTypes is a repeating group in MDReqGrp
type NoMDEntryTypes struct {
	//MDEntryType is a required field for NoMDEntryTypes.
	MDEntryType string `fix:"269"`
}

//MDReqGrp is a fix50sp1 Component
type MDReqGrp struct {
	//NoMDEntryTypes is a required field for MDReqGrp.
	NoMDEntryTypes []NoMDEntryTypes `fix:"267"`
}

func (m *MDReqGrp) SetNoMDEntryTypes(v []NoMDEntryTypes) { m.NoMDEntryTypes = v }
