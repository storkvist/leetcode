package main

import "testing"

var tests = []struct {
	nums []int
	want int
}{
	{[]int{2, 2, 3, 2}, 3},
	{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
	{[]int{1}, 1},
}

func TestSingleNumber(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		if got := singleNumber(nums); got != test.want {
			t.Errorf("singleNumber(%v) = %v; want %v\n", nums, got, test.want)
		}
	}
}
