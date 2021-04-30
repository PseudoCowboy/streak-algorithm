package main

func main() {
	firstUniqChar("loveleetcode")
}

func firstUniqChar(s string) int {
	m := make(map[byte]int)
	for i := range s {
		_, ok := m[s[i]]
		if ok {
			m[s[i]] = -1
		} else {
			m[s[i]] = i
		}
	}
	result := -1
	flag := true
	for _, v := range m {
		if v != -1 && flag {
			result = v
			flag = false
		}
		if v != -1 && v < result {
			result = v
		}
	}
	return result
}

// 解法一
func firstUniqChar1(s string) int {
	result := make([]int, 26)
	for i := 0; i < len(s); i++ {
		result[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if result[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

// 解法二
// 执行用时: 8 ms
// 内存消耗: 5.2 MB
func firstUniqChar2(s string) int {
	charMap := make([][2]int, 26)
	for i := 0; i < 26; i++ {
		charMap[i][0] = -1
		charMap[i][1] = -1
	}
	for i := 0; i < len(s); i++ {
		if charMap[s[i]-'a'][0] == -1 {
			charMap[s[i]-'a'][0] = i
		} else { //已经出现过
			charMap[s[i]-'a'][1] = i
		}
	}
	res := len(s)
	for i := 0; i < 26; i++ {
		//只出现了一次
		if charMap[i][0] >= 0 && charMap[i][1] == -1 {
			if charMap[i][0] < res {
				res = charMap[i][0]
			}
		}
	}
	if res == len(s) {
		return -1
	}
	return res
}
