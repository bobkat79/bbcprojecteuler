package projecteuler

import "errors"

func MultSum(sumnums []int, target int) (retsum int, err error) {
	// Find all the natural number multiples of a slice of numbers up to a target (not inclusive).
	// Then Return the sum of those numbers but deduplicate any numbers that are multiples of the other numbers.
	// Example if the target is 10 and the slice contains 3 and 5, then the multiples will be 3, 5, 6, 9.
	// So Multsum will return the sum of those as 23.
	var accum []int
	for _, n := range sumnums {
		// topdiv will provide the ceiling for the multiplier
		var topdiv int
		if n == 0 {
			return 0, errors.New("zero is not valid for input numbers or target")
		}
		for tch := (target - 1); tch > 2; tch-- {
			// start at the target and work down till we find an even dividing point
			if tch%n == 0 {
				// evenly divisable set as top div and break
				topdiv = (tch / n)
				//fmt.Printf("top divider for %d is %d\n", n, topdiv)
				break
			}
		}
		for i := 1; i < topdiv+1; i++ {
			accum = append(accum, (n * i))
		}
	}
	seen := make(map[int]struct{}, len(accum))
	for _, v := range accum {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		retsum = retsum + v
	}
	return retsum, nil
}
