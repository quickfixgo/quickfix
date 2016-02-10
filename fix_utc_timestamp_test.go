package quickfix_test

import (
	"github.com/quickfixgo/quickfix"
	"testing"
	"time"
)

func TestFIXUTCTimestampRead(t *testing.T) {
	var tests = []struct {
		timeStr          string
		expectedTime     time.Time
		expectedNoMillis bool
	}{
		{"20160208-22:07:16.000", time.Date(2016, time.February, 8, 22, 7, 16, 0, time.UTC), false},
		{"20160208-22:07:16", time.Date(2016, time.February, 8, 22, 7, 16, 0, time.UTC), true},
	}

	for _, test := range tests {
		var f quickfix.FIXUTCTimestamp
		if err := f.Read([]byte(test.timeStr)); err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !f.Value.Equal(test.expectedTime) {
			t.Errorf("For Time expected %v got %v", test.expectedTime, f.Value)
		}

		if f.NoMillis != test.expectedNoMillis {
			t.Errorf("For NoMillis expected %v got %v", test.expectedNoMillis, f.NoMillis)
		}
	}

}
