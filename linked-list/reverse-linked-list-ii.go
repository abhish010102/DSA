/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
    
    c := 1
	h1 := head
	pre := head

	for c != left {
		c++;
        pre = h1
		h1 = h1.Next
	}

	n1 := h1.Next

	for c != right {
		c++
        l1 := n1.Next
		n1.Next = h1
		h1 = n1
		n1 = l1
	}

    if(left!=1){
        pre.Next.Next = n1
        pre.Next = h1
    }else{
        pre.Next=n1
        head=h1
    }

	return head
}