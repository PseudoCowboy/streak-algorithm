package main

func isHappy(n int) bool {
	if n == 0 {
		return false
	}
	res := 0
	num := n
	record := map[int]int{}

	for {
		for num != 0 {
			res += (num % 10) * (num % 10)
			num /= 10
		}
		if _, ok := record[res]; !ok {
			if res == 1 {
				return true
			}
			num = res
			record[res] = res
			res = 0
			continue
		} else {
			return false
		}
	}
}
