package internal

import (
	"sync"
	"time"
)

type EventTimer struct {
	f     func()
	timer *time.Timer
	done  chan struct{}
	wg    sync.WaitGroup
	once  sync.Once
}

func NewEventTimer(task func()) *EventTimer {
	t := &EventTimer{
		f:     task,
		timer: newStoppedTimer(),
		done:  make(chan struct{}),
	}

	t.wg.Add(1)
	go func() {
		defer t.wg.Done()

		for {
			select {

			case <-t.timer.C:
				t.f()

			case <-t.done:
				t.timer.Stop()
				return

			}
		}
	}()

	return t
}

func (t *EventTimer) Stop() {
	if t == nil {
		return
	}

	t.once.Do(func() {
		close(t.done)
	})

	t.wg.Wait()
}

func (t *EventTimer) Reset(timeout time.Duration) {
	if t == nil {
		return
	}

	t.timer.Reset(timeout)
}

func newStoppedTimer() *time.Timer {
	timer := time.NewTimer(time.Second)
	if !timer.Stop() {
		<-timer.C
	}
	return timer
}
