package main

func lemonadeChange(bills []int) bool {
	five := 0
	ten := 0
	for i := range bills {
		if bills[i] == 5 {
			five++
		} else if bills[i] == 10 {
			five--
			ten++
		} else {
			if ten > 0 {
				ten--
				five--
			} else {
				five -= 3
			}
		}
		if five < 0 {
			return false
		}
	}
	return true
}
