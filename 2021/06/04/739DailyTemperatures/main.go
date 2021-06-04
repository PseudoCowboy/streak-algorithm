package main

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	for i := 0; i < len(temperatures)-1; i++ {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				result[i] = j - i
				break
			}
		}
	}

	return result
}

func dailyTemperatures1(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var toCheck []int
	for i, t := range temperatures {
		for len(toCheck) > 0 && temperatures[toCheck[len(toCheck)-1]] < t {
			idx := toCheck[len(toCheck)-1]
			res[idx] = i - idx
			toCheck = toCheck[:len(toCheck)-1]
		}
		toCheck = append(toCheck, i)
	}
	return res
}

func dailyTemperatures2(temperatures []int) []int {
	type Tuple struct {
		Temp  int
		Index int
	}

	res := make([]int, len(temperatures))
	var stack []Tuple

	for i := len(temperatures) - 1; i >= 0; i-- {
		res[i] = 0
		temp := temperatures[i]
		for (len(stack) > 0) && (temp >= stack[len(stack)-1].Temp) {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1].Index - i
		}

		stack = append(stack, Tuple{Temp: temp, Index: i})
	}
	return res
}
