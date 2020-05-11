package main

import "testing"

var tests = []struct {
	N     int
	trust [][]int
	want  int
}{
	{2, [][]int{{1, 2}}, 2},
	{3, [][]int{{1, 3}, {2, 3}}, 3},
	{3, [][]int{{1, 3}, {2, 3}, {3, 1}}, -1},
	{1, [][]int{}, 1},
}

func TestFindJudge(t *testing.T) {
	for _, test := range tests {
		if got := findJudge(test.N, test.trust); got != test.want {
			t.Errorf("findJudge(%v, %v) = %v; want %v", test.N, test.trust, got, test.want)
		}
	}
}
