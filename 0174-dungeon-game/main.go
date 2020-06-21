package main

func calculateMinimumHP(dungeon [][]int) int {
	height := len(dungeon)
	if height == 0 {
		return 1
	}
	width := len(dungeon[0])
	if width == 0 {
		return 1
	}

	dp := make([][]int, height)
	for i := range dp {
		dp[i] = make([]int, width)
	}

	if dungeon[height-1][width-1] > 0 {
		dp[height-1][width-1] = 1
	} else {
		dp[height-1][width-1] = 1 - dungeon[height-1][width-1]
	}

	for i := height - 2; i >= 0; i-- {
		dp[i][width-1] = max(dp[i+1][width-1]-dungeon[i][width-1], 1)
	}
	for i := width - 2; i >= 0; i-- {
		dp[height-1][i] = max(dp[height-1][i+1]-dungeon[height-1][i], 1)
	}

	for i := height - 2; i >= 0; i-- {
		for j := width - 2; j >= 0; j-- {
			minCostOnExit := min(dp[i+1][j], dp[i][j+1])
			dp[i][j] = max(minCostOnExit-dungeon[i][j], 1)
		}

	}

	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
