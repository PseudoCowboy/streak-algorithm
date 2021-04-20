package main

import "fmt"

func main() {
	var num uint32 = 1111110111
	var a uint32

	fmt.Println(a<<1 | num&1)
	fmt.Println(a << 1)

	fmt.Println(num >> 1)

	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)
}

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res<<1 | num&1
		num >>= 1
	}
	return res
}
