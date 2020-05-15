package main

func kadane(A []int) int {
	current, best := 0, A[0]
	for _, x := range A {
		if current > 0 {
			current = current + x
		} else {
			current = x
		}

		if current > best {
			best = current
		}
	}

	return best
}

func maxSubarraySumCircular(A []int) int {
	max := kadane(A)

	sum := 0
	allNegative, minimumNegative := true, 0
	for i, x := range A {
		allNegative = allNegative && (x < 0)
		if allNegative && x < minimumNegative {
			minimumNegative = x
		}

		sum += x
		A[i] = -x
	}

	var min int
	if allNegative {
		min = minimumNegative
	} else {
		min = kadane(A)
	}
	other := sum + min

	if max > other {
		return max
	}

	return other
}

func maxSubarraySumCircularSlow(A []int) int {
	length := len(A)
	maximum := A[0]
	for startPointer := 0; startPointer < length; startPointer++ {
		sum := 0
		for endPointer := startPointer; endPointer-startPointer < length; endPointer++ {

			sum += A[endPointer%length]
			if sum > maximum {
				maximum = sum
			}
		}
	}

	return maximum
}

//func maxSubarraySumCircularSuperSlow(A []int) int {
//	length := len(A)
//	maximum := A[0]
//	for startPointer := 0; startPointer < length; startPointer++ {
//		for endPointer := startPointer; endPointer-startPointer < length; endPointer++ {
//			sum := 0
//			for i := startPointer; i <= endPointer; i++ {
//				sum += A[i%length]
//			}
//
//			if sum > maximum {
//				maximum	= sum
//			}
//		}
//	}
//
//	return maximum
//}
