package main

import "fmt"

func main() {

	a := byte(0)
	b := byte(1)
	fmt.Println(a + b)
}

func addStrings(num1 string, num2 string) string {
	m := len(num1) - 1
	n := len(num2) - 1
	ans := make([]byte, 0)
	digit := byte(0)
	addOne := byte(0)
	for m >= 0 && n >= 0 {
		digit = num1[m] - '0' + num2[n] - '0' + addOne
		addOne = 0
		if digit >= 10 {
			addOne = 1
			digit -= 10
		}
		ans = append(ans, '0'+digit)
		m--
		n--
	}
	for m >= 0 {
		digit = num1[m] - '0' + addOne
		addOne = 0
		if digit >= 10 {
			addOne = 1
			digit -= 10
		}
		ans = append(ans, '0'+digit)
		m--
	}
	for n >= 0 {
		digit = num2[n] - '0' + addOne
		addOne = 0
		if digit >= 10 {
			addOne = 1
			digit -= 10
		}
		ans = append(ans, '0'+digit)
		n--
	}
	if addOne > 0 {
		ans = append(ans, '1')
	}
	reverse(ans)
	return string(ans)
}

func reverse(arr []byte) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
