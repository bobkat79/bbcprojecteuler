package projecteuler

import "testing"

func TestSumSquareDiff(t *testing.T) {
	tables := []struct {
		start  int
		end    int
		expres int
	}{
		{1, 10, 2640},
		{1, 5, 170},
		{0, 16, 0},
		{1, 0, 0},
		{16, 4, 0},
		{1, 100, 25164150},
	}
	for pos, table := range tables {
		res := SumSquareDiff(table.start, table.end)
		if res != table.expres {
			t.Errorf("Test %d failed, expected %d, got %d.\n", pos, table.expres, res)
		}
	}
}
