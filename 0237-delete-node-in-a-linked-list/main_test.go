package main

import (
	"strconv"
	"strings"
	"testing"
)

var tests = []struct {
	head *ListNode
	node int
	want string
}{
	{
		&ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		},
		5,
		"4, 1, 9",
	},
	{
		&ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		},
		1,
		"4, 5, 9",
	},
}

func (n *ListNode) values() string {
	vals := make([]string, 0)
	for n != nil {
		vals = append(vals, strconv.Itoa(n.Val))
		n = n.Next
	}
	return strings.Join(vals, ", ")
}

func TestDeleteNode(t *testing.T) {
	for _, test := range tests {
		node := test.head
		for node.Val != test.node {
			node = node.Next
		}
		deleteNode(node)
		if got := test.head.values(); got != test.want {
			t.Errorf("deleteNode(%v) -> %q; want %q\n", node, got, test.want)
		}
	}
}
