package main

import (
	"strconv"
)

func main() {
	// fmt.Println(interchangeableRectangles([][]int{
	// 	{4, 8}, {3, 6}, {10, 20}, {15, 30},
	// }))

	test()
}

func reversePrefix(word string, ch byte) string {
	sb := []byte(word)
	index := 0
	for i := range sb {
		if sb[i] == ch {
			index = i
			break
		}
	}
	reverse(sb, index)
	return string(sb)
}

func reverse(arr []byte, end int) {
	start := 0
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func interchangeableRectangles(rectangles [][]int) int64 {
	m := make(map[string]int)
	for i := range rectangles {
		g := gcd(rectangles[i][0], rectangles[i][1])
		f := rectangles[i][0] / g
		s := rectangles[i][1] / g
		m[strconv.Itoa(f)+","+strconv.Itoa(s)]++
	}
	ans := int64(0)
	for _, v := range m {
		if v > 1 {
			count := int64(v) - 1
			temp := (count + 1) * (count / 2)

			if count%2 != 0 {
				temp += count/2 + 1
			}
			ans += temp
		}
	}
	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		mod := a % b
		a = b
		b = mod
	}
	return a
}

func test() {
	arr := make([]int, 0)
	test1(arr)
}

func test1(arr []int) {
	arr = append(arr, 1)
}
