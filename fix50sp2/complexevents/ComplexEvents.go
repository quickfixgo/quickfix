package complexevents

import (
	"github.com/quickfixgo/quickfix/fix50sp2/complexeventdates"
)

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
	//ComplexEventDates Component
	ComplexEventDates complexeventdates.Component
}

//Component is a fix50sp2 ComplexEvents Component
type Component struct {
	//NoComplexEvents is a non-required field for ComplexEvents.
	NoComplexEvents []NoComplexEvents `fix:"1483,omitempty"`
}

func New() *Component { return new(Component) }
