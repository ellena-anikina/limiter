package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value       int
	valueAtomic int64
	sync.Mutex
}

func (c *Counter) Inc() {
	c.Lock()
	c.value++
	c.Unlock()
}

func (c *Counter) IncAtomic() {
	atomic.AddInt64(&c.valueAtomic, 1)
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) ValueAtomic() int64 {
	return atomic.LoadInt64(&c.valueAtomic)
}

func main() {
	c := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
			c.IncAtomic()
		}()
	}
	wg.Wait()

	fmt.Println("Final counter value:", c.Value())
	fmt.Println("Final counter value atomic:", c.ValueAtomic())
}
