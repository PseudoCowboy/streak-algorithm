package main

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	used := make(map[string]int)
	for i := range wordList {
		used[wordList[i]] = 1
	}

	stack := make([]string, 0)
	stack = append(stack, beginWord)
	ans := make([][]string, 0)
	for len(stack) > 0 {

	}

	return ans

}
