package palindromenumber

import "testing"

var tests = []struct {
	x          int
	palindrome bool
}{
	{121, true},
	{-121, false},
	{10, false},
	{7, true},
	{0, true},
	{9223372036854775807, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range tests {
		if got := isPalindrome(test.x); got != test.palindrome {
			t.Errorf("isPalindrome(%d) = %v; want %v", test.x, got, test.palindrome)
		}
	}
}

func TestIsPalindrom2(t *testing.T) {
	for _, test := range tests {
		if got := isPalindrome2(test.x); got != test.palindrome {
			t.Errorf("isPalindrome2(%d) = %v; want %v", test.x, got, test.palindrome)
		}
	}
}
