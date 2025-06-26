/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	slow,fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = slow.Next
	slow.Next = nil
	for fast != nil {
        l1 := fast.Next
        fast.Next = slow
        slow = fast
		fast = l1
	}

	for head != nil && slow != nil {
		if slow.Val != head.Val {
			return false
		}
		slow = slow.Next
		head = head.Next
	}

	return true
}