package main

import "math"

func findNumbers(nums []int) int {
	counter := 0
	for _, num := range nums {
		if int(math.Log10(float64(num))+1)%2 == 0 {
			counter++
		}
	}
	return counter
}
