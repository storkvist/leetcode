package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkNode(node *TreeNode, previousMax int) int {
	count, max := 0, previousMax
	if node.Val >= previousMax {
		count++
	}
	if node.Val > max {
		max = node.Val
	}
	if node.Left != nil {
		count += checkNode(node.Left, max)
	}
	if node.Right != nil {
		count += checkNode(node.Right, max)
	}

	return count
}

func goodNodes(root *TreeNode) int {
	return checkNode(root, root.Val)
}
