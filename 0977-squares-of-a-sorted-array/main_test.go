package main

import "testing"

var tests = []struct {
	A    []int
	want []int
}{
	{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
	{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	{[]int{-1, 1}, []int{1, 1}},
	{[]int{0}, []int{0}},
	{[]int{-3, 0, 2}, []int{0, 4, 9}},
	{[]int{-3, -3, -2, 1}, []int{1, 4, 9, 9}},
	{[]int{-3, -3, -2, 1, 1, 1}, []int{1, 1, 1, 4, 9, 9}},
}

func TestSortedSquaresInPlace(t *testing.T) {
	for _, test := range tests {
		A := make([]int, len(test.A))
		copy(A, test.A)
		got := sortedSquaresInPlace(A)
		for i, n := range test.want {
			if n != got[i] {
				t.Errorf("sortedSquaresInPlace(%v) = %v; want %v", test.A, got, test.want)
				break
			}
		}
	}
}

func TestSortedSquaresExtraMemory(t *testing.T) {
	for _, test := range tests {
		A := make([]int, len(test.A))
		copy(A, test.A)
		got := sortedSquaresExtraMemory(A)
		for i, n := range test.want {
			if n != got[i] {
				t.Errorf("sortedSquaresExtraMemory(%v) = %v; want %v", test.A, got, test.want)
				break
			}
		}
	}
}

func BenchmarkSortedSquaresInPlace(b *testing.B) {
	A := make([]int, len(tests[0].A))
	copy(A, tests[0].A)
	for i := 0; i < b.N; i++ {
		sortedSquaresInPlace(A)
	}
}

func BenchmarkSortedSquaresExtraMemory(b *testing.B) {
	A := make([]int, len(tests[0].A))
	copy(A, tests[0].A)
	for i := 0; i < b.N; i++ {
		sortedSquaresExtraMemory(A)
	}
}
