package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	w     []int
	count int
	want  []int
}{
	{[]int{1}, 1, []int{0}},
	{[]int{1, 3}, 5, []int{0, 1, 1, 1, 0}},
}

func TestSolution_PickIndex(t *testing.T) {
	for _, test := range tests {
		fmt.Printf("!!!   w = %+v\n", test.w)
		s := Constructor(test.w)
		var results []int
		for i := 0; i < test.count; i++ {
			results = append(results, s.PickIndex())
		}
		fmt.Printf("results = %+v\n", results)
		fmt.Printf("   want = %+v\n", test.want)
	}
}
