package internal

import (
	"time"
)

type EventTimer struct {
	Task  func()
	timer *time.Timer
}

func (t *EventTimer) Stop() (ok bool) {
	if t.timer != nil {
		ok = t.timer.Stop()
	}
	return
}

func (t *EventTimer) Reset(timeout time.Duration) (ok bool) {
	ok = t.Stop()
	if t.Task != nil && timeout > 0 {
		t.timer = time.AfterFunc(timeout, t.Task)
	}
	return
}
