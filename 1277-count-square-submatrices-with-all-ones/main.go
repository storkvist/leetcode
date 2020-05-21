package main

func min(a, b, c int) int {
	ab := a
	if b < a {
		ab = b
	}

	if ab < c {
		return ab
	}

	return c
}

func countSquares(matrix [][]int) int {
	sum := 0
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var down, left, right int
			if i > 0 && j > 0 {
				down = matrix[i-1][j-1]
			}
			if i > 0 {
				left = matrix[i-1][j]
			}
			if j > 0 {
				right = matrix[i][j-1]
			}
			matrix[i][j] *= min(down, left, right) + 1
			sum += matrix[i][j]
		}
	}

	return sum
}
