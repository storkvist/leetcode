package main

import "testing"

var tests = []struct {
	nums []int
	want []int
}{
	{[]int{1, 2, 3}, []int{2, 1}},
	{[]int{1, 2, 4, 8}, []int{8, 4, 2, 1}},
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
}

func TestLargestDivisibleSubset(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		got := largestDivisibleSubset(nums)
		if len(got) != len(test.want) {
			t.Errorf("largestDibisibleSubset(%v) = %v; want %v\n", test.nums, got, test.want)
			continue
		}
		for i := range test.want {
			if got[i] != test.want[i] {
				t.Errorf("largestDibisibleSubset(%v) = %v; want %v\n", test.nums, got, test.want)
				break
			}
		}
	}
}
