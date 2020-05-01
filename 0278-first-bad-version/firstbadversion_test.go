package firstbadversion

import "testing"

var tests = []struct {
	bad      int
	versions int
}{
	{1, 5},
	{2, 5},
	{3, 5},
	{4, 5},
	{5, 5},
	{1, 6},
	{4, 6},
	{6, 6},
}

func TestFirstBadVersion(t *testing.T) {
	for _, test := range tests {
		badVersion = test.bad
		if got := firstBadVersion(test.versions); badVersion != got {
			t.Errorf("firstBadVersion(%v) = %v; want %v", test.versions, got, badVersion)
		}
	}
}
