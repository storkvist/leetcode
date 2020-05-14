package main

import "testing"

var tests = []struct {
	A    []int
	want bool
}{
	{[]int{2, 1}, false},
	{[]int{3, 5, 5}, false},
	{[]int{0, 3, 2, 1}, true},
	{[]int{}, false},
	{[]int{1, 2, 1}, true},
	{[]int{2, 1, 2}, false},
	{[]int{1}, false},
	{[]int{1, 2}, false},
	{[]int{1, 1}, false},
	{[]int{1, 2, 2, 1}, false},
	{[]int{3, 5, 5}, false},
}

func TestValidMountainArray(t *testing.T) {
	for _, test := range tests {
		if got := validMountainArray(test.A); got != test.want {
			t.Errorf("validMountainArray(%v) = %v; want %v\n", test.A, got, test.want)
		}
	}
}
