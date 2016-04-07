package trdregtimestamps

import (
	"time"
)

func New() *TrdRegTimestamps {
	var m TrdRegTimestamps
	return &m
}

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

func (m *NoTrdRegTimestamps) SetTrdRegTimestamp(v time.Time)    { m.TrdRegTimestamp = &v }
func (m *NoTrdRegTimestamps) SetTrdRegTimestampType(v int)      { m.TrdRegTimestampType = &v }
func (m *NoTrdRegTimestamps) SetTrdRegTimestampOrigin(v string) { m.TrdRegTimestampOrigin = &v }
func (m *NoTrdRegTimestamps) SetDeskType(v string)              { m.DeskType = &v }
func (m *NoTrdRegTimestamps) SetDeskTypeSource(v int)           { m.DeskTypeSource = &v }
func (m *NoTrdRegTimestamps) SetDeskOrderHandlingInst(v string) { m.DeskOrderHandlingInst = &v }

//TrdRegTimestamps is a fix50sp2 Component
type TrdRegTimestamps struct {
	//NoTrdRegTimestamps is a non-required field for TrdRegTimestamps.
	NoTrdRegTimestamps []NoTrdRegTimestamps `fix:"768,omitempty"`
}

func (m *TrdRegTimestamps) SetNoTrdRegTimestamps(v []NoTrdRegTimestamps) { m.NoTrdRegTimestamps = v }
