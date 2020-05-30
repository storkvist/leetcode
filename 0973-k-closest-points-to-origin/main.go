package main

import "sort"

type ByEuclid [][]int

func (b ByEuclid) Len() int      { return len(b) }
func (b ByEuclid) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByEuclid) Less(i, j int) bool {
	bi := b[i][0]*b[i][0] + b[i][1]*b[i][1]
	bj := b[j][0]*b[j][0] + b[j][1]*b[j][1]
	return bi < bj
}

func kClosest(points [][]int, K int) [][]int {
	if K >= len(points) {
		return points
	}
	sort.Sort(ByEuclid(points))
	return points[:K]
}
