package main

import "testing"

var tests = []struct {
	n    int
	want int
}{
	{3, 5},
}

func TestNumTrees(t *testing.T) {
	for _, test := range tests {
		if got := numTrees(test.n); got != test.want {
			t.Errorf("numTrees(%v) = %v; want %v\n", test.n, got, test.want)
		}
	}
}
