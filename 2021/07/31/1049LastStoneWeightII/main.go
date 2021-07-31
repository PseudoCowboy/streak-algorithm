package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40}))
}

func lastStoneWeightII(stones []int) int {
	sort.Ints(stones)
	for len(stones) > 1 {
		temp := make([]int, 0)
		for i := len(stones) - 1; i >= 1; i -= 2 {
			if stones[i]-stones[i-1] > 0 {
				temp = append(temp, stones[i]-stones[i-1])
			}
		}
		if len(stones)%2 == 1 {
			temp = append(temp, stones[len(stones)-1])
		}
		stones = temp
		sort.Ints(stones)
	}
	if len(stones) > 0 {
		return stones[0]
	}
	return 0
}
