package main

import "fmt"

func main() {
	fmt.Println(minimumPerimeter(1000000000))
}

func minimumPerimeter(neededApples int64) int64 {
	n := neededApples/4 + 1
	count := 0
	s := int64(0)
	for s < n {
		count++
		for i := 1; i < count; i++ {
			s += int64(i)
		}
		s += int64(count)
	}
	return int64(count * 8)
}
