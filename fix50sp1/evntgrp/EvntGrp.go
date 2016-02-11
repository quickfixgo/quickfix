package evntgrp

import (
	"time"
)

//NoEvents is a repeating group in EvntGrp
type NoEvents struct {
	//EventType is a non-required field for NoEvents.
	EventType *int `fix:"865"`
	//EventDate is a non-required field for NoEvents.
	EventDate *string `fix:"866"`
	//EventPx is a non-required field for NoEvents.
	EventPx *float64 `fix:"867"`
	//EventText is a non-required field for NoEvents.
	EventText *string `fix:"868"`
	//EventTime is a non-required field for NoEvents.
	EventTime *time.Time `fix:"1145"`
}

//Component is a fix50sp1 EvntGrp Component
type Component struct {
	//NoEvents is a non-required field for EvntGrp.
	NoEvents []NoEvents `fix:"864,omitempty"`
}

func New() *Component { return new(Component) }
