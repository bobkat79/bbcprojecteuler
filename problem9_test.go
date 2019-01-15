package projecteuler

import "testing"

func TestPythTriplet(t *testing.T) {
	tables := []struct {
		target int
		expres int
	}{
		{12, 60},
		{1000, 31875000},
	}
	for pos, table := range tables {
		res := FindPythTriplet(table.target)
		if res != table.expres {
			t.Errorf("test %d failed for target %d, expected %d, got %d", pos+1, table.target, table.expres, res)
		}
	}
}
