package main

import "testing"

var tests = []struct {
	nums   []int
	target int
	want   int
}{
	{[]int{1, 3, 5, 6}, 5, 2},
	{[]int{1, 3, 5, 6}, 2, 1},
	{[]int{1, 3, 5, 6}, 7, 4},
	{[]int{1, 3, 5, 6}, 0, 0},
	{[]int{}, 1, 0},
}

func TestSearchInsert(t *testing.T) {
	for _, test := range tests {
		if got := searchInsert(test.nums, test.target); got != test.want {
			t.Errorf("searchInsert(%v, %v) = %v; want %v\n", test.nums, test.target, got, test.want)
		}
	}
}
