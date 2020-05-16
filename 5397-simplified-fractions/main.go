package main

import "fmt"

func nod(a, b int) int {
	if b == 0 {
		return a
	}

	return nod(b, a%b)
}

func simplifiedFractions(n int) []string {
	result := []string{}
	for i := 1; i < n; i++ {
		for j := i - 1; j <= n; j++ {
			if j > i && nod(i, j) == 1 {
				result = append(result, fmt.Sprintf("%d/%d", i, j))
			}
		}
	}

	return result
}
