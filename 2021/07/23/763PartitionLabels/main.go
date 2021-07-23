package main

func partitionLabels(s string) []int {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]] = i
	}
	pre := 0
	max := 0
	ans := make([]int, 0)
	for i := range s {
		if m[s[i]] > max {
			max = m[s[i]]
		}
		if i == max {
			ans = append(ans, max-pre+1)
			pre = max + 1
		}
	}
	return ans
}
