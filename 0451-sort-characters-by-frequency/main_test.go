package main

import "testing"

var tests = []struct {
	s    string
	want []string
}{
	{"Aabb", []string{"bbAa", "bbaA"}},
	{"", []string{""}},
	{"a", []string{"a"}},
	{"aa", []string{"aa"}},
	{"tree", []string{"eetr", "eert"}},
	{"cccaaa", []string{"aaaccc", "cccaaa"}},
}

func TestFrequencySort(t *testing.T) {
	for _, test := range tests {
		got := frequencySort(test.s)
		ok := false
		for _, s := range test.want {
			ok = ok || got == s
		}
		if !ok {
			t.Errorf("frequencySort(%q) = %q; want %q\n", test.s, got, test.want)
		}
	}
}

func BenchmarkFrequencySort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		frequencySort(tests[0].s)
	}
}

func BenchmarkFrequencySort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		frequencySort1(tests[0].s)
	}
}

func BenchmarkFrequencySort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		frequencySort2(tests[0].s)
	}
}
