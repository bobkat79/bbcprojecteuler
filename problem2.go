package projecteuler

import "errors"

func FibSums(limit int64, filter string) (res int64, err error) {
	// Take a set of the Fibonacci up to the limit and sum the numbers together.
	// You can use the filter arguement to filter what is summed by even numbers or odd numbers.
	// If you pass nil to the filter it will not filter the results.
	if filter != "even" && filter != "odd" && filter != "" {
		return 0, errors.New("invalid filter arguement")
	}
	var lastfib int64
	var curfib int64
	var nextfib int64
	var fibs []int64
	curfib = 1
	lastfib = 1
	for curfib < limit {
		// deal with the filter first
		switch filter {
		case "even":
			if curfib%2 == 0 {
				fibs = append(fibs, curfib)
			}
		case "odd":
			if curfib%2 != 0 {
				fibs = append(fibs, curfib)
			}
		default:
			fibs = append(fibs, curfib)
		}
		// roll the numbers
		nextfib = lastfib + curfib
		lastfib = curfib
		curfib = nextfib
	}
	for _, num := range fibs {
		res = res + int64(num)
	}
	return res, nil
}
