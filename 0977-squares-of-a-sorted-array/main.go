package main

func sortedSquares(A []int) []int {
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
