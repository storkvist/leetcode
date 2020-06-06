package main

import "testing"

var tests = []struct {
	people [][]int
	want   [][]int
}{
	{
		[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
		[][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
	},
}

func TestReconstructQueue(t *testing.T) {
	for _, test := range tests {
		got := reconstructQueue(test.people)
		if len(got) != len(test.want) {
			t.Errorf("reconstructQueue\n    (%+v) \n   = %+v \nwant %+v\n", test.people, got, test.want)
			continue
		}
		for i := range test.want {
			if len(got[i]) != 2 || got[i][0] != test.want[i][0] || got[i][1] != test.want[i][1] {
				t.Errorf("reconstructQueue\n    (%+v) \n   = %+v \nwant %+v\n", test.people, got, test.want)
				break
			}
		}
	}
}
