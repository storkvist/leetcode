package main

import "testing"

var tests = []struct {
	prices []int
	want   []int
}{
	{[]int{100, 80, 60, 70, 60, 75, 85}, []int{1, 1, 1, 2, 1, 4, 6}},
	{[]int{2, 2, 2}, []int{1, 2, 3}},
	{[]int{}, []int{}},
	{[]int{100}, []int{1}},
	{[]int{5, 2, 2, 2, 1}, []int{1, 1, 2, 3, 1}},
	{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{[]int{5, 4, 3, 2, 1}, []int{1, 1, 1, 1, 1}},
	{[]int{1, 2, 2, 2, 1, 2, 2, 2, 1}, []int{1, 2, 3, 4, 1, 6, 7, 8, 1}},
	{[]int{1, 2, 2, 2, 1}, []int{1, 2, 3, 4, 1}},
}

func TestStockSpanner_Next(t *testing.T) {
	for _, test := range tests {
		spanner := Constructor()
		var results []int
		for i := range test.prices {
			got := spanner.Next(test.prices[i])
			results = append(results, got)
			if got != test.want[i] {
				t.Errorf("StockSpanner(%v) -> %v; want %v", test.prices[:i+1], results, test.want[:i+1])
				break
			}
		}
	}
}

func TestStockSpannerSlow_Next(t *testing.T) {
	for _, test := range tests {
		spanner := ConstructorSlow()
		var results []int
		for i := range test.prices {
			got := spanner.Next(test.prices[i])
			results = append(results, got)
			if got != test.want[i] {
				t.Errorf("StockSpannerSlow(%v) -> %v; want %v", test.prices[:i+1], results, test.want[:i+1])
				break
			}
		}
	}
}

func TestStockSpannerSlow2_Next(t *testing.T) {
	for _, test := range tests {
		spanner := ConstructorSlow2()
		var results []int
		for i := range test.prices {
			got := spanner.Next(test.prices[i])
			results = append(results, got)
			if got != test.want[i] {
				t.Errorf("StockSpannerSlow2(%v) -> %v; want %v", test.prices[:i+1], results, test.want[:i+1])
				break
			}
		}
	}
}

func BenchmarkStockSpanner_Next(b *testing.B) {
	for i := 0; i < b.N; i++ {
		spanner := Constructor()
		for _, x := range tests[0].prices {
			spanner.Next(x)
		}
	}
}

func BenchmarkStockSpannerSlow_Next(b *testing.B) {
	for i := 0; i < b.N; i++ {
		spanner := ConstructorSlow()
		for _, x := range tests[0].prices {
			spanner.Next(x)
		}
	}
}

func BenchmarkStockSpannerSlow2_Next(b *testing.B) {
	for i := 0; i < b.N; i++ {
		spanner := ConstructorSlow2()
		for _, x := range tests[0].prices {
			spanner.Next(x)
		}
	}
}
