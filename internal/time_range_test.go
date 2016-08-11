package internal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func TestTimeRangeIsInRange(t *testing.T) {
	start := NewTimeOfDay(3, 0, 0)
	end := NewTimeOfDay(18, 0, 0)

	now := time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)
	assert.True(t, (&TimeRange{start, end}).IsInRange(now))

	now = time.Date(2016, time.August, 10, 18, 0, 0, 0, time.UTC)
	assert.True(t, (&TimeRange{start, end}).IsInRange(now))

	now = time.Date(2016, time.August, 10, 2, 0, 0, 0, time.UTC)
	assert.False(t, (&TimeRange{start, end}).IsInRange(now))

	now = time.Date(2016, time.August, 10, 19, 0, 0, 0, time.UTC)
	assert.False(t, (&TimeRange{start, end}).IsInRange(now))

	now = time.Date(2016, time.August, 10, 18, 0, 1, 0, time.UTC)
	assert.False(t, (&TimeRange{start, end}).IsInRange(now))

	start = NewTimeOfDay(18, 0, 0)
	end = NewTimeOfDay(3, 0, 0)
	now = time.Date(2016, time.August, 10, 18, 0, 0, 0, time.UTC)
	assert.True(t, (&TimeRange{start, end}).IsInRange(now))

	now = time.Date(2016, time.August, 10, 3, 0, 0, 0, time.UTC)
	assert.True(t, (&TimeRange{start, end}).IsInRange(now))

	now = time.Date(2016, time.August, 10, 4, 0, 0, 0, time.UTC)
	assert.False(t, (&TimeRange{start, end}).IsInRange(now))

	now = time.Date(2016, time.August, 10, 17, 0, 0, 0, time.UTC)
	assert.False(t, (&TimeRange{start, end}).IsInRange(now))

	var tr *TimeRange
	assert.True(t, tr.IsInRange(now), "always in range if time range is nil")
}

func TestTimeRangeIsInSameRange(t *testing.T) {

	// start time is less than end time
	start := NewTimeOfDay(3, 0, 0)
	end := NewTimeOfDay(18, 0, 0)

	// same time
	time1 := time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)
	time2 := time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)

	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time1, time2))
	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time2, time1))
	// time 2 in same session but greater
	time1 = time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 11, 0, 0, 0, time.UTC)

	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time1, time2))
	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time2, time1))

	// time 2 in same session but less
	time1 = time.Date(2016, time.August, 10, 11, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)

	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time1, time2))
	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time2, time1))

	// time 1 not in session
	time1 = time.Date(2016, time.August, 10, 19, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)

	assert.False(t, (&TimeRange{start, end}).IsInSameRange(time1, time2))
	assert.False(t, (&TimeRange{start, end}).IsInSameRange(time2, time1))

	// time 2 not in session
	time1 = time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 2, 0, 0, 0, time.UTC)

	assert.False(t, (&TimeRange{start, end}).IsInSameRange(time1, time2))
	assert.False(t, (&TimeRange{start, end}).IsInSameRange(time2, time1))

	// start time is greater than end time
	start = NewTimeOfDay(18, 0, 0)
	end = NewTimeOfDay(3, 0, 0)

	// same session same day
	time1 = time.Date(2016, time.August, 10, 19, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 20, 0, 0, 0, time.UTC)

	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time1, time2))
	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time2, time1))

	// same session time 2 is in next day
	time1 = time.Date(2016, time.August, 10, 19, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 11, 2, 0, 0, 0, time.UTC)

	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time1, time2))
	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time2, time1))

	// same session time 1 is in next day
	time1 = time.Date(2016, time.August, 11, 2, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 19, 0, 0, 0, time.UTC)

	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time1, time2))
	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time2, time1))

	// time 1 is 25 hours greater than time 2
	time1 = time.Date(2016, time.August, 11, 21, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 10, 20, 0, 0, 0, time.UTC)

	assert.False(t, (&TimeRange{start, end}).IsInSameRange(time1, time2))
	assert.False(t, (&TimeRange{start, end}).IsInSameRange(time2, time1))

	// start time is greater than end time
	start = NewTimeOfDay(6, 0, 0)
	end = NewTimeOfDay(6, 0, 0)

	time1 = time.Date(2016, time.January, 13, 19, 10, 0, 0, time.UTC)
	time2 = time.Date(2016, time.January, 14, 19, 6, 0, 0, time.UTC)
	assert.False(t, (&TimeRange{start, end}).IsInSameRange(time1, time2))
	assert.False(t, (&TimeRange{start, end}).IsInSameRange(time2, time1))

	loc, err := time.LoadLocation("America/Chicago")
	require.Nil(t, err)

	time1 = time.Date(2016, time.August, 10, 20, 53, 47, 397094815, loc)
	time2 = time.Date(2016, time.August, 10, 20, 53, 47, 397094992, loc)
	start = NewTimeOfDay(0, 53, 47)
	end = NewTimeOfDay(2, 53, 47)
	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time1, time2))
	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time2, time1))

	start = NewTimeOfDay(23, 31, 0)
	end = NewTimeOfDay(0, 31, 0)
	time1 = time.Date(2016, time.August, 10, 18, 31, 0, 88414271, loc)
	time2 = time.Date(2016, time.August, 10, 18, 31, 0, 88416914, loc)

	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time1, time2))
	assert.True(t, (&TimeRange{start, end}).IsInSameRange(time2, time1))

	var tr *TimeRange
	assert.True(t, tr.IsInSameRange(time1, time2), "always in same range if time range is nil")
}
