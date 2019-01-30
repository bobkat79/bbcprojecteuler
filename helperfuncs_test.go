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

func TestGenTriNum(t *testing.T) {
	tables := []struct {
		lvl    int
		expres int
	}{
		{7, 28},
		{9, 45},
	}
	for _, table := range tables {
		res := GenTriNum(table.lvl)
		if res != table.expres {
			t.Errorf("Generate Triangle number failed for %d, expected %d but got %d\n.", table.lvl, table.expres, res)
		}
	}
}

func TestFastTriNum(t *testing.T) {
	tables := []struct {
		lvl    int
		expres int
	}{
		{7, 28},
		{9, 45},
		{9800, 48024900},
	}
	for _, table := range tables {
		res := FastTriNum(table.lvl)
		if res != table.expres {
			t.Errorf("Generate Fast Triangle number failed for %d, expected %d but got %d\n.", table.lvl, table.expres, res)
		}
	}
}

func TestListFactors(t *testing.T) {
	tables := []struct {
		num    int
		expres []int
	}{
		{10, []int{1, 2, 5, 10}},
		{21, []int{1, 3, 7, 21}},
		{28, []int{1, 2, 4, 7, 14, 28}},
		{10294453, []int{1, 13, 349, 2269, 4537, 29497, 791881, 10294453}},
	}
	for _, table := range tables {
		res := ListFactors(table.num)
		if len(res) != len(table.expres) {
			t.Errorf("Test List Factors failed for %d.  Expected a return slice of size of %d but got %d, received %v\n", table.num, len(table.expres), len(res), res)
			break
		}
		for pos, val := range res {
			if val != table.expres[pos] {
				t.Errorf("Test List Factors failed for number %d, recieved %v, expected %v\n", table.num, table.expres, res)
				break
			}
		}

	}
}

func TestFastFactors(t *testing.T) {
	tables := []struct {
		num    int
		expres int
	}{
		{15, 4},
		{28, 6},
		{48024900, 405},
		{76576500, 576},
	}
	for _, table := range tables {
		res := FastFactors(table.num)
		if res != table.expres {
			t.Errorf("Test Fast Factors failed for num %d, expected %d, got %d\n", table.num, table.expres, res)
		}
	}
}

func TestLargePlus(t *testing.T) {
	tables := []struct {
		lnum1  string
		lnum2  string
		expres string
	}{
		{"371072875339021027987979", "98220837590246510135740250", "98591910465585531163728229"},
		{"44274228917", "4745144573", "49019373490"},
	}
	for pos, table := range tables {
		res := LargePlus(table.lnum1, table.lnum2)
		if res != table.expres {
			t.Errorf("Test Large Plus failed for num %d, expected %s, got %s\n", pos, table.expres, res)
		}
	}
}

func TestCollatzStep(t *testing.T) {
	tables := []struct {
		num    int
		expres int
	}{
		{13, 40},
		{5, 16},
		{100, 50},
		{1567312, 783656},
		{1567311, 4701934},
	}
	for pos, table := range tables {
		res := CollatzStep(table.num)
		if res != table.expres {
			t.Errorf("Failed test %d of CollatzStep, expected %d, got %d.\n", pos, table.expres, res)
		}
	}
}
