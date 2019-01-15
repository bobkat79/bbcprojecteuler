package projecteuler

func SumofPrimes(limit int) int {
	// Find the sum of all the primes up to a limit.
	var accum int
	for i := 2; i < limit; i++ {
		if IsPrime(i) {
			accum += i
		}
	}
	return accum
}
