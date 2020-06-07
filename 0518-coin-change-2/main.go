package main

func change(amount int, coins []int) int {
	table := make([]int, amount+1)
	table[0] = 1

	for i, length := 0, len(coins); i < length; i++ {
		for j := coins[i]; j <= amount; j++ {
			table[j] += table[j-coins[i]]
		}
	}

	return table[amount]
}
