package main

func maxPower(s string) int {
	max := 0
	candidate, count := s[0], 0
	for i := 0; i < len(s); i++ {
		if s[i] == candidate {
			count++
		} else {
			candidate = s[i]
			count = 1
		}

		if count > max {
			max = count
		}
	}

	return max
}
