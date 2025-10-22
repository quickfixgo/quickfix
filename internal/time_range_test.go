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
	r, err := NewUTCTimeRange(NewTimeOfDay(3, 0, 0), NewTimeOfDay(18, 0, 0), []time.Weekday{})
	assert.Nil(t, err)
	assert.Equal(t, NewTimeOfDay(3, 0, 0), r.startTime)
	assert.Equal(t, NewTimeOfDay(18, 0, 0), r.endTime)
	assert.Equal(t, []time.Weekday{}, r.weekdays)
	assert.Nil(t, r.startDay)
	assert.Nil(t, r.endDay)
	assert.Equal(t, time.UTC, r.loc)
}

func TestNewTimeRangeInLocation(t *testing.T) {
	r, err := NewTimeRangeInLocation(NewTimeOfDay(3, 0, 0), NewTimeOfDay(18, 0, 0), []time.Weekday{}, time.Local)
	assert.Nil(t, err)
	assert.Equal(t, NewTimeOfDay(3, 0, 0), r.startTime)
	assert.Equal(t, NewTimeOfDay(18, 0, 0), r.endTime)
	assert.Equal(t, []time.Weekday{}, r.weekdays)
	assert.Nil(t, r.startDay)
	assert.Nil(t, r.endDay)
	assert.Equal(t, time.Local, r.loc)
}

func TestNewTimeRangeInNilLocation(t *testing.T) {
	_, err := NewTimeRangeInLocation(NewTimeOfDay(3, 0, 0), NewTimeOfDay(18, 0, 0), []time.Weekday{}, nil)
	assert.EqualError(t, err, "time: missing Location in call to NewTimeRangeInLocation")
}

func TestNewUTCWeekRange(t *testing.T) {
	r, err := NewUTCWeekRange(NewTimeOfDay(3, 0, 0), NewTimeOfDay(18, 0, 0), time.Monday, time.Wednesday)
	assert.Nil(t, err)
	assert.Equal(t, NewTimeOfDay(3, 0, 0), r.startTime)
	assert.Equal(t, NewTimeOfDay(18, 0, 0), r.endTime)
	assert.NotNil(t, r.startDay)
	assert.NotNil(t, r.endDay)
	assert.Equal(t, []time.Weekday{}, r.weekdays)
	assert.Equal(t, time.Monday, *r.startDay)
	assert.Equal(t, time.Wednesday, *r.endDay)
	assert.Equal(t, time.UTC, r.loc)
}

func TestNewWeekRangeInLocation(t *testing.T) {
	r, err := NewWeekRangeInLocation(NewTimeOfDay(3, 0, 0), NewTimeOfDay(18, 0, 0), time.Monday, time.Wednesday, time.Local)
	assert.Nil(t, err)
	assert.Equal(t, NewTimeOfDay(3, 0, 0), r.startTime)
	assert.Equal(t, NewTimeOfDay(18, 0, 0), r.endTime)
	assert.NotNil(t, r.startDay)
	assert.NotNil(t, r.endDay)
	assert.Equal(t, []time.Weekday{}, r.weekdays)
	assert.Equal(t, time.Monday, *r.startDay)
	assert.Equal(t, time.Wednesday, *r.endDay)
	assert.Equal(t, time.Local, r.loc)
}

func TestNewWeekRangeInNilLocation(t *testing.T) {
	_, err := NewWeekRangeInLocation(NewTimeOfDay(3, 0, 0), NewTimeOfDay(18, 0, 0), time.Monday, time.Wednesday, nil)
	assert.EqualError(t, err, "time: missing Location in call to NewTimeRangeInLocation")
}

func BenchmarkInRange(b *testing.B) {
	start := NewTimeOfDay(3, 0, 0)
	end := NewTimeOfDay(18, 0, 0)
	weekdays := []time.Weekday{}
	tr, _ := NewUTCTimeRange(start, end, weekdays)

	now := time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC)
	for n := 0; n < b.N; n++ {
		tr.IsInRange(now)
	}
}

