package projecteuler

func BiggestPalindrome(ceilr int, ceill int) int {
	// Figure out the largest palindrome from the product of the two ceiling digits.
	var candidatepali []int
	var highpali int
	for r := ceilr; r > int((float32(ceilr) * .1)); r-- {
		for l := ceill; l > int((float32(ceill) * .1)); l-- {
			testceil := l * r
			if IsNumPalindrome(testceil) {
				// sadly brute force was the best I could come up with.
				candidatepali = append(candidatepali, testceil)
			}
		}
	}
	for _, chpali := range candidatepali {
		if chpali > highpali {
			highpali = chpali
		}
	}
	return highpali
}
