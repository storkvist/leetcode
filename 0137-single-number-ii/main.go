package main

func singleNumber(nums []int) int {
	p, q := 0, 0
	for _, n := range nums {
		p = q & (p ^ n)
		q = p | (q ^ n)
	}
	return q
}
