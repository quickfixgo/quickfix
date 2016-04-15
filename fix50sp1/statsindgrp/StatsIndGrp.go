package statsindgrp

func New() *StatsIndGrp {
	var m StatsIndGrp
	return &m
}

//NoStatsIndicators is a repeating group in StatsIndGrp
type NoStatsIndicators struct {
	//StatsType is a non-required field for NoStatsIndicators.
	StatsType *int `fix:"1176"`
}

func (m *NoStatsIndicators) SetStatsType(v int) { m.StatsType = &v }

//StatsIndGrp is a fix50sp1 Component
type StatsIndGrp struct {
	//NoStatsIndicators is a non-required field for StatsIndGrp.
	NoStatsIndicators []NoStatsIndicators `fix:"1175,omitempty"`
}

func (m *StatsIndGrp) SetNoStatsIndicators(v []NoStatsIndicators) { m.NoStatsIndicators = v }
