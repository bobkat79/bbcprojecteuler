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

func GenTriNum(lvl int) int {
	// Generate triangle number based on level.  lvl 7 would be the 7th triangle number or 28.
	// This is the extremely slow brute force method.  Use FastTriNum instead.
	if lvl < 1 {
		return 0
	}
	var accum int
	for i := 1; i <= lvl; i++ {
		accum += i
	}
	return accum
}

func FastTriNum(lvl int) int {
	// So the above triangle number generation is slow although obvious.
	// A little further reading lead me to:
	// http://www.maths.surrey.ac.uk/hosted-sites/R.Knott/runsums/triNbProof.html
	return (lvl * (lvl + 1) / 2)
}

func ListFactors(num int) []int {
	// Return a slice of factorials for a number.  This is slow and uses the most obvious method.
	if num < 1 {
		return []int{}
	}
	var accum []int
	for i := 1; i < (num + 1); i++ {
		if num%i == 0 {
			accum = append(accum, i)
		}
	}
	return accum
}

func FastFactors(num int) int {
	// Returns the number of factors a number has.
	// This uses a more mathematical oriented solution.
	// There is a square root optimization you could use here but it was giving me troubles and added a dependency.
	accum := 1
	trackfact := make(map[int]int)
	// Work through the 2's.
	for num%2 == 0 {
		num = num / 2
		trackfact[2] += 1
	}
	// Work through anything above 3.
	for i := 3; num != 1; {
		for num%i == 0 {
			num = num / i
			trackfact[i] += 1
		}
		i += 2
	}
	for _, val := range trackfact {
		accum = accum * (val + 1)
	}
	return accum
}

func LargePlus(lsnum string, rsnum string) string {
	// Take a very large number (greater than int64) and add it to another very large number.
	// Accepts string type as input and sends back a string.
	// This Algorithm is taken from the algorithms in a nutshell book.
	var carry int
	var ret string
	for lpos, rpos := len(lsnum)-1, len(rsnum)-1; lpos >= 0 || rpos >= 0; lpos, rpos = lpos-1, rpos-1 {
		lnum := 0
		rnum := 0
		total := 0
		// Having the numbers be the same size is not a requirement so we need to check where we are.
		switch {
		case lpos < 0:
			rnum, _ = strconv.Atoi(string(rsnum[rpos]))
		case rpos < 0:
			lnum, _ = strconv.Atoi(string(lsnum[lpos]))
		default:
			rnum, _ = strconv.Atoi(string(rsnum[rpos]))
			lnum, _ = strconv.Atoi(string(lsnum[lpos]))
		}
		total = lnum + rnum + carry
		carry = 0
		// anything less than nine is easy
		if total > 9 {
			carry = 1
			total = total - 10
		}
		stotal := strconv.Itoa(total)
		ret = stotal + ret
	}
	if carry == 1 {
		ret = "1" + ret
	}
	return ret
}

func CollatzStep(num int) int {
	// Take a number and return the next Collatz step.
	if num%2 == 0 {
		return (num / 2)
	}
	return ((num * 3) + 1)
}
