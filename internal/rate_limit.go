package internal

import (
	"time"

	"github.com/juju/ratelimit"
)

const (
	OnceTakeCount       = 1
	DefaultFillInterval = time.Second
)

type RateLimiter struct {
	ratePerSecond int64
	bucket        *ratelimit.Bucket
}

func NewRateLimiter(ratePerSecond int64) *RateLimiter {
	rateLimiter := &RateLimiter{
		ratePerSecond: ratePerSecond,
	}
	if rateLimiter.RateLimitIsOpen() {
		rateLimiter.bucket = ratelimit.NewBucketWithQuantum(DefaultFillInterval, ratePerSecond, ratePerSecond)
	}
	return rateLimiter
}

func (l *RateLimiter) RateLimitIsOpen() bool {
	return l.ratePerSecond > 0
}

func (l *RateLimiter) WaitRateLimit() {
	//if set ratePerSecond zero, not control send rate
	if !l.RateLimitIsOpen() {
		return
	}
	l.bucket.Wait(OnceTakeCount)
}
