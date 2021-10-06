package main

import "fmt"

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i := 0; i < len(original); i++ {
		row := i / n
		col := i % n
		ans[row][col] = original[i]
	}
	return ans
}

func numOfPairs(nums []string, target string) int {
	ans := 0
	for i := range nums {
		for j := range nums {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maxConsecutiveAnswers("FFTFTTTFFF", 1))
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	tc := 0
	n := len(answerKey)
	for i := range answerKey {
		if answerKey[i] == 'T' {
			tc++
		}
	}
	if tc <= k || n-tc <= k {
		return n
	}
	a := check(answerKey, k, 'T', 'F')
	b := check(answerKey, k, 'F', 'T')
	if a > b {
		return a
	}
	return b
	// return a
}

// "FFFTTFTTFT"
// 3

//FFTFTTTFFF

func check(answerKey string, k int, t, f byte) int {
	max := 0
	left, right := 0, 0
	for right < len(answerKey) {
		if answerKey[right] == t {
			if right == len(answerKey)-1 {
				if right-left+1 > max {
					max = right - left + 1
				}
			}
		} else if answerKey[right] == f && k > 0 {
			k--
		} else if answerKey[right] == f && k == 0 {
			if right-left > max {
				fmt.Println(right, left)
				max = right - left
			}
			for answerKey[left] != f && left < right {
				left++
			}
			left++
		}
		right++
	}
	return max
}
