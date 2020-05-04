package main

func bitwiseComplement(N int) int {
	if N == 0 {
		return 0
	}

	i := 0
	for N>>i != 0 {
		i++
	}
	return int((^uint(N)) &^ ((1<<63 - 1) << i))
}
