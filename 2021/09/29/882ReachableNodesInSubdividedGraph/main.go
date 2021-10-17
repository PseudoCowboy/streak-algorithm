package main

import (
	"container/heap"
	"fmt"
)

func main() {
	reachableNodes([][]int{
		{1, 2, 4}, {1, 4, 5}, {1, 3, 1}, {2, 3, 4}, {3, 4, 5},
	}, 17, 5)
}

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	graph := make(map[int][][2]int)
	for _, v := range edges {
		s := v[0]
		e := v[1]
		w := v[2]
		graph[s] = append(graph[s], [2]int{e, w + 1})
		graph[e] = append(graph[e], [2]int{s, w + 1})
	}
	fmt.Print(graph)
	visited := make([]bool, n)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = 1e9
	}
	dist[0] = 0
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, [2]int{0, 0})
	for minHeap.Len() > 0 {
		item := heap.Pop(minHeap).([2]int)
		s := item[0]
		w := item[1]
		if visited[s] {
			continue
		}
		visited[s] = true
		for _, neighbor := range graph[s] {
			fmt.Println(neighbor)
			e := neighbor[0]
			if w+neighbor[1] < dist[e] {
				dist[e] = w + neighbor[1]
				heap.Push(minHeap, [2]int{e, w + neighbor[1]})
			}
		}
	}
	fmt.Println(dist)
	ans := 0
	for _, edge := range edges {
		left := max(0, min(maxMoves-dist[edge[1]], edge[2]))
		fmt.Println(max(0, min(maxMoves-dist[0], edge[2])))
		right := max(0, min(maxMoves-dist[0], edge[2]))
		fmt.Println(left + right)
		ans += min(left+right, edge[2])
	}
	for i := range dist {
		if dist[i] <= maxMoves {
			ans++
		}
	}
	fmt.Println(ans)

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type MinHeap [][2]int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	if h[i][1] == h[j][1] {
		return h[i][0] < h[j][0]
	} else {
		return h[i][1] < h[j][1]
	}
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
