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
	assert.Equal(t, 45244*time.Second, to.d)
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
	assert.Nil(t, r.startDay)
	assert.Nil(t, r.endDay)
	assert.Equal(t, time.UTC, r.loc)
}

func TestNewTimeRangeInLocation(t *testing.T) {
	r := NewTimeRangeInLocation(NewTimeOfDay(3, 0, 0), NewTimeOfDay(18, 0, 0), time.Local)
	assert.Equal(t, NewTimeOfDay(3, 0, 0), r.startTime)
	assert.Equal(t, NewTimeOfDay(18, 0, 0), r.endTime)
	assert.Nil(t, r.startDay)
	assert.Nil(t, r.endDay)
	assert.Equal(t, time.Local, r.loc)
}

func TestNewUTCWeekRange(t *testing.T) {
	r := NewUTCWeekRange(NewTimeOfDay(3, 0, 0), NewTimeOfDay(18, 0, 0), time.Monday, time.Wednesday)
	assert.Equal(t, NewTimeOfDay(3, 0, 0), r.startTime)
	assert.Equal(t, NewTimeOfDay(18, 0, 0), r.endTime)
	assert.NotNil(t, r.startDay)
	assert.NotNil(t, r.endDay)
	assert.Equal(t, time.Monday, *r.startDay)
	assert.Equal(t, time.Wednesday, *r.endDay)
	assert.Equal(t, time.UTC, r.loc)
}

func TestNewWeekRangeInLocation(t *testing.T) {
	r := NewWeekRangeInLocation(NewTimeOfDay(3, 0, 0), NewTimeOfDay(18, 0, 0), time.Monday, time.Wednesday, time.Local)
	assert.Equal(t, NewTimeOfDay(3, 0, 0), r.startTime)
	assert.Equal(t, NewTimeOfDay(18, 0, 0), r.endTime)
	assert.NotNil(t, r.startDay)
	assert.NotNil(t, r.endDay)
	assert.Equal(t, time.Monday, *r.startDay)
	assert.Equal(t, time.Wednesday, *r.endDay)
	assert.Equal(t, time.Local, r.loc)
}

func BenchmarkInRange(b *testing.B) {
	start := NewTimeOfDay(3, 0, 0)
	end := NewTimeOfDay(18, 0, 0)
	tr := NewUTCTimeRange(start, end)

	now := time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)
	for n := 0; n < b.N; n++ {
		tr.IsInRange(now)
	}
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

	start = NewTimeOfDay(0, 0, 0)
	end = NewTimeOfDay(0, 0, 0)
	now = time.Date(2016, time.August, 10, 18, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCTimeRange(start, end).IsInRange(now))

}

