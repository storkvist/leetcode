package main

import "testing"

var tests = []struct {
	A    [][]int
	B    [][]int
	want [][]int
}{
	{
		[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
		[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
		[][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
	},
	{
		[][]int{{0, 2}},
		[][]int{{3, 4}},
		[][]int{},
	},
	{
		[][]int{{0, 2}},
		[][]int{{2, 3}},
		[][]int{{2, 2}},
	},
	{
		[][]int{{3, 10}},
		[][]int{{5, 10}},
		[][]int{{5, 10}},
	},
}

func TestIntervalIntersection(t *testing.T) {
	for _, test := range tests {
		got := intervalIntersection(test.A, test.B)
		if len(got) != len(test.want) || (len(got) > 0 && len(got[0]) != len(test.want[0])) {
			t.Errorf("intervalIntersection(%v, %v) = %v; want %v\n", test.A, test.B, got, test.want)
			continue
		}
		for i := range test.want {
			for j := range test.want[i] {
				if got[i][j] != test.want[i][j] {
					t.Errorf("intervalIntersection(%v, %v) = %v; want %v\n", test.A, test.B, got, test.want)
					break
				}
			}
		}
	}
}

func BenchmarkIntervalIntersection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intervalIntersection(tests[0].A, tests[0].B)
	}
}

func BenchmarkIntervalIntersectionSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intervalIntersectionSlow(tests[0].A, tests[0].B)
	}
}
