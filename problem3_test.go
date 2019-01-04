package projecteuler

import "testing"

func TestLargestPrimeFactor(t *testing.T) {
	tables := []struct {
		num    int
		expret int
	}{
		{13195, 29},
		{147, 7},
		{17, 17},
		{600851475143, 6857},
	}
	for _, table := range tables {
		ret := LargestPrimeFactor(table.num)
		if ret != table.expret {
			t.Errorf("expected %d, got %d", table.expret, ret)
		}
	}
}
