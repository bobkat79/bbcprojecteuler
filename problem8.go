package projecteuler

import (
	"math"
	"strconv"
)

func LargestProductInSeries(lnum string, adj int) (digitlist []int, ret int) {
	// Take a very large number as a string (say 1000 digits) and find the largest product based on a number of adjacencies.
	// Adjacencies being number of digits, so adj of 4 means the product of four sequencial digits 4 x 5 x 4 x 6.
	// Return the largest adjacency.

	var workingset []int
	var accum int = 1

	if adj < 2 || len(lnum) < adj || len(lnum) < 2 {
		return digitlist, ret
	}

	// first calculate the largest adjacency possible so we can short circuit if we find it.
	maxadj := int(math.Pow(9, float64(adj)))
	for i := 0; i < len(lnum)-1; i++ {
		strnum, err := strconv.Atoi(string(lnum[i]))
		if err != nil {
			panic(err)
		}
		if len(workingset) < adj {
			workingset = append(workingset, strnum)
			accum = accum * strnum
			continue
		}
		// filled the workingset, evaluate
		if accum > ret {
			// new maximum, short circuit if it can't get better
			digitlist = workingset
			ret = accum
			if ret == maxadj {
				return digitlist, ret
			}
		}
		// pop front
		workingset = workingset[1:]
		workingset = append(workingset, strnum)
		accum = 1
		for _, j := range workingset {
			accum = accum * j
		}
	}
	return digitlist, ret
}
