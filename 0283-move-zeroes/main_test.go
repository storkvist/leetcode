package main

import "testing"

var tests = []struct {
	nums []int
	want []int
}{
	{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	{[]int{}, []int{}},
	{[]int{0}, []int{0}},
	{[]int{0, 0}, []int{0, 0}},
	{[]int{1}, []int{1}},
	{[]int{2, 1}, []int{2, 1}},
	{[]int{2, 0}, []int{2, 0}},
}

func TestMoveZeroes(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		moveZeroes(nums)
		for i, n := range test.want {
			if nums[i] != n {
				t.Errorf("moveZeroes(%v) -> %v; want %v", test.nums, nums, test.want)
				break
			}
		}
	}
}
