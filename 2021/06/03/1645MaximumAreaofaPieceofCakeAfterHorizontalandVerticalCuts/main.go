package main

import (
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	horizontalCuts = append(horizontalCuts, 0, h)
	verticalCuts = append(verticalCuts, 0, w)

	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	widthMax, hightMax := 0, 0
	for i := 1; i < len(horizontalCuts); i++ {
		if horizontalCuts[i]-horizontalCuts[i-1] > hightMax {
			hightMax = horizontalCuts[i] - horizontalCuts[i-1]
		}
	}
	for i := 1; i < len(verticalCuts); i++ {
		if verticalCuts[i]-verticalCuts[i-1] > widthMax {
			widthMax = verticalCuts[i] - verticalCuts[i-1]
		}
	}

	return (widthMax * hightMax) % 1000000007
}
