package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func (node TreeNode) String() string {
//	return fmt.Sprintf("{🌲 %v  L%v  R%v}", node.Val, node.Left, node.Right)
//}

func isCousins(root *TreeNode, x int, y int) bool {
	depths := make(map[int][2]int)
	foundNodes := 0

	queue := list.New()
	queue.PushBack(root)
	depths[root.Val] = [2]int{1: 0}
	for queue.Len() != 0 {
		element := queue.Front()
		node := element.Value.(*TreeNode)
		nodeDepth := depths[node.Val][1]

		if node.Val == x || node.Val == y {
			foundNodes++
		}
		if foundNodes == 2 {
			break
		}

		if node.Left != nil {
			queue.PushBack(node.Left)
			depths[node.Left.Val] = [2]int{0: node.Val, nodeDepth + 1}
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
			depths[node.Right.Val] = [2]int{0: node.Val, nodeDepth + 1}
		}

		queue.Remove(element)
	}

	return (depths[x][0] != depths[y][0]) && (depths[x][1] == depths[y][1])
}
