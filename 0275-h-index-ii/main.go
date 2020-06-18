package main

func hIndex(citations []int) int {
	length := len(citations)

	if length == 0 {
		return 0
	}

	if length == 1 {
		if citations[0] == 0 {
			return 0
		} else {
			return 1
		}
	}

	left, right := 0, length-1
	for left < right {
		middle := left + (right-left+1)/2
		if citations[middle] > length-middle {
			right = middle - 1
		} else {
			left = middle
		}
	}

	if citations[right] > length-right {
		return length
	}

	if citations[right] == length-right {
		return length - right
	} else {
		return length - right - 1
	}
}
