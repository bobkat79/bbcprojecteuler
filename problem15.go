package projecteuler

import "math/big"

func FindLaticePaths(grid int) int {
	// Takes a square grid size and returns the number of paths you can traverse through it only moving down or to the right.
	// an empty grid while valid will have no paths through it.
	if grid == 0 || grid == 1 {
		return 0
	}
	// This is a binomial coefficient problem.
	val := big.NewInt(0)
	val.Binomial(int64(grid*2), int64(grid))
	ret := val.Int64()
	return int(ret)
}
