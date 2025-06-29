/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head, cv, c := l1, l1, 0

	for l1 != nil && l2 != nil {
		l1.Val += l2.Val + c
		if l1.Val >= 10 {
			c = l1.Val / 10
			l1.Val %= 10
		} else {
			c = 0
		}
		l2 = l2.Next
		cv = l1
		l1 = l1.Next
	}

	if l2 != nil {
		cv.Next = l2
	}

	for cv.Next != nil {
		cv.Next.Val += c
		if cv.Next.Val >= 10 {
			c = cv.Next.Val / 10
			cv.Next.Val %= 10
		} else {
			c = 0
		}
		cv = cv.Next
	}

	if c != 0 {
		nnode := &ListNode{Val: c}
		cv.Next = nnode
	}

	return head
}



func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head, c := &ListNode{}, 0
	cur := head

	for l1 != nil || l2 != nil || c != 0 {
		val1, val2 := 0, 0

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		s := val1 + val2 + c
		c = s / 10
		cur.Next = &ListNode{Val: s % 10}
		cur = cur.Next
	}

	return head.Next
}