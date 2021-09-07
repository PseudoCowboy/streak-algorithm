package main

import "fmt"

func main() {
	fmt.Println(getRow(1))
}

func getRow(rowIndex int) []int {
	pre := make([]int, rowIndex+1)
	ans := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				ans[j] = 1
				continue
			}
			ans[j] = pre[j-1] + pre[j]
		}
		copy(pre, ans)
	}
	return ans
}
