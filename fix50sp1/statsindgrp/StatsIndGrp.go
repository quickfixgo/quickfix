package statsindgrp

//NoStatsIndicators is a repeating group in StatsIndGrp
type NoStatsIndicators struct {
	//StatsType is a non-required field for NoStatsIndicators.
	StatsType *int `fix:"1176"`
}

//Component is a fix50sp1 StatsIndGrp Component
type Component struct {
	//NoStatsIndicators is a non-required field for StatsIndGrp.
	NoStatsIndicators []NoStatsIndicators `fix:"1175,omitempty"`
}

func New() *Component { return new(Component) }
