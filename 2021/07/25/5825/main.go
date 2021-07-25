package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxCompatibilitySum([][]int{
		{1, 1, 1},
		{0, 0, 1},
		{0, 0, 1},
		{0, 1, 0},
	}, [][]int{
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
		{1, 1, 0},
	}))
}

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	n := len(students[0])

	score := 0
	for i := range students {
		ic := 0
		target := 0
		for j := range mentors {
			temp := 0
			if mentors[j][0] == 2 {
				continue
			}
			for k := 0; k < n; k++ {
				if students[i][k] == mentors[j][k] {
					temp++
				}
			}
			if temp > ic {
				ic = temp
				target = j
			}
		}
		fmt.Println(target)
		mentors[target][0] = 2
		score += ic
	}
	return score
}
