package main

func minOperations(logs []string) int {
	stack := 0
	for i := range logs {
		if logs[i] == "./" {
			continue
		} else if logs[i] == "../" {
			if stack > 0 {
				stack--
			}
		} else {
			stack++
		}
	}
	return stack
}
