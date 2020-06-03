package main

import "testing"

var tests = []struct {
	costs [][]int
	want  int
}{
	{[][]int{{1, 1}, {1, 1}}, 2},
	{[][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}, 110},
	{[][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}, 1859},
}

func TestTwoCitySchedCost(t *testing.T) {
	for _, test := range tests {
		if got := twoCitySchedCost(test.costs); got != test.want {
			t.Errorf("twoCitySchedCost(%v) = %v; want %v\n", test.costs, got, test.want)
		}
	}
}
