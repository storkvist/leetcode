package main

import (
	"testing"
)

var tests = []struct {
	A []int
}{
	{[]int{3, 1, 2, 4}},
	{[]int{3, 1}},
	{[]int{2, 4}},
	{[]int{}},
}

func TestSortArrayByParity(t *testing.T) {
	for _, test := range tests {
		length := len(test.A)
		A := make([]int, length)
		copy(A, test.A)
		got := sortArrayByParity(A)
		i := 0
		for i < length && got[i]%2 == 0 {
			i++
		}
		j := i
		for j < length && got[j]%2 == 1 {
			j++
		}
		if (i > j) || (j != length) {
			t.Errorf("sortArrayByParity(%v) = %v; should be sorted", test.A, got)
		}
	}
}

func TestSortArrayByParityExtraMemory(t *testing.T) {
	for _, test := range tests {
		length := len(test.A)
		A := make([]int, length)
		copy(A, test.A)
		got := sortArrayByParityExtraMemory(A)
		i := 0
		for i < length && got[i]%2 == 0 {
			i++
		}
		j := i
		for j < length && got[j]%2 == 1 {
			j++
		}
		if (i > j) || (j != length) {
			t.Errorf("sortArrayByParityExtraMemory(%v) = %v; should be sorted", test.A, got)
		}
	}
}

func BenchmarkSortArrayByParity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArrayByParity(tests[0].A)
	}
}

func BenchmarkSortArrayByParityExtraMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArrayByParityExtraMemory(tests[0].A)
	}
}
