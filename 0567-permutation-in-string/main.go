package main

func checkInclusion(s1 string, s2 string) bool {
	length1, length2 := len(s1), len(s2)
	if length1 > length2 {
		return false
	}

	var s1map, s2map [26]int
	for i := range s1 {
		s1map[s1[i]-'a']++
		s2map[s2[i]-'a']++
	}

	count := 0
	for i := 0; i < 26; i++ {
		if s1map[i] == s2map[i] {
			count++
		}
	}

	for i := 0; i < length2-length1; i++ {
		if count == 26 {
			return true
		}

		left := s2[i] - 'a'
		right := s2[i+length1] - 'a'

		s2map[right]++
		if s2map[right] == s1map[right] {
			count++
		} else if s2map[right] == s1map[right]+1 {
			count--
		}

		s2map[left]--
		if s2map[left] == s1map[left] {
			count++
		} else if s2map[left] == s1map[left]-1 {
			count--
		}
	}

	return count == 26
}

func isEmpty(a map[byte]int) bool {
	if len(a) == 0 {
		return true
	}
	for _, n := range a {
		if n != 0 {
			return false
		}
	}
	return true
}

func checkInclusionSlow(s1 string, s2 string) bool {
	length1, length2 := len(s1), len(s2)
	if length1 > length2 {
		return false
	}

	diff := make(map[byte]int)
	for i := range s1 {
		diff[s1[i]]++
	}

	for i, j := -length1, 0; j < length2; i, j = i+1, j+1 {
		if i >= 0 {
			diff[s2[i]]++
		}
		diff[s2[j]]--

		if isEmpty(diff) {
			return true
		}
	}

	return false
}
