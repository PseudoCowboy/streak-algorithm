package main

func findComplement(num int) int {
	xx := ^0
	for xx&num > 0 {
		xx <<= 1
	}
	return ^xx ^ num
}
