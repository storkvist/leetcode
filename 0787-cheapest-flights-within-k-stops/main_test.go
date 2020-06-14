package main

import "testing"

var tests = []struct {
	n       int
	flights [][]int
	src     int
	dst     int
	K       int
	want    int
}{
	{3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1, 200},
	{3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0, 500},
	{2, [][]int{}, 0, 1, 1, -1},
	{1, [][]int{}, 0, 0, 1, 0},
}

func TestFindCheapestPrice(t *testing.T) {
	for _, test := range tests {
		if got := findCheapestPrice(test.n, test.flights, test.src, test.dst, test.K); got != test.want {
			t.Errorf("findCheapestPrice(%v, %v, %v, %v, %v) = %v; want %v\n", test.n, test.flights, test.src, test.dst, test.K, got, test.want)
		}
	}
}
