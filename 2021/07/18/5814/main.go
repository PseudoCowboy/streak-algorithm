package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	b[0] = 2
	fmt.Println(a, b)

}

func addRungs(rungs []int, dist int) int {
	pre := 0
	ans := 0
	for i := range rungs {
		if rungs[i]-pre > dist {
			ans += (rungs[i] - pre) / dist
			if (rungs[i]-pre)%dist == 0 {
				ans--
			}
		}
		pre = rungs[i]
	}
	return ans
}
