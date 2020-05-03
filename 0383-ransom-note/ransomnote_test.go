package ransomnote

import "testing"

var tests = []struct {
	ransomNote string
	magazine   string
	can        bool
}{
	{"a", "a", true},
	{"a", "b", false},
	{"aa", "ab", false},
	{"aa", "aab", true},
	{"ab", "ba", true},
	{"ab", "ab", true},
}

func TestCanConstruct(t *testing.T) {
	for _, test := range tests {
		if got := canConstruct(test.ransomNote, test.magazine); got != test.can {
			t.Errorf("canConstruct(%v, %v) = %v; want %v", test.ransomNote, test.magazine, got, test.can)
		}
	}
}
