package main

import (
	"fmt"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	ans := 0
	for i := range seats {
		ans += abs(seats[i] - students[i])
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(winnerOfGame("AAAAABBBBBBAAAAA"))
}

func winnerOfGame(colors string) bool {
	a := 0
	b := 0
	temp := 0
	for i := range colors {
		if colors[i] == 'A' {
			temp += 1
			if temp > 2 {
				a += temp - 2
			}
		} else {
			temp = 0
		}
	}
	temp = 0
	for i := range colors {
		if colors[i] == 'B' {
			temp += 1
			if temp > 2 {
				b += temp - 2
			}
		} else {
			temp = 0
		}
	}

	if a-b > 0 {
		return true
	}
	return false
}

func networkBecomesIdle(edges [][]int, patience []int) int {

}
