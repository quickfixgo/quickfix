package statsindgrp

//New returns an initialized StatsIndGrp instance
func New() *StatsIndGrp {
	var m StatsIndGrp
	return &m
}

//NoStatsIndicators is a repeating group in StatsIndGrp
type NoStatsIndicators struct {
	//StatsType is a non-required field for NoStatsIndicators.
	StatsType *int `fix:"1176"`
}

//NewNoStatsIndicators returns an initialized NoStatsIndicators instance
func NewNoStatsIndicators() *NoStatsIndicators {
	var m NoStatsIndicators
	return &m
}

func (m *NoStatsIndicators) SetStatsType(v int) { m.StatsType = &v }

//StatsIndGrp is a fix50sp2 Component
type StatsIndGrp struct {
	//NoStatsIndicators is a non-required field for StatsIndGrp.
	NoStatsIndicators []NoStatsIndicators `fix:"1175,omitempty"`
}

func (m *StatsIndGrp) SetNoStatsIndicators(v []NoStatsIndicators) { m.NoStatsIndicators = v }
