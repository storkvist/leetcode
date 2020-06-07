package main

import "testing"

var tests = []struct {
	amount int
	coins  []int
	want   int
}{
	{5, []int{1, 2, 5}, 4},
	{3, []int{2}, 0},
	{10, []int{10}, 1},
	{0, []int{1, 2, 3}, 1},
	{5, []int{2, 5}, 1},
	{0, []int{}, 1},
	{7, []int{}, 0},
	{1000, []int{3, 5, 7, 8, 9, 10, 11}, 1952879228},
}

func TestChange(t *testing.T) {
	for _, test := range tests {
		if got := change(test.amount, test.coins); got != test.want {
			t.Errorf("change(%d, %v) = %d; want %d\n", test.amount, test.coins, got, test.want)
		}
	}
}
