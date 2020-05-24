package main

const (
	IntMin = 1
	IntMax = 1e8
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	preorderIndex := 0
	return bstFromPreorderUtil(preorder, &preorderIndex, preorder[0], IntMin, IntMax)
}

func bstFromPreorderUtil(preorder []int, preorderIndex *int, key, min, max int) *TreeNode {
	if *preorderIndex >= len(preorder) {
		return nil
	}

	var root *TreeNode
	if key >= min && key <= max {
		root = &TreeNode{Val: key}
		*preorderIndex += 1
		if *preorderIndex < len(preorder) {
			root.Left = bstFromPreorderUtil(preorder, preorderIndex, preorder[*preorderIndex], min, key)
		}
		if *preorderIndex < len(preorder) {
			root.Right = bstFromPreorderUtil(preorder, preorderIndex, preorder[*preorderIndex], key, max)
		}
	}

	return root
}
