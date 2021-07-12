package main

import (
	"sort"
)

func findItinerary(tickets [][]string) []string {
	m := map[string][]string{}
	mUsed := map[string][]bool{}
	for i := range tickets {
		m[tickets[i][0]] = append(m[tickets[i][0]], tickets[i][1])
		mUsed[tickets[i][0]] = append(mUsed[tickets[i][0]], false)
	}
	for _, v := range m {
		sort.Strings(v)
	}
	current := "JFK"
	travel := []string{"JKF"}
	dfs(m, mUsed, current, &travel, len(tickets))
	return travel
}

func dfs(m map[string][]string, mUsed map[string][]bool, current string, travel *[]string, left int) bool {
	if left == 0 {
		return true
	}
	if len(m[current]) == 0 {
		return false
	}
	for i, v := range mUsed[current] {
		if v == false {
			mUsed[current][i] = true
			*travel = append(*travel, m[current][i])
			if dfs(m, mUsed, m[current][i], travel, left-1) {
				return true
			}
			mUsed[current][i] = false
			*travel = (*travel)[:len(*travel)-1]
		}
	}
	return false
}
