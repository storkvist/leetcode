package main

import "testing"

var tests = []struct {
	cells []int
	N     int
	want  []int
}{
	{[]int{0, 1, 0, 1, 1, 0, 0, 1}, 7, []int{0, 0, 1, 1, 0, 0, 0, 0}},
	{[]int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000, []int{0, 0, 1, 1, 1, 1, 1, 0}},
}

func TestPrisonAfterNDays(t *testing.T) {
	for _, test := range tests {
		got := prisonAfterNDays(test.cells, test.N)
		if len(got) != len(test.want) {
			t.Errorf("prisonAfterNDays(%v, %v) = %v; want %v\n", test.cells, test.N, got, test.want)
			continue
		}
		for i := range test.want {
			if got[i] != test.want[i] {
				t.Errorf("prisonAfterNDays(%v, %v) = %v; want %v\n", test.cells, test.N, got, test.want)
				break
			}
		}
	}
}
