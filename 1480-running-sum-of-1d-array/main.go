package main

func runningSum(nums []int) []int {
	prev := 0
	for i := range nums {
		prev, nums[i] = prev+nums[i], prev+nums[i]
	}
	return nums
}
