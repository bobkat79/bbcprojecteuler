package projecteuler

import (
	"testing"
)

func TestPowerDigitSum(t *testing.T) {
	tables := []struct {
		base     int
		exponent int
		expres   int
	}{
		{6, 1, 6},
		{8, 0, 1},
		{2, 15, 26},
		{2, 1000, 1366},
	}
	for _, table := range tables {
		res := PowerDigitSum(table.base, table.exponent)
		if res != table.expres {
			t.Errorf("Expected power digit sum to return %d, got %d", table.expres, res)
		}
	}
}
