package main

import "testing"

var tests = []struct {
	n    int
	want bool
}{
	{0, false},
	{1, true},
	{16, true},
	{218, false},
}

func TestIsPowerOfTwo(t *testing.T) {
	for _, test := range tests {
		if got := isPowerOfTwo(test.n); got != test.want {
			t.Errorf("isPowerOfTwo(%v) = %v; want %v\n", test.n, got, test.want)
		}
	}
}
