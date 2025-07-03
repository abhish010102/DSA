/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	h1, t := head, 0

	for h1 != nil {
		t++
		h1 = h1.Next
	}

	h1 = &ListNode{Next: head}

	ret := make([]*ListNode, k)
	cn := t / k
	for i := 0; i < k; i++ {
		head = h1
		ret[i] = head.Next

		for j := 0; j < cn; j++ {
			head = head.Next
		}

		if t%k > 0 {
			head = head.Next
			t--
		}

		if head != nil {
			h1 = &ListNode{Next: head.Next}
			head.Next = nil
		}
	}

	return ret
}