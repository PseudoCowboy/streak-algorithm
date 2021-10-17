package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	ans := 0
	for i := range seats {
		ans += abs(seats[i] - students[i])
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func winnerOfGame(colors string) bool {
	a := 0
	b := 0
	temp := 0
	for i := range colors {
		if colors[i] == 'A' {
			temp += 1
			if temp > 2 {
				a += temp - 2
			}
		} else {
			temp = 0
		}
	}
	temp = 0
	for i := range colors {
		if colors[i] == 'B' {
			temp += 1
			if temp > 2 {
				b += temp - 2
			}
		} else {
			temp = 0
		}
	}

	if a-b > 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(networkBecomesIdle([][]int{
		{0, 1}, {1, 2},
	}, []int{0, 2, 1}))
	// [0,1],[0,2],[1,2]
	fmt.Println(networkBecomesIdle([][]int{
		{0, 1}, {0, 2}, {1, 2},
	}, []int{0, 10, 10}))
	//[3,8],[4,13],[0,7],[0,4],[1,8],[14,1],[7,2],[13,10],[9,11],[12,14],[0,6],[2,12],[11,5],[6,9],[10,3]
	fmt.Println(networkBecomesIdle([][]int{
		{3, 8}, {4, 13}, {0, 7}, {0, 4}, {1, 8}, {14, 1}, {7, 2}, {13, 10}, {9, 11}, {12, 14}, {0, 6}, {2, 12}, {11, 5}, {6, 9}, {10, 3},
	}, []int{0, 3, 2, 1, 5, 1, 5, 5, 3, 1, 2, 2, 2, 2, 1}))
}

func networkBecomesIdle(edges [][]int, patience []int) int {
	m := make(map[int][]int)
	for i := range edges {
		s := edges[i][0]
		e := edges[i][1]
		m[s] = append(m[s], e)
		m[e] = append(m[e], s)
	}
	dist := make([]int, len(patience))
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[0] = 0
	queue := make([]int, 0)
	used := make(map[string]bool)
	for i := range m[0] {
		dist[m[0][i]] = 1
		queue = append(queue, m[0][i])
		used[strconv.Itoa(0)+","+strconv.Itoa(m[0][i])] = true
	}
	level := 1
	for len(queue) != 0 {
		size := len(queue)
		level++
		for i := range queue {
			for _, v := range m[queue[i]] {
				if level < dist[v] {
					dist[v] = level
				}
				if used[strconv.Itoa(queue[i])+","+strconv.Itoa(v)] == false {
					queue = append(queue, v)
				}
				used[strconv.Itoa(queue[i])+","+strconv.Itoa(v)] = true
			}
		}
		queue = queue[size:]
	}
	fmt.Println(dist)
	fmt.Println(patience)
	far := 0
	extra := math.MaxInt32
	for i := range dist {
		if dist[i] > far {
			far = dist[i]
		}
	}
	in := make([]int, 0)
	for i := range dist {
		if dist[i] == far {
			in = append(in, i)
		}
	}
	fmt.Println(extra)
	for j := range in {
		i := in[j]
		if patience[i] < 2*far {
			remain := 2 * far % patience[i]
			fmt.Println(remain, extra)
			if remain == 0 {
				fmt.Println(((2 * far / patience[i]) - 1) * patience[i])
				if 2*far-((2*far/patience[i])-1)*patience[i] < extra {
					extra = 2*far - ((2*far/patience[i])-1)*patience[i]
				}
				fmt.Println(extra)
			} else {
				if 2*far-(2*far/patience[i])*patience[i] < extra {
					extra = 2*far - (2*far/patience[i])*patience[i]
				}
			}
		}
	}

	if extra == math.MaxInt32 {
		return 2*far + 1
	}
	return 2*far + (2*far - extra) + 1
}
