package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func intervalIntersection(A [][]int, B [][]int) (result [][]int) {
	lengthA, lengthB := len(A), len(B)
	i, j := 0, 0
	for i < lengthA && j < lengthB {
		lo := max(A[i][0], B[j][0])
		hi := min(A[i][1], B[j][1])
		if lo <= hi {
			result = append(result, []int{lo, hi})
		}
		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
		}
	}
	return
}

func intersection(a, b []int) []int {
	if a[0] <= b[0] && a[1] >= b[1] {
		return b
	}
	if a[0] >= b[0] && a[1] <= b[1] {
		return a
	}
	if a[0] >= b[0] && a[0] <= b[1] && a[1] >= b[1] {
		return []int{a[0], b[1]}
	}
	if a[0] <= b[0] && a[1] >= b[0] && a[1] <= b[1] {
		return []int{b[0], a[1]}
	}

	return nil
}

func intervalIntersectionSlow(A [][]int, B [][]int) (result [][]int) {
	lengthA, lengthB := len(A), len(B)
	i, j := 0, 0
	prevJ := 0
	for i < lengthA {
		j = prevJ
		for j < lengthB && B[j][0] <= A[i][1] {
			x := intersection(A[i], B[j])
			if x != nil {
				result = append(result, x)
			}

			prevJ = j
			j++
		}
		i++
	}

	return
}
