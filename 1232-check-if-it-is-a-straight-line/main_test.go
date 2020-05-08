package main

import "testing"

var tests = []struct {
	coordinates [][]int
	want        bool
}{
	{
		[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}},
		true,
	},
	{
		[][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}},
		false,
	},
}

func TestCheckStraightLine(t *testing.T) {
	for _, test := range tests {
		if got := checkStraightLine(test.coordinates); got != test.want {
			t.Errorf("checkStraightLine(%v) = %v; want %v\n", test.coordinates, got, test.want)
		}
	}
}
