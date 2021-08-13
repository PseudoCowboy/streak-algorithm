package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	ans := make([][]string, 0)
	m := make(map[string][]string)
	for i := range strs {
		temp := []byte(strs[i])
		quick(temp)
		m[string(temp)] = append(m[string(temp)], strs[i])
	}
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	arr := []byte{3, 2, 1, 6, 7, 8}
	quick(arr)
	fmt.Println(arr)
}

func quick(arr []byte) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	quick(arr[:left])
	quick(arr[left+1:])
}
