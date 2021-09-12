package main

import (
	"fmt"
	"sort"
)

func main() {
	outerTrees([][]int{
		{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2},
	})

	fmt.Println(checkOnLine([]int{2, 0}, []int{2, 2}, []int{2, 4}))
}

func outerTrees(trees [][]int) [][]int {
	n := len(trees)
	k := 0
	stack := make([][]int, 2*n)
	sort.Slice(trees, func(i, j int) bool {
		if trees[i][0] == trees[j][0] {
			return trees[i][1] < trees[j][1]
		}
		return trees[i][0] < trees[j][0]
	})
	fmt.Println(trees)
	for i := 0; i < n; i++ {
		for k >= 2 && checkOnLine(stack[k-2], stack[k-1], trees[i]) < 0 {
			k--
		}
		stack[k] = trees[i]
		k++
	}
	fmt.Println(stack, k)
	for i := n - 2; i >= 0; i-- {
		for k >= 2 && checkOnLine(stack[k-2], stack[k-1], trees[i]) < 0 {
			k--
		}
		stack[k] = trees[i]
		k++
	}
	return stack[:k-1]
}

func checkOnLine(o, a, b []int) int {
	return (a[0]-o[0])*(b[1]-o[1]) - (a[1]-o[1])*(b[0]-o[0])
}
