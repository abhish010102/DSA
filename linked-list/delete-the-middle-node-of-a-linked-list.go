/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head
	mslow := head

	for fast != nil && fast.Next != nil {
		mslow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if slow != mslow {
		mslow.Next = slow.Next
	} else {
		return nil
	}

	return head
}