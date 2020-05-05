package main

import "testing"

var tests = []struct {
	s string
	i int
}{
	{"leetcode", 0},
	{"loveleetcode", 2},
	{"lovelove", -1},
}

func TestFirstUniqChar(t *testing.T) {
	for _, test := range tests {
		if got := firstUniqChar(test.s); got != test.i {
			t.Errorf("firstUniqChar(%v) = %v; want %v", test.s, got, test.i)
		}
	}
}
