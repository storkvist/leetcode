package main

import "testing"

var tests = []struct {
	nums []int
	want []int
}{
	{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
}

func TestSortColors(t *testing.T) {
	for _, test := range tests {
		got := make([]int, len(test.nums))
		copy(got, test.nums)
		sortColors(got)
		if len(got) != len(test.want) {
			t.Errorf("sortColors(%v) -> %v; want %v\n", test.nums, got, test.want)
			continue
		}
		for i := range test.want {
			if got[i] != test.want[i] {
				t.Errorf("sortColors(%v) -> %v; want %v\n", test.nums, got, test.want)
				break
			}
		}
	}
}
