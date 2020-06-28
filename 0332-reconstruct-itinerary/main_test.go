package main

import "testing"

var tests = []struct {
	tickets [][]string
	want    []string
}{
	{
		[][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
		[]string{"JFK", "MUC", "LHR", "SFO", "SJC"},
	},
	// {
	// 	[][]string{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}},
	// 	[]string{"JFK","ATL","JFK","SFO","ATL","SFO"},
	// },
}

func TestFindItinerary(t *testing.T) {
	for _, test := range tests {
		got := findItinerary(test.tickets)
		if len(got) != len(test.want) {
			t.Errorf("findItinerary(%v) = %v; want %v\n", test.tickets, got, test.want)
			continue
		}
		for i := range test.want {
			if got[i] != test.want[i] {
				t.Errorf("findItinerary(%v) = %v; want %v\n", test.tickets, got, test.want)
				break
			}
		}
	}
}
