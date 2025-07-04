/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return []int{-1, -1}
	}

	f, l, pre, p, min := -1, -1, -1, head.Val, 10000000

	for i := 2; head.Next != nil && head.Next.Next != nil; i++ {
		head = head.Next
		cur := head.Val
		if (p > cur && head.Next.Val > cur) || (p < cur && head.Next.Val < cur) {
			pre = l
			l = i

			if f == -1 {
				f = i
			} else if min > l-pre {
				min = l - pre
			}
		}
		p = cur
	}
    
	if f == -1 || min == 10000000{
		return []int{-1, -1}
	}

	return []int{min, l - f}
}