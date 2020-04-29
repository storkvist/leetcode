package reverseinteger

import "testing"

var tests = []struct {
	x        int
	reversed int
}{
	{123, 321},
	{-123, -321},
	{120, 21},
	{-2147483648, 0},
	{1563847412, 0},
}

func TestReverse(t *testing.T) {
	for _, test := range tests {
		if got := Reverse(test.x); got != test.reversed {
			t.Errorf("Reverse(%d) = %d; want %d", test.x, got, test.reversed)
		}
	}
}
