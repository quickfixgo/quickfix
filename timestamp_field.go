package quickfixgo

import(
  "time"
)

const (
  utcTimestampFormat = "20060102-15:04:05.000"
  utcTimestampNoMillisFormat = "20060102-15:04:05"
)

type UTCTimestampField interface {
  Field

  //field value as timestamp
  UTCTimestampValue() time.Time
}

type utcTimestampFieldBase struct {
  Field
  timeValue time.Time
}

func (f *utcTimestampFieldBase) UTCTimestampValue() time.Time { return f.timeValue}

//returns utc timestamp field in the format YYYYMMDD-HH:MM:SS.sss 
/*func NewUTCTimestampField(tag Tag, value time.Time) UTCTimestampField {
  return &utcTimestampFieldBase{New(tag, value.UTC().Format(utcTimestampFormat)),value}
}

//returns utc timestamp field in the format YYYYMMDD-HH:MM:SS
func NewUTCTimestampFieldNoMillis(tag Tag, value time.Time) UTCTimestampField {
  return &utcTimestampFieldBase{New(tag,value.UTC().Format(utcTimestampNoMillisFormat)),value}
}


//converts a generic field to a utc timestamp field
//check error for convert errors
func ToUTCTimestampField(f Field) (UTCTimestampField, error) {
  //with millisecs
  value, err:=time.Parse(utcTimestampFormat,f.Value())
  if err==nil {
    return NewUTCTimestampField(f.Tag(), value), nil
  }

  //w/o millisecs
  value, err=time.Parse(utcTimestampNoMillisFormat,f.Value())
  if err==nil {
    return NewUTCTimestampFieldNoMillis(f.Tag(), value), nil
  }

  return nil, err
}
*/
