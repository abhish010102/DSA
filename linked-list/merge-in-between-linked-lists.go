/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {

	l1, l2 := list1, list2
	for i := 1; l1 != nil; i++ {
		if i == a {
			l2 = l1
		} else if i-2 == b {
			break
		}
		l1 = l1.Next
	}

	l2.Next = list2

	for list2.Next != nil {
		list2 = list2.Next
	}

	list2.Next = l1
	return list1
}