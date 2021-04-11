package main

import "fmt"

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	alength := len(a)
	blength := len(b)
	var length int
	if alength > blength {
		length = alength
	} else {
		length = blength
	}
	var carrying byte
	result := ""
	// 0 48  1 49  2 50
	for i := 0; i < length+1; i++ {
		var abit byte
		var bbit byte
		if i < alength {
			abit = a[alength-i-1]
		}
		if i < blength {
			bbit = b[blength-i-1]
		}
		switch abit + bbit + carrying {
		case 48:
			result = "0" + result
		case 49:
			result = "1" + result
		case 96:
			result = "0" + result
		case 97:
			result = "1" + result
			carrying = 0
		case 98:
			result = "0" + result
			carrying = 49
		case 145:
			result = "1" + result
			carrying = 0
		case 146:
			result = "0" + result
			carrying = 49
		case 147:
			result = "1" + result
			carrying = 49
		}
	}
	return result
}
