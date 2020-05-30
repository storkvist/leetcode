package main

import (
	"sort"
	"strconv"
	"testing"
)

var tests = []struct {
	points [][]int
	K      int
	want   [][]int
}{
	{
		[][]int{{-41, 72}, {53, 83}, {-95, -31}, {-61, 68}, {32, -56}, {16, 88}, {-81, -48}, {-31, 56}, {-57, -74}, {24, -42}, {-59, 72}, {-86, 40}, {34, -85}, {-45, 22}, {-35, -95}},
		9,
		[][]int{{24, -42}, {-45, 22}, {-31, 56}, {32, -56}, {-41, 72}, {16, 88}, {-61, 68}, {34, -85}, {-59, 72}},
	},
	{[][]int{{0, 0}}, 2, [][]int{{0, 0}}},
	{[][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
	{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2, [][]int{{3, 3}, {-2, 4}}},
}

func TestKClosest(t *testing.T) {
	for _, test := range tests {
		points := clone(test.points)
		if got := kClosest(points, test.K); !match(got, test.want) {
			t.Errorf("kClosest(%v, %v) = \n\t%v; want \n\t%v", test.points, test.K, got, test.want)
		}
	}
}

func clone(a [][]int) (b [][]int) {
	b = make([][]int, len(a))
	for i, pair := range a {
		b[i] = make([]int, 2)
		copy(b[i], pair)
	}
	return b
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type ByHash [][]int

func (b ByHash) Len() int      { return len(b) }
func (b ByHash) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByHash) Less(i, j int) bool {
	b1 := strconv.Itoa(b[i][0]) + "-" + strconv.Itoa(b[i][1])
	b2 := strconv.Itoa(b[j][0]) + "-" + strconv.Itoa(b[j][1])
	return b1 < b2
}

func match(oa, ob [][]int) bool {
	if len(oa) != len(ob) {
		return false
	}
	a := clone(oa)
	b := clone(ob)
	sort.Sort(ByHash(a))
	sort.Sort(ByHash(b))
	for i := range a {
		if !equal(a[i], b[i]) {
			return false
		}
	}
	return true
}
