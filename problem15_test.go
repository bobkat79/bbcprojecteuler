package projecteuler

import (
	"testing"
)

func TestFindLaticePaths(t *testing.T) {
	tables := []struct {
		gridsize int
		expres   int
	}{
		{0, 0},
		{2, 6},
		{20, 137846528820},
	}
	for _, table := range tables {
		res := FindLaticePaths(table.gridsize)
		if res != table.expres {
			t.Errorf("Unexpected return for latice paths, expected %d, got %d.", table.expres, res)
		}
	}
}
