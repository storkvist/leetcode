package main

import (
	"sort"
)

type ByQueuePosition [][]int

func (a ByQueuePosition) Len() int      { return len(a) }
func (a ByQueuePosition) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByQueuePosition) Less(i, j int) bool {
	if a[i][0] != a[j][0] {
		return a[i][0] > a[j][0]
	}
	return a[i][1] < a[j][1]
}

func reconstructQueue(people [][]int) [][]int {
	length := len(people)
	sort.Sort(ByQueuePosition(people))

	result := make([][]int, length+1)
	for _, person := range people {
		position := person[1]
		copy(result[position+1:], result[position:])
		result[position] = person
	}

	return result[:length]
}
