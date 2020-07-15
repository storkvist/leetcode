package main

import "testing"

var tests = []struct {
	hour    int
	minutes int
	want    float64
}{
	{12, 30, 165},
	{3, 30, 75},
	{3, 15, 7.5},
	{4, 50, 155},
	{12, 0, 0},
}

func TestAngleClock(t *testing.T) {
	for _, test := range tests {
		if got := angleClock(test.hour, test.minutes); got != test.want {
			t.Errorf("angleClock(%v, %v) = %v; want %v\n", test.hour, test.minutes, got, test.want)
		}
	}
}
