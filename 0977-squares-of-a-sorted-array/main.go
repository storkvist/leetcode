package main

func sortedSquares(A []int) []int {
	length := len(A)
	if length == 0 {
		return A
	}

	for i := 0; i < length && A[i] < 0; i++ {
		A[i] = -A[i]
	}

	// Sort numbers using the upper 16 bits of the numbers as storage
	leftPointer, rightPointer := 0, length-1
	for k := length - 1; k >= 0; k-- {
		// Take numbers from lower 16 bits
		left, right := A[leftPointer]&0xffff, A[rightPointer]&0xffff

		if left < right {
			A[k] = A[k] | (right << 16)
			rightPointer--
		} else {
			A[k] = A[k] | (left << 16)
			leftPointer++
		}
	}

	// Calculate squares using the upper 16 bits from the numbers
	for i, x := range A {
		number := x >> 16
		A[i] = number * number
	}

	return A
}

func sortedSquaresExtraMemory(A []int) []int {
	result := make([]int, len(A))

	abs := func(n int) int {
		y := n >> 31
		return (n ^ y) - y
	}

	i, j := 0, len(A)-1
	current := j
	for i <= j {
		if abs(A[i]) > abs(A[j]) {
			result[current] = A[i] * A[i]
			i++
		} else {
			result[current] = A[j] * A[j]
			j--
		}

		current--
	}

	return result
}
