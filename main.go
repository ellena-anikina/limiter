package primes

import "context"

func generatePrimes(ctx context.Context, primes chan<- int) {
	defer close(primes)

	n := 2

	select {
	case <-ctx.Done():
		return
	default:
		if isPrime(n) {
			primes <- n
		}
		n++
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
