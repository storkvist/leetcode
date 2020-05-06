package main

func majorityElement(nums []int) int {
	var candidate int
	var sum int
	for _, num := range nums {
		if sum == 0 {
			candidate = num
		}

		if num == candidate {
			sum++
		} else {
			sum--
		}
	}

	return candidate
}

func majorityElementHashMap(nums []int) int {
	var dictionary = make(map[int]int)
	for _, num := range nums {
		dictionary[num]++
	}

	majorCount := len(nums)/2 + len(nums)%2
	for num, count := range dictionary {
		if count >= majorCount {
			return num
		}
	}
	return 0
}
