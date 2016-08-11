package internal

import (
	"errors"
	"time"
)

//TimeOfDay represents the time of day
type TimeOfDay struct {
	hour, minute, second int
}

func (t TimeOfDay) duration() time.Duration {
	return time.Duration(t.second)*time.Second +
		time.Duration(t.minute)*time.Minute +
		time.Duration(t.hour)*time.Hour
}

const shortForm = "15:04:05"

var errParseTime = errors.New("Time must be in the format HH:MM:SS")

//NewTimeOfDay returns a newly initialized TimeOfDay
func NewTimeOfDay(hour, minute, second int) TimeOfDay {
	return TimeOfDay{hour, minute, second}
}

//ParseTimeOfDay parses a TimeOfDay from a string in the format HH:MM:SS
func ParseTimeOfDay(str string) (TimeOfDay, error) {
	t, err := time.Parse(shortForm, str)
	if err != nil {
		return TimeOfDay{}, errParseTime
	}

	return NewTimeOfDay(t.Clock()), nil
}

//TimeRange represents a time band in a given time zone
type TimeRange struct {
	startTime, endTime TimeOfDay
	loc                *time.Location
}

//NewUTCTimeRange returns a time range in UTC
func NewUTCTimeRange(start, end TimeOfDay) *TimeRange {
	return &TimeRange{start, end, time.UTC}
}

//NewTimeRangeInLocation returns a time range in a given location
func NewTimeRangeInLocation(start, end TimeOfDay, loc *time.Location) *TimeRange {
	if loc == nil {
		panic("time: missing Location in call to NewTimeRangeInLocation")
	}

	return &TimeRange{start, end, loc}
}

//IsInRange returns true if time t is within in the time range
func (r *TimeRange) IsInRange(t time.Time) bool {
	if r == nil {
		return true
	}

	t = t.In(r.loc)
	ts := NewTimeOfDay(t.Clock()).duration()
	start := r.startTime.duration()
	end := r.endTime.duration()

	if start > end {
		return !(end < ts && ts < start)
	}
	return start <= ts && ts <= end
}

//IsInSameRange determines if two points in time are in the same time range
func (r *TimeRange) IsInSameRange(t1, t2 time.Time) bool {
	if r == nil {
		return true
	}

	if !(r.IsInRange(t1) && r.IsInRange(t2)) {
		return false
	}

	if t2.Before(t1) {
		t1, t2 = t2, t1
	}

	t1 = t1.In(r.loc)
	t1Time := NewTimeOfDay(t1.Clock())
	sessionEnd := time.Date(t1.Year(), t1.Month(), t1.Day(), r.endTime.hour, r.endTime.minute, r.endTime.second, 0, r.loc)
	if r.startTime.duration() >= r.endTime.duration() && t1Time.duration() >= r.startTime.duration() {
		sessionEnd = sessionEnd.AddDate(0, 0, 1)
	}

	return t2.Before(sessionEnd)
}
