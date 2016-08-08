package internal

import (
	"errors"
	"time"
)

//UTCTimeOfDay represents the time of day in UTC
type UTCTimeOfDay struct {
	hour, minute, second int
}

func (t UTCTimeOfDay) duration() time.Duration {
	return time.Duration(t.second)*time.Second +
		time.Duration(t.minute)*time.Minute +
		time.Duration(t.hour)*time.Hour
}

const shortForm = "15:04:05"

var errParseTime = errors.New("Time must be in the format HH:MM:SS")

//NewTimeOfDay returns a newly initialized UTCTimeOfDay
func NewTimeOfDay(hour, minute, second int) UTCTimeOfDay {
	return UTCTimeOfDay{hour, minute, second}
}

//ParseTimeOfDay parses a UTCTimeOfDay from a string in the format HH:MM:SS
func ParseTimeOfDay(str string) (UTCTimeOfDay, error) {
	t, err := time.Parse(shortForm, str)
	if err != nil {
		return UTCTimeOfDay{}, errParseTime
	}

	return NewTimeOfDay(t.Clock()), nil
}

//TimeRange represents a time band from StartTime to EndTime
type TimeRange struct {
	StartTime, EndTime UTCTimeOfDay
}

//IsInRange returns true if time t is within in the time range
func (r *TimeRange) IsInRange(t time.Time) bool {
	if r == nil {
		return true
	}

	ts := NewTimeOfDay(t.UTC().Clock()).duration()
	start := r.StartTime.duration()
	end := r.EndTime.duration()

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

	t1Time := NewTimeOfDay(t1.UTC().Clock())
	sessionEnd := time.Date(t1.Year(), t1.Month(), t1.Day(), r.EndTime.hour, r.EndTime.minute, r.EndTime.second, 0, time.UTC)
	if r.StartTime.duration() >= r.EndTime.duration() && t1Time.duration() > r.StartTime.duration() {
		sessionEnd = sessionEnd.AddDate(0, 0, 1)
	}

	return t2.Before(sessionEnd)
}
