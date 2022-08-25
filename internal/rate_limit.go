package internal

import (
	"container/list"
	"sync"
	"time"
)

type token struct {
	enableTime time.Time
}

type LimitBucket struct {
	fillInterval time.Duration
	queue        *list.List
	mutex        sync.Mutex
}

func New(fillInterval time.Duration, limit uint64) *LimitBucket {
	bucket := &LimitBucket{
		fillInterval: fillInterval,
		queue:        list.New(),
	}
	now := time.Now()
	for i := 0; i < int(limit); i++ {
		bucket.queue.PushBack(&token{
			enableTime: now,
		})
	}
	return bucket
}

func (bucket *LimitBucket) Wait() {
	for {
		if bucket.tryTakeToken() {
			return
		}
	}
}

func (bucket *LimitBucket) WaitForTimeout(timeout time.Duration) {
	begin := time.Now()
	for {
		if time.Now().After(begin.Add(timeout)) {
			return
		}
		if bucket.tryTakeToken() {
			return
		}
	}
}

func (bucket *LimitBucket) tryTakeToken() bool {
	bucket.mutex.Lock()
	defer bucket.mutex.Unlock()
	front := bucket.queue.Front()
	fristToken := front.Value.(*token)
	now := time.Now()
	if now.After(fristToken.enableTime) {
		bucket.queue.Remove(front)
		bucket.queue.PushBack(&token{
			enableTime: now.Add(bucket.fillInterval),
		})
		return true
	}
	return false
}
