package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	var dictionary = make(map[rune]int)
	for _, letter := range magazine {
		dictionary[letter]++
	}

	for _, letter := range ransomNote {
		if dictionary[letter] == 0 {
			return false
		}

		dictionary[letter]--
	}
	return true
}
