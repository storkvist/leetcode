package main

import "testing"

var tests = []struct {
	preorder []int
	wantNode *TreeNode
	want     []int
}{
	{
		[]int{8, 5, 1, 7, 10, 12},
		&TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 10,
				Right: &TreeNode{
					Val: 12,
				},
			},
		},
		[]int{8, 5, 10, 1, 7, 0, 12},
	},
	{
		[]int{4, 2},
		&TreeNode{},
		[]int{4, 2, 0},
	},
}

func BreadthFirstTraversal(root *TreeNode) []*TreeNode {
	queue := []*TreeNode{root}
	var nodes []*TreeNode
	return breadthFirstTraversalUtil(queue, nodes)
}

func breadthFirstTraversalUtil(queue []*TreeNode, nodes []*TreeNode) []*TreeNode {
	if len(queue) == 0 {
		return nodes
	}
	nodes = append(nodes, queue[0])

	something := queue[0].Left != nil || queue[0].Right != nil
	if queue[0].Left != nil {
		queue = append(queue, queue[0].Left)
	} else if something {
		queue = append(queue, &TreeNode{Val: 0})
	}
	if queue[0].Right != nil {
		queue = append(queue, queue[0].Right)
	} else if something {
		queue = append(queue, &TreeNode{Val: 0})
	}
	return breadthFirstTraversalUtil(queue[1:], nodes)
}

func (root *TreeNode) toArray() (result []int) {
	for _, node := range BreadthFirstTraversal(root) {
		result = append(result, node.Val)
	}
	return
}

func TestTreeNode_toArray(t *testing.T) {
	test := tests[0]
	got := test.wantNode.toArray()
	if len(got) != len(test.want) {
		t.Errorf("(%v).toArray() = %v; want %v\n", test.wantNode, got, test.want)
		return
	}
	for i := range test.want {
		if got[i] != test.want[i] {
			t.Errorf("(%v).toArray() = %v; want %v\n", test.wantNode, got, test.want)
			break
		}
	}
}

func TestNameBSTFromPreorder(t *testing.T) {
	for _, test := range tests {
		got := bstFromPreorder(test.preorder).toArray()
		if len(got) != len(test.want) {
			t.Errorf("(%v).toArray() = %v; want %v\n", test.wantNode, got, test.want)
			return
		}
		for i := range test.want {
			if got[i] != test.want[i] {
				t.Errorf("(%v).toArray() = %v; want %v\n", test.wantNode, got, test.want)
				break
			}
		}
	}
}
