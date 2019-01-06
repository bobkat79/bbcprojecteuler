package projecteuler

import "testing"

func TestSmallestMultiple(t *testing.T) {
	tables := []struct {
		start  int
		end    int
		expres int
	}{
		{1, 10, 2520},
		{1, 5, 60},
		// {1, 20, 232792560},
		{0, 50, 0},
		{25, 16, 0},
	}
	for pos, table := range tables {
		res := SmallestMultiple(table.start, table.end)
		if res != table.expres {
			t.Errorf("failed test %d, expected %d, got %d\n", pos+1, table.expres, res)
		}
	}
}

func benchmarkSmallestMultiple(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SmallestMultiple(1, i)
	}
}

func BenchmarkSmallestMultiple10(b *testing.B) { benchmarkSmallestMultiple(10, b) }
func BenchmarkSmallestMultiple11(b *testing.B) { benchmarkSmallestMultiple(11, b) }
func BenchmarkSmallestMultiple12(b *testing.B) { benchmarkSmallestMultiple(12, b) }
func BenchmarkSmallestMultiple13(b *testing.B) { benchmarkSmallestMultiple(13, b) }
func BenchmarkSmallestMultiple14(b *testing.B) { benchmarkSmallestMultiple(14, b) }
func BenchmarkSmallestMultiple15(b *testing.B) { benchmarkSmallestMultiple(15, b) }
func BenchmarkSmallestMultiple16(b *testing.B) { benchmarkSmallestMultiple(16, b) }
func BenchmarkSmallestMultiple17(b *testing.B) { benchmarkSmallestMultiple(17, b) }
func BenchmarkSmallestMultiple18(b *testing.B) { benchmarkSmallestMultiple(18, b) }
func BenchmarkSmallestMultiple19(b *testing.B) { benchmarkSmallestMultiple(19, b) }
func BenchmarkSmallestMultiple20(b *testing.B) { benchmarkSmallestMultiple(20, b) }
