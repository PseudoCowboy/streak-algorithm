package main

import "fmt"

func main() {
	a := longestCommonPrefix([]string{"ab", "a"})
	fmt.Println(a)
}

func longestCommonPrefix(strs []string) string {
	prefix := ""
	for i := 0; i < len(strs[0]); i++ {
		var temp byte
		for j := range strs {
			if j == 0 {
				continue
			}
			if i > len(strs[j])-1 {
				return prefix
			}
			temp = strs[0][i]
			fmt.Println(temp)
			if strs[j][i] != temp {
				return prefix
			}
		}
		prefix += string(temp)
	}
	return prefix
}

func longestCommonPrefix1(strs []string) string {
	index := 0
	result := ""
	s := ""
	flag := true
	if len(strs) == 0 {
		return result
	}
	for flag == true {
		for i := range strs {
			if index > len(strs[i])-1 {
				flag = false
				break
			}
			if i == 0 {
				s = string(strs[i][index])
			}
			if string(strs[i][index]) != s {
				flag = false
				break
			}
			if i == len(strs)-1 {
				result += s
				s = ""
			}
		}
		index++
	}
	return result
}
