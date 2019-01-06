package projecteuler

import "testing"

func TestBiggestPalindrome(t *testing.T) {
	tables := []struct {
		ceilr  int
		ceill  int
		expres int
	}{
		{99, 99, 9009},
		{999, 999, 906609},
	}
	for pos, table := range tables {
		res := BiggestPalindrome(table.ceilr, table.ceill)
		if res != table.expres {
			t.Errorf("expected %d, returned %d for test %d\n", table.expres, res, pos)
		}
	}
}
