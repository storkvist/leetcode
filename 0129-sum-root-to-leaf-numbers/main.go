package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return sumNumbersUtil(root, 0)
}

func sumNumbersUtil(root *TreeNode, value int) int {
	if root == nil {
		return 0
	}

	value = 10*value + root.Val

	if root.Left == nil && root.Right == nil {
		return value
	}

	return sumNumbersUtil(root.Left, value) + sumNumbersUtil(root.Right, value)
}
