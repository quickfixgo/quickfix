package field

import (
	"github.com/cbusbey/quickfixgo/tag"
	"time"
)

//Container for utctimestamp, knows part of FieldConverter interface
type UTCTimestampValue struct {
	Value    time.Time
	NoMillis bool
}

const (
	utcTimestampFormat         = "20060102-15:04:05.000"
	utcTimestampNoMillisFormat = "20060102-15:04:05"
)

func (f *UTCTimestampValue) ConvertValueFromBytes(bytes []byte) (err error) {
	//with millisecs
	if f.Value, err = time.Parse(utcTimestampFormat, string(bytes)); err == nil {
		return
	}

	//w/o millisecs
	f.Value, err = time.Parse(utcTimestampNoMillisFormat, string(bytes))

	return
}

func (f UTCTimestampValue) ConvertValueToBytes() []byte {
	if f.NoMillis {
		return []byte(f.Value.UTC().Format(utcTimestampNoMillisFormat))
	}

	return []byte(f.Value.UTC().Format(utcTimestampFormat))
}

//Generic utctimestamp Field Type. Implements FieldConverter
type UTCTimestampField struct {
	FieldTag tag.Tag
	UTCTimestampValue
}

func (f UTCTimestampField) Tag() tag.Tag {
	return f.FieldTag
}

func NewUTCTimestampField(tag tag.Tag, value time.Time) *UTCTimestampField {
	field := UTCTimestampField{FieldTag: tag}
	field.Value = value

	return &field
}

func NewUTCTimestampFieldNoMillis(tag tag.Tag, value time.Time) *UTCTimestampField {
	field := UTCTimestampField{FieldTag: tag}
	field.Value = value
	field.NoMillis = true

	return &field
}
