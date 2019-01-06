package projecteuler

func SmallestMultiple(s int, e int) int {
	// Find the smallest number that can be divided by each number from s to e without a remainder

	var mc int
	var t int
	var candidate int
	var multcounter int

	// return 0 if invalid input
	if s < 1 || e < s || e < 1 {
		return 0
	}

	// Brute force yet again, sigh.  I seriously need to read an algorithms book.
	// Also this is dangerous and could be infinite.
	for multcounter = 2; (e - 1) != mc; multcounter++ {
		t = e * multcounter
		mc = 0
		for i := s; i < e; i++ {
			// we don't need to check e because the candidate will always be a multiple of e.
			if t%i != 0 {
				break
			}
			mc += 1
			candidate = t
		}
	}
	return candidate
}
