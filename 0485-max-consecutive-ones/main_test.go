package main

import "testing"

var tests = []struct {
	nums []int
	want int
}{
	{[]int{1, 1, 0, 1, 1, 1}, 3},
	{[]int{1}, 1},
	{[]int{0}, 0},
	{[]int{1, 0, 1, 1, 0, 1}, 2},
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	for _, test := range tests {
		if got := findMaxConsecutiveOnes(test.nums); got != test.want {
			t.Errorf("findMaxConsecutiveOnes(%v) = %v; want %v", test.nums, got, test.want)
		}
	}
}
