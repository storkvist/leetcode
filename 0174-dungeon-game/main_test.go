package main

import "testing"

var tests = []struct {
	dungeon [][]int
	want    int
}{
	{
		[][]int{
			{-2, -3, 3},
			{-5, -10, 1},
			{10, 30, -5},
		},
		7,
	},
	{
		[][]int{
			{-3, 5},
		},
		4,
	},
	{
		[][]int{
			{2, 1},
			{1, -1},
		},
		1,
	},
	{
		[][]int{
			{1, -3, 3},
			{0, -2, 0},
			{-3, -3, -3},
		},
		3,
	},
}

func TestCalculateMinimumHP(t *testing.T) {
	for _, test := range tests {
		if got := calculateMinimumHP(test.dungeon); got != test.want {
			t.Errorf("calculateMinimumHP(%v) = %v; want %v\n", test.dungeon, got, test.want)
		}
	}
}
