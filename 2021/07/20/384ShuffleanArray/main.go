package main

import (
	"math/rand"
	"time"
)

type Solution struct {
	original []int
}

func Constructor(nums []int) Solution {
	rand.Seed(time.Now().UnixNano())
	return Solution{
		original: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.original
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	shuffle := make([]int, len(this.original))
	copy(shuffle, this.original)
	for i := len(shuffle) - 1; i >= 0; i-- {
		random := rand.Intn(i + 1)
		shuffle[i], shuffle[random] = shuffle[random], shuffle[i]
	}
	return shuffle
}

/**
* Your Solution object will be instantiated and called as such:
* obj := Constructor(nums);
* param_1 := obj.Reset();
* param_2 := obj.Shuffle();
 */
