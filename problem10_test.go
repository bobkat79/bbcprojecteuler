package projecteuler

import "testing"

func TestSumofPrimes(t *testing.T) {
	tables := []struct {
		limit  int
		expres int
	}{
		{10, 17},
		{2000000, 142913828922},
	}
	for _, table := range tables {
		res := SumofPrimes(table.limit)
		if res != table.expres {
			t.Errorf("failed test for limit %d, expected %d, got %d\n", table.limit, table.expres, res)
		}
	}
}
