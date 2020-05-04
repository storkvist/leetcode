package numbercomplement

func findComplement(num int) int {
	if num == 0 {
		return 0
	}

	i := 0
	for num>>i != 0 {
		i++
	}
	return int((^uint(num)) &^ ((1<<63 - 1) << i))
}
