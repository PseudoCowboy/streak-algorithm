package main

func reverseWords(s string) string {
	sb := []byte(s)

	count := 0

	for i := 0; i < len(sb); i++ {
		if sb[i] == 32 {
			count++
		} else {
			break
		}
	}
	if count > 0 {
		sb = sb[count:]
	}
	count = 0

	reverse(sb)

	for i := 0; i < len(sb); i++ {
		if sb[i] == 32 {
			count++
		} else {
			break
		}
	}
	if count > 0 {
		sb = sb[count:]
	}
	count = 0

	fast, slow := 0, 0
	flag := 0
	for fast < len(sb) {
		if sb[slow] != 32 {
			fast++
			slow++
		} else {
			if sb[fast] == 32 {
				fast++
				flag = 0
			} else {
				if flag == 0 {
					slow++
					flag = 1
				}
				if fast-slow > 1 {
					sb[fast], sb[slow] = sb[slow], sb[fast]
					count++
					slow++
					fast++
				}
			}
		}
	}
	sb = sb[:len(sb)-count]
	sb = append(sb, byte(32))

	count = 0
	for i := 0; i < len(sb); i++ {
		if sb[i] != 32 {
			count++
		} else {
			reverse(sb[i-count : i])
			count = 0
		}
	}
	return string(sb[:len(sb)-1])

}

func reverse(sb []byte) {
	left, right := 0, len(sb)-1
	for left < right {
		sb[left], sb[right] = sb[right], sb[left]
		left++
		right--
	}
}
