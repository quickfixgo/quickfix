package trdregtimestamps

import (
	"time"
)

//New returns an initialized TrdRegTimestamps instance
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
}

//NewNoTrdRegTimestamps returns an initialized NoTrdRegTimestamps instance
func NewNoTrdRegTimestamps() *NoTrdRegTimestamps {
	var m NoTrdRegTimestamps
	return &m
}

func (m *NoTrdRegTimestamps) SetTrdRegTimestamp(v time.Time)    { m.TrdRegTimestamp = &v }
func (m *NoTrdRegTimestamps) SetTrdRegTimestampType(v int)      { m.TrdRegTimestampType = &v }
func (m *NoTrdRegTimestamps) SetTrdRegTimestampOrigin(v string) { m.TrdRegTimestampOrigin = &v }

//TrdRegTimestamps is a fix44 Component
type TrdRegTimestamps struct {
	//NoTrdRegTimestamps is a non-required field for TrdRegTimestamps.
	NoTrdRegTimestamps []NoTrdRegTimestamps `fix:"768,omitempty"`
}

func (m *TrdRegTimestamps) SetNoTrdRegTimestamps(v []NoTrdRegTimestamps) { m.NoTrdRegTimestamps = v }
