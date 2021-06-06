package main

import (
	"container/heap"
	"math"
)

func networkDelayTime(times [][]int, N int, K int) int {
	ways := make([][][2]int, N+1)
	record := make([]int, N+1)
	curMin := make([]int, N+1)
	for i := range record {
		record[i] = -1
	}
	for i := range curMin {
		curMin[i] = math.MaxInt32
	}
	for _, t := range times {
		// ways[i][0]: delay
		// ways[i][0]: node
		ways[t[0]] = append(ways[t[0]], [2]int{t[2], t[1]})
	}
	light := Heap{}
	record[K] = 0
	curMin[K] = 0
	for _, w := range ways[K] {
		heap.Push(&light, w)
		curMin[w[1]] = w[0]
	}
	for len(light) > 0 {
		node := heap.Pop(&light).([2]int)
		// already reached, continue
		if record[node[1]] != -1 {
			continue
		}
		curDelay := node[0]
		record[node[1]] = curDelay
		for _, w := range ways[node[1]] {
			// already reached, continue
			if record[w[1]] != -1 {
				continue
			}
			if curMin[w[1]] < w[0]+curDelay {
				continue
			}
			curMin[w[1]] = w[0] + curDelay
			heap.Push(&light, [2]int{w[0] + curDelay, w[1]})
		}
	}
	maxDelay := 0
	for i := 1; i <= N; i++ {
		if record[i] == -1 {
			return -1
		}
		if record[i] > maxDelay {
			maxDelay = record[i]
		}
	}
	return maxDelay
}

type Heap [][2]int

func (p Heap) Len() int            { return len(p) }
func (p Heap) Less(i, j int) bool  { return p[i][0] < p[j][0] }
func (p Heap) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *Heap) Push(i interface{}) { *p = append(*p, i.([2]int)) }
func (p *Heap) Pop() interface{}   { v := (*p)[len(*p)-1]; *p = (*p)[:len(*p)-1]; return v }
