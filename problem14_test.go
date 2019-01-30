package projecteuler

import "testing"

func TestCollatzStepCounter(t *testing.T) {
	tables := []struct {
		num    int
		expres int
	}{
		{13, 9},
		{9, 19},
		{97, 118},
	}
	cc := make(chan [2]int, 20)
	for _, table := range tables {
		go CollatzStepCounter(table.num, cc)
		res := <-cc
		if res[1] != table.expres {
			t.Errorf("CollatzStepCounter test failed for number %d, expected %d got %d.\n", table.num, table.expres, res)
		}
	}
}

func TestMostCollatzSteps(t *testing.T) {
	tables := []struct {
		limit  int
		expres int
	}{
		{10, 9},
		{1000, 871},
		{1000000, 231},
	}
	for _, table := range tables {
		res, steps := MostCollatzSteps(table.limit)
		if res != table.expres {
			t.Errorf("MostCollatzSteps test failed for limit %d, expected %d got %d with steps %d.\n", table.limit, table.expres, res, steps)
		}
	}
}
