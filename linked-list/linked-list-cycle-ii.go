/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head

	for fast != nil && fast.Next != nil {
        slow = slow.Next
		fast = fast.Next.Next
        if fast==slow {
            for head!=slow {
                slow=slow.Next;
                head=head.Next;
            }
            return head;
        }
	}
	return nil
}