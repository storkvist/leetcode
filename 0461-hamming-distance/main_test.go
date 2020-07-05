package main

import "testing"

var tests = []struct {
	x    int
	y    int
	want int
}{
	{1, 4, 2},
}

func TestHammingDistance(t *testing.T) {
	for _, test := range tests {
		if got := hammingDistance(test.x, test.y); got != test.want {
			t.Errorf("hammingDistance(%v, %v) = %v; want %v\n", test.x, test.y, got, test.want)
		}
	}
}
