package main

func main() {
	canPartitionKSubsets([]int{1, 1, 1, 1, 2, 2, 2, 2, 2}, 2)
}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	average := sum / k
	if sum%k != 0 {
		return false
	}

	bucket := make([]int, k)
	return dfs(nums, &bucket, average, 0)

}

func dfs(nums []int, bucket *[]int, average int, index int) bool {
	if index == len(nums) {
		for i := range *bucket {
			if (*bucket)[i] != average {
				return false
			}
		}

		return true
	}

	for i := 0; i < len(*bucket); i++ {
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

func canPartitionKSubsets1(nums []int, k int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	average := sum / k
	if sum%k != 0 {
		return false
	}

	used := make([]bool, len(nums))

	return dfs1(nums, used, k, average, 0, 0)
}

func dfs1(nums []int, used []bool, k, average, start, bucket int) bool {
	if k == 0 {
		return true
	}
	if bucket == average {
		return dfs1(nums, used, k-1, average, 0, 0)
	}

	for i := start; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if bucket+nums[i] > average {
			continue
		}
		bucket += nums[i]
		used[i] = true
		if dfs1(nums, used, k, average, i+1, bucket) {
			return true
		}
		bucket -= nums[i]
		used[i] = false
	}

	return false
}
