package main

import (
	"container/heap"
	"fmt"
)

func main() {
	min := &MinHeap{}
	heap.Init(min)
	heap.Push(min, [2]int{2, 2})
	heap.Push(min, [2]int{1, 11})
	info := heap.Pop(min).([2]int)
	fmt.Println(info)

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
