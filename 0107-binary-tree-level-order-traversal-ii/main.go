package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) (result [][]int) {
	if root == nil {
		return
	}

	hash := make(map[int][]int)
	dfs(&hash, 0, root)

	length := len(hash)
	for i := length - 1; i >= 0; i-- {
		r := make([]int, len(hash[i]))
		copy(r, hash[i])
		result = append(result, r)
	}

	return
}

func dfs(hash *map[int][]int, level int, node *TreeNode) {
	if node == nil {
		return
	}

	(*hash)[level] = append((*hash)[level], node.Val)
	dfs(hash, level+1, node.Left)
	dfs(hash, level+1, node.Right)
}
