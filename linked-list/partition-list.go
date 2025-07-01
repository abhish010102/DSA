/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	h1, h2 := &ListNode{}, &ListNode{}
	c1, c2 := h1, h2
	
	for head != nil {
		if head.Val < x {
			c1.Next = head
			c1 = c1.Next
		} else {
			c2.Next = head
			c2 = c2.Next
		}
		head = head.Next
	}
	
	c1.Next = h2.Next
	c2.Next = nil

	return h1.Next

}