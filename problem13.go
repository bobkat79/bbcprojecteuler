package projecteuler

import (
	"strings"
)

func LargeSum(largenum string) string {
	// Take a large number as string, split by newline, sum them together and return the result.
	var accum string
	accum = "0"
	n := strings.Split(largenum, "\n")
	for _, sval := range n {
		accum = LargePlus(sval, accum)
	}
	return accum
}
