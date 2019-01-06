package projecteuler

import "strconv"

func IsPrime(p int) bool {
	if p == 2 || p == 3 {
		return true
	}
	if p%2 == 0 || p%3 == 0 {
		// anything divisable by 2 or 3 is not prime
		return false
	}
	for i := 5; i*i <= p; i += 6 {
		if p%i == 0 || p%(i+2) == 0 {
			return false
		}
	}
	return true
}

func IsNumPalindrome(num int) bool {
	// Determine if a number is a palindrome.
	// First convert it to a string so we can move around it.
	nsl := strconv.Itoa(num)
	for i := 0; i < len(nsl); i++ {
		if nsl[i] != nsl[len(nsl)-i-1] {
			return false
		}
	}
	return true
}
