package internal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {
	to, err := ParseTime("12:34:04")
	assert.Nil(t, err)
	assert.Equal(t, NewTime(12, 34, 4), to)

	_, err = ParseTime("0:0:0")
	assert.NotNil(t, err)

	_, err = ParseTime("00:00")
	assert.NotNil(t, err)

	_, err = ParseTime("0000:00")
	assert.NotNil(t, err)
}

func TestTimeRangeIsInRange(t *testing.T) {
	start := NewTime(3, 0, 0)
	end := NewTime(18, 0, 0)

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

	start = NewTime(18, 0, 0)
	end = NewTime(3, 0, 0)
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
