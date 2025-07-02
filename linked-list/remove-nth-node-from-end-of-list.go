/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h1, h2 := head, head

	for n >= 0 && h2 != nil {
		h2 = h2.Next
		n--
	}
	if n == 0 && h2 == nil {
		return head.Next
	}
	if n >= 0 {
		return nil
	}

	for h2 != nil {
		h1 = h1.Next
		h2 = h2.Next
	}

	if h2 == nil {
		h1.Next = h1.Next.Next
	}

	return head
}