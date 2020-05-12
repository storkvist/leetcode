package main

func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		middle := left + (right-left)/2

		if middle%2 == 0 {
			if nums[middle] == nums[middle+1] {
				left = middle + 2
			} else {
				right = middle
			}
		} else {
			if nums[middle] == nums[middle-1] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}

	return nums[left]
}
