package main

import "math"

func numSquares(n int) int {
	length := n
	if length <= 3 {
		length = 3
	}

	dp := make([]int, length+1)

	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	for i := 4; i <= length; i++ {
		dp[i] = i

		for x := 1; x <= int(math.Ceil(math.Sqrt(float64(i)))); x++ {
			temp := x * x
			if temp > i {
				break
			}

			dp[i] = min(dp[i], 1+dp[i-temp])
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
