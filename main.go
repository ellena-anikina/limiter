package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	sync.Mutex
}

func (c *Counter) Inc() {
	c.Lock()
	c.value++
	c.Unlock()
}

func (c *Counter) Value() int {
	return c.value
}

func main() {
	c := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()

	fmt.Println("Final counter value:", c.Value())
}
