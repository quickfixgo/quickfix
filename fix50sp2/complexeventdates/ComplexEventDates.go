package complexeventdates

import (
	"github.com/quickfixgo/quickfix/fix50sp2/complexeventtimes"
	"time"
)

//New returns an initialized ComplexEventDates instance
func New() *ComplexEventDates {
	var m ComplexEventDates
	return &m
}

//NoComplexEventDates is a repeating group in ComplexEventDates
type NoComplexEventDates struct {
	//ComplexEventStartDate is a non-required field for NoComplexEventDates.
	ComplexEventStartDate *time.Time `fix:"1492"`
	//ComplexEventEndDate is a non-required field for NoComplexEventDates.
	ComplexEventEndDate *time.Time `fix:"1493"`
	//ComplexEventTimes is a non-required component for NoComplexEventDates.
	ComplexEventTimes *complexeventtimes.ComplexEventTimes
}

//NewNoComplexEventDates returns an initialized NoComplexEventDates instance
func NewNoComplexEventDates() *NoComplexEventDates {
	var m NoComplexEventDates
	return &m
}

func (m *NoComplexEventDates) SetComplexEventStartDate(v time.Time) { m.ComplexEventStartDate = &v }
func (m *NoComplexEventDates) SetComplexEventEndDate(v time.Time)   { m.ComplexEventEndDate = &v }
func (m *NoComplexEventDates) SetComplexEventTimes(v complexeventtimes.ComplexEventTimes) {
	m.ComplexEventTimes = &v
}

//ComplexEventDates is a fix50sp2 Component
type ComplexEventDates struct {
	//NoComplexEventDates is a non-required field for ComplexEventDates.
	NoComplexEventDates []NoComplexEventDates `fix:"1491,omitempty"`
}

func (m *ComplexEventDates) SetNoComplexEventDates(v []NoComplexEventDates) { m.NoComplexEventDates = v }
