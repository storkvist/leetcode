package main

func nthUglyNumber(n int) int {
	ugly := make([]int, n)
	ugly[0] = 1

	i2, i3, i5 := 0, 0, 0
	nextMultipleOf2, nextMultipleOf3, nextMultipleOf5 := 2, 3, 5
	next := 1

	for i := 1; i < n; i++ {
		next = min(nextMultipleOf2, nextMultipleOf3, nextMultipleOf5)
		ugly[i] = next
		if next == nextMultipleOf2 {
			i2++
			nextMultipleOf2 = ugly[i2] * 2
		}
		if next == nextMultipleOf3 {
			i3++
			nextMultipleOf3 = ugly[i3] * 3
		}
		if next == nextMultipleOf5 {
			i5++
			nextMultipleOf5 = ugly[i5] * 5
		}
	}

	return next
}

func min(a, b, c int) int {
	return min2(a, min2(b, c))
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
