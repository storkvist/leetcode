package main

import "testing"

var tests = []struct {
	nums []int
	n    int
	want []int
}{
	{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
	{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
	{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
}

func TestShuffle(t *testing.T) {
	for _, test := range tests {
		got := shuffle(test.nums, test.n)
		if len(got) != len(test.want) {
			t.Errorf("shuffle(%v, %v) = %v; want %v\n", test.nums, test.n, got, test.want)
			continue
		}
		for i := range test.want {
			if got[i] != test.want[i] {
				t.Errorf("shuffle(%v, %v) = %v; want %v\n", test.nums, test.n, got, test.want)
				break
			}
		}
	}
}
