package main

func main() {
	slowestKey([]int{19, 22, 28, 29, 66, 81, 93, 97}, "fnfaaxha")
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	arr := make([]int, 26)
	arr[keysPressed[0]-'a'] = releaseTimes[0]
	ans := keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		arr[keysPressed[i]-'a'] = max(releaseTimes[i]-releaseTimes[i-1], arr[keysPressed[i]-'a'])
	}
	count := 0
	for i := 0; i < 26; i++ {
		if arr[i] >= count {
			count = arr[i]
			ans = 'a' + byte(i)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
