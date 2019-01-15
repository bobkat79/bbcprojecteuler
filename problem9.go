package projecteuler

import "math"

func FindPythTriplet(tar int) (prod int) {
	// Find the Pythagorean triplet that adds up to the product in tar.
	// a < b < c and a + b + c = tar
	// return the product of a * b * c
	var a float64
	var b float64
	var c float64
	lima := float64(tar / 3)
	limb := float64(tar / 2)
	for a = lima; a > 1; a-- {
		// find a
		for b = limb; b > 1; b-- {
			if a >= b {
				break
			}
			asq := math.Pow(a, 2)
			bsq := math.Pow(b, 2)
			c = math.Sqrt(asq + bsq)
			if (a + b + c) == float64(tar) {
				return int(a * b * c)
			}
		}
	}

	return 0
}
