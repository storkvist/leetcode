package main

import "testing"

var tests = []struct {
	A    []int
	B    []int
	want int
}{
	{[]int{1, 4, 2}, []int{1, 2, 4}, 2},
	{[]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}, 3},
	{[]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}, 2},
	{
		[]int{15, 14, 1, 7, 15, 1, 12, 18, 9, 15, 1, 20, 18, 15, 16, 18, 11, 8, 11, 18, 11, 11, 17, 20, 16, 20, 15, 15, 9, 18, 16, 4, 16, 1, 13, 10, 10, 20, 4, 18, 17, 3, 8, 1, 8, 19, 14, 10, 10, 12},
		[]int{12, 8, 17, 4, 2, 18, 16, 10, 11, 12, 7, 1, 8, 16, 4, 14, 12, 18, 18, 19, 19, 1, 11, 18, 1, 6, 12, 17, 6, 19, 10, 5, 11, 16, 6, 17, 12, 1, 9, 3, 19, 2, 18, 18, 2, 4, 11, 11, 14, 9, 20, 19, 2, 20, 9, 15, 8, 7, 8, 6, 19, 12, 4, 11, 18, 18, 1, 6, 9, 17, 13, 19, 5, 4, 14, 9, 11, 15, 2, 5, 4, 1, 10, 11, 6, 4, 9, 7, 11, 7, 3, 8, 11, 12, 4, 19, 12, 17, 14, 18},
		23,
	},
}

func TestMaxUncrossedLines(t *testing.T) {
	for _, test := range tests {
		if got := maxUncrossedLines(test.A, test.B); got != test.want {
			t.Errorf("maxUncrossedLines(%v, %v) = %v; want %v", test.A, test.B, got, test.want)
		}
	}
}
