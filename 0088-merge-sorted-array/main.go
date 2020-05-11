package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	fmt.Printf("!!! nums1 = %v, m = %v\n", nums1, m)
	fmt.Printf("!!! nums2 = %v, n = %v\n", nums2, n)

	n1, n2 := m-1, n-1
	for n1 >= 0 || n2 >= 0 {
		fmt.Printf("nums1 = %v, n1 = %v\n", nums1, n1)
		fmt.Printf("nums2 = %v, n2 = %v\n", nums2, n2)
		position := n1 + n2 + 1

		if n1 < 0 {
			nums1[position] = nums2[n2]
			n2--
			continue
		}

		if n2 < 0 {
			nums1[position] = nums1[n1]
			n1--
			continue
		}

		if nums1[n1] > nums2[n2] {
			nums1[position] = nums1[n1]
			n1--
		} else {
			nums1[position] = nums2[n2]
			n2--
		}
	}
}