func TestTimeRangeIsInRange(t *testing.T) {
	testcases := []struct {
		label           string
		start           TimeOfDay
		end             TimeOfDay
		weekdays        []time.Weekday
		location        *time.Location
		now             time.Time
		expectedInRange bool
	}{
		{
			label:           "10AM",
			start:           NewTimeOfDay(3, 0, 0),
			end:             NewTimeOfDay(18, 0, 0),
			weekdays:        []time.Weekday{},
			location:        time.UTC,
			now:             time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "6PM",
			start:           NewTimeOfDay(3, 0, 0),
			end:             NewTimeOfDay(18, 0, 0),
			weekdays:        []time.Weekday{},
			location:        time.UTC,
			now:             time.Date(2016, time.August, 10, 18, 0, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "2AM",
			start:           NewTimeOfDay(3, 0, 0),
			end:             NewTimeOfDay(18, 0, 0),
			weekdays:        []time.Weekday{},
			location:        time.UTC,
			now:             time.Date(2016, time.August, 10, 2, 0, 0, 0, time.UTC),
			expectedInRange: false,
		},
		{
			label:           "7PM",
			start:           NewTimeOfDay(3, 0, 0),
			end:             NewTimeOfDay(18, 0, 0),
			weekdays:        []time.Weekday{},
			location:        time.UTC,
			now:             time.Date(2016, time.August, 10, 19, 0, 0, 0, time.UTC),
			expectedInRange: false,
		},
		{
			label:           "6:01PM",
			start:           NewTimeOfDay(3, 0, 0),
			end:             NewTimeOfDay(18, 0, 0),
			weekdays:        []time.Weekday{},
			location:        time.UTC,
			now:             time.Date(2016, time.August, 10, 18, 1, 0, 0, time.UTC),
			expectedInRange: false,
		},
		{
			label:           "6PM, inverted start/end",
			start:           NewTimeOfDay(18, 0, 0),
			end:             NewTimeOfDay(3, 0, 0),
			weekdays:        []time.Weekday{},
			location:        time.UTC,
			now:             time.Date(2016, time.August, 10, 18, 0, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "3AM, inverted start/end",
			start:           NewTimeOfDay(18, 0, 0),
			end:             NewTimeOfDay(3, 0, 0),
			weekdays:        []time.Weekday{},
			location:        time.UTC,
			now:             time.Date(2016, time.August, 10, 3, 0, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "4AM, inverted start/end",
			start:           NewTimeOfDay(18, 0, 0),
			end:             NewTimeOfDay(3, 0, 0),
			weekdays:        []time.Weekday{},
			location:        time.UTC,
			now:             time.Date(2016, time.August, 10, 4, 0, 0, 0, time.UTC),
			expectedInRange: false,
		},
		{
			label:           "5PM, inverted start/end",
			start:           NewTimeOfDay(18, 0, 0),
			end:             NewTimeOfDay(3, 0, 0),
			weekdays:        []time.Weekday{},
			location:        time.UTC,
			now:             time.Date(2016, time.August, 10, 17, 0, 0, 0, time.UTC),
			expectedInRange: false,
		},
		{
			label:           "3AM in fixed zone of -60",
			start:           NewTimeOfDay(3, 0, 0),
			end:             NewTimeOfDay(5, 0, 0),
			weekdays:        []time.Weekday{},
			location:        time.FixedZone("myzone", -60),
			now:             time.Date(2016, time.August, 10, 3, 0, 0, 0, time.UTC),
			expectedInRange: false,
		},
		{
			label:           "3:01AM in fixed zone of -60",
			start:           NewTimeOfDay(3, 0, 0),
			end:             NewTimeOfDay(5, 0, 0),
			weekdays:        []time.Weekday{},
			location:        time.FixedZone("myzone", -60),
			now:             time.Date(2016, time.August, 10, 3, 1, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "equal start/end",
			start:           NewTimeOfDay(0, 0, 0),
			end:             NewTimeOfDay(0, 0, 0),
			weekdays:        []time.Weekday{},
			location:        time.UTC,
			now:             time.Date(2016, time.August, 10, 18, 0, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "10AM, weekdays in range",
			start:           NewTimeOfDay(3, 0, 0),
			end:             NewTimeOfDay(18, 0, 0),
			weekdays:        []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday},
			location:        time.UTC,
			now:             time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "10AM, weekdays out of range",
			start:           NewTimeOfDay(3, 0, 0),
			end:             NewTimeOfDay(18, 0, 0),
			weekdays:        []time.Weekday{time.Monday, time.Tuesday},
			location:        time.UTC,
			now:             time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC),
			expectedInRange: false,
		},
		{
			label:           "2AM, weekdays in range",
			start:           NewTimeOfDay(3, 0, 0),
			end:             NewTimeOfDay(18, 0, 0),
			weekdays:        []time.Weekday{time.Monday, time.Tuesday, time.Wednesday},
			location:        time.UTC,
			now:             time.Date(2016, time.August, 10, 2, 0, 0, 0, time.UTC),
			expectedInRange: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.label, func(t *testing.T) {
			timeRange, err := NewTimeRangeInLocation(tc.start, tc.end, tc.weekdays, tc.location)
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedInRange, timeRange.IsInRange(tc.now))
		})
	}
}

func TestTimeInRangeNil(t *testing.T) {
	now := time.Date(2016, time.August, 10, 18, 0, 0, 0, time.UTC)

	var tr *TimeRange
	assert.True(t, tr.IsInRange(now), "always in range if the time range is nil")
}

func TestTimeRangeIsInRangeWithDay(t *testing.T) {
	testcases := []struct {
		label           string
		startTime       TimeOfDay
		endTime         TimeOfDay
		startDay        time.Weekday
		endDay          time.Weekday
		weekdays        []time.Weekday
		location        *time.Location
		now             time.Time
		expectedInRange bool
	}{
		{
			label:           "2AM Wednesday",
			startTime:       NewTimeOfDay(3, 0, 0),
			endTime:         NewTimeOfDay(18, 0, 0),
			startDay:        time.Monday,
			endDay:          time.Thursday,
			location:        time.UTC,
			now:             time.Date(2004, time.July, 28, 2, 0, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "6PM Tuesday",
			startTime:       NewTimeOfDay(3, 0, 0),
			endTime:         NewTimeOfDay(18, 0, 0),
			startDay:        time.Monday,
			endDay:          time.Thursday,
			location:        time.UTC,
			now:             time.Date(2004, time.July, 27, 18, 0, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "10PM Tuesday",
			startTime:       NewTimeOfDay(3, 0, 0),
			endTime:         NewTimeOfDay(18, 0, 0),
			startDay:        time.Monday,
			endDay:          time.Thursday,
			location:        time.UTC,
			now:             time.Date(2004, time.July, 27, 22, 0, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "3AM Tuesday",
			startTime:       NewTimeOfDay(3, 0, 0),
			endTime:         NewTimeOfDay(18, 0, 0),
			startDay:        time.Monday,
			endDay:          time.Thursday,
			location:        time.UTC,
			now:             time.Date(2004, time.July, 27, 3, 0, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "2:59:59 Monday",
			startTime:       NewTimeOfDay(3, 0, 0),
			endTime:         NewTimeOfDay(18, 0, 0),
			startDay:        time.Monday,
			endDay:          time.Thursday,
			location:        time.UTC,
			now:             time.Date(2004, time.July, 26, 2, 59, 59, 0, time.UTC),
			expectedInRange: false,
		},
		{
			label:           "6:01PM Thursday",
			startTime:       NewTimeOfDay(3, 0, 0),
			endTime:         NewTimeOfDay(18, 0, 0),
			startDay:        time.Monday,
			endDay:          time.Thursday,
			location:        time.UTC,
			now:             time.Date(2004, time.July, 29, 18, 0, 1, 0, time.UTC),
			expectedInRange: false,
		},
		{
			label:           "2AM Saturday, inverted days",
			startTime:       NewTimeOfDay(3, 0, 0),
			endTime:         NewTimeOfDay(18, 0, 0),
			startDay:        time.Thursday,
			endDay:          time.Monday,
			location:        time.UTC,
			now:             time.Date(2004, time.July, 24, 2, 0, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "2AM Wednesday, inverted days",
			startTime:       NewTimeOfDay(3, 0, 0),
			endTime:         NewTimeOfDay(18, 0, 0),
			startDay:        time.Thursday,
			endDay:          time.Monday,
			location:        time.UTC,
			now:             time.Date(2004, time.July, 28, 2, 0, 0, 0, time.UTC),
			expectedInRange: false,
		},
		{
			label:           "3AM Thursday, inverted days",
			startTime:       NewTimeOfDay(3, 0, 0),
			endTime:         NewTimeOfDay(18, 0, 0),
			startDay:        time.Thursday,
			endDay:          time.Monday,
			location:        time.UTC,
			now:             time.Date(2004, time.July, 22, 3, 0, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "6PM Monday, inverted days",
			startTime:       NewTimeOfDay(3, 0, 0),
			endTime:         NewTimeOfDay(18, 0, 0),
			startDay:        time.Thursday,
			endDay:          time.Monday,
			location:        time.UTC,
			now:             time.Date(2004, time.July, 26, 18, 0, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "2:59AM Thursday, inverted days",
			startTime:       NewTimeOfDay(3, 0, 0),
			endTime:         NewTimeOfDay(18, 0, 0),
			startDay:        time.Thursday,
			endDay:          time.Monday,
			location:        time.UTC,
			now:             time.Date(2004, time.July, 22, 2, 59, 59, 0, time.UTC),
			expectedInRange: false,
		},
		{
			label:           "6:01PM Monday, inverted days",
			startTime:       NewTimeOfDay(3, 0, 0),
			endTime:         NewTimeOfDay(18, 0, 0),
			startDay:        time.Thursday,
			endDay:          time.Monday,
			location:        time.UTC,
			now:             time.Date(2004, time.July, 26, 18, 0, 1, 0, time.UTC),
			expectedInRange: false,
		},
		{
			label:           "8:59AM Sunday, whole week",
			startTime:       NewTimeOfDay(9, 1, 0),
			endTime:         NewTimeOfDay(8, 59, 0),
			startDay:        time.Sunday,
			endDay:          time.Sunday,
			location:        time.UTC,
			now:             time.Date(2006, time.December, 3, 8, 59, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "8:59:01AM Sunday, whole week",
			startTime:       NewTimeOfDay(9, 1, 0),
			endTime:         NewTimeOfDay(8, 59, 0),
			startDay:        time.Sunday,
			endDay:          time.Sunday,
			location:        time.UTC,
			now:             time.Date(2006, time.December, 3, 8, 59, 1, 0, time.UTC),
			expectedInRange: false,
		},
		{
			label:           "9:01AM Sunday, whole week",
			startTime:       NewTimeOfDay(9, 1, 0),
			endTime:         NewTimeOfDay(8, 59, 0),
			startDay:        time.Sunday,
			endDay:          time.Sunday,
			location:        time.UTC,
			now:             time.Date(2006, time.December, 3, 9, 1, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "9AM Sunday, whole week",
			startTime:       NewTimeOfDay(9, 1, 0),
			endTime:         NewTimeOfDay(8, 59, 0),
			startDay:        time.Sunday,
			endDay:          time.Sunday,
			location:        time.UTC,
			now:             time.Date(2006, time.December, 3, 9, 0, 0, 0, time.UTC),
			expectedInRange: false,
		},
		{
			label:           "8:59AM Monday, whole week",
			startTime:       NewTimeOfDay(9, 1, 0),
			endTime:         NewTimeOfDay(8, 59, 0),
			startDay:        time.Sunday,
			endDay:          time.Sunday,
			location:        time.UTC,
			now:             time.Date(2006, time.December, 4, 8, 59, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "8:59:01AM Monday, whole week",
			startTime:       NewTimeOfDay(9, 1, 0),
			endTime:         NewTimeOfDay(8, 59, 0),
			startDay:        time.Sunday,
			endDay:          time.Sunday,
			location:        time.UTC,
			now:             time.Date(2006, time.December, 4, 8, 59, 1, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "9:01AM Monday, whole week",
			startTime:       NewTimeOfDay(9, 1, 0),
			endTime:         NewTimeOfDay(8, 59, 0),
			startDay:        time.Sunday,
			endDay:          time.Sunday,
			location:        time.UTC,
			now:             time.Date(2006, time.December, 4, 9, 1, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "9AM Monday, whole week",
			startTime:       NewTimeOfDay(9, 1, 0),
			endTime:         NewTimeOfDay(8, 59, 0),
			startDay:        time.Sunday,
			endDay:          time.Sunday,
			location:        time.UTC,
			now:             time.Date(2006, time.December, 4, 9, 0, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "8:59AM Sunday, whole week, inverted times",
			startTime:       NewTimeOfDay(8, 59, 0),
			endTime:         NewTimeOfDay(9, 1, 0),
			startDay:        time.Sunday,
			endDay:          time.Sunday,
			location:        time.UTC,
			now:             time.Date(2006, time.December, 3, 8, 59, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "9:01AM Sunday, whole week, inverted times",
			startTime:       NewTimeOfDay(8, 59, 0),
			endTime:         NewTimeOfDay(9, 1, 0),
			startDay:        time.Sunday,
			endDay:          time.Sunday,
			location:        time.UTC,
			now:             time.Date(2006, time.December, 3, 9, 1, 0, 0, time.UTC),
			expectedInRange: true,
		},
		{
			label:           "8:59AM Monday, whole week, inverted times",
			startTime:       NewTimeOfDay(8, 59, 0),
			endTime:         NewTimeOfDay(9, 1, 0),
			startDay:        time.Sunday,
			endDay:          time.Sunday,
			location:        time.UTC,
			now:             time.Date(2006, time.December, 4, 8, 59, 0, 0, time.UTC),
			expectedInRange: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.label, func(t *testing.T) {
			timeRange, err := NewWeekRangeInLocation(tc.startTime, tc.endTime, tc.startDay, tc.endDay, tc.location)
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedInRange, timeRange.IsInRange(tc.now))
		})
	}
}

func TestTimeRangeIsInSameRange(t *testing.T) {
	testcases := []struct {
		label    string
		start    TimeOfDay
		end      TimeOfDay
		location *time.Location
		weekdays []time.Weekday
		examples []struct {
			label    string
			time1    time.Time
			time2    time.Time
			expected bool
		}
	}{
		{
			label:    "start time is less than end time",
			start:    NewTimeOfDay(3, 0, 0),
			end:      NewTimeOfDay(18, 0, 0),
			location: time.UTC,
			weekdays: []time.Weekday{},
			examples: []struct {
				label    string
				time1    time.Time
				time2    time.Time
				expected bool
			}{
				{
					label:    "same time",
					time1:    time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC),
					expected: true,
				},
				{
					label:    "time2 in same session but greater",
					time1:    time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 10, 11, 0, 0, 0, time.UTC),
					expected: true,
				},
				{
					label:    "time2 in same session but less",
					time1:    time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 10, 11, 0, 0, 0, time.UTC),
					expected: true,
				},
				{
					label:    "time1 not in same session",
					time1:    time.Date(2016, time.August, 10, 19, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC),
					expected: false,
				},
				{
					label:    "time2 not in same session",
					time1:    time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 10, 2, 0, 0, 0, time.UTC),
					expected: false,
				},
				{
					label:    "time1 within time range but next day",
					time1:    time.Date(2016, time.August, 11, 10, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC),
					expected: false,
				},
				{
					label:    "time2 within time range but next day",
					time1:    time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 11, 10, 0, 0, 0, time.UTC),
					expected: false,
				},
			},
		},
		{
			label:    "with weekdays, start time is less than end time",
			start:    NewTimeOfDay(3, 0, 0),
			end:      NewTimeOfDay(18, 0, 0),
			location: time.UTC,
			weekdays: []time.Weekday{time.Monday, time.Tuesday, time.Wednesday},
			examples: []struct {
				label    string
				time1    time.Time
				time2    time.Time
				expected bool
			}{
				{
					label:    "same time within range",
					time1:    time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC),
					expected: true,
				},
				{
					label:    "same time outside weekdays range",
					time1:    time.Date(2016, time.August, 11, 10, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 11, 10, 0, 0, 0, time.UTC),
					expected: false,
				},
				{
					label:    "time2 in same session but greater within range",
					time1:    time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 10, 11, 0, 0, 0, time.UTC),
					expected: true,
				},
				{
					label:    "time2 in same session but less within range",
					time1:    time.Date(2016, time.August, 10, 10, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 10, 11, 0, 0, 0, time.UTC),
					expected: true,
				},
				{
					label:    "time2 in same session but greater outside range",
					time1:    time.Date(2016, time.August, 11, 10, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 11, 11, 0, 0, 0, time.UTC),
					expected: false,
				},
				{
					label:    "time2 in same session but less outside range",
					time1:    time.Date(2016, time.August, 11, 10, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 11, 11, 0, 0, 0, time.UTC),
					expected: false,
				},
			},
		},
		{
			label:    "start time is greater than end time",
			start:    NewTimeOfDay(18, 0, 0),
			end:      NewTimeOfDay(3, 0, 0),
			location: time.UTC,
			weekdays: []time.Weekday{},
			examples: []struct {
				label    string
				time1    time.Time
				time2    time.Time
				expected bool
			}{
				{
					label:    "same session, same day",
					time1:    time.Date(2016, time.August, 10, 19, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 10, 20, 0, 0, 0, time.UTC),
					expected: true,
				},
				{
					label:    "same session, time2 is in next day",
					time1:    time.Date(2016, time.August, 10, 19, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 11, 2, 0, 0, 0, time.UTC),
					expected: true,
				},
				{
					label:    "same session, time1 is in next day",
					time1:    time.Date(2016, time.August, 11, 2, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 10, 19, 0, 0, 0, time.UTC),
					expected: true,
				},
				{
					label:    "time1 is 25 hours greater than time2",
					time1:    time.Date(2016, time.August, 11, 21, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 10, 20, 0, 0, 0, time.UTC),
					expected: false,
				},
			},
		},
		{
			label:    "start time is equal to end time",
			start:    NewTimeOfDay(6, 0, 0),
			end:      NewTimeOfDay(6, 0, 0),
			location: time.UTC,
			weekdays: []time.Weekday{},
			examples: []struct {
				label    string
				time1    time.Time
				time2    time.Time
				expected bool
			}{
				{
					label:    "times on different days",
					time1:    time.Date(2016, time.January, 13, 19, 10, 0, 0, time.UTC),
					time2:    time.Date(2016, time.January, 14, 19, 6, 0, 0, time.UTC),
					expected: false,
				},
			},
		},
		{
			label:    "in different time zone",
			start:    NewTimeOfDay(0, 0, 0),
			end:      NewTimeOfDay(2, 0, 0),
			location: time.FixedZone("myzone", -60),
			weekdays: []time.Weekday{},
			examples: []struct {
				label    string
				time1    time.Time
				time2    time.Time
				expected bool
			}{
				{
					label:    "same times in range",
					time1:    time.Date(2016, time.August, 10, 0, 1, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 10, 0, 1, 0, 0, time.UTC),
					expected: true,
				},
			},
		},
		{
			label:    "in different time zone, inverted start/end",
			start:    NewTimeOfDay(2, 0, 0),
			end:      NewTimeOfDay(0, 0, 0),
			location: time.FixedZone("myzone", -60),
			weekdays: []time.Weekday{},
			examples: []struct {
				label    string
				time1    time.Time
				time2    time.Time
				expected bool
			}{
				{
					label:    "same times in range",
					time1:    time.Date(2016, time.August, 10, 2, 1, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 10, 2, 1, 0, 0, time.UTC),
					expected: true,
				},
			},
		},
		{
			label:    "midnight to midnight session in time zone",
			start:    NewTimeOfDay(0, 0, 0),
			end:      NewTimeOfDay(0, 0, 0),
			location: time.FixedZone("custom", -60),
			weekdays: []time.Weekday{},
			examples: []struct {
				label    string
				time1    time.Time
				time2    time.Time
				expected bool
			}{
				{
					label:    "different days UTC but same in fixed zone",
					time1:    time.Date(2016, time.August, 10, 0, 0, 0, 0, time.UTC),
					time2:    time.Date(2016, time.August, 11, 0, 0, 0, 0, time.UTC),
					expected: false,
				},
			},
		},
		{
			label:    "midnight to midnight session in UTC",
			start:    NewTimeOfDay(0, 0, 0),
			end:      NewTimeOfDay(0, 0, 0),
			location: time.UTC,
			weekdays: []time.Weekday{},
			examples: []struct {
				label    string
				time1    time.Time
				time2    time.Time
				expected bool
			}{
				{
					label:    "spanning midnight",
					time1:    time.Date(2016, time.August, 10, 23, 59, 59, 0, time.UTC),
					time2:    time.Date(2016, time.August, 11, 0, 0, 0, 0, time.UTC),
					expected: false,
				},
			},
		},
		{
			label:    "1:49 to 1:49 session",
			start:    NewTimeOfDay(1, 49, 0),
			end:      NewTimeOfDay(1, 49, 0),
			location: time.UTC,
			weekdays: []time.Weekday{},
			examples: []struct {
				label    string
				time1    time.Time
				time2    time.Time
				expected bool
			}{
				{
					label:    "edgecase 1",
					time1:    time.Date(2016, time.August, 16, 1, 48, 21, 0, time.UTC),
					time2:    time.Date(2016, time.August, 16, 1, 49, 02, 0, time.UTC),
					expected: false,
				},
				{
					label:    "edgecase 2",
					time1:    time.Date(2016, time.August, 16, 13, 48, 21, 0, time.UTC),
					time2:    time.Date(2016, time.August, 16, 13, 49, 02, 0, time.UTC),
					expected: true,
				},
			},
		},
		{
			label:    "13:49 to 13:49 session",
			start:    NewTimeOfDay(13, 49, 0),
			end:      NewTimeOfDay(13, 49, 0),
			location: time.UTC,
			weekdays: []time.Weekday{},
			examples: []struct {
				label    string
				time1    time.Time
				time2    time.Time
				expected bool
			}{
				{
					label:    "edgecase 1",
					time1:    time.Date(2016, time.August, 16, 13, 48, 21, 0, time.UTC),
					time2:    time.Date(2016, time.August, 16, 13, 49, 02, 0, time.UTC),
					expected: false,
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.label, func(t *testing.T) {
			var timeRange *TimeRange
			var err error

			if tc.location == time.UTC {
				timeRange, err = NewUTCTimeRange(tc.start, tc.end, tc.weekdays)
			} else {
				timeRange, err = NewTimeRangeInLocation(tc.start, tc.end, tc.weekdays, tc.location)
			}
			assert.Nil(t, err)

			for _, example := range tc.examples {
				t.Run(example.label, func(t *testing.T) {
					assert.Equal(t, example.expected, timeRange.IsInSameRange(example.time1, example.time2))
					assert.Equal(t, example.expected, timeRange.IsInSameRange(example.time2, example.time1))
				})
			}
		})
	}

}

func TestTimeRangeIsInSameRangeWithDay(t *testing.T) {
	testcases := []struct {
		label     string
		startTime TimeOfDay
		endTime   TimeOfDay
		startDay  time.Weekday
		endDay    time.Weekday
		location  *time.Location
		cases     []struct {
			label    string
			time1    time.Time
			time2    time.Time
			expected bool
		}
	}{
		{
			label:     "base example",
			startTime: NewTimeOfDay(3, 0, 0),
			endTime:   NewTimeOfDay(18, 0, 0),
			startDay:  time.Monday,
			endDay:    time.Thursday,
			location:  time.UTC,
			cases: []struct {
				label    string
				time1    time.Time
				time2    time.Time
				expected bool
			}{
				{
					label:    "time1 before day range",
					time1:    time.Date(2004, time.July, 27, 3, 0, 0, 0, time.UTC), // Tuesday
					time2:    time.Date(2004, time.July, 25, 3, 0, 0, 0, time.UTC), // Sunday
					expected: false,
				},
				{
					label:    "time1 after day range",
					time1:    time.Date(2004, time.July, 31, 3, 0, 0, 0, time.UTC), // Saturday
					time2:    time.Date(2004, time.July, 27, 3, 0, 0, 0, time.UTC), // Tuesday
					expected: false,
				},
				{
					label:    "same date/time",
					time1:    time.Date(2004, time.July, 27, 3, 0, 0, 0, time.UTC), // Tuesday
					time2:    time.Date(2004, time.July, 27, 3, 0, 0, 0, time.UTC), // Tuesday
					expected: true,
				},
				{
					label:    "different days within day range",
					time1:    time.Date(2004, time.July, 26, 10, 0, 0, 0, time.UTC), // Monday
					time2:    time.Date(2004, time.July, 27, 3, 0, 0, 0, time.UTC),  // Tuesday
					expected: true,
				},
				{
					label:    "different days within day range 2",
					time1:    time.Date(2004, time.July, 27, 10, 0, 0, 0, time.UTC), // Tuesday
					time2:    time.Date(2004, time.July, 29, 2, 0, 0, 0, time.UTC),  // Thursday
					expected: true,
				},
				{
					label:    "same day, different weeks",
					time1:    time.Date(2004, time.July, 27, 10, 0, 0, 0, time.UTC), // Tuesday
					time2:    time.Date(2004, time.July, 20, 3, 0, 0, 0, time.UTC),  // Tuesday
					expected: false,
				},
				{
					label:    "same day, different weeks inverted",
					time1:    time.Date(2004, time.July, 20, 3, 0, 0, 0, time.UTC),  // Tuesday
					time2:    time.Date(2004, time.July, 27, 10, 0, 0, 0, time.UTC), // Tuesday
					expected: false,
				},
				{
					label:    "different weeks again",
					time1:    time.Date(2004, time.July, 26, 2, 0, 0, 0, time.UTC), // Monday
					time2:    time.Date(2004, time.July, 19, 3, 0, 0, 0, time.UTC), // Monday
					expected: false,
				},
			},
		},
		{
			label:     "start/end time within an hour of midnight, week long session",
			startTime: NewTimeOfDay(0, 5, 0),
			endTime:   NewTimeOfDay(23, 45, 0),
			startDay:  time.Sunday,
			endDay:    time.Saturday,
			location:  time.UTC,
			cases: []struct {
				label    string
				time1    time.Time
				time2    time.Time
				expected bool
			}{
				{
					label:    "check DST change is handled when missing an hour",
					time1:    time.Date(2006, time.April, 4, 0, 0, 0, 0, time.UTC),
					time2:    time.Date(2006, time.April, 3, 1, 0, 0, 0, time.UTC),
					expected: true,
				},
				{
					label:    "check DST change is handled when gaining an hour",
					time1:    time.Date(2006, time.October, 30, 0, 0, 0, 0, time.UTC),
					time2:    time.Date(2006, time.October, 31, 1, 0, 0, 0, time.UTC),
					expected: true,
				},
				{
					label:    "works across year boundary",
					time1:    time.Date(2006, time.December, 31, 10, 10, 10, 0, time.UTC),
					time2:    time.Date(2007, time.January, 1, 10, 10, 10, 0, time.UTC),
					expected: true,
				},
			},
		},
		{
			label:     "when session days are the same",
			startTime: NewTimeOfDay(9, 1, 0),
			endTime:   NewTimeOfDay(8, 59, 0),
			location:  time.UTC,
			cases: []struct {
				label    string
				time1    time.Time
				time2    time.Time
				expected bool
			}{
				{
					label:    "same time",
					time1:    time.Date(2006, time.December, 3, 9, 1, 0, 0, time.UTC),
					time2:    time.Date(2006, time.December, 3, 9, 1, 0, 0, time.UTC),
					expected: true,
				},
				{
					label:    "different week",
					time1:    time.Date(2006, time.December, 3, 9, 1, 0, 0, time.UTC),
					time2:    time.Date(2006, time.December, 10, 9, 1, 0, 0, time.UTC),
					expected: false,
				},
				{
					label:    "different day",
					time1:    time.Date(2006, time.December, 3, 9, 1, 0, 0, time.UTC),
					time2:    time.Date(2006, time.December, 4, 9, 1, 0, 0, time.UTC),
					expected: true,
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.label, func(t *testing.T) {
			var timeRange *TimeRange
			var err error

			if tc.location == time.UTC {
				timeRange, err = NewUTCWeekRange(tc.startTime, tc.endTime, tc.startDay, tc.endDay)
			} else {
				timeRange, err = NewWeekRangeInLocation(tc.startTime, tc.endTime, tc.startDay, tc.endDay, tc.location)
			}
			assert.Nil(t, err)

			for _, c := range tc.cases {
				t.Run(c.label, func(t *testing.T) {
					assert.Equal(t, c.expected, timeRange.IsInSameRange(c.time1, c.time2))
				})
			}
		})
	}
}

func TestTimeRangeIsInSameRangeWhenNil(t *testing.T) {
	time1 := time.Date(2016, time.August, 10, 2, 1, 0, 0, time.UTC)
	time2 := time.Date(2016, time.August, 10, 2, 1, 0, 0, time.UTC)

	var tr *TimeRange
	assert.True(t, tr.IsInSameRange(time1, time2), "always in same range if time range is nil")
}
