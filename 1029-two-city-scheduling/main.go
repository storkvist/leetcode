package main

import "sort"

type ByDiff [][]int

func (d ByDiff) Len() int           { return len(d) }
func (d ByDiff) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d ByDiff) Less(i, j int) bool { return (d[i][0] - d[i][1]) < (d[j][0] - d[j][1]) }

func twoCitySchedCost(costs [][]int) int {
	sort.Sort(ByDiff(costs))
	cost, count := 0, len(costs)/2
	for i := range costs {
		if i < count {
			cost += costs[i][0]
		} else {
			cost += costs[i][1]

		}

	}
	return cost
}
