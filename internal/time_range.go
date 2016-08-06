package internal

import (
	"errors"
	"time"
)

//UTCTimeOnly represents the time of day in UTC
type UTCTimeOnly struct {
	seconds int
}

const shortForm = "15:04:05"

var errParseTime = errors.New("Time must be in the format HH:MM:SS")

//NewTime returns a newly initialized UTCTimeOnly
func NewTime(hour, minute, second int) UTCTimeOnly {
	return UTCTimeOnly{second + minute*60 + hour*60*60}
}

//ParseTime parses a UTCTimeOnly from a string in the format HH:MM:SS
func ParseTime(str string) (UTCTimeOnly, error) {
	t, err := time.Parse(shortForm, str)
	if err != nil {
		return UTCTimeOnly{}, errParseTime
	}

	return NewTime(t.Clock()), nil
}

//TimeRange represents a time band from StartTime to EndTime
type TimeRange struct {
	StartTime, EndTime UTCTimeOnly
}

//IsInRange returns true if time t is within in the time range
func (r *TimeRange) IsInRange(t time.Time) bool {
	if r == nil {
		return true
	}

	ts := NewTime(t.UTC().Clock()).seconds
	start := r.StartTime.seconds
	end := r.EndTime.seconds

	if start > end {
		return !(end < ts && ts < start)
	}
	return start <= ts && ts <= end
}
