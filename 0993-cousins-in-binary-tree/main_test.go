package main

import "testing"

var tests = []struct {
	root *TreeNode
	x    int
	y    int
	want bool
}{
	{
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		4,
		3,
		false,
	},
	{
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		5,
		4,
		true,
	},
	{
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		2,
		3,
		false,
	},
}

func TestIsCousins(t *testing.T) {
	for _, test := range tests {
		if got := isCousins(test.root, test.x, test.y); got != test.want {
			t.Errorf("func isCousins(%v, %v, %v) = %v; want %v\n", test.root, test.x, test.y, got, test.want)
		}
	}
}
