package projecteuler

func FindSpecificPrime(n int) int {
	// Find a specific prime number that is of n in the counts of primes.
	// n 1 would be 2, n 5 would be 11, etc.

	//invalid input
	if n < 1 {
		return 0
	}
	var accum int
	var ret int
	for i := 2; accum < n; i++ {
		if i > 3 && i%2 == 0 {
			continue
		}
		if IsPrime(i) {
			accum += 1
			ret = i
		}
	}
	return ret
}
