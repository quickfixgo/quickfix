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
	//DeskType is a non-required field for NoTrdRegTimestamps.
	DeskType *string `fix:"1033"`
	//DeskTypeSource is a non-required field for NoTrdRegTimestamps.
	DeskTypeSource *int `fix:"1034"`
	//DeskOrderHandlingInst is a non-required field for NoTrdRegTimestamps.
	DeskOrderHandlingInst *string `fix:"1035"`
}

//Component is a fix50 TrdRegTimestamps Component
type Component struct {
	//NoTrdRegTimestamps is a non-required field for TrdRegTimestamps.
	NoTrdRegTimestamps []NoTrdRegTimestamps `fix:"768,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoTrdRegTimestamps(v []NoTrdRegTimestamps) { m.NoTrdRegTimestamps = v }
