package projecteuler

import "testing"

func TestIsPrime(t *testing.T) {
	tables := []struct {
		num    int
		expres bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{3119, true},
		{4049, true},
		{2658, false},
		{2205, false},
	}
	for pos, table := range tables {
		res := IsPrime(table.num)
		if res != table.expres {
			t.Errorf("test %d failed, expected %t, got %t\n.", pos, table.expres, res)
		}
	}
}

func TestIsNumPalindrome(t *testing.T) {
	tables := []struct {
		num    int
		expres bool
	}{
		{9009, true},
		{8118, true},
		{9801, false},
		{76411457, false},
		{1635641, false},
		{653356, true},
	}
	for _, table := range tables {
		res := IsNumPalindrome(table.num)
		if res != table.expres {
			t.Errorf("Number tested %d, expected %t, got %t\n", table.num, table.expres, res)
		}
	}
}
