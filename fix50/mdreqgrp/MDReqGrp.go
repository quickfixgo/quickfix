package mdreqgrp

//New returns an initialized MDReqGrp instance
func New(nomdentrytypes []NoMDEntryTypes) *MDReqGrp {
	var m MDReqGrp
	m.SetNoMDEntryTypes(nomdentrytypes)
	return &m
}

//NoMDEntryTypes is a repeating group in MDReqGrp
type NoMDEntryTypes struct {
	//MDEntryType is a required field for NoMDEntryTypes.
	MDEntryType string `fix:"269"`
}

//NewNoMDEntryTypes returns an initialized NoMDEntryTypes instance
func NewNoMDEntryTypes(mdentrytype string) *NoMDEntryTypes {
	var m NoMDEntryTypes
	m.SetMDEntryType(mdentrytype)
	return &m
}

func (m *NoMDEntryTypes) SetMDEntryType(v string) { m.MDEntryType = v }

//MDReqGrp is a fix50 Component
type MDReqGrp struct {
	//NoMDEntryTypes is a required field for MDReqGrp.
	NoMDEntryTypes []NoMDEntryTypes `fix:"267"`
}

func (m *MDReqGrp) SetNoMDEntryTypes(v []NoMDEntryTypes) { m.NoMDEntryTypes = v }
