package main

func sign(s string) (result [26]int) {
	for _, c := range s {
		result[c-'a']++
	}
	return
}

func diff(a, b [26]int) (result [26]int) {
	for i, n := range a {
		result[i] = n - b[i]
	}
	return
}

func empty(a [26]int) bool {
	for _, n := range a {
		if n != 0 {
			return false
		}
	}

	return true
}

func findAnagrams(s string, p string) []int {
	length, width := len(s), len(p)
	if length == 0 || length < width {
		return []int{}
	}

	var result []int

	x1 := sign(s[:width])
	x2 := sign(p)
	buffer := diff(x1, x2)
	left := s[0]
	for i := 0; i < length-width+1; i++ {
		right := s[i+width-1]
		if i > 0 {
			buffer[left-'a'] -= 1
			buffer[right-'a'] += 1
			left = s[i]
		}

		if empty(buffer) {
			result = append(result, i)
		}
	}

	return result
}
