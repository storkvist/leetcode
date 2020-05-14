package main

func removeDuplicates(nums []int) int {
	last := len(nums) - 1
	current, result := 0, 0
	for current <= last {
		if current == last || nums[current] != nums[current+1] {
			nums[result], nums[current] = nums[current], nums[result]
			result++
		}

		current++
	}

	return result
}
