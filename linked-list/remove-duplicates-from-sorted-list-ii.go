/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nh := &ListNode{Val: 0, Next: head}
	pre, dv := nh, -101

	for head != nil && head.Next != nil {

		if head.Val == head.Next.Val {
			dv = head.Val
		}else{
            pre=head
            head=head.Next
            continue
        }

		for head != nil && head.Val == dv {
			head = head.Next
        }
        pre.Next=head
	}

	return nh.Next
}