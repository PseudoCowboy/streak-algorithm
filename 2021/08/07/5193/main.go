package main

import "fmt"

func main() {
	fmt.Println(makeFancyString("aaabaaaa"))
}

func makeFancyString(s string) string {
	b := []byte(s)
	target := make([]byte, 0)
	target = append(target, b[0])
	temp := 0
	fast := 1
	for fast < len(b) {
		if b[fast] == b[fast-1] {
			temp++
		} else {
			temp = 0
		}
		if temp >= 2 {
			fast++
		} else {
			target = append(target, b[fast])
			fast++
		}
	}

	return string(target)
}
