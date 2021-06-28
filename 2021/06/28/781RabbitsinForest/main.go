package main

import "fmt"

func main() {
	fmt.Println(numRabbits([]int{0, 0, 1, 1, 1}))
}

func numRabbits(answers []int) int {
	if len(answers) == 0 {
		return 0
	}
	m := make(map[int]int)
	sum := 0
	for _, v := range answers {
		if m[v] == 0 {
			m[v] += v
			sum += v + 1
		} else {
			m[v]--
		}
	}

	return sum
}
