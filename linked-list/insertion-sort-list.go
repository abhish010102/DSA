/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	h1, nxt := &ListNode{Next: head}, head.Next
	pre := head

	for nxt != nil {
		if pre.Val <= nxt.Val {
			pre = nxt
			nxt = nxt.Next
			continue
		}

		head = h1
		for head != nxt {
			if head.Next.Val > nxt.Val {
				l1 := nxt.Next
				pre.Next = l1
				nxt.Next = head.Next
				head.Next = nxt
				nxt = l1
				break
			}
			head = head.Next
		}
	}
	return h1.Next
}