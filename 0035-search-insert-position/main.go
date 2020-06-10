package main

func searchInsert(nums []int, target int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}

	if length == 1 {
		if nums[0] < target {
			return 1
		} else {
			return 0
		}
	}

	middle := length / 2
	if target > nums[middle] {
		return middle + searchInsert(nums[middle:], target)
	}
	return searchInsert(nums[:middle], target)
}
