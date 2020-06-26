package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	root *TreeNode
	want int
}{
	{
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		25,
	},
	{
		&TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 0,
			},
		},
		1026,
	},
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("{Val: %v, Left: %v, Right: %v}", t.Val, t.Left, t.Right)
}

func TestSumNumbers(t *testing.T) {
	for _, test := range tests {
		if got := sumNumbers(test.root); got != test.want {
			t.Errorf("sumNumbers(%+v) = %v; want %v\n", test.root, got, test.want)
		}
	}
}
