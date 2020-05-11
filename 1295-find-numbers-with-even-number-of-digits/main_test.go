package main

import "testing"

var tests = []struct {
	nums []int
	want int
}{
	{[]int{12, 345, 2, 6, 7896}, 2},
	{[]int{555, 901, 482, 1771}, 1},
	{[]int{12}, 1},
	{[]int{1}, 0},
}

func TestFindNumbers(t *testing.T) {
	for _, test := range tests {
		if got := findNumbers(test.nums); got != test.want {
			t.Errorf("findNumbers(%v) = %v; want %v", test.nums, got, test.want)
		}
	}
}
