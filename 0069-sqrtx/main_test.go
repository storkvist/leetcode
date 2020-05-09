package main

import "testing"

var tests = []struct {
	x    int
	want int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 1},
	{4, 2},
	{5, 2},
	{6, 2},
	{7, 2},
	{8, 2},
}

func TestMySqrt(t *testing.T) {
	for _, test := range tests {
		if got := mySqrt(test.x); got != test.want {
			t.Errorf("mySqrt(%v) = %v; want %v", test.x, got, test.want)
		}
	}
}
