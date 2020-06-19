package main

func longestDupSubstring(S string) string {
	length := len(S)
	if length == 1 {
		return ""
	}

	left, right := 1, length-1
	foundSubstring := ""
	for left <= right {
		middle := left + (right-left)/2
		if s := searchDupSubstring(S, middle); s == "" {
			right = middle - 1
		} else {
			foundSubstring = s
			left = middle + 1
		}
	}

	return foundSubstring
}

func searchDupSubstring(S string, width int) string {
	substrs := make(map[string]bool)
	for i := 0; i <= len(S)-width; i++ {
		substr := S[i : i+width]
		_, ok := substrs[substr]
		if ok {
			return substr
		}
		substrs[substr] = true
	}

	return ""
}
