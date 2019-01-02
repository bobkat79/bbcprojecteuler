package projecteuler

import "testing"

func TestMultSum(t *testing.T) {
	tables := []struct {
		inums  []int
		target int
		exp    int
	}{
		{[]int{3, 5}, 10, 23},
		{[]int{7, 8, 12}, 36, 197},
		{[]int{3, 5}, 1000, 233168},
	}
	for pos, table := range tables {
		retsum, err := MultSum(table.inums, table.target)
		if err != nil {
			t.Errorf("received error on test: %d.", pos)
		}
		if retsum != table.exp {
			t.Errorf("expected: %d, received: %d, for test %d", table.exp, retsum, pos)
		}
	}
}

func BenchmarkMultSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultSum([]int{3, 5}, 100000)
	}
}
