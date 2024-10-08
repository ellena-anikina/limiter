package worker_pool

func makePool(n int, handler func(int, string)) (handle func(string), wait func()) {
	pool := make(chan int, n)

	for i := 0; i < n; i++ {
		pool <- i
	}
	handle = func(phrase string) {
		id := <-pool
		go func() {
			handler(id, phrase)
			pool <- id
		}()
	}

	wait = func() {
		for i := 0; i < n; i++ {
			pool <- i
		}
	}
	return handle, wait
}
