package sidetrdregts

import (
	"time"
)

//New returns an initialized SideTrdRegTS instance
func New() *SideTrdRegTS {
	var m SideTrdRegTS
	return &m
}

//NoSideTrdRegTS is a repeating group in SideTrdRegTS
type NoSideTrdRegTS struct {
	//SideTrdRegTimestamp is a non-required field for NoSideTrdRegTS.
	SideTrdRegTimestamp *time.Time `fix:"1012"`
	//SideTrdRegTimestampType is a non-required field for NoSideTrdRegTS.
	SideTrdRegTimestampType *int `fix:"1013"`
	//SideTrdRegTimestampSrc is a non-required field for NoSideTrdRegTS.
	SideTrdRegTimestampSrc *string `fix:"1014"`
}

//NewNoSideTrdRegTS returns an initialized NoSideTrdRegTS instance
func NewNoSideTrdRegTS() *NoSideTrdRegTS {
	var m NoSideTrdRegTS
	return &m
}

func (m *NoSideTrdRegTS) SetSideTrdRegTimestamp(v time.Time) { m.SideTrdRegTimestamp = &v }
func (m *NoSideTrdRegTS) SetSideTrdRegTimestampType(v int)   { m.SideTrdRegTimestampType = &v }
func (m *NoSideTrdRegTS) SetSideTrdRegTimestampSrc(v string) { m.SideTrdRegTimestampSrc = &v }

//SideTrdRegTS is a fix50sp1 Component
type SideTrdRegTS struct {
	//NoSideTrdRegTS is a non-required field for SideTrdRegTS.
	NoSideTrdRegTS []NoSideTrdRegTS `fix:"1016,omitempty"`
}

func (m *SideTrdRegTS) SetNoSideTrdRegTS(v []NoSideTrdRegTS) { m.NoSideTrdRegTS = v }
