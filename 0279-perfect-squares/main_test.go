package main

import "testing"

var tests = []struct {
	n    int
	want int
}{
	{12, 3},
	{13, 2},
	{1, 1},
}

func TestNumSquares(t *testing.T) {
	for _, test := range tests {
		if got := numSquares(test.n); got != test.want {
			t.Errorf("numSquares(%v) = %v; want %v\n", test.n, got, test.want)
		}
	}
}
