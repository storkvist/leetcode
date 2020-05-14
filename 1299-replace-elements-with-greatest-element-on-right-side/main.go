package main

func replaceElements(arr []int) []int {
	maximum := -1
	for i := len(arr) - 1; i >= 0; i-- {
		current := arr[i]
		arr[i] = maximum

		if current > maximum {
			maximum = current
		}
	}

	return arr
}

func replaceElementsSlow(arr []int) []int {
	last := len(arr) - 1
	for i := 0; i <= last-1; i++ {
		max := arr[i+1]
		for j := i + 1; j <= last; j++ {
			if arr[j] > max {
				max = arr[j]
			}
		}

		arr[i] = max
	}
	arr[last] = -1

	return arr
}
