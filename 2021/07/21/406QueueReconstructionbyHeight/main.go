package main

import (
	"sort"
)

func main() {
	reconstructQueue([][]int{
		{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2},
	})
}

func reconstructQueue(people [][]int) [][]int {
	ans := make([][]int, len(people))
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})
	indexes := make([]int, len(people))
	for i := range indexes {
		indexes[i] = i
	}

	for i := 0; i < len(people); i++ {
		ans[indexes[people[i][1]]] = people[i]
		indexes = append(indexes[:people[i][1]], indexes[people[i][1]+1:]...)
	}

	return ans
}
