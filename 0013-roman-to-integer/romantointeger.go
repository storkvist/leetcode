package romantointeger

func romanToInt(s string) int {
	var digits = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result, prevNumber := 0, 0
	for _, c := range s {
		currentNumber := digits[c]
		if currentNumber > prevNumber {
			result += currentNumber - 2*prevNumber
		} else {
			result += currentNumber
		}

		prevNumber = currentNumber
	}
	return result
}
