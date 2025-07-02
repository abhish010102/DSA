/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverse(head *ListNode) *ListNode {
    if head == nil{
        return head
    }
    
    l1 := head.Next
    head.Next = nil

    for l1 != nil {
        l2 := l1.Next
        l1.Next = head
        head=l1
        l1 = l2
    } 

    return head
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    l1=reverse(l1);
    l2=reverse(l2);
    return reverse(addTwoNumber(l1,l2));
}

func addTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {

	head, c := &ListNode{}, 0
	cur := head

	for l1 != nil || l2 != nil || c != 0 {
		val1, val2 := 0, 0

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		s := val1 + val2 + c
		c = s / 10
		cur.Next = &ListNode{Val: s % 10}
		cur = cur.Next
	}

	return head.Next
}