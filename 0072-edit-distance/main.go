package main

func minDistance(word1, word2 string) int {
	len1, len2 := len(word1), len(word2)
	var dp [2][]int
	dp[0] = make([]int, len1+1)
	dp[1] = make([]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= len2; i++ {
		for j := 0; j <= len1; j++ {
			if j == 0 {
				dp[i%2][j] = i
			} else if word1[j-1] == word2[i-1] {
				dp[i%2][j] = dp[(i-1)%2][j-1]
			} else {
				dp[i%2][j] = 1 + min(dp[(i-1)%2][j], dp[i%2][j-1], dp[(i-1)%2][j-1])
			}
		}
	}
	return dp[len2%2][len1]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}
