package main

import "testing"

var tests = []struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
	want  []int
}{
	{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
	{[]int{1, 2, 3, 0, 0, 0, 0}, 3, []int{2}, 1, []int{1, 2, 2, 3, 0, 0, 0}},
	{[]int{1, 0, 0, 0}, 1, []int{1, 2}, 2, []int{1, 1, 2, 0}},
	{[]int{0}, 0, []int{1}, 1, []int{1}},
	{[]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 6}},
}

func TestMerge(t *testing.T) {
	for _, test := range tests {
		arr := make([]int, len(test.nums1))
		copy(arr, test.nums1)
		merge(arr, test.m, test.nums2, test.n)
		for i, x := range test.want {
			if x != arr[i] {
				t.Errorf("merge(%v, %v, %v, %v) = %v; want %v", test.nums1, test.m, test.nums2, test.n, arr, test.want)
				break
			}
		}
	}
}
