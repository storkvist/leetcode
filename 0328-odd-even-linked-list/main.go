package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	oddsTail, evensHead, evensTail := head, head.Next, head.Next
	for evensTail != nil && evensTail.Next != nil {
		oddsTail.Next = evensTail.Next
		oddsTail = oddsTail.Next
		evensTail.Next = oddsTail.Next
		evensTail = evensTail.Next
	}
	oddsTail.Next = evensHead
	return head
}

func oddEvenListInitial(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var oddsHead, oddsTail, evensHead, evensTail *ListNode
	i, current := 1, head
	for current != nil {
		if i%2 == 0 {
			if evensHead == nil {
				evensHead = current
			}
			if evensTail == nil {
				evensTail = current
			} else {
				evensTail.Next = current
				evensTail = current
			}
		} else {
			if oddsHead == nil {
				oddsHead = current
			}
			if oddsTail == nil {
				oddsTail = current
			} else {
				oddsTail.Next = current
				oddsTail = current
			}
		}

		current = current.Next
		i++
	}

	if evensTail != nil {
		evensTail.Next = nil
	}

	oddsTail.Next = evensHead
	return oddsHead
}
