package jewelsandstones

func NumJewelsInStones(J string, S string) (num int) {
	jewels := make(map[rune]int)
	for i, j := range J {
		jewels[j] = i
	}

	for _, rune := range S {
		if _, ok := jewels[rune]; ok {
			num++
		}
	}

	return
}
