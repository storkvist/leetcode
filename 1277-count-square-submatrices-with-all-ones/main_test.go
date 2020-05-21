package main

import "testing"

var tests = []struct {
	matrix [][]int
	want   int
}{
	{[][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}, 15},
	{[][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}, 7},
}

func TestCountSquares(t *testing.T) {
	for _, test := range tests {
		if got := countSquares(test.matrix); got != test.want {
			t.Errorf("countSquares(%v) = %v; want %v\n", test.matrix, got, test.want)
		}
	}
}
