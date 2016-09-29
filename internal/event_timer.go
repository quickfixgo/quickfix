package internal

import (
	"sync"
	"time"
)

type EventTimer struct {
	f     func()
	timer *time.Timer
	reset chan time.Duration
	wg    sync.WaitGroup
}

func NewEventTimer(task func()) *EventTimer {
	t := &EventTimer{
		f:     task,
		reset: make(chan time.Duration, 1),
	}

	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		var c <-chan time.Time

		for {
			select {

			case <-c:
				t.f()

			case d, ok := <-t.reset:
				if !ok {
					return
				}

				if t.timer != nil {
					t.timer.Reset(d)
				} else {
					t.timer = time.NewTimer(d)
					c = t.timer.C
				}
			}
		}
	}()

	return t
}

func (t *EventTimer) Stop() {
	if t == nil {
		return
	}

	close(t.reset)
	t.wg.Wait()
}

func (t *EventTimer) Reset(timeout time.Duration) {
	if t == nil {
		return
	}

	t.reset <- timeout
}
