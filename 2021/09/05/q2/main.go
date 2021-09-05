package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numberOfWeakCharacters([][]int{
		// {5, 5}, {6, 3}, {3, 6},
		// {2, 2}, {3, 3},
		// {1, 5}, {10, 4}, {4, 3}, {4, 2},
		{1, 1}, {1, 2}, {2, 1}, {2, 2},
	}))
}
func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i int, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] > properties[j][1]
		}
		return properties[i][0] > properties[j][0]
	})
	dp := make([]int, len(properties))
	arr := make([]int, 0)
	m := make(map[int]int)
	m[properties[0][0]] = properties[0][1]
	arr = append(arr, properties[0][0])
	pre := make(map[int]int)
	two := []int{properties[0][1], properties[0][1]}
	for i := 1; i < len(properties); i++ {
		_, ok := m[properties[i][0]]
		if !ok {
			m[properties[i][0]] = properties[i][1]
			if properties[i][1] > two[0] {
				two[0], two[1] = properties[i][1], two[0]
			} else if properties[i][1] > two[1] {
				two[1] = properties[i][1]
			}
		}
		if properties[i][0] < properties[i-1][0] && properties[i][1] < properties[i-1][1] {
			dp[i] = dp[i-1] + 1
			continue
		}

		if pre[properties[i][0]] > 0 {
			dp[i] = dp[i-1] + 1
			continue
		}
		v := m[properties[i][0]]
		if v < two[0] {
			pre[properties[i][0]] = 1
			dp[i] = dp[i-1] + 1
			continue
		}
		if v == two[0] && properties[i][1] < two[1] {
			pre[properties[i][0]] = 1
			dp[i] = dp[i-1] + 1
			continue
		}

		dp[i] = max(dp[i], dp[i-1])
	}
	fmt.Println(properties)
	fmt.Println(dp)
	return dp[len(dp)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
