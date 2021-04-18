package main

func twoSum(numbers []int, target int) []int {
	m := map[int]int{}
	result := []int{}
	for i := range numbers {
		if val, ok := m[target-numbers[i]]; ok {
			result = []int{val, i + 1}
		} else {
			m[numbers[i]] = i + 1
		}
	}
	return result
}

func twoSum1(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return []int{-1, -1}
}
