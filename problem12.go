package projecteuler

func TriDivisor(div int) int {
	// Find the first triangle number that has a specific number of divisors.
	if div < 1 {
		return 0
	}
	var trival int
	for i := 1; i > 0; i++ {
		testtri := FastTriNum(i)
		factors := FastFactors(testtri)
		if factors >= div {
			trival = testtri
			break
		}
	}
	return trival
}
