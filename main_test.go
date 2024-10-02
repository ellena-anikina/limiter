package primes

import (
	"context"
	"testing"
	"time"
)

func TestGeneratePrimes(t *testing.T) {
	primes := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go generatePrimes(ctx, primes)

	go func() {
		time.Sleep(1 * time.Second) // stop after 1 second
		cancel()
	}()

	for prime := range primes {
		if !isPrime(prime) {
			t.Errorf("%v is not a prime number", prime)
		}
	}
}
