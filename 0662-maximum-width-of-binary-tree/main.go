package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	max, leftmost, rightmost := 1, 1, 1
	root.Val = 1
	queue.PushBack(root)

	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			element := queue.Front()
			current := element.Value.(*TreeNode)
			if i == 0 {
				leftmost = current.Val
			}
			if i == size-1 {
				rightmost = current.Val
			}
			if current.Left != nil {
				current.Left.Val = current.Val * 2
				queue.PushBack(current.Left)
			}
			if current.Right != nil {
				current.Right.Val = current.Val*2 + 1
				queue.PushBack(current.Right)
			}
			queue.Remove(element)
		}
		if rightmost-leftmost+1 > max {
			max = rightmost - leftmost + 1
		}
	}
	return max
}
