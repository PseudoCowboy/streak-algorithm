package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(findDifferentBinaryString([]string{"001", "000", "000"}))
}

func findDifferentBinaryString(nums []string) string {
	m := make(map[string]bool)
	for i := range nums {
		m[nums[i]] = true
	}
	distinct := make([]string, 0)
	for k := range m {
		distinct = append(distinct, k)
	}
	arr := make([]int, len(distinct))
	for i := range distinct {
		v, _ := strconv.ParseInt(distinct[i], 2, 32)
		arr[i] = int(v)
	}

	sort.Ints(arr)
	index := 0
	target := 0
	for i := 0; i <= arr[len(arr)-1]+1; i++ {
		if i < len(arr) && i == arr[index] {
			index++
		} else {
			target = i
			break
		}
	}
	var b string
	for n := target; n > 0; n /= 2 {
		b = strconv.Itoa(n%2) + b
	}
	j := len(distinct[0]) - len(b)
	for i := 0; i < j; i++ {
		b = "0" + b
	}
	return b
}
