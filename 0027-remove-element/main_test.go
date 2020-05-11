package main

import (
	"sort"
	"testing"
)

var tests = []struct {
	nums    []int
	val     int
	wantArr []int
	want    int
}{
	{[]int{3, 2, 2, 3}, 3, []int{2, 2}, 2},
	{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 0, 1, 3, 4}, 5},
	{[]int{}, 0, []int{}, 0},
	{[]int{1}, 1, []int{}, 0},
	{[]int{1}, 2, []int{1}, 1},
}

func TestRemoveElement(t *testing.T) {
	for _, test := range tests {
		arr := make([]int, len(test.nums))
		copy(arr, test.nums)

		got := removeElement(arr, test.val)
		if got != test.want {
			t.Errorf("removeElement(%v, %v) = %v -> %v; want %v -> %v",
				test.nums, test.val, got, arr, test.want, test.wantArr)
			continue
		}
		arr = arr[:got]
		sort.Ints(arr)
		for i, x := range test.wantArr {
			if x != arr[i] {
				t.Errorf("removeElement(%v, %v) = %v -> %v; want %v -> %v",
					test.nums, test.val, got, arr, test.want, test.wantArr)
				continue
			}
		}
	}
}
