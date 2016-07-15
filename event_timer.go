package quickfix

import (
	"time"
)

type eventTimer struct {
	Task  func()
	timer *time.Timer
}

func (t *eventTimer) Stop() (ok bool) {
	if t.timer != nil {
		ok = t.timer.Stop()
	}
	return
}

func (t *eventTimer) Reset(timeout time.Duration) (ok bool) {
	ok = t.Stop()
	if t.Task != nil && timeout > 0 {
		t.timer = time.AfterFunc(timeout, t.Task)
	}
	return
}
