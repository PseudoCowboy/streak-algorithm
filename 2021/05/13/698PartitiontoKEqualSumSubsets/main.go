package main

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	average := sum / k
	bucket := make([]int, k)
	return dfs(nums, &bucket, average, 0)

}

func dfs(nums []int, bucket *[]int, average int, index int) bool {
	if index == len(*bucket) {
		for i := range *bucket {
			if (*bucket)[i] != average {
				return false
			}
		}

		return true
	}

	for i := 0; i < len(nums); i++ {
		if (*bucket)[i]+nums[index] > average {
			continue
		}
		(*bucket)[i] += nums[index]
		if dfs(nums, bucket, average, index+1) {
			return true
		}
		(*bucket)[i] -= nums[index]
	}

	return false
}
