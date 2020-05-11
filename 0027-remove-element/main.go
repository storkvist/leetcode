package main

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)-1
	length := 0
	for left <= right {
		if nums[left] != val {
			left++
			length++
			continue
		}

		if nums[right] == val {
			right--
			continue
		}

		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
		length++
	}

	return length
}
