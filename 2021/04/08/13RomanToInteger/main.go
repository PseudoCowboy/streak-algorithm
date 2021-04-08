package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	skip := false
	result := 0
	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	for i := range s {
		value := string(s[i])
		if skip == true {
			skip = false
			continue
		}
		if i+1 < len(s) && m[value] < m[string(s[i+1])] {
			skip = true
			result += m[string(s[i+1])] - m[value]
		} else {
			result += m[value]
		}
	}
	return result
}

func romanToInt1(s string) int {
	m := map[string]int{"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, "XL": 40, "L": 50, "XC": 90, "C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000}
	result := 0

	skip := false
	for i := range s {
		if skip == true {
			skip = false
			continue
		}

		if i+1 < len(s) && m[string(s[i])+string(s[i+1])] != 0 {
			result += m[string(s[i])+string(s[i+1])]
			skip = true
		} else {
			result += m[string(s[i])]
		}
	}

	return result
}
