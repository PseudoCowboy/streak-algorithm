package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numberOfWeeks1([]int{1, 6, 8}))
}

func numberOfWeeks(milestones []int) int64 {
	sort.Ints(milestones)
	max := int64(milestones[len(milestones)-1])
	count := int64(0)
	for len(milestones) > 1 {
		temp := make([]int, 0)
		for i := len(milestones) - 1; i >= 1; i -= 2 {
			if milestones[i]-milestones[i-1] > 0 {
				temp = append(temp, milestones[i]-milestones[i-1])
			}
			count += 2 * int64(milestones[i-1])
		}
		if len(milestones)%2 == 1 {
			temp = append(temp, milestones[0])
		}
		milestones = temp
		sort.Ints(milestones)
	}

	if len(milestones) > 0 {
		if milestones[0] > 1 {
			if max-count/2 > 1 {
				count++
			} else {
				count += 2
			}
		} else {
			count++
		}
	}
	return count
}

func lastStoneWeightII(milestones []int) int64 {
	sort.Ints(milestones)
	maxv := int64(milestones[len(milestones)-1])
	sum := int64(0)
	for _, v := range milestones {
		sum += int64(v)
	}
	target := sum / 2
	dp := make([]int64, target+1)
	for i := 0; i < len(milestones); i++ {
		for j := target; j >= int64(milestones[i]); j-- {
			dp[j] = max(dp[j], dp[j-int64(milestones[i])]+int64(milestones[i]))
		}
	}
	diff := sum - 2*dp[target]
	count := 2 * dp[target]
	if diff > 1 {
		if maxv-count/2 > 1 {
			count++
		} else {
			count += 2
		}
	} else if diff > 0 {
		count++
	}
	return count
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func numberOfWeeks1(milestones []int) int64 {
	sum := int64(0)
	sort.Ints(milestones)
	for _, v := range milestones {
		sum += int64(v)
	}
	maxv := milestones[len(milestones)-1]
	for i := len(milestones) - 2; i >= 0; i-- {
		if maxv-milestones[i] < 0 {
			return sum
		} else {
			maxv -= milestones[i]
		}
	}
	if maxv > 0 {
		return sum - int64(maxv) + 1
	}
	return sum
}
