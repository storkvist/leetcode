package main

func sortArrayByParity(A []int) []int {
	left, right := 0, len(A)-1
	for left < right {
		if A[left]%2 == 0 {
			left++
			continue
		}
		if A[right]%2 == 1 {
			right--
			continue
		}
		A[left], A[right] = A[right], A[left]
		left++
		right--
	}
	return A
}

func sortArrayByParityExtraMemory(A []int) []int {
	length := len(A)
	result := make([]int, length)
	left, right := 0, length-1
	for _, x := range A {
		if x%2 == 0 {
			result[left] = x
			left++
		} else {
			result[right] = x
			right--
		}
	}

	return result
}
