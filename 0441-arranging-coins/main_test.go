package main

import "testing"

var tests = []struct {
	n    int
	want int
}{
	{5, 2},
	{8, 3},
}

func TestArrangeCoins(t *testing.T) {
	for _, test := range tests {
		if got := arrangeCoins(test.n); got != test.want {
			t.Errorf("arrangeCoins(%v) = %v; want %v\n", test.n, got, test.want)
		}
	}
}
