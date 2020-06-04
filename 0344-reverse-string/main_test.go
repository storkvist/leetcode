package main

import (
	"testing"
)

var tests = []struct {
	s    []byte
	want []byte
}{
	{[]byte("hello"), []byte("olleh")},
	{[]byte("Hannah"), []byte("hannaH")},
	{[]byte("a"), []byte("a")},
	{[]byte{}, []byte{}},
}

func TestReverseString(t *testing.T) {
	for _, test := range tests {
		s := make([]byte, len(test.s))
		copy(s, test.s)
		reverseString(s)
		if len(s) != len(test.want) {
			t.Errorf("reverseString(%v) -> %v; want %v\n", test.s, s, test.want)
			continue
		}
		for i := range test.want {
			if s[i] != test.want[i] {
				t.Errorf("reverseString(%v) -> %v; want %v\n", test.s, s, test.want)
				break
			}
		}
	}
}
