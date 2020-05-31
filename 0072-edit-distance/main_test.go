package main

import "testing"

var tests = []struct {
	word1 string
	word2 string
	want  int
}{
	{"horse", "ros", 3},
	{"intention", "execution", 5},
}

func TestMinDistance(t *testing.T) {
	for _, test := range tests {
		if got := minDistance(test.word1, test.word2); got != test.want {
			t.Errorf("minDistance(%q, %q) = %d; want %d", test.word1, test.word2, got, test.want)
		}
	}
}
