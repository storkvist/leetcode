package main

func firstUniqChar(s string) int {
	var dictionary = make(map[rune]int)
	for _, char := range s {
		dictionary[char]++
	}

	for i, char := range s {
		if dictionary[char] == 1 {
			return i
		}
	}

	return -1
}
