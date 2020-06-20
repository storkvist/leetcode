package main

import "strconv"

func getPermutation(n int, k int) string {
	var numbers []string
	mod := 1
	for i := 1; i <= n; i++ {
		numbers = append(numbers, strconv.Itoa(i))
		mod *= i
	}

	result := ""
	k--
	for i := 0; i < n; i++ {
		mod /= n - i
		currentIndex := k / mod
		k %= mod
		result += numbers[currentIndex]
		copy(numbers[currentIndex:], numbers[currentIndex+1:])
		numbers = numbers[:len(numbers)-1]
	}

	return result
}
