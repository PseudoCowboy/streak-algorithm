package main

func arrayNesting(nums []int) int {
	m := make(map[int]int)
	max := 0
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		start := nums[i]
		next := nums[start]
		current := make([]int, 0)
		current = append(current, start)
		for next != start {
			current = append(current, next)
			next = nums[next]
		}
		for j := range current {
			m[current[j]] = len(current)
		}
		if len(current) > max {
			max = len(current)
		}
	}
	return max
}
