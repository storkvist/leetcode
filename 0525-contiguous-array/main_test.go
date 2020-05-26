package main

import "testing"

var tests = []struct {
	nums []int
	want int
}{
	{[]int{0, 0, 1, 0, 0, 0, 1, 1}, 6},
	{[]int{0, 1}, 2},
	{[]int{0, 1, 0}, 2},
	{[]int{0, 1, 0, 1, 0, 1, 1}, 6},
	{[]int{}, 0},
	{[]int{1}, 0},
	{[]int{0, 0, 0, 1, 1, 1}, 6},
	{[]int{0, 0, 0, 1, 1, 1, 0, 1}, 8},
}

func TestFindMaxLength(t *testing.T) {
	for _, test := range tests {
		if got := findMaxLength(test.nums); got != test.want {
			t.Errorf("findMaxNums(%v) = %v; want %v\n", test.nums, got, test.want)
		}
	}
}
