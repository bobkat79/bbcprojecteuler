package projecteuler

func CollatzStepCounter(num int, c chan [2]int) {
	// Counts the number of Collatz steps to one.
	var count int
	for i := num; i != 1; {
		i = CollatzStep(i)
		count++
	}
	c <- [2]int{num, count}
}

func MostCollatzSteps(limit int) (num int, maxsteps int) {
	// Find the number with the maximum Collatz steps up to limit.
	cc := make(chan [2]int, 20)
	for i := 1; i <= limit; i++ {
		go CollatzStepCounter(i, cc)
	}
	for i := 1; i <= limit; i++ {
		// retrieve from channels and check
		ret := <-cc
		if ret[1] > maxsteps {
			maxsteps = ret[1]
			num = ret[0]
		}
	}
	return num, maxsteps
}
