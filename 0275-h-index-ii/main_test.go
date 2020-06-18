package main

import "testing"

var tests = []struct {
	citations []int
	want      int
}{
	{[]int{0, 1, 3, 5, 6}, 3},
	{[]int{}, 0},
	{[]int{0}, 0},
	{[]int{0, 1}, 1},
	{[]int{0, 1, 1}, 1},
	{[]int{0, 2, 2, 3, 3, 3}, 3},
	{[]int{2, 2, 2, 2, 2, 2}, 2},
}

func TestHIndex(t *testing.T) {
	for _, test := range tests {
		if got := hIndex(test.citations); got != test.want {
			t.Errorf("hIndex(%v) = %v; want %v\n", test.citations, got, test.want)
		}
	}
}
