package main

func maxUncrossedLines(A []int, B []int) int {
	dp := make([][]int, len(A))
	lengthB := len(B)
	for i := range dp {
		dp[i] = make([]int, lengthB)
	}
	return solve(0, 0, A, B, dp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(i, j int, A, B []int, dp [][]int) int {
	lengthA := len(A)
	if i >= lengthA {
		return 0
	}
	lengthB := len(B)
	if j >= lengthB {
		return 0
	}
	if dp[i][j] != 0 {
		return dp[i][j]
	}

	xj := j
	for xj < lengthB && B[xj] != A[i] {
		xj++
	}

	foundEqual := 0
	if xj < lengthB {
		foundEqual = 1
	}

	result := max(solve(i+1, j, A, B, dp), foundEqual+solve(i+1, xj+1, A, B, dp))
	dp[i][j] = result
	return result
}
