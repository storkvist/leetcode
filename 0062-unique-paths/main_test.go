package main

import "testing"

var tests = []struct {
	m    int
	n    int
	want int
}{
	{3, 2, 3},
	{7, 3, 28},
}

func TestUniquePaths(t *testing.T) {
	for _, test := range tests {
		if got := uniquePaths(test.m, test.n); got != test.want {
			t.Errorf("uniquePaths(%v, %v) = %v; want %v\n", test.m, test.n, got, test.want)
		}
	}
}
