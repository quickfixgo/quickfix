package hopgrp

import (
	"time"
)

//New returns an initialized HopGrp instance
func New() *HopGrp {
	var m HopGrp
	return &m
}

//NoHops is a repeating group in HopGrp
type NoHops struct {
	//HopCompID is a non-required field for NoHops.
	HopCompID *string `fix:"628"`
	//HopSendingTime is a non-required field for NoHops.
	HopSendingTime *time.Time `fix:"629"`
	//HopRefID is a non-required field for NoHops.
	HopRefID *int `fix:"630"`
}

//NewNoHops returns an initialized NoHops instance
func NewNoHops() *NoHops {
	var m NoHops
	return &m
}

func (m *NoHops) SetHopCompID(v string)         { m.HopCompID = &v }
func (m *NoHops) SetHopSendingTime(v time.Time) { m.HopSendingTime = &v }
func (m *NoHops) SetHopRefID(v int)             { m.HopRefID = &v }

//HopGrp is a fix50sp2 Component
type HopGrp struct {
	//NoHops is a non-required field for HopGrp.
	NoHops []NoHops `fix:"627,omitempty"`
}

func (m *HopGrp) SetNoHops(v []NoHops) { m.NoHops = v }
