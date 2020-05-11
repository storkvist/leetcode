package main

func duplicateZeros(arr []int) {
	duplicates := 0
	last := len(arr) - 1
	for i := 0; i <= last-duplicates; i++ {
		if arr[i] == 0 {
			if i == last-duplicates {
				arr[last] = 0
				last--
				break
			}
			duplicates++
		}
	}

	for i := last - duplicates; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+duplicates] = 0
			duplicates--
			arr[i+duplicates] = 0
		} else {
			arr[i+duplicates] = arr[i]
		}
	}
}

func duplicateZerosExtraMemory(arr []int) {
	limit := len(arr)
	reference := make([]int, limit)
	copy(reference, arr)
	i, j := 0, 0
	for i < limit {
		if reference[j] == 0 {
			arr[i] = 0

			if i+1 < limit {
				arr[i+1] = 0
			}
			i++
		} else {
			arr[i] = reference[j]
		}

		i++
		j++
	}
}
