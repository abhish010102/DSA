/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
    h1, h2, h3  := head, head, head
    
    for k != 1 && h2 != nil {
		h2 = h2.Next
		k--
	}

    h3=h2;

    for h2.Next !=nil {
        h1=h1.Next
        h2=h2.Next
    }

    h3.Val, h1.Val = h1.Val,h3.Val
    return head
}