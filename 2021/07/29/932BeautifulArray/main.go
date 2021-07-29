package main

import "fmt"

func main() {
	fmt.Println(beautifulArray(10))
}

func beautifulArray(n int) []int {
	ans := make([]int, n)
	for i := range ans {
		ans[i] = i + 1
	}
	ans = helper(ans)
	return ans
}

func helper(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left, right := make([]int, 0), make([]int, 0)
	for i := 0; i < len(arr); i += 2 {
		left = append(left, arr[i])
	}
	for i := 1; i < len(arr); i += 2 {
		right = append(right, arr[i])
	}
	left = helper(left)
	right = helper(right)

	return append(left, right...)

}
