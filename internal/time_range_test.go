package internal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTimeOfDay(t *testing.T) {
	to := NewTimeOfDay(12, 34, 4)
	assert.Equal(t, 12, to.hour)
	assert.Equal(t, 34, to.minute)
	assert.Equal(t, 4, to.second)
}

func TestParseTime(t *testing.T) {
	to, err := ParseTimeOfDay("12:34:04")
	assert.Nil(t, err)
	assert.Equal(t, NewTimeOfDay(12, 34, 4), to)

	_, err = ParseTimeOfDay("0:0:0")
	assert.NotNil(t, err)

	_, err = ParseTimeOfDay("00:00")
	assert.NotNil(t, err)

	_, err = ParseTimeOfDay("0000:00")
	assert.NotNil(t, err)
}

func TestNewUTCTimeRange(t *testing.T) {
	r := NewUTCTimeRange(NewTimeOfDay(3, 0, 0), NewTimeOfDay(18, 0, 0))
	assert.Equal(t, NewTimeOfDay(3, 0, 0), r.startTime)
	assert.Equal(t, NewTimeOfDay(18, 0, 0), r.endTime)
	assert.Equal(t, time.UTC, r.loc)
}

func TestNewTimeRangeInLocation(t *testing.T) {
	r := NewTimeRangeInLocation(NewTimeOfDay(3, 0, 0), NewTimeOfDay(18, 0, 0), time.Local)
	assert.Equal(t, NewTimeOfDay(3, 0, 0), r.startTime)
	assert.Equal(t, NewTimeOfDay(18, 0, 0), r.endTime)
	assert.Equal(t, time.Local, r.loc)
}

func TestTimeRangeIsInRange(t *testing.T) {
	start := NewTimeOfDay(3, 0, 0)
	end := NewTimeOfDay(18, 0, 0)

	now := time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCTimeRange(start, end).IsInRange(now))

	now = time.Date(2016, time.August, 10, 18, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCTimeRange(start, end).IsInRange(now))

	now = time.Date(2016, time.August, 10, 2, 0, 0, 0, time.UTC)
	assert.False(t, NewUTCTimeRange(start, end).IsInRange(now))

	now = time.Date(2016, time.August, 10, 19, 0, 0, 0, time.UTC)
	assert.False(t, NewUTCTimeRange(start, end).IsInRange(now))

	now = time.Date(2016, time.August, 10, 18, 0, 1, 0, time.UTC)
	assert.False(t, NewUTCTimeRange(start, end).IsInRange(now))

	start = NewTimeOfDay(18, 0, 0)
	end = NewTimeOfDay(3, 0, 0)
	now = time.Date(2016, time.August, 10, 18, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCTimeRange(start, end).IsInRange(now))

	now = time.Date(2016, time.August, 10, 3, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCTimeRange(start, end).IsInRange(now))

	now = time.Date(2016, time.August, 10, 4, 0, 0, 0, time.UTC)
	assert.False(t, NewUTCTimeRange(start, end).IsInRange(now))

	now = time.Date(2016, time.August, 10, 17, 0, 0, 0, time.UTC)
	assert.False(t, NewUTCTimeRange(start, end).IsInRange(now))

	var tr *TimeRange
	assert.True(t, tr.IsInRange(now), "always in range if time range is nil")

	loc := time.FixedZone("myzone", -60)
	start = NewTimeOfDay(3, 0, 0)
	end = NewTimeOfDay(5, 0, 0)

	now = time.Date(2016, time.August, 10, 3, 0, 0, 0, time.UTC)
	assert.False(t, NewTimeRangeInLocation(start, end, loc).IsInRange(now))

	now = time.Date(2016, time.August, 10, 3, 1, 0, 0, time.UTC)
	assert.True(t, NewTimeRangeInLocation(start, end, loc).IsInRange(now))
}

