package main

import "testing"

var tests = []struct {
	arr  []int
	want bool
}{
	{[]int{0, 0}, true},
	{[]int{-1, -2}, true},
	{[]int{10, 2, 5, 3}, true},
	{[]int{7, 1, 14, 11}, true},
	{[]int{3, 1, 7, 11}, false},
}

func TestCheckIfExist(t *testing.T) {
	for _, test := range tests {
		if got := checkIfExist(test.arr); got != test.want {
			t.Errorf("checkIfExist(%v) = %v; want %v\n", test.arr, got, test.want)
		}
	}
}
