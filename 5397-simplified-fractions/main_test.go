package main

import "testing"

var tests = []struct {
	n    int
	want []string
}{
	{2, []string{"1/2"}},
	{3, []string{"1/2", "1/3", "2/3"}},
	{4, []string{"1/2", "1/3", "1/4", "2/3", "3/4"}},
	{1, []string{}},
}

func TestSimplifiedFractions(t *testing.T) {
	for _, test := range tests {
		got := simplifiedFractions(test.n)
		if len(got) != len(test.want) {
			t.Errorf("simplifiedFractions(%v) = %v; want %v", test.n, got, test.want)
			continue
		}
		for i, x := range test.want {
			if x != got[i] {
				t.Errorf("simplifiedFractions(%v) = %v; want %v", test.n, got, test.want)
				continue
			}
		}
	}
}
