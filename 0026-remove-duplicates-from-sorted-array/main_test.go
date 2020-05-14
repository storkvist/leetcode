package main

import "testing"

var tests = []struct {
	nums     []int
	wantNums []int
	want     int
}{
	{[]int{1, 1, 2}, []int{1, 2}, 2},
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}, 5},
	{[]int{1, 1}, []int{1}, 1},
	{[]int{1, 2, 2}, []int{1, 2}, 2},
	{[]int{1}, []int{1}, 1},
	{[]int{}, []int{}, 0},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		got := removeDuplicates(nums)
		for i, n := range test.wantNums {
			if n != nums[i] {
				t.Errorf("removeDuplicates(%v) = %v -> %v; want %v", test.nums, got, nums, test.wantNums)
				break
			}
		}
	}
}
