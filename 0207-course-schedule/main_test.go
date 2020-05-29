package main

import "testing"

var tests = []struct {
	numCourses    int
	prerequisites [][]int
	want          bool
}{
	{1, [][]int{}, true},
	{2, [][]int{}, true},
	//{2, [][]int{{1, 0}}, true},
	//{2, [][]int{{1, 0}, {0, 1}}, false},
}

func TestCanFinish(t *testing.T) {
	for _, test := range tests {
		if got := canFinish(test.numCourses, test.prerequisites); got != test.want {
			t.Errorf("canFinish(%v, %v) = %v; want %v\n", test.numCourses, test.prerequisites, got, test.want)
		}
	}
}
