package romantointeger

import "testing"

var tests = []struct {
	s string
	n int
}{
	{"III", 3},
	{"IV", 4},
	{"IX", 9},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
}

func TestRomanToInt(t *testing.T) {
	for _, test := range tests {
		if got := romanToInt(test.s); got != test.n {
			t.Errorf("romanToInt(%v) = %v; want %v\n", test.s, got, test.n)
		}
	}
}
