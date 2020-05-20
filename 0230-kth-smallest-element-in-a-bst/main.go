package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode, k int, values []int) []int {
	if len(values) >= k {
		return values
	}
	if root.Left != nil {
		values = inorderTraversal(root.Left, k, values)
	}
	values = append(values, root.Val)
	if root.Right != nil {
		values = inorderTraversal(root.Right, k, values)
	}
	return values
}

func kthSmallest(root *TreeNode, k int) int {
	values := inorderTraversal(root, k, []int{})
	return values[k-1]
}

func (node *TreeNode) String() string {
	if node == nil {
		return "_"
	}

	return fmt.Sprintf("&{%d %v %v}", node.Val, node.Left, node.Right)
}
