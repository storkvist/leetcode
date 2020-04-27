package twosum

// TwoSum returns indices of the two numbers from given array such that they add up to a specific target.
func TwoSum(nums []int, target int) []int {
	for i, num1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num1+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// TwoSum2 realizes TwoSum with one-pass hash table.
func TwoSum2(nums []int, target int) []int {
	table := make(map[int]int)
	for i, num := range nums {
		j, ok := table[target-num]
		if ok && i != j {
			return []int{j, i}
		}
		table[num] = i
	}
	return nil
}
