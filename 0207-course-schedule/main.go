package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	indegrees := make(map[int]int)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}
	for _, pair := range prerequisites {
		graph[pair[0]] = append(graph[pair[0]], pair[1])
		indegrees[pair[1]]++
	}

	var isolatedNodes []int
	for node := range graph {
		if indegrees[node] == 0 {
			isolatedNodes = append(isolatedNodes, node)
		}
	}

	var topologicalOrdering []int
	for len(isolatedNodes) > 0 {
		length := len(isolatedNodes)
		node := isolatedNodes[length-1]
		isolatedNodes = isolatedNodes[:length-1]
		topologicalOrdering = append(topologicalOrdering, node)

		for _, neighbor := range graph[node] {
			indegrees[neighbor]--
			if indegrees[neighbor] == 0 {
				isolatedNodes = append(isolatedNodes, neighbor)
			}
		}
	}

	if len(topologicalOrdering) >= numCourses {
		return true
	}

	return false
}
