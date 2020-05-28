package main

import "testing"

var tests = []struct {
	num  int
	want []int
}{
	{0, []int{0}},
	{2, []int{0, 1, 1}},
	{5, []int{0, 1, 1, 2, 1, 2}},
}

func TestCountBits(t *testing.T) {
	for _, test := range tests {
		got := countBits(test.num)
		if len(got) != len(test.want) {
			t.Errorf("countBits(%v) = %v; want %v\n", test.num, got, test.want)
			continue
		}
		for i := range test.want {
			if got[i] != test.want[i] {
				t.Errorf("countBits(%v) = %v; want %v\n", test.num, got, test.want)
				break
			}
		}
	}
}
