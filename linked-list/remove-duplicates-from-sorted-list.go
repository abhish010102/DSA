/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if (head==nil || head.Next==nil){
        return head
    }

    pre, nxt := head, head.Next

	for nxt != nil {
        // fmt.Println(pre.Val,nxt.Val)
		if pre.Val == nxt.Val {
			pre.Next = nxt.Next
		} else {
			pre = pre.Next
		}
		nxt = pre.Next
	}
	return head
}