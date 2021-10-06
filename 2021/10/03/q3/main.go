package main

import "fmt"

func main() {
	// fmt.Println(stoneGameIX([]int{5, 1, 2, 4, 3}))
	// fmt.Println(stoneGameIX([]int{2}))
	// fmt.Println(stoneGameIX([]int{2, 1}))
	// fmt.Println(stoneGameIX([]int{2, 3}))
	// fmt.Println(stoneGameIX([]int{2, 3, 1}))

	fmt.Println(stoneGameIX([]int{20, 3, 20, 17, 2, 12, 15, 17, 4}))

}

func stoneGameIX(stones []int) bool {
	if len(stones) == 1 {
		return false
	}
	arr := make([]int, 0)
	for i := range stones {
		if stones[i]%3 != 0 {
			arr = append(arr, stones[i]%3)
		}
	}
	if len(arr) == 0 {
		return false
	}
	if len(arr) == 1 {
		return false
	}
	fmt.Println(arr)

	tc := 0
	for i := range arr {
		if arr[i] == 2 {
			tc++
		}
	}
	thridc := len(stones) - len(arr)
	oc := len(arr) - tc
	afirst := 1
	for tc > 1 && oc > 0 {
		tc -= 2
		oc--
		afirst = -1 * afirst
	}

	if thridc%2 == 1 {
		afirst = -1 * afirst

	}
	if oc == 0 && tc == 0 {
		return false
	}
	if oc == 0 {
		if tc < 3 {
			return false
		} else {
			if afirst == 1 {
				return false
			} else {
				return true
			}
		}
	} else {
		if tc == 1 {
			if afirst == 1 {
				return true
			} else {
				return false
			}
		}
		if oc == 1 {
			return false
		}
		if oc == 2 {
			return false
		}
		if afirst == 1 {
			return false
		}
		return true
	}
}
