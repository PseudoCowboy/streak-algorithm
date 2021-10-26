package main

import "fmt"

func main() {
	m := make(map[int]bool)
	m[1] = true
	v, ok := m[2]
	fmt.Println(v, ok)
	fmt.Println(2 >> 1)
}
