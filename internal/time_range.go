package internal

import (
	"errors"
	"time"
)

//TimeOfDay represents the time of day
type TimeOfDay struct {
	hour, minute, second int
	d                    time.Duration
}

const shortForm = "15:04:05"

var errParseTime = errors.New("Time must be in the format HH:MM:SS")

//NewTimeOfDay returns a newly initialized TimeOfDay
func NewTimeOfDay(hour, minute, second int) TimeOfDay {
	d := time.Duration(second)*time.Second +
		time.Duration(minute)*time.Minute +
		time.Duration(hour)*time.Hour

	return TimeOfDay{hour: hour, minute: minute, second: second, d: d}
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
	startDay, endDay   *time.Weekday
	loc                *time.Location
}

//NewUTCTimeRange returns a time range in UTC
func NewUTCTimeRange(start, end TimeOfDay) *TimeRange {
	return NewTimeRangeInLocation(start, end, time.UTC)
}

//NewTimeRangeInLocation returns a time range in a given location
func NewTimeRangeInLocation(start, end TimeOfDay, loc *time.Location) *TimeRange {
	if loc == nil {
		panic("time: missing Location in call to NewTimeRangeInLocation")
	}

	return &TimeRange{startTime: start, endTime: end, loc: loc}
}

//NewUTCWeekRange returns a weekly TimeRange
func NewUTCWeekRange(startTime, endTime TimeOfDay, startDay, endDay time.Weekday) *TimeRange {
	return NewWeekRangeInLocation(startTime, endTime, startDay, endDay, time.UTC)
}

//NewWeekRangeInLocation returns a time range in a given location
func NewWeekRangeInLocation(startTime, endTime TimeOfDay, startDay, endDay time.Weekday, loc *time.Location) *TimeRange {
	r := NewTimeRangeInLocation(startTime, endTime, loc)
	r.startDay = &startDay
	r.endDay = &endDay

	return r
}

func (r *TimeRange) isInTimeRange(t time.Time) bool {
	t = t.In(r.loc)
	ts := NewTimeOfDay(t.Clock()).d

	if r.startTime.d < r.endTime.d {
		return r.startTime.d <= ts && ts <= r.endTime.d
	}

	return !(r.endTime.d < ts && ts < r.startTime.d)
}

func (r *TimeRange) isInWeekRange(t time.Time) bool {
	t = t.In(r.loc)
	day := t.Weekday()

	if *r.startDay == *r.endDay {
		if day == *r.startDay {
			return r.isInTimeRange(t)
		}

		if r.startTime.d < r.endTime.d {
			return false
		}

		return true
	}

	switch {
	case *r.startDay < *r.endDay:
		if day < *r.startDay || *r.endDay < day {
			return false
		}

	default:
		if *r.endDay < day && day < *r.startDay {
			return false
		}
	}

	timeOfDay := NewTimeOfDay(t.Clock())

	if day == *r.startDay {
		return timeOfDay.d >= r.startTime.d
	}

	if day == *r.endDay {
		return timeOfDay.d <= r.endTime.d
	}

	return true
}

//IsInRange returns true if time t is within in the time range
func (r *TimeRange) IsInRange(t time.Time) bool {
	if r == nil {
		return true
	}

	if r.startDay != nil {
		return r.isInWeekRange(t)
	}

	return r.isInTimeRange(t)
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
	dayOffset := 0

	if r.endDay == nil {
		if r.startTime.d >= r.endTime.d && t1Time.d >= r.startTime.d {
			dayOffset = 1
		}
	} else {
		switch {
		case *r.endDay < t1.Weekday():
			dayOffset = 7 + int(*(r.endDay)-t1.Weekday())

		case t1.Weekday() == *r.endDay:
			if r.endTime.d <= t1Time.d {
				dayOffset = 7
			}

		default:
			dayOffset = int(*(r.endDay) - t1.Weekday())
		}
	}

	sessionEnd := time.Date(t1.Year(), t1.Month(), t1.Day(), r.endTime.hour, r.endTime.minute, r.endTime.second, 0, r.loc)
	sessionEnd = sessionEnd.AddDate(0, 0, dayOffset)

	return t2.Before(sessionEnd)
}
