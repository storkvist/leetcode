package main

import "testing"

var tests = []struct {
	n    int
	k    int
	want string
}{
	{3, 3, "213"},
	{4, 9, "2314"},
}

func TestGetPermutation(t *testing.T) {
	for _, test := range tests {
		if got := getPermutation(test.n, test.k); got != test.want {
			t.Errorf("getPermutation(%v, %v) = %q; want %q\n", test.n, test.k, got, test.want)
		}
	}
}
