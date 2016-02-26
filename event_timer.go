package quickfix

import (
	"time"
)

type eventTimer struct {
	Task  func()
	timer *time.Timer
}

func (t *eventTimer) Reset(timeout time.Duration) (ok bool) {
	if t.timer != nil {
		ok = t.timer.Stop()
	} else {
		ok = true
	}

	if t.Task != nil {
		t.timer = time.AfterFunc(timeout, t.Task)
	}
	return
}
