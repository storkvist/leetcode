package palindromenumber

import (
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	size := int(math.Log10(float64(x))) + 1
	middle := size/2 + size%2
	digits := make([]int, middle)
	for i := 0; i < middle; i++ {
		digits[i] = x % 10
		x = (x - digits[i]) / 10
	}

	ok := true
	var digit int
	for i := middle; x > 0; i++ {
		digit = x % 10
		ok = ok && (digit == digits[size-i-1])
		x = (x - digit) / 10
	}

	return ok
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}

	reverted := 0
	for x > reverted {
		reverted = reverted*10 + x%10
		x = x / 10
	}

	return (x == reverted) || (x == reverted/10)
}
