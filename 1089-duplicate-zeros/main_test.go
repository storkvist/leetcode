package main

import "testing"

var tests = []struct {
	arr  []int
	want []int
}{
	{[]int{8, 4, 5, 0, 0, 0, 0, 7}, []int{8, 4, 5, 0, 0, 0, 0, 0}},
	{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
	{[]int{1, 2, 3}, []int{1, 2, 3}},
	{[]int{1}, []int{1}},
	{[]int{0}, []int{0}},
	{[]int{0, 0, 0}, []int{0, 0, 0}},
	{[]int{0, 1}, []int{0, 0}},
	{[]int{0, 0, 1, 1}, []int{0, 0, 0, 0}},
}

func TestDuplicateZeros(t *testing.T) {
	for _, test := range tests {
		arr := test.arr
		duplicateZeros(arr)
		if len(test.arr) != len(test.want) {
			t.Errorf("duplicateZeros(%v) -> %v; want %v", test.arr, arr, test.want)
		}
		for i, n := range arr {
			if n != test.want[i] {
				t.Errorf("duplicateZeros(%v) -> %v; want %v", test.arr, arr, test.want)
			}
		}
	}
}
