package main

func findJudge(N int, trust [][]int) int {
	connections := make([]int, N+1)

	for _, pair := range trust {
		connections[pair[0]]--
		connections[pair[1]]++
	}

	for i := 1; i <= N; i++ {
		if connections[i] == N-1 {
			return i
		}
	}

	return -1
}

func findJudgeHashMap(N int, trust [][]int) int {
	if N == 1 {
		return 1
	}

	var candidates []int
	connectionsTo := make(map[int]int)
	connectionsFrom := make(map[int]int)

	for _, pair := range trust {
		candidates = append(candidates, pair[1])
		connectionsFrom[pair[0]]++
		connectionsTo[pair[1]]++
	}

	for _, candidate := range candidates {
		if connectionsTo[candidate] == N-1 && connectionsFrom[candidate] == 0 {
			return candidate
		}
	}

	return -1
}
