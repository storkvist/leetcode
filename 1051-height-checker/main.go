package main

import "sort"

func heightChecker(heights []int) int {
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)

	count := 0
	for i, x := range sorted {
		if x != heights[i] {
			count++
		}
	}

	return count
}
