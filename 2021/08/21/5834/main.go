package main

import "fmt"

func main() {
	fmt.Println(minTimeToType("abc"))

}

func minTimeToType(word string) int {
	ans := len(word) + min(int(word[0])-int('a'), int(int('z')-int(word[0]))+1)
	for i := 1; i < len(word); i++ {
		ans += min(abs(int(word[i])-int(word[i-1])), min(int(word[i])-int('a')+int('z')-int(word[i-1])+1, int(word[i-1])-int('a')+int('z')-int(word[i])+1))
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
