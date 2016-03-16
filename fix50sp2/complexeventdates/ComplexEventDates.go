package complexeventdates

import (
	"github.com/quickfixgo/quickfix/fix50sp2/complexeventtimes"
	"time"
)

//NoComplexEventDates is a repeating group in ComplexEventDates
type NoComplexEventDates struct {
	//ComplexEventStartDate is a non-required field for NoComplexEventDates.
	ComplexEventStartDate *time.Time `fix:"1492"`
	//ComplexEventEndDate is a non-required field for NoComplexEventDates.
	ComplexEventEndDate *time.Time `fix:"1493"`
	//ComplexEventTimes is a non-required component for NoComplexEventDates.
	ComplexEventTimes *complexeventtimes.ComplexEventTimes
}

//ComplexEventDates is a fix50sp2 Component
type ComplexEventDates struct {
	//NoComplexEventDates is a non-required field for ComplexEventDates.
	NoComplexEventDates []NoComplexEventDates `fix:"1491,omitempty"`
}

func (m *ComplexEventDates) SetNoComplexEventDates(v []NoComplexEventDates) { m.NoComplexEventDates = v }
