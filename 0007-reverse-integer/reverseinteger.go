package reverseinteger

func Reverse(x int) int {
	maximum := 1<<31 - 1

	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	reversed := 0
	for x > 0 {
		digit := x % 10

		if (maximum-digit)/10 < reversed {
			return 0
		}
		reversed = reversed*10 + digit
		x = (x - digit) / 10
	}

	return sign * reversed
}
