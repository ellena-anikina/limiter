package limiter

import (
	"sync"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	requestsCount := 20
	limiter := NewRateLimiter(requestsCount/4, 100*time.Millisecond)

	var usedTime time.Duration
	var wgR sync.WaitGroup

	wgR.Add(1)
	go func() {
		defer wgR.Done()
		start := time.Now()
		limiter.Run()
		t.Logf("Run time: %v", time.Now().Sub(start))
		usedTime = time.Now().Sub(start)
	}()

	var wg sync.WaitGroup
	for i := 0; i < requestsCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			limiter.AddRequest()
		}()
	}

	wg.Wait()
	close(limiter.requests)
	wgR.Wait()

	t.Logf("Expected time: %v", time.Duration(requestsCount/4*3)*100*time.Millisecond)

	if len(limiter.requests) != 0 {
		t.Errorf("Expected 0 remaining requests, got %d", len(limiter.requests))
	}

	expectedTime := time.Duration(requestsCount/4*3) * 100 * time.Millisecond
	if usedTime < expectedTime || usedTime > expectedTime+200*time.Millisecond {
		t.Errorf("Expected time to be more than %v, got %v", expectedTime, usedTime)
	}
}
