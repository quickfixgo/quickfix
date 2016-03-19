package hopgrp

import (
	"time"
)

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

//HopGrp is a fixt11 Component
type HopGrp struct {
	//NoHops is a non-required field for HopGrp.
	NoHops []NoHops `fix:"627,omitempty"`
}

func (m *HopGrp) SetNoHops(v []NoHops) { m.NoHops = v }
