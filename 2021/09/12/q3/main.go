package main

import "fmt"

func main() {
	fmt.Println(maxProduct("leetcodecom"))
}

func maxProduct(s string) int {
	sb := []byte(s)
	ans := 0
	fmt.Println(find(sb, 4, 5))
	// m := make(map[string]int)
	// for i := 1; i < len(sb); i++ {
	// 	s1 := find(sb, i, i)
	// 	if m[s1] == 0 {
	// 		s1other := filter(sb, s1)
	// 		s1b := []byte(s1other)
	// 		for j := 1; j < len(s1b); j++ {
	// 			o := find(s1b, j, j)
	// 			if len(s1)*len(o) > ans {
	// 				ans = len(s1) * len(o)
	// 			}
	// 			b := find(s1b, j, j+1)
	// 			if len(s1)*len(b) > ans {
	// 				ans = len(s1) * len(b)
	// 			}
	// 		}
	// 		m[s1]++
	// 	}
	// 	s2 := find(sb, i, i+1)
	// 	if m[s2] == 0 {
	// 		s2other := filter(sb, s2)
	// 		s2b := []byte(s2other)
	// 		for j := 1; j < len(s2b); j++ {
	// 			o := find(s2b, j, j)
	// 			if len(s2)*len(o) > ans {
	// 				ans = len(s2) * len(o)
	// 			}
	// 			b := find(s2b, j, j+1)
	// 			if len(s2)*len(b) > ans {
	// 				ans = len(s2) * len(b)
	// 			}
	// 		}
	// 		m[s2]++
	// 	}
	// }

	return ans
}

func filter(s []byte, str string) string {
	ans := make([]byte, 0)
	sstart := 0
	strstart := 0
	for sstart < len(s) && strstart < len(str) {
		if s[sstart] == str[strstart] {
			strstart++
			sstart++
			continue
		}
		ans = append(ans, s[sstart])
		sstart++
	}
	ans = append(ans, s[sstart:]...)
	return string(ans)
}

func find(s []byte, start, end int) string {
	ans := ""
	if start == end {
		ans = string(s[start])
		start--
		end++
	}
	m := make(map[byte]int)
	for start >= 0 {
		m[s[start]]++
		start--
	}
	for i := end; i < len(s); i++ {
		if m[s[i]] > 0 {
			ans = string(s[i]) + ans + string(s[i])
			m[s[i]]--
		}
	}
	return ans
}
