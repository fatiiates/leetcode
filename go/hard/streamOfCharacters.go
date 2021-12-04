package main

import "fmt"

type StreamChecker struct {
	words     []string
	tr        *TrieNode
	trieLevel *TrieNode
	queries   *[]byte
}

type TrieNode struct {
	children    map[byte]*TrieNode
	isEndOfWord bool
}

func (t *TrieNode) insert(word string) {
	node := t

	for i := len(word) - 1; i >= 0; i-- {
		if _, ok := node.children[byte(word[i])]; ok {
			if !node.children[byte(word[i])].isEndOfWord && i == 0 {
				node.children[byte(word[i])].isEndOfWord = true
			}
			node = (*node).children[byte(word[i])]
		} else {
			isEndOfWord := false
			if i == 0 {
				isEndOfWord = true
			}
			node.children[byte(word[i])] = &TrieNode{
				children:    make(map[byte]*TrieNode),
				isEndOfWord: isEndOfWord,
			}
			node = (*node).children[byte(word[i])]
		}
	}
}

func (t *TrieNode) search(s *[]byte) bool {

	node := t
	for i := len(*s) - 1; i >= 0; i-- {
		if _, ok := node.children[(*s)[i]]; !ok {
			return false
		}
		node = node.children[(*s)[i]]
		if node.isEndOfWord {
			return true
		}
	}
	return false
}

func Constructor(words []string) StreamChecker {
	tr := TrieNode{
		children:    make(map[byte]*TrieNode),
		isEndOfWord: false,
	}

	// create trie
	for _, v := range words {
		tr.insert(v)
	}

	return StreamChecker{
		words:     words,
		tr:        &tr,
		trieLevel: &tr,
		queries:   &[]byte{},
	}
}

func (this *StreamChecker) Query(letter byte) bool {

	*this.queries = append(*this.queries, letter) // add new character to query stream
	return this.tr.search(this.queries)           // and search for any suffix in Trie

}

func maxLength(words []string) int {
	mx := 0
	for _, v := range words {
		lmx := len(v)
		if mx < lmx {
			mx = lmx
		}
	}
	return mx
}

func wrtieTrie(t *TrieNode, s string) {

	for k, v := range t.children {
		if v.isEndOfWord {
			fmt.Println(s + string(k))
		}
		if len(v.children) > 0 {
			wrtieTrie(v, s+string(k))
		}
	}
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
