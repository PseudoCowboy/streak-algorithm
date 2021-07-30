package main

type Trie struct {
	isWord   bool
	children map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		isWord:   false,
		children: map[rune]*Trie{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	parent := this
	for _, v := range word {
		if child, ok := parent.children[v]; ok {
			parent = child
		} else {
			newChild := &Trie{isWord: false, children: map[rune]*Trie{}}
			parent.children[v] = newChild
			parent = newChild
		}
	}
	parent.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	parent := this
	for _, v := range word {
		if child, ok := parent.children[v]; ok {
			parent = child
		} else {
			return false
		}
	}
	return parent.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, v := range prefix {
		if child, ok := parent.children[v]; ok {
			parent = child
		} else {
			return false
		}
	}
	return true
}
