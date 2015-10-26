package quickfix

import (
	"time"
)

//UTCTimestampValue is a Container for utctimestamp, implements FieldValue
type UTCTimestampValue struct {
	Value    time.Time
	NoMillis bool
}

const (
	utcTimestampFormat         = "20060102-15:04:05.000"
	utcTimestampNoMillisFormat = "20060102-15:04:05"
)

func (f *UTCTimestampValue) Read(bytes []byte) (err error) {
	//with millisecs
	if f.Value, err = time.Parse(utcTimestampFormat, string(bytes)); err == nil {
		return
	}

	//w/o millisecs
	f.Value, err = time.Parse(utcTimestampNoMillisFormat, string(bytes))

	return
}

func (f UTCTimestampValue) Write() []byte {
	if f.NoMillis {
		return []byte(f.Value.UTC().Format(utcTimestampNoMillisFormat))
	}

	return []byte(f.Value.UTC().Format(utcTimestampFormat))
}

//UTCTimestampField is a generic utctimestamp Field Type. Implements Field
type UTCTimestampField struct {
	tagContainer
	UTCTimestampValue
}

func NewUTCTimestampField(tag Tag, value time.Time) *UTCTimestampField {
	var field UTCTimestampField
	field.tag = tag
	field.Value = value

	return &field
}

func NewUTCTimestampFieldNoMillis(tag Tag, value time.Time) *UTCTimestampField {
	var field UTCTimestampField
	field.tag = tag
	field.Value = value
	field.NoMillis = true

	return &field
}
