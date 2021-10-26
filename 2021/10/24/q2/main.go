package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	strs := []string{"1", "3", "2"}
	sort.Strings(strs)
	fmt.Println(strs)
}

func nextBeautifulNumber(n int) int {
	s := strconv.Itoa(n)
	count := 0
	arr := make([]string, 0)
	for _, v := range s {
		vs := string(v)
		arr = append(arr, vs)
		value, _ := strconv.Atoi(vs)
		count += value
	}
	if count < len(s) {
		sort.Strings(arr)

	} else if count == len(s) {

	} else {

	}
	return 0
}
