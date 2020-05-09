package main

func isPerfectSquare(num int) bool {
	left, right := 1, num/2+num%2
	for left <= right {
		middle := (left + right) / 2
		middleSqr := middle * middle
		if middleSqr == num {
			return true
		}
		if middleSqr > num {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return false
}

func isPerfectSquareLong(num int) bool {
	for i := num/2 + num%2; i >= 1; i-- {
		if num == i*i {
			return true
		}
	}
	return false
}
