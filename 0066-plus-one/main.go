package main

func plusOne(digits []int) []int {
	length := len(digits)
	result := make([]int, length+1)

	extraOne := false
	for i := length; i > 0; i-- {
		result[i] = digits[i-1]
		if i == length {
			result[i]++
		}

		if extraOne {
			result[i]++
			extraOne = false
		}

		if result[i] > 9 {
			result[i] = 0
			extraOne = true
		}
	}

	if extraOne {
		result[0] = 1
	} else {
		return result[1:]
	}

	return result
}
