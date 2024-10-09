package limiter

import (
	"time"
)

type RateLimiter struct {
	burstLimit int
	timeLimit  time.Duration
	requests   chan struct{}
	counter    int
}

func NewRateLimiter(burstLimit int, timeLimit time.Duration) *RateLimiter {
	return &RateLimiter{
		timeLimit:  timeLimit,
		burstLimit: burstLimit,
		requests:   make(chan struct{}),
		counter:    0,
	}
}

func (r *RateLimiter) AddRequest() {
	r.requests <- struct{}{}
}

func (r *RateLimiter) Run() {
	ticker := time.NewTicker(r.timeLimit)
	defer ticker.Stop()

	for {
		if r.counter < r.burstLimit {
			_ = <-r.requests
			r.counter++
			continue
		}
		select {
		case <-ticker.C:
			_, ok := <-r.requests
			if !ok {
				return
			}
		}
	}
}
