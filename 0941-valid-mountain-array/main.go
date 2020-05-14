package main

func validMountainArray(A []int) bool {
	last := len(A) - 1
	left := 0
	for left < last && A[left] < A[left+1] {
		left++
	}
	if left == 0 {
		return false
	}

	right := left
	for right < last && A[right] > A[right+1] {
		right++
	}

	return left != right && right == last
}
