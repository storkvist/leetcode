package main

import (
	"sort"
)

type BSTNode struct {
	data  byte
	freq  int
	left  *BSTNode
	right *BSTNode
}

type dataFreq struct {
	data byte
	freq int
}

type byFreq []dataFreq

func (a byFreq) Len() int           { return len(a) }
func (a byFreq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFreq) Less(i, j int) bool { return a[i].freq < a[j].freq }

func newNode(data byte) *BSTNode {
	return &BSTNode{
		data:  data,
		freq:  1,
		left:  nil,
		right: nil,
	}
}

func insert(root *BSTNode, data byte) *BSTNode {
	if root == nil {
		return newNode(data)
	}
	if root.data == data {
		root.freq++
	} else if root.data < data {
		root.left = insert(root.left, data)
	} else {
		root.right = insert(root.right, data)

	}
	return root
}

func store(root *BSTNode, count []dataFreq, index *int) {
	if root == nil {
		return
	}
	store(root.left, count, index)
	count[*index].freq = root.freq
	count[*index].data = root.data
	*index += 1
	store(root.right, count, index)
}

func frequencySort(s string) string {
	var root *BSTNode
	for _, b := range []byte(s) {
		root = insert(root, b)
	}

	count := make([]dataFreq, len(s))
	index := 0
	store(root, count, &index)

	sort.Sort(byFreq(count))

	result := ""
	length := len(count)
	for i := length - 1; i >= length-index; i-- {
		for count[i].freq > 0 {
			result += string(count[i].data)
			count[i].freq--
		}
	}

	return result
}

func frequencySort1(s string) string {
	dict := make(map[rune]int)
	for _, c := range s {
		dict[c]++
	}

	letters := make([]string, len(s)+1)
	for r, count := range dict {
		seq := ""
		for count > 0 {
			seq += string(r)
			count--
		}
		letters[count] += seq
	}

	result := ""
	for i := len(letters) - 1; i >= 0; i-- {
		result += letters[i]
	}

	return result
}

func frequencySort2(s string) string {
	dict := make(map[rune]int)
	for _, c := range s {
		dict[c]++
	}

	reversed := make(map[int]string)
	var frequencies []int
	for r, freq := range dict {
		reversed[freq] += string(r)
		frequencies = append(frequencies, freq)
	}

	sort.Ints(frequencies)

	result := ""
	prev := 0
	for i := len(frequencies) - 1; i >= 0; i-- {
		frequency := frequencies[i]

		if prev == frequency {
			continue
		}

		for _, c := range reversed[frequency] {
			for j := 0; j < frequency; j++ {
				result += string(c)
			}
		}

		prev = frequency
	}
	return result
}
