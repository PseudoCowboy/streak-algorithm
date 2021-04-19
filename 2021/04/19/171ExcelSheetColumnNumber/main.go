package main

func titleToNumber(columnTitle string) int {
	val, res := 0, 0
	for i := 0; i < len(columnTitle); i++ {
		val = int(columnTitle[i] - 'A' + 1)
		res = res*26 + val
	}
	return res
}
