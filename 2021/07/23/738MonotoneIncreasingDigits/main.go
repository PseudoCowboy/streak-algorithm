package main

func main() {
	monotoneIncreasingDigits(10)
}

func monotoneIncreasingDigits(n int) int {
	arr := make([]int, 0)
	for n != 0 {
		arr = append(arr, n%10)
		n /= 10
	}
	reverse(arr)
	flag := len(arr)
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i+1] < arr[i] {
			flag = i + 1
			arr[i+1] = 9
			arr[i]--
		}
	}
	ans := 0
	for i := flag; i < len(arr); i++ {
		arr[i] = 9
	}
	for i := range arr {
		ans *= 10
		ans += arr[i]
	}
	return ans
}

func reverse(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
