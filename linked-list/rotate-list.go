/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil{
		return head
	}
	t, h1, h2 := 0, head, head

	for t != k && h1 != nil {
		t++
		h1 = h1.Next
	}

	if h1 == nil {
		return rotateRight(head, k%t)
	}

	for h1.Next != nil {
		h1 = h1.Next
		h2 = h2.Next
	}
	h1.Next = head
	head = h2.Next
	h2.Next = nil

	return head
}