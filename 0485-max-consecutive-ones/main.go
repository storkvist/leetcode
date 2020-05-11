package main

func findMaxConsecutiveOnes(nums []int) int {
	result, counter := 0, 0
	for _, num := range nums {
		if num == 1 {
			counter++

			if counter > result {
				result = counter
			}
		} else {
			counter = 0
		}
	}

	return result
}
