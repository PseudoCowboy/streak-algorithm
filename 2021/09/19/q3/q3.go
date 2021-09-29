package main

import (
	"fmt"
	"strconv"
)

func main() {
	d := DetectSquares{
		x:  make(map[int][]int),
		y:  make(map[int][]int),
		xy: make(map[string]int),
	}

	s := []string{"add", "add", "add", "count", "add", "add", "add", "count", "add", "add", "add", "count", "add", "add", "add", "count"}
	t := [][]int{{5, 10}, {10, 5}, {10, 10}, {5, 5}, {3, 0}, {8, 0}, {8, 5}, {3, 5}, {9, 0}, {9, 8}, {1, 8}, {1, 0}, {0, 0}, {8, 0}, {8, 8}, {0, 8}}
	fmt.Println(len(s), len(t))
	for i, v := range s {
		if v == "add" {
			d.Add(t[i])
		} else {
			fmt.Println(d.Count(t[i]))
		}
	}
}

type DetectSquares struct {
	x  map[int][]int
	y  map[int][]int
	xy map[string]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		x:  make(map[int][]int),
		y:  make(map[int][]int),
		xy: make(map[string]int),
	}
}

func (this *DetectSquares) Add(point []int) {
	this.x[point[0]] = append(this.x[point[0]], point[1])
	this.y[point[1]] = append(this.y[point[1]], point[0])
	this.xy[strconv.Itoa(point[0])+","+strconv.Itoa(point[1])]++
}

func (this *DetectSquares) Count(point []int) int {
	count := 0
	for _, v := range this.x[point[0]] {
		if v == point[1] {
			continue
		}
		for _, v2 := range this.y[point[1]] {
			if v2 == point[0] {
				continue
			}
			if abs(point[0]-v2) != abs(point[1]-v) {
				continue
			}
			count += this.xy[strconv.Itoa(v2)+","+strconv.Itoa(v)]
		}
	}
	return count
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
