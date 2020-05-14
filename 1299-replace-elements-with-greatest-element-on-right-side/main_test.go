package main

import "testing"

var tests = []struct {
	arr  []int
	want []int
}{
	{[]int{2}, []int{-1}},
	{[]int{1, 1}, []int{1, -1}},
	{[]int{2, 1}, []int{1, -1}},
	{[]int{1, 2}, []int{2, -1}},
	{[]int{17, 18, 5, 4, 6, 1}, []int{18, 6, 6, 6, 1, -1}},
}

func TestReplaceElements(t *testing.T) {
	for _, test := range tests {
		arr := make([]int, len(test.arr))
		copy(arr, test.arr)
		got := replaceElements(arr)
		for i, n := range test.want {
			if got[i] != n {
				t.Errorf("replaceElements(%v) = %v; want %v", test.arr, got, test.want)
				break
			}
		}
	}
}
