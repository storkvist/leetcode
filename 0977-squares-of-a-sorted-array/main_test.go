package main

import "testing"

var tests = []struct {
	A    []int
	want []int
}{
	{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
	{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	{[]int{-1, 1}, []int{1, 1}},
	{[]int{0}, []int{0}},
}

func TestSortedSquares(t *testing.T) {
	for _, test := range tests {
		got := sortedSquares(test.A)
		if len(got) != len(test.want) {
			t.Errorf("sortedSquares(%v) = %v; want %v", test.A, got, test.want)
		}
		for i, n := range got {
			if n != test.want[i] {
				t.Errorf("sortedSquares(%v) = %v; want %v", test.A, got, test.want)
			}
		}
	}
}
