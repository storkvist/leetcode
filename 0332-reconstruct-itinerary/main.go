package main

import (
	"sort"
)

var routes map[string][]string
var result []string

func findItinerary(tickets [][]string) []string {
	routes = make(map[string][]string)
	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		_, ok := routes[from]
		if !ok {
			routes[from] = make([]string, 0)
		}
		routes[from] = append(routes[from], to)
	}

	for _, destinations := range routes {
		if len(destinations) > 1 {
			sort.Strings(destinations)
		}
	}

	result = make([]string, 0)
	dfs("JFK")

	return result
}

func dfs(from string) {
	for len(routes[from]) > 0 {
		first := routes[from][0]
		routes[from] = routes[from][1:]
		dfs(first)
	}

	newResult := make([]string, len(result)+1)
	copy(newResult[1:], result)
	newResult[0] = from
	result = newResult
}
