package main

import (
	"container/heap"
	"math"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *MinHeap) Push(x interface{}) {

	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h *MaxHeap) Push(x interface{}) {

	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

type MedianFinder struct {
	minHeap *MinHeap
	maxHeap *MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {

	minH := &MinHeap{}
	maxH := &MaxHeap{}
	heap.Init(minH)
	heap.Init(maxH)

	md := MedianFinder{
		minHeap: minH,
		maxHeap: maxH,
	}

	return md
}

func (this *MedianFinder) balance() {

	for this.minHeap.Len()-this.maxHeap.Len() > 1 {
		heap.Push(this.maxHeap, heap.Pop(this.minHeap))
	}

	for this.maxHeap.Len()-this.minHeap.Len() > 1 {
		heap.Push(this.minHeap, heap.Pop(this.maxHeap))
	}
}

func (this *MedianFinder) AddNum(num int) {

	if float64(num) > this.FindMedian() {
		heap.Push(this.minHeap, num)
	} else {
		heap.Push(this.maxHeap, num)
	}

	this.balance()
}

func (this *MedianFinder) FindMedian() float64 {

	if this.minHeap.Len() == 0 && this.minHeap.Len() == 0 {
		return float64(math.MinInt64)

	} else if this.minHeap.Len() == this.maxHeap.Len() {
		return (float64((*this.minHeap)[0]) + float64((*this.maxHeap)[0])) / 2.00

	} else if this.minHeap.Len() > this.maxHeap.Len() {
		return float64((*this.minHeap)[0])

	} else {
		return float64((*this.maxHeap)[0])

	}
}
