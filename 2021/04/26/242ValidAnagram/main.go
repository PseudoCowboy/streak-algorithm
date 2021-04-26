package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	arr := make([]int, 26)
	sbyte := []byte(s)
	tbyte := []byte(t)

	for i := range sbyte {
		arr[sbyte[i]-'a']++
	}

	for i := range tbyte {
		arr[tbyte[i]-'a']--
	}

	for i := range arr {
		if arr[i] != 0 {
			return false
		}
	}

	return true
}
