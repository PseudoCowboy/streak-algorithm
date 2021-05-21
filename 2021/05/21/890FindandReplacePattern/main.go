package main

func findAndReplacePattern(words []string, pattern string) []string {
	result := make([]string, 0)
	for i := range words {
		wmap := make(map[byte]byte)
		pmap := make(map[byte]byte)
		j := 0
		for ; j < len(pattern); j++ {
			wv, wok := wmap[words[i][j]]
			pv, pok := pmap[pattern[j]]
			if wok && !pok {
				break
			}
			if !wok && pok {
				break
			}
			if wok && pok {
				if wv != pattern[j] || pv != words[i][j] {
					break
				}
				continue
			}
			wmap[words[i][j]] = pattern[j]
			pmap[pattern[j]] = words[i][j]
		}

		if j == len(pattern) {
			result = append(result, words[i])
		}
	}
	return result
}