func TestTimeRangeIsInSameRange(t *testing.T) {

	// start time is less than end time
	start := NewTimeOfDay(3, 0, 0)
	end := NewTimeOfDay(18, 0, 0)

	// same time
	time1 := time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)
	time2 := time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)

	assert.True(t, NewUTCTimeRange(start, end).IsInSameRange(time1, time2))
	assert.True(t, NewUTCTimeRange(start, end).IsInSameRange(time2, time1))
	// time 2 in same session but greater
	time1 = time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 11, 0, 0, 0, time.UTC)

	assert.True(t, NewUTCTimeRange(start, end).IsInSameRange(time1, time2))
	assert.True(t, NewUTCTimeRange(start, end).IsInSameRange(time2, time1))

	// time 2 in same session but less
	time1 = time.Date(2016, time.August, 10, 11, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)

	assert.True(t, NewUTCTimeRange(start, end).IsInSameRange(time1, time2))
	assert.True(t, NewUTCTimeRange(start, end).IsInSameRange(time2, time1))

	// time 1 not in session
	time1 = time.Date(2016, time.August, 10, 19, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)

	assert.False(t, NewUTCTimeRange(start, end).IsInSameRange(time1, time2))
	assert.False(t, NewUTCTimeRange(start, end).IsInSameRange(time2, time1))

	// time 2 not in session
	time1 = time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 2, 0, 0, 0, time.UTC)

	assert.False(t, NewUTCTimeRange(start, end).IsInSameRange(time1, time2))
	assert.False(t, NewUTCTimeRange(start, end).IsInSameRange(time2, time1))

	// start time is greater than end time
	start = NewTimeOfDay(18, 0, 0)
	end = NewTimeOfDay(3, 0, 0)

	// same session same day
	time1 = time.Date(2016, time.August, 10, 19, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 20, 0, 0, 0, time.UTC)

	assert.True(t, NewUTCTimeRange(start, end).IsInSameRange(time1, time2))
	assert.True(t, NewUTCTimeRange(start, end).IsInSameRange(time2, time1))

	// same session time 2 is in next day
	time1 = time.Date(2016, time.August, 10, 19, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 11, 2, 0, 0, 0, time.UTC)

	assert.True(t, NewUTCTimeRange(start, end).IsInSameRange(time1, time2))
	assert.True(t, NewUTCTimeRange(start, end).IsInSameRange(time2, time1))

	// same session time 1 is in next day
	time1 = time.Date(2016, time.August, 11, 2, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 19, 0, 0, 0, time.UTC)

	assert.True(t, NewUTCTimeRange(start, end).IsInSameRange(time1, time2))
	assert.True(t, NewUTCTimeRange(start, end).IsInSameRange(time2, time1))

	// time 1 is 25 hours greater (han time 2
	time1 = time.Date(2016, time.August, 11, 21, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 20, 0, 0, 0, time.UTC)

	assert.False(t, NewUTCTimeRange(start, end).IsInSameRange(time1, time2))
	assert.False(t, NewUTCTimeRange(start, end).IsInSameRange(time2, time1))

	// start time is greater than end time
	start = NewTimeOfDay(6, 0, 0)
	end = NewTimeOfDay(6, 0, 0)

	time1 = time.Date(2016, time.January, 13, 19, 10, 0, 0, time.UTC)
	time2 = time.Date(2016, time.January, 14, 19, 6, 0, 0, time.UTC)
	assert.False(t, NewUTCTimeRange(start, end).IsInSameRange(time1, time2))
	assert.False(t, NewUTCTimeRange(start, end).IsInSameRange(time2, time1))

	start = NewTimeOfDay(0, 0, 0)
	end = NewTimeOfDay(2, 0, 0)
	loc := time.FixedZone("myzone", -60)

	time1 = time.Date(2016, time.August, 10, 0, 1, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 0, 1, 0, 0, time.UTC)
	assert.True(t, NewTimeRangeInLocation(start, end, loc).IsInSameRange(time1, time2))
	assert.True(t, NewTimeRangeInLocation(start, end, loc).IsInSameRange(time2, time1))

	start = NewTimeOfDay(2, 0, 0)
	end = NewTimeOfDay(0, 0, 0)
	time1 = time.Date(2016, time.August, 10, 2, 1, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 2, 1, 0, 0, time.UTC)
	assert.True(t, NewTimeRangeInLocation(start, end, loc).IsInSameRange(time1, time2))
	assert.True(t, NewTimeRangeInLocation(start, end, loc).IsInSameRange(time2, time1))

	var tr *TimeRange
	assert.True(t, tr.IsInSameRange(time1, time2), "always in same range if time range is nil")
}
