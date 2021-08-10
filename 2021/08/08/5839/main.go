package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minStoneSum([]int{4, 3, 6, 7}, 3))
}

func minStoneSum(piles []int, k int) int {
	sum := 0
	for i := range piles {
		sum += piles[i]
	}
	remove := 0
	pq := IntHeap(piles)
	heap.Init(&pq)
	for k > 0 {
		k--
		x := heap.Pop(&pq).(int)
		remove += x / 2
		heap.Push(&pq, x-x/2)
	}
	return sum - remove
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
