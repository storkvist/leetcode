package main

import "testing"

var tests = []struct {
	root *TreeNode
	k    int
	want int
}{
	{
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		1,
		1,
	},
	{
		&TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		3,
		3,
	},
}

func TestKthSmallest(t *testing.T) {
	for _, test := range tests {
		if got := kthSmallest(test.root, test.k); got != test.want {
			t.Errorf("kthSmallest(%v, %v) = %v; want %v\n", test.root, test.k, got, test.want)
		}
	}
}
