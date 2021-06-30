package main

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[byte]int)
	for i := range magazine {
		m[magazine[i]]++
	}
	for i := range ransomNote {
		m[ransomNote[i]]--
		if m[ransomNote[i]] < 0 {
			return false
		}
	}

	return true
}
