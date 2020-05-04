package main

import "testing"

var tests = []struct {
	num        int
	complement int
}{
	{5, 2},
	{1, 0},
	{0, 0},
	{7, 0},
	{10, 5},
	{1000000000, 73741823},
}

func TestBitwiseComplement(t *testing.T) {
	for _, test := range tests {
		if got := bitwiseComplement(test.num); got != test.complement {
			t.Errorf("bitwiseComplement(%v) = %v; want %v", test.num, got, test.complement)
		}
	}
}
