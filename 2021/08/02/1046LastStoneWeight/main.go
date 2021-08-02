package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}

func lastStoneWeight2(stones []int) int {
	pq := IntHeap(stones)
	heap.Init(&pq)
	for pq.Len() > 1 {
		x, y := heap.Pop(&pq).(int), heap.Pop(&pq).(int)
		fmt.Println(x, y)
		heap.Push(&pq, x-y)

	}
	return heap.Pop(&pq).(int)
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

func lastStoneWeight(stones []int) int {
	l := len(stones)
	h := make(intHeap, l)
	for i := range stones {
		h[i] = stones[i]
	}
	heap.Init(&h)
	for len(h) > 1 {
		x, y := heap.Pop(&h).(int), heap.Pop(&h).(int)
		if x-y > 0 {
			heap.Push(&h, x-y)
		}
	}
	if h.Len() > 0 {
		return h.Pop().(int)
	}
	return 0
}

func lastStoneWeight1(stones []int) int {
	sort.Ints(stones)
	for len(stones) > 1 {
		stone := stones[len(stones)-1] - stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		if stone > 0 {
			idx := sort.SearchInts(stones, stone)
			if idx >= len(stones) {
				stones = append(stones, stone)
			} else {
				stones = append(stones[:idx], append([]int{stone}, stones[idx:]...)...)
			}
		}
	}
	if len(stones) > 0 {
		return stones[0]
	}
	return 0
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	l := len(*h)
	x := (*h)[l-1]
	*h = (*h)[:l-1]
	return x
}
