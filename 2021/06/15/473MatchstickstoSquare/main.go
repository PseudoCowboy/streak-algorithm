package main

import (
	"sort"
)

func main() {
	makesquare([]int{5, 5, 5, 5, 16, 4, 4, 4, 4, 4, 3, 3, 3, 3, 4})
}

func makesquare(matchsticks []int) bool {
	sum := 0
	max := 0
	used := make([]bool, len(matchsticks))
	for i := range matchsticks {
		sum += matchsticks[i]
		if matchsticks[i] > max {
			max = matchsticks[i]
		}
	}
	if sum%4 != 0 || sum/4 < max {
		return false
	}
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})
	return dfs(matchsticks, &used, 4, 0, sum/4, 0)
}

func dfs(matchsticks []int, used *[]bool, remainSides int, currentValue, sideValue, nextIndex int) bool {
	if remainSides == 0 {
		return true
	}

	if currentValue == sideValue {
		return dfs(matchsticks, used, remainSides-1, 0, sideValue, 0)
	}

	for i := nextIndex; i < len(*used); i++ {
		if !(*used)[i] && matchsticks[i]+currentValue <= sideValue {
			(*used)[i] = true
			if dfs(matchsticks, used, remainSides, matchsticks[i]+currentValue, sideValue, i+1) {
				return true
			}
			(*used)[i] = false
		}
	}
	return false
}
