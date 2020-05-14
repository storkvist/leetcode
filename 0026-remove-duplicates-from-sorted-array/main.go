package main

func removeDuplicates(nums []int) int {
	last := len(nums) - 1
	readPointer, writePointer := 0, 0
	for readPointer <= last {
		if readPointer == last || nums[readPointer] != nums[readPointer+1] {
			nums[writePointer], nums[readPointer] = nums[readPointer], nums[writePointer]
			writePointer++
		}

		readPointer++
	}

	return writePointer
}
