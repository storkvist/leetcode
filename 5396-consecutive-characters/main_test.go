package main

import "testing"

var tests = []struct {
	s    string
	want int
}{
	{"leetcode", 2},
	{"abbcccddddeeeeedcba", 5},
	{"triplepillooooow", 5},
	{"hooraaaaaaaaaaay", 11},
	{"tourist", 1},
	{"s", 1},
	{"aaaaa", 5},
	{"aaaaabaaaaa", 5},
	{"aabb", 2},
}

func TestMaxPower(t *testing.T) {
	for _, test := range tests {
		if got := maxPower(test.s); got != test.want {
			t.Errorf("maxPower(%v) = %v; want %v", test.s, got, test.want)
		}
	}
}
