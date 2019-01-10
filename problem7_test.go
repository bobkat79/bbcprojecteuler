package projecteuler

import "testing"

func TestFindSpecificPrime(t *testing.T) {
	tables := []struct {
		num    int
		expres int
	}{
		{1, 2},
		{5, 11},
		{0, 0},
		{10001, 104743},
	}
	for _, table := range tables {
		res := FindSpecificPrime(table.num)
		if res != table.expres {
			t.Errorf("Prime returned for count %d not what was expected.  Expected %d, got %d.\n", table.num, table.expres, res)
		}
	}
}
