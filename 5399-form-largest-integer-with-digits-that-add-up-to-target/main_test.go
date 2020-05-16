package main

import "testing"

var tests = []struct {
	cost   []int
	target int
	want   string
}{
	{[]int{4, 3, 2, 5, 6, 7, 2, 5, 5}, 9, "7772"},
	{[]int{7, 6, 5, 5, 5, 6, 8, 7, 8}, 12, "85"},
	{[]int{2, 4, 6, 2, 4, 6, 4, 4, 4}, 5, "0"},
	{[]int{6, 10, 15, 40, 40, 40, 40, 40, 40}, 47, "32211"},
	{[]int{1, 1, 1, 1, 1, 1, 1, 3, 2}, 10, ""},
}

func TestLargestNumber(t *testing.T) {
	for _, test := range tests {
		if got := largestNumber(test.cost, test.target); got != test.want {
			t.Errorf("largestNumber(%v, %v) = %v; want %v\n", test.cost, test.target, got, test.want)
		}
	}
}
