package projecteuler

func LargestPrimeFactor(n int) int {
	// Take a number and return the largest prime factor.

	// Going to cheat a little and preload primes under 600 to speed up the low calculations
	sliceoprimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599}
	var primefactors []int
	var highprime int
	for rem := n; rem > 1; {
		for _, pnum := range sliceoprimes {
			if rem%pnum == 0 {
				// evenly divides
				primefactors = append(primefactors, pnum)
				rem = rem / pnum
				break
			}
		}
		// we still have a remainder and we've gone through all the primes up to 600.  Now we have to use brute force.
		if IsPrime(rem) {
			primefactors = append(primefactors, rem)
			break
		}
		for i := 600; i < rem; i++ {
			if !IsPrime(i) {
				// skip non-primes
				continue
			}
			if rem%i == 0 {
				primefactors = append(primefactors, i)
				rem = rem / i
				break
			}
		}

	}
	for _, prime := range primefactors {
		if prime > highprime {
			highprime = prime
		}
	}
	return highprime
}
