package main

import "testing"

var tests = []struct {
	nums []int
	want int
}{
	{[]int{1}, 1},
	{[]int{1, 2, 2}, 1},
	{[]int{1, 1, 2}, 2},
	{[]int{1, 1, 2, 3, 3}, 2},
	{[]int{1, 2, 2, 3, 3}, 1},
	{[]int{1, 1, 2, 2, 3}, 3},
	{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
	{[]int{3, 3, 7, 7, 10, 11, 11}, 10},
}

func TestSingleNonDuplicate(t *testing.T) {
	for _, test := range tests {
		if got := singleNonDuplicate(test.nums); got != test.want {
			t.Errorf("singleNonDuplicate(%v) = %v; want %v", test.nums, got, test.want)
		}
	}
}
