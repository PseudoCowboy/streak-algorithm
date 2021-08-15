package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(latestDayToCross(2, 2, [][]int{
		{1, 1}, {2, 1}, {1, 2}, {2, 2},
	}))
}

var dir = [][]int{
	{1, 1},
	{1, 0},
	{0, 1},
	{1, -1},
	{-1, 1},
	{-1, 0},
	{0, -1},
	{-1, -1},
}

func latestDayToCross(row int, col int, cells [][]int) int {
	matrix := make([][]int, row)
	for i := range matrix {
		matrix[i] = make([]int, col)
	}
	m := make(map[string]int)
	wid := make([][]int, 0)
	max := 0
	for index, cell := range cells {
		x := cell[0] - 1
		y := cell[1] - 1
		matrix[x][y] = 1
		l := check(x, y, matrix, m)
		fmt.Println(wid)
		if len(l) > 0 {
			miny := col
			maxy := -1
			first := l[0]
			used := make(map[int]bool)
			for _, v := range l {
				cminy := wid[v][0]
				cmaxy := wid[v][1]
				if cminy < miny {
					miny = cminy
				}
				if cmaxy > maxy {
					maxy = cmaxy
				}
				for key, value := range m {
					if value == v {
						m[key] = first
					}
				}
				used[v] = true
			}
			if y+1 > maxy {
				maxy = y + 1
			}
			if y+1 < miny {
				miny = y + 1
			}
			m[build(x, y)] = first
			if maxy-miny > max {
				max = maxy - miny
			}
			f := []int{miny, maxy}
			temp := make([][]int, 0)
			temp = append(temp, f)
			for i := range wid {
				if used[i] {
					continue
				}
				temp = append(temp, wid[i])
			}
			copy(wid, temp)
		} else {
			s := build(x, y)
			wid = append(wid, []int{y, y + 1})
			m[s] = len(wid) - 1
			if 1 > max {
				max = 1
			}
		}
		if max == col {
			return index + 1
		}
	}
	return len(cells)
}

func build(a, b int) string {
	return strconv.Itoa(a) + "," + strconv.Itoa(b)
}

func check(x, y int, matrix [][]int, m map[string]int) []int {
	temp := make(map[int]bool)
	ans := make([]int, 0)
	for _, move := range dir {
		nextx := x + move[0]
		nexty := y + move[1]
		if nextx >= 0 && nextx < len(matrix) && nexty >= 0 && nexty < len(matrix[0]) && matrix[nextx][nexty] == 1 {
			s := build(nextx, nexty)
			temp[m[s]] = true
		}
	}
	for k := range temp {
		ans = append(ans, k)
	}
	sort.Ints(ans)
	return ans
}
