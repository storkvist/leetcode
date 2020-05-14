package main

func rec(str string, n int, result *string) {
	if n == 0 {
		*result += str
		return
	}

	limit := len(str)
	if limit <= n {
		return
	}

	min := 0
	for i := 1; i <= n; i++ {
		if str[i] < str[min] {
			min = i
		}
	}

	*result += string(str[min])

	rec(str[min+1:], n-min, result)
}

func removeKdigits(num string, k int) string {
	result := ""
	rec(num, k, &result)

	head, limit := 0, len(result)
	for head < limit && result[head] == '0' {
		head++
	}

	if head == limit {
		return "0"
	}

	return result[head:]
}

func removeKdigitsPassThrough(num string, k int) string {
	if len(num) <= k {
		return "0"
	}

	for i := 0; i < k; i++ {
		last := len(num) - 1
		j := 0
		for j < last {
			if num[j] > num[j+1] {
				break
			}
			j++
		}
		if j == last {
			num = num[:j]
		} else {
			num = num[:j] + num[j+1:]
		}
	}

	head, limit := 0, len(num)
	for head < limit && num[head] == '0' {
		head++
	}

	if head == limit {
		return "0"
	}

	return num[head:]
}
