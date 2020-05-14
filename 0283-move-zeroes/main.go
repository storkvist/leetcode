package main

func moveZeroes(nums []int) {
	length := len(nums)
	readPointer, writePointer := 0, 0
	for readPointer < length {
		if nums[readPointer] != 0 {
			nums[writePointer] = nums[readPointer]
			writePointer++
		}

		readPointer++
	}

	for writePointer < length {
		nums[writePointer] = 0
		writePointer++
	}
}
