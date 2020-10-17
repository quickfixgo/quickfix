package quickfix_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/quickfixgo/quickfix"
)

func TestFIXUTCTimestampWrite(t *testing.T) {
	ts := time.Date(2016, time.February, 8, 22, 7, 16, 954123123, time.UTC)

	var tests = []struct {
		precision quickfix.TimestampPrecision
		val       []byte
	}{
		{quickfix.Millis, []byte("20160208-22:07:16.954")},
		{quickfix.Seconds, []byte("20160208-22:07:16")},
		{quickfix.Micros, []byte("20160208-22:07:16.954123")},
		{quickfix.Nanos, []byte("20160208-22:07:16.954123123")},
	}

	for _, test := range tests {
		var f quickfix.FIXUTCTimestamp
		f.Time = ts
		f.Precision = test.precision

		b := f.Write()

		if !bytes.Equal(b, test.val) {
			t.Errorf("got %s; want %s", b, test.val)
		}
	}
}

func TestFIXUTCTimestampRead(t *testing.T) {
	var tests = []struct {
		timeStr           string
		expectedTime      time.Time
		expectedPrecision quickfix.TimestampPrecision
	}{
		{"20160208-22:07:16.310", time.Date(2016, time.February, 8, 22, 7, 16, 310000000, time.UTC), quickfix.Millis},
		{"20160208-22:07:16", time.Date(2016, time.February, 8, 22, 7, 16, 0, time.UTC), quickfix.Seconds},
		{"20160208-22:07:16.123455", time.Date(2016, time.February, 8, 22, 7, 16, 123455000, time.UTC), quickfix.Micros},
		{"20160208-22:07:16.954123123", time.Date(2016, time.February, 8, 22, 7, 16, 954123123, time.UTC), quickfix.Nanos},
	}

	for _, test := range tests {
		var f quickfix.FIXUTCTimestamp
		if err := f.Read([]byte(test.timeStr)); err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !f.Time.Equal(test.expectedTime) {
			t.Errorf("For Time expected %v got %v", test.expectedTime, f.Time)
		}

		if f.Precision != test.expectedPrecision {
			t.Errorf("For NoMillis expected %v got %v", test.expectedPrecision, f.Precision)
		}
	}
}
