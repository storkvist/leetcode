package jewelsandstones

import "testing"

var tests = []struct {
	J   string
	S   string
	num int
}{
	{"aA", "aAAbbbb", 3},
	{"z", "ZZ", 0},
}

func TestNumJewelsInStones(t *testing.T) {
	for _, test := range tests {
		if got := NumJewelsInStones(test.J, test.S); got != test.num {
			t.Errorf("NumJewelsInStones(%s, %s) = %d; want %d\n", test.J, test.S, got, test.num)
		}
	}
}
