package main

import "fmt"

func main() {
	fmt.Println(customSortString("cba", "abcd"))
}

func customSortString(order string, str string) string {
	count := make([]int, 26)
	for i := range str {
		count[str[i]-'a']++
	}
	index := 0
	result := []byte(str)
	for i := range order {
		for count[order[i]-'a'] > 0 {
			result[index] = order[i]
			index++
			count[order[i]-'a']--
		}
	}
	for i := range count {
		for count[i] > 0 {
			result[index] = byte(i) + 'a'
			count[i]--
			index++
		}
	}
	return string(result)
}
