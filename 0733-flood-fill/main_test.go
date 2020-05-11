package main

import "testing"

var tests = []struct {
	image    [][]int
	sr       int
	sc       int
	newColor int
	want     [][]int
}{
	{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
	{[][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1, [][]int{{0, 0, 0}, {0, 1, 1}}},
}

func TestFloodFill(t *testing.T) {
	for _, test := range tests {
		got := floodFill(test.image, test.sr, test.sc, test.newColor)
		for i, row := range test.want {
			for j, pixel := range row {
				if pixel != got[i][j] {
					t.Errorf(
						"floodFill(%v, %v, %v, %v) = %v; want %v\n",
						test.image, test.sr, test.sc, test.newColor, got, test.want)
				}
			}
		}
	}
}
