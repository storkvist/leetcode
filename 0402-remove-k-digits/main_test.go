package main

import "testing"

var tests = []struct {
	num  string
	k    int
	want string
}{
	{"10", 2, "0"},
	{"10", 1, "0"},
	{"8888188", 3, "8188"},
	{"1432219", 3, "1219"},
	{"10200", 1, "200"},
	{"2000", 1, "0"},
	{"112", 1, "11"},
	{"1173", 2, "11"},
}

func TestRemoveKdigits(t *testing.T) {
	for _, test := range tests {
		if got := removeKdigits(test.num, test.k); got != test.want {
			t.Errorf("removeKdigits(%v, %v) = %v; want %v", test.num, test.k, got, test.want)
		}
	}
}
