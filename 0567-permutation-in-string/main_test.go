package main

import "testing"

var tests = []struct {
	s1   string
	s2   string
	want bool
}{
	{"ab", "eidbaooo", true},
	{"ab", "eidboaoo", false},
	{"abc", "ab", false},
	{"abc", "acb", true},
	{"a", "a", true},
	{"a", "b", false},
	{"ab", "ac", false},
}

func TestCheckInclusion(t *testing.T) {
	for _, test := range tests {
		if got := checkInclusion(test.s1, test.s2); got != test.want {
			t.Errorf("checkInclusion(%v, %v) = %v; want %v", test.s1, test.s2, got, test.want)
		}
	}
}

func TestCheckInclusionSlow(t *testing.T) {
	for _, test := range tests {
		if got := checkInclusionSlow(test.s1, test.s2); got != test.want {
			t.Errorf("checkInclusionSlow(%v, %v) = %v; want %v", test.s1, test.s2, got, test.want)
		}
	}
}

func BenchmarkCheckInclusion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkInclusion(tests[0].s1, tests[0].s2)
	}
}

func BenchmarkCheckInclusionSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkInclusionSlow(tests[0].s1, tests[0].s2)
	}
}
