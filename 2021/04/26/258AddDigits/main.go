package main

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	for num > 9 {
		cur := 0
		for num != 0 {
			cur += num % 10
			num /= 10
		}
		num = cur
	}
	return num
}
