package main

import "fmt"

func sortedSquaresInPlace(A []int) []int {
	fmt.Printf("!!! A = %v\n", A)
	abs := func(n int) int {
		y := n >> 31
		return (n ^ y) - y
	}

	minimumPointer, minimum := 0, abs(A[0])
	for i := 0; i < len(A); i++ {
		current := abs(A[i])
		if current <= minimum {
			minimumPointer, minimum = i, current
		}
	}

	leftPointer, rightPointer := 0, minimumPointer
	for leftPointer < rightPointer {
		A[leftPointer], A[rightPointer] = A[rightPointer], A[leftPointer]
		leftPointer++
		rightPointer--
	}
	fmt.Printf("    A = %v\n", A)

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
