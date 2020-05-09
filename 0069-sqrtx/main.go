package main

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	left, right := 0, x
	for left+1 < right {
		middle := left + (right-left)/2
		check := middle * middle

		if check == x {
			return middle
		}

		if check > x {
			right = middle
		} else {
			left = middle
		}
	}

	return left
}
