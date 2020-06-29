package main

func uniquePaths(m int, n int) int {
	count := make([][]int, m)
	for i := range count {
		count[i] = make([]int, n)
		count[i][0] = 1
	}
	for j := 0; j < n; j++ {
		count[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			count[i][j] = count[i-1][j] + count[i][j-1]
		}
	}

	return count[m-1][n-1]
}
