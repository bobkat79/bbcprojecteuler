package projecteuler

import (
	"math/big"
)

func PowerDigitSum(base int, exponent int) int {
	// get the power of the base as a number then sum the digits
	if base == 0 {
		return 0
	}
	if base == 1 || exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	x := big.NewInt(int64(base))
	y := big.NewInt(int64(exponent))
	x.Exp(x, y, nil)
	var sum int
	powertxt := x.String()
	for _, v := range powertxt {
		sum += int(v - '0')
	}
	return sum
}
