package main

import "fmt"

type Trie struct {
	nodes map[byte]*Trie
	leaf  bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{nodes: make(map[byte]*Trie)}
}

/** Inserts a word into the trie. */
func (trie *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	letter := word[0]
	if trie.nodes[letter] == nil {
		child := Constructor()
		trie.nodes[letter] = &child
	}

	if len(word) > 1 {
		trie.nodes[letter].Insert(word[1:])
	} else {
		trie.nodes[letter].leaf = true
	}
}

/** Returns if the word is in the trie. */
func (trie *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	child, ok := trie.nodes[word[0]]
	if len(word) == 1 {
		return ok && child.leaf
	}

	return ok && child.Search(word[1:])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (trie *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}

	child, ok := trie.nodes[prefix[0]]
	if len(prefix) == 1 {
		return ok
	}

	return ok && child.StartsWith(prefix[1:])
}

func (trie Trie) String() string {
	return fmt.Sprintf("{leaf: %v, nodes: %v}\n", trie.leaf, trie.nodes)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
