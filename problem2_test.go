package projecteuler

import "testing"

func TestFibSums(t *testing.T) {
	tables := []struct {
		filter string
		limit  int64
		result int64
	}{
		{"even", 100, 44},
		{"odd", 100, 187},
		{"", 100, 231},
		{"even", 4000000, 4613732},
	}
	for pos, table := range tables {
		res, err := FibSums(table.limit, table.filter)
		if err != nil {
			t.Errorf("error returned from function.  %v", err)
		}
		if res != table.result {
			t.Errorf("Expected %d, got %d for test %d", table.result, res, pos+1)
		}
	}
}

func BenchmarkFibSums(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibSums(100000, "even")
	}
}
