package projecteuler

import "testing"

func TestTriDivisor(t *testing.T) {
	tables := []struct {
		divisor int
		expres  int
	}{
		{4, 6},
		{5, 28},
		{500, 76576500},
	}
	for _, table := range tables {
		res := TriDivisor(table.divisor)
		if res != table.expres {
			t.Errorf("Failed to find Triangle Divisor for divisor of %d.  Expected %d, but got %d\n", table.divisor, table.expres, res)
		}
	}
}
