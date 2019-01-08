package projecteuler

func SumSquareDiff(s int, e int) int {
	// Find the difference between the sum of squares and the square of sums of numbers from s to e.

	var sqaccum int
	var sumaccum int

	// invalid inputs
	if s > e || s < 1 || e < 1 {
		return 0
	}

	for i := s; i < (e + 1); i++ {
		sumaccum += i
		sqaccum += i * i
	}
	return (sumaccum * sumaccum) - sqaccum
}
