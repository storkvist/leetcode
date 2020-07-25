package main

import "testing"

var tests = []struct {
	nums []int
	want []int
}{
	{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
	{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
	{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
}

func TestRunningSum(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		got := runningSum(nums)
		if len(got) != len(test.want) {
			t.Errorf("runningSum(%v) = %v; want %v\n", test.nums, got, test.want)
			continue
		}
		for i := range got {
			if got[i] != test.want[i] {
				t.Errorf("runningSum(%v) = %v; want %v\n", test.nums, got, test.want)
				break
			}
		}
	}
}