func TestTimeRangeIsInRangeWithDay(t *testing.T) {

	startTime := NewTimeOfDay(3, 0, 0)
	endTime := NewTimeOfDay(18, 0, 0)
	startDay := time.Monday
	endDay := time.Thursday

	now := time.Date(2004, time.July, 28, 2, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2004, time.July, 27, 18, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2004, time.July, 27, 3, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2004, time.July, 26, 2, 59, 59, 0, time.UTC)
	assert.False(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2004, time.July, 29, 18, 0, 1, 0, time.UTC)
	assert.False(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	startDay = time.Thursday
	endDay = time.Monday

	now = time.Date(2004, time.July, 24, 2, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2004, time.July, 28, 2, 0, 0, 0, time.UTC)
	assert.False(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2004, time.July, 22, 3, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2004, time.July, 26, 18, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2004, time.July, 22, 2, 59, 59, 0, time.UTC)
	assert.False(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2004, time.July, 26, 18, 0, 1, 0, time.UTC)
	assert.False(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	startTime = NewTimeOfDay(9, 1, 0)
	endTime = NewTimeOfDay(8, 59, 0)
	startDay = time.Sunday
	endDay = time.Sunday

	now = time.Date(2006, time.December, 3, 8, 59, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2006, time.December, 3, 8, 59, 1, 0, time.UTC)
	assert.False(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2006, time.December, 3, 9, 1, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))
	now = time.Date(2006, time.December, 3, 9, 0, 0, 0, time.UTC)
	assert.False(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2006, time.December, 4, 8, 59, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2006, time.December, 4, 8, 59, 1, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2006, time.December, 4, 9, 1, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2006, time.December, 4, 9, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	startTime = NewTimeOfDay(8, 59, 0)
	endTime = NewTimeOfDay(9, 1, 0)
	startDay = time.Sunday
	endDay = time.Sunday

	now = time.Date(2006, time.December, 3, 8, 59, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2006, time.December, 3, 9, 1, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))

	now = time.Date(2006, time.December, 4, 8, 59, 0, 0, time.UTC)
	assert.False(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInRange(now))
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

	start = NewTimeOfDay(0, 0, 0)
	end = NewTimeOfDay(0, 0, 0)
	time1 = time.Date(2016, time.August, 10, 0, 0, 0, 0, time.UTC)
	time2 = time.Date(2016, time.August, 11, 0, 0, 0, 0, time.UTC)
	assert.False(t, NewTimeRangeInLocation(start, end, loc).IsInSameRange(time1, time2))
	assert.False(t, NewTimeRangeInLocation(start, end, loc).IsInSameRange(time2, time1))

	time1 = time.Date(2016, time.August, 10, 23, 59, 59, 0, time.UTC)
	time2 = time.Date(2016, time.August, 11, 0, 0, 0, 0, time.UTC)
	assert.False(t, NewUTCTimeRange(start, end).IsInSameRange(time1, time2))
	assert.False(t, NewUTCTimeRange(start, end).IsInSameRange(time2, time1))

	start = NewTimeOfDay(1, 49, 0)
	end = NewTimeOfDay(1, 49, 0)
	time1 = time.Date(2016, time.August, 16, 1, 48, 21, 0, time.UTC)
	time2 = time.Date(2016, time.August, 16, 1, 49, 02, 0, time.UTC)

	assert.False(t, NewUTCTimeRange(start, end).IsInSameRange(time1, time2))
	assert.False(t, NewUTCTimeRange(start, end).IsInSameRange(time2, time1))

	start = NewTimeOfDay(1, 49, 0)
	end = NewTimeOfDay(1, 49, 0)
	time1 = time.Date(2016, time.August, 16, 13, 48, 21, 0, time.UTC)
	time2 = time.Date(2016, time.August, 16, 13, 49, 02, 0, time.UTC)

	assert.True(t, NewUTCTimeRange(start, end).IsInSameRange(time1, time2))
	assert.True(t, NewUTCTimeRange(start, end).IsInSameRange(time2, time1))

	start = NewTimeOfDay(13, 49, 0)
	end = NewTimeOfDay(13, 49, 0)
	time1 = time.Date(2016, time.August, 16, 13, 48, 21, 0, time.UTC)
	time2 = time.Date(2016, time.August, 16, 13, 49, 02, 0, time.UTC)

	assert.False(t, NewUTCTimeRange(start, end).IsInSameRange(time1, time2))
	assert.False(t, NewUTCTimeRange(start, end).IsInSameRange(time2, time1))
}

func TestTimeRangeIsInSameRangeWithDay(t *testing.T) {

	startTime := NewTimeOfDay(3, 0, 0)
	endTime := NewTimeOfDay(18, 0, 0)
	startDay := time.Monday
	endDay := time.Thursday

	time1 := time.Date(2004, time.July, 27, 3, 0, 0, 0, time.UTC)
	time2 := time.Date(2004, time.July, 25, 3, 0, 0, 0, time.UTC)
	assert.False(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInSameRange(time1, time2))

	time1 = time.Date(2004, time.July, 31, 3, 0, 0, 0, time.UTC)
	time2 = time.Date(2004, time.July, 27, 3, 0, 0, 0, time.UTC)
	assert.False(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInSameRange(time1, time2))

	time1 = time.Date(2004, time.July, 27, 3, 0, 0, 0, time.UTC)
	time2 = time.Date(2004, time.July, 27, 3, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInSameRange(time1, time2))

	time1 = time.Date(2004, time.July, 26, 10, 0, 0, 0, time.UTC)
	time2 = time.Date(2004, time.July, 27, 3, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInSameRange(time1, time2))

	time1 = time.Date(2004, time.July, 27, 10, 0, 0, 0, time.UTC)
	time2 = time.Date(2004, time.July, 29, 2, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInSameRange(time1, time2))

	time1 = time.Date(2004, time.July, 27, 10, 0, 0, 0, time.UTC)
	time2 = time.Date(2004, time.July, 20, 3, 0, 0, 0, time.UTC)
	assert.False(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInSameRange(time1, time2))

	time1 = time.Date(2004, time.July, 27, 2, 0, 0, 0, time.UTC)
	time2 = time.Date(2004, time.July, 20, 3, 0, 0, 0, time.UTC)
	assert.False(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInSameRange(time1, time2))

	time1 = time.Date(2004, time.July, 26, 2, 0, 0, 0, time.UTC)
	time2 = time.Date(2004, time.July, 19, 3, 0, 0, 0, time.UTC)
	assert.False(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInSameRange(time1, time2))

	// Reset start/end time so that they fall within an hour of midnight
	startTime = NewTimeOfDay(0, 5, 0)
	endTime = NewTimeOfDay(23, 45, 0)

	// Make it a week-long session
	startDay = time.Sunday
	endDay = time.Saturday

	// Check that ST-->DST (Sunday is missing one hour) is handled
	time1 = time.Date(2006, time.April, 4, 0, 0, 0, 0, time.UTC)
	time2 = time.Date(2006, time.April, 3, 1, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInSameRange(time1, time2))

	// Check that DST-->ST (Sunday has an extra hour) is handled
	time1 = time.Date(2006, time.October, 30, 0, 0, 0, 0, time.UTC)
	time2 = time.Date(2006, time.October, 31, 1, 0, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInSameRange(time1, time2))

	// Check that everything works across a year boundary
	time1 = time.Date(2006, time.December, 31, 10, 10, 10, 0, time.UTC)
	time2 = time.Date(2007, time.January, 1, 10, 10, 10, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInSameRange(time1, time2))

	// Session days are the same
	startDay = time.Sunday
	endDay = time.Sunday
	startTime = NewTimeOfDay(9, 1, 0)
	endTime = NewTimeOfDay(8, 59, 0)
	time1 = time.Date(2006, time.December, 3, 9, 1, 0, 0, time.UTC)
	time2 = time.Date(2006, time.December, 3, 9, 1, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInSameRange(time1, time2))

	time2 = time.Date(2006, time.December, 10, 9, 1, 0, 0, time.UTC)
	assert.False(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInSameRange(time1, time2))

	time2 = time.Date(2006, time.December, 4, 9, 1, 0, 0, time.UTC)
	assert.True(t, NewUTCWeekRange(startTime, endTime, startDay, endDay).IsInSameRange(time1, time2))
}
