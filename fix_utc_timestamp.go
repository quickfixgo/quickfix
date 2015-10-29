package quickfix

import (
	"time"
)

//FIXUTCTimestamp is a FIX UTC Timestamp value, implements FieldValue
type FIXUTCTimestamp struct {
	Value    time.Time
	NoMillis bool
}

const (
	utcTimestampFormat         = "20060102-15:04:05.000"
	utcTimestampNoMillisFormat = "20060102-15:04:05"
)

func (f *FIXUTCTimestamp) Read(tv TagValues) (TagValues, error) {
	bytes := tv[0].Value
	var err error
	//with millisecs
	if f.Value, err = time.Parse(utcTimestampFormat, string(bytes)); err == nil {
		return tv[1:], nil
	}

	//w/o millisecs
	if f.Value, err = time.Parse(utcTimestampNoMillisFormat, string(bytes)); err != nil {
		return tv, err
	}

	return tv[1:], nil
}

func (f FIXUTCTimestamp) Write() []byte {
	if f.NoMillis {
		return []byte(f.Value.UTC().Format(utcTimestampNoMillisFormat))
	}

	return []byte(f.Value.UTC().Format(utcTimestampFormat))
}

func (f FIXUTCTimestamp) Clone() FieldValue {
	return &FIXUTCTimestamp{f.Value, f.NoMillis}
}
