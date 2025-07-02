/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	nh := &ListNode{Next: head}
	f := nh

	for f != nil {
		sum := 0
		s := f.Next

		for s != nil {
			sum += s.Val
			if sum == 0 {
				f.Next = s.Next
				sum = 0
			}
			s = s.Next
		}

		f = f.Next
	}

	return nh.Next
}