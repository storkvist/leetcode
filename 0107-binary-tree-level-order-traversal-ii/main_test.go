package main

import "testing"

var tests = []struct {
	root *TreeNode
	want [][]int
}{
	{
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		[][]int{
			{15, 7},
			{9, 20},
			{3},
		},
	},
}

func TestLevelOrderBottom(t *testing.T) {
	for _, test := range tests {
		got := levelOrderBottom(test.root)
		if len(got) != len(test.want) {
			t.Errorf("levelOrderBottom(%+v) = %v; want %v\n", test.root, got, test.want)
			continue
		}
		for i := range test.want {
			if len(got[i]) != len(test.want[i]) {
				t.Errorf("levelOrderBottom(%+v) = %v; want %v\n", test.root, got, test.want)
				break
			}
			for j := range test.want[i] {
				if got[i][j] != test.want[i][j] {
					t.Errorf("levelOrderBottom(%+v) = %v; want %v\n", test.root, got, test.want)
					break
				}
			}
		}
	}
}
