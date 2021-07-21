package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMinArrowShots([][]int{
		{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8},
	}))

}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	cur := points[0][1]
	ans := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] <= cur {
			continue
		}
		ans++
		cur = points[i][1]
	}
	return ans
}
