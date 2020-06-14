package main

const NotProcessed = -2

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	if n == 1 {
		return 0
	}

	var dp [][]int
	dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = -1
		for j := 1; j < n+1; j++ {
			dp[i][j] = NotProcessed
		}
	}
	for i := range dp[src] {
		dp[src][i] = 0
	}

	paths := make(map[int]map[int]int)
	for _, flight := range flights {
		psrc, pdst, pw := flight[0], flight[1], flight[2]
		_, ok := paths[pdst]
		if !ok {
			paths[pdst] = make(map[int]int)
		}
		paths[pdst][psrc] = pw
	}

	return length(&dp, &paths, dst, K+1)
}

func length(dp *[][]int, paths *map[int]map[int]int, dst, K int) int {
	if K < 0 {
		return -1
	}

	if (*dp)[dst][K] != NotProcessed {
		return (*dp)[dst][K]
	}

	var options []int
	for src, w := range (*paths)[dst] {
		l := length(dp, paths, src, K-1)
		(*dp)[src][K-1] = l
		if l > -1 {
			options = append(options, l+w)
		}
	}

	if len(options) == 0 {
		return -1
	}

	return min(options...)
}

func min(nums ...int) int {
	m := nums[0]
	for i, length := 1, len(nums); i < length; i++ {
		if nums[i] < m {
			m = nums[i]
		}
	}
	return m
}
