package main

import "fmt"

func main() {
	fmt.Println(missingRolls([]int{4, 5, 6, 2, 3, 6, 5, 4, 6, 4, 5, 1, 6, 3, 1, 4, 5, 5, 3, 2, 3, 5, 3, 2, 1, 5, 4, 3, 5, 1, 5}, 4, 40))
}

func missingRolls(rolls []int, mean int, n int) []int {
	sum := 0
	for i := range rolls {
		sum += rolls[i]
	}
	if n*6+sum < mean*(len(rolls)+n) || n+sum > mean*(len(rolls)+n) {
		return []int{}
	}
	target := mean*(len(rolls)+n) - sum

	each := target / n
	remain := target % n
	ans := make([]int, n)
	for i := range ans {
		ans[i] = each
	}
	for i := 0; i < remain; i++ {
		ans[i]++
	}

	return ans
}
