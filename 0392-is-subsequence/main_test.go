package main

import "testing"

var tests = []struct {
	s    string
	t    string
	want bool
}{
	{"abc", "ahbgdc", true},
	{"axc", "ahbgdc", false},
	{"", "", true},
	{"abc", "", false},
	{"", "abc", true},
}

func TestIsSubsequence(t *testing.T) {
	for _, test := range tests {
		if got := isSubsequence(test.s, test.t); got != test.want {
			t.Errorf("isSubsequence(%q, %q) = %v; want %v\n", test.s, test.t, got, test.want)
		}
	}
}
