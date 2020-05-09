package main

import "testing"

var tests = []struct {
	num  int
	want bool
}{
	{16, true},
	{14, false},
	{2, false},
	{1, true},
}

func TestIsPerfectSquare(t *testing.T) {
	for _, test := range tests {
		if got := isPerfectSquare(test.num); got != test.want {
			t.Errorf("isPerfectSquare(%v) = %v; want %v", test.num, got, test.want)
		}
	}
}
