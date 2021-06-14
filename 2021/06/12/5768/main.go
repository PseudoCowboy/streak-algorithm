package main

func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for i := range chalk {
		sum += chalk[i]
	}
	k = k % sum
	i := 0
	for k >= 0 {
		k -= chalk[i]
		i++
	}
	return i - 1
}
