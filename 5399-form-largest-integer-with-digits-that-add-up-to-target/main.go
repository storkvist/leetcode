package main

import (
	"fmt"
	"sort"
)

type ByLen []string

func (a ByLen) Len() int {
	return len(a)
}

func (a ByLen) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func (a ByLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func largestNumber(cost []int, target int) string {
	sets := [][9]int{}

	fmt.Printf("cost = %v, target = %v\n", cost, target)
	for i1 := 0; i1 <= target/cost[0]; i1++ {
		for i2 := 0; i2 <= target/cost[1]; i2++ {
			for i3 := 0; i3 <= target/cost[2]; i3++ {
				for i4 := 0; i4 <= target/cost[3]; i4++ {
					for i5 := 0; i5 <= target/cost[4]; i5++ {
						for i6 := 0; i6 <= target/cost[5]; i6++ {
							for i7 := 0; i7 <= target/cost[6]; i7++ {
								for i8 := 0; i8 <= target/cost[7]; i8++ {
									for i9 := 0; i9 <= target/cost[8]; i9++ {
										sum := 0
										sum += i1 * cost[0]
										sum += i2 * cost[1]
										sum += i3 * cost[2]
										sum += i4 * cost[3]
										sum += i5 * cost[4]
										sum += i6 * cost[5]
										sum += i7 * cost[6]
										sum += i8 * cost[7]
										sum += i9 * cost[8]
										if sum == target {
											sets = append(sets, [9]int{i1, i2, i3, i4, i5, i6, i7, i8, i9})
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("sets = %v\n", sets)

	results := []string{}
	for _, set := range sets {
		var result string
		for i := 8; i >= 0; i-- {
			if set[i] > 0 {
				for j := 0; j < set[i]; j++ {
					result = result + string('0'+i+1)
				}
			}
		}
		results = append(results, result)
	}

	if len(results) == 0 {
		return "0"
	}

	sort.Sort(ByLen(results))

	length := len(results[len(results)-1])
	i := len(results) - 1
	longest := []string{}
	for i >= 0 && len(results[i]) == length {
		longest = append(longest, results[i])
		i--
	}

	sort.Strings(longest)

	fmt.Printf("results = %v\n", results)

	return longest[len(longest)-1]
}
