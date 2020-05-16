package main

import (
	"testing"
)

var tests = []struct {
	values []int
	want   []int
}{
	{[]int{1, 2, 3, 4, 5}, []int{1, 3, 5, 2, 4}},
	{[]int{2, 1, 3, 5, 6, 4, 7}, []int{2, 3, 6, 7, 1, 5, 4}},
	{[]int{}, []int{}},
	{[]int{3}, []int{3}},
	{[]int{2, 1}, []int{2, 1}},
}

func valuesToList(values []int) *ListNode {
	length := len(values)
	if length == 0 {
		return nil
	}

	head := &ListNode{}
	current := head
	for i := 0; i < length; i++ {
		current.Val = values[i]

		if i < length-1 {
			current.Next = &ListNode{}
			current = current.Next
		}
	}

	return head
}

func listToValues(head *ListNode) (result []int) {
	if head == nil {
		return []int{}
	}

	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return
}

func TestOddEvenList(t *testing.T) {
	for _, test := range tests {
		list := valuesToList(test.values)
		modified := oddEvenList(list)
		got := listToValues(modified)
		for i, x := range test.want {
			if x != got[i] {
				t.Errorf("oddEvenList(%v) = %v; want %v", test.values, got, test.want)
				break
			}
		}
	}
}

func TestOddEvenListInitial(t *testing.T) {
	for _, test := range tests {
		list := valuesToList(test.values)
		modified := oddEvenListInitial(list)
		got := listToValues(modified)
		for i, x := range test.want {
			if x != got[i] {
				t.Errorf("oddEvenListInitial(%v) = %v; want %v", test.values, got, test.want)
				break
			}
		}
	}
}

func BenchmarkOddEvenList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := valuesToList(tests[0].values)
		oddEvenList(list)
	}
}

func BenchmarkOddEvenListInitial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := valuesToList(tests[0].values)
		oddEvenListInitial(list)
	}
}
