package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findContentChildren([]int{10, 9, 8, 7}, []int{10, 9, 8, 7}))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)

	sort.Ints(s)
	kid := 0
	for i := 0; i < len(s); i++ {
		if kid == len(g) {
			break
		}
		if s[i] >= g[kid] {
			kid++
		}
	}
	return kid
}

func findContentChildren1(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gi, si := 0, 0
	for gi < len(g) && si < len(s) {
		if s[si] >= g[gi] {
			si++
			gi++
		} else {
			si++
		}
	}
	return gi
}
