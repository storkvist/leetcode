package main

import (
	"testing"
)

var tests = []struct {
	root *TreeNode
	want []int
}{
	{
		&TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
		},
		[]int{4, 7, 2, 9, 6, 3, 1},
	},
}

func TestInvertTree(t *testing.T) {
	for _, test := range tests {
		original := test.root.toSlice()
		got := invertTree(test.root).toSlice()
		if len(got) != len(test.want) {
			t.Errorf("invertTree(%v) -> %v; want %v\n", original, got, test.want)
			continue
		}
		for i := range test.want {
			if got[i] != test.want[i] {
				t.Errorf("invertTree(%v) -> %v; want %v\n", original, got, test.want)
				break
			}
		}
	}
}

func breadthFirstTraversal(root *TreeNode) []*TreeNode {
	queue := []*TreeNode{root}
	var nodes []*TreeNode
	return breadthFirstTraversalUtil(queue, nodes)
}

func breadthFirstTraversalUtil(queue []*TreeNode, nodes []*TreeNode) []*TreeNode {
	if len(queue) == 0 {
		return nodes
	}
	nodes = append(nodes, queue[0])
	if queue[0].Left != nil {
		queue = append(queue, queue[0].Left)
	}
	if queue[0].Right != nil {
		queue = append(queue, queue[0].Right)
	}
	return breadthFirstTraversalUtil(queue[1:], nodes)
}

func (t *TreeNode) toSlice() (result []int) {
	for _, node := range breadthFirstTraversal(t) {
		result = append(result, node.Val)
	}
	return
}
