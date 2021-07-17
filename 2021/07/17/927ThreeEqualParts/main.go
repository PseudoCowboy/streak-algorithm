package main

func threeEqualParts(arr []int) []int {
	n, ones, i, count := len(arr), 0, 0, 0
	for _, a := range arr {
		ones += a
	}
	if ones == 0 {
		return []int{0, n - 1}
	}
	if ones%3 != 0 {
		return []int{-1, -1}
	}
	k := ones / 3
	for i < n {
		if arr[i] == 1 {
			break
		}
		i++
	}
	start, j := i, i
	for j < n {
		count += arr[j]
		if count == k+1 {
			break
		}
		j++
	}
	mid := j
	j, count = 0, 0
	for j < n {
		count += arr[j]
		if count == 2*k+1 {
			break
		}
		j++
	}
	end := j
	for end < n && arr[start] == arr[mid] && arr[mid] == arr[end] {
		start++
		mid++
		end++
	}
	if end == n {
		return []int{start - 1, mid}
	}
	return []int{-1, -1}
}
