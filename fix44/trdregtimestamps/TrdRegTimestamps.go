package trdregtimestamps

import (
	"time"
)

//NoTrdRegTimestamps is a repeating group in TrdRegTimestamps
type NoTrdRegTimestamps struct {
	//TrdRegTimestamp is a non-required field for NoTrdRegTimestamps.
	TrdRegTimestamp *time.Time `fix:"769"`
	//TrdRegTimestampType is a non-required field for NoTrdRegTimestamps.
	TrdRegTimestampType *int `fix:"770"`
	//TrdRegTimestampOrigin is a non-required field for NoTrdRegTimestamps.
	TrdRegTimestampOrigin *string `fix:"771"`
}

//Component is a fix44 TrdRegTimestamps Component
type Component struct {
	//NoTrdRegTimestamps is a non-required field for TrdRegTimestamps.
	NoTrdRegTimestamps []NoTrdRegTimestamps `fix:"768,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoTrdRegTimestamps(v []NoTrdRegTimestamps) { m.NoTrdRegTimestamps = v }
