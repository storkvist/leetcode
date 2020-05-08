package main

import (
	"testing"
)

var tests = []struct {
	root *TreeNode
	want [][]int
}{
	{},
	{
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		[][]int{{3}, {9, 20}, {15, 7}},
	},
	{
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 6,
				},
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		[][]int{{3}, {9, 20}, {6, 15, 7}},
	},
	{
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		[][]int{{3}, {9, 20}, {9, 3}},
	},
}

func TestLevelOrder(t *testing.T) {
	for _, test := range tests {
		got := levelOrder(test.root)

		ok := len(got) == len(test.want)
		if ok {
			for i, nodes := range got {
				ok = ok && (len(nodes) == len(test.want[i]))
				if !ok {
					break
				}
				for j, num := range nodes {
					ok = ok && (num == test.want[i][j])
				}
			}
		}

		if !ok {
			t.Errorf("levelOrder(%v) = %v; want %v", test.root, got, test.want)
		}
	}
}
