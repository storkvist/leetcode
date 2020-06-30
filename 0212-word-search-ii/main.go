package main

import "fmt"

var result map[string]bool

func findWords(board [][]byte, words []string) []string {
	result = make(map[string]bool)

	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}

	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(board, visited, "", i, j, trie)
		}
	}

	resultStr := make([]string, 0)
	for str := range result {
		resultStr = append(resultStr, str)
	}
	return resultStr
}

func dfs(board [][]byte, visited [][]bool, str string, i, j int, trie Trie) {
	m, n := len(board), len(board[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if visited[i][j] {
		return
	}
	str += string(board[i][j])
	if !trie.StartsWith(str) {
		return
	}
	if trie.Search(str) {
		result[str] = true
	}
	visited[i][j] = true
	dfs(board, visited, str, i-1, j, trie)
	dfs(board, visited, str, i+1, j, trie)
	dfs(board, visited, str, i, j-1, trie)
	dfs(board, visited, str, i, j+1, trie)
	visited[i][j] = false
}

type Trie struct {
	nodes map[byte]*Trie
	leaf  bool
}

func Constructor() Trie {
	return Trie{nodes: make(map[byte]*Trie)}
}

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
