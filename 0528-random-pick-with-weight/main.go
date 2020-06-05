package main

import (
	"math/rand"
	"sort"
)

type Solution struct {
	total   int
	weights []int
}

func Constructor(w []int) Solution {
	s := Solution{
		weights: make([]int, len(w)),
	}
	for i, e := range w {
		s.total += e
		s.weights[i] = s.total
	}
	return s
}

func (s *Solution) PickIndex() int {
	n := 1 + rand.Intn(s.total)
	return sort.SearchInts(s.weights, n)
}
