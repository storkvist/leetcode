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
//	return fmt.Sprintf("🌲%v", node.Val)
//}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	depths := make(map[*TreeNode]int)
	var flatList []*TreeNode

	list := list.New()
	list.PushBack(root)
	depths[root] = 0
	for list.Len() != 0 {
		element := list.Front()
		node := element.Value.(*TreeNode)

		flatList = append(flatList, node)

		if node.Left != nil {
			list.PushBack(node.Left)
			depths[node.Left] = depths[node] + 1
		}
		if node.Right != nil {
			list.PushBack(node.Right)
			depths[node.Right] = depths[node] + 1
		}

		list.Remove(element)
	}

	var result [][]int
	var currentSlice []int
	var currentDepth int
	for _, node := range flatList {
		if depths[node] != currentDepth {
			result = append(result, currentSlice)
			currentDepth = depths[node]
			currentSlice = []int{node.Val}
		} else {
			currentSlice = append(currentSlice, node.Val)
		}
	}
	result = append(result, currentSlice)

	return result
}
