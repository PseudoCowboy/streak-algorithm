package main

import (
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	num1Arr := strings.Split(num1, "+")
	num2Arr := strings.Split(num2, "+")
	num1real, _ := strconv.Atoi(num1Arr[0])
	num2real, _ := strconv.Atoi(num2Arr[0])
	num1ima, _ := strconv.Atoi(num1Arr[1][:len(num1Arr[1])-1])
	num2ima, _ := strconv.Atoi(num2Arr[1][:len(num2Arr[1])-1])
	resultReal := num1real*num2real - num1ima*num2ima
	resultIma := num1real*num2ima + num2real*num1ima
	return strconv.Itoa(resultReal) + "+" + strconv.Itoa(resultIma) + "i"
}
