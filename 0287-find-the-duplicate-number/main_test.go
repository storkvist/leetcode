package main

import "testing"

var tests = []struct {
	nums []int
	want int
}{
	{[]int{1, 3, 4, 2, 2}, 2},
	{[]int{3, 1, 3, 4, 2}, 3},
}

func TestFindDuplicate(t *testing.T) {
	for _, test := range tests {
		if got := findDuplicate(test.nums); got != test.want {
			t.Errorf("findDuplicate(%v) = %v; want %v\n", test.nums, got, test.want)
		}
	}
}
