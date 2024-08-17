package projecteuler

func LargestProductinGrid(adjc int, g map[int][]int) (maxproduct int, maxset []int, orient string) {
	// Find the largest product from an adjaceny in a grid along with information.
	// The grid is passed in through g as a map where each entry is a row and each position in the slice is column.
	// This expects a grid so check the input grid to ensure it's actually a grid.
	gridrows := len(g)
	var accum int
	// var maxproduct int
	// var orient string
	// maxset := make([]int, adjc)
	for _, row := range g {
		if len(row) != gridrows {
			// grid malformed
			return 0, maxset, orient
		}
	}
	if adjc > gridrows {
		// You can't have more adjacencies than the size of the grid.
		return 0, maxset, orient
	}
	for rpos, row := range g {
		for cpos := range row {
			tempset := make([]int, adjc)
			// check horizontal left to right
			if (cpos + adjc) <= len(row) {
				// step through section and track
				accum = 1
				for i := 0; i < adjc; i++ {
					tempset[i] = row[(cpos + i)]
					accum = accum * row[(cpos+i)]
				}
				if accum > maxproduct {
					maxproduct = accum
					orient = "horizontal"
					maxset = tempset
				}
			}
			// check down
			if (rpos + adjc) <= (len(g) + 1) {
				accum = 1
				for i := 0; i < adjc; i++ {
					tempset[i] = g[rpos+i][cpos]
					accum = accum * g[rpos+i][cpos]
				}
				if accum > maxproduct {
					maxproduct = accum
					orient = "down"
					maxset = tempset
				}
			}
			// check right diagonal
			if (cpos+adjc) <= len(row) && (rpos+adjc) <= (len(g)+1) {
				accum = 1
				for i := 0; i < adjc; i++ {
					tempset[i] = g[rpos+i][cpos+i]
					accum = accum * g[rpos+i][cpos+i]
				}
				if accum > maxproduct {
					maxproduct = accum
					orient = "rdiag"
					maxset = tempset
				}
			}
			// check left diagonal
			if (cpos+1) >= adjc && (rpos+adjc) <= (len(g)+1) {
				accum = 1
				for i := 0; i < adjc; i++ {
					tempset[i] = g[rpos+i][cpos-i]
					accum = accum * g[rpos+i][cpos-i]
				}
				if accum > maxproduct {
					maxproduct = accum
					orient = "ldiag"
					maxset = tempset
				}
			}
		}
	}
	return maxproduct, maxset, orient
}
