package twosum

import (
	"testing"
)

var tests = []struct {
	given  []int
	target int
	wanted []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
}

func TestTwoSum(t *testing.T) {
	for _, test := range tests {
		got := TwoSum(test.given, test.target)
		if len(got) < 2 || got[0] != test.wanted[0] || got[1] != test.wanted[1] {
			t.Errorf("TwoSum(%v, %v) = %v; want %v", test.given, test.target, got, test.wanted)
		}
	}
}

func TestTwoSum2(t *testing.T) {
	for _, test := range tests {
		got := TwoSum2(test.given, test.target)
		if len(got) < 2 || got[0] != test.wanted[0] || got[1] != test.wanted[1] {
			t.Errorf("TwoSum2(%v, %v) = %v; want %v", test.given, test.target, got, test.wanted)
		}
	}
}
