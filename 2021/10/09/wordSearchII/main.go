package main

func main() {
	findWords([][]byte{
		{'a', 'b', 'c', 'e'},
		{'x', 'x', 'c', 'd'},
		{'x', 'x', 'b', 'a'},
	}, []string{"abc", "abcd"})
}

func findWords(board [][]byte, words []string) []string {
	t := &trie{
		isWord:   false,
		children: make(map[byte]*trie),
	}
	for i := range words {
		t.insert(words[i])
	}
	ans := make([]string, 0)
	used := make([][]bool, len(board))
	for i := range board {
		used[i] = make([]bool, len(board[0]))
	}
	for i := range board {
		for j := range board[i] {
			for k := range used {
				for l := range used[k] {
					used[k][l] = false
				}
			}
			dfs(board, &ans, t, i, j, used, "")
		}
	}
	return ans
}

func dfs(board [][]byte, ans *[]string, t *trie, row, col int, used [][]bool, current string) {
	used[row][col] = true
	current += string(board[row][col])
	if !t.isPrefix(current) {
		return
	}
	if t.search(current) {
		*ans = append(*ans, current)
		t.remove(current)
	}

	for _, move := range moves {
		nr := row + move[0]
		nc := col + move[1]
		if nr >= 0 && nr < len(board) && nc >= 0 && nc < len(board[0]) && !used[nr][nc] {
			dfs(board, ans, t, nr, nc, used, current)
			used[nr][nc] = false
		}
	}
}

var moves = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

type trie struct {
	isWord   bool
	children map[byte]*trie
}

func (t *trie) insert(word string) {
	parent := t
	for i := range word {
		if child, ok := parent.children[word[i]]; ok {
			parent = child
		} else {
			newChild := &trie{
				isWord:   false,
				children: make(map[byte]*trie),
			}
			parent.children[word[i]] = newChild
			parent = newChild
		}
	}
	parent.isWord = true
}

func (t *trie) isPrefix(word string) bool {
	parent := t
	for i := range word {
		if child, ok := parent.children[word[i]]; ok {
			parent = child
		} else {
			return false
		}
	}
	return true
}

func (t *trie) remove(word string) {
	parent := t
	for i := range word {
		child := parent.children[word[i]]
		parent = child
	}
	parent.isWord = false
}

func (t *trie) search(word string) bool {
	parent := t
	for i := range word {
		if child, ok := parent.children[word[i]]; ok {
			parent = child
		} else {
			return false
		}
	}
	return parent.isWord
}
