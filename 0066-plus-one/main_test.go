package main

import "testing"

var tests = []struct {
	digits []int
	want   []int
}{
	{[]int{1, 2, 3}, []int{1, 2, 4}},
	{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
	{[]int{1, 2, 9}, []int{1, 3, 0}},
	{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
}

func TestPlusOne(t *testing.T) {
	for _, test := range tests {
		digits := make([]int, len(test.digits))
		copy(digits, test.digits)
		got := plusOne(digits)
		if len(got) != len(test.want) {
			t.Errorf("plusOne(%v) = %v; want %v\n", test.digits, got, test.want)
			continue
		}
		for i := range test.want {
			if got[i] != test.want[i] {
				t.Errorf("plusOne(%v) = %v; want %v\n", test.digits, got, test.want)
				break
			}
		}
	}
}
