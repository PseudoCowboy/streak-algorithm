package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))
}

func minSetSize(arr []int) int {
	m := make(map[int]int)
	for i := range arr {
		m[arr[i]]++
	}
	// interval := make([][]int, 0)
	freq := make([]int, 0)
	for _, v := range m {
		freq = append(freq, v)
	}
	// quickSort(interval)
	sort.Slice(freq, func(i, j int) bool {
		return freq[i] > freq[j]
	})
	fmt.Println(freq)
	ans := 0
	n := len(arr)
	temp := len(arr)
	for i := range freq {
		ans++
		temp -= freq[i]
		if temp <= n/2 {
			break
		}
	}
	return ans
}

func quickSort(interval [][]int) {
	if len(interval) < 2 {
		return
	}
	left, right := 0, len(interval)-1
	for i := 0; i < len(interval); i++ {
		if interval[i][1] > interval[right][1] {
			interval[i], interval[left] = interval[left], interval[i]
			left++
		}
	}
	interval[left], interval[right] = interval[right], interval[left]

	quickSort(interval[:left])
	quickSort(interval[left+1:])

}
