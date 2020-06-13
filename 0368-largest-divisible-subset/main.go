package main

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	length := len(nums)
	if length == 0 {
		return []int{}
	}

	divCount := make([]int, length)
	for i := range divCount {
		divCount[i] = 1
	}
	prev := make([]int, length)
	for i := range prev {
		prev[i] = -1
	}

	maxIndex := 0
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				if divCount[i] < divCount[j]+1 {
					divCount[i] = divCount[j] + 1
					prev[i] = j
				}
			}
		}

		if divCount[maxIndex] < divCount[i] {
			maxIndex = i
		}
	}

	var result []int
	for maxIndex >= 0 {
		result = append(result, nums[maxIndex])
		maxIndex = prev[maxIndex]
	}

	return result
}
