package complexevents

import (
	"github.com/quickfixgo/quickfix/fix50sp2/complexeventdates"
)

//New returns an initialized ComplexEvents instance
func New() *ComplexEvents {
	var m ComplexEvents
	return &m
}

//NoComplexEvents is a repeating group in ComplexEvents
type NoComplexEvents struct {
	//ComplexEventType is a non-required field for NoComplexEvents.
	ComplexEventType *int `fix:"1484"`
	//ComplexOptPayoutAmount is a non-required field for NoComplexEvents.
	ComplexOptPayoutAmount *float64 `fix:"1485"`
	//ComplexEventPrice is a non-required field for NoComplexEvents.
	ComplexEventPrice *float64 `fix:"1486"`
	//ComplexEventPriceBoundaryMethod is a non-required field for NoComplexEvents.
	ComplexEventPriceBoundaryMethod *int `fix:"1487"`
	//ComplexEventPriceBoundaryPrecision is a non-required field for NoComplexEvents.
	ComplexEventPriceBoundaryPrecision *float64 `fix:"1488"`
	//ComplexEventPriceTimeType is a non-required field for NoComplexEvents.
	ComplexEventPriceTimeType *int `fix:"1489"`
	//ComplexEventCondition is a non-required field for NoComplexEvents.
	ComplexEventCondition *int `fix:"1490"`
	//ComplexEventDates is a non-required component for NoComplexEvents.
	ComplexEventDates *complexeventdates.ComplexEventDates
}

//NewNoComplexEvents returns an initialized NoComplexEvents instance
func NewNoComplexEvents() *NoComplexEvents {
	var m NoComplexEvents
	return &m
}

func (m *NoComplexEvents) SetComplexEventType(v int)           { m.ComplexEventType = &v }
func (m *NoComplexEvents) SetComplexOptPayoutAmount(v float64) { m.ComplexOptPayoutAmount = &v }
func (m *NoComplexEvents) SetComplexEventPrice(v float64)      { m.ComplexEventPrice = &v }
func (m *NoComplexEvents) SetComplexEventPriceBoundaryMethod(v int) {
	m.ComplexEventPriceBoundaryMethod = &v
}
func (m *NoComplexEvents) SetComplexEventPriceBoundaryPrecision(v float64) {
	m.ComplexEventPriceBoundaryPrecision = &v
}
func (m *NoComplexEvents) SetComplexEventPriceTimeType(v int) { m.ComplexEventPriceTimeType = &v }
func (m *NoComplexEvents) SetComplexEventCondition(v int)     { m.ComplexEventCondition = &v }
func (m *NoComplexEvents) SetComplexEventDates(v complexeventdates.ComplexEventDates) {
	m.ComplexEventDates = &v
}

//ComplexEvents is a fix50sp2 Component
type ComplexEvents struct {
	//NoComplexEvents is a non-required field for ComplexEvents.
	NoComplexEvents []NoComplexEvents `fix:"1483,omitempty"`
}

func (m *ComplexEvents) SetNoComplexEvents(v []NoComplexEvents) { m.NoComplexEvents = v }
