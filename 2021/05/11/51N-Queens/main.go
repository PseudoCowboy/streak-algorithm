package main

import "fmt"

func solveNQueens(n int) [][]string {
	if n <= 0 {
		return [][]string{{}}
	}
	solution := make([][]int, 0)
	elem := make([]int, n)
	for i := range elem {
		elem[i] = i
	}
	bfs(elem, &solution, []int{}, map[int]bool{}, map[int]bool{})
	fmt.Println(solution)
	realSolution := make([][]string, len(solution))
	for i, s := range solution {
		realSolution[i] = transSolution(s)
	}
	return realSolution
}

func bfs(elem []int, solution *[][]int, current []int, valid1, valid2 map[int]bool) {
	if len(elem) == 0 {
		*solution = append(*solution, current)
		return
	}
	for i := 0; i < len(elem); i++ {
		point := []int{len(current), elem[i]}
		if valid1[point[0]+point[1]] == true {
			continue
		}
		if valid2[point[0]-point[1]] == true {
			continue
		}
		valid1[point[0]+point[1]] = true
		valid2[point[0]-point[1]] = true
		nextElem := make([]int, len(elem)-1)
		copy(nextElem, elem[:i])
		copy(nextElem[i:], elem[i+1:])
		next := make([]int, len(current)+1)
		copy(next, current)
		next[len(current)] = elem[i]
		bfs(nextElem, solution, next, valid1, valid2)
		valid1[point[0]+point[1]] = false
		valid2[point[0]-point[1]] = false
	}
}

func transSolution(pattern []int) []string {
	solution := make([]string, len(pattern))
	b := make([]byte, len(pattern))
	for i := range b {
		b[i] = '.'
	}
	for i, v := range pattern {
		b[v] = 'Q'
		solution[i] = string(b)
		b[v] = '.'
	}
	return solution
}
