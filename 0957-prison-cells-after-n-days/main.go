package main

func prisonAfterNDays(cells []int, N int) []int {
	N = N % 14
	if N == 0 {
		N = 14
	}

	prev := make([]int, len(cells))
	for N > 0 {
		copy(prev, cells)
		copy(cells, next(prev))
		N--
	}

	return cells
}

func next(cells []int) []int {
	result := make([]int, len(cells))
	result[0] = 0
	result[7] = 0
	for i := 1; i < 7; i++ {
		if cells[i-1] == cells[i+1] {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}
	return result
}
