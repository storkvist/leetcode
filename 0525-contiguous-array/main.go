package main

func findMaxLength(nums []int) int {
	sumPositions := make(map[int]int)
	currentSum, maxLength := 0, 0
	for i := range nums {
		currentSum += nums[i]
		if nums[i] == 0 {
			currentSum--
		}
		if currentSum == 0 {
			maxLength = i + 1
		}

		if position, ok := sumPositions[currentSum]; ok {
			if maxLength < i-position {
				maxLength = i - position
			}
		} else {
			sumPositions[currentSum] = i
		}
	}
	return maxLength
}
