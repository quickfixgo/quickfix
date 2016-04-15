package evntgrp

import (
	"time"
)

func New() *EvntGrp {
	var m EvntGrp
	return &m
}

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

func (m *NoEvents) SetEventType(v int)       { m.EventType = &v }
func (m *NoEvents) SetEventDate(v string)    { m.EventDate = &v }
func (m *NoEvents) SetEventPx(v float64)     { m.EventPx = &v }
func (m *NoEvents) SetEventText(v string)    { m.EventText = &v }
func (m *NoEvents) SetEventTime(v time.Time) { m.EventTime = &v }

//EvntGrp is a fix50sp2 Component
type EvntGrp struct {
	//NoEvents is a non-required field for EvntGrp.
	NoEvents []NoEvents `fix:"864,omitempty"`
}

func (m *EvntGrp) SetNoEvents(v []NoEvents) { m.NoEvents = v }
