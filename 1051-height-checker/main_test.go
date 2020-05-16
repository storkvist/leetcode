package main

import "testing"

var tests = []struct {
	heights []int
	want    int
}{
	{[]int{1, 1, 4, 2, 1, 3}, 3},
	{[]int{5, 1, 2, 3, 4}, 5},
	{[]int{1, 2, 3, 4, 5}, 0},
}

func TestHeightChecker(t *testing.T) {
	for _, test := range tests {
		if got := heightChecker(test.heights); got != test.want {
			t.Errorf("heightChecker(%v) = %v; want %v\n", test.heights, got, test.want)
		}
	}
}
