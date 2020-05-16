package main

import "testing"

var tests = []struct {
	root *TreeNode
	want int
}{
	{
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		4,
	},
	{
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		3,
	},
	{
		&TreeNode{
			Val: 1,
		},
		1,
	},
	{
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
		},
		1,
	},
	{
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		4,
	},
}

func TestGoodNodes(t *testing.T) {
	for _, test := range tests {
		if got := goodNodes(test.root); got != test.want {
			t.Errorf("goodNodes(%v) = %v; want %v\n", test.root, got, test.want)
		}
	}
}
