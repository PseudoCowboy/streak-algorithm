package main

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; i++ {
		rest := 0
		for j := i; j < n+i; j++ {
			rest += gas[j%n]
			rest -= cost[j%n]
			if rest < 0 {
				break
			}
		}
		if rest >= 0 {
			return i
		}
	}
	return -1
}
func canCompleteCircuit1(gas []int, cost []int) int {
	totalGas := 0
	cur := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		cur += gas[i] - cost[i]
		totalGas += gas[i] - cost[i]
		if cur < 0 {
			start = i + 1
			cur = 0
		}
	}
	if totalGas < 0 {
		return -1
	}
	return start
}
