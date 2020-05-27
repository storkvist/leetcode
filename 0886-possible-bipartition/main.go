package main

type graph map[int][]int

func possibleBipartition(N int, dislikes [][]int) bool {
	graph := make(graph)
	for _, pair := range dislikes {
		graph[pair[0]] = append(graph[pair[0]], pair[1])
		graph[pair[1]] = append(graph[pair[1]], pair[0])
	}

	colors := make([]int, N+1)
	for i := 1; i <= N; i++ {
		if colors[i] == 0 && !paint(colors, i, graph, 1) {
			return false
		}
	}
	return true
}

func paint(colors []int, current int, graph graph, color int) bool {
	colors[current] = color
	for i := 0; i < len(graph[current]); i++ {
		next := graph[current][i]
		if colors[next] == color {
			return false

		}
		if colors[next] == 0 && !paint(colors, next, graph, -1*color) {
			return false
		}
	}

	return true
}
