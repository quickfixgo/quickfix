package quickfix

import (
	"time"
)

//FIXUTCTimestamp is a FIX UTC Timestamp value, implements FieldValue
type FIXUTCTimestamp struct {
	time.Time
	NoMillis bool
}

const (
	utcTimestampFormat         = "20060102-15:04:05.000"
	utcTimestampNoMillisFormat = "20060102-15:04:05"
)

func (f *FIXUTCTimestamp) Read(bytes []byte) error {
	var err error
	//with millisecs
	if f.Time, err = time.Parse(utcTimestampFormat, string(bytes)); err == nil {
		return nil
	}

	//w/o millisecs
	if f.Time, err = time.Parse(utcTimestampNoMillisFormat, string(bytes)); err != nil {
		return err
	}

	f.NoMillis = true
	return nil
}

func (f FIXUTCTimestamp) Write() []byte {
	if f.NoMillis {
		return []byte(f.UTC().Format(utcTimestampNoMillisFormat))
	}

	return []byte(f.UTC().Format(utcTimestampFormat))
}
