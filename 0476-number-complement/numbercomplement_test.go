package numbercomplement

import "testing"

var tests = []struct {
	num        int
	complement int
}{
	{5, 2},
	{1, 0},
	{0, 0},
	{7, 0},
	{10, 5},
}

func TestFindComplement(t *testing.T) {
	for _, test := range tests {
		if got := findComplement(test.num); got != test.complement {
			t.Errorf("findComplement(%v) = %v; want %v", test.num, got, test.complement)
		}
	}
}
